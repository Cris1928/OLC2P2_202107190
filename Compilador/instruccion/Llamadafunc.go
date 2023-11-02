package instruction

import (
	"OLC2/Compilador/interfaces"

	"fmt"

	"reflect"

	arrayList "github.com/colegno/arraylist"
)

type Llamada struct {
	Id        string
	Parametro *arrayList.List
	Row       int
	Column    int
}

func NewLlamada(Id string, Parametro *arrayList.List, Row int, Column int) Llamada {
	instr := Llamada{Id, Parametro, Row, Column}
	return instr
}

func (p Llamada) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} { //TODO: Verificar si es necesario el return

	gen.AddComment("Llamada de Funcion") //agregamos un comentario
	temp := gen.NewTemp()                //creamos un temporal
	symbol := env.GetFunction(p.Id)      //buscamos la funcion en el entorno

	function := symbol.Value.(interfaces.Symbol).Value.(interfaces.SymbolFunction) //obtenemos la funcion

	if p.Id != function.Id {
		excep := interfaces.NewException("Semantico", "Error en Llamada, No Existe esa Funcion "+p.Id, p.Row, p.Column)                 //creamos una nueva excepcion
		tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column}) //agregamos la excepcion al arbol
		return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}                    //retornamos un value
	}

	if function.Parametro != nil && p.Parametro != nil { //si la funcion tiene parametros y la llamada tambien

		if function.Parametro.Len() == p.Parametro.Len() { //si la cantidad de parametros es igual

			flag := false //creamos una bandera

			for s := 0; s < p.Parametro.Len(); s++ { //recorremos los parametros
				if reflect.TypeOf(p.Parametro.GetValue(s).(ListExpre).Expresion).String() == "function.LlamadaExpresion" { //si el tipo de la expresion es una llamada
					flag = true //cambiamos la bandera
					break       //salimos del ciclo
				}
			}

			if !flag { //si la bandera es falsa
				for s := 0; s < p.Parametro.Len(); s++ { //recorremos los parametros

					instrCall := p.Parametro.GetValue(s).(ListExpre).Compilar(env, tree, gen) //compilamos la expresion
					instrFunc := function.Parametro.GetValue(s).(ListExprefunc)               //obtenemos el parametro

					if instrCall.Type == instrFunc.Type { //si el tipo de la expresion es igual al tipo del parametro
						if s == 0 { //si es el primer parametro
							gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion+1), "+") //agregamos una expresion
							gen.AddStack(temp, instrCall.Value)                                  //agregamos al stack

						} else {
							gen.AddExpression(temp, temp, "1", "+") //agregamos una expresion
							gen.AddStack(temp, instrCall.Value)     //agregamos al stack

						}

					}

				}
			} else {

				var saveTemps []interface{} //creamos un arreglo de interfaces

				for s := 0; s < p.Parametro.Len(); s++ { //recorremos los parametros

					instrCall := p.Parametro.GetValue(s).(ListExpre).Compilar(env, tree, gen) //compilamos la expresion
					instrFunc := function.Parametro.GetValue(s).(ListExprefunc)               //obtenemos el parametro

					if instrCall.Type == instrFunc.Type { //si el tipo de la expresion es igual al tipo del parametro
						auxTemp := gen.NewTemp()                                                                                                                      //creamos un temporal
						gen.AddComment("Guardando Temporal")                                                                                                          //agregamos un comentario
						gen.AddExpression(auxTemp, "P", fmt.Sprintf("%v", env.Posicion), "+")                                                                         //agregamos una expresion
						gen.AddStack(auxTemp, instrCall.Value)                                                                                                        //agregamos al stack
						tree.AddTableSymbol(*interfaces.NewTableSymbol(instrCall.Value, "Temporal", "Local", p.Row, p.Column, "--", fmt.Sprintf("%v", env.Posicion))) //agregamos a la tabla de simbolos
						saveTemps = append(saveTemps, env.Posicion)                                                                                                   //agregamos la posicion al arreglo
						env.NewPos()                                                                                                                                  //nueva posicion

					}

				}

				for s := 0; s < p.Parametro.Len(); s++ { //recorremos los parametros
					gen.AddComment("Llamado de Temporales") //agregamos un comentario

					if s == 0 {
						gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion+1), "+")  //agregamos una expresion
						tempNew := gen.NewTemp()                                              //creamos un temporal
						tempStack := gen.NewTemp()                                            //creamos un temporal
						gen.AddExpression(tempNew, "P", fmt.Sprintf("%v", saveTemps[s]), "+") //agregamos una expresion
						gen.AddExpressionStack(tempStack, tempNew)                            //agregamos una expresion
						gen.AddStack(temp, tempStack)                                         //agregamos al stack

					} else {
						gen.AddExpression(temp, temp, "1", "+")                               //agregamos una expresion
						tempNew := gen.NewTemp()                                              //creamos un temporal
						tempStack := gen.NewTemp()                                            //creamos un temporal
						gen.AddExpression(tempNew, "P", fmt.Sprintf("%v", saveTemps[s]), "+") //agregamos una expresion
						gen.AddExpressionStack(tempStack, tempNew)                            //agregamos una expresion
						gen.AddStack(temp, tempStack)                                         //agregamos al stack

					}

				}

			}

			gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "+") //agregamos una expresion
			gen.PrintFunction(p.Id)                                           //imprimimos la funcion
			gen.AddExpressionStack(temp, "P")                                 //agregamos una expresion
			gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "-") //agregamos una expresion

		} else {
			excep := interfaces.NewException("Semantico", "Error en Llamada, cantidad de Parametros incorrecto.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
		}

	} else {

		if (function.Parametro != nil && p.Parametro == nil) || (function.Parametro == nil && p.Parametro != nil) {
			excep := interfaces.NewException("Semantico", "Error en Llamada, parametros incorrectos.", p.Row, p.Column)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
			return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

		} else {
			gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "+")
			gen.PrintFunction(p.Id)
			gen.AddExpressionStack(temp, "P")
			gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "-")

		}

	}

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
