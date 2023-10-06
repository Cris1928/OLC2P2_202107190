package instruction

import (
	"OLC2/Compilador/interfaces"
	"fmt"
)

type Identifier struct {
	Id     string
	Row    int
	Column int
}

func NewIdentifier(id string, row int, column int) Identifier {
	instr := Identifier{id, row, column}
	return instr
}

func (p Identifier) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	/* Buscar si el id ya existe */
	symbol := env.SearchSymbol(p.Id)

	if symbol.Type == interfaces.NULL {
		symbol = env.GetSymbol(p.Id)

		if symbol.Type == interfaces.NULL {
			excep := interfaces.NewException("Semantico", "No Existe ese Id "+p.Id+"(identificador).", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

		}
	}

	if symbol.Type == interfaces.STRUCT {
		excep := interfaces.NewException("Semantico", "Falta Atributo para Struct (identificador).", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

	} else {

		gen.AddComment("Identificador")
		temp0 := gen.NewTemp()
		temp1 := gen.NewTemp()
		gen.AddExpression(temp0, "P", fmt.Sprintf("%v", symbol.Posicion), "+")
		gen.AddExpressionStack(temp1, temp0)

		return interfaces.Value{Value: temp1, IsTemp: true, Type: symbol.Value.(interfaces.Value).Type, TrueLabel: "", FalseLabel: ""}

	}

}
