package main

import (
	_ "github.com/icattlecoder/godaemon"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
)

type Notice struct {
	ID    uint   `json:"id"`
	Type  string `json:"_type"`
	Date  string `json:"date"`
	Href  string `json:"href"`
	Title string `json:"title"`
}

type Visit struct {
	ID      uint `gorm:"primary_key" json:"-"`
	Visitor uint `json:"visitor"` // visit count
	Request uint `json:"request"` // request count
}

var (
	db    *gorm.DB
	err   error
	visit Visit
	mu    sync.Mutex
)

// middleware to record visitor and request
func recordVisitorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//log.Println(c.Request().RequestURI)
		mu.Lock()
		if c.Request().RequestURI == "/visit" {
			visit.Visitor++
			db.Save(&visit)
		}
		visit.Request += 1
		db.Save(&visit)
		mu.Unlock()
		return next(c)
	}
}

func main() {
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}
	db, err = gorm.Open("sqlite3", "crawler/crawler/bulletin.db")
	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&Visit{})
	visit = Visit{}
	db.FirstOrCreate(&visit) // init database

	e := echo.New()
	e.Use(middleware.Recover())
	//e.Use(middleware.Logger())
	e.Use(recordVisitorMiddleware)

	e.Static("/", path.Join("frontend", "dist"))
	e.GET("/notices", fetchNotices)
	e.GET("/visit", getVisitData)
	e.Logger.Fatal(e.Start(":" + port))
}

func fetchNotices(c echo.Context) error {
	page := c.QueryParam("page")
	var pageInt int
	if pageInt, err = strconv.Atoi(page); err != nil {
		pageInt = 0
	}
	var notices []Notice
	//enable paging
	if err := db.Order("date desc, id desc").Offset(pageInt * 25).Limit(25).Find(&notices).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, notices)
}

func getVisitData(c echo.Context) error {
	return c.JSON(http.StatusOK, visit)
}
