module eye

go 1.12

require (
	github.com/julienschmidt/httprouter v1.2.0
	github.com/BurntSushi/toml v0.3.1
	github.com/bilibili/kratos v0.1.0
	github.com/sirupsen/logrus v1.4.1
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 // indirect
	golang.org/x/sys v0.0.0-20190312061237-fead79001313 // indirect
	google.golang.org/grpc v1.20.1
	go.etcd.io/etcd v0.0.0-20190720005121-fe86a786a4c3
	github.com/gorilla/websocket v1.2.0
	go.etcd.io/etcd v3.3.13+incompatible
	github.com/json-iterator/go v1.1.7
	go.uber.org/zap v1.2.3
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190609082536-301114b31cce
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190612232758-d4e310b4a8a5
	google.golang.org/appengine => github.com/golang/appengine v1.6.0
)
