package main

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"os/exec"
	"strings"
)

const DefaultPotPlayerPath = `C:\Program Files\DAUM\PotPlayer\PotPlayerMini64.exe`

func main() {
	binPath := os.Args[0]
	url := os.Args[1]
	name := os.Args[7]
	customPath, _ := ioutil.ReadFile( filepath.Join(filepath.Dir(binPath), "custom-path.txt"))
	potPlayerPath := string(customPath)
	if len(potPlayerPath) == 0 {
		potPlayerPath = DefaultPotPlayerPath
	}
	var cmd *exec.Cmd
	if strings.HasPrefix(url, "http") {
		cmd = exec.Command(potPlayerPath,  []string{url + "\\" + name}...)
	} else {
		cmd = exec.Command(potPlayerPath,  []string{url}...)
	}
	cmd.Start();
	cmd.Process.Release();
}
