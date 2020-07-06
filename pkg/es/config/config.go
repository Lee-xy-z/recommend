/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:10
 */
package config

import "time"

// Configuration describes the configuration properties needed to connect to an ElasticSearch cluster
type Configuration struct {
	Servers              []string      `mapstructure:"server_urls"`
	Username             string        `mapstructure:"username"`
	Password             string        `mapstructure:"password" json:"-"`
	MaxNumRecommend      int           `mapstructure:"-"`
	MaxRecommendAge      time.Duration `yaml:"max_recommend_age" mapstructure:"-"`
	IndexPrefix          string        `mapstructure:"index_prefix"`
	NumShards            int64         `yaml:"shards" mapstructure:"num_shards"`
	NumReplicas          int64         `yaml:"replicas" mapstructure:"num_replicas"`
	BulkSize             int           `mapstructure:"-"`
	BulkWorkers          int           `mapstructure:"-"`
	BulkFlushInterval    time.Duration `mapstructure:"-"`
	Sniffer              bool          `mapstructure:"sniffer"` //https://github.com/olivere/elastic/wiki/Sniffing
	CreateIndexTemplates bool          `mapstructure:"create_mappings"`
	Version              uint          `mapstructure:"version"`
	Timeout              time.Duration `validate:"min=500" mapstructure:"-"`
}

// ClientBuilder creates new es.Client
type ClientBuilder interface {
}
