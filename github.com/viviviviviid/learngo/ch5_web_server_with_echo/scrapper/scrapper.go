package scrapper

// FLOW
// getPages로 몇 페이지인지 확인
// 각 페이지별로 getPage가 실행. 여기 사이트에서는 9개
// 이전 방식으로는 getPage를 하나하나 처리해줬는데, goroutines을 이용해서 한번에 처리하려함
// getPage 내에 있는 extractJob 또한 하나의 일자리내용을 추출한 뒤 다음 내용을 진행함. -> 이것도 goroutines을 이용
//
// 50개의 extractJob가 동시진행되면, channel을 통해 getPage로 보내고, getPage가 전부 진행되면 channel을 통해 main으로 보냄
// 즉 extractJob <-> getPage, getPage <-> main으로 총 두개의 channel이 필요함
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
	company  string
	title    string
	location string
	sector   string
}

// Scrape Indeed by the term
func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?=&searchword=" + term + "&exp_min=1&exp_max=1&recruitSort=relation&recruitPageCount=50"
	var jobs []extractedJob
	c := make(chan []extractedJob) // 이 채널로 일자리 정보가 여러개 전달되므로, 채널의 타입은 그냥 extractedJob가 아닌 []extractedJob
	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ { // 총 totalPages개 만큼의 goroutine이 생성
		go getPage(i, baseURL, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)                          // extractJob이 goroutines의 channel을 통해 extractedJob struct 형태를 보낼 것임.
	pageURL := url + "&recruitPage=" + strconv.Itoa(page) // strconv.Itoa() : go에서 지원하는 string으로 바꾸는 함수
	fmt.Println("Requesting ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) { // 현재 찾은건 각각의 카드
		go extractJob(card, c) // extracJob 함수에 채널을 인자로 입력
	})
	// extractJob 함수는 카드 하나마다 실행 될거기 때문애, searchCards의 길이만큼 반복문 실행. 현재는 50
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs // return이 아니라 넘겨주기
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	company := cleanString(card.Find(".area_corp>.corp_name>a").Text())
	title := cleanString(card.Find(".area_job>.job_tit>a").Text())
	location := cleanString(card.Find(".area_job>.job_condition>span>a").Text())
	sector := cleanString(card.Find(".area_job>.job_sector>a").Text())
	c <- extractedJob{ // goroutines
		id:       id,
		company:  company,
		title:    title,
		location: location,
		sector:   sector,
	}
}

func getPages(url string) int {
	pages := 0
	// goQuery 내용을 살펴보면, get을 사용할때 에러처리를 해줘야하고, 도큐먼트 만들때도 에러처리해줘야함
	res, err := http.Get(url)
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

func writeJobs(jobs []extractedJob) { // csv 파일 저장 관련 함수 // go standard method 페이지에서 찾을 수 있음
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file) // writer 생성
	defer w.Flush()          // 함수가 끝나는 시점, writer에 데이터 입력

	headers := []string{"Link", "Company", "Title", "Location", "Sector"}

	wErr := w.Write(headers) // Write 메소드는 에러를 반환
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.company, job.title, job.location, job.sector}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
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

// CleanString string
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ") // 빈공간을 제외한 문자열만 반환 제거한 내용을 join으로 합침
}
