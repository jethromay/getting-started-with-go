package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}

	// Our team struct
	type Team struct {
		name    string
		players [2]Person
	}

	// Declare an empty Team
	// var myTeam Team
	// fmt.Println(myTeam)

	players := [...]Person{Person{name: "Jethro"}, Person{name: "John"}}

	arsenal := Team{name: "Arsenal", players: players}

	fmt.Println(arsenal)
}
