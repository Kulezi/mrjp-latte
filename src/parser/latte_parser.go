// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Latte

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

type LatteParser struct {
	*antlr.BaseParser
}

var latteParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func latteParserInit() {
	staticData := &latteParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'class'", "'{'", "'}'", "'extends'", "','", "';'",
		"'['", "']'", "'.'", "'='", "'++'", "'--'", "'return'", "'if'", "'else'",
		"'while'", "'for'", "':'", "'void'", "'[]'", "'int'", "'string'", "'boolean'",
		"'-'", "'!'", "'&&'", "'||'", "'new'", "'self'", "'true'", "'false'",
		"'null'", "'+'", "'*'", "'/'", "'%'", "'<'", "'<='", "'>'", "'>='",
		"'=='", "'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "COMMENT", "MULTICOMMENT",
		"INT", "ID", "WS", "STR",
	}
	staticData.ruleNames = []string{
		"program", "topDef", "fundef", "classdef", "arg", "field", "block",
		"lvalue", "stmt", "type_", "nvtype_", "singular_type_", "item", "expr",
		"addOp", "mulOp", "relOp",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 319, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 4, 0, 36, 8, 0, 11, 0, 12, 0, 37, 1, 1, 1, 1, 3, 1,
		42, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 48, 8, 2, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 5, 3, 57, 8, 3, 10, 3, 12, 3, 60, 9, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 69, 8, 3, 10, 3, 12, 3, 72, 9, 3, 1,
		3, 3, 3, 75, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 83, 8, 4,
		10, 4, 12, 4, 86, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 93, 8, 5, 1,
		6, 1, 6, 5, 6, 97, 8, 6, 10, 6, 12, 6, 100, 9, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 128,
		8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 136, 8, 8, 10, 8, 12, 8,
		139, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 194, 8, 8, 1, 9, 1, 9,
		3, 9, 198, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 204, 8, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 3, 11, 210, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3,
		12, 216, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 241, 8, 13, 10, 13, 12, 13, 244,
		9, 13, 3, 13, 246, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 258, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		5, 13, 285, 8, 13, 10, 13, 12, 13, 288, 9, 13, 3, 13, 290, 8, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 308, 8, 13, 10, 13, 12, 13, 311,
		9, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 0, 1, 26, 17, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 3, 2, 0,
		26, 26, 35, 35, 1, 0, 36, 38, 1, 0, 39, 44, 358, 0, 35, 1, 0, 0, 0, 2,
		41, 1, 0, 0, 0, 4, 43, 1, 0, 0, 0, 6, 74, 1, 0, 0, 0, 8, 76, 1, 0, 0, 0,
		10, 92, 1, 0, 0, 0, 12, 94, 1, 0, 0, 0, 14, 127, 1, 0, 0, 0, 16, 193, 1,
		0, 0, 0, 18, 197, 1, 0, 0, 0, 20, 203, 1, 0, 0, 0, 22, 209, 1, 0, 0, 0,
		24, 215, 1, 0, 0, 0, 26, 257, 1, 0, 0, 0, 28, 312, 1, 0, 0, 0, 30, 314,
		1, 0, 0, 0, 32, 316, 1, 0, 0, 0, 34, 36, 3, 2, 1, 0, 35, 34, 1, 0, 0, 0,
		36, 37, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 1, 1, 0,
		0, 0, 39, 42, 3, 6, 3, 0, 40, 42, 3, 4, 2, 0, 41, 39, 1, 0, 0, 0, 41, 40,
		1, 0, 0, 0, 42, 3, 1, 0, 0, 0, 43, 44, 3, 18, 9, 0, 44, 45, 5, 48, 0, 0,
		45, 47, 5, 1, 0, 0, 46, 48, 3, 8, 4, 0, 47, 46, 1, 0, 0, 0, 47, 48, 1,
		0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 5, 2, 0, 0, 50, 51, 3, 12, 6, 0, 51,
		5, 1, 0, 0, 0, 52, 53, 5, 3, 0, 0, 53, 54, 5, 48, 0, 0, 54, 58, 5, 4, 0,
		0, 55, 57, 3, 10, 5, 0, 56, 55, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56,
		1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 61, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0,
		61, 75, 5, 5, 0, 0, 62, 63, 5, 3, 0, 0, 63, 64, 5, 48, 0, 0, 64, 65, 5,
		6, 0, 0, 65, 66, 5, 48, 0, 0, 66, 70, 5, 4, 0, 0, 67, 69, 3, 10, 5, 0,
		68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1,
		0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 75, 5, 5, 0, 0, 74,
		52, 1, 0, 0, 0, 74, 62, 1, 0, 0, 0, 75, 7, 1, 0, 0, 0, 76, 77, 3, 20, 10,
		0, 77, 84, 5, 48, 0, 0, 78, 79, 5, 7, 0, 0, 79, 80, 3, 20, 10, 0, 80, 81,
		5, 48, 0, 0, 81, 83, 1, 0, 0, 0, 82, 78, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0,
		84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 9, 1, 0, 0, 0, 86, 84, 1, 0,
		0, 0, 87, 88, 3, 20, 10, 0, 88, 89, 5, 48, 0, 0, 89, 90, 5, 8, 0, 0, 90,
		93, 1, 0, 0, 0, 91, 93, 3, 4, 2, 0, 92, 87, 1, 0, 0, 0, 92, 91, 1, 0, 0,
		0, 93, 11, 1, 0, 0, 0, 94, 98, 5, 4, 0, 0, 95, 97, 3, 16, 8, 0, 96, 95,
		1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0,
		99, 101, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 102, 5, 5, 0, 0, 102, 13,
		1, 0, 0, 0, 103, 104, 3, 26, 13, 0, 104, 105, 5, 9, 0, 0, 105, 106, 3,
		26, 13, 0, 106, 107, 5, 10, 0, 0, 107, 128, 1, 0, 0, 0, 108, 109, 3, 26,
		13, 0, 109, 110, 5, 11, 0, 0, 110, 111, 5, 48, 0, 0, 111, 112, 5, 9, 0,
		0, 112, 113, 3, 26, 13, 0, 113, 114, 5, 10, 0, 0, 114, 128, 1, 0, 0, 0,
		115, 116, 3, 26, 13, 0, 116, 117, 5, 11, 0, 0, 117, 118, 5, 48, 0, 0, 118,
		119, 5, 1, 0, 0, 119, 120, 3, 26, 13, 0, 120, 121, 5, 2, 0, 0, 121, 128,
		1, 0, 0, 0, 122, 123, 3, 26, 13, 0, 123, 124, 5, 11, 0, 0, 124, 125, 5,
		48, 0, 0, 125, 128, 1, 0, 0, 0, 126, 128, 5, 48, 0, 0, 127, 103, 1, 0,
		0, 0, 127, 108, 1, 0, 0, 0, 127, 115, 1, 0, 0, 0, 127, 122, 1, 0, 0, 0,
		127, 126, 1, 0, 0, 0, 128, 15, 1, 0, 0, 0, 129, 194, 5, 8, 0, 0, 130, 194,
		3, 12, 6, 0, 131, 132, 3, 20, 10, 0, 132, 137, 3, 24, 12, 0, 133, 134,
		5, 7, 0, 0, 134, 136, 3, 24, 12, 0, 135, 133, 1, 0, 0, 0, 136, 139, 1,
		0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 140, 1, 0, 0,
		0, 139, 137, 1, 0, 0, 0, 140, 141, 5, 8, 0, 0, 141, 194, 1, 0, 0, 0, 142,
		143, 3, 14, 7, 0, 143, 144, 5, 12, 0, 0, 144, 145, 3, 26, 13, 0, 145, 146,
		5, 8, 0, 0, 146, 194, 1, 0, 0, 0, 147, 148, 3, 14, 7, 0, 148, 149, 5, 13,
		0, 0, 149, 150, 5, 8, 0, 0, 150, 194, 1, 0, 0, 0, 151, 152, 3, 14, 7, 0,
		152, 153, 5, 14, 0, 0, 153, 154, 5, 8, 0, 0, 154, 194, 1, 0, 0, 0, 155,
		156, 5, 15, 0, 0, 156, 157, 3, 26, 13, 0, 157, 158, 5, 8, 0, 0, 158, 194,
		1, 0, 0, 0, 159, 160, 5, 15, 0, 0, 160, 194, 5, 8, 0, 0, 161, 162, 5, 16,
		0, 0, 162, 163, 5, 1, 0, 0, 163, 164, 3, 26, 13, 0, 164, 165, 5, 2, 0,
		0, 165, 166, 3, 16, 8, 0, 166, 194, 1, 0, 0, 0, 167, 168, 5, 16, 0, 0,
		168, 169, 5, 1, 0, 0, 169, 170, 3, 26, 13, 0, 170, 171, 5, 2, 0, 0, 171,
		172, 3, 16, 8, 0, 172, 173, 5, 17, 0, 0, 173, 174, 3, 16, 8, 0, 174, 194,
		1, 0, 0, 0, 175, 176, 5, 18, 0, 0, 176, 177, 5, 1, 0, 0, 177, 178, 3, 26,
		13, 0, 178, 179, 5, 2, 0, 0, 179, 180, 3, 16, 8, 0, 180, 194, 1, 0, 0,
		0, 181, 182, 5, 19, 0, 0, 182, 183, 5, 1, 0, 0, 183, 184, 3, 18, 9, 0,
		184, 185, 5, 48, 0, 0, 185, 186, 5, 20, 0, 0, 186, 187, 3, 26, 13, 0, 187,
		188, 5, 2, 0, 0, 188, 189, 3, 16, 8, 0, 189, 194, 1, 0, 0, 0, 190, 191,
		3, 26, 13, 0, 191, 192, 5, 8, 0, 0, 192, 194, 1, 0, 0, 0, 193, 129, 1,
		0, 0, 0, 193, 130, 1, 0, 0, 0, 193, 131, 1, 0, 0, 0, 193, 142, 1, 0, 0,
		0, 193, 147, 1, 0, 0, 0, 193, 151, 1, 0, 0, 0, 193, 155, 1, 0, 0, 0, 193,
		159, 1, 0, 0, 0, 193, 161, 1, 0, 0, 0, 193, 167, 1, 0, 0, 0, 193, 175,
		1, 0, 0, 0, 193, 181, 1, 0, 0, 0, 193, 190, 1, 0, 0, 0, 194, 17, 1, 0,
		0, 0, 195, 198, 3, 20, 10, 0, 196, 198, 5, 21, 0, 0, 197, 195, 1, 0, 0,
		0, 197, 196, 1, 0, 0, 0, 198, 19, 1, 0, 0, 0, 199, 200, 3, 22, 11, 0, 200,
		201, 5, 22, 0, 0, 201, 204, 1, 0, 0, 0, 202, 204, 3, 22, 11, 0, 203, 199,
		1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 21, 1, 0, 0, 0, 205, 210, 5, 48,
		0, 0, 206, 210, 5, 23, 0, 0, 207, 210, 5, 24, 0, 0, 208, 210, 5, 25, 0,
		0, 209, 205, 1, 0, 0, 0, 209, 206, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209,
		208, 1, 0, 0, 0, 210, 23, 1, 0, 0, 0, 211, 216, 5, 48, 0, 0, 212, 213,
		5, 48, 0, 0, 213, 214, 5, 12, 0, 0, 214, 216, 3, 26, 13, 0, 215, 211, 1,
		0, 0, 0, 215, 212, 1, 0, 0, 0, 216, 25, 1, 0, 0, 0, 217, 218, 6, 13, -1,
		0, 218, 219, 5, 26, 0, 0, 219, 258, 3, 26, 13, 18, 220, 221, 5, 27, 0,
		0, 221, 258, 3, 26, 13, 17, 222, 223, 5, 30, 0, 0, 223, 224, 3, 22, 11,
		0, 224, 225, 5, 9, 0, 0, 225, 226, 3, 26, 13, 0, 226, 227, 5, 10, 0, 0,
		227, 258, 1, 0, 0, 0, 228, 229, 5, 30, 0, 0, 229, 258, 3, 22, 11, 0, 230,
		258, 5, 31, 0, 0, 231, 258, 5, 48, 0, 0, 232, 258, 5, 47, 0, 0, 233, 258,
		5, 32, 0, 0, 234, 258, 5, 33, 0, 0, 235, 236, 5, 48, 0, 0, 236, 245, 5,
		1, 0, 0, 237, 242, 3, 26, 13, 0, 238, 239, 5, 7, 0, 0, 239, 241, 3, 26,
		13, 0, 240, 238, 1, 0, 0, 0, 241, 244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0,
		242, 243, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245,
		237, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 258,
		5, 2, 0, 0, 248, 258, 5, 50, 0, 0, 249, 250, 5, 1, 0, 0, 250, 251, 5, 48,
		0, 0, 251, 252, 5, 2, 0, 0, 252, 258, 5, 34, 0, 0, 253, 254, 5, 1, 0, 0,
		254, 255, 3, 26, 13, 0, 255, 256, 5, 2, 0, 0, 256, 258, 1, 0, 0, 0, 257,
		217, 1, 0, 0, 0, 257, 220, 1, 0, 0, 0, 257, 222, 1, 0, 0, 0, 257, 228,
		1, 0, 0, 0, 257, 230, 1, 0, 0, 0, 257, 231, 1, 0, 0, 0, 257, 232, 1, 0,
		0, 0, 257, 233, 1, 0, 0, 0, 257, 234, 1, 0, 0, 0, 257, 235, 1, 0, 0, 0,
		257, 248, 1, 0, 0, 0, 257, 249, 1, 0, 0, 0, 257, 253, 1, 0, 0, 0, 258,
		309, 1, 0, 0, 0, 259, 260, 10, 16, 0, 0, 260, 261, 3, 30, 15, 0, 261, 262,
		3, 26, 13, 17, 262, 308, 1, 0, 0, 0, 263, 264, 10, 15, 0, 0, 264, 265,
		3, 28, 14, 0, 265, 266, 3, 26, 13, 16, 266, 308, 1, 0, 0, 0, 267, 268,
		10, 14, 0, 0, 268, 269, 3, 32, 16, 0, 269, 270, 3, 26, 13, 15, 270, 308,
		1, 0, 0, 0, 271, 272, 10, 13, 0, 0, 272, 273, 5, 28, 0, 0, 273, 308, 3,
		26, 13, 13, 274, 275, 10, 12, 0, 0, 275, 276, 5, 29, 0, 0, 276, 308, 3,
		26, 13, 12, 277, 278, 10, 22, 0, 0, 278, 279, 5, 11, 0, 0, 279, 280, 5,
		48, 0, 0, 280, 289, 5, 1, 0, 0, 281, 286, 3, 26, 13, 0, 282, 283, 5, 7,
		0, 0, 283, 285, 3, 26, 13, 0, 284, 282, 1, 0, 0, 0, 285, 288, 1, 0, 0,
		0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288,
		286, 1, 0, 0, 0, 289, 281, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291,
		1, 0, 0, 0, 291, 308, 5, 2, 0, 0, 292, 293, 10, 21, 0, 0, 293, 294, 5,
		11, 0, 0, 294, 295, 5, 48, 0, 0, 295, 296, 5, 9, 0, 0, 296, 297, 3, 26,
		13, 0, 297, 298, 5, 10, 0, 0, 298, 308, 1, 0, 0, 0, 299, 300, 10, 20, 0,
		0, 300, 301, 5, 11, 0, 0, 301, 308, 5, 48, 0, 0, 302, 303, 10, 19, 0, 0,
		303, 304, 5, 9, 0, 0, 304, 305, 3, 26, 13, 0, 305, 306, 5, 10, 0, 0, 306,
		308, 1, 0, 0, 0, 307, 259, 1, 0, 0, 0, 307, 263, 1, 0, 0, 0, 307, 267,
		1, 0, 0, 0, 307, 271, 1, 0, 0, 0, 307, 274, 1, 0, 0, 0, 307, 277, 1, 0,
		0, 0, 307, 292, 1, 0, 0, 0, 307, 299, 1, 0, 0, 0, 307, 302, 1, 0, 0, 0,
		308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310,
		27, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 313, 7, 0, 0, 0, 313, 29, 1,
		0, 0, 0, 314, 315, 7, 1, 0, 0, 315, 31, 1, 0, 0, 0, 316, 317, 7, 2, 0,
		0, 317, 33, 1, 0, 0, 0, 23, 37, 41, 47, 58, 70, 74, 84, 92, 98, 127, 137,
		193, 197, 203, 209, 215, 242, 245, 257, 286, 289, 307, 309,
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

// LatteParserInit initializes any static state used to implement LatteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLatteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LatteParserInit() {
	staticData := &latteParserStaticData
	staticData.once.Do(latteParserInit)
}

// NewLatteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLatteParser(input antlr.TokenStream) *LatteParser {
	LatteParserInit()
	this := new(LatteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &latteParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// LatteParser tokens.
const (
	LatteParserEOF          = antlr.TokenEOF
	LatteParserT__0         = 1
	LatteParserT__1         = 2
	LatteParserT__2         = 3
	LatteParserT__3         = 4
	LatteParserT__4         = 5
	LatteParserT__5         = 6
	LatteParserT__6         = 7
	LatteParserT__7         = 8
	LatteParserT__8         = 9
	LatteParserT__9         = 10
	LatteParserT__10        = 11
	LatteParserT__11        = 12
	LatteParserT__12        = 13
	LatteParserT__13        = 14
	LatteParserT__14        = 15
	LatteParserT__15        = 16
	LatteParserT__16        = 17
	LatteParserT__17        = 18
	LatteParserT__18        = 19
	LatteParserT__19        = 20
	LatteParserT__20        = 21
	LatteParserT__21        = 22
	LatteParserT__22        = 23
	LatteParserT__23        = 24
	LatteParserT__24        = 25
	LatteParserT__25        = 26
	LatteParserT__26        = 27
	LatteParserT__27        = 28
	LatteParserT__28        = 29
	LatteParserT__29        = 30
	LatteParserT__30        = 31
	LatteParserT__31        = 32
	LatteParserT__32        = 33
	LatteParserT__33        = 34
	LatteParserT__34        = 35
	LatteParserT__35        = 36
	LatteParserT__36        = 37
	LatteParserT__37        = 38
	LatteParserT__38        = 39
	LatteParserT__39        = 40
	LatteParserT__40        = 41
	LatteParserT__41        = 42
	LatteParserT__42        = 43
	LatteParserT__43        = 44
	LatteParserCOMMENT      = 45
	LatteParserMULTICOMMENT = 46
	LatteParserINT          = 47
	LatteParserID           = 48
	LatteParserWS           = 49
	LatteParserSTR          = 50
)

// LatteParser rules.
const (
	LatteParserRULE_program        = 0
	LatteParserRULE_topDef         = 1
	LatteParserRULE_fundef         = 2
	LatteParserRULE_classdef       = 3
	LatteParserRULE_arg            = 4
	LatteParserRULE_field          = 5
	LatteParserRULE_block          = 6
	LatteParserRULE_lvalue         = 7
	LatteParserRULE_stmt           = 8
	LatteParserRULE_type_          = 9
	LatteParserRULE_nvtype_        = 10
	LatteParserRULE_singular_type_ = 11
	LatteParserRULE_item           = 12
	LatteParserRULE_expr           = 13
	LatteParserRULE_addOp          = 14
	LatteParserRULE_mulOp          = 15
	LatteParserRULE_relOp          = 16
)

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
	p.RuleIndex = LatteParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllTopDef() []ITopDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopDefContext); ok {
			len++
		}
	}

	tst := make([]ITopDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopDefContext); ok {
			tst[i] = t.(ITopDefContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) TopDef(i int) ITopDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopDefContext); ok {
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

	return t.(ITopDefContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LatteParserRULE_program)
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281475037528072) != 0 {
		{
			p.SetState(34)
			p.TopDef()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITopDefContext is an interface to support dynamic dispatch.
type ITopDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopDefContext differentiates from other interfaces.
	IsTopDefContext()
}

type TopDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDefContext() *TopDefContext {
	var p = new(TopDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_topDef
	return p
}

func (*TopDefContext) IsTopDefContext() {}

func NewTopDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDefContext {
	var p = new(TopDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_topDef

	return p
}

func (s *TopDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDefContext) Classdef() IClassdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassdefContext)
}

func (s *TopDefContext) Fundef() IFundefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFundefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFundefContext)
}

func (s *TopDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTopDef(s)
	}
}

func (s *TopDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTopDef(s)
	}
}

func (s *TopDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTopDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) TopDef() (localctx ITopDefContext) {
	this := p
	_ = this

	localctx = NewTopDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LatteParserRULE_topDef)

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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LatteParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Classdef()
		}

	case LatteParserT__20, LatteParserT__22, LatteParserT__23, LatteParserT__24, LatteParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Fundef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFundefContext is an interface to support dynamic dispatch.
type IFundefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFundefContext differentiates from other interfaces.
	IsFundefContext()
}

type FundefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFundefContext() *FundefContext {
	var p = new(FundefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_fundef
	return p
}

func (*FundefContext) IsFundefContext() {}

func NewFundefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FundefContext {
	var p = new(FundefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_fundef

	return p
}

func (s *FundefContext) GetParser() antlr.Parser { return s.parser }

func (s *FundefContext) CopyFrom(ctx *FundefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FundefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FundefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunDefContext struct {
	*FundefContext
}

func NewFunDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunDefContext {
	var p = new(FunDefContext)

	p.FundefContext = NewEmptyFundefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FundefContext))

	return p
}

func (s *FunDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunDefContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FunDefContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *FunDefContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunDefContext) Arg() IArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *FunDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterFunDef(s)
	}
}

func (s *FunDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitFunDef(s)
	}
}

func (s *FunDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitFunDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Fundef() (localctx IFundefContext) {
	this := p
	_ = this

	localctx = NewFundefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LatteParserRULE_fundef)
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

	localctx = NewFunDefContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Type_()
	}
	{
		p.SetState(44)
		p.Match(LatteParserID)
	}
	{
		p.SetState(45)
		p.Match(LatteParserT__0)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281475035430912) != 0 {
		{
			p.SetState(46)
			p.Arg()
		}

	}
	{
		p.SetState(49)
		p.Match(LatteParserT__1)
	}
	{
		p.SetState(50)
		p.Block()
	}

	return localctx
}

// IClassdefContext is an interface to support dynamic dispatch.
type IClassdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassdefContext differentiates from other interfaces.
	IsClassdefContext()
}

type ClassdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassdefContext() *ClassdefContext {
	var p = new(ClassdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_classdef
	return p
}

func (*ClassdefContext) IsClassdefContext() {}

func NewClassdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassdefContext {
	var p = new(ClassdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_classdef

	return p
}

func (s *ClassdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassdefContext) CopyFrom(ctx *ClassdefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ClassdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DerivedClassDefContext struct {
	*ClassdefContext
}

func NewDerivedClassDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DerivedClassDefContext {
	var p = new(DerivedClassDefContext)

	p.ClassdefContext = NewEmptyClassdefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ClassdefContext))

	return p
}

func (s *DerivedClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerivedClassDefContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LatteParserID)
}

func (s *DerivedClassDefContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LatteParserID, i)
}

func (s *DerivedClassDefContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *DerivedClassDefContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *DerivedClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterDerivedClassDef(s)
	}
}

func (s *DerivedClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitDerivedClassDef(s)
	}
}

func (s *DerivedClassDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitDerivedClassDef(s)

	default:
		return t.VisitChildren(s)
	}
}

type BaseClassDefContext struct {
	*ClassdefContext
}

func NewBaseClassDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BaseClassDefContext {
	var p = new(BaseClassDefContext)

	p.ClassdefContext = NewEmptyClassdefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ClassdefContext))

	return p
}

func (s *BaseClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseClassDefContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *BaseClassDefContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *BaseClassDefContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
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

	return t.(IFieldContext)
}

func (s *BaseClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterBaseClassDef(s)
	}
}

func (s *BaseClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitBaseClassDef(s)
	}
}

func (s *BaseClassDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitBaseClassDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Classdef() (localctx IClassdefContext) {
	this := p
	_ = this

	localctx = NewClassdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LatteParserRULE_classdef)
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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBaseClassDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(LatteParserT__2)
		}
		{
			p.SetState(53)
			p.Match(LatteParserID)
		}
		{
			p.SetState(54)
			p.Match(LatteParserT__3)
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281475037528064) != 0 {
			{
				p.SetState(55)
				p.Field()
			}

			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(61)
			p.Match(LatteParserT__4)
		}

	case 2:
		localctx = NewDerivedClassDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(LatteParserT__2)
		}
		{
			p.SetState(63)
			p.Match(LatteParserID)
		}
		{
			p.SetState(64)
			p.Match(LatteParserT__5)
		}
		{
			p.SetState(65)
			p.Match(LatteParserID)
		}
		{
			p.SetState(66)
			p.Match(LatteParserT__3)
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281475037528064) != 0 {
			{
				p.SetState(67)
				p.Field()
			}

			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(73)
			p.Match(LatteParserT__4)
		}

	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) AllNvtype_() []INvtype_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INvtype_Context); ok {
			len++
		}
	}

	tst := make([]INvtype_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INvtype_Context); ok {
			tst[i] = t.(INvtype_Context)
			i++
		}
	}

	return tst
}

func (s *ArgContext) Nvtype_(i int) INvtype_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INvtype_Context); ok {
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

	return t.(INvtype_Context)
}

func (s *ArgContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LatteParserID)
}

func (s *ArgContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LatteParserID, i)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitArg(s)
	}
}

func (s *ArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Arg() (localctx IArgContext) {
	this := p
	_ = this

	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LatteParserRULE_arg)
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
		p.SetState(76)
		p.Nvtype_()
	}
	{
		p.SetState(77)
		p.Match(LatteParserID)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LatteParserT__6 {
		{
			p.SetState(78)
			p.Match(LatteParserT__6)
		}
		{
			p.SetState(79)
			p.Nvtype_()
		}
		{
			p.SetState(80)
			p.Match(LatteParserID)
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) CopyFrom(ctx *FieldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ClassMethodDefContext struct {
	*FieldContext
}

func NewClassMethodDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClassMethodDefContext {
	var p = new(ClassMethodDefContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *ClassMethodDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassMethodDefContext) Fundef() IFundefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFundefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFundefContext)
}

func (s *ClassMethodDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterClassMethodDef(s)
	}
}

func (s *ClassMethodDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitClassMethodDef(s)
	}
}

func (s *ClassMethodDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitClassMethodDef(s)

	default:
		return t.VisitChildren(s)
	}
}

type ClassFieldDefContext struct {
	*FieldContext
}

func NewClassFieldDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClassFieldDefContext {
	var p = new(ClassFieldDefContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *ClassFieldDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassFieldDefContext) Nvtype_() INvtype_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INvtype_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INvtype_Context)
}

func (s *ClassFieldDefContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *ClassFieldDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterClassFieldDef(s)
	}
}

func (s *ClassFieldDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitClassFieldDef(s)
	}
}

func (s *ClassFieldDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitClassFieldDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LatteParserRULE_field)

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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewClassFieldDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Nvtype_()
		}
		{
			p.SetState(88)
			p.Match(LatteParserID)
		}
		{
			p.SetState(89)
			p.Match(LatteParserT__7)
		}

	case 2:
		localctx = NewClassMethodDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Fundef()
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LatteParserRULE_block)
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
		p.SetState(94)
		p.Match(LatteParserT__3)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1548128738967826) != 0 {
		{
			p.SetState(95)
			p.Stmt()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
		p.Match(LatteParserT__4)
	}

	return localctx
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_lvalue
	return p
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) CopyFrom(ctx *LvalueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LVArrayRefContext struct {
	*LvalueContext
}

func NewLVArrayRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LVArrayRefContext {
	var p = new(LVArrayRefContext)

	p.LvalueContext = NewEmptyLvalueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LvalueContext))

	return p
}

func (s *LVArrayRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LVArrayRefContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LVArrayRefContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *LVArrayRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterLVArrayRef(s)
	}
}

func (s *LVArrayRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitLVArrayRef(s)
	}
}

func (s *LVArrayRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitLVArrayRef(s)

	default:
		return t.VisitChildren(s)
	}
}

type LVFieldContext struct {
	*LvalueContext
}

func NewLVFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LVFieldContext {
	var p = new(LVFieldContext)

	p.LvalueContext = NewEmptyLvalueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LvalueContext))

	return p
}

func (s *LVFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LVFieldContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LVFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *LVFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterLVField(s)
	}
}

func (s *LVFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitLVField(s)
	}
}

func (s *LVFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitLVField(s)

	default:
		return t.VisitChildren(s)
	}
}

type LVIdContext struct {
	*LvalueContext
}

func NewLVIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LVIdContext {
	var p = new(LVIdContext)

	p.LvalueContext = NewEmptyLvalueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LvalueContext))

	return p
}

func (s *LVIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LVIdContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *LVIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterLVId(s)
	}
}

func (s *LVIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitLVId(s)
	}
}

func (s *LVIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitLVId(s)

	default:
		return t.VisitChildren(s)
	}
}

type LVFieldArrayRefContext struct {
	*LvalueContext
}

func NewLVFieldArrayRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LVFieldArrayRefContext {
	var p = new(LVFieldArrayRefContext)

	p.LvalueContext = NewEmptyLvalueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LvalueContext))

	return p
}

func (s *LVFieldArrayRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LVFieldArrayRefContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LVFieldArrayRefContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *LVFieldArrayRefContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *LVFieldArrayRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterLVFieldArrayRef(s)
	}
}

func (s *LVFieldArrayRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitLVFieldArrayRef(s)
	}
}

func (s *LVFieldArrayRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitLVFieldArrayRef(s)

	default:
		return t.VisitChildren(s)
	}
}

type LVFieldMethodCallContext struct {
	*LvalueContext
}

func NewLVFieldMethodCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LVFieldMethodCallContext {
	var p = new(LVFieldMethodCallContext)

	p.LvalueContext = NewEmptyLvalueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LvalueContext))

	return p
}

func (s *LVFieldMethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LVFieldMethodCallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LVFieldMethodCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *LVFieldMethodCallContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *LVFieldMethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterLVFieldMethodCall(s)
	}
}

func (s *LVFieldMethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitLVFieldMethodCall(s)
	}
}

func (s *LVFieldMethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitLVFieldMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Lvalue() (localctx ILvalueContext) {
	this := p
	_ = this

	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LatteParserRULE_lvalue)

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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLVArrayRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.expr(0)
		}
		{
			p.SetState(104)
			p.Match(LatteParserT__8)
		}
		{
			p.SetState(105)
			p.expr(0)
		}
		{
			p.SetState(106)
			p.Match(LatteParserT__9)
		}

	case 2:
		localctx = NewLVFieldArrayRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.expr(0)
		}
		{
			p.SetState(109)
			p.Match(LatteParserT__10)
		}
		{
			p.SetState(110)
			p.Match(LatteParserID)
		}
		{
			p.SetState(111)
			p.Match(LatteParserT__8)
		}
		{
			p.SetState(112)
			p.expr(0)
		}
		{
			p.SetState(113)
			p.Match(LatteParserT__9)
		}

	case 3:
		localctx = NewLVFieldMethodCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.expr(0)
		}
		{
			p.SetState(116)
			p.Match(LatteParserT__10)
		}
		{
			p.SetState(117)
			p.Match(LatteParserID)
		}
		{
			p.SetState(118)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(119)
			p.expr(0)
		}
		{
			p.SetState(120)
			p.Match(LatteParserT__1)
		}

	case 4:
		localctx = NewLVFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.expr(0)
		}
		{
			p.SetState(123)
			p.Match(LatteParserT__10)
		}
		{
			p.SetState(124)
			p.Match(LatteParserID)
		}

	case 5:
		localctx = NewLVIdContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)
			p.Match(LatteParserID)
		}

	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyFrom(ctx *StmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SBlockStmtContext struct {
	*StmtContext
}

func NewSBlockStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SBlockStmtContext {
	var p = new(SBlockStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SBlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SBlockStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SBlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSBlockStmt(s)
	}
}

func (s *SBlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSBlockStmt(s)
	}
}

func (s *SBlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type SWhileContext struct {
	*StmtContext
}

func NewSWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SWhileContext {
	var p = new(SWhileContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SWhileContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SWhileContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *SWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSWhile(s)
	}
}

func (s *SWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSWhile(s)
	}
}

func (s *SWhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type SCondContext struct {
	*StmtContext
}

func NewSCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SCondContext {
	var p = new(SCondContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SCondContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SCondContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *SCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSCond(s)
	}
}

func (s *SCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSCond(s)
	}
}

func (s *SCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type SVRetContext struct {
	*StmtContext
}

func NewSVRetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SVRetContext {
	var p = new(SVRetContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SVRetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SVRetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSVRet(s)
	}
}

func (s *SVRetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSVRet(s)
	}
}

func (s *SVRetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSVRet(s)

	default:
		return t.VisitChildren(s)
	}
}

type SCondElseContext struct {
	*StmtContext
}

func NewSCondElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SCondElseContext {
	var p = new(SCondElseContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SCondElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SCondElseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SCondElseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *SCondElseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *SCondElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSCondElse(s)
	}
}

func (s *SCondElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSCondElse(s)
	}
}

func (s *SCondElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSCondElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type SAssContext struct {
	*StmtContext
}

func NewSAssContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SAssContext {
	var p = new(SAssContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SAssContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SAssContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *SAssContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SAssContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSAss(s)
	}
}

func (s *SAssContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSAss(s)
	}
}

func (s *SAssContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSAss(s)

	default:
		return t.VisitChildren(s)
	}
}

type SRetContext struct {
	*StmtContext
}

func NewSRetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SRetContext {
	var p = new(SRetContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SRetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SRetContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SRetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSRet(s)
	}
}

func (s *SRetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSRet(s)
	}
}

func (s *SRetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSRet(s)

	default:
		return t.VisitChildren(s)
	}
}

type SExpContext struct {
	*StmtContext
}

func NewSExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SExpContext {
	var p = new(SExpContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSExp(s)
	}
}

func (s *SExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSExp(s)
	}
}

func (s *SExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDeclContext struct {
	*StmtContext
}

func NewSDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDeclContext {
	var p = new(SDeclContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDeclContext) Nvtype_() INvtype_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INvtype_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INvtype_Context)
}

func (s *SDeclContext) AllItem() []IItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IItemContext); ok {
			len++
		}
	}

	tst := make([]IItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IItemContext); ok {
			tst[i] = t.(IItemContext)
			i++
		}
	}

	return tst
}

func (s *SDeclContext) Item(i int) IItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItemContext); ok {
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

	return t.(IItemContext)
}

func (s *SDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSDecl(s)
	}
}

func (s *SDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSDecl(s)
	}
}

func (s *SDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type SForContext struct {
	*StmtContext
}

func NewSForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SForContext {
	var p = new(SForContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SForContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *SForContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *SForContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SForContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *SForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSFor(s)
	}
}

func (s *SForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSFor(s)
	}
}

func (s *SForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSFor(s)

	default:
		return t.VisitChildren(s)
	}
}

type SDecrContext struct {
	*StmtContext
}

func NewSDecrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SDecrContext {
	var p = new(SDecrContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SDecrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SDecrContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *SDecrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSDecr(s)
	}
}

func (s *SDecrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSDecr(s)
	}
}

func (s *SDecrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSDecr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SIncrContext struct {
	*StmtContext
}

func NewSIncrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SIncrContext {
	var p = new(SIncrContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SIncrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SIncrContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *SIncrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSIncr(s)
	}
}

func (s *SIncrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSIncr(s)
	}
}

func (s *SIncrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSIncr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SEmptyContext struct {
	*StmtContext
}

func NewSEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SEmptyContext {
	var p = new(SEmptyContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSEmpty(s)
	}
}

func (s *SEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSEmpty(s)
	}
}

func (s *SEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LatteParserRULE_stmt)
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

	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Match(LatteParserT__7)
		}

	case 2:
		localctx = NewSBlockStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Block()
		}

	case 3:
		localctx = NewSDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Nvtype_()
		}
		{
			p.SetState(132)
			p.Item()
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LatteParserT__6 {
			{
				p.SetState(133)
				p.Match(LatteParserT__6)
			}
			{
				p.SetState(134)
				p.Item()
			}

			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(140)
			p.Match(LatteParserT__7)
		}

	case 4:
		localctx = NewSAssContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(142)
			p.Lvalue()
		}
		{
			p.SetState(143)
			p.Match(LatteParserT__11)
		}
		{
			p.SetState(144)
			p.expr(0)
		}
		{
			p.SetState(145)
			p.Match(LatteParserT__7)
		}

	case 5:
		localctx = NewSIncrContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(147)
			p.Lvalue()
		}
		{
			p.SetState(148)
			p.Match(LatteParserT__12)
		}
		{
			p.SetState(149)
			p.Match(LatteParserT__7)
		}

	case 6:
		localctx = NewSDecrContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(151)
			p.Lvalue()
		}
		{
			p.SetState(152)
			p.Match(LatteParserT__13)
		}
		{
			p.SetState(153)
			p.Match(LatteParserT__7)
		}

	case 7:
		localctx = NewSRetContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(155)
			p.Match(LatteParserT__14)
		}
		{
			p.SetState(156)
			p.expr(0)
		}
		{
			p.SetState(157)
			p.Match(LatteParserT__7)
		}

	case 8:
		localctx = NewSVRetContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(159)
			p.Match(LatteParserT__14)
		}
		{
			p.SetState(160)
			p.Match(LatteParserT__7)
		}

	case 9:
		localctx = NewSCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(161)
			p.Match(LatteParserT__15)
		}
		{
			p.SetState(162)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(163)
			p.expr(0)
		}
		{
			p.SetState(164)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(165)
			p.Stmt()
		}

	case 10:
		localctx = NewSCondElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(167)
			p.Match(LatteParserT__15)
		}
		{
			p.SetState(168)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(169)
			p.expr(0)
		}
		{
			p.SetState(170)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(171)
			p.Stmt()
		}
		{
			p.SetState(172)
			p.Match(LatteParserT__16)
		}
		{
			p.SetState(173)
			p.Stmt()
		}

	case 11:
		localctx = NewSWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(175)
			p.Match(LatteParserT__17)
		}
		{
			p.SetState(176)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(177)
			p.expr(0)
		}
		{
			p.SetState(178)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(179)
			p.Stmt()
		}

	case 12:
		localctx = NewSForContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(181)
			p.Match(LatteParserT__18)
		}
		{
			p.SetState(182)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(183)
			p.Type_()
		}
		{
			p.SetState(184)
			p.Match(LatteParserID)
		}
		{
			p.SetState(185)
			p.Match(LatteParserT__19)
		}
		{
			p.SetState(186)
			p.expr(0)
		}
		{
			p.SetState(187)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(188)
			p.Stmt()
		}

	case 13:
		localctx = NewSExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(190)
			p.expr(0)
		}
		{
			p.SetState(191)
			p.Match(LatteParserT__7)
		}

	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) CopyFrom(ctx *Type_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TVoidContext struct {
	*Type_Context
}

func NewTVoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TVoidContext {
	var p = new(TVoidContext)

	p.Type_Context = NewEmptyType_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_Context))

	return p
}

func (s *TVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTVoid(s)
	}
}

func (s *TVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTVoid(s)
	}
}

func (s *TVoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

type TNonVoidContext struct {
	*Type_Context
}

func NewTNonVoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TNonVoidContext {
	var p = new(TNonVoidContext)

	p.Type_Context = NewEmptyType_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_Context))

	return p
}

func (s *TNonVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TNonVoidContext) Nvtype_() INvtype_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INvtype_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INvtype_Context)
}

func (s *TNonVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTNonVoid(s)
	}
}

func (s *TNonVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTNonVoid(s)
	}
}

func (s *TNonVoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTNonVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LatteParserRULE_type_)

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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LatteParserT__22, LatteParserT__23, LatteParserT__24, LatteParserID:
		localctx = NewTNonVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.Nvtype_()
		}

	case LatteParserT__20:
		localctx = NewTVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
			p.Match(LatteParserT__20)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INvtype_Context is an interface to support dynamic dispatch.
type INvtype_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNvtype_Context differentiates from other interfaces.
	IsNvtype_Context()
}

type Nvtype_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNvtype_Context() *Nvtype_Context {
	var p = new(Nvtype_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_nvtype_
	return p
}

func (*Nvtype_Context) IsNvtype_Context() {}

func NewNvtype_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nvtype_Context {
	var p = new(Nvtype_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_nvtype_

	return p
}

func (s *Nvtype_Context) GetParser() antlr.Parser { return s.parser }

func (s *Nvtype_Context) CopyFrom(ctx *Nvtype_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Nvtype_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nvtype_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TSingularContext struct {
	*Nvtype_Context
}

func NewTSingularContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TSingularContext {
	var p = new(TSingularContext)

	p.Nvtype_Context = NewEmptyNvtype_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Nvtype_Context))

	return p
}

func (s *TSingularContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TSingularContext) Singular_type_() ISingular_type_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingular_type_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingular_type_Context)
}

func (s *TSingularContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTSingular(s)
	}
}

func (s *TSingularContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTSingular(s)
	}
}

func (s *TSingularContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTSingular(s)

	default:
		return t.VisitChildren(s)
	}
}

type TArrayContext struct {
	*Nvtype_Context
}

func NewTArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TArrayContext {
	var p = new(TArrayContext)

	p.Nvtype_Context = NewEmptyNvtype_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Nvtype_Context))

	return p
}

func (s *TArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TArrayContext) Singular_type_() ISingular_type_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingular_type_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingular_type_Context)
}

func (s *TArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTArray(s)
	}
}

func (s *TArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTArray(s)
	}
}

func (s *TArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Nvtype_() (localctx INvtype_Context) {
	this := p
	_ = this

	localctx = NewNvtype_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LatteParserRULE_nvtype_)

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

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Singular_type_()
		}
		{
			p.SetState(200)
			p.Match(LatteParserT__21)
		}

	case 2:
		localctx = NewTSingularContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Singular_type_()
		}

	}

	return localctx
}

// ISingular_type_Context is an interface to support dynamic dispatch.
type ISingular_type_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingular_type_Context differentiates from other interfaces.
	IsSingular_type_Context()
}

type Singular_type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingular_type_Context() *Singular_type_Context {
	var p = new(Singular_type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_singular_type_
	return p
}

func (*Singular_type_Context) IsSingular_type_Context() {}

func NewSingular_type_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Singular_type_Context {
	var p = new(Singular_type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_singular_type_

	return p
}

func (s *Singular_type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Singular_type_Context) CopyFrom(ctx *Singular_type_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Singular_type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Singular_type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TBoolContext struct {
	*Singular_type_Context
}

func NewTBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TBoolContext {
	var p = new(TBoolContext)

	p.Singular_type_Context = NewEmptySingular_type_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Singular_type_Context))

	return p
}

func (s *TBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTBool(s)
	}
}

func (s *TBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTBool(s)
	}
}

func (s *TBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type TStrContext struct {
	*Singular_type_Context
}

func NewTStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TStrContext {
	var p = new(TStrContext)

	p.Singular_type_Context = NewEmptySingular_type_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Singular_type_Context))

	return p
}

func (s *TStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTStr(s)
	}
}

func (s *TStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTStr(s)
	}
}

func (s *TStrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTStr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TIntContext struct {
	*Singular_type_Context
}

func NewTIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TIntContext {
	var p = new(TIntContext)

	p.Singular_type_Context = NewEmptySingular_type_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Singular_type_Context))

	return p
}

func (s *TIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTInt(s)
	}
}

func (s *TIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTInt(s)
	}
}

func (s *TIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type TClassContext struct {
	*Singular_type_Context
}

func NewTClassContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TClassContext {
	var p = new(TClassContext)

	p.Singular_type_Context = NewEmptySingular_type_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Singular_type_Context))

	return p
}

func (s *TClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TClassContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *TClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTClass(s)
	}
}

func (s *TClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTClass(s)
	}
}

func (s *TClassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTClass(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Singular_type_() (localctx ISingular_type_Context) {
	this := p
	_ = this

	localctx = NewSingular_type_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LatteParserRULE_singular_type_)

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

	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LatteParserID:
		localctx = NewTClassContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(LatteParserID)
		}

	case LatteParserT__22:
		localctx = NewTIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(LatteParserT__22)
		}

	case LatteParserT__23:
		localctx = NewTStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(207)
			p.Match(LatteParserT__23)
		}

	case LatteParserT__24:
		localctx = NewTBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(208)
			p.Match(LatteParserT__24)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IItemContext is an interface to support dynamic dispatch.
type IItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemContext differentiates from other interfaces.
	IsItemContext()
}

type ItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemContext() *ItemContext {
	var p = new(ItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_item
	return p
}

func (*ItemContext) IsItemContext() {}

func NewItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemContext {
	var p = new(ItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_item

	return p
}

func (s *ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *ItemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterItem(s)
	}
}

func (s *ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitItem(s)
	}
}

func (s *ItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Item() (localctx IItemContext) {
	this := p
	_ = this

	localctx = NewItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LatteParserRULE_item)

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

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(LatteParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(LatteParserID)
		}
		{
			p.SetState(213)
			p.Match(LatteParserT__11)
		}
		{
			p.SetState(214)
			p.expr(0)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EIdContext struct {
	*ExprContext
}

func NewEIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EIdContext {
	var p = new(EIdContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIdContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *EIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEId(s)
	}
}

func (s *EIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEId(s)
	}
}

func (s *EIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEId(s)

	default:
		return t.VisitChildren(s)
	}
}

type ESelfContext struct {
	*ExprContext
}

func NewESelfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ESelfContext {
	var p = new(ESelfContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ESelfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ESelfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterESelf(s)
	}
}

func (s *ESelfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitESelf(s)
	}
}

func (s *ESelfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitESelf(s)

	default:
		return t.VisitChildren(s)
	}
}

type EFunCallContext struct {
	*ExprContext
}

func NewEFunCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EFunCallContext {
	var p = new(EFunCallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EFunCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EFunCallContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *EFunCallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EFunCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EFunCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEFunCall(s)
	}
}

func (s *EFunCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEFunCall(s)
	}
}

func (s *EFunCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEFunCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENewArrayContext struct {
	*ExprContext
}

func NewENewArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENewArrayContext {
	var p = new(ENewArrayContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ENewArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENewArrayContext) Singular_type_() ISingular_type_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingular_type_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingular_type_Context)
}

func (s *ENewArrayContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ENewArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterENewArray(s)
	}
}

func (s *ENewArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitENewArray(s)
	}
}

func (s *ENewArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitENewArray(s)

	default:
		return t.VisitChildren(s)
	}
}

type EArrayRefContext struct {
	*ExprContext
}

func NewEArrayRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EArrayRefContext {
	var p = new(EArrayRefContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EArrayRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EArrayRefContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EArrayRefContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EArrayRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEArrayRef(s)
	}
}

func (s *EArrayRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEArrayRef(s)
	}
}

func (s *EArrayRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEArrayRef(s)

	default:
		return t.VisitChildren(s)
	}
}

type ERelOpContext struct {
	*ExprContext
}

func NewERelOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ERelOpContext {
	var p = new(ERelOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ERelOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ERelOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ERelOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ERelOpContext) RelOp() IRelOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelOpContext)
}

func (s *ERelOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterERelOp(s)
	}
}

func (s *ERelOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitERelOp(s)
	}
}

func (s *ERelOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitERelOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ETrueContext struct {
	*ExprContext
}

func NewETrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ETrueContext {
	var p = new(ETrueContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ETrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ETrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterETrue(s)
	}
}

func (s *ETrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitETrue(s)
	}
}

func (s *ETrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitETrue(s)

	default:
		return t.VisitChildren(s)
	}
}

type EOrContext struct {
	*ExprContext
}

func NewEOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EOrContext {
	var p = new(EOrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EOrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EOrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEOr(s)
	}
}

func (s *EOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEOr(s)
	}
}

func (s *EOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EIntContext struct {
	*ExprContext
}

func NewEIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EIntContext {
	var p = new(EIntContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIntContext) INT() antlr.TerminalNode {
	return s.GetToken(LatteParserINT, 0)
}

func (s *EIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEInt(s)
	}
}

func (s *EIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEInt(s)
	}
}

func (s *EIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type EStrContext struct {
	*ExprContext
}

func NewEStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EStrContext {
	var p = new(EStrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EStrContext) STR() antlr.TerminalNode {
	return s.GetToken(LatteParserSTR, 0)
}

func (s *EStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEStr(s)
	}
}

func (s *EStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEStr(s)
	}
}

func (s *EStrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEStr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EFieldArrayAccessContext struct {
	*ExprContext
}

func NewEFieldArrayAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EFieldArrayAccessContext {
	var p = new(EFieldArrayAccessContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EFieldArrayAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EFieldArrayAccessContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EFieldArrayAccessContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EFieldArrayAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *EFieldArrayAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEFieldArrayAccess(s)
	}
}

func (s *EFieldArrayAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEFieldArrayAccess(s)
	}
}

func (s *EFieldArrayAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEFieldArrayAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENotOpContext struct {
	*ExprContext
}

func NewENotOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENotOpContext {
	var p = new(ENotOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ENotOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENotOpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ENotOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterENotOp(s)
	}
}

func (s *ENotOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitENotOp(s)
	}
}

func (s *ENotOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitENotOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type EMulOpContext struct {
	*ExprContext
}

func NewEMulOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EMulOpContext {
	var p = new(EMulOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EMulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EMulOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EMulOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EMulOpContext) MulOp() IMulOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulOpContext)
}

func (s *EMulOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEMulOp(s)
	}
}

func (s *EMulOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEMulOp(s)
	}
}

func (s *EMulOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEMulOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAndContext struct {
	*ExprContext
}

func NewEAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAndContext {
	var p = new(EAndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EAndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEAnd(s)
	}
}

func (s *EAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEAnd(s)
	}
}

func (s *EAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type EMethodCallContext struct {
	*ExprContext
}

func NewEMethodCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EMethodCallContext {
	var p = new(EMethodCallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EMethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EMethodCallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EMethodCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EMethodCallContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *EMethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEMethodCall(s)
	}
}

func (s *EMethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEMethodCall(s)
	}
}

func (s *EMethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENegOpContext struct {
	*ExprContext
}

func NewENegOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENegOpContext {
	var p = new(ENegOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ENegOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENegOpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ENegOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterENegOp(s)
	}
}

func (s *ENegOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitENegOp(s)
	}
}

func (s *ENegOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitENegOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type EParenContext struct {
	*ExprContext
}

func NewEParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EParenContext {
	var p = new(EParenContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EParenContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEParen(s)
	}
}

func (s *EParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEParen(s)
	}
}

func (s *EParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEParen(s)

	default:
		return t.VisitChildren(s)
	}
}

type EFalseContext struct {
	*ExprContext
}

func NewEFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EFalseContext {
	var p = new(EFalseContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEFalse(s)
	}
}

func (s *EFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEFalse(s)
	}
}

func (s *EFalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENewContext struct {
	*ExprContext
}

func NewENewContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENewContext {
	var p = new(ENewContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ENewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENewContext) Singular_type_() ISingular_type_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingular_type_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingular_type_Context)
}

func (s *ENewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterENew(s)
	}
}

func (s *ENewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitENew(s)
	}
}

func (s *ENewContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitENew(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAddOpContext struct {
	*ExprContext
}

func NewEAddOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAddOpContext {
	var p = new(EAddOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EAddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAddOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EAddOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *EAddOpContext) AddOp() IAddOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddOpContext)
}

func (s *EAddOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEAddOp(s)
	}
}

func (s *EAddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEAddOp(s)
	}
}

func (s *EAddOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEAddOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENullContext struct {
	*ExprContext
}

func NewENullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENullContext {
	var p = new(ENullContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ENullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENullContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *ENullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterENull(s)
	}
}

func (s *ENullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitENull(s)
	}
}

func (s *ENullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitENull(s)

	default:
		return t.VisitChildren(s)
	}
}

type EFieldAccessContext struct {
	*ExprContext
}

func NewEFieldAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EFieldAccessContext {
	var p = new(EFieldAccessContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EFieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EFieldAccessContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EFieldAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *EFieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEFieldAccess(s)
	}
}

func (s *EFieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEFieldAccess(s)
	}
}

func (s *EFieldAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEFieldAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *LatteParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, LatteParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewENegOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(218)
			p.Match(LatteParserT__25)
		}
		{
			p.SetState(219)
			p.expr(18)
		}

	case 2:
		localctx = NewENotOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(220)
			p.Match(LatteParserT__26)
		}
		{
			p.SetState(221)
			p.expr(17)
		}

	case 3:
		localctx = NewENewArrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(222)
			p.Match(LatteParserT__29)
		}
		{
			p.SetState(223)
			p.Singular_type_()
		}
		{
			p.SetState(224)
			p.Match(LatteParserT__8)
		}
		{
			p.SetState(225)
			p.expr(0)
		}
		{
			p.SetState(226)
			p.Match(LatteParserT__9)
		}

	case 4:
		localctx = NewENewContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(228)
			p.Match(LatteParserT__29)
		}
		{
			p.SetState(229)
			p.Singular_type_()
		}

	case 5:
		localctx = NewESelfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(230)
			p.Match(LatteParserT__30)
		}

	case 6:
		localctx = NewEIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(231)
			p.Match(LatteParserID)
		}

	case 7:
		localctx = NewEIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(232)
			p.Match(LatteParserINT)
		}

	case 8:
		localctx = NewETrueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(233)
			p.Match(LatteParserT__31)
		}

	case 9:
		localctx = NewEFalseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(234)
			p.Match(LatteParserT__32)
		}

	case 10:
		localctx = NewEFunCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(235)
			p.Match(LatteParserID)
		}
		{
			p.SetState(236)
			p.Match(LatteParserT__0)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1548128679362562) != 0 {
			{
				p.SetState(237)
				p.expr(0)
			}
			p.SetState(242)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == LatteParserT__6 {
				{
					p.SetState(238)
					p.Match(LatteParserT__6)
				}
				{
					p.SetState(239)
					p.expr(0)
				}

				p.SetState(244)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(247)
			p.Match(LatteParserT__1)
		}

	case 11:
		localctx = NewEStrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(248)
			p.Match(LatteParserSTR)
		}

	case 12:
		localctx = NewENullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(249)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(250)
			p.Match(LatteParserID)
		}
		{
			p.SetState(251)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(252)
			p.Match(LatteParserT__33)
		}

	case 13:
		localctx = NewEParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(253)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(254)
			p.expr(0)
		}
		{
			p.SetState(255)
			p.Match(LatteParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(307)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEMulOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(259)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(260)
					p.MulOp()
				}
				{
					p.SetState(261)
					p.expr(17)
				}

			case 2:
				localctx = NewEAddOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(264)
					p.AddOp()
				}
				{
					p.SetState(265)
					p.expr(16)
				}

			case 3:
				localctx = NewERelOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(267)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(268)
					p.RelOp()
				}
				{
					p.SetState(269)
					p.expr(15)
				}

			case 4:
				localctx = NewEAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(272)
					p.Match(LatteParserT__27)
				}
				{
					p.SetState(273)
					p.expr(13)
				}

			case 5:
				localctx = NewEOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(275)
					p.Match(LatteParserT__28)
				}
				{
					p.SetState(276)
					p.expr(12)
				}

			case 6:
				localctx = NewEMethodCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(278)
					p.Match(LatteParserT__10)
				}
				{
					p.SetState(279)
					p.Match(LatteParserID)
				}
				{
					p.SetState(280)
					p.Match(LatteParserT__0)
				}
				p.SetState(289)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1548128679362562) != 0 {
					{
						p.SetState(281)
						p.expr(0)
					}
					p.SetState(286)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LatteParserT__6 {
						{
							p.SetState(282)
							p.Match(LatteParserT__6)
						}
						{
							p.SetState(283)
							p.expr(0)
						}

						p.SetState(288)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(291)
					p.Match(LatteParserT__1)
				}

			case 7:
				localctx = NewEFieldArrayAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(293)
					p.Match(LatteParserT__10)
				}
				{
					p.SetState(294)
					p.Match(LatteParserID)
				}
				{
					p.SetState(295)
					p.Match(LatteParserT__8)
				}
				{
					p.SetState(296)
					p.expr(0)
				}
				{
					p.SetState(297)
					p.Match(LatteParserT__9)
				}

			case 8:
				localctx = NewEFieldAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(299)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(300)
					p.Match(LatteParserT__10)
				}
				{
					p.SetState(301)
					p.Match(LatteParserID)
				}

			case 9:
				localctx = NewEArrayRefContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(303)
					p.Match(LatteParserT__8)
				}
				{
					p.SetState(304)
					p.expr(0)
				}
				{
					p.SetState(305)
					p.Match(LatteParserT__9)
				}

			}

		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IAddOpContext is an interface to support dynamic dispatch.
type IAddOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddOpContext differentiates from other interfaces.
	IsAddOpContext()
}

type AddOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOpContext() *AddOpContext {
	var p = new(AddOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_addOp
	return p
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }
func (s *AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterAddOp(s)
	}
}

func (s *AddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitAddOp(s)
	}
}

func (s *AddOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitAddOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) AddOp() (localctx IAddOpContext) {
	this := p
	_ = this

	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LatteParserRULE_addOp)
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
		p.SetState(312)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LatteParserT__25 || _la == LatteParserT__34) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMulOpContext is an interface to support dynamic dispatch.
type IMulOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulOpContext differentiates from other interfaces.
	IsMulOpContext()
}

type MulOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulOpContext() *MulOpContext {
	var p = new(MulOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_mulOp
	return p
}

func (*MulOpContext) IsMulOpContext() {}

func NewMulOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulOpContext {
	var p = new(MulOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_mulOp

	return p
}

func (s *MulOpContext) GetParser() antlr.Parser { return s.parser }
func (s *MulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterMulOp(s)
	}
}

func (s *MulOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitMulOp(s)
	}
}

func (s *MulOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitMulOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) MulOp() (localctx IMulOpContext) {
	this := p
	_ = this

	localctx = NewMulOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LatteParserRULE_mulOp)
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
		p.SetState(314)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&481036337152) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRelOpContext is an interface to support dynamic dispatch.
type IRelOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelOpContext differentiates from other interfaces.
	IsRelOpContext()
}

type RelOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelOpContext() *RelOpContext {
	var p = new(RelOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_relOp
	return p
}

func (*RelOpContext) IsRelOpContext() {}

func NewRelOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelOpContext {
	var p = new(RelOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_relOp

	return p
}

func (s *RelOpContext) GetParser() antlr.Parser { return s.parser }
func (s *RelOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterRelOp(s)
	}
}

func (s *RelOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitRelOp(s)
	}
}

func (s *RelOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitRelOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) RelOp() (localctx IRelOpContext) {
	this := p
	_ = this

	localctx = NewRelOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LatteParserRULE_relOp)
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
		p.SetState(316)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34634616274944) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *LatteParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LatteParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 19)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
