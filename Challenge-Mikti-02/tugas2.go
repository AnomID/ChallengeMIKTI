package main

import "fmt"

type person struct {
	name          string
	heigh, weight float64
}

func (p person) bmi() float64 {
	return (p.weight / (p.heigh * p.heigh))

}

func (p person) info() {
	fmt.Printf("Nama : %s\nTinggi : %.2f cm\nBerat : %.2f kg\nBMI : %f\n\n", p.name, p.heigh, p.weight, p.bmi())

}

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

func printPerson() {
	mark := person{"Mark", 169, 78}
	john := person{"John", 195, 92}

	mark.info()
	john.info()
	mark.markHigherBMI(john)
}

func printPerson2() {
	mark := person{"Mark", 188, 95}
	john := person{"John", 176, 85}

	mark.info()
	john.info()
	mark.markHigherBMI(john)
}

func main() {
	printPerson()
	printPerson2()
}
