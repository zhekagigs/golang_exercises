package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/bearbin/go-age"
)

// Create a program that has multiple string variable and displays the string on one line. [Strings]
func assignmentOne() {
	stringA := "One"
	stringB := "Two"
	stringC := "Three"
	fmt.Printf("%v %v %v\n", stringA, stringB, stringC)
	fmt.Println(strings.Join([]string{stringA, stringB, stringC}, "-"))
}

// Create a program that lets the user input a first name, middle name and last name. Display the person's full name on one line. [Keyboard input]
func assignmentTwo() {
	var firstName string
	var midName string
	var lastName string
	fmt.Println("Enter you first name, middle name and last name. E.g John Percival Smith")
	_, err := fmt.Scanf("%s %s %s", &firstName, &midName, &lastName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("your first name is %s middle name is %s last name is %s\n", firstName, midName, lastName)
}

// Create a program that allows the user to input a number. Check whether the number lies between 1 and 10. [Variables]
func assignmentThree() {
	var userNum int
	fmt.Println("Enter a number")
	_, err := fmt.Scanf("%d", &userNum)
	if err != nil {
		panic(err)
	}
	if 1 <= userNum && userNum <= 10 {
		fmt.Println("Number between 1 and 10")
	} else {
		fmt.Println("Number is out of range")
	}

}

// Create a program that initialises an array with the integer values 1 to 10: [Arrays][Slices][For Loops]
// Display the array content in ascending sequential order 1 to 10.
// Display the array content in descending sequential order 10 to 1.
// Count even numbers and odd numbers in increasing and decreasing sequential order.
// Display the even and odd count sequences to screen.
func assignmentFour() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var even, odd int
	for _, n := range arr {
		fmt.Println(n)
		if n%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}
	fmt.Printf("Odd %d Even %d\n", odd, even)
	even = 0
	odd = 0
	for i := len(arr) - 1; i > -1; i-- {
		fmt.Println(arr[i])
		if arr[i]%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}
	fmt.Printf("Odd %d Even %d\n", odd, even)
}

// Create a program that accepts and sums nine numbers. [Methods][Arrays][Slices][For loops]
// Three single digit numbers from one method.
// Three double digit numbers from a second method.
// Three triple digit numbers from a third method.
// Finally sum all methods into a final sum in the main program.
func assignmentFive() {
	np := NumberProvider{}
	total := np.SingleDigits().DoubleDigits().TripleDigits().Sum()
	fmt.Println(total)
}

type NumberProvider struct {
	values []int
	total  int
}

func (np *NumberProvider) SingleDigits() *NumberProvider {
	np.values = append(np.values, 1, 2, 3)
	return np
}

func (np *NumberProvider) DoubleDigits() *NumberProvider {
	np.values = append(np.values, 11, 12, 13)
	return np
}

func (np *NumberProvider) TripleDigits() *NumberProvider {
	np.values = append(np.values, 111, 112, 113)
	return np
}

func (np *NumberProvider) Sum() int {
	for _, v := range np.values {
		np.total += v
	}
	return np.total
}

func assignmentSix() {
	loc, err := time.LoadLocation("Europe/Volgograd")
	if err != nil {
		panic(err)
	}
	dateOfBirth := time.Date(1987, 7, 14, 10, 15, 00, 00, loc)
	ageToday := age.AgeAt(dateOfBirth, time.Now())
	fmt.Println("Age is ", ageToday)
}

func assignmentSeven() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var dice1 int
	var dice2 int

	for i := range 50 {
		dice1 = rnd.Intn(6) + 1
		dice2 = rnd.Intn(6) + 1
		fmt.Printf("Throw %d result %d -> ", i, dice1+dice2)
		diceName := diceName(dice1 + dice2)
		fmt.Println(diceName)
	}

}

func diceName(sum int) string {
	switch sum {
	case 7, 11:
		return "NATURAL"
	case 2:
		return "SNAKE-EYES-CRAPS"
	case 3, 12:
		return "LOSS-CRAPS"
	default:
		return "NEUTRAL"
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var cities = []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

//Create a program that: [Write File][Read File][I/O Package][I/O]
// Copies the following list of cities to a new file - "Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi".
// Reads a list of cities from the newly created file.
// Displays the list of cities in alphabetical order.

func assignmentEight() {
	fileWriteStrings()
	writerWriteStrings()
	writeCitiesJoined()
	readCities()
}

func fileWriteStrings() {
	f, err := os.Create("local.txt")
	check(err)
	defer f.Close()
	for _, citi := range cities {
		_, err := f.WriteString(citi + "\n")
		check(err)
	}
	err = f.Sync()
	check(err)
	fmt.Println("Done!")
	info, err := os.Stat("local.txt")
	check(err)
	fmt.Printf("File permissions: %v\n", info.Mode().Perm())
}

func writerWriteStrings() {
	f, err := os.Create("local_buffered.txt")
	check(err)
	defer f.Close()
	writer := bufio.NewWriter(f)

	for _, citi := range cities {
		_, err := writer.WriteString(citi + "\n")
		check(err)
	}
	err = writer.Flush()
	check(err)
	fmt.Println("Done!")
	info, err := os.Stat("local_buffered.txt")
	check(err)
	fmt.Printf("File permissions: %v\n", info.Mode().Perm())
}

func writeCitiesJoined() error {
	for i := 0; i < 2000; i++ {
		cities = append(cities, generateCityName())
	}
	content := strings.Join(cities, "\n") + "\n"
	return os.WriteFile("cities_joined.txt", []byte(content), 0644)
}

func readCities() {
	fmt.Println("Reading cities")
	data, err := os.ReadFile("local.txt")
	check(err)
	trimmedData := strings.TrimSpace(string(data))
	split := strings.Split(trimmedData, "\n")
	slices.Sort(split)
	for _, name := range split {
		fmt.Println(name)
	}
}

func main() {
	fmt.Println("Assignment 1")
	assignmentOne()
	fmt.Println("Assignment 2")
	assignmentTwo()
	fmt.Println("Assignment 3")
	assignmentThree()
	fmt.Println("Assignment 4")
	assignmentFour()
	fmt.Println("Assignment 5")
	assignmentFive()
	fmt.Println("Assignment 6")
	assignmentSix()
	fmt.Println("Assignment 7")
	assignmentSeven()
	fmt.Println("Assignment 8")
	assignmentEight()
	// fmt.Println("Assignment 9")
}
