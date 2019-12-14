package cmd

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	. "satellite/global"
	"satellite/parses"
)

var parsesCmd = flag.NewFlagSet(CmdParses, flag.ExitOnError)
var parsesResolver string
var parsesSrc string
var parsesSection string
var parsesName string
var parsesType string

func init() {
	parsesCmd.StringVar(&parsesResolver, "r", "", "parses package resolver, you can choose one from [\"ini\",\"erl\"]")
	parsesCmd.StringVar(&parsesSrc, "i", "", "input file: correspond file need to be parse, such as \"test.ini\" or \"test.erl\"")
	parsesCmd.StringVar(&parsesSection, "s", "", "parameter section: correspond parameter section only for ini, such as \"INIT\"")
	parsesCmd.StringVar(&parsesName, "n", "", "parameter name: correspond parameter need to be parse, such as \"conf\" or \"pi\"")
	parsesCmd.StringVar(&parsesType, "t", "", "parameter type: correspond parameter type need to be parse, such as \"bool\" or \"int\"")
}

func ParseCmdParses() {
	// check args number
	if len(os.Args) == 2 {
		parsesCmd.Usage()
		os.Exit(1)
	}
	// parse command parses
	err := parsesCmd.Parse(os.Args[2:])
	if err != nil {
		log.Println("Error Parse Parses Command.")
		os.Exit(1)
	}
	// handle command parameters
	switch parsesResolver {
	case "INI", "ini":
		err = handleCmdParsesIni(parsesSrc, parsesSection, parsesName, parsesType)
	case "ERL", "erl":
	default:
		err = errors.New("unrecognized resolver")
	}
	if err != nil {
		fmt.Println("Parses failure:", err)
		os.Exit(1)
	}
	fmt.Println("Parses success.")
}

func handleCmdParsesIni(src string, section string, name string, kind string) (err error) {
	// handle command parses ini
	switch kind {
	case "string":
		var value string
		err = parses.GetValueFrom(src, section, name, &value)
		if err != nil {
			return err
		}
		fmt.Println("get value:", value)
	case "int":
		var value int
		err = parses.GetValueFrom(src, section, name, &value)
		if err != nil {
			return err
		}
		fmt.Println("get value:", value)
	case "float64":
		var value float64
		err = parses.GetValueFrom(src, section, name, &value)
		if err != nil {
			return err
		}
		fmt.Println("get value:", value)
	case "bool":
		var value bool
		err = parses.GetValueFrom(src, section, name, &value)
		if err != nil {
			return err
		}
		fmt.Println("get value:", value)
	default:
		err = errors.New("unrecognized type name")
	}
	return err
}

func handleCmdParsesErl(src string, name string, kind string) (err error) {
	// handle command parses erl
	return err
}
