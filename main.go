// @since 2022-04-09 19:15:49
// @author acrazing <joking.young@gmail.com>

package main

import (
	"flag"
	"fmt"
	"github.com/acrazing/universal-ingress-controller/pkg/core"
	_ "github.com/acrazing/universal-ingress-controller/pkg/envoy"
	_ "github.com/acrazing/universal-ingress-controller/pkg/nginx"
	"log"
	"os"
)

var version = "v0.1.0"

func main() {
	var configFile string
	var printVersion bool
	flag.StringVar(&configFile, "config", "", "specify config file, could be a local file(default), or configmap/secret with format [cm|secret]:<name>:<path>.")
	flag.BoolVar(&printVersion, "version", false, "show version number")
	flag.Parse()

	if printVersion {
		fmt.Printf("universal-ingress-controller: %s", version)
		os.Exit(0)
	}

	if configFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := core.NewUniversalIngressController(configFile).Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
