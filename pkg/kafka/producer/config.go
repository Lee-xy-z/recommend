/**
* @Author: Zhangxinyu
* @Date: 2020/7/4 16:42
 */
package producer

import (
	"github.com/Shopify/sarama"
	"time"
)

// Builder builds a new kafka producer
type Builder interface {
	NewProducer() (sarama.AsyncProducer, error)
}

// Configuration describes the configuration properties needed to create a Kafka producer
type Configuration struct {
	Brokers          []string                `mapstructure:"brokers"`
	RequiredAcks     sarama.RequiredAcks     `mapstructure:"required_acks"`
	Compression      sarama.CompressionCodec `mapstructure:"compression"`
	CompressionLevel int                     `mapstructure:"compression_level"`
	ProtocolVersion  string                  `mapstructure:"protocol_version"`
	BatchLinger      time.Duration           `mapstructure:"batch_linger"`
	BatchSize        int                     `mapstructure:"batch_size"`
	BatchMaxMessages int                     `mapstructure:"batch_max_messages"`
}

// NewProducer creates a new asynchronous kafka producer
func (c *Configuration) NewProducer() (sarama.AsyncProducer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.RequiredAcks = c.RequiredAcks
	saramaConfig.Producer.Compression = c.Compression
	saramaConfig.Producer.CompressionLevel = c.CompressionLevel
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.Flush.Bytes = c.BatchSize
	saramaConfig.Producer.Flush.Frequency = c.BatchLinger
	saramaConfig.Producer.Flush.MaxMessages = c.BatchMaxMessages

	if len(c.ProtocolVersion) > 0 {
		ver, err := sarama.ParseKafkaVersion(c.ProtocolVersion)
		if err != nil {
			return nil, err
		}
		saramaConfig.Version = ver
	}

	return sarama.NewAsyncProducer(c.Brokers, saramaConfig)
}
