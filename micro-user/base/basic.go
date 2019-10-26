package base

import (
	"github.com/entere/go-demos/micro-user/base/config"
	"github.com/entere/go-demos/micro-user/base/pkg/db"
	"github.com/spf13/pflag"
)

// Init ...
func Init() {

	// init config
	pflag.Parse()
	cfg := pflag.StringP("config", "c", "", "apiserver config file path.")
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	db.MyDB.Init()
	defer db.MyDB.Close()

}
