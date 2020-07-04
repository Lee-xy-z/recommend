/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:10
 */
package es

import "github.com/Lee-xy-z/recommend/pkg/es/config"

const (
	defaultServerURL = "http://127.0.0.1:9200"
)

type namespaceConfig struct {
	config.Configuration
	namespace string
}

// Options contains various type of Elasticsearch configs and provides the ability
// to bind them to command line flag and apply overlays,so that some configurations
//(e.g. archive) may be underspecified and infer the rest of its parameters from primary
type Options struct {
	Primary namespaceConfig
	others  map[string]*namespaceConfig
}

// NewOptions creates a new Options struct.
func NewOptions(primaryNamespace string) *Options {
	//TODO all default values should be defined via cobra flags
	options := &Options{
		Primary: namespaceConfig{
			Configuration: config.Configuration{
				Servers:  []string{defaultServerURL},
				Username: "",
				Password: "",
			},
			namespace: primaryNamespace,
		},
		others: nil,
	}
	return options
}
