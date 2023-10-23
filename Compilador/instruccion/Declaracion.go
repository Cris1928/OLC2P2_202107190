package instruction

import (
	"OLC2/Compilador/interfaces"
	"fmt"
	// "reflect"
)

type Declaration struct {
	Id        string
	Type      interfaces.TypeExpression
	Expresion interfaces.Expression
	IsMut     bool
	Row       int
	Column    int
}

func NewDeclaration(id string, tipo interfaces.TypeExpression, val interfaces.Expression, isMut bool, row int, column int) Declaration { // ESTE ES EL QUE COMPILA
	instr := Declaration{id, tipo, val, isMut, row, column} // Se crea la instruccion
	return instr
}

func (p Declaration) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} { // ESTE ES EL QUE COMPILA

	/* Buscar si el id ya existe */
	symbol := env.SearchSymbol(p.Id) // Se busca el simbolo

	if symbol.Type != interfaces.NULL { // si existe
		// fmt.Println("No puede agregar")
		excep := interfaces.NewException("Semantico", "Ya Existe ese Id "+p.Id, p.Row, p.Column)                                        // error
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) // agregar a la lista de errores
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}                    // retornar valor nulo
	}

	var result interfaces.Value // valor de la expresion

	if p.Expresion != nil { // si tiene expresion
		result = p.Expresion.Compilar(env, tree, gen) // compilar expresion
		// fmt.Println("entro")
		if p.Type != interfaces.NULL {
			if p.Type != result.Type {
				excep := interfaces.NewException("Semantico", "No se puede declarar, tipo de datos diferentes.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}
		}

	} else {

		result.Type = p.Type
		result.IsTemp = false
		result.Value = "0"
	}

	if result.Type == p.Type || p.Type == interfaces.NULL { // si el tipo es igual al tipo de la expresion o es nulo
		gen.AddComment("Declaracion") // comentario, para saber que es una declaracion

		// fmt.Println("Local")
		temp := gen.NewTemp()                                                                                                              // nuevo temporal, para guardar en el stack
		gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")                                                                 // P + posicion, para obtener la posicion en el stack
		gen.AddStack(temp, result.Value)                                                                                                   // guardar en el stack
		env.AddSymbol(p.Id, result, result.Type, p.IsMut, env.Posicion, env)                                                               // agregar el simbolo a la tabla de simbolos
		tree.AddTableSymbol(*interfaces.NewTableSymbol(p.Id, "Variable", "Local", p.Row, p.Column, "--", fmt.Sprintf("%v", env.Posicion))) // agregar a la tabla de simbolos
		env.NewPos()                                                                                                                       // nueva posicion

	} else {
		excep := interfaces.NewException("Semantico", "Tipo Incorrecto en Declaracion.", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

	return result.Value // retornamos el

}
