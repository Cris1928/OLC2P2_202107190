package interfaces

import (
	"fmt"
	// "reflect"
)

type Environment struct {
	anterior *Environment      //entorno anterior
	variable map[string]Symbol //variables
	Function map[string]Symbol //funciones
	Structs  map[string]Symbol //structs

	Posicion int //posicion
}

func NewEnvironment(anterior *Environment) Environment { // funcion de nuevo enviroment

	env := Environment{anterior, make(map[string]Symbol), make(map[string]Symbol), make(map[string]Symbol), 0} //creamos un nuevo entorno
	return env
}
func (env *Environment) IsAmbit() bool { //funcion para saber si es ambito

	var tmpEnv *Environment //variable temporal para recorrer el entorno
	tmpEnv = env            //asignamos el entorno actual a la variable temporal

	cont := 0 //contador para saber si es ambito

	for { //recorremos el entorno

		if tmpEnv.anterior == nil { //si el entorno actual es el global
			break //rompemos el ciclo
		} else { //si no es el global
			tmpEnv = tmpEnv.anterior //asignamos el entorno anterior a la variable temporal
			cont++                   //aumentamos el contador
		}
	}

	return 0 < cont //si el contador es mayor a 0 retornamos true, si no false
}
func (env *Environment) NewPos() { //funcion para aumentar la posicion
	env.Posicion = env.Posicion + 1 //aumentamos la posicion
}

func (env *Environment) UpdatePos(treePos int, envPos int, isBool bool, newTable *Environment) { //funcion para actualizar la posicion

	newTable.Posicion = envPos //actualizamos la posicion

}

func (env *Environment) AddSymbol(id string, value Value, tipo TypeExpression, isMut bool, pos int, newTable *Environment) { //funcion para agregar simbolo

	if variable, ok := newTable.variable[id]; ok { //si la variable existe en el entorno actual
		fmt.Println("[ADD SYMBOL] La variable " + variable.Id + " ya existe") //imprimimos que la variable ya existe
		return                                                                //retornamos
	}
	newTable.variable[id] = Symbol{Id: id, Type: tipo, Value: value, IsMut: isMut, Posicion: pos} //agregamos la variable al entorno actual
}

func (env *Environment) SearchSymbol(id string) Symbol { //funcion para buscar simbolo
	var tmpEnv *Environment //variable temporal para recorrer el entorno
	tmpEnv = env            //asignamos el entorno actual a la variable temporal

	for { //recorremos el entorno

		if tmpEnv.anterior == nil { //si el entorno actual es el global
			break //rompemos el ciclo
		} else { //si no es el global
			tmpEnv = tmpEnv.anterior //asignamos el entorno anterior a la variable temporal
		}
	}
	if variable, ok := tmpEnv.variable[id]; ok { //si la variable existe en el entorno actual
		return variable //retornamos la variable
	}
	return Symbol{Id: "", Type: NULL, Value: Symbol{Id: "", Type: NULL, Value: 0}}
}

func (env *Environment) GetSymbol(id string) Symbol { //funcion para buscar simbolo

	var tmpEnv *Environment //variable temporal para recorrer el entorno
	tmpEnv = env            //asignamos el entorno actual a la variable temporal

	for {
		if variable, ok := tmpEnv.variable[id]; ok { //si la variable existe en el entorno actual
			return variable //retornamos la variable
		}

		if tmpEnv.anterior == nil { //si el entorno actual es el global
			break
		} else { //si no es el global
			tmpEnv = tmpEnv.anterior //asignamos el entorno anterior a la variable temporal
		}
	}

	return Symbol{Id: "", Type: NULL, Value: Symbol{Id: "", Type: NULL, Value: 0}} //si no se encuentra la variable retornamos un simbolo vacio
}

func (env *Environment) SetSymbol(id string, value Value, mut bool, pos int) Symbol { //funcion para buscar simbolo

	var tmpEnv *Environment //variable temporal para recorrer el entorno
	tmpEnv = env            //asignamos el entorno actual a la variable temporal

	for { //recorremos el entorno
		if variable, ok := tmpEnv.variable[id]; ok { //si la variable existe en el entorno actual
			tmpEnv.variable[id] = Symbol{Id: id, Type: variable.Type, Value: value, IsMut: mut, Posicion: pos} //actualizamos la variable
			return variable                                                                                    //retornamos la variable
		}

		if tmpEnv.anterior == nil { //si el entorno actual es el global
			break //rompemos el ciclo
		} else { //si no es el global
			tmpEnv = tmpEnv.anterior //asignamos el entorno anterior a la variable temporal
		}
	}

	return Symbol{Id: "", Type: NULL, Value: Symbol{Id: "", Type: NULL, Value: 0}} //si no se encuentra la variable retornamos un simbolo vacio
}
func (env *Environment) AddFunction(id string, value Symbol, tipo TypeExpression) { //funcion para agregar simbolo

	if variable, ok := env.Function[id]; ok { //si la variable existe en el entorno actual
		fmt.Println("[ADD SYMBOL] La variable " + variable.Id + " ya existe") //imprimimos que la variable ya existe
		return                                                                //retornamos
	}
	env.Function[id] = Symbol{Id: id, Type: tipo, Value: value, IsMut: true, Posicion: 0} //agregamos la variable al entorno actual

}

func (env *Environment) GetFunction(id string) Symbol { //funcion para buscar simbolo

	var tmpEnv *Environment //variable temporal para recorrer el entorno
	tmpEnv = env            //asignamos el entorno actual a la variable temporal

	for { //recorremos el entorno

		if variable, ok := tmpEnv.Function[id]; ok { //si la variable existe en el entorno actual
			return variable //retornamos la variable
		}

		if tmpEnv.anterior == nil { //si el entorno actual es el global
			break //rompemos el ciclo

		} else { //si no es el global
			tmpEnv = tmpEnv.anterior //asignamos el entorno anterior a la variable temporal
		}
	}

	return Symbol{Id: "", Type: NULL, Value: Symbol{Id: "", Type: NULL, Value: 0}} //si no se encuentra la variable retornamos un simbolo vacio

}

func (env *Environment) AddStructs(id string, value Symbol, tipo TypeExpression) { //funcion para agregar simbolo

	if variable, ok := env.Structs[id]; ok { //si la variable existe en el entorno actual
		fmt.Println("[ADD SYMBOL] La variable " + variable.Id + " ya existe") //imprimimos que la variable ya existe
		return                                                                //retornamos
	}
	env.Structs[id] = Symbol{Id: id, Type: tipo, Value: value, IsMut: true, Posicion: 0} //agregamos la variable al entorno actual

}

func (env *Environment) AddSymbolStruct(id string, value Symbol, tipo TypeExpression, isMut bool, pos int, newTable *Environment) { //funcion para agregar simbolo

	if variable, ok := newTable.variable[id]; ok { //si la variable existe en el entorno actual
		fmt.Println("[ADD SYMBOL] La variable " + variable.Id + " ya existe") //imprimimos que la variable ya existe
		return                                                                //retornamos
	}
	newTable.variable[id] = Symbol{Id: id, Type: tipo, Value: value, IsMut: isMut, Posicion: pos} //agregamos la variable al entorno actual
}

func (env *Environment) GetStructs(id string) Symbol { //funcion para buscar simbolo

	var tmpEnv *Environment //variable temporal para recorrer el entorno
	tmpEnv = env            //asignamos el entorno actual a la variable temporal

	for {

		if variable, ok := tmpEnv.Structs[id]; ok { //si la variable existe en el entorno actual
			return variable //retornamos la variable
		}

		if tmpEnv.anterior == nil { //si el entorno actual es el global
			break //rompemos el ciclo

		} else {
			tmpEnv = tmpEnv.anterior //asignamos el entorno anterior a la variable temporal
		}
	}

	return Symbol{Id: "", Type: NULL, Value: Symbol{Id: "", Type: NULL, Value: 0}} //si no se encuentra la variable retornamos un simbolo vacio

}

/* ADD SYMBOL ARRAY*/
func (env *Environment) AddSymbolArrays(id string, value Symbol, tipo TypeExpression, isMut bool, pos int, newTable *Environment) { //funcion para agregar simbolo

	if variable, ok := newTable.variable[id]; ok { //si la variable existe en el entorno actual
		fmt.Println("[ADD SYMBOL] La variable " + variable.Id + " ya existe") //imprimimos que la variable ya existe
		return                                                                //retornamos
	}
	newTable.variable[id] = Symbol{Id: id, Type: tipo, Value: value, IsMut: isMut, Posicion: pos} //agregamos la variable al entorno actual
}
