//show person1 OMIT
type person struct {
    name string
    cpf  int
}
//end show person1 OMIT

//show person2 OMIT
type person struct {
    name string
    cpf  cpf
}

type cpf int

func (c cpf) isValid() bool {
	// Regra de validação
}

func (c cpf) String() string {
	// Regra de formatação
}
//end show person2 OMIT