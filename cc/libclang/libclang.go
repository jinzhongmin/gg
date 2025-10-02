package libclang

import (
	"github.com/jinzhongmin/gg/cc"
)

func addr[T any](v T) *T { return &v }

// #region BuildSystem.h

func GetBuildSessionTimestamp() uint64 {
	return uint64(clang_getBuildSessionTimestamp.Fn()())
}

type CXVirtualFileOverlay iptr

func CreateVirtualFileOverlay(opt uint32) CXVirtualFileOverlay {
	return clang_VirtualFileOverlay_create.Fn()(opt)
}
func (overlay CXVirtualFileOverlay) AddFileMapping(virtualPath, realPath string) CXErrorCode {
	cVirtualPath := cc.NewString(virtualPath)
	cRealPath := cc.NewString(realPath)
	defer cVirtualPath.Free()
	defer cRealPath.Free()
	return clang_VirtualFileOverlay_addFileMapping.Fn()(overlay, cVirtualPath, cRealPath)
}
func (overlay CXVirtualFileOverlay) SetCaseSensitivity(caseSensitive bool) CXErrorCode {
	cs := int32(0)
	if caseSensitive {
		cs = 1
	}
	return clang_VirtualFileOverlay_setCaseSensitivity.Fn()(overlay, cs)
}
func (overlay CXVirtualFileOverlay) WriteToBuffer(options uint32) (code CXErrorCode, out string) {
	var outBuffer cc.String
	var outSize uint32
	code = clang_VirtualFileOverlay_writeToBuffer.Fn()(overlay, options, &outBuffer, &outSize)
	if outSize > 0 {
		defer outBuffer.Free()
		out = outBuffer.String(uint64(outSize))
	}
	return
}
func (overlay CXVirtualFileOverlay) Dispose() {
	clang_VirtualFileOverlay_dispose.Fn()(overlay)
}
func Free(buf uptr) { clang_free.Fn()(buf) }

type CXModuleMapDescriptor iptr

func CreateModuleMapDescriptor(options uint32) CXModuleMapDescriptor {
	return clang_ModuleMapDescriptor_create.Fn()(options)
}
func (desc CXModuleMapDescriptor) SetFrameworkModuleName(name string) CXErrorCode {
	cName := cc.NewString(name)
	defer cName.Free()
	return clang_ModuleMapDescriptor_setFrameworkModuleName.Fn()(desc, cName)
}

func (desc CXModuleMapDescriptor) SetUmbrellaHeader(name string) CXErrorCode {
	cName := cc.NewString(name)
	defer cName.Free()
	return clang_ModuleMapDescriptor_setUmbrellaHeader.Fn()(desc, cName)
}
func (desc CXModuleMapDescriptor) WriteToBuffer(options uint32) (code CXErrorCode, out string) {
	var outBuffer cc.String
	var outSize uint32
	code = clang_ModuleMapDescriptor_writeToBuffer.Fn()(desc, options, &outBuffer, &outSize)
	if outSize > 0 {
		defer outBuffer.Free()
		out = outBuffer.String(uint64(outSize))
	}
	return
}
func (desc CXModuleMapDescriptor) Dispose() {
	clang_ModuleMapDescriptor_dispose.Fn()(desc)
}

// #endregion

// #region CXCompilationDatabase.h

type CXCompilationDatabase iptr
type CXCompileCommands iptr
type CXCompileCommand iptr

func CreateCompilationDatabaseFromDirectory(buildDir string) (db CXCompilationDatabase, errCode CXCompilationDatabase_Error) {
	cBuildDir := cc.NewString(buildDir)
	defer cBuildDir.Free()
	db = clang_CompilationDatabase_fromDirectory.Fn()(cBuildDir, &errCode)
	return
}
func (db CXCompilationDatabase) Dispose() { clang_CompilationDatabase_dispose.Fn()(db) }
func (db CXCompilationDatabase) GetCompileCommands(completeFileName string) CXCompileCommands {
	cFileName := cc.NewString(completeFileName)
	defer cFileName.Free()
	return clang_CompilationDatabase_getCompileCommands.Fn()(db, cFileName)
}
func (db CXCompilationDatabase) GetAllCompileCommands() CXCompileCommands {
	return clang_CompilationDatabase_getAllCompileCommands.Fn()(db)
}
func (cmds CXCompileCommands) Dispose()        { clang_CompileCommands_dispose.Fn()(cmds) }
func (cmds CXCompileCommands) GetSize() uint32 { return clang_CompileCommands_getSize.Fn()(cmds) }
func (cmds CXCompileCommands) GetCommand(i uint32) CXCompileCommand {
	return clang_CompileCommands_getCommand.Fn()(cmds, i)
}
func (cmd CXCompileCommand) GetDirectory() *CXString {
	return addr(clang_CompileCommand_getDirectory.Fn()(cmd))
}
func (cmd CXCompileCommand) GetFilename() *CXString {
	return addr(clang_CompileCommand_getFilename.Fn()(cmd))
}
func (cmd CXCompileCommand) GetNumArgs() uint32 {
	return clang_CompileCommand_getNumArgs.Fn()(cmd)
}
func (cmd CXCompileCommand) GetArg(i uint32) *CXString {
	return addr(clang_CompileCommand_getArg.Fn()(cmd, i))
}
func (cmd CXCompileCommand) GetNumMappedSources() uint32 {
	return clang_CompileCommand_getNumMappedSources.Fn()(cmd)
}
func (cmd CXCompileCommand) GetMappedSourcePath(i uint32) *CXString {
	return addr(clang_CompileCommand_getMappedSourcePath.Fn()(cmd, i))
}
func (cmd CXCompileCommand) GetMappedSourceContent(i uint32) *CXString {
	return addr(clang_CompileCommand_getMappedSourceContent.Fn()(cmd, i))
}

// #endregion

// #region CXDiagnostic.h

type CXDiagnostic iptr
type CXDiagnosticSet iptr

func (diagSet CXDiagnosticSet) GetNumDiagnostics() uint32 {
	return clang_getNumDiagnosticsInSet.Fn()(diagSet)
}

func (diagSet CXDiagnosticSet) GetDiagnostic(index uint32) CXDiagnostic {
	return clang_getDiagnosticInSet.Fn()(diagSet, index)
}
func (diagSet CXDiagnosticSet) Dispose() {
	clang_disposeDiagnosticSet.Fn()(diagSet)
}
func LoadDiagnostics(file cc.String) (diagSet CXDiagnosticSet, err CXLoadDiag_Error, errMsg CXString) {
	diagSet = clang_loadDiagnostics.Fn()(file, &err, &errMsg)
	return
}
func (diag CXDiagnostic) GetChildDiagnostics() CXDiagnosticSet {
	return clang_getChildDiagnostics.Fn()(diag)
}
func (diag CXDiagnostic) Dispose() { clang_disposeDiagnostic.Fn()(diag) }
func (diag CXDiagnostic) Format(options CXDiagnosticDisplayOptions) *CXString {
	return addr(clang_formatDiagnostic.Fn()(diag, options))
}
func (diag CXDiagnostic) GetSeverity() CXDiagnosticSeverity {
	return clang_getDiagnosticSeverity.Fn()(diag)
}
func (diag CXDiagnostic) GetLocation() *CXSourceLocation {
	return addr(clang_getDiagnosticLocation.Fn()(diag))
}
func (diag CXDiagnostic) GetSpelling() *CXString {
	return addr(clang_getDiagnosticSpelling.Fn()(diag))
}
func (diag CXDiagnostic) GetOption() (disable, option *CXString) {
	var disableOpt CXString
	option = addr(clang_getDiagnosticOption.Fn()(diag, &disableOpt))
	disable = &disableOpt
	return
}
func (diag CXDiagnostic) GetCategory() uint32 {
	return clang_getDiagnosticCategory.Fn()(diag)
}
func (diag CXDiagnostic) GetCategoryText() *CXString {
	return addr(clang_getDiagnosticCategoryText.Fn()(diag))
}
func (diag CXDiagnostic) GetNumRanges() uint32 {
	return clang_getDiagnosticNumRanges.Fn()(diag)
}
func (diag CXDiagnostic) GetRange(index uint32) *CXSourceRange {
	return addr(clang_getDiagnosticRange.Fn()(diag, index))
}
func (diag CXDiagnostic) GetNumFixIts() uint32 {
	return clang_getDiagnosticNumFixIts.Fn()(diag)
}
func (diag CXDiagnostic) GetFixIt(fixIt uint32) (ret *CXString, replaceRange *CXSourceRange) {
	var rng CXSourceRange
	ret = addr(clang_getDiagnosticFixIt.Fn()(diag, fixIt, &rng))
	replaceRange = &rng
	return
}
func DefaultDiagnosticDisplayOptions() CXDiagnosticDisplayOptions {
	return clang_defaultDiagnosticDisplayOptions.Fn()()
}

// func GetDiagnosticCategoryName(category uint32) *CXString {
// 	return addr(clang_getDiagnosticCategoryName.Fn()(category))
// }

// #endregion

// #region CXIndex

type CXIndex iptr

func CreateIndex(excludeDeclarationsFromPCH, displayDiagnostics int32) CXIndex {
	return clang_createIndex.Fn()(excludeDeclarationsFromPCH, displayDiagnostics)
}

func (idx CXIndex) Dispose() { clang_disposeIndex.Fn()(idx) }

type CXIndexOptions struct {
	Size                                uint32
	ThreadBackgroundPriorityForIndexing uint8
	ThreadBackgroundPriorityForEditing  uint8
	_                                   uint64
	_                                   uint64
	PreambleStoragePath                 cc.String
	InvocationEmissionPath              cc.String
}

func CreateIndexWithOptions(opts *CXIndexOptions) CXIndex {
	return clang_createIndexWithOptions.Fn()(opts)
}

func (idx CXIndex) SetGlobalOptions(flag CXGlobalOptFlags) {
	clang_CXIndex_setGlobalOptions.Fn()(idx, flag)
}
func (idx CXIndex) GetGlobalOptions() CXGlobalOptFlags {
	return clang_CXIndex_getGlobalOptions.Fn()(idx)
}
func (idx CXIndex) SetInvocationEmissionPathOption(path string) {
	cs := cc.NewString(path)
	defer cs.Free()
	clang_CXIndex_setInvocationEmissionPathOption.Fn()(idx, cs)
}
func (idx CXIndex) IsFileMultipleIncludeGuarded(file CXFile) bool {
	return clang_isFileMultipleIncludeGuarded.Fn()(idx, file) != 0
}

// func GetFile(tu CXTranslationUnit, file_name string) CXFile {
// 	return tu.GetFile(file_name)
// }
// func GetFileContents(tu CXTranslationUnit, file CXFile) string {
// 	return tu.GetFileContents(file)
// }
// func GetLocation(tu CXTranslationUnit, file CXFile, line, column uint32) CXSourceLocation {
// 	return tu.GetLocation(file, line, column)
// }
// func GetLocationForOffset(tu CXTranslationUnit, file CXFile, offset uint32) CXSourceLocation {
// 	return tu.GetLocationForOffset(file, offset)
// }
// func GetSkippedRanges(tu CXTranslationUnit, file CXFile) *CXSourceRangeList {
// 	return tu.GetSkippedRanges(file)
// }
// func GetAllSkippedRanges(tu CXTranslationUnit) *CXSourceRangeList {
// 	return tu.GetAllSkippedRanges()
// }

func (idx CXIndex) ParseTranslationUnit(
	sourceFilename string, commandLineArgs []string,
	unsaved_files []CXUnsavedFile, flag CXTranslationUnit_Flags) CXTranslationUnit {

	sf := cc.NewString(sourceFilename)
	defer sf.Free()

	cla, clal := cc.NewStringsLen(commandLineArgs)
	defer cla.Free()

	uf, ufl := (*CXUnsavedFile)(nil), int32(len(unsaved_files))
	if ufl > 0 {
		uf = &unsaved_files[0]
	}
	return clang_parseTranslationUnit.Fn()(idx, sf, cla, clal, uf, ufl, flag)
}

// #endregion

// #region CXTranslationUnit

type CXTranslationUnit iptr

func ParseTranslationUnit(idx CXIndex,
	sourceFilename string,
	commandLineArgs []string,
	unsaved_files []CXUnsavedFile, flag CXTranslationUnit_Flags) CXTranslationUnit {
	sf := cc.NewString(sourceFilename)
	cla, clal := cc.NewStringsLen(commandLineArgs)
	defer sf.Free()
	defer cla.Free()

	uf, ufl := (*CXUnsavedFile)(nil), int32(len(unsaved_files))
	if ufl > 0 {
		uf = &unsaved_files[0]
	}
	return clang_parseTranslationUnit.Fn()(idx, sf, cla, clal, uf, ufl, flag)
}
func (tu CXTranslationUnit) Dispose() { clang_disposeTranslationUnit.Fn()(tu) }
func (tu CXTranslationUnit) GetFile(file_name string) CXFile {
	fn := cc.NewString(file_name)
	defer fn.Free()
	return clang_getFile.Fn()(tu, fn)
}
func (tu CXTranslationUnit) GetFileContents(file CXFile) string {
	var size cc.SizeT
	s := clang_getFileContents.Fn()(tu, file, &size)
	return s.Ref(uint64(size))
}
func (tu CXTranslationUnit) GetLocation(file CXFile, line, column uint32) *CXSourceLocation {
	return addr(clang_getLocation.Fn()(tu, file, line, column))
}
func (tu CXTranslationUnit) GetLocationForOffset(file CXFile, offset uint32) *CXSourceLocation {
	return addr(clang_getLocationForOffset.Fn()(tu, file, offset))
}
func (tu CXTranslationUnit) GetSkippedRanges(file CXFile) *CXSourceRangeList {
	return clang_getSkippedRanges.Fn()(tu, file)
}
func (tu CXTranslationUnit) GetAllSkippedRanges() *CXSourceRangeList {
	return clang_getAllSkippedRanges.Fn()(tu)
}
func (tu CXTranslationUnit) GetNumDiagnostics() uint32 {
	return clang_getNumDiagnostics.Fn()(tu)
}
func (tu CXTranslationUnit) GetCursor() *CXCursor {
	return GetTranslationUnitCursor(tu)
}

type CXUnsavedFile struct {
	Filename cc.String
	Contents cc.String
	Length   cc.ULong
}

// #endregion

// #region CXFile

type CXFile iptr

// #endregion

// #region CXSourceLocation

type CXSourceLocation struct {
	DataPtr [2]uptr
	DataInt uint32
}

type CXSourceRange struct {
	DatePtr   [2]uptr
	DataBegin uint32
	DataEnd   uint32
}

type CXSourceRangeList struct {
	Count  uint32
	Ranges *CXSourceRange
}

func (srl *CXSourceRangeList) List() []CXSourceRange {
	return *(*[]CXSourceRange)(cc.Slice(uptr(srl.Ranges), int64(srl.Count)))
}

// #endregion

type CXCursor struct {
	Kind CXCursorKind
	_    int32
	_    [3]iptr
}

func GetNullCursor() *CXCursor { return addr(clang_getNullCursor.Fn()()) }
func GetTranslationUnitCursor(tu CXTranslationUnit) *CXCursor {
	return addr(clang_getTranslationUnitCursor.Fn()(tu))
}
func (cur *CXCursor) Equal(other CXCursor) bool { return clang_equalCursors.Fn()(*cur, other) != 0 }
func (cur *CXCursor) IsNull() bool              { return clang_Cursor_isNull.Fn()(*cur) != 0 }
func (cur *CXCursor) Hash() uint32              { return clang_hashCursor.Fn()(*cur) }
func (cur *CXCursor) GetKind() CXCursorKind     { return clang_getCursorKind.Fn()(*cur) }
func (kid CXCursorKind) IsDeclaration() bool    { return clang_isDeclaration.Fn()(kid) != 0 }
func (cur *CXCursor) IsInvalidDeclaration() bool {
	return clang_isInvalidDeclaration.Fn()(*cur) != 0
}
func (kid CXCursorKind) IsReference() bool            { return clang_isReference.Fn()(kid) != 0 }
func (kid CXCursorKind) IsExpression() bool           { return clang_isExpression.Fn()(kid) != 0 }
func (kid CXCursorKind) IsStatement() bool            { return clang_isStatement.Fn()(kid) != 0 }
func (kid CXCursorKind) IsAttribute() bool            { return clang_isAttribute.Fn()(kid) != 0 }
func (cur *CXCursor) HasAttrs() bool                  { return clang_Cursor_hasAttrs.Fn()(*cur) != 0 }
func (kid CXCursorKind) IsInvalid() bool              { return clang_isInvalid.Fn()(kid) != 0 }
func (kid CXCursorKind) IsTranslationUnit() bool      { return clang_isTranslationUnit.Fn()(kid) != 0 }
func (kid CXCursorKind) IsPreprocessing() bool        { return clang_isPreprocessing.Fn()(kid) != 0 }
func (kid CXCursorKind) IsUnexposed() bool            { return clang_isUnexposed.Fn()(kid) != 0 }
func (cur *CXCursor) GetLinkage() CXLinkageKind       { return clang_getCursorLinkage.Fn()(*cur) }
func (cur *CXCursor) GetVisibility() CXVisibilityKind { return clang_getCursorVisibility.Fn()(*cur) }
func (cur *CXCursor) GetAvailability() CXAvailabilityKind {
	return clang_getCursorAvailability.Fn()(*cur)
}
func (cur *CXCursor) GetPlatformAvailability() (
	always_deprecated bool, deprecated_message CXString,
	always_unavailable bool, unavailable_message CXString,
	availability []CXPlatformAvailability) {

	var _always_deprecated int32
	var _always_unavailable int32

	_availability_size := int32(32)
	_availability := make([]CXPlatformAvailability, _availability_size)

	platforms := clang_getCursorPlatformAvailability.Fn()(
		*cur, &_always_deprecated, &deprecated_message,
		&_always_unavailable, &unavailable_message,
		&_availability[0], _availability_size,
	)
	always_deprecated = _always_deprecated != 0
	always_unavailable = _always_unavailable != 0
	availability = _availability[:platforms]
	return
}
func (availability *CXPlatformAvailability) Dispose() {
	clang_disposeCXPlatformAvailability.Fn()(availability)
}
func (cur *CXCursor) GetVarDeclInitializer() *CXCursor {
	return addr(clang_Cursor_getVarDeclInitializer.Fn()(*cur))
}
func (cur *CXCursor) HasVarDeclGlobalStorage() bool {
	return clang_Cursor_hasVarDeclGlobalStorage.Fn()(*cur) == 1
}
func (cur *CXCursor) HasVarDeclExternalStorage() bool {
	return clang_Cursor_hasVarDeclExternalStorage.Fn()(*cur) == 1
}
func (cur *CXCursor) GetLanguage() CXLanguageKind { return clang_getCursorLanguage.Fn()(*cur) }
func (cur *CXCursor) GetTLSKind() CXTLSKind       { return clang_getCursorTLSKind.Fn()(*cur) }
func (cur *CXCursor) GetTranslationUnit() CXTranslationUnit {
	return clang_Cursor_getTranslationUnit.Fn()(*cur)
}
func CreateCXCursorSet() CXCursorSet { return clang_createCXCursorSet.Fn()() }
func (cset CXCursorSet) Dispose()    { clang_disposeCXCursorSet.Fn()(cset) }
func (cset CXCursorSet) Contains(cursor *CXCursor) bool {
	return clang_CXCursorSet_contains.Fn()(cset, *cursor) != 0
}
func (cset CXCursorSet) Insert(cursor *CXCursor) bool {
	return clang_CXCursorSet_insert.Fn()(cset, *cursor) != 0
}
func (cur *CXCursor) GetSemanticParent() *CXCursor {
	return addr(clang_getCursorSemanticParent.Fn()(*cur))
}
func (cur *CXCursor) GetLexicalParent() *CXCursor {
	return addr(clang_getCursorLexicalParent.Fn()(*cur))
}

type CXCursors []CXCursor

func (cur *CXCursor) GetOverriddenCursors() (overridden CXCursors) {
	var overriddenPtr *CXCursor
	var numOverridden uint32

	clang_getOverriddenCursors.Fn()(*cur, &overriddenPtr, &numOverridden)
	if overriddenPtr == nil || numOverridden == 0 {
		return nil
	}

	return *(*CXCursors)(cc.Slice(uptr(overriddenPtr), int64(numOverridden)))
}

func (overridden CXCursors) Dispose() {
	if len(overridden) == 0 {
		return
	}
	clang_disposeOverriddenCursors.Fn()(&overridden[0])
}
func (cur *CXCursor) GetIncludedFile() CXFile { return clang_getIncludedFile.Fn()(*cur) }

func (cur CXCursor) VisitChildren(
	visitor func(currsor, parent *CXCursor) CXChildVisitResult) CXChildVisitResult {
	fn := cc.Cbk(func(currsor, parent CXCursor, data CXClientData) CXChildVisitResult {
		return visitor(&currsor, &parent)
	})
	defer cc.CbkClose(fn)
	return clang_visitChildren.Fn()(cur, fn, uptr(nil))
}
func (cur *CXCursor) GetSpelling() *CXString {
	return addr(clang_getCursorSpelling.Fn()(*cur))
}

type CXType struct {
	Kind CXTypeKind
	Data [2]uptr
}

func (cur *CXCursor) GetType() *CXType     { return addr(clang_getCursorType.Fn()(*cur)) }
func (typ *CXType) GetSizeOf() cc.Longlong { return clang_Type_getSizeOf.Fn()(*typ) }

type CXClientData uptr

type CXString struct {
	Data         uptr
	PrivateFlags uint32
}

func (cur *CXString) GetCString() string { return clang_getCString.Fn()(*cur).Ref() }
func (cur *CXString) String() string     { return clang_getCString.Fn()(*cur).Ref() }
func (cur *CXString) Dispose()           { clang_disposeString.Fn()(*cur) }

type CXVersion struct {
	Major    int32
	Minor    int32
	Subminor int32
}

type CXPlatformAvailability struct {
	Platform    CXString
	Introduced  CXVersion
	Deprecated  CXVersion
	Obsoleted   CXVersion
	Unavailable int32
	Message     CXString
}

type CXCursorSet iptr
