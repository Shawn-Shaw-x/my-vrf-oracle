package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"myvrf/bindings" // <-- 你实际生成路径
)

func main() {
	// === 配置 ===
	rpcURL := "http://127.0.0.1:8545"
	privateKeyHex := "YOUR_PRIVATE_KEY_HEX"
	coordinatorAddr := common.HexToAddress("0xYourDeployedCoordinatorAddress")

	// === 连接以太坊本地节点 ===
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// === 导入私钥 ===
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}
	publicAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("Wallet:", publicAddr.Hex())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337)) // hardhat/anvil 默认链 ID
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.Context = context.Background()
	auth.GasLimit = 5_000_000

	// === 绑定合约 ===
	coordinator, err := bindings.NewVRFCoordinatorV2(coordinatorAddr, client)
	if err != nil {
		log.Fatalf("Failed to bind coordinator: %v", err)
	}

	// === 注册 keyHash 到合约（只需注册一次） ===
	keyHash := crypto.Keccak256Hash([]byte("test-key"))
	tx, err := coordinator.RegisterKeyHash(auth, keyHash, publicAddr)
	if err != nil {
		log.Fatalf("Failed to register keyHash: %v", err)
	}
	fmt.Println("✅ registerKeyHash tx:", tx.Hash().Hex())

	// === 请求随机数 ===
	tx2, err := coordinator.RequestRandomWords(
		auth,
		keyHash,
		big.NewInt(0),  // subId ignored
		uint16(3),      // confirmations ignored
		uint32(100000), // callbackGasLimit ignored
		uint32(2),      // numWords
	)
	if err != nil {
		log.Fatalf("Failed to request random words: %v", err)
	}
	fmt.Println("📩 requestRandomWords tx:", tx2.Hash().Hex())

	// === 获取最新 requestId ===
	requestId, err := coordinator.RequestCounter(nil)
	if err != nil {
		log.Fatalf("Failed to get request counter: %v", err)
	}
	fmt.Println("🔍 requestId:", requestId)

	// === 模拟随机数生成 ===
	numWords := 2
	randomWords := make([]*big.Int, numWords)
	for i := 0; i < numWords; i++ {
		b := make([]byte, 32)
		rand.Read(b)
		randomWords[i] = new(big.Int).SetBytes(b)
	}

	// === fulfill 回调（模拟 Oracle） ===
	tx3, err := coordinator.FulfillRandomWords(
		auth,
		requestId,
		randomWords,
		keyHash,
	)
	if err != nil {
		log.Fatalf("Failed to fulfill random words: %v", err)
	}
	fmt.Println("📤 fulfillRandomWords tx:", tx3.Hash().Hex())
}
