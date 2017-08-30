package main

import (
	"testing"

	"fmt"

	golua "github.com/Shopify/go-lua"
	"github.com/yuin/gopher-lua"
)

//func TestCandyjs(t *testing.T) {
//	script := os.Args[1]
//	fmt.Printf("Executing %q\n", script)
//	ctx := candyjs.NewContext()
//	ctx.PevalFile(script)
//	fmt.Println("exit")
//}

func TestGopherLua(t *testing.T) {
	//script := os.Args[1]
	//fmt.Printf("Executing %q\n", script)
	L := lua.NewState()
	defer L.Close()
	//if err := L.DoString(`print("hello" .. (1 + 2.5))`); err != nil {
	//	panic(err)
	//}

	var err error
	err = L.DoFile(`tests.lua`)
	//err = L.DoFile(`main.lua`)
	if err != nil {
		fmt.Printf("%v", err)
		//fmt.Errorf("%v", err)
	}

	//if err := L.CallByParam(lua.P{
	//	Fn:      L.GetGlobal("main"),
	//	NRet:    1,
	//	Protect: true,
	//}); err != nil {
	//	panic(err)
	//}
	//
	//ret := L.Get(-1) // returned value
	//fmt.Println(ret)
	//L.Pop(1) // remove received value

	//fmt.Println("GopherLua exit")
}

func TestGoLua(t *testing.T) {
	_ = golua.NewState()
	//fmt.Println("golua start")
	//l := golua.NewState()
	//golua.OpenLibraries(l)
	//if err := golua.DoFile(l, "tests.lua"); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("golua exit")
}
