grammar Latte;

program: topDef+;

topDef: classdef | fundef;

fundef: type_ ID '(' arg? ')' block # FunDef;

classdef:
	'class' ID '{' field* '}'					# BaseClassDef
	| 'class' ID 'extends' ID '{' field* '}'	# DerivedClassDef;

arg: type_ ID ( ',' type_ ID)*;

field: type_ ID ';' # ClassFieldDef | fundef # ClassMethodDef;

block: '{' stmt* '}';

lvalue:
	ID					# LVId
	| ID '[' expr ']'	# LVArrayRef
	| lvalue '.' lvalue	# LVField;

stmt:
	';'												# SEmpty
	| block											# SBlockStmt
	| type_ item ( ',' item)* ';'					# SDecl
	| lvalue '=' expr ';'							# SAss
	| lvalue '[' expr ']' '=' expr ';'				# SAssArray
	| lvalue '++' ';'								# SIncr
	| lvalue '--' ';'								# SDecr
	| 'return' expr ';'								# SRet
	| 'return' ';'									# SVRet
	| 'if' '(' expr ')' stmt						# SCond
	| 'if' '(' expr ')' stmt 'else' stmt			# SCondElse
	| 'while' '(' expr ')' stmt						# SWhile
	| 'for' '(' singular_type_ ID ':' ID ')' stmt	# SFor
	| expr ';'										# SExp;

type_:
	singular_type_ '[]'	# TArray
	| singular_type_	# TSingular;

singular_type_:
	ID			# TClass
	| 'int'		# TInt
	| 'string'	# TStr
	| 'boolean'	# TBool
	| 'void'	# TVoid;

item: ID | ID '=' expr;

expr: ('-' | '!') expr						# EUnOp
	| expr mulOp expr						# EMulOp
	| expr addOp expr						# EAddOp
	| expr relOp expr						# ERelOp
	| <assoc = right> expr '&&' expr		# EAnd
	| <assoc = right> expr '||' expr		# EOr
	| 'new' singular_type_ '[' expr ']'		# ENewArray
	| 'new' ID								# ENewStruct
	| expr '.' expr							# EFieldAccess
	| expr '[' expr ']'						# EArrayRef
	| 'self'								# ESelf
	| ID									# EId
	| INT									# EInt
	| 'true'								# ETrue
	| 'false'								# EFalse
	| expr '(' ( expr ( ',' expr)*)? ')'	# EFunCall
	| STR									# EStr
	| '(' ID ')' 'null'						# ENull
	| '(' expr ')'							# EParen;

addOp: '+' | '-';

mulOp: '*' | '/' | '%';

relOp: '<' | '<=' | '>' | '>=' | '==' | '!=';

COMMENT: ('#' ~[\r\n]* | '//' ~[\r\n]*) -> channel(HIDDEN);
MULTICOMMENT: '/*' .*? '*/' -> channel(HIDDEN);

fragment Letter: Capital | Small;
fragment Capital: [A-Z\u00C0-\u00D6\u00D8-\u00DE];
fragment Small: [a-z\u00DF-\u00F6\u00F8-\u00FF];
fragment Digit: [0-9];

INT: Digit+;
fragment ID_First: Letter | '_';
ID: ID_First (ID_First | Digit)*;

WS: (' ' | '\r' | '\t' | '\n')+ -> skip;

STR: '"' StringCharacters? '"';
fragment StringCharacters: StringCharacter+;
fragment StringCharacter: ~["\\] | '\\' [tnr"\\];
