package member

type Member struct {
	Name  string
	Age   int
	IsVip bool
}

type Member2 struct {
}

func AddMember(name string, age int, isVip bool) *Member {
	return &Member{
		Name:  name,
		Age:   age,
		IsVip: isVip,
	}
}

func (user *Member) SetName(name string) {
	user.Name = name
}
