package main

import "fmt"

func main() {

	integers()

	floatingPointNumbers()

	booleans()

	strings()

	arrays()

	slices()

	maps()

}

func integers() {
	// integers
	fmt.Println("-----------------Integers-----------------")
	fmt.Println(-123)

	var magicNumber int = 3788

	fmt.Printf("Magic number: %d\n", magicNumber)

	var num2 int = 2
	sum := magicNumber + num2

	fmt.Printf("sum: of %d and %d: %d\n", magicNumber, num2, sum)

}

func floatingPointNumbers() {
	// floating point numbers
	fmt.Println("-----------------Floating point numbers-----------------")
	absoluteZero := -256.367
	fmt.Printf("Absolute zero: %f\n", absoluteZero)
}

func booleans() {
	// booleans
	fmt.Println("-----------------Booleans-----------------")
	var myBool bool = 5 > 6
	fmt.Printf("Absolute zero: %t\n", myBool)
}

func strings() {
	// Strings
	fmt.Println("-----------------Strings-----------------")
	var rawStr string = `I am a raw string \n and you 
	can see that I did not escape the "\n" above`

	fmt.Printf("Raw string: %s\n", rawStr)

	var interpretedStr string = "Say \"Hello\" to go"
	fmt.Printf("Interpreted String: %s\n", interpretedStr)
}

func arrays() {
	// Arrays
	fmt.Println("-----------------Arrays-----------------")

	coral := [3]string{"blue coral", "staghorn coral", "pillar coral"}
	fmt.Println(coral)
}

func slices() {
	// Slices
	fmt.Println("-----------------Slices-----------------")

	seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp"}
	fmt.Println(seaCreatures)
	seaCreatures = append(seaCreatures, "seaHorse")
	fmt.Printf("Appended sea creatues %s\n", seaCreatures)
}

func maps() {
	// maps
	fmt.Println("-----------------Maps-----------------")

	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Printf("Sammy shark details:%s\n", sammy)
	fmt.Printf("Sammy's color: %s\n", sammy["color"])
}
