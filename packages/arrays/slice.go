package arrays

import "fmt"

//Slices Example
//There are 2 ways to define a slice
func Slices() {
	//1st method to declare as an array without specifying a size
	var array = []int{10, 20, 30, 40, 50, 60}
	array = array[:3]

	//2nd method to use the make function
	//1st param. is the length, 2nd is the capacity
	arrayMake := make([]int, 6, 10)

	fmt.Println(cap(arrayMake))
	fmt.Println(len(arrayMake))
	fmt.Println(arrayMake)

	fmt.Println(cap(array))
	fmt.Println(len(array))
	fmt.Println(array)

	array = append(array, 100)
	array = append(array, 100)
	fmt.Println(array)

	//copy array's content into arrayMake
	copy(arrayMake, array)
	fmt.Println(arrayMake)
}

//Range Example
//The range keyword is used in for loop to iterate over items of an array, slice, channel or map.
// it returns the index of the item as integer. With maps, it returns the key of the next key-value pair.
//Range either returns one value or two.
func Range() {
	numbers := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}
	for i, num := range numbers {
		fmt.Println(i)
		fmt.Println(numbers[i])
		fmt.Println(num)
	}
}

//Map Example
//var map_variable map[key_data_type]value_data_type
func Map() {
	countryCapitalMap := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
	}

	for key, value := range countryCapitalMap {
		fmt.Println(key, value)
	}

	delete(countryCapitalMap, "France")

	fmt.Printf("\n\n\n")

	for key, value := range countryCapitalMap {
		fmt.Println(key, value)
	}

}
