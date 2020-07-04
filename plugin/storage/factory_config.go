/**
* @Author: Zhangxinyu
* @Date: 2020-07-03 17:11
 */
package storage

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	RcdStorageTypeEnvVar = "RCD_STORAGE_TYPE"

	rcdStorageFlag = "rcd-storage.type"
)

// FactoryConfig tells the Factory which types of backends it needs to create for different storage types.
type FactoryConfig struct {
	RcdWriterTypes []string
	RcdReaderTypes string
}

// FactoryConfigFromEnvAndCLI reads the desired types of storage backends from RCD_STORAGE_TYPE and
// DEPENDENCY_STORAGE_TYPE environment variables. Allowed values:
// * `elasticsearch` - built-in
// * `kafka` - built-in
// * `memory` - built-in
// It also parse the args --rcd-storage.type flag. But use environment variables whenever possible.
func FactoryConfigFromEnvAndCLI(args []string, log io.Writer) FactoryConfig {

	rcdStorageType := os.Getenv(RcdStorageTypeEnvVar)
	if rcdStorageType == "" {
		// check command line for --rcd-storage.type flag
		rcdStorageType = rcdStorageTypeFromArgs(args, log)
	}
	if rcdStorageType == "" {
		rcdStorageType = elasticsearchStorageType
	}
	rcdWriterTypes := strings.Split(rcdStorageType, ",")
	if len(rcdWriterTypes) > 1 {
		fmt.Fprintf(log, "WARNING: multiple rcd storage types have been specified. "+
			"Only the first type (%s) will be used for reading and archiving.\n\n", rcdWriterTypes[0])
	}

	return FactoryConfig{
		RcdWriterTypes: rcdWriterTypes,
		RcdReaderTypes: rcdWriterTypes[0],
	}

}

func rcdStorageTypeFromArgs(args []string, log io.Writer) string {
	for i, token := range args {
		if i == 0 {
			continue // skip app name; easier than dealing with +-1 offset
		}
		if !strings.HasPrefix(token, rcdStorageFlag) {
			continue
		}
		fmt.Fprintf(log, "WARNING: please use environment variable %s whenever possible instead of command line %s", RcdStorageTypeEnvVar, token)

		if token == rcdStorageFlag && i < len(args)-1 {
			return args[i+1]
		}
		if strings.HasPrefix(token, rcdStorageFlag+"=") {
			return token[(len(rcdStorageFlag) + 1):]
		}

		break
	}
	return ""
}
