## Eye

```
eye是一个基于httprouter为路由，底层为官方标准库net/http为网络框架的企业级标准web框架，微服务框架。可直接开箱即用
```

## 集成：

``` shell
etcdv3
zap
json
钉钉报警
mod
toml
redis 未来更新
mysql 未来更新
```

## 接口

``` go
// pkg/etcd/etcd.go
func NewEtcdClient(addr []string, timeout int64)*EtcdClien
初始化etcd接口，

func (e *EtcdClient)GetValue(Key string)(*clientv3.GetResponse,error)
服务发现接口

func (e *EtcdClient)Register(key string,value []byte)error
服务注册接口

func (e *EtcdClient)Put(key string,value []byte)error
更新服务接口

func (e *EtcdClient)WatchArray(key []string)
watch监听

租约续约我设置为1秒，为了防止机器宕机造成太长时间服务一直存在etcd的问题

//pkg/json/json.go
func Marshal(v interface{}) ([]byte, error)
json的marshal接口

func Unmarshal(data []byte, v interface{}) error
json的unmarshal接口
```

## 配置文件说明

``` toml
[Config]
Ip = "127.0.0.1"                 # 当前的运行的ip
Port = 8080                      # 当前的运行的端口
PprofAddr = "127.0.0.1:10000"    # pprof的监控地址
ReadTimeout = 10         
WriteTimeout = 10
IdleTimeout  = 10
Lenth        = 10
Env          = "test"             # test 测试   pro生产环境



[EtcdConfig]
EtcdAddr = ["127.0.0.1:2379",]     # etcd的地址
DialTimeout = 5                    # 链接超时

[Log]
LogLevel = "debug"                 # 日志级别
LogFile = "./framework.log"        # 打印的日志路径
IsDebug = 1
DingUrl = "http://"                # 钉钉报警链接地址
```

