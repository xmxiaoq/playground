package main

import (
	"fmt"

	golua "github.com/Shopify/go-lua"
	"github.com/xmxiaoq/gofs"
	"github.com/xmxiaoq/golog"
	"github.com/yuin/gopher-lua"
	"go.uber.org/zap"
	"net"
	"os"
	"strconv"
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
