package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Hello, World!")

	// var name string = "ayushi"
	// var name2 = "anula"
	// var name3 string
	// fmt.Println(name, name2, name3)

	// name = "arundhati"
	// fmt.Println(name)

	// name4 := "khushi"
	// fmt.Println(name4)

	// //int

	// var age int = 20
	// var age2 = 30
	// age3 := 40

	// fmt.Println(age, age2, age3)

	// //bit memory
	// var rednum int8 = 55
	// var bluenum int8 = 127
	// var yellownum int16 = 225
	// var pinknum uint16 = 335
	// fmt.Println(rednum, bluenum, yellownum, pinknum)

	// //float
	// var scoreone float32 = -222.645
	// fmt.Println(scoreone)

	// someone := 643.753

	// fmt.Println(someone)

	// //print

	// fmt.Print("hello ")
	// fmt.Print("world!\n")
	// fmt.Print("HI ANULA!!\n")
	// var name = "Ayushi"
	// var height = 172
	// var score = 3428.342
	// fmt.Print("My name is ", name, " and my height is ", height)

	// fmt.Printf("\nheight is of type %T", height)
	// fmt.Printf("\nyou scored %0.3f points\n", score)
	// fmt.Printf("\nMy name is %v and my height is %v", name, height)
	// fmt.Printf("\nMy name is %q and my height is %q", name, height)

	// var str = fmt.Sprintf("My name is %v and my height is %v", name, height)

	// fmt.Println("\nThe saved string is: ", str)

	// var ages [3]int = [3]int{20, 25, 30}
	// var names = [5]string{"Ayushi", "Anula", "Arundhati", "Khushi"}
	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// //Slices
	// var scores = []int{100, 50, 0, 60}

	// scores[2] = 35
	// scores = append(scores, 55)
	// fmt.Println(scores, len(scores))

	// // Slice ranges
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]
	// fmt.Println(rangeOne)
	// fmt.Println(rangeTwo)
	// fmt.Println(rangeThree)

	greetings := "Good Morning!"
	fmt.Println(strings.Contains(greetings, "oo"))

	fmt.Println((strings.ReplaceAll(greetings, "Morning", "Afternoon")))
	fmt.Println(greetings)
	fmt.Println(strings.ToUpper(greetings))
}
