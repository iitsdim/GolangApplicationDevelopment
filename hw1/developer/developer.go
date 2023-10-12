package developer

type Developer struct {
	position string
	salary   float64
	address  string
}

func (a *Developer) SetSalary(salary float64) {
	a.salary = salary
}

func (a *Developer) SetPosition(position string) {
	a.position = position
}

func (a *Developer) SetAddress(address string) {
	a.address = address
}

func (a *Developer) GetSalary() float64 {
	return a.salary
}

func (a *Developer) GetPosition() string {
	return a.position
}

func (a *Developer) GetAddress() string {
	return a.address
}
