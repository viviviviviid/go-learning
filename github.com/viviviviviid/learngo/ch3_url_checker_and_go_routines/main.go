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
// func main() {
// 	// chan : channel의 타입, chan bool : 메인과의 채널로 bool 값을 보내주겠다
// 	c := make(chan string)
// 	people := [2]string{"nico", "flynn"}
//
// 	for _, person := range people {
// 		// goroutines의 go로는 아래와 같이 못함.
// 		// result := go isSexy(person) // return문은 안먹히므로 -> 채널로만 진행해야함
// 		go isSexy(person, c)
// 	}
// 	fmt.Println("Waiting for message")
// 	result := <-c                                  // 채널을 통해 값을 받는 방법
// 	fmt.Println("Received this message: ", result) // "<-" 이건 채널을 통해 값을 받고 있다는 뜻이고, 이때는 메인문이 멈춰짐 // blocking operation
// 	// 하나의 값을 받으면 다음 라인으로 넘어감 // 그전까지는 멈춤
// 	fmt.Println("Received this message: ", <-c) // 하나의 채널로 두개 받기 // 다른 형식
// 	// fmt.Println(<-c) // 하나의 채널로 두개만 보냈기에 세개째에는 오류가 발생함. // 메세지는 기다리고 있지만 goroutines이 종료되었기에.
// }
//
// func isSexy(person string, c chan string) { // 채널을 통해 어떤 타입을 보낼지 go에게 인지시켜 줘야함
// 	time.Sleep(time.Second * 5)
// 	// return true // return은 안먹히므로, 채널로 보내줘야함
// 	c <- person + " is sexy" // 이게 채널을 통해 값을 메인으로 보내는 방법
// }

// / @title Loop of Channel
func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		// 하나하나 채널로 받던 이전 내용과는 다르게, 반복문을 사용해서 깔끔하게 처리.
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
