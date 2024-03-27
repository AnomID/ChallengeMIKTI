package main

import "fmt"

// Create struck person for object creation
type person struct {
	name          string
	heigh, weight float64
}

// Create function for BMI formula from object use struct person
func (p person) bmi() float64 {
	return (p.weight / (p.heigh * p.heigh))

}

// Create output information for data person and value BMI
func (p person) info() {
	fmt.Printf("Nama : %s\nTinggi : %.2f m\nBerat : %.2f kg\nBMI : %f\n\n", p.name, p.heigh, p.weight, p.bmi())

}

// function to compare BMI Mark and John from object use struct person p and q
func (p person) markHigherBMI(q person) bool {

	mark := p.bmi()
	john := q.bmi()
	if mark > john {
		mark := bool(true)
		fmt.Println("Hasil dari boolean BMI Mark lebih tinggi dari John : ", mark)
		fmt.Println("\n-----=====-----")
	} else {
		mark := bool(false)
		fmt.Println("Hasil dari boolean BMI Mark lebih rendah dari John : ", mark)
		fmt.Println("\n-----=====-----")
	}
	return bool(false)
}

// function output for data 1
//
//	Mark = {name, height, weight}
//
// John = {name, height, weight}
func printPerson() {
	mark := person{"Mark", 1.69, 78}
	john := person{"John", 1.95, 92}

	mark.info()
	john.info()
	mark.markHigherBMI(john)
}

// function output for data 2
func printPerson2() {
	mark := person{"Mark", 1.88, 95}
	john := person{"John", 1.76, 85}

	mark.info()
	john.info()
	mark.markHigherBMI(john)
}

// Function Main
func main() {
	printPerson()
	printPerson2()
}
