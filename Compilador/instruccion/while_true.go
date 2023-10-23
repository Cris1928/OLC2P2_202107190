package instruction

import (
	"OLC2/Compilador/interfaces"

	arrayList "github.com/colegno/arraylist"
)

type while_t struct {
	Instrucciones *arrayList.List
	Row           int
	Column        int
}

func NewWtrue(instruccion *arrayList.List, row int, column int) while_t {
	instr := while_t{instruccion, row, column}
	return instr
}

func (p while_t) Compilar(env *interfaces.Environment, tree *interfaces.Arbol, gen *interfaces.Generator) interface{} { //funcion compilar

	Linicio := gen.NewLabel()                     //creamos el label inicio
	Lfinal := gen.NewLabel()                      //creamos el label final
	tree.AddDisplay(Linicio, Lfinal, "-1", false) //agregamos el display al arbol de traduccion para el while true

	gen.AddComment("While") //agregamos un comentario
	gen.AddLabel(Linicio)   //agregamos el label inicio

	var newTable interfaces.Environment                                           //nuevo entorno
	newTable = interfaces.NewEnvironment(env)                                     //creamos un nuevo entorno
	newTable.UpdatePos(tree.GetPos(), env.Posicion, env.Posicion != 0, &newTable) //actualizamos la posicion

	for _, s := range p.Instrucciones.ToArray() { //recorremos las instrucciones
		s.(interfaces.Instruction).Compilar(&newTable, tree, gen)

	}
	gen.AddGoto(Linicio)
	gen.AddLabel(Lfinal)
	tree.RestPosDisplay()

	return interfaces.Value{Value: "", IsTemp: false, Type: interfaces.NULL, TrueLabel: "", FalseLabel: ""}
}
