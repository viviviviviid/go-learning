package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/m/viviviviviid/learngo/ch5_web_server_with_echo/scrapper"
)

// echo 이용해서 서버돌리기 // echo docs가면 다양하게 볼게있음

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

const fileName string = "jobs.csv"

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName) // 다끝나고 서버에 저장된 파일 삭제
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
