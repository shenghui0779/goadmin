package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"

	log "github.com/gogap/logrus"
)

type Etcd struct {
	Kv clientv3.KV
}

func NewEtcd(etcdaddress string) *Etcd {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdaddress},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Error(err)
		return nil
	}
	kv := clientv3.NewKV(cli)
	etcd := new(Etcd)
	etcd.Kv = kv
	return etcd
}

func (etcdcli *Etcd) Get(k string) (string, error) {
	getResp, err := etcdcli.Kv.Get(context.TODO(), k)
	if err != nil {
		return "", err
	}
	// 输出本次的Revision
	if len(getResp.Kvs) == 0 {
		return "", fmt.Errorf("Key Not Found")
	}

	return string(getResp.Kvs[0].Value), nil
}

func (etcdcli *Etcd) Save(k, v string) error {
	_, err := etcdcli.Kv.Put(context.TODO(), k, v)
	return err
}

func (etcdcli *Etcd) GetAll(k string) ([]string, error) {
	rangeResp, err := etcdcli.Kv.Get(context.TODO(), k, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	if len(rangeResp.Kvs) == 0 {
		return nil, fmt.Errorf("Key Not Found")
	}
	events := make([]string, 0)
	for _, v := range rangeResp.Kvs {
		events = append(events, string(v.Value))
	}
	return events, nil
}
