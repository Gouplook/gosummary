package gofunc

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender int
	Height int

	Country string
	City    string
}

type Options func(*Person)

func WithPersonProperty(name string, age, gender, height int) Options {
	return func(p *Person) {
		p.Name = name
		p.Age = age
		p.Gender = gender
		p.Height = height
	}
}

func WithRegional(country, city string) Options {
	return func(p *Person) {
		p.Country = country
		p.City = city
	}
}

func NewPerson(opt ...Options) *Person {
	p := new(Person)
	p.Country = "china"
	p.City = "beijing"
	for _, o := range opt {
		o(p)
	}
	return p
}

func Demo_main() {

	// 默认值方式
	person := NewPerson(WithPersonProperty("dongxiaojian", 18, 1, 180))
	fmt.Printf("%+v\n", person)

	// 设置值
	person2 := NewPerson(WithPersonProperty("dongxiaojian", 18, 1, 180), WithRegional("china", "hebei"))
	fmt.Printf("%+v\n", person2)
}
