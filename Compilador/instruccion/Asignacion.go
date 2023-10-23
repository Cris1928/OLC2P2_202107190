package instruction

import (
	"OLC2/Compilador/interfaces"
	"fmt"
)

type Assignment struct {
	Id        string
	Expresion interfaces.Expression
	Row       int
	Column    int
}

func NewAssignment(id string, val interfaces.Expression, row int, column int) Assignment {
	instr := Assignment{id, val, row, column}
	return instr
}

func (p Assignment) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} {

	/* Buscar si el id ya existe */
	symbol := env.GetSymbol(p.Id) //

	if symbol.Type == interfaces.NULL { // si no existe
		excep := interfaces.NewException("Semantico", "No Existe ese Id "+p.Id, p.Row, p.Column)                                        // error
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) // agregar a la lista de errores
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}                    // retornar valor nulo
	}

	var result interfaces.Value                   // valor de la expresion
	result = p.Expresion.Compilar(env, tree, gen) // compilar expresion

	if symbol.IsMut { // si es mutable
		gen.AddComment("Asignacion") // comentario
		// fmt.Println(symbol.Type)
		// fmt.Println(symbol.Value.(interfaces.Value).Type )
		if symbol.Type == result.Type || symbol.Type == interfaces.NULL { // si el tipo es igual al tipo de la expresion o es nulo
			symbol.IsMut = true                                                   // es mutable
			env.SetSymbol(p.Id, result, true, symbol.Posicion)                    // actualizar el simbolo
			temp := gen.NewTemp()                                                 // nuevo temporal
			gen.AddExpression(temp, "P", fmt.Sprintf("%v", symbol.Posicion), "+") // P + posicion
			gen.AddStack(temp, result.Value)                                      // guardar en el stack

		} else {
			excep := interfaces.NewException("Semantico", "No se puede asignar, tipo de datos diferentes.", p.Row, p.Column)                // error
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) // agregar a la lista de errores
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}                    // retornar valor nulo

		}

	} else {
		excep := interfaces.NewException("Semantico", "No se puede asignar a "+p.Id+", no es mutable.", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	return result.Value

}
