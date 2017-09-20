package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	fmt.Println(os.Getenv("HOME"))
	checkError(err)
	defer src.Close()
	dst, err := os.Create(dstName)
	checkError(err)
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	org := os.Args[1]
	srcFile, _ := filepath.Abs(org)
	binFilePath = os.Getenv("HOME") + "/.mybin/" + org
	if _, err := CopyFile(os.Getenv("HOME")+"/.mybin/"+org, srcFile); err == nil {
		cmd, _ := exec.Command("chmod", "777", binFilePath)
		fmt.Println(os.Args[1], "  命令载入成功")
	}
}
