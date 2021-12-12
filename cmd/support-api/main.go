package main

import "flag"

var configFile = flag.String("f", "config/open.yml", "set config file which viper will loading.")

func main() {
	flag.Parse()
}
