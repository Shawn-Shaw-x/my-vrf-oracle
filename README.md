## 链下随机数

以太坊「链下随机数」（off-chain randomness），指的是不在链上生成，而是由链外系统（如预言机节点）生成，并最终提交到链上使用的一种随机数生成方式。它是为了解决链上无法安全生成真正随机数的问题而提出的。

### github链接
[链下 Oracle（go）](https://github.com/Shawn-Shaw-x/my-vrf-oracle.git)
[链上合约（solidity）](https://github.com/Shawn-Shaw-x/my-vrf.git)

## 链上随机数
在 `solidity` 语言中，我们可以采用 `hash` 函数来进行生成随机数。如：

```js
randomBytes = keccak256(abi.encodePacked(block.timestamp, msg.sender, blockhash(block.number-1)));
```
但是，在以太坊中，这样的随机数生成并不是安全的。我们知道，以太坊中中交易是会被验证者（矿工）打包的，这也就意味着，验证交易的物理机器在他们手上。矿工可以控制`block.timestamp` 等参数进行预测随机数。再不济他也可以离线生成随机数，等生成到所需的随机数后方才使用。这样在诸如抽奖、铸造稀有 `NFT` 等场景中，是有很大安全隐患的。
## 链下随机数
链下随机数是一种避免矿工节点作弊的方案。他的核心思想是：将随机数放到链下进行生成，通过生成证明后上传到链上，确保随机数是不可预测、可验证、安全的。
### 流程图分析
![image.png](https://img.learnblockchain.cn/attachments/2025/04/VXM0danK680885b46a700.png)
### 有几方进行交互、如何交互
在简化版的 `ChainLink` 随机数预言机实现中。我们一般认为有三方进行交互。分别是：1. 链上的消费者（用户合约）。2. 链上的 `ChainLink-VRF`协调者合约。3. 链下的随机数预言机系统。

在这三方进行交互的过程中，我们可以对其流程进行简化：
1. **消费者合约调用协调者合约**：

    消费者合约首先需要发送一个请求给 `chainlink` 的协调者合约进行请求随机数服务，请求包含随机数的需求个数。
    
2. **协调者合约释放请求事件日志**：

    协调者收到消费者的请求后，生成唯一请求 `id`，记录消费饿者的合约地址，生成 `preSeed` 种子，向区块链上释放一个请求的事件日志。
   
3. **链下预言机监听到协调者的随机数请求时间**：

    链下预言机一直在监听协调者的请求事件事件，当发现链上出现了这个事件的时候，预言机会去获取事件相应的 `preSeed`。
  
4. **预言机生成随机数，并生成证明，上传到协调者合约**：

    当获取到事件中携带的 `preSeed` 后，预言机通过本机 `hash` 函数生成一个（或多个）随机数，并且携带一份证明（`proof`），确保这份随机数是由协调者传过来的 `preSeed` 生成的，是不可预测、可验证的。然后上传给到协调者合约。
    
5. **协调者合约将随机数转发到消费者合约**：

    最后，随机数和随机数相应的证明（`proof`）流转到了协调者合约中。在这里，协调者可以不对随机数进行存储，但是必须要对随机数的证明（`proof`）进行验证，以确保链下预言机不出现作恶的情况。然后，调用消费者的相应回调函数，直接将随机数转发给消费者进行处理即可。

### 怎么保证验证者节点不作恶
在 `chainlink` 的架构中，通过将随机数生成服务转移到链下进行，矿工没办法进行操纵随机数生成。即保证了验证者节点不作恶。
### 怎么保证链下预言机不作恶
在 `chainlink` 的预言机架构中，保证链下预言机不作恶的最核心的要点就是 `VRF（variable random function）`，节点返回随机数的同时，必须要携带一个可以证明的凭证，在链上协调者合约中，会对这个凭证来进行验证，确保这个随机数是由协调者出示的种子经过特定算法生成的且不可伪造。

其次，`chainlink` 在设计中，也将引入了 `staking、slashing` 质押罚没机制。通过奖惩来确保不会出现预言机节点作恶的情况。
## 链下随机数极简样例

### 合约项目创建
#### 消费者 MyVRFConsumer.sol：
    
```js
/// @title MyVRFConsumer - 一个使用 VRF 随机数的消费者示例合约
/// @notice 通过协调器请求链上随机数，接收回调并处理结果
contract MyVRFConsumer is VRFConsumerBaseV2 {
    // --- 状态变量 ---

    /// @dev VRF协调器合约实例
    VRFCoordinatorV2 public COORDINATOR;

    // --- 事件定义 ---

    /// @notice 当 fulfillRandomWords 被调用时触发，表示已获得随机数
    event GainRandomnessEvent(
        uint256 indexed requestId,
        uint256[] indexed randomWords
    );

    // --- 构造函数 ---

    /// @param coordinator 协调器合约地址
    constructor(address coordinator)
    VRFConsumerBaseV2(coordinator)
    {
        COORDINATOR = VRFCoordinatorV2(coordinator);
    }

    // --- 外部函数 ---

    /// @notice 用户调用该函数请求链上 VRF 随机数
    /// @param numWords 请求的随机数个数
    function requestRandomWords(
        uint32 numWords
    ) external {
        COORDINATOR.requestRandomWords(numWords);
    }

    // --- 实现回调 ---

    /// @notice 协调器调用该函数传回随机数
    /// @dev VRFCoordinator 合约通过 rawFulfillRandomWords 调用该函数
    /// @param requestId 请求编号
    /// @param randomWords 返回的随机数数组
    function fulfillRandomWords(
        uint256 requestId,
        uint256[] memory randomWords
    ) internal override {
        // 此处可添加游戏逻辑、状态保存、奖励逻辑等
        emit GainRandomnessEvent(requestId, randomWords);
    }
}
```
#### 链上协同合约 VRFCoodinatorV2.sol：

```js
/// @title VRFCoordinatorV2 - 一个最简版链上 VRF 协调器合约
/// @notice 接收消费者随机数请求，抛出事件供链下预言机监听，并接收链下 fulfill 回调将随机数发给消费者
/// @dev 不包含订阅系统、权限控制、链下证明、验证等，仅为链下回调流程的最小实现示例
contract VRFCoordinatorV2 {
    // --- 数据结构 ---

    /// @dev 请求结构体，记录每一次 requestRandomWords 调用信息
    struct Request {
        address requester;    // 请求方地址（消费者）
        uint32 numWords;      // 请求的随机数个数
        bool fulfilled;       // 是否已被 fulfill 处理
    }

    /// @dev 请求记录表，按 requestId 索引
    mapping(uint256 => Request) public requests;

    /// @dev 请求 ID 计数器，自增生成唯一 ID
    uint256 public requestCounter;

    // --- 事件定义 ---

    /// @notice 当有请求发起时抛出，链下预言机监听后生成随机数
    event RandomWordsRequested(
        uint256 indexed requestId,
        address indexed requester,
        uint32 numWords
    );

    /// @notice 当 fulfill 被链下预言机调用时触发，表示随机数已交付
    event RandomWordsFulfilled(
        uint256 indexed requestId,
        uint256[] randomWords
    );

    // --- 外部函数 ---

    /// @notice 消费者合约调用以请求链上 VRF 随机数
    /// @param numWords 请求的随机数个数
    /// @return requestId 本次请求的唯一 ID
    function requestRandomWords(
        uint32 numWords
    ) external returns (uint256 requestId) {
        requestId = ++requestCounter;

        // 保存请求信息
        requests[requestId] = Request({
            requester: msg.sender,
            numWords: numWords,
            fulfilled: false
        });

        // 抛出事件供链下监听
        emit RandomWordsRequested(requestId, msg.sender, numWords);
    }

    /// @notice 被链下预言机节点调用，传入随机数并回调给消费者
    /// @dev 假设链下已完成证明校验，此处不再验证
    /// @param requestId 请求编号
    /// @param randomWords 生成的随机数数组
    function fulfillRandomWords(
        uint256 requestId,
        uint256[] calldata randomWords
    ) external {
        Request storage req = requests[requestId];
        require(!req.fulfilled, "Already fulfilled");

        req.fulfilled = true;

        // 回调请求方合约，将随机数传递回去
        VRFConsumerBaseV2(req.requester).rawFulfillRandomWords(
            requestId,
            randomWords
        );

        // 抛出事件供监听与确认
        emit RandomWordsFulfilled(requestId, randomWords);
    }
}
```
#### 部署脚本VRFCoordinatorScript.s.sol:

```js
contract DeployVRFCoordinatorScript is Script {
    function run() external {
        vm.startBroadcast();

        // 部署 VRFCoordinatorV2
        VRFCoordinatorV2 coordinator = new VRFCoordinatorV2();
        console2.log("VRFCoordinatorV2 deployed at:", address(coordinator));

        // consumer 部署
        MyVRFConsumer consumer = new MyVRFConsumer(address(coordinator));
        consumer.requestRandomWords(2);

        vm.stopBroadcast();
    }
}
```
#### 测试脚本 VRFTest.t.sol:

```js
// 用于测试链下 fulfillRandomWords 调用
contract OracleSimulator {
    VRFCoordinatorV2 coordinator;

    constructor(address _coordinator) {
        coordinator = VRFCoordinatorV2(_coordinator);
    }

    function fulfill(uint256 requestId, uint32 numWords ) external {
        uint256[] memory words = new uint256[](numWords);
        for (uint i = 0; i < numWords; i++) {
            words[i] = uint256(keccak256(abi.encodePacked(block.timestamp, i)));
        }
        console.logString("---------- off-chain oracle random nums-------");
        for (uint i = 0; i < numWords; i++) {
            console.logUint(words[i]);
        }
        console.logString("---------- off-chain oracle random nums-------");


    coordinator.fulfillRandomWords(requestId, words);
    }
}

contract VRFTest is Test {
    VRFCoordinatorV2 coordinator;
    MyVRFConsumer consumer;
    OracleSimulator oracle;

    address oracleAddr;

    function setUp() public {
        coordinator = new VRFCoordinatorV2();
        consumer = new MyVRFConsumer(address(coordinator));
        oracle = new OracleSimulator(address(coordinator));
        oracleAddr = address(oracle);

    }

    function testVRFIntegration() public {
        // 1.consumer 请求随机数
        consumer.requestRandomWords(
            2 // numWords
        );

        // 2.模拟链上 coodinator 获取到最新 requestId
        uint256 requestId = coordinator.requestCounter();
        (
            address requester,
            uint32 numWords,
            bool fulfilled
        ) = coordinator.requests(requestId);

        // 3.模拟链下捕捉完链上事件，生成随机数后发起 fulfill（OracleSimulator 调用）
        vm.prank(oracleAddr);
        oracle.fulfill(requestId, numWords);

        // 4.consumer 验证 fulfilled 标志
        (address requester1,
        uint32 numWords1,
        bool fulfilled1
        ) = coordinator.requests(requestId);
        console.logAddress(requester1);
        console.logUint(numWords1);
        console.logBool(fulfilled1);

    assertTrue(fulfilled1);


    }
}
```
### 预言机项目创建
#### main.go：

```js
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
```
#### 脚本命令 Makefile：
```js
GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

VRF_ABI_ARTIFACT := ./abis/VRFCoodinatorV2.sol/VRFCoordinatorV2.json


my-vrf-oracle:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/my-vrf-oracle

clean:
	rm cmd/my-vrf-oracle

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings: binding-vrf


binding-vrf:
	$(eval temp := $(shell mktemp))

	cat $(VRF_ABI_ARTIFACT) \
    	| jq -r .bytecode.object > $(temp)

	cat $(VRF_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg bindings \
		--abi - \
		--out bindings/VRFCoordinatorV2.go \
		--type VRFCoordinatorV2 \
		--bin $(temp)

		rm $(temp)


.PHONY: \
	my-vrf-oracle \
	bindings \
	binding-vrf \
	clean \
	test \
	lint
```
### 测试
#### 构建合约：
```js
forge build
```

![image.png](https://img.learnblockchain.cn/attachments/2025/04/5vtszbmj68087963f1004.png)
#### 转移 ABI 文件：
将`out/VRFCoodinatorV2.sol/VRFCoordinatorV2.json`中的 `JSON` 文件转移到 `Go` 项目中的 `abis` 文件夹下。

![image.png](https://img.learnblockchain.cn/attachments/2025/04/RNWfK7BA68087a0391fc8.png)

![image.png](https://img.learnblockchain.cn/attachments/2025/04/Y0R5P4ez68087a237c7c0.png)
#### 生成 go 语言 bindings
控制台执行命令
```js
make bindings
```

![image.png](https://img.learnblockchain.cn/attachments/2025/04/OavdoA7z68087a8177a52.png)
未安装 abigen 的需要先安装：
```js
 go install abigen
```
`binding` 完成后会自动生成文件

![image.png](https://img.learnblockchain.cn/attachments/2025/04/FpY23Pxz68087ae2eefee.png)

#### 启动 anvil
在合约项目目录下，控制台执行

```js
anvil -vvvv
```

![image.png](https://img.learnblockchain.cn/attachments/2025/04/9eAmtEBY68087b8026284.png)
#### 部署合约：
在合约项目目录下，控制台执行
```js
forge script script/VRFCoordinatorScript.s.sol \
  --rpc-url 127.0.0.1:8545 \
  --private-key $PRIVATE_KEY \
  --broadcast \
  -vvvv
```
![image.png](https://img.learnblockchain.cn/attachments/2025/04/oEhdDnFY68087be67f274.png)


#### 启动 go 程序：
main.go中配置好相关参数：

![image.png](https://img.learnblockchain.cn/attachments/2025/04/WRBaO3cm68087cd05b5bf.png)

配置好参数后可以直接在 `IDE` 中启动 `main` 函数

![image.png](https://img.learnblockchain.cn/attachments/2025/04/GPQJllrE68087d1b82943.png)
#### 消费者请求随机数：
回到控制台，合约项目的目录下，执行
```js
cast send 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512 \
    "requestRandomWords(uint32)" 2 \
    --rpc-url http://localhost:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```   

![image.png](https://img.learnblockchain.cn/attachments/2025/04/vJYN1bSE68087dc091699.png)
此时，回到 `go` 项目的控制台中，可以看到

![image.png](https://img.learnblockchain.cn/attachments/2025/04/4cP2fCwQ68087e61db40c.png)
我们可以使用 `cast` 命令，查看合约的事件日志

```js
cast logs \
  --from-block 0 \
  --address 0x5FbDB2315678afecb367f032d93F642f64180aa3 \
  --rpc-url http://localhost:8545
```

![image.png](https://img.learnblockchain.cn/attachments/2025/04/RAKtBBFd68087f0016f8e.png)
可以观察到，我们 `go` 项目生成的和 `go` 项目上传到合约中的随机数是一致的。
