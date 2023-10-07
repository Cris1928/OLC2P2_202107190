parser grammar Swiftgrammar;

options {
    tokenVocab=Swiftlexer;
    }


@header {

import "OLC2/Compilador/interfaces"
import "OLC2/Compilador/instruccion"
import "OLC2/Compilador/expression"

import arrayList "github.com/colegno/arraylist"

}

start returns [*arrayList.List lista]:
instrucciones EOF {$lista = $instrucciones.l};


instrucciones returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : list +=instruccion+   {
        listInt := localctx.(*InstruccionesContext).GetList()
        for _, e := range listInt {
            $l.Add(e.GetInstr())
        }
    }
;

instruccion returns [interfaces.Instruction instr]:
instruccion_println  { $instr = $instruccion_println.instr               }
//|instruccion_entorno
|instruccion_structs_declaracion 
|instruccion_declaracion { $instr = $instruccion_declaracion.instr           }
|instruccion_asignacion{ $instr = $instruccion_asignacion.instr            }
|instr_structs_assignment
|instruccion_if
|instruccion_for_in
|instruccion_while
|instruccion_while_true
|instruccion_switch 
|instruccion_break
|instruccion_continue
|instruccion_func
|instruccion_llamada
|instruccion_return
|instr_structs_declaration
//|instruccion_arrays
//|instruccion_db

;

/*PRINT*/
instruccion_println returns [interfaces.Instruction instr]:
PRINT TK_PARA primitivo TK_PARC  { $instr = instruction.NewPrintln(nil, $primitivo.p,       "-1",         $PRINT.line, localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn()) }
|PRINT TK_PARA STRING TK_COMA list_expression TK_PARC { $instr = instruction.NewPrintln($list_expression.l, nil,$STRING.text[1:len($STRING.text)-1], $PRINT.line, localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn()) }
|PRINT TK_PARA expressions TK_PARC { $instr = instruction.NewPrintln(nil, $expressions.p,       "-1",         $PRINT.line, localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn()) }
;

/*ENTORNO 
instruccion_entorno:
 TK_LLAVEA instrucciones TK_LLAVEC;
*/ 
 /*DECLARACION DE VARIABLE*/
 instruccion_declaracion returns [interfaces.Instruction instr]:
 VAR ID TK_IGUAL expressions    { $instr = instruction.NewDeclaration($ID.text, interfaces.NULL,      $expressions.p, true,  $VAR.line, localctx.(*Instruccion_declaracionContext).Get_VAR().GetColumn()) }
 |VAR ID TK_DOSPUNTOS instruccion_tipo  { $instr = instruction.NewDeclaration($ID.text, $instruccion_tipo.tipo_exp,  nil,           true,  $VAR.line, localctx.(*Instruccion_declaracionContext).Get_VAR().GetColumn()) }
 |VAR ID TK_DOSPUNTOS instruccion_tipo TK_IGUAL expressions  { $instr = instruction.NewDeclaration($ID.text, $instruccion_tipo.tipo_exp, $expressions.p, true,  $VAR.line, localctx.(*Instruccion_declaracionContext).Get_VAR().GetColumn()) }
 |LET ID TK_IGUAL expressions                { $instr = instruction.NewDeclaration($ID.text, interfaces.NULL,      $expressions.p, false, $LET.line, localctx.(*Instruccion_declaracionContext).Get_LET().GetColumn()) }
 |LET ID TK_DOSPUNTOS instruccion_tipo           { $instr = instruction.NewDeclaration($ID.text, $instruccion_tipo.tipo_exp, nil,            false, $LET.line, localctx.(*Instruccion_declaracionContext).Get_LET().GetColumn()) }
 |LET ID TK_DOSPUNTOS instruccion_tipo TK_IGUAL expressions   { $instr = instruction.NewDeclaration($ID.text, $instruccion_tipo.tipo_exp, $expressions.p, false, $LET.line, localctx.(*Instruccion_declaracionContext).Get_LET().GetColumn()) }
;
 
/*ASIGNACION DE VARIABLE*/
instruccion_asignacion returns [interfaces.Instruction instr]:
ID TK_IGUAL expressions  { $instr = instruction.NewAssignment($ID.text, $expressions.p, $ID.line, localctx.(*Instruccion_asignacionContext).Get_ID().GetColumn()) }
;

/*CONTROL IF*/
instruccion_if:
IF expressions TK_LLAVEA left_instr=instrucciones TK_LLAVEC
|IF expressions TK_LLAVEA left_instr=instrucciones TK_LLAVEC ELSE TK_LLAVEA right_instr=instrucciones TK_LLAVEC
|IF expressions TK_LLAVEA left_instr=instrucciones TK_LLAVEC ELSE instr_else_if
;

instr_else_if:
e +=instruccion_if*;

/*TERNARIO IF */
instruccion_ternario:
IF cond=expressions TK_LLAVEA left_instr=expressions TK_LLAVEC  
|IF cond=expressions TK_LLAVEA left_instr=expressions TK_LLAVEC ELSE TK_LLAVEA right_instr=expressions TK_LLAVEC
|IF cond=expressions TK_LLAVEA left_instr=expressions TK_LLAVEC ELSE instr_else_if_ternario
;

instr_else_if_ternario:
e +=instruccion_ternario*;

/*CONTROL SWIFT*/
instruccion_switch:
SWITCH expressions TK_LLAVEA list_case block_default TK_LLAVEC
|SWITCH expressions TK_LLAVEA block_default TK_LLAVEC
;

list_case:
e+=instruccion_case+;

instruccion_case:
list_expre_case TK_DOSPUNTOS TK_LLAVEA instrucciones TK_LLAVEC
|list_expre_case TK_DOSPUNTOS block_instr_switch TK_COMA
;

list_expre_case:
e+=block_case+;

block_case:
expressions TK_BARRA
|expressions
;

block_instr_switch:
list += instruccion;

/*DEFAULT*/
instr_default:
CASE ID TK_DOSPUNTOS TK_LLAVEA instrucciones TK_LLAVEC
|CASE ID TK_DOSPUNTOS block_instr_switch
;

block_default:
e += instr_default+;

/*CONTROL SWITCH TERNARIO*/
instruccion_switch_ternario:
SWITCH expressions TK_LLAVEA list_case_ternario instr_default_ter TK_LLAVEC
|SWITCH expressions TK_LLAVEA instr_default_ter TK_LLAVEC
;


list_case_ternario:
e+=instr_case_ter+;

/*CASE */
instr_case_ter:
list_expre_case_ter TK_DOSPUNTOS expressions 
|list_expre_case_ter TK_DOSPUNTOS TK_LLAVEA expressions TK_LLAVEC
;

/* List Expression Case */
list_expre_case_ter:
e+=block_case_ter+;

block_case_ter:
expressions TK_BARRA
|expressions
;


/*DEFAULT*/

instr_default_ter:
CASE ID TK_DOSPUNTOS expressions
|CASE ID TK_DOSPUNTOS TK_LLAVEA expressions TK_LLAVEC
;

/*WHILE*/
instruccion_while:
WHILE expressions TK_LLAVEA instrucciones TK_LLAVEC;
/*FOR */
instruccion_for_in:
FOR ID IN left=expressions TK_TRIPLEPUNTO right=expressions TK_LLAVEA instrucciones TK_LLAVEC
|FOR ID IN left=expressions TK_LLAVEA instrucciones TK_LLAVEC
;
/*wHILE TRUE */
instruccion_while_true:
WHILE TRUE TK_LLAVEA instrucciones TK_LLAVEC
;

/*wHILE TRUE TERNARIO*/
instruccion_while_true_ternario:
WHILE TRUE TK_LLAVEA instrucciones TK_LLAVEC
;

/*BREAK*/
instruccion_break:
BREAK
|BREAK expressions
;

/*CONTINUE*/
instruccion_continue:
CONTINUE    
;

/*RETURN */
instruccion_return:
RETURN expressions
;

/*FUNCIONES*/
instruccion_func:
FUNC ID TK_PARA TK_PARC TK_LLAVEA instrucciones TK_LLAVEC
|FUNC ID TK_PARA TK_PARC TK_MENOSMAYOR instruccion_tipo TK_LLAVEA instrucciones TK_LLAVEC
|FUNC ID TK_PARA list_function_parameters TK_PARC TK_LLAVEA instrucciones TK_LLAVEC
|FUNC ID TK_PARA list_function_parameters TK_PARC TK_MENOSMAYOR instruccion_tipo TK_LLAVEA instrucciones TK_LLAVEC
;

list_function_parameters:
e+=block_parameters_fn+;

block_parameters_fn:
ID TK_DOSPUNTOS instruccion_tipo TK_COMA
|ID TK_DOSPUNTOS instruccion_tipo
;

/* LLAMADA FUNCION */
instruccion_llamada:
ID TK_PARA TK_PARC
|ID TK_PARA list_expression TK_PARC
;

instr_llamada_expre:
ID TK_PARA TK_PARC
|ID TK_PARA list_expression TK_PARC
;

/*STRUCTS DECLARATION*/
instruccion_structs_declaracion:
STRUCT ID TK_LLAVEA list_struct_parameters TK_LLAVEC
;

list_struct_parameters:
e+=block_structs_parameters+
;

block_structs_parameters:
ID TK_DOSPUNTOS instruccion_tipo TK_COMA
|ID TK_DOSPUNTOS instruccion_tipo
;

/*STRUCT DECLARACION */
/*
instr_arrays:
VAR TK_DOSPUNTOS list_arrays_definition TK_IGUAL list_arrays_datos
|LET ID TK_DOSPUNTOS list_arrays_definition TK_IGUAL list_arrays_datos
;

list_arrays_datos:
   e +=  block_dimensiones_datos
  ;

block_dimensiones_datos: 
block_array_dimensionUno_datos
;

block_array_dimensionUno_datos:
TK_CORA list_array_dimUno TK_CORC
;

list_array_dimUno:
e+= block_array_dimUno_datos+
;

block_array_dimUno_datos:
expressions TK_COMA 
|expressions
;
list_arrays_definition: 
e += block_dimensiones+
;


block_dimensiones:
block_array_dimensionUno
;

block_array_dimensionUno:
TK_CORA instruccion_tipo ; */


/*STRUCTS DECLARACION*/
instr_arrays_identifier:
ID list_arrays_parameters_id;

list_arrays_parameters_id:
e+= block_arrays_identifier+ 
;


block_arrays_identifier:
TK_CORA expressions TK_CORC
;

/*STRUCTS DECLARACION*/
instr_structs_declaration:
VAR left=ID TK_IGUAL right=ID TK_LLAVEA list_struct_parameters_decla TK_LLAVEC  
|LET left=ID TK_IGUAL right=ID TK_LLAVEA list_struct_parameters_decla TK_LLAVEC
;

list_struct_parameters_decla:
e += block_structs_parameters_decla+
;

block_structs_parameters_decla:
ID TK_DOSPUNTOS expressions 
|ID TK_DOSPUNTOS expressions TK_COMA
;

/*STRUCT DECLARACION */
instr_structs_identifier:
ID list_struct_parameters_id
;

list_struct_parameters_id:
e+= block_structs_identifier+
;

block_structs_identifier:
TK_PUNTO ID
;
/*STRUCT ASIGNACION*/
instr_structs_assignment:
ID list_struct_parameters_id TK_IGUAL expressions 
;

 /*TIPO */
instruccion_tipo  returns [interfaces.TypeExpression tipo_exp]:
R_INT   {$tipo_exp = interfaces.INTEGER}
|R_FLOAT  {$tipo_exp = interfaces.FLOAT}
|R_STRING  {$tipo_exp = interfaces.STRING}
|R_BOOL {$tipo_exp = interfaces.BOOLEAN}
|R_CHAR {$tipo_exp = interfaces.CHAR}
//|TK_AMPERSAND ID
;

/* List Expression Case */

list_expression returns [*arrayList.List l]
  @init{
    $l =  arrayList.New()
  }
  : e += block_list_expression+  {
        listInt := localctx.(*List_expressionContext).GetE()
        for _, e := range listInt {
            $l.Add(e.GetP())
        }
    }
;

block_list_expression  returns [interfaces.Expression p]:
expressions TK_COMA  { $p =  instruction.NewListExpre($expressions.p) }
|expressions     { $p =  instruction.NewListExpre($expressions.p) }
;

/*EXPRESIONES*/
expressions returns [interfaces.Expression p]:
expre_log         { $p = $expre_log.p }
//|expre_casteo
|expre_arit                   { $p = $expre_arit.p }
|expre_rel                    { $p = $expre_rel.p }
;

expre_log returns [interfaces.Expression p]:
op='!' left=expre_log                                                { $p = expression.NewOperacion($left.p, $op.text, nil,      true,  $op.line, localctx.(*Expre_logContext).GetOp().GetColumn()) }
|left=expre_log op=('&&'|'||') right=expre_log                       { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_logContext).GetOp().GetColumn()) }
|expre_rel                                                           { $p = $expre_rel.p }
;

expre_rel returns [interfaces.Expression p]:
left=expre_rel op=('<'|'<='|'>='|'>'|'!='|'==') right=expre_rel       { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_relContext).GetOp().GetColumn()) }
| expre_arit                                                          { $p = $expre_arit.p }
;

expre_arit  returns [interfaces.Expression p]:
op='-' left=expre_arit                                    { $p = expression.NewOperacion($left.p, $op.text, nil,      true,  $op.line, localctx.(*Expre_aritContext).GetOp().GetColumn()) }
| left=expre_arit op=('*'|'/'|'%') right=expre_arit       { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_aritContext).GetOp().GetColumn()) }
| left=expre_arit op=('+'|'-') right=expre_arit           { $p = expression.NewOperacion($left.p, $op.text, $right.p, false, $op.line, localctx.(*Expre_aritContext).GetOp().GetColumn()) }
| expre_valor                                             { $p = $expre_valor.p }
| TK_PARA expressions TK_PARC                             { $p = $expressions.p }
;

expre_valor returns [interfaces.Expression p]:
 primitivo                                           { $p = $primitivo.p }
 ;

 primitivo returns[interfaces.Expression p] :
 NUMBER{
              num,err := strconv.Atoi($NUMBER.text)
                if err!= nil{
                    fmt.Println(err)
                }

            $p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.NULL, $NUMBER.line, localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())
       } 
 |DOUBLE {  
                num,err := strconv.ParseFloat($DOUBLE.text, 64)
                if err!= nil{
                    fmt.Println(err)
                }
            $p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.NULL, $DOUBLE.line, localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())
              }
 |STRING 
 { 
              str:= $STRING.text[1:len($STRING.text)-1]
              $p = expression.NewPrimitivo(str, interfaces.STRING, interfaces.NULL, $STRING.line, localctx.(*PrimitivoContext).Get_STRING().GetColumn())
            
            }
 |BOOLEAN { 
          
              exp,_ := strconv.ParseBool($BOOLEAN.text)
              $p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.NULL, $BOOLEAN.line, localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())
            }
 |CHAR   {

            str:= $CHAR.text[1]
            $p = expression.NewPrimitivo(string(str), interfaces.CHAR, interfaces.NULL, $CHAR.line, localctx.(*PrimitivoContext).Get_CHAR().GetColumn())
          
          }
 | instr_llamada_expre 
 | instr_structs_identifier
 | instr_arrays_identifier 
 | ID  { $p = instruction.NewIdentifier($ID.text, $ID.line, localctx.(*PrimitivoContext).Get_ID().GetColumn()) }
 //| nativa_expre  
// | primitivo_casteo 
 | instruccion_ternario 
 | instruccion_switch_ternario 
 ///| instr_loop_ternario 
 ;

 