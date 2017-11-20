package main

import (
	"fmt"
	"net"
	"os"

	"strconv"

	"bytes"

	"io/ioutil"

	"encoding/binary"

	golua "github.com/Shopify/go-lua"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"github.com/spf13/cast"
	"github.com/xmxiaoq/gofs"
	"github.com/xmxiaoq/golog"
	"github.com/xmxiaoq/playground/pb/gogofaster_out"
	"github.com/xtaci/kcp-go"
	"github.com/yuin/charsetutil"
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

			bufSend := bytes.NewBuffer(nil)
			buf := make([]byte, 1024)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					sugar.Errorw("kcp read error", "err", err)
					continue
				}
				if n <= 0 {
					//_, err := conn.Write([]byte("from server"))
					//if err != nil {
					//	sugar.Errorw("write error", "err", err)
					//}
					continue
				}

				len := binary.BigEndian.Uint16(buf)
				samp := &Sample3{}
				_, err = samp.UnmarshalMsg(buf[2:len])
				if err != nil {
					sugar.Errorw("decode error", "err", err)
				}
				continue

				str, err := charsetutil.Decode(buf[2:n], "utf8")
				if err != nil {
					sugar.Errorw("decode error", "err", err)
					continue
				}
				sugar.Infow("read ok", "n", n, "len", len, "buf", str)
				continue

				//str, err = Decode(buf[0:n], "utf8")
				//if err != nil {
				//	sugar.Errorw("decode error", "err", err)
				//	continue
				//}
				//
				//sugar.Infow("read2 ok", "n", n, "buf", str)

				cast.ToString(100)

				msg := &proto3_proto.Message{Name: "数据"}
				msgBuf, err := msg.Marshal()
				if err == nil {
					bufSend.Reset()
					nn := int16(n)
					binary.Write(bufSend, binary.BigEndian, nn)
					binary.Write(bufSend, binary.BigEndian, msgBuf)
					m, err := conn.Write(bufSend.Bytes())
					if err != nil {
						sugar.Errorw("write error", "err", err)
					} else {
						sugar.Infow("write ok", "m", m)
					}

				}
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
	logger = golog.Logger
	sugar = golog.Sugar
	defer logger.Sync() // flushes buffer, if any

	var name = "xiaoq"
	logger.Info("logger info")
	sugar.Infof("sugar info %s", name)
	sugar.Infow("", "中文byte", []byte("中文"), "中文string", "中文")

	fsutil := gofs.OsAfero
	fsutil.Exists("")

	fs := gofs.OsFs
	fs.Name()

	str, err := Decode([]byte("\xa35 for Pepp\xe9"), "latin1")
	if err != nil {
		sugar.Errorw("", "err", err)
		return
	}

	sugar.Infow("decode ok", "str", str)

	//for {
	//	var s string
	//	fmt.Scan(&s)
	//	if s == "c" {
	//		break
	//	}
	//}

	//DoGopherLua()
	//DoGoLua()
	//Tcp()
	//UdpServer()
	//UdpClient()
	KcpServer()
}
