package main

// go program express errors with values. an error is any type that implements the error interface
import "fmt"

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return x / y, nil
}
func errorHandling() {
	i, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)

	name := "shamir"

	number := 10.000

	fmt.Printf("hello my name is %v and my number is %f", name, number)

	//%v is a placeholder for the value of the variable name

	greeting := fmt.Sprintf("hello my name is %v", name)
	// Now you can use 'greeting' later, e.g., fmt.Println(greeting)
	fmt.Println(greeting)
}
