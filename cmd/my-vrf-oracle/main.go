package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"my-vrf-oracle/bindings"
)

// EthereumClient 封装 go-ethereum 的 ethclient.Client
// 提供链上连接功能
type EthereumClient struct {
	*ethclient.Client
}

// NewEthereumClient 创建以太坊 WebSocket 客户端连接
func NewEthereumClient() (*EthereumClient, error) {
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return &EthereumClient{client}, nil
}

// NewContractAuth 创建合约调用账户（Transactor）
func NewContractAuth() *bind.TransactOpts {
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // anvil 默认
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.Context = context.Background()
	auth.GasLimit = 5_000_000

	fmt.Println("🔐Wallet:", crypto.PubkeyToAddress(privateKey.PublicKey).Hex())
	return auth
}

// InitContractBinding 绑定合约地址并构建事件订阅过滤器
func InitContractBinding(client *EthereumClient) (*bindings.VRFCoordinatorV2, ethereum.FilterQuery) {
	coordinatorAddr := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	coordinator, err := bindings.NewVRFCoordinatorV2(coordinatorAddr, client)
	if err != nil {
		log.Fatalf("Failed to bind coordinator: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{coordinatorAddr},
	}
	return coordinator, query
}

func main() {
	client, err := NewEthereumClient()
	if err != nil {
		log.Fatal(err)
	}
	auth := NewContractAuth()
	coordinator, query := InitContractBinding(client)

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}
	fmt.Println("📡 Listening for RandomWordsRequested events...")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		parsedABI := readVRFCoordinatorABI()
		eventSig := parsedABI.Events["RandomWordsRequested"].ID

		for {
			select {
			case vLog := <-logs:
				if vLog.Topics[0] != eventSig {
					continue
				}

				event, err := coordinator.ParseRandomWordsRequested(vLog)
				if err != nil {
					log.Printf("⚠️ Failed to parse event: %v", err)
					continue
				}

				fmt.Printf("📥 New request received:\n")
				fmt.Printf("  - RequestId: %d\n", event.RequestId)
				fmt.Printf("  - Requester: %s\n", event.Requester.Hex())
				fmt.Printf("  - NumWords: %d\n", event.NumWords)

				randomWords := generateRandomNums(event)
				tx, err := coordinator.FulfillRandomWords(auth, event.RequestId, randomWords)
				if err != nil {
					log.Printf("❌ Failed to fulfill random words: %v", err)
					continue
				}
				fmt.Println("✅ fulfillRandomWords tx:", tx.Hash().Hex())

			case <-interrupt:
				fmt.Println("🛑 Gracefully shutting down listener.")
				sub.Unsubscribe()
				client.Close()
				return
			}
		}
	}()

	// 阻塞主线程
	select {}
}

// generateRandomNums 生成指定数量的 256-bit 随机数数组
func generateRandomNums(event *bindings.VRFCoordinatorV2RandomWordsRequested) []*big.Int {
	randomWords := make([]*big.Int, event.NumWords)
	for i := uint32(0); i < event.NumWords; i++ {
		b := make([]byte, 32)
		_, err := rand.Read(b)
		if err != nil {
			log.Println("❌ Failed to generate randomness:", err)
			continue
		}
		randomWords[i] = new(big.Int).SetBytes(b)
		fmt.Printf("    randomWords[%d] = 0x%s\n", i, randomWords[i].Text(16))
	}
	return randomWords
}

// readVRFCoordinatorABI 从文件中读取并解析 ABI
func readVRFCoordinatorABI() abi.ABI {
	abiBytes, err := os.ReadFile("./abis/VRFCoodinatorV2.sol/VRFCoordinatorV2.json")
	if err != nil {
		log.Fatalf("❌ Failed to read ABI file: %v", err)
	}

	var raw interface{}
	if err := json.Unmarshal(abiBytes, &raw); err != nil {
		log.Fatalf("❌ Failed to unmarshal JSON: %v", err)
	}

	var abiString string
	switch v := raw.(type) {
	case map[string]interface{}:
		if abiField, ok := v["abi"]; ok {
			abiFieldBytes, _ := json.Marshal(abiField)
			abiString = string(abiFieldBytes)
		} else {
			abiString = string(abiBytes)
		}
	default:
		abiString = string(abiBytes)
	}

	parsedABI, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatalf("❌ Failed to parse ABI: %v", err)
	}
	return parsedABI
}
