package instruction

import (
	"OLC2/Compilador/interfaces"
	"fmt"
)

type Return struct {
	Expression interfaces.Expression
	Row        int
	Column     int
}

func NewReturn(expresion interfaces.Expression, row int, column int) Return {
	instr := Return{expresion, row, column}
	return instr
}

func (p Return) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} { // funcion compilar, el interface{} es para que retorne cualquier tipo de dato

	gen.AddComment("Return") //agregamos un comentario

	var result interfaces.Value                    //variable para guardar el resultado de la expresion
	result = p.Expression.Compilar(env, tree, gen) //obtenemos el resultado de la expresion

	if tree.IsReturn { //si ya se encontro un return en el metodo

		es := fmt.Sprintf("%v", tree.PosReturn) //obtenemos la posicion del return
		display := tree.GetDisplay(es)          //obtenemos el display del return
		if display.Type == result.Type {        //si el tipo del display es igual al tipo del resultado
			gen.AddStack("P", result.Value) //agregamos el valor del resultado al stack
			gen.AddGoto(display.LFinal)     //agregamos un goto al label final
		} else {
			excep := interfaces.NewException("Semantico", "Error en return, tipos de datos incorrectos a Retornar ", p.Row, p.Column)       //creamos una nueva excepcion
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) //agregamos la excepcion al arbol
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}                    //retornamos un value
		}

	} else {

		excep := interfaces.NewException("Semantico", "Error en return, tipos de datos incorrectos a Retornar ", p.Row, p.Column)       //creamos una nueva excepcion
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) //agregamos la excepcion al arbol
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}                    //retornamos un value

	}

	return result
}
