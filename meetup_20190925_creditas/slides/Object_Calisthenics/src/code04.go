type friends struct {
	data []string
}

type person struct {
	name    string
	friends *friends
}

func (f *friends) Add(name string) {
	// Regra para adicionar
}

func (f *friends) Remove(name string) {
	// Regra para remover
}