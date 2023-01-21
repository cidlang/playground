# A simple one-pass compiler

This code is based in the chapter 2 of the book "Compilers: principles, techniques and tools" by Aho, Sethi, and Ullman. At the end of this chapter we build a simple one-pass compiler.

## Infix to postfix translator specification

<pre>
  init → list <b>eof</b>
  list → expr ; list
       | ϵ
  expr → expr + term       { print('+') }
       | expr - term       { print('-') }
       | term
  term → term * factor     { print('*') }
       | term / factor     { print('/') }
       | term <b>div</b> factor   { print('DIV') }
       | term <b>mod</b> factor   { print('MOD') }
       | factor
factor → ( expr )
       | <b>id</b>                { print(id.lexeme) }
       | num               { print(num.value) }
</pre>

## Description of tokens

| Lexeme                                                    | Token          | Attribute value               |
|-----------------------------------------------------------|----------------|-------------------------------|
| Blank space                                               |                |                               |
| Digit sequence                                            | NUM            | Numeric value of the sequence |
| div                                                       | DIV            |                               |
| mod                                                       | MOD            |                               |
| Other sequence of a letter followed by letters and digits | ID             | Index of `symboltable`        |
| End-of-file character                                     | END            |                               |
| Any character                                             | This character | NONE                          |

# Syntax-directed translation scheme after removal of left recursion

<pre>
       init → list <b>eof</b>
       list → expr ; list
            | ϵ
       expr → term moreterms
  moreterms → + term { print('+') } moreterms
            | - term { print('-') } moreterms
            | ϵ
       term → factor morefactors
morefactors → * factor { print('*') } morefactors
            | / factor { print('/') } morefactors
            | <b>div</b> factor { print('DIV') } morefactors
            | <b>mod</b> factor { print('MOD') } morefactors
            | ϵ
     factor → ( expr )
            | <b>id</b> { print(id.lexeme) }
            | <b>num</b> { print(num.value) }
 </pre>