// Code generated by "stringer -type=TokenType,astTokenType,pBlockType,opcode -output stringer.go"; DO NOT EDIT.

package goal

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NONE-0]
	_ = x[EOF-1]
	_ = x[ERROR-2]
	_ = x[ADVERB-3]
	_ = x[IDENT-4]
	_ = x[LEFTBRACE-5]
	_ = x[LEFTBRACKET-6]
	_ = x[LEFTPAREN-7]
	_ = x[NEWLINE-8]
	_ = x[NUMBER-9]
	_ = x[RIGHTBRACE-10]
	_ = x[RIGHTBRACKET-11]
	_ = x[RIGHTPAREN-12]
	_ = x[SEMICOLON-13]
	_ = x[STRING-14]
	_ = x[VERB-15]
}

const _TokenType_name = "NONEEOFERRORADVERBIDENTLEFTBRACELEFTBRACKETLEFTPARENNEWLINENUMBERRIGHTBRACERIGHTBRACKETRIGHTPARENSEMICOLONSTRINGVERB"

var _TokenType_index = [...]uint8{0, 4, 7, 12, 18, 23, 32, 43, 52, 59, 65, 75, 87, 97, 106, 112, 116}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[astSEP-0]
	_ = x[astEOF-1]
	_ = x[astCLOSE-2]
	_ = x[astNUMBER-3]
	_ = x[astSTRING-4]
	_ = x[astIDENT-5]
	_ = x[astVERB-6]
	_ = x[astADVERB-7]
}

const _astTokenType_name = "astSEPastEOFastCLOSEastNUMBERastSTRINGastIDENTastVERBastADVERB"

var _astTokenType_index = [...]uint8{0, 6, 12, 20, 29, 38, 46, 53, 62}

func (i astTokenType) String() string {
	if i < 0 || i >= astTokenType(len(_astTokenType_index)-1) {
		return "astTokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _astTokenType_name[_astTokenType_index[i]:_astTokenType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[astLAMBDA-0]
	_ = x[astARGS-1]
	_ = x[astSEQ-2]
	_ = x[astLIST-3]
}

const _pBlockType_name = "pLAMBDApARGSpSEQpLIST"

var _pBlockType_index = [...]uint8{0, 7, 12, 16, 21}

func (i astBlockType) String() string {
	if i < 0 || i >= astBlockType(len(_pBlockType_index)-1) {
		return "pBlockType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _pBlockType_name[_pBlockType_index[i]:_pBlockType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opNop-0]
	_ = x[opConst-1]
	_ = x[opNil-2]
	_ = x[opGlobal-3]
	_ = x[opLocal-4]
	_ = x[opAssignGlobal-5]
	_ = x[opAssignLocal-6]
	_ = x[opAdverb-7]
	_ = x[opVariadic-8]
	_ = x[opLambda-9]
	_ = x[opApply-10]
	_ = x[opApply2-11]
	_ = x[opApplyN-12]
	_ = x[opDrop-13]
}

const _opcode_name = "opNopopConstopNilopGlobalopLocalopAssignGlobalopAssignLocalopAdverbopVariadicopLambdaopApplyopApply2opApplyNopDrop"

var _opcode_index = [...]uint8{0, 5, 12, 17, 25, 32, 46, 59, 67, 77, 85, 92, 100, 108, 114}

func (i opcode) String() string {
	if i < 0 || i >= opcode(len(_opcode_index)-1) {
		return "opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _opcode_name[_opcode_index[i]:_opcode_index[i+1]]
}
