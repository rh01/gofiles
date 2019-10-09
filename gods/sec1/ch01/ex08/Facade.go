// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/**
Facade is used to abstract subsystem interfaces with a helper.
The facade design pattern is used in scenarios when the number
of interfaces increases and the system gets complicated.
Facade is an entry point to different subsystems,
and it simplifies the dependencies between the systems.

The facade pattern provides an interface that hides the implementation details of the hidden code.
A loosely coupled principle can be realized with a facade pattern.
You can use a facade to improve poorly designed APIs.
In SOA, a service facade can be used to incorporate changes
to the contract and implementation.

The facade pattern is made up of the facade class, module classes, and a client:
The facade delegates the requests from the client to the module classes.
The facade class hides the complexities of the subsystem logic and rules.
Module classes implement the behaviors and functionalities of the module subsystem.
The client invokes the facade method. The facade class functionality can be
spread across multiple packages and assemblies.

For example, account, customer, and transaction are the classes that
have account, customer, and transaction creation methods.

BranchManagerFacade can be used by the client to create an
account, customer, and transaction:
*/

//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

//Account struct
type Account struct {
	id          string
	accountType string
}

//Account class method create - creates account given AccountType
func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

//Account class method getById given id string
func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by Id")
	return account
}

//Account class method deleteById given id string
func (account *Account) deleteById(id string) {
	fmt.Println("delete account by id")
}

//Customer struct
type Customer struct {
	name string
	id   int
}

//Customer class method create - create Customer given name
func (customer *Customer) create(name string) *Customer {
	fmt.Println("creating customer")
	customer.name = name
	return customer
}

//Transaction struct
type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

//Transaction class method create Transaction
func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

//BranchManagerFacade struct
type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

//method NewBranchManagerFacade
func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

//BranchManagerFacade class method createCustomerAccount
func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)
	return customer, account
}

//BranchManagerFacade class method createTransaction
func (facade *BranchManagerFacade) createTransaction(srcAccountId string, destAccountId string, amount float32) *Transaction {
	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

//main method
func main() {
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account
	customer, account = facade.createCustomerAccount("Thomas Smith",
		"Savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	var transaction = facade.createTransaction("21456", "87345", 1000)
	fmt.Println(transaction.amount)
}
