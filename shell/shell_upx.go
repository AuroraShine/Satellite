package shell

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"satellite/exec"
	"satellite/utils"
)

func shellUpx(src string, dest string) (err error) {
	var b bool
	// check source file exist...
	b, err = utils.PathExist(src)
	if err != nil {
		log.Println("Error src file path:", err)
		return err
	}
	if !b {
		log.Println("Error src file path:Not exist")
		return err
	}
	// according different system
	if runtime.GOOS == "windows" {
		// get executable dir
		exe, err := os.Executable()
		if err != nil {
			log.Println("Error get executable path:", err)
			return err
		}
		dir := filepath.Dir(exe)
		// get upx path
		cmd := dir + "\\tools\\upx.exe -9 -o " + dest + " " + src
		log.Println("Exec command:", cmd)
		// exec command
		r, err := exec.ExecCmd(cmd)
		if err != nil {
			log.Println("Error exec command:", err)
			return err
		}
		log.Println("Exec result:", r)
	} else {
		// exec command
		cmd := "../tools/upx -9 -o " + dest + " " + src
		log.Println("Exec command:", cmd)
		r, err := exec.ExecShell(cmd)
		if err != nil {
			log.Println("Error exec command:", err)
			return err
		}
		log.Println("Exec result:", r)
	}
	return err
}
