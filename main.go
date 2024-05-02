package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	fmt.Println("1. Built-in function")
	fmt.Println("2. areaOfRectangle")
	fmt.Println("3. searchingProgram")
	fmt.Println("4. rangeOfValue")
	fmt.Println("5. Exit")
	var numberString string
	fmt.Print("Input number: ")
	fmt.Scan(&numberString)
	switch numberString {
	case "1":
		clearScreen()
		builtInFunction()
	case "2":
		clearScreen()
		areaOfRectangle()
	case "3":
		clearScreen()
		searchingProgram()
	case "4":
		clearScreen()
		rangeOfValue()
	case "5":
		clearScreen()
		os.Exit(0)
	default:
		fmt.Println("Wrong character inputted")
		time.Sleep(2 * time.Second)
		clearScreen()
		main()
	}
}

func builtInFunction() {
	fmt.Println("===================== 1. Built-in Function =====================")
	// 1. Append built-in function
	fmt.Println("1. Append built-in function")
	var sliceAppend = []string{"How", "are", "you"}
	fmt.Println("Before append: ", sliceAppend)
	sliceAppend = append(sliceAppend, "Mr.", "Muhammad", "bin", "Ishaq")
	fmt.Println("After append: ", sliceAppend)

	//2. Cap built-in function is to retrun the capacity of a variable particularly array, slice, and channel
	fmt.Println("\n2. Cap built-in function is to retrun the capacity of a variable particularly array, slice, and channel")
	var arrayCap = [5]int{1, 2, 3}
	fmt.Println("arrayCap: ", arrayCap)
	fmt.Println("Cap of arrayCap: ", cap(arrayCap))

	//3. clear built-in function is to clear variable with map and slice data type
	fmt.Println("\n3. clear built-in function is to clear variable with map and slice data type")
	var mapClear = map[string]any{
		"name": "Farhan Nur Hakim",
		"age":  22,
	}
	fmt.Println(mapClear)
	for key, val := range mapClear {
		fmt.Println(key, ":\t", val)
	}
	clear(mapClear)
	fmt.Println(mapClear)

	//4. complex built-in function is to construct a complex value from two floating-point values
	fmt.Println("\n4. complex built-in function is to construct a complex value from two floating-point values")
	var realValue float64 = 2.32
	var imaginaryValue float64 = 4.33
	var complexValue complex128 = complex(realValue, imaginaryValue)
	fmt.Println("Complex value: ", complexValue)

	//5. copy built-in function is to copy elements from a source slice into a destination slice
	fmt.Println("\n5. copy built-in function is to copy elements from a source slice into a destination slice")
	var sourceCopy = []string{"Say", "my", "name"}
	var destinationCopy = []string{"My", "name", "is", "Heisenberg"}
	fmt.Println("Destination before copy: ", destinationCopy)
	copy(destinationCopy, sourceCopy)
	fmt.Println("Destination after copy: ", destinationCopy)

	//6. imag built-in function is to return imaginary part of the complex number
	fmt.Println("\n6. imag built-in function is to return imaginary part of the complex number")
	var imaginaryPart float64 = imag(complexValue)
	fmt.Println("Imaginary part of this complex value: ", complexValue, " is ", imaginaryPart)

	//7. len built-in function is to return the length of a variable
	fmt.Println("\n7. len built-in function is to return the length of a variable")
	var mapLen = map[string]any{
		"name":    "Ahmad bin Hambal",
		"age":     25,
		"hobbies": []string{"writing", "read science", "jogging"},
	}
	fmt.Println("Data length of this variable: ", mapLen, " is ", len(mapLen))

	//8. delete built in function is to delete an item in a map based on its key
	fmt.Println("\n8. delete built in function is to delete an item in a map based on its key")
	var mapDel = map[string]any{
		"name":         "Ahmad bin Hambal",
		"age":          25,
		"hobbies":      []string{"writing", "read science", "jogging"},
		"phone number": "089878768678688",
		"address":      "Jakarta, Indonesia",
	}
	fmt.Println("mapDel before delete: ", mapDel)
	delete(mapDel, "phone number")
	fmt.Println("mapDel after delete: ", mapDel)

	//9. make built-in function is to initialize an object with slice, map, and chan data type
	fmt.Println("9. make built-in function is to initialize an object with slice, map, and chan data type")
	var mapMake map[string]any
	mapMake = make(map[string]any)
	mapMake["name"] = "KH Agus Salim"
	fmt.Println(mapMake)

	//10. max built-in function returns the largest value of arguments
	fmt.Println("\n10. max built-in function returns the largest value of arguments")
	maxValue := max(1, 4, 6, 5, 9, 8)
	fmt.Println("maxValue: ", maxValue)

	//11. min built-in function returns the smallest value
	fmt.Println("\n11. min built-in function returns the smallest value")
	minValue := min(8, 7, 5, 2, 1, 9)
	fmt.Println("minValue: ", minValue)

	fmt.Println("\n\n1. re-run this function")
	fmt.Println("2. go to main()")
	var numberString string
	fmt.Print("Input number: ")
	fmt.Scan(&numberString)
	switch numberString {
	case "1":
		clearScreen()
		builtInFunction()
	case "2":
		clearScreen()
		main()
	default:
		fmt.Println("Because the character you inputted is wrong, you will be directed to main()")
		time.Sleep(2 * time.Second)
		clearScreen()
		main()
	}

}

type Rectangle struct {
	width  float64
	height float64
}

func areaOfRectangle() {
	fmt.Println("===================== 2. areaOfRectangle =====================")

	var rectangle = Rectangle{20.1, 10.5}
	fmt.Println("Result :", rectangle.areaOfRectangle(rectangle.width, rectangle.height))

	fmt.Println("1. re-run this function")
	fmt.Println("2. go to main()")
	var numberString string
	fmt.Print("Input number: ")
	fmt.Scan(&numberString)
	switch numberString {
	case "1":
		clearScreen()
		areaOfRectangle()
	case "2":
		clearScreen()
		main()
	default:
		fmt.Println("Because the character you inputted is wrong, you will be directed to main()")
		time.Sleep(2 * time.Second)
		clearScreen()
		main()
	}
}
func (rectangle Rectangle) areaOfRectangle(width float64, height float64) float64 {
	result := width * height
	return result
}

func searchingProgram() {
	fmt.Println("===================== 3. searchingProgram =====================")
	var searchedNames []string
	var search string
	var amountString string
	var amount int

	for {
		fmt.Print("Input characters you want to search on list of names: ")
		fmt.Scan(&search)
		fmt.Print("Input amount of names you want to search: ")
		fmt.Scan(&amountString)
		tempAmount, errAmount := strconv.Atoi(amountString)
		amount = tempAmount
		if errAmount != nil {
			fmt.Println("Data type should be integer")
		} else if amount <= 0 {
			fmt.Println("The amount should be more than 0")
		} else {
			break
		}
	}

	searchName(search, amount, func(name, search string) bool {
		return strings.Contains(name, search)
	}, &searchedNames)
	fmt.Println("searchedNames: ", searchedNames)
	fmt.Println("1. re-run this function")
	fmt.Println("2. go to main()")
	var numberString string
	fmt.Print("Input number: ")
	fmt.Scan(&numberString)
	switch numberString {
	case "1":
		clearScreen()
		searchingProgram()
	case "2":
		clearScreen()
		main()
	default:
		fmt.Println("Because the character you inputted is wrong, you will be directed to main()")
		time.Sleep(2 * time.Second)
		clearScreen()
		main()
	}

}
func searchName(search string, amount int, callback func(string, string) bool, searchedNames *[]string) {
	var names = []string{
		"Abigail", "Alexandra", "Alison",
		"Amanda", "Angela", "Bella",
		"Carol", "Caroline", "Carolyn",
		"Deirdre", "Diana", "Elizabeth",
		"Ella", "Faith", "Olivia", "Penelope",
	}
	for _, name := range names {
		if filtered := callback(strings.ToLower(name), search); filtered && len(*searchedNames) < amount {
			*searchedNames = append(*searchedNames, name)
		}
	}
}

func rangeOfValue() {
	fmt.Println("===================== 4. rangeOfValue =====================")
	var numbers []int
	var amountOfNumbersStr string
	var amountOfNumbers int
	var numberStr string
	var firstValueStr string
	var lastValueStr string
	var firstValue int
	var lastValue int
	for {
		fmt.Print("Input amount of numbers:")
		fmt.Scan(&amountOfNumbersStr)
		tempAmountOfNumbers, errAmountOfNumbers := strconv.Atoi(amountOfNumbersStr)
		if errAmountOfNumbers != nil {
			fmt.Println("Data types should be integer")
		} else if tempAmountOfNumbers <= 5 {
			fmt.Println("Amount of numbers should be more than 5")
		} else {
			amountOfNumbers = tempAmountOfNumbers
			break
		}
	}
	for i := 0; i < amountOfNumbers; i++ {
		for {
			fmt.Print("Input number at index ", i, ":")
			fmt.Scan(&numberStr)
			tempNumber, errNumber := strconv.Atoi(numberStr)
			if errNumber != nil {
				fmt.Println("Data type should be integer")
			} else {
				numbers = append(numbers, tempNumber)
				break
			}

		}
	}
	fmt.Println("Numbers: ", numbers)
	for {
		fmt.Print("Input first limit value:")
		fmt.Scan(&firstValueStr)
		tempFirstValue, errFirstValue := strconv.Atoi(firstValueStr)
		if errFirstValue != nil {
			fmt.Println("Data type should be integer")
		} else {
			firstValue = tempFirstValue
			break
		}
	}
	for {
		fmt.Print("Input last limit value:")
		fmt.Scan(&lastValueStr)
		tempLastValue, errLastValue := strconv.Atoi(lastValueStr)
		if errLastValue != nil {
			fmt.Println("Data type should be integer")
		} else {
			lastValue = tempLastValue
			break
		}
	}
	result := valueSelection(firstValue, lastValue, numbers)
	fmt.Println(result)
	fmt.Println("1. re-run this function")
	fmt.Println("2. go to main()")
	var numberString string
	fmt.Print("Input number: ")
	fmt.Scan(&numberString)
	switch numberString {
	case "1":
		clearScreen()
		rangeOfValue()
	case "2":
		clearScreen()
		main()
	default:
		fmt.Println("Because the character you inputted is wrong, you will be directed to main()")
		time.Sleep(2 * time.Second)
		clearScreen()
		main()
	}

}
func valueSelection(firstValue int, lastValue int, numbers []int) []int {

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	var selectedNumbers []int
	if firstValue >= lastValue {
		fmt.Println("First value should be less than last value")
	} else {
		for _, val := range numbers {
			if firstValue < val && lastValue > val {
				selectedNumbers = append(selectedNumbers, val)
			}
		}
	}
	return selectedNumbers
}
