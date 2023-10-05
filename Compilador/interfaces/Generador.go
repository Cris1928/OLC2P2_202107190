package interfaces

import (
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

type Generator struct {
	temporal int
	label    int
	heap     int
	stack    int
	code     *arrayList.List
	native   *arrayList.List
	tempList *arrayList.List
}

func NewGenerator() *Generator {
	generator := Generator{
		temporal: 0,
		label:    0,
		heap:     0,
		stack:    0,
		code:     arrayList.New(),
		native:   arrayList.New(),
		tempList: arrayList.New()}
	return &generator
}
func (g Generator) GetCode() *arrayList.List {
	return g.code
}
func (g Generator) GetNative() *arrayList.List {
	return g.native
}
func (g Generator) GetTemp() *arrayList.List {
	return g.tempList
}
func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.temporal)
	g.temporal = g.temporal + 1
	g.tempList.Add(temp)
	return temp
}
func (g *Generator) NewLabel() string {
	temp := g.label
	g.label = g.label + 1
	return "L" + fmt.Sprintf("%v", temp)
}
func (g *Generator) AddLabel(label string) {
	g.code.Add(label + ":")
}
func (g *Generator) AddIf(left string, right string, operator string, label string) {
	g.code.Add("if(" + left + " " + operator + " " + right + ") goto " + label + ";")
}
func (g *Generator) AddGoto(label string) {
	g.code.Add("goto " + label + ";")
}
func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	g.code.Add(target + " = " + left + " " + operator + " " + right + ";")
}
func (g *Generator) AddExpression1(target string, result string) {
	g.code.Add(target + " = " + result + ";")
}
func (g *Generator) AddStack(temp string, result string) {
	g.code.Add("stack[(int)" + temp + "] = " + result + ";")
}
func (g *Generator) SaveTemp(pos string, temp string) {
	g.code.Add("stack[(int)" + pos + "] = " + temp + ";")
}
func (g *Generator) AddExpressionStack(target string, pos string) {
	g.code.Add(target + " = stack[(int)" + pos + "];")
}

func (g *Generator) AddHeap(temp string, ascii string) {
	g.code.Add("heap[(int)" + temp + "] = " + ascii + ";")
}
func (g *Generator) AddExpressionHeap(target string, right string) {
	g.code.Add(target + " = heap[(int)" + right + "];")
}

func (g *Generator) AddPrintf(typePrint string, value string) {
	g.code.Add("printf(\"%" + typePrint + "\"," + value + ");")
}

func (g *Generator) AddConcaPrintf(conca string) {
	g.code.Add(conca)
}

func (g *Generator) AddConca(conca string) {
	g.code.Add(conca)
}

func (g *Generator) AddComment(comment string) {
	g.code.Add("/************ " + comment + " ************/")
}

func (g *Generator) AddPrintfString() {

	temp := g.NewTemp()
	auxTemp := g.NewTemp()
	EV := g.NewLabel()
	EF := g.NewLabel()
	g.native.Add("\r/************ NATIVE PRINTF STRING ************/")
	g.native.Add("\rvoid printfString(){")
	g.native.Add(temp + " = P + 1;")
	g.native.Add(auxTemp + " = stack[(int) " + temp + "];")
	g.native.Add(EF + ":")

	temp = g.NewTemp()

	g.native.Add(temp + " = heap[(int)" + auxTemp + "];")
	g.native.Add("if(" + temp + " == -1) goto " + EV + ";")
	g.native.Add("printf(\"%c\"," + " (int)" + fmt.Sprintf("%v", temp) + ");")

	g.native.Add(auxTemp + " = " + auxTemp + " + 1;")
	g.native.Add("goto " + EF + ";")
	g.native.Add(EV + ":")
	g.native.Add("return;\n}")

}

func (g *Generator) PrintfString() {
	g.code.Add("printfString();")
}
func (g *Generator) AddConcatString() {

	conca := ""
	conca1 := ""
	temp := g.NewTemp()    // t2
	auxTemp := g.NewTemp() // t3
	EV := g.NewLabel()     // L0
	EF := g.NewLabel()     // L1
	g.native.Add("\r/************ NATIVE CONCAT STRING ************/")
	g.native.Add("\rvoid concatString(){")

	g.native.Add(temp + " = H + 0;")         // t2
	g.native.Add(auxTemp + " = P + 1;")      // t3
	conca += "stack[(int)P] = " + temp + ";" // 26

	temp = g.NewTemp()        // t4
	newAuxTemp := g.NewTemp() // t5
	g.native.Add(newAuxTemp + " = stack[(int) " + auxTemp + "];")
	g.native.Add(temp + " = P + 2;")
	conca1 += newAuxTemp + " = stack[(int)" + temp + "];" // 15
	g.native.Add(EV + ":")                                // E0
	temp = g.NewTemp()                                    // t6
	g.native.Add(temp + " = heap[(int)" + newAuxTemp + "];")
	g.native.Add("if (" + temp + " == -1) goto " + EF + ";")
	g.native.Add("heap[(int)H] = " + temp + ";")
	g.native.Add("H = H + 1;")
	g.native.Add(newAuxTemp + " = " + newAuxTemp + "+ 1;")
	g.native.Add("goto " + EV + ";")
	g.native.Add(EF + ":")
	g.native.Add(conca1)
	newLabel := g.NewLabel() // L2
	EV = g.NewLabel()        // L3
	g.native.Add(EV + ":")
	g.native.Add(temp + " = heap[(int)" + newAuxTemp + "];")
	g.native.Add("if (" + temp + " == -1) goto " + newLabel + "; ")
	g.native.Add("heap[(int)H] = " + temp + ";")
	g.native.Add("H = H + 1;")
	g.native.Add(newAuxTemp + " = " + newAuxTemp + "+ 1;")
	g.native.Add("goto " + EV + ";")
	g.native.Add(newLabel + ":")
	g.native.Add("heap[(int)H] = -1;")
	g.native.Add("H = H + 1;")
	g.native.Add(conca)
	g.native.Add("return;\n}")
}

func (g *Generator) ConcatString() {
	g.code.Add("concatString();")
}
func (g *Generator) AddCompareString() {

	conca := ""
	// conca1 := "" // conca1
	temp := g.NewTemp()       // t2
	auxTemp := g.NewTemp()    // t3
	newTemp := g.NewTemp()    // t4
	newLabel0 := g.NewLabel() // L0
	newLabel1 := g.NewLabel() // L1
	newLabel2 := g.NewLabel() // L2
	newLabel3 := g.NewLabel() // L3
	g.native.Add("\r/************ NATIVE COMPARE STRING ************/")
	g.native.Add("\rvoid compareString(){")

	g.native.Add(temp + " = P + 1;")
	g.native.Add(auxTemp + " = stack[(int)" + temp + "];")
	g.native.Add(temp + " = " + temp + " + 1;")
	g.native.Add(newTemp + " = stack[(int)" + temp + "];")
	g.native.Add(newLabel1 + ":")
	temp = g.NewTemp() // t5
	g.native.Add(temp + " = heap[(int)" + auxTemp + "];")
	conca += auxTemp + " = " + auxTemp + " + 1;"
	auxTemp = g.NewTemp() // t6
	g.native.Add(auxTemp + " = heap[(int)" + newTemp + "];")
	g.native.Add("if (" + temp + " != " + auxTemp + ") goto " + newLabel3 + ";")
	g.native.Add("if (" + temp + " == -1) goto " + newLabel2 + ";")
	g.native.Add(conca)
	g.native.Add(newTemp + " = " + newTemp + " + 1;")
	g.native.Add("goto " + newLabel1 + ";")
	g.native.Add(newLabel2 + ":")
	g.native.Add("stack[(int)P] = 1;")
	g.native.Add("goto " + newLabel0 + ";")
	g.native.Add(newLabel3 + ":")
	g.native.Add("stack[(int)P] = 0;")
	g.native.Add(newLabel0 + ":")
	g.native.Add("return;\n}")
}

func (g *Generator) CompareString() {
	g.code.Add("compareString();")
}

func (g *Generator) AddBoolean(EV string, EF string, newLabel string) {

	g.AddLabel(EV)
	g.AddPrintf("c", "(char)116")
	g.AddPrintf("c", "(char)114")
	g.AddPrintf("c", "(char)117")
	g.AddPrintf("c", "(char)101")
	g.AddGoto(newLabel)
	g.AddLabel(EF)
	g.AddPrintf("c", "(char)102")
	g.AddPrintf("c", "(char)97")
	g.AddPrintf("c", "(char)108")
	g.AddPrintf("c", "(char)115")
	g.AddPrintf("c", "(char)101")
	g.AddGoto(newLabel)
	g.AddLabel(newLabel)

}

func (g *Generator) Boolean(EV string, EF string, newLabel string, temp string) {

	g.AddLabel(EV)
	g.AddExpression(temp, "1", "0", "+")
	g.AddGoto(newLabel)
	g.AddLabel(EF)
	g.AddExpression(temp, "0", "0", "+")
	g.AddGoto(newLabel)
	g.AddLabel(newLabel)

}

func (g *Generator) StachHeap() {
	g.code.Add("P = 0; H = 0;\n")

}

func (g *Generator) PrintFunction(name string) {
	g.code.Add(name + "();")
}

func (g *Generator) AddFunction(tipo string, nombre string) {
	g.code.Add("\r" + tipo + " " + nombre + " {")

}

func (g *Generator) AddFunctionEnd(flag bool) {

	if flag {
		g.code.Add("\n\treturn 0;\r}")

	} else {
		g.code.Add("\n\treturn;\r}")
	}

}
