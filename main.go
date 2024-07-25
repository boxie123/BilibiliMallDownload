package main

import (
	"log"
	"net/http"

	"github.com/boxie123/BilibiliMallDownload/utils"
	login "github.com/boxie123/GoBilibiliLogin"
)

func main() {
	cookie, _, _ := login.Login()

	client := &http.Client{}
	buyWorkVOList := utils.GetWorksList(client, cookie)
	for _, buyWorkVo := range buyWorkVOList {
		fileUrlMap, err := utils.GetFileUrl(client, cookie, buyWorkVo)
		if err != nil {
			log.Printf("获取文件url失败：%s\n%v", buyWorkVo.NickName, err)
		}
		utils.DownloadFiles(fileUrlMap)
	}
}
