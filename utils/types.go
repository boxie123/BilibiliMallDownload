package utils

// bilibili api 普遍返回数据格式
type ApiResponseCommon struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 配置信息
type ConfigInfo struct {
	AccessKey string `json:"accessKey"`
	Cookie    string `json:"cookie"`
}

type ShopWorksData struct {
	Works      Works         `json:"works"`
	ItemsFeeds []interface{} `json:"itemsFeeds"`
	SubTabList []SubTab      `json:"subTabList"`
}

type Works struct {
	BuyWorkVOList []BuyWorkVO `json:"buyWorkVOList"`
	Total         int         `json:"total"`
}

type BuyWorkVO struct {
	OrderId             int64    `json:"orderId"`
	BizOrderType        int      `json:"bizOrderType"`
	ShopMid             int      `json:"shopMid"`
	NickName            string   `json:"nickName"`
	Face                string   `json:"face"`
	ShopId              int      `json:"shopId"`
	Status              int      `json:"status"`
	WorkImgList         []string `json:"workImgList"`
	ImgTotal            int      `json:"imgTotal"`
	DownloadEndTime     int64    `json:"downloadEndTime"`
	FileList            []File   `json:"fileList"`
	IsPermanentDownload bool     `json:"isPermanentDownload"`
}

type File struct {
	FileName       string      `json:"fileName"`
	FileUrl        string      `json:"fileUrl"`
	FileSize       string      `json:"fileSize"`
	RawFileSize    string      `json:"rawFileSize"`
	FileType       string      `json:"fileType"`
	BucketType     int         `json:"bucketType"`
	SourceType     string      `json:"sourceType"`
	DownloadStatus int         `json:"downloadStatus"`
	Duration       int         `json:"duration"`
	AttachmentId   interface{} `json:"attachmentId"`
	Covers         []string    `json:"covers"`
	Description    string      `json:"description"`
	PreSignedUrl   interface{} `json:"preSignedUrl"`
}

type SubTab struct {
	SubTab     int    `json:"subTab"`
	SubTabName string `json:"subTabName"`
}

type DownloadInfo struct {
	URL      string
	PkgName  string
	FileName string
}
