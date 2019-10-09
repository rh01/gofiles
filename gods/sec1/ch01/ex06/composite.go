//main package has examples shown\
package main

//1.The component interface defines the default behavior of all objects and
//behaviors for accessing the components of the composite.
//2.The composite and component classes implement the component interface.
//3.The client interacts with the component interface to invoke methods in the composite.

// in Hands-On Data Structures and algorithms with Go book
// importing fmt package
import (
	"fmt"
)

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet struct
type Leaflet struct {
	name string
}

// Leaflet class method perform
func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

// Branch struct
type Branch struct {
	leafs    []Leaflet
	branches []Branch
	name     string
}

//Branch class method perform
//The perform method of the Branch class calls the perform method
//on branch and leafs, as seen in the code:
func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}
	for _, branch := range branch.branches {
		branch.perform()
	}
}

// Branch class method add leaflet
func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

//Branch class method addBranch branch
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

//Branch class method getLeaflets
func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

// main method
func main() {
	var branch = &Branch{name: "branch 1"}
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}
	var branch2 = Branch{name: "branch 2"}
	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)
	branch.perform()
}
