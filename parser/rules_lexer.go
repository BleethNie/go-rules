// Code generated from E:/02_Resource/01_Code/jc_go_manager/go-rules/parser/Rules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type RulesLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var RulesLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ruleslexerLexerInit() {
	staticData := &RulesLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'['", "']'", "'&&'", "'||'", "'=='", "'!='", "'<'", "'<='",
		"'>'", "'>='", "'in'", "'+'", "'-'", "'*'", "'/'", "'%'", "'contains'",
		"'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "AND", "OR", "EQ", "NEQ", "LT", "LTE", "GT", "GTE",
		"IN", "ADD", "SUB", "MUL", "DIV", "MOD", "CONTAINS", "LPARENT", "RPARENT",
		"BooleanLiteral", "NumberLiteral", "DurationLiteral", "StringLiteral",
		"Identifier", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "AND", "OR", "EQ", "NEQ", "LT", "LTE", "GT",
		"GTE", "IN", "ADD", "SUB", "MUL", "DIV", "MOD", "CONTAINS", "LPARENT",
		"RPARENT", "BooleanLiteral", "NumberLiteral", "DurationLiteral", "StringLiteral",
		"Identifier", "BOOL", "DIGIT", "DIGIT_SEQUENCE", "SIGN", "INTEGER",
		"FLOAT", "TIME_UNIT", "DURATION", "CHAR_SEQUENCE", "ID_START", "ID_CHAR",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 216, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		20, 1, 20, 1, 21, 1, 21, 3, 21, 134, 8, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		5, 23, 140, 8, 23, 10, 23, 12, 23, 143, 9, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 5, 24, 149, 8, 24, 10, 24, 12, 24, 152, 9, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 163, 8, 25, 1, 26, 1,
		26, 1, 27, 4, 27, 168, 8, 27, 11, 27, 12, 27, 169, 1, 28, 1, 28, 1, 29,
		3, 29, 175, 8, 29, 1, 29, 1, 29, 1, 30, 4, 30, 180, 8, 30, 11, 30, 12,
		30, 181, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31,
		192, 8, 31, 1, 32, 4, 32, 195, 8, 32, 11, 32, 12, 32, 196, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 34, 3, 34, 204, 8, 34, 1, 35, 1, 35, 3, 35, 208, 8, 35,
		1, 36, 4, 36, 211, 8, 36, 11, 36, 12, 36, 212, 1, 36, 1, 36, 0, 0, 37,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 0, 53, 0, 55, 0, 57, 0, 59,
		0, 61, 0, 63, 0, 65, 0, 67, 0, 69, 0, 71, 0, 73, 26, 1, 0, 6, 1, 0, 48,
		57, 2, 0, 43, 43, 45, 45, 3, 0, 104, 104, 109, 109, 115, 115, 4, 0, 10,
		10, 13, 13, 34, 34, 92, 92, 3, 0, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10,
		13, 13, 32, 32, 216, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 73, 1, 0, 0,
		0, 1, 75, 1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 81, 1, 0,
		0, 0, 9, 84, 1, 0, 0, 0, 11, 87, 1, 0, 0, 0, 13, 90, 1, 0, 0, 0, 15, 93,
		1, 0, 0, 0, 17, 95, 1, 0, 0, 0, 19, 98, 1, 0, 0, 0, 21, 100, 1, 0, 0, 0,
		23, 103, 1, 0, 0, 0, 25, 106, 1, 0, 0, 0, 27, 108, 1, 0, 0, 0, 29, 110,
		1, 0, 0, 0, 31, 112, 1, 0, 0, 0, 33, 114, 1, 0, 0, 0, 35, 116, 1, 0, 0,
		0, 37, 125, 1, 0, 0, 0, 39, 127, 1, 0, 0, 0, 41, 129, 1, 0, 0, 0, 43, 133,
		1, 0, 0, 0, 45, 135, 1, 0, 0, 0, 47, 137, 1, 0, 0, 0, 49, 146, 1, 0, 0,
		0, 51, 162, 1, 0, 0, 0, 53, 164, 1, 0, 0, 0, 55, 167, 1, 0, 0, 0, 57, 171,
		1, 0, 0, 0, 59, 174, 1, 0, 0, 0, 61, 179, 1, 0, 0, 0, 63, 191, 1, 0, 0,
		0, 65, 194, 1, 0, 0, 0, 67, 200, 1, 0, 0, 0, 69, 203, 1, 0, 0, 0, 71, 207,
		1, 0, 0, 0, 73, 210, 1, 0, 0, 0, 75, 76, 5, 44, 0, 0, 76, 2, 1, 0, 0, 0,
		77, 78, 5, 91, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 93, 0, 0, 80, 6, 1,
		0, 0, 0, 81, 82, 5, 38, 0, 0, 82, 83, 5, 38, 0, 0, 83, 8, 1, 0, 0, 0, 84,
		85, 5, 124, 0, 0, 85, 86, 5, 124, 0, 0, 86, 10, 1, 0, 0, 0, 87, 88, 5,
		61, 0, 0, 88, 89, 5, 61, 0, 0, 89, 12, 1, 0, 0, 0, 90, 91, 5, 33, 0, 0,
		91, 92, 5, 61, 0, 0, 92, 14, 1, 0, 0, 0, 93, 94, 5, 60, 0, 0, 94, 16, 1,
		0, 0, 0, 95, 96, 5, 60, 0, 0, 96, 97, 5, 61, 0, 0, 97, 18, 1, 0, 0, 0,
		98, 99, 5, 62, 0, 0, 99, 20, 1, 0, 0, 0, 100, 101, 5, 62, 0, 0, 101, 102,
		5, 61, 0, 0, 102, 22, 1, 0, 0, 0, 103, 104, 5, 105, 0, 0, 104, 105, 5,
		110, 0, 0, 105, 24, 1, 0, 0, 0, 106, 107, 5, 43, 0, 0, 107, 26, 1, 0, 0,
		0, 108, 109, 5, 45, 0, 0, 109, 28, 1, 0, 0, 0, 110, 111, 5, 42, 0, 0, 111,
		30, 1, 0, 0, 0, 112, 113, 5, 47, 0, 0, 113, 32, 1, 0, 0, 0, 114, 115, 5,
		37, 0, 0, 115, 34, 1, 0, 0, 0, 116, 117, 5, 99, 0, 0, 117, 118, 5, 111,
		0, 0, 118, 119, 5, 110, 0, 0, 119, 120, 5, 116, 0, 0, 120, 121, 5, 97,
		0, 0, 121, 122, 5, 105, 0, 0, 122, 123, 5, 110, 0, 0, 123, 124, 5, 115,
		0, 0, 124, 36, 1, 0, 0, 0, 125, 126, 5, 40, 0, 0, 126, 38, 1, 0, 0, 0,
		127, 128, 5, 41, 0, 0, 128, 40, 1, 0, 0, 0, 129, 130, 3, 51, 25, 0, 130,
		42, 1, 0, 0, 0, 131, 134, 3, 59, 29, 0, 132, 134, 3, 61, 30, 0, 133, 131,
		1, 0, 0, 0, 133, 132, 1, 0, 0, 0, 134, 44, 1, 0, 0, 0, 135, 136, 3, 65,
		32, 0, 136, 46, 1, 0, 0, 0, 137, 141, 5, 34, 0, 0, 138, 140, 3, 67, 33,
		0, 139, 138, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145,
		5, 34, 0, 0, 145, 48, 1, 0, 0, 0, 146, 150, 3, 69, 34, 0, 147, 149, 3,
		71, 35, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0,
		0, 0, 150, 151, 1, 0, 0, 0, 151, 50, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0,
		153, 154, 5, 116, 0, 0, 154, 155, 5, 114, 0, 0, 155, 156, 5, 117, 0, 0,
		156, 163, 5, 101, 0, 0, 157, 158, 5, 102, 0, 0, 158, 159, 5, 97, 0, 0,
		159, 160, 5, 108, 0, 0, 160, 161, 5, 115, 0, 0, 161, 163, 5, 101, 0, 0,
		162, 153, 1, 0, 0, 0, 162, 157, 1, 0, 0, 0, 163, 52, 1, 0, 0, 0, 164, 165,
		7, 0, 0, 0, 165, 54, 1, 0, 0, 0, 166, 168, 3, 53, 26, 0, 167, 166, 1, 0,
		0, 0, 168, 169, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0,
		170, 56, 1, 0, 0, 0, 171, 172, 7, 1, 0, 0, 172, 58, 1, 0, 0, 0, 173, 175,
		3, 57, 28, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 1,
		0, 0, 0, 176, 177, 3, 55, 27, 0, 177, 60, 1, 0, 0, 0, 178, 180, 3, 59,
		29, 0, 179, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0,
		181, 182, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 5, 46, 0, 0, 184,
		185, 3, 55, 27, 0, 185, 62, 1, 0, 0, 0, 186, 187, 5, 110, 0, 0, 187, 192,
		5, 115, 0, 0, 188, 189, 5, 109, 0, 0, 189, 192, 5, 115, 0, 0, 190, 192,
		7, 2, 0, 0, 191, 186, 1, 0, 0, 0, 191, 188, 1, 0, 0, 0, 191, 190, 1, 0,
		0, 0, 192, 64, 1, 0, 0, 0, 193, 195, 3, 59, 29, 0, 194, 193, 1, 0, 0, 0,
		195, 196, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		198, 1, 0, 0, 0, 198, 199, 3, 63, 31, 0, 199, 66, 1, 0, 0, 0, 200, 201,
		8, 3, 0, 0, 201, 68, 1, 0, 0, 0, 202, 204, 7, 4, 0, 0, 203, 202, 1, 0,
		0, 0, 204, 70, 1, 0, 0, 0, 205, 208, 3, 69, 34, 0, 206, 208, 2, 48, 57,
		0, 207, 205, 1, 0, 0, 0, 207, 206, 1, 0, 0, 0, 208, 72, 1, 0, 0, 0, 209,
		211, 7, 5, 0, 0, 210, 209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 210,
		1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 215, 6, 36,
		0, 0, 215, 74, 1, 0, 0, 0, 13, 0, 133, 141, 150, 162, 169, 174, 181, 191,
		196, 203, 207, 212, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RulesLexerInit initializes any static state used to implement RulesLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRulesLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RulesLexerInit() {
	staticData := &RulesLexerLexerStaticData
	staticData.once.Do(ruleslexerLexerInit)
}

// NewRulesLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRulesLexer(input antlr.CharStream) *RulesLexer {
	RulesLexerInit()
	l := new(RulesLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &RulesLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
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
