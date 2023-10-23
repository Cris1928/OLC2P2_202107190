package instruction

import (
	"OLC2/Compilador/interfaces"
	"fmt"

	// "reflect"
	arrayList "github.com/colegno/arraylist"
)

type For struct {
	Id            string
	Type          interfaces.TypeExpression
	Left          interfaces.Expression
	Right         interfaces.Expression
	Instrucciones *arrayList.List
	Row           int
	Column        int
}

func NewFor(id string, tipo interfaces.TypeExpression, left interfaces.Expression, right interfaces.Expression, instrucciones *arrayList.List, row int, column int) For { //
	instr := For{id, tipo, left, right, instrucciones, row, column}
	return instr
}

func (p For) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} { //

	var newTable interfaces.Environment // declaracion de tabla de simbolos
	newTable = interfaces.NewEnvironment(env)
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable) // actualizacion de posicion

	if p.Type == interfaces.INTEGER { // si es entero

		gen.AddComment("For - Integer")               // comentario
		left := p.Left.Compilar(&newTable, tree, gen) // compilar la expresion izquierda
		if left.Type == interfaces.EXCEPTION {        // si es una excepcion
			return left // retornar la excepcion
		}

		right := p.Right.Compilar(&newTable, tree, gen) // compilar la expresion derecha
		if right.Type == interfaces.EXCEPTION {         // si es una excepcion
			return right // retornar la excepcion
		}

		symbol := newTable.GetSymbol(p.Id) // obtener el simbolo

		if symbol.Type == interfaces.NULL { // si el simbolo no existe

			gen.AddComment("Declaracion") // comentario

			temp := gen.NewTemp()                                                                                                                         // genera un nuevo temporal                                                                                                                        // genera un nuevo temporal                                                                                                                         // nuevo temporal
			gen.AddExpression(temp, "P", fmt.Sprintf("%v", newTable.Posicion), "+")                                                                       // aqui                                                               // agregar expresion
			gen.AddStack(temp, left.Value)                                                                                                                // agregar a la pila
			newTable.AddSymbol(p.Id, left, left.Type, true, newTable.Posicion, &newTable)                                                                 // agregar el simbolo a la tabla de simbolos
			tree.AddTableSymbol(*interfaces.NewTableSymbol(p.Id, "Variable - for", "Local", p.Row, p.Column, "--", fmt.Sprintf("%v", newTable.Posicion))) // agregar el simbolo a la tabla de simbolos del arbol
			newTable.NewPos()                                                                                                                             // nueva posicion

		}

		Linicio := gen.NewLabel()         // Label Inicio, para el goto
		gen.AddLabel(Linicio)             // agregar el label
		symbol = newTable.GetSymbol(p.Id) // obtener el simbolo, para obtener la posicion
		gen.AddComment("Identificador")   // comentario, para saber que es un identificador

		temp := gen.NewTemp()                                            // nuevo temporal, para obtener la posicion
		gen.AddExpressionStack(temp, fmt.Sprintf("%v", symbol.Posicion)) // agregar expresion, para obtener la posicion

		EV := gen.NewLabel()                         // Label para el if, para saber si se cumple la condicion
		Lfinal := gen.NewLabel()                     // Label final, para el goto
		Lincre := gen.NewLabel()                     // Label para el incremento
		tree.AddDisplay(Lincre, Lfinal, "-1", false) // Display
		gen.AddComment("Relacional <")               // comentario
		gen.AddIf(temp, right.Value, "<", EV)        // agregar if, para saber si se cumple la condicion
		gen.AddGoto(Lfinal)                          // agregar goto, para salir del for

		gen.AddLabel(EV) // agregar el label, para saber si se cumple la condicion

		// var newTable interfaces.Environment
		// newTable = interfaces.NewEnvironment(env)

		var SecondTable interfaces.Environment                                                        // declaracion de tabla de simbolos, para el for
		SecondTable = interfaces.NewEnvironment(&newTable)                                            // se crea la tabla de simbolos, para el for
		SecondTable.UpdatePos(tree.GetPos(), newTable.Posicion, newTable.Posicion != 0, &SecondTable) // se actualiza la posicion de la tabla de simbolos

		for _, s := range p.Instrucciones.ToArray() { // se recorren las instrucciones
			s.(interfaces.Instruction).Compilar(&SecondTable, tree, gen) // se compila cada instruccion, en la tabla de simbolos del for
		}

		pos := fmt.Sprintf("%v", tree.PosDisplay-1) // se obtiene la posicion del display, para obtener el display
		display := tree.GetDisplay(pos)             // se obtiene el display, para obtener el temporal del break
		if display.IsTemp {                         // si es temporal, se agrega la expresion
			excep := interfaces.NewException("Semantico", "Error en For, Sentencia de Control incorrecta Break.", p.Row, p.Column)          // se crea una nueva excepcion
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) // se agrega a la lista de errores
		}

		gen.AddComment("Incremento For")                                 // comentario
		gen.AddLabel(Lincre)                                             // agregar el label, para saber si se cumple la condicion
		symbol = newTable.GetSymbol(p.Id)                                // obtener el simbolo, para obtener la posicion
		gen.AddComment("Identificador")                                  // comentario, para saber que es un identificador
		temp = gen.NewTemp()                                             // nuevo temporal, para obtener la posicion
		gen.AddExpressionStack(temp, fmt.Sprintf("%v", symbol.Posicion)) // agregar expresion, para obtener la posicion

		gen.AddComment("Asignacion")                           // comentario, para saber que es una asignacion
		symbol.IsMut = true                                    // es mutable
		gen.AddExpression(temp, temp, "1", "+")                // P + posicion
		left.Value = temp                                      // se actualiza el valor de la expresion izquierda
		newTable.SetSymbol(p.Id, left, true, symbol.Posicion)  // actualizar el simbolo
		gen.AddStack(fmt.Sprintf("%v", symbol.Posicion), temp) // guardar en el stack, para actualizar el valor de la variable

		gen.AddGoto(Linicio) // agregar goto, para salir del for
		// gen.AddIf()

		gen.AddLabel(Lfinal)  // agregar el label, para saber si se cumple la condicion
		tree.RestPosDisplay() // se restaura la posicion del display

	} else if p.Type == interfaces.STRING { // si es string

		gen.AddComment("For - String") // comentario

		left := p.Left.Compilar(&newTable, tree, gen) // compilar la expresion izquierda, para obtener el valor
		if left.Type == interfaces.EXCEPTION {        // si es una excepcion
			return left // retornar la excepcion
		}

		symbol := newTable.GetSymbol(p.Id) // obtener el simbolo, para obtener la posicion
		temp := gen.NewTemp()              // nuevo temporal, para obtener la posicion
		secondTemp := gen.NewTemp()        // nuevo temporal, para obtener el valor de la expresion izquierda

		if symbol.Type == interfaces.NULL { // si el simbolo no existe

			gen.AddComment("Identificador")                                                                                                               // comentario, para saber que es un identificador
			gen.AddExpression(temp, "P", fmt.Sprintf("%v", newTable.Posicion), "+")                                                                       // agregar expresion, para obtener la posicion
			gen.AddStack(temp, left.Value)                                                                                                                // agregar a la pila
			left.Type = interfaces.CHAR                                                                                                                   // se actualiza el tipo de la expresion izquierda
			newTable.AddSymbol(p.Id, left, interfaces.CHAR, true, newTable.Posicion, &newTable)                                                           // agregar el simbolo a la tabla de simbolos
			tree.AddTableSymbol(*interfaces.NewTableSymbol(p.Id, "Variable - for", "Local", p.Row, p.Column, "--", fmt.Sprintf("%v", newTable.Posicion))) // agregar el simbolo a la tabla de simbolos del arbol
			newTable.NewPos()                                                                                                                             // nueva posicion

			gen.AddExpressionStack(secondTemp, temp) // agregar expresion, para obtener el valor de la expresion izquierda
		}

		Linicio := gen.NewLabel() // Label Inicio, para el goto
		gen.AddLabel(Linicio)     // agregar el label

		thirdTemp := gen.NewTemp() // nuevo temporal, para obtener la posicion

		gen.AddExpressionHeap(thirdTemp, secondTemp)        // agregar expresion, para obtener el valor de la expresion izquierda
		gen.AddStack(temp, secondTemp)                      // agregar a la pila
		gen.AddExpression(secondTemp, secondTemp, "1", "+") // agregar expresion, para obtener el valor de la expresion izquierda

		gen.AddComment("Add If") // comentario

		EV := gen.NewLabel()                      // Label para el if, para saber si se cumple la condicion
		tree.AddDisplay(Linicio, EV, "-1", false) // Display
		gen.AddIf(thirdTemp, "-1", "==", EV)      // agregar if, para saber si se cumple la condicion

		for _, s := range p.Instrucciones.ToArray() { // se recorren las instrucciones
			s.(interfaces.Instruction).Compilar(&newTable, tree, gen) // se compila cada instruccion, en la tabla de simbolos del for

		}
		pos := fmt.Sprintf("%v", tree.PosDisplay-1) // se obtiene la posicion del display, para obtener el display
		display := tree.GetDisplay(pos)             // se obtiene el display, para obtener el temporal del break

		if display.IsTemp { // si es temporal, se agrega la expresion
			excep := interfaces.NewException("Semantico", "Error en For, Sentencia de Control incorrecta Break.", p.Row, p.Column)          // se crea una nueva excepcion
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) // se agrega a la lista de errores
		}

		gen.AddGoto(Linicio)  // agregar goto, para salir del for
		gen.AddLabel(EV)      // agregar el label, para saber si se cumple la condicion
		tree.RestPosDisplay() // se restaura la posicion del display

	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
