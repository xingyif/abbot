package driver

import (
	"fmt"

	"arhat.dev/abbot/pkg/types"
)

type key struct {
	name string
	os   string
}

type factory struct {
	newDriver FactoryFunc
	newConfig ConfigFactoryFunc
}

type (
	FactoryFunc       func(cfg interface{}) (types.Driver, error)
	ConfigFactoryFunc func() interface{}
)

var supportedDrivers = make(map[key]factory)

func Register(name, os string, newDriver FactoryFunc, newDriverConfig ConfigFactoryFunc) {
	supportedDrivers[key{
		name: name,
		os:   os,
	}] = factory{
		newDriver: newDriver,
		newConfig: newDriverConfig,
	}
}

func NewDriver(name, os string, cfg interface{}) (types.Driver, error) {
	f, ok := supportedDrivers[key{
		name: name,
		os:   os,
	}]
	if !ok {
		return nil, fmt.Errorf("driver for %s on %s not found", name, os)
	}

	return f.newDriver(cfg)
}

func NewConfig(name, os string) (interface{}, error) {
	f, ok := supportedDrivers[key{
		name: name,
		os:   os,
	}]
	if !ok {
		return nil, fmt.Errorf("driver config for %s on %s not found", name, os)
	}

	return f.newConfig(), nil
}