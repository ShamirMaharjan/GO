package main

import "fmt"

type dog struct {
	name string
	age  int
}

func printStructure(d dog) {
	fmt.Println(d.name)
	fmt.Println(d.age)
}

func structure() {
	dog1 := dog{
		name: "dog1",
		age:  10,
	}
	printStructure(dog1)

}

//anonymous struct
//this is a anonymous structure it can be used only once
// myCar := struct {
// 	Make string
// 	Model string
// }{
// 	Make: "Toyota",
// 	Model: "Corolla",
// }

//embedded struct
//car is a embedded structure, so truck can access all the fields of car
//like truct.Make,
//we dont need to use truct.car.Make
//we can access all the fields of car directly
type car struct {
	Make  string
	Model string
	Year  int
}

type truck struct {
	car
	Trailer bool
}
