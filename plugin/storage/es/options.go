/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:10
 */
package es

import (
	"flag"
	"github.com/Lee-xy-z/recommend/pkg/es/config"
	"github.com/spf13/viper"
	"strings"
	"time"
)

const (
	suffixUsername = ".username"
	suffixPassword = ".password"
	suffixSniffer  = ".sniffer"
	suffixTimeout  = ".timeout"

	suffixServerURLs          = ".server-urls"
	suffixMaxRecommendAge     = ".max-recommend-age"
	suffixMaxNumRecommend     = ".max-num-recommendations"
	suffixNumShards           = ".num-shards"
	suffixNumReplicas         = ".num-replicas"
	suffixBulkSize            = ".bulk.size"
	suffixBulkWorkers         = ".bulk.workers"
	suffixBulkFlushInterval   = ".bulk.flush-interval"
	suffixIndexPrefix         = ".index-prefix"
	suffixCreateIndexTemplate = ".create-index-templates"
	suffixVersion             = ".version"

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

// AddFlags adds flags for Options
func (opt *Options) AddFlags(flagSet *flag.FlagSet) {
	addFlags(flagSet, &opt.Primary)
}

// InitFromViper initializes Options with properties from viper
func (opt *Options) InitFromViper(v *viper.Viper) {
	initFromViper(&opt.Primary, v)
}

func initFromViper(cfg *namespaceConfig, v *viper.Viper) {
	cfg.Username = v.GetString(cfg.namespace + suffixUsername)
	cfg.Password = v.GetString(cfg.namespace + suffixPassword)
	cfg.Sniffer = v.GetBool(cfg.namespace + suffixSniffer)
	cfg.Servers = strings.Split(stripWhiteSpace(v.GetString(cfg.namespace+suffixServerURLs)), ",")
	cfg.MaxNumRecommend = v.GetInt(cfg.namespace + suffixMaxNumRecommend)
	cfg.MaxRecommendAge = v.GetDuration(cfg.namespace + suffixMaxRecommendAge)
	cfg.NumShards = v.GetInt64(cfg.namespace + suffixNumShards)
	cfg.NumReplicas = v.GetInt64(cfg.namespace + suffixNumReplicas)
	cfg.BulkSize = v.GetInt(cfg.namespace + suffixBulkSize)
	cfg.BulkWorkers = v.GetInt(cfg.namespace + suffixBulkWorkers)
	cfg.BulkFlushInterval = v.GetDuration(cfg.namespace + suffixBulkFlushInterval)
	cfg.Timeout = v.GetDuration(cfg.namespace + suffixTimeout)
	cfg.IndexPrefix = v.GetString(cfg.namespace + suffixIndexPrefix)
	cfg.Version = uint(v.GetInt(cfg.namespace + suffixVersion))
	cfg.CreateIndexTemplates = v.GetBool(cfg.namespace + suffixCreateIndexTemplate)
}

// stripWhiteSpace removes all whitespace characters from a string
func stripWhiteSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}

func addFlags(flagSet *flag.FlagSet, nsConfig *namespaceConfig) {
	flagSet.String(nsConfig.namespace+suffixUsername, nsConfig.Username, "The username required by Elasticsearch.")
	flagSet.String(nsConfig.namespace+suffixPassword, nsConfig.Password, "The password required by Elasticsearch")
	flagSet.Bool(nsConfig.namespace+suffixSniffer, nsConfig.Sniffer, "The sniffer config for Elasticsearch; client uses sniffing process to find all nodes automatically, disable if not required")
	flagSet.String(nsConfig.namespace+suffixServerURLs, defaultServerURL, "The comma-separated list of Elasticsearch server, must be full url i.e. http://localhost:9200")
	flagSet.Duration(nsConfig.namespace+suffixTimeout, nsConfig.Timeout, "Timeout used for queries. A Timeout of zero means no timeout")
	flagSet.Duration(nsConfig.namespace+suffixMaxRecommendAge, nsConfig.MaxRecommendAge, "The maximum lookback for recommend in Elasticsearch")
	flagSet.Int(nsConfig.namespace+suffixMaxNumRecommend, nsConfig.MaxNumRecommend, "The maximum number of recommendations to fetch at a time per query in Elasticsearch")
	flagSet.Int64(nsConfig.namespace+suffixNumShards, nsConfig.NumShards, "The number of shards per index in Elasticsearch")
	flagSet.Int64(nsConfig.namespace+suffixNumReplicas, nsConfig.NumReplicas, "The number of replicas per index in Elasticsearch")
	flagSet.Int(nsConfig.namespace+suffixBulkSize, nsConfig.BulkSize, "The number of bytes that bulk requests can take up before the bulk processor decides to commit")
	flagSet.Int(nsConfig.namespace+suffixBulkWorkers, nsConfig.BulkWorkers, "The number of workers that are able to receive bulk requests and eventually commit them to Elasticsearch")
	flagSet.Duration(nsConfig.namespace+suffixBulkFlushInterval, nsConfig.BulkFlushInterval, "A time.Duration after which bulk requests are committed, regardless of other thresholds. Set to zero to disable. By default, this is disabled.")
	flagSet.String(nsConfig.namespace+suffixIndexPrefix, nsConfig.IndexPrefix, "Optional prefix of recommend indices. For example \"production\" creates \"production-recommend-*\".")
	flagSet.Bool(nsConfig.namespace+suffixCreateIndexTemplate, nsConfig.CreateIndexTemplates, "Creates index templates at application startup. Set to false when templates are installed manually.")
	flagSet.Uint(nsConfig.namespace+suffixVersion, 0, "The major Elasticsearch version. If not specified, the value will be auto-detected from Elasticsearch.")
}

// NewOptions creates a new Options struct.
func NewOptions(primaryNamespace string, otherNamespaces ...string) *Options {
	//TODO all default values should be defined via cobra flags
	options := &Options{
		Primary: namespaceConfig{
			Configuration: config.Configuration{
				Servers:              []string{defaultServerURL},
				Username:             "",
				Password:             "",
				MaxNumRecommend:      10000,
				MaxRecommendAge:      72 * time.Hour,
				NumShards:            5,
				NumReplicas:          1,
				BulkSize:             5 * 1000 * 1000,
				Sniffer:              false,
				BulkFlushInterval:    time.Millisecond * 200,
				BulkWorkers:          1,
				CreateIndexTemplates: true,
			},
			namespace: primaryNamespace,
		},
		others: make(map[string]*namespaceConfig, len(otherNamespaces)),
	}

	for _, namespace := range otherNamespaces {
		options.others[namespace] = &namespaceConfig{namespace: namespace}
	}
	return options
}
