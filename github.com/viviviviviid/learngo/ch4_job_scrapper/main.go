package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

func main() {
	getPages()
}

func getPages() int {
	// goQuery 내용을 살펴보면, get을 사용할때 에러처리를 해줘야하고, 도큐먼트 만들때도 에러처리해줘야함
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // 이렇게 닫아주어야, 메모리가 유실되는것을 막을 수 있음
	// defer은 함수가 끝난 뒤 실행되는 내용. ".on" 으로 생각하는게 편할듯

	// 도큐먼트 만들때도 에러처리 (goQuery 지침)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".page page_move track_event").Each()

	return 0
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
