package main

import (
	"fmt"
	"jvmgo/ch03/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

/**
启动虚拟机
*/
func startJvm(cmd *Cmd) {
	//fmt.Printf("classpath:%s class:%s args:%v\n",cmd.cpOption,cmd.class,cmd.args)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
	}
	fmt.Printf("class data:%v\n", classData)
}
