lexer grammar Swiftlexer;

PRINT: 'print' ;
/* Declaracion */
VAR: 'var' ;
LET: 'let' ;
/* Control */
IF: 'if' ;
ELSE: 'else' ;
/* Ciclos */
FOR: 'for' ;
WHILE: 'while' ;
SWITCH: 'switch' ;
CASE: 'case' ;
DEFAULT: 'default' ;
IN: 'in' ;
/* Transferencia de Control */
BREAK: 'break' ;
CONTINUE: 'continue' ;
RETURN: 'return' ;

/* Funciones */
FUNC: 'func' ;
/* STRUCT */
STRUCT: 'struct' ;
/* TIPOS */
R_INT: 'Int' ;
R_FLOAT: 'Float' ;
R_DOUBLE: 'Double' ;
R_BOOL: 'Bool' ;
R_CHAR: 'Character' ;
R_STRING: 'String' ;

TRUE: 'true' ;

NUMBER: [0-9]+;
DOUBLE: [0-9]+'.'[0-9]+;
CHAR:   '\''~["]'\'';
STRING: '"'~["]*'"';
BOOLEAN: ('true'|'false');
ID: ([a-zA-Z_])[a-zA-Z0-9_]*;

TK_TRIPLEPUNTO:   '...';
TK_PUNTO:        '.';
TK_PUNTOCOMA:    ';';
TK_COMA:         ',';
TK_DOSPUNTOS:    ':';
TK_IGUAL:        '=';
TK_IGUALIGUAL:   '==';
TK_MAYORIGUAL:   '>=';
TK_IGUALMAYOR:   '=>';
TK_MENOSMAYOR:   '->';
TK_MENORIGUAL:   '<=';
TK_DIFIGUAL:     '!=';
TK_MAYOR:        '>';
TK_MENOR:        '<';
TK_MULT:         '*';
TK_DIV:          '/';
TK_MODULO:       '%';
TK_MAS:          '+';
TK_MENOS:        '-';
TK_PARA:         '(';
TK_PARC:         ')';
TK_LLAVEA:       '{';
TK_LLAVEC:       '}';
TK_CORA:         '[';
TK_CORC:         ']';
TK_AND:          '&&';
TK_AMPERSAND:    '&';
TK_OR:           '||';
TK_BARRA:        '|';
TK_NOT:          '!';


WHITESPACE: [ \r\n\t]+ -> skip;
TK_MULTI:    '/*' (~[/])+ '*/' -> skip;
TK_LINE:    '//'(~[\n])+ -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;