// Generated from c:/OLC2P2_202107190/Compilador/Gram/Swiftgrammar.g4 by ANTLR 4.13.1


import "OLC2/Compilador/interfaces"
import "OLC2/Compilador/instruccion"
import "OLC2/Compilador/expression"
import "OLC2/Compilador/instruccion/ternario"
import arrayList "github.com/colegno/arraylist"


import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link Swiftgrammar}.
 */
public interface SwiftgrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#start}.
	 * @param ctx the parse tree
	 */
	void enterStart(Swiftgrammar.StartContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#start}.
	 * @param ctx the parse tree
	 */
	void exitStart(Swiftgrammar.StartContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instrucciones}.
	 * @param ctx the parse tree
	 */
	void enterInstrucciones(Swiftgrammar.InstruccionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instrucciones}.
	 * @param ctx the parse tree
	 */
	void exitInstrucciones(Swiftgrammar.InstruccionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion(Swiftgrammar.InstruccionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion(Swiftgrammar.InstruccionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_println}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_println(Swiftgrammar.Instruccion_printlnContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_println}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_println(Swiftgrammar.Instruccion_printlnContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_declaracion}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_declaracion(Swiftgrammar.Instruccion_declaracionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_declaracion}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_declaracion(Swiftgrammar.Instruccion_declaracionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_asignacion}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_asignacion(Swiftgrammar.Instruccion_asignacionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_asignacion}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_asignacion(Swiftgrammar.Instruccion_asignacionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_if}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_if(Swiftgrammar.Instruccion_ifContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_if}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_if(Swiftgrammar.Instruccion_ifContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_else_if}.
	 * @param ctx the parse tree
	 */
	void enterInstr_else_if(Swiftgrammar.Instr_else_ifContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_else_if}.
	 * @param ctx the parse tree
	 */
	void exitInstr_else_if(Swiftgrammar.Instr_else_ifContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_ternario}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_ternario(Swiftgrammar.Instruccion_ternarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_ternario}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_ternario(Swiftgrammar.Instruccion_ternarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_else_if_ternario}.
	 * @param ctx the parse tree
	 */
	void enterInstr_else_if_ternario(Swiftgrammar.Instr_else_if_ternarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_else_if_ternario}.
	 * @param ctx the parse tree
	 */
	void exitInstr_else_if_ternario(Swiftgrammar.Instr_else_if_ternarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_switch}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_switch(Swiftgrammar.Instruccion_switchContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_switch}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_switch(Swiftgrammar.Instruccion_switchContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_case}.
	 * @param ctx the parse tree
	 */
	void enterList_case(Swiftgrammar.List_caseContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_case}.
	 * @param ctx the parse tree
	 */
	void exitList_case(Swiftgrammar.List_caseContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_case}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_case(Swiftgrammar.Instruccion_caseContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_case}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_case(Swiftgrammar.Instruccion_caseContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_expre_case}.
	 * @param ctx the parse tree
	 */
	void enterList_expre_case(Swiftgrammar.List_expre_caseContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_expre_case}.
	 * @param ctx the parse tree
	 */
	void exitList_expre_case(Swiftgrammar.List_expre_caseContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_case}.
	 * @param ctx the parse tree
	 */
	void enterBlock_case(Swiftgrammar.Block_caseContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_case}.
	 * @param ctx the parse tree
	 */
	void exitBlock_case(Swiftgrammar.Block_caseContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_instr_switch}.
	 * @param ctx the parse tree
	 */
	void enterBlock_instr_switch(Swiftgrammar.Block_instr_switchContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_instr_switch}.
	 * @param ctx the parse tree
	 */
	void exitBlock_instr_switch(Swiftgrammar.Block_instr_switchContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_default}.
	 * @param ctx the parse tree
	 */
	void enterInstr_default(Swiftgrammar.Instr_defaultContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_default}.
	 * @param ctx the parse tree
	 */
	void exitInstr_default(Swiftgrammar.Instr_defaultContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_default}.
	 * @param ctx the parse tree
	 */
	void enterBlock_default(Swiftgrammar.Block_defaultContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_default}.
	 * @param ctx the parse tree
	 */
	void exitBlock_default(Swiftgrammar.Block_defaultContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_switch_ternario}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_switch_ternario(Swiftgrammar.Instruccion_switch_ternarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_switch_ternario}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_switch_ternario(Swiftgrammar.Instruccion_switch_ternarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_case_ternario}.
	 * @param ctx the parse tree
	 */
	void enterList_case_ternario(Swiftgrammar.List_case_ternarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_case_ternario}.
	 * @param ctx the parse tree
	 */
	void exitList_case_ternario(Swiftgrammar.List_case_ternarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_case_ter}.
	 * @param ctx the parse tree
	 */
	void enterInstr_case_ter(Swiftgrammar.Instr_case_terContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_case_ter}.
	 * @param ctx the parse tree
	 */
	void exitInstr_case_ter(Swiftgrammar.Instr_case_terContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_expre_case_ter}.
	 * @param ctx the parse tree
	 */
	void enterList_expre_case_ter(Swiftgrammar.List_expre_case_terContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_expre_case_ter}.
	 * @param ctx the parse tree
	 */
	void exitList_expre_case_ter(Swiftgrammar.List_expre_case_terContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_case_ter}.
	 * @param ctx the parse tree
	 */
	void enterBlock_case_ter(Swiftgrammar.Block_case_terContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_case_ter}.
	 * @param ctx the parse tree
	 */
	void exitBlock_case_ter(Swiftgrammar.Block_case_terContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_default_ter}.
	 * @param ctx the parse tree
	 */
	void enterInstr_default_ter(Swiftgrammar.Instr_default_terContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_default_ter}.
	 * @param ctx the parse tree
	 */
	void exitInstr_default_ter(Swiftgrammar.Instr_default_terContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_while}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_while(Swiftgrammar.Instruccion_whileContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_while}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_while(Swiftgrammar.Instruccion_whileContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_for_in}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_for_in(Swiftgrammar.Instruccion_for_inContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_for_in}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_for_in(Swiftgrammar.Instruccion_for_inContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_while_true}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_while_true(Swiftgrammar.Instruccion_while_trueContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_while_true}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_while_true(Swiftgrammar.Instruccion_while_trueContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_while_true_ternario}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_while_true_ternario(Swiftgrammar.Instruccion_while_true_ternarioContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_while_true_ternario}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_while_true_ternario(Swiftgrammar.Instruccion_while_true_ternarioContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_break}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_break(Swiftgrammar.Instruccion_breakContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_break}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_break(Swiftgrammar.Instruccion_breakContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_continue}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_continue(Swiftgrammar.Instruccion_continueContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_continue}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_continue(Swiftgrammar.Instruccion_continueContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_return}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_return(Swiftgrammar.Instruccion_returnContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_return}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_return(Swiftgrammar.Instruccion_returnContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_func}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_func(Swiftgrammar.Instruccion_funcContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_func}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_func(Swiftgrammar.Instruccion_funcContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_function_parameters}.
	 * @param ctx the parse tree
	 */
	void enterList_function_parameters(Swiftgrammar.List_function_parametersContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_function_parameters}.
	 * @param ctx the parse tree
	 */
	void exitList_function_parameters(Swiftgrammar.List_function_parametersContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_parameters_fn}.
	 * @param ctx the parse tree
	 */
	void enterBlock_parameters_fn(Swiftgrammar.Block_parameters_fnContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_parameters_fn}.
	 * @param ctx the parse tree
	 */
	void exitBlock_parameters_fn(Swiftgrammar.Block_parameters_fnContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_llamada}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_llamada(Swiftgrammar.Instruccion_llamadaContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_llamada}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_llamada(Swiftgrammar.Instruccion_llamadaContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_llamada_expre}.
	 * @param ctx the parse tree
	 */
	void enterInstr_llamada_expre(Swiftgrammar.Instr_llamada_expreContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_llamada_expre}.
	 * @param ctx the parse tree
	 */
	void exitInstr_llamada_expre(Swiftgrammar.Instr_llamada_expreContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_structs_declaracion}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_structs_declaracion(Swiftgrammar.Instruccion_structs_declaracionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_structs_declaracion}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_structs_declaracion(Swiftgrammar.Instruccion_structs_declaracionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_struct_parameters}.
	 * @param ctx the parse tree
	 */
	void enterList_struct_parameters(Swiftgrammar.List_struct_parametersContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_struct_parameters}.
	 * @param ctx the parse tree
	 */
	void exitList_struct_parameters(Swiftgrammar.List_struct_parametersContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_structs_parameters}.
	 * @param ctx the parse tree
	 */
	void enterBlock_structs_parameters(Swiftgrammar.Block_structs_parametersContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_structs_parameters}.
	 * @param ctx the parse tree
	 */
	void exitBlock_structs_parameters(Swiftgrammar.Block_structs_parametersContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_arrays_identifier}.
	 * @param ctx the parse tree
	 */
	void enterInstr_arrays_identifier(Swiftgrammar.Instr_arrays_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_arrays_identifier}.
	 * @param ctx the parse tree
	 */
	void exitInstr_arrays_identifier(Swiftgrammar.Instr_arrays_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_arrays_parameters_id}.
	 * @param ctx the parse tree
	 */
	void enterList_arrays_parameters_id(Swiftgrammar.List_arrays_parameters_idContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_arrays_parameters_id}.
	 * @param ctx the parse tree
	 */
	void exitList_arrays_parameters_id(Swiftgrammar.List_arrays_parameters_idContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_arrays_identifier}.
	 * @param ctx the parse tree
	 */
	void enterBlock_arrays_identifier(Swiftgrammar.Block_arrays_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_arrays_identifier}.
	 * @param ctx the parse tree
	 */
	void exitBlock_arrays_identifier(Swiftgrammar.Block_arrays_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_structs_declaration}.
	 * @param ctx the parse tree
	 */
	void enterInstr_structs_declaration(Swiftgrammar.Instr_structs_declarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_structs_declaration}.
	 * @param ctx the parse tree
	 */
	void exitInstr_structs_declaration(Swiftgrammar.Instr_structs_declarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_struct_parameters_decla}.
	 * @param ctx the parse tree
	 */
	void enterList_struct_parameters_decla(Swiftgrammar.List_struct_parameters_declaContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_struct_parameters_decla}.
	 * @param ctx the parse tree
	 */
	void exitList_struct_parameters_decla(Swiftgrammar.List_struct_parameters_declaContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_structs_parameters_decla}.
	 * @param ctx the parse tree
	 */
	void enterBlock_structs_parameters_decla(Swiftgrammar.Block_structs_parameters_declaContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_structs_parameters_decla}.
	 * @param ctx the parse tree
	 */
	void exitBlock_structs_parameters_decla(Swiftgrammar.Block_structs_parameters_declaContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_structs_identifier}.
	 * @param ctx the parse tree
	 */
	void enterInstr_structs_identifier(Swiftgrammar.Instr_structs_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_structs_identifier}.
	 * @param ctx the parse tree
	 */
	void exitInstr_structs_identifier(Swiftgrammar.Instr_structs_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_struct_parameters_id}.
	 * @param ctx the parse tree
	 */
	void enterList_struct_parameters_id(Swiftgrammar.List_struct_parameters_idContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_struct_parameters_id}.
	 * @param ctx the parse tree
	 */
	void exitList_struct_parameters_id(Swiftgrammar.List_struct_parameters_idContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_structs_identifier}.
	 * @param ctx the parse tree
	 */
	void enterBlock_structs_identifier(Swiftgrammar.Block_structs_identifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_structs_identifier}.
	 * @param ctx the parse tree
	 */
	void exitBlock_structs_identifier(Swiftgrammar.Block_structs_identifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instr_structs_assignment}.
	 * @param ctx the parse tree
	 */
	void enterInstr_structs_assignment(Swiftgrammar.Instr_structs_assignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instr_structs_assignment}.
	 * @param ctx the parse tree
	 */
	void exitInstr_structs_assignment(Swiftgrammar.Instr_structs_assignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#instruccion_tipo}.
	 * @param ctx the parse tree
	 */
	void enterInstruccion_tipo(Swiftgrammar.Instruccion_tipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#instruccion_tipo}.
	 * @param ctx the parse tree
	 */
	void exitInstruccion_tipo(Swiftgrammar.Instruccion_tipoContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#list_expression}.
	 * @param ctx the parse tree
	 */
	void enterList_expression(Swiftgrammar.List_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#list_expression}.
	 * @param ctx the parse tree
	 */
	void exitList_expression(Swiftgrammar.List_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#block_list_expression}.
	 * @param ctx the parse tree
	 */
	void enterBlock_list_expression(Swiftgrammar.Block_list_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#block_list_expression}.
	 * @param ctx the parse tree
	 */
	void exitBlock_list_expression(Swiftgrammar.Block_list_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#expressions}.
	 * @param ctx the parse tree
	 */
	void enterExpressions(Swiftgrammar.ExpressionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#expressions}.
	 * @param ctx the parse tree
	 */
	void exitExpressions(Swiftgrammar.ExpressionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#expre_log}.
	 * @param ctx the parse tree
	 */
	void enterExpre_log(Swiftgrammar.Expre_logContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#expre_log}.
	 * @param ctx the parse tree
	 */
	void exitExpre_log(Swiftgrammar.Expre_logContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#expre_rel}.
	 * @param ctx the parse tree
	 */
	void enterExpre_rel(Swiftgrammar.Expre_relContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#expre_rel}.
	 * @param ctx the parse tree
	 */
	void exitExpre_rel(Swiftgrammar.Expre_relContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#expre_arit}.
	 * @param ctx the parse tree
	 */
	void enterExpre_arit(Swiftgrammar.Expre_aritContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#expre_arit}.
	 * @param ctx the parse tree
	 */
	void exitExpre_arit(Swiftgrammar.Expre_aritContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#expre_valor}.
	 * @param ctx the parse tree
	 */
	void enterExpre_valor(Swiftgrammar.Expre_valorContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#expre_valor}.
	 * @param ctx the parse tree
	 */
	void exitExpre_valor(Swiftgrammar.Expre_valorContext ctx);
	/**
	 * Enter a parse tree produced by {@link Swiftgrammar#primitivo}.
	 * @param ctx the parse tree
	 */
	void enterPrimitivo(Swiftgrammar.PrimitivoContext ctx);
	/**
	 * Exit a parse tree produced by {@link Swiftgrammar#primitivo}.
	 * @param ctx the parse tree
	 */
	void exitPrimitivo(Swiftgrammar.PrimitivoContext ctx);
}