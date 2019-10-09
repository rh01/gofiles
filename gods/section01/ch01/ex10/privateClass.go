// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

import (
	"encoding/json"
	"fmt"
)

// importing fmt and encoding/json packages

//AccountDetails struct
type AccountDetails struct {
	id          string
	accountType string
}

//Account struct
type Account struct {
	details      *AccountDetails
	CustomerName string
}

// Account class method setDetails
func (account *Account) setDetails(id string, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

//Account class method getId
func (account *Account) getId() string {
	return account.details.id
}

//Account class method getAccountType
func (account *Account) getAccountType() string {
	return account.details.accountType
}

// main method
func main() {
	var account *Account = &Account{CustomerName: "John Smith"}
	account.setDetails("4532", "current")
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden", string(jsonAccount))
	fmt.Println("Account Id", account.getId())
	fmt.Println("Account Type", account.getAccountType())
}
