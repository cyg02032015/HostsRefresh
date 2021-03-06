package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var RAW_URLS = []string{
	"github.com",
	"api.github.com",
	"github.map.fastly.net",
	"github.githubassets.com",
	"github.global.ssl.fastly.net",
	"raw.githubusercontent.com",
	"camo.githubusercontent.com",
	"avatars.githubusercontent.com",
	"avatars0.githubusercontent.com",
	"avatars1.githubusercontent.com",
	"avatars2.githubusercontent.com",
	"avatars3.githubusercontent.com",
	"avatars4.githubusercontent.com",
	"avatars5.githubusercontent.com",
	"avatars6.githubusercontent.com",
	"avatars7.githubusercontent.com",
	"avatars8.githubusercontent.com",
	"favicons.githubusercontent.com",
	"developer.apple.com",
	"devstreaming-cdn.apple.com",
}

const IPADDRESS_PREFIX = ".ipaddress.com"

func Hosts() {
	var content string = "# Host Start\n"
	for _, rawUrl := range RAW_URLS {
		ip := getIp(rawUrl)
		content += ip + " " + rawUrl + "\n"
		// content += rawUrl + " " + ip + "\n"
	}
	content += "# Host End"

	writeToFile(content)
}

func getIp(rawUrl string) string {
	url := makeIpaddressUrl(rawUrl)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("error url get: ", err)
		return ""
	}
	defer resp.Body.Close()

	re, err := regexp.Compile("\\b([0-9]{1,3}\\.){3}[0-9]{1,3}")
	if err != nil {
		log.Println("regx err: ", err.Error())
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read all : ", err.Error())
	}
	ip := re.FindString(string(bytes))
	return ip

}

func makeIpaddressUrl(rawUrl string) (ipaddress string) {
	dotCount := strings.Count(rawUrl, ".")
	if dotCount > 1 {
		rawUrlList := strings.Split(rawUrl, ".")
		tempUrl := rawUrlList[dotCount-1] + "." + rawUrlList[dotCount]
		ipaddress = "https://" + tempUrl + IPADDRESS_PREFIX + "/" + rawUrl
	} else {
		ipaddress = "https://" + rawUrl + IPADDRESS_PREFIX
	}
	return
}

func writeToFile(content string) {
	file, err := os.OpenFile("hosts", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Println("open file err:", err)
	}
	defer file.Close()

	oldContentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("read host err:", err)
	}
	oldContent := string(oldContentBytes)
	if strings.Compare(content, oldContent) == 0 {
		log.Println("未发生更改")
		return
	}
	file.Truncate(0)
	file.Seek(0, 0)
	file.WriteString(content)

	bytes, err := ioutil.ReadFile("README_template.md")
	if err != nil {
		log.Println("open README_template.md err, ", err)
	}
	readMeContent := strings.Replace(string(bytes), "{{hosts}}", content, -1)
	// 当放在github workflows中执行的时候使用的是其他地区的服务器，Local并不代表本地时间
	time.Local = time.FixedZone("CST", 8*60*60)
	time := time.Now().Local().Format("2006-01-02 15:04:05")
	readMeContent = strings.ReplaceAll(readMeContent, "{{time}}", time)

	readMeFile, err := os.OpenFile("README.md", os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Println("open README.md err, ", err)
	}
	defer readMeFile.Close()
	readMeFile.WriteString(readMeContent)
}

func main() {
	//timeTicker := time.NewTicker(time.Hour * 2)
	//for {
	//	Hosts()
	//	<-timeTicker.C
	//}
	Hosts()
}
