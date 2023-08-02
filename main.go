package main

import (
	"flag"

	"github.com/arbrix/uadmin"
	"github.com/arbrix/vilnaCMS/storage"
)

var pathToDBConf = flag.String("pathToDBConf", "local/conf.json", "set config path for db")

func main() {
	flag.Parse()
	db, err := storage.InitDB(*pathToDBConf)
	if err != nil {
		panic(err)
	}

	uadmin.Database = db

	uadmin.StartServer()
}
