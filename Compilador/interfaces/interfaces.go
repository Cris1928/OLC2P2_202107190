package interfaces

import (
	arrayList "github.com/colegno/arraylist"
)

type Symbol struct {
	Id       string
	Type     TypeExpression
	Value    interface{}
	IsMut    bool
	Posicion int
}

type SymbolFunction struct {
	Id            string
	Type          TypeExpression
	Instrucciones *arrayList.List
	Parametro     *arrayList.List
	IsMut         bool
}

type SymbolStructs struct {
	Id         string
	Type       TypeExpression
	Parameters *arrayList.List
	Value      *arrayList.List
	Size       int
}

type SymbolArrays struct {
	Id   string
	Type TypeExpression
	Tam  *arrayList.List
	Size int
}

type ValueStructs struct {
	Id       string
	Type     TypeExpression
	Posicion int
}

type Value struct {
	Value      string
	IsTemp     bool
	Type       TypeExpression
	TrueLabel  string
	FalseLabel string
}

type Expression interface {
	Compilar(env *Environment, tree *Arbol, gen *Generator) Value
}

type Instruction interface {
	Compilar(env *Environment, tree *Arbol, gen *Generator) interface{}
}
