/**
* @Author: Zhangxinyu
* @Date: 2020/7/4 21:24
 */
package plugin

import (
	"flag"
	"github.com/spf13/viper"
)

// Configurable interface can be implemented by plugins that require external configuration,
// such as CLI flags, config files, or environment variables.
type Configurable interface {
	// AddFlags adds CLI flags for configuring this component.
	AddFlags(flagSet *flag.FlagSet)

	// InitFromViper initializes this component with properties from spf13/viper.
	InitFromViper(v *viper.Viper)
}
