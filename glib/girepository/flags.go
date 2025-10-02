package girepository

type ArrayType int32

const (
	ArrayTypeC         ArrayType = 0
	ArrayTypeArray     ArrayType = 1
	ArrayTypePtrArray  ArrayType = 2
	ArrayTypeByteArray ArrayType = 3
)

type Direction int32

const (
	DirectionIn    Direction = 0
	DirectionOut   Direction = 1
	DirectionInOut Direction = 2
)

type ScopeType int32

const (
	ScopeTypeInvalid  ScopeType = 0
	ScopeTypeCall     ScopeType = 1
	ScopeTypeAsync    ScopeType = 2
	ScopeTypeNotified ScopeType = 3
	ScopeTypeForever  ScopeType = 4
)

type Transfer int32

const (
	TransferNothing    Transfer = 0
	TransferContainer  Transfer = 1
	TransferEverything Transfer = 2
)

type TypeTag int32

const (
	TypeTagVoid      TypeTag = 0
	TypeTagBoolean   TypeTag = 1
	TypeTagInt8      TypeTag = 2
	TypeTagUint8     TypeTag = 3
	TypeTagInt16     TypeTag = 4
	TypeTagUint16    TypeTag = 5
	TypeTagInt32     TypeTag = 6
	TypeTagUint32    TypeTag = 7
	TypeTagInt64     TypeTag = 8
	TypeTagUint64    TypeTag = 9
	TypeTagFloat     TypeTag = 10
	TypeTagDouble    TypeTag = 11
	TypeTagGType     TypeTag = 12
	TypeTagUTF8      TypeTag = 13
	TypeTagFilename  TypeTag = 14
	TypeTagArray     TypeTag = 15
	TypeTagInterface TypeTag = 16
	TypeTagGList     TypeTag = 17
	TypeTagGSList    TypeTag = 18
	TypeTagGHash     TypeTag = 19
	TypeTagError     TypeTag = 20
	TypeTagUnichar   TypeTag = 21
)

type RepositoryLoadFlags int32

const (
	RepositoryLoadFlagNone RepositoryLoadFlags = 0
	RepositoryLoadFlagLazy RepositoryLoadFlags = 1 << 0
)

type FieldInfoFlags int32

const (
	FieldIsReadable FieldInfoFlags = 1 << 0
	FieldIsWritable FieldInfoFlags = 1 << 1
)

type FunctionInfoFlags int32

const (
	FunctionIsMethod      FunctionInfoFlags = 1 << 0
	FunctionIsConstructor FunctionInfoFlags = 1 << 1
	FunctionIsGetter      FunctionInfoFlags = 1 << 2
	FunctionIsSetter      FunctionInfoFlags = 1 << 3
	FunctionWrapsVFunc    FunctionInfoFlags = 1 << 4
)

type VFuncInfoFlags int32

const (
	VFuncMustChainUp     VFuncInfoFlags = 1 << 0
	VFuncMustOverride    VFuncInfoFlags = 1 << 1
	VFuncMustNotOverride VFuncInfoFlags = 1 << 2
)
