package main

/***
cmd类
*/
import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool     //显示帮助信息
	versionFlag bool     //查看版本号
	cpOption    string   //
	XjreOption  string   //指定目录位置
	class       string   //类名
	args        []string //主程序参数

}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre") //配置指向jre目录的参数
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args)
}
