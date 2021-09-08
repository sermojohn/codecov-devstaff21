package greet

type Greeter struct{}

func (g *Greeter) GetGeneralGreeting() string {
	return "hello"
}

func (g *Greeter) GetMorningGreeting() string {
	return "good morning"
}

func (g *Greeter) GetSubject() string {
	return "devstaff"
}
