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

func (p while_ter) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {

	Linicio := gen.NewLabel()
	Lfinal := gen.NewLabel()
	temp := gen.NewTemp()
	var isType interfaces.TypeExpression
	tree.AddDisplay(Linicio, Lfinal, "-1", false) // Display

	gen.AddComment("while")
	gen.AddLabel(Linicio)

	var newTable interfaces.Environment
	newTable = interfaces.NewEnvironment(env)
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable)

	for _, s := range p.Instrucciones.ToArray() {
		s.(interfaces.Instruction).Compilar(&newTable, tree, gen)

	}

	gen.AddGoto(Linicio)
	gen.AddLabel(Lfinal)
	gen.AddComment("Break temp")
	pos := fmt.Sprintf("%v", tree.PosDisplay-1)
	display := tree.GetDisplay(pos)
	if display.IsTemp {
		gen.AddExpression1(temp, display.Temp)
		isType = display.Type
	} else {

		excep := interfaces.NewException("Semantico", "Error en while, Break no trea una Expresion.", p.Row, p.Column)
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
	}
	tree.RestPosDisplay()

	return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
}
