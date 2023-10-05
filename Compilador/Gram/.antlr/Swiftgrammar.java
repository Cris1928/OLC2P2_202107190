// Generated from c:\OLC2P2_202107190\Compilador\Gram\Swiftgrammar.g4 by ANTLR 4.9.2


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

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Swiftgrammar extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

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
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRINT) | (1L << VAR) | (1L << LET) | (1L << IF) | (1L << FOR) | (1L << WHILE) | (1L << SWITCH) | (1L << BREAK) | (1L << CONTINUE) | (1L << RETURN) | (1L << FUNC) | (1L << STRUCT) | (1L << ID))) != 0) );

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

	public static class InstruccionContext extends ParserRuleContext {
		public interfaces.Instruction instr;
		public Instruccion_printlnContext instruccion_println;
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
			setState(145);
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
				instruccion_declaracion();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(132);
				instruccion_asignacion();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(133);
				instr_structs_assignment();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(134);
				instruccion_if();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(135);
				instruccion_for_in();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(136);
				instruccion_while();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(137);
				instruccion_while_true();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(138);
				instruccion_switch();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(139);
				instruccion_break();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(140);
				instruccion_continue();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(141);
				instruccion_func();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(142);
				instruccion_llamada();
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(143);
				instruccion_return();
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(144);
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
			setState(167);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(147);
				((Instruccion_printlnContext)_localctx).PRINT = match(PRINT);
				setState(148);
				match(TK_PARA);
				setState(149);
				((Instruccion_printlnContext)_localctx).primitivo = primitivo();
				setState(150);
				match(TK_PARC);
				 _localctx.instr = instruction.NewPrintln(nil, ((Instruccion_printlnContext)_localctx).primitivo.p,       "-1",         (((Instruccion_printlnContext)_localctx).PRINT!=null?((Instruccion_printlnContext)_localctx).PRINT.getLine():0), localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn()) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(153);
				((Instruccion_printlnContext)_localctx).PRINT = match(PRINT);
				setState(154);
				match(TK_PARA);
				setState(155);
				((Instruccion_printlnContext)_localctx).STRING = match(STRING);
				setState(156);
				match(TK_COMA);
				setState(157);
				((Instruccion_printlnContext)_localctx).list_expression = list_expression();
				setState(158);
				match(TK_PARC);
				 _localctx.instr = instruction.NewPrintln(((Instruccion_printlnContext)_localctx).list_expression.l, nil, (((Instruccion_printlnContext)_localctx).STRING!=null?((Instruccion_printlnContext)_localctx).STRING.getText():null)[1:len((((Instruccion_printlnContext)_localctx).STRING!=null?((Instruccion_printlnContext)_localctx).STRING.getText():null))-1], (((Instruccion_printlnContext)_localctx).PRINT!=null?((Instruccion_printlnContext)_localctx).PRINT.getLine():0), localctx.(*Instruccion_printlnContext).Get_PRINT().GetColumn()) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(161);
				((Instruccion_printlnContext)_localctx).PRINT = match(PRINT);
				setState(162);
				match(TK_PARA);
				setState(163);
				((Instruccion_printlnContext)_localctx).expressions = expressions();
				setState(164);
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

	public static class Instruccion_declaracionContext extends ParserRuleContext {
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
			setState(199);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(169);
				match(VAR);
				setState(170);
				match(ID);
				setState(171);
				match(TK_IGUAL);
				setState(172);
				expressions();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(173);
				match(VAR);
				setState(174);
				match(ID);
				setState(175);
				match(TK_DOSPUNTOS);
				setState(176);
				instruccion_tipo();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(177);
				match(VAR);
				setState(178);
				match(ID);
				setState(179);
				match(TK_DOSPUNTOS);
				setState(180);
				instruccion_tipo();
				setState(181);
				match(TK_IGUAL);
				setState(182);
				expressions();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(184);
				match(LET);
				setState(185);
				match(ID);
				setState(186);
				match(TK_IGUAL);
				setState(187);
				expressions();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(188);
				match(LET);
				setState(189);
				match(ID);
				setState(190);
				match(TK_DOSPUNTOS);
				setState(191);
				instruccion_tipo();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(192);
				match(LET);
				setState(193);
				match(ID);
				setState(194);
				match(TK_DOSPUNTOS);
				setState(195);
				instruccion_tipo();
				setState(196);
				match(TK_IGUAL);
				setState(197);
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

	public static class Instruccion_asignacionContext extends ParserRuleContext {
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
			setState(201);
			match(ID);
			setState(202);
			match(TK_IGUAL);
			setState(203);
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
			setState(229);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(205);
				match(IF);
				setState(206);
				expressions();
				setState(207);
				match(TK_LLAVEA);
				setState(208);
				((Instruccion_ifContext)_localctx).left_instr = instrucciones();
				setState(209);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(211);
				match(IF);
				setState(212);
				expressions();
				setState(213);
				match(TK_LLAVEA);
				setState(214);
				((Instruccion_ifContext)_localctx).left_instr = instrucciones();
				setState(215);
				match(TK_LLAVEC);
				setState(216);
				match(ELSE);
				setState(217);
				match(TK_LLAVEA);
				setState(218);
				((Instruccion_ifContext)_localctx).right_instr = instrucciones();
				setState(219);
				match(TK_LLAVEC);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(221);
				match(IF);
				setState(222);
				expressions();
				setState(223);
				match(TK_LLAVEA);
				setState(224);
				((Instruccion_ifContext)_localctx).left_instr = instrucciones();
				setState(225);
				match(TK_LLAVEC);
				setState(226);
				match(ELSE);
				setState(227);
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
			setState(234);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(231);
					((Instr_else_ifContext)_localctx).instruccion_if = instruccion_if();
					((Instr_else_ifContext)_localctx).e.add(((Instr_else_ifContext)_localctx).instruccion_if);
					}
					} 
				}
				setState(236);
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
			setState(261);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(237);
				match(IF);
				setState(238);
				((Instruccion_ternarioContext)_localctx).cond = expressions();
				setState(239);
				match(TK_LLAVEA);
				setState(240);
				((Instruccion_ternarioContext)_localctx).left_instr = expressions();
				setState(241);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(243);
				match(IF);
				setState(244);
				((Instruccion_ternarioContext)_localctx).cond = expressions();
				setState(245);
				match(TK_LLAVEA);
				setState(246);
				((Instruccion_ternarioContext)_localctx).left_instr = expressions();
				setState(247);
				match(TK_LLAVEC);
				setState(248);
				match(ELSE);
				setState(249);
				match(TK_LLAVEA);
				setState(250);
				((Instruccion_ternarioContext)_localctx).right_instr = expressions();
				setState(251);
				match(TK_LLAVEC);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(253);
				match(IF);
				setState(254);
				((Instruccion_ternarioContext)_localctx).cond = expressions();
				setState(255);
				match(TK_LLAVEA);
				setState(256);
				((Instruccion_ternarioContext)_localctx).left_instr = expressions();
				setState(257);
				match(TK_LLAVEC);
				setState(258);
				match(ELSE);
				setState(259);
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
			setState(266);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(263);
					((Instr_else_if_ternarioContext)_localctx).instruccion_ternario = instruccion_ternario();
					((Instr_else_if_ternarioContext)_localctx).e.add(((Instr_else_if_ternarioContext)_localctx).instruccion_ternario);
					}
					} 
				}
				setState(268);
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
			setState(282);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(269);
				match(SWITCH);
				setState(270);
				expressions();
				setState(271);
				match(TK_LLAVEA);
				setState(272);
				list_case();
				setState(273);
				block_default();
				setState(274);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(276);
				match(SWITCH);
				setState(277);
				expressions();
				setState(278);
				match(TK_LLAVEA);
				setState(279);
				block_default();
				setState(280);
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
			setState(285); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(284);
				((List_caseContext)_localctx).instruccion_case = instruccion_case();
				((List_caseContext)_localctx).e.add(((List_caseContext)_localctx).instruccion_case);
				}
				}
				setState(287); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IF) | (1L << SWITCH) | (1L << NUMBER) | (1L << DOUBLE) | (1L << CHAR) | (1L << STRING) | (1L << BOOLEAN) | (1L << ID) | (1L << TK_MENOS) | (1L << TK_PARA) | (1L << TK_NOT))) != 0) );
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
			setState(300);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(289);
				list_expre_case();
				setState(290);
				match(TK_DOSPUNTOS);
				setState(291);
				match(TK_LLAVEA);
				setState(292);
				instrucciones();
				setState(293);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(295);
				list_expre_case();
				setState(296);
				match(TK_DOSPUNTOS);
				setState(297);
				block_instr_switch();
				setState(298);
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
			setState(303); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(302);
				((List_expre_caseContext)_localctx).block_case = block_case();
				((List_expre_caseContext)_localctx).e.add(((List_expre_caseContext)_localctx).block_case);
				}
				}
				setState(305); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IF) | (1L << SWITCH) | (1L << NUMBER) | (1L << DOUBLE) | (1L << CHAR) | (1L << STRING) | (1L << BOOLEAN) | (1L << ID) | (1L << TK_MENOS) | (1L << TK_PARA) | (1L << TK_NOT))) != 0) );
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
			setState(311);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(307);
				expressions();
				setState(308);
				match(TK_BARRA);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(310);
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
			setState(313);
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
			setState(326);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(315);
				match(CASE);
				setState(316);
				match(ID);
				setState(317);
				match(TK_DOSPUNTOS);
				setState(318);
				match(TK_LLAVEA);
				setState(319);
				instrucciones();
				setState(320);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(322);
				match(CASE);
				setState(323);
				match(ID);
				setState(324);
				match(TK_DOSPUNTOS);
				setState(325);
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
			setState(329); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(328);
				((Block_defaultContext)_localctx).instr_default = instr_default();
				((Block_defaultContext)_localctx).e.add(((Block_defaultContext)_localctx).instr_default);
				}
				}
				setState(331); 
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
			setState(346);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(333);
				match(SWITCH);
				setState(334);
				expressions();
				setState(335);
				match(TK_LLAVEA);
				setState(336);
				list_case_ternario();
				setState(337);
				instr_default_ter();
				setState(338);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(340);
				match(SWITCH);
				setState(341);
				expressions();
				setState(342);
				match(TK_LLAVEA);
				setState(343);
				instr_default_ter();
				setState(344);
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
			setState(349); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(348);
				((List_case_ternarioContext)_localctx).instr_case_ter = instr_case_ter();
				((List_case_ternarioContext)_localctx).e.add(((List_case_ternarioContext)_localctx).instr_case_ter);
				}
				}
				setState(351); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IF) | (1L << SWITCH) | (1L << NUMBER) | (1L << DOUBLE) | (1L << CHAR) | (1L << STRING) | (1L << BOOLEAN) | (1L << ID) | (1L << TK_MENOS) | (1L << TK_PARA) | (1L << TK_NOT))) != 0) );
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
			setState(363);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(353);
				list_expre_case_ter();
				setState(354);
				match(TK_DOSPUNTOS);
				setState(355);
				expressions();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(357);
				list_expre_case_ter();
				setState(358);
				match(TK_DOSPUNTOS);
				setState(359);
				match(TK_LLAVEA);
				setState(360);
				expressions();
				setState(361);
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
			setState(366); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(365);
				((List_expre_case_terContext)_localctx).block_case_ter = block_case_ter();
				((List_expre_case_terContext)_localctx).e.add(((List_expre_case_terContext)_localctx).block_case_ter);
				}
				}
				setState(368); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IF) | (1L << SWITCH) | (1L << NUMBER) | (1L << DOUBLE) | (1L << CHAR) | (1L << STRING) | (1L << BOOLEAN) | (1L << ID) | (1L << TK_MENOS) | (1L << TK_PARA) | (1L << TK_NOT))) != 0) );
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
			setState(374);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(370);
				expressions();
				setState(371);
				match(TK_BARRA);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(373);
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
			setState(387);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(376);
				match(CASE);
				setState(377);
				match(ID);
				setState(378);
				match(TK_DOSPUNTOS);
				setState(379);
				expressions();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(380);
				match(CASE);
				setState(381);
				match(ID);
				setState(382);
				match(TK_DOSPUNTOS);
				setState(383);
				match(TK_LLAVEA);
				setState(384);
				expressions();
				setState(385);
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
			setState(389);
			match(WHILE);
			setState(390);
			expressions();
			setState(391);
			match(TK_LLAVEA);
			setState(392);
			instrucciones();
			setState(393);
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
			setState(413);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(395);
				match(FOR);
				setState(396);
				match(ID);
				setState(397);
				match(IN);
				setState(398);
				((Instruccion_for_inContext)_localctx).left = expressions();
				setState(399);
				match(TK_TRIPLEPUNTO);
				setState(400);
				((Instruccion_for_inContext)_localctx).right = expressions();
				setState(401);
				match(TK_LLAVEA);
				setState(402);
				instrucciones();
				setState(403);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(405);
				match(FOR);
				setState(406);
				match(ID);
				setState(407);
				match(IN);
				setState(408);
				((Instruccion_for_inContext)_localctx).left = expressions();
				setState(409);
				match(TK_LLAVEA);
				setState(410);
				instrucciones();
				setState(411);
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
			setState(415);
			match(WHILE);
			setState(416);
			match(TRUE);
			setState(417);
			match(TK_LLAVEA);
			setState(418);
			instrucciones();
			setState(419);
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
			setState(421);
			match(WHILE);
			setState(422);
			match(TRUE);
			setState(423);
			match(TK_LLAVEA);
			setState(424);
			instrucciones();
			setState(425);
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
			setState(430);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(427);
				match(BREAK);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(428);
				match(BREAK);
				setState(429);
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
			setState(432);
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
			setState(434);
			match(RETURN);
			setState(435);
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
			setState(475);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(437);
				match(FUNC);
				setState(438);
				match(ID);
				setState(439);
				match(TK_PARA);
				setState(440);
				match(TK_PARC);
				setState(441);
				match(TK_LLAVEA);
				setState(442);
				instrucciones();
				setState(443);
				match(TK_LLAVEC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(445);
				match(FUNC);
				setState(446);
				match(ID);
				setState(447);
				match(TK_PARA);
				setState(448);
				match(TK_PARC);
				setState(449);
				match(TK_MENOSMAYOR);
				setState(450);
				instruccion_tipo();
				setState(451);
				match(TK_LLAVEA);
				setState(452);
				instrucciones();
				setState(453);
				match(TK_LLAVEC);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(455);
				match(FUNC);
				setState(456);
				match(ID);
				setState(457);
				match(TK_PARA);
				setState(458);
				list_function_parameters();
				setState(459);
				match(TK_PARC);
				setState(460);
				match(TK_LLAVEA);
				setState(461);
				instrucciones();
				setState(462);
				match(TK_LLAVEC);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(464);
				match(FUNC);
				setState(465);
				match(ID);
				setState(466);
				match(TK_PARA);
				setState(467);
				list_function_parameters();
				setState(468);
				match(TK_PARC);
				setState(469);
				match(TK_MENOSMAYOR);
				setState(470);
				instruccion_tipo();
				setState(471);
				match(TK_LLAVEA);
				setState(472);
				instrucciones();
				setState(473);
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
			setState(478); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(477);
				((List_function_parametersContext)_localctx).block_parameters_fn = block_parameters_fn();
				((List_function_parametersContext)_localctx).e.add(((List_function_parametersContext)_localctx).block_parameters_fn);
				}
				}
				setState(480); 
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
			setState(490);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(482);
				match(ID);
				setState(483);
				match(TK_DOSPUNTOS);
				setState(484);
				instruccion_tipo();
				setState(485);
				match(TK_COMA);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(487);
				match(ID);
				setState(488);
				match(TK_DOSPUNTOS);
				setState(489);
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
			setState(500);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(492);
				match(ID);
				setState(493);
				match(TK_PARA);
				setState(494);
				match(TK_PARC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(495);
				match(ID);
				setState(496);
				match(TK_PARA);
				setState(497);
				list_expression();
				setState(498);
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
			setState(510);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(502);
				match(ID);
				setState(503);
				match(TK_PARA);
				setState(504);
				match(TK_PARC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(505);
				match(ID);
				setState(506);
				match(TK_PARA);
				setState(507);
				list_expression();
				setState(508);
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
			setState(512);
			match(STRUCT);
			setState(513);
			match(ID);
			setState(514);
			match(TK_LLAVEA);
			setState(515);
			list_struct_parameters();
			setState(516);
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
			setState(519); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(518);
				((List_struct_parametersContext)_localctx).block_structs_parameters = block_structs_parameters();
				((List_struct_parametersContext)_localctx).e.add(((List_struct_parametersContext)_localctx).block_structs_parameters);
				}
				}
				setState(521); 
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
			setState(531);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(523);
				match(ID);
				setState(524);
				match(TK_DOSPUNTOS);
				setState(525);
				instruccion_tipo();
				setState(526);
				match(TK_COMA);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(528);
				match(ID);
				setState(529);
				match(TK_DOSPUNTOS);
				setState(530);
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
			setState(533);
			match(ID);
			setState(534);
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
			setState(537); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(536);
					((List_arrays_parameters_idContext)_localctx).block_arrays_identifier = block_arrays_identifier();
					((List_arrays_parameters_idContext)_localctx).e.add(((List_arrays_parameters_idContext)_localctx).block_arrays_identifier);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(539); 
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
			setState(541);
			match(TK_CORA);
			setState(542);
			expressions();
			setState(543);
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
			setState(561);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(545);
				match(VAR);
				setState(546);
				((Instr_structs_declarationContext)_localctx).left = match(ID);
				setState(547);
				match(TK_IGUAL);
				setState(548);
				((Instr_structs_declarationContext)_localctx).right = match(ID);
				setState(549);
				match(TK_LLAVEA);
				setState(550);
				list_struct_parameters_decla();
				setState(551);
				match(TK_LLAVEC);
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(553);
				match(LET);
				setState(554);
				((Instr_structs_declarationContext)_localctx).left = match(ID);
				setState(555);
				match(TK_IGUAL);
				setState(556);
				((Instr_structs_declarationContext)_localctx).right = match(ID);
				setState(557);
				match(TK_LLAVEA);
				setState(558);
				list_struct_parameters_decla();
				setState(559);
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
			setState(564); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(563);
				((List_struct_parameters_declaContext)_localctx).block_structs_parameters_decla = block_structs_parameters_decla();
				((List_struct_parameters_declaContext)_localctx).e.add(((List_struct_parameters_declaContext)_localctx).block_structs_parameters_decla);
				}
				}
				setState(566); 
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
			setState(576);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(568);
				match(ID);
				setState(569);
				match(TK_DOSPUNTOS);
				setState(570);
				expressions();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(571);
				match(ID);
				setState(572);
				match(TK_DOSPUNTOS);
				setState(573);
				expressions();
				setState(574);
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
			setState(578);
			match(ID);
			setState(579);
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
			setState(582); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(581);
					((List_struct_parameters_idContext)_localctx).block_structs_identifier = block_structs_identifier();
					((List_struct_parameters_idContext)_localctx).e.add(((List_struct_parameters_idContext)_localctx).block_structs_identifier);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(584); 
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
			setState(586);
			match(TK_PUNTO);
			setState(587);
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
			setState(589);
			match(ID);
			setState(590);
			list_struct_parameters_id();
			setState(591);
			match(TK_IGUAL);
			setState(592);
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

	public static class Instruccion_tipoContext extends ParserRuleContext {
		public TerminalNode R_INT() { return getToken(Swiftgrammar.R_INT, 0); }
		public TerminalNode R_FLOAT() { return getToken(Swiftgrammar.R_FLOAT, 0); }
		public TerminalNode R_STRING() { return getToken(Swiftgrammar.R_STRING, 0); }
		public TerminalNode R_BOOL() { return getToken(Swiftgrammar.R_BOOL, 0); }
		public TerminalNode R_CHAR() { return getToken(Swiftgrammar.R_CHAR, 0); }
		public TerminalNode TK_AMPERSAND() { return getToken(Swiftgrammar.TK_AMPERSAND, 0); }
		public TerminalNode ID() { return getToken(Swiftgrammar.ID, 0); }
		public Instruccion_tipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_tipo; }
	}

	public final Instruccion_tipoContext instruccion_tipo() throws RecognitionException {
		Instruccion_tipoContext _localctx = new Instruccion_tipoContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_instruccion_tipo);
		try {
			setState(601);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case R_INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(594);
				match(R_INT);
				}
				break;
			case R_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(595);
				match(R_FLOAT);
				}
				break;
			case R_STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(596);
				match(R_STRING);
				}
				break;
			case R_BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(597);
				match(R_BOOL);
				}
				break;
			case R_CHAR:
				enterOuterAlt(_localctx, 5);
				{
				setState(598);
				match(R_CHAR);
				}
				break;
			case TK_AMPERSAND:
				enterOuterAlt(_localctx, 6);
				{
				setState(599);
				match(TK_AMPERSAND);
				setState(600);
				match(ID);
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
			setState(604); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(603);
				((List_expressionContext)_localctx).block_list_expression = block_list_expression();
				((List_expressionContext)_localctx).e.add(((List_expressionContext)_localctx).block_list_expression);
				}
				}
				setState(606); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IF) | (1L << SWITCH) | (1L << NUMBER) | (1L << DOUBLE) | (1L << CHAR) | (1L << STRING) | (1L << BOOLEAN) | (1L << ID) | (1L << TK_MENOS) | (1L << TK_PARA) | (1L << TK_NOT))) != 0) );

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
			setState(617);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(610);
				((Block_list_expressionContext)_localctx).expressions = expressions();
				setState(611);
				match(TK_COMA);
				 _localctx.p =  instruction.NewListExpre(((Block_list_expressionContext)_localctx).expressions.p) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(614);
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
			setState(628);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(619);
				((ExpressionsContext)_localctx).expre_log = expre_log(0);
				 _localctx.p = ((ExpressionsContext)_localctx).expre_log.p 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(622);
				((ExpressionsContext)_localctx).expre_arit = expre_arit(0);
				 _localctx.p = ((ExpressionsContext)_localctx).expre_arit.p 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(625);
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
			setState(638);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_NOT:
				{
				setState(631);
				((Expre_logContext)_localctx).op = match(TK_NOT);
				setState(632);
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
				setState(635);
				((Expre_logContext)_localctx).expre_rel = expre_rel(0);
				 _localctx.p = ((Expre_logContext)_localctx).expre_rel.p 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(647);
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
					_localctx.left = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_expre_log);
					setState(640);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(641);
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
					setState(642);
					((Expre_logContext)_localctx).right = expre_log(3);
					 _localctx.p = expression.NewOperacion(((Expre_logContext)_localctx).left.p, (((Expre_logContext)_localctx).op!=null?((Expre_logContext)_localctx).op.getText():null), ((Expre_logContext)_localctx).right.p, false, (((Expre_logContext)_localctx).op!=null?((Expre_logContext)_localctx).op.getLine():0), localctx.(*Expre_logContext).GetOp().GetColumn()) 
					}
					} 
				}
				setState(649);
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
			setState(651);
			((Expre_relContext)_localctx).expre_arit = expre_arit(0);
			 _localctx.p = ((Expre_relContext)_localctx).expre_arit.p 
			}
			_ctx.stop = _input.LT(-1);
			setState(661);
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
					_localctx.left = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_expre_rel);
					setState(654);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(655);
					((Expre_relContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_IGUALIGUAL) | (1L << TK_MAYORIGUAL) | (1L << TK_MENORIGUAL) | (1L << TK_DIFIGUAL) | (1L << TK_MAYOR) | (1L << TK_MENOR))) != 0)) ) {
						((Expre_relContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(656);
					((Expre_relContext)_localctx).right = expre_rel(3);
					 _localctx.p = expression.NewOperacion(((Expre_relContext)_localctx).left.p, (((Expre_relContext)_localctx).op!=null?((Expre_relContext)_localctx).op.getText():null), ((Expre_relContext)_localctx).right.p, false, (((Expre_relContext)_localctx).op!=null?((Expre_relContext)_localctx).op.getLine():0), localctx.(*Expre_relContext).GetOp().GetColumn()) 
					}
					} 
				}
				setState(663);
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
			setState(677);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_MENOS:
				{
				setState(665);
				((Expre_aritContext)_localctx).op = match(TK_MENOS);
				setState(666);
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
				setState(669);
				((Expre_aritContext)_localctx).expre_valor = expre_valor();
				 _localctx.p = ((Expre_aritContext)_localctx).expre_valor.p 
				}
				break;
			case TK_PARA:
				{
				setState(672);
				match(TK_PARA);
				setState(673);
				((Expre_aritContext)_localctx).expressions = expressions();
				setState(674);
				match(TK_PARC);
				 _localctx.p = ((Expre_aritContext)_localctx).expressions.p 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(691);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,44,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(689);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
					case 1:
						{
						_localctx = new Expre_aritContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expre_arit);
						setState(679);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(680);
						((Expre_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_MULT) | (1L << TK_DIV) | (1L << TK_MODULO))) != 0)) ) {
							((Expre_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(681);
						((Expre_aritContext)_localctx).right = expre_arit(5);
						 _localctx.p = expression.NewOperacion(((Expre_aritContext)_localctx).left.p, (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getText():null), ((Expre_aritContext)_localctx).right.p, false, (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getLine():0), localctx.(*Expre_aritContext).GetOp().GetColumn()) 
						}
						break;
					case 2:
						{
						_localctx = new Expre_aritContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expre_arit);
						setState(684);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(685);
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
						setState(686);
						((Expre_aritContext)_localctx).right = expre_arit(4);
						 _localctx.p = expression.NewOperacion(((Expre_aritContext)_localctx).left.p, (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getText():null), ((Expre_aritContext)_localctx).right.p, false, (((Expre_aritContext)_localctx).op!=null?((Expre_aritContext)_localctx).op.getLine():0), localctx.(*Expre_aritContext).GetOp().GetColumn()) 
						}
						break;
					}
					} 
				}
				setState(693);
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
			setState(694);
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

	public static class PrimitivoContext extends ParserRuleContext {
		public interfaces.Expression p;
		public Token NUMBER;
		public Token DOUBLE;
		public Token STRING;
		public Token BOOLEAN;
		public Token CHAR;
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
			setState(713);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,45,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(697);
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
				setState(699);
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
				setState(701);
				((PrimitivoContext)_localctx).STRING = match(STRING);
				 
				              str:= (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null)[1:len((((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null))-1]
				              _localctx.p = expression.NewPrimitivo(str, interfaces.STRING, interfaces.NULL, (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getLine():0), localctx.(*PrimitivoContext).Get_STRING().GetColumn())
				            
				            
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(703);
				((PrimitivoContext)_localctx).BOOLEAN = match(BOOLEAN);
				 
				          
				              exp,_ := strconv.ParseBool((((PrimitivoContext)_localctx).BOOLEAN!=null?((PrimitivoContext)_localctx).BOOLEAN.getText():null))
				              _localctx.p = expression.NewPrimitivo(exp, interfaces.BOOLEAN, interfaces.NULL, (((PrimitivoContext)_localctx).BOOLEAN!=null?((PrimitivoContext)_localctx).BOOLEAN.getLine():0), localctx.(*PrimitivoContext).Get_BOOLEAN().GetColumn())
				            
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(705);
				((PrimitivoContext)_localctx).CHAR = match(CHAR);


				            str:= (((PrimitivoContext)_localctx).CHAR!=null?((PrimitivoContext)_localctx).CHAR.getText():null)[1]
				            _localctx.p = expression.NewPrimitivo(string(str), interfaces.CHAR, interfaces.NULL, (((PrimitivoContext)_localctx).CHAR!=null?((PrimitivoContext)_localctx).CHAR.getLine():0), localctx.(*PrimitivoContext).Get_CHAR().GetColumn())
				          
				          
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(707);
				instr_llamada_expre();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(708);
				instr_structs_identifier();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(709);
				instr_arrays_identifier();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(710);
				match(ID);
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(711);
				instruccion_ternario();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(712);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3@\u02ce\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\3\2\3\2\3\2"+
		"\3\2\3\3\6\3|\n\3\r\3\16\3}\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\u0094\n\4\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5\u00aa"+
		"\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u00ca\n\6\3"+
		"\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u00e8\n\b\3\t\7\t\u00eb\n"+
		"\t\f\t\16\t\u00ee\13\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u0108\n\n\3\13\7"+
		"\13\u010b\n\13\f\13\16\13\u010e\13\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\5\f\u011d\n\f\3\r\6\r\u0120\n\r\r\r\16\r\u0121\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u012f\n\16"+
		"\3\17\6\17\u0132\n\17\r\17\16\17\u0133\3\20\3\20\3\20\3\20\5\20\u013a"+
		"\n\20\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\5\22\u0149\n\22\3\23\6\23\u014c\n\23\r\23\16\23\u014d\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u015d\n\24\3\25"+
		"\6\25\u0160\n\25\r\25\16\25\u0161\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\3\26\3\26\5\26\u016e\n\26\3\27\6\27\u0171\n\27\r\27\16\27\u0172\3"+
		"\30\3\30\3\30\3\30\5\30\u0179\n\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\5\31\u0186\n\31\3\32\3\32\3\32\3\32\3\32\3\32\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\5\33\u01a0\n\33\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3\35"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\5\36\u01b1\n\36\3\37\3\37\3 \3 \3"+
		" \3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\5!\u01de\n!\3\"\6\"\u01e1"+
		"\n\"\r\"\16\"\u01e2\3#\3#\3#\3#\3#\3#\3#\3#\5#\u01ed\n#\3$\3$\3$\3$\3"+
		"$\3$\3$\3$\5$\u01f7\n$\3%\3%\3%\3%\3%\3%\3%\3%\5%\u0201\n%\3&\3&\3&\3"+
		"&\3&\3&\3\'\6\'\u020a\n\'\r\'\16\'\u020b\3(\3(\3(\3(\3(\3(\3(\3(\5(\u0216"+
		"\n(\3)\3)\3)\3*\6*\u021c\n*\r*\16*\u021d\3+\3+\3+\3+\3,\3,\3,\3,\3,\3"+
		",\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\5,\u0234\n,\3-\6-\u0237\n-\r-\16-\u0238"+
		"\3.\3.\3.\3.\3.\3.\3.\3.\5.\u0243\n.\3/\3/\3/\3\60\6\60\u0249\n\60\r\60"+
		"\16\60\u024a\3\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3"+
		"\63\3\63\3\63\3\63\5\63\u025c\n\63\3\64\6\64\u025f\n\64\r\64\16\64\u0260"+
		"\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\65\5\65\u026c\n\65\3\66\3\66"+
		"\3\66\3\66\3\66\3\66\3\66\3\66\3\66\5\66\u0277\n\66\3\67\3\67\3\67\3\67"+
		"\3\67\3\67\3\67\3\67\5\67\u0281\n\67\3\67\3\67\3\67\3\67\3\67\7\67\u0288"+
		"\n\67\f\67\16\67\u028b\13\67\38\38\38\38\38\38\38\38\38\78\u0296\n8\f"+
		"8\168\u0299\138\39\39\39\39\39\39\39\39\39\39\39\39\39\59\u02a8\n9\39"+
		"\39\39\39\39\39\39\39\39\39\79\u02b4\n9\f9\169\u02b7\139\3:\3:\3:\3;\3"+
		";\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\5;\u02cc\n;\3;\2\5lnp<\2\4"+
		"\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNP"+
		"RTVXZ\\^`bdfhjlnprt\2\6\4\299;;\4\2&\'*-\3\2.\60\3\2\61\62\2\u02e7\2v"+
		"\3\2\2\2\4{\3\2\2\2\6\u0093\3\2\2\2\b\u00a9\3\2\2\2\n\u00c9\3\2\2\2\f"+
		"\u00cb\3\2\2\2\16\u00e7\3\2\2\2\20\u00ec\3\2\2\2\22\u0107\3\2\2\2\24\u010c"+
		"\3\2\2\2\26\u011c\3\2\2\2\30\u011f\3\2\2\2\32\u012e\3\2\2\2\34\u0131\3"+
		"\2\2\2\36\u0139\3\2\2\2 \u013b\3\2\2\2\"\u0148\3\2\2\2$\u014b\3\2\2\2"+
		"&\u015c\3\2\2\2(\u015f\3\2\2\2*\u016d\3\2\2\2,\u0170\3\2\2\2.\u0178\3"+
		"\2\2\2\60\u0185\3\2\2\2\62\u0187\3\2\2\2\64\u019f\3\2\2\2\66\u01a1\3\2"+
		"\2\28\u01a7\3\2\2\2:\u01b0\3\2\2\2<\u01b2\3\2\2\2>\u01b4\3\2\2\2@\u01dd"+
		"\3\2\2\2B\u01e0\3\2\2\2D\u01ec\3\2\2\2F\u01f6\3\2\2\2H\u0200\3\2\2\2J"+
		"\u0202\3\2\2\2L\u0209\3\2\2\2N\u0215\3\2\2\2P\u0217\3\2\2\2R\u021b\3\2"+
		"\2\2T\u021f\3\2\2\2V\u0233\3\2\2\2X\u0236\3\2\2\2Z\u0242\3\2\2\2\\\u0244"+
		"\3\2\2\2^\u0248\3\2\2\2`\u024c\3\2\2\2b\u024f\3\2\2\2d\u025b\3\2\2\2f"+
		"\u025e\3\2\2\2h\u026b\3\2\2\2j\u0276\3\2\2\2l\u0280\3\2\2\2n\u028c\3\2"+
		"\2\2p\u02a7\3\2\2\2r\u02b8\3\2\2\2t\u02cb\3\2\2\2vw\5\4\3\2wx\7\2\2\3"+
		"xy\b\2\1\2y\3\3\2\2\2z|\5\6\4\2{z\3\2\2\2|}\3\2\2\2}{\3\2\2\2}~\3\2\2"+
		"\2~\177\3\2\2\2\177\u0080\b\3\1\2\u0080\5\3\2\2\2\u0081\u0082\5\b\5\2"+
		"\u0082\u0083\b\4\1\2\u0083\u0094\3\2\2\2\u0084\u0094\5J&\2\u0085\u0094"+
		"\5\n\6\2\u0086\u0094\5\f\7\2\u0087\u0094\5b\62\2\u0088\u0094\5\16\b\2"+
		"\u0089\u0094\5\64\33\2\u008a\u0094\5\62\32\2\u008b\u0094\5\66\34\2\u008c"+
		"\u0094\5\26\f\2\u008d\u0094\5:\36\2\u008e\u0094\5<\37\2\u008f\u0094\5"+
		"@!\2\u0090\u0094\5F$\2\u0091\u0094\5> \2\u0092\u0094\5V,\2\u0093\u0081"+
		"\3\2\2\2\u0093\u0084\3\2\2\2\u0093\u0085\3\2\2\2\u0093\u0086\3\2\2\2\u0093"+
		"\u0087\3\2\2\2\u0093\u0088\3\2\2\2\u0093\u0089\3\2\2\2\u0093\u008a\3\2"+
		"\2\2\u0093\u008b\3\2\2\2\u0093\u008c\3\2\2\2\u0093\u008d\3\2\2\2\u0093"+
		"\u008e\3\2\2\2\u0093\u008f\3\2\2\2\u0093\u0090\3\2\2\2\u0093\u0091\3\2"+
		"\2\2\u0093\u0092\3\2\2\2\u0094\7\3\2\2\2\u0095\u0096\7\3\2\2\u0096\u0097"+
		"\7\63\2\2\u0097\u0098\5t;\2\u0098\u0099\7\64\2\2\u0099\u009a\b\5\1\2\u009a"+
		"\u00aa\3\2\2\2\u009b\u009c\7\3\2\2\u009c\u009d\7\63\2\2\u009d\u009e\7"+
		"\35\2\2\u009e\u009f\7#\2\2\u009f\u00a0\5f\64\2\u00a0\u00a1\7\64\2\2\u00a1"+
		"\u00a2\b\5\1\2\u00a2\u00aa\3\2\2\2\u00a3\u00a4\7\3\2\2\u00a4\u00a5\7\63"+
		"\2\2\u00a5\u00a6\5j\66\2\u00a6\u00a7\7\64\2\2\u00a7\u00a8\b\5\1\2\u00a8"+
		"\u00aa\3\2\2\2\u00a9\u0095\3\2\2\2\u00a9\u009b\3\2\2\2\u00a9\u00a3\3\2"+
		"\2\2\u00aa\t\3\2\2\2\u00ab\u00ac\7\4\2\2\u00ac\u00ad\7\37\2\2\u00ad\u00ae"+
		"\7%\2\2\u00ae\u00ca\5j\66\2\u00af\u00b0\7\4\2\2\u00b0\u00b1\7\37\2\2\u00b1"+
		"\u00b2\7$\2\2\u00b2\u00ca\5d\63\2\u00b3\u00b4\7\4\2\2\u00b4\u00b5\7\37"+
		"\2\2\u00b5\u00b6\7$\2\2\u00b6\u00b7\5d\63\2\u00b7\u00b8\7%\2\2\u00b8\u00b9"+
		"\5j\66\2\u00b9\u00ca\3\2\2\2\u00ba\u00bb\7\5\2\2\u00bb\u00bc\7\37\2\2"+
		"\u00bc\u00bd\7%\2\2\u00bd\u00ca\5j\66\2\u00be\u00bf\7\5\2\2\u00bf\u00c0"+
		"\7\37\2\2\u00c0\u00c1\7$\2\2\u00c1\u00ca\5d\63\2\u00c2\u00c3\7\5\2\2\u00c3"+
		"\u00c4\7\37\2\2\u00c4\u00c5\7$\2\2\u00c5\u00c6\5d\63\2\u00c6\u00c7\7%"+
		"\2\2\u00c7\u00c8\5j\66\2\u00c8\u00ca\3\2\2\2\u00c9\u00ab\3\2\2\2\u00c9"+
		"\u00af\3\2\2\2\u00c9\u00b3\3\2\2\2\u00c9\u00ba\3\2\2\2\u00c9\u00be\3\2"+
		"\2\2\u00c9\u00c2\3\2\2\2\u00ca\13\3\2\2\2\u00cb\u00cc\7\37\2\2\u00cc\u00cd"+
		"\7%\2\2\u00cd\u00ce\5j\66\2\u00ce\r\3\2\2\2\u00cf\u00d0\7\6\2\2\u00d0"+
		"\u00d1\5j\66\2\u00d1\u00d2\7\65\2\2\u00d2\u00d3\5\4\3\2\u00d3\u00d4\7"+
		"\66\2\2\u00d4\u00e8\3\2\2\2\u00d5\u00d6\7\6\2\2\u00d6\u00d7\5j\66\2\u00d7"+
		"\u00d8\7\65\2\2\u00d8\u00d9\5\4\3\2\u00d9\u00da\7\66\2\2\u00da\u00db\7"+
		"\7\2\2\u00db\u00dc\7\65\2\2\u00dc\u00dd\5\4\3\2\u00dd\u00de\7\66\2\2\u00de"+
		"\u00e8\3\2\2\2\u00df\u00e0\7\6\2\2\u00e0\u00e1\5j\66\2\u00e1\u00e2\7\65"+
		"\2\2\u00e2\u00e3\5\4\3\2\u00e3\u00e4\7\66\2\2\u00e4\u00e5\7\7\2\2\u00e5"+
		"\u00e6\5\20\t\2\u00e6\u00e8\3\2\2\2\u00e7\u00cf\3\2\2\2\u00e7\u00d5\3"+
		"\2\2\2\u00e7\u00df\3\2\2\2\u00e8\17\3\2\2\2\u00e9\u00eb\5\16\b\2\u00ea"+
		"\u00e9\3\2\2\2\u00eb\u00ee\3\2\2\2\u00ec\u00ea\3\2\2\2\u00ec\u00ed\3\2"+
		"\2\2\u00ed\21\3\2\2\2\u00ee\u00ec\3\2\2\2\u00ef\u00f0\7\6\2\2\u00f0\u00f1"+
		"\5j\66\2\u00f1\u00f2\7\65\2\2\u00f2\u00f3\5j\66\2\u00f3\u00f4\7\66\2\2"+
		"\u00f4\u0108\3\2\2\2\u00f5\u00f6\7\6\2\2\u00f6\u00f7\5j\66\2\u00f7\u00f8"+
		"\7\65\2\2\u00f8\u00f9\5j\66\2\u00f9\u00fa\7\66\2\2\u00fa\u00fb\7\7\2\2"+
		"\u00fb\u00fc\7\65\2\2\u00fc\u00fd\5j\66\2\u00fd\u00fe\7\66\2\2\u00fe\u0108"+
		"\3\2\2\2\u00ff\u0100\7\6\2\2\u0100\u0101\5j\66\2\u0101\u0102\7\65\2\2"+
		"\u0102\u0103\5j\66\2\u0103\u0104\7\66\2\2\u0104\u0105\7\7\2\2\u0105\u0106"+
		"\5\24\13\2\u0106\u0108\3\2\2\2\u0107\u00ef\3\2\2\2\u0107\u00f5\3\2\2\2"+
		"\u0107\u00ff\3\2\2\2\u0108\23\3\2\2\2\u0109\u010b\5\22\n\2\u010a\u0109"+
		"\3\2\2\2\u010b\u010e\3\2\2\2\u010c\u010a\3\2\2\2\u010c\u010d\3\2\2\2\u010d"+
		"\25\3\2\2\2\u010e\u010c\3\2\2\2\u010f\u0110\7\n\2\2\u0110\u0111\5j\66"+
		"\2\u0111\u0112\7\65\2\2\u0112\u0113\5\30\r\2\u0113\u0114\5$\23\2\u0114"+
		"\u0115\7\66\2\2\u0115\u011d\3\2\2\2\u0116\u0117\7\n\2\2\u0117\u0118\5"+
		"j\66\2\u0118\u0119\7\65\2\2\u0119\u011a\5$\23\2\u011a\u011b\7\66\2\2\u011b"+
		"\u011d\3\2\2\2\u011c\u010f\3\2\2\2\u011c\u0116\3\2\2\2\u011d\27\3\2\2"+
		"\2\u011e\u0120\5\32\16\2\u011f\u011e\3\2\2\2\u0120\u0121\3\2\2\2\u0121"+
		"\u011f\3\2\2\2\u0121\u0122\3\2\2\2\u0122\31\3\2\2\2\u0123\u0124\5\34\17"+
		"\2\u0124\u0125\7$\2\2\u0125\u0126\7\65\2\2\u0126\u0127\5\4\3\2\u0127\u0128"+
		"\7\66\2\2\u0128\u012f\3\2\2\2\u0129\u012a\5\34\17\2\u012a\u012b\7$\2\2"+
		"\u012b\u012c\5 \21\2\u012c\u012d\7#\2\2\u012d\u012f\3\2\2\2\u012e\u0123"+
		"\3\2\2\2\u012e\u0129\3\2\2\2\u012f\33\3\2\2\2\u0130\u0132\5\36\20\2\u0131"+
		"\u0130\3\2\2\2\u0132\u0133\3\2\2\2\u0133\u0131\3\2\2\2\u0133\u0134\3\2"+
		"\2\2\u0134\35\3\2\2\2\u0135\u0136\5j\66\2\u0136\u0137\7<\2\2\u0137\u013a"+
		"\3\2\2\2\u0138\u013a\5j\66\2\u0139\u0135\3\2\2\2\u0139\u0138\3\2\2\2\u013a"+
		"\37\3\2\2\2\u013b\u013c\5\6\4\2\u013c!\3\2\2\2\u013d\u013e\7\13\2\2\u013e"+
		"\u013f\7\37\2\2\u013f\u0140\7$\2\2\u0140\u0141\7\65\2\2\u0141\u0142\5"+
		"\4\3\2\u0142\u0143\7\66\2\2\u0143\u0149\3\2\2\2\u0144\u0145\7\13\2\2\u0145"+
		"\u0146\7\37\2\2\u0146\u0147\7$\2\2\u0147\u0149\5 \21\2\u0148\u013d\3\2"+
		"\2\2\u0148\u0144\3\2\2\2\u0149#\3\2\2\2\u014a\u014c\5\"\22\2\u014b\u014a"+
		"\3\2\2\2\u014c\u014d\3\2\2\2\u014d\u014b\3\2\2\2\u014d\u014e\3\2\2\2\u014e"+
		"%\3\2\2\2\u014f\u0150\7\n\2\2\u0150\u0151\5j\66\2\u0151\u0152\7\65\2\2"+
		"\u0152\u0153\5(\25\2\u0153\u0154\5\60\31\2\u0154\u0155\7\66\2\2\u0155"+
		"\u015d\3\2\2\2\u0156\u0157\7\n\2\2\u0157\u0158\5j\66\2\u0158\u0159\7\65"+
		"\2\2\u0159\u015a\5\60\31\2\u015a\u015b\7\66\2\2\u015b\u015d\3\2\2\2\u015c"+
		"\u014f\3\2\2\2\u015c\u0156\3\2\2\2\u015d\'\3\2\2\2\u015e\u0160\5*\26\2"+
		"\u015f\u015e\3\2\2\2\u0160\u0161\3\2\2\2\u0161\u015f\3\2\2\2\u0161\u0162"+
		"\3\2\2\2\u0162)\3\2\2\2\u0163\u0164\5,\27\2\u0164\u0165\7$\2\2\u0165\u0166"+
		"\5j\66\2\u0166\u016e\3\2\2\2\u0167\u0168\5,\27\2\u0168\u0169\7$\2\2\u0169"+
		"\u016a\7\65\2\2\u016a\u016b\5j\66\2\u016b\u016c\7\66\2\2\u016c\u016e\3"+
		"\2\2\2\u016d\u0163\3\2\2\2\u016d\u0167\3\2\2\2\u016e+\3\2\2\2\u016f\u0171"+
		"\5.\30\2\u0170\u016f\3\2\2\2\u0171\u0172\3\2\2\2\u0172\u0170\3\2\2\2\u0172"+
		"\u0173\3\2\2\2\u0173-\3\2\2\2\u0174\u0175\5j\66\2\u0175\u0176\7<\2\2\u0176"+
		"\u0179\3\2\2\2\u0177\u0179\5j\66\2\u0178\u0174\3\2\2\2\u0178\u0177\3\2"+
		"\2\2\u0179/\3\2\2\2\u017a\u017b\7\13\2\2\u017b\u017c\7\37\2\2\u017c\u017d"+
		"\7$\2\2\u017d\u0186\5j\66\2\u017e\u017f\7\13\2\2\u017f\u0180\7\37\2\2"+
		"\u0180\u0181\7$\2\2\u0181\u0182\7\65\2\2\u0182\u0183\5j\66\2\u0183\u0184"+
		"\7\66\2\2\u0184\u0186\3\2\2\2\u0185\u017a\3\2\2\2\u0185\u017e\3\2\2\2"+
		"\u0186\61\3\2\2\2\u0187\u0188\7\t\2\2\u0188\u0189\5j\66\2\u0189\u018a"+
		"\7\65\2\2\u018a\u018b\5\4\3\2\u018b\u018c\7\66\2\2\u018c\63\3\2\2\2\u018d"+
		"\u018e\7\b\2\2\u018e\u018f\7\37\2\2\u018f\u0190\7\r\2\2\u0190\u0191\5"+
		"j\66\2\u0191\u0192\7 \2\2\u0192\u0193\5j\66\2\u0193\u0194\7\65\2\2\u0194"+
		"\u0195\5\4\3\2\u0195\u0196\7\66\2\2\u0196\u01a0\3\2\2\2\u0197\u0198\7"+
		"\b\2\2\u0198\u0199\7\37\2\2\u0199\u019a\7\r\2\2\u019a\u019b\5j\66\2\u019b"+
		"\u019c\7\65\2\2\u019c\u019d\5\4\3\2\u019d\u019e\7\66\2\2\u019e\u01a0\3"+
		"\2\2\2\u019f\u018d\3\2\2\2\u019f\u0197\3\2\2\2\u01a0\65\3\2\2\2\u01a1"+
		"\u01a2\7\t\2\2\u01a2\u01a3\7\31\2\2\u01a3\u01a4\7\65\2\2\u01a4\u01a5\5"+
		"\4\3\2\u01a5\u01a6\7\66\2\2\u01a6\67\3\2\2\2\u01a7\u01a8\7\t\2\2\u01a8"+
		"\u01a9\7\31\2\2\u01a9\u01aa\7\65\2\2\u01aa\u01ab\5\4\3\2\u01ab\u01ac\7"+
		"\66\2\2\u01ac9\3\2\2\2\u01ad\u01b1\7\16\2\2\u01ae\u01af\7\16\2\2\u01af"+
		"\u01b1\5j\66\2\u01b0\u01ad\3\2\2\2\u01b0\u01ae\3\2\2\2\u01b1;\3\2\2\2"+
		"\u01b2\u01b3\7\17\2\2\u01b3=\3\2\2\2\u01b4\u01b5\7\20\2\2\u01b5\u01b6"+
		"\5j\66\2\u01b6?\3\2\2\2\u01b7\u01b8\7\21\2\2\u01b8\u01b9\7\37\2\2\u01b9"+
		"\u01ba\7\63\2\2\u01ba\u01bb\7\64\2\2\u01bb\u01bc\7\65\2\2\u01bc\u01bd"+
		"\5\4\3\2\u01bd\u01be\7\66\2\2\u01be\u01de\3\2\2\2\u01bf\u01c0\7\21\2\2"+
		"\u01c0\u01c1\7\37\2\2\u01c1\u01c2\7\63\2\2\u01c2\u01c3\7\64\2\2\u01c3"+
		"\u01c4\7)\2\2\u01c4\u01c5\5d\63\2\u01c5\u01c6\7\65\2\2\u01c6\u01c7\5\4"+
		"\3\2\u01c7\u01c8\7\66\2\2\u01c8\u01de\3\2\2\2\u01c9\u01ca\7\21\2\2\u01ca"+
		"\u01cb\7\37\2\2\u01cb\u01cc\7\63\2\2\u01cc\u01cd\5B\"\2\u01cd\u01ce\7"+
		"\64\2\2\u01ce\u01cf\7\65\2\2\u01cf\u01d0\5\4\3\2\u01d0\u01d1\7\66\2\2"+
		"\u01d1\u01de\3\2\2\2\u01d2\u01d3\7\21\2\2\u01d3\u01d4\7\37\2\2\u01d4\u01d5"+
		"\7\63\2\2\u01d5\u01d6\5B\"\2\u01d6\u01d7\7\64\2\2\u01d7\u01d8\7)\2\2\u01d8"+
		"\u01d9\5d\63\2\u01d9\u01da\7\65\2\2\u01da\u01db\5\4\3\2\u01db\u01dc\7"+
		"\66\2\2\u01dc\u01de\3\2\2\2\u01dd\u01b7\3\2\2\2\u01dd\u01bf\3\2\2\2\u01dd"+
		"\u01c9\3\2\2\2\u01dd\u01d2\3\2\2\2\u01deA\3\2\2\2\u01df\u01e1\5D#\2\u01e0"+
		"\u01df\3\2\2\2\u01e1\u01e2\3\2\2\2\u01e2\u01e0\3\2\2\2\u01e2\u01e3\3\2"+
		"\2\2\u01e3C\3\2\2\2\u01e4\u01e5\7\37\2\2\u01e5\u01e6\7$\2\2\u01e6\u01e7"+
		"\5d\63\2\u01e7\u01e8\7#\2\2\u01e8\u01ed\3\2\2\2\u01e9\u01ea\7\37\2\2\u01ea"+
		"\u01eb\7$\2\2\u01eb\u01ed\5d\63\2\u01ec\u01e4\3\2\2\2\u01ec\u01e9\3\2"+
		"\2\2\u01edE\3\2\2\2\u01ee\u01ef\7\37\2\2\u01ef\u01f0\7\63\2\2\u01f0\u01f7"+
		"\7\64\2\2\u01f1\u01f2\7\37\2\2\u01f2\u01f3\7\63\2\2\u01f3\u01f4\5f\64"+
		"\2\u01f4\u01f5\7\64\2\2\u01f5\u01f7\3\2\2\2\u01f6\u01ee\3\2\2\2\u01f6"+
		"\u01f1\3\2\2\2\u01f7G\3\2\2\2\u01f8\u01f9\7\37\2\2\u01f9\u01fa\7\63\2"+
		"\2\u01fa\u0201\7\64\2\2\u01fb\u01fc\7\37\2\2\u01fc\u01fd\7\63\2\2\u01fd"+
		"\u01fe\5f\64\2\u01fe\u01ff\7\64\2\2\u01ff\u0201\3\2\2\2\u0200\u01f8\3"+
		"\2\2\2\u0200\u01fb\3\2\2\2\u0201I\3\2\2\2\u0202\u0203\7\22\2\2\u0203\u0204"+
		"\7\37\2\2\u0204\u0205\7\65\2\2\u0205\u0206\5L\'\2\u0206\u0207\7\66\2\2"+
		"\u0207K\3\2\2\2\u0208\u020a\5N(\2\u0209\u0208\3\2\2\2\u020a\u020b\3\2"+
		"\2\2\u020b\u0209\3\2\2\2\u020b\u020c\3\2\2\2\u020cM\3\2\2\2\u020d\u020e"+
		"\7\37\2\2\u020e\u020f\7$\2\2\u020f\u0210\5d\63\2\u0210\u0211\7#\2\2\u0211"+
		"\u0216\3\2\2\2\u0212\u0213\7\37\2\2\u0213\u0214\7$\2\2\u0214\u0216\5d"+
		"\63\2\u0215\u020d\3\2\2\2\u0215\u0212\3\2\2\2\u0216O\3\2\2\2\u0217\u0218"+
		"\7\37\2\2\u0218\u0219\5R*\2\u0219Q\3\2\2\2\u021a\u021c\5T+\2\u021b\u021a"+
		"\3\2\2\2\u021c\u021d\3\2\2\2\u021d\u021b\3\2\2\2\u021d\u021e\3\2\2\2\u021e"+
		"S\3\2\2\2\u021f\u0220\7\67\2\2\u0220\u0221\5j\66\2\u0221\u0222\78\2\2"+
		"\u0222U\3\2\2\2\u0223\u0224\7\4\2\2\u0224\u0225\7\37\2\2\u0225\u0226\7"+
		"%\2\2\u0226\u0227\7\37\2\2\u0227\u0228\7\65\2\2\u0228\u0229\5X-\2\u0229"+
		"\u022a\7\66\2\2\u022a\u0234\3\2\2\2\u022b\u022c\7\5\2\2\u022c\u022d\7"+
		"\37\2\2\u022d\u022e\7%\2\2\u022e\u022f\7\37\2\2\u022f\u0230\7\65\2\2\u0230"+
		"\u0231\5X-\2\u0231\u0232\7\66\2\2\u0232\u0234\3\2\2\2\u0233\u0223\3\2"+
		"\2\2\u0233\u022b\3\2\2\2\u0234W\3\2\2\2\u0235\u0237\5Z.\2\u0236\u0235"+
		"\3\2\2\2\u0237\u0238\3\2\2\2\u0238\u0236\3\2\2\2\u0238\u0239\3\2\2\2\u0239"+
		"Y\3\2\2\2\u023a\u023b\7\37\2\2\u023b\u023c\7$\2\2\u023c\u0243\5j\66\2"+
		"\u023d\u023e\7\37\2\2\u023e\u023f\7$\2\2\u023f\u0240\5j\66\2\u0240\u0241"+
		"\7#\2\2\u0241\u0243\3\2\2\2\u0242\u023a\3\2\2\2\u0242\u023d\3\2\2\2\u0243"+
		"[\3\2\2\2\u0244\u0245\7\37\2\2\u0245\u0246\5^\60\2\u0246]\3\2\2\2\u0247"+
		"\u0249\5`\61\2\u0248\u0247\3\2\2\2\u0249\u024a\3\2\2\2\u024a\u0248\3\2"+
		"\2\2\u024a\u024b\3\2\2\2\u024b_\3\2\2\2\u024c\u024d\7!\2\2\u024d\u024e"+
		"\7\37\2\2\u024ea\3\2\2\2\u024f\u0250\7\37\2\2\u0250\u0251\5^\60\2\u0251"+
		"\u0252\7%\2\2\u0252\u0253\5j\66\2\u0253c\3\2\2\2\u0254\u025c\7\23\2\2"+
		"\u0255\u025c\7\24\2\2\u0256\u025c\7\30\2\2\u0257\u025c\7\26\2\2\u0258"+
		"\u025c\7\27\2\2\u0259\u025a\7:\2\2\u025a\u025c\7\37\2\2\u025b\u0254\3"+
		"\2\2\2\u025b\u0255\3\2\2\2\u025b\u0256\3\2\2\2\u025b\u0257\3\2\2\2\u025b"+
		"\u0258\3\2\2\2\u025b\u0259\3\2\2\2\u025ce\3\2\2\2\u025d\u025f\5h\65\2"+
		"\u025e\u025d\3\2\2\2\u025f\u0260\3\2\2\2\u0260\u025e\3\2\2\2\u0260\u0261"+
		"\3\2\2\2\u0261\u0262\3\2\2\2\u0262\u0263\b\64\1\2\u0263g\3\2\2\2\u0264"+
		"\u0265\5j\66\2\u0265\u0266\7#\2\2\u0266\u0267\b\65\1\2\u0267\u026c\3\2"+
		"\2\2\u0268\u0269\5j\66\2\u0269\u026a\b\65\1\2\u026a\u026c\3\2\2\2\u026b"+
		"\u0264\3\2\2\2\u026b\u0268\3\2\2\2\u026ci\3\2\2\2\u026d\u026e\5l\67\2"+
		"\u026e\u026f\b\66\1\2\u026f\u0277\3\2\2\2\u0270\u0271\5p9\2\u0271\u0272"+
		"\b\66\1\2\u0272\u0277\3\2\2\2\u0273\u0274\5n8\2\u0274\u0275\b\66\1\2\u0275"+
		"\u0277\3\2\2\2\u0276\u026d\3\2\2\2\u0276\u0270\3\2\2\2\u0276\u0273\3\2"+
		"\2\2\u0277k\3\2\2\2\u0278\u0279\b\67\1\2\u0279\u027a\7=\2\2\u027a\u027b"+
		"\5l\67\5\u027b\u027c\b\67\1\2\u027c\u0281\3\2\2\2\u027d\u027e\5n8\2\u027e"+
		"\u027f\b\67\1\2\u027f\u0281\3\2\2\2\u0280\u0278\3\2\2\2\u0280\u027d\3"+
		"\2\2\2\u0281\u0289\3\2\2\2\u0282\u0283\f\4\2\2\u0283\u0284\t\2\2\2\u0284"+
		"\u0285\5l\67\5\u0285\u0286\b\67\1\2\u0286\u0288\3\2\2\2\u0287\u0282\3"+
		"\2\2\2\u0288\u028b\3\2\2\2\u0289\u0287\3\2\2\2\u0289\u028a\3\2\2\2\u028a"+
		"m\3\2\2\2\u028b\u0289\3\2\2\2\u028c\u028d\b8\1\2\u028d\u028e\5p9\2\u028e"+
		"\u028f\b8\1\2\u028f\u0297\3\2\2\2\u0290\u0291\f\4\2\2\u0291\u0292\t\3"+
		"\2\2\u0292\u0293\5n8\5\u0293\u0294\b8\1\2\u0294\u0296\3\2\2\2\u0295\u0290"+
		"\3\2\2\2\u0296\u0299\3\2\2\2\u0297\u0295\3\2\2\2\u0297\u0298\3\2\2\2\u0298"+
		"o\3\2\2\2\u0299\u0297\3\2\2\2\u029a\u029b\b9\1\2\u029b\u029c\7\62\2\2"+
		"\u029c\u029d\5p9\7\u029d\u029e\b9\1\2\u029e\u02a8\3\2\2\2\u029f\u02a0"+
		"\5r:\2\u02a0\u02a1\b9\1\2\u02a1\u02a8\3\2\2\2\u02a2\u02a3\7\63\2\2\u02a3"+
		"\u02a4\5j\66\2\u02a4\u02a5\7\64\2\2\u02a5\u02a6\b9\1\2\u02a6\u02a8\3\2"+
		"\2\2\u02a7\u029a\3\2\2\2\u02a7\u029f\3\2\2\2\u02a7\u02a2\3\2\2\2\u02a8"+
		"\u02b5\3\2\2\2\u02a9\u02aa\f\6\2\2\u02aa\u02ab\t\4\2\2\u02ab\u02ac\5p"+
		"9\7\u02ac\u02ad\b9\1\2\u02ad\u02b4\3\2\2\2\u02ae\u02af\f\5\2\2\u02af\u02b0"+
		"\t\5\2\2\u02b0\u02b1\5p9\6\u02b1\u02b2\b9\1\2\u02b2\u02b4\3\2\2\2\u02b3"+
		"\u02a9\3\2\2\2\u02b3\u02ae\3\2\2\2\u02b4\u02b7\3\2\2\2\u02b5\u02b3\3\2"+
		"\2\2\u02b5\u02b6\3\2\2\2\u02b6q\3\2\2\2\u02b7\u02b5\3\2\2\2\u02b8\u02b9"+
		"\5t;\2\u02b9\u02ba\b:\1\2\u02bas\3\2\2\2\u02bb\u02bc\7\32\2\2\u02bc\u02cc"+
		"\b;\1\2\u02bd\u02be\7\33\2\2\u02be\u02cc\b;\1\2\u02bf\u02c0\7\35\2\2\u02c0"+
		"\u02cc\b;\1\2\u02c1\u02c2\7\36\2\2\u02c2\u02cc\b;\1\2\u02c3\u02c4\7\34"+
		"\2\2\u02c4\u02cc\b;\1\2\u02c5\u02cc\5H%\2\u02c6\u02cc\5\\/\2\u02c7\u02cc"+
		"\5P)\2\u02c8\u02cc\7\37\2\2\u02c9\u02cc\5\22\n\2\u02ca\u02cc\5&\24\2\u02cb"+
		"\u02bb\3\2\2\2\u02cb\u02bd\3\2\2\2\u02cb\u02bf\3\2\2\2\u02cb\u02c1\3\2"+
		"\2\2\u02cb\u02c3\3\2\2\2\u02cb\u02c5\3\2\2\2\u02cb\u02c6\3\2\2\2\u02cb"+
		"\u02c7\3\2\2\2\u02cb\u02c8\3\2\2\2\u02cb\u02c9\3\2\2\2\u02cb\u02ca\3\2"+
		"\2\2\u02ccu\3\2\2\2\60}\u0093\u00a9\u00c9\u00e7\u00ec\u0107\u010c\u011c"+
		"\u0121\u012e\u0133\u0139\u0148\u014d\u015c\u0161\u016d\u0172\u0178\u0185"+
		"\u019f\u01b0\u01dd\u01e2\u01ec\u01f6\u0200\u020b\u0215\u021d\u0233\u0238"+
		"\u0242\u024a\u025b\u0260\u026b\u0276\u0280\u0289\u0297\u02a7\u02b3\u02b5"+
		"\u02cb";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}