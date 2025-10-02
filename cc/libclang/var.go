package libclang

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

func init() { cc.Open("libclang.*") }

type uptr = unsafe.Pointer
type iptr = uintptr

var (
	// #region BuildSystem.h
	clang_getBuildSessionTimestamp                   = cc.DlFunc[func() cc.ULonglong, cc.ULonglong]{Name: "clang_getBuildSessionTimestamp"}
	clang_VirtualFileOverlay_create                  = cc.DlFunc[func(uint32) CXVirtualFileOverlay, CXVirtualFileOverlay]{Name: "clang_VirtualFileOverlay_create"}
	clang_VirtualFileOverlay_addFileMapping          = cc.DlFunc[func(CXVirtualFileOverlay, cc.String, cc.String) CXErrorCode, CXErrorCode]{Name: "clang_VirtualFileOverlay_addFileMapping"}
	clang_VirtualFileOverlay_setCaseSensitivity      = cc.DlFunc[func(CXVirtualFileOverlay, int32) CXErrorCode, CXErrorCode]{Name: "clang_VirtualFileOverlay_setCaseSensitivity"}
	clang_VirtualFileOverlay_writeToBuffer           = cc.DlFunc[func(CXVirtualFileOverlay, uint32, *cc.String, *uint32) CXErrorCode, CXErrorCode]{Name: "clang_VirtualFileOverlay_writeToBuffer"}
	clang_free                                       = cc.DlFunc[func(uptr), cc.Void]{Name: "clang_free"}
	clang_VirtualFileOverlay_dispose                 = cc.DlFunc[func(CXVirtualFileOverlay), cc.Void]{Name: "clang_VirtualFileOverlay_dispose"}
	clang_ModuleMapDescriptor_create                 = cc.DlFunc[func(uint32) CXModuleMapDescriptor, CXModuleMapDescriptor]{Name: "clang_ModuleMapDescriptor_create"}
	clang_ModuleMapDescriptor_setFrameworkModuleName = cc.DlFunc[func(CXModuleMapDescriptor, cc.String) CXErrorCode, CXErrorCode]{Name: "clang_ModuleMapDescriptor_setFrameworkModuleName"}
	clang_ModuleMapDescriptor_setUmbrellaHeader      = cc.DlFunc[func(CXModuleMapDescriptor, cc.String) CXErrorCode, CXErrorCode]{Name: "clang_ModuleMapDescriptor_setUmbrellaHeader"}
	clang_ModuleMapDescriptor_writeToBuffer          = cc.DlFunc[func(CXModuleMapDescriptor, uint32, *cc.String, *uint32) CXErrorCode, CXErrorCode]{Name: "clang_ModuleMapDescriptor_writeToBuffer"}
	clang_ModuleMapDescriptor_dispose                = cc.DlFunc[func(CXModuleMapDescriptor), cc.Void]{Name: "clang_ModuleMapDescriptor_dispose"}
	// #endregion

	// #region CXCompilationDatabase.h
	clang_CompilationDatabase_fromDirectory         = cc.DlFunc[func(cc.String, *CXCompilationDatabase_Error) CXCompilationDatabase, CXCompilationDatabase]{Name: "clang_CompilationDatabase_fromDirectory"}
	clang_CompilationDatabase_dispose               = cc.DlFunc[func(CXCompilationDatabase), cc.Void]{Name: "clang_CompilationDatabase_dispose"}
	clang_CompilationDatabase_getCompileCommands    = cc.DlFunc[func(CXCompilationDatabase, cc.String) CXCompileCommands, CXCompileCommands]{Name: "clang_CompilationDatabase_getCompileCommands"}
	clang_CompilationDatabase_getAllCompileCommands = cc.DlFunc[func(CXCompilationDatabase) CXCompileCommands, CXCompileCommands]{Name: "clang_CompilationDatabase_getAllCompileCommands"}
	clang_CompileCommands_dispose                   = cc.DlFunc[func(CXCompileCommands), cc.Void]{Name: "clang_CompileCommands_dispose"}
	clang_CompileCommands_getSize                   = cc.DlFunc[func(CXCompileCommands) uint32, uint32]{Name: "clang_CompileCommands_getSize"}
	clang_CompileCommands_getCommand                = cc.DlFunc[func(CXCompileCommands, uint32) CXCompileCommand, CXCompileCommand]{Name: "clang_CompileCommands_getCommand"}
	clang_CompileCommand_getDirectory               = cc.DlFunc[func(CXCompileCommand) CXString, CXString]{Name: "clang_CompileCommand_getDirectory"}
	clang_CompileCommand_getFilename                = cc.DlFunc[func(CXCompileCommand) CXString, CXString]{Name: "clang_CompileCommand_getFilename"}
	clang_CompileCommand_getNumArgs                 = cc.DlFunc[func(CXCompileCommand) uint32, uint32]{Name: "clang_CompileCommand_getNumArgs"}
	clang_CompileCommand_getArg                     = cc.DlFunc[func(CXCompileCommand, uint32) CXString, CXString]{Name: "clang_CompileCommand_getArg"}
	clang_CompileCommand_getNumMappedSources        = cc.DlFunc[func(CXCompileCommand) uint32, uint32]{Name: "clang_CompileCommand_getNumMappedSources"}
	clang_CompileCommand_getMappedSourcePath        = cc.DlFunc[func(CXCompileCommand, uint32) CXString, CXString]{Name: "clang_CompileCommand_getMappedSourcePath"}
	clang_CompileCommand_getMappedSourceContent     = cc.DlFunc[func(CXCompileCommand, uint32) CXString, CXString]{Name: "clang_CompileCommand_getMappedSourceContent"}
	// #endregion

	// #region CXDiagnostic.h
	clang_getNumDiagnosticsInSet          = cc.DlFunc[func(CXDiagnosticSet) uint32, uint32]{Name: "clang_getNumDiagnosticsInSet"}
	clang_getDiagnosticInSet              = cc.DlFunc[func(CXDiagnosticSet, uint32) CXDiagnostic, CXDiagnostic]{Name: "clang_getDiagnosticInSet"}
	clang_loadDiagnostics                 = cc.DlFunc[func(cc.String, *CXLoadDiag_Error, *CXString) CXDiagnosticSet, CXDiagnosticSet]{Name: "clang_loadDiagnostics"} // 修正：file→cc.String
	clang_disposeDiagnosticSet            = cc.DlFunc[func(CXDiagnosticSet), cc.Void]{Name: "clang_disposeDiagnosticSet"}
	clang_getChildDiagnostics             = cc.DlFunc[func(CXDiagnostic) CXDiagnosticSet, CXDiagnosticSet]{Name: "clang_getChildDiagnostics"}
	clang_disposeDiagnostic               = cc.DlFunc[func(CXDiagnostic), cc.Void]{Name: "clang_disposeDiagnostic"}
	clang_formatDiagnostic                = cc.DlFunc[func(CXDiagnostic, CXDiagnosticDisplayOptions) CXString, CXString]{Name: "clang_formatDiagnostic"}
	clang_defaultDiagnosticDisplayOptions = cc.DlFunc[func() CXDiagnosticDisplayOptions, CXDiagnosticDisplayOptions]{Name: "clang_defaultDiagnosticDisplayOptions"}
	clang_getDiagnosticSeverity           = cc.DlFunc[func(CXDiagnostic) CXDiagnosticSeverity, CXDiagnosticSeverity]{Name: "clang_getDiagnosticSeverity"}
	clang_getDiagnosticLocation           = cc.DlFunc[func(CXDiagnostic) CXSourceLocation, CXSourceLocation]{Name: "clang_getDiagnosticLocation"}
	clang_getDiagnosticSpelling           = cc.DlFunc[func(CXDiagnostic) CXString, CXString]{Name: "clang_getDiagnosticSpelling"}
	clang_getDiagnosticOption             = cc.DlFunc[func(CXDiagnostic, *CXString) CXString, CXString]{Name: "clang_getDiagnosticOption"}
	clang_getDiagnosticCategory           = cc.DlFunc[func(CXDiagnostic) uint32, uint32]{Name: "clang_getDiagnosticCategory"}
	clang_getDiagnosticCategoryName       = cc.DlFunc[func(uint32) CXString, CXString]{Name: "clang_getDiagnosticCategoryName"}
	clang_getDiagnosticCategoryText       = cc.DlFunc[func(CXDiagnostic) CXString, CXString]{Name: "clang_getDiagnosticCategoryText"}
	clang_getDiagnosticNumRanges          = cc.DlFunc[func(CXDiagnostic) uint32, uint32]{Name: "clang_getDiagnosticNumRanges"}
	clang_getDiagnosticRange              = cc.DlFunc[func(CXDiagnostic, uint32) CXSourceRange, CXSourceRange]{Name: "clang_getDiagnosticRange"}
	clang_getDiagnosticNumFixIts          = cc.DlFunc[func(CXDiagnostic) uint32, uint32]{Name: "clang_getDiagnosticNumFixIts"}
	clang_getDiagnosticFixIt              = cc.DlFunc[func(CXDiagnostic, uint32, *CXSourceRange) CXString, CXString]{Name: "clang_getDiagnosticFixIt"}
	// #endregion

	clang_createIndex                             = cc.DlFunc[func(int32, int32) CXIndex, CXIndex]{Name: "clang_createIndex"}
	clang_disposeIndex                            = cc.DlFunc[func(CXIndex), cc.Void]{Name: "clang_disposeIndex"}
	clang_createIndexWithOptions                  = cc.DlFunc[func(*CXIndexOptions) CXIndex, CXIndex]{Name: "clang_createIndexWithOptions"}
	clang_CXIndex_setGlobalOptions                = cc.DlFunc[func(CXIndex, CXGlobalOptFlags), cc.Void]{Name: "clang_CXIndex_setGlobalOptions"}
	clang_CXIndex_getGlobalOptions                = cc.DlFunc[func(CXIndex) CXGlobalOptFlags, CXGlobalOptFlags]{Name: "clang_CXIndex_getGlobalOptions"}
	clang_CXIndex_setInvocationEmissionPathOption = cc.DlFunc[func(CXIndex, cc.String), cc.Void]{Name: "clang_CXIndex_setInvocationEmissionPathOption"}
	clang_isFileMultipleIncludeGuarded            = cc.DlFunc[func(CXIndex, CXFile) uint32, uint32]{Name: "clang_isFileMultipleIncludeGuarded"}
	clang_getFile                                 = cc.DlFunc[func(CXTranslationUnit, cc.String) CXFile, CXFile]{Name: "clang_getFile"}
	clang_getFileContents                         = cc.DlFunc[func(CXTranslationUnit, CXFile, *cc.SizeT) cc.String, cc.String]{Name: "clang_getFileContents"}
	clang_getLocation                             = cc.DlFunc[func(CXTranslationUnit, CXFile, uint32, uint32) CXSourceLocation, CXSourceLocation]{Name: "clang_getLocation"}
	clang_getLocationForOffset                    = cc.DlFunc[func(CXTranslationUnit, CXFile, uint32) CXSourceLocation, CXSourceLocation]{Name: "clang_getLocationForOffset"}
	clang_getSkippedRanges                        = cc.DlFunc[func(CXTranslationUnit, CXFile) *CXSourceRangeList, *CXSourceRangeList]{Name: "clang_getSkippedRanges"}
	clang_getAllSkippedRanges                     = cc.DlFunc[func(CXTranslationUnit) *CXSourceRangeList, *CXSourceRangeList]{Name: "clang_getAllSkippedRanges"}
	clang_getNumDiagnostics                       = cc.DlFunc[func(CXTranslationUnit) uint32, uint32]{Name: "clang_getNumDiagnostics"}

	clang_parseTranslationUnit   = cc.DlFunc[func(CXIndex, cc.String, cc.Strings, int32, *CXUnsavedFile, int32, CXTranslationUnit_Flags) CXTranslationUnit, CXTranslationUnit]{Name: "clang_parseTranslationUnit"}
	clang_disposeTranslationUnit = cc.DlFunc[func(CXTranslationUnit), cc.Void]{Name: "clang_disposeTranslationUnit"}

	clang_getNullCursor            = cc.DlFunc[func() CXCursor, CXCursor]{Name: "clang_getNullCursor"}
	clang_getTranslationUnitCursor = cc.DlFunc[func(CXTranslationUnit) CXCursor, CXCursor]{Name: "clang_getTranslationUnitCursor"}
	clang_equalCursors             = cc.DlFunc[func(CXCursor, CXCursor) uint32, uint32]{Name: "clang_equalCursors"}
	clang_Cursor_isNull            = cc.DlFunc[func(CXCursor) int32, int32]{Name: "clang_Cursor_isNull"}
	clang_hashCursor               = cc.DlFunc[func(CXCursor) uint32, uint32]{Name: "clang_hashCursor"}
	clang_getCursorKind            = cc.DlFunc[func(CXCursor) CXCursorKind, CXCursorKind]{Name: "clang_getCursorKind"}
	clang_isDeclaration            = cc.DlFunc[func(CXCursorKind) uint32, uint32]{Name: "clang_isDeclaration"}
	clang_isInvalidDeclaration     = cc.DlFunc[func(CXCursor) uint32, uint32]{Name: "clang_isInvalidDeclaration"}
	clang_isReference              = cc.DlFunc[func(CXCursorKind) uint32, uint32]{Name: "clang_isReference"}
	clang_isExpression             = cc.DlFunc[func(CXCursorKind) uint32, uint32]{Name: "clang_isExpression"}

	clang_isStatement                      = cc.DlFunc[func(CXCursorKind) uint32, uint32]{Name: "clang_isStatement"}
	clang_isAttribute                      = cc.DlFunc[func(CXCursorKind) uint32, uint32]{Name: "clang_isAttribute"}
	clang_Cursor_hasAttrs                  = cc.DlFunc[func(CXCursor) uint32, uint32]{Name: "clang_Cursor_hasAttrs"}
	clang_isInvalid                        = cc.DlFunc[func(CXCursorKind) uint32, uint32]{Name: "clang_isInvalid"}
	clang_isTranslationUnit                = cc.DlFunc[func(CXCursorKind) uint32, uint32]{Name: "clang_isTranslationUnit"}
	clang_isPreprocessing                  = cc.DlFunc[func(CXCursorKind) uint32, uint32]{Name: "clang_isPreprocessing"}
	clang_isUnexposed                      = cc.DlFunc[func(CXCursorKind) uint32, uint32]{Name: "clang_isUnexposed"}
	clang_getCursorLinkage                 = cc.DlFunc[func(CXCursor) CXLinkageKind, CXLinkageKind]{Name: "clang_getCursorLinkage"}
	clang_getCursorVisibility              = cc.DlFunc[func(CXCursor) CXVisibilityKind, CXVisibilityKind]{Name: "clang_getCursorVisibility"}
	clang_getCursorAvailability            = cc.DlFunc[func(CXCursor) CXAvailabilityKind, CXAvailabilityKind]{Name: "clang_getCursorAvailability"}
	clang_getCursorPlatformAvailability    = cc.DlFunc[func(CXCursor, *int32, *CXString, *int32, *CXString, *CXPlatformAvailability, int32) int32, int32]{Name: "clang_getCursorPlatformAvailability"}
	clang_disposeCXPlatformAvailability    = cc.DlFunc[func(*CXPlatformAvailability), cc.Void]{Name: "clang_disposeCXPlatformAvailability"}
	clang_Cursor_getVarDeclInitializer     = cc.DlFunc[func(CXCursor) CXCursor, CXCursor]{Name: "clang_Cursor_getVarDeclInitializer"}
	clang_Cursor_hasVarDeclGlobalStorage   = cc.DlFunc[func(CXCursor) int32, int32]{Name: "clang_Cursor_hasVarDeclGlobalStorage"}
	clang_Cursor_hasVarDeclExternalStorage = cc.DlFunc[func(CXCursor) int32, int32]{Name: "clang_Cursor_hasVarDeclExternalStorage"}
	clang_getCursorLanguage                = cc.DlFunc[func(CXCursor) CXLanguageKind, CXLanguageKind]{Name: "clang_getCursorLanguage"}
	clang_getCursorTLSKind                 = cc.DlFunc[func(CXCursor) CXTLSKind, CXTLSKind]{Name: "clang_getCursorTLSKind"}
	clang_Cursor_getTranslationUnit        = cc.DlFunc[func(CXCursor) CXTranslationUnit, CXTranslationUnit]{Name: "clang_Cursor_getTranslationUnit"}
	clang_createCXCursorSet                = cc.DlFunc[func() CXCursorSet, CXCursorSet]{Name: "clang_createCXCursorSet"}
	clang_disposeCXCursorSet               = cc.DlFunc[func(CXCursorSet), cc.Void]{Name: "clang_disposeCXCursorSet"}
	clang_CXCursorSet_contains             = cc.DlFunc[func(CXCursorSet, CXCursor) uint32, uint32]{Name: "clang_CXCursorSet_contains"}
	clang_CXCursorSet_insert               = cc.DlFunc[func(CXCursorSet, CXCursor) uint32, uint32]{Name: "clang_CXCursorSet_insert"}
	clang_getCursorSemanticParent          = cc.DlFunc[func(CXCursor) CXCursor, CXCursor]{Name: "clang_getCursorSemanticParent"}
	clang_getCursorLexicalParent           = cc.DlFunc[func(CXCursor) CXCursor, CXCursor]{Name: "clang_getCursorLexicalParent"}
	clang_getOverriddenCursors             = cc.DlFunc[func(CXCursor, **CXCursor, *uint32), cc.Void]{Name: "clang_getOverriddenCursors"}
	clang_disposeOverriddenCursors         = cc.DlFunc[func(*CXCursor), cc.Void]{Name: "clang_disposeOverriddenCursors"}
	clang_getIncludedFile                  = cc.DlFunc[func(CXCursor) CXFile, CXFile]{Name: "clang_getIncludedFile"}

	clang_visitChildren     = cc.DlFunc[func(CXCursor, uptr, uptr) CXChildVisitResult, CXChildVisitResult]{Name: "clang_visitChildren"}
	clang_getCursorSpelling = cc.DlFunc[func(CXCursor) CXString, CXString]{Name: "clang_getCursorSpelling"}

	clang_getCursorType  = cc.DlFunc[func(CXCursor) CXType, CXType]{Name: "clang_getCursorType"}
	clang_Type_getSizeOf = cc.DlFunc[func(CXType) cc.Longlong, cc.Longlong]{Name: "clang_Type_getSizeOf"}

	clang_getCString    = cc.DlFunc[func(CXString) cc.String, cc.String]{Name: "clang_getCString"}
	clang_disposeString = cc.DlFunc[func(CXString), cc.Void]{Name: "clang_disposeString"}
)
