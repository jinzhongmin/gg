package treesitter

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

func addr[T any](a T) *T { return &a }

func C() *Language { return tree_sitter_c.Fn()() }

type Language struct{}
type Parser struct{}
type Tree struct{}
type Query struct{}
type QueryCursor struct{}
type LookaheadIterator struct{}

type Point struct{ Row, Column uint32 }

type Range struct {
	StartPoint Point
	EndPoint   Point
	StartByte  uint32
	EndByte    uint32
}

type Input struct {
	Payload  uptr
	Read     cc.Func // const char *(*read)(void *payload, uint32_t byte_index, TSPoint position, uint32_t *bytes_read);
	Encoding InputEncoding
	Decode   cc.Func // typedef uint32_t (*DecodeFunction)(const uint8_t *string, uint32_t length, int32_t *code_point);
}

type ParseState struct {
	Payload           uptr
	CurrentByteOffset uint32
	HasError          cc.Bool
}

type ParseOptions struct {
	Payload          uptr
	ProgressCallback cc.Func // bool (*progress_callback)(TSParseState *state);
}

type Logger struct {
	Payload uptr
	Log     cc.Func // void (*log)(void *payload, TSLogType log_type, const char *buffer);
}

type InputEdit struct {
	StartByte  uint32
	OldEndbyte uint32
	NewEndByte uint32

	StartPoint  Point
	OldEndPoint Point
	NewEndPoint Point
}

type Node struct {
	Context [4]uint32
	ID      uptr
	Tree    *Tree
}

type TreeCursor struct {
	Tree    uptr
	ID      uptr
	Context [3]uint32
}

type QueryCapture struct {
	Node  Node
	Index uint32
}

type QueryMatch struct {
	ID           uint32
	PatternIndex uint16
	CaptureCount uint16
	Captures     *QueryCapture
}

func (qm QueryMatch) ListCaptures() []QueryCapture {
	if qm.Captures == nil || qm.CaptureCount == 0 {
		return nil
	}
	return unsafe.Slice(qm.Captures, qm.CaptureCount)[:]
}

type QueryPredicateStep struct {
	Type    QueryPredicateStepType
	ValueID uint32
}

type QueryCursorState struct {
	Payload           uptr
	CurrentByteOffset uint32
}

type QueryCursorOptions struct {
	Payload          uptr
	ProgressCallback cc.Func // bool (*progress_callback)(TSQueryCursorState *state);
}

type LanguageMetadata struct {
	MajorVersion uint8
	MinorVersion uint8
	PatchVersion uint8
}

func NewParser() *Parser              { return ts_parser_new.Fn()() }
func (p *Parser) Delete()             { ts_parser_delete.Fn()(p) }
func (p *Parser) Language() *Language { return ts_parser_language.Fn()(p) }
func (p *Parser) SetLanguage(lang *Language) bool {
	return ts_parser_set_language.Fn()(p, lang) == cc.True
}
func (p *Parser) SetIncludedRanges(ranges []Range) bool {
	r, n := (*Range)(nil), uint32(len(ranges))
	if n > 0 {
		r = &ranges[0]
	}
	return ts_parser_set_included_ranges.Fn()(p, r, n) == cc.True
}
func (p *Parser) IncludedRanges() []Range {
	n := uint32(0)
	r := ts_parser_included_ranges.Fn()(p, &n)
	if n == 0 || r == nil {
		return nil
	}
	return unsafe.Slice(r, n)[:]
}
func (p *Parser) Parse(oldTree *Tree, input Input) *Tree {
	return ts_parser_parse.Fn()(p, oldTree, input)
}
func (p *Parser) ParseWithOptions(oldTree *Tree, input Input, opts ParseOptions) *Tree {
	return ts_parser_parse_with_options.Fn()(p, oldTree, input, opts)
}
func (p *Parser) ParseString(oldTree *Tree, str string) *Tree {
	s, n := cc.NewStringLen(str)
	defer s.Free()
	return ts_parser_parse_string.Fn()(p, oldTree, s, uint32(n))
}
func (p *Parser) ParseStringByte(oldTree *Tree, str []byte) *Tree {
	s, n := cc.String(0), len(str)
	if n > 0 {
		s = (cc.String)(uptr(&str[0]))
	}
	return ts_parser_parse_string.Fn()(p, oldTree, s, uint32(n))
}
func (p *Parser) ParseStringEncoding(oldTree *Tree, str string, enc InputEncoding) *Tree {
	s, n := cc.NewStringLen(str)
	defer s.Free()
	return ts_parser_parse_string_encoding.Fn()(p, oldTree, s, uint32(n), enc)
}
func (p *Parser) Reset()                     { ts_parser_reset.Fn()(p) }
func (p *Parser) SetTimeoutMicros(ms uint64) { ts_parser_set_timeout_micros.Fn()(p, ms) }
func (p *Parser) TimeoutMicros() uint64      { return ts_parser_timeout_micros.Fn()(p) }
func (p *Parser) SetCancellationFlag(flag *uint64) {
	ts_parser_set_cancellation_flag.Fn()(p, flag)
}
func (p *Parser) CancellationFlag() *uint64 {
	return ts_parser_cancellation_flag.Fn()(p)
}
func (p *Parser) SetLogger(logger Logger) { ts_parser_set_logger.Fn()(p, logger) }
func (p *Parser) Logger() *Logger         { return addr(ts_parser_logger.Fn()(p)) }
func (p *Parser) PrintDotGraphs(fd int32) { ts_parser_print_dot_graphs.Fn()(p, fd) }

func (t *Tree) Copy() *Tree     { return ts_tree_copy.Fn()(t) }
func (t *Tree) Delete()         { ts_tree_delete.Fn()(t) }
func (t *Tree) RootNode() *Node { return addr(ts_tree_root_node.Fn()(t)) }
func (t *Tree) RootNodeWithOffset(offset_bytes uint32, offset_extent Point) *Node {
	return addr(ts_tree_root_node_with_offset.Fn()(t, offset_bytes, offset_extent))
}
func (t *Tree) Language() *Language { return ts_tree_language.Fn()(t) }
func (t *Tree) IncludedRanges() []Range {
	l := uint32(0)
	r := ts_tree_included_ranges.Fn()(t, &l)
	return unsafe.Slice(r, l)
}
func (t *Tree) Edit(edit *InputEdit) { ts_tree_edit.Fn()(t, edit) }
func (t *Tree) GetChangedRanges(newTree *Tree) []Range {
	l := uint32(0)
	r := ts_tree_get_changed_ranges.Fn()(t, newTree, &l)
	return unsafe.Slice(r, l)
}
func (t *Tree) PrintDotGraph(file_descriptor int32) { ts_tree_print_dot_graph.Fn()(t, file_descriptor) }

func (n *Node) Type() string            { return ts_node_type.Fn()(*n).String() }
func (n *Node) Symbol() Symbol          { return ts_node_symbol.Fn()(*n) }
func (n *Node) Language() *Language     { return ts_node_language.Fn()(*n) }
func (n *Node) GrammarType() string     { return ts_node_grammar_type.Fn()(*n).String() }
func (n *Node) GrammarSymbol() Symbol   { return ts_node_grammar_symbol.Fn()(*n) }
func (n *Node) StartByte() uint32       { return ts_node_start_byte.Fn()(*n) }
func (n *Node) StartPoint() *Point      { return addr(ts_node_start_point.Fn()(*n)) }
func (n *Node) EndByte() uint32         { return ts_node_end_byte.Fn()(*n) }
func (n *Node) EndPoint() *Point        { return addr(ts_node_end_point.Fn()(*n)) }
func (n *Node) String() string          { return ts_node_string.Fn()(*n).String() }
func (n *Node) IsNull() bool            { return ts_node_is_null.Fn()(*n) != 0 }
func (n *Node) IsNamed() bool           { return ts_node_is_named.Fn()(*n) != 0 }
func (n *Node) IsMissing() bool         { return ts_node_is_missing.Fn()(*n) != 0 }
func (n *Node) IsExtra() bool           { return ts_node_is_extra.Fn()(*n) != 0 }
func (n *Node) HasChanges() bool        { return ts_node_has_changes.Fn()(*n) != 0 }
func (n *Node) HasError() bool          { return ts_node_has_error.Fn()(*n) != 0 }
func (n *Node) IsError() bool           { return ts_node_is_error.Fn()(*n) != 0 }
func (n *Node) ParseState() StateId     { return ts_node_parse_state.Fn()(*n) }
func (n *Node) NextParseState() StateId { return ts_node_next_parse_state.Fn()(*n) }
func (n *Node) Parent() *Node           { return addr(ts_node_parent.Fn()(*n)) }
func (n *Node) ChildWithDescendant(descendant *Node) *Node {
	return addr(ts_node_child_with_descendant.Fn()(*n, *descendant))
}
func (n *Node) Child(idx uint32) *Node { return addr(ts_node_child.Fn()(*n, idx)) }
func (n *Node) FieldNameForChild(idx uint32) string {
	return ts_node_field_name_for_child.Fn()(*n, idx).String()
}
func (n *Node) FieldNameForNamedChild(idx uint32) string {
	return ts_node_field_name_for_named_child.Fn()(*n, idx).String()
}
func (n *Node) ChildCount() uint32          { return ts_node_child_count.Fn()(*n) }
func (n *Node) NamedChild(idx uint32) *Node { return addr(ts_node_named_child.Fn()(*n, idx)) }
func (n *Node) NamedChildCount() uint32     { return ts_node_named_child_count.Fn()(*n) }
func (n *Node) ChildByFieldName(name string) *Node {
	nm, l := cc.NewStringLen(name)
	defer nm.Free()
	return addr(ts_node_child_by_field_name.Fn()(*n, nm, uint32(l)))
}
func (n *Node) ChildByFieldId(id FieldId) *Node { return addr(ts_node_child_by_field_id.Fn()(*n, id)) }
func (n *Node) NextSibling() *Node              { return addr(ts_node_next_sibling.Fn()(*n)) }
func (n *Node) PrevSibling() *Node              { return addr(ts_node_prev_sibling.Fn()(*n)) }
func (n *Node) NextNamedSibling() *Node         { return addr(ts_node_next_named_sibling.Fn()(*n)) }
func (n *Node) PrevNamedSibling() *Node         { return addr(ts_node_prev_named_sibling.Fn()(*n)) }
func (n *Node) FirstChildForByte(byte uint32) *Node {
	return addr(ts_node_first_child_for_byte.Fn()(*n, byte))
}
func (n *Node) FirstNamedChildForByte(byte uint32) *Node {
	return addr(ts_node_first_named_child_for_byte.Fn()(*n, byte))
}
func (n *Node) DescendantCount() uint32 { return ts_node_descendant_count.Fn()(*n) }
func (n *Node) DescendantForByteRange(start, end uint32) *Node {
	return addr(ts_node_descendant_for_byte_range.Fn()(*n, start, end))
}
func (n *Node) DescendantForPointRange(start, end Point) *Node {
	return addr(ts_node_descendant_for_point_range.Fn()(*n, start, end))
}
func (n *Node) NamedDescendantForByteRange(start, end uint32) *Node {
	return addr(ts_node_named_descendant_for_byte_range.Fn()(*n, start, end))
}
func (n *Node) NamedDescendantForPointRange(start, end Point) *Node {
	return addr(ts_node_named_descendant_for_point_range.Fn()(*n, start, end))
}
func (n *Node) Edit(edit *InputEdit) { ts_node_edit.Fn()(n, edit) }
func (n *Node) Eq(other *Node) bool  { return ts_node_eq.Fn()(*n, *other) != 0 }

func (n *Node) NewTreeCursor() *TreeCursor     { return addr(ts_tree_cursor_new.Fn()(*n)) }
func (tc *TreeCursor) Delete()                 { ts_tree_cursor_delete.Fn()(tc) }
func (tc *TreeCursor) Reset(node *Node)        { ts_tree_cursor_reset.Fn()(tc, *node) }
func (tc *TreeCursor) ResetTo(dst *TreeCursor) { ts_tree_cursor_reset_to.Fn()(dst, tc) }
func (tc *TreeCursor) CurrentNode() *Node      { return addr(ts_tree_cursor_current_node.Fn()(tc)) }
func (tc *TreeCursor) CurrentFieldName() string {
	return ts_tree_cursor_current_field_name.Fn()(tc).String()
}
func (tc *TreeCursor) CurrentFieldId() FieldId {
	return ts_tree_cursor_current_field_id.Fn()(tc)
}
func (tc *TreeCursor) GotoParent() bool      { return ts_tree_cursor_goto_parent.Fn()(tc) != 0 }
func (tc *TreeCursor) GotoNextSibling() bool { return ts_tree_cursor_goto_next_sibling.Fn()(tc) != 0 }
func (tc *TreeCursor) GotoPreviousSibling() bool {
	return ts_tree_cursor_goto_previous_sibling.Fn()(tc) != 0
}
func (tc *TreeCursor) GotoFirstChild() bool { return ts_tree_cursor_goto_first_child.Fn()(tc) != 0 }
func (tc *TreeCursor) GotoLastChild() bool  { return ts_tree_cursor_goto_last_child.Fn()(tc) != 0 }
func (tc *TreeCursor) GotoDescendant(goalDescendantIndex uint32) {
	ts_tree_cursor_goto_descendant.Fn()(tc, goalDescendantIndex)
}
func (tc *TreeCursor) CurrentDescendantIndex() uint32 {
	return ts_tree_cursor_current_descendant_index.Fn()(tc)
}
func (tc *TreeCursor) CurrentDepth() uint32 { return ts_tree_cursor_current_depth.Fn()(tc) }
func (tc *TreeCursor) GotoFirstChildForByte(goalByte uint32) int64 {
	return ts_tree_cursor_goto_first_child_for_byte.Fn()(tc, goalByte)
}
func (tc *TreeCursor) GotoFirstChildForPoint(goalPoint Point) int64 {
	return ts_tree_cursor_goto_first_child_for_point.Fn()(tc, goalPoint)
}
func (tc *TreeCursor) Copy() *TreeCursor { return addr(ts_tree_cursor_copy.Fn()(tc)) }

func NewQuery(lang *Language, source string) (query *Query, errOffset uint32, errType QueryError) {
	ss, sl := cc.NewStringLen(source)
	defer ss.Free()
	query = ts_query_new.Fn()(lang, ss, uint32(sl), &errOffset, &errType)
	return
}
func (q *Query) Delete()              { ts_query_delete.Fn()(q) }
func (q *Query) PatternCount() uint32 { return ts_query_pattern_count.Fn()(q) }
func (q *Query) CaptureCount() uint32 { return ts_query_capture_count.Fn()(q) }
func (q *Query) StringCount() uint32  { return ts_query_string_count.Fn()(q) }
func (q *Query) StartByteForPattern(patternIndex uint32) uint32 {
	return ts_query_start_byte_for_pattern.Fn()(q, patternIndex)
}
func (q *Query) EndByteForPattern(patternIndex uint32) uint32 {
	return ts_query_end_byte_for_pattern.Fn()(q, patternIndex)
}
func (q *Query) PredicatesForPattern(patternIndex uint32) []QueryPredicateStep {
	count := uint32(0)
	p := ts_query_predicates_for_pattern.Fn()(q, patternIndex, &count)
	return unsafe.Slice(p, count)[:]
}
func (q *Query) IsPatternRooted(patternIndex uint32) bool {
	return ts_query_is_pattern_rooted.Fn()(q, patternIndex) != 0
}
func (q *Query) IsPatternNonLocal(patternIndex uint32) bool {
	return ts_query_is_pattern_non_local.Fn()(q, patternIndex) != 0
}
func (q *Query) IsPatternGuaranteedAtStep(byteOffset uint32) bool {
	return ts_query_is_pattern_guaranteed_at_step.Fn()(q, byteOffset) != 0
}
func (q *Query) NameForId(index uint32) string {
	len := uint32(0)
	ss := ts_query_capture_name_for_id.Fn()(q, index, &len)
	return ss.String(uint64(len))
}
func (q *Query) CaptureQuantifierForId(patternIdx, captureIdx uint32) Quantifier {
	return ts_query_capture_quantifier_for_id.Fn()(q, patternIdx, captureIdx)
}
func (q *Query) StringValueForId(index uint32) string {
	len := uint32(0)
	ss := ts_query_string_value_for_id.Fn()(q, index, &len)
	return ss.String(uint64(len))
}
func (q *Query) DisableCapture(name string) {
	ss, len := cc.NewStringLen(name)
	defer ss.Free()
	ts_query_disable_capture.Fn()(q, ss, uint32(len))
}
func (q *Query) DisablePattern(patternIndex uint32) { ts_query_disable_pattern.Fn()(q, patternIndex) }

func NewQueryCursor() *QueryCursor                   { return ts_query_cursor_new.Fn()() }
func (q *QueryCursor) Delete()                       { ts_query_cursor_delete.Fn()(q) }
func (q *QueryCursor) Exec(query *Query, node *Node) { ts_query_cursor_exec.Fn()(q, query, *node) }
func (q *QueryCursor) ExecWithOptions(query *Query, node *Node, options *QueryCursorOptions) {
	ts_query_cursor_exec_with_options.Fn()(q, query, *node, options)
}
func (q *QueryCursor) DidExceedMatchLimit() bool {
	return ts_query_cursor_did_exceed_match_limit.Fn()(q) != 0
}
func (q *QueryCursor) MatchLimit() uint32         { return ts_query_cursor_match_limit.Fn()(q) }
func (q *QueryCursor) SetMatchLimit(limit uint32) { ts_query_cursor_set_match_limit.Fn()(q, limit) }
func (q *QueryCursor) SetTimeoutMicros(ms uint64) {
	ts_query_cursor_set_timeout_micros.Fn()(q, ms)
}
func (q *QueryCursor) GetTimeoutMicros() uint64 {
	return ts_query_cursor_timeout_micros.Fn()(q)
}
func (q *QueryCursor) SetByteRange(start, end uint32) bool {
	return ts_query_cursor_set_byte_range.Fn()(q, start, end) != 0
}
func (q *QueryCursor) SetPointRange(start, end Point) bool {
	return ts_query_cursor_set_point_range.Fn()(q, start, end) != 0
}
func (q *QueryCursor) NextMatch() (match *QueryMatch, matched bool) {
	var _match QueryMatch
	matched = ts_query_cursor_next_match.Fn()(q, &_match) != 0
	if !matched {
		return nil, false
	}
	return &_match, true
}
func (q *QueryCursor) RemoveMatch(matchId uint32) {
	ts_query_cursor_remove_match.Fn()(q, matchId)
}
func (q *QueryCursor) NextCapture(matchId uint32) (match QueryMatch, captureIndex uint32, caoture bool) {
	caoture = ts_query_cursor_next_capture.Fn()(q, &match, &captureIndex) != 0
	return
}
func (q *QueryCursor) SetMaxStartDepth(maxStartDepth uint32) {
	ts_query_cursor_set_max_start_depth.Fn()(q, maxStartDepth)
}

func (l *Language) Copy() *Language     { return ts_language_copy.Fn()(l) }
func (l *Language) Delete()             { ts_language_delete.Fn()(l) }
func (l *Language) SymbolCount() uint32 { return ts_language_symbol_count.Fn()(l) }
func (l *Language) StateCount() uint32  { return ts_language_state_count.Fn()(l) }
func (l *Language) SymbolForName(str string, isNamed bool) Symbol {
	ss, sl := cc.NewStringLen(str)
	defer ss.Free()
	in := 1
	if !isNamed {
		in = 0
	}
	return ts_language_symbol_for_name.Fn()(l, ss, uint32(sl), int32(in))
}
func (l *Language) FieldCount() uint32 { return ts_language_field_count.Fn()(l) }
func (l *Language) FieldNameForId(id FieldId) string {
	return ts_language_field_name_for_id.Fn()(l, id).String()
}
func (l *Language) FieldIdForName(name string) FieldId {
	ns, nl := cc.NewStringLen(name)
	defer ns.Free()
	return ts_language_field_id_for_name.Fn()(l, ns, uint32(nl))
}
func (l *Language) Supertypes() []Symbol {
	n := uint32(0)
	p := ts_language_supertypes.Fn()(l, &n)
	return unsafe.Slice(p, n)
}
func (l *Language) Subtypes(supertype Symbol) []Symbol {
	n := uint32(0)
	p := ts_language_subtypes.Fn()(l, supertype, &n)
	return unsafe.Slice(p, n)
}
func (l *Language) SymbolName(symbol Symbol) string {
	return ts_language_symbol_name.Fn()(l, symbol).String()
}
func (l *Language) SymbolType(symbol Symbol) SymbolType {
	return ts_language_symbol_type.Fn()(l, symbol)
}
func (l *Language) Version() uint32             { return ts_language_version.Fn()(l) }
func (l *Language) AbiVersion() uint32          { return ts_language_abi_version.Fn()(l) }
func (l *Language) Metadata() *LanguageMetadata { return ts_language_metadata.Fn()(l) }
func (l *Language) NextState(state StateId, symbol Symbol) StateId {
	return ts_language_next_state.Fn()(l, state, symbol)
}
func (l *Language) Name() string { return ts_language_name.Fn()(l).String() }

func NewtLookaheadIterator(lang *Language, state StateId) *LookaheadIterator {
	return ts_lookahead_iterator_new.Fn()(lang, state)
}
func (l *LookaheadIterator) Delete() { ts_lookahead_iterator_delete.Fn()(l) }
func (l *LookaheadIterator) ResetState(state StateId) bool {
	return ts_lookahead_iterator_reset_state.Fn()(l, state) != 0
}
func (l *LookaheadIterator) Reset(lang *Language, state StateId) bool {
	return ts_lookahead_iterator_reset.Fn()(l, lang, state) != 0
}
func (l *LookaheadIterator) Language() *Language { return ts_lookahead_iterator_language.Fn()(l) }
func (l *LookaheadIterator) Next() bool          { return ts_lookahead_iterator_next.Fn()(l) != 0 }
func (l *LookaheadIterator) CurrentSymbol() Symbol {
	return ts_lookahead_iterator_current_symbol.Fn()(l)
}
func (l *LookaheadIterator) CurrentSymbolName() string {
	return ts_lookahead_iterator_current_symbol_name.Fn()(l).String()
}

type WasmEngine struct{}
type WasmStore struct{}

type WasmError struct {
	Kind    WasmErrorKind
	Message cc.String
}

func NewWasmStore(engine *WasmEngine) (store *WasmStore, err *WasmError) {
	var _err WasmError
	_store := ts_wasm_store_new.Fn()(engine, &_err)
	if _err.Kind == WasmErrorKindNone {
		return _store, nil
	}
	return _store, &_err
}
func (s *WasmStore) Delete() { ts_wasm_store_delete.Fn()(s) }
func (s *WasmStore) LoadLanguage(name string, wasm []byte) (lang *Language, err *WasmError) {
	var _err WasmError
	ns := cc.NewString(name)
	defer ns.Free()
	p, l := (*byte)(nil), len(wasm)
	if l > 0 {
		p = &wasm[0]
	}
	lang = ts_wasm_store_load_language.Fn()(s, ns, p, uint32(l), &_err)
	if _err.Kind == WasmErrorKindNone {
		return lang, nil
	}
	return lang, &_err
}
func (s *WasmStore) LanguageCount() uint64      { return ts_wasm_store_language_count.Fn()(s) }
func (l *Language) IsWasm() bool                { return ts_language_is_wasm.Fn()(l) != 0 }
func (p *Parser) SetWasmStore(store *WasmStore) { ts_parser_set_wasm_store.Fn()(p, store) }
func (p *Parser) TakeWasmStore() *WasmStore     { return ts_parser_take_wasm_store.Fn()(p) }
