package main

import (
	"fmt"
	"github.com/hassanh12/rivercrossing/state"
)

func main() {
	fmt.Println(state.ViewState())
	fmt.Println(state.PutInBoat())
	fmt.Println(state.CrossRiver())
}
