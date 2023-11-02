// Code generated from Swiftgrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgrammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "OLC2/Compilador/interfaces"
import "OLC2/Compilador/instruccion"
import "OLC2/Compilador/expression"
import "OLC2/Compilador/instruccion/ternario"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Swiftgrammar struct {
	*antlr.BaseParser
}

var SwiftgrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftgrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'print'", "'var'", "'let'", "'if'", "'else'", "'for'", "'while'",
		"'switch'", "'case'", "'default'", "'in'", "'break'", "'continue'",
		"'return'", "'func'", "'struct'", "'Int'", "'Float'", "'Double'", "'Bool'",
		"'Character'", "'String'", "'true'", "", "", "", "", "", "", "'...'",
		"'.'", "';'", "','", "':'", "'='", "'=='", "'>='", "'=>'", "'->'", "'<='",
		"'!='", "'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "'&&'", "'&'", "'||'", "'|'", "'!'",
	}
	staticData.SymbolicNames = []string{
		"", "PRINT", "VAR", "LET", "IF", "ELSE", "FOR", "WHILE", "SWITCH", "CASE",
		"DEFAULT", "IN", "BREAK", "CONTINUE", "RETURN", "FUNC", "STRUCT", "R_INT",
		"R_FLOAT", "R_DOUBLE", "R_BOOL", "R_CHAR", "R_STRING", "TRUE", "NUMBER",
		"DOUBLE", "CHAR", "STRING", "BOOLEAN", "ID", "TK_TRIPLEPUNTO", "TK_PUNTO",
		"TK_PUNTOCOMA", "TK_COMA", "TK_DOSPUNTOS", "TK_IGUAL", "TK_IGUALIGUAL",
		"TK_MAYORIGUAL", "TK_IGUALMAYOR", "TK_MENOSMAYOR", "TK_MENORIGUAL",
		"TK_DIFIGUAL", "TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV", "TK_MODULO",
		"TK_MAS", "TK_MENOS", "TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC",
		"TK_CORA", "TK_CORC", "TK_AND", "TK_AMPERSAND", "TK_OR", "TK_BARRA",
		"TK_NOT", "WHITESPACE", "TK_MULTI", "TK_LINE",
	}
	staticData.RuleNames = []string{
		"start", "instrucciones", "instruccion", "instruccion_println", "instruccion_declaracion",
		"instruccion_asignacion", "instruccion_if", "instr_else_if", "instruccion_ternario",
		"instr_else_if_ternario", "instruccion_switch", "list_case", "instruccion_case",
		"list_expre_case", "block_case", "block_instr_switch", "instr_default",
		"block_default", "instruccion_switch_ternario", "list_case_ternario",
		"instr_case_ter", "list_expre_case_ter", "block_case_ter", "instr_default_ter",
		"instruccion_while", "instruccion_for_in", "instruccion_while_true",
		"instruccion_while_true_ternario", "instruccion_break", "instruccion_continue",
		"instruccion_return", "instruccion_func", "list_function_parameters",
		"block_parameters_fn", "instruccion_llamada", "instr_llamada_expre",
		"instruccion_structs_declaracion", "list_struct_parameters", "block_structs_parameters",
		"instr_arrays_identifier", "list_arrays_parameters_id", "block_arrays_identifier",
		"instr_structs_declaration", "list_struct_parameters_decla", "block_structs_parameters_decla",
		"instr_structs_identifier", "list_struct_parameters_id", "block_structs_identifier",
		"instr_structs_assignment", "instruccion_tipo", "list_expression", "block_list_expression",
		"expressions", "expre_log", "expre_rel", "expre_arit", "expre_valor",
		"primitivo",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 62, 777, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 122, 8, 1, 11, 1, 12, 1, 123, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		162, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 184,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 226, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 260, 8, 6, 1, 7, 5, 7, 263, 8, 7, 10, 7,
		12, 7, 266, 9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 297, 8, 8, 1, 9,
		5, 9, 300, 8, 9, 10, 9, 12, 9, 303, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 320, 8, 10, 1, 11, 4, 11, 323, 8, 11, 11, 11, 12, 11, 324, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3,
		12, 338, 8, 12, 1, 13, 4, 13, 341, 8, 13, 11, 13, 12, 13, 342, 1, 14, 1,
		14, 1, 14, 1, 14, 3, 14, 349, 8, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 364, 8,
		16, 1, 17, 4, 17, 367, 8, 17, 11, 17, 12, 17, 368, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 384, 8, 18, 1, 19, 4, 19, 387, 8, 19, 11, 19, 12, 19, 388, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 401,
		8, 20, 1, 21, 4, 21, 404, 8, 21, 11, 21, 12, 21, 405, 1, 22, 1, 22, 1,
		22, 1, 22, 3, 22, 412, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 425, 8, 23, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 3, 25, 454, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 3, 28, 473, 8, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 3, 31, 523, 8, 31, 1, 32, 4, 32, 526, 8, 32, 11, 32, 12, 32, 527,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 3, 33, 543, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 3, 34, 553, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 3, 35, 563, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 37, 4, 37, 572, 8, 37, 11, 37, 12, 37, 573, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 584, 8, 38, 1, 39,
		1, 39, 1, 39, 1, 40, 4, 40, 590, 8, 40, 11, 40, 12, 40, 591, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 614, 8,
		42, 1, 43, 4, 43, 617, 8, 43, 11, 43, 12, 43, 618, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 629, 8, 44, 1, 45, 1, 45, 1,
		45, 1, 46, 4, 46, 635, 8, 46, 11, 46, 12, 46, 636, 1, 47, 1, 47, 1, 47,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 657, 8, 49, 1, 50, 4, 50, 660, 8,
		50, 11, 50, 12, 50, 661, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51,
		1, 51, 1, 51, 3, 51, 673, 8, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1,
		52, 1, 52, 1, 52, 1, 52, 3, 52, 684, 8, 52, 1, 53, 1, 53, 1, 53, 1, 53,
		1, 53, 1, 53, 1, 53, 1, 53, 3, 53, 694, 8, 53, 1, 53, 1, 53, 1, 53, 1,
		53, 1, 53, 5, 53, 701, 8, 53, 10, 53, 12, 53, 704, 9, 53, 1, 54, 1, 54,
		1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 5, 54, 715, 8, 54, 10,
		54, 12, 54, 718, 9, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55,
		1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 3, 55, 733, 8, 55, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 5, 55, 745,
		8, 55, 10, 55, 12, 55, 748, 9, 55, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1,
		57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57,
		1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 3, 57, 775,
		8, 57, 1, 57, 0, 3, 106, 108, 110, 58, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90,
		92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 0, 4, 2, 0, 55,
		55, 57, 57, 2, 0, 36, 37, 40, 43, 1, 0, 44, 46, 1, 0, 47, 48, 802, 0, 116,
		1, 0, 0, 0, 2, 121, 1, 0, 0, 0, 4, 161, 1, 0, 0, 0, 6, 183, 1, 0, 0, 0,
		8, 225, 1, 0, 0, 0, 10, 227, 1, 0, 0, 0, 12, 259, 1, 0, 0, 0, 14, 264,
		1, 0, 0, 0, 16, 296, 1, 0, 0, 0, 18, 301, 1, 0, 0, 0, 20, 319, 1, 0, 0,
		0, 22, 322, 1, 0, 0, 0, 24, 337, 1, 0, 0, 0, 26, 340, 1, 0, 0, 0, 28, 348,
		1, 0, 0, 0, 30, 350, 1, 0, 0, 0, 32, 363, 1, 0, 0, 0, 34, 366, 1, 0, 0,
		0, 36, 383, 1, 0, 0, 0, 38, 386, 1, 0, 0, 0, 40, 400, 1, 0, 0, 0, 42, 403,
		1, 0, 0, 0, 44, 411, 1, 0, 0, 0, 46, 424, 1, 0, 0, 0, 48, 426, 1, 0, 0,
		0, 50, 453, 1, 0, 0, 0, 52, 455, 1, 0, 0, 0, 54, 462, 1, 0, 0, 0, 56, 472,
		1, 0, 0, 0, 58, 474, 1, 0, 0, 0, 60, 476, 1, 0, 0, 0, 62, 522, 1, 0, 0,
		0, 64, 525, 1, 0, 0, 0, 66, 542, 1, 0, 0, 0, 68, 552, 1, 0, 0, 0, 70, 562,
		1, 0, 0, 0, 72, 564, 1, 0, 0, 0, 74, 571, 1, 0, 0, 0, 76, 583, 1, 0, 0,
		0, 78, 585, 1, 0, 0, 0, 80, 589, 1, 0, 0, 0, 82, 593, 1, 0, 0, 0, 84, 613,
		1, 0, 0, 0, 86, 616, 1, 0, 0, 0, 88, 628, 1, 0, 0, 0, 90, 630, 1, 0, 0,
		0, 92, 634, 1, 0, 0, 0, 94, 638, 1, 0, 0, 0, 96, 641, 1, 0, 0, 0, 98, 656,
		1, 0, 0, 0, 100, 659, 1, 0, 0, 0, 102, 672, 1, 0, 0, 0, 104, 683, 1, 0,
		0, 0, 106, 693, 1, 0, 0, 0, 108, 705, 1, 0, 0, 0, 110, 732, 1, 0, 0, 0,
		112, 749, 1, 0, 0, 0, 114, 774, 1, 0, 0, 0, 116, 117, 3, 2, 1, 0, 117,
		118, 5, 0, 0, 1, 118, 119, 6, 0, -1, 0, 119, 1, 1, 0, 0, 0, 120, 122, 3,
		4, 2, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0,
		0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 6, 1, -1, 0, 126,
		3, 1, 0, 0, 0, 127, 128, 3, 6, 3, 0, 128, 129, 6, 2, -1, 0, 129, 162, 1,
		0, 0, 0, 130, 162, 3, 72, 36, 0, 131, 132, 3, 8, 4, 0, 132, 133, 6, 2,
		-1, 0, 133, 162, 1, 0, 0, 0, 134, 135, 3, 10, 5, 0, 135, 136, 6, 2, -1,
		0, 136, 162, 1, 0, 0, 0, 137, 162, 3, 96, 48, 0, 138, 139, 3, 12, 6, 0,
		139, 140, 6, 2, -1, 0, 140, 162, 1, 0, 0, 0, 141, 142, 3, 50, 25, 0, 142,
		143, 6, 2, -1, 0, 143, 162, 1, 0, 0, 0, 144, 145, 3, 48, 24, 0, 145, 146,
		6, 2, -1, 0, 146, 162, 1, 0, 0, 0, 147, 148, 3, 52, 26, 0, 148, 149, 6,
		2, -1, 0, 149, 162, 1, 0, 0, 0, 150, 162, 3, 20, 10, 0, 151, 162, 3, 56,
		28, 0, 152, 162, 3, 58, 29, 0, 153, 154, 3, 62, 31, 0, 154, 155, 6, 2,
		-1, 0, 155, 162, 1, 0, 0, 0, 156, 162, 3, 68, 34, 0, 157, 158, 3, 60, 30,
		0, 158, 159, 6, 2, -1, 0, 159, 162, 1, 0, 0, 0, 160, 162, 3, 84, 42, 0,
		161, 127, 1, 0, 0, 0, 161, 130, 1, 0, 0, 0, 161, 131, 1, 0, 0, 0, 161,
		134, 1, 0, 0, 0, 161, 137, 1, 0, 0, 0, 161, 138, 1, 0, 0, 0, 161, 141,
		1, 0, 0, 0, 161, 144, 1, 0, 0, 0, 161, 147, 1, 0, 0, 0, 161, 150, 1, 0,
		0, 0, 161, 151, 1, 0, 0, 0, 161, 152, 1, 0, 0, 0, 161, 153, 1, 0, 0, 0,
		161, 156, 1, 0, 0, 0, 161, 157, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162,
		5, 1, 0, 0, 0, 163, 164, 5, 1, 0, 0, 164, 165, 5, 49, 0, 0, 165, 166, 3,
		114, 57, 0, 166, 167, 5, 50, 0, 0, 167, 168, 6, 3, -1, 0, 168, 184, 1,
		0, 0, 0, 169, 170, 5, 1, 0, 0, 170, 171, 5, 49, 0, 0, 171, 172, 5, 27,
		0, 0, 172, 173, 5, 33, 0, 0, 173, 174, 3, 100, 50, 0, 174, 175, 5, 50,
		0, 0, 175, 176, 6, 3, -1, 0, 176, 184, 1, 0, 0, 0, 177, 178, 5, 1, 0, 0,
		178, 179, 5, 49, 0, 0, 179, 180, 3, 104, 52, 0, 180, 181, 5, 50, 0, 0,
		181, 182, 6, 3, -1, 0, 182, 184, 1, 0, 0, 0, 183, 163, 1, 0, 0, 0, 183,
		169, 1, 0, 0, 0, 183, 177, 1, 0, 0, 0, 184, 7, 1, 0, 0, 0, 185, 186, 5,
		2, 0, 0, 186, 187, 5, 29, 0, 0, 187, 188, 5, 35, 0, 0, 188, 189, 3, 104,
		52, 0, 189, 190, 6, 4, -1, 0, 190, 226, 1, 0, 0, 0, 191, 192, 5, 2, 0,
		0, 192, 193, 5, 29, 0, 0, 193, 194, 5, 34, 0, 0, 194, 195, 3, 98, 49, 0,
		195, 196, 6, 4, -1, 0, 196, 226, 1, 0, 0, 0, 197, 198, 5, 2, 0, 0, 198,
		199, 5, 29, 0, 0, 199, 200, 5, 34, 0, 0, 200, 201, 3, 98, 49, 0, 201, 202,
		5, 35, 0, 0, 202, 203, 3, 104, 52, 0, 203, 204, 6, 4, -1, 0, 204, 226,
		1, 0, 0, 0, 205, 206, 5, 3, 0, 0, 206, 207, 5, 29, 0, 0, 207, 208, 5, 35,
		0, 0, 208, 209, 3, 104, 52, 0, 209, 210, 6, 4, -1, 0, 210, 226, 1, 0, 0,
		0, 211, 212, 5, 3, 0, 0, 212, 213, 5, 29, 0, 0, 213, 214, 5, 34, 0, 0,
		214, 215, 3, 98, 49, 0, 215, 216, 6, 4, -1, 0, 216, 226, 1, 0, 0, 0, 217,
		218, 5, 3, 0, 0, 218, 219, 5, 29, 0, 0, 219, 220, 5, 34, 0, 0, 220, 221,
		3, 98, 49, 0, 221, 222, 5, 35, 0, 0, 222, 223, 3, 104, 52, 0, 223, 224,
		6, 4, -1, 0, 224, 226, 1, 0, 0, 0, 225, 185, 1, 0, 0, 0, 225, 191, 1, 0,
		0, 0, 225, 197, 1, 0, 0, 0, 225, 205, 1, 0, 0, 0, 225, 211, 1, 0, 0, 0,
		225, 217, 1, 0, 0, 0, 226, 9, 1, 0, 0, 0, 227, 228, 5, 29, 0, 0, 228, 229,
		5, 35, 0, 0, 229, 230, 3, 104, 52, 0, 230, 231, 6, 5, -1, 0, 231, 11, 1,
		0, 0, 0, 232, 233, 5, 4, 0, 0, 233, 234, 3, 104, 52, 0, 234, 235, 5, 51,
		0, 0, 235, 236, 3, 2, 1, 0, 236, 237, 5, 52, 0, 0, 237, 238, 6, 6, -1,
		0, 238, 260, 1, 0, 0, 0, 239, 240, 5, 4, 0, 0, 240, 241, 3, 104, 52, 0,
		241, 242, 5, 51, 0, 0, 242, 243, 3, 2, 1, 0, 243, 244, 5, 52, 0, 0, 244,
		245, 5, 5, 0, 0, 245, 246, 5, 51, 0, 0, 246, 247, 3, 2, 1, 0, 247, 248,
		5, 52, 0, 0, 248, 249, 6, 6, -1, 0, 249, 260, 1, 0, 0, 0, 250, 251, 5,
		4, 0, 0, 251, 252, 3, 104, 52, 0, 252, 253, 5, 51, 0, 0, 253, 254, 3, 2,
		1, 0, 254, 255, 5, 52, 0, 0, 255, 256, 5, 5, 0, 0, 256, 257, 3, 14, 7,
		0, 257, 258, 6, 6, -1, 0, 258, 260, 1, 0, 0, 0, 259, 232, 1, 0, 0, 0, 259,
		239, 1, 0, 0, 0, 259, 250, 1, 0, 0, 0, 260, 13, 1, 0, 0, 0, 261, 263, 3,
		12, 6, 0, 262, 261, 1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0,
		0, 264, 265, 1, 0, 0, 0, 265, 267, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267,
		268, 6, 7, -1, 0, 268, 15, 1, 0, 0, 0, 269, 270, 5, 4, 0, 0, 270, 271,
		3, 104, 52, 0, 271, 272, 5, 51, 0, 0, 272, 273, 3, 104, 52, 0, 273, 274,
		5, 52, 0, 0, 274, 275, 6, 8, -1, 0, 275, 297, 1, 0, 0, 0, 276, 277, 5,
		4, 0, 0, 277, 278, 3, 104, 52, 0, 278, 279, 5, 51, 0, 0, 279, 280, 3, 104,
		52, 0, 280, 281, 5, 52, 0, 0, 281, 282, 5, 5, 0, 0, 282, 283, 5, 51, 0,
		0, 283, 284, 3, 104, 52, 0, 284, 285, 5, 52, 0, 0, 285, 286, 6, 8, -1,
		0, 286, 297, 1, 0, 0, 0, 287, 288, 5, 4, 0, 0, 288, 289, 3, 104, 52, 0,
		289, 290, 5, 51, 0, 0, 290, 291, 3, 104, 52, 0, 291, 292, 5, 52, 0, 0,
		292, 293, 5, 5, 0, 0, 293, 294, 3, 18, 9, 0, 294, 295, 6, 8, -1, 0, 295,
		297, 1, 0, 0, 0, 296, 269, 1, 0, 0, 0, 296, 276, 1, 0, 0, 0, 296, 287,
		1, 0, 0, 0, 297, 17, 1, 0, 0, 0, 298, 300, 3, 16, 8, 0, 299, 298, 1, 0,
		0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0,
		302, 304, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 305, 6, 9, -1, 0, 305,
		19, 1, 0, 0, 0, 306, 307, 5, 8, 0, 0, 307, 308, 3, 104, 52, 0, 308, 309,
		5, 51, 0, 0, 309, 310, 3, 22, 11, 0, 310, 311, 3, 34, 17, 0, 311, 312,
		5, 52, 0, 0, 312, 320, 1, 0, 0, 0, 313, 314, 5, 8, 0, 0, 314, 315, 3, 104,
		52, 0, 315, 316, 5, 51, 0, 0, 316, 317, 3, 34, 17, 0, 317, 318, 5, 52,
		0, 0, 318, 320, 1, 0, 0, 0, 319, 306, 1, 0, 0, 0, 319, 313, 1, 0, 0, 0,
		320, 21, 1, 0, 0, 0, 321, 323, 3, 24, 12, 0, 322, 321, 1, 0, 0, 0, 323,
		324, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 23, 1,
		0, 0, 0, 326, 327, 3, 26, 13, 0, 327, 328, 5, 34, 0, 0, 328, 329, 5, 51,
		0, 0, 329, 330, 3, 2, 1, 0, 330, 331, 5, 52, 0, 0, 331, 338, 1, 0, 0, 0,
		332, 333, 3, 26, 13, 0, 333, 334, 5, 34, 0, 0, 334, 335, 3, 30, 15, 0,
		335, 336, 5, 33, 0, 0, 336, 338, 1, 0, 0, 0, 337, 326, 1, 0, 0, 0, 337,
		332, 1, 0, 0, 0, 338, 25, 1, 0, 0, 0, 339, 341, 3, 28, 14, 0, 340, 339,
		1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0,
		0, 0, 343, 27, 1, 0, 0, 0, 344, 345, 3, 104, 52, 0, 345, 346, 5, 58, 0,
		0, 346, 349, 1, 0, 0, 0, 347, 349, 3, 104, 52, 0, 348, 344, 1, 0, 0, 0,
		348, 347, 1, 0, 0, 0, 349, 29, 1, 0, 0, 0, 350, 351, 3, 4, 2, 0, 351, 31,
		1, 0, 0, 0, 352, 353, 5, 9, 0, 0, 353, 354, 5, 29, 0, 0, 354, 355, 5, 34,
		0, 0, 355, 356, 5, 51, 0, 0, 356, 357, 3, 2, 1, 0, 357, 358, 5, 52, 0,
		0, 358, 364, 1, 0, 0, 0, 359, 360, 5, 9, 0, 0, 360, 361, 5, 29, 0, 0, 361,
		362, 5, 34, 0, 0, 362, 364, 3, 30, 15, 0, 363, 352, 1, 0, 0, 0, 363, 359,
		1, 0, 0, 0, 364, 33, 1, 0, 0, 0, 365, 367, 3, 32, 16, 0, 366, 365, 1, 0,
		0, 0, 367, 368, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0,
		369, 35, 1, 0, 0, 0, 370, 371, 5, 8, 0, 0, 371, 372, 3, 104, 52, 0, 372,
		373, 5, 51, 0, 0, 373, 374, 3, 38, 19, 0, 374, 375, 3, 46, 23, 0, 375,
		376, 5, 52, 0, 0, 376, 384, 1, 0, 0, 0, 377, 378, 5, 8, 0, 0, 378, 379,
		3, 104, 52, 0, 379, 380, 5, 51, 0, 0, 380, 381, 3, 46, 23, 0, 381, 382,
		5, 52, 0, 0, 382, 384, 1, 0, 0, 0, 383, 370, 1, 0, 0, 0, 383, 377, 1, 0,
		0, 0, 384, 37, 1, 0, 0, 0, 385, 387, 3, 40, 20, 0, 386, 385, 1, 0, 0, 0,
		387, 388, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389,
		39, 1, 0, 0, 0, 390, 391, 3, 42, 21, 0, 391, 392, 5, 34, 0, 0, 392, 393,
		3, 104, 52, 0, 393, 401, 1, 0, 0, 0, 394, 395, 3, 42, 21, 0, 395, 396,
		5, 34, 0, 0, 396, 397, 5, 51, 0, 0, 397, 398, 3, 104, 52, 0, 398, 399,
		5, 52, 0, 0, 399, 401, 1, 0, 0, 0, 400, 390, 1, 0, 0, 0, 400, 394, 1, 0,
		0, 0, 401, 41, 1, 0, 0, 0, 402, 404, 3, 44, 22, 0, 403, 402, 1, 0, 0, 0,
		404, 405, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406,
		43, 1, 0, 0, 0, 407, 408, 3, 104, 52, 0, 408, 409, 5, 58, 0, 0, 409, 412,
		1, 0, 0, 0, 410, 412, 3, 104, 52, 0, 411, 407, 1, 0, 0, 0, 411, 410, 1,
		0, 0, 0, 412, 45, 1, 0, 0, 0, 413, 414, 5, 9, 0, 0, 414, 415, 5, 29, 0,
		0, 415, 416, 5, 34, 0, 0, 416, 425, 3, 104, 52, 0, 417, 418, 5, 9, 0, 0,
		418, 419, 5, 29, 0, 0, 419, 420, 5, 34, 0, 0, 420, 421, 5, 51, 0, 0, 421,
		422, 3, 104, 52, 0, 422, 423, 5, 52, 0, 0, 423, 425, 1, 0, 0, 0, 424, 413,
		1, 0, 0, 0, 424, 417, 1, 0, 0, 0, 425, 47, 1, 0, 0, 0, 426, 427, 5, 7,
		0, 0, 427, 428, 3, 104, 52, 0, 428, 429, 5, 51, 0, 0, 429, 430, 3, 2, 1,
		0, 430, 431, 5, 52, 0, 0, 431, 432, 6, 24, -1, 0, 432, 49, 1, 0, 0, 0,
		433, 434, 5, 6, 0, 0, 434, 435, 5, 29, 0, 0, 435, 436, 5, 11, 0, 0, 436,
		437, 3, 104, 52, 0, 437, 438, 5, 30, 0, 0, 438, 439, 3, 104, 52, 0, 439,
		440, 5, 51, 0, 0, 440, 441, 3, 2, 1, 0, 441, 442, 5, 52, 0, 0, 442, 443,
		6, 25, -1, 0, 443, 454, 1, 0, 0, 0, 444, 445, 5, 6, 0, 0, 445, 446, 5,
		29, 0, 0, 446, 447, 5, 11, 0, 0, 447, 448, 3, 104, 52, 0, 448, 449, 5,
		51, 0, 0, 449, 450, 3, 2, 1, 0, 450, 451, 5, 52, 0, 0, 451, 452, 6, 25,
		-1, 0, 452, 454, 1, 0, 0, 0, 453, 433, 1, 0, 0, 0, 453, 444, 1, 0, 0, 0,
		454, 51, 1, 0, 0, 0, 455, 456, 5, 7, 0, 0, 456, 457, 5, 23, 0, 0, 457,
		458, 5, 51, 0, 0, 458, 459, 3, 2, 1, 0, 459, 460, 5, 52, 0, 0, 460, 461,
		6, 26, -1, 0, 461, 53, 1, 0, 0, 0, 462, 463, 5, 7, 0, 0, 463, 464, 5, 23,
		0, 0, 464, 465, 5, 51, 0, 0, 465, 466, 3, 2, 1, 0, 466, 467, 5, 52, 0,
		0, 467, 468, 6, 27, -1, 0, 468, 55, 1, 0, 0, 0, 469, 473, 5, 12, 0, 0,
		470, 471, 5, 12, 0, 0, 471, 473, 3, 104, 52, 0, 472, 469, 1, 0, 0, 0, 472,
		470, 1, 0, 0, 0, 473, 57, 1, 0, 0, 0, 474, 475, 5, 13, 0, 0, 475, 59, 1,
		0, 0, 0, 476, 477, 5, 14, 0, 0, 477, 478, 3, 104, 52, 0, 478, 479, 6, 30,
		-1, 0, 479, 61, 1, 0, 0, 0, 480, 481, 5, 15, 0, 0, 481, 482, 5, 29, 0,
		0, 482, 483, 5, 49, 0, 0, 483, 484, 5, 50, 0, 0, 484, 485, 5, 51, 0, 0,
		485, 486, 3, 2, 1, 0, 486, 487, 5, 52, 0, 0, 487, 488, 6, 31, -1, 0, 488,
		523, 1, 0, 0, 0, 489, 490, 5, 15, 0, 0, 490, 491, 5, 29, 0, 0, 491, 492,
		5, 49, 0, 0, 492, 493, 5, 50, 0, 0, 493, 494, 5, 39, 0, 0, 494, 495, 3,
		98, 49, 0, 495, 496, 5, 51, 0, 0, 496, 497, 3, 2, 1, 0, 497, 498, 5, 52,
		0, 0, 498, 499, 6, 31, -1, 0, 499, 523, 1, 0, 0, 0, 500, 501, 5, 15, 0,
		0, 501, 502, 5, 29, 0, 0, 502, 503, 5, 49, 0, 0, 503, 504, 3, 64, 32, 0,
		504, 505, 5, 50, 0, 0, 505, 506, 5, 51, 0, 0, 506, 507, 3, 2, 1, 0, 507,
		508, 5, 52, 0, 0, 508, 509, 6, 31, -1, 0, 509, 523, 1, 0, 0, 0, 510, 511,
		5, 15, 0, 0, 511, 512, 5, 29, 0, 0, 512, 513, 5, 49, 0, 0, 513, 514, 3,
		64, 32, 0, 514, 515, 5, 50, 0, 0, 515, 516, 5, 39, 0, 0, 516, 517, 3, 98,
		49, 0, 517, 518, 5, 51, 0, 0, 518, 519, 3, 2, 1, 0, 519, 520, 5, 52, 0,
		0, 520, 521, 6, 31, -1, 0, 521, 523, 1, 0, 0, 0, 522, 480, 1, 0, 0, 0,
		522, 489, 1, 0, 0, 0, 522, 500, 1, 0, 0, 0, 522, 510, 1, 0, 0, 0, 523,
		63, 1, 0, 0, 0, 524, 526, 3, 66, 33, 0, 525, 524, 1, 0, 0, 0, 526, 527,
		1, 0, 0, 0, 527, 525, 1, 0, 0, 0, 527, 528, 1, 0, 0, 0, 528, 529, 1, 0,
		0, 0, 529, 530, 6, 32, -1, 0, 530, 65, 1, 0, 0, 0, 531, 532, 5, 29, 0,
		0, 532, 533, 5, 34, 0, 0, 533, 534, 3, 98, 49, 0, 534, 535, 5, 33, 0, 0,
		535, 536, 6, 33, -1, 0, 536, 543, 1, 0, 0, 0, 537, 538, 5, 29, 0, 0, 538,
		539, 5, 34, 0, 0, 539, 540, 3, 98, 49, 0, 540, 541, 6, 33, -1, 0, 541,
		543, 1, 0, 0, 0, 542, 531, 1, 0, 0, 0, 542, 537, 1, 0, 0, 0, 543, 67, 1,
		0, 0, 0, 544, 545, 5, 29, 0, 0, 545, 546, 5, 49, 0, 0, 546, 553, 5, 50,
		0, 0, 547, 548, 5, 29, 0, 0, 548, 549, 5, 49, 0, 0, 549, 550, 3, 100, 50,
		0, 550, 551, 5, 50, 0, 0, 551, 553, 1, 0, 0, 0, 552, 544, 1, 0, 0, 0, 552,
		547, 1, 0, 0, 0, 553, 69, 1, 0, 0, 0, 554, 555, 5, 29, 0, 0, 555, 556,
		5, 49, 0, 0, 556, 563, 5, 50, 0, 0, 557, 558, 5, 29, 0, 0, 558, 559, 5,
		49, 0, 0, 559, 560, 3, 100, 50, 0, 560, 561, 5, 50, 0, 0, 561, 563, 1,
		0, 0, 0, 562, 554, 1, 0, 0, 0, 562, 557, 1, 0, 0, 0, 563, 71, 1, 0, 0,
		0, 564, 565, 5, 16, 0, 0, 565, 566, 5, 29, 0, 0, 566, 567, 5, 51, 0, 0,
		567, 568, 3, 74, 37, 0, 568, 569, 5, 52, 0, 0, 569, 73, 1, 0, 0, 0, 570,
		572, 3, 76, 38, 0, 571, 570, 1, 0, 0, 0, 572, 573, 1, 0, 0, 0, 573, 571,
		1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 75, 1, 0, 0, 0, 575, 576, 5, 29,
		0, 0, 576, 577, 5, 34, 0, 0, 577, 578, 3, 98, 49, 0, 578, 579, 5, 33, 0,
		0, 579, 584, 1, 0, 0, 0, 580, 581, 5, 29, 0, 0, 581, 582, 5, 34, 0, 0,
		582, 584, 3, 98, 49, 0, 583, 575, 1, 0, 0, 0, 583, 580, 1, 0, 0, 0, 584,
		77, 1, 0, 0, 0, 585, 586, 5, 29, 0, 0, 586, 587, 3, 80, 40, 0, 587, 79,
		1, 0, 0, 0, 588, 590, 3, 82, 41, 0, 589, 588, 1, 0, 0, 0, 590, 591, 1,
		0, 0, 0, 591, 589, 1, 0, 0, 0, 591, 592, 1, 0, 0, 0, 592, 81, 1, 0, 0,
		0, 593, 594, 5, 53, 0, 0, 594, 595, 3, 104, 52, 0, 595, 596, 5, 54, 0,
		0, 596, 83, 1, 0, 0, 0, 597, 598, 5, 2, 0, 0, 598, 599, 5, 29, 0, 0, 599,
		600, 5, 35, 0, 0, 600, 601, 5, 29, 0, 0, 601, 602, 5, 51, 0, 0, 602, 603,
		3, 86, 43, 0, 603, 604, 5, 52, 0, 0, 604, 614, 1, 0, 0, 0, 605, 606, 5,
		3, 0, 0, 606, 607, 5, 29, 0, 0, 607, 608, 5, 35, 0, 0, 608, 609, 5, 29,
		0, 0, 609, 610, 5, 51, 0, 0, 610, 611, 3, 86, 43, 0, 611, 612, 5, 52, 0,
		0, 612, 614, 1, 0, 0, 0, 613, 597, 1, 0, 0, 0, 613, 605, 1, 0, 0, 0, 614,
		85, 1, 0, 0, 0, 615, 617, 3, 88, 44, 0, 616, 615, 1, 0, 0, 0, 617, 618,
		1, 0, 0, 0, 618, 616, 1, 0, 0, 0, 618, 619, 1, 0, 0, 0, 619, 87, 1, 0,
		0, 0, 620, 621, 5, 29, 0, 0, 621, 622, 5, 34, 0, 0, 622, 629, 3, 104, 52,
		0, 623, 624, 5, 29, 0, 0, 624, 625, 5, 34, 0, 0, 625, 626, 3, 104, 52,
		0, 626, 627, 5, 33, 0, 0, 627, 629, 1, 0, 0, 0, 628, 620, 1, 0, 0, 0, 628,
		623, 1, 0, 0, 0, 629, 89, 1, 0, 0, 0, 630, 631, 5, 29, 0, 0, 631, 632,
		3, 92, 46, 0, 632, 91, 1, 0, 0, 0, 633, 635, 3, 94, 47, 0, 634, 633, 1,
		0, 0, 0, 635, 636, 1, 0, 0, 0, 636, 634, 1, 0, 0, 0, 636, 637, 1, 0, 0,
		0, 637, 93, 1, 0, 0, 0, 638, 639, 5, 31, 0, 0, 639, 640, 5, 29, 0, 0, 640,
		95, 1, 0, 0, 0, 641, 642, 5, 29, 0, 0, 642, 643, 3, 92, 46, 0, 643, 644,
		5, 35, 0, 0, 644, 645, 3, 104, 52, 0, 645, 97, 1, 0, 0, 0, 646, 647, 5,
		17, 0, 0, 647, 657, 6, 49, -1, 0, 648, 649, 5, 18, 0, 0, 649, 657, 6, 49,
		-1, 0, 650, 651, 5, 22, 0, 0, 651, 657, 6, 49, -1, 0, 652, 653, 5, 20,
		0, 0, 653, 657, 6, 49, -1, 0, 654, 655, 5, 21, 0, 0, 655, 657, 6, 49, -1,
		0, 656, 646, 1, 0, 0, 0, 656, 648, 1, 0, 0, 0, 656, 650, 1, 0, 0, 0, 656,
		652, 1, 0, 0, 0, 656, 654, 1, 0, 0, 0, 657, 99, 1, 0, 0, 0, 658, 660, 3,
		102, 51, 0, 659, 658, 1, 0, 0, 0, 660, 661, 1, 0, 0, 0, 661, 659, 1, 0,
		0, 0, 661, 662, 1, 0, 0, 0, 662, 663, 1, 0, 0, 0, 663, 664, 6, 50, -1,
		0, 664, 101, 1, 0, 0, 0, 665, 666, 3, 104, 52, 0, 666, 667, 5, 33, 0, 0,
		667, 668, 6, 51, -1, 0, 668, 673, 1, 0, 0, 0, 669, 670, 3, 104, 52, 0,
		670, 671, 6, 51, -1, 0, 671, 673, 1, 0, 0, 0, 672, 665, 1, 0, 0, 0, 672,
		669, 1, 0, 0, 0, 673, 103, 1, 0, 0, 0, 674, 675, 3, 106, 53, 0, 675, 676,
		6, 52, -1, 0, 676, 684, 1, 0, 0, 0, 677, 678, 3, 110, 55, 0, 678, 679,
		6, 52, -1, 0, 679, 684, 1, 0, 0, 0, 680, 681, 3, 108, 54, 0, 681, 682,
		6, 52, -1, 0, 682, 684, 1, 0, 0, 0, 683, 674, 1, 0, 0, 0, 683, 677, 1,
		0, 0, 0, 683, 680, 1, 0, 0, 0, 684, 105, 1, 0, 0, 0, 685, 686, 6, 53, -1,
		0, 686, 687, 5, 59, 0, 0, 687, 688, 3, 106, 53, 3, 688, 689, 6, 53, -1,
		0, 689, 694, 1, 0, 0, 0, 690, 691, 3, 108, 54, 0, 691, 692, 6, 53, -1,
		0, 692, 694, 1, 0, 0, 0, 693, 685, 1, 0, 0, 0, 693, 690, 1, 0, 0, 0, 694,
		702, 1, 0, 0, 0, 695, 696, 10, 2, 0, 0, 696, 697, 7, 0, 0, 0, 697, 698,
		3, 106, 53, 3, 698, 699, 6, 53, -1, 0, 699, 701, 1, 0, 0, 0, 700, 695,
		1, 0, 0, 0, 701, 704, 1, 0, 0, 0, 702, 700, 1, 0, 0, 0, 702, 703, 1, 0,
		0, 0, 703, 107, 1, 0, 0, 0, 704, 702, 1, 0, 0, 0, 705, 706, 6, 54, -1,
		0, 706, 707, 3, 110, 55, 0, 707, 708, 6, 54, -1, 0, 708, 716, 1, 0, 0,
		0, 709, 710, 10, 2, 0, 0, 710, 711, 7, 1, 0, 0, 711, 712, 3, 108, 54, 3,
		712, 713, 6, 54, -1, 0, 713, 715, 1, 0, 0, 0, 714, 709, 1, 0, 0, 0, 715,
		718, 1, 0, 0, 0, 716, 714, 1, 0, 0, 0, 716, 717, 1, 0, 0, 0, 717, 109,
		1, 0, 0, 0, 718, 716, 1, 0, 0, 0, 719, 720, 6, 55, -1, 0, 720, 721, 5,
		48, 0, 0, 721, 722, 3, 110, 55, 5, 722, 723, 6, 55, -1, 0, 723, 733, 1,
		0, 0, 0, 724, 725, 3, 112, 56, 0, 725, 726, 6, 55, -1, 0, 726, 733, 1,
		0, 0, 0, 727, 728, 5, 49, 0, 0, 728, 729, 3, 104, 52, 0, 729, 730, 5, 50,
		0, 0, 730, 731, 6, 55, -1, 0, 731, 733, 1, 0, 0, 0, 732, 719, 1, 0, 0,
		0, 732, 724, 1, 0, 0, 0, 732, 727, 1, 0, 0, 0, 733, 746, 1, 0, 0, 0, 734,
		735, 10, 4, 0, 0, 735, 736, 7, 2, 0, 0, 736, 737, 3, 110, 55, 5, 737, 738,
		6, 55, -1, 0, 738, 745, 1, 0, 0, 0, 739, 740, 10, 3, 0, 0, 740, 741, 7,
		3, 0, 0, 741, 742, 3, 110, 55, 4, 742, 743, 6, 55, -1, 0, 743, 745, 1,
		0, 0, 0, 744, 734, 1, 0, 0, 0, 744, 739, 1, 0, 0, 0, 745, 748, 1, 0, 0,
		0, 746, 744, 1, 0, 0, 0, 746, 747, 1, 0, 0, 0, 747, 111, 1, 0, 0, 0, 748,
		746, 1, 0, 0, 0, 749, 750, 3, 114, 57, 0, 750, 751, 6, 56, -1, 0, 751,
		113, 1, 0, 0, 0, 752, 753, 5, 24, 0, 0, 753, 775, 6, 57, -1, 0, 754, 755,
		5, 25, 0, 0, 755, 775, 6, 57, -1, 0, 756, 757, 5, 27, 0, 0, 757, 775, 6,
		57, -1, 0, 758, 759, 5, 28, 0, 0, 759, 775, 6, 57, -1, 0, 760, 761, 5,
		26, 0, 0, 761, 775, 6, 57, -1, 0, 762, 775, 3, 70, 35, 0, 763, 775, 3,
		90, 45, 0, 764, 775, 3, 78, 39, 0, 765, 766, 5, 29, 0, 0, 766, 775, 6,
		57, -1, 0, 767, 768, 3, 16, 8, 0, 768, 769, 6, 57, -1, 0, 769, 775, 1,
		0, 0, 0, 770, 775, 3, 36, 18, 0, 771, 772, 3, 54, 27, 0, 772, 773, 6, 57,
		-1, 0, 773, 775, 1, 0, 0, 0, 774, 752, 1, 0, 0, 0, 774, 754, 1, 0, 0, 0,
		774, 756, 1, 0, 0, 0, 774, 758, 1, 0, 0, 0, 774, 760, 1, 0, 0, 0, 774,
		762, 1, 0, 0, 0, 774, 763, 1, 0, 0, 0, 774, 764, 1, 0, 0, 0, 774, 765,
		1, 0, 0, 0, 774, 767, 1, 0, 0, 0, 774, 770, 1, 0, 0, 0, 774, 771, 1, 0,
		0, 0, 775, 115, 1, 0, 0, 0, 46, 123, 161, 183, 225, 259, 264, 296, 301,
		319, 324, 337, 342, 348, 363, 368, 383, 388, 400, 405, 411, 424, 453, 472,
		522, 527, 542, 552, 562, 573, 583, 591, 613, 618, 628, 636, 656, 661, 672,
		683, 693, 702, 716, 732, 744, 746, 774,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftgrammarInit initializes any static state used to implement Swiftgrammar. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftgrammar(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftgrammarInit() {
	staticData := &SwiftgrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftgrammar produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftgrammar(input antlr.TokenStream) *Swiftgrammar {
	SwiftgrammarInit()
	this := new(Swiftgrammar)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftgrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Swiftgrammar.g4"

	return this
}

// Swiftgrammar tokens.
const (
	SwiftgrammarEOF            = antlr.TokenEOF
	SwiftgrammarPRINT          = 1
	SwiftgrammarVAR            = 2
	SwiftgrammarLET            = 3
	SwiftgrammarIF             = 4
	SwiftgrammarELSE           = 5
	SwiftgrammarFOR            = 6
	SwiftgrammarWHILE          = 7
	SwiftgrammarSWITCH         = 8
	SwiftgrammarCASE           = 9
	SwiftgrammarDEFAULT        = 10
	SwiftgrammarIN             = 11
	SwiftgrammarBREAK          = 12
	SwiftgrammarCONTINUE       = 13
	SwiftgrammarRETURN         = 14
	SwiftgrammarFUNC           = 15
	SwiftgrammarSTRUCT         = 16
	SwiftgrammarR_INT          = 17
	SwiftgrammarR_FLOAT        = 18
	SwiftgrammarR_DOUBLE       = 19
	SwiftgrammarR_BOOL         = 20
	SwiftgrammarR_CHAR         = 21
	SwiftgrammarR_STRING       = 22
	SwiftgrammarTRUE           = 23
	SwiftgrammarNUMBER         = 24
	SwiftgrammarDOUBLE         = 25
	SwiftgrammarCHAR           = 26
	SwiftgrammarSTRING         = 27
	SwiftgrammarBOOLEAN        = 28
	SwiftgrammarID             = 29
	SwiftgrammarTK_TRIPLEPUNTO = 30
	SwiftgrammarTK_PUNTO       = 31
	SwiftgrammarTK_PUNTOCOMA   = 32
	SwiftgrammarTK_COMA        = 33
	SwiftgrammarTK_DOSPUNTOS   = 34
	SwiftgrammarTK_IGUAL       = 35
	SwiftgrammarTK_IGUALIGUAL  = 36
	SwiftgrammarTK_MAYORIGUAL  = 37
	SwiftgrammarTK_IGUALMAYOR  = 38
	SwiftgrammarTK_MENOSMAYOR  = 39
	SwiftgrammarTK_MENORIGUAL  = 40
	SwiftgrammarTK_DIFIGUAL    = 41
	SwiftgrammarTK_MAYOR       = 42
	SwiftgrammarTK_MENOR       = 43
	SwiftgrammarTK_MULT        = 44
	SwiftgrammarTK_DIV         = 45
	SwiftgrammarTK_MODULO      = 46
	SwiftgrammarTK_MAS         = 47
	SwiftgrammarTK_MENOS       = 48
	SwiftgrammarTK_PARA        = 49
	SwiftgrammarTK_PARC        = 50
	SwiftgrammarTK_LLAVEA      = 51
	SwiftgrammarTK_LLAVEC      = 52
	SwiftgrammarTK_CORA        = 53
	SwiftgrammarTK_CORC        = 54
	SwiftgrammarTK_AND         = 55
	SwiftgrammarTK_AMPERSAND   = 56
	SwiftgrammarTK_OR          = 57
	SwiftgrammarTK_BARRA       = 58
	SwiftgrammarTK_NOT         = 59
	SwiftgrammarWHITESPACE     = 60
	SwiftgrammarTK_MULTI       = 61
	SwiftgrammarTK_LINE        = 62
)

// Swiftgrammar rules.
const (
	SwiftgrammarRULE_start                           = 0
	SwiftgrammarRULE_instrucciones                   = 1
	SwiftgrammarRULE_instruccion                     = 2
	SwiftgrammarRULE_instruccion_println             = 3
	SwiftgrammarRULE_instruccion_declaracion         = 4
	SwiftgrammarRULE_instruccion_asignacion          = 5
	SwiftgrammarRULE_instruccion_if                  = 6
	SwiftgrammarRULE_instr_else_if                   = 7
	SwiftgrammarRULE_instruccion_ternario            = 8
	SwiftgrammarRULE_instr_else_if_ternario          = 9
	SwiftgrammarRULE_instruccion_switch              = 10
	SwiftgrammarRULE_list_case                       = 11
	SwiftgrammarRULE_instruccion_case                = 12
	SwiftgrammarRULE_list_expre_case                 = 13
	SwiftgrammarRULE_block_case                      = 14
	SwiftgrammarRULE_block_instr_switch              = 15
	SwiftgrammarRULE_instr_default                   = 16
	SwiftgrammarRULE_block_default                   = 17
	SwiftgrammarRULE_instruccion_switch_ternario     = 18
	SwiftgrammarRULE_list_case_ternario              = 19
	SwiftgrammarRULE_instr_case_ter                  = 20
	SwiftgrammarRULE_list_expre_case_ter             = 21
	SwiftgrammarRULE_block_case_ter                  = 22
	SwiftgrammarRULE_instr_default_ter               = 23
	SwiftgrammarRULE_instruccion_while               = 24
	SwiftgrammarRULE_instruccion_for_in              = 25
	SwiftgrammarRULE_instruccion_while_true          = 26
	SwiftgrammarRULE_instruccion_while_true_ternario = 27
	SwiftgrammarRULE_instruccion_break               = 28
	SwiftgrammarRULE_instruccion_continue            = 29
	SwiftgrammarRULE_instruccion_return              = 30
	SwiftgrammarRULE_instruccion_func                = 31
	SwiftgrammarRULE_list_function_parameters        = 32
	SwiftgrammarRULE_block_parameters_fn             = 33
	SwiftgrammarRULE_instruccion_llamada             = 34
	SwiftgrammarRULE_instr_llamada_expre             = 35
	SwiftgrammarRULE_instruccion_structs_declaracion = 36
	SwiftgrammarRULE_list_struct_parameters          = 37
	SwiftgrammarRULE_block_structs_parameters        = 38
	SwiftgrammarRULE_instr_arrays_identifier         = 39
	SwiftgrammarRULE_list_arrays_parameters_id       = 40
	SwiftgrammarRULE_block_arrays_identifier         = 41
	SwiftgrammarRULE_instr_structs_declaration       = 42
	SwiftgrammarRULE_list_struct_parameters_decla    = 43
	SwiftgrammarRULE_block_structs_parameters_decla  = 44
	SwiftgrammarRULE_instr_structs_identifier        = 45
	SwiftgrammarRULE_list_struct_parameters_id       = 46
	SwiftgrammarRULE_block_structs_identifier        = 47
	SwiftgrammarRULE_instr_structs_assignment        = 48
	SwiftgrammarRULE_instruccion_tipo                = 49
	SwiftgrammarRULE_list_expression                 = 50
	SwiftgrammarRULE_block_list_expression           = 51
	SwiftgrammarRULE_expressions                     = 52
	SwiftgrammarRULE_expre_log                       = 53
	SwiftgrammarRULE_expre_rel                       = 54
	SwiftgrammarRULE_expre_arit                      = 55
	SwiftgrammarRULE_expre_valor                     = 56
	SwiftgrammarRULE_primitivo                       = 57
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// Getter signatures
	Instrucciones() IInstruccionesContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *StartContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Swiftgrammar) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftgrammarRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	{
		p.SetState(117)
		p.Match(SwiftgrammarEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_instrucciones().GetL()

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetList returns the list rule context list.
	GetList() []IInstruccionContext

	// SetList sets the list rule context list.
	SetList([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// Getter signatures
	AllInstruccion() []IInstruccionContext
	Instruccion(i int) IInstruccionContext

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	list         []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instrucciones
	return p
}

func InitEmptyInstruccionesContext(p *InstruccionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instrucciones
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetList() []IInstruccionContext { return s.list }

func (s *InstruccionesContext) SetList(v []IInstruccionContext) { s.list = v }

func (s *InstruccionesContext) GetL() *arrayList.List { return s.l }

func (s *InstruccionesContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionContext); ok {
			tst[i] = t.(IInstruccionContext)
			i++
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Swiftgrammar) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftgrammarRULE_instrucciones)

	localctx.(*InstruccionesContext).l = arrayList.New()

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		//goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536998366) != 0) {
		{
			p.SetState(120)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).list = append(localctx.(*InstruccionesContext).list, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			//goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetList()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion_println returns the _instruccion_println rule contexts.
	Get_instruccion_println() IInstruccion_printlnContext

	// Get_instruccion_declaracion returns the _instruccion_declaracion rule contexts.
	Get_instruccion_declaracion() IInstruccion_declaracionContext

	// Get_instruccion_asignacion returns the _instruccion_asignacion rule contexts.
	Get_instruccion_asignacion() IInstruccion_asignacionContext

	// Get_instruccion_if returns the _instruccion_if rule contexts.
	Get_instruccion_if() IInstruccion_ifContext

	// Get_instruccion_for_in returns the _instruccion_for_in rule contexts.
	Get_instruccion_for_in() IInstruccion_for_inContext

	// Get_instruccion_while returns the _instruccion_while rule contexts.
	Get_instruccion_while() IInstruccion_whileContext

	// Get_instruccion_while_true returns the _instruccion_while_true rule contexts.
	Get_instruccion_while_true() IInstruccion_while_trueContext

	// Get_instruccion_func returns the _instruccion_func rule contexts.
	Get_instruccion_func() IInstruccion_funcContext

	// Get_instruccion_return returns the _instruccion_return rule contexts.
	Get_instruccion_return() IInstruccion_returnContext

	// Set_instruccion_println sets the _instruccion_println rule contexts.
	Set_instruccion_println(IInstruccion_printlnContext)

	// Set_instruccion_declaracion sets the _instruccion_declaracion rule contexts.
	Set_instruccion_declaracion(IInstruccion_declaracionContext)

	// Set_instruccion_asignacion sets the _instruccion_asignacion rule contexts.
	Set_instruccion_asignacion(IInstruccion_asignacionContext)

	// Set_instruccion_if sets the _instruccion_if rule contexts.
	Set_instruccion_if(IInstruccion_ifContext)

	// Set_instruccion_for_in sets the _instruccion_for_in rule contexts.
	Set_instruccion_for_in(IInstruccion_for_inContext)

	// Set_instruccion_while sets the _instruccion_while rule contexts.
	Set_instruccion_while(IInstruccion_whileContext)

	// Set_instruccion_while_true sets the _instruccion_while_true rule contexts.
	Set_instruccion_while_true(IInstruccion_while_trueContext)

	// Set_instruccion_func sets the _instruccion_func rule contexts.
	Set_instruccion_func(IInstruccion_funcContext)

	// Set_instruccion_return sets the _instruccion_return rule contexts.
	Set_instruccion_return(IInstruccion_returnContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	Instruccion_println() IInstruccion_printlnContext
	Instruccion_structs_declaracion() IInstruccion_structs_declaracionContext
	Instruccion_declaracion() IInstruccion_declaracionContext
	Instruccion_asignacion() IInstruccion_asignacionContext
	Instr_structs_assignment() IInstr_structs_assignmentContext
	Instruccion_if() IInstruccion_ifContext
	Instruccion_for_in() IInstruccion_for_inContext
	Instruccion_while() IInstruccion_whileContext
	Instruccion_while_true() IInstruccion_while_trueContext
	Instruccion_switch() IInstruccion_switchContext
	Instruccion_break() IInstruccion_breakContext
	Instruccion_continue() IInstruccion_continueContext
	Instruccion_func() IInstruccion_funcContext
	Instruccion_llamada() IInstruccion_llamadaContext
	Instruccion_return() IInstruccion_returnContext
	Instr_structs_declaration() IInstr_structs_declarationContext

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	antlr.BaseParserRuleContext
	parser                   antlr.Parser
	instr                    interfaces.Instruction
	_instruccion_println     IInstruccion_printlnContext
	_instruccion_declaracion IInstruccion_declaracionContext
	_instruccion_asignacion  IInstruccion_asignacionContext
	_instruccion_if          IInstruccion_ifContext
	_instruccion_for_in      IInstruccion_for_inContext
	_instruccion_while       IInstruccion_whileContext
	_instruccion_while_true  IInstruccion_while_trueContext
	_instruccion_func        IInstruccion_funcContext
	_instruccion_return      IInstruccion_returnContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion
	return p
}

func InitEmptyInstruccionContext(p *InstruccionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_instruccion_println() IInstruccion_printlnContext {
	return s._instruccion_println
}

func (s *InstruccionContext) Get_instruccion_declaracion() IInstruccion_declaracionContext {
	return s._instruccion_declaracion
}

func (s *InstruccionContext) Get_instruccion_asignacion() IInstruccion_asignacionContext {
	return s._instruccion_asignacion
}

func (s *InstruccionContext) Get_instruccion_if() IInstruccion_ifContext { return s._instruccion_if }

func (s *InstruccionContext) Get_instruccion_for_in() IInstruccion_for_inContext {
	return s._instruccion_for_in
}

func (s *InstruccionContext) Get_instruccion_while() IInstruccion_whileContext {
	return s._instruccion_while
}

func (s *InstruccionContext) Get_instruccion_while_true() IInstruccion_while_trueContext {
	return s._instruccion_while_true
}

func (s *InstruccionContext) Get_instruccion_func() IInstruccion_funcContext {
	return s._instruccion_func
}

func (s *InstruccionContext) Get_instruccion_return() IInstruccion_returnContext {
	return s._instruccion_return
}

func (s *InstruccionContext) Set_instruccion_println(v IInstruccion_printlnContext) {
	s._instruccion_println = v
}

func (s *InstruccionContext) Set_instruccion_declaracion(v IInstruccion_declaracionContext) {
	s._instruccion_declaracion = v
}

func (s *InstruccionContext) Set_instruccion_asignacion(v IInstruccion_asignacionContext) {
	s._instruccion_asignacion = v
}

func (s *InstruccionContext) Set_instruccion_if(v IInstruccion_ifContext) { s._instruccion_if = v }

func (s *InstruccionContext) Set_instruccion_for_in(v IInstruccion_for_inContext) {
	s._instruccion_for_in = v
}

func (s *InstruccionContext) Set_instruccion_while(v IInstruccion_whileContext) {
	s._instruccion_while = v
}

func (s *InstruccionContext) Set_instruccion_while_true(v IInstruccion_while_trueContext) {
	s._instruccion_while_true = v
}

func (s *InstruccionContext) Set_instruccion_func(v IInstruccion_funcContext) {
	s._instruccion_func = v
}

func (s *InstruccionContext) Set_instruccion_return(v IInstruccion_returnContext) {
	s._instruccion_return = v
}

func (s *InstruccionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *InstruccionContext) Instruccion_println() IInstruccion_printlnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_printlnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_printlnContext)
}

func (s *InstruccionContext) Instruccion_structs_declaracion() IInstruccion_structs_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_structs_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_structs_declaracionContext)
}

func (s *InstruccionContext) Instruccion_declaracion() IInstruccion_declaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_declaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_declaracionContext)
}

func (s *InstruccionContext) Instruccion_asignacion() IInstruccion_asignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_asignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_asignacionContext)
}

func (s *InstruccionContext) Instr_structs_assignment() IInstr_structs_assignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_structs_assignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_structs_assignmentContext)
}

func (s *InstruccionContext) Instruccion_if() IInstruccion_ifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_ifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_ifContext)
}

func (s *InstruccionContext) Instruccion_for_in() IInstruccion_for_inContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_for_inContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_for_inContext)
}

func (s *InstruccionContext) Instruccion_while() IInstruccion_whileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_whileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_whileContext)
}

func (s *InstruccionContext) Instruccion_while_true() IInstruccion_while_trueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_while_trueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_while_trueContext)
}

func (s *InstruccionContext) Instruccion_switch() IInstruccion_switchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_switchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_switchContext)
}

func (s *InstruccionContext) Instruccion_break() IInstruccion_breakContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_breakContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_breakContext)
}

func (s *InstruccionContext) Instruccion_continue() IInstruccion_continueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_continueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_continueContext)
}

func (s *InstruccionContext) Instruccion_func() IInstruccion_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_funcContext)
}

func (s *InstruccionContext) Instruccion_llamada() IInstruccion_llamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_llamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_llamadaContext)
}

func (s *InstruccionContext) Instruccion_return() IInstruccion_returnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_returnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_returnContext)
}

func (s *InstruccionContext) Instr_structs_declaration() IInstr_structs_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_structs_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_structs_declarationContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Swiftgrammar) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftgrammarRULE_instruccion)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)

			var _x = p.Instruccion_println()

			localctx.(*InstruccionContext)._instruccion_println = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instruccion_println().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Instruccion_structs_declaracion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)

			var _x = p.Instruccion_declaracion()

			localctx.(*InstruccionContext)._instruccion_declaracion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instruccion_declaracion().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(134)

			var _x = p.Instruccion_asignacion()

			localctx.(*InstruccionContext)._instruccion_asignacion = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instruccion_asignacion().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(137)
			p.Instr_structs_assignment()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(138)

			var _x = p.Instruccion_if()

			localctx.(*InstruccionContext)._instruccion_if = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instruccion_if().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(141)

			var _x = p.Instruccion_for_in()

			localctx.(*InstruccionContext)._instruccion_for_in = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instruccion_for_in().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(144)

			var _x = p.Instruccion_while()

			localctx.(*InstruccionContext)._instruccion_while = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instruccion_while().GetInstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(147)

			var _x = p.Instruccion_while_true()

			localctx.(*InstruccionContext)._instruccion_while_true = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instruccion_while_true().GetInstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(150)
			p.Instruccion_switch()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(151)
			p.Instruccion_break()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(152)
			p.Instruccion_continue()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(153)

			var _x = p.Instruccion_func()

			localctx.(*InstruccionContext)._instruccion_func = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instruccion_func().GetInstr()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(156)
			p.Instruccion_llamada()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(157)

			var _x = p.Instruccion_return()

			localctx.(*InstruccionContext)._instruccion_return = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_instruccion_return().GetInstr()

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(160)
			p.Instr_structs_declaration()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_printlnContext is an interface to support dynamic dispatch.
type IInstruccion_printlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_list_expression returns the _list_expression rule contexts.
	Get_list_expression() IList_expressionContext

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_list_expression sets the _list_expression rule contexts.
	Set_list_expression(IList_expressionContext)

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	TK_PARA() antlr.TerminalNode
	Primitivo() IPrimitivoContext
	TK_PARC() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TK_COMA() antlr.TerminalNode
	List_expression() IList_expressionContext
	Expressions() IExpressionsContext

	// IsInstruccion_printlnContext differentiates from other interfaces.
	IsInstruccion_printlnContext()
}

type Instruccion_printlnContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Instruction
	_PRINT           antlr.Token
	_primitivo       IPrimitivoContext
	_STRING          antlr.Token
	_list_expression IList_expressionContext
	_expressions     IExpressionsContext
}

func NewEmptyInstruccion_printlnContext() *Instruccion_printlnContext {
	var p = new(Instruccion_printlnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_println
	return p
}

func InitEmptyInstruccion_printlnContext(p *Instruccion_printlnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_println
}

func (*Instruccion_printlnContext) IsInstruccion_printlnContext() {}

func NewInstruccion_printlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_printlnContext {
	var p = new(Instruccion_printlnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_println

	return p
}

func (s *Instruccion_printlnContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_printlnContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *Instruccion_printlnContext) Get_STRING() antlr.Token { return s._STRING }

func (s *Instruccion_printlnContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *Instruccion_printlnContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *Instruccion_printlnContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Instruccion_printlnContext) Get_list_expression() IList_expressionContext {
	return s._list_expression
}

func (s *Instruccion_printlnContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instruccion_printlnContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Instruccion_printlnContext) Set_list_expression(v IList_expressionContext) {
	s._list_expression = v
}

func (s *Instruccion_printlnContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instruccion_printlnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_printlnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_printlnContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarPRINT, 0)
}

func (s *Instruccion_printlnContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARA, 0)
}

func (s *Instruccion_printlnContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Instruccion_printlnContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARC, 0)
}

func (s *Instruccion_printlnContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarSTRING, 0)
}

func (s *Instruccion_printlnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_COMA, 0)
}

func (s *Instruccion_printlnContext) List_expression() IList_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Instruccion_printlnContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_printlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_printlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_printlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_println(s)
	}
}

func (s *Instruccion_printlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_println(s)
	}
}

func (p *Swiftgrammar) Instruccion_println() (localctx IInstruccion_printlnContext) {
	localctx = NewInstruccion_printlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftgrammarRULE_instruccion_println)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)

			var _m = p.Match(SwiftgrammarPRINT)

			localctx.(*Instruccion_printlnContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)

			var _x = p.Primitivo()

			localctx.(*Instruccion_printlnContext)._primitivo = _x
		}
		{
			p.SetState(166)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_printlnContext).instr = instruction.NewPrintln(nil, localctx.(*Instruccion_printlnContext).Get_primitivo().GetP(), "-1", (func() int {
			if localctx.(*Instruccion_printlnContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_printlnContext).Get_PRINT().GetLine()
			}
		}()), localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)

			var _m = p.Match(SwiftgrammarPRINT)

			localctx.(*Instruccion_printlnContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)

			var _m = p.Match(SwiftgrammarSTRING)

			localctx.(*Instruccion_printlnContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(SwiftgrammarTK_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)

			var _x = p.List_expression()

			localctx.(*Instruccion_printlnContext)._list_expression = _x
		}
		{
			p.SetState(174)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_printlnContext).instr = instruction.NewPrintln(localctx.(*Instruccion_printlnContext).Get_list_expression().GetL(), nil, (func() string {
			if localctx.(*Instruccion_printlnContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_printlnContext).Get_STRING().GetText()
			}
		}())[1:len((func() string {
			if localctx.(*Instruccion_printlnContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_printlnContext).Get_STRING().GetText()
			}
		}()))-1], (func() int {
			if localctx.(*Instruccion_printlnContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_printlnContext).Get_PRINT().GetLine()
			}
		}()), localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)

			var _m = p.Match(SwiftgrammarPRINT)

			localctx.(*Instruccion_printlnContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)

			var _x = p.Expressions()

			localctx.(*Instruccion_printlnContext)._expressions = _x
		}
		{
			p.SetState(180)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_printlnContext).instr = instruction.NewPrintln(nil, localctx.(*Instruccion_printlnContext).Get_expressions().GetP(), "-1", (func() int {
			if localctx.(*Instruccion_printlnContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_printlnContext).Get_PRINT().GetLine()
			}
		}()), localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_declaracionContext is an interface to support dynamic dispatch.
type IInstruccion_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_instruccion_tipo returns the _instruccion_tipo rule contexts.
	Get_instruccion_tipo() IInstruccion_tipoContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_instruccion_tipo sets the _instruccion_tipo rule contexts.
	Set_instruccion_tipo(IInstruccion_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	TK_IGUAL() antlr.TerminalNode
	Expressions() IExpressionsContext
	TK_DOSPUNTOS() antlr.TerminalNode
	Instruccion_tipo() IInstruccion_tipoContext
	LET() antlr.TerminalNode

	// IsInstruccion_declaracionContext differentiates from other interfaces.
	IsInstruccion_declaracionContext()
}

type Instruccion_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser            antlr.Parser
	instr             interfaces.Instruction
	_VAR              antlr.Token
	_ID               antlr.Token
	_expressions      IExpressionsContext
	_instruccion_tipo IInstruccion_tipoContext
	_LET              antlr.Token
}

func NewEmptyInstruccion_declaracionContext() *Instruccion_declaracionContext {
	var p = new(Instruccion_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_declaracion
	return p
}

func InitEmptyInstruccion_declaracionContext(p *Instruccion_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_declaracion
}

func (*Instruccion_declaracionContext) IsInstruccion_declaracionContext() {}

func NewInstruccion_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_declaracionContext {
	var p = new(Instruccion_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_declaracion

	return p
}

func (s *Instruccion_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_declaracionContext) Get_VAR() antlr.Token { return s._VAR }

func (s *Instruccion_declaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *Instruccion_declaracionContext) Get_LET() antlr.Token { return s._LET }

func (s *Instruccion_declaracionContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *Instruccion_declaracionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instruccion_declaracionContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *Instruccion_declaracionContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instruccion_declaracionContext) Get_instruccion_tipo() IInstruccion_tipoContext {
	return s._instruccion_tipo
}

func (s *Instruccion_declaracionContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instruccion_declaracionContext) Set_instruccion_tipo(v IInstruccion_tipoContext) {
	s._instruccion_tipo = v
}

func (s *Instruccion_declaracionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_declaracionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_declaracionContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarVAR, 0)
}

func (s *Instruccion_declaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instruccion_declaracionContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_IGUAL, 0)
}

func (s *Instruccion_declaracionContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_declaracionContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DOSPUNTOS, 0)
}

func (s *Instruccion_declaracionContext) Instruccion_tipo() IInstruccion_tipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_tipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_tipoContext)
}

func (s *Instruccion_declaracionContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarLET, 0)
}

func (s *Instruccion_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_declaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_declaracion(s)
	}
}

func (s *Instruccion_declaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_declaracion(s)
	}
}

func (p *Swiftgrammar) Instruccion_declaracion() (localctx IInstruccion_declaracionContext) {
	localctx = NewInstruccion_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftgrammarRULE_instruccion_declaracion)
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)

			var _m = p.Match(SwiftgrammarVAR)

			localctx.(*Instruccion_declaracionContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

			var _x = p.Expressions()

			localctx.(*Instruccion_declaracionContext)._expressions = _x
		}
		localctx.(*Instruccion_declaracionContext).instr = instruction.NewDeclaration((func() string {
			if localctx.(*Instruccion_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instruccion_declaracionContext).Get_expressions().GetP(), true, (func() int {
			if localctx.(*Instruccion_declaracionContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_VAR().GetLine()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_VAR().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)

			var _m = p.Match(SwiftgrammarVAR)

			localctx.(*Instruccion_declaracionContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)

			var _x = p.Instruccion_tipo()

			localctx.(*Instruccion_declaracionContext)._instruccion_tipo = _x
		}
		localctx.(*Instruccion_declaracionContext).instr = instruction.NewDeclaration((func() string {
			if localctx.(*Instruccion_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_instruccion_tipo().GetTipo_exp(), nil, true, (func() int {
			if localctx.(*Instruccion_declaracionContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_VAR().GetLine()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_VAR().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)

			var _m = p.Match(SwiftgrammarVAR)

			localctx.(*Instruccion_declaracionContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)

			var _x = p.Instruccion_tipo()

			localctx.(*Instruccion_declaracionContext)._instruccion_tipo = _x
		}
		{
			p.SetState(201)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)

			var _x = p.Expressions()

			localctx.(*Instruccion_declaracionContext)._expressions = _x
		}
		localctx.(*Instruccion_declaracionContext).instr = instruction.NewDeclaration((func() string {
			if localctx.(*Instruccion_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_instruccion_tipo().GetTipo_exp(), localctx.(*Instruccion_declaracionContext).Get_expressions().GetP(), true, (func() int {
			if localctx.(*Instruccion_declaracionContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_VAR().GetLine()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_VAR().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(205)

			var _m = p.Match(SwiftgrammarLET)

			localctx.(*Instruccion_declaracionContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)

			var _x = p.Expressions()

			localctx.(*Instruccion_declaracionContext)._expressions = _x
		}
		localctx.(*Instruccion_declaracionContext).instr = instruction.NewDeclaration((func() string {
			if localctx.(*Instruccion_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_ID().GetText()
			}
		}()), interfaces.NULL, localctx.(*Instruccion_declaracionContext).Get_expressions().GetP(), false, (func() int {
			if localctx.(*Instruccion_declaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_LET().GetLine()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_LET().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(211)

			var _m = p.Match(SwiftgrammarLET)

			localctx.(*Instruccion_declaracionContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)

			var _x = p.Instruccion_tipo()

			localctx.(*Instruccion_declaracionContext)._instruccion_tipo = _x
		}
		localctx.(*Instruccion_declaracionContext).instr = instruction.NewDeclaration((func() string {
			if localctx.(*Instruccion_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_instruccion_tipo().GetTipo_exp(), nil, false, (func() int {
			if localctx.(*Instruccion_declaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_LET().GetLine()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_LET().GetColumn())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(217)

			var _m = p.Match(SwiftgrammarLET)

			localctx.(*Instruccion_declaracionContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)

			var _x = p.Instruccion_tipo()

			localctx.(*Instruccion_declaracionContext)._instruccion_tipo = _x
		}
		{
			p.SetState(221)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)

			var _x = p.Expressions()

			localctx.(*Instruccion_declaracionContext)._expressions = _x
		}
		localctx.(*Instruccion_declaracionContext).instr = instruction.NewDeclaration((func() string {
			if localctx.(*Instruccion_declaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_instruccion_tipo().GetTipo_exp(), localctx.(*Instruccion_declaracionContext).Get_expressions().GetP(), false, (func() int {
			if localctx.(*Instruccion_declaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_declaracionContext).Get_LET().GetLine()
			}
		}()), localctx.(*Instruccion_declaracionContext).Get_LET().GetColumn())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_asignacionContext is an interface to support dynamic dispatch.
type IInstruccion_asignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	TK_IGUAL() antlr.TerminalNode
	Expressions() IExpressionsContext

	// IsInstruccion_asignacionContext differentiates from other interfaces.
	IsInstruccion_asignacionContext()
}

type Instruccion_asignacionContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_ID          antlr.Token
	_expressions IExpressionsContext
}

func NewEmptyInstruccion_asignacionContext() *Instruccion_asignacionContext {
	var p = new(Instruccion_asignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_asignacion
	return p
}

func InitEmptyInstruccion_asignacionContext(p *Instruccion_asignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_asignacion
}

func (*Instruccion_asignacionContext) IsInstruccion_asignacionContext() {}

func NewInstruccion_asignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_asignacionContext {
	var p = new(Instruccion_asignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_asignacion

	return p
}

func (s *Instruccion_asignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_asignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *Instruccion_asignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instruccion_asignacionContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instruccion_asignacionContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instruccion_asignacionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_asignacionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_asignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instruccion_asignacionContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_IGUAL, 0)
}

func (s *Instruccion_asignacionContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_asignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_asignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_asignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_asignacion(s)
	}
}

func (s *Instruccion_asignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_asignacion(s)
	}
}

func (p *Swiftgrammar) Instruccion_asignacion() (localctx IInstruccion_asignacionContext) {
	localctx = NewInstruccion_asignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftgrammarRULE_instruccion_asignacion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)

		var _m = p.Match(SwiftgrammarID)

		localctx.(*Instruccion_asignacionContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(SwiftgrammarTK_IGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)

		var _x = p.Expressions()

		localctx.(*Instruccion_asignacionContext)._expressions = _x
	}
	localctx.(*Instruccion_asignacionContext).instr = instruction.NewAssignment((func() string {
		if localctx.(*Instruccion_asignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Instruccion_asignacionContext).Get_ID().GetText()
		}
	}()), localctx.(*Instruccion_asignacionContext).Get_expressions().GetP(), (func() int {
		if localctx.(*Instruccion_asignacionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Instruccion_asignacionContext).Get_ID().GetLine()
		}
	}()), localctx.(*Instruccion_asignacionContext).Get_ID().GetColumn())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_ifContext is an interface to support dynamic dispatch.
type IInstruccion_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IInstruccionesContext

	// GetRight_instr returns the right_instr rule contexts.
	GetRight_instr() IInstruccionesContext

	// Get_instr_else_if returns the _instr_else_if rule contexts.
	Get_instr_else_if() IInstr_else_ifContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IInstruccionesContext)

	// SetRight_instr sets the right_instr rule contexts.
	SetRight_instr(IInstruccionesContext)

	// Set_instr_else_if sets the _instr_else_if rule contexts.
	Set_instr_else_if(IInstr_else_ifContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expressions() IExpressionsContext
	AllTK_LLAVEA() []antlr.TerminalNode
	TK_LLAVEA(i int) antlr.TerminalNode
	AllTK_LLAVEC() []antlr.TerminalNode
	TK_LLAVEC(i int) antlr.TerminalNode
	AllInstrucciones() []IInstruccionesContext
	Instrucciones(i int) IInstruccionesContext
	ELSE() antlr.TerminalNode
	Instr_else_if() IInstr_else_ifContext

	// IsInstruccion_ifContext differentiates from other interfaces.
	IsInstruccion_ifContext()
}

type Instruccion_ifContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_IF            antlr.Token
	_expressions   IExpressionsContext
	left_instr     IInstruccionesContext
	right_instr    IInstruccionesContext
	_instr_else_if IInstr_else_ifContext
}

func NewEmptyInstruccion_ifContext() *Instruccion_ifContext {
	var p = new(Instruccion_ifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_if
	return p
}

func InitEmptyInstruccion_ifContext(p *Instruccion_ifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_if
}

func (*Instruccion_ifContext) IsInstruccion_ifContext() {}

func NewInstruccion_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_ifContext {
	var p = new(Instruccion_ifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_if

	return p
}

func (s *Instruccion_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_ifContext) Get_IF() antlr.Token { return s._IF }

func (s *Instruccion_ifContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *Instruccion_ifContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instruccion_ifContext) GetLeft_instr() IInstruccionesContext { return s.left_instr }

func (s *Instruccion_ifContext) GetRight_instr() IInstruccionesContext { return s.right_instr }

func (s *Instruccion_ifContext) Get_instr_else_if() IInstr_else_ifContext { return s._instr_else_if }

func (s *Instruccion_ifContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instruccion_ifContext) SetLeft_instr(v IInstruccionesContext) { s.left_instr = v }

func (s *Instruccion_ifContext) SetRight_instr(v IInstruccionesContext) { s.right_instr = v }

func (s *Instruccion_ifContext) Set_instr_else_if(v IInstr_else_ifContext) { s._instr_else_if = v }

func (s *Instruccion_ifContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_ifContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_ifContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarIF, 0)
}

func (s *Instruccion_ifContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_ifContext) AllTK_LLAVEA() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammarTK_LLAVEA)
}

func (s *Instruccion_ifContext) TK_LLAVEA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, i)
}

func (s *Instruccion_ifContext) AllTK_LLAVEC() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammarTK_LLAVEC)
}

func (s *Instruccion_ifContext) TK_LLAVEC(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, i)
}

func (s *Instruccion_ifContext) AllInstrucciones() []IInstruccionesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccionesContext); ok {
			len++
		}
	}

	tst := make([]IInstruccionesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccionesContext); ok {
			tst[i] = t.(IInstruccionesContext)
			i++
		}
	}

	return tst
}

func (s *Instruccion_ifContext) Instrucciones(i int) IInstruccionesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instruccion_ifContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarELSE, 0)
}

func (s *Instruccion_ifContext) Instr_else_if() IInstr_else_ifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_else_ifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_else_ifContext)
}

func (s *Instruccion_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_if(s)
	}
}

func (s *Instruccion_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_if(s)
	}
}

func (p *Swiftgrammar) Instruccion_if() (localctx IInstruccion_ifContext) {
	localctx = NewInstruccion_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftgrammarRULE_instruccion_if)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)

			var _m = p.Match(SwiftgrammarIF)

			localctx.(*Instruccion_ifContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)

			var _x = p.Expressions()

			localctx.(*Instruccion_ifContext)._expressions = _x
		}
		{
			p.SetState(234)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_ifContext).left_instr = _x
		}
		{
			p.SetState(236)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_ifContext).instr = instruction.NewIf(localctx.(*Instruccion_ifContext).Get_expressions().GetP(), localctx.(*Instruccion_ifContext).GetLeft_instr().GetL(), nil, nil, (func() int {
			if localctx.(*Instruccion_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_ifContext).Get_IF().GetLine()
			}
		}()), localctx.(*Instruccion_ifContext).Get_IF().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)

			var _m = p.Match(SwiftgrammarIF)

			localctx.(*Instruccion_ifContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)

			var _x = p.Expressions()

			localctx.(*Instruccion_ifContext)._expressions = _x
		}
		{
			p.SetState(241)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_ifContext).left_instr = _x
		}
		{
			p.SetState(243)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(SwiftgrammarELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_ifContext).right_instr = _x
		}
		{
			p.SetState(247)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_ifContext).instr = instruction.NewIf(localctx.(*Instruccion_ifContext).Get_expressions().GetP(), localctx.(*Instruccion_ifContext).GetLeft_instr().GetL(), localctx.(*Instruccion_ifContext).GetRight_instr().GetL(), nil, (func() int {
			if localctx.(*Instruccion_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_ifContext).Get_IF().GetLine()
			}
		}()), localctx.(*Instruccion_ifContext).Get_IF().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(250)

			var _m = p.Match(SwiftgrammarIF)

			localctx.(*Instruccion_ifContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)

			var _x = p.Expressions()

			localctx.(*Instruccion_ifContext)._expressions = _x
		}
		{
			p.SetState(252)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_ifContext).left_instr = _x
		}
		{
			p.SetState(254)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(SwiftgrammarELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)

			var _x = p.Instr_else_if()

			localctx.(*Instruccion_ifContext)._instr_else_if = _x
		}
		localctx.(*Instruccion_ifContext).instr = instruction.NewIf(localctx.(*Instruccion_ifContext).Get_expressions().GetP(), localctx.(*Instruccion_ifContext).GetLeft_instr().GetL(), nil, localctx.(*Instruccion_ifContext).Get_instr_else_if().GetL(), (func() int {
			if localctx.(*Instruccion_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_ifContext).Get_IF().GetLine()
			}
		}()), localctx.(*Instruccion_ifContext).Get_IF().GetColumn())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_else_ifContext is an interface to support dynamic dispatch.
type IInstr_else_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion_if returns the _instruccion_if rule contexts.
	Get_instruccion_if() IInstruccion_ifContext

	// Set_instruccion_if sets the _instruccion_if rule contexts.
	Set_instruccion_if(IInstruccion_ifContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccion_ifContext

	// SetE sets the e rule context list.
	SetE([]IInstruccion_ifContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// Getter signatures
	AllInstruccion_if() []IInstruccion_ifContext
	Instruccion_if(i int) IInstruccion_ifContext

	// IsInstr_else_ifContext differentiates from other interfaces.
	IsInstr_else_ifContext()
}

type Instr_else_ifContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	l               *arrayList.List
	_instruccion_if IInstruccion_ifContext
	e               []IInstruccion_ifContext
}

func NewEmptyInstr_else_ifContext() *Instr_else_ifContext {
	var p = new(Instr_else_ifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_else_if
	return p
}

func InitEmptyInstr_else_ifContext(p *Instr_else_ifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_else_if
}

func (*Instr_else_ifContext) IsInstr_else_ifContext() {}

func NewInstr_else_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_else_ifContext {
	var p = new(Instr_else_ifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_else_if

	return p
}

func (s *Instr_else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_else_ifContext) Get_instruccion_if() IInstruccion_ifContext { return s._instruccion_if }

func (s *Instr_else_ifContext) Set_instruccion_if(v IInstruccion_ifContext) { s._instruccion_if = v }

func (s *Instr_else_ifContext) GetE() []IInstruccion_ifContext { return s.e }

func (s *Instr_else_ifContext) SetE(v []IInstruccion_ifContext) { s.e = v }

func (s *Instr_else_ifContext) GetL() *arrayList.List { return s.l }

func (s *Instr_else_ifContext) SetL(v *arrayList.List) { s.l = v }

func (s *Instr_else_ifContext) AllInstruccion_if() []IInstruccion_ifContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccion_ifContext); ok {
			len++
		}
	}

	tst := make([]IInstruccion_ifContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccion_ifContext); ok {
			tst[i] = t.(IInstruccion_ifContext)
			i++
		}
	}

	return tst
}

func (s *Instr_else_ifContext) Instruccion_if(i int) IInstruccion_ifContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_ifContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_ifContext)
}

func (s *Instr_else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_else_if(s)
	}
}

func (s *Instr_else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_else_if(s)
	}
}

func (p *Swiftgrammar) Instr_else_if() (localctx IInstr_else_ifContext) {
	localctx = NewInstr_else_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftgrammarRULE_instr_else_if)

	localctx.(*Instr_else_ifContext).l = arrayList.New()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		//goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
	//	goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(261)

				var _x = p.Instruccion_if()

				localctx.(*Instr_else_ifContext)._instruccion_if = _x
			}
			localctx.(*Instr_else_ifContext).e = append(localctx.(*Instr_else_ifContext).e, localctx.(*Instr_else_ifContext)._instruccion_if)

		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			//goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
		//	goto errorExit
		}
	}

	listInt := localctx.(*Instr_else_ifContext).GetE()
	for _, e := range listInt {
		localctx.(*Instr_else_ifContext).l.Add(e.GetInstr())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_ternarioContext is an interface to support dynamic dispatch.
type IInstruccion_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionsContext

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IExpressionsContext

	// GetRight_instr returns the right_instr rule contexts.
	GetRight_instr() IExpressionsContext

	// Get_instr_else_if_ternario returns the _instr_else_if_ternario rule contexts.
	Get_instr_else_if_ternario() IInstr_else_if_ternarioContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionsContext)

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IExpressionsContext)

	// SetRight_instr sets the right_instr rule contexts.
	SetRight_instr(IExpressionsContext)

	// Set_instr_else_if_ternario sets the _instr_else_if_ternario rule contexts.
	Set_instr_else_if_ternario(IInstr_else_if_ternarioContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	IF() antlr.TerminalNode
	AllTK_LLAVEA() []antlr.TerminalNode
	TK_LLAVEA(i int) antlr.TerminalNode
	AllTK_LLAVEC() []antlr.TerminalNode
	TK_LLAVEC(i int) antlr.TerminalNode
	AllExpressions() []IExpressionsContext
	Expressions(i int) IExpressionsContext
	ELSE() antlr.TerminalNode
	Instr_else_if_ternario() IInstr_else_if_ternarioContext

	// IsInstruccion_ternarioContext differentiates from other interfaces.
	IsInstruccion_ternarioContext()
}

type Instruccion_ternarioContext struct {
	antlr.BaseParserRuleContext
	parser                  antlr.Parser
	p                       interfaces.Expression
	_IF                     antlr.Token
	cond                    IExpressionsContext
	left_instr              IExpressionsContext
	right_instr             IExpressionsContext
	_instr_else_if_ternario IInstr_else_if_ternarioContext
}

func NewEmptyInstruccion_ternarioContext() *Instruccion_ternarioContext {
	var p = new(Instruccion_ternarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_ternario
	return p
}

func InitEmptyInstruccion_ternarioContext(p *Instruccion_ternarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_ternario
}

func (*Instruccion_ternarioContext) IsInstruccion_ternarioContext() {}

func NewInstruccion_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_ternarioContext {
	var p = new(Instruccion_ternarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_ternario

	return p
}

func (s *Instruccion_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_ternarioContext) Get_IF() antlr.Token { return s._IF }

func (s *Instruccion_ternarioContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *Instruccion_ternarioContext) GetCond() IExpressionsContext { return s.cond }

func (s *Instruccion_ternarioContext) GetLeft_instr() IExpressionsContext { return s.left_instr }

func (s *Instruccion_ternarioContext) GetRight_instr() IExpressionsContext { return s.right_instr }

func (s *Instruccion_ternarioContext) Get_instr_else_if_ternario() IInstr_else_if_ternarioContext {
	return s._instr_else_if_ternario
}

func (s *Instruccion_ternarioContext) SetCond(v IExpressionsContext) { s.cond = v }

func (s *Instruccion_ternarioContext) SetLeft_instr(v IExpressionsContext) { s.left_instr = v }

func (s *Instruccion_ternarioContext) SetRight_instr(v IExpressionsContext) { s.right_instr = v }

func (s *Instruccion_ternarioContext) Set_instr_else_if_ternario(v IInstr_else_if_ternarioContext) {
	s._instr_else_if_ternario = v
}

func (s *Instruccion_ternarioContext) GetP() interfaces.Expression { return s.p }

func (s *Instruccion_ternarioContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instruccion_ternarioContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarIF, 0)
}

func (s *Instruccion_ternarioContext) AllTK_LLAVEA() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammarTK_LLAVEA)
}

func (s *Instruccion_ternarioContext) TK_LLAVEA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, i)
}

func (s *Instruccion_ternarioContext) AllTK_LLAVEC() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammarTK_LLAVEC)
}

func (s *Instruccion_ternarioContext) TK_LLAVEC(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, i)
}

func (s *Instruccion_ternarioContext) AllExpressions() []IExpressionsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionsContext); ok {
			len++
		}
	}

	tst := make([]IExpressionsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionsContext); ok {
			tst[i] = t.(IExpressionsContext)
			i++
		}
	}

	return tst
}

func (s *Instruccion_ternarioContext) Expressions(i int) IExpressionsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_ternarioContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarELSE, 0)
}

func (s *Instruccion_ternarioContext) Instr_else_if_ternario() IInstr_else_if_ternarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_else_if_ternarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_else_if_ternarioContext)
}

func (s *Instruccion_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_ternario(s)
	}
}

func (s *Instruccion_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_ternario(s)
	}
}

func (p *Swiftgrammar) Instruccion_ternario() (localctx IInstruccion_ternarioContext) {
	localctx = NewInstruccion_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftgrammarRULE_instruccion_ternario)
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(269)

			var _m = p.Match(SwiftgrammarIF)

			localctx.(*Instruccion_ternarioContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).cond = _x
		}
		{
			p.SetState(271)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).left_instr = _x
		}
		{
			p.SetState(273)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_ternarioContext).p = ternario.NewIf(localctx.(*Instruccion_ternarioContext).GetCond().GetP(), localctx.(*Instruccion_ternarioContext).GetLeft_instr().GetP(), nil, nil, (func() int {
			if localctx.(*Instruccion_ternarioContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_ternarioContext).Get_IF().GetLine()
			}
		}()), localctx.(*Instruccion_ternarioContext).Get_IF().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)

			var _m = p.Match(SwiftgrammarIF)

			localctx.(*Instruccion_ternarioContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).cond = _x
		}
		{
			p.SetState(278)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).left_instr = _x
		}
		{
			p.SetState(280)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(SwiftgrammarELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).right_instr = _x
		}
		{
			p.SetState(284)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_ternarioContext).p = ternario.NewIf(localctx.(*Instruccion_ternarioContext).GetCond().GetP(), localctx.(*Instruccion_ternarioContext).GetLeft_instr().GetP(), localctx.(*Instruccion_ternarioContext).GetRight_instr().GetP(), nil, (func() int {
			if localctx.(*Instruccion_ternarioContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_ternarioContext).Get_IF().GetLine()
			}
		}()), localctx.(*Instruccion_ternarioContext).Get_IF().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(287)

			var _m = p.Match(SwiftgrammarIF)

			localctx.(*Instruccion_ternarioContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).cond = _x
		}
		{
			p.SetState(289)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).left_instr = _x
		}
		{
			p.SetState(291)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Match(SwiftgrammarELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)

			var _x = p.Instr_else_if_ternario()

			localctx.(*Instruccion_ternarioContext)._instr_else_if_ternario = _x
		}
		localctx.(*Instruccion_ternarioContext).p = ternario.NewIf(localctx.(*Instruccion_ternarioContext).GetCond().GetP(), localctx.(*Instruccion_ternarioContext).GetLeft_instr().GetP(), nil, localctx.(*Instruccion_ternarioContext).Get_instr_else_if_ternario().GetL(), (func() int {
			if localctx.(*Instruccion_ternarioContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_ternarioContext).Get_IF().GetLine()
			}
		}()), localctx.(*Instruccion_ternarioContext).Get_IF().GetColumn())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_else_if_ternarioContext is an interface to support dynamic dispatch.
type IInstr_else_if_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion_ternario returns the _instruccion_ternario rule contexts.
	Get_instruccion_ternario() IInstruccion_ternarioContext

	// Set_instruccion_ternario sets the _instruccion_ternario rule contexts.
	Set_instruccion_ternario(IInstruccion_ternarioContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccion_ternarioContext

	// SetE sets the e rule context list.
	SetE([]IInstruccion_ternarioContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// Getter signatures
	AllInstruccion_ternario() []IInstruccion_ternarioContext
	Instruccion_ternario(i int) IInstruccion_ternarioContext

	// IsInstr_else_if_ternarioContext differentiates from other interfaces.
	IsInstr_else_if_ternarioContext()
}

type Instr_else_if_ternarioContext struct {
	antlr.BaseParserRuleContext
	parser                antlr.Parser
	l                     *arrayList.List
	_instruccion_ternario IInstruccion_ternarioContext
	e                     []IInstruccion_ternarioContext
}

func NewEmptyInstr_else_if_ternarioContext() *Instr_else_if_ternarioContext {
	var p = new(Instr_else_if_ternarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_else_if_ternario
	return p
}

func InitEmptyInstr_else_if_ternarioContext(p *Instr_else_if_ternarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_else_if_ternario
}

func (*Instr_else_if_ternarioContext) IsInstr_else_if_ternarioContext() {}

func NewInstr_else_if_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_else_if_ternarioContext {
	var p = new(Instr_else_if_ternarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_else_if_ternario

	return p
}

func (s *Instr_else_if_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_else_if_ternarioContext) Get_instruccion_ternario() IInstruccion_ternarioContext {
	return s._instruccion_ternario
}

func (s *Instr_else_if_ternarioContext) Set_instruccion_ternario(v IInstruccion_ternarioContext) {
	s._instruccion_ternario = v
}

func (s *Instr_else_if_ternarioContext) GetE() []IInstruccion_ternarioContext { return s.e }

func (s *Instr_else_if_ternarioContext) SetE(v []IInstruccion_ternarioContext) { s.e = v }

func (s *Instr_else_if_ternarioContext) GetL() *arrayList.List { return s.l }

func (s *Instr_else_if_ternarioContext) SetL(v *arrayList.List) { s.l = v }

func (s *Instr_else_if_ternarioContext) AllInstruccion_ternario() []IInstruccion_ternarioContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccion_ternarioContext); ok {
			len++
		}
	}

	tst := make([]IInstruccion_ternarioContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccion_ternarioContext); ok {
			tst[i] = t.(IInstruccion_ternarioContext)
			i++
		}
	}

	return tst
}

func (s *Instr_else_if_ternarioContext) Instruccion_ternario(i int) IInstruccion_ternarioContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_ternarioContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_ternarioContext)
}

func (s *Instr_else_if_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_else_if_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_else_if_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_else_if_ternario(s)
	}
}

func (s *Instr_else_if_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_else_if_ternario(s)
	}
}

func (p *Swiftgrammar) Instr_else_if_ternario() (localctx IInstr_else_if_ternarioContext) {
	localctx = NewInstr_else_if_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftgrammarRULE_instr_else_if_ternario)

	localctx.(*Instr_else_if_ternarioContext).l = arrayList.New()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		//goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		//goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(298)

				var _x = p.Instruccion_ternario()

				localctx.(*Instr_else_if_ternarioContext)._instruccion_ternario = _x
			}
			localctx.(*Instr_else_if_ternarioContext).e = append(localctx.(*Instr_else_if_ternarioContext).e, localctx.(*Instr_else_if_ternarioContext)._instruccion_ternario)

		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
		//	goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			//goto errorExit
		}
	}

	listInt := localctx.(*Instr_else_if_ternarioContext).GetE()
	for _, e := range listInt {
		localctx.(*Instr_else_if_ternarioContext).l.Add(e.GetP())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_switchContext is an interface to support dynamic dispatch.
type IInstruccion_switchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expressions() IExpressionsContext
	TK_LLAVEA() antlr.TerminalNode
	List_case() IList_caseContext
	Block_default() IBlock_defaultContext
	TK_LLAVEC() antlr.TerminalNode

	// IsInstruccion_switchContext differentiates from other interfaces.
	IsInstruccion_switchContext()
}

type Instruccion_switchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccion_switchContext() *Instruccion_switchContext {
	var p = new(Instruccion_switchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_switch
	return p
}

func InitEmptyInstruccion_switchContext(p *Instruccion_switchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_switch
}

func (*Instruccion_switchContext) IsInstruccion_switchContext() {}

func NewInstruccion_switchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_switchContext {
	var p = new(Instruccion_switchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_switch

	return p
}

func (s *Instruccion_switchContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_switchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarSWITCH, 0)
}

func (s *Instruccion_switchContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_switchContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instruccion_switchContext) List_case() IList_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_caseContext)
}

func (s *Instruccion_switchContext) Block_default() IBlock_defaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_defaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_defaultContext)
}

func (s *Instruccion_switchContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instruccion_switchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_switchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_switchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_switch(s)
	}
}

func (s *Instruccion_switchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_switch(s)
	}
}

func (p *Swiftgrammar) Instruccion_switch() (localctx IInstruccion_switchContext) {
	localctx = NewInstruccion_switchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftgrammarRULE_instruccion_switch)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(306)
			p.Match(SwiftgrammarSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Expressions()
		}
		{
			p.SetState(308)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.List_case()
		}
		{
			p.SetState(310)
			p.Block_default()
		}
		{
			p.SetState(311)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.Match(SwiftgrammarSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Expressions()
		}
		{
			p.SetState(315)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			p.Block_default()
		}
		{
			p.SetState(317)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_caseContext is an interface to support dynamic dispatch.
type IList_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion_case returns the _instruccion_case rule contexts.
	Get_instruccion_case() IInstruccion_caseContext

	// Set_instruccion_case sets the _instruccion_case rule contexts.
	Set_instruccion_case(IInstruccion_caseContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccion_caseContext

	// SetE sets the e rule context list.
	SetE([]IInstruccion_caseContext)

	// Getter signatures
	AllInstruccion_case() []IInstruccion_caseContext
	Instruccion_case(i int) IInstruccion_caseContext

	// IsList_caseContext differentiates from other interfaces.
	IsList_caseContext()
}

type List_caseContext struct {
	antlr.BaseParserRuleContext
	parser            antlr.Parser
	_instruccion_case IInstruccion_caseContext
	e                 []IInstruccion_caseContext
}

func NewEmptyList_caseContext() *List_caseContext {
	var p = new(List_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_case
	return p
}

func InitEmptyList_caseContext(p *List_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_case
}

func (*List_caseContext) IsList_caseContext() {}

func NewList_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_caseContext {
	var p = new(List_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_case

	return p
}

func (s *List_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *List_caseContext) Get_instruccion_case() IInstruccion_caseContext {
	return s._instruccion_case
}

func (s *List_caseContext) Set_instruccion_case(v IInstruccion_caseContext) { s._instruccion_case = v }

func (s *List_caseContext) GetE() []IInstruccion_caseContext { return s.e }

func (s *List_caseContext) SetE(v []IInstruccion_caseContext) { s.e = v }

func (s *List_caseContext) AllInstruccion_case() []IInstruccion_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstruccion_caseContext); ok {
			len++
		}
	}

	tst := make([]IInstruccion_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstruccion_caseContext); ok {
			tst[i] = t.(IInstruccion_caseContext)
			i++
		}
	}

	return tst
}

func (s *List_caseContext) Instruccion_case(i int) IInstruccion_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_caseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_caseContext)
}

func (s *List_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_case(s)
	}
}

func (s *List_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_case(s)
	}
}

func (p *Swiftgrammar) List_case() (localctx IList_caseContext) {
	localctx = NewList_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftgrammarRULE_list_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520464) != 0) {
		{
			p.SetState(321)

			var _x = p.Instruccion_case()

			localctx.(*List_caseContext)._instruccion_case = _x
		}
		localctx.(*List_caseContext).e = append(localctx.(*List_caseContext).e, localctx.(*List_caseContext)._instruccion_case)

		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_caseContext is an interface to support dynamic dispatch.
type IInstruccion_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	List_expre_case() IList_expre_caseContext
	TK_DOSPUNTOS() antlr.TerminalNode
	TK_LLAVEA() antlr.TerminalNode
	Instrucciones() IInstruccionesContext
	TK_LLAVEC() antlr.TerminalNode
	Block_instr_switch() IBlock_instr_switchContext
	TK_COMA() antlr.TerminalNode

	// IsInstruccion_caseContext differentiates from other interfaces.
	IsInstruccion_caseContext()
}

type Instruccion_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccion_caseContext() *Instruccion_caseContext {
	var p = new(Instruccion_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_case
	return p
}

func InitEmptyInstruccion_caseContext(p *Instruccion_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_case
}

func (*Instruccion_caseContext) IsInstruccion_caseContext() {}

func NewInstruccion_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_caseContext {
	var p = new(Instruccion_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_case

	return p
}

func (s *Instruccion_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_caseContext) List_expre_case() IList_expre_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_expre_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_expre_caseContext)
}

func (s *Instruccion_caseContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DOSPUNTOS, 0)
}

func (s *Instruccion_caseContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instruccion_caseContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instruccion_caseContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instruccion_caseContext) Block_instr_switch() IBlock_instr_switchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_instr_switchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_instr_switchContext)
}

func (s *Instruccion_caseContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_COMA, 0)
}

func (s *Instruccion_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_case(s)
	}
}

func (s *Instruccion_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_case(s)
	}
}

func (p *Swiftgrammar) Instruccion_case() (localctx IInstruccion_caseContext) {
	localctx = NewInstruccion_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftgrammarRULE_instruccion_case)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.List_expre_case()
		}
		{
			p.SetState(327)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Instrucciones()
		}
		{
			p.SetState(330)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.List_expre_case()
		}
		{
			p.SetState(333)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Block_instr_switch()
		}
		{
			p.SetState(335)
			p.Match(SwiftgrammarTK_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_expre_caseContext is an interface to support dynamic dispatch.
type IList_expre_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_case returns the _block_case rule contexts.
	Get_block_case() IBlock_caseContext

	// Set_block_case sets the _block_case rule contexts.
	Set_block_case(IBlock_caseContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_caseContext

	// SetE sets the e rule context list.
	SetE([]IBlock_caseContext)

	// Getter signatures
	AllBlock_case() []IBlock_caseContext
	Block_case(i int) IBlock_caseContext

	// IsList_expre_caseContext differentiates from other interfaces.
	IsList_expre_caseContext()
}

type List_expre_caseContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	_block_case IBlock_caseContext
	e           []IBlock_caseContext
}

func NewEmptyList_expre_caseContext() *List_expre_caseContext {
	var p = new(List_expre_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_expre_case
	return p
}

func InitEmptyList_expre_caseContext(p *List_expre_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_expre_case
}

func (*List_expre_caseContext) IsList_expre_caseContext() {}

func NewList_expre_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expre_caseContext {
	var p = new(List_expre_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_expre_case

	return p
}

func (s *List_expre_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expre_caseContext) Get_block_case() IBlock_caseContext { return s._block_case }

func (s *List_expre_caseContext) Set_block_case(v IBlock_caseContext) { s._block_case = v }

func (s *List_expre_caseContext) GetE() []IBlock_caseContext { return s.e }

func (s *List_expre_caseContext) SetE(v []IBlock_caseContext) { s.e = v }

func (s *List_expre_caseContext) AllBlock_case() []IBlock_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_caseContext); ok {
			len++
		}
	}

	tst := make([]IBlock_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_caseContext); ok {
			tst[i] = t.(IBlock_caseContext)
			i++
		}
	}

	return tst
}

func (s *List_expre_caseContext) Block_case(i int) IBlock_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_caseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_caseContext)
}

func (s *List_expre_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expre_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expre_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_expre_case(s)
	}
}

func (s *List_expre_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_expre_case(s)
	}
}

func (p *Swiftgrammar) List_expre_case() (localctx IList_expre_caseContext) {
	localctx = NewList_expre_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftgrammarRULE_list_expre_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520464) != 0) {
		{
			p.SetState(339)

			var _x = p.Block_case()

			localctx.(*List_expre_caseContext)._block_case = _x
		}
		localctx.(*List_expre_caseContext).e = append(localctx.(*List_expre_caseContext).e, localctx.(*List_expre_caseContext)._block_case)

		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_caseContext is an interface to support dynamic dispatch.
type IBlock_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expressions() IExpressionsContext
	TK_BARRA() antlr.TerminalNode

	// IsBlock_caseContext differentiates from other interfaces.
	IsBlock_caseContext()
}

type Block_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_caseContext() *Block_caseContext {
	var p = new(Block_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_case
	return p
}

func InitEmptyBlock_caseContext(p *Block_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_case
}

func (*Block_caseContext) IsBlock_caseContext() {}

func NewBlock_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_caseContext {
	var p = new(Block_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_case

	return p
}

func (s *Block_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_caseContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_caseContext) TK_BARRA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_BARRA, 0)
}

func (s *Block_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_case(s)
	}
}

func (s *Block_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_case(s)
	}
}

func (p *Swiftgrammar) Block_case() (localctx IBlock_caseContext) {
	localctx = NewBlock_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftgrammarRULE_block_case)
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.Expressions()
		}
		{
			p.SetState(345)
			p.Match(SwiftgrammarTK_BARRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)
			p.Expressions()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_instr_switchContext is an interface to support dynamic dispatch.
type IBlock_instr_switchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetList returns the list rule context list.
	GetList() []IInstruccionContext

	// SetList sets the list rule context list.
	SetList([]IInstruccionContext)

	// Getter signatures
	Instruccion() IInstruccionContext

	// IsBlock_instr_switchContext differentiates from other interfaces.
	IsBlock_instr_switchContext()
}

type Block_instr_switchContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	_instruccion IInstruccionContext
	list         []IInstruccionContext
}

func NewEmptyBlock_instr_switchContext() *Block_instr_switchContext {
	var p = new(Block_instr_switchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_instr_switch
	return p
}

func InitEmptyBlock_instr_switchContext(p *Block_instr_switchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_instr_switch
}

func (*Block_instr_switchContext) IsBlock_instr_switchContext() {}

func NewBlock_instr_switchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_instr_switchContext {
	var p = new(Block_instr_switchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_instr_switch

	return p
}

func (s *Block_instr_switchContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_instr_switchContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *Block_instr_switchContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *Block_instr_switchContext) GetList() []IInstruccionContext { return s.list }

func (s *Block_instr_switchContext) SetList(v []IInstruccionContext) { s.list = v }

func (s *Block_instr_switchContext) Instruccion() IInstruccionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *Block_instr_switchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_instr_switchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_instr_switchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_instr_switch(s)
	}
}

func (s *Block_instr_switchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_instr_switch(s)
	}
}

func (p *Swiftgrammar) Block_instr_switch() (localctx IBlock_instr_switchContext) {
	localctx = NewBlock_instr_switchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftgrammarRULE_block_instr_switch)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)

		var _x = p.Instruccion()

		localctx.(*Block_instr_switchContext)._instruccion = _x
	}
	localctx.(*Block_instr_switchContext).list = append(localctx.(*Block_instr_switchContext).list, localctx.(*Block_instr_switchContext)._instruccion)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_defaultContext is an interface to support dynamic dispatch.
type IInstr_defaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	ID() antlr.TerminalNode
	TK_DOSPUNTOS() antlr.TerminalNode
	TK_LLAVEA() antlr.TerminalNode
	Instrucciones() IInstruccionesContext
	TK_LLAVEC() antlr.TerminalNode
	Block_instr_switch() IBlock_instr_switchContext

	// IsInstr_defaultContext differentiates from other interfaces.
	IsInstr_defaultContext()
}

type Instr_defaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr_defaultContext() *Instr_defaultContext {
	var p = new(Instr_defaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_default
	return p
}

func InitEmptyInstr_defaultContext(p *Instr_defaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_default
}

func (*Instr_defaultContext) IsInstr_defaultContext() {}

func NewInstr_defaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_defaultContext {
	var p = new(Instr_defaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_default

	return p
}

func (s *Instr_defaultContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_defaultContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarCASE, 0)
}

func (s *Instr_defaultContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instr_defaultContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DOSPUNTOS, 0)
}

func (s *Instr_defaultContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instr_defaultContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instr_defaultContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instr_defaultContext) Block_instr_switch() IBlock_instr_switchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_instr_switchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_instr_switchContext)
}

func (s *Instr_defaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_defaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_defaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_default(s)
	}
}

func (s *Instr_defaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_default(s)
	}
}

func (p *Swiftgrammar) Instr_default() (localctx IInstr_defaultContext) {
	localctx = NewInstr_defaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftgrammarRULE_instr_default)
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(352)
			p.Match(SwiftgrammarCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Instrucciones()
		}
		{
			p.SetState(357)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
			p.Match(SwiftgrammarCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)
			p.Block_instr_switch()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_defaultContext is an interface to support dynamic dispatch.
type IBlock_defaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_default returns the _instr_default rule contexts.
	Get_instr_default() IInstr_defaultContext

	// Set_instr_default sets the _instr_default rule contexts.
	Set_instr_default(IInstr_defaultContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_defaultContext

	// SetE sets the e rule context list.
	SetE([]IInstr_defaultContext)

	// Getter signatures
	AllInstr_default() []IInstr_defaultContext
	Instr_default(i int) IInstr_defaultContext

	// IsBlock_defaultContext differentiates from other interfaces.
	IsBlock_defaultContext()
}

type Block_defaultContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	_instr_default IInstr_defaultContext
	e              []IInstr_defaultContext
}

func NewEmptyBlock_defaultContext() *Block_defaultContext {
	var p = new(Block_defaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_default
	return p
}

func InitEmptyBlock_defaultContext(p *Block_defaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_default
}

func (*Block_defaultContext) IsBlock_defaultContext() {}

func NewBlock_defaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_defaultContext {
	var p = new(Block_defaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_default

	return p
}

func (s *Block_defaultContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_defaultContext) Get_instr_default() IInstr_defaultContext { return s._instr_default }

func (s *Block_defaultContext) Set_instr_default(v IInstr_defaultContext) { s._instr_default = v }

func (s *Block_defaultContext) GetE() []IInstr_defaultContext { return s.e }

func (s *Block_defaultContext) SetE(v []IInstr_defaultContext) { s.e = v }

func (s *Block_defaultContext) AllInstr_default() []IInstr_defaultContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstr_defaultContext); ok {
			len++
		}
	}

	tst := make([]IInstr_defaultContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstr_defaultContext); ok {
			tst[i] = t.(IInstr_defaultContext)
			i++
		}
	}

	return tst
}

func (s *Block_defaultContext) Instr_default(i int) IInstr_defaultContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_defaultContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_defaultContext)
}

func (s *Block_defaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_defaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_defaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_default(s)
	}
}

func (s *Block_defaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_default(s)
	}
}

func (p *Swiftgrammar) Block_default() (localctx IBlock_defaultContext) {
	localctx = NewBlock_defaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftgrammarRULE_block_default)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftgrammarCASE {
		{
			p.SetState(365)

			var _x = p.Instr_default()

			localctx.(*Block_defaultContext)._instr_default = _x
		}
		localctx.(*Block_defaultContext).e = append(localctx.(*Block_defaultContext).e, localctx.(*Block_defaultContext)._instr_default)

		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_switch_ternarioContext is an interface to support dynamic dispatch.
type IInstruccion_switch_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expressions() IExpressionsContext
	TK_LLAVEA() antlr.TerminalNode
	List_case_ternario() IList_case_ternarioContext
	Instr_default_ter() IInstr_default_terContext
	TK_LLAVEC() antlr.TerminalNode

	// IsInstruccion_switch_ternarioContext differentiates from other interfaces.
	IsInstruccion_switch_ternarioContext()
}

type Instruccion_switch_ternarioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccion_switch_ternarioContext() *Instruccion_switch_ternarioContext {
	var p = new(Instruccion_switch_ternarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_switch_ternario
	return p
}

func InitEmptyInstruccion_switch_ternarioContext(p *Instruccion_switch_ternarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_switch_ternario
}

func (*Instruccion_switch_ternarioContext) IsInstruccion_switch_ternarioContext() {}

func NewInstruccion_switch_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_switch_ternarioContext {
	var p = new(Instruccion_switch_ternarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_switch_ternario

	return p
}

func (s *Instruccion_switch_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_switch_ternarioContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarSWITCH, 0)
}

func (s *Instruccion_switch_ternarioContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_switch_ternarioContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instruccion_switch_ternarioContext) List_case_ternario() IList_case_ternarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_case_ternarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_case_ternarioContext)
}

func (s *Instruccion_switch_ternarioContext) Instr_default_ter() IInstr_default_terContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_default_terContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_default_terContext)
}

func (s *Instruccion_switch_ternarioContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instruccion_switch_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_switch_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_switch_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_switch_ternario(s)
	}
}

func (s *Instruccion_switch_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_switch_ternario(s)
	}
}

func (p *Swiftgrammar) Instruccion_switch_ternario() (localctx IInstruccion_switch_ternarioContext) {
	localctx = NewInstruccion_switch_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftgrammarRULE_instruccion_switch_ternario)
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(370)
			p.Match(SwiftgrammarSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)
			p.Expressions()
		}
		{
			p.SetState(372)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.List_case_ternario()
		}
		{
			p.SetState(374)
			p.Instr_default_ter()
		}
		{
			p.SetState(375)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(377)
			p.Match(SwiftgrammarSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.Expressions()
		}
		{
			p.SetState(379)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Instr_default_ter()
		}
		{
			p.SetState(381)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_case_ternarioContext is an interface to support dynamic dispatch.
type IList_case_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instr_case_ter returns the _instr_case_ter rule contexts.
	Get_instr_case_ter() IInstr_case_terContext

	// Set_instr_case_ter sets the _instr_case_ter rule contexts.
	Set_instr_case_ter(IInstr_case_terContext)

	// GetE returns the e rule context list.
	GetE() []IInstr_case_terContext

	// SetE sets the e rule context list.
	SetE([]IInstr_case_terContext)

	// Getter signatures
	AllInstr_case_ter() []IInstr_case_terContext
	Instr_case_ter(i int) IInstr_case_terContext

	// IsList_case_ternarioContext differentiates from other interfaces.
	IsList_case_ternarioContext()
}

type List_case_ternarioContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	_instr_case_ter IInstr_case_terContext
	e               []IInstr_case_terContext
}

func NewEmptyList_case_ternarioContext() *List_case_ternarioContext {
	var p = new(List_case_ternarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_case_ternario
	return p
}

func InitEmptyList_case_ternarioContext(p *List_case_ternarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_case_ternario
}

func (*List_case_ternarioContext) IsList_case_ternarioContext() {}

func NewList_case_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_case_ternarioContext {
	var p = new(List_case_ternarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_case_ternario

	return p
}

func (s *List_case_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *List_case_ternarioContext) Get_instr_case_ter() IInstr_case_terContext {
	return s._instr_case_ter
}

func (s *List_case_ternarioContext) Set_instr_case_ter(v IInstr_case_terContext) {
	s._instr_case_ter = v
}

func (s *List_case_ternarioContext) GetE() []IInstr_case_terContext { return s.e }

func (s *List_case_ternarioContext) SetE(v []IInstr_case_terContext) { s.e = v }

func (s *List_case_ternarioContext) AllInstr_case_ter() []IInstr_case_terContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstr_case_terContext); ok {
			len++
		}
	}

	tst := make([]IInstr_case_terContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstr_case_terContext); ok {
			tst[i] = t.(IInstr_case_terContext)
			i++
		}
	}

	return tst
}

func (s *List_case_ternarioContext) Instr_case_ter(i int) IInstr_case_terContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_case_terContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_case_terContext)
}

func (s *List_case_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_case_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_case_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_case_ternario(s)
	}
}

func (s *List_case_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_case_ternario(s)
	}
}

func (p *Swiftgrammar) List_case_ternario() (localctx IList_case_ternarioContext) {
	localctx = NewList_case_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftgrammarRULE_list_case_ternario)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520464) != 0) {
		{
			p.SetState(385)

			var _x = p.Instr_case_ter()

			localctx.(*List_case_ternarioContext)._instr_case_ter = _x
		}
		localctx.(*List_case_ternarioContext).e = append(localctx.(*List_case_ternarioContext).e, localctx.(*List_case_ternarioContext)._instr_case_ter)

		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_case_terContext is an interface to support dynamic dispatch.
type IInstr_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	List_expre_case_ter() IList_expre_case_terContext
	TK_DOSPUNTOS() antlr.TerminalNode
	Expressions() IExpressionsContext
	TK_LLAVEA() antlr.TerminalNode
	TK_LLAVEC() antlr.TerminalNode

	// IsInstr_case_terContext differentiates from other interfaces.
	IsInstr_case_terContext()
}

type Instr_case_terContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr_case_terContext() *Instr_case_terContext {
	var p = new(Instr_case_terContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_case_ter
	return p
}

func InitEmptyInstr_case_terContext(p *Instr_case_terContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_case_ter
}

func (*Instr_case_terContext) IsInstr_case_terContext() {}

func NewInstr_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_case_terContext {
	var p = new(Instr_case_terContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_case_ter

	return p
}

func (s *Instr_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_case_terContext) List_expre_case_ter() IList_expre_case_terContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_expre_case_terContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_expre_case_terContext)
}

func (s *Instr_case_terContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DOSPUNTOS, 0)
}

func (s *Instr_case_terContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_case_terContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instr_case_terContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instr_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_case_ter(s)
	}
}

func (s *Instr_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_case_ter(s)
	}
}

func (p *Swiftgrammar) Instr_case_ter() (localctx IInstr_case_terContext) {
	localctx = NewInstr_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftgrammarRULE_instr_case_ter)
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(390)
			p.List_expre_case_ter()
		}
		{
			p.SetState(391)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Expressions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(394)
			p.List_expre_case_ter()
		}
		{
			p.SetState(395)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Expressions()
		}
		{
			p.SetState(398)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_expre_case_terContext is an interface to support dynamic dispatch.
type IList_expre_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_case_ter returns the _block_case_ter rule contexts.
	Get_block_case_ter() IBlock_case_terContext

	// Set_block_case_ter sets the _block_case_ter rule contexts.
	Set_block_case_ter(IBlock_case_terContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_case_terContext

	// SetE sets the e rule context list.
	SetE([]IBlock_case_terContext)

	// Getter signatures
	AllBlock_case_ter() []IBlock_case_terContext
	Block_case_ter(i int) IBlock_case_terContext

	// IsList_expre_case_terContext differentiates from other interfaces.
	IsList_expre_case_terContext()
}

type List_expre_case_terContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	_block_case_ter IBlock_case_terContext
	e               []IBlock_case_terContext
}

func NewEmptyList_expre_case_terContext() *List_expre_case_terContext {
	var p = new(List_expre_case_terContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_expre_case_ter
	return p
}

func InitEmptyList_expre_case_terContext(p *List_expre_case_terContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_expre_case_ter
}

func (*List_expre_case_terContext) IsList_expre_case_terContext() {}

func NewList_expre_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expre_case_terContext {
	var p = new(List_expre_case_terContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_expre_case_ter

	return p
}

func (s *List_expre_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expre_case_terContext) Get_block_case_ter() IBlock_case_terContext {
	return s._block_case_ter
}

func (s *List_expre_case_terContext) Set_block_case_ter(v IBlock_case_terContext) {
	s._block_case_ter = v
}

func (s *List_expre_case_terContext) GetE() []IBlock_case_terContext { return s.e }

func (s *List_expre_case_terContext) SetE(v []IBlock_case_terContext) { s.e = v }

func (s *List_expre_case_terContext) AllBlock_case_ter() []IBlock_case_terContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_case_terContext); ok {
			len++
		}
	}

	tst := make([]IBlock_case_terContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_case_terContext); ok {
			tst[i] = t.(IBlock_case_terContext)
			i++
		}
	}

	return tst
}

func (s *List_expre_case_terContext) Block_case_ter(i int) IBlock_case_terContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_case_terContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_case_terContext)
}

func (s *List_expre_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expre_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expre_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_expre_case_ter(s)
	}
}

func (s *List_expre_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_expre_case_ter(s)
	}
}

func (p *Swiftgrammar) List_expre_case_ter() (localctx IList_expre_case_terContext) {
	localctx = NewList_expre_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftgrammarRULE_list_expre_case_ter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520464) != 0) {
		{
			p.SetState(402)

			var _x = p.Block_case_ter()

			localctx.(*List_expre_case_terContext)._block_case_ter = _x
		}
		localctx.(*List_expre_case_terContext).e = append(localctx.(*List_expre_case_terContext).e, localctx.(*List_expre_case_terContext)._block_case_ter)

		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_case_terContext is an interface to support dynamic dispatch.
type IBlock_case_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expressions() IExpressionsContext
	TK_BARRA() antlr.TerminalNode

	// IsBlock_case_terContext differentiates from other interfaces.
	IsBlock_case_terContext()
}

type Block_case_terContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_case_terContext() *Block_case_terContext {
	var p = new(Block_case_terContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_case_ter
	return p
}

func InitEmptyBlock_case_terContext(p *Block_case_terContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_case_ter
}

func (*Block_case_terContext) IsBlock_case_terContext() {}

func NewBlock_case_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_case_terContext {
	var p = new(Block_case_terContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_case_ter

	return p
}

func (s *Block_case_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_case_terContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_case_terContext) TK_BARRA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_BARRA, 0)
}

func (s *Block_case_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_case_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_case_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_case_ter(s)
	}
}

func (s *Block_case_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_case_ter(s)
	}
}

func (p *Swiftgrammar) Block_case_ter() (localctx IBlock_case_terContext) {
	localctx = NewBlock_case_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftgrammarRULE_block_case_ter)
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(407)
			p.Expressions()
		}
		{
			p.SetState(408)
			p.Match(SwiftgrammarTK_BARRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.Expressions()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_default_terContext is an interface to support dynamic dispatch.
type IInstr_default_terContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	ID() antlr.TerminalNode
	TK_DOSPUNTOS() antlr.TerminalNode
	Expressions() IExpressionsContext
	TK_LLAVEA() antlr.TerminalNode
	TK_LLAVEC() antlr.TerminalNode

	// IsInstr_default_terContext differentiates from other interfaces.
	IsInstr_default_terContext()
}

type Instr_default_terContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr_default_terContext() *Instr_default_terContext {
	var p = new(Instr_default_terContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_default_ter
	return p
}

func InitEmptyInstr_default_terContext(p *Instr_default_terContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_default_ter
}

func (*Instr_default_terContext) IsInstr_default_terContext() {}

func NewInstr_default_terContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_default_terContext {
	var p = new(Instr_default_terContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_default_ter

	return p
}

func (s *Instr_default_terContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_default_terContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarCASE, 0)
}

func (s *Instr_default_terContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instr_default_terContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DOSPUNTOS, 0)
}

func (s *Instr_default_terContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_default_terContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instr_default_terContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instr_default_terContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_default_terContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_default_terContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_default_ter(s)
	}
}

func (s *Instr_default_terContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_default_ter(s)
	}
}

func (p *Swiftgrammar) Instr_default_ter() (localctx IInstr_default_terContext) {
	localctx = NewInstr_default_terContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SwiftgrammarRULE_instr_default_ter)
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(413)
			p.Match(SwiftgrammarCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Expressions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(417)
			p.Match(SwiftgrammarCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Expressions()
		}
		{
			p.SetState(422)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_whileContext is an interface to support dynamic dispatch.
type IInstruccion_whileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expressions() IExpressionsContext
	TK_LLAVEA() antlr.TerminalNode
	Instrucciones() IInstruccionesContext
	TK_LLAVEC() antlr.TerminalNode

	// IsInstruccion_whileContext differentiates from other interfaces.
	IsInstruccion_whileContext()
}

type Instruccion_whileContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_WHILE         antlr.Token
	_expressions   IExpressionsContext
	_instrucciones IInstruccionesContext
}

func NewEmptyInstruccion_whileContext() *Instruccion_whileContext {
	var p = new(Instruccion_whileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_while
	return p
}

func InitEmptyInstruccion_whileContext(p *Instruccion_whileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_while
}

func (*Instruccion_whileContext) IsInstruccion_whileContext() {}

func NewInstruccion_whileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_whileContext {
	var p = new(Instruccion_whileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_while

	return p
}

func (s *Instruccion_whileContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_whileContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *Instruccion_whileContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *Instruccion_whileContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instruccion_whileContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instruccion_whileContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instruccion_whileContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instruccion_whileContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_whileContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_whileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarWHILE, 0)
}

func (s *Instruccion_whileContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_whileContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instruccion_whileContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instruccion_whileContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instruccion_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_whileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_whileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_while(s)
	}
}

func (s *Instruccion_whileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_while(s)
	}
}

func (p *Swiftgrammar) Instruccion_while() (localctx IInstruccion_whileContext) {
	localctx = NewInstruccion_whileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftgrammarRULE_instruccion_while)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)

		var _m = p.Match(SwiftgrammarWHILE)

		localctx.(*Instruccion_whileContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)

		var _x = p.Expressions()

		localctx.(*Instruccion_whileContext)._expressions = _x
	}
	{
		p.SetState(428)
		p.Match(SwiftgrammarTK_LLAVEA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(429)

		var _x = p.Instrucciones()

		localctx.(*Instruccion_whileContext)._instrucciones = _x
	}
	{
		p.SetState(430)
		p.Match(SwiftgrammarTK_LLAVEC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*Instruccion_whileContext).instr = instruction.NewWhile(localctx.(*Instruccion_whileContext).Get_expressions().GetP(), localctx.(*Instruccion_whileContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instruccion_whileContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*Instruccion_whileContext).Get_WHILE().GetLine()
		}
	}()), localctx.(*Instruccion_whileContext).Get_WHILE().GetColumn())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_for_inContext is an interface to support dynamic dispatch.
type IInstruccion_for_inContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FOR returns the _FOR token.
	Get_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionsContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionsContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionsContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	TK_TRIPLEPUNTO() antlr.TerminalNode
	TK_LLAVEA() antlr.TerminalNode
	Instrucciones() IInstruccionesContext
	TK_LLAVEC() antlr.TerminalNode
	AllExpressions() []IExpressionsContext
	Expressions(i int) IExpressionsContext

	// IsInstruccion_for_inContext differentiates from other interfaces.
	IsInstruccion_for_inContext()
}

type Instruccion_for_inContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_FOR           antlr.Token
	_ID            antlr.Token
	left           IExpressionsContext
	right          IExpressionsContext
	_instrucciones IInstruccionesContext
}

func NewEmptyInstruccion_for_inContext() *Instruccion_for_inContext {
	var p = new(Instruccion_for_inContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_for_in
	return p
}

func InitEmptyInstruccion_for_inContext(p *Instruccion_for_inContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_for_in
}

func (*Instruccion_for_inContext) IsInstruccion_for_inContext() {}

func NewInstruccion_for_inContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_for_inContext {
	var p = new(Instruccion_for_inContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_for_in

	return p
}

func (s *Instruccion_for_inContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_for_inContext) Get_FOR() antlr.Token { return s._FOR }

func (s *Instruccion_for_inContext) Get_ID() antlr.Token { return s._ID }

func (s *Instruccion_for_inContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *Instruccion_for_inContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instruccion_for_inContext) GetLeft() IExpressionsContext { return s.left }

func (s *Instruccion_for_inContext) GetRight() IExpressionsContext { return s.right }

func (s *Instruccion_for_inContext) Get_instrucciones() IInstruccionesContext {
	return s._instrucciones
}

func (s *Instruccion_for_inContext) SetLeft(v IExpressionsContext) { s.left = v }

func (s *Instruccion_for_inContext) SetRight(v IExpressionsContext) { s.right = v }

func (s *Instruccion_for_inContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instruccion_for_inContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_for_inContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_for_inContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarFOR, 0)
}

func (s *Instruccion_for_inContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instruccion_for_inContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarIN, 0)
}

func (s *Instruccion_for_inContext) TK_TRIPLEPUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_TRIPLEPUNTO, 0)
}

func (s *Instruccion_for_inContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instruccion_for_inContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instruccion_for_inContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instruccion_for_inContext) AllExpressions() []IExpressionsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionsContext); ok {
			len++
		}
	}

	tst := make([]IExpressionsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionsContext); ok {
			tst[i] = t.(IExpressionsContext)
			i++
		}
	}

	return tst
}

func (s *Instruccion_for_inContext) Expressions(i int) IExpressionsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_for_inContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_for_inContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_for_inContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_for_in(s)
	}
}

func (s *Instruccion_for_inContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_for_in(s)
	}
}

func (p *Swiftgrammar) Instruccion_for_in() (localctx IInstruccion_for_inContext) {
	localctx = NewInstruccion_for_inContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftgrammarRULE_instruccion_for_in)
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)

			var _m = p.Match(SwiftgrammarFOR)

			localctx.(*Instruccion_for_inContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_for_inContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Match(SwiftgrammarIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)

			var _x = p.Expressions()

			localctx.(*Instruccion_for_inContext).left = _x
		}
		{
			p.SetState(437)
			p.Match(SwiftgrammarTK_TRIPLEPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)

			var _x = p.Expressions()

			localctx.(*Instruccion_for_inContext).right = _x
		}
		{
			p.SetState(439)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(441)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_for_inContext).instr = instruction.NewFor((func() string {
			if localctx.(*Instruccion_for_inContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_for_inContext).Get_ID().GetText()
			}
		}()), interfaces.INTEGER, localctx.(*Instruccion_for_inContext).GetLeft().GetP(), localctx.(*Instruccion_for_inContext).GetRight().GetP(), localctx.(*Instruccion_for_inContext).Get_instrucciones().GetL(), (func() int {
			if localctx.(*Instruccion_for_inContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_for_inContext).Get_FOR().GetLine()
			}
		}()), localctx.(*Instruccion_for_inContext).Get_FOR().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(444)

			var _m = p.Match(SwiftgrammarFOR)

			localctx.(*Instruccion_for_inContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_for_inContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.Match(SwiftgrammarIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)

			var _x = p.Expressions()

			localctx.(*Instruccion_for_inContext).left = _x
		}
		{
			p.SetState(448)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_for_inContext)._instrucciones = _x
		}
		{
			p.SetState(450)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_for_inContext).instr = instruction.NewFor((func() string {
			if localctx.(*Instruccion_for_inContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_for_inContext).Get_ID().GetText()
			}
		}()), interfaces.STRING, localctx.(*Instruccion_for_inContext).GetLeft().GetP(), nil, localctx.(*Instruccion_for_inContext).Get_instrucciones().GetL(), (func() int {
			if localctx.(*Instruccion_for_inContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_for_inContext).Get_FOR().GetLine()
			}
		}()), localctx.(*Instruccion_for_inContext).Get_FOR().GetColumn())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_while_trueContext is an interface to support dynamic dispatch.
type IInstruccion_while_trueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	TK_LLAVEA() antlr.TerminalNode
	Instrucciones() IInstruccionesContext
	TK_LLAVEC() antlr.TerminalNode

	// IsInstruccion_while_trueContext differentiates from other interfaces.
	IsInstruccion_while_trueContext()
}

type Instruccion_while_trueContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	instr          interfaces.Instruction
	_WHILE         antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstruccion_while_trueContext() *Instruccion_while_trueContext {
	var p = new(Instruccion_while_trueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_while_true
	return p
}

func InitEmptyInstruccion_while_trueContext(p *Instruccion_while_trueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_while_true
}

func (*Instruccion_while_trueContext) IsInstruccion_while_trueContext() {}

func NewInstruccion_while_trueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_while_trueContext {
	var p = new(Instruccion_while_trueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_while_true

	return p
}

func (s *Instruccion_while_trueContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_while_trueContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *Instruccion_while_trueContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *Instruccion_while_trueContext) Get_instrucciones() IInstruccionesContext {
	return s._instrucciones
}

func (s *Instruccion_while_trueContext) Set_instrucciones(v IInstruccionesContext) {
	s._instrucciones = v
}

func (s *Instruccion_while_trueContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_while_trueContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_while_trueContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarWHILE, 0)
}

func (s *Instruccion_while_trueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTRUE, 0)
}

func (s *Instruccion_while_trueContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instruccion_while_trueContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instruccion_while_trueContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instruccion_while_trueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_while_trueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_while_trueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_while_true(s)
	}
}

func (s *Instruccion_while_trueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_while_true(s)
	}
}

func (p *Swiftgrammar) Instruccion_while_true() (localctx IInstruccion_while_trueContext) {
	localctx = NewInstruccion_while_trueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SwiftgrammarRULE_instruccion_while_true)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)

		var _m = p.Match(SwiftgrammarWHILE)

		localctx.(*Instruccion_while_trueContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.Match(SwiftgrammarTRUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(457)
		p.Match(SwiftgrammarTK_LLAVEA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)

		var _x = p.Instrucciones()

		localctx.(*Instruccion_while_trueContext)._instrucciones = _x
	}
	{
		p.SetState(459)
		p.Match(SwiftgrammarTK_LLAVEC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*Instruccion_while_trueContext).instr = instruction.NewWtrue(localctx.(*Instruccion_while_trueContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instruccion_while_trueContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*Instruccion_while_trueContext).Get_WHILE().GetLine()
		}
	}()), localctx.(*Instruccion_while_trueContext).Get_WHILE().GetColumn())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_while_true_ternarioContext is an interface to support dynamic dispatch.
type IInstruccion_while_true_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	WHILE() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	TK_LLAVEA() antlr.TerminalNode
	Instrucciones() IInstruccionesContext
	TK_LLAVEC() antlr.TerminalNode

	// IsInstruccion_while_true_ternarioContext differentiates from other interfaces.
	IsInstruccion_while_true_ternarioContext()
}

type Instruccion_while_true_ternarioContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	p              interfaces.Expression
	_WHILE         antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyInstruccion_while_true_ternarioContext() *Instruccion_while_true_ternarioContext {
	var p = new(Instruccion_while_true_ternarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_while_true_ternario
	return p
}

func InitEmptyInstruccion_while_true_ternarioContext(p *Instruccion_while_true_ternarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_while_true_ternario
}

func (*Instruccion_while_true_ternarioContext) IsInstruccion_while_true_ternarioContext() {}

func NewInstruccion_while_true_ternarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_while_true_ternarioContext {
	var p = new(Instruccion_while_true_ternarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_while_true_ternario

	return p
}

func (s *Instruccion_while_true_ternarioContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_while_true_ternarioContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *Instruccion_while_true_ternarioContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *Instruccion_while_true_ternarioContext) Get_instrucciones() IInstruccionesContext {
	return s._instrucciones
}

func (s *Instruccion_while_true_ternarioContext) Set_instrucciones(v IInstruccionesContext) {
	s._instrucciones = v
}

func (s *Instruccion_while_true_ternarioContext) GetP() interfaces.Expression { return s.p }

func (s *Instruccion_while_true_ternarioContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Instruccion_while_true_ternarioContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarWHILE, 0)
}

func (s *Instruccion_while_true_ternarioContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTRUE, 0)
}

func (s *Instruccion_while_true_ternarioContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instruccion_while_true_ternarioContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instruccion_while_true_ternarioContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instruccion_while_true_ternarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_while_true_ternarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_while_true_ternarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_while_true_ternario(s)
	}
}

func (s *Instruccion_while_true_ternarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_while_true_ternario(s)
	}
}

func (p *Swiftgrammar) Instruccion_while_true_ternario() (localctx IInstruccion_while_true_ternarioContext) {
	localctx = NewInstruccion_while_true_ternarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SwiftgrammarRULE_instruccion_while_true_ternario)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)

		var _m = p.Match(SwiftgrammarWHILE)

		localctx.(*Instruccion_while_true_ternarioContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(463)
		p.Match(SwiftgrammarTRUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(464)
		p.Match(SwiftgrammarTK_LLAVEA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)

		var _x = p.Instrucciones()

		localctx.(*Instruccion_while_true_ternarioContext)._instrucciones = _x
	}
	{
		p.SetState(466)
		p.Match(SwiftgrammarTK_LLAVEC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*Instruccion_while_true_ternarioContext).p = ternario.NewWhileter(localctx.(*Instruccion_while_true_ternarioContext).Get_instrucciones().GetL(), (func() int {
		if localctx.(*Instruccion_while_true_ternarioContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*Instruccion_while_true_ternarioContext).Get_WHILE().GetLine()
		}
	}()), localctx.(*Instruccion_while_true_ternarioContext).Get_WHILE().GetColumn())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_breakContext is an interface to support dynamic dispatch.
type IInstruccion_breakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode
	Expressions() IExpressionsContext

	// IsInstruccion_breakContext differentiates from other interfaces.
	IsInstruccion_breakContext()
}

type Instruccion_breakContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccion_breakContext() *Instruccion_breakContext {
	var p = new(Instruccion_breakContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_break
	return p
}

func InitEmptyInstruccion_breakContext(p *Instruccion_breakContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_break
}

func (*Instruccion_breakContext) IsInstruccion_breakContext() {}

func NewInstruccion_breakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_breakContext {
	var p = new(Instruccion_breakContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_break

	return p
}

func (s *Instruccion_breakContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_breakContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarBREAK, 0)
}

func (s *Instruccion_breakContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_breakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_breakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_breakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_break(s)
	}
}

func (s *Instruccion_breakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_break(s)
	}
}

func (p *Swiftgrammar) Instruccion_break() (localctx IInstruccion_breakContext) {
	localctx = NewInstruccion_breakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SwiftgrammarRULE_instruccion_break)
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(469)
			p.Match(SwiftgrammarBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(470)
			p.Match(SwiftgrammarBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.Expressions()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_continueContext is an interface to support dynamic dispatch.
type IInstruccion_continueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsInstruccion_continueContext differentiates from other interfaces.
	IsInstruccion_continueContext()
}

type Instruccion_continueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccion_continueContext() *Instruccion_continueContext {
	var p = new(Instruccion_continueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_continue
	return p
}

func InitEmptyInstruccion_continueContext(p *Instruccion_continueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_continue
}

func (*Instruccion_continueContext) IsInstruccion_continueContext() {}

func NewInstruccion_continueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_continueContext {
	var p = new(Instruccion_continueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_continue

	return p
}

func (s *Instruccion_continueContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_continueContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarCONTINUE, 0)
}

func (s *Instruccion_continueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_continueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_continueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_continue(s)
	}
}

func (s *Instruccion_continueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_continue(s)
	}
}

func (p *Swiftgrammar) Instruccion_continue() (localctx IInstruccion_continueContext) {
	localctx = NewInstruccion_continueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftgrammarRULE_instruccion_continue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(474)
		p.Match(SwiftgrammarCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_returnContext is an interface to support dynamic dispatch.
type IInstruccion_returnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expressions() IExpressionsContext

	// IsInstruccion_returnContext differentiates from other interfaces.
	IsInstruccion_returnContext()
}

type Instruccion_returnContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        interfaces.Instruction
	_RETURN      antlr.Token
	_expressions IExpressionsContext
}

func NewEmptyInstruccion_returnContext() *Instruccion_returnContext {
	var p = new(Instruccion_returnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_return
	return p
}

func InitEmptyInstruccion_returnContext(p *Instruccion_returnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_return
}

func (*Instruccion_returnContext) IsInstruccion_returnContext() {}

func NewInstruccion_returnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_returnContext {
	var p = new(Instruccion_returnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_return

	return p
}

func (s *Instruccion_returnContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_returnContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *Instruccion_returnContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *Instruccion_returnContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Instruccion_returnContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Instruccion_returnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_returnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_returnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarRETURN, 0)
}

func (s *Instruccion_returnContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instruccion_returnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_returnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_returnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_return(s)
	}
}

func (s *Instruccion_returnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_return(s)
	}
}

func (p *Swiftgrammar) Instruccion_return() (localctx IInstruccion_returnContext) {
	localctx = NewInstruccion_returnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SwiftgrammarRULE_instruccion_return)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(476)

		var _m = p.Match(SwiftgrammarRETURN)

		localctx.(*Instruccion_returnContext)._RETURN = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(477)

		var _x = p.Expressions()

		localctx.(*Instruccion_returnContext)._expressions = _x
	}
	localctx.(*Instruccion_returnContext).instr = instruction.NewReturn(localctx.(*Instruccion_returnContext).Get_expressions().GetP(), (func() int {
		if localctx.(*Instruccion_returnContext).Get_RETURN() == nil {
			return 0
		} else {
			return localctx.(*Instruccion_returnContext).Get_RETURN().GetLine()
		}
	}()), localctx.(*Instruccion_returnContext).Get_RETURN().GetColumn())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_funcContext is an interface to support dynamic dispatch.
type IInstruccion_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FUNC returns the _FUNC token.
	Get_FUNC() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FUNC sets the _FUNC token.
	Set_FUNC(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_instruccion_tipo returns the _instruccion_tipo rule contexts.
	Get_instruccion_tipo() IInstruccion_tipoContext

	// Get_list_function_parameters returns the _list_function_parameters rule contexts.
	Get_list_function_parameters() IList_function_parametersContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_instruccion_tipo sets the _instruccion_tipo rule contexts.
	Set_instruccion_tipo(IInstruccion_tipoContext)

	// Set_list_function_parameters sets the _list_function_parameters rule contexts.
	Set_list_function_parameters(IList_function_parametersContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	TK_PARA() antlr.TerminalNode
	TK_PARC() antlr.TerminalNode
	TK_LLAVEA() antlr.TerminalNode
	Instrucciones() IInstruccionesContext
	TK_LLAVEC() antlr.TerminalNode
	TK_MENOSMAYOR() antlr.TerminalNode
	Instruccion_tipo() IInstruccion_tipoContext
	List_function_parameters() IList_function_parametersContext

	// IsInstruccion_funcContext differentiates from other interfaces.
	IsInstruccion_funcContext()
}

type Instruccion_funcContext struct {
	antlr.BaseParserRuleContext
	parser                    antlr.Parser
	instr                     interfaces.Instruction
	_FUNC                     antlr.Token
	_ID                       antlr.Token
	_instrucciones            IInstruccionesContext
	_instruccion_tipo         IInstruccion_tipoContext
	_list_function_parameters IList_function_parametersContext
}

func NewEmptyInstruccion_funcContext() *Instruccion_funcContext {
	var p = new(Instruccion_funcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_func
	return p
}

func InitEmptyInstruccion_funcContext(p *Instruccion_funcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_func
}

func (*Instruccion_funcContext) IsInstruccion_funcContext() {}

func NewInstruccion_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_funcContext {
	var p = new(Instruccion_funcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_func

	return p
}

func (s *Instruccion_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_funcContext) Get_FUNC() antlr.Token { return s._FUNC }

func (s *Instruccion_funcContext) Get_ID() antlr.Token { return s._ID }

func (s *Instruccion_funcContext) Set_FUNC(v antlr.Token) { s._FUNC = v }

func (s *Instruccion_funcContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Instruccion_funcContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instruccion_funcContext) Get_instruccion_tipo() IInstruccion_tipoContext {
	return s._instruccion_tipo
}

func (s *Instruccion_funcContext) Get_list_function_parameters() IList_function_parametersContext {
	return s._list_function_parameters
}

func (s *Instruccion_funcContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instruccion_funcContext) Set_instruccion_tipo(v IInstruccion_tipoContext) {
	s._instruccion_tipo = v
}

func (s *Instruccion_funcContext) Set_list_function_parameters(v IList_function_parametersContext) {
	s._list_function_parameters = v
}

func (s *Instruccion_funcContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_funcContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_funcContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarFUNC, 0)
}

func (s *Instruccion_funcContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instruccion_funcContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARA, 0)
}

func (s *Instruccion_funcContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARC, 0)
}

func (s *Instruccion_funcContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instruccion_funcContext) Instrucciones() IInstruccionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instruccion_funcContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instruccion_funcContext) TK_MENOSMAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_MENOSMAYOR, 0)
}

func (s *Instruccion_funcContext) Instruccion_tipo() IInstruccion_tipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_tipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_tipoContext)
}

func (s *Instruccion_funcContext) List_function_parameters() IList_function_parametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_function_parametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_function_parametersContext)
}

func (s *Instruccion_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_funcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_func(s)
	}
}

func (s *Instruccion_funcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_func(s)
	}
}

func (p *Swiftgrammar) Instruccion_func() (localctx IInstruccion_funcContext) {
	localctx = NewInstruccion_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SwiftgrammarRULE_instruccion_func)
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(480)

			var _m = p.Match(SwiftgrammarFUNC)

			localctx.(*Instruccion_funcContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_funcContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(482)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_funcContext)._instrucciones = _x
		}
		{
			p.SetState(486)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_funcContext).instr = instruction.NewFunction((func() string {
			if localctx.(*Instruccion_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_funcContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*Instruccion_funcContext).Get_instrucciones().GetL(), interfaces.NULL, (func() int {
			if localctx.(*Instruccion_funcContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_funcContext).Get_FUNC().GetLine()
			}
		}()), localctx.(*Instruccion_funcContext).Get_FUNC().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(489)

			var _m = p.Match(SwiftgrammarFUNC)

			localctx.(*Instruccion_funcContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_funcContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Match(SwiftgrammarTK_MENOSMAYOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)

			var _x = p.Instruccion_tipo()

			localctx.(*Instruccion_funcContext)._instruccion_tipo = _x
		}
		{
			p.SetState(495)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(496)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_funcContext)._instrucciones = _x
		}
		{
			p.SetState(497)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_funcContext).instr = instruction.NewFunction((func() string {
			if localctx.(*Instruccion_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_funcContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*Instruccion_funcContext).Get_instrucciones().GetL(), localctx.(*Instruccion_funcContext).Get_instruccion_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Instruccion_funcContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_funcContext).Get_FUNC().GetLine()
			}
		}()), localctx.(*Instruccion_funcContext).Get_FUNC().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(500)

			var _m = p.Match(SwiftgrammarFUNC)

			localctx.(*Instruccion_funcContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(501)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_funcContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(502)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)

			var _x = p.List_function_parameters()

			localctx.(*Instruccion_funcContext)._list_function_parameters = _x
		}
		{
			p.SetState(504)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(506)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_funcContext)._instrucciones = _x
		}
		{
			p.SetState(507)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_funcContext).instr = instruction.NewFunction((func() string {
			if localctx.(*Instruccion_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Instruccion_funcContext).Get_list_function_parameters().GetL(), localctx.(*Instruccion_funcContext).Get_instrucciones().GetL(), interfaces.NULL, (func() int {
			if localctx.(*Instruccion_funcContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_funcContext).Get_FUNC().GetLine()
			}
		}()), localctx.(*Instruccion_funcContext).Get_FUNC().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(510)

			var _m = p.Match(SwiftgrammarFUNC)

			localctx.(*Instruccion_funcContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(511)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_funcContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(513)

			var _x = p.List_function_parameters()

			localctx.(*Instruccion_funcContext)._list_function_parameters = _x
		}
		{
			p.SetState(514)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(515)
			p.Match(SwiftgrammarTK_MENOSMAYOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(516)

			var _x = p.Instruccion_tipo()

			localctx.(*Instruccion_funcContext)._instruccion_tipo = _x
		}
		{
			p.SetState(517)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(518)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_funcContext)._instrucciones = _x
		}
		{
			p.SetState(519)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_funcContext).instr = instruction.NewFunction((func() string {
			if localctx.(*Instruccion_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Instruccion_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Instruccion_funcContext).Get_list_function_parameters().GetL(), localctx.(*Instruccion_funcContext).Get_instrucciones().GetL(), localctx.(*Instruccion_funcContext).Get_instruccion_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Instruccion_funcContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*Instruccion_funcContext).Get_FUNC().GetLine()
			}
		}()), localctx.(*Instruccion_funcContext).Get_FUNC().GetColumn())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_function_parametersContext is an interface to support dynamic dispatch.
type IList_function_parametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_parameters_fn returns the _block_parameters_fn rule contexts.
	Get_block_parameters_fn() IBlock_parameters_fnContext

	// Set_block_parameters_fn sets the _block_parameters_fn rule contexts.
	Set_block_parameters_fn(IBlock_parameters_fnContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_parameters_fnContext

	// SetE sets the e rule context list.
	SetE([]IBlock_parameters_fnContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// Getter signatures
	AllBlock_parameters_fn() []IBlock_parameters_fnContext
	Block_parameters_fn(i int) IBlock_parameters_fnContext

	// IsList_function_parametersContext differentiates from other interfaces.
	IsList_function_parametersContext()
}

type List_function_parametersContext struct {
	antlr.BaseParserRuleContext
	parser               antlr.Parser
	l                    *arrayList.List
	_block_parameters_fn IBlock_parameters_fnContext
	e                    []IBlock_parameters_fnContext
}

func NewEmptyList_function_parametersContext() *List_function_parametersContext {
	var p = new(List_function_parametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_function_parameters
	return p
}

func InitEmptyList_function_parametersContext(p *List_function_parametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_function_parameters
}

func (*List_function_parametersContext) IsList_function_parametersContext() {}

func NewList_function_parametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_function_parametersContext {
	var p = new(List_function_parametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_function_parameters

	return p
}

func (s *List_function_parametersContext) GetParser() antlr.Parser { return s.parser }

func (s *List_function_parametersContext) Get_block_parameters_fn() IBlock_parameters_fnContext {
	return s._block_parameters_fn
}

func (s *List_function_parametersContext) Set_block_parameters_fn(v IBlock_parameters_fnContext) {
	s._block_parameters_fn = v
}

func (s *List_function_parametersContext) GetE() []IBlock_parameters_fnContext { return s.e }

func (s *List_function_parametersContext) SetE(v []IBlock_parameters_fnContext) { s.e = v }

func (s *List_function_parametersContext) GetL() *arrayList.List { return s.l }

func (s *List_function_parametersContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_function_parametersContext) AllBlock_parameters_fn() []IBlock_parameters_fnContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_parameters_fnContext); ok {
			len++
		}
	}

	tst := make([]IBlock_parameters_fnContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_parameters_fnContext); ok {
			tst[i] = t.(IBlock_parameters_fnContext)
			i++
		}
	}

	return tst
}

func (s *List_function_parametersContext) Block_parameters_fn(i int) IBlock_parameters_fnContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_parameters_fnContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_parameters_fnContext)
}

func (s *List_function_parametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_function_parametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_function_parametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_function_parameters(s)
	}
}

func (s *List_function_parametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_function_parameters(s)
	}
}

func (p *Swiftgrammar) List_function_parameters() (localctx IList_function_parametersContext) {
	localctx = NewList_function_parametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SwiftgrammarRULE_list_function_parameters)

	localctx.(*List_function_parametersContext).l = arrayList.New()

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(525)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		//goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftgrammarID {
		{
			p.SetState(524)

			var _x = p.Block_parameters_fn()

			localctx.(*List_function_parametersContext)._block_parameters_fn = _x
		}
		localctx.(*List_function_parametersContext).e = append(localctx.(*List_function_parametersContext).e, localctx.(*List_function_parametersContext)._block_parameters_fn)

		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	//		goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_function_parametersContext).GetE()
	for _, e := range listInt {
		localctx.(*List_function_parametersContext).l.Add(e.GetInstr())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_parameters_fnContext is an interface to support dynamic dispatch.
type IBlock_parameters_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instruccion_tipo returns the _instruccion_tipo rule contexts.
	Get_instruccion_tipo() IInstruccion_tipoContext

	// Set_instruccion_tipo sets the _instruccion_tipo rule contexts.
	Set_instruccion_tipo(IInstruccion_tipoContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	TK_DOSPUNTOS() antlr.TerminalNode
	Instruccion_tipo() IInstruccion_tipoContext
	TK_COMA() antlr.TerminalNode

	// IsBlock_parameters_fnContext differentiates from other interfaces.
	IsBlock_parameters_fnContext()
}

type Block_parameters_fnContext struct {
	antlr.BaseParserRuleContext
	parser            antlr.Parser
	instr             interfaces.Instruction
	_ID               antlr.Token
	_instruccion_tipo IInstruccion_tipoContext
}

func NewEmptyBlock_parameters_fnContext() *Block_parameters_fnContext {
	var p = new(Block_parameters_fnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_parameters_fn
	return p
}

func InitEmptyBlock_parameters_fnContext(p *Block_parameters_fnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_parameters_fn
}

func (*Block_parameters_fnContext) IsBlock_parameters_fnContext() {}

func NewBlock_parameters_fnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_parameters_fnContext {
	var p = new(Block_parameters_fnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_parameters_fn

	return p
}

func (s *Block_parameters_fnContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_parameters_fnContext) Get_ID() antlr.Token { return s._ID }

func (s *Block_parameters_fnContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Block_parameters_fnContext) Get_instruccion_tipo() IInstruccion_tipoContext {
	return s._instruccion_tipo
}

func (s *Block_parameters_fnContext) Set_instruccion_tipo(v IInstruccion_tipoContext) {
	s._instruccion_tipo = v
}

func (s *Block_parameters_fnContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Block_parameters_fnContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Block_parameters_fnContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Block_parameters_fnContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DOSPUNTOS, 0)
}

func (s *Block_parameters_fnContext) Instruccion_tipo() IInstruccion_tipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_tipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_tipoContext)
}

func (s *Block_parameters_fnContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_COMA, 0)
}

func (s *Block_parameters_fnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_parameters_fnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_parameters_fnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_parameters_fn(s)
	}
}

func (s *Block_parameters_fnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_parameters_fn(s)
	}
}

func (p *Swiftgrammar) Block_parameters_fn() (localctx IBlock_parameters_fnContext) {
	localctx = NewBlock_parameters_fnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SwiftgrammarRULE_block_parameters_fn)
	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(531)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Block_parameters_fnContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(532)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(533)

			var _x = p.Instruccion_tipo()

			localctx.(*Block_parameters_fnContext)._instruccion_tipo = _x
		}
		{
			p.SetState(534)
			p.Match(SwiftgrammarTK_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Block_parameters_fnContext).instr = instruction.NewListExprefunc((func() string {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetText()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_instruccion_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetLine()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_ID().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(537)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Block_parameters_fnContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(538)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(539)

			var _x = p.Instruccion_tipo()

			localctx.(*Block_parameters_fnContext)._instruccion_tipo = _x
		}
		localctx.(*Block_parameters_fnContext).instr = instruction.NewListExprefunc((func() string {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetText()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_instruccion_tipo().GetTipo_exp(), (func() int {
			if localctx.(*Block_parameters_fnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Block_parameters_fnContext).Get_ID().GetLine()
			}
		}()), localctx.(*Block_parameters_fnContext).Get_ID().GetColumn())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_llamadaContext is an interface to support dynamic dispatch.
type IInstruccion_llamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TK_PARA() antlr.TerminalNode
	TK_PARC() antlr.TerminalNode
	List_expression() IList_expressionContext

	// IsInstruccion_llamadaContext differentiates from other interfaces.
	IsInstruccion_llamadaContext()
}

type Instruccion_llamadaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccion_llamadaContext() *Instruccion_llamadaContext {
	var p = new(Instruccion_llamadaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_llamada
	return p
}

func InitEmptyInstruccion_llamadaContext(p *Instruccion_llamadaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_llamada
}

func (*Instruccion_llamadaContext) IsInstruccion_llamadaContext() {}

func NewInstruccion_llamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_llamadaContext {
	var p = new(Instruccion_llamadaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_llamada

	return p
}

func (s *Instruccion_llamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_llamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instruccion_llamadaContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARA, 0)
}

func (s *Instruccion_llamadaContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARC, 0)
}

func (s *Instruccion_llamadaContext) List_expression() IList_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Instruccion_llamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_llamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_llamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_llamada(s)
	}
}

func (s *Instruccion_llamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_llamada(s)
	}
}

func (p *Swiftgrammar) Instruccion_llamada() (localctx IInstruccion_llamadaContext) {
	localctx = NewInstruccion_llamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SwiftgrammarRULE_instruccion_llamada)
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(544)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(546)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(547)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(548)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
			p.List_expression()
		}
		{
			p.SetState(550)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_llamada_expreContext is an interface to support dynamic dispatch.
type IInstr_llamada_expreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TK_PARA() antlr.TerminalNode
	TK_PARC() antlr.TerminalNode
	List_expression() IList_expressionContext

	// IsInstr_llamada_expreContext differentiates from other interfaces.
	IsInstr_llamada_expreContext()
}

type Instr_llamada_expreContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr_llamada_expreContext() *Instr_llamada_expreContext {
	var p = new(Instr_llamada_expreContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_llamada_expre
	return p
}

func InitEmptyInstr_llamada_expreContext(p *Instr_llamada_expreContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_llamada_expre
}

func (*Instr_llamada_expreContext) IsInstr_llamada_expreContext() {}

func NewInstr_llamada_expreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_llamada_expreContext {
	var p = new(Instr_llamada_expreContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_llamada_expre

	return p
}

func (s *Instr_llamada_expreContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_llamada_expreContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instr_llamada_expreContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARA, 0)
}

func (s *Instr_llamada_expreContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARC, 0)
}

func (s *Instr_llamada_expreContext) List_expression() IList_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Instr_llamada_expreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_llamada_expreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_llamada_expreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_llamada_expre(s)
	}
}

func (s *Instr_llamada_expreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_llamada_expre(s)
	}
}

func (p *Swiftgrammar) Instr_llamada_expre() (localctx IInstr_llamada_expreContext) {
	localctx = NewInstr_llamada_expreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SwiftgrammarRULE_instr_llamada_expre)
	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(554)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(555)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(556)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(557)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(558)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(559)
			p.List_expression()
		}
		{
			p.SetState(560)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_structs_declaracionContext is an interface to support dynamic dispatch.
type IInstruccion_structs_declaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	TK_LLAVEA() antlr.TerminalNode
	List_struct_parameters() IList_struct_parametersContext
	TK_LLAVEC() antlr.TerminalNode

	// IsInstruccion_structs_declaracionContext differentiates from other interfaces.
	IsInstruccion_structs_declaracionContext()
}

type Instruccion_structs_declaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccion_structs_declaracionContext() *Instruccion_structs_declaracionContext {
	var p = new(Instruccion_structs_declaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_structs_declaracion
	return p
}

func InitEmptyInstruccion_structs_declaracionContext(p *Instruccion_structs_declaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_structs_declaracion
}

func (*Instruccion_structs_declaracionContext) IsInstruccion_structs_declaracionContext() {}

func NewInstruccion_structs_declaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_structs_declaracionContext {
	var p = new(Instruccion_structs_declaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_structs_declaracion

	return p
}

func (s *Instruccion_structs_declaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_structs_declaracionContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarSTRUCT, 0)
}

func (s *Instruccion_structs_declaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instruccion_structs_declaracionContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instruccion_structs_declaracionContext) List_struct_parameters() IList_struct_parametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_struct_parametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_struct_parametersContext)
}

func (s *Instruccion_structs_declaracionContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instruccion_structs_declaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_structs_declaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_structs_declaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_structs_declaracion(s)
	}
}

func (s *Instruccion_structs_declaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_structs_declaracion(s)
	}
}

func (p *Swiftgrammar) Instruccion_structs_declaracion() (localctx IInstruccion_structs_declaracionContext) {
	localctx = NewInstruccion_structs_declaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SwiftgrammarRULE_instruccion_structs_declaracion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(564)
		p.Match(SwiftgrammarSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(565)
		p.Match(SwiftgrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(566)
		p.Match(SwiftgrammarTK_LLAVEA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(567)
		p.List_struct_parameters()
	}
	{
		p.SetState(568)
		p.Match(SwiftgrammarTK_LLAVEC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_struct_parametersContext is an interface to support dynamic dispatch.
type IList_struct_parametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_structs_parameters returns the _block_structs_parameters rule contexts.
	Get_block_structs_parameters() IBlock_structs_parametersContext

	// Set_block_structs_parameters sets the _block_structs_parameters rule contexts.
	Set_block_structs_parameters(IBlock_structs_parametersContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_structs_parametersContext

	// SetE sets the e rule context list.
	SetE([]IBlock_structs_parametersContext)

	// Getter signatures
	AllBlock_structs_parameters() []IBlock_structs_parametersContext
	Block_structs_parameters(i int) IBlock_structs_parametersContext

	// IsList_struct_parametersContext differentiates from other interfaces.
	IsList_struct_parametersContext()
}

type List_struct_parametersContext struct {
	antlr.BaseParserRuleContext
	parser                    antlr.Parser
	_block_structs_parameters IBlock_structs_parametersContext
	e                         []IBlock_structs_parametersContext
}

func NewEmptyList_struct_parametersContext() *List_struct_parametersContext {
	var p = new(List_struct_parametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_struct_parameters
	return p
}

func InitEmptyList_struct_parametersContext(p *List_struct_parametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_struct_parameters
}

func (*List_struct_parametersContext) IsList_struct_parametersContext() {}

func NewList_struct_parametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_struct_parametersContext {
	var p = new(List_struct_parametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_struct_parameters

	return p
}

func (s *List_struct_parametersContext) GetParser() antlr.Parser { return s.parser }

func (s *List_struct_parametersContext) Get_block_structs_parameters() IBlock_structs_parametersContext {
	return s._block_structs_parameters
}

func (s *List_struct_parametersContext) Set_block_structs_parameters(v IBlock_structs_parametersContext) {
	s._block_structs_parameters = v
}

func (s *List_struct_parametersContext) GetE() []IBlock_structs_parametersContext { return s.e }

func (s *List_struct_parametersContext) SetE(v []IBlock_structs_parametersContext) { s.e = v }

func (s *List_struct_parametersContext) AllBlock_structs_parameters() []IBlock_structs_parametersContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_structs_parametersContext); ok {
			len++
		}
	}

	tst := make([]IBlock_structs_parametersContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_structs_parametersContext); ok {
			tst[i] = t.(IBlock_structs_parametersContext)
			i++
		}
	}

	return tst
}

func (s *List_struct_parametersContext) Block_structs_parameters(i int) IBlock_structs_parametersContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_structs_parametersContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_structs_parametersContext)
}

func (s *List_struct_parametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_struct_parametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_struct_parametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_struct_parameters(s)
	}
}

func (s *List_struct_parametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_struct_parameters(s)
	}
}

func (p *Swiftgrammar) List_struct_parameters() (localctx IList_struct_parametersContext) {
	localctx = NewList_struct_parametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SwiftgrammarRULE_list_struct_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(571)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftgrammarID {
		{
			p.SetState(570)

			var _x = p.Block_structs_parameters()

			localctx.(*List_struct_parametersContext)._block_structs_parameters = _x
		}
		localctx.(*List_struct_parametersContext).e = append(localctx.(*List_struct_parametersContext).e, localctx.(*List_struct_parametersContext)._block_structs_parameters)

		p.SetState(573)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_structs_parametersContext is an interface to support dynamic dispatch.
type IBlock_structs_parametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TK_DOSPUNTOS() antlr.TerminalNode
	Instruccion_tipo() IInstruccion_tipoContext
	TK_COMA() antlr.TerminalNode

	// IsBlock_structs_parametersContext differentiates from other interfaces.
	IsBlock_structs_parametersContext()
}

type Block_structs_parametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_structs_parametersContext() *Block_structs_parametersContext {
	var p = new(Block_structs_parametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_structs_parameters
	return p
}

func InitEmptyBlock_structs_parametersContext(p *Block_structs_parametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_structs_parameters
}

func (*Block_structs_parametersContext) IsBlock_structs_parametersContext() {}

func NewBlock_structs_parametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_structs_parametersContext {
	var p = new(Block_structs_parametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_structs_parameters

	return p
}

func (s *Block_structs_parametersContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_structs_parametersContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Block_structs_parametersContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DOSPUNTOS, 0)
}

func (s *Block_structs_parametersContext) Instruccion_tipo() IInstruccion_tipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_tipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_tipoContext)
}

func (s *Block_structs_parametersContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_COMA, 0)
}

func (s *Block_structs_parametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_structs_parametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_structs_parametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_structs_parameters(s)
	}
}

func (s *Block_structs_parametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_structs_parameters(s)
	}
}

func (p *Swiftgrammar) Block_structs_parameters() (localctx IBlock_structs_parametersContext) {
	localctx = NewBlock_structs_parametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SwiftgrammarRULE_block_structs_parameters)
	p.SetState(583)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(575)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(576)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(577)
			p.Instruccion_tipo()
		}
		{
			p.SetState(578)
			p.Match(SwiftgrammarTK_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(580)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(581)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(582)
			p.Instruccion_tipo()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_arrays_identifierContext is an interface to support dynamic dispatch.
type IInstr_arrays_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	List_arrays_parameters_id() IList_arrays_parameters_idContext

	// IsInstr_arrays_identifierContext differentiates from other interfaces.
	IsInstr_arrays_identifierContext()
}

type Instr_arrays_identifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr_arrays_identifierContext() *Instr_arrays_identifierContext {
	var p = new(Instr_arrays_identifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_arrays_identifier
	return p
}

func InitEmptyInstr_arrays_identifierContext(p *Instr_arrays_identifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_arrays_identifier
}

func (*Instr_arrays_identifierContext) IsInstr_arrays_identifierContext() {}

func NewInstr_arrays_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_arrays_identifierContext {
	var p = new(Instr_arrays_identifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_arrays_identifier

	return p
}

func (s *Instr_arrays_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_arrays_identifierContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instr_arrays_identifierContext) List_arrays_parameters_id() IList_arrays_parameters_idContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_arrays_parameters_idContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_arrays_parameters_idContext)
}

func (s *Instr_arrays_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_arrays_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_arrays_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_arrays_identifier(s)
	}
}

func (s *Instr_arrays_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_arrays_identifier(s)
	}
}

func (p *Swiftgrammar) Instr_arrays_identifier() (localctx IInstr_arrays_identifierContext) {
	localctx = NewInstr_arrays_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SwiftgrammarRULE_instr_arrays_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(585)
		p.Match(SwiftgrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(586)
		p.List_arrays_parameters_id()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_arrays_parameters_idContext is an interface to support dynamic dispatch.
type IList_arrays_parameters_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_arrays_identifier returns the _block_arrays_identifier rule contexts.
	Get_block_arrays_identifier() IBlock_arrays_identifierContext

	// Set_block_arrays_identifier sets the _block_arrays_identifier rule contexts.
	Set_block_arrays_identifier(IBlock_arrays_identifierContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_arrays_identifierContext

	// SetE sets the e rule context list.
	SetE([]IBlock_arrays_identifierContext)

	// Getter signatures
	AllBlock_arrays_identifier() []IBlock_arrays_identifierContext
	Block_arrays_identifier(i int) IBlock_arrays_identifierContext

	// IsList_arrays_parameters_idContext differentiates from other interfaces.
	IsList_arrays_parameters_idContext()
}

type List_arrays_parameters_idContext struct {
	antlr.BaseParserRuleContext
	parser                   antlr.Parser
	_block_arrays_identifier IBlock_arrays_identifierContext
	e                        []IBlock_arrays_identifierContext
}

func NewEmptyList_arrays_parameters_idContext() *List_arrays_parameters_idContext {
	var p = new(List_arrays_parameters_idContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_arrays_parameters_id
	return p
}

func InitEmptyList_arrays_parameters_idContext(p *List_arrays_parameters_idContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_arrays_parameters_id
}

func (*List_arrays_parameters_idContext) IsList_arrays_parameters_idContext() {}

func NewList_arrays_parameters_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_arrays_parameters_idContext {
	var p = new(List_arrays_parameters_idContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_arrays_parameters_id

	return p
}

func (s *List_arrays_parameters_idContext) GetParser() antlr.Parser { return s.parser }

func (s *List_arrays_parameters_idContext) Get_block_arrays_identifier() IBlock_arrays_identifierContext {
	return s._block_arrays_identifier
}

func (s *List_arrays_parameters_idContext) Set_block_arrays_identifier(v IBlock_arrays_identifierContext) {
	s._block_arrays_identifier = v
}

func (s *List_arrays_parameters_idContext) GetE() []IBlock_arrays_identifierContext { return s.e }

func (s *List_arrays_parameters_idContext) SetE(v []IBlock_arrays_identifierContext) { s.e = v }

func (s *List_arrays_parameters_idContext) AllBlock_arrays_identifier() []IBlock_arrays_identifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_arrays_identifierContext); ok {
			len++
		}
	}

	tst := make([]IBlock_arrays_identifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_arrays_identifierContext); ok {
			tst[i] = t.(IBlock_arrays_identifierContext)
			i++
		}
	}

	return tst
}

func (s *List_arrays_parameters_idContext) Block_arrays_identifier(i int) IBlock_arrays_identifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_arrays_identifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_arrays_identifierContext)
}

func (s *List_arrays_parameters_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_arrays_parameters_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_arrays_parameters_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_arrays_parameters_id(s)
	}
}

func (s *List_arrays_parameters_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_arrays_parameters_id(s)
	}
}

func (p *Swiftgrammar) List_arrays_parameters_id() (localctx IList_arrays_parameters_idContext) {
	localctx = NewList_arrays_parameters_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SwiftgrammarRULE_list_arrays_parameters_id)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(588)

				var _x = p.Block_arrays_identifier()

				localctx.(*List_arrays_parameters_idContext)._block_arrays_identifier = _x
			}
			localctx.(*List_arrays_parameters_idContext).e = append(localctx.(*List_arrays_parameters_idContext).e, localctx.(*List_arrays_parameters_idContext)._block_arrays_identifier)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_arrays_identifierContext is an interface to support dynamic dispatch.
type IBlock_arrays_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_CORA() antlr.TerminalNode
	Expressions() IExpressionsContext
	TK_CORC() antlr.TerminalNode

	// IsBlock_arrays_identifierContext differentiates from other interfaces.
	IsBlock_arrays_identifierContext()
}

type Block_arrays_identifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_arrays_identifierContext() *Block_arrays_identifierContext {
	var p = new(Block_arrays_identifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_arrays_identifier
	return p
}

func InitEmptyBlock_arrays_identifierContext(p *Block_arrays_identifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_arrays_identifier
}

func (*Block_arrays_identifierContext) IsBlock_arrays_identifierContext() {}

func NewBlock_arrays_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_arrays_identifierContext {
	var p = new(Block_arrays_identifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_arrays_identifier

	return p
}

func (s *Block_arrays_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_arrays_identifierContext) TK_CORA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_CORA, 0)
}

func (s *Block_arrays_identifierContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_arrays_identifierContext) TK_CORC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_CORC, 0)
}

func (s *Block_arrays_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_arrays_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_arrays_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_arrays_identifier(s)
	}
}

func (s *Block_arrays_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_arrays_identifier(s)
	}
}

func (p *Swiftgrammar) Block_arrays_identifier() (localctx IBlock_arrays_identifierContext) {
	localctx = NewBlock_arrays_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SwiftgrammarRULE_block_arrays_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(593)
		p.Match(SwiftgrammarTK_CORA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(594)
		p.Expressions()
	}
	{
		p.SetState(595)
		p.Match(SwiftgrammarTK_CORC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_structs_declarationContext is an interface to support dynamic dispatch.
type IInstr_structs_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left token.
	GetLeft() antlr.Token

	// GetRight returns the right token.
	GetRight() antlr.Token

	// SetLeft sets the left token.
	SetLeft(antlr.Token)

	// SetRight sets the right token.
	SetRight(antlr.Token)

	// Getter signatures
	VAR() antlr.TerminalNode
	TK_IGUAL() antlr.TerminalNode
	TK_LLAVEA() antlr.TerminalNode
	List_struct_parameters_decla() IList_struct_parameters_declaContext
	TK_LLAVEC() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	LET() antlr.TerminalNode

	// IsInstr_structs_declarationContext differentiates from other interfaces.
	IsInstr_structs_declarationContext()
}

type Instr_structs_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   antlr.Token
	right  antlr.Token
}

func NewEmptyInstr_structs_declarationContext() *Instr_structs_declarationContext {
	var p = new(Instr_structs_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_structs_declaration
	return p
}

func InitEmptyInstr_structs_declarationContext(p *Instr_structs_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_structs_declaration
}

func (*Instr_structs_declarationContext) IsInstr_structs_declarationContext() {}

func NewInstr_structs_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_structs_declarationContext {
	var p = new(Instr_structs_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_structs_declaration

	return p
}

func (s *Instr_structs_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_structs_declarationContext) GetLeft() antlr.Token { return s.left }

func (s *Instr_structs_declarationContext) GetRight() antlr.Token { return s.right }

func (s *Instr_structs_declarationContext) SetLeft(v antlr.Token) { s.left = v }

func (s *Instr_structs_declarationContext) SetRight(v antlr.Token) { s.right = v }

func (s *Instr_structs_declarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarVAR, 0)
}

func (s *Instr_structs_declarationContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_IGUAL, 0)
}

func (s *Instr_structs_declarationContext) TK_LLAVEA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEA, 0)
}

func (s *Instr_structs_declarationContext) List_struct_parameters_decla() IList_struct_parameters_declaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_struct_parameters_declaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_struct_parameters_declaContext)
}

func (s *Instr_structs_declarationContext) TK_LLAVEC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_LLAVEC, 0)
}

func (s *Instr_structs_declarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftgrammarID)
}

func (s *Instr_structs_declarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, i)
}

func (s *Instr_structs_declarationContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarLET, 0)
}

func (s *Instr_structs_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_structs_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_structs_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_structs_declaration(s)
	}
}

func (s *Instr_structs_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_structs_declaration(s)
	}
}

func (p *Swiftgrammar) Instr_structs_declaration() (localctx IInstr_structs_declarationContext) {
	localctx = NewInstr_structs_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SwiftgrammarRULE_instr_structs_declaration)
	p.SetState(613)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammarVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(597)
			p.Match(SwiftgrammarVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(598)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instr_structs_declarationContext).left = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(599)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(600)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instr_structs_declarationContext).right = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(601)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)
			p.List_struct_parameters_decla()
		}
		{
			p.SetState(603)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftgrammarLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(605)
			p.Match(SwiftgrammarLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(606)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instr_structs_declarationContext).left = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(607)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(608)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instr_structs_declarationContext).right = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(609)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(610)
			p.List_struct_parameters_decla()
		}
		{
			p.SetState(611)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_struct_parameters_declaContext is an interface to support dynamic dispatch.
type IList_struct_parameters_declaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_structs_parameters_decla returns the _block_structs_parameters_decla rule contexts.
	Get_block_structs_parameters_decla() IBlock_structs_parameters_declaContext

	// Set_block_structs_parameters_decla sets the _block_structs_parameters_decla rule contexts.
	Set_block_structs_parameters_decla(IBlock_structs_parameters_declaContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_structs_parameters_declaContext

	// SetE sets the e rule context list.
	SetE([]IBlock_structs_parameters_declaContext)

	// Getter signatures
	AllBlock_structs_parameters_decla() []IBlock_structs_parameters_declaContext
	Block_structs_parameters_decla(i int) IBlock_structs_parameters_declaContext

	// IsList_struct_parameters_declaContext differentiates from other interfaces.
	IsList_struct_parameters_declaContext()
}

type List_struct_parameters_declaContext struct {
	antlr.BaseParserRuleContext
	parser                          antlr.Parser
	_block_structs_parameters_decla IBlock_structs_parameters_declaContext
	e                               []IBlock_structs_parameters_declaContext
}

func NewEmptyList_struct_parameters_declaContext() *List_struct_parameters_declaContext {
	var p = new(List_struct_parameters_declaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_struct_parameters_decla
	return p
}

func InitEmptyList_struct_parameters_declaContext(p *List_struct_parameters_declaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_struct_parameters_decla
}

func (*List_struct_parameters_declaContext) IsList_struct_parameters_declaContext() {}

func NewList_struct_parameters_declaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_struct_parameters_declaContext {
	var p = new(List_struct_parameters_declaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_struct_parameters_decla

	return p
}

func (s *List_struct_parameters_declaContext) GetParser() antlr.Parser { return s.parser }

func (s *List_struct_parameters_declaContext) Get_block_structs_parameters_decla() IBlock_structs_parameters_declaContext {
	return s._block_structs_parameters_decla
}

func (s *List_struct_parameters_declaContext) Set_block_structs_parameters_decla(v IBlock_structs_parameters_declaContext) {
	s._block_structs_parameters_decla = v
}

func (s *List_struct_parameters_declaContext) GetE() []IBlock_structs_parameters_declaContext {
	return s.e
}

func (s *List_struct_parameters_declaContext) SetE(v []IBlock_structs_parameters_declaContext) {
	s.e = v
}

func (s *List_struct_parameters_declaContext) AllBlock_structs_parameters_decla() []IBlock_structs_parameters_declaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_structs_parameters_declaContext); ok {
			len++
		}
	}

	tst := make([]IBlock_structs_parameters_declaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_structs_parameters_declaContext); ok {
			tst[i] = t.(IBlock_structs_parameters_declaContext)
			i++
		}
	}

	return tst
}

func (s *List_struct_parameters_declaContext) Block_structs_parameters_decla(i int) IBlock_structs_parameters_declaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_structs_parameters_declaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_structs_parameters_declaContext)
}

func (s *List_struct_parameters_declaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_struct_parameters_declaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_struct_parameters_declaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_struct_parameters_decla(s)
	}
}

func (s *List_struct_parameters_declaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_struct_parameters_decla(s)
	}
}

func (p *Swiftgrammar) List_struct_parameters_decla() (localctx IList_struct_parameters_declaContext) {
	localctx = NewList_struct_parameters_declaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SwiftgrammarRULE_list_struct_parameters_decla)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftgrammarID {
		{
			p.SetState(615)

			var _x = p.Block_structs_parameters_decla()

			localctx.(*List_struct_parameters_declaContext)._block_structs_parameters_decla = _x
		}
		localctx.(*List_struct_parameters_declaContext).e = append(localctx.(*List_struct_parameters_declaContext).e, localctx.(*List_struct_parameters_declaContext)._block_structs_parameters_decla)

		p.SetState(618)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_structs_parameters_declaContext is an interface to support dynamic dispatch.
type IBlock_structs_parameters_declaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TK_DOSPUNTOS() antlr.TerminalNode
	Expressions() IExpressionsContext
	TK_COMA() antlr.TerminalNode

	// IsBlock_structs_parameters_declaContext differentiates from other interfaces.
	IsBlock_structs_parameters_declaContext()
}

type Block_structs_parameters_declaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_structs_parameters_declaContext() *Block_structs_parameters_declaContext {
	var p = new(Block_structs_parameters_declaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_structs_parameters_decla
	return p
}

func InitEmptyBlock_structs_parameters_declaContext(p *Block_structs_parameters_declaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_structs_parameters_decla
}

func (*Block_structs_parameters_declaContext) IsBlock_structs_parameters_declaContext() {}

func NewBlock_structs_parameters_declaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_structs_parameters_declaContext {
	var p = new(Block_structs_parameters_declaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_structs_parameters_decla

	return p
}

func (s *Block_structs_parameters_declaContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_structs_parameters_declaContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Block_structs_parameters_declaContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DOSPUNTOS, 0)
}

func (s *Block_structs_parameters_declaContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_structs_parameters_declaContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_COMA, 0)
}

func (s *Block_structs_parameters_declaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_structs_parameters_declaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_structs_parameters_declaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_structs_parameters_decla(s)
	}
}

func (s *Block_structs_parameters_declaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_structs_parameters_decla(s)
	}
}

func (p *Swiftgrammar) Block_structs_parameters_decla() (localctx IBlock_structs_parameters_declaContext) {
	localctx = NewBlock_structs_parameters_declaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SwiftgrammarRULE_block_structs_parameters_decla)
	p.SetState(628)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(620)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)
			p.Expressions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(623)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(624)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(625)
			p.Expressions()
		}
		{
			p.SetState(626)
			p.Match(SwiftgrammarTK_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_structs_identifierContext is an interface to support dynamic dispatch.
type IInstr_structs_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	List_struct_parameters_id() IList_struct_parameters_idContext

	// IsInstr_structs_identifierContext differentiates from other interfaces.
	IsInstr_structs_identifierContext()
}

type Instr_structs_identifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr_structs_identifierContext() *Instr_structs_identifierContext {
	var p = new(Instr_structs_identifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_structs_identifier
	return p
}

func InitEmptyInstr_structs_identifierContext(p *Instr_structs_identifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_structs_identifier
}

func (*Instr_structs_identifierContext) IsInstr_structs_identifierContext() {}

func NewInstr_structs_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_structs_identifierContext {
	var p = new(Instr_structs_identifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_structs_identifier

	return p
}

func (s *Instr_structs_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_structs_identifierContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instr_structs_identifierContext) List_struct_parameters_id() IList_struct_parameters_idContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_struct_parameters_idContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_struct_parameters_idContext)
}

func (s *Instr_structs_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_structs_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_structs_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_structs_identifier(s)
	}
}

func (s *Instr_structs_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_structs_identifier(s)
	}
}

func (p *Swiftgrammar) Instr_structs_identifier() (localctx IInstr_structs_identifierContext) {
	localctx = NewInstr_structs_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SwiftgrammarRULE_instr_structs_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(630)
		p.Match(SwiftgrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(631)
		p.List_struct_parameters_id()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_struct_parameters_idContext is an interface to support dynamic dispatch.
type IList_struct_parameters_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_structs_identifier returns the _block_structs_identifier rule contexts.
	Get_block_structs_identifier() IBlock_structs_identifierContext

	// Set_block_structs_identifier sets the _block_structs_identifier rule contexts.
	Set_block_structs_identifier(IBlock_structs_identifierContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_structs_identifierContext

	// SetE sets the e rule context list.
	SetE([]IBlock_structs_identifierContext)

	// Getter signatures
	AllBlock_structs_identifier() []IBlock_structs_identifierContext
	Block_structs_identifier(i int) IBlock_structs_identifierContext

	// IsList_struct_parameters_idContext differentiates from other interfaces.
	IsList_struct_parameters_idContext()
}

type List_struct_parameters_idContext struct {
	antlr.BaseParserRuleContext
	parser                    antlr.Parser
	_block_structs_identifier IBlock_structs_identifierContext
	e                         []IBlock_structs_identifierContext
}

func NewEmptyList_struct_parameters_idContext() *List_struct_parameters_idContext {
	var p = new(List_struct_parameters_idContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_struct_parameters_id
	return p
}

func InitEmptyList_struct_parameters_idContext(p *List_struct_parameters_idContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_struct_parameters_id
}

func (*List_struct_parameters_idContext) IsList_struct_parameters_idContext() {}

func NewList_struct_parameters_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_struct_parameters_idContext {
	var p = new(List_struct_parameters_idContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_struct_parameters_id

	return p
}

func (s *List_struct_parameters_idContext) GetParser() antlr.Parser { return s.parser }

func (s *List_struct_parameters_idContext) Get_block_structs_identifier() IBlock_structs_identifierContext {
	return s._block_structs_identifier
}

func (s *List_struct_parameters_idContext) Set_block_structs_identifier(v IBlock_structs_identifierContext) {
	s._block_structs_identifier = v
}

func (s *List_struct_parameters_idContext) GetE() []IBlock_structs_identifierContext { return s.e }

func (s *List_struct_parameters_idContext) SetE(v []IBlock_structs_identifierContext) { s.e = v }

func (s *List_struct_parameters_idContext) AllBlock_structs_identifier() []IBlock_structs_identifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_structs_identifierContext); ok {
			len++
		}
	}

	tst := make([]IBlock_structs_identifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_structs_identifierContext); ok {
			tst[i] = t.(IBlock_structs_identifierContext)
			i++
		}
	}

	return tst
}

func (s *List_struct_parameters_idContext) Block_structs_identifier(i int) IBlock_structs_identifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_structs_identifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_structs_identifierContext)
}

func (s *List_struct_parameters_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_struct_parameters_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_struct_parameters_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_struct_parameters_id(s)
	}
}

func (s *List_struct_parameters_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_struct_parameters_id(s)
	}
}

func (p *Swiftgrammar) List_struct_parameters_id() (localctx IList_struct_parameters_idContext) {
	localctx = NewList_struct_parameters_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SwiftgrammarRULE_list_struct_parameters_id)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(634)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(633)

				var _x = p.Block_structs_identifier()

				localctx.(*List_struct_parameters_idContext)._block_structs_identifier = _x
			}
			localctx.(*List_struct_parameters_idContext).e = append(localctx.(*List_struct_parameters_idContext).e, localctx.(*List_struct_parameters_idContext)._block_structs_identifier)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(636)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_structs_identifierContext is an interface to support dynamic dispatch.
type IBlock_structs_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_PUNTO() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsBlock_structs_identifierContext differentiates from other interfaces.
	IsBlock_structs_identifierContext()
}

type Block_structs_identifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_structs_identifierContext() *Block_structs_identifierContext {
	var p = new(Block_structs_identifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_structs_identifier
	return p
}

func InitEmptyBlock_structs_identifierContext(p *Block_structs_identifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_structs_identifier
}

func (*Block_structs_identifierContext) IsBlock_structs_identifierContext() {}

func NewBlock_structs_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_structs_identifierContext {
	var p = new(Block_structs_identifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_structs_identifier

	return p
}

func (s *Block_structs_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_structs_identifierContext) TK_PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PUNTO, 0)
}

func (s *Block_structs_identifierContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Block_structs_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_structs_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_structs_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_structs_identifier(s)
	}
}

func (s *Block_structs_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_structs_identifier(s)
	}
}

func (p *Swiftgrammar) Block_structs_identifier() (localctx IBlock_structs_identifierContext) {
	localctx = NewBlock_structs_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SwiftgrammarRULE_block_structs_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(638)
		p.Match(SwiftgrammarTK_PUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(639)
		p.Match(SwiftgrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstr_structs_assignmentContext is an interface to support dynamic dispatch.
type IInstr_structs_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	List_struct_parameters_id() IList_struct_parameters_idContext
	TK_IGUAL() antlr.TerminalNode
	Expressions() IExpressionsContext

	// IsInstr_structs_assignmentContext differentiates from other interfaces.
	IsInstr_structs_assignmentContext()
}

type Instr_structs_assignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstr_structs_assignmentContext() *Instr_structs_assignmentContext {
	var p = new(Instr_structs_assignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_structs_assignment
	return p
}

func InitEmptyInstr_structs_assignmentContext(p *Instr_structs_assignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instr_structs_assignment
}

func (*Instr_structs_assignmentContext) IsInstr_structs_assignmentContext() {}

func NewInstr_structs_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_structs_assignmentContext {
	var p = new(Instr_structs_assignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instr_structs_assignment

	return p
}

func (s *Instr_structs_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_structs_assignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *Instr_structs_assignmentContext) List_struct_parameters_id() IList_struct_parameters_idContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IList_struct_parameters_idContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IList_struct_parameters_idContext)
}

func (s *Instr_structs_assignmentContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_IGUAL, 0)
}

func (s *Instr_structs_assignmentContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Instr_structs_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_structs_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_structs_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstr_structs_assignment(s)
	}
}

func (s *Instr_structs_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstr_structs_assignment(s)
	}
}

func (p *Swiftgrammar) Instr_structs_assignment() (localctx IInstr_structs_assignmentContext) {
	localctx = NewInstr_structs_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SwiftgrammarRULE_instr_structs_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(641)
		p.Match(SwiftgrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(642)
		p.List_struct_parameters_id()
	}
	{
		p.SetState(643)
		p.Match(SwiftgrammarTK_IGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(644)
		p.Expressions()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstruccion_tipoContext is an interface to support dynamic dispatch.
type IInstruccion_tipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo_exp returns the tipo_exp attribute.
	GetTipo_exp() interfaces.TypeExpression

	// SetTipo_exp sets the tipo_exp attribute.
	SetTipo_exp(interfaces.TypeExpression)

	// Getter signatures
	R_INT() antlr.TerminalNode
	R_FLOAT() antlr.TerminalNode
	R_STRING() antlr.TerminalNode
	R_BOOL() antlr.TerminalNode
	R_CHAR() antlr.TerminalNode

	// IsInstruccion_tipoContext differentiates from other interfaces.
	IsInstruccion_tipoContext()
}

type Instruccion_tipoContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	tipo_exp interfaces.TypeExpression
}

func NewEmptyInstruccion_tipoContext() *Instruccion_tipoContext {
	var p = new(Instruccion_tipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_tipo
	return p
}

func InitEmptyInstruccion_tipoContext(p *Instruccion_tipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_instruccion_tipo
}

func (*Instruccion_tipoContext) IsInstruccion_tipoContext() {}

func NewInstruccion_tipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_tipoContext {
	var p = new(Instruccion_tipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_instruccion_tipo

	return p
}

func (s *Instruccion_tipoContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_tipoContext) GetTipo_exp() interfaces.TypeExpression { return s.tipo_exp }

func (s *Instruccion_tipoContext) SetTipo_exp(v interfaces.TypeExpression) { s.tipo_exp = v }

func (s *Instruccion_tipoContext) R_INT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarR_INT, 0)
}

func (s *Instruccion_tipoContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarR_FLOAT, 0)
}

func (s *Instruccion_tipoContext) R_STRING() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarR_STRING, 0)
}

func (s *Instruccion_tipoContext) R_BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarR_BOOL, 0)
}

func (s *Instruccion_tipoContext) R_CHAR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarR_CHAR, 0)
}

func (s *Instruccion_tipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_tipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_tipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterInstruccion_tipo(s)
	}
}

func (s *Instruccion_tipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitInstruccion_tipo(s)
	}
}

func (p *Swiftgrammar) Instruccion_tipo() (localctx IInstruccion_tipoContext) {
	localctx = NewInstruccion_tipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SwiftgrammarRULE_instruccion_tipo)
	p.SetState(656)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammarR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(646)
			p.Match(SwiftgrammarR_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_tipoContext).tipo_exp = interfaces.INTEGER

	case SwiftgrammarR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(648)
			p.Match(SwiftgrammarR_FLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_tipoContext).tipo_exp = interfaces.FLOAT

	case SwiftgrammarR_STRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(650)
			p.Match(SwiftgrammarR_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_tipoContext).tipo_exp = interfaces.STRING

	case SwiftgrammarR_BOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(652)
			p.Match(SwiftgrammarR_BOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_tipoContext).tipo_exp = interfaces.BOOLEAN

	case SwiftgrammarR_CHAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(654)
			p.Match(SwiftgrammarR_CHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Instruccion_tipoContext).tipo_exp = interfaces.CHAR

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IList_expressionContext is an interface to support dynamic dispatch.
type IList_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block_list_expression returns the _block_list_expression rule contexts.
	Get_block_list_expression() IBlock_list_expressionContext

	// Set_block_list_expression sets the _block_list_expression rule contexts.
	Set_block_list_expression(IBlock_list_expressionContext)

	// GetE returns the e rule context list.
	GetE() []IBlock_list_expressionContext

	// SetE sets the e rule context list.
	SetE([]IBlock_list_expressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// Getter signatures
	AllBlock_list_expression() []IBlock_list_expressionContext
	Block_list_expression(i int) IBlock_list_expressionContext

	// IsList_expressionContext differentiates from other interfaces.
	IsList_expressionContext()
}

type List_expressionContext struct {
	antlr.BaseParserRuleContext
	parser                 antlr.Parser
	l                      *arrayList.List
	_block_list_expression IBlock_list_expressionContext
	e                      []IBlock_list_expressionContext
}

func NewEmptyList_expressionContext() *List_expressionContext {
	var p = new(List_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_expression
	return p
}

func InitEmptyList_expressionContext(p *List_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_list_expression
}

func (*List_expressionContext) IsList_expressionContext() {}

func NewList_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expressionContext {
	var p = new(List_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_list_expression

	return p
}

func (s *List_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expressionContext) Get_block_list_expression() IBlock_list_expressionContext {
	return s._block_list_expression
}

func (s *List_expressionContext) Set_block_list_expression(v IBlock_list_expressionContext) {
	s._block_list_expression = v
}

func (s *List_expressionContext) GetE() []IBlock_list_expressionContext { return s.e }

func (s *List_expressionContext) SetE(v []IBlock_list_expressionContext) { s.e = v }

func (s *List_expressionContext) GetL() *arrayList.List { return s.l }

func (s *List_expressionContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_expressionContext) AllBlock_list_expression() []IBlock_list_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlock_list_expressionContext); ok {
			len++
		}
	}

	tst := make([]IBlock_list_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlock_list_expressionContext); ok {
			tst[i] = t.(IBlock_list_expressionContext)
			i++
		}
	}

	return tst
}

func (s *List_expressionContext) Block_list_expression(i int) IBlock_list_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_list_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_list_expressionContext)
}

func (s *List_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterList_expression(s)
	}
}

func (s *List_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitList_expression(s)
	}
}

func (p *Swiftgrammar) List_expression() (localctx IList_expressionContext) {
	localctx = NewList_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SwiftgrammarRULE_list_expression)

	localctx.(*List_expressionContext).l = arrayList.New()

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(659)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		//goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520464) != 0) {
		{
			p.SetState(658)

			var _x = p.Block_list_expression()

			localctx.(*List_expressionContext)._block_list_expression = _x
		}
		localctx.(*List_expressionContext).e = append(localctx.(*List_expressionContext).e, localctx.(*List_expressionContext)._block_list_expression)

		p.SetState(661)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			//goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*List_expressionContext).GetE()
	for _, e := range listInt {
		localctx.(*List_expressionContext).l.Add(e.GetP())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlock_list_expressionContext is an interface to support dynamic dispatch.
type IBlock_list_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	Expressions() IExpressionsContext
	TK_COMA() antlr.TerminalNode

	// IsBlock_list_expressionContext differentiates from other interfaces.
	IsBlock_list_expressionContext()
}

type Block_list_expressionContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	_expressions IExpressionsContext
}

func NewEmptyBlock_list_expressionContext() *Block_list_expressionContext {
	var p = new(Block_list_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_list_expression
	return p
}

func InitEmptyBlock_list_expressionContext(p *Block_list_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_block_list_expression
}

func (*Block_list_expressionContext) IsBlock_list_expressionContext() {}

func NewBlock_list_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_list_expressionContext {
	var p = new(Block_list_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_block_list_expression

	return p
}

func (s *Block_list_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_list_expressionContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Block_list_expressionContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Block_list_expressionContext) GetP() interfaces.Expression { return s.p }

func (s *Block_list_expressionContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Block_list_expressionContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Block_list_expressionContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_COMA, 0)
}

func (s *Block_list_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_list_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_list_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterBlock_list_expression(s)
	}
}

func (s *Block_list_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitBlock_list_expression(s)
	}
}

func (p *Swiftgrammar) Block_list_expression() (localctx IBlock_list_expressionContext) {
	localctx = NewBlock_list_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SwiftgrammarRULE_block_list_expression)
	p.SetState(672)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(665)

			var _x = p.Expressions()

			localctx.(*Block_list_expressionContext)._expressions = _x
		}
		{
			p.SetState(666)
			p.Match(SwiftgrammarTK_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Block_list_expressionContext).p = instruction.NewListExpre(localctx.(*Block_list_expressionContext).Get_expressions().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(669)

			var _x = p.Expressions()

			localctx.(*Block_list_expressionContext)._expressions = _x
		}
		localctx.(*Block_list_expressionContext).p = instruction.NewListExpre(localctx.(*Block_list_expressionContext).Get_expressions().GetP())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expre_log returns the _expre_log rule contexts.
	Get_expre_log() IExpre_logContext

	// Get_expre_arit returns the _expre_arit rule contexts.
	Get_expre_arit() IExpre_aritContext

	// Get_expre_rel returns the _expre_rel rule contexts.
	Get_expre_rel() IExpre_relContext

	// Set_expre_log sets the _expre_log rule contexts.
	Set_expre_log(IExpre_logContext)

	// Set_expre_arit sets the _expre_arit rule contexts.
	Set_expre_arit(IExpre_aritContext)

	// Set_expre_rel sets the _expre_rel rule contexts.
	Set_expre_rel(IExpre_relContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	Expre_log() IExpre_logContext
	Expre_arit() IExpre_aritContext
	Expre_rel() IExpre_relContext

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expression
	_expre_log  IExpre_logContext
	_expre_arit IExpre_aritContext
	_expre_rel  IExpre_relContext
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expressions
	return p
}

func InitEmptyExpressionsContext(p *ExpressionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expressions
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) Get_expre_log() IExpre_logContext { return s._expre_log }

func (s *ExpressionsContext) Get_expre_arit() IExpre_aritContext { return s._expre_arit }

func (s *ExpressionsContext) Get_expre_rel() IExpre_relContext { return s._expre_rel }

func (s *ExpressionsContext) Set_expre_log(v IExpre_logContext) { s._expre_log = v }

func (s *ExpressionsContext) Set_expre_arit(v IExpre_aritContext) { s._expre_arit = v }

func (s *ExpressionsContext) Set_expre_rel(v IExpre_relContext) { s._expre_rel = v }

func (s *ExpressionsContext) GetP() interfaces.Expression { return s.p }

func (s *ExpressionsContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ExpressionsContext) Expre_log() IExpre_logContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpre_logContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpre_logContext)
}

func (s *ExpressionsContext) Expre_arit() IExpre_aritContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpre_aritContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *ExpressionsContext) Expre_rel() IExpre_relContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpre_relContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpre_relContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *Swiftgrammar) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SwiftgrammarRULE_expressions)
	p.SetState(683)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(674)

			var _x = p.expre_log(0)

			localctx.(*ExpressionsContext)._expre_log = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_log().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(677)

			var _x = p.expre_arit(0)

			localctx.(*ExpressionsContext)._expre_arit = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_arit().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(680)

			var _x = p.expre_rel(0)

			localctx.(*ExpressionsContext)._expre_rel = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_rel().GetP()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpre_logContext is an interface to support dynamic dispatch.
type IExpre_logContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpre_logContext

	// Get_expre_rel returns the _expre_rel rule contexts.
	Get_expre_rel() IExpre_relContext

	// GetRight returns the right rule contexts.
	GetRight() IExpre_logContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpre_logContext)

	// Set_expre_rel sets the _expre_rel rule contexts.
	Set_expre_rel(IExpre_relContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpre_logContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	TK_NOT() antlr.TerminalNode
	AllExpre_log() []IExpre_logContext
	Expre_log(i int) IExpre_logContext
	Expre_rel() IExpre_relContext
	TK_AND() antlr.TerminalNode
	TK_OR() antlr.TerminalNode

	// IsExpre_logContext differentiates from other interfaces.
	IsExpre_logContext()
}

type Expre_logContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expression
	left       IExpre_logContext
	op         antlr.Token
	_expre_rel IExpre_relContext
	right      IExpre_logContext
}

func NewEmptyExpre_logContext() *Expre_logContext {
	var p = new(Expre_logContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expre_log
	return p
}

func InitEmptyExpre_logContext(p *Expre_logContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expre_log
}

func (*Expre_logContext) IsExpre_logContext() {}

func NewExpre_logContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_logContext {
	var p = new(Expre_logContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_expre_log

	return p
}

func (s *Expre_logContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_logContext) GetOp() antlr.Token { return s.op }

func (s *Expre_logContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_logContext) GetLeft() IExpre_logContext { return s.left }

func (s *Expre_logContext) Get_expre_rel() IExpre_relContext { return s._expre_rel }

func (s *Expre_logContext) GetRight() IExpre_logContext { return s.right }

func (s *Expre_logContext) SetLeft(v IExpre_logContext) { s.left = v }

func (s *Expre_logContext) Set_expre_rel(v IExpre_relContext) { s._expre_rel = v }

func (s *Expre_logContext) SetRight(v IExpre_logContext) { s.right = v }

func (s *Expre_logContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_logContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_logContext) TK_NOT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_NOT, 0)
}

func (s *Expre_logContext) AllExpre_log() []IExpre_logContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpre_logContext); ok {
			len++
		}
	}

	tst := make([]IExpre_logContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpre_logContext); ok {
			tst[i] = t.(IExpre_logContext)
			i++
		}
	}

	return tst
}

func (s *Expre_logContext) Expre_log(i int) IExpre_logContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpre_logContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpre_logContext)
}

func (s *Expre_logContext) Expre_rel() IExpre_relContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpre_relContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpre_relContext)
}

func (s *Expre_logContext) TK_AND() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_AND, 0)
}

func (s *Expre_logContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_OR, 0)
}

func (s *Expre_logContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_logContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_logContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterExpre_log(s)
	}
}

func (s *Expre_logContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitExpre_log(s)
	}
}

func (p *Swiftgrammar) Expre_log() (localctx IExpre_logContext) {
	return p.expre_log(0)
}

func (p *Swiftgrammar) expre_log(_p int) (localctx IExpre_logContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpre_logContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_logContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 106
	p.EnterRecursionRule(localctx, 106, SwiftgrammarRULE_expre_log, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(693)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammarTK_NOT:
		{
			p.SetState(686)

			var _m = p.Match(SwiftgrammarTK_NOT)

			localctx.(*Expre_logContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(687)

			var _x = p.expre_log(3)

			localctx.(*Expre_logContext).left = _x
		}
		localctx.(*Expre_logContext).p = expression.NewOperacion(localctx.(*Expre_logContext).GetLeft().GetP(), (func() string {
			if localctx.(*Expre_logContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Expre_logContext).GetOp().GetText()
			}
		}()), nil, true, (func() int {
			if localctx.(*Expre_logContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Expre_logContext).GetOp().GetLine()
			}
		}()), localctx.(*Expre_logContext).GetOp().GetColumn())

	case SwiftgrammarIF, SwiftgrammarWHILE, SwiftgrammarSWITCH, SwiftgrammarNUMBER, SwiftgrammarDOUBLE, SwiftgrammarCHAR, SwiftgrammarSTRING, SwiftgrammarBOOLEAN, SwiftgrammarID, SwiftgrammarTK_MENOS, SwiftgrammarTK_PARA:
		{
			p.SetState(690)

			var _x = p.expre_rel(0)

			localctx.(*Expre_logContext)._expre_rel = _x
		}
		localctx.(*Expre_logContext).p = localctx.(*Expre_logContext).Get_expre_rel().GetP()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(702)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpre_logContext(p, _parentctx, _parentState)
			localctx.(*Expre_logContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftgrammarRULE_expre_log)
			p.SetState(695)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(696)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expre_logContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SwiftgrammarTK_AND || _la == SwiftgrammarTK_OR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expre_logContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(697)

				var _x = p.expre_log(3)

				localctx.(*Expre_logContext).right = _x
			}
			localctx.(*Expre_logContext).p = expression.NewOperacion(localctx.(*Expre_logContext).GetLeft().GetP(), (func() string {
				if localctx.(*Expre_logContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_logContext).GetOp().GetText()
				}
			}()), localctx.(*Expre_logContext).GetRight().GetP(), false, (func() int {
				if localctx.(*Expre_logContext).GetOp() == nil {
					return 0
				} else {
					return localctx.(*Expre_logContext).GetOp().GetLine()
				}
			}()), localctx.(*Expre_logContext).GetOp().GetColumn())

		}
		p.SetState(704)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpre_relContext is an interface to support dynamic dispatch.
type IExpre_relContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpre_relContext

	// Get_expre_arit returns the _expre_arit rule contexts.
	Get_expre_arit() IExpre_aritContext

	// GetRight returns the right rule contexts.
	GetRight() IExpre_relContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpre_relContext)

	// Set_expre_arit sets the _expre_arit rule contexts.
	Set_expre_arit(IExpre_aritContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpre_relContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	Expre_arit() IExpre_aritContext
	AllExpre_rel() []IExpre_relContext
	Expre_rel(i int) IExpre_relContext
	TK_MENOR() antlr.TerminalNode
	TK_MENORIGUAL() antlr.TerminalNode
	TK_MAYORIGUAL() antlr.TerminalNode
	TK_MAYOR() antlr.TerminalNode
	TK_DIFIGUAL() antlr.TerminalNode
	TK_IGUALIGUAL() antlr.TerminalNode

	// IsExpre_relContext differentiates from other interfaces.
	IsExpre_relContext()
}

type Expre_relContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expression
	left        IExpre_relContext
	_expre_arit IExpre_aritContext
	op          antlr.Token
	right       IExpre_relContext
}

func NewEmptyExpre_relContext() *Expre_relContext {
	var p = new(Expre_relContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expre_rel
	return p
}

func InitEmptyExpre_relContext(p *Expre_relContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expre_rel
}

func (*Expre_relContext) IsExpre_relContext() {}

func NewExpre_relContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_relContext {
	var p = new(Expre_relContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_expre_rel

	return p
}

func (s *Expre_relContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_relContext) GetOp() antlr.Token { return s.op }

func (s *Expre_relContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_relContext) GetLeft() IExpre_relContext { return s.left }

func (s *Expre_relContext) Get_expre_arit() IExpre_aritContext { return s._expre_arit }

func (s *Expre_relContext) GetRight() IExpre_relContext { return s.right }

func (s *Expre_relContext) SetLeft(v IExpre_relContext) { s.left = v }

func (s *Expre_relContext) Set_expre_arit(v IExpre_aritContext) { s._expre_arit = v }

func (s *Expre_relContext) SetRight(v IExpre_relContext) { s.right = v }

func (s *Expre_relContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_relContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_relContext) Expre_arit() IExpre_aritContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpre_aritContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *Expre_relContext) AllExpre_rel() []IExpre_relContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpre_relContext); ok {
			len++
		}
	}

	tst := make([]IExpre_relContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpre_relContext); ok {
			tst[i] = t.(IExpre_relContext)
			i++
		}
	}

	return tst
}

func (s *Expre_relContext) Expre_rel(i int) IExpre_relContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpre_relContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpre_relContext)
}

func (s *Expre_relContext) TK_MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_MENOR, 0)
}

func (s *Expre_relContext) TK_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_MENORIGUAL, 0)
}

func (s *Expre_relContext) TK_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_MAYORIGUAL, 0)
}

func (s *Expre_relContext) TK_MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_MAYOR, 0)
}

func (s *Expre_relContext) TK_DIFIGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DIFIGUAL, 0)
}

func (s *Expre_relContext) TK_IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_IGUALIGUAL, 0)
}

func (s *Expre_relContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_relContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_relContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterExpre_rel(s)
	}
}

func (s *Expre_relContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitExpre_rel(s)
	}
}

func (p *Swiftgrammar) Expre_rel() (localctx IExpre_relContext) {
	return p.expre_rel(0)
}

func (p *Swiftgrammar) expre_rel(_p int) (localctx IExpre_relContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpre_relContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_relContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 108
	p.EnterRecursionRule(localctx, 108, SwiftgrammarRULE_expre_rel, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(706)

		var _x = p.expre_arit(0)

		localctx.(*Expre_relContext)._expre_arit = _x
	}
	localctx.(*Expre_relContext).p = localctx.(*Expre_relContext).Get_expre_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(716)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpre_relContext(p, _parentctx, _parentState)
			localctx.(*Expre_relContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftgrammarRULE_expre_rel)
			p.SetState(709)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(710)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expre_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16698832846848) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expre_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(711)

				var _x = p.expre_rel(3)

				localctx.(*Expre_relContext).right = _x
			}
			localctx.(*Expre_relContext).p = expression.NewOperacion(localctx.(*Expre_relContext).GetLeft().GetP(), (func() string {
				if localctx.(*Expre_relContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expre_relContext).GetOp().GetText()
				}
			}()), localctx.(*Expre_relContext).GetRight().GetP(), false, (func() int {
				if localctx.(*Expre_relContext).GetOp() == nil {
					return 0
				} else {
					return localctx.(*Expre_relContext).GetOp().GetLine()
				}
			}()), localctx.(*Expre_relContext).GetOp().GetColumn())

		}
		p.SetState(718)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpre_aritContext is an interface to support dynamic dispatch.
type IExpre_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpre_aritContext

	// Get_expre_valor returns the _expre_valor rule contexts.
	Get_expre_valor() IExpre_valorContext

	// Get_expressions returns the _expressions rule contexts.
	Get_expressions() IExpressionsContext

	// GetRight returns the right rule contexts.
	GetRight() IExpre_aritContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpre_aritContext)

	// Set_expre_valor sets the _expre_valor rule contexts.
	Set_expre_valor(IExpre_valorContext)

	// Set_expressions sets the _expressions rule contexts.
	Set_expressions(IExpressionsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpre_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	TK_MENOS() antlr.TerminalNode
	AllExpre_arit() []IExpre_aritContext
	Expre_arit(i int) IExpre_aritContext
	Expre_valor() IExpre_valorContext
	TK_PARA() antlr.TerminalNode
	Expressions() IExpressionsContext
	TK_PARC() antlr.TerminalNode
	TK_MULT() antlr.TerminalNode
	TK_DIV() antlr.TerminalNode
	TK_MODULO() antlr.TerminalNode
	TK_MAS() antlr.TerminalNode

	// IsExpre_aritContext differentiates from other interfaces.
	IsExpre_aritContext()
}

type Expre_aritContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            interfaces.Expression
	left         IExpre_aritContext
	op           antlr.Token
	_expre_valor IExpre_valorContext
	_expressions IExpressionsContext
	right        IExpre_aritContext
}

func NewEmptyExpre_aritContext() *Expre_aritContext {
	var p = new(Expre_aritContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expre_arit
	return p
}

func InitEmptyExpre_aritContext(p *Expre_aritContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expre_arit
}

func (*Expre_aritContext) IsExpre_aritContext() {}

func NewExpre_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_aritContext {
	var p = new(Expre_aritContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_expre_arit

	return p
}

func (s *Expre_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expre_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expre_aritContext) GetLeft() IExpre_aritContext { return s.left }

func (s *Expre_aritContext) Get_expre_valor() IExpre_valorContext { return s._expre_valor }

func (s *Expre_aritContext) Get_expressions() IExpressionsContext { return s._expressions }

func (s *Expre_aritContext) GetRight() IExpre_aritContext { return s.right }

func (s *Expre_aritContext) SetLeft(v IExpre_aritContext) { s.left = v }

func (s *Expre_aritContext) Set_expre_valor(v IExpre_valorContext) { s._expre_valor = v }

func (s *Expre_aritContext) Set_expressions(v IExpressionsContext) { s._expressions = v }

func (s *Expre_aritContext) SetRight(v IExpre_aritContext) { s.right = v }

func (s *Expre_aritContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_aritContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_aritContext) TK_MENOS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_MENOS, 0)
}

func (s *Expre_aritContext) AllExpre_arit() []IExpre_aritContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpre_aritContext); ok {
			len++
		}
	}

	tst := make([]IExpre_aritContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpre_aritContext); ok {
			tst[i] = t.(IExpre_aritContext)
			i++
		}
	}

	return tst
}

func (s *Expre_aritContext) Expre_arit(i int) IExpre_aritContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpre_aritContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpre_aritContext)
}

func (s *Expre_aritContext) Expre_valor() IExpre_valorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpre_valorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpre_valorContext)
}

func (s *Expre_aritContext) TK_PARA() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARA, 0)
}

func (s *Expre_aritContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Expre_aritContext) TK_PARC() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_PARC, 0)
}

func (s *Expre_aritContext) TK_MULT() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_MULT, 0)
}

func (s *Expre_aritContext) TK_DIV() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_DIV, 0)
}

func (s *Expre_aritContext) TK_MODULO() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_MODULO, 0)
}

func (s *Expre_aritContext) TK_MAS() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarTK_MAS, 0)
}

func (s *Expre_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterExpre_arit(s)
	}
}

func (s *Expre_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitExpre_arit(s)
	}
}

func (p *Swiftgrammar) Expre_arit() (localctx IExpre_aritContext) {
	return p.expre_arit(0)
}

func (p *Swiftgrammar) expre_arit(_p int) (localctx IExpre_aritContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpre_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpre_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 110
	p.EnterRecursionRule(localctx, 110, SwiftgrammarRULE_expre_arit, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(732)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammarTK_MENOS:
		{
			p.SetState(720)

			var _m = p.Match(SwiftgrammarTK_MENOS)

			localctx.(*Expre_aritContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(721)

			var _x = p.expre_arit(5)

			localctx.(*Expre_aritContext).left = _x
		}
		localctx.(*Expre_aritContext).p = expression.NewOperacion(localctx.(*Expre_aritContext).GetLeft().GetP(), (func() string {
			if localctx.(*Expre_aritContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*Expre_aritContext).GetOp().GetText()
			}
		}()), nil, true, (func() int {
			if localctx.(*Expre_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Expre_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Expre_aritContext).GetOp().GetColumn())

	case SwiftgrammarIF, SwiftgrammarWHILE, SwiftgrammarSWITCH, SwiftgrammarNUMBER, SwiftgrammarDOUBLE, SwiftgrammarCHAR, SwiftgrammarSTRING, SwiftgrammarBOOLEAN, SwiftgrammarID:
		{
			p.SetState(724)

			var _x = p.Expre_valor()

			localctx.(*Expre_aritContext)._expre_valor = _x
		}
		localctx.(*Expre_aritContext).p = localctx.(*Expre_aritContext).Get_expre_valor().GetP()

	case SwiftgrammarTK_PARA:
		{
			p.SetState(727)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(728)

			var _x = p.Expressions()

			localctx.(*Expre_aritContext)._expressions = _x
		}
		{
			p.SetState(729)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Expre_aritContext).p = localctx.(*Expre_aritContext).Get_expressions().GetP()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(746)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(744)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpre_aritContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammarRULE_expre_arit)
				p.SetState(734)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(735)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expre_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&123145302310912) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expre_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(736)

					var _x = p.expre_arit(5)

					localctx.(*Expre_aritContext).right = _x
				}
				localctx.(*Expre_aritContext).p = expression.NewOperacion(localctx.(*Expre_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expre_aritContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expre_aritContext).GetOp().GetColumn())

			case 2:
				localctx = NewExpre_aritContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammarRULE_expre_arit)
				p.SetState(739)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(740)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expre_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftgrammarTK_MAS || _la == SwiftgrammarTK_MENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expre_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(741)

					var _x = p.expre_arit(4)

					localctx.(*Expre_aritContext).right = _x
				}
				localctx.(*Expre_aritContext).p = expression.NewOperacion(localctx.(*Expre_aritContext).GetLeft().GetP(), (func() string {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expre_aritContext).GetRight().GetP(), false, (func() int {
					if localctx.(*Expre_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expre_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expre_aritContext).GetOp().GetColumn())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(748)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpre_valorContext is an interface to support dynamic dispatch.
type IExpre_valorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	Primitivo() IPrimitivoContext

	// IsExpre_valorContext differentiates from other interfaces.
	IsExpre_valorContext()
}

type Expre_valorContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expression
	_primitivo IPrimitivoContext
}

func NewEmptyExpre_valorContext() *Expre_valorContext {
	var p = new(Expre_valorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expre_valor
	return p
}

func InitEmptyExpre_valorContext(p *Expre_valorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_expre_valor
}

func (*Expre_valorContext) IsExpre_valorContext() {}

func NewExpre_valorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expre_valorContext {
	var p = new(Expre_valorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_expre_valor

	return p
}

func (s *Expre_valorContext) GetParser() antlr.Parser { return s.parser }

func (s *Expre_valorContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Expre_valorContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expre_valorContext) GetP() interfaces.Expression { return s.p }

func (s *Expre_valorContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expre_valorContext) Primitivo() IPrimitivoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimitivoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Expre_valorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expre_valorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expre_valorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterExpre_valor(s)
	}
}

func (s *Expre_valorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitExpre_valor(s)
	}
}

func (p *Swiftgrammar) Expre_valor() (localctx IExpre_valorContext) {
	localctx = NewExpre_valorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, SwiftgrammarRULE_expre_valor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(749)

		var _x = p.Primitivo()

		localctx.(*Expre_valorContext)._primitivo = _x
	}
	localctx.(*Expre_valorContext).p = localctx.(*Expre_valorContext).Get_primitivo().GetP()

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_DOUBLE returns the _DOUBLE token.
	Get_DOUBLE() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_BOOLEAN returns the _BOOLEAN token.
	Get_BOOLEAN() antlr.Token

	// Get_CHAR returns the _CHAR token.
	Get_CHAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_DOUBLE sets the _DOUBLE token.
	Set_DOUBLE(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_BOOLEAN sets the _BOOLEAN token.
	Set_BOOLEAN(antlr.Token)

	// Set_CHAR sets the _CHAR token.
	Set_CHAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instruccion_ternario returns the _instruccion_ternario rule contexts.
	Get_instruccion_ternario() IInstruccion_ternarioContext

	// Get_instruccion_while_true_ternario returns the _instruccion_while_true_ternario rule contexts.
	Get_instruccion_while_true_ternario() IInstruccion_while_true_ternarioContext

	// Set_instruccion_ternario sets the _instruccion_ternario rule contexts.
	Set_instruccion_ternario(IInstruccion_ternarioContext)

	// Set_instruccion_while_true_ternario sets the _instruccion_while_true_ternario rule contexts.
	Set_instruccion_while_true_ternario(IInstruccion_while_true_ternarioContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	NUMBER() antlr.TerminalNode
	DOUBLE() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	Instr_llamada_expre() IInstr_llamada_expreContext
	Instr_structs_identifier() IInstr_structs_identifierContext
	Instr_arrays_identifier() IInstr_arrays_identifierContext
	ID() antlr.TerminalNode
	Instruccion_ternario() IInstruccion_ternarioContext
	Instruccion_switch_ternario() IInstruccion_switch_ternarioContext
	Instruccion_while_true_ternario() IInstruccion_while_true_ternarioContext

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	antlr.BaseParserRuleContext
	parser                           antlr.Parser
	p                                interfaces.Expression
	_NUMBER                          antlr.Token
	_DOUBLE                          antlr.Token
	_STRING                          antlr.Token
	_BOOLEAN                         antlr.Token
	_CHAR                            antlr.Token
	_ID                              antlr.Token
	_instruccion_ternario            IInstruccion_ternarioContext
	_instruccion_while_true_ternario IInstruccion_while_true_ternarioContext
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_primitivo
	return p
}

func InitEmptyPrimitivoContext(p *PrimitivoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftgrammarRULE_primitivo
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftgrammarRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *PrimitivoContext) Get_DOUBLE() antlr.Token { return s._DOUBLE }

func (s *PrimitivoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitivoContext) Get_BOOLEAN() antlr.Token { return s._BOOLEAN }

func (s *PrimitivoContext) Get_CHAR() antlr.Token { return s._CHAR }

func (s *PrimitivoContext) Get_ID() antlr.Token { return s._ID }

func (s *PrimitivoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitivoContext) Set_DOUBLE(v antlr.Token) { s._DOUBLE = v }

func (s *PrimitivoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitivoContext) Set_BOOLEAN(v antlr.Token) { s._BOOLEAN = v }

func (s *PrimitivoContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *PrimitivoContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PrimitivoContext) Get_instruccion_ternario() IInstruccion_ternarioContext {
	return s._instruccion_ternario
}

func (s *PrimitivoContext) Get_instruccion_while_true_ternario() IInstruccion_while_true_ternarioContext {
	return s._instruccion_while_true_ternario
}

func (s *PrimitivoContext) Set_instruccion_ternario(v IInstruccion_ternarioContext) {
	s._instruccion_ternario = v
}

func (s *PrimitivoContext) Set_instruccion_while_true_ternario(v IInstruccion_while_true_ternarioContext) {
	s._instruccion_while_true_ternario = v
}

func (s *PrimitivoContext) GetP() interfaces.Expression { return s.p }

func (s *PrimitivoContext) SetP(v interfaces.Expression) { s.p = v }

func (s *PrimitivoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarNUMBER, 0)
}

func (s *PrimitivoContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarDOUBLE, 0)
}

func (s *PrimitivoContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarSTRING, 0)
}

func (s *PrimitivoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarBOOLEAN, 0)
}

func (s *PrimitivoContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarCHAR, 0)
}

func (s *PrimitivoContext) Instr_llamada_expre() IInstr_llamada_expreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_llamada_expreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_llamada_expreContext)
}

func (s *PrimitivoContext) Instr_structs_identifier() IInstr_structs_identifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_structs_identifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_structs_identifierContext)
}

func (s *PrimitivoContext) Instr_arrays_identifier() IInstr_arrays_identifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstr_arrays_identifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstr_arrays_identifierContext)
}

func (s *PrimitivoContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftgrammarID, 0)
}

func (s *PrimitivoContext) Instruccion_ternario() IInstruccion_ternarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_ternarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_ternarioContext)
}

func (s *PrimitivoContext) Instruccion_switch_ternario() IInstruccion_switch_ternarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_switch_ternarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_switch_ternarioContext)
}

func (s *PrimitivoContext) Instruccion_while_true_ternario() IInstruccion_while_true_ternarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruccion_while_true_ternarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruccion_while_true_ternarioContext)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftgrammarListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Swiftgrammar) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, SwiftgrammarRULE_primitivo)
	p.SetState(774)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(752)

			var _m = p.Match(SwiftgrammarNUMBER)

			localctx.(*PrimitivoContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}

		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(754)

			var _m = p.Match(SwiftgrammarDOUBLE)

			localctx.(*PrimitivoContext)._DOUBLE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_DOUBLE() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_DOUBLE().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(756)

			var _m = p.Match(SwiftgrammarSTRING)

			localctx.(*PrimitivoContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(str, interfaces.STRING, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_STRING().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(758)

			var _m = p.Match(SwiftgrammarBOOLEAN)

			localctx.(*PrimitivoContext)._BOOLEAN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		exp, _ := strconv.ParseBool((func() string {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetText()
			}
		}()))
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(760)

			var _m = p.Match(SwiftgrammarCHAR)

			localctx.(*PrimitivoContext)._CHAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CHAR().GetText()
			}
		}())[1]
		localctx.(*PrimitivoContext).p = expression.NewPrimitivo(string(str), interfaces.CHAR, interfaces.NULL, (func() int {
			if localctx.(*PrimitivoContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_CHAR().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_CHAR().GetColumn())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(762)
			p.Instr_llamada_expre()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(763)
			p.Instr_structs_identifier()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(764)
			p.Instr_arrays_identifier()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(765)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*PrimitivoContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrimitivoContext).p = instruction.NewIdentifier((func() string {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetText()
			}
		}()), (func() int {
			if localctx.(*PrimitivoContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*PrimitivoContext).Get_ID().GetLine()
			}
		}()), localctx.(*PrimitivoContext).Get_ID().GetColumn())

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(767)

			var _x = p.Instruccion_ternario()

			localctx.(*PrimitivoContext)._instruccion_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instruccion_ternario().GetP()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(770)
			p.Instruccion_switch_ternario()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(771)

			var _x = p.Instruccion_while_true_ternario()

			localctx.(*PrimitivoContext)._instruccion_while_true_ternario = _x
		}
		localctx.(*PrimitivoContext).p = localctx.(*PrimitivoContext).Get_instruccion_while_true_ternario().GetP()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *Swiftgrammar) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 53:
		var t *Expre_logContext = nil
		if localctx != nil {
			t = localctx.(*Expre_logContext)
		}
		return p.Expre_log_Sempred(t, predIndex)

	case 54:
		var t *Expre_relContext = nil
		if localctx != nil {
			t = localctx.(*Expre_relContext)
		}
		return p.Expre_rel_Sempred(t, predIndex)

	case 55:
		var t *Expre_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expre_aritContext)
		}
		return p.Expre_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Swiftgrammar) Expre_log_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Swiftgrammar) Expre_rel_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Swiftgrammar) Expre_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
