package main

import (
	"GolangApplicationDevelopment/hw1/designer"
	"fmt"
)

func main() {
	des := designer.Designer{}
	des.SetSalary(100)
	des.SetAddress("Barcelona")
	des.SetPosition("Senior Designer")
	fmt.Println("position:", des.GetPosition(), "salary:", des.GetSalary(), "location:", des.GetAddress())
}
