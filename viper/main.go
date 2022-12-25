package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("ContentDir", "content")
	fmt.Println(viper.GetString("ContentDir"))

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/etc/viper")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(viper.GetString("name"))

	var info string
	demo := pflag.NewFlagSet("demo", pflag.ContinueOnError)
	demo.StringVarP(&info, "info", "d", "demo", "demo info")
	demo.Parse(os.Args[1:])

	viper.BindPFlags(demo)

	fmt.Println(viper.GetString("info"))
}
