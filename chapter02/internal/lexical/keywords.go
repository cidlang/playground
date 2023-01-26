package lexical

type keywords []Symbol

var defaultKeywords = keywords{
	{lexeme: "div", component: DIV},
	{lexeme: "mod", component: MOD},
}
