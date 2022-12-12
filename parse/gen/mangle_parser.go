// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package gen // Mangle
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MangleParser struct {
	*antlr.BaseParser
}

var mangleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mangleParserInit() {
	staticData := &mangleParserStaticData
	staticData.literalNames = []string{
		"", "'.'", "'descr'", "'inclusion'", "':'", "'{'", "'}'", "", "", "'Package'",
		"'Use'", "'Decl'", "'bound'", "'let'", "'do'", "'('", "')'", "'['",
		"']'", "'='", "'!='", "','", "'!'", "'<'", "'<='", "':-'", "'\\n'",
		"'|>'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "WHITESPACE", "COMMENT", "PACKAGE", "USE",
		"DECL", "BOUND", "LET", "DO", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET",
		"EQ", "BANGEQ", "COMMA", "BANG", "LESS", "LESSEQ", "COLONDASH", "NEWLINE",
		"PIPEGREATER", "NUMBER", "VARIABLE", "NAME", "CONSTANT", "STRING",
	}
	staticData.ruleNames = []string{
		"start", "program", "packageDecl", "useDecl", "decl", "descrBlock",
		"boundsBlock", "constraintsBlock", "clause", "clauseBody", "transform",
		"letStmt", "literalOrFml", "term", "atom", "atoms",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 236, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 1, 3, 1, 37, 8, 1, 1, 1, 5, 1, 40, 8, 1, 10, 1, 12,
		1, 43, 9, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9, 1, 1, 2,
		1, 2, 1, 2, 3, 2, 55, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 62, 8,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 69, 8, 4, 1, 4, 5, 4, 72, 8, 4,
		10, 4, 12, 4, 75, 9, 4, 1, 4, 3, 4, 78, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 90, 8, 6, 10, 6, 12, 6, 93, 9, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 103, 8, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 110, 8, 9, 10, 9, 12, 9, 113, 9, 9, 1, 9,
		1, 9, 3, 9, 117, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		125, 8, 10, 10, 10, 12, 10, 128, 9, 10, 3, 10, 130, 8, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 135, 8, 10, 10, 10, 12, 10, 138, 9, 10, 3, 10, 140, 8, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 150, 8,
		12, 1, 12, 1, 12, 3, 12, 154, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 165, 8, 13, 10, 13, 12, 13, 168, 9,
		13, 3, 13, 170, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 177, 8,
		13, 10, 13, 12, 13, 180, 9, 13, 3, 13, 182, 8, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 194, 8, 13, 10,
		13, 12, 13, 197, 9, 13, 3, 13, 199, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 211, 8, 13, 10, 13, 12,
		13, 214, 9, 13, 3, 13, 216, 8, 13, 1, 13, 3, 13, 219, 8, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 227, 8, 15, 10, 15, 12, 15, 230,
		9, 15, 3, 15, 232, 8, 15, 1, 15, 1, 15, 1, 15, 0, 0, 16, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 1, 2, 0, 19, 20, 23, 24,
		255, 0, 32, 1, 0, 0, 0, 2, 36, 1, 0, 0, 0, 4, 51, 1, 0, 0, 0, 6, 58, 1,
		0, 0, 0, 8, 65, 1, 0, 0, 0, 10, 81, 1, 0, 0, 0, 12, 84, 1, 0, 0, 0, 14,
		96, 1, 0, 0, 0, 16, 99, 1, 0, 0, 0, 18, 106, 1, 0, 0, 0, 20, 139, 1, 0,
		0, 0, 22, 141, 1, 0, 0, 0, 24, 153, 1, 0, 0, 0, 26, 218, 1, 0, 0, 0, 28,
		220, 1, 0, 0, 0, 30, 222, 1, 0, 0, 0, 32, 33, 3, 2, 1, 0, 33, 34, 5, 0,
		0, 1, 34, 1, 1, 0, 0, 0, 35, 37, 3, 4, 2, 0, 36, 35, 1, 0, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 41, 1, 0, 0, 0, 38, 40, 3, 6, 3, 0, 39, 38, 1, 0, 0, 0,
		40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 48, 1,
		0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 47, 3, 8, 4, 0, 45, 47, 3, 16, 8, 0, 46,
		44, 1, 0, 0, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0,
		0, 48, 49, 1, 0, 0, 0, 49, 3, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 52, 5,
		9, 0, 0, 52, 54, 5, 30, 0, 0, 53, 55, 3, 30, 15, 0, 54, 53, 1, 0, 0, 0,
		54, 55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 5, 22, 0, 0, 57, 5, 1,
		0, 0, 0, 58, 59, 5, 10, 0, 0, 59, 61, 5, 30, 0, 0, 60, 62, 3, 30, 15, 0,
		61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 5,
		22, 0, 0, 64, 7, 1, 0, 0, 0, 65, 66, 5, 11, 0, 0, 66, 68, 3, 28, 14, 0,
		67, 69, 3, 10, 5, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 73, 1,
		0, 0, 0, 70, 72, 3, 12, 6, 0, 71, 70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73,
		71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0,
		0, 76, 78, 3, 14, 7, 0, 77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79,
		1, 0, 0, 0, 79, 80, 5, 1, 0, 0, 80, 9, 1, 0, 0, 0, 81, 82, 5, 2, 0, 0,
		82, 83, 3, 30, 15, 0, 83, 11, 1, 0, 0, 0, 84, 85, 5, 12, 0, 0, 85, 86,
		5, 17, 0, 0, 86, 91, 3, 26, 13, 0, 87, 88, 5, 21, 0, 0, 88, 90, 3, 26,
		13, 0, 89, 87, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91,
		92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 95, 5, 18,
		0, 0, 95, 13, 1, 0, 0, 0, 96, 97, 5, 3, 0, 0, 97, 98, 3, 30, 15, 0, 98,
		15, 1, 0, 0, 0, 99, 102, 3, 28, 14, 0, 100, 101, 5, 25, 0, 0, 101, 103,
		3, 18, 9, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 1, 0,
		0, 0, 104, 105, 5, 1, 0, 0, 105, 17, 1, 0, 0, 0, 106, 111, 3, 24, 12, 0,
		107, 108, 5, 21, 0, 0, 108, 110, 3, 24, 12, 0, 109, 107, 1, 0, 0, 0, 110,
		113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 116,
		1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 115, 5, 27, 0, 0, 115, 117, 3, 20,
		10, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 19, 1, 0, 0, 0,
		118, 119, 5, 14, 0, 0, 119, 129, 3, 26, 13, 0, 120, 121, 5, 21, 0, 0, 121,
		126, 3, 22, 11, 0, 122, 123, 5, 21, 0, 0, 123, 125, 3, 22, 11, 0, 124,
		122, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127,
		1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 120, 1, 0,
		0, 0, 129, 130, 1, 0, 0, 0, 130, 140, 1, 0, 0, 0, 131, 136, 3, 22, 11,
		0, 132, 133, 5, 21, 0, 0, 133, 135, 3, 22, 11, 0, 134, 132, 1, 0, 0, 0,
		135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 118, 1, 0, 0, 0, 139, 131,
		1, 0, 0, 0, 140, 21, 1, 0, 0, 0, 141, 142, 5, 13, 0, 0, 142, 143, 5, 29,
		0, 0, 143, 144, 5, 19, 0, 0, 144, 145, 3, 26, 13, 0, 145, 23, 1, 0, 0,
		0, 146, 149, 3, 26, 13, 0, 147, 148, 7, 0, 0, 0, 148, 150, 3, 26, 13, 0,
		149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 154, 1, 0, 0, 0, 151,
		152, 5, 22, 0, 0, 152, 154, 3, 26, 13, 0, 153, 146, 1, 0, 0, 0, 153, 151,
		1, 0, 0, 0, 154, 25, 1, 0, 0, 0, 155, 219, 5, 29, 0, 0, 156, 219, 5, 31,
		0, 0, 157, 219, 5, 28, 0, 0, 158, 219, 5, 32, 0, 0, 159, 160, 5, 30, 0,
		0, 160, 169, 5, 15, 0, 0, 161, 166, 3, 26, 13, 0, 162, 163, 5, 21, 0, 0,
		163, 165, 3, 26, 13, 0, 164, 162, 1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166,
		164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168, 166,
		1, 0, 0, 0, 169, 161, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 1, 0,
		0, 0, 171, 219, 5, 16, 0, 0, 172, 181, 5, 17, 0, 0, 173, 178, 3, 26, 13,
		0, 174, 175, 5, 21, 0, 0, 175, 177, 3, 26, 13, 0, 176, 174, 1, 0, 0, 0,
		177, 180, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 181, 173, 1, 0, 0, 0, 181, 182,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 219, 5, 18, 0, 0, 184, 198, 5, 17,
		0, 0, 185, 186, 3, 26, 13, 0, 186, 187, 5, 4, 0, 0, 187, 195, 3, 26, 13,
		0, 188, 189, 5, 21, 0, 0, 189, 190, 3, 26, 13, 0, 190, 191, 5, 4, 0, 0,
		191, 192, 3, 26, 13, 0, 192, 194, 1, 0, 0, 0, 193, 188, 1, 0, 0, 0, 194,
		197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 199,
		1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 185, 1, 0, 0, 0, 198, 199, 1, 0,
		0, 0, 199, 200, 1, 0, 0, 0, 200, 219, 5, 18, 0, 0, 201, 215, 5, 5, 0, 0,
		202, 203, 3, 26, 13, 0, 203, 204, 5, 4, 0, 0, 204, 212, 3, 26, 13, 0, 205,
		206, 5, 21, 0, 0, 206, 207, 3, 26, 13, 0, 207, 208, 5, 4, 0, 0, 208, 209,
		3, 26, 13, 0, 209, 211, 1, 0, 0, 0, 210, 205, 1, 0, 0, 0, 211, 214, 1,
		0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 216, 1, 0, 0,
		0, 214, 212, 1, 0, 0, 0, 215, 202, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216,
		217, 1, 0, 0, 0, 217, 219, 5, 6, 0, 0, 218, 155, 1, 0, 0, 0, 218, 156,
		1, 0, 0, 0, 218, 157, 1, 0, 0, 0, 218, 158, 1, 0, 0, 0, 218, 159, 1, 0,
		0, 0, 218, 172, 1, 0, 0, 0, 218, 184, 1, 0, 0, 0, 218, 201, 1, 0, 0, 0,
		219, 27, 1, 0, 0, 0, 220, 221, 3, 26, 13, 0, 221, 29, 1, 0, 0, 0, 222,
		231, 5, 17, 0, 0, 223, 228, 3, 28, 14, 0, 224, 225, 5, 21, 0, 0, 225, 227,
		3, 28, 14, 0, 226, 224, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1,
		0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0,
		0, 231, 223, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233,
		234, 5, 18, 0, 0, 234, 31, 1, 0, 0, 0, 30, 36, 41, 46, 48, 54, 61, 68,
		73, 77, 91, 102, 111, 116, 126, 129, 136, 139, 149, 153, 166, 169, 178,
		181, 195, 198, 212, 215, 218, 228, 231,
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

// MangleParserInit initializes any static state used to implement MangleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMangleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MangleParserInit() {
	staticData := &mangleParserStaticData
	staticData.once.Do(mangleParserInit)
}

// NewMangleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMangleParser(input antlr.TokenStream) *MangleParser {
	MangleParserInit()
	this := new(MangleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mangleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// MangleParser tokens.
const (
	MangleParserEOF         = antlr.TokenEOF
	MangleParserT__0        = 1
	MangleParserT__1        = 2
	MangleParserT__2        = 3
	MangleParserT__3        = 4
	MangleParserT__4        = 5
	MangleParserT__5        = 6
	MangleParserWHITESPACE  = 7
	MangleParserCOMMENT     = 8
	MangleParserPACKAGE     = 9
	MangleParserUSE         = 10
	MangleParserDECL        = 11
	MangleParserBOUND       = 12
	MangleParserLET         = 13
	MangleParserDO          = 14
	MangleParserLPAREN      = 15
	MangleParserRPAREN      = 16
	MangleParserLBRACKET    = 17
	MangleParserRBRACKET    = 18
	MangleParserEQ          = 19
	MangleParserBANGEQ      = 20
	MangleParserCOMMA       = 21
	MangleParserBANG        = 22
	MangleParserLESS        = 23
	MangleParserLESSEQ      = 24
	MangleParserCOLONDASH   = 25
	MangleParserNEWLINE     = 26
	MangleParserPIPEGREATER = 27
	MangleParserNUMBER      = 28
	MangleParserVARIABLE    = 29
	MangleParserNAME        = 30
	MangleParserCONSTANT    = 31
	MangleParserSTRING      = 32
)

// MangleParser rules.
const (
	MangleParserRULE_start            = 0
	MangleParserRULE_program          = 1
	MangleParserRULE_packageDecl      = 2
	MangleParserRULE_useDecl          = 3
	MangleParserRULE_decl             = 4
	MangleParserRULE_descrBlock       = 5
	MangleParserRULE_boundsBlock      = 6
	MangleParserRULE_constraintsBlock = 7
	MangleParserRULE_clause           = 8
	MangleParserRULE_clauseBody       = 9
	MangleParserRULE_transform        = 10
	MangleParserRULE_letStmt          = 11
	MangleParserRULE_literalOrFml     = 12
	MangleParserRULE_term             = 13
	MangleParserRULE_atom             = 14
	MangleParserRULE_atoms            = 15
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Program() IProgramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(MangleParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MangleParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Program()
	}
	{
		p.SetState(33)
		p.Match(MangleParserEOF)
	}

	return localctx
}

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PackageDecl() IPackageDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageDeclContext)
}

func (s *ProgramContext) AllUseDecl() []IUseDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUseDeclContext); ok {
			len++
		}
	}

	tst := make([]IUseDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUseDeclContext); ok {
			tst[i] = t.(IUseDeclContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) UseDecl(i int) IUseDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseDeclContext)
}

func (s *ProgramContext) AllDecl() []IDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclContext); ok {
			len++
		}
	}

	tst := make([]IDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclContext); ok {
			tst[i] = t.(IDeclContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Decl(i int) IDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ProgramContext) AllClause() []IClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClauseContext); ok {
			len++
		}
	}

	tst := make([]IClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClauseContext); ok {
			tst[i] = t.(IClauseContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Clause(i int) IClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MangleParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MangleParserPACKAGE {
		{
			p.SetState(35)
			p.PackageDecl()
		}

	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MangleParserUSE {
		{
			p.SetState(38)
			p.UseDecl()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8321632288) != 0 {
		p.SetState(46)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MangleParserDECL:
			{
				p.SetState(44)
				p.Decl()
			}

		case MangleParserT__4, MangleParserLBRACKET, MangleParserNUMBER, MangleParserVARIABLE, MangleParserNAME, MangleParserCONSTANT, MangleParserSTRING:
			{
				p.SetState(45)
				p.Clause()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPackageDeclContext is an interface to support dynamic dispatch.
type IPackageDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageDeclContext differentiates from other interfaces.
	IsPackageDeclContext()
}

type PackageDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageDeclContext() *PackageDeclContext {
	var p = new(PackageDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_packageDecl
	return p
}

func (*PackageDeclContext) IsPackageDeclContext() {}

func NewPackageDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageDeclContext {
	var p = new(PackageDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_packageDecl

	return p
}

func (s *PackageDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageDeclContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(MangleParserPACKAGE, 0)
}

func (s *PackageDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(MangleParserNAME, 0)
}

func (s *PackageDeclContext) BANG() antlr.TerminalNode {
	return s.GetToken(MangleParserBANG, 0)
}

func (s *PackageDeclContext) Atoms() IAtomsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomsContext)
}

func (s *PackageDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterPackageDecl(s)
	}
}

func (s *PackageDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitPackageDecl(s)
	}
}

func (s *PackageDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitPackageDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) PackageDecl() (localctx IPackageDeclContext) {
	this := p
	_ = this

	localctx = NewPackageDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MangleParserRULE_packageDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(MangleParserPACKAGE)
	}
	{
		p.SetState(52)
		p.Match(MangleParserNAME)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MangleParserLBRACKET {
		{
			p.SetState(53)
			p.Atoms()
		}

	}
	{
		p.SetState(56)
		p.Match(MangleParserBANG)
	}

	return localctx
}

// IUseDeclContext is an interface to support dynamic dispatch.
type IUseDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUseDeclContext differentiates from other interfaces.
	IsUseDeclContext()
}

type UseDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseDeclContext() *UseDeclContext {
	var p = new(UseDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_useDecl
	return p
}

func (*UseDeclContext) IsUseDeclContext() {}

func NewUseDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseDeclContext {
	var p = new(UseDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_useDecl

	return p
}

func (s *UseDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *UseDeclContext) USE() antlr.TerminalNode {
	return s.GetToken(MangleParserUSE, 0)
}

func (s *UseDeclContext) NAME() antlr.TerminalNode {
	return s.GetToken(MangleParserNAME, 0)
}

func (s *UseDeclContext) BANG() antlr.TerminalNode {
	return s.GetToken(MangleParserBANG, 0)
}

func (s *UseDeclContext) Atoms() IAtomsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomsContext)
}

func (s *UseDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterUseDecl(s)
	}
}

func (s *UseDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitUseDecl(s)
	}
}

func (s *UseDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitUseDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) UseDecl() (localctx IUseDeclContext) {
	this := p
	_ = this

	localctx = NewUseDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MangleParserRULE_useDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(MangleParserUSE)
	}
	{
		p.SetState(59)
		p.Match(MangleParserNAME)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MangleParserLBRACKET {
		{
			p.SetState(60)
			p.Atoms()
		}

	}
	{
		p.SetState(63)
		p.Match(MangleParserBANG)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) DECL() antlr.TerminalNode {
	return s.GetToken(MangleParserDECL, 0)
}

func (s *DeclContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *DeclContext) DescrBlock() IDescrBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescrBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescrBlockContext)
}

func (s *DeclContext) AllBoundsBlock() []IBoundsBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoundsBlockContext); ok {
			len++
		}
	}

	tst := make([]IBoundsBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoundsBlockContext); ok {
			tst[i] = t.(IBoundsBlockContext)
			i++
		}
	}

	return tst
}

func (s *DeclContext) BoundsBlock(i int) IBoundsBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoundsBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoundsBlockContext)
}

func (s *DeclContext) ConstraintsBlock() IConstraintsBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstraintsBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstraintsBlockContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (s *DeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MangleParserRULE_decl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(MangleParserDECL)
	}
	{
		p.SetState(66)
		p.Atom()
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MangleParserT__1 {
		{
			p.SetState(67)
			p.DescrBlock()
		}

	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MangleParserBOUND {
		{
			p.SetState(70)
			p.BoundsBlock()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MangleParserT__2 {
		{
			p.SetState(76)
			p.ConstraintsBlock()
		}

	}
	{
		p.SetState(79)
		p.Match(MangleParserT__0)
	}

	return localctx
}

// IDescrBlockContext is an interface to support dynamic dispatch.
type IDescrBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescrBlockContext differentiates from other interfaces.
	IsDescrBlockContext()
}

type DescrBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescrBlockContext() *DescrBlockContext {
	var p = new(DescrBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_descrBlock
	return p
}

func (*DescrBlockContext) IsDescrBlockContext() {}

func NewDescrBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescrBlockContext {
	var p = new(DescrBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_descrBlock

	return p
}

func (s *DescrBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *DescrBlockContext) Atoms() IAtomsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomsContext)
}

func (s *DescrBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescrBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescrBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterDescrBlock(s)
	}
}

func (s *DescrBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitDescrBlock(s)
	}
}

func (s *DescrBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitDescrBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) DescrBlock() (localctx IDescrBlockContext) {
	this := p
	_ = this

	localctx = NewDescrBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MangleParserRULE_descrBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(MangleParserT__1)
	}
	{
		p.SetState(82)
		p.Atoms()
	}

	return localctx
}

// IBoundsBlockContext is an interface to support dynamic dispatch.
type IBoundsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoundsBlockContext differentiates from other interfaces.
	IsBoundsBlockContext()
}

type BoundsBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundsBlockContext() *BoundsBlockContext {
	var p = new(BoundsBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_boundsBlock
	return p
}

func (*BoundsBlockContext) IsBoundsBlockContext() {}

func NewBoundsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundsBlockContext {
	var p = new(BoundsBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_boundsBlock

	return p
}

func (s *BoundsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundsBlockContext) BOUND() antlr.TerminalNode {
	return s.GetToken(MangleParserBOUND, 0)
}

func (s *BoundsBlockContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MangleParserLBRACKET, 0)
}

func (s *BoundsBlockContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *BoundsBlockContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *BoundsBlockContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MangleParserRBRACKET, 0)
}

func (s *BoundsBlockContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MangleParserCOMMA)
}

func (s *BoundsBlockContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MangleParserCOMMA, i)
}

func (s *BoundsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundsBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterBoundsBlock(s)
	}
}

func (s *BoundsBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitBoundsBlock(s)
	}
}

func (s *BoundsBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitBoundsBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) BoundsBlock() (localctx IBoundsBlockContext) {
	this := p
	_ = this

	localctx = NewBoundsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MangleParserRULE_boundsBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(MangleParserBOUND)
	}
	{
		p.SetState(85)
		p.Match(MangleParserLBRACKET)
	}
	{
		p.SetState(86)
		p.Term()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MangleParserCOMMA {
		{
			p.SetState(87)
			p.Match(MangleParserCOMMA)
		}
		{
			p.SetState(88)
			p.Term()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(MangleParserRBRACKET)
	}

	return localctx
}

// IConstraintsBlockContext is an interface to support dynamic dispatch.
type IConstraintsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstraintsBlockContext differentiates from other interfaces.
	IsConstraintsBlockContext()
}

type ConstraintsBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraintsBlockContext() *ConstraintsBlockContext {
	var p = new(ConstraintsBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_constraintsBlock
	return p
}

func (*ConstraintsBlockContext) IsConstraintsBlockContext() {}

func NewConstraintsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintsBlockContext {
	var p = new(ConstraintsBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_constraintsBlock

	return p
}

func (s *ConstraintsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintsBlockContext) Atoms() IAtomsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomsContext)
}

func (s *ConstraintsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstraintsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstraintsBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterConstraintsBlock(s)
	}
}

func (s *ConstraintsBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitConstraintsBlock(s)
	}
}

func (s *ConstraintsBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitConstraintsBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) ConstraintsBlock() (localctx IConstraintsBlockContext) {
	this := p
	_ = this

	localctx = NewConstraintsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MangleParserRULE_constraintsBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(MangleParserT__2)
	}
	{
		p.SetState(97)
		p.Atoms()
	}

	return localctx
}

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_clause
	return p
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ClauseContext) COLONDASH() antlr.TerminalNode {
	return s.GetToken(MangleParserCOLONDASH, 0)
}

func (s *ClauseContext) ClauseBody() IClauseBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClauseBodyContext)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitClause(s)
	}
}

func (s *ClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) Clause() (localctx IClauseContext) {
	this := p
	_ = this

	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MangleParserRULE_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Atom()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MangleParserCOLONDASH {
		{
			p.SetState(100)
			p.Match(MangleParserCOLONDASH)
		}
		{
			p.SetState(101)
			p.ClauseBody()
		}

	}
	{
		p.SetState(104)
		p.Match(MangleParserT__0)
	}

	return localctx
}

// IClauseBodyContext is an interface to support dynamic dispatch.
type IClauseBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClauseBodyContext differentiates from other interfaces.
	IsClauseBodyContext()
}

type ClauseBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseBodyContext() *ClauseBodyContext {
	var p = new(ClauseBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_clauseBody
	return p
}

func (*ClauseBodyContext) IsClauseBodyContext() {}

func NewClauseBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseBodyContext {
	var p = new(ClauseBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_clauseBody

	return p
}

func (s *ClauseBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseBodyContext) AllLiteralOrFml() []ILiteralOrFmlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralOrFmlContext); ok {
			len++
		}
	}

	tst := make([]ILiteralOrFmlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralOrFmlContext); ok {
			tst[i] = t.(ILiteralOrFmlContext)
			i++
		}
	}

	return tst
}

func (s *ClauseBodyContext) LiteralOrFml(i int) ILiteralOrFmlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralOrFmlContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralOrFmlContext)
}

func (s *ClauseBodyContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MangleParserCOMMA)
}

func (s *ClauseBodyContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MangleParserCOMMA, i)
}

func (s *ClauseBodyContext) PIPEGREATER() antlr.TerminalNode {
	return s.GetToken(MangleParserPIPEGREATER, 0)
}

func (s *ClauseBodyContext) Transform() ITransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransformContext)
}

func (s *ClauseBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterClauseBody(s)
	}
}

func (s *ClauseBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitClauseBody(s)
	}
}

func (s *ClauseBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitClauseBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) ClauseBody() (localctx IClauseBodyContext) {
	this := p
	_ = this

	localctx = NewClauseBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MangleParserRULE_clauseBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.LiteralOrFml()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MangleParserCOMMA {
		{
			p.SetState(107)
			p.Match(MangleParserCOMMA)
		}
		{
			p.SetState(108)
			p.LiteralOrFml()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == MangleParserPIPEGREATER {
		{
			p.SetState(114)
			p.Match(MangleParserPIPEGREATER)
		}
		{
			p.SetState(115)
			p.Transform()
		}

	}

	return localctx
}

// ITransformContext is an interface to support dynamic dispatch.
type ITransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransformContext differentiates from other interfaces.
	IsTransformContext()
}

type TransformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformContext() *TransformContext {
	var p = new(TransformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_transform
	return p
}

func (*TransformContext) IsTransformContext() {}

func NewTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransformContext {
	var p = new(TransformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_transform

	return p
}

func (s *TransformContext) GetParser() antlr.Parser { return s.parser }

func (s *TransformContext) DO() antlr.TerminalNode {
	return s.GetToken(MangleParserDO, 0)
}

func (s *TransformContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TransformContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MangleParserCOMMA)
}

func (s *TransformContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MangleParserCOMMA, i)
}

func (s *TransformContext) AllLetStmt() []ILetStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILetStmtContext); ok {
			len++
		}
	}

	tst := make([]ILetStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILetStmtContext); ok {
			tst[i] = t.(ILetStmtContext)
			i++
		}
	}

	return tst
}

func (s *TransformContext) LetStmt(i int) ILetStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetStmtContext)
}

func (s *TransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterTransform(s)
	}
}

func (s *TransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitTransform(s)
	}
}

func (s *TransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) Transform() (localctx ITransformContext) {
	this := p
	_ = this

	localctx = NewTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MangleParserRULE_transform)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MangleParserDO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.Match(MangleParserDO)
		}
		{
			p.SetState(119)
			p.Term()
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == MangleParserCOMMA {
			{
				p.SetState(120)
				p.Match(MangleParserCOMMA)
			}
			{
				p.SetState(121)
				p.LetStmt()
			}
			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == MangleParserCOMMA {
				{
					p.SetState(122)
					p.Match(MangleParserCOMMA)
				}
				{
					p.SetState(123)
					p.LetStmt()
				}

				p.SetState(128)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}

	case MangleParserLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.LetStmt()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MangleParserCOMMA {
			{
				p.SetState(132)
				p.Match(MangleParserCOMMA)
			}
			{
				p.SetState(133)
				p.LetStmt()
			}

			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILetStmtContext is an interface to support dynamic dispatch.
type ILetStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLetStmtContext differentiates from other interfaces.
	IsLetStmtContext()
}

type LetStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetStmtContext() *LetStmtContext {
	var p = new(LetStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_letStmt
	return p
}

func (*LetStmtContext) IsLetStmtContext() {}

func NewLetStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetStmtContext {
	var p = new(LetStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_letStmt

	return p
}

func (s *LetStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *LetStmtContext) LET() antlr.TerminalNode {
	return s.GetToken(MangleParserLET, 0)
}

func (s *LetStmtContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(MangleParserVARIABLE, 0)
}

func (s *LetStmtContext) EQ() antlr.TerminalNode {
	return s.GetToken(MangleParserEQ, 0)
}

func (s *LetStmtContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *LetStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterLetStmt(s)
	}
}

func (s *LetStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitLetStmt(s)
	}
}

func (s *LetStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitLetStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) LetStmt() (localctx ILetStmtContext) {
	this := p
	_ = this

	localctx = NewLetStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MangleParserRULE_letStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(MangleParserLET)
	}
	{
		p.SetState(142)
		p.Match(MangleParserVARIABLE)
	}
	{
		p.SetState(143)
		p.Match(MangleParserEQ)
	}
	{
		p.SetState(144)
		p.Term()
	}

	return localctx
}

// ILiteralOrFmlContext is an interface to support dynamic dispatch.
type ILiteralOrFmlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralOrFmlContext differentiates from other interfaces.
	IsLiteralOrFmlContext()
}

type LiteralOrFmlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralOrFmlContext() *LiteralOrFmlContext {
	var p = new(LiteralOrFmlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_literalOrFml
	return p
}

func (*LiteralOrFmlContext) IsLiteralOrFmlContext() {}

func NewLiteralOrFmlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralOrFmlContext {
	var p = new(LiteralOrFmlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_literalOrFml

	return p
}

func (s *LiteralOrFmlContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralOrFmlContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *LiteralOrFmlContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *LiteralOrFmlContext) EQ() antlr.TerminalNode {
	return s.GetToken(MangleParserEQ, 0)
}

func (s *LiteralOrFmlContext) BANGEQ() antlr.TerminalNode {
	return s.GetToken(MangleParserBANGEQ, 0)
}

func (s *LiteralOrFmlContext) LESS() antlr.TerminalNode {
	return s.GetToken(MangleParserLESS, 0)
}

func (s *LiteralOrFmlContext) LESSEQ() antlr.TerminalNode {
	return s.GetToken(MangleParserLESSEQ, 0)
}

func (s *LiteralOrFmlContext) BANG() antlr.TerminalNode {
	return s.GetToken(MangleParserBANG, 0)
}

func (s *LiteralOrFmlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralOrFmlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralOrFmlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterLiteralOrFml(s)
	}
}

func (s *LiteralOrFmlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitLiteralOrFml(s)
	}
}

func (s *LiteralOrFmlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitLiteralOrFml(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) LiteralOrFml() (localctx ILiteralOrFmlContext) {
	this := p
	_ = this

	localctx = NewLiteralOrFmlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MangleParserRULE_literalOrFml)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MangleParserT__4, MangleParserLBRACKET, MangleParserNUMBER, MangleParserVARIABLE, MangleParserNAME, MangleParserCONSTANT, MangleParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Term()
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26738688) != 0 {
			{
				p.SetState(147)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26738688) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(148)
				p.Term()
			}

		}

	case MangleParserBANG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Match(MangleParserBANG)
		}
		{
			p.SetState(152)
			p.Term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) CopyFrom(ctx *TermContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StrContext struct {
	*TermContext
}

func NewStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrContext {
	var p = new(StrContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) STRING() antlr.TerminalNode {
	return s.GetToken(MangleParserSTRING, 0)
}

func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitStr(s)
	}
}

func (s *StrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitStr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ApplContext struct {
	*TermContext
}

func NewApplContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ApplContext {
	var p = new(ApplContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *ApplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApplContext) NAME() antlr.TerminalNode {
	return s.GetToken(MangleParserNAME, 0)
}

func (s *ApplContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MangleParserLPAREN, 0)
}

func (s *ApplContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MangleParserRPAREN, 0)
}

func (s *ApplContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ApplContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ApplContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MangleParserCOMMA)
}

func (s *ApplContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MangleParserCOMMA, i)
}

func (s *ApplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterAppl(s)
	}
}

func (s *ApplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitAppl(s)
	}
}

func (s *ApplContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitAppl(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarContext struct {
	*TermContext
}

func NewVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarContext {
	var p = new(VarContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(MangleParserVARIABLE, 0)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitVar(s)
	}
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstContext struct {
	*TermContext
}

func NewConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstContext {
	var p = new(ConstContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *ConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstContext) CONSTANT() antlr.TerminalNode {
	return s.GetToken(MangleParserCONSTANT, 0)
}

func (s *ConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterConst(s)
	}
}

func (s *ConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitConst(s)
	}
}

func (s *ConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitConst(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumContext struct {
	*TermContext
}

func NewNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumContext {
	var p = new(NumContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MangleParserNUMBER, 0)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitNum(s)
	}
}

func (s *NumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitNum(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListContext struct {
	*TermContext
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MangleParserLBRACKET, 0)
}

func (s *ListContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MangleParserRBRACKET, 0)
}

func (s *ListContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MangleParserCOMMA)
}

func (s *ListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MangleParserCOMMA, i)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

type MapContext struct {
	*TermContext
}

func NewMapContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapContext {
	var p = new(MapContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *MapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MangleParserLBRACKET, 0)
}

func (s *MapContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MangleParserRBRACKET, 0)
}

func (s *MapContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *MapContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *MapContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MangleParserCOMMA)
}

func (s *MapContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MangleParserCOMMA, i)
}

func (s *MapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterMap(s)
	}
}

func (s *MapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitMap(s)
	}
}

func (s *MapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitMap(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructContext struct {
	*TermContext
}

func NewStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructContext {
	var p = new(StructContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *StructContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *StructContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MangleParserCOMMA)
}

func (s *StructContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MangleParserCOMMA, i)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (s *StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MangleParserRULE_term)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Match(MangleParserVARIABLE)
		}

	case 2:
		localctx = NewConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(MangleParserCONSTANT)
		}

	case 3:
		localctx = NewNumContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.Match(MangleParserNUMBER)
		}

	case 4:
		localctx = NewStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.Match(MangleParserSTRING)
		}

	case 5:
		localctx = NewApplContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(159)
			p.Match(MangleParserNAME)
		}
		{
			p.SetState(160)
			p.Match(MangleParserLPAREN)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8321630240) != 0 {
			{
				p.SetState(161)
				p.Term()
			}
			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == MangleParserCOMMA {
				{
					p.SetState(162)
					p.Match(MangleParserCOMMA)
				}
				{
					p.SetState(163)
					p.Term()
				}

				p.SetState(168)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(171)
			p.Match(MangleParserRPAREN)
		}

	case 6:
		localctx = NewListContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(172)
			p.Match(MangleParserLBRACKET)
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8321630240) != 0 {
			{
				p.SetState(173)
				p.Term()
			}
			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == MangleParserCOMMA {
				{
					p.SetState(174)
					p.Match(MangleParserCOMMA)
				}
				{
					p.SetState(175)
					p.Term()
				}

				p.SetState(180)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(183)
			p.Match(MangleParserRBRACKET)
		}

	case 7:
		localctx = NewMapContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(184)
			p.Match(MangleParserLBRACKET)
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8321630240) != 0 {
			{
				p.SetState(185)
				p.Term()
			}
			{
				p.SetState(186)
				p.Match(MangleParserT__3)
			}
			{
				p.SetState(187)
				p.Term()
			}
			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == MangleParserCOMMA {
				{
					p.SetState(188)
					p.Match(MangleParserCOMMA)
				}
				{
					p.SetState(189)
					p.Term()
				}
				{
					p.SetState(190)
					p.Match(MangleParserT__3)
				}
				{
					p.SetState(191)
					p.Term()
				}

				p.SetState(197)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(200)
			p.Match(MangleParserRBRACKET)
		}

	case 8:
		localctx = NewStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(201)
			p.Match(MangleParserT__4)
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8321630240) != 0 {
			{
				p.SetState(202)
				p.Term()
			}
			{
				p.SetState(203)
				p.Match(MangleParserT__3)
			}
			{
				p.SetState(204)
				p.Term()
			}
			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == MangleParserCOMMA {
				{
					p.SetState(205)
					p.Match(MangleParserCOMMA)
				}
				{
					p.SetState(206)
					p.Term()
				}
				{
					p.SetState(207)
					p.Match(MangleParserT__3)
				}
				{
					p.SetState(208)
					p.Term()
				}

				p.SetState(214)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(217)
			p.Match(MangleParserT__5)
		}

	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MangleParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Term()
	}

	return localctx
}

// IAtomsContext is an interface to support dynamic dispatch.
type IAtomsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomsContext differentiates from other interfaces.
	IsAtomsContext()
}

type AtomsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomsContext() *AtomsContext {
	var p = new(AtomsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MangleParserRULE_atoms
	return p
}

func (*AtomsContext) IsAtomsContext() {}

func NewAtomsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomsContext {
	var p = new(AtomsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MangleParserRULE_atoms

	return p
}

func (s *AtomsContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomsContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MangleParserLBRACKET, 0)
}

func (s *AtomsContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MangleParserRBRACKET, 0)
}

func (s *AtomsContext) AllAtom() []IAtomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtomContext); ok {
			len++
		}
	}

	tst := make([]IAtomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtomContext); ok {
			tst[i] = t.(IAtomContext)
			i++
		}
	}

	return tst
}

func (s *AtomsContext) Atom(i int) IAtomContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *AtomsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MangleParserCOMMA)
}

func (s *AtomsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MangleParserCOMMA, i)
}

func (s *AtomsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.EnterAtoms(s)
	}
}

func (s *AtomsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MangleListener); ok {
		listenerT.ExitAtoms(s)
	}
}

func (s *AtomsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MangleVisitor:
		return t.VisitAtoms(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MangleParser) Atoms() (localctx IAtomsContext) {
	this := p
	_ = this

	localctx = NewAtomsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MangleParserRULE_atoms)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(MangleParserLBRACKET)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8321630240) != 0 {
		{
			p.SetState(223)
			p.Atom()
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == MangleParserCOMMA {
			{
				p.SetState(224)
				p.Match(MangleParserCOMMA)
			}
			{
				p.SetState(225)
				p.Atom()
			}

			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(233)
		p.Match(MangleParserRBRACKET)
	}

	return localctx
}
