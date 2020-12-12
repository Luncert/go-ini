grammar Ini;

main: (section (MultiNewLine section)*)? EOF;

section: sectionHeader sectionBody;

sectionHeader: SectionHeader;

sectionBody: (valueLine | commentLine)*;

valueLine: MultiNewLine variableName '=' value?;

variableName: Identifier;

value: basicValue | listValue;

listValue: basicValue (',' basicValue)+;

basicValue: boolValue | stringValue | numberValue;

commentLine: MultiNewLine CommentLine;

SectionHeader: SectionHeaderStart SectionName SectionHeaderEnd;

SectionHeaderStart: '[';

SectionHeaderEnd: ']';

fragment SectionName: ~[[\]]+;

boolValue: BooleanLiteral;

BooleanLiteral: ([Yy] 'es') | ([Nn] 'o');

// string value

stringValue: StringLiteral;

StringLiteral: '"' ~["]* '"';

// number value

numberValue: integerValue | decimalValue;

integerValue: IntegerLiteral;

decimalValue: DecimalLiteral;

IntegerLiteral: '0' | '0x' HexDigit+ | '-'? NonZeroDigit Digits?;

DecimalLiteral: '-'? Digits '.' Digits? | '.' Digits;

fragment HexDigit: [1-9a-fA-F];

fragment Digits: Digit+;

fragment Digit: '0' | NonZeroDigit;

fragment NonZeroDigit: [1-9];

// identifier

Identifier: Letter LetterOrDigitOrUnderscore*;

fragment Letter: [a-zA-Z];

fragment LetterOrDigitOrUnderscore: [a-zA-Z0-9_];

CommentLine: NewLine ';' CommentLiteral;

fragment CommentLiteral: ~[\n]*;

MultiNewLine: NewLine+;

NewLine: [\r]? [\n];

// whitespace

WS: [ \t\r]+ -> skip;