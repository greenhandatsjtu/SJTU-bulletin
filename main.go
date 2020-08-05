package main

import (
	_ "github.com/icattlecoder/godaemon"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strconv"
)

type Notice struct {
	Type  string `json:"_type"`
	Date  string `json:"date"`
	Href  string `json:"href"`
	Title string `json:"title"`
}

var (
	db  *gorm.DB
	err error
)

func main() {
	db, err = gorm.Open("sqlite3", "crawler/crawler/bulletin.db")
	if err != nil {
		log.Println(err)
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Static("/", ""+"frontend/dist")
	e.GET("/notices", fetchNotices)
	_ = e.Start(":8080")
}

func fetchNotices(c echo.Context) error {
	page := c.QueryParam("page")
	var pageInt int
	if pageInt, err = strconv.Atoi(page); err != nil {
		pageInt = 0
	}
	var notices []Notice
	//enable paging
	if err := db.Order("date desc").Offset(pageInt * 25).Limit(25).Find(&notices).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, notices)
}
