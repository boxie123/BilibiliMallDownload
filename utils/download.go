package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const (
	ApiMyShopWorks     = "https://mall.bilibili.com/mall-up-c/my/shop/works"
	ApiDeliveryFileUrl = "https://mall.bilibili.com/mall-up-c/order/process/delivery/file/url"
)

func GetWorksList(client *http.Client, cookie string) []BuyWorkVO {
	paramsMap := map[string]interface{}{
		"subTab":           0,
		"pageSize":         10,
		"pageNum":          1,
		"bizOrderTypeList": []int{2, 3, 5},
	}

	data, err := PostApiResponseData(client, cookie, ApiMyShopWorks, paramsMap)
	if err != nil {
		panic(err)
	}
	buyWorkVOList, err := parseWorksList(data)
	if err != nil {
		panic(err)
	}

	return buyWorkVOList
}

func parseWorksList(data map[string]interface{}) ([]BuyWorkVO, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var shopWorksData ShopWorksData
	err = json.Unmarshal(jsonBytes, &shopWorksData)
	if err != nil {
		return nil, err
	}

	return shopWorksData.Works.BuyWorkVOList, nil
}

func GetFileUrl(client *http.Client, cookie string, buyWorkVO BuyWorkVO) ([]DownloadInfo, error) {
	allInfo := []DownloadInfo{}
	for _, file := range buyWorkVO.FileList {
		paramsMap := map[string]interface{}{
			"fileUrl":    file.FileUrl,
			"fileName":   file.FileName,
			"orderId":    buyWorkVO.OrderId,
			"bucketType": file.BucketType,
		}
		data, err := PostApiResponseData(client, cookie, ApiDeliveryFileUrl, paramsMap)
		if err != nil {
			return nil, err
		}

		fileUrl, ok := data["url"].(string)
		if !ok {
			panic("Error: fileUrl is not an string")
		}
		singleInfo := DownloadInfo{
			URL:      fileUrl,
			FileName: file.FileName,
			PkgName:  buyWorkVO.NickName + "_" + fmt.Sprintf("%d", buyWorkVO.OrderId),
		}
		allInfo = append(allInfo, singleInfo)
	}
	return allInfo, nil
}

func DownloadFiles(downloadInfoList []DownloadInfo) {
	var wg sync.WaitGroup
	for _, downloadInfo := range downloadInfoList {
		wg.Add(1)
		go func(downloadInfo DownloadInfo) {
			defer wg.Done()
			err := downloadFile(downloadInfo)
			if err != nil {
				log.Printf("%s 下载失败：%s\n", downloadInfo.FileName, downloadInfo.URL)
			} else {
				log.Printf("%s/%s 下载成功\n", downloadInfo.PkgName, downloadInfo.FileName)
			}
		}(downloadInfo)
	}
	wg.Wait()
}

func downloadFile(info DownloadInfo) error {
	resp, err := http.Get(info.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	dirPath := filepath.Join(".", "data", "mall", info.PkgName)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm)
	}

	filePath := filepath.Join(dirPath, info.FileName)
	err = os.WriteFile(filePath, body, 0644)
	if err != nil {
		return err
	}

	return nil
}
