package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type skillSet struct {
	name  []string
	years []int
}

type person struct {
	lastName  string
	firstName string
	contactInfo
	skillSet
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p *person) addSkill(skill string, year int) {
	(*p).skillSet.name = append(p.skillSet.name, skill)
	(*p).skillSet.years = append(p.skillSet.years, year)
}

func (p person) changeSkillYear(skill string, year int) {
	for i := 0; i < len(p.skillSet.name); i++ {
		if p.skillSet.name[i] == skill {
			p.skillSet.years[i] = year
		}
	}
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Harrison",
		contactInfo: contactInfo{
			zipCode: 123456,
			email:   "test@test.com",
		},
		skillSet: skillSet{
			name: []string{
				"HTML", "JavaScript", "Java",
			},
			years: []int{
				1, 2, 3,
			},
		},
	}
	alex.updateName("Alexander")
	alex.addSkill("CSS", 2)
	alex.changeSkillYear("Java", 10)
	alex.print()
}
