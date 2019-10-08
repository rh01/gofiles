package model

type student struct {
	name string
	age  int
}

func (s *student) GetName() string {
	return s.name
}

func NewStudent(name string, age int) *student {
	return &student{
		name:name,
		age:age,
	}
}