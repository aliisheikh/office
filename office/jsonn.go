package main

import "fmt"

type Companyy struct {
	address
	company string
}

type address struct {
	street  string
	city    string
	state   string
	zipcode int
}

type departmentt struct {
	namee    string
	head     bool
	name     string
	position string
	namme    string
}
type Employee struct {
	id       int
	name     string
	position string
	skills   []string
}

type projects struct {
	ProjectId   string
	Name        string
	Description string
	Team        []string
	Status      string
	name        string
	projectId   string
	team        interface{}
	description string
	status      string
}

//	type marketing struct {
//		name: "Marketing",
//		head: {
//		"name": "Emily Brown",
//		"position": "CMO"
//	}
type pro struct {
	projectId   string
	name        string
	description string
	team        []string
	status      string
}

func main() {

	co1 := Companyy{
		address: address{
			street:  "123 Main Street",
			city:    "Tech ville",
			state:   "Techsylvania",
			zipcode: 12345,
		},

		company: "TechCorp",
	}

	fmt.Println(co1)
	fmt.Println("\n")
	dep := departmentt{
		namee:    "Engineering",
		head:     true,
		name:     "John Doe",
		position: "CTO",
	}
	fmt.Println(dep)
	fmt.Println("\n")
	//{
	//
	//}
	emp1 := Employee{
		id:       1,
		name:     "Jane Smith",
		position: "Software Engineer",
		skills:   []string{"JavaScript", "Python", "React"},
	}

	emp2 := Employee{
		id:       2,
		name:     "Alex Johnson",
		position: "Frontend Developer",
		skills:   []string{"HTML", "CSS", "JavaScript"},
	}

	fmt.Println(emp1)
	fmt.Println(emp2)

	fmt.Println("\n")

	p := projects{
		projectId:   "proj123",
		name:        "E-commerce Platform",
		description: "Developing an online shopping platform",
		team:        []string{"Jane Smith", "Alex Johnson"},
		status:      "In Progress",
	}
	fmt.Println(p)
	fmt.Println("\n")
	d := departmentt{
		namee:    "Marketing",
		head:     true,
		name:     "Emily Brown",
		position: "CMO",
	}
	fmt.Println(d)
	fmt.Println("\n")
	pro1 := projects{
		projectId:   "proj456",
		name:        "Social Media Campaign",
		description: "Launching a new social media campaign",
		team:        []string{"Michael Johnson"},
		status:      "Completed",
	}
	fmt.Println(pro1)
	fmt.Println("\n")

}
