// Code generated from Rules.g4 by ANTLR 4.8. DO NOT EDIT.

package _

import (
	"fmt"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []int32{
	4, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 28, 218,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 23, 3, 23, 5, 23, 136, 10, 23, 3, 24, 3, 24, 3, 25, 3, 25, 7, 25,
	142, 10, 25, 12, 25, 14, 25, 145, 11, 25, 3, 25, 3, 25, 3, 26, 3, 26, 7,
	26, 151, 10, 26, 12, 26, 14, 26, 154, 11, 26, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 165, 10, 27, 3, 28, 3, 28, 3,
	29, 6, 29, 170, 10, 29, 13, 29, 14, 29, 171, 3, 30, 3, 30, 3, 31, 5, 31,
	177, 10, 31, 3, 31, 3, 31, 3, 32, 6, 32, 182, 10, 32, 13, 32, 14, 32, 183,
	3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 194, 10,
	33, 3, 34, 6, 34, 197, 10, 34, 13, 34, 14, 34, 198, 3, 34, 3, 34, 3, 35,
	3, 35, 3, 36, 5, 36, 206, 10, 36, 3, 37, 3, 37, 5, 37, 210, 10, 37, 3,
	38, 6, 38, 213, 10, 38, 13, 38, 14, 38, 214, 3, 38, 3, 38, 2, 2, 39, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 2, 55, 2, 57, 2, 59, 2, 61,
	2, 63, 2, 65, 2, 67, 2, 69, 2, 71, 2, 73, 2, 75, 28, 3, 2, 8, 3, 2, 50,
	59, 4, 2, 45, 45, 47, 47, 5, 2, 106, 106, 111, 111, 117, 117, 6, 2, 12,
	12, 15, 15, 36, 36, 94, 94, 5, 2, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12,
	15, 15, 34, 34, 2, 218, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2,
	31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2,
	2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2,
	2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 75, 3, 2,
	2, 2, 3, 77, 3, 2, 2, 2, 5, 79, 3, 2, 2, 2, 7, 81, 3, 2, 2, 2, 9, 83, 3,
	2, 2, 2, 11, 86, 3, 2, 2, 2, 13, 89, 3, 2, 2, 2, 15, 92, 3, 2, 2, 2, 17,
	95, 3, 2, 2, 2, 19, 97, 3, 2, 2, 2, 21, 100, 3, 2, 2, 2, 23, 102, 3, 2,
	2, 2, 25, 105, 3, 2, 2, 2, 27, 108, 3, 2, 2, 2, 29, 110, 3, 2, 2, 2, 31,
	112, 3, 2, 2, 2, 33, 114, 3, 2, 2, 2, 35, 116, 3, 2, 2, 2, 37, 118, 3,
	2, 2, 2, 39, 127, 3, 2, 2, 2, 41, 129, 3, 2, 2, 2, 43, 131, 3, 2, 2, 2,
	45, 135, 3, 2, 2, 2, 47, 137, 3, 2, 2, 2, 49, 139, 3, 2, 2, 2, 51, 148,
	3, 2, 2, 2, 53, 164, 3, 2, 2, 2, 55, 166, 3, 2, 2, 2, 57, 169, 3, 2, 2,
	2, 59, 173, 3, 2, 2, 2, 61, 176, 3, 2, 2, 2, 63, 181, 3, 2, 2, 2, 65, 193,
	3, 2, 2, 2, 67, 196, 3, 2, 2, 2, 69, 202, 3, 2, 2, 2, 71, 205, 3, 2, 2,
	2, 73, 209, 3, 2, 2, 2, 75, 212, 3, 2, 2, 2, 77, 78, 7, 46, 2, 2, 78, 4,
	3, 2, 2, 2, 79, 80, 7, 93, 2, 2, 80, 6, 3, 2, 2, 2, 81, 82, 7, 95, 2, 2,
	82, 8, 3, 2, 2, 2, 83, 84, 7, 40, 2, 2, 84, 85, 7, 40, 2, 2, 85, 10, 3,
	2, 2, 2, 86, 87, 7, 126, 2, 2, 87, 88, 7, 126, 2, 2, 88, 12, 3, 2, 2, 2,
	89, 90, 7, 63, 2, 2, 90, 91, 7, 63, 2, 2, 91, 14, 3, 2, 2, 2, 92, 93, 7,
	35, 2, 2, 93, 94, 7, 63, 2, 2, 94, 16, 3, 2, 2, 2, 95, 96, 7, 62, 2, 2,
	96, 18, 3, 2, 2, 2, 97, 98, 7, 62, 2, 2, 98, 99, 7, 63, 2, 2, 99, 20, 3,
	2, 2, 2, 100, 101, 7, 64, 2, 2, 101, 22, 3, 2, 2, 2, 102, 103, 7, 64, 2,
	2, 103, 104, 7, 63, 2, 2, 104, 24, 3, 2, 2, 2, 105, 106, 7, 107, 2, 2,
	106, 107, 7, 112, 2, 2, 107, 26, 3, 2, 2, 2, 108, 109, 7, 45, 2, 2, 109,
	28, 3, 2, 2, 2, 110, 111, 7, 47, 2, 2, 111, 30, 3, 2, 2, 2, 112, 113, 7,
	44, 2, 2, 113, 32, 3, 2, 2, 2, 114, 115, 7, 49, 2, 2, 115, 34, 3, 2, 2,
	2, 116, 117, 7, 39, 2, 2, 117, 36, 3, 2, 2, 2, 118, 119, 7, 101, 2, 2,
	119, 120, 7, 113, 2, 2, 120, 121, 7, 112, 2, 2, 121, 122, 7, 118, 2, 2,
	122, 123, 7, 99, 2, 2, 123, 124, 7, 107, 2, 2, 124, 125, 7, 112, 2, 2,
	125, 126, 7, 117, 2, 2, 126, 38, 3, 2, 2, 2, 127, 128, 7, 42, 2, 2, 128,
	40, 3, 2, 2, 2, 129, 130, 7, 43, 2, 2, 130, 42, 3, 2, 2, 2, 131, 132, 5,
	53, 27, 2, 132, 44, 3, 2, 2, 2, 133, 136, 5, 61, 31, 2, 134, 136, 5, 63,
	32, 2, 135, 133, 3, 2, 2, 2, 135, 134, 3, 2, 2, 2, 136, 46, 3, 2, 2, 2,
	137, 138, 5, 67, 34, 2, 138, 48, 3, 2, 2, 2, 139, 143, 7, 36, 2, 2, 140,
	142, 5, 69, 35, 2, 141, 140, 3, 2, 2, 2, 142, 145, 3, 2, 2, 2, 143, 141,
	3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 146, 3, 2, 2, 2, 145, 143, 3, 2,
	2, 2, 146, 147, 7, 36, 2, 2, 147, 50, 3, 2, 2, 2, 148, 152, 5, 71, 36,
	2, 149, 151, 5, 73, 37, 2, 150, 149, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2,
	152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 52, 3, 2, 2, 2, 154, 152,
	3, 2, 2, 2, 155, 156, 7, 118, 2, 2, 156, 157, 7, 116, 2, 2, 157, 158, 7,
	119, 2, 2, 158, 165, 7, 103, 2, 2, 159, 160, 7, 104, 2, 2, 160, 161, 7,
	99, 2, 2, 161, 162, 7, 110, 2, 2, 162, 163, 7, 117, 2, 2, 163, 165, 7,
	103, 2, 2, 164, 155, 3, 2, 2, 2, 164, 159, 3, 2, 2, 2, 165, 54, 3, 2, 2,
	2, 166, 167, 9, 2, 2, 2, 167, 56, 3, 2, 2, 2, 168, 170, 5, 55, 28, 2, 169,
	168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172,
	3, 2, 2, 2, 172, 58, 3, 2, 2, 2, 173, 174, 9, 3, 2, 2, 174, 60, 3, 2, 2,
	2, 175, 177, 5, 59, 30, 2, 176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2,
	177, 178, 3, 2, 2, 2, 178, 179, 5, 57, 29, 2, 179, 62, 3, 2, 2, 2, 180,
	182, 5, 61, 31, 2, 181, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 181,
	3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 7, 48,
	2, 2, 186, 187, 5, 57, 29, 2, 187, 64, 3, 2, 2, 2, 188, 189, 7, 112, 2,
	2, 189, 194, 7, 117, 2, 2, 190, 191, 7, 111, 2, 2, 191, 194, 7, 117, 2,
	2, 192, 194, 9, 4, 2, 2, 193, 188, 3, 2, 2, 2, 193, 190, 3, 2, 2, 2, 193,
	192, 3, 2, 2, 2, 194, 66, 3, 2, 2, 2, 195, 197, 5, 61, 31, 2, 196, 195,
	3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2,
	2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 5, 65, 33, 2, 201, 68, 3, 2, 2, 2,
	202, 203, 10, 5, 2, 2, 203, 70, 3, 2, 2, 2, 204, 206, 9, 6, 2, 2, 205,
	204, 3, 2, 2, 2, 206, 72, 3, 2, 2, 2, 207, 210, 5, 71, 36, 2, 208, 210,
	4, 50, 59, 2, 209, 207, 3, 2, 2, 2, 209, 208, 3, 2, 2, 2, 210, 74, 3, 2,
	2, 2, 211, 213, 9, 7, 2, 2, 212, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2,
	214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216,
	217, 8, 38, 2, 2, 217, 76, 3, 2, 2, 2, 15, 2, 135, 143, 152, 164, 171,
	176, 183, 193, 198, 205, 209, 214, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.Deserialize(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'['", "']'", "'&&'", "'||'", "'=='", "'!='", "'<'", "'<='",
	"'>'", "'>='", "'in'", "'+'", "'-'", "'*'", "'/'", "'%'", "'contains'",
	"'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "AND", "OR", "EQ", "NEQ", "LT", "LTE", "GT", "GTE", "IN",
	"ADD", "SUB", "MUL", "DIV", "MOD", "CONTAINS", "LPARENT", "RPARENT", "BooleanLiteral",
	"NumberLiteral", "DurationLiteral", "StringLiteral", "Identifier", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "AND", "OR", "EQ", "NEQ", "LT", "LTE", "GT", "GTE",
	"IN", "ADD", "SUB", "MUL", "DIV", "MOD", "CONTAINS", "LPARENT", "RPARENT",
	"BooleanLiteral", "NumberLiteral", "DurationLiteral", "StringLiteral",
	"Identifier", "BOOL", "DIGIT", "DIGIT_SEQUENCE", "SIGN", "INTEGER", "FLOAT",
	"TIME_UNIT", "DURATION", "CHAR_SEQUENCE", "ID_START", "ID_CHAR", "WS",
}

type RulesLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewRulesLexer(input antlr.CharStream) *RulesLexer {

	l := new(RulesLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Rules.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RulesLexer tokens.
const (
	RulesLexerT__0            = 1
	RulesLexerT__1            = 2
	RulesLexerT__2            = 3
	RulesLexerAND             = 4
	RulesLexerOR              = 5
	RulesLexerEQ              = 6
	RulesLexerNEQ             = 7
	RulesLexerLT              = 8
	RulesLexerLTE             = 9
	RulesLexerGT              = 10
	RulesLexerGTE             = 11
	RulesLexerIN              = 12
	RulesLexerADD             = 13
	RulesLexerSUB             = 14
	RulesLexerMUL             = 15
	RulesLexerDIV             = 16
	RulesLexerMOD             = 17
	RulesLexerCONTAINS        = 18
	RulesLexerLPARENT         = 19
	RulesLexerRPARENT         = 20
	RulesLexerBooleanLiteral  = 21
	RulesLexerNumberLiteral   = 22
	RulesLexerDurationLiteral = 23
	RulesLexerStringLiteral   = 24
	RulesLexerIdentifier      = 25
	RulesLexerWS              = 26
)
