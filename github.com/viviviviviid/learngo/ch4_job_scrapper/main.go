package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractJob struct {
	id       string
	title    string
	location string
	sector   string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=blockchain&exp_min=1&exp_max=1&recruitSort=relation&recruitPageCount=3"

func main() {
	totlaPages := getPages()

	for i := 0; i < totlaPages; i++ {
		getPage(i + 1)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page) // strconv.Itoa() : go에서 지원하는 string으로 바꾸는 함수
	fmt.Println("Requesting ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) { // 현재 찾은건 각각의 카드
		id, _ := card.Attr("value")
		fmt.Println(id)
		title := card.Find(".area_job>.job_tit>a").Text()
		fmt.Println(title)
		location := card.Find(".area_job>.job_condition>span>a").Text()
		fmt.Println(location)
		sector := card.Find(".area_job>.job_sector>a").Text()
		fmt.Println(sector)
	})
}

func getPages() int {
	pages := 0
	// goQuery 내용을 살펴보면, get을 사용할때 에러처리를 해줘야하고, 도큐먼트 만들때도 에러처리해줘야함
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // 이렇게 닫아주어야, 메모리가 유실되는것을 막을 수 있음
	// defer은 함수가 끝난 뒤 실행되는 내용. ".on" 으로 생각하는게 편할듯

	// 도큐먼트 만들때도 에러처리 (goQuery 지침)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// Each 위에 마우스 올려보면, 함수를 받는다는걸 확인가능. 복붙하기
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {

	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	// getPages함수의 첫 코드라인, Get위에 마우스 올려보면, response 값으로 *http.Response인 포인터값이 온다는 것을 확인할 수 있음
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status", res.StatusCode)
	}
}

func cleanString(str string) string {

	return str
}
