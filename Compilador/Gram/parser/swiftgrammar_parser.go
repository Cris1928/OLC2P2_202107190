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
		4, 1, 62, 735, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		150, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 172,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 214, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 245, 8, 6, 1, 7, 5, 7, 248, 8, 7, 10, 7, 12, 7, 251, 9, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 277, 8, 8, 1, 9, 5, 9, 280, 8, 9, 10, 9, 12, 9, 283, 9, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 298, 8, 10, 1, 11, 4, 11, 301, 8, 11, 11, 11, 12, 11,
		302, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 316, 8, 12, 1, 13, 4, 13, 319, 8, 13, 11, 13, 12, 13,
		320, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 327, 8, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3,
		16, 342, 8, 16, 1, 17, 4, 17, 345, 8, 17, 11, 17, 12, 17, 346, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 3, 18, 362, 8, 18, 1, 19, 4, 19, 365, 8, 19, 11, 19, 12, 19, 366,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3,
		20, 379, 8, 20, 1, 21, 4, 21, 382, 8, 21, 11, 21, 12, 21, 383, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 390, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 403, 8, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 3, 25, 429, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 3, 28, 446,
		8, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 3, 31, 491, 8, 31, 1, 32, 4, 32, 494, 8, 32, 11, 32, 12,
		32, 495, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33,
		506, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3,
		34, 516, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		3, 35, 526, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 4,
		37, 535, 8, 37, 11, 37, 12, 37, 536, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 3, 38, 547, 8, 38, 1, 39, 1, 39, 1, 39, 1, 40, 4,
		40, 553, 8, 40, 11, 40, 12, 40, 554, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 577, 8, 42, 1, 43, 4, 43, 580, 8,
		43, 11, 43, 12, 43, 581, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 3, 44, 592, 8, 44, 1, 45, 1, 45, 1, 45, 1, 46, 4, 46, 598, 8, 46,
		11, 46, 12, 46, 599, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		3, 49, 620, 8, 49, 1, 50, 4, 50, 623, 8, 50, 11, 50, 12, 50, 624, 1, 50,
		1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 636, 8,
		51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52,
		647, 8, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 3,
		53, 657, 8, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 5, 53, 664, 8, 53, 10,
		53, 12, 53, 667, 9, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54,
		1, 54, 1, 54, 5, 54, 678, 8, 54, 10, 54, 12, 54, 681, 9, 54, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55,
		1, 55, 3, 55, 696, 8, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 5, 55, 708, 8, 55, 10, 55, 12, 55, 711, 9, 55,
		1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1,
		57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 3, 57,
		733, 8, 57, 1, 57, 0, 3, 106, 108, 110, 58, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
		88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 0, 4, 2,
		0, 55, 55, 57, 57, 2, 0, 36, 37, 40, 43, 1, 0, 44, 46, 1, 0, 47, 48, 759,
		0, 116, 1, 0, 0, 0, 2, 121, 1, 0, 0, 0, 4, 149, 1, 0, 0, 0, 6, 171, 1,
		0, 0, 0, 8, 213, 1, 0, 0, 0, 10, 215, 1, 0, 0, 0, 12, 244, 1, 0, 0, 0,
		14, 249, 1, 0, 0, 0, 16, 276, 1, 0, 0, 0, 18, 281, 1, 0, 0, 0, 20, 297,
		1, 0, 0, 0, 22, 300, 1, 0, 0, 0, 24, 315, 1, 0, 0, 0, 26, 318, 1, 0, 0,
		0, 28, 326, 1, 0, 0, 0, 30, 328, 1, 0, 0, 0, 32, 341, 1, 0, 0, 0, 34, 344,
		1, 0, 0, 0, 36, 361, 1, 0, 0, 0, 38, 364, 1, 0, 0, 0, 40, 378, 1, 0, 0,
		0, 42, 381, 1, 0, 0, 0, 44, 389, 1, 0, 0, 0, 46, 402, 1, 0, 0, 0, 48, 404,
		1, 0, 0, 0, 50, 428, 1, 0, 0, 0, 52, 430, 1, 0, 0, 0, 54, 436, 1, 0, 0,
		0, 56, 445, 1, 0, 0, 0, 58, 447, 1, 0, 0, 0, 60, 449, 1, 0, 0, 0, 62, 490,
		1, 0, 0, 0, 64, 493, 1, 0, 0, 0, 66, 505, 1, 0, 0, 0, 68, 515, 1, 0, 0,
		0, 70, 525, 1, 0, 0, 0, 72, 527, 1, 0, 0, 0, 74, 534, 1, 0, 0, 0, 76, 546,
		1, 0, 0, 0, 78, 548, 1, 0, 0, 0, 80, 552, 1, 0, 0, 0, 82, 556, 1, 0, 0,
		0, 84, 576, 1, 0, 0, 0, 86, 579, 1, 0, 0, 0, 88, 591, 1, 0, 0, 0, 90, 593,
		1, 0, 0, 0, 92, 597, 1, 0, 0, 0, 94, 601, 1, 0, 0, 0, 96, 604, 1, 0, 0,
		0, 98, 619, 1, 0, 0, 0, 100, 622, 1, 0, 0, 0, 102, 635, 1, 0, 0, 0, 104,
		646, 1, 0, 0, 0, 106, 656, 1, 0, 0, 0, 108, 668, 1, 0, 0, 0, 110, 695,
		1, 0, 0, 0, 112, 712, 1, 0, 0, 0, 114, 732, 1, 0, 0, 0, 116, 117, 3, 2,
		1, 0, 117, 118, 5, 0, 0, 1, 118, 119, 6, 0, -1, 0, 119, 1, 1, 0, 0, 0,
		120, 122, 3, 4, 2, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123,
		121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126,
		6, 1, -1, 0, 126, 3, 1, 0, 0, 0, 127, 128, 3, 6, 3, 0, 128, 129, 6, 2,
		-1, 0, 129, 150, 1, 0, 0, 0, 130, 150, 3, 72, 36, 0, 131, 132, 3, 8, 4,
		0, 132, 133, 6, 2, -1, 0, 133, 150, 1, 0, 0, 0, 134, 135, 3, 10, 5, 0,
		135, 136, 6, 2, -1, 0, 136, 150, 1, 0, 0, 0, 137, 150, 3, 96, 48, 0, 138,
		150, 3, 12, 6, 0, 139, 150, 3, 50, 25, 0, 140, 150, 3, 48, 24, 0, 141,
		150, 3, 52, 26, 0, 142, 150, 3, 20, 10, 0, 143, 150, 3, 56, 28, 0, 144,
		150, 3, 58, 29, 0, 145, 150, 3, 62, 31, 0, 146, 150, 3, 68, 34, 0, 147,
		150, 3, 60, 30, 0, 148, 150, 3, 84, 42, 0, 149, 127, 1, 0, 0, 0, 149, 130,
		1, 0, 0, 0, 149, 131, 1, 0, 0, 0, 149, 134, 1, 0, 0, 0, 149, 137, 1, 0,
		0, 0, 149, 138, 1, 0, 0, 0, 149, 139, 1, 0, 0, 0, 149, 140, 1, 0, 0, 0,
		149, 141, 1, 0, 0, 0, 149, 142, 1, 0, 0, 0, 149, 143, 1, 0, 0, 0, 149,
		144, 1, 0, 0, 0, 149, 145, 1, 0, 0, 0, 149, 146, 1, 0, 0, 0, 149, 147,
		1, 0, 0, 0, 149, 148, 1, 0, 0, 0, 150, 5, 1, 0, 0, 0, 151, 152, 5, 1, 0,
		0, 152, 153, 5, 49, 0, 0, 153, 154, 3, 114, 57, 0, 154, 155, 5, 50, 0,
		0, 155, 156, 6, 3, -1, 0, 156, 172, 1, 0, 0, 0, 157, 158, 5, 1, 0, 0, 158,
		159, 5, 49, 0, 0, 159, 160, 5, 27, 0, 0, 160, 161, 5, 33, 0, 0, 161, 162,
		3, 100, 50, 0, 162, 163, 5, 50, 0, 0, 163, 164, 6, 3, -1, 0, 164, 172,
		1, 0, 0, 0, 165, 166, 5, 1, 0, 0, 166, 167, 5, 49, 0, 0, 167, 168, 3, 104,
		52, 0, 168, 169, 5, 50, 0, 0, 169, 170, 6, 3, -1, 0, 170, 172, 1, 0, 0,
		0, 171, 151, 1, 0, 0, 0, 171, 157, 1, 0, 0, 0, 171, 165, 1, 0, 0, 0, 172,
		7, 1, 0, 0, 0, 173, 174, 5, 2, 0, 0, 174, 175, 5, 29, 0, 0, 175, 176, 5,
		35, 0, 0, 176, 177, 3, 104, 52, 0, 177, 178, 6, 4, -1, 0, 178, 214, 1,
		0, 0, 0, 179, 180, 5, 2, 0, 0, 180, 181, 5, 29, 0, 0, 181, 182, 5, 34,
		0, 0, 182, 183, 3, 98, 49, 0, 183, 184, 6, 4, -1, 0, 184, 214, 1, 0, 0,
		0, 185, 186, 5, 2, 0, 0, 186, 187, 5, 29, 0, 0, 187, 188, 5, 34, 0, 0,
		188, 189, 3, 98, 49, 0, 189, 190, 5, 35, 0, 0, 190, 191, 3, 104, 52, 0,
		191, 192, 6, 4, -1, 0, 192, 214, 1, 0, 0, 0, 193, 194, 5, 3, 0, 0, 194,
		195, 5, 29, 0, 0, 195, 196, 5, 35, 0, 0, 196, 197, 3, 104, 52, 0, 197,
		198, 6, 4, -1, 0, 198, 214, 1, 0, 0, 0, 199, 200, 5, 3, 0, 0, 200, 201,
		5, 29, 0, 0, 201, 202, 5, 34, 0, 0, 202, 203, 3, 98, 49, 0, 203, 204, 6,
		4, -1, 0, 204, 214, 1, 0, 0, 0, 205, 206, 5, 3, 0, 0, 206, 207, 5, 29,
		0, 0, 207, 208, 5, 34, 0, 0, 208, 209, 3, 98, 49, 0, 209, 210, 5, 35, 0,
		0, 210, 211, 3, 104, 52, 0, 211, 212, 6, 4, -1, 0, 212, 214, 1, 0, 0, 0,
		213, 173, 1, 0, 0, 0, 213, 179, 1, 0, 0, 0, 213, 185, 1, 0, 0, 0, 213,
		193, 1, 0, 0, 0, 213, 199, 1, 0, 0, 0, 213, 205, 1, 0, 0, 0, 214, 9, 1,
		0, 0, 0, 215, 216, 5, 29, 0, 0, 216, 217, 5, 35, 0, 0, 217, 218, 3, 104,
		52, 0, 218, 219, 6, 5, -1, 0, 219, 11, 1, 0, 0, 0, 220, 221, 5, 4, 0, 0,
		221, 222, 3, 104, 52, 0, 222, 223, 5, 51, 0, 0, 223, 224, 3, 2, 1, 0, 224,
		225, 5, 52, 0, 0, 225, 245, 1, 0, 0, 0, 226, 227, 5, 4, 0, 0, 227, 228,
		3, 104, 52, 0, 228, 229, 5, 51, 0, 0, 229, 230, 3, 2, 1, 0, 230, 231, 5,
		52, 0, 0, 231, 232, 5, 5, 0, 0, 232, 233, 5, 51, 0, 0, 233, 234, 3, 2,
		1, 0, 234, 235, 5, 52, 0, 0, 235, 245, 1, 0, 0, 0, 236, 237, 5, 4, 0, 0,
		237, 238, 3, 104, 52, 0, 238, 239, 5, 51, 0, 0, 239, 240, 3, 2, 1, 0, 240,
		241, 5, 52, 0, 0, 241, 242, 5, 5, 0, 0, 242, 243, 3, 14, 7, 0, 243, 245,
		1, 0, 0, 0, 244, 220, 1, 0, 0, 0, 244, 226, 1, 0, 0, 0, 244, 236, 1, 0,
		0, 0, 245, 13, 1, 0, 0, 0, 246, 248, 3, 12, 6, 0, 247, 246, 1, 0, 0, 0,
		248, 251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250,
		15, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 252, 253, 5, 4, 0, 0, 253, 254, 3,
		104, 52, 0, 254, 255, 5, 51, 0, 0, 255, 256, 3, 104, 52, 0, 256, 257, 5,
		52, 0, 0, 257, 277, 1, 0, 0, 0, 258, 259, 5, 4, 0, 0, 259, 260, 3, 104,
		52, 0, 260, 261, 5, 51, 0, 0, 261, 262, 3, 104, 52, 0, 262, 263, 5, 52,
		0, 0, 263, 264, 5, 5, 0, 0, 264, 265, 5, 51, 0, 0, 265, 266, 3, 104, 52,
		0, 266, 267, 5, 52, 0, 0, 267, 277, 1, 0, 0, 0, 268, 269, 5, 4, 0, 0, 269,
		270, 3, 104, 52, 0, 270, 271, 5, 51, 0, 0, 271, 272, 3, 104, 52, 0, 272,
		273, 5, 52, 0, 0, 273, 274, 5, 5, 0, 0, 274, 275, 3, 18, 9, 0, 275, 277,
		1, 0, 0, 0, 276, 252, 1, 0, 0, 0, 276, 258, 1, 0, 0, 0, 276, 268, 1, 0,
		0, 0, 277, 17, 1, 0, 0, 0, 278, 280, 3, 16, 8, 0, 279, 278, 1, 0, 0, 0,
		280, 283, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282,
		19, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 284, 285, 5, 8, 0, 0, 285, 286, 3,
		104, 52, 0, 286, 287, 5, 51, 0, 0, 287, 288, 3, 22, 11, 0, 288, 289, 3,
		34, 17, 0, 289, 290, 5, 52, 0, 0, 290, 298, 1, 0, 0, 0, 291, 292, 5, 8,
		0, 0, 292, 293, 3, 104, 52, 0, 293, 294, 5, 51, 0, 0, 294, 295, 3, 34,
		17, 0, 295, 296, 5, 52, 0, 0, 296, 298, 1, 0, 0, 0, 297, 284, 1, 0, 0,
		0, 297, 291, 1, 0, 0, 0, 298, 21, 1, 0, 0, 0, 299, 301, 3, 24, 12, 0, 300,
		299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303,
		1, 0, 0, 0, 303, 23, 1, 0, 0, 0, 304, 305, 3, 26, 13, 0, 305, 306, 5, 34,
		0, 0, 306, 307, 5, 51, 0, 0, 307, 308, 3, 2, 1, 0, 308, 309, 5, 52, 0,
		0, 309, 316, 1, 0, 0, 0, 310, 311, 3, 26, 13, 0, 311, 312, 5, 34, 0, 0,
		312, 313, 3, 30, 15, 0, 313, 314, 5, 33, 0, 0, 314, 316, 1, 0, 0, 0, 315,
		304, 1, 0, 0, 0, 315, 310, 1, 0, 0, 0, 316, 25, 1, 0, 0, 0, 317, 319, 3,
		28, 14, 0, 318, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 318, 1, 0,
		0, 0, 320, 321, 1, 0, 0, 0, 321, 27, 1, 0, 0, 0, 322, 323, 3, 104, 52,
		0, 323, 324, 5, 58, 0, 0, 324, 327, 1, 0, 0, 0, 325, 327, 3, 104, 52, 0,
		326, 322, 1, 0, 0, 0, 326, 325, 1, 0, 0, 0, 327, 29, 1, 0, 0, 0, 328, 329,
		3, 4, 2, 0, 329, 31, 1, 0, 0, 0, 330, 331, 5, 9, 0, 0, 331, 332, 5, 29,
		0, 0, 332, 333, 5, 34, 0, 0, 333, 334, 5, 51, 0, 0, 334, 335, 3, 2, 1,
		0, 335, 336, 5, 52, 0, 0, 336, 342, 1, 0, 0, 0, 337, 338, 5, 9, 0, 0, 338,
		339, 5, 29, 0, 0, 339, 340, 5, 34, 0, 0, 340, 342, 3, 30, 15, 0, 341, 330,
		1, 0, 0, 0, 341, 337, 1, 0, 0, 0, 342, 33, 1, 0, 0, 0, 343, 345, 3, 32,
		16, 0, 344, 343, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0,
		346, 347, 1, 0, 0, 0, 347, 35, 1, 0, 0, 0, 348, 349, 5, 8, 0, 0, 349, 350,
		3, 104, 52, 0, 350, 351, 5, 51, 0, 0, 351, 352, 3, 38, 19, 0, 352, 353,
		3, 46, 23, 0, 353, 354, 5, 52, 0, 0, 354, 362, 1, 0, 0, 0, 355, 356, 5,
		8, 0, 0, 356, 357, 3, 104, 52, 0, 357, 358, 5, 51, 0, 0, 358, 359, 3, 46,
		23, 0, 359, 360, 5, 52, 0, 0, 360, 362, 1, 0, 0, 0, 361, 348, 1, 0, 0,
		0, 361, 355, 1, 0, 0, 0, 362, 37, 1, 0, 0, 0, 363, 365, 3, 40, 20, 0, 364,
		363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366, 367,
		1, 0, 0, 0, 367, 39, 1, 0, 0, 0, 368, 369, 3, 42, 21, 0, 369, 370, 5, 34,
		0, 0, 370, 371, 3, 104, 52, 0, 371, 379, 1, 0, 0, 0, 372, 373, 3, 42, 21,
		0, 373, 374, 5, 34, 0, 0, 374, 375, 5, 51, 0, 0, 375, 376, 3, 104, 52,
		0, 376, 377, 5, 52, 0, 0, 377, 379, 1, 0, 0, 0, 378, 368, 1, 0, 0, 0, 378,
		372, 1, 0, 0, 0, 379, 41, 1, 0, 0, 0, 380, 382, 3, 44, 22, 0, 381, 380,
		1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 383, 384, 1, 0,
		0, 0, 384, 43, 1, 0, 0, 0, 385, 386, 3, 104, 52, 0, 386, 387, 5, 58, 0,
		0, 387, 390, 1, 0, 0, 0, 388, 390, 3, 104, 52, 0, 389, 385, 1, 0, 0, 0,
		389, 388, 1, 0, 0, 0, 390, 45, 1, 0, 0, 0, 391, 392, 5, 9, 0, 0, 392, 393,
		5, 29, 0, 0, 393, 394, 5, 34, 0, 0, 394, 403, 3, 104, 52, 0, 395, 396,
		5, 9, 0, 0, 396, 397, 5, 29, 0, 0, 397, 398, 5, 34, 0, 0, 398, 399, 5,
		51, 0, 0, 399, 400, 3, 104, 52, 0, 400, 401, 5, 52, 0, 0, 401, 403, 1,
		0, 0, 0, 402, 391, 1, 0, 0, 0, 402, 395, 1, 0, 0, 0, 403, 47, 1, 0, 0,
		0, 404, 405, 5, 7, 0, 0, 405, 406, 3, 104, 52, 0, 406, 407, 5, 51, 0, 0,
		407, 408, 3, 2, 1, 0, 408, 409, 5, 52, 0, 0, 409, 49, 1, 0, 0, 0, 410,
		411, 5, 6, 0, 0, 411, 412, 5, 29, 0, 0, 412, 413, 5, 11, 0, 0, 413, 414,
		3, 104, 52, 0, 414, 415, 5, 30, 0, 0, 415, 416, 3, 104, 52, 0, 416, 417,
		5, 51, 0, 0, 417, 418, 3, 2, 1, 0, 418, 419, 5, 52, 0, 0, 419, 429, 1,
		0, 0, 0, 420, 421, 5, 6, 0, 0, 421, 422, 5, 29, 0, 0, 422, 423, 5, 11,
		0, 0, 423, 424, 3, 104, 52, 0, 424, 425, 5, 51, 0, 0, 425, 426, 3, 2, 1,
		0, 426, 427, 5, 52, 0, 0, 427, 429, 1, 0, 0, 0, 428, 410, 1, 0, 0, 0, 428,
		420, 1, 0, 0, 0, 429, 51, 1, 0, 0, 0, 430, 431, 5, 7, 0, 0, 431, 432, 5,
		23, 0, 0, 432, 433, 5, 51, 0, 0, 433, 434, 3, 2, 1, 0, 434, 435, 5, 52,
		0, 0, 435, 53, 1, 0, 0, 0, 436, 437, 5, 7, 0, 0, 437, 438, 5, 23, 0, 0,
		438, 439, 5, 51, 0, 0, 439, 440, 3, 2, 1, 0, 440, 441, 5, 52, 0, 0, 441,
		55, 1, 0, 0, 0, 442, 446, 5, 12, 0, 0, 443, 444, 5, 12, 0, 0, 444, 446,
		3, 104, 52, 0, 445, 442, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 57, 1,
		0, 0, 0, 447, 448, 5, 13, 0, 0, 448, 59, 1, 0, 0, 0, 449, 450, 5, 14, 0,
		0, 450, 451, 3, 104, 52, 0, 451, 61, 1, 0, 0, 0, 452, 453, 5, 15, 0, 0,
		453, 454, 5, 29, 0, 0, 454, 455, 5, 49, 0, 0, 455, 456, 5, 50, 0, 0, 456,
		457, 5, 51, 0, 0, 457, 458, 3, 2, 1, 0, 458, 459, 5, 52, 0, 0, 459, 491,
		1, 0, 0, 0, 460, 461, 5, 15, 0, 0, 461, 462, 5, 29, 0, 0, 462, 463, 5,
		49, 0, 0, 463, 464, 5, 50, 0, 0, 464, 465, 5, 39, 0, 0, 465, 466, 3, 98,
		49, 0, 466, 467, 5, 51, 0, 0, 467, 468, 3, 2, 1, 0, 468, 469, 5, 52, 0,
		0, 469, 491, 1, 0, 0, 0, 470, 471, 5, 15, 0, 0, 471, 472, 5, 29, 0, 0,
		472, 473, 5, 49, 0, 0, 473, 474, 3, 64, 32, 0, 474, 475, 5, 50, 0, 0, 475,
		476, 5, 51, 0, 0, 476, 477, 3, 2, 1, 0, 477, 478, 5, 52, 0, 0, 478, 491,
		1, 0, 0, 0, 479, 480, 5, 15, 0, 0, 480, 481, 5, 29, 0, 0, 481, 482, 5,
		49, 0, 0, 482, 483, 3, 64, 32, 0, 483, 484, 5, 50, 0, 0, 484, 485, 5, 39,
		0, 0, 485, 486, 3, 98, 49, 0, 486, 487, 5, 51, 0, 0, 487, 488, 3, 2, 1,
		0, 488, 489, 5, 52, 0, 0, 489, 491, 1, 0, 0, 0, 490, 452, 1, 0, 0, 0, 490,
		460, 1, 0, 0, 0, 490, 470, 1, 0, 0, 0, 490, 479, 1, 0, 0, 0, 491, 63, 1,
		0, 0, 0, 492, 494, 3, 66, 33, 0, 493, 492, 1, 0, 0, 0, 494, 495, 1, 0,
		0, 0, 495, 493, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 65, 1, 0, 0, 0,
		497, 498, 5, 29, 0, 0, 498, 499, 5, 34, 0, 0, 499, 500, 3, 98, 49, 0, 500,
		501, 5, 33, 0, 0, 501, 506, 1, 0, 0, 0, 502, 503, 5, 29, 0, 0, 503, 504,
		5, 34, 0, 0, 504, 506, 3, 98, 49, 0, 505, 497, 1, 0, 0, 0, 505, 502, 1,
		0, 0, 0, 506, 67, 1, 0, 0, 0, 507, 508, 5, 29, 0, 0, 508, 509, 5, 49, 0,
		0, 509, 516, 5, 50, 0, 0, 510, 511, 5, 29, 0, 0, 511, 512, 5, 49, 0, 0,
		512, 513, 3, 100, 50, 0, 513, 514, 5, 50, 0, 0, 514, 516, 1, 0, 0, 0, 515,
		507, 1, 0, 0, 0, 515, 510, 1, 0, 0, 0, 516, 69, 1, 0, 0, 0, 517, 518, 5,
		29, 0, 0, 518, 519, 5, 49, 0, 0, 519, 526, 5, 50, 0, 0, 520, 521, 5, 29,
		0, 0, 521, 522, 5, 49, 0, 0, 522, 523, 3, 100, 50, 0, 523, 524, 5, 50,
		0, 0, 524, 526, 1, 0, 0, 0, 525, 517, 1, 0, 0, 0, 525, 520, 1, 0, 0, 0,
		526, 71, 1, 0, 0, 0, 527, 528, 5, 16, 0, 0, 528, 529, 5, 29, 0, 0, 529,
		530, 5, 51, 0, 0, 530, 531, 3, 74, 37, 0, 531, 532, 5, 52, 0, 0, 532, 73,
		1, 0, 0, 0, 533, 535, 3, 76, 38, 0, 534, 533, 1, 0, 0, 0, 535, 536, 1,
		0, 0, 0, 536, 534, 1, 0, 0, 0, 536, 537, 1, 0, 0, 0, 537, 75, 1, 0, 0,
		0, 538, 539, 5, 29, 0, 0, 539, 540, 5, 34, 0, 0, 540, 541, 3, 98, 49, 0,
		541, 542, 5, 33, 0, 0, 542, 547, 1, 0, 0, 0, 543, 544, 5, 29, 0, 0, 544,
		545, 5, 34, 0, 0, 545, 547, 3, 98, 49, 0, 546, 538, 1, 0, 0, 0, 546, 543,
		1, 0, 0, 0, 547, 77, 1, 0, 0, 0, 548, 549, 5, 29, 0, 0, 549, 550, 3, 80,
		40, 0, 550, 79, 1, 0, 0, 0, 551, 553, 3, 82, 41, 0, 552, 551, 1, 0, 0,
		0, 553, 554, 1, 0, 0, 0, 554, 552, 1, 0, 0, 0, 554, 555, 1, 0, 0, 0, 555,
		81, 1, 0, 0, 0, 556, 557, 5, 53, 0, 0, 557, 558, 3, 104, 52, 0, 558, 559,
		5, 54, 0, 0, 559, 83, 1, 0, 0, 0, 560, 561, 5, 2, 0, 0, 561, 562, 5, 29,
		0, 0, 562, 563, 5, 35, 0, 0, 563, 564, 5, 29, 0, 0, 564, 565, 5, 51, 0,
		0, 565, 566, 3, 86, 43, 0, 566, 567, 5, 52, 0, 0, 567, 577, 1, 0, 0, 0,
		568, 569, 5, 3, 0, 0, 569, 570, 5, 29, 0, 0, 570, 571, 5, 35, 0, 0, 571,
		572, 5, 29, 0, 0, 572, 573, 5, 51, 0, 0, 573, 574, 3, 86, 43, 0, 574, 575,
		5, 52, 0, 0, 575, 577, 1, 0, 0, 0, 576, 560, 1, 0, 0, 0, 576, 568, 1, 0,
		0, 0, 577, 85, 1, 0, 0, 0, 578, 580, 3, 88, 44, 0, 579, 578, 1, 0, 0, 0,
		580, 581, 1, 0, 0, 0, 581, 579, 1, 0, 0, 0, 581, 582, 1, 0, 0, 0, 582,
		87, 1, 0, 0, 0, 583, 584, 5, 29, 0, 0, 584, 585, 5, 34, 0, 0, 585, 592,
		3, 104, 52, 0, 586, 587, 5, 29, 0, 0, 587, 588, 5, 34, 0, 0, 588, 589,
		3, 104, 52, 0, 589, 590, 5, 33, 0, 0, 590, 592, 1, 0, 0, 0, 591, 583, 1,
		0, 0, 0, 591, 586, 1, 0, 0, 0, 592, 89, 1, 0, 0, 0, 593, 594, 5, 29, 0,
		0, 594, 595, 3, 92, 46, 0, 595, 91, 1, 0, 0, 0, 596, 598, 3, 94, 47, 0,
		597, 596, 1, 0, 0, 0, 598, 599, 1, 0, 0, 0, 599, 597, 1, 0, 0, 0, 599,
		600, 1, 0, 0, 0, 600, 93, 1, 0, 0, 0, 601, 602, 5, 31, 0, 0, 602, 603,
		5, 29, 0, 0, 603, 95, 1, 0, 0, 0, 604, 605, 5, 29, 0, 0, 605, 606, 3, 92,
		46, 0, 606, 607, 5, 35, 0, 0, 607, 608, 3, 104, 52, 0, 608, 97, 1, 0, 0,
		0, 609, 610, 5, 17, 0, 0, 610, 620, 6, 49, -1, 0, 611, 612, 5, 18, 0, 0,
		612, 620, 6, 49, -1, 0, 613, 614, 5, 22, 0, 0, 614, 620, 6, 49, -1, 0,
		615, 616, 5, 20, 0, 0, 616, 620, 6, 49, -1, 0, 617, 618, 5, 21, 0, 0, 618,
		620, 6, 49, -1, 0, 619, 609, 1, 0, 0, 0, 619, 611, 1, 0, 0, 0, 619, 613,
		1, 0, 0, 0, 619, 615, 1, 0, 0, 0, 619, 617, 1, 0, 0, 0, 620, 99, 1, 0,
		0, 0, 621, 623, 3, 102, 51, 0, 622, 621, 1, 0, 0, 0, 623, 624, 1, 0, 0,
		0, 624, 622, 1, 0, 0, 0, 624, 625, 1, 0, 0, 0, 625, 626, 1, 0, 0, 0, 626,
		627, 6, 50, -1, 0, 627, 101, 1, 0, 0, 0, 628, 629, 3, 104, 52, 0, 629,
		630, 5, 33, 0, 0, 630, 631, 6, 51, -1, 0, 631, 636, 1, 0, 0, 0, 632, 633,
		3, 104, 52, 0, 633, 634, 6, 51, -1, 0, 634, 636, 1, 0, 0, 0, 635, 628,
		1, 0, 0, 0, 635, 632, 1, 0, 0, 0, 636, 103, 1, 0, 0, 0, 637, 638, 3, 106,
		53, 0, 638, 639, 6, 52, -1, 0, 639, 647, 1, 0, 0, 0, 640, 641, 3, 110,
		55, 0, 641, 642, 6, 52, -1, 0, 642, 647, 1, 0, 0, 0, 643, 644, 3, 108,
		54, 0, 644, 645, 6, 52, -1, 0, 645, 647, 1, 0, 0, 0, 646, 637, 1, 0, 0,
		0, 646, 640, 1, 0, 0, 0, 646, 643, 1, 0, 0, 0, 647, 105, 1, 0, 0, 0, 648,
		649, 6, 53, -1, 0, 649, 650, 5, 59, 0, 0, 650, 651, 3, 106, 53, 3, 651,
		652, 6, 53, -1, 0, 652, 657, 1, 0, 0, 0, 653, 654, 3, 108, 54, 0, 654,
		655, 6, 53, -1, 0, 655, 657, 1, 0, 0, 0, 656, 648, 1, 0, 0, 0, 656, 653,
		1, 0, 0, 0, 657, 665, 1, 0, 0, 0, 658, 659, 10, 2, 0, 0, 659, 660, 7, 0,
		0, 0, 660, 661, 3, 106, 53, 3, 661, 662, 6, 53, -1, 0, 662, 664, 1, 0,
		0, 0, 663, 658, 1, 0, 0, 0, 664, 667, 1, 0, 0, 0, 665, 663, 1, 0, 0, 0,
		665, 666, 1, 0, 0, 0, 666, 107, 1, 0, 0, 0, 667, 665, 1, 0, 0, 0, 668,
		669, 6, 54, -1, 0, 669, 670, 3, 110, 55, 0, 670, 671, 6, 54, -1, 0, 671,
		679, 1, 0, 0, 0, 672, 673, 10, 2, 0, 0, 673, 674, 7, 1, 0, 0, 674, 675,
		3, 108, 54, 3, 675, 676, 6, 54, -1, 0, 676, 678, 1, 0, 0, 0, 677, 672,
		1, 0, 0, 0, 678, 681, 1, 0, 0, 0, 679, 677, 1, 0, 0, 0, 679, 680, 1, 0,
		0, 0, 680, 109, 1, 0, 0, 0, 681, 679, 1, 0, 0, 0, 682, 683, 6, 55, -1,
		0, 683, 684, 5, 48, 0, 0, 684, 685, 3, 110, 55, 5, 685, 686, 6, 55, -1,
		0, 686, 696, 1, 0, 0, 0, 687, 688, 3, 112, 56, 0, 688, 689, 6, 55, -1,
		0, 689, 696, 1, 0, 0, 0, 690, 691, 5, 49, 0, 0, 691, 692, 3, 104, 52, 0,
		692, 693, 5, 50, 0, 0, 693, 694, 6, 55, -1, 0, 694, 696, 1, 0, 0, 0, 695,
		682, 1, 0, 0, 0, 695, 687, 1, 0, 0, 0, 695, 690, 1, 0, 0, 0, 696, 709,
		1, 0, 0, 0, 697, 698, 10, 4, 0, 0, 698, 699, 7, 2, 0, 0, 699, 700, 3, 110,
		55, 5, 700, 701, 6, 55, -1, 0, 701, 708, 1, 0, 0, 0, 702, 703, 10, 3, 0,
		0, 703, 704, 7, 3, 0, 0, 704, 705, 3, 110, 55, 4, 705, 706, 6, 55, -1,
		0, 706, 708, 1, 0, 0, 0, 707, 697, 1, 0, 0, 0, 707, 702, 1, 0, 0, 0, 708,
		711, 1, 0, 0, 0, 709, 707, 1, 0, 0, 0, 709, 710, 1, 0, 0, 0, 710, 111,
		1, 0, 0, 0, 711, 709, 1, 0, 0, 0, 712, 713, 3, 114, 57, 0, 713, 714, 6,
		56, -1, 0, 714, 113, 1, 0, 0, 0, 715, 716, 5, 24, 0, 0, 716, 733, 6, 57,
		-1, 0, 717, 718, 5, 25, 0, 0, 718, 733, 6, 57, -1, 0, 719, 720, 5, 27,
		0, 0, 720, 733, 6, 57, -1, 0, 721, 722, 5, 28, 0, 0, 722, 733, 6, 57, -1,
		0, 723, 724, 5, 26, 0, 0, 724, 733, 6, 57, -1, 0, 725, 733, 3, 70, 35,
		0, 726, 733, 3, 90, 45, 0, 727, 733, 3, 78, 39, 0, 728, 729, 5, 29, 0,
		0, 729, 733, 6, 57, -1, 0, 730, 733, 3, 16, 8, 0, 731, 733, 3, 36, 18,
		0, 732, 715, 1, 0, 0, 0, 732, 717, 1, 0, 0, 0, 732, 719, 1, 0, 0, 0, 732,
		721, 1, 0, 0, 0, 732, 723, 1, 0, 0, 0, 732, 725, 1, 0, 0, 0, 732, 726,
		1, 0, 0, 0, 732, 727, 1, 0, 0, 0, 732, 728, 1, 0, 0, 0, 732, 730, 1, 0,
		0, 0, 732, 731, 1, 0, 0, 0, 733, 115, 1, 0, 0, 0, 46, 123, 149, 171, 213,
		244, 249, 276, 281, 297, 302, 315, 320, 326, 341, 346, 361, 366, 378, 383,
		389, 402, 428, 445, 490, 495, 505, 515, 525, 536, 546, 554, 576, 581, 591,
		599, 619, 624, 635, 646, 656, 665, 679, 695, 707, 709, 732,
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
	//	goto errorExit
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
		//	goto errorExit
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

	// Set_instruccion_println sets the _instruccion_println rule contexts.
	Set_instruccion_println(IInstruccion_printlnContext)

	// Set_instruccion_declaracion sets the _instruccion_declaracion rule contexts.
	Set_instruccion_declaracion(IInstruccion_declaracionContext)

	// Set_instruccion_asignacion sets the _instruccion_asignacion rule contexts.
	Set_instruccion_asignacion(IInstruccion_asignacionContext)

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

func (s *InstruccionContext) Set_instruccion_println(v IInstruccion_printlnContext) {
	s._instruccion_println = v
}

func (s *InstruccionContext) Set_instruccion_declaracion(v IInstruccion_declaracionContext) {
	s._instruccion_declaracion = v
}

func (s *InstruccionContext) Set_instruccion_asignacion(v IInstruccion_asignacionContext) {
	s._instruccion_asignacion = v
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
	p.SetState(149)
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
			p.Instruccion_if()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(139)
			p.Instruccion_for_in()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(140)
			p.Instruccion_while()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(141)
			p.Instruccion_while_true()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(142)
			p.Instruccion_switch()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(143)
			p.Instruccion_break()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(144)
			p.Instruccion_continue()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(145)
			p.Instruccion_func()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(146)
			p.Instruccion_llamada()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(147)
			p.Instruccion_return()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(148)
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
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)

			var _m = p.Match(SwiftgrammarPRINT)

			localctx.(*Instruccion_printlnContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)

			var _x = p.Primitivo()

			localctx.(*Instruccion_printlnContext)._primitivo = _x
		}
		{
			p.SetState(154)
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
			p.SetState(157)

			var _m = p.Match(SwiftgrammarPRINT)

			localctx.(*Instruccion_printlnContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)

			var _m = p.Match(SwiftgrammarSTRING)

			localctx.(*Instruccion_printlnContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(SwiftgrammarTK_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)

			var _x = p.List_expression()

			localctx.(*Instruccion_printlnContext)._list_expression = _x
		}
		{
			p.SetState(162)
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
			p.SetState(165)

			var _m = p.Match(SwiftgrammarPRINT)

			localctx.(*Instruccion_printlnContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)

			var _x = p.Expressions()

			localctx.(*Instruccion_printlnContext)._expressions = _x
		}
		{
			p.SetState(168)
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
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)

			var _m = p.Match(SwiftgrammarVAR)

			localctx.(*Instruccion_declaracionContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)

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
			p.SetState(179)

			var _m = p.Match(SwiftgrammarVAR)

			localctx.(*Instruccion_declaracionContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)

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
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

			var _x = p.Instruccion_tipo()

			localctx.(*Instruccion_declaracionContext)._instruccion_tipo = _x
		}
		{
			p.SetState(189)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)

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
			p.SetState(193)

			var _m = p.Match(SwiftgrammarLET)

			localctx.(*Instruccion_declaracionContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)

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
			p.SetState(199)

			var _m = p.Match(SwiftgrammarLET)

			localctx.(*Instruccion_declaracionContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instruccion_declaracionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)

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
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)

			var _x = p.Instruccion_tipo()

			localctx.(*Instruccion_declaracionContext)._instruccion_tipo = _x
		}
		{
			p.SetState(209)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)

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
		p.SetState(215)

		var _m = p.Match(SwiftgrammarID)

		localctx.(*Instruccion_asignacionContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(SwiftgrammarTK_IGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)

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

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IInstruccionesContext

	// GetRight_instr returns the right_instr rule contexts.
	GetRight_instr() IInstruccionesContext

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IInstruccionesContext)

	// SetRight_instr sets the right_instr rule contexts.
	SetRight_instr(IInstruccionesContext)

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
	parser      antlr.Parser
	left_instr  IInstruccionesContext
	right_instr IInstruccionesContext
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

func (s *Instruccion_ifContext) GetLeft_instr() IInstruccionesContext { return s.left_instr }

func (s *Instruccion_ifContext) GetRight_instr() IInstruccionesContext { return s.right_instr }

func (s *Instruccion_ifContext) SetLeft_instr(v IInstruccionesContext) { s.left_instr = v }

func (s *Instruccion_ifContext) SetRight_instr(v IInstruccionesContext) { s.right_instr = v }

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
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.Match(SwiftgrammarIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.Expressions()
		}
		{
			p.SetState(222)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_ifContext).left_instr = _x
		}
		{
			p.SetState(224)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.Match(SwiftgrammarIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Expressions()
		}
		{
			p.SetState(228)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_ifContext).left_instr = _x
		}
		{
			p.SetState(230)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(SwiftgrammarELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_ifContext).right_instr = _x
		}
		{
			p.SetState(234)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(236)
			p.Match(SwiftgrammarIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Expressions()
		}
		{
			p.SetState(238)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_ifContext).left_instr = _x
		}
		{
			p.SetState(240)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(SwiftgrammarELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Instr_else_if()
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

	// Getter signatures
	AllInstruccion_if() []IInstruccion_ifContext
	Instruccion_if(i int) IInstruccion_ifContext

	// IsInstr_else_ifContext differentiates from other interfaces.
	IsInstr_else_ifContext()
}

type Instr_else_ifContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(246)

				var _x = p.Instruccion_if()

				localctx.(*Instr_else_ifContext)._instruccion_if = _x
			}
			localctx.(*Instr_else_ifContext).e = append(localctx.(*Instr_else_ifContext).e, localctx.(*Instr_else_ifContext)._instruccion_if)

		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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

// IInstruccion_ternarioContext is an interface to support dynamic dispatch.
type IInstruccion_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IExpressionsContext

	// GetLeft_instr returns the left_instr rule contexts.
	GetLeft_instr() IExpressionsContext

	// GetRight_instr returns the right_instr rule contexts.
	GetRight_instr() IExpressionsContext

	// SetCond sets the cond rule contexts.
	SetCond(IExpressionsContext)

	// SetLeft_instr sets the left_instr rule contexts.
	SetLeft_instr(IExpressionsContext)

	// SetRight_instr sets the right_instr rule contexts.
	SetRight_instr(IExpressionsContext)

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
	parser      antlr.Parser
	cond        IExpressionsContext
	left_instr  IExpressionsContext
	right_instr IExpressionsContext
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

func (s *Instruccion_ternarioContext) GetCond() IExpressionsContext { return s.cond }

func (s *Instruccion_ternarioContext) GetLeft_instr() IExpressionsContext { return s.left_instr }

func (s *Instruccion_ternarioContext) GetRight_instr() IExpressionsContext { return s.right_instr }

func (s *Instruccion_ternarioContext) SetCond(v IExpressionsContext) { s.cond = v }

func (s *Instruccion_ternarioContext) SetLeft_instr(v IExpressionsContext) { s.left_instr = v }

func (s *Instruccion_ternarioContext) SetRight_instr(v IExpressionsContext) { s.right_instr = v }

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
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Match(SwiftgrammarIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).cond = _x
		}
		{
			p.SetState(254)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).left_instr = _x
		}
		{
			p.SetState(256)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.Match(SwiftgrammarIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).cond = _x
		}
		{
			p.SetState(260)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).left_instr = _x
		}
		{
			p.SetState(262)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Match(SwiftgrammarELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).right_instr = _x
		}
		{
			p.SetState(266)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(268)
			p.Match(SwiftgrammarIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).cond = _x
		}
		{
			p.SetState(270)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)

			var _x = p.Expressions()

			localctx.(*Instruccion_ternarioContext).left_instr = _x
		}
		{
			p.SetState(272)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.Match(SwiftgrammarELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.Instr_else_if_ternario()
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

	// Getter signatures
	AllInstruccion_ternario() []IInstruccion_ternarioContext
	Instruccion_ternario(i int) IInstruccion_ternarioContext

	// IsInstr_else_if_ternarioContext differentiates from other interfaces.
	IsInstr_else_if_ternarioContext()
}

type Instr_else_if_ternarioContext struct {
	antlr.BaseParserRuleContext
	parser                antlr.Parser
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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(278)

				var _x = p.Instruccion_ternario()

				localctx.(*Instr_else_if_ternarioContext)._instruccion_ternario = _x
			}
			localctx.(*Instr_else_if_ternarioContext).e = append(localctx.(*Instr_else_if_ternarioContext).e, localctx.(*Instr_else_if_ternarioContext)._instruccion_ternario)

		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Match(SwiftgrammarSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Expressions()
		}
		{
			p.SetState(286)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.List_case()
		}
		{
			p.SetState(288)
			p.Block_default()
		}
		{
			p.SetState(289)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(291)
			p.Match(SwiftgrammarSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Expressions()
		}
		{
			p.SetState(293)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Block_default()
		}
		{
			p.SetState(295)
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
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520336) != 0) {
		{
			p.SetState(299)

			var _x = p.Instruccion_case()

			localctx.(*List_caseContext)._instruccion_case = _x
		}
		localctx.(*List_caseContext).e = append(localctx.(*List_caseContext).e, localctx.(*List_caseContext)._instruccion_case)

		p.SetState(302)
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
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.List_expre_case()
		}
		{
			p.SetState(305)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Instrucciones()
		}
		{
			p.SetState(308)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.List_expre_case()
		}
		{
			p.SetState(311)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Block_instr_switch()
		}
		{
			p.SetState(313)
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
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520336) != 0) {
		{
			p.SetState(317)

			var _x = p.Block_case()

			localctx.(*List_expre_caseContext)._block_case = _x
		}
		localctx.(*List_expre_caseContext).e = append(localctx.(*List_expre_caseContext).e, localctx.(*List_expre_caseContext)._block_case)

		p.SetState(320)
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
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.Expressions()
		}
		{
			p.SetState(323)
			p.Match(SwiftgrammarTK_BARRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(325)
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
		p.SetState(328)

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
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Match(SwiftgrammarCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Instrucciones()
		}
		{
			p.SetState(335)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.Match(SwiftgrammarCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
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
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftgrammarCASE {
		{
			p.SetState(343)

			var _x = p.Instr_default()

			localctx.(*Block_defaultContext)._instr_default = _x
		}
		localctx.(*Block_defaultContext).e = append(localctx.(*Block_defaultContext).e, localctx.(*Block_defaultContext)._instr_default)

		p.SetState(346)
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
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.Match(SwiftgrammarSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Expressions()
		}
		{
			p.SetState(350)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.List_case_ternario()
		}
		{
			p.SetState(352)
			p.Instr_default_ter()
		}
		{
			p.SetState(353)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(355)
			p.Match(SwiftgrammarSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Expressions()
		}
		{
			p.SetState(357)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
			p.Instr_default_ter()
		}
		{
			p.SetState(359)
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
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520336) != 0) {
		{
			p.SetState(363)

			var _x = p.Instr_case_ter()

			localctx.(*List_case_ternarioContext)._instr_case_ter = _x
		}
		localctx.(*List_case_ternarioContext).e = append(localctx.(*List_case_ternarioContext).e, localctx.(*List_case_ternarioContext)._instr_case_ter)

		p.SetState(366)
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
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(368)
			p.List_expre_case_ter()
		}
		{
			p.SetState(369)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Expressions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.List_expre_case_ter()
		}
		{
			p.SetState(373)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Expressions()
		}
		{
			p.SetState(376)
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
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520336) != 0) {
		{
			p.SetState(380)

			var _x = p.Block_case_ter()

			localctx.(*List_expre_case_terContext)._block_case_ter = _x
		}
		localctx.(*List_expre_case_terContext).e = append(localctx.(*List_expre_case_terContext).e, localctx.(*List_expre_case_terContext)._block_case_ter)

		p.SetState(383)
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
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)
			p.Expressions()
		}
		{
			p.SetState(386)
			p.Match(SwiftgrammarTK_BARRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(388)
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
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.Match(SwiftgrammarCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.Expressions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.Match(SwiftgrammarCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.Expressions()
		}
		{
			p.SetState(400)
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
	parser antlr.Parser
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
		p.SetState(404)
		p.Match(SwiftgrammarWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)
		p.Expressions()
	}
	{
		p.SetState(406)
		p.Match(SwiftgrammarTK_LLAVEA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Instrucciones()
	}
	{
		p.SetState(408)
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

// IInstruccion_for_inContext is an interface to support dynamic dispatch.
type IInstruccion_for_inContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExpressionsContext

	// GetRight returns the right rule contexts.
	GetRight() IExpressionsContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpressionsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpressionsContext)

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
	parser antlr.Parser
	left   IExpressionsContext
	right  IExpressionsContext
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

func (s *Instruccion_for_inContext) GetLeft() IExpressionsContext { return s.left }

func (s *Instruccion_for_inContext) GetRight() IExpressionsContext { return s.right }

func (s *Instruccion_for_inContext) SetLeft(v IExpressionsContext) { s.left = v }

func (s *Instruccion_for_inContext) SetRight(v IExpressionsContext) { s.right = v }

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
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(410)
			p.Match(SwiftgrammarFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.Match(SwiftgrammarIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)

			var _x = p.Expressions()

			localctx.(*Instruccion_for_inContext).left = _x
		}
		{
			p.SetState(414)
			p.Match(SwiftgrammarTK_TRIPLEPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)

			var _x = p.Expressions()

			localctx.(*Instruccion_for_inContext).right = _x
		}
		{
			p.SetState(416)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.Instrucciones()
		}
		{
			p.SetState(418)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(420)
			p.Match(SwiftgrammarFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.Match(SwiftgrammarIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)

			var _x = p.Expressions()

			localctx.(*Instruccion_for_inContext).left = _x
		}
		{
			p.SetState(424)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.Instrucciones()
		}
		{
			p.SetState(426)
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

// IInstruccion_while_trueContext is an interface to support dynamic dispatch.
type IInstruccion_while_trueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	parser antlr.Parser
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
		p.SetState(430)
		p.Match(SwiftgrammarWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
		p.Match(SwiftgrammarTRUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(432)
		p.Match(SwiftgrammarTK_LLAVEA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Instrucciones()
	}
	{
		p.SetState(434)
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

// IInstruccion_while_true_ternarioContext is an interface to support dynamic dispatch.
type IInstruccion_while_true_ternarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	parser antlr.Parser
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
		p.SetState(436)
		p.Match(SwiftgrammarWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(437)
		p.Match(SwiftgrammarTRUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(438)
		p.Match(SwiftgrammarTK_LLAVEA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(439)
		p.Instrucciones()
	}
	{
		p.SetState(440)
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
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(442)
			p.Match(SwiftgrammarBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(443)
			p.Match(SwiftgrammarBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
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
		p.SetState(447)
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

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expressions() IExpressionsContext

	// IsInstruccion_returnContext differentiates from other interfaces.
	IsInstruccion_returnContext()
}

type Instruccion_returnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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
		p.SetState(449)
		p.Match(SwiftgrammarRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(450)
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

// IInstruccion_funcContext is an interface to support dynamic dispatch.
type IInstruccion_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	parser antlr.Parser
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
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.Match(SwiftgrammarFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(456)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.Instrucciones()
		}
		{
			p.SetState(458)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(460)
			p.Match(SwiftgrammarFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(464)
			p.Match(SwiftgrammarTK_MENOSMAYOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(465)
			p.Instruccion_tipo()
		}
		{
			p.SetState(466)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.Instrucciones()
		}
		{
			p.SetState(468)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(470)
			p.Match(SwiftgrammarFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)
			p.List_function_parameters()
		}
		{
			p.SetState(474)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(475)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.Instrucciones()
		}
		{
			p.SetState(477)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(479)
			p.Match(SwiftgrammarFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(482)
			p.List_function_parameters()
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
			p.Match(SwiftgrammarTK_MENOSMAYOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.Instruccion_tipo()
		}
		{
			p.SetState(486)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(487)
			p.Instrucciones()
		}
		{
			p.SetState(488)
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

	// Getter signatures
	AllBlock_parameters_fn() []IBlock_parameters_fnContext
	Block_parameters_fn(i int) IBlock_parameters_fnContext

	// IsList_function_parametersContext differentiates from other interfaces.
	IsList_function_parametersContext()
}

type List_function_parametersContext struct {
	antlr.BaseParserRuleContext
	parser               antlr.Parser
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftgrammarID {
		{
			p.SetState(492)

			var _x = p.Block_parameters_fn()

			localctx.(*List_function_parametersContext)._block_parameters_fn = _x
		}
		localctx.(*List_function_parametersContext).e = append(localctx.(*List_function_parametersContext).e, localctx.(*List_function_parametersContext)._block_parameters_fn)

		p.SetState(495)
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

// IBlock_parameters_fnContext is an interface to support dynamic dispatch.
type IBlock_parameters_fnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	parser antlr.Parser
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
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(497)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(499)
			p.Instruccion_tipo()
		}
		{
			p.SetState(500)
			p.Match(SwiftgrammarTK_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(502)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
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
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(507)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(508)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(509)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(510)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(511)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)
			p.List_expression()
		}
		{
			p.SetState(513)
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
	p.SetState(525)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(517)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(518)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(519)
			p.Match(SwiftgrammarTK_PARC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(520)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(521)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)
			p.List_expression()
		}
		{
			p.SetState(523)
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
		p.SetState(527)
		p.Match(SwiftgrammarSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(528)
		p.Match(SwiftgrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(529)
		p.Match(SwiftgrammarTK_LLAVEA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(530)
		p.List_struct_parameters()
	}
	{
		p.SetState(531)
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
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftgrammarID {
		{
			p.SetState(533)

			var _x = p.Block_structs_parameters()

			localctx.(*List_struct_parametersContext)._block_structs_parameters = _x
		}
		localctx.(*List_struct_parametersContext).e = append(localctx.(*List_struct_parametersContext).e, localctx.(*List_struct_parametersContext)._block_structs_parameters)

		p.SetState(536)
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
	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(538)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(539)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.Instruccion_tipo()
		}
		{
			p.SetState(541)
			p.Match(SwiftgrammarTK_COMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(543)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(544)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)
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
		p.SetState(548)
		p.Match(SwiftgrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(549)
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
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(551)

				var _x = p.Block_arrays_identifier()

				localctx.(*List_arrays_parameters_idContext)._block_arrays_identifier = _x
			}
			localctx.(*List_arrays_parameters_idContext).e = append(localctx.(*List_arrays_parameters_idContext).e, localctx.(*List_arrays_parameters_idContext)._block_arrays_identifier)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(554)
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
		p.SetState(556)
		p.Match(SwiftgrammarTK_CORA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(557)
		p.Expressions()
	}
	{
		p.SetState(558)
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
	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammarVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(560)
			p.Match(SwiftgrammarVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(561)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instr_structs_declarationContext).left = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(562)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(563)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instr_structs_declarationContext).right = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(564)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(565)
			p.List_struct_parameters_decla()
		}
		{
			p.SetState(566)
			p.Match(SwiftgrammarTK_LLAVEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SwiftgrammarLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(568)
			p.Match(SwiftgrammarLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(569)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instr_structs_declarationContext).left = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(570)
			p.Match(SwiftgrammarTK_IGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(571)

			var _m = p.Match(SwiftgrammarID)

			localctx.(*Instr_structs_declarationContext).right = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(572)
			p.Match(SwiftgrammarTK_LLAVEA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(573)
			p.List_struct_parameters_decla()
		}
		{
			p.SetState(574)
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
	p.SetState(579)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftgrammarID {
		{
			p.SetState(578)

			var _x = p.Block_structs_parameters_decla()

			localctx.(*List_struct_parameters_declaContext)._block_structs_parameters_decla = _x
		}
		localctx.(*List_struct_parameters_declaContext).e = append(localctx.(*List_struct_parameters_declaContext).e, localctx.(*List_struct_parameters_declaContext)._block_structs_parameters_decla)

		p.SetState(581)
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
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(583)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(584)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(585)
			p.Expressions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(586)
			p.Match(SwiftgrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(587)
			p.Match(SwiftgrammarTK_DOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(588)
			p.Expressions()
		}
		{
			p.SetState(589)
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
		p.SetState(593)
		p.Match(SwiftgrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(594)
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
	p.SetState(597)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(596)

				var _x = p.Block_structs_identifier()

				localctx.(*List_struct_parameters_idContext)._block_structs_identifier = _x
			}
			localctx.(*List_struct_parameters_idContext).e = append(localctx.(*List_struct_parameters_idContext).e, localctx.(*List_struct_parameters_idContext)._block_structs_identifier)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(599)
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
		p.SetState(601)
		p.Match(SwiftgrammarTK_PUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(602)
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
		p.SetState(604)
		p.Match(SwiftgrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(605)
		p.List_struct_parameters_id()
	}
	{
		p.SetState(606)
		p.Match(SwiftgrammarTK_IGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(607)
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
	p.SetState(619)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammarR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(609)
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
			p.SetState(611)
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
			p.SetState(613)
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
			p.SetState(615)
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
			p.SetState(617)
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
	p.SetState(622)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
	//	goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&577305178290520336) != 0) {
		{
			p.SetState(621)

			var _x = p.Block_list_expression()

			localctx.(*List_expressionContext)._block_list_expression = _x
		}
		localctx.(*List_expressionContext).e = append(localctx.(*List_expressionContext).e, localctx.(*List_expressionContext)._block_list_expression)

		p.SetState(624)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
		//	goto errorExit
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
	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(628)

			var _x = p.Expressions()

			localctx.(*Block_list_expressionContext)._expressions = _x
		}
		{
			p.SetState(629)
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
			p.SetState(632)

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
	p.SetState(646)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(637)

			var _x = p.expre_log(0)

			localctx.(*ExpressionsContext)._expre_log = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_log().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(640)

			var _x = p.expre_arit(0)

			localctx.(*ExpressionsContext)._expre_arit = _x
		}
		localctx.(*ExpressionsContext).p = localctx.(*ExpressionsContext).Get_expre_arit().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(643)

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
	p.SetState(656)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammarTK_NOT:
		{
			p.SetState(649)

			var _m = p.Match(SwiftgrammarTK_NOT)

			localctx.(*Expre_logContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(650)

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

	case SwiftgrammarIF, SwiftgrammarSWITCH, SwiftgrammarNUMBER, SwiftgrammarDOUBLE, SwiftgrammarCHAR, SwiftgrammarSTRING, SwiftgrammarBOOLEAN, SwiftgrammarID, SwiftgrammarTK_MENOS, SwiftgrammarTK_PARA:
		{
			p.SetState(653)

			var _x = p.expre_rel(0)

			localctx.(*Expre_logContext)._expre_rel = _x
		}
		localctx.(*Expre_logContext).p = localctx.(*Expre_logContext).Get_expre_rel().GetP()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(665)
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
			p.SetState(658)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(659)

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
				p.SetState(660)

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
		p.SetState(667)
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
		p.SetState(669)

		var _x = p.expre_arit(0)

		localctx.(*Expre_relContext)._expre_arit = _x
	}
	localctx.(*Expre_relContext).p = localctx.(*Expre_relContext).Get_expre_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(679)
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
			p.SetState(672)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(673)

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
				p.SetState(674)

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
		p.SetState(681)
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
	p.SetState(695)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftgrammarTK_MENOS:
		{
			p.SetState(683)

			var _m = p.Match(SwiftgrammarTK_MENOS)

			localctx.(*Expre_aritContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(684)

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

	case SwiftgrammarIF, SwiftgrammarSWITCH, SwiftgrammarNUMBER, SwiftgrammarDOUBLE, SwiftgrammarCHAR, SwiftgrammarSTRING, SwiftgrammarBOOLEAN, SwiftgrammarID:
		{
			p.SetState(687)

			var _x = p.Expre_valor()

			localctx.(*Expre_aritContext)._expre_valor = _x
		}
		localctx.(*Expre_aritContext).p = localctx.(*Expre_aritContext).Get_expre_valor().GetP()

	case SwiftgrammarTK_PARA:
		{
			p.SetState(690)
			p.Match(SwiftgrammarTK_PARA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(691)

			var _x = p.Expressions()

			localctx.(*Expre_aritContext)._expressions = _x
		}
		{
			p.SetState(692)
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
	p.SetState(709)
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
			p.SetState(707)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpre_aritContext(p, _parentctx, _parentState)
				localctx.(*Expre_aritContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftgrammarRULE_expre_arit)
				p.SetState(697)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(698)

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
					p.SetState(699)

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
				p.SetState(702)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(703)

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
					p.SetState(704)

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
		p.SetState(711)
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
		p.SetState(712)

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

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	p        interfaces.Expression
	_NUMBER  antlr.Token
	_DOUBLE  antlr.Token
	_STRING  antlr.Token
	_BOOLEAN antlr.Token
	_CHAR    antlr.Token
	_ID      antlr.Token
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
	p.SetState(732)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(715)

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
			p.SetState(717)

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
			p.SetState(719)

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
			p.SetState(721)

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
			p.SetState(723)

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
			p.SetState(725)
			p.Instr_llamada_expre()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(726)
			p.Instr_structs_identifier()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(727)
			p.Instr_arrays_identifier()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(728)

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
			p.SetState(730)
			p.Instruccion_ternario()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(731)
			p.Instruccion_switch_ternario()
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
