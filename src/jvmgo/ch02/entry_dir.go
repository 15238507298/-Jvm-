package main

import (
	"io/ioutil"
	_ "io/ioutil"
	"path/filepath"
)
import _ "path/filepath"

//创建目录结构体
type DirEntry struct {
	absDir string
}

//利用文件相对路径获取文件绝对路径
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path) //将path转换为绝对路径
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

//读取文件二进制数据
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	filename := filepath.Join(self.absDir, className) //将self.absDir与className拼接起来
	data, err := ioutil.ReadFile(filename)            //将拼接起来的文件路径读取文件二进制数据，返回到data中
	return data, self, err
}

//返回文件目录
func (self *DirEntry) String() string {
	return self.absDir
}
