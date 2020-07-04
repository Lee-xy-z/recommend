/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:10
 */
package config

// Configuration describes the configuration properties needed to connect to an ElasticSearch cluster
type Configuration struct {
	Servers     []string `mapstructure:"server_urls"`
	Username    string   `mapstructure:"username"`
	Password    string   `mapstructure:"password" json:"-"`
	IndexPrefix string   `mapstructure:"index_prefix"`
}

// ClientBuilder creates new es.Client
type ClientBuilder interface {
}
