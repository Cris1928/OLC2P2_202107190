package interfaces

import (
	"fmt"
	"time"

	arrayList "github.com/colegno/arraylist"
)

type Arbol struct { //estructura del arbol
	Code         *arrayList.List    //lista de instrucciones
	StackGlobal  int                //stack global
	IsPrimitive  bool               //is primitive
	IsCocant     bool               //is cocant
	IsCompareStr bool               //is compare str
	_Exception   *arrayList.List    //lista de excepciones
	Display      map[string]Display //display
	PosDisplay   int                //posicion del display
	IsReturn     bool               //is return
	PosReturn    int                //posicion del return
	_TableSymbol *arrayList.List    //tabla de simbolos
	_TableDB     *arrayList.List    //tabla de bases de datos
}

type Exception struct { //estructura de la excepcion
	Tipo        string //tipo de la excepcion
	Descripcion string //descripcion de la excepcion
	Row         int    //fila de la excepcion
	Column      int    //columna de la excepcion
	Time        string //hora de la excepcion
}

type TableSymbol struct { //estructura de la tabla de simbolos
	Name     string //nombre de la variable
	Tipo     string //tipo de la variable
	Ambito   string //ambito de la variable
	Row      int    //fila de la variable
	Column   int    //columna de la variable
	Size     string //size de la variable
	Posicion string //posicion de la variable
}

type TableDB struct { //estructura de la tabla de simbolos
	Name   string //nombre de la variable
	Row    int    //fila de la variable
	Column int    //columna de la variable
}

type Display struct { //estructura del display
	Temp    string         //temporal
	IsTemp  bool           //is temporal
	LInicio string         //label inicio
	LFinal  string         //label final
	Type    TypeExpression //tipo de expresion
}

func NewArbol() *Arbol { //funcion nuevo arbol
	tree := Arbol{ //se crea el arbol
		Code:         arrayList.New(),          //se crea la lista de instrucciones
		StackGlobal:  0,                        //se crea el stack global
		IsPrimitive:  false,                    //se crea el is primitive
		IsCocant:     false,                    //se crea el is cocant
		IsCompareStr: false,                    //se crea el is compare str
		_Exception:   arrayList.New(),          //se crea la lista de excepciones
		Display:      make(map[string]Display), //se crea el display
		PosDisplay:   0,                        //se crea la posicion del display
		IsReturn:     false,                    //se crea el is return
		PosReturn:    -1,                       //se crea la posicion del return
		_TableSymbol: arrayList.New(),          //se crea la tabla de simbolos
		_TableDB:     arrayList.New(),          //se crea la tabla de bases de datos
	}
	return &tree
}

func NewException(tipo string, descripcion string, row int, column int) *Exception { //funcion nuevo arbol
	t := time.Now()                                    //obtiene la hora actual
	hora := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", //se crea el formato de la hora
		t.Year(), t.Month(), t.Day(), //se crea el formato de la hora
		t.Hour(), t.Minute(), t.Second()) //se crea el formato de la hora

	e := Exception{Tipo: tipo, Descripcion: descripcion, Row: row, Column: column, Time: hora} //se crea la excepcion
	return &e                                                                                  //se retorna la excepcion
}

func (a *Arbol) AddCode(input string) { //funcion para agregar codigo
	a.Code.Add(input) //se agrega el codigo
}

func (a Arbol) GetCode() *arrayList.List { //funcion para obtener el codigo
	return a.Code //se retorna el codigo
}

func (a *Arbol) AddException(e Exception) { //funcion para agregar excepcion
	t := time.Now()                                    //obtiene la hora actual
	hora := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d", //se crea el formato de la hora
		t.Day(), t.Month(), t.Year(), //se crea el formato de la hora
		t.Hour(), t.Minute(), t.Second()) //se crea el formato de la hora
	e.Time = hora       //se agrega la hora a la excepcion
	a._Exception.Add(e) //se agrega la excepcion
}

func (a Arbol) GetException() *arrayList.List { //funcion para obtener las excepciones
	return a._Exception //se retorna las excepciones
}

func (a Arbol) GetPos() int { //funcion para obtener la posicion del stack global
	return a.StackGlobal //se retorna la posicion del stack global
}

func (a *Arbol) NewPos() { //funcion para obtener la posicion del stack global
	a.StackGlobal = a.StackGlobal + 1 //se aumenta la posicion del stack global
}

func (a *Arbol) AddDisplay(labelInicio string, labelFinal string, temp string, isTemp bool) { //funcion para agregar display

	pos := fmt.Sprintf("%v", a.PosDisplay)                                                                     //se crea la posicion del display
	a.Display[pos] = Display{Temp: temp, IsTemp: isTemp, LInicio: labelInicio, LFinal: labelFinal, Type: NULL} //se agrega el display
	a.PosDisplay = a.PosDisplay + 1                                                                            //se aumenta la posicion del display
}

func (a *Arbol) SetDisplayTemp(pos string, temp string, isTemp bool, isType TypeExpression) { //funcion para agregar display

	if display, ok := a.Display[pos]; ok { //se verifica si existe el display

		a.Display[pos] = Display{Temp: temp, IsTemp: isTemp, LInicio: display.LInicio, LFinal: display.LFinal, Type: isType} //se agrega el display
	}

}

func (a *Arbol) GetDisplay(pos string) Display { //funcion para obtener el display

	if display, ok := a.Display[pos]; ok { //se verifica si existe el display

		return display //se retorna el display
	}

	return Display{Temp: "-1", IsTemp: false, LInicio: "-1", LFinal: "-1"} //se retorna el display, el -1 significa que no existe

}

func (a *Arbol) RestPosDisplay() { //funcion para obtener la posicion del stack global
	a.PosDisplay = a.PosDisplay - 1 //se aumenta la posicion del stack global
}

/** TABLE SYMBOL*/
func NewTableSymbol(name string, tipo string, ambito string, row int, column int, size string, pos string) *TableSymbol { //funcion nuevo arbol

	e := TableSymbol{ //se crea la tabla de simbolos
		Name:     name,   //se crea el nombre de la variable
		Tipo:     tipo,   //se crea el tipo de la variable
		Ambito:   ambito, //se crea el ambito de la variable
		Row:      row,    //se crea la fila de la variable
		Column:   column, //se crea la columna de la variable
		Size:     size,   //se crea el size de la variable
		Posicion: pos,    //se crea la posicion de la variable
	}
	return &e //se retorna la tabla de simbolos
}

func (a *Arbol) AddTableSymbol(e TableSymbol) { //funcion para agregar tabla de simbolos
	a._TableSymbol.Add(e) //se agrega la tabla de simbolos
}

/* Get TableSymbol */
func (a Arbol) GetTableSymbol() *arrayList.List { //funcion para obtener la tabla de simbolos
	return a._TableSymbol //se retorna la tabla de simbolos
}

/** TABLE DB*/
func NewTableDB(name string, row int, column int) *TableDB { //funcion nuevo arbol

	e := TableDB{ //se crea la tabla de simbolos
		Name:   name,   //se crea el nombre de la variable
		Row:    row,    //se crea la fila de la variable
		Column: column, //se crea la columna de la variable
	}
	return &e
}

/* Add TableSymbol */
func (a *Arbol) AddTableDB(e TableDB) { //funcion para agregar tabla de simbolos
	a._TableDB.Add(e) //se agrega la tabla de simbolos
}

/* Get TableSymbol */
func (a Arbol) GetTableDB() *arrayList.List { //funcion para obtener la tabla de simbolos
	return a._TableDB //se retorna la tabla de simbolos
}
