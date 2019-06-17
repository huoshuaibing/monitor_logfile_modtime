package main
import (
	"os"
	"time"
	"log"
	"github.com/royeo/dingrobot"
)
func GetFileModTime(path string) int64 {
	f,err := os.Open(path)
	if err != nil{
		println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()
	fi,err := f.Stat()
	if err != nil{
		println("stat fileinfo err")
		return time.Now().Unix()
	}
	return fi.ModTime().Unix()
}
func send2dingtalk(){
	webhook := "https://oapi.dingtalk.com/robot/send?access_token=598f161d26735e2a09bff7c7abbe84a8e0243c"
	robot := dingrobot.NewRobot(webhook)
	content := "xxx服务已停止写日志"
	atMobiles := []string{"17600147211"}
	isAtAll := false
	err := robot.SendText(content,atMobiles,isAtAll)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	loopWorker()
}

func loopWorker() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			t1 := GetFileModTime("/Users/edz/Desktop/1.py")
			now := time.Now()
			timestamp1 := now.Unix()
			if timestamp1 - t1 > 1800 {
				send2dingtalk()
			}else{
				return
			}
		}
	}
}
