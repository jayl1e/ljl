grammar Ljl;

file: top* EOF;
top :   funcDefine
    |   funcExtern ';'
    |   typeDefine
    ;

typeDefine: 'type' ID  'struct' '{' (structMem )* '}' # structDefine
          | 'type' ID '=' typ # typeAlias
          ;

structMem: ID ':' typ;

funcExtern: 'extern' funcProto;
funcDefine : funcProto block;

funcProto:  'fn' ID '(' paramList ')' '->' typ;

paramList:
         | param (',' param)*
         ;

param: ID ':' typ;
typ: ID
    | '(' ID ',' ID (',' ID)* ')'
    | '(' ')'
    ;

block: '{' expr* '}';

expr: block  # blockExpr
    | expr op=('*'|'/') expr # mulDivExpr
    | expr op=('+'|'-') expr # addSubExpr
    | expr op=('>'|'<'|'>='|'<=') expr # cmpExpr
    | expr op=('=='| '!=') expr # eqExpr
    | '!' expr #notExpr
    | expr '&&' expr # andExpr
    | expr '||' expr # orExpr
    | ID tupleVal # callExpr
    | ID # valRefExpr
    | literal # litExpr
    | 'if' expr 'then' expr 'else' expr # ifExpr
    | tupleVal #tupleExpr
    | expr '.' ID #memberExpr
    ;

tupleVal: '(' ( expr (','expr)*)? ')';

literal: INT
        | FLOAT
        | STRING
        ;

INT: '-'? [0-9]+;
FLOAT: '-'? [0-9]+ '.' [0-9]+;
STRING: '"' .*? '"';
ID: [a-zA-Z_][a-zA-Z0-9_]*;
COMMENT: '/*' .*? '*/' -> skip;
WS: [ \t\r\n]+ ->skip;
