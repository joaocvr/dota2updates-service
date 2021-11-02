package endpoint

import (
	"dota2updates-service/commom"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) GetNews(c *gin.Context) {
	count := "10"
	if c.Query("count") != "" {
		count = c.Query("count")
	}

	response, err := http.Get(fmt.Sprintf("http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002/?appid=570&count=%s&format=json", count))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "no such host"})
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad response body"})
		return
	}

	appNews := commom.AppNews{}
	err = json.Unmarshal(responseData, &appNews)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "can't json unmarshal"})
		return
	}

	appNews.ConvertTaggedContent()

	c.JSON(http.StatusOK, appNews)
}
