package main

import (
	"fmt"

	"net"
	"os"
	"strconv"

	"encoding/json"
	"time"

	"github.com/GianlucaGuarini/go-observable"
	golua "github.com/Shopify/go-lua"
	"github.com/asaskevich/EventBus"
	"github.com/chuckpreslar/emission"
	"github.com/davecgh/go-spew/spew"
	"github.com/k0kubun/pp"
	"github.com/kr/pretty"
	"github.com/shurcooL/go-goon"
	"github.com/y0ssar1an/q"
	"github.com/yuin/gopher-lua"
	"go.uber.org/zap"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

type UserInfo struct {
	Id     int
	Name   string
	Items  [6]int
	Equips [2]int
	Cards  []int
}

func TestJsonArray() {

}

func TestJsonUnmarshal() {
	var info UserInfo

	buf := []byte(`{"id":1, "name":"xq", "users":[1,2,4,7], "Equips":[1,2,4,7]}`)
	err := json.Unmarshal(buf, &info)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info)
	}

	var infoPtr *UserInfo = &UserInfo{Id: 2, Name: "xq2"}
	err = json.Unmarshal(buf, &infoPtr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v \n", infoPtr)
	}
}

//func TestLumerjack() {
//	l := &lumberjack.Logger{Filename: `D:\GOPATH\src\github.com\xmxiaoq\playground\playground.log`}
//	logger := zap.New(
//		//zap.NewJSONEncoder(),
//		zap.NewTextEncoder(zap.TextTimeFormat(time.RFC3339)),
//		zap.DebugLevel,
//		//zap.DiscardOutput,
//		zap.Output(zap.Tee(zap.AddSync(l), zap.AddSync(os.Stdout))),
//	)
//
//	logger.Info("Failed to fetch URL.404",
//		zap.String(`"url"`, `this is: = (づ｡◕‿‿◕｡)づ"充值"`),
//		zap.Int("attempt", 10),
//		zap.Duration("backoff", time.Duration(100)),
//	)
//
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, syscall.SIGHUP)
//
//	go func() {
//		for {
//			<-c
//			l.Rotate()
//		}
//	}()
//}

func TestChanel() {
	ch := make(chan string)
	for i := 0; i < 1; i++ {
		k := i
		go func() {
			fmt.Println("func bengin: ", k)
			for m := range ch {
				fmt.Println("processed:", m)
			}
			fmt.Println("func end: ", k)
		}()

		ch <- fmt.Sprintf("cmd.%d", i)
	}

	//time.Sleep(1 * time.Second)
	fmt.Println("main exit")

}

func TestPrinter() {
	m := map[string]string{"foo": "bar", "hello": "world"}
	var infoPtr *UserInfo = &UserInfo{Id: 2, Name: "xq2", Cards: []int{6, 7, 8}}
	infoPtr.Cards = make([]int, 0, 10)
	infoPtr.Cards = append(infoPtr.Cards, 16, 17, 18)

	pp.Println(m)
	pp.Println(infoPtr)
	//var s string
	//s = pp.Sprint(infoPtr)
	//fmt.Println(s)

	pretty.Println(m)
	pretty.Println(infoPtr)

	spew.Dump(m)
	spew.Dump(infoPtr)

	q.Q(m)
	q.Q(infoPtr)

	type Inner struct {
		Field1 string
		Field2 int
	}
	type Lang struct {
		Name  string
		Year  int
		URL   string
		Inner *Inner
	}

	x := Lang{
		Name: "Go",
		Year: 2009,
		URL:  "http",
		Inner: &Inner{
			Field1: "Secret!",
		},
	}

	goon.Dump(x)
	goon.Dump(m)
	goon.Dump(infoPtr)

	//l.Warn(infoPtr)
}

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func TestClosure() {

	data2 := []*field{{"one1"}, {"two2"}, {"three3"}}
	//data := []field{{"one"}, {"two"}, {"three"}}
	//
	//for _, v := range data {
	//	go v.print()
	//}

	for _, v := range data2 {
		//v.print()
		//v := v
		go v.print()
		//go func() {
		//	fmt.Println(v.name)
		//}()
	}

	fmt.Println(`exit`)
	time.Sleep(3 * time.Second)
}

func TestSlice() {
	var s0 []int
	if s0 == nil {
		fmt.Println("s0 = null")
		s0 = append(s0, 10)
	}

	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) //prints 3 3 [1 2 3]

	type User struct {
		name string
		id   int
	}
	type Data struct {
		users []User
	}

	// 1 ,3, 4, 5, 6
	//
	data := Data{users: []User{{"1", 1}, {"2", 2}, {"3", 3}, {"4", 4}, {"5", 5}}}
	for key, val := range data.users {
		if key%2 == 1 {
			data.users = append(data.users[:key], data.users[key+1:]...)
		}
		fmt.Println(key, val.name)
	}

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [2 3]
}

func TestSliceDelete() {
	type User struct {
		name string
		id   int
	}
	type Data struct {
		users []User
	}

	// 1 ,3, 4, 5, 6
	//
	data := Data{users: []User{{"1", 1}, {"2", 2}, {"3", 3}, {"4", 4}, {"5", 5}, {"6", 6}}}
	n := len(data.users)
	for key, val := range data.users {
		if key%2 == 1 {
			//data.users = append(data.users[:key], data.users[key+1:]...)
			copy(data.users[key:], data.users[key+1:])
			data.users = data.users[:n-1]
		}
		fmt.Println(key, val.name)
	}
}

//func TestDeferWhenFatal() {
//	defer func() {
//		logg.Println("defer")
//	}()
//
//	log.Println("print")
//	log.Fatalln("fatal")
//}

func TestObservable() {
	o := observable.New()
	n := 0

	onFoo := func() {
		n++
	}

	o.One("foo", onFoo)

	o.Trigger("foo").Trigger("foo").Trigger("foo")

	if n != 1 {
		fmt.Errorf("The counter is %d instead of being %d", n, 1)
	}
}

func TestEmitterRemoveOnce() {
	event := "test"
	flag := false
	fn := func() { flag = !flag }

	emission.NewEmitter().
		Once(event, fn).
		RemoveListener(event, fn).
		Emit(event)

	if flag {
		fmt.Println("Failed to remove Listener for Once")
	}
}

func TestEventBus() {
	event := "test"
	flag := 0
	fn := func() { flag = flag + 1 }

	bus := EventBus.New()
	bus.SubscribeOnce(event, fn)
	bus.Subscribe(event, fn)
	bus.Subscribe(event, fn)

	bus.Publish(event)

	fmt.Printf("flag=%v \n", flag)
	//if flag {
	//	fmt.Println("Failed to remove Listener for Once")
	//}
}

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

func Tcp() {
	const (
		CONN_HOST = "localhost"
		CONN_PORT = "8888"
		CONN_TYPE = "tcp"
	)

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		fmt.Println("new connetion")
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte(strconv.Itoa(reqLen) + " Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}

func main() {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = append(cfg.OutputPaths, "playground.log")
	logger, _ = cfg.Build()
	sugar = logger.Sugar()
	defer logger.Sync() // flushes buffer, if any

	//fmt.Println(101%10, 101/10)
	//TestClosure()

	//TestJsonArray()
	//TestJsonUnmarshal()
	//TestChanel()
	//TestPrinter()
	//TestLumerjack()
	//TestSlice()
	//TestSliceDelete()
	//TestDeferWhenFatal()
	//TestObservable()
	//TestEmitterRemoveOnce()
	TestEventBus()

	//for {
	//	var s string
	//	fmt.Scan(&s)
	//	if s == "c" {
	//		break
	//	}
	//}

	//DoGopherLua()
	//DoGoLua()
	Tcp()
}
