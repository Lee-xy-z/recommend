package main

import (
	"fmt"
	"github.com/Lee-xy-z/recommend/plugin/storage"
	"log"
	"os"
)

func main() {

	//svc := flags.NewService()

	storageFactory, err := storage.NewFactory(storage.FactoryConfigFromEnvAndCLI(os.Args, os.Stderr))
	if err != nil {
		log.Fatalf("Cannot initialize storage factory: %v", err)
	}
	fmt.Println(storageFactory)

	/*command := &cobra.Command{
		Use:   "recommend",
		Short: "Commands for recommend.",
		Long:  "These commands correspond to recommend.",
	}

	command.AddCommand(version.Command())

	command.Execute()

	routers.RegisterAPI()*/

}
