package main

import (
	"fmt"
	"time"
)

// Employee ...
type Employee struct {
	ID      int
	Name    string
	Age     *time.Time
	Salary  float64
	Company string
}

func main() {
	usingStructs()
	usingStructsAnonymous()
}

func usingStructs() {
	cl := Employee{}
	//forma de acesso
	cl.ID = 1
	cl.Name = "Diego dos Santos"
	cl.Age = nil
	cl.Salary = 100.55
	cl.Company = "Fliper"
	fmt.Println("o nome é:", cl.Name, " trabalha na empresa: ", cl.Company)
	//oura forma de popular structs
	cl1 := Employee{
		ID:      1,
		Name:    "Francisco Oliveira",
		Age:     nil,
		Salary:  2000.50,
		Company: "Iron Mountain",
	}
	fmt.Println("o nome é:", cl1.Name, " trabalha na empresa: ", cl1.Company)

}

func usingStructsAnonymous() {
	inferData("Diego", "Santos")
	inferData("Francisco", "Oliveira")
}

func inferData(fN, lN string) {
	name1 := struct{ FirstName, LastName string }{FirstName: fN, LastName: lN}
	fmt.Println("O nome é:", name1.FirstName, name1.LastName)
}
