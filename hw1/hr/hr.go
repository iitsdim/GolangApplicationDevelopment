package hr

type HR struct {
	position string
	salary   float64
	address  string
}

func (a *HR) SetSalary(salary float64) {
	a.salary = salary
}

func (a *HR) SetPosition(position string) {
	a.position = position
}

func (a *HR) SetAddress(address string) {
	a.address = address
}

func (a *HR) GetSalary() float64 {
	return a.salary
}

func (a *HR) GetPosition() string {
	return a.position
}

func (a *HR) GetAddress() string {
	return a.address
}
