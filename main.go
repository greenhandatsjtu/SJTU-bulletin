package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

type Notice struct {
	Type  string `json:"_type"`
	Date  string `json:"date"`
	Href  string `json:"href"`
	Title string `json:"title"`
}

func main() {
	db, err := gorm.Open("sqlite3", "crawler/crawler/bulletin.db")
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&Notice{})

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Static("/", ""+"frontend/dist")
	e.GET("/notices", func(c echo.Context) error {
		var notices []Notice
		if err := db.Order("date desc").Find(&notices).Error; err != nil {
			log.Println(err)
			return c.JSON(http.StatusNotFound, nil)
		}
		return c.JSON(http.StatusOK, notices)
	})
	_ = e.Start(":8080")
}
