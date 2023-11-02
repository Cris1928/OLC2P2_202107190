package ternario

import (
	"OLC2/Compilador/interfaces"

	arrayList "github.com/colegno/arraylist"
)

type If struct {
	Condicion   interfaces.Expression
	InstrIf     interfaces.Expression
	InstrElse   interfaces.Expression
	InstrElseIf *arrayList.List
	Row         int
	Column      int
}

func NewIf(condicion interfaces.Expression, instrIf interfaces.Expression, instrElse interfaces.Expression, instrElseIf *arrayList.List, row int, column int) If { // aqui se crea el objeto
	instr := If{condicion, instrIf, instrElse, instrElseIf, row, column} // se crea el objeto
	return instr                                                         // se retorna el objeto
}

func (p If) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value { //funcion compilar, recibe el entorno, el arbol y el generador de codigo

	gen.AddComment("Ternario - If")             //agregamos un comentarios
	var cond, result interfaces.Value           //creamos dos variables de tipo value
	cond = p.Condicion.Compilar(env, tree, gen) // compilamos la condicion

	if cond.Type == interfaces.EXCEPTION { // si la condicion es de tipo excepcion
		return cond // retornamos la condicion
	}

	if cond.Type == interfaces.BOOLEAN { // si la condicion es de tipo booleano

		EV := gen.NewLabel()       // creamos un nuevo label
		EF := gen.NewLabel()       // creamos un nuevo label
		newLabel := gen.NewLabel() // creamos un nuevo label
		temp := gen.NewTemp()      // creamos un nuevo temporal

		gen.AddExpression(temp, "0", "1", "-") // agregamos una expresion

		var isType interfaces.TypeExpression = interfaces.NULL // creamos una variable de tipo TypeExpression

		gen.AddIf(cond.Value, "1", "==", EV) // agregamos un if
		gen.AddGoto(EF)                      // agregamos un goto

		gen.AddLabel(EV)                              // agregamos un label
		result = p.InstrIf.Compilar(env, tree, gen)   // compilamos la instruccion if y la guardamos en result
		isType = result.Type                          // obtenemos el tipo de la instruccion if
		gen.AddExpression(temp, result.Value, "", "") // agregamos una expresion

		gen.AddGoto(newLabel) // agregamos un goto

		gen.AddLabel(EF) // agregamos un label

		if p.InstrElse != nil { // si la instruccion else no es nula
			gen.AddComment("Ternario - If (else)")        // agregamos un comentario
			result = p.InstrElse.Compilar(env, tree, gen) // compilamos la instruccion else y la guardamos en result

			gen.AddExpression(temp, result.Value, "", "") // agregamos una expresion

			if isType != result.Type { // si el tipo de la instruccion if es diferente al tipo de la instruccion else
				excep := interfaces.NewException("Semantico", "Tipos de Datos incorrectos en If  .", p.Row, p.Column)                           // creamos una nueva excepcion
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) // agregamos la excepcion al arbol
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}                    // retornamos un value
			}
		}

		if p.InstrElseIf != nil { // si la instruccion else if no es nula
			for _, s := range p.InstrElseIf.ToArray() { // recorremos la instruccion else if
				gen.AddComment("Ternario - If (else if)")   // agregamos un comentario
				newInstr := s.(If).Compilar(env, tree, gen) // compilamos la instruccion else if y la guardamos en newInstr
				if newInstr.Type != result.Type {           // si el tipo de la instruccion else if es diferente al tipo de la instruccion if
					excep := interfaces.NewException("Semantico", "Tipos Diferentes en If  .", p.Row, p.Column)                                     // creamos una nueva excepcion
					tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) // agregamos la excepcion al arbol
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}                    // retornamos un value
				}
				gen.AddComment("Ternario - Retorno de Temp") // agregamos un comentario
				newLabel1 := gen.NewLabel()                  // creamos un nuevo label
				newLabel2 := gen.NewLabel()                  // creamos un nuevo label

				gen.AddIf(newInstr.Value, "-1", "!=", newLabel1) // agregamos un if
				gen.AddGoto(newLabel2)                           // agregamos un goto

				gen.AddLabel(newLabel1)                         // agregamos un label
				gen.AddExpression(temp, newInstr.Value, "", "") // agregamos una expresion
				gen.AddGoto(newLabel2)                          // agregamos un goto
				gen.AddLabel(newLabel2)                         // agregamos un label

			}
		}
		gen.AddGoto(newLabel)

		gen.AddLabel(newLabel)
		return interfaces.Value{Value: temp, IsTemp: false, Type: result.Type, TrueLabel: "", FalseLabel: ""}

	} else {
		excep := interfaces.NewException("Semantico", "Tipo de Dato no Booleano en If  .", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}

}
