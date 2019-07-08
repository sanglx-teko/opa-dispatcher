package model

import (
	etcdClient "github.com/coreos/etcd/clientv3"
	"github.com/sanglx-teko/opa-dispatcher/model"
)

// IConfigurationManager ...
type IConfigurationManager interface {
	InitWithConfig(etcdClient.Config) error
	GetServiceConfig() map[string]*model.ServiceGroup
}

// ServiceConfig ...
type ServiceConfig struct {
	Name         string
	ServiceGroup string
}
