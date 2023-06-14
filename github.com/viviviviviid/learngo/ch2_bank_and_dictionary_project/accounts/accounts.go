package accounts

import "errors"

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
// func (a Account) Deposit(amount int) {
// 	a.balance += amount
// }
// Balance of your Account
// func (a Account) Balance() int {
// 	return a.balance
// }

// / @title 위 'method' 만들기에서 문제점 찾기 + pointer Receiver
// / @dev 기존 방식으로는 deposit을 해도 이 함수 내에서만 적용된 것처럼 보임. 이걸 해결할 것임.
// Deposit x amount on your amount.
func (a *Account) Deposit(amount int) {
	// Receiver에서 Account 앞에 *가 붙었는데, 이게 큰 차이점을 가짐. => "pointer receiver"
	// 이 pointer receiver 있다면, account.Deposit(10)에서 account의 값을 직접 변경하는 것이고,
	// 이게 없다면 이 메소드를 호출한 코드라인인 account.Deposit(10)에서 account의 복사본을 가져와 쓰는 것임.
	// 없을때에는 복사본의 값을 증가시키기에, 실제 바뀐게 없을 것임. 즉 Balance() 메소드를 이용해도 그대로인 것.
	a.balance += amount
}

// Balance of your Account
func (a Account) Balance() int { // 얘는 반환만 하는 것이므로 굳이 직접 접근할 필요가 없음
	return a.balance
}

// / @title 에러핸들링
// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error { // 에러를 리턴
	// golang에는 try-catch문이 없음
	if a.balance < amount {
		var errNoMoney = errors.New("Can't withdraw") // 외부에 둬도 상관없음 // 보기편하게 갖다 놓은 것 뿐
		// 위와같이 에러변수의 이름은 err000 이런형태가 되어야함
		return errNoMoney
	}
	a.balance -= amount
	return nil // null, none과 같음
}
