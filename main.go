package main

//antlr4 -Dlanguage=Go -o parser Swiftlexer.g4
//antlr4 -Dlanguage=Go -o parser Swiftgrammar.g4
import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"

	"OLC2/Compilador/interfaces"

	"OLC2/Compilador/Gram/parser"

	"github.com/antlr4-go/antlr/v4"
)

var CODE_OUT_ string = ""
var CODE_HEAD string = ""
var TSGlobalError []interface{}
var TSGlobalSymbol []interface{}
var TSGlobalDB []interface{}

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./Cliente", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	text := ""

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"CODE_INPUT":  text,
			"CODE_OUT":    "",
			"Tabla_Error": nil,
		})
	})

	app.Post("/compilar", Execute)
	_ = app.Listen(":3000")
}

type getInput struct {
	Input string `form:"editor"`
}

func Execute(c *fiber.Ctx) error {
	data := new(getInput)
	fmt.Println(data)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	is := antlr.NewInputStream(data.Input)

	lexer := parser.NewSwiftlexer(is) // aqui van tokens
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewSwiftgrammar(stream)

	p.BuildParseTrees = true
	tree := p.Start_()

	result := NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(result, tree)

	return c.Render("index", fiber.Map{
		"CODE_INPUT":   data.Input,
		"CODE_OUT":     CODE_OUT_,
		"Tabla_Error":  TSGlobalError,
		"Tabla_Symbol": TSGlobalSymbol,
		"Tabla_DB":     TSGlobalDB,
	})
}

/* ANTLR*/

type TreeShapeListener struct {
	*parser.BaseSwiftgrammarListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}
func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetLista()

	TSGlobalError = nil
	TSGlobalSymbol = nil
	TSGlobalDB = nil
	/* SALIDA */
	var _salida string
	_salida = ""
	CODE_OUT_ = ""
	/* TREE */
	var tree *interfaces.Arbol
	tree = interfaces.NewArbol()
	/* ENVIRONMENT */
	var globalEnv interfaces.Environment
	globalEnv = interfaces.NewEnvironment(nil)

	/* GENERATOR */
	var gen *interfaces.Generator
	gen = interfaces.NewGenerator()

	globalEnv.UpdatePos(0, 0, true, &globalEnv)
	gen.AddComment("Main")
	gen.AddFunction("int", "main()")
	gen.StachHeap()
	for _, s := range result.ToArray() {
		newInstr := s.(interfaces.Instruction)
		fmt.Println(reflect.TypeOf(newInstr).String())
		s.(interfaces.Instruction).Compilar(&globalEnv, tree, gen)

		if reflect.TypeOf(s).String() == "transferencia.Break" {
			excep := interfaces.NewException("Semantico", "Sentencia Break fuera de Ciclo.", -1, -1)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			break
		}
		if reflect.TypeOf(s).String() == "transferencia.Continue" {
			excep := interfaces.NewException("Semantico", "Sentencia Continue fuera de Ciclo.", -1, -1)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			break
		}
		if reflect.TypeOf(s).String() == "transferencia.Return" {
			excep := interfaces.NewException("Semantico", "Sentencia Return fuera de Ciclo.", -1, -1)
			tree.AddException(interfaces.Exception{Tipo: excep.Tipo, Descripcion: excep.Descripcion, Row: excep.Row, Column: excep.Row})
			break
		}

	}

	gen.AddFunctionEnd(true)

	_salida += "#include <stdio.h>\n"
	_salida += "#include <math.h>\n"
	_salida += "double heap[23111998];\n"
	_salida += "double stack[23111998];\n"
	_salida += "double P;\n"
	_salida += "double H;\n"

	if gen.GetTemp().Len() != 0 {
		_salida += "double "

		_salida += fmt.Sprintf("%v", gen.GetTemp().GetValue(0))
		gen.GetTemp().RemoveAtIndex(0)

		for _, s := range gen.GetTemp().ToArray() {
			_salida += ", "
			_salida += fmt.Sprintf("%v", s)
		}

		_salida += ";\n\n"

	}

	for _, s := range gen.GetNative().ToArray() {
		_salida += "\t" + fmt.Sprintf("%v", s)
		_salida += "\n"
	}
	CODE_HEAD = _salida
	for _, s := range gen.GetCode().ToArray() {
		_salida += "\t" + fmt.Sprintf("%v", s)
		_salida += "\n"
	}

	fmt.Println("----------")
	var OutException string
	OutException = ""

	for _, s := range tree.GetException().ToArray() {
		OutException += fmt.Sprintf("%v", s)
		m := make(map[string]string)
		m["Id"] = fmt.Sprintf("%v", s.(interfaces.Exception).Tipo)
		m["Descripcion"] = fmt.Sprintf("%v", s.(interfaces.Exception).Descripcion)
		m["Row"] = fmt.Sprintf("%v", s.(interfaces.Exception).Row)
		m["Column"] = fmt.Sprintf("%v", s.(interfaces.Exception).Column)
		m["Time"] = fmt.Sprintf("%v", s.(interfaces.Exception).Time)

		TSGlobalError = append(TSGlobalError, m)
	}

	fmt.Println(OutException)
	fmt.Println("----------")
	OutException = ""

	for _, s := range tree.GetTableSymbol().ToArray() {
		OutException += fmt.Sprintf("%v", s)
		m := make(map[string]string)
		m["Name"] = fmt.Sprintf("%v", s.(interfaces.TableSymbol).Name)
		m["Tipo"] = fmt.Sprintf("%v", s.(interfaces.TableSymbol).Tipo)
		m["Ambito"] = fmt.Sprintf("%v", s.(interfaces.TableSymbol).Ambito)
		m["Row"] = fmt.Sprintf("%v", s.(interfaces.TableSymbol).Row)
		m["Column"] = fmt.Sprintf("%v", s.(interfaces.TableSymbol).Column)
		m["Size"] = fmt.Sprintf("%v", s.(interfaces.TableSymbol).Size)
		m["Posicion"] = fmt.Sprintf("%v", s.(interfaces.TableSymbol).Posicion)

		TSGlobalSymbol = append(TSGlobalSymbol, m)
	}

	fmt.Println(OutException)

	fmt.Println("----------")

	auxCont := 1
	for _, s := range tree.GetTableDB().ToArray() {

		m := make(map[string]string)
		m["Id"] = fmt.Sprintf("%v", auxCont)
		m["Name"] = fmt.Sprintf("%v", s.(interfaces.TableDB).Name)
		m["Row"] = fmt.Sprintf("%v", s.(interfaces.TableDB).Row)
		m["Column"] = fmt.Sprintf("%v", s.(interfaces.TableDB).Column)

		TSGlobalDB = append(TSGlobalDB, m)
		auxCont++
	}

	CODE_OUT_ = _salida

}
