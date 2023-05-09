package config

import (
	"github.com/hashicorp/consul/api"
)

type Consul struct {
	addr   string
	scheme string

	client *api.Client
}

func NewConsul(addr string, schema string) *Consul {
	return &Consul{
		addr:   addr,
		scheme: schema,
	}
}

func (c *Consul) Connect() error {
	config := &api.Config{
		Address: c.addr,
		Scheme:  c.scheme,
	}

	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *Consul) GetValue(key string) ([]byte, error) {
	kv := c.client.KV()

	result, _, err := kv.Get(key, nil)
	if err != nil {
		return nil, err
	}

	return result.Value, nil
}
