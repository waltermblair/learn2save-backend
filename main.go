package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/waltermblair/learn2save-backend/boundaries"
	"github.com/waltermblair/learn2save-backend/common"
	"net/http"
	"github.com/gin-contrib/cors"
)

func main() {

	log := logrus.New()
	config, err := common.LoadConfig()
	if err != nil {
		log.Error("unable to load config: ", err)
	}
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.ReactURL}
	router.Use(cors.New(corsConfig))
	db, err := boundaries.NewDBClient(config)
	if err != nil {
		log.Error("unable to connect to db: ", err)
	}

	router.GET("/accounts/:id", func(c *gin.Context) {
		id := c.Param("id")
		account, err := db.GetAccountByID(id)
		if err != nil {
			log.Error("unable to fetch account: ", err)
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			log.WithFields(logrus.Fields{
				"id": id,
				"account": account,
			}).Info("fetched account")
			c.JSON(http.StatusOK, account)
		}
	})

	err = router.Run(); if err != nil {
		log.Fatal("fatal error: ", err)
	}

}
