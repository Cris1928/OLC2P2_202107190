package expression

import (
	"OLC2/Compilador/interfaces"
	"fmt"

	// "math"
	// "strconv"
	"reflect"
)

type Aritmetica struct {
	left     interfaces.Expression
	Operator string
	right    interfaces.Expression
	Unario   bool
	Row      int
	Column   int
}

func NewOperacion(left interfaces.Expression, Operator string, right interfaces.Expression, unario bool, row int, column int) Aritmetica {
	exp := Aritmetica{left, Operator, right, unario, row, column}
	return exp
}
func (p Aritmetica) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interfaces.Value {
	var exp_left interfaces.Value
	var exp_right interfaces.Value

	switch p.Operator {
	case "+":
		{
			if p.Unario { // si es unario solo se compila el lado izquierdo
				exp_left = p.left.Compilar(env, tree, gen) // se compila el lado izquierdo
			} else { // si no es unario se compila el lado izquierdo y derecho
				if reflect.TypeOf(p.right).String() == "function.LlamadaExpresion" { // si el lado derecho es una llamada a funcion
					exp_left = p.left.Compilar(env, tree, gen)                                                                                                   // se compila el lado izquierdo
					temp := gen.NewTemp()                                                                                                                        // se crea un temporal
					gen.AddComment("Guardando Temporal")                                                                                                         // se agrega un comentario
					gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")                                                                           // se agrega la expresion
					aux := env.Posicion                                                                                                                          // se guarda la posicion
					gen.AddStack(temp, exp_left.Value)                                                                                                           // se agrega el valor del temporal al stack
					tree.AddTableSymbol(*interfaces.NewTableSymbol(exp_left.Value, "Temporal", "Local", p.Row, p.Column, "--", fmt.Sprintf("%v", env.Posicion))) // se agrega el temporal a la tabla de simbolos
					env.NewPos()                                                                                                                                 // se crea una nueva posicion

					exp_right = p.right.Compilar(env, tree, gen) // se compila el lado derecho

					temp = gen.NewTemp()                                      // se crea un temporal
					temp1 := gen.NewTemp()                                    // se crea un temporal
					gen.AddExpression(temp, "P", fmt.Sprintf("%v", aux), "+") // se agrega la expresion
					gen.AddExpressionStack(temp1, temp)                       // se agrega la expresion
					exp_left.Value = temp1                                    // se guarda el valor del temporal en el valor del lado izquierdo
				} else { // si no es una llamada a funcion
					exp_left = p.left.Compilar(env, tree, gen)   // se compila el lado izquierdo
					exp_right = p.right.Compilar(env, tree, gen) // se compila el lado derecho
				}
			}
			gen.AddComment("Aritmetica +")
			temp := gen.NewTemp()
			isType := interfaces.NULL
			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "+")
				isType = interfaces.INTEGER

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "+")
				isType = interfaces.FLOAT

			} else if exp_left.Type == interfaces.STRING && exp_right.Type == interfaces.STRING {
				if !tree.IsCocant {
					gen.AddConcatString()
					tree.IsCocant = true
				}
				gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")
				gen.AddExpression(temp, temp, "1", "+")
				gen.AddStack(temp, exp_left.Value)
				gen.AddExpression(temp, temp, "1", "+")
				gen.AddStack(temp, exp_right.Value)
				gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "+")
				gen.ConcatString()
				temp = gen.NewTemp()
				gen.AddExpressionStack(temp, "P")
				gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "-")

				isType = interfaces.STRING
			} else {
				excep := interfaces.NewException("Semantico", "No es posible Sumar, Tipos de datos Incorrectos.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}
			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}
	case "-":
		{
			temp := gen.NewTemp()
			isType := interfaces.NULL

			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
				gen.AddComment("Aritmetica -")
				/* ************************************************************** INTEGER ************************************************************** */
				if exp_left.Type == interfaces.INTEGER {
					gen.AddExpression(temp, "0", exp_left.Value, "-")
					isType = interfaces.INTEGER
					/* ************************************************************** FLOAT ************************************************************** */
				} else if exp_left.Type == interfaces.FLOAT {
					gen.AddExpression(temp, "0", exp_left.Value, "-")
					isType = interfaces.FLOAT

				} else {
					excep := interfaces.NewException("Semantico", "No es posible Restar, Tipos de datos Incorrectos.", p.Row, p.Column)
					tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}

			} else {

				if reflect.TypeOf(p.right).String() == "function.LlamadaExpresion" {
					exp_left = p.left.Compilar(env, tree, gen)
					temp := gen.NewTemp()
					gen.AddComment("Guardando Temporal")
					gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")
					aux := env.Posicion
					gen.AddStack(temp, exp_left.Value)
					tree.AddTableSymbol(*interfaces.NewTableSymbol(exp_left.Value, "Temporal", "Local", p.Row, p.Column, "--", fmt.Sprintf("%v", env.Posicion)))

					env.NewPos()

					exp_right = p.right.Compilar(env, tree, gen)

					temp = gen.NewTemp()
					temp1 := gen.NewTemp()
					gen.AddExpression(temp, "P", fmt.Sprintf("%v", aux), "+")
					gen.AddExpressionStack(temp1, temp)
					exp_left.Value = temp1
				} else {
					exp_left = p.left.Compilar(env, tree, gen)
					exp_right = p.right.Compilar(env, tree, gen)
				}

				gen.AddComment("Aritmetica -")

				/* ************************************************************** INTEGER ************************************************************** */
				if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
					gen.AddExpression(temp, exp_left.Value, exp_right.Value, "-")
					isType = interfaces.INTEGER
					/* ************************************************************** FLOAT ************************************************************** */
				} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
					gen.AddExpression(temp, exp_left.Value, exp_right.Value, "-")
					isType = interfaces.FLOAT

				} else {
					excep := interfaces.NewException("Semantico", "No es posible Restar, Tipos de datos Incorrectos.", p.Row, p.Column)
					tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}

			}

			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}

	case "*":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {

				if reflect.TypeOf(p.right).String() == "function.LlamadaExpresion" {

					exp_left = p.left.Compilar(env, tree, gen)
					gen.AddComment("Guardando Temporal")
					temp := gen.NewTemp()
					gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")
					aux := env.Posicion
					gen.AddStack(temp, exp_left.Value)
					tree.AddTableSymbol(*interfaces.NewTableSymbol(exp_left.Value, "Temporal", "Local", p.Row, p.Column, "--", fmt.Sprintf("%v", env.Posicion)))

					env.NewPos()

					exp_right = p.right.Compilar(env, tree, gen)

					temp = gen.NewTemp()
					temp1 := gen.NewTemp()
					gen.AddExpression(temp, "P", fmt.Sprintf("%v", aux), "+")
					gen.AddExpressionStack(temp1, temp)
					exp_left.Value = temp1

				} else {
					exp_left = p.left.Compilar(env, tree, gen)
					exp_right = p.right.Compilar(env, tree, gen)
				}

			}

			gen.AddComment("Aritmetica *")

			temp := gen.NewTemp()
			isType := interfaces.NULL
			/* ************************************************************** INTEGER ************************************************************** */
			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "*")
				isType = interfaces.INTEGER
				/* ************************************************************** FLOAT ************************************************************** */
			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "*")
				isType = interfaces.FLOAT

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Multiplicar, Tipos de datos Incorrectos.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}
	case "%":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}

			gen.AddComment("Aritmetica %")

			temp := gen.NewTemp()
			isType := interfaces.NULL
			/* ************************************************************** INTEGER ************************************************************** */
			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
				val := "fmod(" + fmt.Sprintf("%v", exp_left.Value) + "," + fmt.Sprintf("%v", exp_right.Value) + ")"
				gen.AddExpression(temp, val, "0", "+")
				isType = interfaces.INTEGER
				/* ************************************************************** FLOAT ************************************************************** */
			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
				val := "fmod(" + fmt.Sprintf("%v", exp_left.Value) + "," + fmt.Sprintf("%v", exp_right.Value) + ")"
				gen.AddExpression(temp, val, "0", "+")
				isType = interfaces.FLOAT

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Multiplicar, Tipos de datos Incorrectos.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}

	case "/":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Aritmetica °|°")

			if !exp_right.IsTemp {
				if exp_right.Value == "0" {
					excep := interfaces.NewException("Semantico", "No es posible Dividir entre 0.", p.Row, p.Column)
					tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
				}
			}
			temp := gen.NewTemp()
			isType := interfaces.NULL
			/* ************************************************************** INTEGER ************************************************************** */
			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "/")
				isType = interfaces.INTEGER
				/* ************************************************************** FLOAT ************************************************************** */
			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {
				gen.AddExpression(temp, exp_left.Value, exp_right.Value, "/")
				isType = interfaces.FLOAT

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Dividir.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			return interfaces.Value{Value: temp, IsTemp: true, Type: isType, TrueLabel: "", FalseLabel: ""}
		}

	case "<":
		{

			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional <")
			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, "<", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, "<", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.BOOLEAN && exp_right.Type == interfaces.BOOLEAN {

				gen.AddIf(exp_left.Value, exp_right.Value, "<", EV)
				gen.AddGoto(EF)

			} else if (exp_left.Type == interfaces.STRING && exp_right.Type == interfaces.STRING) || (exp_left.Type == interfaces.CHAR && exp_right.Type == interfaces.CHAR) {
				_left := gen.NewTemp()
				_right := gen.NewTemp()
				gen.AddExpressionHeap(_left, exp_left.Value)
				gen.AddExpressionHeap(_right, exp_right.Value)
				gen.AddIf(_left, _right, "<", EV)
				gen.AddGoto(EF)

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Comparar <.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			newLabel := gen.NewLabel()
			newTemp := gen.NewTemp()
			gen.Boolean(EV, EF, newLabel, newTemp)

			return interfaces.Value{Value: newTemp, IsTemp: true, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}

		}
	case ">":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional >")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, ">", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, ">", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.BOOLEAN && exp_right.Type == interfaces.BOOLEAN {

				gen.AddIf(exp_left.Value, exp_right.Value, ">", EV)
				gen.AddGoto(EF)

			} else if (exp_left.Type == interfaces.STRING && exp_right.Type == interfaces.STRING) || (exp_left.Type == interfaces.CHAR && exp_right.Type == interfaces.CHAR) {
				_left := gen.NewTemp()
				_right := gen.NewTemp()
				gen.AddExpressionHeap(_left, exp_left.Value)
				gen.AddExpressionHeap(_right, exp_right.Value)
				gen.AddIf(_left, _right, ">", EV)
				gen.AddGoto(EF)

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Comparar >.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			newLabel := gen.NewLabel()
			newTemp := gen.NewTemp()
			gen.Boolean(EV, EF, newLabel, newTemp)

			return interfaces.Value{Value: newTemp, IsTemp: true, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}

		}

	case "<=":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional <=")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, "<=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, "<=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.BOOLEAN && exp_right.Type == interfaces.BOOLEAN {

				gen.AddIf(exp_left.Value, exp_right.Value, "<=", EV)
				gen.AddGoto(EF)

			} else if (exp_left.Type == interfaces.STRING && exp_right.Type == interfaces.STRING) || (exp_left.Type == interfaces.CHAR && exp_right.Type == interfaces.CHAR) {
				_left := gen.NewTemp()
				_right := gen.NewTemp()
				gen.AddExpressionHeap(_left, exp_left.Value)
				gen.AddExpressionHeap(_right, exp_right.Value)
				gen.AddIf(_left, _right, "<=", EV)
				gen.AddGoto(EF)

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Comparar <=.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			newLabel := gen.NewLabel()
			newTemp := gen.NewTemp()
			gen.Boolean(EV, EF, newLabel, newTemp)

			return interfaces.Value{Value: newTemp, IsTemp: true, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}

		}
	case ">=":
		{
			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional >=")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, ">=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, ">=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.BOOLEAN && exp_right.Type == interfaces.BOOLEAN {

				gen.AddIf(exp_left.Value, exp_right.Value, ">=", EV)
				gen.AddGoto(EF)

			} else if (exp_left.Type == interfaces.STRING && exp_right.Type == interfaces.STRING) || (exp_left.Type == interfaces.CHAR && exp_right.Type == interfaces.CHAR) {
				_left := gen.NewTemp()
				_right := gen.NewTemp()
				gen.AddExpressionHeap(_left, exp_left.Value)
				gen.AddExpressionHeap(_right, exp_right.Value)
				gen.AddIf(_left, _right, ">=", EV)
				gen.AddGoto(EF)

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Comparar >=.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			newLabel := gen.NewLabel()
			newTemp := gen.NewTemp()
			gen.Boolean(EV, EF, newLabel, newTemp)

			return interfaces.Value{Value: newTemp, IsTemp: true, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}

		}
	case "==":
		{

			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional ==")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, "==", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, "==", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.STRING && exp_right.Type == interfaces.STRING {

				if !tree.IsCompareStr {
					gen.AddCompareString()
					tree.IsCompareStr = true
				}

				temp := gen.NewTemp()
				gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")
				gen.AddExpression(temp, temp, "1", "+")
				gen.AddStack(temp, exp_left.Value)
				gen.AddExpression(temp, temp, "1", "+")
				gen.AddStack(temp, exp_right.Value)
				gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "+")
				gen.CompareString()
				temp = gen.NewTemp()
				gen.AddExpressionStack(temp, "P")
				gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "-")

				gen.AddIf(temp, "1", "==", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.CHAR && exp_right.Type == interfaces.CHAR {
				_left := gen.NewTemp()
				_right := gen.NewTemp()
				gen.AddExpressionHeap(_left, exp_left.Value)
				gen.AddExpressionHeap(_right, exp_right.Value)
				gen.AddIf(_left, _right, "==", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.BOOLEAN && exp_right.Type == interfaces.BOOLEAN {

				gen.AddIf(exp_left.Value, exp_right.Value, "==", EV)
				gen.AddGoto(EF)

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Comparar ==.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			newLabel := gen.NewLabel()
			newTemp := gen.NewTemp()
			gen.Boolean(EV, EF, newLabel, newTemp)

			return interfaces.Value{Value: newTemp, IsTemp: true, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}

		}
	case "!=":
		{

			if p.Unario == true {
				exp_left = p.left.Compilar(env, tree, gen)
			} else {
				exp_left = p.left.Compilar(env, tree, gen)
				exp_right = p.right.Compilar(env, tree, gen)
			}
			gen.AddComment("Relacional !=")

			EV := gen.NewLabel()
			EF := gen.NewLabel()

			if exp_left.Type == interfaces.INTEGER && exp_right.Type == interfaces.INTEGER {

				gen.AddIf(exp_left.Value, exp_right.Value, "!=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.FLOAT && exp_right.Type == interfaces.FLOAT {

				gen.AddIf(exp_left.Value, exp_right.Value, "!=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.STRING && exp_right.Type == interfaces.STRING {

				if !tree.IsCompareStr {
					gen.AddCompareString()
					tree.IsCompareStr = true
				}

				temp := gen.NewTemp()
				gen.AddExpression(temp, "P", fmt.Sprintf("%v", env.Posicion), "+")
				gen.AddExpression(temp, temp, "1", "+")
				gen.AddStack(temp, exp_left.Value)
				gen.AddExpression(temp, temp, "1", "+")
				gen.AddStack(temp, exp_right.Value)
				gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "+")
				gen.CompareString()
				temp = gen.NewTemp()
				gen.AddExpressionStack(temp, "P")
				gen.AddExpression("P", "P", fmt.Sprintf("%v", env.Posicion), "-")

				gen.AddIf(temp, "0", "==", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.CHAR && exp_right.Type == interfaces.CHAR {
				_left := gen.NewTemp()
				_right := gen.NewTemp()
				gen.AddExpressionHeap(_left, exp_left.Value)
				gen.AddExpressionHeap(_right, exp_right.Value)
				gen.AddIf(_left, _right, "!=", EV)
				gen.AddGoto(EF)

			} else if exp_left.Type == interfaces.BOOLEAN && exp_right.Type == interfaces.BOOLEAN {

				gen.AddIf(exp_left.Value, exp_right.Value, "!=", EV)
				gen.AddGoto(EF)

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Comparar !=.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

			newLabel := gen.NewLabel()
			newTemp := gen.NewTemp()
			gen.Boolean(EV, EF, newLabel, newTemp)

			return interfaces.Value{Value: newTemp, IsTemp: true, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}

		}
	case "&&":
		{

			exp_left = p.left.Compilar(env, tree, gen)
			exp_right := p.right.Compilar(env, tree, gen)
			gen.AddComment("Logico &&")
			if exp_left.Type == interfaces.BOOLEAN {

				if exp_right.Type == interfaces.BOOLEAN {
					temp := gen.NewTemp()
					EV1 := gen.NewLabel()      // L0
					EF1 := gen.NewLabel()      // L1
					EV2 := gen.NewLabel()      // L2
					EF2 := gen.NewLabel()      // L3
					newLabel := gen.NewLabel() // L4

					gen.AddIf(exp_left.Value, "1", "==", EV1) // L0
					gen.AddGoto(EF1)                          // L1

					gen.AddLabel(EV1) // L0

					gen.AddIf(exp_right.Value, "1", "==", EV2) // L2
					gen.AddGoto(EF2)                           // L3

					gen.AddLabel(EV2) // L1
					gen.AddExpression(temp, "1", "0", "+")
					gen.AddGoto(newLabel)

					gen.AddLabel(EF1) // L1
					gen.AddExpression(temp, "0", "0", "+")
					gen.AddGoto(newLabel)

					gen.AddLabel(EF2) // L3
					gen.AddExpression(temp, "0", "0", "+")
					gen.AddGoto(newLabel)

					gen.AddLabel(newLabel) // L3

					return interfaces.Value{Value: temp, IsTemp: true, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}

				} else {
					excep := interfaces.NewException("Semantico", "No es posible Comparar &&; Tipo de Dato no Booleano en expresion Der.", p.Row, p.Column)
					tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

				}

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Comparar &&; Tipo de Dato no Booleano en expresion Izq.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

		}
	case "||":
		{

			exp_left = p.left.Compilar(env, tree, gen)
			exp_right := p.right.Compilar(env, tree, gen)
			gen.AddComment("Logico ||")
			if exp_left.Type == interfaces.BOOLEAN {

				if exp_right.Type == interfaces.BOOLEAN {

					temp := gen.NewTemp()
					EV1 := gen.NewLabel()      // L0
					EF1 := gen.NewLabel()      // L1
					EV2 := gen.NewLabel()      // L2
					EF2 := gen.NewLabel()      // L3
					newLabel := gen.NewLabel() // L4

					gen.AddIf(exp_left.Value, "1", "==", EV1) // L0
					gen.AddGoto(EF1)                          // L1

					gen.AddLabel(EF1) // L1

					gen.AddIf(exp_right.Value, "1", "==", EV2) // L2
					gen.AddGoto(EF2)                           // L3

					gen.AddLabel(EV1) // L0
					gen.AddExpression(temp, "1", "0", "+")
					gen.AddGoto(newLabel)

					gen.AddLabel(EV2) // L2
					gen.AddExpression(temp, "1", "0", "+")
					gen.AddGoto(newLabel)

					gen.AddLabel(EF2) // L3
					gen.AddExpression(temp, "0", "0", "+")
					gen.AddGoto(newLabel)

					gen.AddLabel(newLabel) // L3

					return interfaces.Value{Value: temp, IsTemp: true, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}

				} else {
					excep := interfaces.NewException("Semantico", "No es posible Comparar ||; Tipo de Dato no Booleano en expresion Der.", p.Row, p.Column)
					tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
					return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}

				}

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Comparar ||; Tipo de Dato no Booleano en expresion Izq.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

		}
	case "!":
		{
			if p.Unario {
				exp_left = p.left.Compilar(env, tree, gen)
			}

			gen.AddComment("Logico !")
			if exp_left.Type == interfaces.BOOLEAN {

				temp := gen.NewTemp()
				EV := gen.NewLabel()       // L0
				EF := gen.NewLabel()       // L1
				newLabel := gen.NewLabel() // L2

				gen.AddIf(exp_left.Value, "1", "==", EV) // L0
				gen.AddGoto(EF)                          // L1

				gen.AddLabel(EV) // L0
				gen.AddExpression(temp, "0", "0", "+")
				gen.AddGoto(newLabel)

				gen.AddLabel(EF) // L1
				gen.AddExpression(temp, "1", "0", "+")
				gen.AddGoto(newLabel)

				gen.AddLabel(newLabel) // L32

				return interfaces.Value{Value: temp, IsTemp: true, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}

			} else {
				excep := interfaces.NewException("Semantico", "No es posible Comparar ||; Tipo de Dato no Booleano.", p.Row, p.Column)
				tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Column})
				return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.EXCEPTION, TrueLabel: "", FalseLabel: ""}
			}

		}
	}
	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
