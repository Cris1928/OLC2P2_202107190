package interfaces

import (
	"fmt"
	"time"

	arrayList "github.com/colegno/arraylist"
)

type Arbol struct {
	Code         *arrayList.List
	StackGlobal  int
	IsPrimitive  bool
	IsCocant     bool
	IsCompareStr bool
	_Exception   *arrayList.List
	Display      map[string]Display
	PosDisplay   int
	IsReturn     bool
	PosReturn    int
	_TableSymbol *arrayList.List
	_TableDB     *arrayList.List
}

type Exception struct {
	Tipo        string
	Descripcion string
	Row         int
	Column      int
	Time        string
}

type TableSymbol struct {
	Name     string
	Tipo     string
	Ambito   string
	Row      int
	Column   int
	Size     string
	Posicion string
}

type TableDB struct {
	Name   string
	Row    int
	Column int
}

type Display struct {
	Temp    string
	IsTemp  bool
	LInicio string
	LFinal  string
	Type    TypeExpression
}

func NewArbol() *Arbol {
	tree := Arbol{
		Code:         arrayList.New(),
		StackGlobal:  0,
		IsPrimitive:  false,
		IsCocant:     false,
		IsCompareStr: false,
		_Exception:   arrayList.New(),
		Display:      make(map[string]Display),
		PosDisplay:   0,
		IsReturn:     false,
		PosReturn:    -1,
		_TableSymbol: arrayList.New(),
		_TableDB:     arrayList.New(),
	}
	return &tree
}

func NewException(tipo string, descripcion string, row int, column int) *Exception {
	t := time.Now()
	hora := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	e := Exception{Tipo: tipo, Descripcion: descripcion, Row: row, Column: column, Time: hora}
	return &e
}

func (a *Arbol) AddCode(input string) {
	a.Code.Add(input)
}

func (a Arbol) GetCode() *arrayList.List {
	return a.Code
}

func (a *Arbol) AddException(e Exception) {
	t := time.Now()
	hora := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second())
	e.Time = hora
	a._Exception.Add(e)
}

func (a Arbol) GetException() *arrayList.List {
	return a._Exception
}

func (a Arbol) GetPos() int {
	return a.StackGlobal
}

func (a *Arbol) NewPos() {
	a.StackGlobal = a.StackGlobal + 1
}

func (a *Arbol) AddDisplay(labelInicio string, labelFinal string, temp string, isTemp bool) {

	pos := fmt.Sprintf("%v", a.PosDisplay)
	a.Display[pos] = Display{Temp: temp, IsTemp: isTemp, LInicio: labelInicio, LFinal: labelFinal, Type: NULL}
	a.PosDisplay = a.PosDisplay + 1
}

func (a *Arbol) SetDisplayTemp(pos string, temp string, isTemp bool, isType TypeExpression) {

	if display, ok := a.Display[pos]; ok {

		a.Display[pos] = Display{Temp: temp, IsTemp: isTemp, LInicio: display.LInicio, LFinal: display.LFinal, Type: isType}
	}

}

func (a *Arbol) GetDisplay(pos string) Display {

	if display, ok := a.Display[pos]; ok {

		return display
	}

	return Display{Temp: "-1", IsTemp: false, LInicio: "-1", LFinal: "-1"}

}

func (a *Arbol) RestPosDisplay() {
	a.PosDisplay = a.PosDisplay - 1
}

/** TABLE SYMBOL*/
func NewTableSymbol(name string, tipo string, ambito string, row int, column int, size string, pos string) *TableSymbol {

	e := TableSymbol{
		Name:     name,
		Tipo:     tipo,
		Ambito:   ambito,
		Row:      row,
		Column:   column,
		Size:     size,
		Posicion: pos,
	}
	return &e
}

func (a *Arbol) AddTableSymbol(e TableSymbol) {
	a._TableSymbol.Add(e)
}

/* Get TableSymbol */
func (a Arbol) GetTableSymbol() *arrayList.List {
	return a._TableSymbol
}

/** TABLE DB*/
func NewTableDB(name string, row int, column int) *TableDB {

	e := TableDB{
		Name:   name,
		Row:    row,
		Column: column,
	}
	return &e
}

/* Add TableSymbol */
func (a *Arbol) AddTableDB(e TableDB) {
	a._TableDB.Add(e)
}

/* Get TableSymbol */
func (a Arbol) GetTableDB() *arrayList.List {
	return a._TableDB
}
