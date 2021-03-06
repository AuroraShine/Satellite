package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"satellite/cmd"
	. "satellite/global"
	_ "satellite/logging"
)

func init() {
	// start multi-cpu
	core := runtime.NumCPU()
	runtime.GOMAXPROCS(core)
	// start debug pprof
	go func() {
		_ = http.ListenAndServe(":10514", nil)
	}()
}

func main() {
	// check command args number
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	// switch command execute
	switch os.Args[1] {
	case CmdHelp:
		flag.Usage()
	case CmdPacket:
		cmd.ParseCmdPack()
	case CmdUnpack:
		cmd.ParseCmdUnpack()
	case CmdCompress:
		cmd.ParseCmdComp()
	case CmdDecompress:
		cmd.ParseCmdDeComp()
	case CmdTcp:
		cmd.ParseCmdTcp()
	case CmdUdp:
		cmd.ParseCmdUdp()
	case CmdHttp:
		cmd.ParseCmdHttp()
	case CmdHttps:
		cmd.ParseCmdHttps()
	case CmdFtp:
		cmd.ParseCmdFtp()
	case CmdRpc:
		cmd.ParseCmdRpc()
	case CmdQRCode:
		cmd.ParseCmdQRCode()
	case CmdShell:
		cmd.ParseCmdShell()
	case CmdParses:
		cmd.ParseCmdParses()
	default:
		fmt.Println("Unrecognized command~")
		os.Exit(1)
	}
}
