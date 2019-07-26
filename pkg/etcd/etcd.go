package etcd

import (
	"context"
	"errors"
	"eye/pkg/logger"
	"go.etcd.io/etcd/clientv3"
	"time"
)


var(
	ValueCh = make(chan clientv3.WatchResponse,100)
)

type EtcdClient struct {
	Cli *clientv3.Client
	Lease clientv3.Lease
	LeaseResp *clientv3.LeaseGrantResponse
	keepAliveChan <- chan *clientv3.LeaseKeepAliveResponse
	*clientv3.LeaseKeepAliveResponse
	Stop   chan int
}

// 初始化etcd客户端
func NewEtcdClient(addr []string, timeout int64)*EtcdClient{
	etcdClient := &EtcdClient{
		Stop:make(chan int),
	}
	var err error
	etcdClient.Cli, err = clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: time.Duration(timeout) * time.Second,
	})
	if err != nil {
		logger.Error("connect etcd failed, err:%v", err)
		return nil
	}
	etcdClient.Lease = clientv3.NewLease(etcdClient.Cli)

	etcdClient.LeaseResp,err = etcdClient.Lease.Grant(context.TODO(), 1)
	if err != nil {
		logger.Errorf("etcdclient lease grant err %v",err)
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	etcdClient.keepAliveChan,err = etcdClient.Lease.KeepAlive(ctx,etcdClient.LeaseResp.ID)
	if err != nil {
		logger.Errorf("etcd client lease keepalive err: %v",err)
		return nil
	}
	go func() {
		t := time.NewTicker(1*time.Second)
		for {
			select {
			case <-t.C:
				etcdClient.Keepalive()
			}
		}
	}()
	cancel()
	return etcdClient
}

func (e *EtcdClient)Keepalive(){
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_,err = e.Lease.KeepAliveOnce(ctx,e.LeaseResp.ID)
	if err != nil {
		logger.Errorf("etcd client lease keepalive err: %v",err)
		return
	}
	cancel()
}

// 服务发现
func (e *EtcdClient)GetValue(Key string)(*clientv3.GetResponse,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := e.Cli.Get(ctx, Key)
	if err != nil {
		logger.Errorf("client get from etcd failed, err:%v", err)
		return nil,err
	}
	cancel()
	return resp,nil
}

// 服务注册
func (e *EtcdClient)Register(key string,value []byte)error{
	if resp,err := e.GetValue(key);err != nil {
		logger.Errorf("etcd get key err :%v",err)
		return err
	}else {
		for _, v := range resp.Kvs {
			if string(v.Key) == key {
				return errors.New("this key is aready register")
			}
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := e.Cli.Put(ctx, key, string(value),clientv3.WithLease(e.LeaseResp.ID))
	cancel()
	if err != nil {
		logger.Errorf("etcd register key:%v value :%v err",key,value,err)
		return err
	}
	return nil
}

// 更新
func (e *EtcdClient)Put(key string,value []byte)error{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := e.Cli.Put(ctx, key, string(value),clientv3.WithLease(e.LeaseResp.ID))
	cancel()
	if err != nil {
		logger.Errorf("etcd register key:%v value :%v err",key,value,err)
		return err
	}
	return nil
}

// watch多个服务
func (e *EtcdClient)WatchArray(key []string){
	for _,v := range key {
		go func(key string) {
			if err := recover(); err != nil {
				logger.Errorf("watchArray goroutine panic err :%v",err)
			}
			e.watch(v)
		}(v)
	}
}

// watch单个服务
func (e *EtcdClient)watch(key string) {
	for {
		select {
			case <- e.Stop:
				logger.Error("this watch stop")
				return
			case ValueCh <- <- e.Cli.Watch(context.Background(), key):
		}
	}
}