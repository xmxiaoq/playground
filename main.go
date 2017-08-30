package main

import (
	"fmt"

	golua "github.com/Shopify/go-lua"
	"github.com/xmxiaoq/gofs"
	"github.com/xmxiaoq/golog"
	"github.com/yuin/gopher-lua"
	"go.uber.org/zap"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func DoGopherLua() {
	L := lua.NewState()
	defer L.Close()
	var err error
	err = L.DoFile(`tests.lua`)
	//err = L.DoFile(`main.lua`)
	if err != nil {
		sugar.Infof("%v", err)
	}
}

func DoGoLua() {
	l := golua.NewState()
	golua.OpenLibraries(l)

	if err := golua.DoFile(l, "collections.lua"); err != nil {
		panic(err)
	}

	if err := golua.DoFile(l, "tests.lua"); err != nil {
		panic(err)
	}
}

func main() {
	logger = golog.Logger
	sugar = golog.Sugar
	defer logger.Sync() // flushes buffer, if any

	var name = "xiaoq"
	logger.Info("logger info")
	sugar.Infof("sugar info %s", name)

	fsutil := gofs.OsAfero
	fsutil.Exists("")

	fs := gofs.OsFs
	fs.Name()

	for {
		var s string
		fmt.Scan(&s)
		if s == "c" {
			break
		}
	}

	//DoGopherLua()
	DoGoLua()
}
