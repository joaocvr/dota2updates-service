package commom

type AppNews struct {
	News News `json:"appnews"`
}

type News struct {
	AppID     int64      `json:"appid"`
	NewsItems []NewsItem `json:"newsitems"`
	Count     int64      `json:"count"`
}

type NewsItem struct {
	GID           string `json:"gid"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	IsExternalURL bool   `json:"is_external_url"`
	Author        string `json:"author"`
	Contents      string `json:"contents"`
	FeedLabel     string `json:"feedlabel"`
	Date          int64  `json:"date"`
	FeedName      string `json:"feedname"`
}
