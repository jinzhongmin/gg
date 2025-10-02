package libclang

type CXErrorCode uint32

const (
	/**
	 * No error.
	 */
	CXError_Success CXErrorCode = 0

	/**
	 * A generic error code, no further details are available.
	 *
	 * Errors of this kind can get their own specific error codes in future
	 * libclang versions.
	 */
	CXError_Failure CXErrorCode = 1

	/**
	 * libclang crashed while performing the requested operation.
	 */
	CXError_Crashed CXErrorCode = 2

	/**
	 * The function detected that the arguments violate the function
	 * contract.
	 */
	CXError_InvalidArguments CXErrorCode = 3

	/**
	 * An AST deserialization error has occurred.
	 */
	CXError_ASTReadError CXErrorCode = 4
)

type CXCompilationDatabase_Error uint32

const (
	/*
	 * No error occurred
	 */
	CXCompilationDatabase_NoError CXCompilationDatabase_Error = 0

	/*
	 * Database can not be loaded
	 */
	CXCompilationDatabase_CanNotLoadDatabase CXCompilationDatabase_Error = 1
)

type CXDiagnosticSeverity uint32

const (
	/**
	 * A diagnostic that has been suppressed, e.g., by a command-line
	 * option.
	 */
	CXDiagnostic_Ignored CXDiagnosticSeverity = 0

	/**
	 * This diagnostic is a note that should be attached to the
	 * previous (non-note) diagnostic.
	 */
	CXDiagnostic_Note CXDiagnosticSeverity = 1

	/**
	 * This diagnostic indicates suspicious code that may not be
	 * wrong.
	 */
	CXDiagnostic_Warning CXDiagnosticSeverity = 2

	/**
	 * This diagnostic indicates that the code is ill-formed.
	 */
	CXDiagnostic_Error CXDiagnosticSeverity = 3

	/**
	 * This diagnostic indicates that the code is ill-formed such
	 * that future parser recovery is unlikely to produce useful
	 * results.
	 */
	CXDiagnostic_Fatal CXDiagnosticSeverity = 4
)

type CXLoadDiag_Error uint32

const (
	/**
	 * Indicates that no error occurred.
	 */
	CXLoadDiag_None CXLoadDiag_Error = 0

	/**
	 * Indicates that an unknown error occurred while attempting to
	 * deserialize diagnostics.
	 */
	CXLoadDiag_Unknown CXLoadDiag_Error = 1

	/**
	 * Indicates that the file containing the serialized diagnostics
	 * could not be opened.
	 */
	CXLoadDiag_CannotLoad CXLoadDiag_Error = 2

	/**
	 * Indicates that the serialized diagnostics file is invalid or
	 * corrupt.
	 */
	CXLoadDiag_InvalidFile CXLoadDiag_Error = 3
)

type CXDiagnosticDisplayOptions uint32

const (
	/**
	 * Display the source-location information where the
	 * diagnostic was located.
	 *
	 * When set, diagnostics will be prefixed by the file, line, and
	 * (optionally) column to which the diagnostic refers. For example,
	 *
	 * \code
	 * test.c:28: warning: extra tokens at end of #endif directive
	 * \endcode
	 *
	 * This option corresponds to the clang flag \c -fshow-source-location.
	 */
	CXDiagnostic_DisplaySourceLocation CXDiagnosticDisplayOptions = 0x01

	/**
	 * If displaying the source-location information of the
	 * diagnostic, also include the column number.
	 *
	 * This option corresponds to the clang flag \c -fshow-column.
	 */
	CXDiagnostic_DisplayColumn CXDiagnosticDisplayOptions = 0x02

	/**
	 * If displaying the source-location information of the
	 * diagnostic, also include information about source ranges in a
	 * machine-parsable format.
	 *
	 * This option corresponds to the clang flag
	 * \c -fdiagnostics-print-source-range-info.
	 */
	CXDiagnostic_DisplaySourceRanges CXDiagnosticDisplayOptions = 0x04

	/**
	 * Display the option name associated with this diagnostic, if any.
	 *
	 * The option name displayed (e.g., -Wconversion) will be placed in brackets
	 * after the diagnostic text. This option corresponds to the clang flag
	 * \c -fdiagnostics-show-option.
	 */
	CXDiagnostic_DisplayOption CXDiagnosticDisplayOptions = 0x08

	/**
	 * Display the category number associated with this diagnostic, if any.
	 *
	 * The category number is displayed within brackets after the diagnostic text.
	 * This option corresponds to the clang flag
	 * \c -fdiagnostics-show-category=id.
	 */
	CXDiagnostic_DisplayCategoryId CXDiagnosticDisplayOptions = 0x10

	/**
	 * Display the category name associated with this diagnostic, if any.
	 *
	 * The category name is displayed within brackets after the diagnostic text.
	 * This option corresponds to the clang flag
	 * \c -fdiagnostics-show-category=name.
	 */
	CXDiagnostic_DisplayCategoryName CXDiagnosticDisplayOptions = 0x20
)

type CXTranslationUnit_Flags int32

const (
	CXTranslationUnit_None                                 CXTranslationUnit_Flags = 0x0
	CXTranslationUnit_DetailedPreprocessingRecord          CXTranslationUnit_Flags = 0x01
	CXTranslationUnit_Incomplete                           CXTranslationUnit_Flags = 0x02
	CXTranslationUnit_PrecompiledPreamble                  CXTranslationUnit_Flags = 0x04
	CXTranslationUnit_CacheCompletionResults               CXTranslationUnit_Flags = 0x08
	CXTranslationUnit_ForSerialization                     CXTranslationUnit_Flags = 0x10
	CXTranslationUnit_CXXChainedPCH                        CXTranslationUnit_Flags = 0x20
	CXTranslationUnit_SkipFunctionBodies                   CXTranslationUnit_Flags = 0x40
	CXTranslationUnit_IncludeBriefCommentsInCodeCompletion CXTranslationUnit_Flags = 0x80
	CXTranslationUnit_CreatePreambleOnFirstParse           CXTranslationUnit_Flags = 0x100
	CXTranslationUnit_KeepGoing                            CXTranslationUnit_Flags = 0x200
	CXTranslationUnit_SingleFileParse                      CXTranslationUnit_Flags = 0x400
	CXTranslationUnit_LimitSkipFunctionBodiesToPreamble    CXTranslationUnit_Flags = 0x800
	CXTranslationUnit_IncludeAttributedTypes               CXTranslationUnit_Flags = 0x1000
	CXTranslationUnit_VisitImplicitAttributes              CXTranslationUnit_Flags = 0x2000
	CXTranslationUnit_IgnoreNonErrorsFromIncludedFiles     CXTranslationUnit_Flags = 0x4000
	CXTranslationUnit_RetainExcludedConditionalBlocks      CXTranslationUnit_Flags = 0x8000
)

type CXGlobalOptFlags uint32

const (
	CXGlobalOpt_None                                CXGlobalOptFlags = 0x0
	CXGlobalOpt_ThreadBackgroundPriorityForIndexing CXGlobalOptFlags = 0x1
	CXGlobalOpt_ThreadBackgroundPriorityForEditing  CXGlobalOptFlags = 0x2
	CXGlobalOpt_ThreadBackgroundPriorityForAll      CXGlobalOptFlags = 0x3
)

type CXCursorKind uint32

const (

	/* Declarations */
	/**
	 * A declaration whose specific kind is not exposed via this
	 * interface.
	 *
	 * Unexposed declarations have the same operations as any other kind
	 * of declaration; one can extract their location information
	 * spelling, find their definitions, etc. However, the specific kind
	 * of the declaration is not reported.
	 */
	CXCursor_UnexposedDecl CXCursorKind = 1
	/** A C or C++ struct. */
	CXCursor_StructDecl CXCursorKind = 2
	/** A C or C++ union. */
	CXCursor_UnionDecl CXCursorKind = 3
	/** A C++ class. */
	CXCursor_ClassDecl CXCursorKind = 4
	/** An enumeration. */
	CXCursor_EnumDecl CXCursorKind = 5
	/**
	 * A field (in C) or non-static data member (in C++) in a
	 * struct, union, or C++ class.
	 */
	CXCursor_FieldDecl CXCursorKind = 6
	/** An enumerator constant. */
	CXCursor_EnumConstantDecl CXCursorKind = 7
	/** A function. */
	CXCursor_FunctionDecl CXCursorKind = 8
	/** A variable. */
	CXCursor_VarDecl CXCursorKind = 9
	/** A function or method parameter. */
	CXCursor_ParmDecl CXCursorKind = 10
	/** An Objective-C \@interface. */
	CXCursor_ObjCInterfaceDecl CXCursorKind = 11
	/** An Objective-C \@interface for a category. */
	CXCursor_ObjCCategoryDecl CXCursorKind = 12
	/** An Objective-C \@protocol declaration. */
	CXCursor_ObjCProtocolDecl CXCursorKind = 13
	/** An Objective-C \@property declaration. */
	CXCursor_ObjCPropertyDecl CXCursorKind = 14
	/** An Objective-C instance variable. */
	CXCursor_ObjCIvarDecl CXCursorKind = 15
	/** An Objective-C instance method. */
	CXCursor_ObjCInstanceMethodDecl CXCursorKind = 16
	/** An Objective-C class method. */
	CXCursor_ObjCClassMethodDecl CXCursorKind = 17
	/** An Objective-C \@implementation. */
	CXCursor_ObjCImplementationDecl CXCursorKind = 18
	/** An Objective-C \@implementation for a category. */
	CXCursor_ObjCCategoryImplDecl CXCursorKind = 19
	/** A typedef. */
	CXCursor_TypedefDecl CXCursorKind = 20
	/** A C++ class method. */
	CXCursor_CXXMethod CXCursorKind = 21
	/** A C++ namespace. */
	CXCursor_Namespace CXCursorKind = 22
	/** A linkage specification, e.g. 'extern "C"'. */
	CXCursor_LinkageSpec CXCursorKind = 23
	/** A C++ constructor. */
	CXCursor_Constructor CXCursorKind = 24
	/** A C++ destructor. */
	CXCursor_Destructor CXCursorKind = 25
	/** A C++ conversion function. */
	CXCursor_ConversionFunction CXCursorKind = 26
	/** A C++ template type parameter. */
	CXCursor_TemplateTypeParameter CXCursorKind = 27
	/** A C++ non-type template parameter. */
	CXCursor_NonTypeTemplateParameter CXCursorKind = 28
	/** A C++ template template parameter. */
	CXCursor_TemplateTemplateParameter CXCursorKind = 29
	/** A C++ function template. */
	CXCursor_FunctionTemplate CXCursorKind = 30
	/** A C++ class template. */
	CXCursor_ClassTemplate CXCursorKind = 31
	/** A C++ class template partial specialization. */
	CXCursor_ClassTemplatePartialSpecialization CXCursorKind = 32
	/** A C++ namespace alias declaration. */
	CXCursor_NamespaceAlias CXCursorKind = 33
	/** A C++ using directive. */
	CXCursor_UsingDirective CXCursorKind = 34
	/** A C++ using declaration. */
	CXCursor_UsingDeclaration CXCursorKind = 35
	/** A C++ alias declaration */
	CXCursor_TypeAliasDecl CXCursorKind = 36
	/** An Objective-C \@synthesize definition. */
	CXCursor_ObjCSynthesizeDecl CXCursorKind = 37
	/** An Objective-C \@dynamic definition. */
	CXCursor_ObjCDynamicDecl CXCursorKind = 38
	/** An access specifier. */
	CXCursor_CXXAccessSpecifier CXCursorKind = 39

	CXCursor_FirstDecl CXCursorKind = CXCursor_UnexposedDecl
	CXCursor_LastDecl  CXCursorKind = CXCursor_CXXAccessSpecifier

	/* References */
	CXCursor_FirstRef          CXCursorKind = 40 /* Decl references */
	CXCursor_ObjCSuperClassRef CXCursorKind = 40
	CXCursor_ObjCProtocolRef   CXCursorKind = 41
	CXCursor_ObjCClassRef      CXCursorKind = 42
	/**
	 * A reference to a type declaration.
	 *
	 * A type reference occurs anywhere where a type is named but not
	 * declared. For example, given:
	 *
	 * \code
	 * typedef unsigned size_type;
	 * size_type size;
	 * \endcode
	 *
	 * The typedef is a declaration of size_type (CXCursor_TypedefDecl)
	 * while the type of the variable "size" is referenced. The cursor
	 * referenced by the type of size is the typedef for size_type.
	 */
	CXCursor_TypeRef          CXCursorKind = 43
	CXCursor_CXXBaseSpecifier CXCursorKind = 44
	/**
	 * A reference to a class template, function template, template
	 * template parameter, or class template partial specialization.
	 */
	CXCursor_TemplateRef CXCursorKind = 45
	/**
	 * A reference to a namespace or namespace alias.
	 */
	CXCursor_NamespaceRef CXCursorKind = 46
	/**
	 * A reference to a member of a struct, union, or class that occurs in
	 * some non-expression context, e.g., a designated initializer.
	 */
	CXCursor_MemberRef CXCursorKind = 47
	/**
	 * A reference to a labeled statement.
	 *
	 * This cursor kind is used to describe the jump to "start_over" in the
	 * goto statement in the following example:
	 *
	 * \code
	 *   start_over:
	 *     ++counter;
	 *
	 *     goto start_over;
	 * \endcode
	 *
	 * A label reference cursor refers to a label statement.
	 */
	CXCursor_LabelRef CXCursorKind = 48

	/**
	 * A reference to a set of overloaded functions or function templates
	 * that has not yet been resolved to a specific function or function template.
	 *
	 * An overloaded declaration reference cursor occurs in C++ templates where
	 * a dependent name refers to a function. For example:
	 *
	 * \code
	 * template<typename T> void swap(T&, T&);
	 *
	 * struct X { ... };
	 * void swap(X&, X&);
	 *
	 * template<typename T>
	 * void reverse(T* first, T* last) {
	 *   while (first < last - 1) {
	 *     swap(*first, *--last);
	 *     ++first;
	 *   }
	 * }
	 *
	 * struct Y { };
	 * void swap(Y&, Y&);
	 * \endcode
	 *
	 * Here, the identifier "swap" is associated with an overloaded declaration
	 * reference. In the template definition, "swap" refers to either of the two
	 * "swap" functions declared above, so both results will be available. At
	 * instantiation time, "swap" may also refer to other functions found via
	 * argument-dependent lookup (e.g., the "swap" function at the end of the
	 * example).
	 *
	 * The functions \c clang_getNumOverloadedDecls() and
	 * \c clang_getOverloadedDecl() can be used to retrieve the definitions
	 * referenced by this cursor.
	 */
	CXCursor_OverloadedDeclRef CXCursorKind = 49

	/**
	 * A reference to a variable that occurs in some non-expression
	 * context, e.g., a C++ lambda capture list.
	 */
	CXCursor_VariableRef CXCursorKind = 50

	CXCursor_LastRef CXCursorKind = CXCursor_VariableRef

	/* Error conditions */
	CXCursor_FirstInvalid   CXCursorKind = 70
	CXCursor_InvalidFile    CXCursorKind = 70
	CXCursor_NoDeclFound    CXCursorKind = 71
	CXCursor_NotImplemented CXCursorKind = 72
	CXCursor_InvalidCode    CXCursorKind = 73
	CXCursor_LastInvalid    CXCursorKind = CXCursor_InvalidCode

	/* Expressions */
	CXCursor_FirstExpr CXCursorKind = 100

	/**
	 * An expression whose specific kind is not exposed via this
	 * interface.
	 *
	 * Unexposed expressions have the same operations as any other kind
	 * of expression; one can extract their location information
	 * spelling, children, etc. However, the specific kind of the
	 * expression is not reported.
	 */
	CXCursor_UnexposedExpr CXCursorKind = 100

	/**
	 * An expression that refers to some value declaration, such
	 * as a function, variable, or enumerator.
	 */
	CXCursor_DeclRefExpr CXCursorKind = 101

	/**
	 * An expression that refers to a member of a struct, union
	 * class, Objective-C class, etc.
	 */
	CXCursor_MemberRefExpr CXCursorKind = 102

	/** An expression that calls a function. */
	CXCursor_CallExpr CXCursorKind = 103

	/** An expression that sends a message to an Objective-C
	  object or class. */
	CXCursor_ObjCMessageExpr CXCursorKind = 104

	/** An expression that represents a block literal. */
	CXCursor_BlockExpr CXCursorKind = 105

	/** An integer literal.
	 */
	CXCursor_IntegerLiteral CXCursorKind = 106

	/** A floating point number literal.
	 */
	CXCursor_FloatingLiteral CXCursorKind = 107

	/** An imaginary number literal.
	 */
	CXCursor_ImaginaryLiteral CXCursorKind = 108

	/** A string literal.
	 */
	CXCursor_StringLiteral CXCursorKind = 109

	/** A character literal.
	 */
	CXCursor_CharacterLiteral CXCursorKind = 110

	/** A parenthesized expression, e.g. "(1)".
	 *
	 * This AST node is only formed if full location information is requested.
	 */
	CXCursor_ParenExpr CXCursorKind = 111

	/** This represents the unary-expression's (except sizeof and
	 * alignof).
	 */
	CXCursor_UnaryOperator CXCursorKind = 112

	/** [C99 6.5.2.1] Array Subscripting.
	 */
	CXCursor_ArraySubscriptExpr CXCursorKind = 113

	/** A builtin binary operation expression such as "x + y" or
	 * "x <= y".
	 */
	CXCursor_BinaryOperator CXCursorKind = 114

	/** Compound assignment such as "+=".
	 */
	CXCursor_CompoundAssignOperator CXCursorKind = 115

	/** The ?: ternary operator.
	 */
	CXCursor_ConditionalOperator CXCursorKind = 116

	/** An explicit cast in C (C99 6.5.4) or a C-style cast in C++
	 * (C++ [expr.cast]), which uses the syntax (Type)expr.
	 *
	 * For example: (int)f.
	 */
	CXCursor_CStyleCastExpr CXCursorKind = 117

	/** [C99 6.5.2.5]
	 */
	CXCursor_CompoundLiteralExpr CXCursorKind = 118

	/** Describes an C or C++ initializer list.
	 */
	CXCursor_InitListExpr CXCursorKind = 119

	/** The GNU address of label extension, representing &&label.
	 */
	CXCursor_AddrLabelExpr CXCursorKind = 120

	/** This is the GNU Statement Expression extension: ({int X=4; X;})
	 */
	CXCursor_StmtExpr CXCursorKind = 121

	/** Represents a C11 generic selection.
	 */
	CXCursor_GenericSelectionExpr CXCursorKind = 122

	/** Implements the GNU __null extension, which is a name for a null
	 * pointer constant that has integral type (e.g., int or long) and is the same
	 * size and alignment as a pointer.
	 *
	 * The __null extension is typically only used by system headers, which define
	 * NULL as __null in C++ rather than using 0 (which is an integer that may not
	 * match the size of a pointer).
	 */
	CXCursor_GNUNullExpr CXCursorKind = 123

	/** C++'s static_cast<> expression.
	 */
	CXCursor_CXXStaticCastExpr CXCursorKind = 124

	/** C++'s dynamic_cast<> expression.
	 */
	CXCursor_CXXDynamicCastExpr CXCursorKind = 125

	/** C++'s reinterpret_cast<> expression.
	 */
	CXCursor_CXXReinterpretCastExpr CXCursorKind = 126

	/** C++'s const_cast<> expression.
	 */
	CXCursor_CXXConstCastExpr CXCursorKind = 127

	/** Represents an explicit C++ type conversion that uses "functional"
	 * notion (C++ [expr.type.conv]).
	 *
	 * Example:
	 * \code
	 *   x CXCursorKind = int(0.5);
	 * \endcode
	 */
	CXCursor_CXXFunctionalCastExpr CXCursorKind = 128

	/** A C++ typeid expression (C++ [expr.typeid]).
	 */
	CXCursor_CXXTypeidExpr CXCursorKind = 129

	/** [C++ 2.13.5] C++ Boolean Literal.
	 */
	CXCursor_CXXBoolLiteralExpr CXCursorKind = 130

	/** [C++0x 2.14.7] C++ Pointer Literal.
	 */
	CXCursor_CXXNullPtrLiteralExpr CXCursorKind = 131

	/** Represents the "this" expression in C++
	 */
	CXCursor_CXXThisExpr CXCursorKind = 132

	/** [C++ 15] C++ Throw Expression.
	 *
	 * This handles 'throw' and 'throw' assignment-expression. When
	 * assignment-expression isn't present, Op will be null.
	 */
	CXCursor_CXXThrowExpr CXCursorKind = 133

	/** A new expression for memory allocation and constructor calls, e.g:
	 * "new CXXNewExpr(foo)".
	 */
	CXCursor_CXXNewExpr CXCursorKind = 134

	/** A delete expression for memory deallocation and destructor calls
	 * e.g. "delete[] pArray".
	 */
	CXCursor_CXXDeleteExpr CXCursorKind = 135

	/** A unary expression. (noexcept, sizeof, or other traits)
	 */
	CXCursor_UnaryExpr CXCursorKind = 136

	/** An Objective-C string literal i.e. @"foo".
	 */
	CXCursor_ObjCStringLiteral CXCursorKind = 137

	/** An Objective-C \@encode expression.
	 */
	CXCursor_ObjCEncodeExpr CXCursorKind = 138

	/** An Objective-C \@selector expression.
	 */
	CXCursor_ObjCSelectorExpr CXCursorKind = 139

	/** An Objective-C \@protocol expression.
	 */
	CXCursor_ObjCProtocolExpr CXCursorKind = 140

	/** An Objective-C "bridged" cast expression, which casts between
	 * Objective-C pointers and C pointers, transferring ownership in the process.
	 *
	 * \code
	 *   NSString *str CXCursorKind = (__bridge_transfer NSString *)CFCreateString();
	 * \endcode
	 */
	CXCursor_ObjCBridgedCastExpr CXCursorKind = 141

	/** Represents a C++0x pack expansion that produces a sequence of
	 * expressions.
	 *
	 * A pack expansion expression contains a pattern (which itself is an
	 * expression) followed by an ellipsis. For example:
	 *
	 * \code
	 * template<typename F, typename ...Types>
	 * void forward(F f, Types &&...args) {
	 *  f(static_cast<Types&&>(args)...);
	 * }
	 * \endcode
	 */
	CXCursor_PackExpansionExpr CXCursorKind = 142

	/** Represents an expression that computes the length of a parameter
	 * pack.
	 *
	 * \code
	 * template<typename ...Types>
	 * struct count {
	 *   static const unsigned value CXCursorKind = sizeof...(Types);
	 * };
	 * \endcode
	 */
	CXCursor_SizeOfPackExpr CXCursorKind = 143

	/* Represents a C++ lambda expression that produces a local function
	 * object.
	 *
	 * \code
	 * void abssort(float *x, unsigned N) {
	 *   std::sort(x, x + N
	 *             [](float a, float b) {
	 *               return std::abs(a) < std::abs(b);
	 *             });
	 * }
	 * \endcode
	 */
	CXCursor_LambdaExpr CXCursorKind = 144

	/** Objective-c Boolean Literal.
	 */
	CXCursor_ObjCBoolLiteralExpr CXCursorKind = 145

	/** Represents the "self" expression in an Objective-C method.
	 */
	CXCursor_ObjCSelfExpr CXCursorKind = 146

	/** OpenMP 5.0 [2.1.5, Array Section].
	 * OpenACC 3.3 [2.7.1, Data Specification for Data Clauses (Sub Arrays)]
	 */
	CXCursor_ArraySectionExpr CXCursorKind = 147

	/** Represents an @available(...) check.
	 */
	CXCursor_ObjCAvailabilityCheckExpr CXCursorKind = 148

	/**
	 * Fixed point literal
	 */
	CXCursor_FixedPointLiteral CXCursorKind = 149

	/** OpenMP 5.0 [2.1.4, Array Shaping].
	 */
	CXCursor_OMPArrayShapingExpr CXCursorKind = 150

	/**
	 * OpenMP 5.0 [2.1.6 Iterators]
	 */
	CXCursor_OMPIteratorExpr CXCursorKind = 151

	/** OpenCL's addrspace_cast<> expression.
	 */
	CXCursor_CXXAddrspaceCastExpr CXCursorKind = 152

	/**
	 * Expression that references a C++20 concept.
	 */
	CXCursor_ConceptSpecializationExpr CXCursorKind = 153

	/**
	 * Expression that references a C++20 requires expression.
	 */
	CXCursor_RequiresExpr CXCursorKind = 154

	/**
	 * Expression that references a C++20 parenthesized list aggregate
	 * initializer.
	 */
	CXCursor_CXXParenListInitExpr CXCursorKind = 155

	/**
	 *  Represents a C++26 pack indexing expression.
	 */
	CXCursor_PackIndexingExpr CXCursorKind = 156

	CXCursor_LastExpr CXCursorKind = CXCursor_PackIndexingExpr

	/* Statements */
	CXCursor_FirstStmt CXCursorKind = 200
	/**
	 * A statement whose specific kind is not exposed via this
	 * interface.
	 *
	 * Unexposed statements have the same operations as any other kind of
	 * statement; one can extract their location information, spelling
	 * children, etc. However, the specific kind of the statement is not
	 * reported.
	 */
	CXCursor_UnexposedStmt CXCursorKind = 200

	/** A labelled statement in a function.
	 *
	 * This cursor kind is used to describe the "start_over:" label statement in
	 * the following example:
	 *
	 * \code
	 *   start_over:
	 *     ++counter;
	 * \endcode
	 *
	 */
	CXCursor_LabelStmt CXCursorKind = 201

	/** A group of statements like { stmt stmt }.
	 *
	 * This cursor kind is used to describe compound statements, e.g. function
	 * bodies.
	 */
	CXCursor_CompoundStmt CXCursorKind = 202

	/** A case statement.
	 */
	CXCursor_CaseStmt CXCursorKind = 203

	/** A default statement.
	 */
	CXCursor_DefaultStmt CXCursorKind = 204

	/** An if statement
	 */
	CXCursor_IfStmt CXCursorKind = 205

	/** A switch statement.
	 */
	CXCursor_SwitchStmt CXCursorKind = 206

	/** A while statement.
	 */
	CXCursor_WhileStmt CXCursorKind = 207

	/** A do statement.
	 */
	CXCursor_DoStmt CXCursorKind = 208

	/** A for statement.
	 */
	CXCursor_ForStmt CXCursorKind = 209

	/** A goto statement.
	 */
	CXCursor_GotoStmt CXCursorKind = 210

	/** An indirect goto statement.
	 */
	CXCursor_IndirectGotoStmt CXCursorKind = 211

	/** A continue statement.
	 */
	CXCursor_ContinueStmt CXCursorKind = 212

	/** A break statement.
	 */
	CXCursor_BreakStmt CXCursorKind = 213

	/** A return statement.
	 */
	CXCursor_ReturnStmt CXCursorKind = 214

	/** A GCC inline assembly statement extension.
	 */
	CXCursor_GCCAsmStmt CXCursorKind = 215
	CXCursor_AsmStmt    CXCursorKind = CXCursor_GCCAsmStmt

	/** Objective-C's overall \@try-\@catch-\@finally statement.
	 */
	CXCursor_ObjCAtTryStmt CXCursorKind = 216

	/** Objective-C's \@catch statement.
	 */
	CXCursor_ObjCAtCatchStmt CXCursorKind = 217

	/** Objective-C's \@finally statement.
	 */
	CXCursor_ObjCAtFinallyStmt CXCursorKind = 218

	/** Objective-C's \@throw statement.
	 */
	CXCursor_ObjCAtThrowStmt CXCursorKind = 219

	/** Objective-C's \@synchronized statement.
	 */
	CXCursor_ObjCAtSynchronizedStmt CXCursorKind = 220

	/** Objective-C's autorelease pool statement.
	 */
	CXCursor_ObjCAutoreleasePoolStmt CXCursorKind = 221

	/** Objective-C's collection statement.
	 */
	CXCursor_ObjCForCollectionStmt CXCursorKind = 222

	/** C++'s catch statement.
	 */
	CXCursor_CXXCatchStmt CXCursorKind = 223

	/** C++'s try statement.
	 */
	CXCursor_CXXTryStmt CXCursorKind = 224

	/** C++'s for (* : *) statement.
	 */
	CXCursor_CXXForRangeStmt CXCursorKind = 225

	/** Windows Structured Exception Handling's try statement.
	 */
	CXCursor_SEHTryStmt CXCursorKind = 226

	/** Windows Structured Exception Handling's except statement.
	 */
	CXCursor_SEHExceptStmt CXCursorKind = 227

	/** Windows Structured Exception Handling's finally statement.
	 */
	CXCursor_SEHFinallyStmt CXCursorKind = 228

	/** A MS inline assembly statement extension.
	 */
	CXCursor_MSAsmStmt CXCursorKind = 229

	/** The null statement ";": C99 6.8.3p3.
	 *
	 * This cursor kind is used to describe the null statement.
	 */
	CXCursor_NullStmt CXCursorKind = 230

	/** Adaptor class for mixing declarations with statements and
	 * expressions.
	 */
	CXCursor_DeclStmt CXCursorKind = 231

	/** OpenMP parallel directive.
	 */
	CXCursor_OMPParallelDirective CXCursorKind = 232

	/** OpenMP SIMD directive.
	 */
	CXCursor_OMPSimdDirective CXCursorKind = 233

	/** OpenMP for directive.
	 */
	CXCursor_OMPForDirective CXCursorKind = 234

	/** OpenMP sections directive.
	 */
	CXCursor_OMPSectionsDirective CXCursorKind = 235

	/** OpenMP section directive.
	 */
	CXCursor_OMPSectionDirective CXCursorKind = 236

	/** OpenMP single directive.
	 */
	CXCursor_OMPSingleDirective CXCursorKind = 237

	/** OpenMP parallel for directive.
	 */
	CXCursor_OMPParallelForDirective CXCursorKind = 238

	/** OpenMP parallel sections directive.
	 */
	CXCursor_OMPParallelSectionsDirective CXCursorKind = 239

	/** OpenMP task directive.
	 */
	CXCursor_OMPTaskDirective CXCursorKind = 240

	/** OpenMP master directive.
	 */
	CXCursor_OMPMasterDirective CXCursorKind = 241

	/** OpenMP critical directive.
	 */
	CXCursor_OMPCriticalDirective CXCursorKind = 242

	/** OpenMP taskyield directive.
	 */
	CXCursor_OMPTaskyieldDirective CXCursorKind = 243

	/** OpenMP barrier directive.
	 */
	CXCursor_OMPBarrierDirective CXCursorKind = 244

	/** OpenMP taskwait directive.
	 */
	CXCursor_OMPTaskwaitDirective CXCursorKind = 245

	/** OpenMP flush directive.
	 */
	CXCursor_OMPFlushDirective CXCursorKind = 246

	/** Windows Structured Exception Handling's leave statement.
	 */
	CXCursor_SEHLeaveStmt CXCursorKind = 247

	/** OpenMP ordered directive.
	 */
	CXCursor_OMPOrderedDirective CXCursorKind = 248

	/** OpenMP atomic directive.
	 */
	CXCursor_OMPAtomicDirective CXCursorKind = 249

	/** OpenMP for SIMD directive.
	 */
	CXCursor_OMPForSimdDirective CXCursorKind = 250

	/** OpenMP parallel for SIMD directive.
	 */
	CXCursor_OMPParallelForSimdDirective CXCursorKind = 251

	/** OpenMP target directive.
	 */
	CXCursor_OMPTargetDirective CXCursorKind = 252

	/** OpenMP teams directive.
	 */
	CXCursor_OMPTeamsDirective CXCursorKind = 253

	/** OpenMP taskgroup directive.
	 */
	CXCursor_OMPTaskgroupDirective CXCursorKind = 254

	/** OpenMP cancellation point directive.
	 */
	CXCursor_OMPCancellationPointDirective CXCursorKind = 255

	/** OpenMP cancel directive.
	 */
	CXCursor_OMPCancelDirective CXCursorKind = 256

	/** OpenMP target data directive.
	 */
	CXCursor_OMPTargetDataDirective CXCursorKind = 257

	/** OpenMP taskloop directive.
	 */
	CXCursor_OMPTaskLoopDirective CXCursorKind = 258

	/** OpenMP taskloop simd directive.
	 */
	CXCursor_OMPTaskLoopSimdDirective CXCursorKind = 259

	/** OpenMP distribute directive.
	 */
	CXCursor_OMPDistributeDirective CXCursorKind = 260

	/** OpenMP target enter data directive.
	 */
	CXCursor_OMPTargetEnterDataDirective CXCursorKind = 261

	/** OpenMP target exit data directive.
	 */
	CXCursor_OMPTargetExitDataDirective CXCursorKind = 262

	/** OpenMP target parallel directive.
	 */
	CXCursor_OMPTargetParallelDirective CXCursorKind = 263

	/** OpenMP target parallel for directive.
	 */
	CXCursor_OMPTargetParallelForDirective CXCursorKind = 264

	/** OpenMP target update directive.
	 */
	CXCursor_OMPTargetUpdateDirective CXCursorKind = 265

	/** OpenMP distribute parallel for directive.
	 */
	CXCursor_OMPDistributeParallelForDirective CXCursorKind = 266

	/** OpenMP distribute parallel for simd directive.
	 */
	CXCursor_OMPDistributeParallelForSimdDirective CXCursorKind = 267

	/** OpenMP distribute simd directive.
	 */
	CXCursor_OMPDistributeSimdDirective CXCursorKind = 268

	/** OpenMP target parallel for simd directive.
	 */
	CXCursor_OMPTargetParallelForSimdDirective CXCursorKind = 269

	/** OpenMP target simd directive.
	 */
	CXCursor_OMPTargetSimdDirective CXCursorKind = 270

	/** OpenMP teams distribute directive.
	 */
	CXCursor_OMPTeamsDistributeDirective CXCursorKind = 271

	/** OpenMP teams distribute simd directive.
	 */
	CXCursor_OMPTeamsDistributeSimdDirective CXCursorKind = 272

	/** OpenMP teams distribute parallel for simd directive.
	 */
	CXCursor_OMPTeamsDistributeParallelForSimdDirective CXCursorKind = 273

	/** OpenMP teams distribute parallel for directive.
	 */
	CXCursor_OMPTeamsDistributeParallelForDirective CXCursorKind = 274

	/** OpenMP target teams directive.
	 */
	CXCursor_OMPTargetTeamsDirective CXCursorKind = 275

	/** OpenMP target teams distribute directive.
	 */
	CXCursor_OMPTargetTeamsDistributeDirective CXCursorKind = 276

	/** OpenMP target teams distribute parallel for directive.
	 */
	CXCursor_OMPTargetTeamsDistributeParallelForDirective CXCursorKind = 277

	/** OpenMP target teams distribute parallel for simd directive.
	 */
	CXCursor_OMPTargetTeamsDistributeParallelForSimdDirective CXCursorKind = 278

	/** OpenMP target teams distribute simd directive.
	 */
	CXCursor_OMPTargetTeamsDistributeSimdDirective CXCursorKind = 279

	/** C++2a std::bit_cast expression.
	 */
	CXCursor_BuiltinBitCastExpr CXCursorKind = 280

	/** OpenMP master taskloop directive.
	 */
	CXCursor_OMPMasterTaskLoopDirective CXCursorKind = 281

	/** OpenMP parallel master taskloop directive.
	 */
	CXCursor_OMPParallelMasterTaskLoopDirective CXCursorKind = 282

	/** OpenMP master taskloop simd directive.
	 */
	CXCursor_OMPMasterTaskLoopSimdDirective CXCursorKind = 283

	/** OpenMP parallel master taskloop simd directive.
	 */
	CXCursor_OMPParallelMasterTaskLoopSimdDirective CXCursorKind = 284

	/** OpenMP parallel master directive.
	 */
	CXCursor_OMPParallelMasterDirective CXCursorKind = 285

	/** OpenMP depobj directive.
	 */
	CXCursor_OMPDepobjDirective CXCursorKind = 286

	/** OpenMP scan directive.
	 */
	CXCursor_OMPScanDirective CXCursorKind = 287

	/** OpenMP tile directive.
	 */
	CXCursor_OMPTileDirective CXCursorKind = 288

	/** OpenMP canonical loop.
	 */
	CXCursor_OMPCanonicalLoop CXCursorKind = 289

	/** OpenMP interop directive.
	 */
	CXCursor_OMPInteropDirective CXCursorKind = 290

	/** OpenMP dispatch directive.
	 */
	CXCursor_OMPDispatchDirective CXCursorKind = 291

	/** OpenMP masked directive.
	 */
	CXCursor_OMPMaskedDirective CXCursorKind = 292

	/** OpenMP unroll directive.
	 */
	CXCursor_OMPUnrollDirective CXCursorKind = 293

	/** OpenMP metadirective directive.
	 */
	CXCursor_OMPMetaDirective CXCursorKind = 294

	/** OpenMP loop directive.
	 */
	CXCursor_OMPGenericLoopDirective CXCursorKind = 295

	/** OpenMP teams loop directive.
	 */
	CXCursor_OMPTeamsGenericLoopDirective CXCursorKind = 296

	/** OpenMP target teams loop directive.
	 */
	CXCursor_OMPTargetTeamsGenericLoopDirective CXCursorKind = 297

	/** OpenMP parallel loop directive.
	 */
	CXCursor_OMPParallelGenericLoopDirective CXCursorKind = 298

	/** OpenMP target parallel loop directive.
	 */
	CXCursor_OMPTargetParallelGenericLoopDirective CXCursorKind = 299

	/** OpenMP parallel masked directive.
	 */
	CXCursor_OMPParallelMaskedDirective CXCursorKind = 300

	/** OpenMP masked taskloop directive.
	 */
	CXCursor_OMPMaskedTaskLoopDirective CXCursorKind = 301

	/** OpenMP masked taskloop simd directive.
	 */
	CXCursor_OMPMaskedTaskLoopSimdDirective CXCursorKind = 302

	/** OpenMP parallel masked taskloop directive.
	 */
	CXCursor_OMPParallelMaskedTaskLoopDirective CXCursorKind = 303

	/** OpenMP parallel masked taskloop simd directive.
	 */
	CXCursor_OMPParallelMaskedTaskLoopSimdDirective CXCursorKind = 304

	/** OpenMP error directive.
	 */
	CXCursor_OMPErrorDirective CXCursorKind = 305

	/** OpenMP scope directive.
	 */
	CXCursor_OMPScopeDirective CXCursorKind = 306

	/** OpenMP reverse directive.
	 */
	CXCursor_OMPReverseDirective CXCursorKind = 307

	/** OpenMP interchange directive.
	 */
	CXCursor_OMPInterchangeDirective CXCursorKind = 308

	/** OpenMP assume directive.
	 */
	CXCursor_OMPAssumeDirective CXCursorKind = 309

	/** OpenACC Compute Construct.
	 */
	CXCursor_OpenACCComputeConstruct CXCursorKind = 320

	/** OpenACC Loop Construct.
	 */
	CXCursor_OpenACCLoopConstruct CXCursorKind = 321

	/** OpenACC Combined Constructs.
	 */
	CXCursor_OpenACCCombinedConstruct CXCursorKind = 322

	/** OpenACC data Construct.
	 */
	CXCursor_OpenACCDataConstruct CXCursorKind = 323

	/** OpenACC enter data Construct.
	 */
	CXCursor_OpenACCEnterDataConstruct CXCursorKind = 324

	/** OpenACC exit data Construct.
	 */
	CXCursor_OpenACCExitDataConstruct CXCursorKind = 325

	/** OpenACC host_data Construct.
	 */
	CXCursor_OpenACCHostDataConstruct CXCursorKind = 326

	/** OpenACC wait Construct.
	 */
	CXCursor_OpenACCWaitConstruct CXCursorKind = 327

	/** OpenACC init Construct.
	 */
	CXCursor_OpenACCInitConstruct CXCursorKind = 328

	/** OpenACC shutdown Construct.
	 */
	CXCursor_OpenACCShutdownConstruct CXCursorKind = 329

	/** OpenACC set Construct.
	 */
	CXCursor_OpenACCSetConstruct CXCursorKind = 330

	/** OpenACC update Construct.
	 */
	CXCursor_OpenACCUpdateConstruct CXCursorKind = 331

	CXCursor_LastStmt CXCursorKind = CXCursor_OpenACCUpdateConstruct

	/**
	 * Cursor that represents the translation unit itself.
	 *
	 * The translation unit cursor exists primarily to act as the root
	 * cursor for traversing the contents of a translation unit.
	 */
	CXCursor_TranslationUnit CXCursorKind = 350

	/* Attributes */
	CXCursor_FirstAttr CXCursorKind = 400
	/**
	 * An attribute whose specific kind is not exposed via this
	 * interface.
	 */
	CXCursor_UnexposedAttr CXCursorKind = 400

	CXCursor_IBActionAttr              CXCursorKind = 401
	CXCursor_IBOutletAttr              CXCursorKind = 402
	CXCursor_IBOutletCollectionAttr    CXCursorKind = 403
	CXCursor_CXXFinalAttr              CXCursorKind = 404
	CXCursor_CXXOverrideAttr           CXCursorKind = 405
	CXCursor_AnnotateAttr              CXCursorKind = 406
	CXCursor_AsmLabelAttr              CXCursorKind = 407
	CXCursor_PackedAttr                CXCursorKind = 408
	CXCursor_PureAttr                  CXCursorKind = 409
	CXCursor_ConstAttr                 CXCursorKind = 410
	CXCursor_NoDuplicateAttr           CXCursorKind = 411
	CXCursor_CUDAConstantAttr          CXCursorKind = 412
	CXCursor_CUDADeviceAttr            CXCursorKind = 413
	CXCursor_CUDAGlobalAttr            CXCursorKind = 414
	CXCursor_CUDAHostAttr              CXCursorKind = 415
	CXCursor_CUDASharedAttr            CXCursorKind = 416
	CXCursor_VisibilityAttr            CXCursorKind = 417
	CXCursor_DLLExport                 CXCursorKind = 418
	CXCursor_DLLImport                 CXCursorKind = 419
	CXCursor_NSReturnsRetained         CXCursorKind = 420
	CXCursor_NSReturnsNotRetained      CXCursorKind = 421
	CXCursor_NSReturnsAutoreleased     CXCursorKind = 422
	CXCursor_NSConsumesSelf            CXCursorKind = 423
	CXCursor_NSConsumed                CXCursorKind = 424
	CXCursor_ObjCException             CXCursorKind = 425
	CXCursor_ObjCNSObject              CXCursorKind = 426
	CXCursor_ObjCIndependentClass      CXCursorKind = 427
	CXCursor_ObjCPreciseLifetime       CXCursorKind = 428
	CXCursor_ObjCReturnsInnerPointer   CXCursorKind = 429
	CXCursor_ObjCRequiresSuper         CXCursorKind = 430
	CXCursor_ObjCRootClass             CXCursorKind = 431
	CXCursor_ObjCSubclassingRestricted CXCursorKind = 432
	CXCursor_ObjCExplicitProtocolImpl  CXCursorKind = 433
	CXCursor_ObjCDesignatedInitializer CXCursorKind = 434
	CXCursor_ObjCRuntimeVisible        CXCursorKind = 435
	CXCursor_ObjCBoxable               CXCursorKind = 436
	CXCursor_FlagEnum                  CXCursorKind = 437
	CXCursor_ConvergentAttr            CXCursorKind = 438
	CXCursor_WarnUnusedAttr            CXCursorKind = 439
	CXCursor_WarnUnusedResultAttr      CXCursorKind = 440
	CXCursor_AlignedAttr               CXCursorKind = 441
	CXCursor_LastAttr                  CXCursorKind = CXCursor_AlignedAttr

	/* Preprocessing */
	CXCursor_PreprocessingDirective CXCursorKind = 500
	CXCursor_MacroDefinition        CXCursorKind = 501
	CXCursor_MacroExpansion         CXCursorKind = 502
	CXCursor_MacroInstantiation     CXCursorKind = CXCursor_MacroExpansion
	CXCursor_InclusionDirective     CXCursorKind = 503
	CXCursor_FirstPreprocessing     CXCursorKind = CXCursor_PreprocessingDirective
	CXCursor_LastPreprocessing      CXCursorKind = CXCursor_InclusionDirective

	/* Extra Declarations */
	/**
	 * A module import declaration.
	 */
	CXCursor_ModuleImportDecl      CXCursorKind = 600
	CXCursor_TypeAliasTemplateDecl CXCursorKind = 601
	/**
	 * A static_assert or _Static_assert node
	 */
	CXCursor_StaticAssert CXCursorKind = 602
	/**
	 * a friend declaration.
	 */
	CXCursor_FriendDecl CXCursorKind = 603
	/**
	 * a concept declaration.
	 */
	CXCursor_ConceptDecl CXCursorKind = 604

	CXCursor_FirstExtraDecl CXCursorKind = CXCursor_ModuleImportDecl
	CXCursor_LastExtraDecl  CXCursorKind = CXCursor_ConceptDecl

	/**
	 * A code completion overload candidate.
	 */
	CXCursor_OverloadCandidate CXCursorKind = 700
)

type CXChildVisitResult uint32

const (
	/**
	 * Terminates the cursor traversal.
	 */
	CXChildVisit_Break CXChildVisitResult = iota
	/**
	 * Continues the cursor traversal with the next sibling of
	 * the cursor just visited, without visiting its children.
	 */
	CXChildVisit_Continue
	/**
	 * Recursively traverse the children of this cursor, using
	 * the same visitor and client data.
	 */
	CXChildVisit_Recurse
)

type CXLinkageKind uint32

const (

	/** This value indicates that no linkage information is available
	 * for a provided CXCursor. */
	CXLinkage_Invalid CXLinkageKind = iota
	/**
	 * This is the linkage for variables, parameters, and so on that
	 *  have automatic storage.  This covers normal (non-extern) local variables.
	 */
	CXLinkage_NoLinkage
	/** This is the linkage for static variables and static functions. */
	CXLinkage_Internal
	/** This is the linkage for entities with external linkage that live
	 * in C++ anonymous namespaces.*/
	CXLinkage_UniqueExternal
	/** This is the linkage for entities with true, external linkage. */
	CXLinkage_External
)

type CXVisibilityKind uint32

const (
	/** This value indicates that no visibility information is available
	 * for a provided CXCursor. */
	CXVisibility_Invalid CXVisibilityKind = iota

	/** Symbol not seen by the linker. */
	CXVisibility_Hidden
	/** Symbol seen by the linker but resolves to a symbol inside this object. */
	CXVisibility_Protected
	/** Symbol seen by the linker and acts like a normal symbol. */
	CXVisibility_Default
)

type CXAvailabilityKind uint32

const (
	/**
	 * The entity is available.
	 */
	CXAvailability_Available CXAvailabilityKind = iota
	/**
	 * The entity is available, but has been deprecated (and its use is
	 * not recommended).
	 */
	CXAvailability_Deprecated
	/**
	 * The entity is not available; any use of it will be an error.
	 */
	CXAvailability_NotAvailable
	/**
	 * The entity is available, but not accessible; any use of it will be
	 * an error.
	 */
	CXAvailability_NotAccessible
)

type CXLanguageKind uint32

const (
	CXLanguage_Invalid CXLanguageKind = iota
	CXLanguage_C
	CXLanguage_ObjC
	CXLanguage_CPlusPlus
)

type CXTLSKind uint32

const (
	CXTLS_None CXTLSKind = iota
	CXTLS_Dynamic
	CXTLS_Static
)

type CXTypeKind uint32

const (
	/**
	 * Represents an invalid type (e.g., where no type is available).
	 */
	CXType_Invalid CXTypeKind = 0

	/**
	 * A type whose specific kind is not exposed via this
	 * interface.
	 */
	CXType_Unexposed CXTypeKind = 1

	/* Builtin types */
	CXType_Void         CXTypeKind = 2
	CXType_Bool         CXTypeKind = 3
	CXType_Char_U       CXTypeKind = 4
	CXType_UChar        CXTypeKind = 5
	CXType_Char16       CXTypeKind = 6
	CXType_Char32       CXTypeKind = 7
	CXType_UShort       CXTypeKind = 8
	CXType_UInt         CXTypeKind = 9
	CXType_ULong        CXTypeKind = 10
	CXType_ULongLong    CXTypeKind = 11
	CXType_UInt128      CXTypeKind = 12
	CXType_Char_S       CXTypeKind = 13
	CXType_SChar        CXTypeKind = 14
	CXType_WChar        CXTypeKind = 15
	CXType_Short        CXTypeKind = 16
	CXType_Int          CXTypeKind = 17
	CXType_Long         CXTypeKind = 18
	CXType_LongLong     CXTypeKind = 19
	CXType_Int128       CXTypeKind = 20
	CXType_Float        CXTypeKind = 21
	CXType_Double       CXTypeKind = 22
	CXType_LongDouble   CXTypeKind = 23
	CXType_NullPtr      CXTypeKind = 24
	CXType_Overload     CXTypeKind = 25
	CXType_Dependent    CXTypeKind = 26
	CXType_ObjCId       CXTypeKind = 27
	CXType_ObjCClass    CXTypeKind = 28
	CXType_ObjCSel      CXTypeKind = 29
	CXType_Float128     CXTypeKind = 30
	CXType_Half         CXTypeKind = 31
	CXType_Float16      CXTypeKind = 32
	CXType_ShortAccum   CXTypeKind = 33
	CXType_Accum        CXTypeKind = 34
	CXType_LongAccum    CXTypeKind = 35
	CXType_UShortAccum  CXTypeKind = 36
	CXType_UAccum       CXTypeKind = 37
	CXType_ULongAccum   CXTypeKind = 38
	CXType_BFloat16     CXTypeKind = 39
	CXType_Ibm128       CXTypeKind = 40
	CXType_FirstBuiltin CXTypeKind = CXType_Void
	CXType_LastBuiltin  CXTypeKind = CXType_Ibm128

	CXType_Complex             CXTypeKind = 100
	CXType_Pointer             CXTypeKind = 101
	CXType_BlockPointer        CXTypeKind = 102
	CXType_LValueReference     CXTypeKind = 103
	CXType_RValueReference     CXTypeKind = 104
	CXType_Record              CXTypeKind = 105
	CXType_Enum                CXTypeKind = 106
	CXType_Typedef             CXTypeKind = 107
	CXType_ObjCInterface       CXTypeKind = 108
	CXType_ObjCObjectPointer   CXTypeKind = 109
	CXType_FunctionNoProto     CXTypeKind = 110
	CXType_FunctionProto       CXTypeKind = 111
	CXType_ConstantArray       CXTypeKind = 112
	CXType_Vector              CXTypeKind = 113
	CXType_IncompleteArray     CXTypeKind = 114
	CXType_VariableArray       CXTypeKind = 115
	CXType_DependentSizedArray CXTypeKind = 116
	CXType_MemberPointer       CXTypeKind = 117
	CXType_Auto                CXTypeKind = 118

	/**
	 * Represents a type that was referred to using an elaborated type keyword.
	 *
	 * E.g., struct S, or via a qualified name, e.g., N::M::type, or both.
	 */
	CXType_Elaborated CXTypeKind = 119

	/* OpenCL PipeType. */
	CXType_Pipe CXTypeKind = 120

	/* OpenCL builtin types. */
	CXType_OCLImage1dRO               CXTypeKind = 121
	CXType_OCLImage1dArrayRO          CXTypeKind = 122
	CXType_OCLImage1dBufferRO         CXTypeKind = 123
	CXType_OCLImage2dRO               CXTypeKind = 124
	CXType_OCLImage2dArrayRO          CXTypeKind = 125
	CXType_OCLImage2dDepthRO          CXTypeKind = 126
	CXType_OCLImage2dArrayDepthRO     CXTypeKind = 127
	CXType_OCLImage2dMSAARO           CXTypeKind = 128
	CXType_OCLImage2dArrayMSAARO      CXTypeKind = 129
	CXType_OCLImage2dMSAADepthRO      CXTypeKind = 130
	CXType_OCLImage2dArrayMSAADepthRO CXTypeKind = 131
	CXType_OCLImage3dRO               CXTypeKind = 132
	CXType_OCLImage1dWO               CXTypeKind = 133
	CXType_OCLImage1dArrayWO          CXTypeKind = 134
	CXType_OCLImage1dBufferWO         CXTypeKind = 135
	CXType_OCLImage2dWO               CXTypeKind = 136
	CXType_OCLImage2dArrayWO          CXTypeKind = 137
	CXType_OCLImage2dDepthWO          CXTypeKind = 138
	CXType_OCLImage2dArrayDepthWO     CXTypeKind = 139
	CXType_OCLImage2dMSAAWO           CXTypeKind = 140
	CXType_OCLImage2dArrayMSAAWO      CXTypeKind = 141
	CXType_OCLImage2dMSAADepthWO      CXTypeKind = 142
	CXType_OCLImage2dArrayMSAADepthWO CXTypeKind = 143
	CXType_OCLImage3dWO               CXTypeKind = 144
	CXType_OCLImage1dRW               CXTypeKind = 145
	CXType_OCLImage1dArrayRW          CXTypeKind = 146
	CXType_OCLImage1dBufferRW         CXTypeKind = 147
	CXType_OCLImage2dRW               CXTypeKind = 148
	CXType_OCLImage2dArrayRW          CXTypeKind = 149
	CXType_OCLImage2dDepthRW          CXTypeKind = 150
	CXType_OCLImage2dArrayDepthRW     CXTypeKind = 151
	CXType_OCLImage2dMSAARW           CXTypeKind = 152
	CXType_OCLImage2dArrayMSAARW      CXTypeKind = 153
	CXType_OCLImage2dMSAADepthRW      CXTypeKind = 154
	CXType_OCLImage2dArrayMSAADepthRW CXTypeKind = 155
	CXType_OCLImage3dRW               CXTypeKind = 156
	CXType_OCLSampler                 CXTypeKind = 157
	CXType_OCLEvent                   CXTypeKind = 158
	CXType_OCLQueue                   CXTypeKind = 159
	CXType_OCLReserveID               CXTypeKind = 160

	CXType_ObjCObject    CXTypeKind = 161
	CXType_ObjCTypeParam CXTypeKind = 162
	CXType_Attributed    CXTypeKind = 163

	CXType_OCLIntelSubgroupAVCMcePayload                        CXTypeKind = 164
	CXType_OCLIntelSubgroupAVCImePayload                        CXTypeKind = 165
	CXType_OCLIntelSubgroupAVCRefPayload                        CXTypeKind = 166
	CXType_OCLIntelSubgroupAVCSicPayload                        CXTypeKind = 167
	CXType_OCLIntelSubgroupAVCMceResult                         CXTypeKind = 168
	CXType_OCLIntelSubgroupAVCImeResult                         CXTypeKind = 169
	CXType_OCLIntelSubgroupAVCRefResult                         CXTypeKind = 170
	CXType_OCLIntelSubgroupAVCSicResult                         CXTypeKind = 171
	CXType_OCLIntelSubgroupAVCImeResultSingleReferenceStreamout CXTypeKind = 172
	CXType_OCLIntelSubgroupAVCImeResultDualReferenceStreamout   CXTypeKind = 173
	CXType_OCLIntelSubgroupAVCImeSingleReferenceStreamin        CXTypeKind = 174
	CXType_OCLIntelSubgroupAVCImeDualReferenceStreamin          CXTypeKind = 175

	/* Old aliases for AVC OpenCL extension types. */
	CXType_OCLIntelSubgroupAVCImeResultSingleRefStreamout CXTypeKind = 172
	CXType_OCLIntelSubgroupAVCImeResultDualRefStreamout   CXTypeKind = 173
	CXType_OCLIntelSubgroupAVCImeSingleRefStreamin        CXTypeKind = 174
	CXType_OCLIntelSubgroupAVCImeDualRefStreamin          CXTypeKind = 175

	CXType_ExtVector        CXTypeKind = 176
	CXType_Atomic           CXTypeKind = 177
	CXType_BTFTagAttributed CXTypeKind = 178

	/* HLSL Types */
	CXType_HLSLResource           CXTypeKind = 179
	CXType_HLSLAttributedResource CXTypeKind = 180
)
