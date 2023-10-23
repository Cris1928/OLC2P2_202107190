package ternario

import (
	"OLC2/Compilador/interfaces"
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

type while_ter struct {
	Instrucciones *arrayList.List
	Row           int
	Column        int
}

// NewWhileter
func NewWhileter(instruccion *arrayList.List, row int, column int) while_ter {
	instr := while_ter{instruccion, row, column}
	return instr
}

func (p while_ter) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value { // ESTE ES EL QUE COMPILA

	Linicio := gen.NewLabel() // Label Inicio
	Lfinal := gen.NewLabel()  // Label Final del while
	temp := gen.NewTemp()     // Temporal para el Break
	var isType interfaces.TypeExpression
	tree.AddDisplay(Linicio, Lfinal, "-1", false) // Display

	gen.AddComment("while") // Comentario
	gen.AddLabel(Linicio)   // Label Inicio

	var newTable interfaces.Environment                                           // Nueva tabla de simbolos
	newTable = interfaces.NewEnvironment(env)                                     // Se crea la nueva tabla de simbolos
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable) // Se actualiza la posicion de la tabla de simbolos

	for _, s := range p.Instrucciones.ToArray() { // Se recorren las instrucciones
		s.(interfaces.Instruction).Compilar(&newTable, tree, gen) // Se compila cada instruccion

	} // Fin for

	gen.AddGoto(Linicio)                        // Se agrega un goto al inicio del while, para que se repita
	gen.AddLabel(Lfinal)                        // Se agrega el label final del while, para que se salga del while
	gen.AddComment("Break temp")                // Se agrega un comentario, para saber que es el break
	pos := fmt.Sprintf("%v", tree.PosDisplay-1) // Se obtiene la posicion del display, para obtener el display
	display := tree.GetDisplay(pos)             // Se obtiene el display, para obtener el temporal del break
	if display.IsTemp {                         // Si es temporal, se agrega la expresion
		gen.AddExpression1(temp, display.Temp) // Se agrega la expresion, para que se salga del while
		isType = display.Type                  // Se obtiene el tipo, para retornarlo
	} else { // Si no es temporal, se agrega una excepcion

		excep := interfaces.NewException("Semantico", "Error en while, Break no trea una Expresion.", p.Row, p.Column) // Se crea una nueva excepcion
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}
	tree.RestPosDisplay()

	return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""} //
}
