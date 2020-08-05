package main

import (
	_ "github.com/icattlecoder/godaemon"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"path"
	"strconv"
)

type Notice struct {
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

type User struct {
	gorm.Model
	IP    string `gorm:"unique;not null"`
	Count uint   // visit count
}

var (
	db    *gorm.DB
	err   error
	visit Visit
)

// middleware to record visitor and request
func recordVisitorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//log.Println(c.Request().RequestURI)
		if c.Request().RequestURI == "/" {
			visit.Visitor += 1
			var user User
			if err = db.Where(User{IP: c.RealIP()}).First(&user).Error; err != nil {
				db.Create(&User{IP: c.RealIP(), Count: 1})
			} else {
				user.Count += 1
				db.Save(&user)
			}
		}
		visit.Request += 1
		db.Save(&visit)
		return next(c)
	}
}

func main() {
	db, err = gorm.Open("sqlite3", "crawler/crawler/bulletin.db")
	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&Visit{}, &User{})
	visit = Visit{}
	db.FirstOrCreate(&visit) // init database

	e := echo.New()
	e.Use(middleware.Recover())
	//e.Use(middleware.Logger())
	e.Use(recordVisitorMiddleware)

	e.Static("/", path.Join("frontend", "dist"))
	e.GET("/notices", fetchNotices)
	e.GET("/visit", GetVisitData)
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

func GetVisitData(c echo.Context) error {
	var visit Visit
	db.Find(&visit)
	return c.JSON(http.StatusOK, visit)
}
