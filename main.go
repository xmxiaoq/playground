package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/GianlucaGuarini/go-observable"
	golua "github.com/Shopify/go-lua"
	"github.com/asaskevich/EventBus"
	"github.com/chuckpreslar/emission"
	"github.com/davecgh/go-spew/spew"
	"github.com/k0kubun/pp"
	"github.com/kr/pretty"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/shurcooL/go-goon"
	"github.com/spf13/cast"
	"github.com/tinylib/msgp/msgp"
	"github.com/xmxiaoq/playground/msg"
	"github.com/xtaci/kcp-go"
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

func UdpClient() {
	<-serverListened
	sugar.Infow("client start")

	ip := net.ParseIP("127.0.0.1")
	//srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 8888}
	conn, err := net.DialUDP("udp", nil, dstAddr)
	if err != nil {
		sugar.Errorw("DialUDP fail", "err", err)
		return
	}
	defer conn.Close()

	numWrite, err := conn.Write([]byte("aaa"))
	if err != nil {
		sugar.Errorw("client write error", "err", err, "numWrite", numWrite)
		return
	}
	sugar.Infow("client write", "buf", "aaa", "numWrite", numWrite)
}

func UdpServer() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		print(err)
		return
	}

	go func() {
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			sugar.Errorw("can not connect", "err", err)
			return
		}
		sugar.Infow("listen on", "addr", addr)
		defer conn.Close()

		close(serverListened)

		for {
			buf := make([]byte, 0, 1024)
			numRead, remote, err := conn.ReadFromUDP(buf)
			if err != nil {
				sugar.Errorw("can not connect", "err", err)
				continue
			}
			if numRead <= 0 {
				continue
			}
			sugar.Infow("server read", "numRead", numRead, "remote", remote, "buf", buf[0:numRead])

			numWrite, err := conn.WriteToUDP(buf[0:numRead], remote)
			sugar.Infow("server write", "numWrite", numWrite, "err", err)
		}
	}()
}

var serverListened chan bool = make(chan bool)

func KcpServer() {
	lis, err := kcp.Listen("192.168.0.149:10000")
	if err != nil {
		sugar.Errorw("kcp listen error", "err", err)
		return
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			sugar.Errorw("kcp accept error", "err", err)
			return
		}

		sugar.Infow("new connection")

		go func() {
			defer conn.Close()

			//_, err = conn.Write([]byte("中文"))
			//if err != nil {
			//	sugar.Errorw("write error", "err", err)
			//}

			membuf := bytes.NewBuffer(nil)
			enc := ffjson.NewEncoder(membuf)
			msgpR := msgp.NewReader(conn)
			//bufSend := bytes.NewBuffer(nil)
			for {
				//header := make([]byte, 2)
				//n, err := io.ReadFull(conn, header)
				//if err != nil {
				//	sugar.Errorw("kcp read error", "err", err)
				//	continue
				//}
				//
				//if n <= 0 {
				//	//_, err := conn.Write([]byte("from server"))
				//	//if err != nil {
				//	//	sugar.Errorw("write error", "err", err)
				//	//}
				//	continue
				//}
				//
				//dataLen := binary.BigEndian.Uint16(header)
				//buf := make([]byte, dataLen)
				//
				//_, err = io.ReadFull(conn, buf)
				//if err != nil {
				//	sugar.Errorw("msg length error", "err", err)
				//	continue
				//}
				//
				//samp := &Sample3List{}
				//_, err = samp.UnmarshalMsg(buf)
				//if err != nil {
				//	sugar.Errorw("decode error", "err", err)
				//}
				//continue

				//samp := &Sample3List{}
				samp := &msg.Sample3{}
				err = samp.DecodeMsg(msgpR)
				if err != nil {
					sugar.Errorw("decode error", "err", err)
				}
				sugar.Infow("DecodeMsg ok", "samp", samp)

				membuf.Reset()
				err = enc.Encode(samp)
				if err != nil {
					sugar.Errorw("decode error", "err", err)
				}
				sugar.Infow("DecodeMsg ok", "sampJson", string(membuf.Bytes()))
				continue

				//str, err := charsetutil.Decode(buf[2:n], "utf8")
				//if err != nil {
				//	sugar.Errorw("decode error", "err", err)
				//	continue
				//}
				//sugar.Infow("read ok", "n", n, "dataLen", dataLen, "buf", str)
				//continue

				//str, err = Decode(buf[0:n], "utf8")
				//if err != nil {
				//	sugar.Errorw("decode error", "err", err)
				//	continue
				//}
				//
				//sugar.Infow("read2 ok", "n", n, "buf", str)

				cast.ToString(100)

				//msg := &proto3_proto.Message{Name: "数据"}
				//msgBuf, err := msg.Marshal()
				//if err == nil {
				//	bufSend.Reset()
				//	nn := int16(n)
				//	binary.Write(bufSend, binary.BigEndian, nn)
				//	binary.Write(bufSend, binary.BigEndian, msgBuf)
				//	m, err := conn.Write(bufSend.Bytes())
				//	if err != nil {
				//		sugar.Errorw("write error", "err", err)
				//	} else {
				//		sugar.Infow("write ok", "m", m)
				//	}
				//
				//}
			}
		}()
	}
}

func Decode(buf []byte, enc string) (string, error) {
	r, err := charset.NewReader(enc, bytes.NewReader(buf))
	if err != nil {
		return "", err
	}
	result, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func main() {
	//cfg := zap.NewDevelopmentConfig()
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = append(cfg.OutputPaths, "playground.log")
	logger, _ = cfg.Build()
	sugar = logger.Sugar()
	defer logger.Sync() // flushes buffer, if any

	//fmt.Println(101%10, 101/10)
	//TestClosure()
	sugar.Infow("", "中文byte", []byte("中文"), "中文string", "中文")

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
	//TestEventBus()

	str, err := Decode([]byte("\xa35 for Pepp\xe9"), "latin1")
	if err != nil {
		sugar.Errorw("Decode error", "err", err)
		return
	}
	sugar.Infow(`Decode "ok`, "str", str)

	//for {
	//	var s string
	//	fmt.Scan(&s)
	//	if s == "c" {
	//		break
	//	}
	//}

	//obj := &msg.Sample3{}
	//obj.Age = 10
	//obj.FirstName = "中文"
	//obj.LastName = "字符"
	//buf, err := obj.MarshalMsg(nil)
	//if err != nil {
	//	sugar.Errorw("MarshalMsg error", "err", err)
	//	return
	//}
	//sugar.Infow("MarshalMsg ok", "buf", buf)
	//
	//obj2 := &msg.Sample3{}
	//_, err = obj2.UnmarshalMsg(buf)
	//if err != nil {
	//	sugar.Errorw("UnmarshalMsg error", "err", err)
	//	return
	//}
	//sugar.Infow("UnmarshalMsg ok", "buf", buf)

	//DoGopherLua()
	//DoGoLua()
	//Tcp()
	//UdpServer()
	//UdpClient()
	KcpServer()
}
