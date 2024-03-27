package main

import (
	"fmt"
)

func main() {
	print1()
	printBonus1()
	printBonus2()
}

// Belajar Struct dan Interface

// type print interface {
// 	participant()
// 	mean()
// 	rule()
// 	bonus(q participant)
// }

// Create struct for data participants
type participant struct {
	name                   string
	point1, point2, point3 float64
}

// func printInterface (printPrint print) {
// 	fmt.Println("Nama Peserta : %s", print.mean())

// }

// Create struct for mean
func (p participant) mean() float64 {
	return (p.point1 + p.point2 + p.point3) / 3
}

// Create struct for condition rule >= 100
func (p participant) rule() bool {
	data := p.mean()
	if data >= 100 {
		fmt.Printf("-> %s memenuhi syarat >= 100 untuk lanjut ke tahap selanjutnya \n Dengan nilai : %.2f\n", p.name, p.mean())
		return bool(true)
	} else {
		fmt.Printf("-> %s tidak memenuhi syarat >= 100 untuk lanjut ke tahap selanjutnya \n dengan nilai : %.2f\n", p.name, p.mean())
		return bool(false)
	}
}

//  Try to make condition fo rule but success failed
// func (p participant) bonus(q participant) participant {
// 	if p.mean() == q.mean() && p.rule() {
// 		fmt.Printf("Hasil dari perolehan nilai SERI\nDengan nilai : \nTim Lumba-Lumba : %f \nTim Koala : %f\n ", p.mean(), q.mean())
// 	} else if p.mean() < q.mean() && q.rule() {
// 		fmt.Printf("Tropi diperoleh oleh tim %s dengan nilai : %f\n ", q.name, q.mean())
// 	} else if p.mean() > q.mean() && p.rule() {
// 		fmt.Printf("Tropi diperoleh oleh tim %s dengan nilai : %f\n ", p.name, p.mean())
// 	} else {
// 		fmt.Printf("Tidak ada pemenang trofi")
// 	}
// 	return participant{}
// }

// Make condition rule for who win, lose, or tie
func (p participant) bonus(q participant) participant {
	if p.rule() {
		if p.mean() == q.mean() {
			fmt.Printf("=> Hasil dari perolehan nilai SERI\nDengan nilai : \nTim Lumba-Lumba : %f \nTim Koala : %f\n ", p.mean(), q.mean())
		} else if p.mean() > q.mean() {
			fmt.Printf("=> Tropi diperoleh oleh tim %s dengan nilai : %f\n ", p.name, p.mean())
		} else if p.mean() < q.mean() {
			fmt.Printf("=> Tropi diperoleh oleh tim %s dengan nilai : %f\n ", q.name, q.mean())
		}
	} else if q.rule() {
		if q.mean() > q.mean() {
			fmt.Printf("=> Tropi diperoleh oleh tim %s dengan nilai : %f\n ", q.name, q.mean())
		}
	} else {
		fmt.Printf("=> Tidak ada pemenang trofi")
	}
	return participant{}
}

func (p participant) info(q participant) participant {
	fmt.Printf("\nNilai dari masing-masing tim : \n%s : %.2f, %.2f, %.2f (%.2f)\n%s : %.2f, %.2f, %.2f (%.2f)\n", p.name, p.point1, p.point2, p.point3, p.mean(), q.name, q.point1, q.point2, q.point3, q.mean())
	return participant{}
}

func print1() {
	participant1 := participant{"Lumba-Lumba", 96, 108, 89}
	participant2 := participant{"Koala", 88, 91, 110}
	// fmt.Printf("\nNilai dari masing-masing tim : \n%s : %.2f, %.2f, %.2f (%.2f)\n%s : %.2f, %.2f, %.2f (%.2f)\n", participant1.name, participant1.point1, participant1.point2, participant1.point3, participant1.mean(), participant2.name, participant2.point1, participant2.point2, participant2.point3, participant2.mean())
	participant1.info(participant2)
	participant2.rule()
	participant1.bonus(participant2)
	fmt.Println()
}

func printBonus1() {
	participant3 := participant{"Lumba-Lumba", 97, 112, 101}
	participant4 := participant{"Koala", 109, 95, 123}
	// fmt.Printf("\n--BONUS 1--\nNilai dari masing-masing tim : \n%s : %.2f, %.2f, %.2f (%.2f)\n%s : %.2f, %.2f, %.2f (%.2f)\n", participant3.name, participant3.point1, participant3.point2, participant3.point3, participant3.mean(), participant4.name, participant4.point1, participant4.point2, participant4.point3, participant4.mean())
	participant3.info(participant4)
	participant4.rule()
	participant3.bonus(participant4)
	fmt.Println()
}

func printBonus2() {
	participant5 := participant{"Lumba-Lumba", 97, 112, 101}
	participant6 := participant{"Koala", 109, 95, 106}
	// fmt.Printf("\n--BONUS 2--\nNilai dari masing-masing tim : \n%s : %.2f, %.2f, %.2f (%.2f)\n%s : %.2f, %.2f, %.2f (%.2f)\n", participant5.name, participant5.point1, participant5.point2, participant5.point3, participant5.mean(), participant6.name, participant6.point1, participant6.point2, participant6.point3, participant6.mean())
	participant5.info(participant6)
	participant6.rule()
	participant5.bonus(participant6)
	fmt.Println()
}
