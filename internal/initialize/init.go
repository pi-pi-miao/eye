package initialize

import (
	"time"
	"fmt"
	"errors"
	"io/ioutil"
	"net/http"
	"eye/internal/router/httpApi"
	"eye/pkg/logger"
	"github.com/BurntSushi/toml"
	"github.com/julienschmidt/httprouter"
)

var (
	Config *eye
)

type (
	eye struct {
		ConfPath string
		handler  http.Handler
		Config   conf
		EtcdConfig etcdConfig
		Log       log
	}
	conf struct {
		Ip           string
		Port         int
		Route        string
		PprofAddr    string
		ReadTimeout  int
		WriteTimeout int
		IdleTimeout  int
		Lenth        int
		Env          string
	}
	etcdConfig struct {
		EtcdAddr     []string
		DialTimeout   int64
	}
	log struct {
		LogLevel string
		LogFile  string
		IsDebug  int
		DingUrl  string
	}
)

func NewEye(path string)*eye{
	return &eye{
		ConfPath:path,
		EtcdConfig:etcdConfig{
			EtcdAddr:make([]string,0,10),
		},
	}
}

func (f *eye)InitAll(){
	f.InitConfig().
		InitLog().
		init().
		run()
}

func (f *eye) InitConfig() *eye {
	configBytes, err := ioutil.ReadFile(f.ConfPath)
	if err != nil {
		fmt.Printf("read config err %v", err)
		panic(err)
	}
	fmt.Println("config is ", string(configBytes))
	if _, err := toml.Decode(string(configBytes), f); err != nil {
		fmt.Printf("toml decode err %v", err)
		panic(err)
	}
	fmt.Println("config is", f.Config,"etcd config is ",f.EtcdConfig)
	return f
}

func (f *eye)InitLog()*eye {
	if len(f.Log.DingUrl) == 0 && f.Log.IsDebug < 0 && len(f.Log.DingUrl) == 0 && len(f.Log.LogLevel) == 0 {
		panic(errors.New("please read config file and do InitConfig() function"))
		return nil
	}
	if err := logger.InitLogger(&logger.LogConfig{
		LogLevel:f.Log.LogLevel,
		LogFile:f.Log.LogFile,
		IsDebug:f.Log.IsDebug,
	},f.Config.Env,f.Log.DingUrl);err != nil {
		fmt.Printf("set logger file and level err %v",err)
		panic(err)
	}
	logger.Debug("this log init succ")
	return f
}

func (f *eye)init()*eye{
	handler := httprouter.New()
	f.handler = handler
	httpApi.Register(handler)
	return f
}

func (f *eye)run(){
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("framwork main goroutine panic err is %v", err)
		}
		logger.SugaredLogger.Sync()
	}()
	go func(f *eye) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("framwork main goroutine panic err is %v", err)
			}
		}()
		fmt.Println(http.ListenAndServe(f.Config.PprofAddr, nil))
	}(f)
	httpServer := &http.Server{
		ReadTimeout:  time.Duration(f.Config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(f.Config.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(f.Config.IdleTimeout) * time.Second,
		Handler:      f.handler,
		Addr:         fmt.Sprintf("%v:%v", f.Config.Ip, f.Config.Port),
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		logger.Errorf("this server ip is %v port is %v start err is %v", f.Config.Ip, f.Config.Port, err)
		panic(err)
	}
}