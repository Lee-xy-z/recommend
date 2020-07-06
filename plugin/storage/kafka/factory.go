/**
* @Author: Zhangxinyu
* @Date: 2020/7/4 11:34
 */
package kafka

import (
	"flag"
	"github.com/Lee-xy-z/recommend/pkg/kafka/producer"
	"github.com/Lee-xy-z/recommend/storage/rcdstore"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Factory implements storage.Factory and creates write-only storage component backend by kafka.
type Factory struct {
	options Options
	logger  *zap.Logger

	marshaller Marshaller
	producer   sarama.AsyncProducer

	producer.Builder
}

// NewFactory creates a new Factory
func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) CreateRcdWriter() (rcdstore.Writer, error) {
	return nil, nil
}

// AddFlags implements plugin.Configurable
func (f *Factory) AddFlags(flagSet *flag.FlagSet) {
	f.options.AddFlags(flagSet)
}

// InitFromViper implements plugin.Configurable
func (f *Factory) InitFromViper(v *viper.Viper) {
	f.options.InitFromViper(v)
	f.Builder = &f.options.Config
}
