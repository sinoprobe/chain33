package testnode

var cfgstring = `Title="chain33"

[log]
# 日志级别，支持debug(dbug)/info/warn/error(eror)/crit
loglevel = "debug"
logConsoleLevel = "info"
# 日志文件名，可带目录，所有生成的日志文件都放到此目录下
logFile = "logs/chain33.log"
# 单个日志文件的最大值（单位：兆）
maxFileSize = 20
# 最多保存的历史日志文件个数
maxBackups = 20
# 最多保存的历史日志消息（单位：天）
maxAge = 28
# 日志文件名是否使用本地事件（否则使用UTC时间）
localTime = true
# 历史日志文件是否压缩（压缩格式为gz）
compress = false
# 是否打印调用源文件和行号
callerFile = true
# 是否打印调用方法
callerFunction = true

[blockchain]
defCacheSize=128
maxFetchBlockNum=128
timeoutSeconds=5
batchBlockNum=128
driver="memdb"
dbPath="datadir"
dbCache=64
isStrongConsistency=true
singleMode=true
batchsync=false
isRecordBlockSequence=true
isParaChain=false
enableTxQuickIndex=false


[p2p]
seeds=["47.104.125.151:13802","47.104.125.97:13802","47.104.125.177:13802"]
enable=true
isSeed=true
serverStart=true
msgCacheSize=10240
driver="memdb"
dbPath="datadir/addrbook"
dbCache=4
grpcLogFile="grpc33.log"
version=216
verMix=216
verMax=217

[rpc]
jrpcBindAddr="localhost:8801"
grpcBindAddr="localhost:8802"
whitelist=["127.0.0.1"]
jrpcFuncWhitelist=["*"]
grpcFuncWhitelist=["*"]

[mempool]
poolCacheSize=10240
minTxFee=1000000
maxTxNumPerAccount=10000

[consensus]
name="solo"
minerstart=true
genesisBlockTime=1514533394
genesis="14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"
hotkeyAddr="12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"

[consensus.sub.solo]
genesis="14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"
genesisBlockTime=1514533394

[consensus.sub.ticket]
genesisBlockTime=1514533394
[[consensus.sub.ticket.genesis]]
minerAddr="12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"
returnAddr="14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"
count=10000

[[consensus.sub.ticket.genesis]]
minerAddr="1PUiGcbsccfxW3zuvHXZBJfznziph5miAo"
returnAddr="1EbDHAXpoiewjPLX9uqoz38HsKqMXayZrF"
count=10000

[[consensus.sub.ticket.genesis]]
minerAddr="1EDnnePAZN48aC2hiTDzhkczfF39g1pZZX"
returnAddr="1KcCVZLSQYRUwE5EXTsAoQs9LuJW6xwfQa"
count=10000

[store]
name="mavl"
driver="memdb"
dbPath="datadir/mavltree"
dbCache=128

[store.sub.mavl]
enableMavlPrefix=false
enableMVCC=false

[wallet]
minFee=1000000
driver="memdb"
dbPath="datadir/wallet"
dbCache=16
signType="secp256k1"

[wallet.sub.ticket]
minerwhitelist=["*"]

[exec]
isFree=true
minExecFee=0
enableStat=false
enableMVCC=false

[exec.sub.cert]
# 是否启用证书验证和签名
enable=false
# 加密文件路径
cryptoPath="authdir/crypto"
# 带证书签名类型，支持"auth_ecdsa", "auth_sm2"
signType="auth_ecdsa"
`
