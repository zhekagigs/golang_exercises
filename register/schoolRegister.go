package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bearbin/go-age"
)

var (
	firstNames = []string{"John", "Emma", "Michael", "Sophia", "William", "Olivia", "James", "Ava", "Benjamin", "Isabella"}
	lastNames  = []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"}
)

type Register struct {
	students []Printable
}

type Printable interface {
	String() string
}

func (s Student) String() string {
	return fmt.Sprintf("%v at age %v", s.fullname, s.age)
}

func (r *Register) DisplayAll() {
	for _, item := range r.students {
		fmt.Println(item.String())
	}
}

type Student struct {
	fullname    string
	dateOfBirth time.Time
	age         int
}

func getRandomStudent() Student {
	// Generate random name
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	fullname := firstName + " " + lastName

	// Generate random date of birth (between 18 and 25 years ago)
	now := time.Now()
	yearsAgo := rand.Intn(8) + 18 // 18 to 25 years
	year := now.Year() - yearsAgo
	month := time.Month(rand.Intn(12) + 1) // 1 to 12
	day := rand.Intn(28) + 1
	dateOfBirth := time.Date(year, month, day, 10, 15, 00, 00, time.UTC)
	ageToday := age.AgeAt(dateOfBirth, now)

	return Student{
		fullname:    fullname,
		dateOfBirth: dateOfBirth,
		age:         ageToday,
	}
}

// Create a school register program that lists 10 pupils - full name, date of birth and age. [Structures][Arrays][Interfaces]
func main() {
	reg := Register{}
	for i := 0; i < 10; i++ {
		student := getRandomStudent()
		fmt.Printf("Student %d:\n", i+1)
		fmt.Printf("  Name: %s\n", student.fullname)
		fmt.Printf("  Date of Birth: %s\n", student.dateOfBirth.Format("2006-01-02"))
		fmt.Printf("  Age: %d\n\n", student.age)
		reg.students = append(reg.students, student)
	}
	reg.DisplayAll()

}
