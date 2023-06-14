package main

import (
	"fmt"

	"github.com/m/viviviviviid/learngo/ch2_bank_and_dictionary_project/banking"
)

/// @title public한 struct를 사용할 때 -> 대신 이건 아무나 접근가능하고, 누구든 변경가능함.
// // 하지만 우린 은행관련이기 때문에, 아무나가 아니어야함
// func main() {
// 	account := banking.Account{Owner: "nicolas", Balance: 1000}
// 	fmt.Println(account)
// }

// / @title go에서 constructor 만들기
func main() {
	account := banking.NewAccount("nico")
	fmt.Println(account)
	// account.balance = 10000 이런식으로 public하게 사용불가.
	// account.balance는 private 하기 때문에. // banking.go 파일 참고.
}
