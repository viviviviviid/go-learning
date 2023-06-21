package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	sector   string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=blockchain&exp_min=1&exp_max=1&recruitSort=relation&recruitPageCount=3"

func main() {
	var jobs []extractedJob
	totlaPages := getPages()

	for i := 0; i < totlaPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...) // 하나의 배열로 합침 // ... 이 없으면 배열안에 배열이 추가되는 형태
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func writeJobs(jobs []extractedJob) { // csv 파일 저장 관련 함수 // go standard method 페이지에서 찾을 수 있음
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file) // writer 생성
	defer w.Flush()          // 함수가 끝나는 시점, writer에 데이터 입력

	headers := []string{"ID", "Title", "Location", "Sector"}

	wErr := w.Write(headers) // Write 메소드는 에러를 반환
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=blockchain&exp_min=1&exp_max=1&recruitSort=relation&recruitPageCount=50" + job.id, job.title, job.location, job.sector}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
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
		job := extractJob(card)  // struct 내용을 여기에 저장
		jobs = append(jobs, job) // 추출될떄마다 내용을 jobs에 업데이트
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("value")
	title := card.Find(".area_job>.job_tit>a").Text()
	location := card.Find(".area_job>.job_condition>span>a").Text()
	sector := card.Find(".area_job>.job_sector>a").Text()
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		sector:   sector,
	}
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

	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ") // 빈공간을 제외한 문자열만 반환 제거한 내용을 join으로 합침
}
