package treesitter

// typedef uint16_t TSStateId;
// typedef uint16_t TSSymbol;
// typedef uint16_t TSFieldId;

type StateId uint16
type Symbol uint16
type FieldId uint16

type InputEncoding uint32
type SymbolType uint32
type LogType uint32
type Quantifier uint32
type QueryPredicateStepType uint32
type QueryError uint32
type WasmErrorKind uint32

const (
	InputEncodingUTF8 InputEncoding = iota
	InputEncodingUTF16LE
	InputEncodingUTF16BE
	InputEncodingCustom
)
const (
	SymbolTypeRegular SymbolType = iota
	SymbolTypeAnonymous
	SymbolTypeSupertype
	SymbolTypeAuxiliary
)
const (
	LogTypeParse LogType = iota
	LogTypeLex
)
const (
	QuantifierZero Quantifier = iota // must match the array initialization value
	QuantifierZeroOrOne
	QuantifierZeroOrMore
	QuantifierOne
	QuantifierOneOrMore
)
const (
	QueryPredicateStepTypeDone QueryPredicateStepType = iota
	QueryPredicateStepTypeCapture
	QueryPredicateStepTypeString
)
const (
	QueryErrorNone QueryError = iota
	QueryErrorSyntax
	QueryErrorNodeType
	QueryErrorField
	QueryErrorCapture
	QueryErrorStructure
	QueryErrorLanguage
)
const (
	WasmErrorKindNone WasmErrorKind = iota
	WasmErrorKindParse
	WasmErrorKindCompile
	WasmErrorKindInstantiate
	WasmErrorKindAllocate
)
