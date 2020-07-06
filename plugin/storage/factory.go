/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:05
 */
package storage

import (
	"flag"
	"fmt"
	"github.com/Lee-xy-z/recommend/plugin"
	"github.com/Lee-xy-z/recommend/plugin/storage/es"
	"github.com/Lee-xy-z/recommend/plugin/storage/kafka"
	"github.com/Lee-xy-z/recommend/plugin/storage/memory"
	"github.com/Lee-xy-z/recommend/storage"
	"github.com/Lee-xy-z/recommend/storage/rcdstore"
	"github.com/spf13/viper"
)

const (
	elasticsearchStorageType = "elasticsearch"
	memoryStorageType        = "memory"
	kafkaStorageType         = "kafka"
)

// AllStorageTypes defines all available storage backends
var AllStorageTypes = []string{elasticsearchStorageType, memoryStorageType, kafkaStorageType}

// Factory implements storage.Factory interface as a meta-factory for storage components.
type Factory struct {
	FactoryConfig
	factories map[string]storage.Factory
}

// New Factory creates the meta-factory.
func NewFactory(config FactoryConfig) (*Factory, error) {
	f := &Factory{FactoryConfig: config}
	uniqueTypes := map[string]struct{}{
		f.RcdReaderTypes: {},
	}

	for _, storageType := range f.RcdWriterTypes {
		uniqueTypes[storageType] = struct{}{}
	}

	f.factories = make(map[string]storage.Factory)

	for t := range uniqueTypes {
		ff, err := f.getFactoryOfType(t)
		if err != nil {
			return nil, err
		}

		f.factories[t] = ff
	}
	return f, nil
}

// CreateRcdWriter implements storage.Factory
func (f *Factory) CreateRcdWriter() (rcdstore.Writer, error) {
	return nil, nil
}

func (f *Factory) getFactoryOfType(factoryType string) (storage.Factory, error) {
	switch factoryType {
	case elasticsearchStorageType:
		return es.NewFactory(), nil
	case memoryStorageType:
		return memory.NewFactory(), nil
	case kafkaStorageType:
		return kafka.NewFactory(), nil
	default:
		return nil, fmt.Errorf("unknown storage type %s. Valid types are %v", factoryType, AllStorageTypes)
	}

}

// AddFlags implements plugin.Configurable
func (f *Factory) AddFlags(flagSet *flag.FlagSet) {

}

// InitFromViper implements plugin.Configurable
func (f *Factory) InitFromViper(v *viper.Viper) {
	for _, factory := range f.factories {
		if conf, ok := factory.(plugin.Configurable); ok {
			conf.InitFromViper(v)
		}
	}
}
