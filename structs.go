package main

//Struct is a collection of different types of data types
import "fmt" //why didn't he add this part in the tutorial

type Student struct {
	name string;
	level int;
	dept string;
	school string;
	courses []string;

}

func main() {
	student1 := Student{
//You don't need to rewrite the key name e.g name := Ben, you can write Ben directly but it's not advisable
		name:"Benita",
		level: 300,
		dept: "Chemical Engineering",
		school: "FUPRE",
		courses:[]string{
			"Che 313", 
			"Get 311", 
			"Che 311", 
			"Mee 317",
		},//why does it have to have a comma here?
	}
	fmt.Println(student1)
	fmt.Println(student1.courses[2])
	student1.courses[4] = "Che 312"
	fmt.Println(student1.courses[4])
}