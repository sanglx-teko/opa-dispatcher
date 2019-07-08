package configurationmanager

import (
	"encoding/json"
	"sync"

	etcdClient "github.com/coreos/etcd/clientv3"
	m "github.com/sanglx-teko/opa-dispatcher/cores/configurationmanager/model"
	model "github.com/sanglx-teko/opa-dispatcher/model"
	"golang.org/x/net/context"
)

type configurationManager struct {
	// etcd client
	eClient *etcdClient.Client
	ctx     context.Context

	serviceConfig     map[string]*model.ServiceGroup
	serviceConfigLock sync.RWMutex
}

func (c *configurationManager) updateServiceConf() {
	if resp, err := c.eClient.Get(c.ctx, "/common/service"); err == nil && resp.Count > 0 {
		obj := map[string]*model.ServiceGroup{}
		if err = json.Unmarshal(resp.Kvs[0].Value, &obj); err == nil {
			c.serviceConfigLock.RLock()
			defer c.serviceConfigLock.RUnlock()
			c.serviceConfig = obj
		}
	}
}

// InitWithConfig configuration manager interface
func (c *configurationManager) InitWithConfig(conf etcdClient.Config) error {
	client, e := etcdClient.New(conf)
	if e != nil {
		return e
	}
	c.eClient, c.ctx = client, context.Background()
	c.updateServiceConf()
	go c.watchServiceConf()
	return nil
}

// GetServiceConfig configuration manager interface
func (c *configurationManager) GetServiceConfig() (result map[string]*model.ServiceGroup) {
	c.serviceConfigLock.RLock()
	defer c.serviceConfigLock.RUnlock()
	result = c.serviceConfig
	return
}

func (c *configurationManager) watchServiceConf() {
	defer func() {
		if e := recover(); e != nil {

		}
	}()

	watchan := c.eClient.Watch(c.ctx, "/common/service")
	for v := range watchan {
		if v.Err() != nil {
			continue
		}
		c.updateServiceConf()
	}

}

// Instance get instance of configuration manager
var Instance m.IConfigurationManager = &configurationManager{}
