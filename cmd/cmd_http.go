package cmd

import (
	"flag"
	"log"
	"os"
	. "satellite/global"
	"satellite/nets"
)

var httpCmd = flag.NewFlagSet(CmdHttp, flag.ExitOnError)
var httpIp string
var httpPort string

func init() {
	httpCmd.StringVar(&httpIp, "ip", "127.0.0.1", "ip address: ipv4 address witch http server listen, such as \"127.0.0.1\"")
	httpCmd.StringVar(&httpPort, "port", "14514", "port: port number witch http server listen, such as \"14514\"")
}

func ParseCmdHttp() {
	// check args number
	if len(os.Args) == 2 {
		httpCmd.Usage()
		os.Exit(1)
	}
	// parse command http
	err := httpCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Http Command.")
		os.Exit(1)
	}
	// handle command parameters
	handleCmdHttp(httpIp, httpPort)
}

func handleCmdHttp(ip string, port string) {
	nets.StartHttpServer(ip, port)
}
