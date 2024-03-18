package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

//누구나 나의 Balance를 0 마음대로 조정하지 못하도록 대문자를 소문자로 지정후 퍼블릭을 만든다
//실제 메모리 account를 return
// NewAccount create Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}

//balance ++ 증가 /-- 감소
//Deposit x amount onyour account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance of account
func (a Account) Balance() int {
	return a.balance
}