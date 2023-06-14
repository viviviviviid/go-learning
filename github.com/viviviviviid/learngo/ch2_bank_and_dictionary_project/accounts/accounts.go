package accounts

/// @title public한 struct, 즉 아무나 접근 가능
// Account struct
// type Account struct {
// 	Owner   string
// 	Balance int
// }
// 뭔가를 export 하려면 위와 같이 주석을 무조건 달아줘야함. // 무조건 첫 단어는 struct 이름으로 작성해야함.
// 내부 인자의 이름이 소문자로 시작한다면 private. 즉 export해서 외부에서 사용하려면 대문자로 변경해줘야함. owner -> Owner

// / @title NewAccount용
// Account struct
type Account struct {
	owner   string
	balance int
}

// / @title go에서 constructor 만들기. function을 만들어서 object를 return. -> 실제 메모리를 return.
// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 새로 만들어진 복사본을 return하는게 아닌, 메모리 주소를 return.
	// 이렇게 한다면 main.go에서, account.balance = 10000 이런식으로 public하게 사용불가.
	// account.balance는 private 하기 때문에.
}

// / @title 'method' 만들기
// func 다음에 나오는 (a Account)는 Receiver로 이름을 아무렇게나 지을수는 있지만,
// 규칙에 의거하여 현재 이용하려는 struct인 Account의 첫글자인 A의 소문자를 사용해야함. ex (b, Banking)
// Deposit x amount on your amount.
func (a Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your Account
func (a Account) Balance() int {
	return a.balance
}
