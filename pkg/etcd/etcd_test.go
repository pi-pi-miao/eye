package etcd

import (
	"eye/internal/initialize"
	"fmt"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"sync"
	"testing"
	"time"
)

var(
	wg sync.WaitGroup
	prefix = "everest.firstleap.cn/framework"
	key    = "webserver"
	etcdkey = fmt.Sprintf("%v/%v",prefix,key)
	value   = "127.0.0.1:8081"
	value1 = "127.0.0.1:8082"

)

type etcdConfig struct {
	key   string      `json:"etcdkey"`
	value []byte      `json:"value"`
}

func newEtcdConfig()*etcdConfig {
	return &etcdConfig{}
}

func TestEtcdClient_Register(t *testing.T) {
	initialize.NewEye("/Users/wanglei/Desktop/framework/configs/application.toml").
		InitConfig().
		InitLog()
	if err := NewEtcdClient([]string{"127.0.0.1:2379"},5).Register(etcdkey,[]byte(value)); err != nil {
		t.Errorf("etcd register err :%v",err)
	}
}


func TestEtcdClient_GetValue(t *testing.T) {
	TestEtcdClient_Register(t)
	time.Sleep(3*time.Second)
	if resp,err := NewEtcdClient([]string{"127.0.0.1:2379"},5).GetValue(etcdkey); err != nil {
		t.Errorf("etcd server get key err : %v",err)
	}else {
		fmt.Println("resp is ",resp.Kvs)
		for _, v := range resp.Kvs {
			fmt.Println("value is ",string(v.Value))
			if string(v.Key) == etcdkey {
				fmt.Println("result is ",string(v.Value),"key is ",string(v.Key))
			}
		}
	}
}

func TestEtcdClient_Put(t *testing.T) {
	//TestEtcdClient_Register(t)
	if err := NewEtcdClient([]string{"127.0.0.1:2379"},5).Put(etcdkey,[]byte(value1));err != nil {
		t.Errorf("etcd put err :%v",err)
	}
	TestEtcdClient_GetValue(t)
}

func TestEtcdClient_WatchArray(t *testing.T) {
	TestEtcdClient_Register(t)
	cli := NewEtcdClient([]string{"127.0.0.1:2379"},5)
	//wg.Add(1)
	go func(cli *EtcdClient) {
		fmt.Println("start 1")
		keys := make([]string,0,10)
		keys = append(keys,etcdkey)
		cli.WatchArray(keys)
	}(cli)
	go func() {
		fmt.Println("start 2")
		for wresp := range ValueCh{
			fmt.Println("start 3")
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					continue
				}

				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == etcdkey {
					fmt.Println("watch value is ",string(ev.Kv.Value))
					close(cli.Stop)
					return
				}
				fmt.Printf("get config from etcd, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()
	//cli.Put(etcdkey,[]byte(value1))
	//<- cli.Stop

	time.Sleep(3*time.Second)
}

