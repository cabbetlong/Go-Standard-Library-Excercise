package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	tempDir, err := ioutil.TempDir("", "*") // 第一个参数为空串表示在操作系统temp目录下创建
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(tempDir)

	tempFile, err := ioutil.TempFile(tempDir, "test_*.txt") // 生成文件名为: test_xxxxx.txt 的临时文件
	if err != nil {
		fmt.Println(err)
		return
	}
	tempFile.Close()
	defer os.Remove(tempFile.Name())

	b := []byte("asdf qwer zxcv")
	ioutil.WriteFile(tempFile.Name(), b, os.ModeAppend)

	content, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("len: %d, file content: %s", len(content), content)
}
