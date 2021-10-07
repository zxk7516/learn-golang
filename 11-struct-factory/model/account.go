package model

import "fmt"

type Account struct {
	Name    string
	Pwd     string
	Balance float64
}


func NewAccount(name string,pwd string, balance float64) *Account {
  if len(name) < 6 || len(name) > 10 {
    fmt.Println("wrong accoutn length")
    return nil
  }
  if len(pwd) != 6 {
    fmt.Println("wrong pwd lenght")
    return nil
  }
  if balance < 0 {
    fmt.Println("wrong balance")
    return nil
  }
  return &Account{
    Name: name,
    Pwd: pwd,
    Balance: balance,
  }
}


func (acc *Account) Deposite(pwd string, amount float64) {
	if pwd != acc.Pwd {
		fmt.Println("password error")
		return
	}
	if amount <= 0 {
		fmt.Println("error")
		return
	}
	acc.Balance += amount
	fmt.Println("success")
}
func (acc *Account) Withdraw(pwd string, amount float64) {
	if pwd != acc.Pwd {
		fmt.Println("password error")
		return
	}
	if amount <= 0 || amount > acc.Balance {
		fmt.Println("not enouth")
		return
	}
	acc.Balance -= amount
	fmt.Println("success")
}

func (acc *Account) Query(pwd string) float64 {

	return acc.Balance

}

func _test() {
	acc := Account{"zxk", "123456", 888.99}
	fmt.Println("account :", acc.Query("123456"))

}
