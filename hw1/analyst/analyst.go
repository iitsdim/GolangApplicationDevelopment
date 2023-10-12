package analyst

type Analyst struct {
	position string
	salary   float64
	address  string
}

func (a *Analyst) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Analyst) SetPosition(position string) {
	a.position = position
}

func (a *Analyst) SetAddress(address string) {
	a.address = address
}

func (a *Analyst) GetSalary() float64 {
	return a.salary
}

func (a *Analyst) GetPosition() string {
	return a.position
}

func (a *Analyst) GetAddress() string {
	return a.address
}
