package instruction

import (
	"OLC2/Compilador/interfaces"

	arrayList "github.com/colegno/arraylist"
)

type If struct {
	Condicion   interfaces.Expression
	InstrIf     *arrayList.List
	InstrElse   *arrayList.List
	InstrElseIf *arrayList.List
	Row         int
	Column      int
}

func NewIf(condicion interfaces.Expression, instrIf *arrayList.List, instrElse *arrayList.List, instrElseIf *arrayList.List, row int, column int) If { // aqui se crea el objeto
	instr := If{condicion, instrIf, instrElse, instrElseIf, row, column}
	return instr
}

func (p If) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} { // ESTE ES EL QUE COMPILA

	var cond interfaces.Value                   //creamos una variable de tipo value
	cond = p.Condicion.Compilar(env, tree, gen) // compilamos la condicion

	if cond.Type == interfaces.EXCEPTION { // si la condicion es de tipo excepcion
		return cond // retornamos la condicion
	}

	gen.AddComment("Control - If") //agregamos un comentarios

	if cond.Type == interfaces.BOOLEAN { // si la condicion es de tipo booleano

		var newTable interfaces.Environment                                           //creamos un nuevo entorno
		newTable = interfaces.NewEnvironment(env)                                     //creamos un nuevo entorno
		newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable) //actualizamos la posicion

		EV := gen.NewLabel()       // creamos un nuevo label
		EF := gen.NewLabel()       // creamos un nuevo label
		newLabel := gen.NewLabel() // creamos un nuevo label

		gen.AddIf(cond.Value, "1", "==", EV) // agregamos un if
		gen.AddGoto(EF)                      // agregamos un goto

		gen.AddLabel(EV)                        // agregamos un label
		for _, s := range p.InstrIf.ToArray() { // recorremos las instrucciones
			s.(interfaces.Instruction).Compilar(&newTable, tree, gen) // compilamos cada instruccion
		}
		gen.AddGoto(newLabel) // agregamos un goto

		gen.AddLabel(EF) // agregamos un label

		if p.InstrElse != nil { // si la instruccion else no es nula
			gen.AddComment("Control - If (else)")                                         // agregamos un comentario
			newTable = interfaces.NewEnvironment(env)                                     //creamos un nuevo entorno
			newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable) //actualizamos la posicion

			for _, s := range p.InstrElse.ToArray() { // recorremos las instrucciones

				s.(interfaces.Instruction).Compilar(&newTable, tree, gen) // compilamos cada instruccion

			}

		}

		if p.InstrElseIf != nil { // si la instruccion else if no es nula
			gen.AddComment("Control - If (else if)")    // agregamos un comentario
			for _, s := range p.InstrElseIf.ToArray() { // recorremos las instrucciones
				s.(interfaces.Instruction).Compilar(env, tree, gen) // compilamos cada instruccion

			}
		}
		gen.AddGoto(newLabel) // agregamos un goto

		gen.AddLabel(newLabel) // agregamos un label

	} else {
		excep := interfaces.NewException("Semantico", "Error en If. Tipo de Dato no Booleano.", p.Row, p.Column)                        // creamos una nueva excepcion
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) // agregamos la excepcion al arbol
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}                    // retornamos un value
	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""} // retornamos un value
}
