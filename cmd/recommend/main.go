package main

import (
	"fmt"
	"github.com/Lee-xy-z/recommend/plugin/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	//svc := flags.NewService()

	storageFactory, err := storage.NewFactory(storage.FactoryConfigFromEnvAndCLI(os.Args, os.Stderr))
	if err != nil {
		log.Fatalf("Cannot initialize storage factory: %v", err)
	}

	v := viper.New()
	command := &cobra.Command{
		Use:   "recommend",
		Short: "Recommend consumes from Kafka and writes to storage.",
		Long:  `Recommend consumes events from a particular Kafka topic and writes them handled to a configured storage.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			storageFactory.InitFromViper(v)

			return nil
		},
	}
	fmt.Println(command)

	/*command := &cobra.Command{
		Use:   "recommend",
		Short: "Commands for recommend.",
		Long:  "These commands correspond to recommend.",
	}

	command.AddCommand(version.Command())

	command.Execute()

	routers.RegisterAPI()*/

}
