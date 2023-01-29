package lexical

type keywords []Symbol

var defaultKeywords = keywords{
	symbol{lexeme: "div", component: DIV},
	symbol{lexeme: "mod", component: MOD},
}
