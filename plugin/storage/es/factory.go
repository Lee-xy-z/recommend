/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:09
 */
package es

import (
	"flag"
	"github.com/Lee-xy-z/recommend/pkg/es"
	"github.com/Lee-xy-z/recommend/pkg/es/config"
	"github.com/Lee-xy-z/recommend/storage/rcdstore"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	primaryNamespace = "es"
	archiveNamespace = "es-archive"
)

// Factory implements storage.Factory for Elasticsearch backend.
type Factory struct {
	Options *Options

	logger        *zap.Logger
	primaryConfig config.ClientBuilder
	primaryClient es.Client
}

// NewFactory creates a new Factory
func NewFactory() *Factory {
	return &Factory{
		Options: NewOptions(primaryNamespace, archiveNamespace),
	}
}

// CreateRcdWriter implements storage.Factory
func (f *Factory) CreateRcdWriter() (rcdstore.Writer, error) {
	return createRcdWriter(f.logger, f.primaryClient, f.primaryConfig, false)
}

func createRcdWriter(logger *zap.Logger, client es.Client, cfg config.ClientBuilder, archive bool) (rcdstore.Writer, error) {
	return nil, nil
}

// AddFlags implements plugin.Configurable
func (f *Factory) AddFlags(flagSet *flag.FlagSet) {
	f.Options.AddFlags(flagSet)
}

// InitFromViper implements plugin.Configurable
func (f *Factory) InitFromViper(v *viper.Viper) {
	f.Options.InitFromViper(v)

}
