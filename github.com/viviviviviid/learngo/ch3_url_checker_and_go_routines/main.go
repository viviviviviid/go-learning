package main

import (
	"fmt"
	"time"
)

/// @title hitURL start
// var errRequestFailed = errors.New("Request Failed")
//
// func main() {
// 	// map으로 초기화 된 map을 생성. 만약 초기화 안하고 값을 넣으려 한다면 panic이라는 원인불명의 에러가 발생
// 	var results = make(map[string]string)
// 	urls := []string{
// 		"https://www.airbnb.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.reddit.com/",
// 		"https://www.google.com/",
// 		"https://soundcloud.com/",
// 		"https://www.facebook.com/",
// 		"https://www.instagram.com/",
// 		"https://academy.nomadcoders.co/",
// 		"https://error.erorrorroro.com",
// 	}
//
// 	// 두가지 값이 나오는데, 첫번째는 인덱스, 두번째가 내용
// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}
// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}
//
// }
//
// func hitURL(url string) error {
// 	// hit : 특정 웹사이트의 파일 1개에 접속하는 것.
// 	//
// 	// golang standard package // https://pkg.go.dev/net/http@go1.20.5
// 	// resp, err := http.Get("http://example.com/")
// 	// resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
// 	// resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		return errRequestFailed
// 	}
// 	return nil
// }

// / @title Start goroutines
// func main() {
// 	go sexyCount("nico")
// 	go sexyCount("flynn")
// 	time.Sleep(time.Second * 5)
// }
//
// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep((time.Second))
// 	}
// }

// / @title Channel(like pipe) of goroutines
func main() {
	// chan : channel의 타입, chan bool : 메인과의 채널로 bool 값을 보내주겠다
	channel := make(chan bool)
	people := [2]string{"nico", "flynn"}

	for _, person := range people {
		// goroutines의 go로는 아래와 같이 못함.
		// result := go isSexy(person) // return문은 안먹히므로 -> 채널로만 진행해야함
		go isSexy(person, channel)
	}
	// result := <-channel // 채널을 통해 값을 받는 방법 // 메세지를 하나 받을때까지 기다림.
	// fmt.Println(result)
	// 아래와 같이 해도 됨.
	fmt.Println("채널 통해 받은 후", <-channel)
	fmt.Println("채널 통해 받은 후", <-channel) // 하나의 채널로 두개 받기
	// fmt.Println(<-channel) // 하나의 채널로 두개만 보냈기에 세개째에는 오류가 발생함. // 메세지는 기다리고 있지만 goroutines이 종료되었기에.
}

func isSexy(person string, channel chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println("채널로 보내기 전", person)
	// 이걸론 안먹히므로, 채널로 보내줘야함
	// return true
	channel <- true // 이게 채널을 통해 값을 메인으로 보내는 방법
}
