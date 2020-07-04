/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:05
 */
package storage

import (
	"github.com/Lee-xy-z/recommend/plugin/storage/es"
	"github.com/Lee-xy-z/recommend/storage"
	"github.com/Lee-xy-z/recommend/storage/rcdstore"
)

const (
	elasticsearchStorageType = "elasticsearch"
	memoryStorageType        = "memory"
	kafkaStorageType         = "kafka"
)

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
	//var writers []rcdstore.Writer

	return nil, nil
}

func (f *Factory) getFactoryOfType(factoryType string) (storage.Factory, error) {
	switch factoryType {
	case elasticsearchStorageType:
		return es.NewFactory(), nil
	}
	return nil, nil
}
