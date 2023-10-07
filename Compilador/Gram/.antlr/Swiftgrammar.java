// Generated from c:/OLC2P2_202107190/Compilador/Gram/Swiftgrammar.g4 by ANTLR 4.13.1


import "OLC2/Compilador/interfaces"
import "OLC2/Compilador/instruccion"
import "OLC2/Compilador/expression"

import arrayList "github.com/colegno/arraylist"


import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Swiftgrammar extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PRINT=1, VAR=2, LET=3, IF=4, ELSE=5, FOR=6, WHILE=7, SWITCH=8, CASE=9, 
		DEFAULT=10, IN=11, BREAK=12, CONTINUE=13, RETURN=14, FUNC=15, STRUCT=16, 
		R_INT=17, R_FLOAT=18, R_DOUBLE=19, R_BOOL=20, R_CHAR=21, R_STRING=22, 
		TRUE=23, NUMBER=24, DOUBLE=25, CHAR=26, STRING=27, BOOLEAN=28, ID=29, 
		TK_TRIPLEPUNTO=30, TK_PUNTO=31, TK_PUNTOCOMA=32, TK_COMA=33, TK_DOSPUNTOS=34, 
		TK_IGUAL=35, TK_IGUALIGUAL=36, TK_MAYORIGUAL=37, TK_IGUALMAYOR=38, TK_MENOSMAYOR=39, 
		TK_MENORIGUAL=40, TK_DIFIGUAL=41, TK_MAYOR=42, TK_MENOR=43, TK_MULT=44, 
		TK_DIV=45, TK_MODULO=46, TK_MAS=47, TK_MENOS=48, TK_PARA=49, TK_PARC=50, 
		TK_LLAVEA=51, TK_LLAVEC=52, TK_CORA=53, TK_CORC=54, TK_AND=55, TK_AMPERSAND=56, 
		TK_OR=57, TK_BARRA=58, TK_NOT=59, WHITESPACE=60, TK_MULTI=61, TK_LINE=62;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_instruccion_println = 3, 
		RULE_instruccion_declaracion = 4, RULE_instruccion_asignacion = 5, RULE_instruccion_if = 6, 
		RULE_instr_else_if = 7, RULE_instruccion_ternario = 8, RULE_instr_else_if_ternario = 9, 
		RULE_instruccion_switch = 10, RULE_list_case = 11, RULE_instruccion_case = 12, 
		RULE_list_expre_case = 13, RULE_block_case = 14, RULE_block_instr_switch = 15, 
		RULE_instr_default = 16, RULE_block_default = 17, RULE_instruccion_switch_ternario = 18, 
		RULE_list_case_ternario = 19, RULE_instr_case_ter = 20, RULE_list_expre_case_ter = 21, 
		RULE_block_case_ter = 22, RULE_instr_default_ter = 23, RULE_instruccion_while = 24, 
		RULE_instruccion_for_in = 25, RULE_instruccion_while_true = 26, RULE_instruccion_while_true_ternario = 27, 
		RULE_instruccion_break = 28, RULE_instruccion_continue = 29, RULE_instruccion_return = 30, 
		RULE_instruccion_func = 31, RULE_list_function_parameters = 32, RULE_block_parameters_fn = 33, 
		RULE_instruccion_llamada = 34, RULE_instr_llamada_expre = 35, RULE_instruccion_structs_declaracion = 36, 
		RULE_list_struct_parameters = 37, RULE_block_structs_parameters = 38, 
		RULE_instr_arrays_identifier = 39, RULE_list_arrays_parameters_id = 40, 
		RULE_block_arrays_identifier = 41, RULE_instr_structs_declaration = 42, 
		RULE_list_struct_parameters_decla = 43, RULE_block_structs_parameters_decla = 44, 
		RULE_instr_structs_identifier = 45, RULE_list_struct_parameters_id = 46, 
		RULE_block_structs_identifier = 47, RULE_instr_structs_assignment = 48, 
		RULE_instruccion_tipo = 49, RULE_list_expression = 50, RULE_block_list_expression = 51, 
		RULE_expressions = 52, RULE_expre_log = 53, RULE_expre_rel = 54, RULE_expre_arit = 55, 
		RULE_expre_valor = 56, RULE_primitivo = 57;
	private static String[] makeRuleNames() {
		return new String[] {
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
			"primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'print'", "'var'", "'let'", "'if'", "'else'", "'for'", "'while'", 
			"'switch'", "'case'", "'default'", "'in'", "'break'", "'continue'", "'return'", 
			"'func'", "'struct'", "'Int'", "'Float'", "'Double'", "'Bool'", "'Character'", 
			"'String'", "'true'", null, null, null, null, null, null, "'...'", "'.'", 
			"';'", "','", "':'", "'='", "'=='", "'>='", "'=>'", "'->'", "'<='", "'!='", 
			"'>'", "'<'", "'*'", "'/'", "'%'", "'+'", "'-'", "'('", "')'", "'{'", 
			"'}'", "'['", "']'", "'&&'", "'&'", "'||'", "'|'", "'!'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINT", "VAR", "LET", "IF", "ELSE", "FOR", "WHILE", "SWITCH", 
			"CASE", "DEFAULT", "IN", "BREAK", "CONTINUE", "RETURN", "FUNC", "STRUCT", 
			"R_INT", "R_FLOAT", "R_DOUBLE", "R_BOOL", "R_CHAR", "R_STRING", "TRUE", 
			"NUMBER", "DOUBLE", "CHAR", "STRING", "BOOLEAN", "ID", "TK_TRIPLEPUNTO", 
			"TK_PUNTO", "TK_PUNTOCOMA", "TK_COMA", "TK_DOSPUNTOS", "TK_IGUAL", "TK_IGUALIGUAL", 
			"TK_MAYORIGUAL", "TK_IGUALMAYOR", "TK_MENOSMAYOR", "TK_MENORIGUAL", "TK_DIFIGUAL", 
			"TK_MAYOR", "TK_MENOR", "TK_MULT", "TK_DIV", "TK_MODULO", "TK_MAS", "TK_MENOS", 
			"TK_PARA", "TK_PARC", "TK_LLAVEA", "TK_LLAVEC", "TK_CORA", "TK_CORC", 
			"TK_AND", "TK_AMPERSAND", "TK_OR", "TK_BARRA", "TK_NOT", "WHITESPACE", 
			"TK_MULTI", "TK_LINE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Swiftgrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Swiftgrammar(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StartContext extends ParserRuleContext {
		public *arrayList.List lista;
		public InstruccionesContext instrucciones;
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode EOF() { return getToken(Swiftgrammar.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			((StartContext)_localctx).instrucciones = instrucciones();
			setState(117);
			match(EOF);
			_localctx.lista = ((StartContext)_localctx).instrucciones.l
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstruccionesContext extends ParserRuleContext {
		public *arrayList.List l;
		public InstruccionContext instruccion;
		public List<InstruccionContext> list = new ArrayList<InstruccionContext>();
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucciones);

		    _localctx.l =  arrayList.New()
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(121); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(120);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).list.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(123); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 536998366L) != 0) );

			        listInt := localctx.(*InstruccionesContext).GetList()
			        for _, e := range listInt {
			            _localctx.l.Add(e.GetInstr())
			        }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstruccionContext extends ParserRuleContext {
		public interfaces.Instruction instr;
		public Instruccion_printlnContext instruccion_println;
		public Instruccion_declaracionContext instruccion_declaracion;
		public Instruccion_asignacionContext instruccion_asignacion;
		public Instruccion_printlnContext instruccion_println() {
			return getRuleContext(Instruccion_printlnContext.class,0);
		}
		public Instruccion_structs_declaracionContext instruccion_structs_declaracion() {
			return getRuleContext(Instruccion_structs_declaracionContext.class,0);
		}
		public Instruccion_declaracionContext instruccion_declaracion() {
			return getRuleContext(Instruccion_declaracionContext.class,0);
		}
		public Instruccion_asignacionContext instruccion_asignacion() {
			return getRuleContext(Instruccion_asignacionContext.class,0);
		}
		public Instr_structs_assignmentContext instr_structs_assignment() {
			return getRuleContext(Instr_structs_assignmentContext.class,0);
		}
		public Instruccion_ifContext instruccion_if() {
			return getRuleContext(Instruccion_ifContext.class,0);
		}
		public Instruccion_for_inContext instruccion_for_in() {
			return getRuleContext(Instruccion_for_inContext.class,0);
		}
		public Instruccion_whileContext instruccion_while() {
			return getRuleContext(Instruccion_whileContext.class,0);
		}
		public Instruccion_while_trueContext instruccion_while_true() {
			return getRuleContext(Instruccion_while_trueContext.class,0);
		}
		public Instruccion_switchContext instruccion_switch() {
			return getRuleContext(Instruccion_switchContext.class,0);
		}
		public Instruccion_breakContext instruccion_break() {
			return getRuleContext(Instruccion_breakContext.class,0);
		}
		public Instruccion_continueContext instruccion_continue() {
			return getRuleContext(Instruccion_continueContext.class,0);
		}
		public Instruccion_funcContext instruccion_func() {
			return getRuleContext(Instruccion_funcContext.class,0);
		}
		public Instruccion_llamadaContext instruccion_llamada() {
			return getRuleContext(Instruccion_llamadaContext.class,0);
		}
		public Instruccion_returnContext instruccion_return() {
			return getRuleContext(Instruccion_returnContext.class,0);
		}
		public Instr_structs_declarationContext instr_structs_declaration() {
			return getRuleContext(Instr_structs_declarationContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			setState(149);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(127);
				((InstruccionContext)_localctx).instruccion_println = instruccion_println();
				 _localctx.instr = ((InstruccionContext)_localctx).instruccion_println.instr               
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(130);
				instruccion_structs_declaracion();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(131);
				((InstruccionContext)_localctx).instruccion_declaracion = instruccion_declaracion();
				 _localctx.instr = ((InstruccionContext)_localctx).instruccion_declaracion.instr           
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(134);
				((InstruccionContext)_localctx).instruccion_asignacion = instruccion_asignacion();
				 _localctx.instr = ((InstruccionContext)_localctx).instruccion_asignacion.instr            
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(137);
				instr_structs_assignment();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(138);
				instruccion_if();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(139);
				instruccion_for_in();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(140);
				instruccion_while();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(141);
				instruccion_while_true();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(142);
				instruccion_switch();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(143);
				instruccion_break();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(144);
				instruccion_continue();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(145);
				instruccion_func();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(146);
				instruccion_llamada();
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(147);
				instruccion_return();
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(148);
				instr_structs_declaration();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_printlnContext extends ParserRuleContext {
		public interfaces.Instruction instr;
		public Token PRINT;
		public PrimitivoContext primitivo;
		public Token STRING;
		public List_expressionContext list_expression;
		public ExpressionsContext expressions;
		public TerminalNode PRINT() { return getToken(Swiftgrammar.PRINT, 0); }
		public TerminalNode TK_PARA() { return getToken(Swiftgrammar.TK_PARA, 0); }
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
		}
		public TerminalNode TK_PARC() { return getToken(Swiftgrammar.TK_PARC, 0); }
		public TerminalNode STRING() { return getToken(Swiftgrammar.STRING, 0); }
		public TerminalNode TK_COMA() { return getToken(Swiftgrammar.TK_COMA, 0); }
		public List_expressionContext list_expression() {
			return getRuleContext(List_expressionContext.class,0);
		}
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public Instruccion_printlnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_println; }
	}

	public final Instruccion_printlnContext instruccion_println() throws RecognitionException {
		Instruccion_printlnContext _localctx = new Instruccion_printlnContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_instruccion_println);
		try {
			setState(171);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(151);
				((Instruccion_printlnContext)_localctx).PRINT = match(PRINT);
				setState(152);
				match(TK_PARA);
				setState(153);
				((Instruccion_printlnContext)_localctx).primitivo = primitivo();
				setState(154);
				match(TK_PARC);
				 _localctx.instr = instruction.NewPrintln(nil, ((Instruccion_printlnContext)_localctx).primitivo.p,       "-1",         (((Instruccion_printlnContext)_localctx).PRINT!=null?((Instruccion_printlnContext)_localctx).PRINT.getLine():0), localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn()) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(157);
				((Instruccion_printlnContext)_localctx).PRINT = match(PRINT);
				setState(158);
				match(TK_PARA);
				setState(159);
				((Instruccion_printlnContext)_localctx).STRING = match(STRING);
				setState(160);
				match(TK_COMA);
				setState(161);
				((Instruccion_printlnContext)_localctx).list_expression = list_expression();
				setState(162);
				match(TK_PARC);
				 _localctx.instr = instruction.NewPrintln(((Instruccion_printlnContext)_localctx).list_expression.l, nil,(((Instruccion_printlnContext)_localctx).STRING!=null?((Instruccion_printlnContext)_localctx).STRING.getText():null)[1:len((((Instruccion_printlnContext)_localctx).STRING!=null?((Instruccion_printlnContext)_localctx).STRING.getText():null))-1], (((Instruccion_printlnContext)_localctx).PRINT!=null?((Instruccion_printlnContext)_localctx).PRINT.getLine():0), localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn()) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(165);
				((Instruccion_printlnContext)_localctx).PRINT = match(PRINT);
				setState(166);
				match(TK_PARA);
				setState(167);
				((Instruccion_printlnContext)_localctx).expressions = expressions();
				setState(168);
				match(TK_PARC);
				 _localctx.instr = instruction.NewPrintln(nil, ((Instruccion_printlnContext)_localctx).expressions.p,       "-1",         (((Instruccion_printlnContext)_localctx).PRINT!=null?((Instruccion_printlnContext)_localctx).PRINT.getLine():0), localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn()) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_declaracionContext extends ParserRuleContext {
		public interfaces.Instruction instr;
		public Token VAR;
		public Token ID;
		public ExpressionsContext expressions;
		public Instruccion_tipoContext instruccion_tipo;
		public Token LET;
		public TerminalNode VAR() { return getToken(Swiftgrammar.VAR, 0); }
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_IGUAL() { return getToken(Swiftgrammar.TK_IGUAL, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_DOSPUNTOS() { return getToken(Swiftgrammar.TK_DOSPUNTOS, 0); }
		public Instruccion_tipoContext instruccion_tipo() {
			return getRuleContext(Instruccion_tipoContext.class,0);
		}
		public TerminalNode LET() { return getToken(Swiftgrammar.LET, 0); }
		public Instruccion_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_declaracion; }
	}

	public final Instruccion_declaracionContext instruccion_declaracion() throws RecognitionException {
		Instruccion_declaracionContext _localctx = new Instruccion_declaracionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_instruccion_declaracion);
		try {
			setState(213);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(173);
				((Instruccion_declaracionContext)_localctx).VAR = match(VAR);
				setState(174);
				((Instruccion_declaracionContext)_localctx).ID = match(ID);
				setState(175);
				match(TK_IGUAL);
				setState(176);
				((Instruccion_declaracionContext)_localctx).expressions = expressions();
				 _localctx.instr = instruction.NewDeclaration((((Instruccion_declaracionContext)_localctx).ID!=null?((Instruccion_declaracionContext)_localctx).ID.getText():null), interfaces.NULL,      ((Instruccion_declaracionContext)_localctx).expressions.p, true,  (((Instruccion_declaracionContext)_localctx).VAR!=null?((Instruccion_declaracionContext)_localctx).VAR.getLine():0), localctx.(*Instruccion_declaracionContext).Get_VAR().GetColumn()) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(179);
				((Instruccion_declaracionContext)_localctx).VAR = match(VAR);
				setState(180);
				((Instruccion_declaracionContext)_localctx).ID = match(ID);
				setState(181);
				match(TK_DOSPUNTOS);
				setState(182);
				((Instruccion_declaracionContext)_localctx).instruccion_tipo = instruccion_tipo();
				 _localctx.instr = instruction.NewDeclaration((((Instruccion_declaracionContext)_localctx).ID!=null?((Instruccion_declaracionContext)_localctx).ID.getText():null), ((Instruccion_declaracionContext)_localctx).instruccion_tipo.tipo_exp,  nil,           true,  (((Instruccion_declaracionContext)_localctx).VAR!=null?((Instruccion_declaracionContext)_localctx).VAR.getLine():0), localctx.(*Instruccion_declaracionContext).Get_VAR().GetColumn()) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(185);
				((Instruccion_declaracionContext)_localctx).VAR = match(VAR);
				setState(186);
				((Instruccion_declaracionContext)_localctx).ID = match(ID);
				setState(187);
				match(TK_DOSPUNTOS);
				setState(188);
				((Instruccion_declaracionContext)_localctx).instruccion_tipo = instruccion_tipo();
				setState(189);
				match(TK_IGUAL);
				setState(190);
				((Instruccion_declaracionContext)_localctx).expressions = expressions();
				 _localctx.instr = instruction.NewDeclaration((((Instruccion_declaracionContext)_localctx).ID!=null?((Instruccion_declaracionContext)_localctx).ID.getText():null), ((Instruccion_declaracionContext)_localctx).instruccion_tipo.tipo_exp, ((Instruccion_declaracionContext)_localctx).expressions.p, true,  (((Instruccion_declaracionContext)_localctx).VAR!=null?((Instruccion_declaracionContext)_localctx).VAR.getLine():0), localctx.(*Instruccion_declaracionContext).Get_VAR().GetColumn()) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(193);
				((Instruccion_declaracionContext)_localctx).LET = match(LET);
				setState(194);
				((Instruccion_declaracionContext)_localctx).ID = match(ID);
				setState(195);
				match(TK_IGUAL);
				setState(196);
				((Instruccion_declaracionContext)_localctx).expressions = expressions();
				 _localctx.instr = instruction.NewDeclaration((((Instruccion_declaracionContext)_localctx).ID!=null?((Instruccion_declaracionContext)_localctx).ID.getText():null), interfaces.NULL,      ((Instruccion_declaracionContext)_localctx).expressions.p, false, (((Instruccion_declaracionContext)_localctx).LET!=null?((Instruccion_declaracionContext)_localctx).LET.getLine():0), localctx.(*Instruccion_declaracionContext).Get_LET().GetColumn()) 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(199);
				((Instruccion_declaracionContext)_localctx).LET = match(LET);
				setState(200);
				((Instruccion_declaracionContext)_localctx).ID = match(ID);
				setState(201);
				match(TK_DOSPUNTOS);
				setState(202);
				((Instruccion_declaracionContext)_localctx).instruccion_tipo = instruccion_tipo();
				 _localctx.instr = instruction.NewDeclaration((((Instruccion_declaracionContext)_localctx).ID!=null?((Instruccion_declaracionContext)_localctx).ID.getText():null), ((Instruccion_declaracionContext)_localctx).instruccion_tipo.tipo_exp, nil,            false, (((Instruccion_declaracionContext)_localctx).LET!=null?((Instruccion_declaracionContext)_localctx).LET.getLine():0), localctx.(*Instruccion_declaracionContext).Get_LET().GetColumn()) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(205);
				((Instruccion_declaracionContext)_localctx).LET = match(LET);
				setState(206);
				((Instruccion_declaracionContext)_localctx).ID = match(ID);
				setState(207);
				match(TK_DOSPUNTOS);
				setState(208);
				((Instruccion_declaracionContext)_localctx).instruccion_tipo = instruccion_tipo();
				setState(209);
				match(TK_IGUAL);
				setState(210);
				((Instruccion_declaracionContext)_localctx).expressions = expressions();
				 _localctx.instr = instruction.NewDeclaration((((Instruccion_declaracionContext)_localctx).ID!=null?((Instruccion_declaracionContext)_localctx).ID.getText():null), ((Instruccion_declaracionContext)_localctx).instruccion_tipo.tipo_exp, ((Instruccion_declaracionContext)_localctx).expressions.p, false, (((Instruccion_declaracionContext)_localctx).LET!=null?((Instruccion_declaracionContext)_localctx).LET.getLine():0), localctx.(*Instruccion_declaracionContext).Get_LET().GetColumn()) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_asignacionContext extends ParserRuleContext {
		public interfaces.Instruction instr;
		public Token ID;
		public ExpressionsContext expressions;
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_IGUAL() { return getToken(Swiftgrammar.TK_IGUAL, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public Instruccion_asignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_asignacion; }
	}

	public final Instruccion_asignacionContext instruccion_asignacion() throws RecognitionException {
		Instruccion_asignacionContext _localctx = new Instruccion_asignacionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_instruccion_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
			((Instruccion_asignacionContext)_localctx).ID = match(ID);
			setState(216);
			match(TK_IGUAL);
			setState(217);
			((Instruccion_asignacionContext)_localctx).expressions = expressions();
			 _localctx.instr = instruction.NewAssignment((((Instruccion_asignacionContext)_localctx).ID!=null?((Instruccion_asignacionContext)_localctx).ID.getText():null), ((Instruccion_asignacionContext)_localctx).expressions.p, (((Instruccion_asignacionContext)_localctx).ID!=null?((Instruccion_asignacionContext)_localctx).ID.getLine():0), localctx.(*Instruccion_asignacionContext).Get_ID().GetColumn()) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_ifContext extends ParserRuleContext {
		public InstruccionesContext left_instr;
		public InstruccionesContext right_instr;
		public TerminalNode IF() { return getToken(Swiftgrammar.IF, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public List<TerminalNode> TK_LLAVEA() { return getTokens(Swiftgrammar.TK_LLAVEA); }
		public TerminalNode TK_LLAVEA(int i) {
			return getToken(Swiftgrammar.TK_LLAVEA, i);
		}
		public List<TerminalNode> TK_LLAVEC() { return getTokens(Swiftgrammar.TK_LLAVEC); }
		public TerminalNode TK_LLAVEC(int i) {
			return getToken(Swiftgrammar.TK_LLAVEC, i);
		}
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(Swiftgrammar.ELSE, 0); }
		public Instr_else_ifContext instr_else_if() {
			return getRuleContext(Instr_else_ifContext.class,0);
		}
		public Instruccion_ifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_if; }
	}

	public final Instruccion_ifContext instruccion_if() throws RecognitionException {
		Instruccion_ifContext _localctx = new Instruccion_ifContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_instruccion_if);
		try {
			setState(244);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(220);
				match(IF);
				setState(221);
				expressions();
				setState(222);
				match(TK_LLAVEA);
				setState(223);
				((Instruccion_ifContext)_localctx).left_instr = instrucciones();
				setState(224);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(226);
				match(IF);
				setState(227);
				expressions();
				setState(228);
				match(TK_LLAVEA);
				setState(229);
				((Instruccion_ifContext)_localctx).left_instr = instrucciones();
				setState(230);
				match(TK_LLAVEC);
				setState(231);
				match(ELSE);
				setState(232);
				match(TK_LLAVEA);
				setState(233);
				((Instruccion_ifContext)_localctx).right_instr = instrucciones();
				setState(234);
				match(TK_LLAVEC);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(236);
				match(IF);
				setState(237);
				expressions();
				setState(238);
				match(TK_LLAVEA);
				setState(239);
				((Instruccion_ifContext)_localctx).left_instr = instrucciones();
				setState(240);
				match(TK_LLAVEC);
				setState(241);
				match(ELSE);
				setState(242);
				instr_else_if();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_else_ifContext extends ParserRuleContext {
		public Instruccion_ifContext instruccion_if;
		public List<Instruccion_ifContext> e = new ArrayList<Instruccion_ifContext>();
		public List<Instruccion_ifContext> instruccion_if() {
			return getRuleContexts(Instruccion_ifContext.class);
		}
		public Instruccion_ifContext instruccion_if(int i) {
			return getRuleContext(Instruccion_ifContext.class,i);
		}
		public Instr_else_ifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_else_if; }
	}

	public final Instr_else_ifContext instr_else_if() throws RecognitionException {
		Instr_else_ifContext _localctx = new Instr_else_ifContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_instr_else_if);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(249);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(246);
					((Instr_else_ifContext)_localctx).instruccion_if = instruccion_if();
					((Instr_else_ifContext)_localctx).e.add(((Instr_else_ifContext)_localctx).instruccion_if);
					}
					} 
				}
				setState(251);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_ternarioContext extends ParserRuleContext {
		public ExpressionsContext cond;
		public ExpressionsContext left_instr;
		public ExpressionsContext right_instr;
		public TerminalNode IF() { return getToken(Swiftgrammar.IF, 0); }
		public List<TerminalNode> TK_LLAVEA() { return getTokens(Swiftgrammar.TK_LLAVEA); }
		public TerminalNode TK_LLAVEA(int i) {
			return getToken(Swiftgrammar.TK_LLAVEA, i);
		}
		public List<TerminalNode> TK_LLAVEC() { return getTokens(Swiftgrammar.TK_LLAVEC); }
		public TerminalNode TK_LLAVEC(int i) {
			return getToken(Swiftgrammar.TK_LLAVEC, i);
		}
		public List<ExpressionsContext> expressions() {
			return getRuleContexts(ExpressionsContext.class);
		}
		public ExpressionsContext expressions(int i) {
			return getRuleContext(ExpressionsContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(Swiftgrammar.ELSE, 0); }
		public Instr_else_if_ternarioContext instr_else_if_ternario() {
			return getRuleContext(Instr_else_if_ternarioContext.class,0);
		}
		public Instruccion_ternarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_ternario; }
	}

	public final Instruccion_ternarioContext instruccion_ternario() throws RecognitionException {
		Instruccion_ternarioContext _localctx = new Instruccion_ternarioContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_instruccion_ternario);
		try {
			setState(276);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(252);
				match(IF);
				setState(253);
				((Instruccion_ternarioContext)_localctx).cond = expressions();
				setState(254);
				match(TK_LLAVEA);
				setState(255);
				((Instruccion_ternarioContext)_localctx).left_instr = expressions();
				setState(256);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(258);
				match(IF);
				setState(259);
				((Instruccion_ternarioContext)_localctx).cond = expressions();
				setState(260);
				match(TK_LLAVEA);
				setState(261);
				((Instruccion_ternarioContext)_localctx).left_instr = expressions();
				setState(262);
				match(TK_LLAVEC);
				setState(263);
				match(ELSE);
				setState(264);
				match(TK_LLAVEA);
				setState(265);
				((Instruccion_ternarioContext)_localctx).right_instr = expressions();
				setState(266);
				match(TK_LLAVEC);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(268);
				match(IF);
				setState(269);
				((Instruccion_ternarioContext)_localctx).cond = expressions();
				setState(270);
				match(TK_LLAVEA);
				setState(271);
				((Instruccion_ternarioContext)_localctx).left_instr = expressions();
				setState(272);
				match(TK_LLAVEC);
				setState(273);
				match(ELSE);
				setState(274);
				instr_else_if_ternario();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_else_if_ternarioContext extends ParserRuleContext {
		public Instruccion_ternarioContext instruccion_ternario;
		public List<Instruccion_ternarioContext> e = new ArrayList<Instruccion_ternarioContext>();
		public List<Instruccion_ternarioContext> instruccion_ternario() {
			return getRuleContexts(Instruccion_ternarioContext.class);
		}
		public Instruccion_ternarioContext instruccion_ternario(int i) {
			return getRuleContext(Instruccion_ternarioContext.class,i);
		}
		public Instr_else_if_ternarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_else_if_ternario; }
	}

	public final Instr_else_if_ternarioContext instr_else_if_ternario() throws RecognitionException {
		Instr_else_if_ternarioContext _localctx = new Instr_else_if_ternarioContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_instr_else_if_ternario);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(281);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(278);
					((Instr_else_if_ternarioContext)_localctx).instruccion_ternario = instruccion_ternario();
					((Instr_else_if_ternarioContext)_localctx).e.add(((Instr_else_if_ternarioContext)_localctx).instruccion_ternario);
					}
					} 
				}
				setState(283);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_switchContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(Swiftgrammar.SWITCH, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public List_caseContext list_case() {
			return getRuleContext(List_caseContext.class,0);
		}
		public Block_defaultContext block_default() {
			return getRuleContext(Block_defaultContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Instruccion_switchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_switch; }
	}

	public final Instruccion_switchContext instruccion_switch() throws RecognitionException {
		Instruccion_switchContext _localctx = new Instruccion_switchContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_instruccion_switch);
		try {
			setState(297);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(284);
				match(SWITCH);
				setState(285);
				expressions();
				setState(286);
				match(TK_LLAVEA);
				setState(287);
				list_case();
				setState(288);
				block_default();
				setState(289);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(291);
				match(SWITCH);
				setState(292);
				expressions();
				setState(293);
				match(TK_LLAVEA);
				setState(294);
				block_default();
				setState(295);
				match(TK_LLAVEC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_caseContext extends ParserRuleContext {
		public Instruccion_caseContext instruccion_case;
		public List<Instruccion_caseContext> e = new ArrayList<Instruccion_caseContext>();
		public List<Instruccion_caseContext> instruccion_case() {
			return getRuleContexts(Instruccion_caseContext.class);
		}
		public Instruccion_caseContext instruccion_case(int i) {
			return getRuleContext(Instruccion_caseContext.class,i);
		}
		public List_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_case; }
	}

	public final List_caseContext list_case() throws RecognitionException {
		List_caseContext _localctx = new List_caseContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_list_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(300); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(299);
				((List_caseContext)_localctx).instruccion_case = instruccion_case();
				((List_caseContext)_localctx).e.add(((List_caseContext)_localctx).instruccion_case);
				}
				}
				setState(302); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 577305178290520336L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_caseContext extends ParserRuleContext {
		public List_expre_caseContext list_expre_case() {
			return getRuleContext(List_expre_caseContext.class,0);
		}
		public TerminalNode TK_DOSPUNTOS() { return getToken(Swiftgrammar.TK_DOSPUNTOS, 0); }
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Block_instr_switchContext block_instr_switch() {
			return getRuleContext(Block_instr_switchContext.class,0);
		}
		public TerminalNode TK_COMA() { return getToken(Swiftgrammar.TK_COMA, 0); }
		public Instruccion_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_case; }
	}

	public final Instruccion_caseContext instruccion_case() throws RecognitionException {
		Instruccion_caseContext _localctx = new Instruccion_caseContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_instruccion_case);
		try {
			setState(315);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(304);
				list_expre_case();
				setState(305);
				match(TK_DOSPUNTOS);
				setState(306);
				match(TK_LLAVEA);
				setState(307);
				instrucciones();
				setState(308);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(310);
				list_expre_case();
				setState(311);
				match(TK_DOSPUNTOS);
				setState(312);
				block_instr_switch();
				setState(313);
				match(TK_COMA);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_expre_caseContext extends ParserRuleContext {
		public Block_caseContext block_case;
		public List<Block_caseContext> e = new ArrayList<Block_caseContext>();
		public List<Block_caseContext> block_case() {
			return getRuleContexts(Block_caseContext.class);
		}
		public Block_caseContext block_case(int i) {
			return getRuleContext(Block_caseContext.class,i);
		}
		public List_expre_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_expre_case; }
	}

	public final List_expre_caseContext list_expre_case() throws RecognitionException {
		List_expre_caseContext _localctx = new List_expre_caseContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_list_expre_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(318); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(317);
				((List_expre_caseContext)_localctx).block_case = block_case();
				((List_expre_caseContext)_localctx).e.add(((List_expre_caseContext)_localctx).block_case);
				}
				}
				setState(320); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 577305178290520336L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_caseContext extends ParserRuleContext {
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_BARRA() { return getToken(Swiftgrammar.TK_BARRA, 0); }
		public Block_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_case; }
	}

	public final Block_caseContext block_case() throws RecognitionException {
		Block_caseContext _localctx = new Block_caseContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_block_case);
		try {
			setState(326);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(322);
				expressions();
				setState(323);
				match(TK_BARRA);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(325);
				expressions();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_instr_switchContext extends ParserRuleContext {
		public InstruccionContext instruccion;
		public List<InstruccionContext> list = new ArrayList<InstruccionContext>();
		public InstruccionContext instruccion() {
			return getRuleContext(InstruccionContext.class,0);
		}
		public Block_instr_switchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_instr_switch; }
	}

	public final Block_instr_switchContext block_instr_switch() throws RecognitionException {
		Block_instr_switchContext _localctx = new Block_instr_switchContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_block_instr_switch);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(328);
			((Block_instr_switchContext)_localctx).instruccion = instruccion();
			((Block_instr_switchContext)_localctx).list.add(((Block_instr_switchContext)_localctx).instruccion);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_defaultContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(Swiftgrammar.CASE, 0); }
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_DOSPUNTOS() { return getToken(Swiftgrammar.TK_DOSPUNTOS, 0); }
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Block_instr_switchContext block_instr_switch() {
			return getRuleContext(Block_instr_switchContext.class,0);
		}
		public Instr_defaultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_default; }
	}

	public final Instr_defaultContext instr_default() throws RecognitionException {
		Instr_defaultContext _localctx = new Instr_defaultContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_instr_default);
		try {
			setState(341);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(330);
				match(CASE);
				setState(331);
				match(ID);
				setState(332);
				match(TK_DOSPUNTOS);
				setState(333);
				match(TK_LLAVEA);
				setState(334);
				instrucciones();
				setState(335);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(337);
				match(CASE);
				setState(338);
				match(ID);
				setState(339);
				match(TK_DOSPUNTOS);
				setState(340);
				block_instr_switch();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_defaultContext extends ParserRuleContext {
		public Instr_defaultContext instr_default;
		public List<Instr_defaultContext> e = new ArrayList<Instr_defaultContext>();
		public List<Instr_defaultContext> instr_default() {
			return getRuleContexts(Instr_defaultContext.class);
		}
		public Instr_defaultContext instr_default(int i) {
			return getRuleContext(Instr_defaultContext.class,i);
		}
		public Block_defaultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_default; }
	}

	public final Block_defaultContext block_default() throws RecognitionException {
		Block_defaultContext _localctx = new Block_defaultContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_block_default);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(344); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(343);
				((Block_defaultContext)_localctx).instr_default = instr_default();
				((Block_defaultContext)_localctx).e.add(((Block_defaultContext)_localctx).instr_default);
				}
				}
				setState(346); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_switch_ternarioContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(Swiftgrammar.SWITCH, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public List_case_ternarioContext list_case_ternario() {
			return getRuleContext(List_case_ternarioContext.class,0);
		}
		public Instr_default_terContext instr_default_ter() {
			return getRuleContext(Instr_default_terContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Instruccion_switch_ternarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_switch_ternario; }
	}

	public final Instruccion_switch_ternarioContext instruccion_switch_ternario() throws RecognitionException {
		Instruccion_switch_ternarioContext _localctx = new Instruccion_switch_ternarioContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_instruccion_switch_ternario);
		try {
			setState(361);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(348);
				match(SWITCH);
				setState(349);
				expressions();
				setState(350);
				match(TK_LLAVEA);
				setState(351);
				list_case_ternario();
				setState(352);
				instr_default_ter();
				setState(353);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(355);
				match(SWITCH);
				setState(356);
				expressions();
				setState(357);
				match(TK_LLAVEA);
				setState(358);
				instr_default_ter();
				setState(359);
				match(TK_LLAVEC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_case_ternarioContext extends ParserRuleContext {
		public Instr_case_terContext instr_case_ter;
		public List<Instr_case_terContext> e = new ArrayList<Instr_case_terContext>();
		public List<Instr_case_terContext> instr_case_ter() {
			return getRuleContexts(Instr_case_terContext.class);
		}
		public Instr_case_terContext instr_case_ter(int i) {
			return getRuleContext(Instr_case_terContext.class,i);
		}
		public List_case_ternarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_case_ternario; }
	}

	public final List_case_ternarioContext list_case_ternario() throws RecognitionException {
		List_case_ternarioContext _localctx = new List_case_ternarioContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_list_case_ternario);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(364); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(363);
				((List_case_ternarioContext)_localctx).instr_case_ter = instr_case_ter();
				((List_case_ternarioContext)_localctx).e.add(((List_case_ternarioContext)_localctx).instr_case_ter);
				}
				}
				setState(366); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 577305178290520336L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_case_terContext extends ParserRuleContext {
		public List_expre_case_terContext list_expre_case_ter() {
			return getRuleContext(List_expre_case_terContext.class,0);
		}
		public TerminalNode TK_DOSPUNTOS() { return getToken(Swiftgrammar.TK_DOSPUNTOS, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Instr_case_terContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_case_ter; }
	}

	public final Instr_case_terContext instr_case_ter() throws RecognitionException {
		Instr_case_terContext _localctx = new Instr_case_terContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_instr_case_ter);
		try {
			setState(378);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(368);
				list_expre_case_ter();
				setState(369);
				match(TK_DOSPUNTOS);
				setState(370);
				expressions();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(372);
				list_expre_case_ter();
				setState(373);
				match(TK_DOSPUNTOS);
				setState(374);
				match(TK_LLAVEA);
				setState(375);
				expressions();
				setState(376);
				match(TK_LLAVEC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_expre_case_terContext extends ParserRuleContext {
		public Block_case_terContext block_case_ter;
		public List<Block_case_terContext> e = new ArrayList<Block_case_terContext>();
		public List<Block_case_terContext> block_case_ter() {
			return getRuleContexts(Block_case_terContext.class);
		}
		public Block_case_terContext block_case_ter(int i) {
			return getRuleContext(Block_case_terContext.class,i);
		}
		public List_expre_case_terContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_expre_case_ter; }
	}

	public final List_expre_case_terContext list_expre_case_ter() throws RecognitionException {
		List_expre_case_terContext _localctx = new List_expre_case_terContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_list_expre_case_ter);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(381); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(380);
				((List_expre_case_terContext)_localctx).block_case_ter = block_case_ter();
				((List_expre_case_terContext)_localctx).e.add(((List_expre_case_terContext)_localctx).block_case_ter);
				}
				}
				setState(383); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 577305178290520336L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_case_terContext extends ParserRuleContext {
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_BARRA() { return getToken(Swiftgrammar.TK_BARRA, 0); }
		public Block_case_terContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_case_ter; }
	}

	public final Block_case_terContext block_case_ter() throws RecognitionException {
		Block_case_terContext _localctx = new Block_case_terContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_block_case_ter);
		try {
			setState(389);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(385);
				expressions();
				setState(386);
				match(TK_BARRA);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(388);
				expressions();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_default_terContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(Swiftgrammar.CASE, 0); }
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_DOSPUNTOS() { return getToken(Swiftgrammar.TK_DOSPUNTOS, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Instr_default_terContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_default_ter; }
	}

	public final Instr_default_terContext instr_default_ter() throws RecognitionException {
		Instr_default_terContext _localctx = new Instr_default_terContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_instr_default_ter);
		try {
			setState(402);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(391);
				match(CASE);
				setState(392);
				match(ID);
				setState(393);
				match(TK_DOSPUNTOS);
				setState(394);
				expressions();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(395);
				match(CASE);
				setState(396);
				match(ID);
				setState(397);
				match(TK_DOSPUNTOS);
				setState(398);
				match(TK_LLAVEA);
				setState(399);
				expressions();
				setState(400);
				match(TK_LLAVEC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_whileContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(Swiftgrammar.WHILE, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Instruccion_whileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_while; }
	}

	public final Instruccion_whileContext instruccion_while() throws RecognitionException {
		Instruccion_whileContext _localctx = new Instruccion_whileContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_instruccion_while);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(404);
			match(WHILE);
			setState(405);
			expressions();
			setState(406);
			match(TK_LLAVEA);
			setState(407);
			instrucciones();
			setState(408);
			match(TK_LLAVEC);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_for_inContext extends ParserRuleContext {
		public ExpressionsContext left;
		public ExpressionsContext right;
		public TerminalNode FOR() { return getToken(Swiftgrammar.FOR, 0); }
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode IN() { return getToken(Swiftgrammar.IN, 0); }
		public TerminalNode TK_TRIPLEPUNTO() { return getToken(Swiftgrammar.TK_TRIPLEPUNTO, 0); }
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public List<ExpressionsContext> expressions() {
			return getRuleContexts(ExpressionsContext.class);
		}
		public ExpressionsContext expressions(int i) {
			return getRuleContext(ExpressionsContext.class,i);
		}
		public Instruccion_for_inContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_for_in; }
	}

	public final Instruccion_for_inContext instruccion_for_in() throws RecognitionException {
		Instruccion_for_inContext _localctx = new Instruccion_for_inContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_instruccion_for_in);
		try {
			setState(428);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(410);
				match(FOR);
				setState(411);
				match(ID);
				setState(412);
				match(IN);
				setState(413);
				((Instruccion_for_inContext)_localctx).left = expressions();
				setState(414);
				match(TK_TRIPLEPUNTO);
				setState(415);
				((Instruccion_for_inContext)_localctx).right = expressions();
				setState(416);
				match(TK_LLAVEA);
				setState(417);
				instrucciones();
				setState(418);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(420);
				match(FOR);
				setState(421);
				match(ID);
				setState(422);
				match(IN);
				setState(423);
				((Instruccion_for_inContext)_localctx).left = expressions();
				setState(424);
				match(TK_LLAVEA);
				setState(425);
				instrucciones();
				setState(426);
				match(TK_LLAVEC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_while_trueContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(Swiftgrammar.WHILE, 0); }
		public TerminalNode TRUE() { return getToken(Swiftgrammar.TRUE, 0); }
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Instruccion_while_trueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_while_true; }
	}

	public final Instruccion_while_trueContext instruccion_while_true() throws RecognitionException {
		Instruccion_while_trueContext _localctx = new Instruccion_while_trueContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_instruccion_while_true);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(430);
			match(WHILE);
			setState(431);
			match(TRUE);
			setState(432);
			match(TK_LLAVEA);
			setState(433);
			instrucciones();
			setState(434);
			match(TK_LLAVEC);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_while_true_ternarioContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(Swiftgrammar.WHILE, 0); }
		public TerminalNode TRUE() { return getToken(Swiftgrammar.TRUE, 0); }
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Instruccion_while_true_ternarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_while_true_ternario; }
	}

	public final Instruccion_while_true_ternarioContext instruccion_while_true_ternario() throws RecognitionException {
		Instruccion_while_true_ternarioContext _localctx = new Instruccion_while_true_ternarioContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_instruccion_while_true_ternario);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(436);
			match(WHILE);
			setState(437);
			match(TRUE);
			setState(438);
			match(TK_LLAVEA);
			setState(439);
			instrucciones();
			setState(440);
			match(TK_LLAVEC);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_breakContext extends ParserRuleContext {
		public TerminalNode BREAK() { return getToken(Swiftgrammar.BREAK, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public Instruccion_breakContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_break; }
	}

	public final Instruccion_breakContext instruccion_break() throws RecognitionException {
		Instruccion_breakContext _localctx = new Instruccion_breakContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_instruccion_break);
		try {
			setState(445);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(442);
				match(BREAK);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(443);
				match(BREAK);
				setState(444);
				expressions();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_continueContext extends ParserRuleContext {
		public TerminalNode CONTINUE() { return getToken(Swiftgrammar.CONTINUE, 0); }
		public Instruccion_continueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_continue; }
	}

	public final Instruccion_continueContext instruccion_continue() throws RecognitionException {
		Instruccion_continueContext _localctx = new Instruccion_continueContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_instruccion_continue);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(447);
			match(CONTINUE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_returnContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(Swiftgrammar.RETURN, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public Instruccion_returnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_return; }
	}

	public final Instruccion_returnContext instruccion_return() throws RecognitionException {
		Instruccion_returnContext _localctx = new Instruccion_returnContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_instruccion_return);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(449);
			match(RETURN);
			setState(450);
			expressions();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_funcContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(Swiftgrammar.FUNC, 0); }
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_PARA() { return getToken(Swiftgrammar.TK_PARA, 0); }
		public TerminalNode TK_PARC() { return getToken(Swiftgrammar.TK_PARC, 0); }
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public TerminalNode TK_MENOSMAYOR() { return getToken(Swiftgrammar.TK_MENOSMAYOR, 0); }
		public Instruccion_tipoContext instruccion_tipo() {
			return getRuleContext(Instruccion_tipoContext.class,0);
		}
		public List_function_parametersContext list_function_parameters() {
			return getRuleContext(List_function_parametersContext.class,0);
		}
		public Instruccion_funcContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_func; }
	}

	public final Instruccion_funcContext instruccion_func() throws RecognitionException {
		Instruccion_funcContext _localctx = new Instruccion_funcContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_instruccion_func);
		try {
			setState(490);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(452);
				match(FUNC);
				setState(453);
				match(ID);
				setState(454);
				match(TK_PARA);
				setState(455);
				match(TK_PARC);
				setState(456);
				match(TK_LLAVEA);
				setState(457);
				instrucciones();
				setState(458);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(460);
				match(FUNC);
				setState(461);
				match(ID);
				setState(462);
				match(TK_PARA);
				setState(463);
				match(TK_PARC);
				setState(464);
				match(TK_MENOSMAYOR);
				setState(465);
				instruccion_tipo();
				setState(466);
				match(TK_LLAVEA);
				setState(467);
				instrucciones();
				setState(468);
				match(TK_LLAVEC);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(470);
				match(FUNC);
				setState(471);
				match(ID);
				setState(472);
				match(TK_PARA);
				setState(473);
				list_function_parameters();
				setState(474);
				match(TK_PARC);
				setState(475);
				match(TK_LLAVEA);
				setState(476);
				instrucciones();
				setState(477);
				match(TK_LLAVEC);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(479);
				match(FUNC);
				setState(480);
				match(ID);
				setState(481);
				match(TK_PARA);
				setState(482);
				list_function_parameters();
				setState(483);
				match(TK_PARC);
				setState(484);
				match(TK_MENOSMAYOR);
				setState(485);
				instruccion_tipo();
				setState(486);
				match(TK_LLAVEA);
				setState(487);
				instrucciones();
				setState(488);
				match(TK_LLAVEC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_function_parametersContext extends ParserRuleContext {
		public Block_parameters_fnContext block_parameters_fn;
		public List<Block_parameters_fnContext> e = new ArrayList<Block_parameters_fnContext>();
		public List<Block_parameters_fnContext> block_parameters_fn() {
			return getRuleContexts(Block_parameters_fnContext.class);
		}
		public Block_parameters_fnContext block_parameters_fn(int i) {
			return getRuleContext(Block_parameters_fnContext.class,i);
		}
		public List_function_parametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_function_parameters; }
	}

	public final List_function_parametersContext list_function_parameters() throws RecognitionException {
		List_function_parametersContext _localctx = new List_function_parametersContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_list_function_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(493); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(492);
				((List_function_parametersContext)_localctx).block_parameters_fn = block_parameters_fn();
				((List_function_parametersContext)_localctx).e.add(((List_function_parametersContext)_localctx).block_parameters_fn);
				}
				}
				setState(495); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_parameters_fnContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_DOSPUNTOS() { return getToken(Swiftgrammar.TK_DOSPUNTOS, 0); }
		public Instruccion_tipoContext instruccion_tipo() {
			return getRuleContext(Instruccion_tipoContext.class,0);
		}
		public TerminalNode TK_COMA() { return getToken(Swiftgrammar.TK_COMA, 0); }
		public Block_parameters_fnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_parameters_fn; }
	}

	public final Block_parameters_fnContext block_parameters_fn() throws RecognitionException {
		Block_parameters_fnContext _localctx = new Block_parameters_fnContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_block_parameters_fn);
		try {
			setState(505);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(497);
				match(ID);
				setState(498);
				match(TK_DOSPUNTOS);
				setState(499);
				instruccion_tipo();
				setState(500);
				match(TK_COMA);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(502);
				match(ID);
				setState(503);
				match(TK_DOSPUNTOS);
				setState(504);
				instruccion_tipo();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_llamadaContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_PARA() { return getToken(Swiftgrammar.TK_PARA, 0); }
		public TerminalNode TK_PARC() { return getToken(Swiftgrammar.TK_PARC, 0); }
		public List_expressionContext list_expression() {
			return getRuleContext(List_expressionContext.class,0);
		}
		public Instruccion_llamadaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_llamada; }
	}

	public final Instruccion_llamadaContext instruccion_llamada() throws RecognitionException {
		Instruccion_llamadaContext _localctx = new Instruccion_llamadaContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_instruccion_llamada);
		try {
			setState(515);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(507);
				match(ID);
				setState(508);
				match(TK_PARA);
				setState(509);
				match(TK_PARC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(510);
				match(ID);
				setState(511);
				match(TK_PARA);
				setState(512);
				list_expression();
				setState(513);
				match(TK_PARC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_llamada_expreContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_PARA() { return getToken(Swiftgrammar.TK_PARA, 0); }
		public TerminalNode TK_PARC() { return getToken(Swiftgrammar.TK_PARC, 0); }
		public List_expressionContext list_expression() {
			return getRuleContext(List_expressionContext.class,0);
		}
		public Instr_llamada_expreContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_llamada_expre; }
	}

	public final Instr_llamada_expreContext instr_llamada_expre() throws RecognitionException {
		Instr_llamada_expreContext _localctx = new Instr_llamada_expreContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_instr_llamada_expre);
		try {
			setState(525);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(517);
				match(ID);
				setState(518);
				match(TK_PARA);
				setState(519);
				match(TK_PARC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(520);
				match(ID);
				setState(521);
				match(TK_PARA);
				setState(522);
				list_expression();
				setState(523);
				match(TK_PARC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_structs_declaracionContext extends ParserRuleContext {
		public TerminalNode STRUCT() { return getToken(Swiftgrammar.STRUCT, 0); }
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public List_struct_parametersContext list_struct_parameters() {
			return getRuleContext(List_struct_parametersContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public Instruccion_structs_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_structs_declaracion; }
	}

	public final Instruccion_structs_declaracionContext instruccion_structs_declaracion() throws RecognitionException {
		Instruccion_structs_declaracionContext _localctx = new Instruccion_structs_declaracionContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_instruccion_structs_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(527);
			match(STRUCT);
			setState(528);
			match(ID);
			setState(529);
			match(TK_LLAVEA);
			setState(530);
			list_struct_parameters();
			setState(531);
			match(TK_LLAVEC);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_struct_parametersContext extends ParserRuleContext {
		public Block_structs_parametersContext block_structs_parameters;
		public List<Block_structs_parametersContext> e = new ArrayList<Block_structs_parametersContext>();
		public List<Block_structs_parametersContext> block_structs_parameters() {
			return getRuleContexts(Block_structs_parametersContext.class);
		}
		public Block_structs_parametersContext block_structs_parameters(int i) {
			return getRuleContext(Block_structs_parametersContext.class,i);
		}
		public List_struct_parametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_struct_parameters; }
	}

	public final List_struct_parametersContext list_struct_parameters() throws RecognitionException {
		List_struct_parametersContext _localctx = new List_struct_parametersContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_list_struct_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(534); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(533);
				((List_struct_parametersContext)_localctx).block_structs_parameters = block_structs_parameters();
				((List_struct_parametersContext)_localctx).e.add(((List_struct_parametersContext)_localctx).block_structs_parameters);
				}
				}
				setState(536); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_structs_parametersContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_DOSPUNTOS() { return getToken(Swiftgrammar.TK_DOSPUNTOS, 0); }
		public Instruccion_tipoContext instruccion_tipo() {
			return getRuleContext(Instruccion_tipoContext.class,0);
		}
		public TerminalNode TK_COMA() { return getToken(Swiftgrammar.TK_COMA, 0); }
		public Block_structs_parametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_structs_parameters; }
	}

	public final Block_structs_parametersContext block_structs_parameters() throws RecognitionException {
		Block_structs_parametersContext _localctx = new Block_structs_parametersContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_block_structs_parameters);
		try {
			setState(546);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(538);
				match(ID);
				setState(539);
				match(TK_DOSPUNTOS);
				setState(540);
				instruccion_tipo();
				setState(541);
				match(TK_COMA);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(543);
				match(ID);
				setState(544);
				match(TK_DOSPUNTOS);
				setState(545);
				instruccion_tipo();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_arrays_identifierContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public List_arrays_parameters_idContext list_arrays_parameters_id() {
			return getRuleContext(List_arrays_parameters_idContext.class,0);
		}
		public Instr_arrays_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_arrays_identifier; }
	}

	public final Instr_arrays_identifierContext instr_arrays_identifier() throws RecognitionException {
		Instr_arrays_identifierContext _localctx = new Instr_arrays_identifierContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_instr_arrays_identifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(548);
			match(ID);
			setState(549);
			list_arrays_parameters_id();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_arrays_parameters_idContext extends ParserRuleContext {
		public Block_arrays_identifierContext block_arrays_identifier;
		public List<Block_arrays_identifierContext> e = new ArrayList<Block_arrays_identifierContext>();
		public List<Block_arrays_identifierContext> block_arrays_identifier() {
			return getRuleContexts(Block_arrays_identifierContext.class);
		}
		public Block_arrays_identifierContext block_arrays_identifier(int i) {
			return getRuleContext(Block_arrays_identifierContext.class,i);
		}
		public List_arrays_parameters_idContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_arrays_parameters_id; }
	}

	public final List_arrays_parameters_idContext list_arrays_parameters_id() throws RecognitionException {
		List_arrays_parameters_idContext _localctx = new List_arrays_parameters_idContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_list_arrays_parameters_id);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(552); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(551);
					((List_arrays_parameters_idContext)_localctx).block_arrays_identifier = block_arrays_identifier();
					((List_arrays_parameters_idContext)_localctx).e.add(((List_arrays_parameters_idContext)_localctx).block_arrays_identifier);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(554); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_arrays_identifierContext extends ParserRuleContext {
		public TerminalNode TK_CORA() { return getToken(Swiftgrammar.TK_CORA, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_CORC() { return getToken(Swiftgrammar.TK_CORC, 0); }
		public Block_arrays_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_arrays_identifier; }
	}

	public final Block_arrays_identifierContext block_arrays_identifier() throws RecognitionException {
		Block_arrays_identifierContext _localctx = new Block_arrays_identifierContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_block_arrays_identifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(556);
			match(TK_CORA);
			setState(557);
			expressions();
			setState(558);
			match(TK_CORC);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_structs_declarationContext extends ParserRuleContext {
		public Token left;
		public Token right;
		public TerminalNode VAR() { return getToken(Swiftgrammar.VAR, 0); }
		public TerminalNode TK_IGUAL() { return getToken(Swiftgrammar.TK_IGUAL, 0); }
		public TerminalNode TK_LLAVEA() { return getToken(Swiftgrammar.TK_LLAVEA, 0); }
		public List_struct_parameters_declaContext list_struct_parameters_decla() {
			return getRuleContext(List_struct_parameters_declaContext.class,0);
		}
		public TerminalNode TK_LLAVEC() { return getToken(Swiftgrammar.TK_LLAVEC, 0); }
		public List<TerminalNode> ID() { return getTokens(Swiftgrammar.ID); }
		public TerminalNode ID(int i) {
			return getToken(Swiftgrammar.ID, i);
		}
		public TerminalNode LET() { return getToken(Swiftgrammar.LET, 0); }
		public Instr_structs_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_structs_declaration; }
	}

	public final Instr_structs_declarationContext instr_structs_declaration() throws RecognitionException {
		Instr_structs_declarationContext _localctx = new Instr_structs_declarationContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_instr_structs_declaration);
		try {
			setState(576);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(560);
				match(VAR);
				setState(561);
				((Instr_structs_declarationContext)_localctx).left = match(ID);
				setState(562);
				match(TK_IGUAL);
				setState(563);
				((Instr_structs_declarationContext)_localctx).right = match(ID);
				setState(564);
				match(TK_LLAVEA);
				setState(565);
				list_struct_parameters_decla();
				setState(566);
				match(TK_LLAVEC);
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(568);
				match(LET);
				setState(569);
				((Instr_structs_declarationContext)_localctx).left = match(ID);
				setState(570);
				match(TK_IGUAL);
				setState(571);
				((Instr_structs_declarationContext)_localctx).right = match(ID);
				setState(572);
				match(TK_LLAVEA);
				setState(573);
				list_struct_parameters_decla();
				setState(574);
				match(TK_LLAVEC);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_struct_parameters_declaContext extends ParserRuleContext {
		public Block_structs_parameters_declaContext block_structs_parameters_decla;
		public List<Block_structs_parameters_declaContext> e = new ArrayList<Block_structs_parameters_declaContext>();
		public List<Block_structs_parameters_declaContext> block_structs_parameters_decla() {
			return getRuleContexts(Block_structs_parameters_declaContext.class);
		}
		public Block_structs_parameters_declaContext block_structs_parameters_decla(int i) {
			return getRuleContext(Block_structs_parameters_declaContext.class,i);
		}
		public List_struct_parameters_declaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_struct_parameters_decla; }
	}

	public final List_struct_parameters_declaContext list_struct_parameters_decla() throws RecognitionException {
		List_struct_parameters_declaContext _localctx = new List_struct_parameters_declaContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_list_struct_parameters_decla);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(579); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(578);
				((List_struct_parameters_declaContext)_localctx).block_structs_parameters_decla = block_structs_parameters_decla();
				((List_struct_parameters_declaContext)_localctx).e.add(((List_struct_parameters_declaContext)_localctx).block_structs_parameters_decla);
				}
				}
				setState(581); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_structs_parameters_declaContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public TerminalNode TK_DOSPUNTOS() { return getToken(Swiftgrammar.TK_DOSPUNTOS, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_COMA() { return getToken(Swiftgrammar.TK_COMA, 0); }
		public Block_structs_parameters_declaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_structs_parameters_decla; }
	}

	public final Block_structs_parameters_declaContext block_structs_parameters_decla() throws RecognitionException {
		Block_structs_parameters_declaContext _localctx = new Block_structs_parameters_declaContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_block_structs_parameters_decla);
		try {
			setState(591);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(583);
				match(ID);
				setState(584);
				match(TK_DOSPUNTOS);
				setState(585);
				expressions();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(586);
				match(ID);
				setState(587);
				match(TK_DOSPUNTOS);
				setState(588);
				expressions();
				setState(589);
				match(TK_COMA);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_structs_identifierContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public List_struct_parameters_idContext list_struct_parameters_id() {
			return getRuleContext(List_struct_parameters_idContext.class,0);
		}
		public Instr_structs_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_structs_identifier; }
	}

	public final Instr_structs_identifierContext instr_structs_identifier() throws RecognitionException {
		Instr_structs_identifierContext _localctx = new Instr_structs_identifierContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_instr_structs_identifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(593);
			match(ID);
			setState(594);
			list_struct_parameters_id();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_struct_parameters_idContext extends ParserRuleContext {
		public Block_structs_identifierContext block_structs_identifier;
		public List<Block_structs_identifierContext> e = new ArrayList<Block_structs_identifierContext>();
		public List<Block_structs_identifierContext> block_structs_identifier() {
			return getRuleContexts(Block_structs_identifierContext.class);
		}
		public Block_structs_identifierContext block_structs_identifier(int i) {
			return getRuleContext(Block_structs_identifierContext.class,i);
		}
		public List_struct_parameters_idContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_struct_parameters_id; }
	}

	public final List_struct_parameters_idContext list_struct_parameters_id() throws RecognitionException {
		List_struct_parameters_idContext _localctx = new List_struct_parameters_idContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_list_struct_parameters_id);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(597); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(596);
					((List_struct_parameters_idContext)_localctx).block_structs_identifier = block_structs_identifier();
					((List_struct_parameters_idContext)_localctx).e.add(((List_struct_parameters_idContext)_localctx).block_structs_identifier);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(599); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_structs_identifierContext extends ParserRuleContext {
		public TerminalNode TK_PUNTO() { return getToken(Swiftgrammar.TK_PUNTO, 0); }
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public Block_structs_identifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_structs_identifier; }
	}

	public final Block_structs_identifierContext block_structs_identifier() throws RecognitionException {
		Block_structs_identifierContext _localctx = new Block_structs_identifierContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_block_structs_identifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(601);
			match(TK_PUNTO);
			setState(602);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instr_structs_assignmentContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public List_struct_parameters_idContext list_struct_parameters_id() {
			return getRuleContext(List_struct_parameters_idContext.class,0);
		}
		public TerminalNode TK_IGUAL() { return getToken(Swiftgrammar.TK_IGUAL, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public Instr_structs_assignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instr_structs_assignment; }
	}

	public final Instr_structs_assignmentContext instr_structs_assignment() throws RecognitionException {
		Instr_structs_assignmentContext _localctx = new Instr_structs_assignmentContext(_ctx, getState());
		enterRule(_localctx, 96, RULE_instr_structs_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(604);
			match(ID);
			setState(605);
			list_struct_parameters_id();
			setState(606);
			match(TK_IGUAL);
			setState(607);
			expressions();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Instruccion_tipoContext extends ParserRuleContext {
		public interfaces.TypeExpression tipo_exp;
		public TerminalNode R_INT() { return getToken(Swiftgrammar.R_INT, 0); }
		public TerminalNode R_FLOAT() { return getToken(Swiftgrammar.R_FLOAT, 0); }
		public TerminalNode R_STRING() { return getToken(Swiftgrammar.R_STRING, 0); }
		public TerminalNode R_BOOL() { return getToken(Swiftgrammar.R_BOOL, 0); }
		public TerminalNode R_CHAR() { return getToken(Swiftgrammar.R_CHAR, 0); }
		public Instruccion_tipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_tipo; }
	}

	public final Instruccion_tipoContext instruccion_tipo() throws RecognitionException {
		Instruccion_tipoContext _localctx = new Instruccion_tipoContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_instruccion_tipo);
		try {
			setState(619);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case R_INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(609);
				match(R_INT);
				_localctx.tipo_exp = interfaces.INTEGER
				}
				break;
			case R_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(611);
				match(R_FLOAT);
				_localctx.tipo_exp = interfaces.FLOAT
				}
				break;
			case R_STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(613);
				match(R_STRING);
				_localctx.tipo_exp = interfaces.STRING
				}
				break;
			case R_BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(615);
				match(R_BOOL);
				_localctx.tipo_exp = interfaces.BOOLEAN
				}
				break;
			case R_CHAR:
				enterOuterAlt(_localctx, 5);
				{
				setState(617);
				match(R_CHAR);
				_localctx.tipo_exp = interfaces.CHAR
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_expressionContext extends ParserRuleContext {
		public *arrayList.List l;
		public Block_list_expressionContext block_list_expression;
		public List<Block_list_expressionContext> e = new ArrayList<Block_list_expressionContext>();
		public List<Block_list_expressionContext> block_list_expression() {
			return getRuleContexts(Block_list_expressionContext.class);
		}
		public Block_list_expressionContext block_list_expression(int i) {
			return getRuleContext(Block_list_expressionContext.class,i);
		}
		public List_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_expression; }
	}

	public final List_expressionContext list_expression() throws RecognitionException {
		List_expressionContext _localctx = new List_expressionContext(_ctx, getState());
		enterRule(_localctx, 100, RULE_list_expression);

		    _localctx.l =  arrayList.New()
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(622); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(621);
				((List_expressionContext)_localctx).block_list_expression = block_list_expression();
				((List_expressionContext)_localctx).e.add(((List_expressionContext)_localctx).block_list_expression);
				}
				}
				setState(624); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 577305178290520336L) != 0) );

			        listInt := localctx.(*List_expressionContext).GetE()
			        for _, e := range listInt {
			            _localctx.l.Add(e.GetP())
			        }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Block_list_expressionContext extends ParserRuleContext {
		public interfaces.Expression p;
		public ExpressionsContext expressions;
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_COMA() { return getToken(Swiftgrammar.TK_COMA, 0); }
		public Block_list_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block_list_expression; }
	}

	public final Block_list_expressionContext block_list_expression() throws RecognitionException {
		Block_list_expressionContext _localctx = new Block_list_expressionContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_block_list_expression);
		try {
			setState(635);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(628);
				((Block_list_expressionContext)_localctx).expressions = expressions();
				setState(629);
				match(TK_COMA);
				 _localctx.p =  instruction.NewListExpre(((Block_list_expressionContext)_localctx).expressions.p) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(632);
				((Block_list_expressionContext)_localctx).expressions = expressions();
				 _localctx.p =  instruction.NewListExpre(((Block_list_expressionContext)_localctx).expressions.p) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionsContext extends ParserRuleContext {
		public interfaces.Expression p;
		public Expre_logContext expre_log;
		public Expre_aritContext expre_arit;
		public Expre_relContext expre_rel;
		public Expre_logContext expre_log() {
			return getRuleContext(Expre_logContext.class,0);
		}
		public Expre_aritContext expre_arit() {
			return getRuleContext(Expre_aritContext.class,0);
		}
		public Expre_relContext expre_rel() {
			return getRuleContext(Expre_relContext.class,0);
		}
		public ExpressionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressions; }
	}

	public final ExpressionsContext expressions() throws RecognitionException {
		ExpressionsContext _localctx = new ExpressionsContext(_ctx, getState());
		enterRule(_localctx, 104, RULE_expressions);
		try {
			setState(646);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(637);
				((ExpressionsContext)_localctx).expre_log = expre_log(0);
				 _localctx.p = ((ExpressionsContext)_localctx).expre_log.p 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(640);
				((ExpressionsContext)_localctx).expre_arit = expre_arit(0);
				 _localctx.p = ((ExpressionsContext)_localctx).expre_arit.p 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(643);
				((ExpressionsContext)_localctx).expre_rel = expre_rel(0);
				 _localctx.p = ((ExpressionsContext)_localctx).expre_rel.p 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Expre_logContext extends ParserRuleContext {
		public interfaces.Expression p;
		public Expre_logContext left;
		public Token op;
		public Expre_relContext expre_rel;
		public Expre_logContext right;
		public TerminalNode TK_NOT() { return getToken(Swiftgrammar.TK_NOT, 0); }
		public List<Expre_logContext> expre_log() {
			return getRuleContexts(Expre_logContext.class);
		}
		public Expre_logContext expre_log(int i) {
			return getRuleContext(Expre_logContext.class,i);
		}
		public Expre_relContext expre_rel() {
			return getRuleContext(Expre_relContext.class,0);
		}
		public TerminalNode TK_AND() { return getToken(Swiftgrammar.TK_AND, 0); }
		public TerminalNode TK_OR() { return getToken(Swiftgrammar.TK_OR, 0); }
		public Expre_logContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expre_log; }
	}

	public final Expre_logContext expre_log() throws RecognitionException {
		return expre_log(0);
	}

	private Expre_logContext expre_log(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expre_logContext _localctx = new Expre_logContext(_ctx, _parentState);
		Expre_logContext _prevctx = _localctx;
		int _startState = 106;
		enterRecursionRule(_localctx, 106, RULE_expre_log, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(656);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_NOT:
				{
				setState(649);
				((Expre_logContext)_localctx).op = match(TK_NOT);
				setState(650);
				((Expre_logContext)_localctx).left = expre_log(3);
				 _localctx.p = expression.NewOperacion(((Expre_logContext)_localctx).left.p, (((Expre_logContext)_localctx).op!=null?((Expre_logContext)_localctx).op.getText():null), nil,      true,  (((Expre_logContext)_localctx).op!=null?((Expre_logContext)_localctx).op.getLine():0), localctx.(*Expre_logContext).GetOp().GetColumn()) 
				}
				break;
			case IF:
			case SWITCH:
			case NUMBER:
			case DOUBLE:
			case CHAR:
			case STRING:
			case BOOLEAN:
			case ID:
			case TK_MENOS:
			case TK_PARA:
				{
				setState(653);
				((Expre_logContext)_localctx).expre_rel = expre_rel(0);
				 _localctx.p = ((Expre_logContext)_localctx).expre_rel.p 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(665);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Expre_logContext(_parentctx, _parentState);
					_localctx.left = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_expre_log);
					setState(658);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(659);
					((Expre_logContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !(_la==TK_AND || _la==TK_OR) ) {
						((Expre_logContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(660);
					((Expre_logContext)_localctx).right = expre_log(3);
					 _localctx.p = expression.NewOperacion(((Expre_logContext)_localctx).left.p, (((Expre_logContext)_localctx).op!=null?((Expre_logContext)_localctx).op.getText():null), ((Expre_logContext)_localctx).right.p, false, (((Expre_logContext)_localctx).op!=null?((Expre_logContext)_localctx).op.getLine():0), localctx.(*Expre_logContext).GetOp().GetColumn()) 
					}
					} 
				}
				setState(667);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Expre_relContext extends ParserRuleContext {
		public interfaces.Expression p;
		public Expre_relContext left;
		public Expre_aritContext expre_arit;
		public Token op;
		public Expre_relContext right;
		public Expre_aritContext expre_arit() {
			return getRuleContext(Expre_aritContext.class,0);
		}
		public List<Expre_relContext> expre_rel() {
			return getRuleContexts(Expre_relContext.class);
		}
		public Expre_relContext expre_rel(int i) {
			return getRuleContext(Expre_relContext.class,i);
		}
		public TerminalNode TK_MENOR() { return getToken(Swiftgrammar.TK_MENOR, 0); }
		public TerminalNode TK_MENORIGUAL() { return getToken(Swiftgrammar.TK_MENORIGUAL, 0); }
		public TerminalNode TK_MAYORIGUAL() { return getToken(Swiftgrammar.TK_MAYORIGUAL, 0); }
		public TerminalNode TK_MAYOR() { return getToken(Swiftgrammar.TK_MAYOR, 0); }
		public TerminalNode TK_DIFIGUAL() { return getToken(Swiftgrammar.TK_DIFIGUAL, 0); }
		public TerminalNode TK_IGUALIGUAL() { return getToken(Swiftgrammar.TK_IGUALIGUAL, 0); }
		public Expre_relContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expre_rel; }
	}

	public final Expre_relContext expre_rel() throws RecognitionException {
		return expre_rel(0);
	}

	private Expre_relContext expre_rel(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expre_relContext _localctx = new Expre_relContext(_ctx, _parentState);
		Expre_relContext _prevctx = _localctx;
		int _startState = 108;
		enterRecursionRule(_localctx, 108, RULE_expre_rel, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(669);
			((Expre_relContext)_localctx).expre_arit = expre_arit(0);
			 _localctx.p = ((Expre_relContext)_localctx).expre_arit.p 
			}
			_ctx.stop = _input.LT(-1);
			setState(679);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Expre_relContext(_parentctx, _parentState);
					_localctx.left = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_expre_rel);
					setState(672);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(673);
					((Expre_relContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 16698832846848L) != 0)) ) {
						((Expre_relContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(674);
					((Expre_relContext)_localctx).right = expre_rel(3);
					 _localctx.p = expression.NewOperacion(((Expre_relContext)_localctx).left.p, (((Expre_relContext)_localctx).op!=null?((Expre_relContext)_localctx).op.getText():null), ((Expre_relContext)_localctx).right.p, false, (((Expre_relContext)_localctx).op!=null?((Expre_relContext)_localctx).op.getLine():0), localctx.(*Expre_relContext).GetOp().GetColumn()) 
					}
					} 
				}
				setState(681);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Expre_aritContext extends ParserRuleContext {
		public interfaces.Expression p;
		public Expre_aritContext left;
		public Token op;
		public Expre_valorContext expre_valor;
		public ExpressionsContext expressions;
		public Expre_aritContext right;
		public TerminalNode TK_MENOS() { return getToken(Swiftgrammar.TK_MENOS, 0); }
		public List<Expre_aritContext> expre_arit() {
			return getRuleContexts(Expre_aritContext.class);
		}
		public Expre_aritContext expre_arit(int i) {
			return getRuleContext(Expre_aritContext.class,i);
		}
		public Expre_valorContext expre_valor() {
			return getRuleContext(Expre_valorContext.class,0);
		}
		public TerminalNode TK_PARA() { return getToken(Swiftgrammar.TK_PARA, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode TK_PARC() { return getToken(Swiftgrammar.TK_PARC, 0); }
		public TerminalNode TK_MULT() { return getToken(Swiftgrammar.TK_MULT, 0); }
		public TerminalNode TK_DIV() { return getToken(Swiftgrammar.TK_DIV, 0); }
		public TerminalNode TK_MODULO() { return getToken(Swiftgrammar.TK_MODULO, 0); }
		public TerminalNode TK_MAS() { return getToken(Swiftgrammar.TK_MAS, 0); }
		public Expre_aritContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expre_arit; }
	}

	public final Expre_aritContext expre_arit() throws RecognitionException {
		return expre_arit(0);
	}

	private Expre_aritContext expre_arit(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expre_aritContext _localctx = new Expre_aritContext(_ctx, _parentState);
		Expre_aritContext _prevctx = _localctx;
		int _startState = 110;
		enterRecursionRule(_localctx, 110, RULE_expre_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(695);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_MENOS:
				{
				setState(683);
				((Expre_aritContext)_localctx).op = match(TK_MENOS);
				setState(684);
				((Expre_aritContext)_localctx).left = expre_arit(5);
				 _localctx.p = expression.NewOperacion(((Expre_aritContext)_localctx).left.p, (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getText():null), nil,      true,  (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getLine():0), localctx.(*Expre_aritContext).GetOp().GetColumn()) 
				}
				break;
			case IF:
			case SWITCH:
			case NUMBER:
			case DOUBLE:
			case CHAR:
			case STRING:
			case BOOLEAN:
			case ID:
				{
				setState(687);
				((Expre_aritContext)_localctx).expre_valor = expre_valor();
				 _localctx.p = ((Expre_aritContext)_localctx).expre_valor.p 
				}
				break;
			case TK_PARA:
				{
				setState(690);
				match(TK_PARA);
				setState(691);
				((Expre_aritContext)_localctx).expressions = expressions();
				setState(692);
				match(TK_PARC);
				 _localctx.p = ((Expre_aritContext)_localctx).expressions.p 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(709);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,44,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(707);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
					case 1:
						{
						_localctx = new Expre_aritContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expre_arit);
						setState(697);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(698);
						((Expre_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 123145302310912L) != 0)) ) {
							((Expre_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(699);
						((Expre_aritContext)_localctx).right = expre_arit(5);
						 _localctx.p = expression.NewOperacion(((Expre_aritContext)_localctx).left.p, (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getText():null), ((Expre_aritContext)_localctx).right.p, false, (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getLine():0), localctx.(*Expre_aritContext).GetOp().GetColumn()) 
						}
						break;
					case 2:
						{
						_localctx = new Expre_aritContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expre_arit);
						setState(702);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(703);
						((Expre_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==TK_MAS || _la==TK_MENOS) ) {
							((Expre_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(704);
						((Expre_aritContext)_localctx).right = expre_arit(4);
						 _localctx.p = expression.NewOperacion(((Expre_aritContext)_localctx).left.p, (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getText():null), ((Expre_aritContext)_localctx).right.p, false, (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getLine():0), localctx.(*Expre_aritContext).GetOp().GetColumn()) 
						}
						break;
					}
					} 
				}
				setState(711);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,44,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Expre_valorContext extends ParserRuleContext {
		public interfaces.Expression p;
		public PrimitivoContext primitivo;
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
		}
		public Expre_valorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expre_valor; }
	}

	public final Expre_valorContext expre_valor() throws RecognitionException {
		Expre_valorContext _localctx = new Expre_valorContext(_ctx, getState());
		enterRule(_localctx, 112, RULE_expre_valor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(712);
			((Expre_valorContext)_localctx).primitivo = primitivo();
			 _localctx.p = ((Expre_valorContext)_localctx).primitivo.p 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrimitivoContext extends ParserRuleContext {
		public interfaces.Expression p;
		public Token NUMBER;
		public Token DOUBLE;
		public Token STRING;
		public Token BOOLEAN;
		public Token CHAR;
		public Token ID;
		public TerminalNode NUMBER() { return getToken(Swiftgrammar.NUMBER, 0); }
		public TerminalNode DOUBLE() { return getToken(Swiftgrammar.DOUBLE, 0); }
		public TerminalNode STRING() { return getToken(Swiftgrammar.STRING, 0); }
		public TerminalNode BOOLEAN() { return getToken(Swiftgrammar.BOOLEAN, 0); }
		public TerminalNode CHAR() { return getToken(Swiftgrammar.CHAR, 0); }
		public Instr_llamada_expreContext instr_llamada_expre() {
			return getRuleContext(Instr_llamada_expreContext.class,0);
		}
		public Instr_structs_identifierContext instr_structs_identifier() {
			return getRuleContext(Instr_structs_identifierContext.class,0);
		}
		public Instr_arrays_identifierContext instr_arrays_identifier() {
			return getRuleContext(Instr_arrays_identifierContext.class,0);
		}
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public Instruccion_ternarioContext instruccion_ternario() {
			return getRuleContext(Instruccion_ternarioContext.class,0);
		}
		public Instruccion_switch_ternarioContext instruccion_switch_ternario() {
			return getRuleContext(Instruccion_switch_ternarioContext.class,0);
		}
		public PrimitivoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivo; }
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 114, RULE_primitivo);
		try {
			setState(732);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,45,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(715);
				((PrimitivoContext)_localctx).NUMBER = match(NUMBER);

				              num,err := strconv.Atoi((((PrimitivoContext)_localctx).NUMBER!=null?((PrimitivoContext)_localctx).NUMBER.getText():null))
				                if err!= nil{
				                    fmt.Println(err)
				                }

				            _localctx.p = expression.NewPrimitivo(num, interfaces.INTEGER, interfaces.NULL, (((PrimitivoContext)_localctx).NUMBER!=null?((PrimitivoContext)_localctx).NUMBER.getLine():0), localctx.(*PrimitivoContext).Get_NUMBER().GetColumn())
				       
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(717);
				((PrimitivoContext)_localctx).DOUBLE = match(DOUBLE);
				  
				                num,err := strconv.ParseFloat((((PrimitivoContext)_localctx).DOUBLE!=null?((PrimitivoContext)_localctx).DOUBLE.getText():null), 64)
				                if err!= nil{
				                    fmt.Println(err)
				                }
				            _localctx.p = expression.NewPrimitivo(num, interfaces.FLOAT, interfaces.NULL, (((PrimitivoContext)_localctx).DOUBLE!=null?((PrimitivoContext)_localctx).DOUBLE.getLine():0), localctx.(*PrimitivoContext).Get_DOUBLE().GetColumn())
				              
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(719);
				((PrimitivoContext)_localctx).STRING = match(STRING);
				 
				              str:= (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null)[1:len((((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null))-1]
				              _localctx.p = expression.NewPrimitivo(str, interfaces.STRING, interfaces.NULL, (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getLine():0), localctx.(*PrimitivoContext).Get_STRING().GetColumn())
				            
				            
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(721);
				((PrimitivoContext)_localctx).BOOLEAN = match(BOOLEAN);
				 
				          
				              exp,_ := strconv.ParseBool((((PrimitivoContext)_localctx).BOOLEAN!=null?((PrimitivoContext)_localctx).BOOLEAN.getText():null))
				              _localctx.p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.NULL, (((PrimitivoContext)_localctx).BOOLEAN!=null?((PrimitivoContext)_localctx).BOOLEAN.getLine():0), localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())
				            
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(723);
				((PrimitivoContext)_localctx).CHAR = match(CHAR);


				            str:= (((PrimitivoContext)_localctx).CHAR!=null?((PrimitivoContext)_localctx).CHAR.getText():null)[1]
				            _localctx.p = expression.NewPrimitivo(string(str), interfaces.CHAR, interfaces.NULL, (((PrimitivoContext)_localctx).CHAR!=null?((PrimitivoContext)_localctx).CHAR.getLine():0), localctx.(*PrimitivoContext).Get_CHAR().GetColumn())
				          
				          
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(725);
				instr_llamada_expre();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(726);
				instr_structs_identifier();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(727);
				instr_arrays_identifier();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(728);
				((PrimitivoContext)_localctx).ID = match(ID);
				 _localctx.p = instruction.NewIdentifier((((PrimitivoContext)_localctx).ID!=null?((PrimitivoContext)_localctx).ID.getText():null), (((PrimitivoContext)_localctx).ID!=null?((PrimitivoContext)_localctx).ID.getLine():0), localctx.(*PrimitivoContext).Get_ID().GetColumn()) 
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(730);
				instruccion_ternario();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(731);
				instruccion_switch_ternario();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 53:
			return expre_log_sempred((Expre_logContext)_localctx, predIndex);
		case 54:
			return expre_rel_sempred((Expre_relContext)_localctx, predIndex);
		case 55:
			return expre_arit_sempred((Expre_aritContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expre_log_sempred(Expre_logContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expre_rel_sempred(Expre_relContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expre_arit_sempred(Expre_aritContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001>\u02df\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007+\u0002,\u0007,\u0002"+
		"-\u0007-\u0002.\u0007.\u0002/\u0007/\u00020\u00070\u00021\u00071\u0002"+
		"2\u00072\u00023\u00073\u00024\u00074\u00025\u00075\u00026\u00076\u0002"+
		"7\u00077\u00028\u00078\u00029\u00079\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0004\u0001z\b\u0001\u000b\u0001\f\u0001{\u0001"+
		"\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002\u0096\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003\u00ac\b\u0003\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u00d6\b\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u00f5"+
		"\b\u0006\u0001\u0007\u0005\u0007\u00f8\b\u0007\n\u0007\f\u0007\u00fb\t"+
		"\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u0115"+
		"\b\b\u0001\t\u0005\t\u0118\b\t\n\t\f\t\u011b\t\t\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0003\n\u012a\b\n\u0001\u000b\u0004\u000b\u012d\b\u000b\u000b"+
		"\u000b\f\u000b\u012e\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u013c\b\f\u0001\r\u0004\r\u013f"+
		"\b\r\u000b\r\f\r\u0140\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0003\u000e\u0147\b\u000e\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u0156\b\u0010\u0001\u0011"+
		"\u0004\u0011\u0159\b\u0011\u000b\u0011\f\u0011\u015a\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003"+
		"\u0012\u016a\b\u0012\u0001\u0013\u0004\u0013\u016d\b\u0013\u000b\u0013"+
		"\f\u0013\u016e\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014"+
		"\u017b\b\u0014\u0001\u0015\u0004\u0015\u017e\b\u0015\u000b\u0015\f\u0015"+
		"\u017f\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u0186"+
		"\b\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0003"+
		"\u0017\u0193\b\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0003\u0019\u01ad\b\u0019\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001"+
		"\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0003\u001c\u01be\b\u001c\u0001\u001d\u0001\u001d\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u01eb\b\u001f\u0001"+
		" \u0004 \u01ee\b \u000b \f \u01ef\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0003!\u01fa\b!\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0003\"\u0204\b\"\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0003#\u020e\b#\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001%\u0004%\u0217\b%\u000b%\f%\u0218\u0001&\u0001&\u0001"+
		"&\u0001&\u0001&\u0001&\u0001&\u0001&\u0003&\u0223\b&\u0001\'\u0001\'\u0001"+
		"\'\u0001(\u0004(\u0229\b(\u000b(\f(\u022a\u0001)\u0001)\u0001)\u0001)"+
		"\u0001*\u0001*\u0001*\u0001*\u0001*\u0001*\u0001*\u0001*\u0001*\u0001"+
		"*\u0001*\u0001*\u0001*\u0001*\u0001*\u0001*\u0003*\u0241\b*\u0001+\u0004"+
		"+\u0244\b+\u000b+\f+\u0245\u0001,\u0001,\u0001,\u0001,\u0001,\u0001,\u0001"+
		",\u0001,\u0003,\u0250\b,\u0001-\u0001-\u0001-\u0001.\u0004.\u0256\b.\u000b"+
		".\f.\u0257\u0001/\u0001/\u0001/\u00010\u00010\u00010\u00010\u00010\u0001"+
		"1\u00011\u00011\u00011\u00011\u00011\u00011\u00011\u00011\u00011\u0003"+
		"1\u026c\b1\u00012\u00042\u026f\b2\u000b2\f2\u0270\u00012\u00012\u0001"+
		"3\u00013\u00013\u00013\u00013\u00013\u00013\u00033\u027c\b3\u00014\u0001"+
		"4\u00014\u00014\u00014\u00014\u00014\u00014\u00014\u00034\u0287\b4\u0001"+
		"5\u00015\u00015\u00015\u00015\u00015\u00015\u00015\u00035\u0291\b5\u0001"+
		"5\u00015\u00015\u00015\u00015\u00055\u0298\b5\n5\f5\u029b\t5\u00016\u0001"+
		"6\u00016\u00016\u00016\u00016\u00016\u00016\u00016\u00056\u02a6\b6\n6"+
		"\f6\u02a9\t6\u00017\u00017\u00017\u00017\u00017\u00017\u00017\u00017\u0001"+
		"7\u00017\u00017\u00017\u00017\u00037\u02b8\b7\u00017\u00017\u00017\u0001"+
		"7\u00017\u00017\u00017\u00017\u00017\u00017\u00057\u02c4\b7\n7\f7\u02c7"+
		"\t7\u00018\u00018\u00018\u00019\u00019\u00019\u00019\u00019\u00019\u0001"+
		"9\u00019\u00019\u00019\u00019\u00019\u00019\u00019\u00019\u00019\u0001"+
		"9\u00039\u02dd\b9\u00019\u0000\u0003jln:\u0000\u0002\u0004\u0006\b\n\f"+
		"\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:"+
		"<>@BDFHJLNPRTVXZ\\^`bdfhjlnpr\u0000\u0004\u0002\u00007799\u0002\u0000"+
		"$%(+\u0001\u0000,.\u0001\u0000/0\u02f7\u0000t\u0001\u0000\u0000\u0000"+
		"\u0002y\u0001\u0000\u0000\u0000\u0004\u0095\u0001\u0000\u0000\u0000\u0006"+
		"\u00ab\u0001\u0000\u0000\u0000\b\u00d5\u0001\u0000\u0000\u0000\n\u00d7"+
		"\u0001\u0000\u0000\u0000\f\u00f4\u0001\u0000\u0000\u0000\u000e\u00f9\u0001"+
		"\u0000\u0000\u0000\u0010\u0114\u0001\u0000\u0000\u0000\u0012\u0119\u0001"+
		"\u0000\u0000\u0000\u0014\u0129\u0001\u0000\u0000\u0000\u0016\u012c\u0001"+
		"\u0000\u0000\u0000\u0018\u013b\u0001\u0000\u0000\u0000\u001a\u013e\u0001"+
		"\u0000\u0000\u0000\u001c\u0146\u0001\u0000\u0000\u0000\u001e\u0148\u0001"+
		"\u0000\u0000\u0000 \u0155\u0001\u0000\u0000\u0000\"\u0158\u0001\u0000"+
		"\u0000\u0000$\u0169\u0001\u0000\u0000\u0000&\u016c\u0001\u0000\u0000\u0000"+
		"(\u017a\u0001\u0000\u0000\u0000*\u017d\u0001\u0000\u0000\u0000,\u0185"+
		"\u0001\u0000\u0000\u0000.\u0192\u0001\u0000\u0000\u00000\u0194\u0001\u0000"+
		"\u0000\u00002\u01ac\u0001\u0000\u0000\u00004\u01ae\u0001\u0000\u0000\u0000"+
		"6\u01b4\u0001\u0000\u0000\u00008\u01bd\u0001\u0000\u0000\u0000:\u01bf"+
		"\u0001\u0000\u0000\u0000<\u01c1\u0001\u0000\u0000\u0000>\u01ea\u0001\u0000"+
		"\u0000\u0000@\u01ed\u0001\u0000\u0000\u0000B\u01f9\u0001\u0000\u0000\u0000"+
		"D\u0203\u0001\u0000\u0000\u0000F\u020d\u0001\u0000\u0000\u0000H\u020f"+
		"\u0001\u0000\u0000\u0000J\u0216\u0001\u0000\u0000\u0000L\u0222\u0001\u0000"+
		"\u0000\u0000N\u0224\u0001\u0000\u0000\u0000P\u0228\u0001\u0000\u0000\u0000"+
		"R\u022c\u0001\u0000\u0000\u0000T\u0240\u0001\u0000\u0000\u0000V\u0243"+
		"\u0001\u0000\u0000\u0000X\u024f\u0001\u0000\u0000\u0000Z\u0251\u0001\u0000"+
		"\u0000\u0000\\\u0255\u0001\u0000\u0000\u0000^\u0259\u0001\u0000\u0000"+
		"\u0000`\u025c\u0001\u0000\u0000\u0000b\u026b\u0001\u0000\u0000\u0000d"+
		"\u026e\u0001\u0000\u0000\u0000f\u027b\u0001\u0000\u0000\u0000h\u0286\u0001"+
		"\u0000\u0000\u0000j\u0290\u0001\u0000\u0000\u0000l\u029c\u0001\u0000\u0000"+
		"\u0000n\u02b7\u0001\u0000\u0000\u0000p\u02c8\u0001\u0000\u0000\u0000r"+
		"\u02dc\u0001\u0000\u0000\u0000tu\u0003\u0002\u0001\u0000uv\u0005\u0000"+
		"\u0000\u0001vw\u0006\u0000\uffff\uffff\u0000w\u0001\u0001\u0000\u0000"+
		"\u0000xz\u0003\u0004\u0002\u0000yx\u0001\u0000\u0000\u0000z{\u0001\u0000"+
		"\u0000\u0000{y\u0001\u0000\u0000\u0000{|\u0001\u0000\u0000\u0000|}\u0001"+
		"\u0000\u0000\u0000}~\u0006\u0001\uffff\uffff\u0000~\u0003\u0001\u0000"+
		"\u0000\u0000\u007f\u0080\u0003\u0006\u0003\u0000\u0080\u0081\u0006\u0002"+
		"\uffff\uffff\u0000\u0081\u0096\u0001\u0000\u0000\u0000\u0082\u0096\u0003"+
		"H$\u0000\u0083\u0084\u0003\b\u0004\u0000\u0084\u0085\u0006\u0002\uffff"+
		"\uffff\u0000\u0085\u0096\u0001\u0000\u0000\u0000\u0086\u0087\u0003\n\u0005"+
		"\u0000\u0087\u0088\u0006\u0002\uffff\uffff\u0000\u0088\u0096\u0001\u0000"+
		"\u0000\u0000\u0089\u0096\u0003`0\u0000\u008a\u0096\u0003\f\u0006\u0000"+
		"\u008b\u0096\u00032\u0019\u0000\u008c\u0096\u00030\u0018\u0000\u008d\u0096"+
		"\u00034\u001a\u0000\u008e\u0096\u0003\u0014\n\u0000\u008f\u0096\u0003"+
		"8\u001c\u0000\u0090\u0096\u0003:\u001d\u0000\u0091\u0096\u0003>\u001f"+
		"\u0000\u0092\u0096\u0003D\"\u0000\u0093\u0096\u0003<\u001e\u0000\u0094"+
		"\u0096\u0003T*\u0000\u0095\u007f\u0001\u0000\u0000\u0000\u0095\u0082\u0001"+
		"\u0000\u0000\u0000\u0095\u0083\u0001\u0000\u0000\u0000\u0095\u0086\u0001"+
		"\u0000\u0000\u0000\u0095\u0089\u0001\u0000\u0000\u0000\u0095\u008a\u0001"+
		"\u0000\u0000\u0000\u0095\u008b\u0001\u0000\u0000\u0000\u0095\u008c\u0001"+
		"\u0000\u0000\u0000\u0095\u008d\u0001\u0000\u0000\u0000\u0095\u008e\u0001"+
		"\u0000\u0000\u0000\u0095\u008f\u0001\u0000\u0000\u0000\u0095\u0090\u0001"+
		"\u0000\u0000\u0000\u0095\u0091\u0001\u0000\u0000\u0000\u0095\u0092\u0001"+
		"\u0000\u0000\u0000\u0095\u0093\u0001\u0000\u0000\u0000\u0095\u0094\u0001"+
		"\u0000\u0000\u0000\u0096\u0005\u0001\u0000\u0000\u0000\u0097\u0098\u0005"+
		"\u0001\u0000\u0000\u0098\u0099\u00051\u0000\u0000\u0099\u009a\u0003r9"+
		"\u0000\u009a\u009b\u00052\u0000\u0000\u009b\u009c\u0006\u0003\uffff\uffff"+
		"\u0000\u009c\u00ac\u0001\u0000\u0000\u0000\u009d\u009e\u0005\u0001\u0000"+
		"\u0000\u009e\u009f\u00051\u0000\u0000\u009f\u00a0\u0005\u001b\u0000\u0000"+
		"\u00a0\u00a1\u0005!\u0000\u0000\u00a1\u00a2\u0003d2\u0000\u00a2\u00a3"+
		"\u00052\u0000\u0000\u00a3\u00a4\u0006\u0003\uffff\uffff\u0000\u00a4\u00ac"+
		"\u0001\u0000\u0000\u0000\u00a5\u00a6\u0005\u0001\u0000\u0000\u00a6\u00a7"+
		"\u00051\u0000\u0000\u00a7\u00a8\u0003h4\u0000\u00a8\u00a9\u00052\u0000"+
		"\u0000\u00a9\u00aa\u0006\u0003\uffff\uffff\u0000\u00aa\u00ac\u0001\u0000"+
		"\u0000\u0000\u00ab\u0097\u0001\u0000\u0000\u0000\u00ab\u009d\u0001\u0000"+
		"\u0000\u0000\u00ab\u00a5\u0001\u0000\u0000\u0000\u00ac\u0007\u0001\u0000"+
		"\u0000\u0000\u00ad\u00ae\u0005\u0002\u0000\u0000\u00ae\u00af\u0005\u001d"+
		"\u0000\u0000\u00af\u00b0\u0005#\u0000\u0000\u00b0\u00b1\u0003h4\u0000"+
		"\u00b1\u00b2\u0006\u0004\uffff\uffff\u0000\u00b2\u00d6\u0001\u0000\u0000"+
		"\u0000\u00b3\u00b4\u0005\u0002\u0000\u0000\u00b4\u00b5\u0005\u001d\u0000"+
		"\u0000\u00b5\u00b6\u0005\"\u0000\u0000\u00b6\u00b7\u0003b1\u0000\u00b7"+
		"\u00b8\u0006\u0004\uffff\uffff\u0000\u00b8\u00d6\u0001\u0000\u0000\u0000"+
		"\u00b9\u00ba\u0005\u0002\u0000\u0000\u00ba\u00bb\u0005\u001d\u0000\u0000"+
		"\u00bb\u00bc\u0005\"\u0000\u0000\u00bc\u00bd\u0003b1\u0000\u00bd\u00be"+
		"\u0005#\u0000\u0000\u00be\u00bf\u0003h4\u0000\u00bf\u00c0\u0006\u0004"+
		"\uffff\uffff\u0000\u00c0\u00d6\u0001\u0000\u0000\u0000\u00c1\u00c2\u0005"+
		"\u0003\u0000\u0000\u00c2\u00c3\u0005\u001d\u0000\u0000\u00c3\u00c4\u0005"+
		"#\u0000\u0000\u00c4\u00c5\u0003h4\u0000\u00c5\u00c6\u0006\u0004\uffff"+
		"\uffff\u0000\u00c6\u00d6\u0001\u0000\u0000\u0000\u00c7\u00c8\u0005\u0003"+
		"\u0000\u0000\u00c8\u00c9\u0005\u001d\u0000\u0000\u00c9\u00ca\u0005\"\u0000"+
		"\u0000\u00ca\u00cb\u0003b1\u0000\u00cb\u00cc\u0006\u0004\uffff\uffff\u0000"+
		"\u00cc\u00d6\u0001\u0000\u0000\u0000\u00cd\u00ce\u0005\u0003\u0000\u0000"+
		"\u00ce\u00cf\u0005\u001d\u0000\u0000\u00cf\u00d0\u0005\"\u0000\u0000\u00d0"+
		"\u00d1\u0003b1\u0000\u00d1\u00d2\u0005#\u0000\u0000\u00d2\u00d3\u0003"+
		"h4\u0000\u00d3\u00d4\u0006\u0004\uffff\uffff\u0000\u00d4\u00d6\u0001\u0000"+
		"\u0000\u0000\u00d5\u00ad\u0001\u0000\u0000\u0000\u00d5\u00b3\u0001\u0000"+
		"\u0000\u0000\u00d5\u00b9\u0001\u0000\u0000\u0000\u00d5\u00c1\u0001\u0000"+
		"\u0000\u0000\u00d5\u00c7\u0001\u0000\u0000\u0000\u00d5\u00cd\u0001\u0000"+
		"\u0000\u0000\u00d6\t\u0001\u0000\u0000\u0000\u00d7\u00d8\u0005\u001d\u0000"+
		"\u0000\u00d8\u00d9\u0005#\u0000\u0000\u00d9\u00da\u0003h4\u0000\u00da"+
		"\u00db\u0006\u0005\uffff\uffff\u0000\u00db\u000b\u0001\u0000\u0000\u0000"+
		"\u00dc\u00dd\u0005\u0004\u0000\u0000\u00dd\u00de\u0003h4\u0000\u00de\u00df"+
		"\u00053\u0000\u0000\u00df\u00e0\u0003\u0002\u0001\u0000\u00e0\u00e1\u0005"+
		"4\u0000\u0000\u00e1\u00f5\u0001\u0000\u0000\u0000\u00e2\u00e3\u0005\u0004"+
		"\u0000\u0000\u00e3\u00e4\u0003h4\u0000\u00e4\u00e5\u00053\u0000\u0000"+
		"\u00e5\u00e6\u0003\u0002\u0001\u0000\u00e6\u00e7\u00054\u0000\u0000\u00e7"+
		"\u00e8\u0005\u0005\u0000\u0000\u00e8\u00e9\u00053\u0000\u0000\u00e9\u00ea"+
		"\u0003\u0002\u0001\u0000\u00ea\u00eb\u00054\u0000\u0000\u00eb\u00f5\u0001"+
		"\u0000\u0000\u0000\u00ec\u00ed\u0005\u0004\u0000\u0000\u00ed\u00ee\u0003"+
		"h4\u0000\u00ee\u00ef\u00053\u0000\u0000\u00ef\u00f0\u0003\u0002\u0001"+
		"\u0000\u00f0\u00f1\u00054\u0000\u0000\u00f1\u00f2\u0005\u0005\u0000\u0000"+
		"\u00f2\u00f3\u0003\u000e\u0007\u0000\u00f3\u00f5\u0001\u0000\u0000\u0000"+
		"\u00f4\u00dc\u0001\u0000\u0000\u0000\u00f4\u00e2\u0001\u0000\u0000\u0000"+
		"\u00f4\u00ec\u0001\u0000\u0000\u0000\u00f5\r\u0001\u0000\u0000\u0000\u00f6"+
		"\u00f8\u0003\f\u0006\u0000\u00f7\u00f6\u0001\u0000\u0000\u0000\u00f8\u00fb"+
		"\u0001\u0000\u0000\u0000\u00f9\u00f7\u0001\u0000\u0000\u0000\u00f9\u00fa"+
		"\u0001\u0000\u0000\u0000\u00fa\u000f\u0001\u0000\u0000\u0000\u00fb\u00f9"+
		"\u0001\u0000\u0000\u0000\u00fc\u00fd\u0005\u0004\u0000\u0000\u00fd\u00fe"+
		"\u0003h4\u0000\u00fe\u00ff\u00053\u0000\u0000\u00ff\u0100\u0003h4\u0000"+
		"\u0100\u0101\u00054\u0000\u0000\u0101\u0115\u0001\u0000\u0000\u0000\u0102"+
		"\u0103\u0005\u0004\u0000\u0000\u0103\u0104\u0003h4\u0000\u0104\u0105\u0005"+
		"3\u0000\u0000\u0105\u0106\u0003h4\u0000\u0106\u0107\u00054\u0000\u0000"+
		"\u0107\u0108\u0005\u0005\u0000\u0000\u0108\u0109\u00053\u0000\u0000\u0109"+
		"\u010a\u0003h4\u0000\u010a\u010b\u00054\u0000\u0000\u010b\u0115\u0001"+
		"\u0000\u0000\u0000\u010c\u010d\u0005\u0004\u0000\u0000\u010d\u010e\u0003"+
		"h4\u0000\u010e\u010f\u00053\u0000\u0000\u010f\u0110\u0003h4\u0000\u0110"+
		"\u0111\u00054\u0000\u0000\u0111\u0112\u0005\u0005\u0000\u0000\u0112\u0113"+
		"\u0003\u0012\t\u0000\u0113\u0115\u0001\u0000\u0000\u0000\u0114\u00fc\u0001"+
		"\u0000\u0000\u0000\u0114\u0102\u0001\u0000\u0000\u0000\u0114\u010c\u0001"+
		"\u0000\u0000\u0000\u0115\u0011\u0001\u0000\u0000\u0000\u0116\u0118\u0003"+
		"\u0010\b\u0000\u0117\u0116\u0001\u0000\u0000\u0000\u0118\u011b\u0001\u0000"+
		"\u0000\u0000\u0119\u0117\u0001\u0000\u0000\u0000\u0119\u011a\u0001\u0000"+
		"\u0000\u0000\u011a\u0013\u0001\u0000\u0000\u0000\u011b\u0119\u0001\u0000"+
		"\u0000\u0000\u011c\u011d\u0005\b\u0000\u0000\u011d\u011e\u0003h4\u0000"+
		"\u011e\u011f\u00053\u0000\u0000\u011f\u0120\u0003\u0016\u000b\u0000\u0120"+
		"\u0121\u0003\"\u0011\u0000\u0121\u0122\u00054\u0000\u0000\u0122\u012a"+
		"\u0001\u0000\u0000\u0000\u0123\u0124\u0005\b\u0000\u0000\u0124\u0125\u0003"+
		"h4\u0000\u0125\u0126\u00053\u0000\u0000\u0126\u0127\u0003\"\u0011\u0000"+
		"\u0127\u0128\u00054\u0000\u0000\u0128\u012a\u0001\u0000\u0000\u0000\u0129"+
		"\u011c\u0001\u0000\u0000\u0000\u0129\u0123\u0001\u0000\u0000\u0000\u012a"+
		"\u0015\u0001\u0000\u0000\u0000\u012b\u012d\u0003\u0018\f\u0000\u012c\u012b"+
		"\u0001\u0000\u0000\u0000\u012d\u012e\u0001\u0000\u0000\u0000\u012e\u012c"+
		"\u0001\u0000\u0000\u0000\u012e\u012f\u0001\u0000\u0000\u0000\u012f\u0017"+
		"\u0001\u0000\u0000\u0000\u0130\u0131\u0003\u001a\r\u0000\u0131\u0132\u0005"+
		"\"\u0000\u0000\u0132\u0133\u00053\u0000\u0000\u0133\u0134\u0003\u0002"+
		"\u0001\u0000\u0134\u0135\u00054\u0000\u0000\u0135\u013c\u0001\u0000\u0000"+
		"\u0000\u0136\u0137\u0003\u001a\r\u0000\u0137\u0138\u0005\"\u0000\u0000"+
		"\u0138\u0139\u0003\u001e\u000f\u0000\u0139\u013a\u0005!\u0000\u0000\u013a"+
		"\u013c\u0001\u0000\u0000\u0000\u013b\u0130\u0001\u0000\u0000\u0000\u013b"+
		"\u0136\u0001\u0000\u0000\u0000\u013c\u0019\u0001\u0000\u0000\u0000\u013d"+
		"\u013f\u0003\u001c\u000e\u0000\u013e\u013d\u0001\u0000\u0000\u0000\u013f"+
		"\u0140\u0001\u0000\u0000\u0000\u0140\u013e\u0001\u0000\u0000\u0000\u0140"+
		"\u0141\u0001\u0000\u0000\u0000\u0141\u001b\u0001\u0000\u0000\u0000\u0142"+
		"\u0143\u0003h4\u0000\u0143\u0144\u0005:\u0000\u0000\u0144\u0147\u0001"+
		"\u0000\u0000\u0000\u0145\u0147\u0003h4\u0000\u0146\u0142\u0001\u0000\u0000"+
		"\u0000\u0146\u0145\u0001\u0000\u0000\u0000\u0147\u001d\u0001\u0000\u0000"+
		"\u0000\u0148\u0149\u0003\u0004\u0002\u0000\u0149\u001f\u0001\u0000\u0000"+
		"\u0000\u014a\u014b\u0005\t\u0000\u0000\u014b\u014c\u0005\u001d\u0000\u0000"+
		"\u014c\u014d\u0005\"\u0000\u0000\u014d\u014e\u00053\u0000\u0000\u014e"+
		"\u014f\u0003\u0002\u0001\u0000\u014f\u0150\u00054\u0000\u0000\u0150\u0156"+
		"\u0001\u0000\u0000\u0000\u0151\u0152\u0005\t\u0000\u0000\u0152\u0153\u0005"+
		"\u001d\u0000\u0000\u0153\u0154\u0005\"\u0000\u0000\u0154\u0156\u0003\u001e"+
		"\u000f\u0000\u0155\u014a\u0001\u0000\u0000\u0000\u0155\u0151\u0001\u0000"+
		"\u0000\u0000\u0156!\u0001\u0000\u0000\u0000\u0157\u0159\u0003 \u0010\u0000"+
		"\u0158\u0157\u0001\u0000\u0000\u0000\u0159\u015a\u0001\u0000\u0000\u0000"+
		"\u015a\u0158\u0001\u0000\u0000\u0000\u015a\u015b\u0001\u0000\u0000\u0000"+
		"\u015b#\u0001\u0000\u0000\u0000\u015c\u015d\u0005\b\u0000\u0000\u015d"+
		"\u015e\u0003h4\u0000\u015e\u015f\u00053\u0000\u0000\u015f\u0160\u0003"+
		"&\u0013\u0000\u0160\u0161\u0003.\u0017\u0000\u0161\u0162\u00054\u0000"+
		"\u0000\u0162\u016a\u0001\u0000\u0000\u0000\u0163\u0164\u0005\b\u0000\u0000"+
		"\u0164\u0165\u0003h4\u0000\u0165\u0166\u00053\u0000\u0000\u0166\u0167"+
		"\u0003.\u0017\u0000\u0167\u0168\u00054\u0000\u0000\u0168\u016a\u0001\u0000"+
		"\u0000\u0000\u0169\u015c\u0001\u0000\u0000\u0000\u0169\u0163\u0001\u0000"+
		"\u0000\u0000\u016a%\u0001\u0000\u0000\u0000\u016b\u016d\u0003(\u0014\u0000"+
		"\u016c\u016b\u0001\u0000\u0000\u0000\u016d\u016e\u0001\u0000\u0000\u0000"+
		"\u016e\u016c\u0001\u0000\u0000\u0000\u016e\u016f\u0001\u0000\u0000\u0000"+
		"\u016f\'\u0001\u0000\u0000\u0000\u0170\u0171\u0003*\u0015\u0000\u0171"+
		"\u0172\u0005\"\u0000\u0000\u0172\u0173\u0003h4\u0000\u0173\u017b\u0001"+
		"\u0000\u0000\u0000\u0174\u0175\u0003*\u0015\u0000\u0175\u0176\u0005\""+
		"\u0000\u0000\u0176\u0177\u00053\u0000\u0000\u0177\u0178\u0003h4\u0000"+
		"\u0178\u0179\u00054\u0000\u0000\u0179\u017b\u0001\u0000\u0000\u0000\u017a"+
		"\u0170\u0001\u0000\u0000\u0000\u017a\u0174\u0001\u0000\u0000\u0000\u017b"+
		")\u0001\u0000\u0000\u0000\u017c\u017e\u0003,\u0016\u0000\u017d\u017c\u0001"+
		"\u0000\u0000\u0000\u017e\u017f\u0001\u0000\u0000\u0000\u017f\u017d\u0001"+
		"\u0000\u0000\u0000\u017f\u0180\u0001\u0000\u0000\u0000\u0180+\u0001\u0000"+
		"\u0000\u0000\u0181\u0182\u0003h4\u0000\u0182\u0183\u0005:\u0000\u0000"+
		"\u0183\u0186\u0001\u0000\u0000\u0000\u0184\u0186\u0003h4\u0000\u0185\u0181"+
		"\u0001\u0000\u0000\u0000\u0185\u0184\u0001\u0000\u0000\u0000\u0186-\u0001"+
		"\u0000\u0000\u0000\u0187\u0188\u0005\t\u0000\u0000\u0188\u0189\u0005\u001d"+
		"\u0000\u0000\u0189\u018a\u0005\"\u0000\u0000\u018a\u0193\u0003h4\u0000"+
		"\u018b\u018c\u0005\t\u0000\u0000\u018c\u018d\u0005\u001d\u0000\u0000\u018d"+
		"\u018e\u0005\"\u0000\u0000\u018e\u018f\u00053\u0000\u0000\u018f\u0190"+
		"\u0003h4\u0000\u0190\u0191\u00054\u0000\u0000\u0191\u0193\u0001\u0000"+
		"\u0000\u0000\u0192\u0187\u0001\u0000\u0000\u0000\u0192\u018b\u0001\u0000"+
		"\u0000\u0000\u0193/\u0001\u0000\u0000\u0000\u0194\u0195\u0005\u0007\u0000"+
		"\u0000\u0195\u0196\u0003h4\u0000\u0196\u0197\u00053\u0000\u0000\u0197"+
		"\u0198\u0003\u0002\u0001\u0000\u0198\u0199\u00054\u0000\u0000\u01991\u0001"+
		"\u0000\u0000\u0000\u019a\u019b\u0005\u0006\u0000\u0000\u019b\u019c\u0005"+
		"\u001d\u0000\u0000\u019c\u019d\u0005\u000b\u0000\u0000\u019d\u019e\u0003"+
		"h4\u0000\u019e\u019f\u0005\u001e\u0000\u0000\u019f\u01a0\u0003h4\u0000"+
		"\u01a0\u01a1\u00053\u0000\u0000\u01a1\u01a2\u0003\u0002\u0001\u0000\u01a2"+
		"\u01a3\u00054\u0000\u0000\u01a3\u01ad\u0001\u0000\u0000\u0000\u01a4\u01a5"+
		"\u0005\u0006\u0000\u0000\u01a5\u01a6\u0005\u001d\u0000\u0000\u01a6\u01a7"+
		"\u0005\u000b\u0000\u0000\u01a7\u01a8\u0003h4\u0000\u01a8\u01a9\u00053"+
		"\u0000\u0000\u01a9\u01aa\u0003\u0002\u0001\u0000\u01aa\u01ab\u00054\u0000"+
		"\u0000\u01ab\u01ad\u0001\u0000\u0000\u0000\u01ac\u019a\u0001\u0000\u0000"+
		"\u0000\u01ac\u01a4\u0001\u0000\u0000\u0000\u01ad3\u0001\u0000\u0000\u0000"+
		"\u01ae\u01af\u0005\u0007\u0000\u0000\u01af\u01b0\u0005\u0017\u0000\u0000"+
		"\u01b0\u01b1\u00053\u0000\u0000\u01b1\u01b2\u0003\u0002\u0001\u0000\u01b2"+
		"\u01b3\u00054\u0000\u0000\u01b35\u0001\u0000\u0000\u0000\u01b4\u01b5\u0005"+
		"\u0007\u0000\u0000\u01b5\u01b6\u0005\u0017\u0000\u0000\u01b6\u01b7\u0005"+
		"3\u0000\u0000\u01b7\u01b8\u0003\u0002\u0001\u0000\u01b8\u01b9\u00054\u0000"+
		"\u0000\u01b97\u0001\u0000\u0000\u0000\u01ba\u01be\u0005\f\u0000\u0000"+
		"\u01bb\u01bc\u0005\f\u0000\u0000\u01bc\u01be\u0003h4\u0000\u01bd\u01ba"+
		"\u0001\u0000\u0000\u0000\u01bd\u01bb\u0001\u0000\u0000\u0000\u01be9\u0001"+
		"\u0000\u0000\u0000\u01bf\u01c0\u0005\r\u0000\u0000\u01c0;\u0001\u0000"+
		"\u0000\u0000\u01c1\u01c2\u0005\u000e\u0000\u0000\u01c2\u01c3\u0003h4\u0000"+
		"\u01c3=\u0001\u0000\u0000\u0000\u01c4\u01c5\u0005\u000f\u0000\u0000\u01c5"+
		"\u01c6\u0005\u001d\u0000\u0000\u01c6\u01c7\u00051\u0000\u0000\u01c7\u01c8"+
		"\u00052\u0000\u0000\u01c8\u01c9\u00053\u0000\u0000\u01c9\u01ca\u0003\u0002"+
		"\u0001\u0000\u01ca\u01cb\u00054\u0000\u0000\u01cb\u01eb\u0001\u0000\u0000"+
		"\u0000\u01cc\u01cd\u0005\u000f\u0000\u0000\u01cd\u01ce\u0005\u001d\u0000"+
		"\u0000\u01ce\u01cf\u00051\u0000\u0000\u01cf\u01d0\u00052\u0000\u0000\u01d0"+
		"\u01d1\u0005\'\u0000\u0000\u01d1\u01d2\u0003b1\u0000\u01d2\u01d3\u0005"+
		"3\u0000\u0000\u01d3\u01d4\u0003\u0002\u0001\u0000\u01d4\u01d5\u00054\u0000"+
		"\u0000\u01d5\u01eb\u0001\u0000\u0000\u0000\u01d6\u01d7\u0005\u000f\u0000"+
		"\u0000\u01d7\u01d8\u0005\u001d\u0000\u0000\u01d8\u01d9\u00051\u0000\u0000"+
		"\u01d9\u01da\u0003@ \u0000\u01da\u01db\u00052\u0000\u0000\u01db\u01dc"+
		"\u00053\u0000\u0000\u01dc\u01dd\u0003\u0002\u0001\u0000\u01dd\u01de\u0005"+
		"4\u0000\u0000\u01de\u01eb\u0001\u0000\u0000\u0000\u01df\u01e0\u0005\u000f"+
		"\u0000\u0000\u01e0\u01e1\u0005\u001d\u0000\u0000\u01e1\u01e2\u00051\u0000"+
		"\u0000\u01e2\u01e3\u0003@ \u0000\u01e3\u01e4\u00052\u0000\u0000\u01e4"+
		"\u01e5\u0005\'\u0000\u0000\u01e5\u01e6\u0003b1\u0000\u01e6\u01e7\u0005"+
		"3\u0000\u0000\u01e7\u01e8\u0003\u0002\u0001\u0000\u01e8\u01e9\u00054\u0000"+
		"\u0000\u01e9\u01eb\u0001\u0000\u0000\u0000\u01ea\u01c4\u0001\u0000\u0000"+
		"\u0000\u01ea\u01cc\u0001\u0000\u0000\u0000\u01ea\u01d6\u0001\u0000\u0000"+
		"\u0000\u01ea\u01df\u0001\u0000\u0000\u0000\u01eb?\u0001\u0000\u0000\u0000"+
		"\u01ec\u01ee\u0003B!\u0000\u01ed\u01ec\u0001\u0000\u0000\u0000\u01ee\u01ef"+
		"\u0001\u0000\u0000\u0000\u01ef\u01ed\u0001\u0000\u0000\u0000\u01ef\u01f0"+
		"\u0001\u0000\u0000\u0000\u01f0A\u0001\u0000\u0000\u0000\u01f1\u01f2\u0005"+
		"\u001d\u0000\u0000\u01f2\u01f3\u0005\"\u0000\u0000\u01f3\u01f4\u0003b"+
		"1\u0000\u01f4\u01f5\u0005!\u0000\u0000\u01f5\u01fa\u0001\u0000\u0000\u0000"+
		"\u01f6\u01f7\u0005\u001d\u0000\u0000\u01f7\u01f8\u0005\"\u0000\u0000\u01f8"+
		"\u01fa\u0003b1\u0000\u01f9\u01f1\u0001\u0000\u0000\u0000\u01f9\u01f6\u0001"+
		"\u0000\u0000\u0000\u01faC\u0001\u0000\u0000\u0000\u01fb\u01fc\u0005\u001d"+
		"\u0000\u0000\u01fc\u01fd\u00051\u0000\u0000\u01fd\u0204\u00052\u0000\u0000"+
		"\u01fe\u01ff\u0005\u001d\u0000\u0000\u01ff\u0200\u00051\u0000\u0000\u0200"+
		"\u0201\u0003d2\u0000\u0201\u0202\u00052\u0000\u0000\u0202\u0204\u0001"+
		"\u0000\u0000\u0000\u0203\u01fb\u0001\u0000\u0000\u0000\u0203\u01fe\u0001"+
		"\u0000\u0000\u0000\u0204E\u0001\u0000\u0000\u0000\u0205\u0206\u0005\u001d"+
		"\u0000\u0000\u0206\u0207\u00051\u0000\u0000\u0207\u020e\u00052\u0000\u0000"+
		"\u0208\u0209\u0005\u001d\u0000\u0000\u0209\u020a\u00051\u0000\u0000\u020a"+
		"\u020b\u0003d2\u0000\u020b\u020c\u00052\u0000\u0000\u020c\u020e\u0001"+
		"\u0000\u0000\u0000\u020d\u0205\u0001\u0000\u0000\u0000\u020d\u0208\u0001"+
		"\u0000\u0000\u0000\u020eG\u0001\u0000\u0000\u0000\u020f\u0210\u0005\u0010"+
		"\u0000\u0000\u0210\u0211\u0005\u001d\u0000\u0000\u0211\u0212\u00053\u0000"+
		"\u0000\u0212\u0213\u0003J%\u0000\u0213\u0214\u00054\u0000\u0000\u0214"+
		"I\u0001\u0000\u0000\u0000\u0215\u0217\u0003L&\u0000\u0216\u0215\u0001"+
		"\u0000\u0000\u0000\u0217\u0218\u0001\u0000\u0000\u0000\u0218\u0216\u0001"+
		"\u0000\u0000\u0000\u0218\u0219\u0001\u0000\u0000\u0000\u0219K\u0001\u0000"+
		"\u0000\u0000\u021a\u021b\u0005\u001d\u0000\u0000\u021b\u021c\u0005\"\u0000"+
		"\u0000\u021c\u021d\u0003b1\u0000\u021d\u021e\u0005!\u0000\u0000\u021e"+
		"\u0223\u0001\u0000\u0000\u0000\u021f\u0220\u0005\u001d\u0000\u0000\u0220"+
		"\u0221\u0005\"\u0000\u0000\u0221\u0223\u0003b1\u0000\u0222\u021a\u0001"+
		"\u0000\u0000\u0000\u0222\u021f\u0001\u0000\u0000\u0000\u0223M\u0001\u0000"+
		"\u0000\u0000\u0224\u0225\u0005\u001d\u0000\u0000\u0225\u0226\u0003P(\u0000"+
		"\u0226O\u0001\u0000\u0000\u0000\u0227\u0229\u0003R)\u0000\u0228\u0227"+
		"\u0001\u0000\u0000\u0000\u0229\u022a\u0001\u0000\u0000\u0000\u022a\u0228"+
		"\u0001\u0000\u0000\u0000\u022a\u022b\u0001\u0000\u0000\u0000\u022bQ\u0001"+
		"\u0000\u0000\u0000\u022c\u022d\u00055\u0000\u0000\u022d\u022e\u0003h4"+
		"\u0000\u022e\u022f\u00056\u0000\u0000\u022fS\u0001\u0000\u0000\u0000\u0230"+
		"\u0231\u0005\u0002\u0000\u0000\u0231\u0232\u0005\u001d\u0000\u0000\u0232"+
		"\u0233\u0005#\u0000\u0000\u0233\u0234\u0005\u001d\u0000\u0000\u0234\u0235"+
		"\u00053\u0000\u0000\u0235\u0236\u0003V+\u0000\u0236\u0237\u00054\u0000"+
		"\u0000\u0237\u0241\u0001\u0000\u0000\u0000\u0238\u0239\u0005\u0003\u0000"+
		"\u0000\u0239\u023a\u0005\u001d\u0000\u0000\u023a\u023b\u0005#\u0000\u0000"+
		"\u023b\u023c\u0005\u001d\u0000\u0000\u023c\u023d\u00053\u0000\u0000\u023d"+
		"\u023e\u0003V+\u0000\u023e\u023f\u00054\u0000\u0000\u023f\u0241\u0001"+
		"\u0000\u0000\u0000\u0240\u0230\u0001\u0000\u0000\u0000\u0240\u0238\u0001"+
		"\u0000\u0000\u0000\u0241U\u0001\u0000\u0000\u0000\u0242\u0244\u0003X,"+
		"\u0000\u0243\u0242\u0001\u0000\u0000\u0000\u0244\u0245\u0001\u0000\u0000"+
		"\u0000\u0245\u0243\u0001\u0000\u0000\u0000\u0245\u0246\u0001\u0000\u0000"+
		"\u0000\u0246W\u0001\u0000\u0000\u0000\u0247\u0248\u0005\u001d\u0000\u0000"+
		"\u0248\u0249\u0005\"\u0000\u0000\u0249\u0250\u0003h4\u0000\u024a\u024b"+
		"\u0005\u001d\u0000\u0000\u024b\u024c\u0005\"\u0000\u0000\u024c\u024d\u0003"+
		"h4\u0000\u024d\u024e\u0005!\u0000\u0000\u024e\u0250\u0001\u0000\u0000"+
		"\u0000\u024f\u0247\u0001\u0000\u0000\u0000\u024f\u024a\u0001\u0000\u0000"+
		"\u0000\u0250Y\u0001\u0000\u0000\u0000\u0251\u0252\u0005\u001d\u0000\u0000"+
		"\u0252\u0253\u0003\\.\u0000\u0253[\u0001\u0000\u0000\u0000\u0254\u0256"+
		"\u0003^/\u0000\u0255\u0254\u0001\u0000\u0000\u0000\u0256\u0257\u0001\u0000"+
		"\u0000\u0000\u0257\u0255\u0001\u0000\u0000\u0000\u0257\u0258\u0001\u0000"+
		"\u0000\u0000\u0258]\u0001\u0000\u0000\u0000\u0259\u025a\u0005\u001f\u0000"+
		"\u0000\u025a\u025b\u0005\u001d\u0000\u0000\u025b_\u0001\u0000\u0000\u0000"+
		"\u025c\u025d\u0005\u001d\u0000\u0000\u025d\u025e\u0003\\.\u0000\u025e"+
		"\u025f\u0005#\u0000\u0000\u025f\u0260\u0003h4\u0000\u0260a\u0001\u0000"+
		"\u0000\u0000\u0261\u0262\u0005\u0011\u0000\u0000\u0262\u026c\u00061\uffff"+
		"\uffff\u0000\u0263\u0264\u0005\u0012\u0000\u0000\u0264\u026c\u00061\uffff"+
		"\uffff\u0000\u0265\u0266\u0005\u0016\u0000\u0000\u0266\u026c\u00061\uffff"+
		"\uffff\u0000\u0267\u0268\u0005\u0014\u0000\u0000\u0268\u026c\u00061\uffff"+
		"\uffff\u0000\u0269\u026a\u0005\u0015\u0000\u0000\u026a\u026c\u00061\uffff"+
		"\uffff\u0000\u026b\u0261\u0001\u0000\u0000\u0000\u026b\u0263\u0001\u0000"+
		"\u0000\u0000\u026b\u0265\u0001\u0000\u0000\u0000\u026b\u0267\u0001\u0000"+
		"\u0000\u0000\u026b\u0269\u0001\u0000\u0000\u0000\u026cc\u0001\u0000\u0000"+
		"\u0000\u026d\u026f\u0003f3\u0000\u026e\u026d\u0001\u0000\u0000\u0000\u026f"+
		"\u0270\u0001\u0000\u0000\u0000\u0270\u026e\u0001\u0000\u0000\u0000\u0270"+
		"\u0271\u0001\u0000\u0000\u0000\u0271\u0272\u0001\u0000\u0000\u0000\u0272"+
		"\u0273\u00062\uffff\uffff\u0000\u0273e\u0001\u0000\u0000\u0000\u0274\u0275"+
		"\u0003h4\u0000\u0275\u0276\u0005!\u0000\u0000\u0276\u0277\u00063\uffff"+
		"\uffff\u0000\u0277\u027c\u0001\u0000\u0000\u0000\u0278\u0279\u0003h4\u0000"+
		"\u0279\u027a\u00063\uffff\uffff\u0000\u027a\u027c\u0001\u0000\u0000\u0000"+
		"\u027b\u0274\u0001\u0000\u0000\u0000\u027b\u0278\u0001\u0000\u0000\u0000"+
		"\u027cg\u0001\u0000\u0000\u0000\u027d\u027e\u0003j5\u0000\u027e\u027f"+
		"\u00064\uffff\uffff\u0000\u027f\u0287\u0001\u0000\u0000\u0000\u0280\u0281"+
		"\u0003n7\u0000\u0281\u0282\u00064\uffff\uffff\u0000\u0282\u0287\u0001"+
		"\u0000\u0000\u0000\u0283\u0284\u0003l6\u0000\u0284\u0285\u00064\uffff"+
		"\uffff\u0000\u0285\u0287\u0001\u0000\u0000\u0000\u0286\u027d\u0001\u0000"+
		"\u0000\u0000\u0286\u0280\u0001\u0000\u0000\u0000\u0286\u0283\u0001\u0000"+
		"\u0000\u0000\u0287i\u0001\u0000\u0000\u0000\u0288\u0289\u00065\uffff\uffff"+
		"\u0000\u0289\u028a\u0005;\u0000\u0000\u028a\u028b\u0003j5\u0003\u028b"+
		"\u028c\u00065\uffff\uffff\u0000\u028c\u0291\u0001\u0000\u0000\u0000\u028d"+
		"\u028e\u0003l6\u0000\u028e\u028f\u00065\uffff\uffff\u0000\u028f\u0291"+
		"\u0001\u0000\u0000\u0000\u0290\u0288\u0001\u0000\u0000\u0000\u0290\u028d"+
		"\u0001\u0000\u0000\u0000\u0291\u0299\u0001\u0000\u0000\u0000\u0292\u0293"+
		"\n\u0002\u0000\u0000\u0293\u0294\u0007\u0000\u0000\u0000\u0294\u0295\u0003"+
		"j5\u0003\u0295\u0296\u00065\uffff\uffff\u0000\u0296\u0298\u0001\u0000"+
		"\u0000\u0000\u0297\u0292\u0001\u0000\u0000\u0000\u0298\u029b\u0001\u0000"+
		"\u0000\u0000\u0299\u0297\u0001\u0000\u0000\u0000\u0299\u029a\u0001\u0000"+
		"\u0000\u0000\u029ak\u0001\u0000\u0000\u0000\u029b\u0299\u0001\u0000\u0000"+
		"\u0000\u029c\u029d\u00066\uffff\uffff\u0000\u029d\u029e\u0003n7\u0000"+
		"\u029e\u029f\u00066\uffff\uffff\u0000\u029f\u02a7\u0001\u0000\u0000\u0000"+
		"\u02a0\u02a1\n\u0002\u0000\u0000\u02a1\u02a2\u0007\u0001\u0000\u0000\u02a2"+
		"\u02a3\u0003l6\u0003\u02a3\u02a4\u00066\uffff\uffff\u0000\u02a4\u02a6"+
		"\u0001\u0000\u0000\u0000\u02a5\u02a0\u0001\u0000\u0000\u0000\u02a6\u02a9"+
		"\u0001\u0000\u0000\u0000\u02a7\u02a5\u0001\u0000\u0000\u0000\u02a7\u02a8"+
		"\u0001\u0000\u0000\u0000\u02a8m\u0001\u0000\u0000\u0000\u02a9\u02a7\u0001"+
		"\u0000\u0000\u0000\u02aa\u02ab\u00067\uffff\uffff\u0000\u02ab\u02ac\u0005"+
		"0\u0000\u0000\u02ac\u02ad\u0003n7\u0005\u02ad\u02ae\u00067\uffff\uffff"+
		"\u0000\u02ae\u02b8\u0001\u0000\u0000\u0000\u02af\u02b0\u0003p8\u0000\u02b0"+
		"\u02b1\u00067\uffff\uffff\u0000\u02b1\u02b8\u0001\u0000\u0000\u0000\u02b2"+
		"\u02b3\u00051\u0000\u0000\u02b3\u02b4\u0003h4\u0000\u02b4\u02b5\u0005"+
		"2\u0000\u0000\u02b5\u02b6\u00067\uffff\uffff\u0000\u02b6\u02b8\u0001\u0000"+
		"\u0000\u0000\u02b7\u02aa\u0001\u0000\u0000\u0000\u02b7\u02af\u0001\u0000"+
		"\u0000\u0000\u02b7\u02b2\u0001\u0000\u0000\u0000\u02b8\u02c5\u0001\u0000"+
		"\u0000\u0000\u02b9\u02ba\n\u0004\u0000\u0000\u02ba\u02bb\u0007\u0002\u0000"+
		"\u0000\u02bb\u02bc\u0003n7\u0005\u02bc\u02bd\u00067\uffff\uffff\u0000"+
		"\u02bd\u02c4\u0001\u0000\u0000\u0000\u02be\u02bf\n\u0003\u0000\u0000\u02bf"+
		"\u02c0\u0007\u0003\u0000\u0000\u02c0\u02c1\u0003n7\u0004\u02c1\u02c2\u0006"+
		"7\uffff\uffff\u0000\u02c2\u02c4\u0001\u0000\u0000\u0000\u02c3\u02b9\u0001"+
		"\u0000\u0000\u0000\u02c3\u02be\u0001\u0000\u0000\u0000\u02c4\u02c7\u0001"+
		"\u0000\u0000\u0000\u02c5\u02c3\u0001\u0000\u0000\u0000\u02c5\u02c6\u0001"+
		"\u0000\u0000\u0000\u02c6o\u0001\u0000\u0000\u0000\u02c7\u02c5\u0001\u0000"+
		"\u0000\u0000\u02c8\u02c9\u0003r9\u0000\u02c9\u02ca\u00068\uffff\uffff"+
		"\u0000\u02caq\u0001\u0000\u0000\u0000\u02cb\u02cc\u0005\u0018\u0000\u0000"+
		"\u02cc\u02dd\u00069\uffff\uffff\u0000\u02cd\u02ce\u0005\u0019\u0000\u0000"+
		"\u02ce\u02dd\u00069\uffff\uffff\u0000\u02cf\u02d0\u0005\u001b\u0000\u0000"+
		"\u02d0\u02dd\u00069\uffff\uffff\u0000\u02d1\u02d2\u0005\u001c\u0000\u0000"+
		"\u02d2\u02dd\u00069\uffff\uffff\u0000\u02d3\u02d4\u0005\u001a\u0000\u0000"+
		"\u02d4\u02dd\u00069\uffff\uffff\u0000\u02d5\u02dd\u0003F#\u0000\u02d6"+
		"\u02dd\u0003Z-\u0000\u02d7\u02dd\u0003N\'\u0000\u02d8\u02d9\u0005\u001d"+
		"\u0000\u0000\u02d9\u02dd\u00069\uffff\uffff\u0000\u02da\u02dd\u0003\u0010"+
		"\b\u0000\u02db\u02dd\u0003$\u0012\u0000\u02dc\u02cb\u0001\u0000\u0000"+
		"\u0000\u02dc\u02cd\u0001\u0000\u0000\u0000\u02dc\u02cf\u0001\u0000\u0000"+
		"\u0000\u02dc\u02d1\u0001\u0000\u0000\u0000\u02dc\u02d3\u0001\u0000\u0000"+
		"\u0000\u02dc\u02d5\u0001\u0000\u0000\u0000\u02dc\u02d6\u0001\u0000\u0000"+
		"\u0000\u02dc\u02d7\u0001\u0000\u0000\u0000\u02dc\u02d8\u0001\u0000\u0000"+
		"\u0000\u02dc\u02da\u0001\u0000\u0000\u0000\u02dc\u02db\u0001\u0000\u0000"+
		"\u0000\u02dds\u0001\u0000\u0000\u0000.{\u0095\u00ab\u00d5\u00f4\u00f9"+
		"\u0114\u0119\u0129\u012e\u013b\u0140\u0146\u0155\u015a\u0169\u016e\u017a"+
		"\u017f\u0185\u0192\u01ac\u01bd\u01ea\u01ef\u01f9\u0203\u020d\u0218\u0222"+
		"\u022a\u0240\u0245\u024f\u0257\u026b\u0270\u027b\u0286\u0290\u0299\u02a7"+
		"\u02b7\u02c3\u02c5\u02dc";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}