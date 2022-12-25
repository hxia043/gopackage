package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	info    string
	flagvar = pflag.Int("flagname", 1234, "help message for flagname")
)

func main() {
	// Use CommandLine as default flagset
	/*
		version := pflag.StringP("version", "v", "1.0.0", "test version")
		pflag.StringVarP(&info, "info", "i", "", "test info")
		pflag.Parse()

		fmt.Printf("argument number is: %v\n", pflag.NArg())
		fmt.Printf("argument list is: %v\n", pflag.Args())
		fmt.Printf("the first argument is: %v\n", pflag.Arg(0))

		fn, _ := pflag.CommandLine.GetInt("flagname")
		fmt.Println("flagvar is: ", fn)

		fmt.Println("version is: ", *version)
		fmt.Println("info is: ", info)
	*/

	// customize flagset
	demo := pflag.NewFlagSet("demo", pflag.ContinueOnError)
	demoString := demo.StringP("info", "d", "demo", "demo info")
	demo.Parse(os.Args[1:])

	fmt.Println("info is: ", *demoString)
}
