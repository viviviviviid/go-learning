package main

import (
	"fmt"
	"strings"

	"github.com/labstack/echo"
	"github.com/m/viviviviviid/learngo/ch5_web_server_with_echo/scrapper"
)

// 서버 관련

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	fmt.Println(c.FormValue("term"))
	return nil
}

func main() {
	// scrapper.Scrape("python")
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
