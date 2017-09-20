package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	//	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type Watch struct {
	watch *fsnotify.Watcher
}


type Conf struct {
	User   string   `json:user`
	Pwd    string   `json:pwd`
	Ignore interface{} `json:ignore`
}

const (
	confFilePath = "./conf.json"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getConf(path string) (c Conf){
	data, err := ioutil.ReadFile("/Users/zhangyunfeng/code/learn/go/test/conf.json")
	checkErr(err)
	//解析配置文件
	err = json.Unmarshal(data, &c)
	checkErr(err)
	return
}

var localPath = os.Args[1]
var localFilePath = "./" + localPath
var serverFilePath = os.Args[2]
var notWatchList = []string{"static", "web", ".git", "node_modules", ".idea"}

func isIn(str string, array []string) bool {
	for _, v := range array {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

//监控目录
func (w *Watch) watchDir(dir string) {
	//通过Walk来遍历目录下的所有子目录
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//这里判断是否为目录，只需监控目录即可
		//目录下的文件也在监控范围内，不需要一个一个加
		if info.IsDir() {
			if isIn(path, notWatchList) {
				// fmt.Println("ignore flodar ------------> ", path)
				return nil
			}
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = w.watch.Add(path)
			if err != nil {
				return err
			}
			fmt.Println("监控 : ", path)
		}
		return nil
	})
	go func() {
		for {
			select {
			case ev := <-w.watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						fmt.Println("创建文件 : ", ev.Name)
						//这里获取新创建文件的信息，如果是目录，则加入监控中
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Add(ev.Name)
							fmt.Println("添加监控 : ", ev.Name)
							scpUpload(ev.Name, localPath)
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						fmt.Println("写入文件 : ", ev.Name)
						scpUpload(ev.Name, localPath)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						fmt.Println("删除文件 : ", ev.Name)
						//如果删除文件是目录，则移除监控
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.watch.Remove(ev.Name)
							fmt.Println("删除监控 : ", ev.Name)
							scpUpload(ev.Name, localPath)
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						fmt.Println("重命名文件 : ", ev.Name)
						//如果重命名文件是目录，则移除监控
						//注意这里无法使用os.Stat来判断是否是目录了
						//因为重命名后，go已经无法找到原文件来获取信息了
						//所以这里就简单粗爆的直接remove好了
						w.watch.Remove(ev.Name)
						scpUpload(ev.Name, localPath)
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						fmt.Println("修改权限 : ", ev.Name)
						scpUpload(ev.Name, localPath)
					}
				}
			case err := <-w.watch.Errors:
				{
					fmt.Println("error : ", err)
					return
				}
			}
		}
	}()
}
func scpUpload(evName string, localPath string) {
	keyPath := strings.Split(evName, localPath)
	fmt.Println("scp " + evName + " yfzhang@relay01.hillinsight.com:" + serverFilePath + keyPath[len(keyPath)-1])
	cmd := exec.Command("scp", evName, "yfzhang@relay01.hillinsight.com:"+serverFilePath+keyPath[len(keyPath)-1])
	cmd.Run()
}

func main() {
	conf := getConf(confFilePath)
	fmt.Println(conf)
	//watch, _ := fsnotify.NewWatcher()
	//w := Watch{
	//	watch: watch,
	//}
	//w.watchDir(localFilePath)
	//select {}
}
