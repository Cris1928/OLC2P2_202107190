package instruction

import (
	"OLC2/Compilador/interfaces"
	"fmt"
	"reflect"

	arrayList "github.com/colegno/arraylist"
)

type Println struct {
	Expression *arrayList.List
	Condicion  interfaces.Expression
	Formato    string
	Row        int
	Column     int
}

func NewPrintln(val *arrayList.List, cond interfaces.Expression, formato string, row int, column int) Println {
	exp := Println{val, cond, formato, row, column}
	return exp
}

func (p Println) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} { //env tendra la tabla de simbolos, tree tendra el arbol de instrucciones, gen tendra el generador de codigo

	// var conca string = "\n"

	var result interfaces.Value
	//fmt.Println("p.condicion=" + fmt.Sprintf("%v", p.Condicion))
	if p.Condicion == nil {
		var condBool bool = false

		newPos := -1
		contExpre := 0
		p.Formato = p.Formato + "{}"
		fmt.Println("p.formato=" + p.Formato)

		for i := 0; i < len(p.Formato); i++ {

			if p.Formato[i] == 123 { //  Caracter ASCII {

				newPos = devolverPos(p.Formato, i)
				if newPos == -2 {
					excep := interfaces.NewException("Semantico", "Formato incorrecto, hace falta \",\".", p.Row, p.Column)
					tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}
				condBool = true
				i = newPos
				if newPos == i && condBool {

					if contExpre < p.Expression.Len() {

						result = p.Expression.GetValue(contExpre).(ListExpre).Compilar(env, tree, gen)

						if result.Type == interfaces.EXCEPTION {
							return result.Value
						}

						if result.Type == interfaces.BOOLEAN {
							gen.AddComment("Printf Boolean")

							if result.IsTemp {

								newLabelEV := gen.NewLabel()
								newLabelEF := gen.NewLabel()

								gen.AddIf(result.Value, "1", "==", newLabelEV)
								gen.AddGoto(newLabelEF)

								newLabel := gen.NewLabel()
								gen.AddBoolean(newLabelEV, newLabelEF, newLabel)

							}

						} else if result.Type == interfaces.INTEGER {
							gen.AddComment("Printf Integer")
							gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))

						} else if result.Type == interfaces.FLOAT {
							gen.AddComment("Printf Float")
							gen.AddPrintf("f", "(double)"+fmt.Sprintf("%v", result.Value))

						} else if result.Type == interfaces.STRING {

							gen.AddComment("Printf String")

							if !tree.IsPrimitive {
								gen.AddPrintfString()
								tree.IsPrimitive = true
							}

							temp := gen.NewTemp()
							gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")
							gen.AddExpression(temp, temp, "1", "+")

							if result.IsTemp { // si es temporal
								gen.AddStack(temp, result.Value)                                  // se agrega a la pila
								gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "+") // se aumenta el puntero
								gen.PrintfString()                                                // se imprime
							}

							temp = gen.NewTemp()
							gen.AddExpressionStack(temp, "P")
							gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "-")

						} else if result.Type == interfaces.CHAR {
							gen.AddComment("Printf Char")
							temp := gen.NewTemp()
							gen.AddExpressionHeap(temp, result.Value)
							gen.AddPrintf("c", "(char)"+temp)
						} else if result.Type == interfaces.NULL {

							temp := gen.NewTemp()
							// t2 = H + 0;
							gen.AddExpression(temp, "H", "0", "+")
							// heap[(int)H] = 78;
							gen.AddHeap("H", "78")
							// H = H + 1;
							gen.AddExpression("H", "H", "1", "+")
							// heap[(int)H] = 85;
							gen.AddHeap("H", "85")
							// H = H + 1;
							gen.AddExpression("H", "H", "1", "+")
							// heap[(int)H] = 76;
							gen.AddHeap("H", "76")
							// H = H + 1;
							gen.AddExpression("H", "H", "1", "+")
							// heap[(int)H] = 79;
							gen.AddHeap("H", "79")
							// H = H + 1;
							gen.AddExpression("H", "H", "1", "+")
							// heap[(int)H] = -1;
							gen.AddHeap("H", "-1")
							// H = H + 1;
							gen.AddExpression("H", "H", "1", "+")
						}

						contExpre++
						condBool = false
					}
				}
			}

			if !condBool && p.Formato[i] != 125 { // ASCII }

				gen.AddComment("Printf ")

				auxTemp := gen.NewTemp()
				gen.AddExpression(auxTemp, "H", "0", "+")

				gen.AddHeap("H", fmt.Sprintf("%v", p.Formato[i]))
				gen.AddExpression("H", "H", "1", "+")
				gen.AddHeap("H", "-1")
				gen.AddExpression("H", "H", "1", "+")

				gen.AddComment("Printf")
				if !tree.IsPrimitive {
					gen.AddPrintfString()
					tree.IsPrimitive = true
				}
				temp := gen.NewTemp()
				gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")
				gen.AddExpression(temp, temp, "1", "+")

				gen.AddStack(temp, auxTemp)
				gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "+")
				gen.PrintfString()

				temp = gen.NewTemp()
				gen.AddExpressionStack(temp, "P")
				gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "-")

			}
		}

		if p.Expression.Len() != contExpre {
			excep := interfaces.NewException("Semantico", "Formato incorrecto, cantidad incorrecta de expreciones, hace falta.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}

	} else {

		result = p.Condicion.Compilar(env, tree, gen)

		if result.Type == interfaces.EXCEPTION {
			return result.Value
		}

		if result.Type == interfaces.STRING || (result.Type == interfaces.STRING && reflect.TypeOf(p.Condicion).String() == "expression.Primitivo") {
			gen.AddComment("Printf String")

			if !tree.IsPrimitive {
				gen.AddPrintfString()
				tree.IsPrimitive = true
			}

			temp := gen.NewTemp()
			gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")
			gen.AddExpression(temp, temp, "1", "+")

			if result.IsTemp {
				gen.AddStack(temp, result.Value)
				gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "+")
				gen.PrintfString()
			}

			temp = gen.NewTemp()
			gen.AddExpressionStack(temp, "P")
			gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "-")
			//}else if result.Type =={

		} else if result.Type == interfaces.INTEGER {
			gen.AddComment("Printf Integer")
			gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
		} else if result.Type == interfaces.FLOAT {
			gen.AddComment("Printf Float")
			gen.AddPrintf("f", "(double)"+fmt.Sprintf("%v", result.Value))
		} else {
			excep := interfaces.NewException("Semantico", "Formato incorrecto, Tipo de Dato Incorrecto.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}
	}

	gen.AddComment("Salto de Linea \\n")
	gen.AddPrintf("c", "10")

	return result.Value
}
func devolverPos(text string, pos int) int {

	var newPos int = -2
	if text[pos] == 123 { // ASCII  -> {
		pos++
	}

	for i := pos; pos < len(text); i++ {

		if text[pos] != 32 { // ASCII  -> espace

			if text[pos] == 125 { // ASCII -> }
				return pos
			} else {
				if text[pos] != 58 && text[pos] != 63 { // ASCCI 58->[:] 63->[?]
					return newPos
				}
			}
		}

		pos++
	}

	return newPos
}
