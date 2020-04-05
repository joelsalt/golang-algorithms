package notes

// 2. STRUCTS
func structs() {
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	vegeta := Saiyan{"Vegeta", 6000}
	krillin := Saiyan{}

	println(goku.Name, vegeta.Name, krillin.Name)
}

type Saiyan struct {
	Name  string
	Power int
}

func pointerExample() {
	goku := &Saiyan{"Goku", 9000} // '&' gets the memory address of variable or struct
	super(goku)
	println(goku.Power)
}

func super(superSaiyan *Saiyan) { // '*X' represents pointer to value of type X
	superSaiyan.Power += 10000
}

func (s *Saiyan) methodOnStruct() { // '*Saiyan' is the receiver of the Super method
	s.Power += 10000
}

func methodOnStructExample() {
	vegeta := Saiyan{"Vegeta", 5000}
	vegeta.methodOnStruct()
	println(vegeta.Power)
}

// Use factories to return an instance of a struct, like a constructor
func newSaiyan(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}
