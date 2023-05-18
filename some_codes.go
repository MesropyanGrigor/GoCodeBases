package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const aInt int = 42

func main2() {
	other := 12.3
	m := make(map[string]int)
	//m := new(map[string]int)
	m["theAnswer"] = 34
	m["otherAnswer"] = 456
	fmt.Println(m)
	//i1, i2, i3 := 12, 45, 68
	//implicitly will infer type as a integerp
	fmt.Printf("The type is %T\n", aInt)
	fmt.Printf("The type is %T\n", other)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text: ")
	input, _ := reader.ReadString('\n')

	fmt.Println("You entered: ", input)

	fmt.Println("Enter a new number: ")
	number, err := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(number), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value of number: ", aFloat)
	}

	colors := []string{"RED", "GREEN", "BLUE"}

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

}
func main3() {
	println(addValues(12, 23))
	println(addAllValues(12, 23, 34))

}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

type Counter struct {
	var count int
}

func main() {
	count := Counter{0}

	println(count.addCount())
	count.addCount()
	count.addCount()
	count.addCount()
	count.addCount()
	println(count.addCount())
}

func (this Counter) addCount() int {
	this.count += 1
	return this.count
}
