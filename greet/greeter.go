package greet

type Greeter struct{}

func (g *Greeter) GetGreeting() string {
	return "hello"
}

func (g *Greeter) GetSubject() string {
	return "devstaff"
}
