// Code generated from Swiftgrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgrammar

import "github.com/antlr4-go/antlr/v4"

// BaseSwiftgrammarListener is a complete listener for a parse tree produced by Swiftgrammar.
type BaseSwiftgrammarListener struct{}

var _ SwiftgrammarListener = &BaseSwiftgrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftgrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftgrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftgrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftgrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSwiftgrammarListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSwiftgrammarListener) ExitStart(ctx *StartContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseSwiftgrammarListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseSwiftgrammarListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterInstruccion_println is called when production instruccion_println is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_println(ctx *Instruccion_printlnContext) {}

// ExitInstruccion_println is called when production instruccion_println is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_println(ctx *Instruccion_printlnContext) {}

// EnterInstruccion_declaracion is called when production instruccion_declaracion is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_declaracion(ctx *Instruccion_declaracionContext) {
}

// ExitInstruccion_declaracion is called when production instruccion_declaracion is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_declaracion(ctx *Instruccion_declaracionContext) {}

// EnterInstruccion_asignacion is called when production instruccion_asignacion is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_asignacion(ctx *Instruccion_asignacionContext) {}

// ExitInstruccion_asignacion is called when production instruccion_asignacion is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_asignacion(ctx *Instruccion_asignacionContext) {}

// EnterInstruccion_if is called when production instruccion_if is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_if(ctx *Instruccion_ifContext) {}

// ExitInstruccion_if is called when production instruccion_if is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_if(ctx *Instruccion_ifContext) {}

// EnterInstr_else_if is called when production instr_else_if is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_else_if(ctx *Instr_else_ifContext) {}

// ExitInstr_else_if is called when production instr_else_if is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_else_if(ctx *Instr_else_ifContext) {}

// EnterInstruccion_ternario is called when production instruccion_ternario is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_ternario(ctx *Instruccion_ternarioContext) {}

// ExitInstruccion_ternario is called when production instruccion_ternario is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_ternario(ctx *Instruccion_ternarioContext) {}

// EnterInstr_else_if_ternario is called when production instr_else_if_ternario is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_else_if_ternario(ctx *Instr_else_if_ternarioContext) {}

// ExitInstr_else_if_ternario is called when production instr_else_if_ternario is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_else_if_ternario(ctx *Instr_else_if_ternarioContext) {}

// EnterInstruccion_switch is called when production instruccion_switch is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_switch(ctx *Instruccion_switchContext) {}

// ExitInstruccion_switch is called when production instruccion_switch is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_switch(ctx *Instruccion_switchContext) {}

// EnterList_case is called when production list_case is entered.
func (s *BaseSwiftgrammarListener) EnterList_case(ctx *List_caseContext) {}

// ExitList_case is called when production list_case is exited.
func (s *BaseSwiftgrammarListener) ExitList_case(ctx *List_caseContext) {}

// EnterInstruccion_case is called when production instruccion_case is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_case(ctx *Instruccion_caseContext) {}

// ExitInstruccion_case is called when production instruccion_case is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_case(ctx *Instruccion_caseContext) {}

// EnterList_expre_case is called when production list_expre_case is entered.
func (s *BaseSwiftgrammarListener) EnterList_expre_case(ctx *List_expre_caseContext) {}

// ExitList_expre_case is called when production list_expre_case is exited.
func (s *BaseSwiftgrammarListener) ExitList_expre_case(ctx *List_expre_caseContext) {}

// EnterBlock_case is called when production block_case is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_case(ctx *Block_caseContext) {}

// ExitBlock_case is called when production block_case is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_case(ctx *Block_caseContext) {}

// EnterBlock_instr_switch is called when production block_instr_switch is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_instr_switch(ctx *Block_instr_switchContext) {}

// ExitBlock_instr_switch is called when production block_instr_switch is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_instr_switch(ctx *Block_instr_switchContext) {}

// EnterInstr_default is called when production instr_default is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_default(ctx *Instr_defaultContext) {}

// ExitInstr_default is called when production instr_default is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_default(ctx *Instr_defaultContext) {}

// EnterBlock_default is called when production block_default is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_default(ctx *Block_defaultContext) {}

// ExitBlock_default is called when production block_default is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_default(ctx *Block_defaultContext) {}

// EnterInstruccion_switch_ternario is called when production instruccion_switch_ternario is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_switch_ternario(ctx *Instruccion_switch_ternarioContext) {
}

// ExitInstruccion_switch_ternario is called when production instruccion_switch_ternario is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_switch_ternario(ctx *Instruccion_switch_ternarioContext) {
}

// EnterList_case_ternario is called when production list_case_ternario is entered.
func (s *BaseSwiftgrammarListener) EnterList_case_ternario(ctx *List_case_ternarioContext) {}

// ExitList_case_ternario is called when production list_case_ternario is exited.
func (s *BaseSwiftgrammarListener) ExitList_case_ternario(ctx *List_case_ternarioContext) {}

// EnterInstr_case_ter is called when production instr_case_ter is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_case_ter(ctx *Instr_case_terContext) {}

// ExitInstr_case_ter is called when production instr_case_ter is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_case_ter(ctx *Instr_case_terContext) {}

// EnterList_expre_case_ter is called when production list_expre_case_ter is entered.
func (s *BaseSwiftgrammarListener) EnterList_expre_case_ter(ctx *List_expre_case_terContext) {}

// ExitList_expre_case_ter is called when production list_expre_case_ter is exited.
func (s *BaseSwiftgrammarListener) ExitList_expre_case_ter(ctx *List_expre_case_terContext) {}

// EnterBlock_case_ter is called when production block_case_ter is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_case_ter(ctx *Block_case_terContext) {}

// ExitBlock_case_ter is called when production block_case_ter is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_case_ter(ctx *Block_case_terContext) {}

// EnterInstr_default_ter is called when production instr_default_ter is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_default_ter(ctx *Instr_default_terContext) {}

// ExitInstr_default_ter is called when production instr_default_ter is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_default_ter(ctx *Instr_default_terContext) {}

// EnterInstruccion_while is called when production instruccion_while is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_while(ctx *Instruccion_whileContext) {}

// ExitInstruccion_while is called when production instruccion_while is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_while(ctx *Instruccion_whileContext) {}

// EnterInstruccion_for_in is called when production instruccion_for_in is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_for_in(ctx *Instruccion_for_inContext) {}

// ExitInstruccion_for_in is called when production instruccion_for_in is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_for_in(ctx *Instruccion_for_inContext) {}

// EnterInstruccion_while_true is called when production instruccion_while_true is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_while_true(ctx *Instruccion_while_trueContext) {}

// ExitInstruccion_while_true is called when production instruccion_while_true is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_while_true(ctx *Instruccion_while_trueContext) {}

// EnterInstruccion_while_true_ternario is called when production instruccion_while_true_ternario is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_while_true_ternario(ctx *Instruccion_while_true_ternarioContext) {
}

// ExitInstruccion_while_true_ternario is called when production instruccion_while_true_ternario is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_while_true_ternario(ctx *Instruccion_while_true_ternarioContext) {
}

// EnterInstruccion_break is called when production instruccion_break is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_break(ctx *Instruccion_breakContext) {}

// ExitInstruccion_break is called when production instruccion_break is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_break(ctx *Instruccion_breakContext) {}

// EnterInstruccion_continue is called when production instruccion_continue is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_continue(ctx *Instruccion_continueContext) {}

// ExitInstruccion_continue is called when production instruccion_continue is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_continue(ctx *Instruccion_continueContext) {}

// EnterInstruccion_return is called when production instruccion_return is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_return(ctx *Instruccion_returnContext) {}

// ExitInstruccion_return is called when production instruccion_return is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_return(ctx *Instruccion_returnContext) {}

// EnterInstruccion_func is called when production instruccion_func is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_func(ctx *Instruccion_funcContext) {}

// ExitInstruccion_func is called when production instruccion_func is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_func(ctx *Instruccion_funcContext) {}

// EnterList_function_parameters is called when production list_function_parameters is entered.
func (s *BaseSwiftgrammarListener) EnterList_function_parameters(ctx *List_function_parametersContext) {
}

// ExitList_function_parameters is called when production list_function_parameters is exited.
func (s *BaseSwiftgrammarListener) ExitList_function_parameters(ctx *List_function_parametersContext) {
}

// EnterBlock_parameters_fn is called when production block_parameters_fn is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_parameters_fn(ctx *Block_parameters_fnContext) {}

// ExitBlock_parameters_fn is called when production block_parameters_fn is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_parameters_fn(ctx *Block_parameters_fnContext) {}

// EnterInstruccion_llamada is called when production instruccion_llamada is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_llamada(ctx *Instruccion_llamadaContext) {}

// ExitInstruccion_llamada is called when production instruccion_llamada is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_llamada(ctx *Instruccion_llamadaContext) {}

// EnterInstr_llamada_expre is called when production instr_llamada_expre is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_llamada_expre(ctx *Instr_llamada_expreContext) {}

// ExitInstr_llamada_expre is called when production instr_llamada_expre is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_llamada_expre(ctx *Instr_llamada_expreContext) {}

// EnterInstruccion_structs_declaracion is called when production instruccion_structs_declaracion is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_structs_declaracion(ctx *Instruccion_structs_declaracionContext) {
}

// ExitInstruccion_structs_declaracion is called when production instruccion_structs_declaracion is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_structs_declaracion(ctx *Instruccion_structs_declaracionContext) {
}

// EnterList_struct_parameters is called when production list_struct_parameters is entered.
func (s *BaseSwiftgrammarListener) EnterList_struct_parameters(ctx *List_struct_parametersContext) {}

// ExitList_struct_parameters is called when production list_struct_parameters is exited.
func (s *BaseSwiftgrammarListener) ExitList_struct_parameters(ctx *List_struct_parametersContext) {}

// EnterBlock_structs_parameters is called when production block_structs_parameters is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_structs_parameters(ctx *Block_structs_parametersContext) {
}

// ExitBlock_structs_parameters is called when production block_structs_parameters is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_structs_parameters(ctx *Block_structs_parametersContext) {
}

// EnterInstr_arrays_identifier is called when production instr_arrays_identifier is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_arrays_identifier(ctx *Instr_arrays_identifierContext) {
}

// ExitInstr_arrays_identifier is called when production instr_arrays_identifier is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_arrays_identifier(ctx *Instr_arrays_identifierContext) {}

// EnterList_arrays_parameters_id is called when production list_arrays_parameters_id is entered.
func (s *BaseSwiftgrammarListener) EnterList_arrays_parameters_id(ctx *List_arrays_parameters_idContext) {
}

// ExitList_arrays_parameters_id is called when production list_arrays_parameters_id is exited.
func (s *BaseSwiftgrammarListener) ExitList_arrays_parameters_id(ctx *List_arrays_parameters_idContext) {
}

// EnterBlock_arrays_identifier is called when production block_arrays_identifier is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_arrays_identifier(ctx *Block_arrays_identifierContext) {
}

// ExitBlock_arrays_identifier is called when production block_arrays_identifier is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_arrays_identifier(ctx *Block_arrays_identifierContext) {}

// EnterInstr_structs_declaration is called when production instr_structs_declaration is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_structs_declaration(ctx *Instr_structs_declarationContext) {
}

// ExitInstr_structs_declaration is called when production instr_structs_declaration is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_structs_declaration(ctx *Instr_structs_declarationContext) {
}

// EnterList_struct_parameters_decla is called when production list_struct_parameters_decla is entered.
func (s *BaseSwiftgrammarListener) EnterList_struct_parameters_decla(ctx *List_struct_parameters_declaContext) {
}

// ExitList_struct_parameters_decla is called when production list_struct_parameters_decla is exited.
func (s *BaseSwiftgrammarListener) ExitList_struct_parameters_decla(ctx *List_struct_parameters_declaContext) {
}

// EnterBlock_structs_parameters_decla is called when production block_structs_parameters_decla is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_structs_parameters_decla(ctx *Block_structs_parameters_declaContext) {
}

// ExitBlock_structs_parameters_decla is called when production block_structs_parameters_decla is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_structs_parameters_decla(ctx *Block_structs_parameters_declaContext) {
}

// EnterInstr_structs_identifier is called when production instr_structs_identifier is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_structs_identifier(ctx *Instr_structs_identifierContext) {
}

// ExitInstr_structs_identifier is called when production instr_structs_identifier is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_structs_identifier(ctx *Instr_structs_identifierContext) {
}

// EnterList_struct_parameters_id is called when production list_struct_parameters_id is entered.
func (s *BaseSwiftgrammarListener) EnterList_struct_parameters_id(ctx *List_struct_parameters_idContext) {
}

// ExitList_struct_parameters_id is called when production list_struct_parameters_id is exited.
func (s *BaseSwiftgrammarListener) ExitList_struct_parameters_id(ctx *List_struct_parameters_idContext) {
}

// EnterBlock_structs_identifier is called when production block_structs_identifier is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_structs_identifier(ctx *Block_structs_identifierContext) {
}

// ExitBlock_structs_identifier is called when production block_structs_identifier is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_structs_identifier(ctx *Block_structs_identifierContext) {
}

// EnterInstr_structs_assignment is called when production instr_structs_assignment is entered.
func (s *BaseSwiftgrammarListener) EnterInstr_structs_assignment(ctx *Instr_structs_assignmentContext) {
}

// ExitInstr_structs_assignment is called when production instr_structs_assignment is exited.
func (s *BaseSwiftgrammarListener) ExitInstr_structs_assignment(ctx *Instr_structs_assignmentContext) {
}

// EnterInstruccion_tipo is called when production instruccion_tipo is entered.
func (s *BaseSwiftgrammarListener) EnterInstruccion_tipo(ctx *Instruccion_tipoContext) {}

// ExitInstruccion_tipo is called when production instruccion_tipo is exited.
func (s *BaseSwiftgrammarListener) ExitInstruccion_tipo(ctx *Instruccion_tipoContext) {}

// EnterList_expression is called when production list_expression is entered.
func (s *BaseSwiftgrammarListener) EnterList_expression(ctx *List_expressionContext) {}

// ExitList_expression is called when production list_expression is exited.
func (s *BaseSwiftgrammarListener) ExitList_expression(ctx *List_expressionContext) {}

// EnterBlock_list_expression is called when production block_list_expression is entered.
func (s *BaseSwiftgrammarListener) EnterBlock_list_expression(ctx *Block_list_expressionContext) {}

// ExitBlock_list_expression is called when production block_list_expression is exited.
func (s *BaseSwiftgrammarListener) ExitBlock_list_expression(ctx *Block_list_expressionContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseSwiftgrammarListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseSwiftgrammarListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterExpre_log is called when production expre_log is entered.
func (s *BaseSwiftgrammarListener) EnterExpre_log(ctx *Expre_logContext) {}

// ExitExpre_log is called when production expre_log is exited.
func (s *BaseSwiftgrammarListener) ExitExpre_log(ctx *Expre_logContext) {}

// EnterExpre_rel is called when production expre_rel is entered.
func (s *BaseSwiftgrammarListener) EnterExpre_rel(ctx *Expre_relContext) {}

// ExitExpre_rel is called when production expre_rel is exited.
func (s *BaseSwiftgrammarListener) ExitExpre_rel(ctx *Expre_relContext) {}

// EnterExpre_arit is called when production expre_arit is entered.
func (s *BaseSwiftgrammarListener) EnterExpre_arit(ctx *Expre_aritContext) {}

// ExitExpre_arit is called when production expre_arit is exited.
func (s *BaseSwiftgrammarListener) ExitExpre_arit(ctx *Expre_aritContext) {}

// EnterExpre_valor is called when production expre_valor is entered.
func (s *BaseSwiftgrammarListener) EnterExpre_valor(ctx *Expre_valorContext) {}

// ExitExpre_valor is called when production expre_valor is exited.
func (s *BaseSwiftgrammarListener) ExitExpre_valor(ctx *Expre_valorContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BaseSwiftgrammarListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BaseSwiftgrammarListener) ExitPrimitivo(ctx *PrimitivoContext) {}
