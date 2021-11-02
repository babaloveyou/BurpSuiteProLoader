package main

import (
	"flag"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {

	var jdk string
	flag.StringVar(&jdk, "jdk", "", "")
	flag.Parse()
	var burpsuite_pro_jar string = ""
	filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		name := strings.ToLower(info.Name())
		if strings.HasPrefix(name, "burpsuite_pro") {
			burpsuite_pro_jar = info.Name()
			return nil
		}
		return nil
	})
	if burpsuite_pro_jar == "" {
		os.WriteFile("error.log", []byte("burpsuite_pro.jar not found \n"), 0666)
		return

	}

	if jdk == "8" {
		jdk8path := os.Getenv("jdk8")

		cmd := exec.Command(jdk8path+"\\java.exe", "-Xbootclasspath/p:burploader.jar", "-jar", burpsuite_pro_jar)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Start()

	} else {

		cmd := exec.Command("java", "-Dfile.encoding=utf-8", "-javaagent:burploader.jar", "-noverify", "-javaagent:BurpSuiteChs.jar", "-jar", burpsuite_pro_jar)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Start()

	}

}
