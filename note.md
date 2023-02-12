
//tmp
// curl https://bsc-mainnet.nodereal.io/v1/your-api-key \
// -X POST \
// -H "Content-Type: application/json" \
// -d '{"jsonrpc":"2.0","method":"txpool_content","params":[],"id":1}'

//https://docs.nodereal.io/reference/txpool_content
// {
// 	"jsonrpc":"2.0",
// 	"id":1,
// 	"result":{
// 	  "pending":{
// 		"0x0000412b8ccb938cd2a3f473c3ce466cdedee049":{
// 		  "36096":{
// 			"blockHash":null,
// 			"blockNumber":null,
// 			"from":"0x0000412b8ccb938cd2a3f473c3ce466cdedee049",
// 			"gas":"0x2dc6c0",
// 			"gasPrice":"0x138eca480",
// 			"hash":"0xb6721e4888adc621158aed7586ea80db4f96529cf5bad16ae1795bbae5eba9a3",
// 			"input":"0x000000000300010026f2271084ee532a0d4238f5fc4a1e8c043f8749ed4f274dbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c05000100022710f2b9155e3c9756a18ef6572",
// 			"nonce":"0x8d00",
// 			"to":"0x05ccc798a921800441291636640880045cebe7c6",
// 			"transactionIndex":null,
// 			"value":"0x0",
// 			"type":"0x0",
// 			"v":"0x94",
// 			"r":"0x19eba5a20f3e4d0cdd5f8b5d93b1127e7bbc7426267398e4cbef5c8928648b34",
// 			"s":"0x44d235ded909b9ac279a1a350507f554d0eb1846c3cf1b36e9d1b535e2eee524"
// 		  }
// 		}
// 	  },
// 	  "queued":{
// 		"0xa267828beb9ed84b0e91d3341fbb973153b9b925":{
// 		  "315743":{
// 			"blockHash":null,
// 			"blockNumber":null,
// 			"from":"0xa267828beb9ed84b0e91d3341fbb973153b9b925",
// 			"gas":"0x5cc60",
// 			"gasPrice":"0x147d35700",
// 			"hash":"0xe13dbbce3cfc30f18027c3cc0b8857ee5d412e6d20bded9a4e8ed1f52cde747d",
// 			"input":"0x8ca0d39c00000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000",
// 			"nonce":"0x4d15f",
// 			"to":"0xb400af9ff00e7ba1f6aac4eeff475e965f9cba1f",
// 			"transactionIndex":null,
// 			"value":"0x0",
// 			"type":"0x0",
// 			"v":"0x94",
// 			"r":"0xdb838ea3f72e1bf066cab996e448698399130a4c159d4fe1507f342a107c47f8",
// 			"s":"0x60f7ba3eee2a382852d7b1971838c3ec9702dca705811170b6c4696d7b1fa14c"
// 		  }
// 		}
// 	  }
// 	}
//   }

// which componenet you want to learn with high performance update and insert
// sqlite3  light convinient
// mysql
// postgress
//
// address,nonce,gas,gasPrice,hash,to,method,input,value,sampleTime,blocktime,transactionIndex

//### 40ms per 100

// msgPerS = 100 * 25

// bytesPerS = msgPerS * 100

// bytesPerS


### 遇到的问题

1. 多个ws 链接，同一个endpoints的时候，需要增加链接方式
2. 多个ws 不同的链接的时候，需要外部统一调度管理

解决办法: 1. 提供通用的调用入口，完成数据请求发送。
```go
type PoolInterface interface {
    
	PendingPool() []string
	QueuedPool() []string
}

```


### 思考
```shell
## 代码模块的可重用性
将代码可重用性用到极致 
网络尤其是链的监听，查询，功能相似性很大。这里做成通用功能，可以极大的提高效率，减少时间上的浪费

## 常用的接口
1. 监听块的情况
2. 通过块信息查询tx 
3. 通过tx 查from to value 
```