package notes

// 3. COMPOSITION
type Person struct {
	Name string
}

func (p *Person) Introduce() {
	println("My name is " + p.Name)
}

type Superhero struct {
	*Person
	HeroName string
}

func compositionExample() {
	spiderman := Superhero{
		Person:   &Person{"Miles Morales"},
		HeroName: "SpiderMan",
	}
	spiderman.Introduce()
}
