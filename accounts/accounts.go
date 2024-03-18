package banking

// Account struct
type Account struct {
	owner   string
	balance int
}

//누구나 나의 Balance를 0 마음대로 조정하지 못하도록 대문자를 소문자로 지정후 퍼블릭을 만든다
// NewAccount create Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}