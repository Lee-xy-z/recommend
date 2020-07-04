/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:09
 */
package es

import (
	"github.com/Lee-xy-z/recommend/pkg/es"
	"github.com/Lee-xy-z/recommend/pkg/es/config"
	"github.com/Lee-xy-z/recommend/storage/rcdstore"
	"go.uber.org/zap"
)

const primaryNamespace = "es"

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
		Options: NewOptions(primaryNamespace),
	}
}

// CreateRcdWriter implements storage.Factory
func (f *Factory) CreateRcdWriter() (rcdstore.Writer, error) {
	return createRcdWriter(f.logger, f.primaryClient, f.primaryConfig, false)
}

func createRcdWriter(logger *zap.Logger, client es.Client, cfg config.ClientBuilder, archive bool) (rcdstore.Writer, error) {
	return nil, nil
}
