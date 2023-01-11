package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	// types

	name := "Gourav"
	fmt.Println(name)
	fmt.Printf("%T \n", name)

	age := 10
	fmt.Println(age)
	fmt.Printf("%T \n", age)

	//bufio, strconv, os, strings package

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating of pizza")

	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}
	fmt.Println("thank for the rating - ", input)
	fmt.Printf("type of rating %T \n", input)

	numint, err := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

	if err != nil {
		panic(err)
	}

	fmt.Println("thank for the rating - int - ", numint+1)
	fmt.Printf("type of rating %T \n", numint)

	// datetime

	presentTime := time.Now()

	fmt.Println("Present Time", presentTime)
	fmt.Println("Foramtted Present Time", presentTime.Format("01-02-2006"))
	fmt.Println("Foramtted Present Time", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdTime := time.Date(2020, time.October, 21, 21, 53, 0, 0, time.Local)
	fmt.Println("Created Time", createdTime)
	fmt.Println("Created Time", createdTime.Format("01-02-2006 15:04:05 Monday"))

	// pointers

	number := 10
	pointer := &number
	fmt.Println("pointer address", pointer)
	fmt.Println("pointer value", *pointer)

	*pointer = *pointer + 2
	fmt.Println("pointer value", *pointer)
	fmt.Println("pointer value", number)

	// arrays

	fruits := []string{"apple", "orange"}
	fmt.Println("Fruits array", fruits)
	fmt.Println("Fruits array", len(fruits))

	//slices

	fruits = append(fruits, "mango")
	fmt.Println("Fruits array", fruits)

	fruits = fruits[0:2]
	fmt.Println("Fruits array", fruits)

	//make

	highscores := make([]int, 4)

	highscores[0] = 23
	highscores[1] = 1
	highscores[2] = 32
	highscores[3] = 10

	fmt.Println("high scores", highscores)
	sort.Ints(highscores)

	fmt.Println("sorted high scores", highscores)
	fmt.Println("sorted high scores", sort.IntsAreSorted(highscores))

	// maps

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "python"

	fmt.Println("Maps", languages)

	for key, value := range languages {
		fmt.Println("values - ", key, value)
	}

	// structs
	userobj := User{"Gourav", 20}
	fmt.Println("User", userobj)

	//if
	if userobj.Age > 10 {
		fmt.Println("Print Age", userobj.Age)
	} else {
		fmt.Println("Less than 10")
	}

	// switch

	switch userobj.Age {
	case 1:
		fmt.Println("switch - 1")
	case 20:
		fmt.Println("switch - 20")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("i value", i)
	}
	// defer

	defer fmt.Println("Defer in the last")

	//functions

	result := sum(2, 3)
	fmt.Println("functions - sum", result)

	result1, values := nvalues(2, 3, 4, 5, 6, 7)
	fmt.Println("functions - nvalues", result1, values)

	// methods part of struct

	getName := userobj.GetName()

	fmt.Println("method - name", getName)

	// files
	content := "Files"
	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("File Length", length)
	defer file.Close()
	ReadFile("./myfile.txt")
}

type User struct {
	Name string
	Age  int
}

func sum(a int, b int) int {
	return a + b
}

func nvalues(values ...int) (int, []int) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, values
}

func (u User) GetName() string {
	return u.Name
}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Databyte", string(databyte))
}
