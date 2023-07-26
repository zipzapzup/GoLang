package main

import "fmt"

type School struct {
	Name        string
	Established int
	Subjects    []string
	Students    map[string]int
}

func main() {

	fmt.Println("School Struct")
	var newSchool School                                 // initialise new struct
	newSchool = School{Name: "SchoolA", Established: 10} // assignment operation

	fmt.Println(newSchool)

	newSchool.Subjects = append(newSchool.Subjects, "Math", "English", "Chemistry", "Science", "Art")

	fmt.Println(newSchool)

	newSchool.Students = make(map[string]int)
	newSchool.Students["Bob"] = 2
	newSchool.Students["John"] = 2
	newSchool.Students["Stuart"] = 2

	fmt.Println(newSchool)

	for _, item := range newSchool.Subjects {
		fmt.Println(item)
	}

	for key, value := range newSchool.Students {
		fmt.Println(key, value)
	}

	type School1 struct {
		Name        string
		Established int
		Subjects    []string
		Students    map[string]int
	}

	var Sydney School1 = School1{Name: "Sydney", Subjects: []string{"History", "Maths"}}
	Sydney.Students = map[string]int{"John": 99, "Betty": 100}
	fmt.Println(Sydney)

	addSubject("Taekwondo", &newSchool)
	fmt.Println(newSchool)
}

// func to add subjects
// when passing a pointer, we can append the subject since this is a reference to the actual variable
func addSubject(name string, s *School) {
	s.Subjects = append(s.Subjects, name)
}
