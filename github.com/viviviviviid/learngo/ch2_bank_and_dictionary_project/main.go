package main

import (
	"fmt"

	mydict "github.com/m/viviviviviid/learngo/ch2_bank_and_dictionary_project/dict"
)

// / @title
func main() {
	word := "hello"
	definition := "Greeting"

	dictionary := mydict.Dictionary{}
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	value, err := dictionary.Search(word)
	fmt.Println(value)
	err2 := dictionary.Add(word, definition)
	if err2 != nil { // 이미 존재하므로 에러
		fmt.Println(err2)
	}

}

// / @title Search in dict
// func main() {
// 	dictionary := mydict.Dictionary{"first": "First word"}
// 	// 기본
// 	// definition, err := dictionary.Search("first")

// 	// 없는 second를 찾아서 에러 발생시키기.
// 	definition, err := dictionary.Search("second")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(definition)
// 	}
// }

// / @title accounts.go
// func main() {
// 	account := accounts.NewAccount("nico")
// 	// account.balance = 10000 이런식으로 public하게 사용불가.
// 	// account.balance는 private 하기 때문에. // banking.go 파일 참고.
// 	// 그래서 메소드를 만들고 이용해야함.
// 	account.Deposit(20)            // accounts.go에서 만든 메소드
// 	fmt.Println(account.Balance()) // .Balance도 accounts.go에서 만든 메소드

// 	err := account.Withdraw(10) // 에러 핸들링을 통해서, 금액이 음수값이 안나오게 에러 코드를 출력
// 	if err != nil {             // 에러가 있다면
// 		log.Fatalln(err) // 종료시키면서 에러코드 출력 // error을 체크하도록 강제하는 golang
// 	}
// 	fmt.Println(account.Balance(), account.Owner())
// 	account.ChangeOwner("minseok")
// 	fmt.Println(account.Balance(), account.Owner())
// 	fmt.Println(account)
// }

/// @title public한 struct를 사용할 때 -> 대신 이건 아무나 접근가능하고, 누구든 변경가능함.
// // 하지만 우린 은행관련이기 때문에, 아무나가 아니어야함
// func main() {
// 	account := account.Account{Owner: "nicolas", Balance: 1000}
// 	fmt.Println(account)
// }
