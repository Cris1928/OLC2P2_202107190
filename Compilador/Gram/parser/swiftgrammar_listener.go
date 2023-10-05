// Code generated from Swiftgrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Swiftgrammar

import "github.com/antlr4-go/antlr/v4"

// SwiftgrammarListener is a complete listener for a parse tree produced by Swiftgrammar.
type SwiftgrammarListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInstruccion_println is called when entering the instruccion_println production.
	EnterInstruccion_println(c *Instruccion_printlnContext)

	// EnterInstruccion_declaracion is called when entering the instruccion_declaracion production.
	EnterInstruccion_declaracion(c *Instruccion_declaracionContext)

	// EnterInstruccion_asignacion is called when entering the instruccion_asignacion production.
	EnterInstruccion_asignacion(c *Instruccion_asignacionContext)

	// EnterInstruccion_if is called when entering the instruccion_if production.
	EnterInstruccion_if(c *Instruccion_ifContext)

	// EnterInstr_else_if is called when entering the instr_else_if production.
	EnterInstr_else_if(c *Instr_else_ifContext)

	// EnterInstruccion_ternario is called when entering the instruccion_ternario production.
	EnterInstruccion_ternario(c *Instruccion_ternarioContext)

	// EnterInstr_else_if_ternario is called when entering the instr_else_if_ternario production.
	EnterInstr_else_if_ternario(c *Instr_else_if_ternarioContext)

	// EnterInstruccion_switch is called when entering the instruccion_switch production.
	EnterInstruccion_switch(c *Instruccion_switchContext)

	// EnterList_case is called when entering the list_case production.
	EnterList_case(c *List_caseContext)

	// EnterInstruccion_case is called when entering the instruccion_case production.
	EnterInstruccion_case(c *Instruccion_caseContext)

	// EnterList_expre_case is called when entering the list_expre_case production.
	EnterList_expre_case(c *List_expre_caseContext)

	// EnterBlock_case is called when entering the block_case production.
	EnterBlock_case(c *Block_caseContext)

	// EnterBlock_instr_switch is called when entering the block_instr_switch production.
	EnterBlock_instr_switch(c *Block_instr_switchContext)

	// EnterInstr_default is called when entering the instr_default production.
	EnterInstr_default(c *Instr_defaultContext)

	// EnterBlock_default is called when entering the block_default production.
	EnterBlock_default(c *Block_defaultContext)

	// EnterInstruccion_switch_ternario is called when entering the instruccion_switch_ternario production.
	EnterInstruccion_switch_ternario(c *Instruccion_switch_ternarioContext)

	// EnterList_case_ternario is called when entering the list_case_ternario production.
	EnterList_case_ternario(c *List_case_ternarioContext)

	// EnterInstr_case_ter is called when entering the instr_case_ter production.
	EnterInstr_case_ter(c *Instr_case_terContext)

	// EnterList_expre_case_ter is called when entering the list_expre_case_ter production.
	EnterList_expre_case_ter(c *List_expre_case_terContext)

	// EnterBlock_case_ter is called when entering the block_case_ter production.
	EnterBlock_case_ter(c *Block_case_terContext)

	// EnterInstr_default_ter is called when entering the instr_default_ter production.
	EnterInstr_default_ter(c *Instr_default_terContext)

	// EnterInstruccion_while is called when entering the instruccion_while production.
	EnterInstruccion_while(c *Instruccion_whileContext)

	// EnterInstruccion_for_in is called when entering the instruccion_for_in production.
	EnterInstruccion_for_in(c *Instruccion_for_inContext)

	// EnterInstruccion_while_true is called when entering the instruccion_while_true production.
	EnterInstruccion_while_true(c *Instruccion_while_trueContext)

	// EnterInstruccion_while_true_ternario is called when entering the instruccion_while_true_ternario production.
	EnterInstruccion_while_true_ternario(c *Instruccion_while_true_ternarioContext)

	// EnterInstruccion_break is called when entering the instruccion_break production.
	EnterInstruccion_break(c *Instruccion_breakContext)

	// EnterInstruccion_continue is called when entering the instruccion_continue production.
	EnterInstruccion_continue(c *Instruccion_continueContext)

	// EnterInstruccion_return is called when entering the instruccion_return production.
	EnterInstruccion_return(c *Instruccion_returnContext)

	// EnterInstruccion_func is called when entering the instruccion_func production.
	EnterInstruccion_func(c *Instruccion_funcContext)

	// EnterList_function_parameters is called when entering the list_function_parameters production.
	EnterList_function_parameters(c *List_function_parametersContext)

	// EnterBlock_parameters_fn is called when entering the block_parameters_fn production.
	EnterBlock_parameters_fn(c *Block_parameters_fnContext)

	// EnterInstruccion_llamada is called when entering the instruccion_llamada production.
	EnterInstruccion_llamada(c *Instruccion_llamadaContext)

	// EnterInstr_llamada_expre is called when entering the instr_llamada_expre production.
	EnterInstr_llamada_expre(c *Instr_llamada_expreContext)

	// EnterInstruccion_structs_declaracion is called when entering the instruccion_structs_declaracion production.
	EnterInstruccion_structs_declaracion(c *Instruccion_structs_declaracionContext)

	// EnterList_struct_parameters is called when entering the list_struct_parameters production.
	EnterList_struct_parameters(c *List_struct_parametersContext)

	// EnterBlock_structs_parameters is called when entering the block_structs_parameters production.
	EnterBlock_structs_parameters(c *Block_structs_parametersContext)

	// EnterInstr_arrays_identifier is called when entering the instr_arrays_identifier production.
	EnterInstr_arrays_identifier(c *Instr_arrays_identifierContext)

	// EnterList_arrays_parameters_id is called when entering the list_arrays_parameters_id production.
	EnterList_arrays_parameters_id(c *List_arrays_parameters_idContext)

	// EnterBlock_arrays_identifier is called when entering the block_arrays_identifier production.
	EnterBlock_arrays_identifier(c *Block_arrays_identifierContext)

	// EnterInstr_structs_declaration is called when entering the instr_structs_declaration production.
	EnterInstr_structs_declaration(c *Instr_structs_declarationContext)

	// EnterList_struct_parameters_decla is called when entering the list_struct_parameters_decla production.
	EnterList_struct_parameters_decla(c *List_struct_parameters_declaContext)

	// EnterBlock_structs_parameters_decla is called when entering the block_structs_parameters_decla production.
	EnterBlock_structs_parameters_decla(c *Block_structs_parameters_declaContext)

	// EnterInstr_structs_identifier is called when entering the instr_structs_identifier production.
	EnterInstr_structs_identifier(c *Instr_structs_identifierContext)

	// EnterList_struct_parameters_id is called when entering the list_struct_parameters_id production.
	EnterList_struct_parameters_id(c *List_struct_parameters_idContext)

	// EnterBlock_structs_identifier is called when entering the block_structs_identifier production.
	EnterBlock_structs_identifier(c *Block_structs_identifierContext)

	// EnterInstr_structs_assignment is called when entering the instr_structs_assignment production.
	EnterInstr_structs_assignment(c *Instr_structs_assignmentContext)

	// EnterInstruccion_tipo is called when entering the instruccion_tipo production.
	EnterInstruccion_tipo(c *Instruccion_tipoContext)

	// EnterList_expression is called when entering the list_expression production.
	EnterList_expression(c *List_expressionContext)

	// EnterBlock_list_expression is called when entering the block_list_expression production.
	EnterBlock_list_expression(c *Block_list_expressionContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterExpre_log is called when entering the expre_log production.
	EnterExpre_log(c *Expre_logContext)

	// EnterExpre_rel is called when entering the expre_rel production.
	EnterExpre_rel(c *Expre_relContext)

	// EnterExpre_arit is called when entering the expre_arit production.
	EnterExpre_arit(c *Expre_aritContext)

	// EnterExpre_valor is called when entering the expre_valor production.
	EnterExpre_valor(c *Expre_valorContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitInstruccion_println is called when exiting the instruccion_println production.
	ExitInstruccion_println(c *Instruccion_printlnContext)

	// ExitInstruccion_declaracion is called when exiting the instruccion_declaracion production.
	ExitInstruccion_declaracion(c *Instruccion_declaracionContext)

	// ExitInstruccion_asignacion is called when exiting the instruccion_asignacion production.
	ExitInstruccion_asignacion(c *Instruccion_asignacionContext)

	// ExitInstruccion_if is called when exiting the instruccion_if production.
	ExitInstruccion_if(c *Instruccion_ifContext)

	// ExitInstr_else_if is called when exiting the instr_else_if production.
	ExitInstr_else_if(c *Instr_else_ifContext)

	// ExitInstruccion_ternario is called when exiting the instruccion_ternario production.
	ExitInstruccion_ternario(c *Instruccion_ternarioContext)

	// ExitInstr_else_if_ternario is called when exiting the instr_else_if_ternario production.
	ExitInstr_else_if_ternario(c *Instr_else_if_ternarioContext)

	// ExitInstruccion_switch is called when exiting the instruccion_switch production.
	ExitInstruccion_switch(c *Instruccion_switchContext)

	// ExitList_case is called when exiting the list_case production.
	ExitList_case(c *List_caseContext)

	// ExitInstruccion_case is called when exiting the instruccion_case production.
	ExitInstruccion_case(c *Instruccion_caseContext)

	// ExitList_expre_case is called when exiting the list_expre_case production.
	ExitList_expre_case(c *List_expre_caseContext)

	// ExitBlock_case is called when exiting the block_case production.
	ExitBlock_case(c *Block_caseContext)

	// ExitBlock_instr_switch is called when exiting the block_instr_switch production.
	ExitBlock_instr_switch(c *Block_instr_switchContext)

	// ExitInstr_default is called when exiting the instr_default production.
	ExitInstr_default(c *Instr_defaultContext)

	// ExitBlock_default is called when exiting the block_default production.
	ExitBlock_default(c *Block_defaultContext)

	// ExitInstruccion_switch_ternario is called when exiting the instruccion_switch_ternario production.
	ExitInstruccion_switch_ternario(c *Instruccion_switch_ternarioContext)

	// ExitList_case_ternario is called when exiting the list_case_ternario production.
	ExitList_case_ternario(c *List_case_ternarioContext)

	// ExitInstr_case_ter is called when exiting the instr_case_ter production.
	ExitInstr_case_ter(c *Instr_case_terContext)

	// ExitList_expre_case_ter is called when exiting the list_expre_case_ter production.
	ExitList_expre_case_ter(c *List_expre_case_terContext)

	// ExitBlock_case_ter is called when exiting the block_case_ter production.
	ExitBlock_case_ter(c *Block_case_terContext)

	// ExitInstr_default_ter is called when exiting the instr_default_ter production.
	ExitInstr_default_ter(c *Instr_default_terContext)

	// ExitInstruccion_while is called when exiting the instruccion_while production.
	ExitInstruccion_while(c *Instruccion_whileContext)

	// ExitInstruccion_for_in is called when exiting the instruccion_for_in production.
	ExitInstruccion_for_in(c *Instruccion_for_inContext)

	// ExitInstruccion_while_true is called when exiting the instruccion_while_true production.
	ExitInstruccion_while_true(c *Instruccion_while_trueContext)

	// ExitInstruccion_while_true_ternario is called when exiting the instruccion_while_true_ternario production.
	ExitInstruccion_while_true_ternario(c *Instruccion_while_true_ternarioContext)

	// ExitInstruccion_break is called when exiting the instruccion_break production.
	ExitInstruccion_break(c *Instruccion_breakContext)

	// ExitInstruccion_continue is called when exiting the instruccion_continue production.
	ExitInstruccion_continue(c *Instruccion_continueContext)

	// ExitInstruccion_return is called when exiting the instruccion_return production.
	ExitInstruccion_return(c *Instruccion_returnContext)

	// ExitInstruccion_func is called when exiting the instruccion_func production.
	ExitInstruccion_func(c *Instruccion_funcContext)

	// ExitList_function_parameters is called when exiting the list_function_parameters production.
	ExitList_function_parameters(c *List_function_parametersContext)

	// ExitBlock_parameters_fn is called when exiting the block_parameters_fn production.
	ExitBlock_parameters_fn(c *Block_parameters_fnContext)

	// ExitInstruccion_llamada is called when exiting the instruccion_llamada production.
	ExitInstruccion_llamada(c *Instruccion_llamadaContext)

	// ExitInstr_llamada_expre is called when exiting the instr_llamada_expre production.
	ExitInstr_llamada_expre(c *Instr_llamada_expreContext)

	// ExitInstruccion_structs_declaracion is called when exiting the instruccion_structs_declaracion production.
	ExitInstruccion_structs_declaracion(c *Instruccion_structs_declaracionContext)

	// ExitList_struct_parameters is called when exiting the list_struct_parameters production.
	ExitList_struct_parameters(c *List_struct_parametersContext)

	// ExitBlock_structs_parameters is called when exiting the block_structs_parameters production.
	ExitBlock_structs_parameters(c *Block_structs_parametersContext)

	// ExitInstr_arrays_identifier is called when exiting the instr_arrays_identifier production.
	ExitInstr_arrays_identifier(c *Instr_arrays_identifierContext)

	// ExitList_arrays_parameters_id is called when exiting the list_arrays_parameters_id production.
	ExitList_arrays_parameters_id(c *List_arrays_parameters_idContext)

	// ExitBlock_arrays_identifier is called when exiting the block_arrays_identifier production.
	ExitBlock_arrays_identifier(c *Block_arrays_identifierContext)

	// ExitInstr_structs_declaration is called when exiting the instr_structs_declaration production.
	ExitInstr_structs_declaration(c *Instr_structs_declarationContext)

	// ExitList_struct_parameters_decla is called when exiting the list_struct_parameters_decla production.
	ExitList_struct_parameters_decla(c *List_struct_parameters_declaContext)

	// ExitBlock_structs_parameters_decla is called when exiting the block_structs_parameters_decla production.
	ExitBlock_structs_parameters_decla(c *Block_structs_parameters_declaContext)

	// ExitInstr_structs_identifier is called when exiting the instr_structs_identifier production.
	ExitInstr_structs_identifier(c *Instr_structs_identifierContext)

	// ExitList_struct_parameters_id is called when exiting the list_struct_parameters_id production.
	ExitList_struct_parameters_id(c *List_struct_parameters_idContext)

	// ExitBlock_structs_identifier is called when exiting the block_structs_identifier production.
	ExitBlock_structs_identifier(c *Block_structs_identifierContext)

	// ExitInstr_structs_assignment is called when exiting the instr_structs_assignment production.
	ExitInstr_structs_assignment(c *Instr_structs_assignmentContext)

	// ExitInstruccion_tipo is called when exiting the instruccion_tipo production.
	ExitInstruccion_tipo(c *Instruccion_tipoContext)

	// ExitList_expression is called when exiting the list_expression production.
	ExitList_expression(c *List_expressionContext)

	// ExitBlock_list_expression is called when exiting the block_list_expression production.
	ExitBlock_list_expression(c *Block_list_expressionContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitExpre_log is called when exiting the expre_log production.
	ExitExpre_log(c *Expre_logContext)

	// ExitExpre_rel is called when exiting the expre_rel production.
	ExitExpre_rel(c *Expre_relContext)

	// ExitExpre_arit is called when exiting the expre_arit production.
	ExitExpre_arit(c *Expre_aritContext)

	// ExitExpre_valor is called when exiting the expre_valor production.
	ExitExpre_valor(c *Expre_valorContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)
}
