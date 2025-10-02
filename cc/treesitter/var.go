package treesitter

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

func init() { cc.Open("libtree-sitter-0.*") }

type uptr = unsafe.Pointer
type iptr = uintptr

var (
	tree_sitter_c = cc.DlFunc[func() *Language, *Language]{Name: "tree_sitter_c"}

	ts_parser_new                   = cc.DlFunc[func() *Parser, *Parser]{Name: "ts_parser_new"}
	ts_parser_delete                = cc.DlFunc[func(*Parser), cc.Void]{Name: "ts_parser_delete"}
	ts_parser_language              = cc.DlFunc[func(*Parser) *Language, *Language]{Name: "ts_parser_language"}
	ts_parser_set_language          = cc.DlFunc[func(*Parser, *Language) cc.Bool, cc.Bool]{Name: "ts_parser_set_language"}
	ts_parser_set_included_ranges   = cc.DlFunc[func(*Parser, *Range, uint32) cc.Bool, cc.Bool]{Name: "ts_parser_set_included_ranges"}
	ts_parser_included_ranges       = cc.DlFunc[func(*Parser, *uint32) *Range, *Range]{Name: "ts_parser_included_ranges"}
	ts_parser_parse                 = cc.DlFunc[func(*Parser, *Tree, Input) *Tree, *Tree]{Name: "ts_parser_parse"}
	ts_parser_parse_with_options    = cc.DlFunc[func(*Parser, *Tree, Input, ParseOptions) *Tree, *Tree]{Name: "ts_parser_parse_with_options"}
	ts_parser_parse_string          = cc.DlFunc[func(*Parser, *Tree, cc.String, uint32) *Tree, *Tree]{Name: "ts_parser_parse_string"}
	ts_parser_parse_string_encoding = cc.DlFunc[func(*Parser, *Tree, cc.String, uint32, InputEncoding) *Tree, *Tree]{Name: "ts_parser_parse_string_encoding"}
	ts_parser_reset                 = cc.DlFunc[func(*Parser), cc.Void]{Name: "ts_parser_reset"}
	ts_parser_set_timeout_micros    = cc.DlFunc[func(*Parser, uint64), cc.Void]{Name: "ts_parser_set_timeout_micros"}
	ts_parser_timeout_micros        = cc.DlFunc[func(*Parser) uint64, uint64]{Name: "ts_parser_timeout_micros"}
	ts_parser_set_cancellation_flag = cc.DlFunc[func(*Parser, *uint64), cc.Void]{Name: "ts_parser_set_cancellation_flag"}
	ts_parser_cancellation_flag     = cc.DlFunc[func(*Parser) *uint64, *uint64]{Name: "ts_parser_cancellation_flag"}
	ts_parser_set_logger            = cc.DlFunc[func(*Parser, Logger), cc.Void]{Name: "ts_parser_set_logger"}
	ts_parser_logger                = cc.DlFunc[func(*Parser) Logger, Logger]{Name: "ts_parser_logger"}
	ts_parser_print_dot_graphs      = cc.DlFunc[func(*Parser, int32), cc.Void]{Name: "ts_parser_print_dot_graphs"}

	ts_tree_copy                  = cc.DlFunc[func(*Tree) *Tree, *Tree]{Name: "ts_tree_copy"}
	ts_tree_delete                = cc.DlFunc[func(*Tree), cc.Void]{Name: "ts_tree_delete"}
	ts_tree_root_node             = cc.DlFunc[func(*Tree) Node, Node]{Name: "ts_tree_root_node"}
	ts_tree_root_node_with_offset = cc.DlFunc[func(*Tree, uint32, Point) Node, Node]{Name: "ts_tree_root_node_with_offset"}
	ts_tree_language              = cc.DlFunc[func(*Tree) *Language, *Language]{Name: "ts_tree_language"}
	ts_tree_included_ranges       = cc.DlFunc[func(*Tree, *uint32) *Range, *Range]{Name: "ts_tree_included_ranges"}
	ts_tree_edit                  = cc.DlFunc[func(*Tree, *InputEdit), cc.Void]{Name: "ts_tree_edit"}
	ts_tree_get_changed_ranges    = cc.DlFunc[func(*Tree, *Tree, *uint32) *Range, *Range]{Name: "ts_tree_get_changed_ranges"}
	ts_tree_print_dot_graph       = cc.DlFunc[func(*Tree, int32), cc.Void]{Name: "ts_tree_print_dot_graph"}

	ts_node_type                             = cc.DlFunc[func(Node) cc.String, cc.String]{Name: "ts_node_type"}
	ts_node_symbol                           = cc.DlFunc[func(Node) Symbol, Symbol]{Name: "ts_node_symbol"}
	ts_node_language                         = cc.DlFunc[func(Node) *Language, *Language]{Name: "ts_node_language"}
	ts_node_grammar_type                     = cc.DlFunc[func(Node) cc.String, cc.String]{Name: "ts_node_grammar_type"}
	ts_node_grammar_symbol                   = cc.DlFunc[func(Node) Symbol, Symbol]{Name: "ts_node_grammar_symbol"}
	ts_node_start_byte                       = cc.DlFunc[func(Node) uint32, uint32]{Name: "ts_node_start_byte"}
	ts_node_start_point                      = cc.DlFunc[func(Node) Point, Point]{Name: "ts_node_start_point"}
	ts_node_end_byte                         = cc.DlFunc[func(Node) uint32, uint32]{Name: "ts_node_end_byte"}
	ts_node_end_point                        = cc.DlFunc[func(Node) Point, Point]{Name: "ts_node_end_point"}
	ts_node_string                           = cc.DlFunc[func(Node) cc.String, cc.String]{Name: "ts_node_string"}
	ts_node_is_null                          = cc.DlFunc[func(Node) int32, int32]{Name: "ts_node_is_null"}
	ts_node_is_named                         = cc.DlFunc[func(Node) int32, int32]{Name: "ts_node_is_named"}
	ts_node_is_missing                       = cc.DlFunc[func(Node) int32, int32]{Name: "ts_node_is_missing"}
	ts_node_is_extra                         = cc.DlFunc[func(Node) int32, int32]{Name: "ts_node_is_extra"}
	ts_node_has_changes                      = cc.DlFunc[func(Node) int32, int32]{Name: "ts_node_has_changes"}
	ts_node_has_error                        = cc.DlFunc[func(Node) int32, int32]{Name: "ts_node_has_error"}
	ts_node_is_error                         = cc.DlFunc[func(Node) int32, int32]{Name: "ts_node_is_error"}
	ts_node_parse_state                      = cc.DlFunc[func(Node) StateId, StateId]{Name: "ts_node_parse_state"}
	ts_node_next_parse_state                 = cc.DlFunc[func(Node) StateId, StateId]{Name: "ts_node_next_parse_state"}
	ts_node_parent                           = cc.DlFunc[func(Node) Node, Node]{Name: "ts_node_parent"}
	ts_node_child_with_descendant            = cc.DlFunc[func(Node, Node) Node, Node]{Name: "ts_node_child_with_descendant"}
	ts_node_child                            = cc.DlFunc[func(Node, uint32) Node, Node]{Name: "ts_node_child"}
	ts_node_field_name_for_child             = cc.DlFunc[func(Node, uint32) cc.String, cc.String]{Name: "ts_node_field_name_for_child"}
	ts_node_field_name_for_named_child       = cc.DlFunc[func(Node, uint32) cc.String, cc.String]{Name: "ts_node_field_name_for_named_child"}
	ts_node_child_count                      = cc.DlFunc[func(Node) uint32, uint32]{Name: "ts_node_child_count"}
	ts_node_named_child                      = cc.DlFunc[func(Node, uint32) Node, Node]{Name: "ts_node_named_child"}
	ts_node_named_child_count                = cc.DlFunc[func(Node) uint32, uint32]{Name: "ts_node_named_child_count"}
	ts_node_child_by_field_name              = cc.DlFunc[func(Node, cc.String, uint32) Node, Node]{Name: "ts_node_child_by_field_name"}
	ts_node_child_by_field_id                = cc.DlFunc[func(Node, FieldId) Node, Node]{Name: "ts_node_child_by_field_id"}
	ts_node_next_sibling                     = cc.DlFunc[func(Node) Node, Node]{Name: "ts_node_next_sibling"}
	ts_node_prev_sibling                     = cc.DlFunc[func(Node) Node, Node]{Name: "ts_node_prev_sibling"}
	ts_node_next_named_sibling               = cc.DlFunc[func(Node) Node, Node]{Name: "ts_node_next_named_sibling"}
	ts_node_prev_named_sibling               = cc.DlFunc[func(Node) Node, Node]{Name: "ts_node_prev_named_sibling"}
	ts_node_first_child_for_byte             = cc.DlFunc[func(Node, uint32) Node, Node]{Name: "ts_node_first_child_for_byte"}
	ts_node_first_named_child_for_byte       = cc.DlFunc[func(Node, uint32) Node, Node]{Name: "ts_node_first_named_child_for_byte"}
	ts_node_descendant_count                 = cc.DlFunc[func(Node) uint32, uint32]{Name: "ts_node_descendant_count"}
	ts_node_descendant_for_byte_range        = cc.DlFunc[func(Node, uint32, uint32) Node, Node]{Name: "ts_node_descendant_for_byte_range"}
	ts_node_descendant_for_point_range       = cc.DlFunc[func(Node, Point, Point) Node, Node]{Name: "ts_node_descendant_for_point_range"}
	ts_node_named_descendant_for_byte_range  = cc.DlFunc[func(Node, uint32, uint32) Node, Node]{Name: "ts_node_named_descendant_for_byte_range"}
	ts_node_named_descendant_for_point_range = cc.DlFunc[func(Node, Point, Point) Node, Node]{Name: "ts_node_named_descendant_for_point_range"}
	ts_node_edit                             = cc.DlFunc[func(*Node, *InputEdit), cc.Void]{Name: "ts_node_edit"}
	ts_node_eq                               = cc.DlFunc[func(Node, Node) int32, int32]{Name: "ts_node_eq"}

	ts_tree_cursor_new                        = cc.DlFunc[func(Node) TreeCursor, TreeCursor]{Name: "ts_tree_cursor_new"}
	ts_tree_cursor_delete                     = cc.DlFunc[func(*TreeCursor), cc.Void]{Name: "ts_tree_cursor_delete"}
	ts_tree_cursor_reset                      = cc.DlFunc[func(*TreeCursor, Node), cc.Void]{Name: "ts_tree_cursor_reset"}
	ts_tree_cursor_reset_to                   = cc.DlFunc[func(*TreeCursor, *TreeCursor), cc.Void]{Name: "ts_tree_cursor_reset_to"}
	ts_tree_cursor_current_node               = cc.DlFunc[func(*TreeCursor) Node, Node]{Name: "ts_tree_cursor_current_node"}
	ts_tree_cursor_current_field_name         = cc.DlFunc[func(*TreeCursor) cc.String, cc.String]{Name: "ts_tree_cursor_current_field_name"}
	ts_tree_cursor_current_field_id           = cc.DlFunc[func(*TreeCursor) FieldId, FieldId]{Name: "ts_tree_cursor_current_field_id"}
	ts_tree_cursor_goto_parent                = cc.DlFunc[func(*TreeCursor) int32, int32]{Name: "ts_tree_cursor_goto_parent"}
	ts_tree_cursor_goto_next_sibling          = cc.DlFunc[func(*TreeCursor) int32, int32]{Name: "ts_tree_cursor_goto_next_sibling"}
	ts_tree_cursor_goto_previous_sibling      = cc.DlFunc[func(*TreeCursor) int32, int32]{Name: "ts_tree_cursor_goto_previous_sibling"}
	ts_tree_cursor_goto_first_child           = cc.DlFunc[func(*TreeCursor) int32, int32]{Name: "ts_tree_cursor_goto_first_child"}
	ts_tree_cursor_goto_last_child            = cc.DlFunc[func(*TreeCursor) int32, int32]{Name: "ts_tree_cursor_goto_last_child"}
	ts_tree_cursor_goto_descendant            = cc.DlFunc[func(*TreeCursor, uint32), cc.Void]{Name: "ts_tree_cursor_goto_descendant"}
	ts_tree_cursor_current_descendant_index   = cc.DlFunc[func(*TreeCursor) uint32, uint32]{Name: "ts_tree_cursor_current_descendant_index"}
	ts_tree_cursor_current_depth              = cc.DlFunc[func(*TreeCursor) uint32, uint32]{Name: "ts_tree_cursor_current_depth"}
	ts_tree_cursor_goto_first_child_for_byte  = cc.DlFunc[func(*TreeCursor, uint32) int64, int64]{Name: "ts_tree_cursor_goto_first_child_for_byte"}
	ts_tree_cursor_goto_first_child_for_point = cc.DlFunc[func(*TreeCursor, Point) int64, int64]{Name: "ts_tree_cursor_goto_first_child_for_point"}
	ts_tree_cursor_copy                       = cc.DlFunc[func(*TreeCursor) TreeCursor, TreeCursor]{Name: "ts_tree_cursor_copy"}

	ts_query_new                           = cc.DlFunc[func(*Language, cc.String, uint32, *uint32, *QueryError) *Query, *Query]{Name: "ts_query_new"}
	ts_query_delete                        = cc.DlFunc[func(*Query), cc.Void]{Name: "ts_query_delete"}
	ts_query_pattern_count                 = cc.DlFunc[func(*Query) uint32, uint32]{Name: "ts_query_pattern_count"}
	ts_query_capture_count                 = cc.DlFunc[func(*Query) uint32, uint32]{Name: "ts_query_capture_count"}
	ts_query_string_count                  = cc.DlFunc[func(*Query) uint32, uint32]{Name: "ts_query_string_count"}
	ts_query_start_byte_for_pattern        = cc.DlFunc[func(*Query, uint32) uint32, uint32]{Name: "ts_query_start_byte_for_pattern"}
	ts_query_end_byte_for_pattern          = cc.DlFunc[func(*Query, uint32) uint32, uint32]{Name: "ts_query_end_byte_for_pattern"}
	ts_query_predicates_for_pattern        = cc.DlFunc[func(*Query, uint32, *uint32) *QueryPredicateStep, *QueryPredicateStep]{Name: "ts_query_predicates_for_pattern"}
	ts_query_is_pattern_rooted             = cc.DlFunc[func(*Query, uint32) int32, int32]{Name: "ts_query_is_pattern_rooted"}
	ts_query_is_pattern_non_local          = cc.DlFunc[func(*Query, uint32) int32, int32]{Name: "ts_query_is_pattern_non_local"}
	ts_query_is_pattern_guaranteed_at_step = cc.DlFunc[func(*Query, uint32) int32, int32]{Name: "ts_query_is_pattern_guaranteed_at_step"}
	ts_query_capture_name_for_id           = cc.DlFunc[func(*Query, uint32, *uint32) cc.String, cc.String]{Name: "ts_query_capture_name_for_id"}
	ts_query_capture_quantifier_for_id     = cc.DlFunc[func(*Query, uint32, uint32) Quantifier, Quantifier]{Name: "ts_query_capture_quantifier_for_id"}
	ts_query_string_value_for_id           = cc.DlFunc[func(*Query, uint32, *uint32) cc.String, cc.String]{Name: "ts_query_string_value_for_id"}
	ts_query_disable_capture               = cc.DlFunc[func(*Query, cc.String, uint32), cc.Void]{Name: "ts_query_disable_capture"}
	ts_query_disable_pattern               = cc.DlFunc[func(*Query, uint32), cc.Void]{Name: "ts_query_disable_pattern"}

	ts_query_cursor_new                    = cc.DlFunc[func() *QueryCursor, *QueryCursor]{Name: "ts_query_cursor_new"}
	ts_query_cursor_delete                 = cc.DlFunc[func(*QueryCursor), cc.Void]{Name: "ts_query_cursor_delete"}
	ts_query_cursor_exec                   = cc.DlFunc[func(*QueryCursor, *Query, Node), cc.Void]{Name: "ts_query_cursor_exec"}
	ts_query_cursor_exec_with_options      = cc.DlFunc[func(*QueryCursor, *Query, Node, *QueryCursorOptions), cc.Void]{Name: "ts_query_cursor_exec_with_options"}
	ts_query_cursor_did_exceed_match_limit = cc.DlFunc[func(*QueryCursor) int32, int32]{Name: "ts_query_cursor_did_exceed_match_limit"}
	ts_query_cursor_match_limit            = cc.DlFunc[func(*QueryCursor) uint32, uint32]{Name: "ts_query_cursor_match_limit"}
	ts_query_cursor_set_match_limit        = cc.DlFunc[func(*QueryCursor, uint32), cc.Void]{Name: "ts_query_cursor_set_match_limit"}
	ts_query_cursor_set_timeout_micros     = cc.DlFunc[func(*QueryCursor, uint64), cc.Void]{Name: "ts_query_cursor_set_timeout_micros"}
	ts_query_cursor_timeout_micros         = cc.DlFunc[func(*QueryCursor) uint64, uint64]{Name: "ts_query_cursor_timeout_micros"}
	ts_query_cursor_set_byte_range         = cc.DlFunc[func(*QueryCursor, uint32, uint32) int32, int32]{Name: "ts_query_cursor_set_byte_range"}
	ts_query_cursor_set_point_range        = cc.DlFunc[func(*QueryCursor, Point, Point) int32, int32]{Name: "ts_query_cursor_set_point_range"}
	ts_query_cursor_next_match             = cc.DlFunc[func(*QueryCursor, *QueryMatch) int32, int32]{Name: "ts_query_cursor_next_match"}
	ts_query_cursor_remove_match           = cc.DlFunc[func(*QueryCursor, uint32), cc.Void]{Name: "ts_query_cursor_remove_match"}
	ts_query_cursor_next_capture           = cc.DlFunc[func(*QueryCursor, *QueryMatch, *uint32) int32, int32]{Name: "ts_query_cursor_next_capture"}
	ts_query_cursor_set_max_start_depth    = cc.DlFunc[func(*QueryCursor, uint32), cc.Void]{Name: "ts_query_cursor_set_max_start_depth"}

	ts_language_copy              = cc.DlFunc[func(*Language) *Language, *Language]{Name: "ts_language_copy"}
	ts_language_delete            = cc.DlFunc[func(*Language), cc.Void]{Name: "ts_language_delete"}
	ts_language_symbol_count      = cc.DlFunc[func(*Language) uint32, uint32]{Name: "ts_language_symbol_count"}
	ts_language_state_count       = cc.DlFunc[func(*Language) uint32, uint32]{Name: "ts_lants_language_state_countguage_symbol_count"}
	ts_language_symbol_for_name   = cc.DlFunc[func(*Language, cc.String, uint32, int32) Symbol, Symbol]{Name: "ts_language_symbol_for_name"}
	ts_language_field_count       = cc.DlFunc[func(*Language) uint32, uint32]{Name: "ts_language_field_count"}
	ts_language_field_name_for_id = cc.DlFunc[func(*Language, FieldId) cc.String, cc.String]{Name: "ts_language_field_name_for_id"}
	ts_language_field_id_for_name = cc.DlFunc[func(*Language, cc.String, uint32) FieldId, FieldId]{Name: "ts_language_field_id_for_name"}
	ts_language_supertypes        = cc.DlFunc[func(*Language, *uint32) *Symbol, *Symbol]{Name: "ts_language_supertypes"}
	ts_language_subtypes          = cc.DlFunc[func(*Language, Symbol, *uint32) *Symbol, *Symbol]{Name: "ts_language_subtypes"}
	ts_language_symbol_name       = cc.DlFunc[func(*Language, Symbol) cc.String, cc.String]{Name: "ts_language_symbol_name"}
	ts_language_symbol_type       = cc.DlFunc[func(*Language, Symbol) SymbolType, SymbolType]{Name: "ts_language_symbol_type"}
	ts_language_version           = cc.DlFunc[func(*Language) uint32, uint32]{Name: "ts_language_version"}
	ts_language_abi_version       = cc.DlFunc[func(*Language) uint32, uint32]{Name: "ts_language_abi_version"}
	ts_language_metadata          = cc.DlFunc[func(*Language) *LanguageMetadata, *LanguageMetadata]{Name: "ts_language_metadata"}
	ts_language_next_state        = cc.DlFunc[func(*Language, StateId, Symbol) StateId, StateId]{Name: "ts_language_next_state"}
	ts_language_name              = cc.DlFunc[func(*Language) cc.String, cc.String]{Name: "ts_language_name"}

	ts_lookahead_iterator_new                 = cc.DlFunc[func(*Language, StateId) *LookaheadIterator, *LookaheadIterator]{Name: "ts_lookahead_iterator_new"}
	ts_lookahead_iterator_delete              = cc.DlFunc[func(*LookaheadIterator), cc.Void]{Name: "ts_lookahead_iterator_delete"}
	ts_lookahead_iterator_reset_state         = cc.DlFunc[func(*LookaheadIterator, StateId) int32, int32]{Name: "ts_lookahead_iterator_reset_state"}
	ts_lookahead_iterator_reset               = cc.DlFunc[func(*LookaheadIterator, *Language, StateId) int32, int32]{Name: "ts_lookahead_iterator_reset"}
	ts_lookahead_iterator_language            = cc.DlFunc[func(*LookaheadIterator) *Language, *Language]{Name: "ts_lookahead_iterator_language"}
	ts_lookahead_iterator_next                = cc.DlFunc[func(*LookaheadIterator) int32, int32]{Name: "ts_lookahead_iterator_next"}
	ts_lookahead_iterator_current_symbol      = cc.DlFunc[func(*LookaheadIterator) Symbol, Symbol]{Name: "ts_lookahead_iterator_current_symbol"}
	ts_lookahead_iterator_current_symbol_name = cc.DlFunc[func(*LookaheadIterator) cc.String, cc.String]{Name: "ts_lookahead_iterator_current_symbol_name"}

	ts_wasm_store_new            = cc.DlFunc[func(*WasmEngine, *WasmError) *WasmStore, *WasmStore]{Name: "ts_wasm_store_new"}
	ts_wasm_store_delete         = cc.DlFunc[func(*WasmStore), cc.Void]{Name: "ts_wasm_store_delete"}
	ts_wasm_store_load_language  = cc.DlFunc[func(*WasmStore, cc.String, *byte, uint32, *WasmError) *Language, *Language]{Name: "ts_wasm_store_load_language"}
	ts_wasm_store_language_count = cc.DlFunc[func(*WasmStore) uint64, uint64]{Name: "ts_wasm_store_language_count"}
	ts_language_is_wasm          = cc.DlFunc[func(*Language) int32, int32]{Name: "ts_language_is_wasm"}
	ts_parser_set_wasm_store     = cc.DlFunc[func(*Parser, *WasmStore), cc.Void]{Name: "ts_parser_set_wasm_store"}
	ts_parser_take_wasm_store    = cc.DlFunc[func(*Parser) *WasmStore, *WasmStore]{Name: "ts_parser_take_wasm_store"}

// /**
//  * Set the allocation functions used by the library.
//  *
//  * By default, Tree-sitter uses the standard libc allocation functions,
//  * but aborts the process when an allocation fails. This function lets
//  * you supply alternative allocation functions at runtime.
//  *
//  * If you pass `NULL` for any parameter, Tree-sitter will switch back to
//  * its default implementation of that function.
//  *
//  * If you call this function after the library has already been used, then
//  * you must ensure that either:
//  *  1. All the existing objects have been freed.
//  *  2. The new allocator shares its state with the old one, so it is capable
//  *     of freeing memory that was allocated by the old allocator.
//  */
// void ts_set_allocator(
//   void *(*new_malloc)(size_t),
// 	void *(*new_calloc)(size_t, size_t),
// 	void *(*new_realloc)(void *, size_t),
// 	void (*new_free)(void *)
// );

)
