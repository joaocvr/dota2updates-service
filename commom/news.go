package commom

import (
	"fmt"
	"regexp"
	"strings"
)

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
	Images        []Image
}

type Image struct {
	Tag     string
	TagHTML string
	Source  string
}

func (a *AppNews) ConvertTaggedContent() {

	for i, _ := range a.News.NewsItems {

		item := &a.News.NewsItems[i]

		if item.Contents != "" {
			convertTaggedContentImages(item)

			convertTaggedContentURL(item)

			convertTaggedContentBold(item)

			convertTaggedContentHeader(item)

			convertTaggedContentList(item)
		}
	}
}

func convertTaggedContentImages(item *NewsItem) {
	rgxImg := regexp.MustCompile(`\[img\](.*?)\[/img\]`)
	tagsImg := rgxImg.FindAllString(item.Contents, -1)
	images := []Image{}
	for _, tag := range tagsImg {
		image := Image{}
		image.Tag = tag
		image.Source = strings.Replace(strings.Replace(tag, "[img]", "", 1), "[/img]", "", 1)
		image.TagHTML = fmt.Sprintf(`<img src="%s">`, image.Source)
		images = append(images, image)

		item.Contents = strings.ReplaceAll(item.Contents, tag, image.TagHTML)
	}
	item.Images = images
}

func convertTaggedContentURL(item *NewsItem) {
	rgxUrl := regexp.MustCompile(`\[url=(.*?)\](.*?)\[/url\]`)
	tagsUrl := rgxUrl.FindAllString(item.Contents, -1)
	for _, tag := range tagsUrl {
		rgxTitle := regexp.MustCompile(`\](.*?)\[`)
		title := rgxTitle.FindAllString(tag, -1)[0]
		title = strings.TrimSuffix(strings.TrimPrefix(title, "]"), "[")

		rgxHref := regexp.MustCompile(`=(.*?)\]`)
		href := rgxHref.FindAllString(tag, -1)[0]
		href = strings.TrimSuffix(strings.TrimPrefix(href, "="), "]")

		link := fmt.Sprintf(`<a href="%s">%s</a>`, href, title)
		item.Contents = strings.ReplaceAll(item.Contents, tag, link)
	}
}

func convertTaggedContentBold(item *NewsItem) {
	item.Contents = strings.ReplaceAll(strings.ReplaceAll(item.Contents, "[b]", "<strong>"), "[/b]", "</strong>")
}

func convertTaggedContentHeader(item *NewsItem) {
	item.Contents = strings.ReplaceAll(strings.ReplaceAll(item.Contents, "[h1]", "<h1>"), "[/h1]", "</h1>")
	item.Contents = strings.ReplaceAll(strings.ReplaceAll(item.Contents, "[h2]", "<h2>"), "[/h2]", "</h2>")
	item.Contents = strings.ReplaceAll(strings.ReplaceAll(item.Contents, "[h3]", "<h3>"), "[/h3]", "</h3>")
	item.Contents = strings.ReplaceAll(strings.ReplaceAll(item.Contents, "[h4]", "<h4>"), "[/h4]", "</h4>")
	item.Contents = strings.ReplaceAll(strings.ReplaceAll(item.Contents, "[h5]", "<h5>"), "[/h5]", "</h5>")
}

func convertTaggedContentList(item *NewsItem) {
	item.Contents = strings.ReplaceAll(strings.ReplaceAll(item.Contents, "[list]", "<ol>"), "[/list]", "</ol>")
	item.Contents = strings.ReplaceAll(item.Contents, "[*]", "<br/>- ")
}
