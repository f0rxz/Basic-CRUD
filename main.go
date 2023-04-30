package main

import (
	"backend/storage"
	"backend/storage/repos"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	router.GET("/quests", func(c *gin.Context) {
		fullParam := c.Query("full")
		//экземпляры моделей
		var (
			//comm     repos.Comment
			card     repos.Card
			cardFull repos.CardFull
		)

		var store = storage.NewStorage()
		card.CreateCard(store, "gfdgdf", "fdsfd", 5, "8:00", 5, "fdsfsd", []byte{}, "genre")
		if fullParam == "true" {
			//тут полная карточка выводится должна
			arr, err := cardFull.GetAllCardsFull(store)
			if err != nil {
				panic(err)
			}
			myStr, err := json.Marshal(arr)
			if err != nil {
				panic(err)
			}
			c.String(http.StatusOK, string(myStr))
		} else {
			//тут короткая
			arr, err := card.GetAllCards(store)
			if err != nil {
				panic(err)
			}
			myStr, err := json.Marshal(arr)
			if err != nil {
				panic(err)
			}
			c.String(http.StatusOK, string(myStr))
		}
	})
	router.Run(":8080")
}
