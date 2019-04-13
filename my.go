
//line my.go.y:2
package main
import __yyfmt__ "fmt"
//line my.go.y:2
		const IF = 57346
const THEN = 57347
const ELSE = 57348
const ELIF = 57349
const FI = 57350
const CASE = 57351
const ESAC = 57352
const FOR = 57353
const SELECT = 57354
const WHILE = 57355
const UNTIL = 57356
const DO = 57357
const DONE = 57358
const FUNCTION = 57359
const COPROC = 57360
const COND_START = 57361
const COND_END = 57362
const COND_ERROR = 57363
const IN = 57364
const BANG = 57365
const TIME = 57366
const TIMEOPT = 57367
const TIMEIGN = 57368
const WORD = 57369
const ASSIGNMENT_WORD = 57370
const REDIR_WORD = 57371
const NUMBER = 57372
const ARITH_CMD = 57373
const ARITH_FOR_EXPRS = 57374
const COND_CMD = 57375
const AND_AND = 57376
const OR_OR = 57377
const GREATER_GREATER = 57378
const LESS_LESS = 57379
const LESS_AND = 57380
const LESS_LESS_LESS = 57381
const GREATER_AND = 57382
const SEMI_SEMI = 57383
const SEMI_AND = 57384
const SEMI_SEMI_AND = 57385
const LESS_LESS_MINUS = 57386
const AND_GREATER = 57387
const AND_GREATER_GREATER = 57388
const LESS_GREATER = 57389
const GREATER_BAR = 57390
const BAR_AND = 57391
const yacc_EOF = 57392

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IF",
	"THEN",
	"ELSE",
	"ELIF",
	"FI",
	"CASE",
	"ESAC",
	"FOR",
	"SELECT",
	"WHILE",
	"UNTIL",
	"DO",
	"DONE",
	"FUNCTION",
	"COPROC",
	"COND_START",
	"COND_END",
	"COND_ERROR",
	"IN",
	"BANG",
	"TIME",
	"TIMEOPT",
	"TIMEIGN",
	"WORD",
	"ASSIGNMENT_WORD",
	"REDIR_WORD",
	"NUMBER",
	"ARITH_CMD",
	"ARITH_FOR_EXPRS",
	"COND_CMD",
	"AND_AND",
	"OR_OR",
	"GREATER_GREATER",
	"LESS_LESS",
	"LESS_AND",
	"LESS_LESS_LESS",
	"GREATER_AND",
	"SEMI_SEMI",
	"SEMI_AND",
	"SEMI_SEMI_AND",
	"LESS_LESS_MINUS",
	"AND_GREATER",
	"AND_GREATER_GREATER",
	"LESS_GREATER",
	"GREATER_BAR",
	"BAR_AND",
	"'&'",
	"';'",
	"'\\n'",
	"yacc_EOF",
	"'|'",
	"'>'",
	"'<'",
	"'-'",
	"'{'",
	"'}'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line my.go.y:987


//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 711

var yyAct = [...]int{

	79, 80, 256, 274, 234, 246, 77, 144, 7, 132,
	6, 206, 8, 33, 304, 205, 17, 67, 72, 304,
	14, 327, 82, 160, 146, 306, 303, 147, 78, 81,
	75, 83, 348, 346, 217, 142, 273, 68, 92, 93,
	94, 219, 276, 342, 336, 314, 302, 300, 73, 37,
	285, 265, 85, 276, 35, 242, 34, 36, 20, 21,
	161, 66, 158, 328, 41, 320, 65, 143, 216, 7,
	7, 143, 133, 154, 37, 275, 40, 218, 143, 35,
	209, 34, 36, 20, 21, 152, 275, 211, 311, 41,
	95, 139, 149, 134, 135, 136, 137, 143, 143, 78,
	143, 40, 143, 39, 75, 38, 329, 150, 321, 70,
	69, 71, 266, 326, 148, 297, 155, 143, 151, 276,
	156, 157, 252, 210, 239, 143, 184, 153, 39, 237,
	38, 312, 189, 59, 60, 188, 70, 69, 71, 57,
	58, 197, 7, 7, 193, 194, 203, 13, 186, 195,
	196, 185, 143, 201, 202, 214, 215, 208, 298, 143,
	221, 143, 190, 139, 78, 253, 143, 240, 175, 207,
	138, 174, 238, 307, 308, 309, 204, 172, 187, 87,
	171, 183, 212, 213, 280, 281, 282, 128, 220, 182,
	127, 181, 180, 110, 113, 116, 115, 117, 176, 7,
	7, 114, 133, 125, 112, 111, 124, 173, 63, 64,
	232, 233, 108, 109, 231, 179, 229, 129, 88, 247,
	249, 78, 139, 89, 250, 207, 178, 177, 223, 224,
	225, 226, 227, 126, 170, 230, 169, 168, 269, 270,
	271, 272, 263, 167, 166, 139, 243, 248, 248, 235,
	278, 207, 201, 202, 288, 289, 290, 251, 292, 258,
	259, 260, 261, 262, 165, 268, 236, 164, 199, 200,
	198, 163, 267, 131, 277, 130, 123, 63, 64, 305,
	122, 287, 248, 248, 284, 121, 120, 119, 286, 118,
	70, 69, 71, 61, 62, 97, 96, 296, 322, 323,
	91, 90, 84, 74, 324, 162, 347, 100, 103, 106,
	105, 107, 330, 331, 345, 104, 310, 332, 102, 101,
	341, 333, 334, 317, 318, 319, 98, 99, 337, 335,
	339, 340, 313, 325, 301, 344, 299, 283, 264, 241,
	228, 222, 145, 140, 349, 279, 255, 257, 254, 343,
	257, 316, 315, 259, 260, 258, 37, 338, 248, 248,
	159, 35, 291, 34, 36, 20, 21, 56, 10, 30,
	31, 41, 245, 244, 24, 9, 12, 23, 15, 29,
	32, 45, 44, 40, 16, 28, 27, 26, 46, 49,
	52, 51, 53, 25, 19, 22, 50, 54, 55, 48,
	47, 18, 2, 70, 69, 71, 141, 42, 43, 4,
	39, 37, 38, 11, 1, 0, 35, 0, 34, 36,
	20, 21, 0, 0, 30, 31, 41, 0, 0, 0,
	9, 12, 0, 0, 29, 32, 45, 44, 40, 0,
	0, 0, 0, 46, 49, 52, 51, 53, 0, 0,
	0, 50, 54, 55, 48, 47, 0, 0, 0, 3,
	5, 37, 42, 43, 0, 39, 35, 38, 34, 36,
	20, 21, 0, 0, 30, 31, 41, 0, 0, 0,
	9, 12, 0, 0, 29, 32, 45, 44, 40, 0,
	0, 0, 0, 46, 49, 52, 51, 53, 0, 0,
	0, 50, 54, 55, 48, 47, 0, 0, 0, 143,
	0, 37, 42, 43, 0, 39, 35, 38, 34, 36,
	20, 21, 0, 0, 30, 31, 41, 0, 0, 0,
	9, 12, 0, 0, 29, 32, 45, 44, 40, 0,
	0, 0, 0, 46, 49, 52, 51, 53, 0, 0,
	0, 50, 54, 55, 48, 47, 0, 0, 0, 0,
	0, 37, 42, 43, 0, 39, 35, 38, 34, 36,
	20, 21, 0, 0, 30, 31, 41, 0, 0, 0,
	0, 0, 0, 0, 29, 32, 45, 44, 40, 0,
	0, 0, 0, 46, 49, 52, 51, 53, 0, 0,
	0, 50, 54, 55, 48, 47, 0, 0, 0, 143,
	0, 37, 42, 43, 0, 39, 35, 38, 34, 36,
	20, 21, 0, 0, 0, 0, 41, 0, 0, 0,
	0, 0, 0, 0, 86, 32, 45, 44, 40, 0,
	0, 0, 0, 46, 49, 52, 51, 53, 0, 0,
	0, 50, 54, 55, 48, 47, 0, 0, 0, 0,
	0, 0, 42, 43, 0, 39, 0, 38, 76, 32,
	45, 44, 0, 0, 0, 0, 0, 46, 49, 52,
	51, 53, 0, 45, 44, 50, 54, 55, 48, 47,
	46, 49, 52, 51, 53, 0, 42, 43, 50, 54,
	55, 48, 47, 0, 0, 0, 0, 0, 0, 42,
	43,
}
var yyPact = [...]int{

	407, -1000, 87, -1000, 81, -1000, 243, -1000, 12, 352,
	352, -1000, 278, 641, 654, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -29,
	275, 607, -1000, -1000, 191, 274, 273, -1000, -1000, -1000,
	-1000, 57, 269, 268, 271, 157, 262, 260, 259, 258,
	253, 249, 176, 160, 248, 246, -1000, -1000, -1000, -1000,
	-1000, 507, 507, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 144, -1000, -1000, 654, -1000, 328,
	-1000, 457, 327, -37, -33, 654, 70, 641, 34, 58,
	-1000, 11, 355, -38, 1, 285, -1000, -1000, 244, 240,
	237, 217, 216, 210, 209, 207, 150, 141, 200, 199,
	188, 165, 164, 162, 154, 99, 121, 105, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 174, 174, 457, 457, 557, 557, -1000, -1000,
	-1000, -1000, 218, -1000, -1000, -1000, -1000, -46, 45, 654,
	654, 65, -1000, -1000, -1000, -1000, 46, 19, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 507, 507, -1000, -1000, 12, 12, 325, -1000, -1000,
	-1000, -1000, -1000, 324, 45, -1000, -1000, 654, 654, -1000,
	-1000, 239, 114, 109, 323, -4, -1000, -1000, -1000, 239,
	107, 340, -1000, 457, 457, 457, 457, 457, -1000, -1000,
	45, 654, 322, -8, 85, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 26, -1000, 335, 143, 321, 457, -9,
	85, -1000, -1000, -1000, -1000, -1000, 354, -1000, 119, 119,
	119, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 100, 320,
	-12, 318, -13, -1000, -35, 92, -1000, 15, 132, -1000,
	-1000, -1000, -1000, -1000, 218, -1000, -1000, 73, 316, -14,
	344, -1000, 346, -1000, -1000, -1000, 50, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 86, -40, -1000, -1000, -1000, -1000,
	48, -1000, -1000, -1000, -1000, -1000, -1000, 457, 457, 457,
	-1000, -1000, 313, -15, -1000, 457, -1000, -1000, -1000, -1000,
	304, -16, 343, 298, -26, -1000, -1000, -1000, 457, 290,
	-27, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 414, 413, 12, 7, 1, 406, 35, 0, 402,
	9, 147, 20, 401, 395, 394, 393, 387, 386, 385,
	384, 378, 11, 377, 2, 374, 13, 6, 16, 4,
	3, 5, 373, 372, 368, 37, 367, 29,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 1, 1, 29, 29, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 28, 28, 28, 27, 27, 11, 11, 2,
	2, 2, 2, 2, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 13, 13, 13, 13, 13,
	13, 13, 13, 19, 19, 19, 19, 14, 14, 14,
	14, 14, 14, 14, 14, 15, 15, 15, 21, 21,
	21, 22, 22, 25, 20, 20, 20, 20, 20, 23,
	23, 23, 16, 17, 18, 24, 24, 24, 33, 33,
	31, 31, 31, 31, 32, 32, 32, 32, 32, 32,
	30, 30, 5, 8, 8, 6, 6, 6, 7, 7,
	7, 7, 7, 7, 36, 36, 35, 35, 35, 37,
	37, 9, 9, 9, 10, 10, 10, 10, 10, 4,
	4, 4, 4, 4, 3, 3, 3, 34, 34, 34,
}
var yyR2 = [...]int{

	0, 2, 1, 2, 2, 1, 1, 2, 2, 2,
	3, 3, 3, 3, 2, 3, 3, 2, 3, 3,
	2, 3, 3, 2, 3, 3, 2, 3, 3, 2,
	3, 3, 2, 3, 3, 2, 3, 3, 2, 3,
	3, 2, 3, 3, 2, 3, 3, 2, 3, 3,
	2, 2, 1, 1, 1, 1, 2, 1, 2, 1,
	1, 2, 1, 1, 1, 1, 5, 5, 1, 1,
	1, 1, 1, 1, 1, 6, 6, 7, 7, 10,
	10, 9, 9, 7, 7, 5, 5, 6, 6, 7,
	7, 10, 10, 9, 9, 6, 7, 6, 5, 6,
	4, 1, 2, 3, 2, 3, 3, 4, 2, 5,
	7, 6, 3, 1, 3, 4, 6, 5, 1, 2,
	4, 4, 5, 5, 2, 3, 2, 3, 2, 3,
	1, 3, 2, 1, 2, 3, 3, 3, 4, 4,
	4, 4, 4, 1, 1, 1, 1, 1, 1, 0,
	2, 1, 2, 2, 4, 4, 3, 3, 1, 1,
	2, 2, 2, 2, 4, 4, 1, 1, 2, 3,
}
var yyChk = [...]int{

	-1000, -1, -9, 52, 2, 53, -10, -4, -3, 23,
	-34, -2, 24, -11, -12, -21, -20, -28, -13, -15,
	13, 14, -14, -23, -25, -16, -17, -18, -19, 27,
	17, 18, 28, -26, 11, 9, 12, 4, 60, 58,
	31, 19, 55, 56, 30, 29, 36, 48, 47, 37,
	44, 39, 38, 40, 45, 46, -36, 52, 53, 52,
	53, 50, 51, 34, 35, 54, 49, -4, -35, 52,
	51, 53, -4, -35, 25, -28, 27, -27, -26, -8,
	-5, -37, -8, 60, 27, -12, 27, -11, 27, 32,
	27, 27, -8, -8, -8, 33, 27, 27, 55, 56,
	36, 48, 47, 37, 44, 39, 38, 40, 55, 56,
	36, 48, 47, 37, 44, 39, 38, 40, 27, 27,
	27, 27, 27, 27, 30, 27, 57, 30, 27, 57,
	27, 27, -10, -10, -37, -37, -37, -37, 26, -26,
	15, -6, -7, 52, -4, 15, 61, 60, -37, -27,
	-12, -37, 51, -35, 15, 58, -37, -37, 51, 5,
	61, 59, 20, 27, 27, 27, 27, 27, 27, 27,
	27, 30, 27, 57, 30, 27, 57, 27, 27, 27,
	27, 27, 27, 27, 27, 30, 27, 57, 30, 27,
	57, 50, 51, -10, -10, -3, -3, -8, 52, 50,
	51, 34, 35, -8, -37, 61, -22, -12, -27, 15,
	58, 22, -37, -37, -8, -8, 22, 15, 58, 22,
	-37, -8, 16, -37, -37, -37, -37, -37, 16, -22,
	-37, -27, -8, -8, -29, -35, 27, 15, 58, 15,
	58, 16, 59, -37, -32, -33, -31, -5, -37, -5,
	-29, -35, 15, 58, 8, 6, -24, 7, -7, -7,
	-7, -7, -7, -22, 16, 59, 27, -35, -37, -8,
	-8, -8, -8, 10, -30, 60, 27, -37, -31, 10,
	41, 42, 43, 16, -7, 59, -35, -37, -5, -5,
	-8, 8, -8, 50, 51, 52, -37, 15, 58, 16,
	59, 16, 59, 61, 54, -30, 10, 41, 42, 43,
	-37, 15, 58, 16, 59, 8, 5, -37, -37, -37,
	15, 58, -8, -8, -8, -37, 27, 61, 15, 58,
	-8, -8, -8, -8, -8, 16, 59, -8, -37, -5,
	-5, 16, 59, 6, -24, 16, 59, 16, 59, -8,
}
var yyDef = [...]int{

	0, -2, 0, 2, 0, 5, 151, 158, 159, 0,
	0, 166, 167, 59, 60, 62, 63, 57, 64, 65,
	149, 149, 68, 69, 70, 71, 72, 73, 74, 52,
	0, 0, 53, 54, 0, 0, 0, 149, 149, 149,
	113, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 1, 144, 145, 3,
	4, 152, 153, 149, 149, 149, 149, 160, 163, 146,
	147, 148, 161, 162, 168, 58, 52, 61, 55, 0,
	133, 0, 0, 0, 149, 104, 52, 108, 149, 0,
	149, 149, 0, 0, 0, 0, 8, 9, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 14, 17,
	20, 23, 26, 29, 32, 38, 47, 35, 41, 44,
	50, 51, 156, 157, 0, 0, 0, 0, 169, 56,
	149, 132, 134, 150, 143, 149, 149, 0, 0, 105,
	106, 0, 149, 149, 149, 149, 0, 0, 149, 149,
	103, 112, 114, 10, 11, 15, 18, 21, 24, 27,
	30, 33, 39, 48, 36, 42, 45, 12, 13, 16,
	19, 22, 25, 28, 31, 34, 40, 49, 37, 43,
	46, 0, 0, 154, 155, 164, 165, 0, 149, 149,
	149, 149, 149, 0, 0, 149, 100, 101, 107, 149,
	149, 0, 0, 0, 0, 0, 149, 149, 149, 0,
	0, 0, 66, 135, 136, 137, 0, 0, 67, 98,
	0, 102, 0, 0, 0, 149, 6, 149, 149, 149,
	149, 85, 86, 0, 149, 0, 118, 0, 0, 0,
	0, 149, 149, 149, 109, 149, 0, 149, 142, 140,
	141, 138, 139, 99, 75, 76, 7, 149, 0, 0,
	0, 0, 0, 95, 0, 0, 130, 0, 119, 97,
	124, 126, 128, 87, 0, 88, 149, 0, 0, 0,
	0, 111, 0, 149, 149, 149, 0, 149, 149, 77,
	78, 83, 84, 149, 0, 0, 96, 125, 127, 129,
	0, 149, 149, 89, 90, 110, 149, 0, 0, 0,
	149, 149, 0, 0, 120, 121, 131, 149, 149, 149,
	0, 0, 115, 0, 0, 81, 82, 122, 123, 0,
	0, 93, 94, 149, 117, 79, 80, 91, 92, 116,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	52, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 50, 3,
	60, 61, 3, 3, 3, 57, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 51,
	56, 3, 55, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 58, 54, 59,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 53,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{
}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:49
		{
				  /* Case of regular command.  Discard the error
			     safety net,and return the command just parsed. */
				  global_command = yyDollar[1].command;
				  eof_encountered = 0;
				  /* discard_parser_constructs (0); */
				  if (parser_state & PST_CMDSUBST)
				    parser_state |= PST_EOFTOKEN;
				  YYACCEPT;
				}
	case 2:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:60
		{
				  /* Case of regular command, but not a very
			     interesting one.  Return a NULL command. */
				  global_command = (COMMAND *)NULL;
				  if (parser_state & PST_CMDSUBST)
				    parser_state |= PST_EOFTOKEN;
				  YYACCEPT;
				}
	case 3:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:69
		{
				  /* Error during parsing.  Return NULL command. */
				  global_command = (COMMAND *)NULL;
				  eof_encountered = 0;
				  /* discard_parser_constructs (1); */
				  if (interactive && parse_and_execute_level == 0)
				    {
				      YYACCEPT;
				    }
				  else
				    {
				      YYABORT;
				    }
				}
	case 4:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:84
		{
				  /* EOF after an error.  Do ignoreeof or not.  Really
			     only interesting in non-interactive shells */
				  global_command = (COMMAND *)NULL;
				  if (last_command_exit_value == 0)
				    last_command_exit_value = EX_BADUSAGE;	/* force error return */
				  handle_eof_input_unit ();
				  if (interactive && parse_and_execute_level == 0)
				    {
				      YYACCEPT;
				    }
				  else
				    {
				      YYABORT;
				    }
				}
	case 5:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:101
		{
				  /* Case of EOF seen by itself.  Do ignoreeof or
			     not. */
				  global_command = (COMMAND *)NULL;
				  handle_eof_input_unit ();
				  YYACCEPT;
				}
	case 6:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:111
		{ yyVAL.word_list = make_word_list (yyDollar[1].word, (WORD_LIST *)NULL); }
	case 7:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:113
		{ yyVAL.word_list = make_word_list (yyDollar[2].word, yyDollar[1].word_list); }
	case 8:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:117
		{
				  source.dest = 1;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_output_direction, redir, 0);
				}
	case 9:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:123
		{
				  source.dest = 0;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_input_direction, redir, 0);
				}
	case 10:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:129
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_output_direction, redir, 0);
				}
	case 11:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:135
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_input_direction, redir, 0);
				}
	case 12:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:141
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_output_direction, redir, REDIR_VARASSIGN);
				}
	case 13:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:147
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_input_direction, redir, REDIR_VARASSIGN);
				}
	case 14:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:153
		{
				  source.dest = 1;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_appending_to, redir, 0);
				}
	case 15:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:159
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_appending_to, redir, 0);
				}
	case 16:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:165
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_appending_to, redir, REDIR_VARASSIGN);
				}
	case 17:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:171
		{
				  source.dest = 1;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_output_force, redir, 0);
				}
	case 18:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:177
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_output_force, redir, 0);
				}
	case 19:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:183
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_output_force, redir, REDIR_VARASSIGN);
				}
	case 20:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:189
		{
				  source.dest = 0;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_input_output, redir, 0);
				}
	case 21:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:195
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_input_output, redir, 0);
				}
	case 22:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:201
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_input_output, redir, REDIR_VARASSIGN);
				}
	case 23:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:207
		{
				  source.dest = 0;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_reading_until, redir, 0);
				  push_heredoc (yyVAL.redirect);
				}
	case 24:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:214
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_reading_until, redir, 0);
				  push_heredoc (yyVAL.redirect);
				}
	case 25:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:221
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_reading_until, redir, REDIR_VARASSIGN);
				  push_heredoc (yyVAL.redirect);
				}
	case 26:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:228
		{
				  source.dest = 0;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_deblank_reading_until, redir, 0);
				  push_heredoc (yyVAL.redirect);
				}
	case 27:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:235
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_deblank_reading_until, redir, 0);
				  push_heredoc (yyVAL.redirect);
				}
	case 28:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:242
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_deblank_reading_until, redir, REDIR_VARASSIGN);
				  push_heredoc (yyVAL.redirect);
				}
	case 29:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:249
		{
				  source.dest = 0;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_reading_string, redir, 0);
				}
	case 30:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:255
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_reading_string, redir, 0);
				}
	case 31:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:261
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_reading_string, redir, REDIR_VARASSIGN);
				}
	case 32:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:267
		{
				  source.dest = 0;
				  redir.dest = yyDollar[2].number;
				  yyVAL.redirect = make_redirection (source, r_duplicating_input, redir, 0);
				}
	case 33:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:273
		{
				  source.dest = yyDollar[1].number;
				  redir.dest = yyDollar[3].number;
				  yyVAL.redirect = make_redirection (source, r_duplicating_input, redir, 0);
				}
	case 34:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:279
		{
				  source.filename = yyDollar[1].word;
				  redir.dest = yyDollar[3].number;
				  yyVAL.redirect = make_redirection (source, r_duplicating_input, redir, REDIR_VARASSIGN);
				}
	case 35:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:285
		{
				  source.dest = 1;
				  redir.dest = yyDollar[2].number;
				  yyVAL.redirect = make_redirection (source, r_duplicating_output, redir, 0);
				}
	case 36:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:291
		{
				  source.dest = yyDollar[1].number;
				  redir.dest = yyDollar[3].number;
				  yyVAL.redirect = make_redirection (source, r_duplicating_output, redir, 0);
				}
	case 37:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:297
		{
				  source.filename = yyDollar[1].word;
				  redir.dest = yyDollar[3].number;
				  yyVAL.redirect = make_redirection (source, r_duplicating_output, redir, REDIR_VARASSIGN);
				}
	case 38:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:303
		{
				  source.dest = 0;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_duplicating_input_word, redir, 0);
				}
	case 39:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:309
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_duplicating_input_word, redir, 0);
				}
	case 40:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:315
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_duplicating_input_word, redir, REDIR_VARASSIGN);
				}
	case 41:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:321
		{
				  source.dest = 1;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_duplicating_output_word, redir, 0);
				}
	case 42:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:327
		{
				  source.dest = yyDollar[1].number;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_duplicating_output_word, redir, 0);
				}
	case 43:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:333
		{
				  source.filename = yyDollar[1].word;
				  redir.filename = yyDollar[3].word;
				  yyVAL.redirect = make_redirection (source, r_duplicating_output_word, redir, REDIR_VARASSIGN);
				}
	case 44:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:339
		{
				  source.dest = 1;
				  redir.dest = 0;
				  yyVAL.redirect = make_redirection (source, r_close_this, redir, 0);
				}
	case 45:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:345
		{
				  source.dest = yyDollar[1].number;
				  redir.dest = 0;
				  yyVAL.redirect = make_redirection (source, r_close_this, redir, 0);
				}
	case 46:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:351
		{
				  source.filename = yyDollar[1].word;
				  redir.dest = 0;
				  yyVAL.redirect = make_redirection (source, r_close_this, redir, REDIR_VARASSIGN);
				}
	case 47:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:357
		{
				  source.dest = 0;
				  redir.dest = 0;
				  yyVAL.redirect = make_redirection (source, r_close_this, redir, 0);
				}
	case 48:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:363
		{
				  source.dest = yyDollar[1].number;
				  redir.dest = 0;
				  yyVAL.redirect = make_redirection (source, r_close_this, redir, 0);
				}
	case 49:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:369
		{
				  source.filename = yyDollar[1].word;
				  redir.dest = 0;
				  yyVAL.redirect = make_redirection (source, r_close_this, redir, REDIR_VARASSIGN);
				}
	case 50:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:375
		{
				  source.dest = 1;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_err_and_out, redir, 0);
				}
	case 51:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:381
		{
				  source.dest = 1;
				  redir.filename = yyDollar[2].word;
				  yyVAL.redirect = make_redirection (source, r_append_err_and_out, redir, 0);
				}
	case 52:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:389
		{ yyVAL.element.word = yyDollar[1].word; yyVAL.element.redirect = 0; }
	case 53:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:391
		{ yyVAL.element.word = yyDollar[1].word; yyVAL.element.redirect = 0; }
	case 54:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:393
		{ yyVAL.element.redirect = yyDollar[1].redirect; yyVAL.element.word = 0; }
	case 55:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:397
		{
				  yyVAL.redirect = yyDollar[1].redirect;
				}
	case 56:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:401
		{
				  register REDIRECT *t;
	
				  for (t = yyDollar[1].redirect; t->next; t = t->next)
				    ;
				  t->next = yyDollar[2].redirect;
				  yyVAL.redirect = yyDollar[1].redirect;
				}
	case 57:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:412
		{ yyVAL.command = make_simple_command (yyDollar[1].element, (COMMAND *)NULL); }
	case 58:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:414
		{ yyVAL.command = make_simple_command (yyDollar[2].element, yyDollar[1].command); }
	case 59:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:418
		{ yyVAL.command = clean_simple_command (yyDollar[1].command); }
	case 60:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:420
		{ yyVAL.command = yyDollar[1].command; }
	case 61:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:422
		{
				  COMMAND *tc;
	
				  tc = yyDollar[1].command;
				  if (tc && tc->redirects)
				    {
				      register REDIRECT *t;
				      for (t = tc->redirects; t->next; t = t->next)
					;
				      t->next = yyDollar[2].redirect;
				    }
				  else if (tc)
				    tc->redirects = yyDollar[2].redirect;
				  yyVAL.command = yyDollar[1].command;
				}
	case 62:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:438
		{ yyVAL.command = yyDollar[1].command; }
	case 63:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:440
		{ yyVAL.command = yyDollar[1].command; }
	case 64:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:444
		{ yyVAL.command = yyDollar[1].command; }
	case 65:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:446
		{ yyVAL.command = yyDollar[1].command; }
	case 66:
		yyDollar = yyS[yypt-5:yypt+1]
		//line my.go.y:448
		{ yyVAL.command = make_while_command (yyDollar[2].command, yyDollar[4].command); }
	case 67:
		yyDollar = yyS[yypt-5:yypt+1]
		//line my.go.y:450
		{ yyVAL.command = make_until_command (yyDollar[2].command, yyDollar[4].command); }
	case 68:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:452
		{ yyVAL.command = yyDollar[1].command; }
	case 69:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:454
		{ yyVAL.command = yyDollar[1].command; }
	case 70:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:456
		{ yyVAL.command = yyDollar[1].command; }
	case 71:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:458
		{ yyVAL.command = yyDollar[1].command; }
	case 72:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:460
		{ yyVAL.command = yyDollar[1].command; }
	case 73:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:462
		{ yyVAL.command = yyDollar[1].command; }
	case 74:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:464
		{ yyVAL.command = yyDollar[1].command; }
	case 75:
		yyDollar = yyS[yypt-6:yypt+1]
		//line my.go.y:468
		{
				  yyVAL.command = make_for_command (yyDollar[2].word, add_string_to_list ("\"$@\"", (WORD_LIST *)NULL), yyDollar[5].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 76:
		yyDollar = yyS[yypt-6:yypt+1]
		//line my.go.y:473
		{
				  yyVAL.command = make_for_command (yyDollar[2].word, add_string_to_list ("\"$@\"", (WORD_LIST *)NULL), yyDollar[5].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 77:
		yyDollar = yyS[yypt-7:yypt+1]
		//line my.go.y:478
		{
				  yyVAL.command = make_for_command (yyDollar[2].word, add_string_to_list ("\"$@\"", (WORD_LIST *)NULL), yyDollar[6].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 78:
		yyDollar = yyS[yypt-7:yypt+1]
		//line my.go.y:483
		{
				  yyVAL.command = make_for_command (yyDollar[2].word, add_string_to_list ("\"$@\"", (WORD_LIST *)NULL), yyDollar[6].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 79:
		yyDollar = yyS[yypt-10:yypt+1]
		//line my.go.y:488
		{
				  yyVAL.command = make_for_command (yyDollar[2].word, REVERSE_LIST (yyDollar[5].word_list, WORD_LIST *), yyDollar[9].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 80:
		yyDollar = yyS[yypt-10:yypt+1]
		//line my.go.y:493
		{
				  yyVAL.command = make_for_command (yyDollar[2].word, REVERSE_LIST (yyDollar[5].word_list, WORD_LIST *), yyDollar[9].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 81:
		yyDollar = yyS[yypt-9:yypt+1]
		//line my.go.y:498
		{
				  yyVAL.command = make_for_command (yyDollar[2].word, (WORD_LIST *)NULL, yyDollar[8].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 82:
		yyDollar = yyS[yypt-9:yypt+1]
		//line my.go.y:503
		{
				  yyVAL.command = make_for_command (yyDollar[2].word, (WORD_LIST *)NULL, yyDollar[8].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 83:
		yyDollar = yyS[yypt-7:yypt+1]
		//line my.go.y:510
		{
					  yyVAL.command = make_arith_for_command (yyDollar[2].word_list, yyDollar[6].command, arith_for_lineno);
					  if (yyVAL.command == 0) YYERROR;
					  if (word_top > 0) word_top--;
					}
	case 84:
		yyDollar = yyS[yypt-7:yypt+1]
		//line my.go.y:516
		{
					  yyVAL.command = make_arith_for_command (yyDollar[2].word_list, yyDollar[6].command, arith_for_lineno);
					  if (yyVAL.command == 0) YYERROR;
					  if (word_top > 0) word_top--;
					}
	case 85:
		yyDollar = yyS[yypt-5:yypt+1]
		//line my.go.y:522
		{
					  yyVAL.command = make_arith_for_command (yyDollar[2].word_list, yyDollar[4].command, arith_for_lineno);
					  if (yyVAL.command == 0) YYERROR;
					  if (word_top > 0) word_top--;
					}
	case 86:
		yyDollar = yyS[yypt-5:yypt+1]
		//line my.go.y:528
		{
					  yyVAL.command = make_arith_for_command (yyDollar[2].word_list, yyDollar[4].command, arith_for_lineno);
					  if (yyVAL.command == 0) YYERROR;
					  if (word_top > 0) word_top--;
					}
	case 87:
		yyDollar = yyS[yypt-6:yypt+1]
		//line my.go.y:536
		{
				  yyVAL.command = make_select_command (yyDollar[2].word, add_string_to_list ("\"$@\"", (WORD_LIST *)NULL), yyDollar[5].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 88:
		yyDollar = yyS[yypt-6:yypt+1]
		//line my.go.y:541
		{
				  yyVAL.command = make_select_command (yyDollar[2].word, add_string_to_list ("\"$@\"", (WORD_LIST *)NULL), yyDollar[5].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 89:
		yyDollar = yyS[yypt-7:yypt+1]
		//line my.go.y:546
		{
				  yyVAL.command = make_select_command (yyDollar[2].word, add_string_to_list ("\"$@\"", (WORD_LIST *)NULL), yyDollar[6].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 90:
		yyDollar = yyS[yypt-7:yypt+1]
		//line my.go.y:551
		{
				  yyVAL.command = make_select_command (yyDollar[2].word, add_string_to_list ("\"$@\"", (WORD_LIST *)NULL), yyDollar[6].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 91:
		yyDollar = yyS[yypt-10:yypt+1]
		//line my.go.y:556
		{
				  yyVAL.command = make_select_command (yyDollar[2].word, REVERSE_LIST (yyDollar[5].word_list, WORD_LIST *), yyDollar[9].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 92:
		yyDollar = yyS[yypt-10:yypt+1]
		//line my.go.y:561
		{
				  yyVAL.command = make_select_command (yyDollar[2].word, REVERSE_LIST (yyDollar[5].word_list, WORD_LIST *), yyDollar[9].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 93:
		yyDollar = yyS[yypt-9:yypt+1]
		//line my.go.y:566
		{
				  yyVAL.command = make_select_command (yyDollar[2].word, (WORD_LIST *)NULL, yyDollar[8].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 94:
		yyDollar = yyS[yypt-9:yypt+1]
		//line my.go.y:571
		{
				  yyVAL.command = make_select_command (yyDollar[2].word, (WORD_LIST *)NULL, yyDollar[8].command, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 95:
		yyDollar = yyS[yypt-6:yypt+1]
		//line my.go.y:578
		{
				  yyVAL.command = make_case_command (yyDollar[2].word, (PATTERN_LIST *)NULL, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 96:
		yyDollar = yyS[yypt-7:yypt+1]
		//line my.go.y:583
		{
				  yyVAL.command = make_case_command (yyDollar[2].word, yyDollar[5].pattern, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 97:
		yyDollar = yyS[yypt-6:yypt+1]
		//line my.go.y:588
		{
				  yyVAL.command = make_case_command (yyDollar[2].word, yyDollar[5].pattern, word_lineno[word_top]);
				  if (word_top > 0) word_top--;
				}
	case 98:
		yyDollar = yyS[yypt-5:yypt+1]
		//line my.go.y:595
		{ yyVAL.command = make_function_def (yyDollar[1].word, yyDollar[5].command, function_dstart, function_bstart); }
	case 99:
		yyDollar = yyS[yypt-6:yypt+1]
		//line my.go.y:598
		{ yyVAL.command = make_function_def (yyDollar[2].word, yyDollar[6].command, function_dstart, function_bstart); }
	case 100:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:601
		{ yyVAL.command = make_function_def (yyDollar[2].word, yyDollar[4].command, function_dstart, function_bstart); }
	case 101:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:605
		{ yyVAL.command = yyDollar[1].command; }
	case 102:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:607
		{
				  COMMAND *tc;
	
				  tc = yyDollar[1].command;
				  /* According to Posix.2 3.9.5, redirections
			     specified after the body of a function should
			     be attached to the function and performed when
			     the function is executed, not as part of the
			     function definition command. */
				  /* XXX - I don't think it matters, but we might
			     want to change this in the future to avoid
			     problems differentiating between a function
			     definition with a redirection and a function
			     definition containing a single command with a
			     redirection.  The two are semantically equivalent,
			     though -- the only difference is in how the
			     command printing code displays the redirections. */
				  if (tc && tc->redirects)
				    {
				      register REDIRECT *t;
				      for (t = tc->redirects; t->next; t = t->next)
					;
				      t->next = yyDollar[2].redirect;
				    }
				  else if (tc)
				    tc->redirects = yyDollar[2].redirect;
				  yyVAL.command = yyDollar[1].command;
				}
	case 103:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:638
		{
				  yyVAL.command = make_subshell_command (yyDollar[2].command);
				  yyVAL.command->flags |= CMD_WANT_SUBSHELL;
				}
	case 104:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:645
		{
				  yyVAL.command = make_coproc_command ("COPROC", yyDollar[2].command);
				  yyVAL.command->flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
				}
	case 105:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:650
		{
				  COMMAND *tc;
	
				  tc = yyDollar[2].command;
				  if (tc && tc->redirects)
				    {
				      register REDIRECT *t;
				      for (t = tc->redirects; t->next; t = t->next)
					;
				      t->next = yyDollar[3].redirect;
				    }
				  else if (tc)
				    tc->redirects = yyDollar[3].redirect;
				  yyVAL.command = make_coproc_command ("COPROC", yyDollar[2].command);
				  yyVAL.command->flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
				}
	case 106:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:667
		{
				  yyVAL.command = make_coproc_command (yyDollar[2].word->word, yyDollar[3].command);
				  yyVAL.command->flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
				}
	case 107:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:672
		{
				  COMMAND *tc;
	
				  tc = yyDollar[3].command;
				  if (tc && tc->redirects)
				    {
				      register REDIRECT *t;
				      for (t = tc->redirects; t->next; t = t->next)
					;
				      t->next = yyDollar[4].redirect;
				    }
				  else if (tc)
				    tc->redirects = yyDollar[4].redirect;
				  yyVAL.command = make_coproc_command (yyDollar[2].word->word, yyDollar[3].command);
				  yyVAL.command->flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
				}
	case 108:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:689
		{
				  yyVAL.command = make_coproc_command ("COPROC", clean_simple_command (yyDollar[2].command));
				  yyVAL.command->flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
				}
	case 109:
		yyDollar = yyS[yypt-5:yypt+1]
		//line my.go.y:696
		{ yyVAL.command = make_if_command (yyDollar[2].command, yyDollar[4].command, (COMMAND *)NULL); }
	case 110:
		yyDollar = yyS[yypt-7:yypt+1]
		//line my.go.y:698
		{ yyVAL.command = make_if_command (yyDollar[2].command, yyDollar[4].command, yyDollar[6].command); }
	case 111:
		yyDollar = yyS[yypt-6:yypt+1]
		//line my.go.y:700
		{ yyVAL.command = make_if_command (yyDollar[2].command, yyDollar[4].command, yyDollar[5].command); }
	case 112:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:705
		{ yyVAL.command = make_group_command (yyDollar[2].command); }
	case 113:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:709
		{ yyVAL.command = make_arith_command (yyDollar[1].word_list); }
	case 114:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:713
		{ yyVAL.command = yyDollar[2].command; }
	case 115:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:717
		{ yyVAL.command = make_if_command (yyDollar[2].command, yyDollar[4].command, (COMMAND *)NULL); }
	case 116:
		yyDollar = yyS[yypt-6:yypt+1]
		//line my.go.y:719
		{ yyVAL.command = make_if_command (yyDollar[2].command, yyDollar[4].command, yyDollar[6].command); }
	case 117:
		yyDollar = yyS[yypt-5:yypt+1]
		//line my.go.y:721
		{ yyVAL.command = make_if_command (yyDollar[2].command, yyDollar[4].command, yyDollar[5].command); }
	case 119:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:726
		{ yyDollar[2].pattern->next = yyDollar[1].pattern; yyVAL.pattern = yyDollar[2].pattern; }
	case 120:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:730
		{ yyVAL.pattern = make_pattern_list (yyDollar[2].word_list, yyDollar[4].command); }
	case 121:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:732
		{ yyVAL.pattern = make_pattern_list (yyDollar[2].word_list, (COMMAND *)NULL); }
	case 122:
		yyDollar = yyS[yypt-5:yypt+1]
		//line my.go.y:734
		{ yyVAL.pattern = make_pattern_list (yyDollar[3].word_list, yyDollar[5].command); }
	case 123:
		yyDollar = yyS[yypt-5:yypt+1]
		//line my.go.y:736
		{ yyVAL.pattern = make_pattern_list (yyDollar[3].word_list, (COMMAND *)NULL); }
	case 124:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:740
		{ yyVAL.pattern = yyDollar[1].pattern; }
	case 125:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:742
		{ yyDollar[2].pattern->next = yyDollar[1].pattern; yyVAL.pattern = yyDollar[2].pattern; }
	case 126:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:744
		{ yyDollar[1].pattern->flags |= CASEPAT_FALLTHROUGH; yyVAL.pattern = yyDollar[1].pattern; }
	case 127:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:746
		{ yyDollar[2].pattern->flags |= CASEPAT_FALLTHROUGH; yyDollar[2].pattern->next = yyDollar[1].pattern; yyVAL.pattern = yyDollar[2].pattern; }
	case 128:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:748
		{ yyDollar[1].pattern->flags |= CASEPAT_TESTNEXT; yyVAL.pattern = yyDollar[1].pattern; }
	case 129:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:750
		{ yyDollar[2].pattern->flags |= CASEPAT_TESTNEXT; yyDollar[2].pattern->next = yyDollar[1].pattern; yyVAL.pattern = yyDollar[2].pattern; }
	case 130:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:754
		{ yyVAL.word_list = make_word_list (yyDollar[1].word, (WORD_LIST *)NULL); }
	case 131:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:756
		{ yyVAL.word_list = make_word_list (yyDollar[3].word, yyDollar[1].word_list); }
	case 132:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:765
		{
				  yyVAL.command = yyDollar[2].command;
				  if (need_here_doc)
				    gather_here_documents ();
				 }
	case 134:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:774
		{
				  yyVAL.command = yyDollar[2].command;
				}
	case 136:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:781
		{
				  if (yyDollar[1].command->type == cm_connection)
				    yyVAL.command = connect_async_list (yyDollar[1].command, (COMMAND *)NULL, '&');
				  else
				    yyVAL.command = command_connect (yyDollar[1].command, (COMMAND *)NULL, '&');
				}
	case 138:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:792
		{ yyVAL.command = command_connect (yyDollar[1].command, yyDollar[4].command, AND_AND); }
	case 139:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:794
		{ yyVAL.command = command_connect (yyDollar[1].command, yyDollar[4].command, OR_OR); }
	case 140:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:796
		{
				  if (yyDollar[1].command->type == cm_connection)
				    yyVAL.command = connect_async_list (yyDollar[1].command, yyDollar[4].command, '&');
				  else
				    yyVAL.command = command_connect (yyDollar[1].command, yyDollar[4].command, '&');
				}
	case 141:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:803
		{ yyVAL.command = command_connect (yyDollar[1].command, yyDollar[4].command, ';'); }
	case 142:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:805
		{ yyVAL.command = command_connect (yyDollar[1].command, yyDollar[4].command, ';'); }
	case 143:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:807
		{ yyVAL.command = yyDollar[1].command; }
	case 146:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:815
		{ yyVAL.number = '\n'; }
	case 147:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:817
		{ yyVAL.number = ';'; }
	case 148:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:819
		{ yyVAL.number = yacc_EOF; }
	case 151:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:833
		{
				  yyVAL.command = yyDollar[1].command;
				  if (need_here_doc)
				    gather_here_documents ();
				  if ((parser_state & PST_CMDSUBST) && current_token == shell_eof_token)
				    {
				      global_command = yyDollar[1].command;
				      eof_encountered = 0;
				      rewind_input_string ();
				      YYACCEPT;
				    }
				}
	case 152:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:846
		{
				  if (yyDollar[1].command->type == cm_connection)
				    yyVAL.command = connect_async_list (yyDollar[1].command, (COMMAND *)NULL, '&');
				  else
				    yyVAL.command = command_connect (yyDollar[1].command, (COMMAND *)NULL, '&');
				  if (need_here_doc)
				    gather_here_documents ();
				  if ((parser_state & PST_CMDSUBST) && current_token == shell_eof_token)
				    {
				      global_command = yyDollar[1].command;
				      eof_encountered = 0;
				      rewind_input_string ();
				      YYACCEPT;
				    }
				}
	case 153:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:862
		{
				  yyVAL.command = yyDollar[1].command;
				  if (need_here_doc)
				    gather_here_documents ();
				  if ((parser_state & PST_CMDSUBST) && current_token == shell_eof_token)
				    {
				      global_command = yyDollar[1].command;
				      eof_encountered = 0;
				      rewind_input_string ();
				      YYACCEPT;
				    }
				}
	case 154:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:877
		{ yyVAL.command = command_connect (yyDollar[1].command, yyDollar[4].command, AND_AND); }
	case 155:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:879
		{ yyVAL.command = command_connect (yyDollar[1].command, yyDollar[4].command, OR_OR); }
	case 156:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:881
		{
				  if (yyDollar[1].command->type == cm_connection)
				    yyVAL.command = connect_async_list (yyDollar[1].command, yyDollar[3].command, '&');
				  else
				    yyVAL.command = command_connect (yyDollar[1].command, yyDollar[3].command, '&');
				}
	case 157:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:888
		{ yyVAL.command = command_connect (yyDollar[1].command, yyDollar[3].command, ';'); }
	case 158:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:891
		{ yyVAL.command = yyDollar[1].command; }
	case 159:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:895
		{ yyVAL.command = yyDollar[1].command; }
	case 160:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:897
		{
				  if (yyDollar[2].command)
				    yyDollar[2].command->flags ^= CMD_INVERT_RETURN;	/* toggle */
				  yyVAL.command = yyDollar[2].command;
				}
	case 161:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:903
		{
				  if (yyDollar[2].command)
				    yyDollar[2].command->flags |= yyDollar[1].number;
				  yyVAL.command = yyDollar[2].command;
				}
	case 162:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:909
		{
				  ELEMENT x;
	
				  /* Boy, this is unclean.  `time' by itself can
			     time a null command.  We cheat and push a
			     newline back if the list_terminator was a newline
			     to avoid the double-newline problem (one to
			     terminate this, one to terminate the command) */
				  x.word = 0;
				  x.redirect = 0;
				  yyVAL.command = make_simple_command (x, (COMMAND *)NULL);
				  yyVAL.command->flags |= yyDollar[1].number;
				  /* XXX - let's cheat and push a newline back */
				  if (yyDollar[2].number == '\n')
				    token_to_read = '\n';
				  else if (yyDollar[2].number == ';')
				    token_to_read = ';';
				  parser_state &= ~PST_REDIRLIST;	/* make_simple_command sets this */
				}
	case 163:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:929
		{
				  ELEMENT x;
	
				  /* This is just as unclean.  Posix says that `!'
			     by itself should be equivalent to `false'.
			     We cheat and push a
			     newline back if the list_terminator was a newline
			     to avoid the double-newline problem (one to
			     terminate this, one to terminate the command) */
				  x.word = 0;
				  x.redirect = 0;
				  yyVAL.command = make_simple_command (x, (COMMAND *)NULL);
				  yyVAL.command->flags |= CMD_INVERT_RETURN;
				  /* XXX - let's cheat and push a newline back */
				  if (yyDollar[2].number == '\n')
				    token_to_read = '\n';
				  if (yyDollar[2].number == ';')
				    token_to_read = ';';
				  parser_state &= ~PST_REDIRLIST;	/* make_simple_command sets this */
				}
	case 164:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:952
		{ yyVAL.command = command_connect (yyDollar[1].command, yyDollar[4].command, '|'); }
	case 165:
		yyDollar = yyS[yypt-4:yypt+1]
		//line my.go.y:954
		{
				  /* Make cmd1 |& cmd2 equivalent to cmd1 2>&1 | cmd2 */
				  COMMAND *tc;
				  REDIRECTEE rd, sd;
				  REDIRECT *r;
	
				  tc = yyDollar[1].command->type == cm_simple ? (COMMAND *)yyDollar[1].command->value.Simple : yyDollar[1].command;
				  sd.dest = 2;
				  rd.dest = 1;
				  r = make_redirection (sd, r_duplicating_output, rd, 0);
				  if (tc->redirects)
				    {
				      register REDIRECT *t;
				      for (t = tc->redirects; t->next; t = t->next)
					;
				      t->next = r;
				    }
				  else
				    tc->redirects = r;
	
				  yyVAL.command = command_connect (yyDollar[1].command, yyDollar[4].command, '|');
				}
	case 166:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:977
		{ yyVAL.command = yyDollar[1].command; }
	case 167:
		yyDollar = yyS[yypt-1:yypt+1]
		//line my.go.y:981
		{ yyVAL.number = CMD_TIME_PIPELINE; }
	case 168:
		yyDollar = yyS[yypt-2:yypt+1]
		//line my.go.y:983
		{ yyVAL.number = CMD_TIME_PIPELINE|CMD_TIME_POSIX; }
	case 169:
		yyDollar = yyS[yypt-3:yypt+1]
		//line my.go.y:985
		{ yyVAL.number = CMD_TIME_PIPELINE|CMD_TIME_POSIX; }
	}
	goto yystack /* stack new state and value */
}
