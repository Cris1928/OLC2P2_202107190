package instruction

import (
	"OLC2/Compilador/interfaces"
)

type ListExprefunc struct {
	Id     string
	Type   interfaces.TypeExpression
	Row    int
	Column int
}

func NewListExprefunc(val string, tipo interfaces.TypeExpression, row int, column int) ListExprefunc {
	exp := ListExprefunc{val, tipo, row, column}
	return exp
}

func (p ListExprefunc) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	return p
}
