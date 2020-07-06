/**
* @Author: Zhangxinyu
* @Date: 2020/7/4 16:40
 */
package kafka

import (
	"flag"
	"fmt"
	"github.com/Lee-xy-z/recommend/pkg/kafka/producer"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"log"
	"strings"
)

const (
	// EncodingJSON is used for recommendations encoded as JSON.
	EncodingJSON = "json"
	// EncodingProto is used for recommendations encoded as Protobuf.
	EncodingProto          = "protobuf"
	configPrefix           = "kafka.producer"
	suffixBrokers          = ".brokers"
	suffixTopic            = ".topic"
	suffixEncoding         = ".encoding"
	suffixRequiredAcks     = ".required-acks"
	suffixCompression      = ".compression"
	suffixCompressionLevel = ".compression-level"
	suffixProtocolVersion  = ".protocol-version"
	suffixBatchSize        = ".batch-size"
	suffixBatchMaxMessages = ".batch-max-messages"
	suffixBatchLinger      = ".batch-linger"

	defaultBroker           = "127.0.0.1:9092"
	defaultTopic            = "recommend"
	defaultRequiredAcks     = "local"
	defaultCompression      = "none"
	defaultCompressionLevel = 0
	defaultBatchSize        = 0
	defaultBatchMaxMessages = 0
	defaultBatchLinger      = 0
	defaultEncoding         = EncodingJSON
)

var (
	// AllEncoding is a list of all supported encodings.
	AllEncodings = []string{EncodingJSON, EncodingProto}

	// requiredAcks is mapping of sarama supported requiredAcks
	requiredAcks = map[string]sarama.RequiredAcks{
		"noack": sarama.NoResponse,
		"local": sarama.WaitForLocal,
		"all":   sarama.WaitForAll,
	}
	// compressionModes is a mapping of supported CompressionType to compressionCodec along with default, min, max, compression level
	compressionModes = map[string]struct {
		compressor              sarama.CompressionCodec
		defaultCompressionLevel int
		minCompressionLevel     int
		maxCompressionLevel     int
	}{
		"none": {
			compressor:              sarama.CompressionNone,
			defaultCompressionLevel: 0,
		},
		"gzip": {
			compressor:              sarama.CompressionGZIP,
			defaultCompressionLevel: 6,
			minCompressionLevel:     1,
			maxCompressionLevel:     9,
		},
		"snappy": {
			compressor:              sarama.CompressionSnappy,
			defaultCompressionLevel: 0,
		},
		"lz4": {
			compressor:              sarama.CompressionLZ4,
			defaultCompressionLevel: 9,
			minCompressionLevel:     1,
			maxCompressionLevel:     17,
		},
		"zstd": {
			compressor:              sarama.CompressionZSTD,
			defaultCompressionLevel: 3,
			minCompressionLevel:     -131072,
			maxCompressionLevel:     22,
		},
	}
)

// Options stores the configuration options for Kafka
type Options struct {
	Config   producer.Configuration
	Topic    string `mapstructure:"topic"`
	Encoding string `mapstructure:"encoding"`
}

// AddFlags adds flags for Options
func (opt *Options) AddFlags(flagSet *flag.FlagSet) {
	flagSet.String(configPrefix+suffixBrokers, defaultBroker, "The comma-separated list of kafka brokers. i.e. '127.0.0.1:9092,0.0.0:1234'")
	flagSet.String(configPrefix+suffixTopic, defaultTopic, "The name of the kafka topic")
	flagSet.String(configPrefix+suffixProtocolVersion, "", "Kafka protocol version - must be supported by kafka server")
	flagSet.String(configPrefix+suffixEncoding, defaultEncoding, fmt.Sprintf(`Encoding of recommendations ("%s" or "%s") sent to kafka.`, EncodingJSON, EncodingProto))
	flagSet.String(configPrefix+suffixRequiredAcks, defaultRequiredAcks, "Required kafka broker acknowledgement. i.e. noack, local, all")
	flagSet.String(configPrefix+suffixCompression, defaultCompression, "Type of compression (none, gzip, snappy, lz4, zstd) to use on messages")
	flagSet.Int(configPrefix+suffixCompressionLevel, defaultCompressionLevel, "compression level to use on messages. gzip = 1-9 (default = 6), snappy = none, lz4 = 1-17 (default = 9), zstd = -131072 - 22 (default = 3)")
	flagSet.Duration(configPrefix+suffixBatchLinger, defaultBatchLinger, "Time interval to wait before sending records to Kafka. Higher value reduce request to Kafka but increase latency and possibility of data loss in case of process restart. See https://kafka.apache.org/documentation/")
	flagSet.Int(configPrefix+suffixBatchMaxMessages, defaultBatchMaxMessages, "Number of message to batch before sending records to Kafka. Higher value reduce request to Kafka but increase latency and possibility of data loss in case of process restart. See https://kafka.apache.org/documentation/")
	flagSet.Int(configPrefix+suffixBatchSize, defaultBatchSize, "Number of bytes to batch before sending records to Kafka. Higher value reduce request to Kafka but increase latency and possibility of data loss in case of process restart. See https://kafka.apache.org/documentation/")
}

// InitFromViper initializes Options with properties from viper
func (opt *Options) InitFromViper(v *viper.Viper) {
	requiredAcks, err := getRequiredAcks(v.GetString(configPrefix + suffixRequiredAcks))
	if err != nil {
		log.Fatal(err)
	}

	compressionMode := strings.ToLower(v.GetString(configPrefix + suffixCompression))
	compressionCodec, err := getCompressionMode(compressionMode)
	if err != nil {
		log.Fatal(err)
	}

	compressionLevel, err := getCompressionLevel(compressionMode, v.GetInt(configPrefix+suffixCompressionLevel))
	if err != nil {
		log.Fatal(err)
	}

	opt.Config = producer.Configuration{
		Brokers:          strings.Split(stripWithSpace(v.GetString(configPrefix+suffixBrokers)), ","),
		RequiredAcks:     requiredAcks,
		Compression:      compressionCodec,
		CompressionLevel: compressionLevel,
		ProtocolVersion:  v.GetString(configPrefix + suffixProtocolVersion),
		BatchLinger:      v.GetDuration(configPrefix + suffixBatchLinger),
		BatchSize:        v.GetInt(configPrefix + suffixBatchSize),
		BatchMaxMessages: v.GetInt(configPrefix + suffixBatchMaxMessages),
	}

	opt.Topic = v.GetString(configPrefix + suffixTopic)
	opt.Encoding = v.GetString(configPrefix + suffixEncoding)
}

// stripWithSpace removes all whitespace character from a string
func stripWithSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}

// getCompressionLevel to get compression level from compression type
func getCompressionLevel(mode string, compressionLevel int) (int, error) {
	compressionModeData, ok := compressionModes[mode]
	if !ok {
		return 0, fmt.Errorf("cannot find compression mode for compressionMode %v", mode)
	}
	if compressionLevel == defaultCompressionLevel {
		return compressionModeData.defaultCompressionLevel, nil
	}

	if compressionModeData.minCompressionLevel > compressionLevel || compressionModeData.maxCompressionLevel < compressionLevel {
		return 0, fmt.Errorf("compression level %d for '%s' is not within valid range [%d, %d]", compressionLevel, mode, compressionModeData.minCompressionLevel, compressionModeData.maxCompressionLevel)
	}

	return compressionLevel, nil
}

// getCompressionMode maps input modes to sarama CompressionCodec
func getCompressionMode(mode string) (sarama.CompressionCodec, error) {
	compressionMode, ok := compressionModes[mode]
	if !ok {
		return 0, fmt.Errorf("unknow compression mode: %v", mode)
	}
	return compressionMode.compressor, nil
}

// getRequiredAcks maps input ack values to sarama requiredAcks
func getRequiredAcks(acks string) (sarama.RequiredAcks, error) {
	requiredAcks, ok := requiredAcks[strings.ToLower(acks)]
	if !ok {
		return 0, fmt.Errorf("unknow Required Ack: %s", acks)
	}
	return requiredAcks, nil
}
