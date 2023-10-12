package main

import (
	"GolangApplicationDevelopment/hw1/analyst"
	"GolangApplicationDevelopment/hw1/designer"
	"fmt"
)

func main() {
	des := designer.Designer{}
	des.SetSalary(100)
	des.SetAddress("Barcelona")
	des.SetPosition("Senior Designer")
	fmt.Println("position:", des.GetPosition(), "salary:", des.GetSalary(), "location:", des.GetAddress())

	x := analyst.Analyst{}
	x.SetSalary(1000)
	x.SetAddress("Barcelona")
	x.SetPosition("Business Analyst")
	fmt.Println("position:", x.GetPosition(), "salary:", x.GetSalary(), "location:", x.GetAddress())

}
