package glib

import (
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

type UPtr = unsafe.Pointer
type IPtr = uintptr

func anyptr(a interface{}) UPtr { return (*(*[2]UPtr)((UPtr(&a))))[1] }
func vcbu(fn interface{}) UPtr  { return cc.Cbk(fn) }

func init() { cc.Open("libglib-2*") }

var (
	// #region Alloc
	malloc = cc.DlFunc[func(uint64) UPtr, UPtr]{Name: "malloc"}
	memset = cc.DlFunc[func(UPtr, int32, uint64), cc.Void]{Name: "memset"}
	// #endregion

	// #region Array
	g_array_new                      = cc.DlFunc[func(zeroTerminated, clear bool, elementSize uint32) UPtr, UPtr]{Name: "g_array_new"}
	g_array_new_take                 = cc.DlFunc[func(data UPtr, length uint64, clear bool, elementSize uint64) UPtr, UPtr]{Name: "g_array_new_take"}
	g_array_new_take_zero_terminated = cc.DlFunc[func(data UPtr, clear bool, elementSize uint64) UPtr, UPtr]{Name: "g_array_new_take_zero_terminated"}
	g_array_steal                    = cc.DlFunc[func(array UPtr, length *uint64) UPtr, UPtr]{Name: "g_array_steal"}
	g_array_sized_new                = cc.DlFunc[func(zeroTerminated, clear bool, elementSize, reservedSize uint32) UPtr, UPtr]{Name: "g_array_sized_new"}
	g_array_copy                     = cc.DlFunc[func(array UPtr) UPtr, UPtr]{Name: "g_array_copy"}
	g_array_free                     = cc.DlFunc[func(array UPtr, freeSegment bool) UPtr, UPtr]{Name: "g_array_free"}
	g_array_ref                      = cc.DlFunc[func(array UPtr) UPtr, UPtr]{Name: "g_array_ref"}
	g_array_unref                    = cc.DlFunc[func(array UPtr), cc.Void]{Name: "g_array_unref"}
	g_array_get_element_size         = cc.DlFunc[func(array UPtr) uint32, uint32]{Name: "g_array_get_element_size"}
	g_array_append_vals              = cc.DlFunc[func(array UPtr, data UPtr, length uint32) UPtr, UPtr]{Name: "g_array_append_vals"}
	g_array_prepend_vals             = cc.DlFunc[func(array UPtr, data UPtr, length uint32) UPtr, UPtr]{Name: "g_array_prepend_vals"}
	g_array_insert_vals              = cc.DlFunc[func(array UPtr, index uint32, data UPtr, length uint32) UPtr, UPtr]{Name: "g_array_insert_vals"}
	g_array_set_size                 = cc.DlFunc[func(array UPtr, length uint32) UPtr, UPtr]{Name: "g_array_set_size"}
	g_array_remove_index             = cc.DlFunc[func(array UPtr, index uint32) UPtr, UPtr]{Name: "g_array_remove_index"}
	g_array_remove_index_fast        = cc.DlFunc[func(array UPtr, index uint32) UPtr, UPtr]{Name: "g_array_remove_index_fast"}
	g_array_remove_range             = cc.DlFunc[func(array UPtr, index uint32, length uint32) UPtr, UPtr]{Name: "g_array_remove_range"}
	g_array_sort                     = cc.DlFunc[func(array UPtr, compareFunc UPtr), cc.Void]{Name: "g_array_sort"}
	g_array_binary_search            = cc.DlFunc[func(array UPtr, target UPtr, compareFunc UPtr, outMatchIndex *uint32) bool, bool]{Name: "g_array_binary_search"}
	g_array_set_clear_func           = cc.DlFunc[func(array UPtr, clearFunc cc.Func), cc.Void]{Name: "g_array_set_clear_func"}

	g_ptr_array_new                            = cc.DlFunc[func() UPtr, UPtr]{Name: "g_ptr_array_new"}
	g_ptr_array_new_with_free_func             = cc.DlFunc[func(elementFreeFunc cc.Func) UPtr, UPtr]{Name: "g_ptr_array_new_with_free_func"}
	g_ptr_array_new_take                       = cc.DlFunc[func(data UPtr, length uint64, elementFreeFunc cc.Func) UPtr, UPtr]{Name: "g_ptr_array_new_take"}
	g_ptr_array_new_from_array                 = cc.DlFunc[func(data UPtr, length uint64, copyFunc UPtr, copyFuncUserData UPtr, elementFreeFunc cc.Func) UPtr, UPtr]{Name: "g_ptr_array_new_from_array"}
	g_ptr_array_steal                          = cc.DlFunc[func(array UPtr, length *uint64) UPtr, UPtr]{Name: "g_ptr_array_steal"}
	g_ptr_array_copy                           = cc.DlFunc[func(array UPtr, copyFunc UPtr, userData UPtr) UPtr, UPtr]{Name: "g_ptr_array_copy"}
	g_ptr_array_sized_new                      = cc.DlFunc[func(reservedSize uint32) UPtr, UPtr]{Name: "g_ptr_array_sized_new"}
	g_ptr_array_new_full                       = cc.DlFunc[func(reservedSize uint32, elementFreeFunc cc.Func) UPtr, UPtr]{Name: "g_ptr_array_new_full"}
	g_ptr_array_new_null_terminated            = cc.DlFunc[func(reservedSize uint32, elementFreeFunc cc.Func, nullTerminated bool) UPtr, UPtr]{Name: "g_ptr_array_new_null_terminated"}
	g_ptr_array_new_take_null_terminated       = cc.DlFunc[func(data UPtr, elementFreeFunc cc.Func) UPtr, UPtr]{Name: "g_ptr_array_new_take_null_terminated"}
	g_ptr_array_new_from_null_terminated_array = cc.DlFunc[func(data UPtr, copyFunc UPtr, copyFuncUserData UPtr, elementFreeFunc cc.Func) UPtr, UPtr]{Name: "g_ptr_array_new_from_null_terminated_array"}
	g_ptr_array_free                           = cc.DlFunc[func(array UPtr, freeSegment bool) UPtr, UPtr]{Name: "g_ptr_array_free"}
	g_ptr_array_ref                            = cc.DlFunc[func(array UPtr) UPtr, UPtr]{Name: "g_ptr_array_ref"}
	g_ptr_array_unref                          = cc.DlFunc[func(array UPtr), cc.Void]{Name: "g_ptr_array_unref"}
	g_ptr_array_set_free_func                  = cc.DlFunc[func(array UPtr, elementFreeFunc cc.Func), cc.Void]{Name: "g_ptr_array_set_free_func"}
	g_ptr_array_set_size                       = cc.DlFunc[func(array UPtr, length int32), cc.Void]{Name: "g_ptr_array_set_size"}
	g_ptr_array_remove_index                   = cc.DlFunc[func(array UPtr, index uint32) UPtr, UPtr]{Name: "g_ptr_array_remove_index"}
	g_ptr_array_remove_index_fast              = cc.DlFunc[func(array UPtr, index uint32) UPtr, UPtr]{Name: "g_ptr_array_remove_index_fast"}
	g_ptr_array_steal_index                    = cc.DlFunc[func(array UPtr, index uint32) UPtr, UPtr]{Name: "g_ptr_array_steal_index"}
	g_ptr_array_steal_index_fast               = cc.DlFunc[func(array UPtr, index uint32) UPtr, UPtr]{Name: "g_ptr_array_steal_index_fast"}
	g_ptr_array_remove                         = cc.DlFunc[func(array UPtr, data UPtr) bool, bool]{Name: "g_ptr_array_remove"}
	g_ptr_array_remove_fast                    = cc.DlFunc[func(array UPtr, data UPtr) bool, bool]{Name: "g_ptr_array_remove_fast"}
	g_ptr_array_remove_range                   = cc.DlFunc[func(array UPtr, index uint32, length uint32) UPtr, UPtr]{Name: "g_ptr_array_remove_range"}
	g_ptr_array_add                            = cc.DlFunc[func(array UPtr, data UPtr), cc.Void]{Name: "g_ptr_array_add"}
	g_ptr_array_extend                         = cc.DlFunc[func(arrayToExtend, array UPtr, copyFunc UPtr, userData UPtr), cc.Void]{Name: "g_ptr_array_extend"}
	g_ptr_array_extend_and_steal               = cc.DlFunc[func(arrayToExtend, array UPtr), cc.Void]{Name: "g_ptr_array_extend_and_steal"}
	g_ptr_array_insert                         = cc.DlFunc[func(array UPtr, index int32, data UPtr), cc.Void]{Name: "g_ptr_array_insert"}
	g_ptr_array_sort                           = cc.DlFunc[func(array UPtr, compareFunc UPtr), cc.Void]{Name: "g_ptr_array_sort"}
	g_ptr_array_sort_with_data                 = cc.DlFunc[func(array UPtr, compareFunc UPtr, userData UPtr), cc.Void]{Name: "g_ptr_array_sort_with_data"}
	g_ptr_array_sort_values                    = cc.DlFunc[func(array UPtr, compareFunc UPtr), cc.Void]{Name: "g_ptr_array_sort_values"}
	g_ptr_array_sort_values_with_data          = cc.DlFunc[func(array UPtr, compareFunc UPtr, userData UPtr), cc.Void]{Name: "g_ptr_array_sort_values_with_data"}
	g_ptr_array_foreach                        = cc.DlFunc[func(array UPtr, fn UPtr, userData UPtr), cc.Void]{Name: "g_ptr_array_foreach"}
	g_ptr_array_find                           = cc.DlFunc[func(haystack UPtr, needle UPtr, index *uint32) bool, bool]{Name: "g_ptr_array_find"}
	g_ptr_array_find_with_equal_func           = cc.DlFunc[func(haystack UPtr, needle UPtr, equalFunc UPtr, index *uint32) bool, bool]{Name: "g_ptr_array_find_with_equal_func"}
	g_ptr_array_is_null_terminated             = cc.DlFunc[func(array UPtr) bool, bool]{Name: "g_ptr_array_is_null_terminated"}

	g_byte_array_new               = cc.DlFunc[func() *GByteArray, *GByteArray]{Name: "g_byte_array_new"}
	g_byte_array_new_take          = cc.DlFunc[func(data UPtr, length uint64) *GByteArray, *GByteArray]{Name: "g_byte_array_new_take"}
	g_byte_array_steal             = cc.DlFunc[func(array *GByteArray, length *uint64) UPtr, UPtr]{Name: "g_byte_array_steal"}
	g_byte_array_sized_new         = cc.DlFunc[func(reservedSize uint32) *GByteArray, *GByteArray]{Name: "g_byte_array_sized_new"}
	g_byte_array_free              = cc.DlFunc[func(array *GByteArray, freeSegment bool) *byte, *byte]{Name: "g_byte_array_free"}
	g_byte_array_free_to_bytes     = cc.DlFunc[func(array *GByteArray) *GBytes, *GBytes]{Name: "g_byte_array_free_to_bytes"}
	g_byte_array_ref               = cc.DlFunc[func(array *GByteArray) *GByteArray, *GByteArray]{Name: "g_byte_array_ref"}
	g_byte_array_unref             = cc.DlFunc[func(array *GByteArray), cc.Void]{Name: "g_byte_array_unref"}
	g_byte_array_append            = cc.DlFunc[func(array *GByteArray, data UPtr, length uint32) *GByteArray, *GByteArray]{Name: "g_byte_array_append"}
	g_byte_array_prepend           = cc.DlFunc[func(array *GByteArray, data UPtr, length uint32) *GByteArray, *GByteArray]{Name: "g_byte_array_prepend"}
	g_byte_array_set_size          = cc.DlFunc[func(array *GByteArray, length uint32) *GByteArray, *GByteArray]{Name: "g_byte_array_set_size"}
	g_byte_array_remove_index      = cc.DlFunc[func(array *GByteArray, index uint32) *GByteArray, *GByteArray]{Name: "g_byte_array_remove_index"}
	g_byte_array_remove_index_fast = cc.DlFunc[func(array *GByteArray, index uint32) *GByteArray, *GByteArray]{Name: "g_byte_array_remove_index_fast"}
	g_byte_array_remove_range      = cc.DlFunc[func(array *GByteArray, index uint32, length uint32) *GByteArray, *GByteArray]{Name: "g_byte_array_remove_range"}
	g_byte_array_sort              = cc.DlFunc[func(array *GByteArray, compareFunc UPtr), cc.Void]{Name: "g_byte_array_sort"}
	g_byte_array_sort_with_data    = cc.DlFunc[func(array *GByteArray, compareFunc UPtr, userData UPtr), cc.Void]{Name: "g_byte_array_sort_with_data"}
	// #endregion

	// #region AsyncQueue
	g_async_queue_new                  = cc.DlFunc[func() *GAsyncQueue, *GAsyncQueue]{Name: "g_async_queue_new"}
	g_async_queue_new_full             = cc.DlFunc[func(itemFreeFunc cc.Func) *GAsyncQueue, *GAsyncQueue]{Name: "g_async_queue_new_full"}
	g_async_queue_lock                 = cc.DlFunc[func(q *GAsyncQueue), cc.Void]{Name: "g_async_queue_lock"}
	g_async_queue_unlock               = cc.DlFunc[func(q *GAsyncQueue), cc.Void]{Name: "g_async_queue_unlock"}
	g_async_queue_ref                  = cc.DlFunc[func(q *GAsyncQueue) *GAsyncQueue, *GAsyncQueue]{Name: "g_async_queue_ref"}
	g_async_queue_unref                = cc.DlFunc[func(q *GAsyncQueue), cc.Void]{Name: "g_async_queue_unref"}
	g_async_queue_push                 = cc.DlFunc[func(q *GAsyncQueue, data UPtr), cc.Void]{Name: "g_async_queue_push"}
	g_async_queue_push_unlocked        = cc.DlFunc[func(q *GAsyncQueue, data UPtr), cc.Void]{Name: "g_async_queue_push_unlocked"}
	g_async_queue_push_sorted          = cc.DlFunc[func(q *GAsyncQueue, data UPtr, cmp UPtr, userData UPtr), cc.Void]{Name: "g_async_queue_push_sorted"}
	g_async_queue_push_sorted_unlocked = cc.DlFunc[func(q *GAsyncQueue, data UPtr, cmp UPtr, userData UPtr), cc.Void]{Name: "g_async_queue_push_sorted_unlocked"}
	g_async_queue_pop                  = cc.DlFunc[func(q *GAsyncQueue) UPtr, UPtr]{Name: "g_async_queue_pop"}
	g_async_queue_pop_unlocked         = cc.DlFunc[func(q *GAsyncQueue) UPtr, UPtr]{Name: "g_async_queue_pop_unlocked"}
	g_async_queue_try_pop              = cc.DlFunc[func(q *GAsyncQueue) UPtr, UPtr]{Name: "g_async_queue_try_pop"}
	g_async_queue_try_pop_unlocked     = cc.DlFunc[func(q *GAsyncQueue) UPtr, UPtr]{Name: "g_async_queue_try_pop_unlocked"}
	g_async_queue_timeout_pop          = cc.DlFunc[func(q *GAsyncQueue, timeout uint64) UPtr, UPtr]{Name: "g_async_queue_timeout_pop"}
	g_async_queue_timeout_pop_unlocked = cc.DlFunc[func(q *GAsyncQueue, timeout uint64) UPtr, UPtr]{Name: "g_async_queue_timeout_pop_unlocked"}
	g_async_queue_length               = cc.DlFunc[func(q *GAsyncQueue) int32, int32]{Name: "g_async_queue_length"}
	g_async_queue_length_unlocked      = cc.DlFunc[func(q *GAsyncQueue) int32, int32]{Name: "g_async_queue_length_unlocked"}
	g_async_queue_sort                 = cc.DlFunc[func(q *GAsyncQueue, cmp UPtr, userData UPtr), cc.Void]{Name: "g_async_queue_sort"}
	g_async_queue_sort_unlocked        = cc.DlFunc[func(q *GAsyncQueue, cmp UPtr, userData UPtr), cc.Void]{Name: "g_async_queue_sort_unlocked"}
	g_async_queue_remove               = cc.DlFunc[func(q *GAsyncQueue, item UPtr) bool, bool]{Name: "g_async_queue_remove"}
	g_async_queue_remove_unlocked      = cc.DlFunc[func(q *GAsyncQueue, item UPtr) bool, bool]{Name: "g_async_queue_remove_unlocked"}
	g_async_queue_push_front           = cc.DlFunc[func(q *GAsyncQueue, item UPtr), cc.Void]{Name: "g_async_queue_push_front"}
	g_async_queue_push_front_unlocked  = cc.DlFunc[func(q *GAsyncQueue, item UPtr), cc.Void]{Name: "g_async_queue_push_front_unlocked"}
	// #endregion

	// #region Bytes
	g_bytes_new            = cc.DlFunc[func(data *byte, size uint64) *GBytes, *GBytes]{Name: "g_bytes_new"}
	g_bytes_new_static     = cc.DlFunc[func(data *byte, size uint64) *GBytes, *GBytes]{Name: "g_bytes_new_static"}
	g_bytes_compare        = cc.DlFunc[func(*GBytes, *GBytes) int, int]{Name: "g_bytes_compare"}
	g_bytes_equal          = cc.DlFunc[func(*GBytes, *GBytes) bool, bool]{Name: "g_bytes_equal"}
	g_bytes_get_data       = cc.DlFunc[func(*GBytes, *uint64) *byte, *byte]{Name: "g_bytes_get_data"}
	g_bytes_get_region     = cc.DlFunc[func(*GBytes, uint64, uint64) *byte, *byte]{Name: "g_bytes_get_region"}
	g_bytes_get_size       = cc.DlFunc[func(*GBytes) uint64, uint64]{Name: "g_bytes_get_size"}
	g_bytes_hash           = cc.DlFunc[func(*GBytes) uint32, uint32]{Name: "g_bytes_hash"}
	g_bytes_new_from_bytes = cc.DlFunc[func(*GBytes, uint64, uint64) *GBytes, *GBytes]{Name: "g_bytes_new_from_bytes"}
	g_bytes_ref            = cc.DlFunc[func(*GBytes) *GBytes, *GBytes]{Name: "g_bytes_ref"}
	g_bytes_unref          = cc.DlFunc[func(*GBytes), cc.Void]{Name: "g_bytes_unref"}
	g_bytes_unref_to_array = cc.DlFunc[func(*GBytes) *GByteArray, *GByteArray]{Name: "g_bytes_unref_to_array"}
	g_bytes_unref_to_data  = cc.DlFunc[func(*GBytes, *uint64) *byte, *byte]{Name: "g_bytes_unref_to_data"}
	// #endregion

	// #region DateTime
	g_date_time_unref                     = cc.DlFunc[func(*GDateTime), cc.Void]{Name: "g_date_time_unref"}
	g_date_time_ref                       = cc.DlFunc[func(*GDateTime) *GDateTime, *GDateTime]{Name: "g_date_time_ref"}
	g_date_time_new_now                   = cc.DlFunc[func(*GTimeZone) *GDateTime, *GDateTime]{Name: "g_date_time_new_now"}
	g_date_time_new_now_local             = cc.DlFunc[func() *GDateTime, *GDateTime]{Name: "g_date_time_new_now_local"}
	g_date_time_new_now_utc               = cc.DlFunc[func() *GDateTime, *GDateTime]{Name: "g_date_time_new_now_utc"}
	g_date_time_new_from_unix_local       = cc.DlFunc[func(int64) *GDateTime, *GDateTime]{Name: "g_date_time_new_from_unix_local"}
	g_date_time_new_from_unix_utc         = cc.DlFunc[func(int64) *GDateTime, *GDateTime]{Name: "g_date_time_new_from_unix_utc"}
	g_date_time_new_from_unix_local_usec  = cc.DlFunc[func(int64) *GDateTime, *GDateTime]{Name: "g_date_time_new_from_unix_local_usec"}         // GLIB_AVAILABLE_IN_2_80
	g_date_time_new_from_unix_utc_usec    = cc.DlFunc[func(int64) *GDateTime, *GDateTime]{Name: "g_date_time_new_from_unix_utc_usec"}           // GLIB_AVAILABLE_IN_2_80
	g_date_time_new_from_iso8601          = cc.DlFunc[func(cc.String, *GTimeZone) *GDateTime, *GDateTime]{Name: "g_date_time_new_from_iso8601"} // GLIB_AVAILABLE_IN_2_56
	g_date_time_new                       = cc.DlFunc[func(*GTimeZone, int, int, int, int, int, float64) *GDateTime, *GDateTime]{Name: "g_date_time_new"}
	g_date_time_new_local                 = cc.DlFunc[func(int, int, int, int, int, float64) *GDateTime, *GDateTime]{Name: "g_date_time_new_local"}
	g_date_time_new_utc                   = cc.DlFunc[func(int, int, int, int, int, float64) *GDateTime, *GDateTime]{Name: "g_date_time_new_utc"}
	g_date_time_add                       = cc.DlFunc[func(*GDateTime, GTimeSpan) *GDateTime, *GDateTime]{Name: "g_date_time_add"}
	g_date_time_add_years                 = cc.DlFunc[func(*GDateTime, int) *GDateTime, *GDateTime]{Name: "g_date_time_add_years"}
	g_date_time_add_months                = cc.DlFunc[func(*GDateTime, int) *GDateTime, *GDateTime]{Name: "g_date_time_add_months"}
	g_date_time_add_weeks                 = cc.DlFunc[func(*GDateTime, int) *GDateTime, *GDateTime]{Name: "g_date_time_add_weeks"}
	g_date_time_add_days                  = cc.DlFunc[func(*GDateTime, int) *GDateTime, *GDateTime]{Name: "g_date_time_add_days"}
	g_date_time_add_hours                 = cc.DlFunc[func(*GDateTime, int) *GDateTime, *GDateTime]{Name: "g_date_time_add_hours"}
	g_date_time_add_minutes               = cc.DlFunc[func(*GDateTime, int) *GDateTime, *GDateTime]{Name: "g_date_time_add_minutes"}
	g_date_time_add_seconds               = cc.DlFunc[func(*GDateTime, float64) *GDateTime, *GDateTime]{Name: "g_date_time_add_seconds"}
	g_date_time_add_full                  = cc.DlFunc[func(*GDateTime, int, int, int, int, int, float64) *GDateTime, *GDateTime]{Name: "g_date_time_add_full"}
	g_date_time_compare                   = cc.DlFunc[func(*GDateTime, *GDateTime) int, int]{Name: "g_date_time_compare"} // gconstpointer 对应 cc.UIntPtr
	g_date_time_difference                = cc.DlFunc[func(*GDateTime, *GDateTime) GTimeSpan, GTimeSpan]{Name: "g_date_time_difference"}
	g_date_time_hash                      = cc.DlFunc[func(*GDateTime) uint32, uint32]{Name: "g_date_time_hash"}            // gconstpointer 对应 cc.UIntPtr
	g_date_time_equal                     = cc.DlFunc[func(*GDateTime, *GDateTime) int32, int32]{Name: "g_date_time_equal"} // gconstpointer 对应 cc.UIntPtr，gboolean 对应 int32
	g_date_time_get_ymd                   = cc.DlFunc[func(*GDateTime, *int, *int, *int), cc.Void]{Name: "g_date_time_get_ymd"}
	g_date_time_get_year                  = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_year"}
	g_date_time_get_month                 = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_month"}
	g_date_time_get_day_of_month          = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_day_of_month"}
	g_date_time_get_week_numbering_year   = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_week_numbering_year"}
	g_date_time_get_week_of_year          = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_week_of_year"}
	g_date_time_get_day_of_week           = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_day_of_week"}
	g_date_time_get_day_of_year           = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_day_of_year"}
	g_date_time_get_hour                  = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_hour"}
	g_date_time_get_minute                = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_minute"}
	g_date_time_get_second                = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_second"}
	g_date_time_get_microsecond           = cc.DlFunc[func(*GDateTime) int, int]{Name: "g_date_time_get_microsecond"}
	g_date_time_get_seconds               = cc.DlFunc[func(*GDateTime) float64, float64]{Name: "g_date_time_get_seconds"}
	g_date_time_to_unix                   = cc.DlFunc[func(*GDateTime) int64, int64]{Name: "g_date_time_to_unix"}
	g_date_time_to_unix_usec              = cc.DlFunc[func(*GDateTime) int64, int64]{Name: "g_date_time_to_unix_usec"} // GLIB_AVAILABLE_IN_2_80
	g_date_time_get_utc_offset            = cc.DlFunc[func(*GDateTime) GTimeSpan, GTimeSpan]{Name: "g_date_time_get_utc_offset"}
	g_date_time_get_timezone              = cc.DlFunc[func(*GDateTime) *GTimeZone, *GTimeZone]{Name: "g_date_time_get_timezone"} // GLIB_AVAILABLE_IN_2_58
	g_date_time_get_timezone_abbreviation = cc.DlFunc[func(*GDateTime) cc.String, cc.String]{Name: "g_date_time_get_timezone_abbreviation"}
	g_date_time_is_daylight_savings       = cc.DlFunc[func(*GDateTime) int32, int32]{Name: "g_date_time_is_daylight_savings"} // gboolean 对应 int32
	g_date_time_to_timezone               = cc.DlFunc[func(*GDateTime, *GTimeZone) *GDateTime, *GDateTime]{Name: "g_date_time_to_timezone"}
	g_date_time_to_local                  = cc.DlFunc[func(*GDateTime) *GDateTime, *GDateTime]{Name: "g_date_time_to_local"}
	g_date_time_to_utc                    = cc.DlFunc[func(*GDateTime) *GDateTime, *GDateTime]{Name: "g_date_time_to_utc"}
	g_date_time_format                    = cc.DlFunc[func(*GDateTime, cc.String) cc.String, cc.String]{Name: "g_date_time_format"}
	g_date_time_format_iso8601            = cc.DlFunc[func(*GDateTime) cc.String, cc.String]{Name: "g_date_time_format_iso8601"} // GLIB_AVAILABLE_IN_2_62
	// #endregion

	// #region Env
	g_getenv           = cc.DlFunc[func(cc.String) cc.String, cc.String]{Name: "g_getenv"}
	g_setenv           = cc.DlFunc[func(cc.String, cc.String, int32) int32, int32]{Name: "g_setenv"}
	g_unsetenv         = cc.DlFunc[func(cc.String), cc.Void]{Name: "g_unsetenv"}
	g_listenv          = cc.DlFunc[func() cc.Strings, cc.Strings]{Name: "g_listenv"}
	g_get_environ      = cc.DlFunc[func() cc.Strings, cc.Strings]{Name: "g_get_environ"}
	g_environ_getenv   = cc.DlFunc[func(cc.Strings, cc.String) cc.String, cc.String]{Name: "g_environ_getenv"}
	g_environ_setenv   = cc.DlFunc[func(cc.Strings, cc.String, cc.String, int32) cc.Strings, cc.Strings]{Name: "g_environ_setenv"}
	g_environ_unsetenv = cc.DlFunc[func(cc.Strings, cc.String) cc.Strings, cc.Strings]{Name: "g_environ_unsetenv"}
	// #endregion

	// #region Error
	g_error_new                    = cc.DlFunc[func(GQuark, int32, cc.String) *GError, *GError]{Name: "g_error_new"}
	g_error_domain_register        = cc.DlFunc[func(cc.String, uint, UPtr, UPtr, UPtr) GQuark, GQuark]{Name: "g_error_domain_register"}
	g_error_domain_register_static = cc.DlFunc[func(cc.String, uint, UPtr, UPtr, UPtr) GQuark, GQuark]{Name: "g_error_domain_register_static"}
	g_error_copy                   = cc.DlFunc[func(*GError) *GError, *GError]{Name: "g_error_copy"}
	g_error_free                   = cc.DlFunc[func(*GError), cc.Void]{Name: "g_error_free"}
	g_error_matches                = cc.DlFunc[func(*GError, GQuark, int32) bool, bool]{Name: "g_error_matches"}
	// #endregion

	// #region List
	g_list_alloc     = cc.DlFunc[func() UPtr, UPtr]{Name: "g_list_alloc"}
	g_list_free      = cc.DlFunc[func(UPtr), cc.Void]{Name: "g_list_free"}
	g_list_free_1    = cc.DlFunc[func(UPtr), cc.Void]{Name: "g_list_free_1"}
	g_list_free_full = cc.DlFunc[func(UPtr, UPtr), cc.Void]{Name: "g_list_free_full"}
	g_list_append    = cc.DlFunc[func(UPtr, UPtr) UPtr, UPtr]{Name: "g_list_append"}
	g_list_prepend   = cc.DlFunc[func(UPtr, UPtr) UPtr, UPtr]{Name: "g_list_prepend"}
	g_list_insert    = cc.DlFunc[func(UPtr, UPtr, int32) UPtr, UPtr]{Name: "g_list_insert"}

	g_list_insert_sorted           = cc.DlFunc[func(UPtr, UPtr, UPtr) UPtr, UPtr]{Name: "g_list_insert_sorted"}
	g_list_insert_sorted_with_data = cc.DlFunc[func(UPtr, UPtr, UPtr, UPtr) UPtr, UPtr]{Name: "g_list_insert_sorted_with_data"}
	g_list_insert_before           = cc.DlFunc[func(UPtr, UPtr, UPtr) UPtr, UPtr]{Name: "g_list_insert_before"}
	g_list_insert_before_link      = cc.DlFunc[func(UPtr, UPtr, UPtr) UPtr, UPtr]{Name: "g_list_insert_before_link"}
	g_list_concat                  = cc.DlFunc[func(UPtr, UPtr) UPtr, UPtr]{Name: "g_list_concat"}
	g_list_remove                  = cc.DlFunc[func(UPtr, UPtr) UPtr, UPtr]{Name: "g_list_remove"}
	g_list_remove_all              = cc.DlFunc[func(UPtr, UPtr) UPtr, UPtr]{Name: "g_list_remove_all"}
	g_list_remove_link             = cc.DlFunc[func(UPtr, UPtr) UPtr, UPtr]{Name: "g_list_remove_link"}
	g_list_delete_link             = cc.DlFunc[func(UPtr, UPtr) UPtr, UPtr]{Name: "g_list_delete_link"}

	g_list_reverse        = cc.DlFunc[func(UPtr) UPtr, UPtr]{Name: "g_list_reverse"}
	g_list_copy           = cc.DlFunc[func(UPtr) UPtr, UPtr]{Name: "g_list_copy"}
	g_list_copy_deep      = cc.DlFunc[func(UPtr, UPtr, UPtr) UPtr, UPtr]{Name: "g_list_copy_deep"}
	g_list_nth            = cc.DlFunc[func(UPtr, uint32) UPtr, UPtr]{Name: "g_list_nth"}
	g_list_nth_prev       = cc.DlFunc[func(UPtr, uint32) UPtr, UPtr]{Name: "g_list_nth_prev"}
	g_list_find           = cc.DlFunc[func(UPtr, UPtr) UPtr, UPtr]{Name: "g_list_find"}
	g_list_find_custom    = cc.DlFunc[func(UPtr, UPtr, UPtr) UPtr, UPtr]{Name: "g_list_find_custom"}
	g_list_position       = cc.DlFunc[func(UPtr, UPtr) int32, int32]{Name: "g_list_position"}
	g_list_index          = cc.DlFunc[func(UPtr, UPtr) int32, int32]{Name: "g_list_index"}
	g_list_last           = cc.DlFunc[func(UPtr) UPtr, UPtr]{Name: "g_list_last"}
	g_list_first          = cc.DlFunc[func(UPtr) UPtr, UPtr]{Name: "g_list_first"}
	g_list_length         = cc.DlFunc[func(UPtr) uint32, uint32]{Name: "g_list_length"}
	g_list_foreach        = cc.DlFunc[func(UPtr, UPtr, UPtr), cc.Void]{Name: "g_list_foreach"}
	g_list_sort           = cc.DlFunc[func(UPtr, UPtr) UPtr, UPtr]{Name: "g_list_sort"}
	g_list_sort_with_data = cc.DlFunc[func(UPtr, UPtr, UPtr) UPtr, UPtr]{Name: "g_list_sort_with_data"}
	g_list_nth_data       = cc.DlFunc[func(UPtr, uint32) UPtr, UPtr]{Name: "g_list_nth_data"}
	g_clear_list          = cc.DlFunc[func(UPtr, UPtr), cc.Void]{Name: "g_clear_list"}
	// #endregion

	// #region Main
	g_main_context_new                            = cc.DlFunc[func() *GMainContext, *GMainContext]{Name: "g_main_context_new"}
	g_main_context_new_with_flags                 = cc.DlFunc[func(flags GMainContextFlags) *GMainContext, *GMainContext]{Name: "g_main_context_new_with_flags"}
	g_main_context_ref                            = cc.DlFunc[func(context *GMainContext) *GMainContext, *GMainContext]{Name: "g_main_context_ref"}
	g_main_context_unref                          = cc.DlFunc[func(context *GMainContext), cc.Void]{Name: "g_main_context_unref"}
	g_main_context_default                        = cc.DlFunc[func() *GMainContext, *GMainContext]{Name: "g_main_context_default"}
	g_main_context_iteration                      = cc.DlFunc[func(context *GMainContext, mayBlock bool) bool, bool]{Name: "g_main_context_iteration"}
	g_main_context_pending                        = cc.DlFunc[func(context *GMainContext) bool, bool]{Name: "g_main_context_pending"}
	g_main_context_find_source_by_id              = cc.DlFunc[func(context *GMainContext, sourceID uint32) *GSource, *GSource]{Name: "g_main_context_find_source_by_id"}
	g_main_context_find_source_by_user_data       = cc.DlFunc[func(context *GMainContext, userData UPtr) *GSource, *GSource]{Name: "g_main_context_find_source_by_user_data"}
	g_main_context_find_source_by_funcs_user_data = cc.DlFunc[func(context *GMainContext, funcs *GSourceFuncs, userData UPtr) *GSource, *GSource]{Name: "g_main_context_find_source_by_funcs_user_data"}
	g_main_context_wakeup                         = cc.DlFunc[func(context *GMainContext), cc.Void]{Name: "g_main_context_wakeup"}
	g_main_context_acquire                        = cc.DlFunc[func(context *GMainContext) bool, bool]{Name: "g_main_context_acquire"}
	g_main_context_release                        = cc.DlFunc[func(context *GMainContext), cc.Void]{Name: "g_main_context_release"}
	g_main_context_is_owner                       = cc.DlFunc[func(context *GMainContext) bool, bool]{Name: "g_main_context_is_owner"}
	g_main_context_wait                           = cc.DlFunc[func(context *GMainContext, cond UPtr, mutex UPtr) bool, bool]{Name: "g_main_context_wait"}
	g_main_context_prepare                        = cc.DlFunc[func(context *GMainContext, priority *int32) bool, bool]{Name: "g_main_context_prepare"}
	g_main_context_query                          = cc.DlFunc[func(context *GMainContext, maxPriority int32, timeout *int32, fds *GPollFD, nFds int32) int32, int32]{Name: "g_main_context_query"}
	g_main_context_check                          = cc.DlFunc[func(context *GMainContext, maxPriority int32, fds *GPollFD, nFds int32) bool, bool]{Name: "g_main_context_check"}
	g_main_context_dispatch                       = cc.DlFunc[func(context *GMainContext), cc.Void]{Name: "g_main_context_dispatch"}
	g_main_context_set_poll_func                  = cc.DlFunc[func(context *GMainContext, fn UPtr), cc.Void]{Name: "g_main_context_set_poll_func"}
	g_main_context_get_poll_func                  = cc.DlFunc[func(context *GMainContext) UPtr, UPtr]{Name: "g_main_context_get_poll_func"}
	g_main_context_add_poll                       = cc.DlFunc[func(context *GMainContext, fd *GPollFD, priority int32), cc.Void]{Name: "g_main_context_add_poll"}
	g_main_context_remove_poll                    = cc.DlFunc[func(context *GMainContext, fd *GPollFD), cc.Void]{Name: "g_main_context_remove_poll"}
	g_main_depth                                  = cc.DlFunc[func() int32, int32]{Name: "g_main_depth"}
	g_main_current_source                         = cc.DlFunc[func() *GSource, *GSource]{Name: "g_main_current_source"}
	g_main_context_push_thread_default            = cc.DlFunc[func(context *GMainContext), cc.Void]{Name: "g_main_context_push_thread_default"}
	g_main_context_pop_thread_default             = cc.DlFunc[func(context *GMainContext), cc.Void]{Name: "g_main_context_pop_thread_default"}
	g_main_context_get_thread_default             = cc.DlFunc[func() *GMainContext, *GMainContext]{Name: "g_main_context_get_thread_default"}
	g_main_context_ref_thread_default             = cc.DlFunc[func() *GMainContext, *GMainContext]{Name: "g_main_context_ref_thread_default"}
	g_main_context_invoke_full                    = cc.DlFunc[func(context *GMainContext, priority int32, fn UPtr, data UPtr, notify UPtr), cc.Void]{Name: "g_main_context_invoke_full"}
	g_main_context_invoke                         = cc.DlFunc[func(context *GMainContext, fn UPtr, data UPtr), cc.Void]{Name: "g_main_context_invoke"}
	g_main_loop_new                               = cc.DlFunc[func(context *GMainContext, isRunning bool) *GMainLoop, *GMainLoop]{Name: "g_main_loop_new"}
	g_main_loop_run                               = cc.DlFunc[func(loop *GMainLoop), cc.Void]{Name: "g_main_loop_run"}
	g_main_loop_quit                              = cc.DlFunc[func(loop *GMainLoop), cc.Void]{Name: "g_main_loop_quit"}
	g_main_loop_ref                               = cc.DlFunc[func(loop *GMainLoop) *GMainLoop, *GMainLoop]{Name: "g_main_loop_ref"}
	g_main_loop_unref                             = cc.DlFunc[func(loop *GMainLoop), cc.Void]{Name: "g_main_loop_unref"}
	g_main_loop_is_running                        = cc.DlFunc[func(loop *GMainLoop) bool, bool]{Name: "g_main_loop_is_running"}
	g_main_loop_get_context                       = cc.DlFunc[func(loop *GMainLoop) *GMainContext, *GMainContext]{Name: "g_main_loop_get_context"}
	g_source_new                                  = cc.DlFunc[func(sourceFuncs *GSourceFuncs, structSize uint32) *GSource, *GSource]{Name: "g_source_new"}
	g_source_set_dispose_function                 = cc.DlFunc[func(source *GSource, dispose UPtr), cc.Void]{Name: "g_source_set_dispose_function"}
	g_source_ref                                  = cc.DlFunc[func(source *GSource) *GSource, *GSource]{Name: "g_source_ref"}
	g_source_unref                                = cc.DlFunc[func(source *GSource), cc.Void]{Name: "g_source_unref"}
	g_source_attach                               = cc.DlFunc[func(source *GSource, context *GMainContext) uint32, uint32]{Name: "g_source_attach"}
	g_source_destroy                              = cc.DlFunc[func(source *GSource), cc.Void]{Name: "g_source_destroy"}
	g_source_set_priority                         = cc.DlFunc[func(source *GSource, priority int32), cc.Void]{Name: "g_source_set_priority"}
	g_source_get_priority                         = cc.DlFunc[func(source *GSource) int32, int32]{Name: "g_source_get_priority"}
	g_source_set_can_recurse                      = cc.DlFunc[func(source *GSource, canRecurse bool), cc.Void]{Name: "g_source_set_can_recurse"}
	g_source_get_can_recurse                      = cc.DlFunc[func(source *GSource) bool, bool]{Name: "g_source_get_can_recurse"}
	g_source_get_id                               = cc.DlFunc[func(source *GSource) uint32, uint32]{Name: "g_source_get_id"}
	g_source_get_context                          = cc.DlFunc[func(source *GSource) *GMainContext, *GMainContext]{Name: "g_source_get_context"}
	g_source_set_callback                         = cc.DlFunc[func(source *GSource, fn UPtr, data UPtr, notify UPtr), cc.Void]{Name: "g_source_set_callback"}
	g_source_set_funcs                            = cc.DlFunc[func(source *GSource, funcs *GSourceFuncs), cc.Void]{Name: "g_source_set_funcs"}
	g_source_is_destroyed                         = cc.DlFunc[func(source *GSource) bool, bool]{Name: "g_source_is_destroyed"}
	g_source_set_name                             = cc.DlFunc[func(source *GSource, name cc.String), cc.Void]{Name: "g_source_set_name"}
	g_source_set_static_name                      = cc.DlFunc[func(source *GSource, name cc.String), cc.Void]{Name: "g_source_set_static_name"}
	g_source_get_name                             = cc.DlFunc[func(source *GSource) cc.String, cc.String]{Name: "g_source_get_name"}
	g_source_set_name_by_id                       = cc.DlFunc[func(tag uint32, name cc.String), cc.Void]{Name: "g_source_set_name_by_id"}
	g_source_set_ready_time                       = cc.DlFunc[func(source *GSource, readyTime int64), cc.Void]{Name: "g_source_set_ready_time"}
	g_source_get_ready_time                       = cc.DlFunc[func(source *GSource) int64, int64]{Name: "g_source_get_ready_time"}
	g_source_add_poll                             = cc.DlFunc[func(source *GSource, fd *GPollFD), cc.Void]{Name: "g_source_add_poll"}
	g_source_remove_poll                          = cc.DlFunc[func(source *GSource, fd *GPollFD), cc.Void]{Name: "g_source_remove_poll"}
	g_source_add_child_source                     = cc.DlFunc[func(source *GSource, child *GSource), cc.Void]{Name: "g_source_add_child_source"}
	g_source_remove_child_source                  = cc.DlFunc[func(source *GSource, child *GSource), cc.Void]{Name: "g_source_remove_child_source"}
	// g_source_get_current_time                     func(source *GSource, timeval *GTimeVal)

	g_source_get_time                  = cc.DlFunc[func(source *GSource) int64, int64]{Name: "g_source_get_time"}
	g_idle_source_new                  = cc.DlFunc[func() *GSource, *GSource]{Name: "g_idle_source_new"}
	g_child_watch_source_new           = cc.DlFunc[func(pid GPid) *GSource, *GSource]{Name: "g_child_watch_source_new"}
	g_timeout_source_new               = cc.DlFunc[func(interval uint32) *GSource, *GSource]{Name: "g_timeout_source_new"}
	g_timeout_source_new_seconds       = cc.DlFunc[func(interval uint32) *GSource, *GSource]{Name: "g_timeout_source_new_seconds"}
	g_source_remove                    = cc.DlFunc[func(tag uint32) bool, bool]{Name: "g_source_remove"}
	g_source_remove_by_user_data       = cc.DlFunc[func(userData UPtr) bool, bool]{Name: "g_source_remove_by_user_data"}
	g_source_remove_by_funcs_user_data = cc.DlFunc[func(funcs *GSourceFuncs, userData UPtr) bool, bool]{Name: "g_source_remove_by_funcs_user_data"}
	g_timeout_add_full                 = cc.DlFunc[func(priority int32, interval uint32, fn UPtr, data UPtr, notify UPtr) uint32, uint32]{Name: "g_timeout_add_full"}
	g_timeout_add                      = cc.DlFunc[func(interval uint32, fn UPtr, data UPtr) uint32, uint32]{Name: "g_timeout_add"}
	g_timeout_add_once                 = cc.DlFunc[func(interval uint32, fn UPtr, data UPtr) uint32, uint32]{Name: "g_timeout_add_once"}
	g_timeout_add_seconds_full         = cc.DlFunc[func(priority int32, interval uint32, fn UPtr, data UPtr, notify UPtr) uint32, uint32]{Name: "g_timeout_add_seconds_full"}
	g_timeout_add_seconds              = cc.DlFunc[func(interval uint32, fn UPtr, data UPtr) uint32, uint32]{Name: "g_timeout_add_seconds"}
	g_timeout_add_seconds_once         = cc.DlFunc[func(interval uint32, fn UPtr, data UPtr) uint32, uint32]{Name: "g_timeout_add_seconds_once"}
	g_child_watch_add_full             = cc.DlFunc[func(priority int32, pid GPid, fn UPtr, data UPtr, notify UPtr) uint32, uint32]{Name: "g_child_watch_add_full"}
	g_child_watch_add                  = cc.DlFunc[func(pid GPid, fn UPtr, data UPtr) uint32, uint32]{Name: "g_child_watch_add"}
	g_idle_add                         = cc.DlFunc[func(fn UPtr, data UPtr) uint32, uint32]{Name: "g_idle_add"}
	g_idle_add_full                    = cc.DlFunc[func(priority int32, fn UPtr, data UPtr, notify UPtr) uint32, uint32]{Name: "g_idle_add_full"}
	g_idle_add_once                    = cc.DlFunc[func(fn UPtr, data UPtr) uint32, uint32]{Name: "g_idle_add_once"}
	g_idle_remove_by_data              = cc.DlFunc[func(data UPtr) bool, bool]{Name: "g_idle_remove_by_data"}
	// #endregion

	// #region Mem
	g_free               = cc.DlFunc[func(UPtr), cc.Void]{Name: "g_free"}
	g_free_sized         = cc.DlFunc[func(UPtr, uint64), cc.Void]{Name: "g_free_sized"}
	g_clear_pointer      = cc.DlFunc[func(*UPtr, cc.Func), cc.Void]{Name: "g_clear_pointer"}
	g_malloc             = cc.DlFunc[func(uint64) UPtr, UPtr]{Name: "g_malloc"}
	g_malloc0            = cc.DlFunc[func(uint64) UPtr, UPtr]{Name: "g_malloc0"}
	g_realloc            = cc.DlFunc[func(UPtr, uint64) UPtr, UPtr]{Name: "g_realloc"}
	g_try_malloc         = cc.DlFunc[func(uint64) UPtr, UPtr]{Name: "g_try_malloc"}
	g_try_malloc0        = cc.DlFunc[func(uint64) UPtr, UPtr]{Name: "g_try_malloc0"}
	g_try_realloc        = cc.DlFunc[func(UPtr, uint64) UPtr, UPtr]{Name: "g_try_realloc"}
	g_malloc_n           = cc.DlFunc[func(uint64, uint64) UPtr, UPtr]{Name: "g_malloc_n"}
	g_malloc0_n          = cc.DlFunc[func(uint64, uint64) UPtr, UPtr]{Name: "g_malloc0_n"}
	g_realloc_n          = cc.DlFunc[func(UPtr, uint64, uint64) UPtr, UPtr]{Name: "g_realloc_n"}
	g_try_malloc_n       = cc.DlFunc[func(uint64, uint64) UPtr, UPtr]{Name: "g_try_malloc_n"}
	g_try_malloc0_n      = cc.DlFunc[func(uint64, uint64) UPtr, UPtr]{Name: "g_try_malloc0_n"}
	g_try_realloc_n      = cc.DlFunc[func(UPtr, uint64, uint64) UPtr, UPtr]{Name: "g_try_realloc_n"}
	g_aligned_alloc      = cc.DlFunc[func(uint64, uint64, uint64) UPtr, UPtr]{Name: "g_aligned_alloc"}
	g_aligned_alloc0     = cc.DlFunc[func(uint64, uint64, uint64) UPtr, UPtr]{Name: "g_aligned_alloc0"}
	g_aligned_free       = cc.DlFunc[func(UPtr), cc.Void]{Name: "g_aligned_free"}
	g_aligned_free_sized = cc.DlFunc[func(UPtr, uint64, uint64), cc.Void]{Name: "g_aligned_free_sized"}
	// #endregion

	// #region OptionContext
	g_option_context_new                        = cc.DlFunc[func(cc.String) *GOptionContext, *GOptionContext]{Name: "g_option_context_new"}
	g_option_context_set_summary                = cc.DlFunc[func(*GOptionContext, cc.String), cc.Void]{Name: "g_option_context_set_summary"}
	g_option_context_get_summary                = cc.DlFunc[func(*GOptionContext) cc.String, cc.String]{Name: "g_option_context_get_summary"}
	g_option_context_set_description            = cc.DlFunc[func(*GOptionContext, cc.String), cc.Void]{Name: "g_option_context_set_description"}
	g_option_context_get_description            = cc.DlFunc[func(*GOptionContext) cc.String, cc.String]{Name: "g_option_context_get_description"}
	g_option_context_free                       = cc.DlFunc[func(*GOptionContext), cc.Void]{Name: "g_option_context_free"}
	g_option_context_set_help_enabled           = cc.DlFunc[func(*GOptionContext, int32), cc.Void]{Name: "g_option_context_set_help_enabled"}
	g_option_context_get_help_enabled           = cc.DlFunc[func(*GOptionContext) int32, int32]{Name: "g_option_context_get_help_enabled"}
	g_option_context_set_ignore_unknown_options = cc.DlFunc[func(*GOptionContext, int32), cc.Void]{Name: "g_option_context_set_ignore_unknown_options"}
	g_option_context_get_ignore_unknown_options = cc.DlFunc[func(*GOptionContext) int32, int32]{Name: "g_option_context_get_ignore_unknown_options"}
	g_option_context_set_strict_posix           = cc.DlFunc[func(*GOptionContext, bool), cc.Void]{Name: "g_option_context_set_strict_posix"}
	g_option_context_get_strict_posix           = cc.DlFunc[func(*GOptionContext) int32, int32]{Name: "g_option_context_get_strict_posix"}
	g_option_context_add_main_entries           = cc.DlFunc[func(*GOptionContext, *GOptionEntry, cc.String), cc.Void]{Name: "g_option_context_add_main_entries"}
	g_option_context_parse                      = cc.DlFunc[func(*GOptionContext, *int32, *cc.Strings, **GError) int32, int32]{Name: "g_option_context_parse"}
	g_option_context_parse_strv                 = cc.DlFunc[func(*GOptionContext, *cc.Strings, **GError) int32, int32]{Name: "g_option_context_parse_strv"}
	g_option_context_set_translate_func         = cc.DlFunc[func(*GOptionContext, UPtr, UPtr, UPtr), cc.Void]{Name: "g_option_context_set_translate_func"}
	g_option_context_set_translation_domain     = cc.DlFunc[func(*GOptionContext, cc.String), cc.Void]{Name: "g_option_context_set_translation_domain"}
	g_option_context_add_group                  = cc.DlFunc[func(*GOptionContext, *GOptionGroup), cc.Void]{Name: "g_option_context_add_group"}
	g_option_context_set_main_group             = cc.DlFunc[func(*GOptionContext, *GOptionGroup), cc.Void]{Name: "g_option_context_set_main_group"}
	g_option_context_get_main_group             = cc.DlFunc[func(*GOptionContext) *GOptionGroup, *GOptionGroup]{Name: "g_option_context_get_main_group"}
	g_option_context_get_help                   = cc.DlFunc[func(*GOptionContext, int32, *GOptionGroup) cc.String, cc.String]{Name: "g_option_context_get_help"}

	g_option_group_new                    = cc.DlFunc[func(cc.String, cc.String, cc.String, UPtr, UPtr) *GOptionGroup, *GOptionGroup]{Name: "g_option_group_new"}
	g_option_group_set_parse_hooks        = cc.DlFunc[func(*GOptionGroup, cc.Func, cc.Func), cc.Void]{Name: "g_option_group_set_parse_hooks"}
	g_option_group_set_error_hook         = cc.DlFunc[func(*GOptionGroup, cc.Func), cc.Void]{Name: "g_option_group_set_error_hook"}
	g_option_group_ref                    = cc.DlFunc[func(*GOptionGroup) *GOptionGroup, *GOptionGroup]{Name: "g_option_group_ref"}
	g_option_group_unref                  = cc.DlFunc[func(*GOptionGroup), cc.Void]{Name: "g_option_group_unref"}
	g_option_group_add_entries            = cc.DlFunc[func(*GOptionGroup, *GOptionEntry), cc.Void]{Name: "g_option_group_add_entries"}
	g_option_group_set_translate_func     = cc.DlFunc[func(*GOptionGroup, UPtr, UPtr, UPtr), cc.Void]{Name: "g_option_group_set_translate_func"}
	g_option_group_set_translation_domain = cc.DlFunc[func(*GOptionGroup, cc.String), cc.Void]{Name: "g_option_group_set_translation_domain"}
	// #endregion

	// #region Quark
	g_quark_try_string         = cc.DlFunc[func(cc.String) GQuark, GQuark]{Name: "g_quark_try_string"}
	g_quark_from_static_string = cc.DlFunc[func(cc.String) GQuark, GQuark]{Name: "g_quark_from_static_string"}
	g_quark_from_string        = cc.DlFunc[func(cc.String) GQuark, GQuark]{Name: "g_quark_from_string"}
	g_quark_to_string          = cc.DlFunc[func(GQuark) cc.String, cc.String]{Name: "g_quark_to_string"}
	// #endregion

	// #region SList
	g_slist_alloc                   = cc.DlFunc[func() UPtr, UPtr]{Name: "g_slist_alloc"}
	g_slist_free                    = cc.DlFunc[func(list UPtr), cc.Void]{Name: "g_slist_free"}
	g_slist_free_1                  = cc.DlFunc[func(list UPtr), cc.Void]{Name: "g_slist_free_1"}
	g_slist_free_full               = cc.DlFunc[func(list UPtr, freeFunc UPtr), cc.Void]{Name: "g_slist_free_full"}
	g_slist_append                  = cc.DlFunc[func(list UPtr, data UPtr) UPtr, UPtr]{Name: "g_slist_append"}
	g_slist_prepend                 = cc.DlFunc[func(list UPtr, data UPtr) UPtr, UPtr]{Name: "g_slist_prepend"}
	g_slist_insert                  = cc.DlFunc[func(list UPtr, data UPtr, position int32) UPtr, UPtr]{Name: "g_slist_insert"}
	g_slist_insert_sorted           = cc.DlFunc[func(list UPtr, data UPtr, compareFunc UPtr) UPtr, UPtr]{Name: "g_slist_insert_sorted"}
	g_slist_insert_sorted_with_data = cc.DlFunc[func(list UPtr, data UPtr, compareDataFunc UPtr, userData UPtr) UPtr, UPtr]{Name: "g_slist_insert_sorted_with_data"}
	g_slist_insert_before           = cc.DlFunc[func(slist UPtr, sibling UPtr, data UPtr) UPtr, UPtr]{Name: "g_slist_insert_before"}
	g_slist_concat                  = cc.DlFunc[func(list1 UPtr, list2 UPtr) UPtr, UPtr]{Name: "g_slist_concat"}
	g_slist_remove                  = cc.DlFunc[func(list UPtr, data UPtr) UPtr, UPtr]{Name: "g_slist_remove"}
	g_slist_remove_all              = cc.DlFunc[func(list UPtr, data UPtr) UPtr, UPtr]{Name: "g_slist_remove_all"}
	g_slist_remove_link             = cc.DlFunc[func(list UPtr, link UPtr) UPtr, UPtr]{Name: "g_slist_remove_link"}
	g_slist_delete_link             = cc.DlFunc[func(list UPtr, link UPtr) UPtr, UPtr]{Name: "g_slist_delete_link"}
	g_slist_reverse                 = cc.DlFunc[func(list UPtr) UPtr, UPtr]{Name: "g_slist_reverse"}
	g_slist_copy                    = cc.DlFunc[func(list UPtr) UPtr, UPtr]{Name: "g_slist_copy"}
	g_slist_copy_deep               = cc.DlFunc[func(list UPtr, copyFunc UPtr, userData UPtr) UPtr, UPtr]{Name: "g_slist_copy_deep"}
	g_slist_nth                     = cc.DlFunc[func(list UPtr, n uint32) UPtr, UPtr]{Name: "g_slist_nth"}
	g_slist_find                    = cc.DlFunc[func(list UPtr, data UPtr) UPtr, UPtr]{Name: "g_slist_find"}
	g_slist_find_custom             = cc.DlFunc[func(list UPtr, data UPtr, compareFunc UPtr) UPtr, UPtr]{Name: "g_slist_find_custom"}
	g_slist_position                = cc.DlFunc[func(list UPtr, llink UPtr) int32, int32]{Name: "g_slist_position"}
	g_slist_index                   = cc.DlFunc[func(list UPtr, data UPtr) int32, int32]{Name: "g_slist_index"}
	g_slist_last                    = cc.DlFunc[func(list UPtr) UPtr, UPtr]{Name: "g_slist_last"}
	g_slist_length                  = cc.DlFunc[func(list UPtr) uint32, uint32]{Name: "g_slist_length"}
	g_slist_foreach                 = cc.DlFunc[func(list UPtr, fn UPtr, userData UPtr), cc.Void]{Name: "g_slist_foreach"}
	g_slist_sort                    = cc.DlFunc[func(list UPtr, compareFunc UPtr) UPtr, UPtr]{Name: "g_slist_sort"}
	g_slist_sort_with_data          = cc.DlFunc[func(list UPtr, compareDataFunc UPtr, userData UPtr) UPtr, UPtr]{Name: "g_slist_sort_with_data"}
	g_slist_nth_data                = cc.DlFunc[func(list UPtr, n uint32) UPtr, UPtr]{Name: "g_slist_nth_data"}
	g_clear_slist                   = cc.DlFunc[func(slistPtr *UPtr, destroy UPtr), cc.Void]{Name: "g_clear_slist"}
	// #endregion

	// #region String
	g_string_new                = cc.DlFunc[func(cc.String) *GString, *GString]{Name: "g_string_new"}
	g_string_new_len            = cc.DlFunc[func(cc.String, int) *GString, *GString]{Name: "g_string_new_len"}
	g_string_new_take           = cc.DlFunc[func(cc.String) *GString, *GString]{Name: "g_string_new_take"}
	g_string_sized_new          = cc.DlFunc[func(uint) *GString, *GString]{Name: "g_string_sized_new"}
	g_string_free               = cc.DlFunc[func(*GString, bool) UPtr, UPtr]{Name: "g_string_free"}
	g_string_free_and_steal     = cc.DlFunc[func(*GString) UPtr, UPtr]{Name: "g_string_free_and_steal"}
	g_string_free_to_bytes      = cc.DlFunc[func(*GString) *GBytes, *GBytes]{Name: "g_string_free_to_bytes"}
	g_string_equal              = cc.DlFunc[func(*GString, *GString) bool, bool]{Name: "g_string_equal"}
	g_string_hash               = cc.DlFunc[func(*GString) uint, uint]{Name: "g_string_hash"}
	g_string_assign             = cc.DlFunc[func(*GString, cc.String) *GString, *GString]{Name: "g_string_assign"}
	g_string_truncate           = cc.DlFunc[func(*GString, uint) *GString, *GString]{Name: "g_string_truncate"}
	g_string_set_size           = cc.DlFunc[func(*GString, uint) *GString, *GString]{Name: "g_string_set_size"}
	g_string_insert_len         = cc.DlFunc[func(*GString, int, cc.String, int) *GString, *GString]{Name: "g_string_insert_len"}
	g_string_append             = cc.DlFunc[func(*GString, cc.String) *GString, *GString]{Name: "g_string_append"}
	g_string_append_len         = cc.DlFunc[func(*GString, cc.String, int) *GString, *GString]{Name: "g_string_append_len"}
	g_string_append_c           = cc.DlFunc[func(*GString, byte) *GString, *GString]{Name: "g_string_append_c"}
	g_string_append_unichar     = cc.DlFunc[func(*GString, uint32) *GString, *GString]{Name: "g_string_append_unichar"}
	g_string_prepend            = cc.DlFunc[func(*GString, cc.String) *GString, *GString]{Name: "g_string_prepend"}
	g_string_prepend_c          = cc.DlFunc[func(*GString, byte) *GString, *GString]{Name: "g_string_prepend_c"}
	g_string_prepend_unichar    = cc.DlFunc[func(*GString, uint32) *GString, *GString]{Name: "g_string_prepend_unichar"}
	g_string_prepend_len        = cc.DlFunc[func(*GString, cc.String, int) *GString, *GString]{Name: "g_string_prepend_len"}
	g_string_insert             = cc.DlFunc[func(*GString, int, cc.String) *GString, *GString]{Name: "g_string_insert"}
	g_string_insert_c           = cc.DlFunc[func(*GString, int, byte) *GString, *GString]{Name: "g_string_insert_c"}
	g_string_insert_unichar     = cc.DlFunc[func(*GString, int, uint32) *GString, *GString]{Name: "g_string_insert_unichar"}
	g_string_overwrite          = cc.DlFunc[func(*GString, uint, cc.String) *GString, *GString]{Name: "g_string_overwrite"}
	g_string_overwrite_len      = cc.DlFunc[func(*GString, uint, cc.String, int) *GString, *GString]{Name: "g_string_overwrite_len"}
	g_string_erase              = cc.DlFunc[func(*GString, int, int) *GString, *GString]{Name: "g_string_erase"}
	g_string_replace            = cc.DlFunc[func(*GString, cc.String, cc.String, uint32) uint32, uint32]{Name: "g_string_replace"}
	g_string_ascii_down         = cc.DlFunc[func(*GString) *GString, *GString]{Name: "g_string_ascii_down"}
	g_string_ascii_up           = cc.DlFunc[func(*GString) *GString, *GString]{Name: "g_string_ascii_up"}
	g_string_printf             = cc.DlFunc[func(p *GString, f cc.String), cc.Void]{Name: "g_string_printf"}
	g_string_append_printf      = cc.DlFunc[func(p *GString, f cc.String), cc.Void]{Name: "g_string_append_printf"}
	g_string_append_uri_escaped = cc.DlFunc[func(*GString, cc.String, cc.String, bool) *GString, *GString]{Name: "g_string_append_uri_escaped"}
	g_string_down               = cc.DlFunc[func(*GString) *GString, *GString]{Name: "g_string_down"}
	g_string_up                 = cc.DlFunc[func(*GString) *GString, *GString]{Name: "g_string_up"}
	// #endregion

	// #region TimeZone
	g_time_zone_new              = cc.DlFunc[func(cc.String) *GTimeZone, *GTimeZone]{Name: "g_time_zone_new"}
	g_time_zone_new_identifier   = cc.DlFunc[func(cc.String) *GTimeZone, *GTimeZone]{Name: "g_time_zone_new_identifier"} // GLIB_AVAILABLE_IN_2_68
	g_time_zone_new_utc          = cc.DlFunc[func() *GTimeZone, *GTimeZone]{Name: "g_time_zone_new_utc"}
	g_time_zone_new_local        = cc.DlFunc[func() *GTimeZone, *GTimeZone]{Name: "g_time_zone_new_local"}
	g_time_zone_new_offset       = cc.DlFunc[func(int32) *GTimeZone, *GTimeZone]{Name: "g_time_zone_new_offset"} // GLIB_AVAILABLE_IN_2_58
	g_time_zone_ref              = cc.DlFunc[func(*GTimeZone) *GTimeZone, *GTimeZone]{Name: "g_time_zone_ref"}
	g_time_zone_unref            = cc.DlFunc[func(*GTimeZone), cc.Void]{Name: "g_time_zone_unref"}
	g_time_zone_find_interval    = cc.DlFunc[func(*GTimeZone, GTimeType, int64) int, int]{Name: "g_time_zone_find_interval"}
	g_time_zone_adjust_time      = cc.DlFunc[func(*GTimeZone, GTimeType, *int64) int, int]{Name: "g_time_zone_adjust_time"}
	g_time_zone_get_abbreviation = cc.DlFunc[func(*GTimeZone, int) cc.String, cc.String]{Name: "g_time_zone_get_abbreviation"}
	g_time_zone_get_offset       = cc.DlFunc[func(*GTimeZone, int) int32, int32]{Name: "g_time_zone_get_offset"}
	g_time_zone_is_dst           = cc.DlFunc[func(*GTimeZone, int) int32, int32]{Name: "g_time_zone_is_dst"}            // gboolean 对应 int32
	g_time_zone_get_identifier   = cc.DlFunc[func(*GTimeZone) cc.String, cc.String]{Name: "g_time_zone_get_identifier"} // GLIB_AVAILABLE_IN_2_58
	// #endregion

	// #region Variant
	g_variant_type_new            = cc.DlFunc[func(cc.String) *GVariantType, *GVariantType]{Name: "g_variant_type_new"}
	g_variant_type_new_array      = cc.DlFunc[func(*GVariantType) *GVariantType, *GVariantType]{Name: "g_variant_type_new_array"}
	g_variant_type_new_dict_entry = cc.DlFunc[func(*GVariantType, *GVariantType) *GVariantType, *GVariantType]{Name: "g_variant_type_new_dict_entry"}
	g_variant_type_new_maybe      = cc.DlFunc[func(*GVariantType) *GVariantType, *GVariantType]{Name: "g_variant_type_new_maybe"}
	g_variant_type_new_tuple      = cc.DlFunc[func(**GVariantType, int32) *GVariantType, *GVariantType]{Name: "g_variant_type_new_tuple"}

	g_variant_type_copy              = cc.DlFunc[func(*GVariantType) *GVariantType, *GVariantType]{Name: "g_variant_type_copy"}
	g_variant_type_dup_string        = cc.DlFunc[func(*GVariantType) cc.String, cc.String]{Name: "g_variant_type_dup_string"}
	g_variant_type_element           = cc.DlFunc[func(*GVariantType) *GVariantType, *GVariantType]{Name: "g_variant_type_element"}
	g_variant_type_equal             = cc.DlFunc[func(*GVariantType, *GVariantType) int32, int32]{Name: "g_variant_type_equal"}
	g_variant_type_first             = cc.DlFunc[func(*GVariantType) *GVariantType, *GVariantType]{Name: "g_variant_type_first"}
	g_variant_type_free              = cc.DlFunc[func(*GVariantType), cc.Void]{Name: "g_variant_type_free"}
	g_variant_type_get_string_length = cc.DlFunc[func(*GVariantType) uint, uint]{Name: "g_variant_type_get_string_length"}
	g_variant_type_hash              = cc.DlFunc[func(UPtr) uint32, uint32]{Name: "g_variant_type_hash"}
	g_variant_type_is_array          = cc.DlFunc[func(*GVariantType) int32, int32]{Name: "g_variant_type_is_array"}
	g_variant_type_is_basic          = cc.DlFunc[func(*GVariantType) int32, int32]{Name: "g_variant_type_is_basic"}
	g_variant_type_is_container      = cc.DlFunc[func(*GVariantType) int32, int32]{Name: "g_variant_type_is_container"}
	g_variant_type_is_definite       = cc.DlFunc[func(*GVariantType) int32, int32]{Name: "g_variant_type_is_definite"}
	g_variant_type_is_dict_entry     = cc.DlFunc[func(*GVariantType) int32, int32]{Name: "g_variant_type_is_dict_entry"}
	g_variant_type_is_maybe          = cc.DlFunc[func(*GVariantType) int32, int32]{Name: "g_variant_type_is_maybe"}
	g_variant_type_is_subtype_of     = cc.DlFunc[func(*GVariantType, *GVariantType) int32, int32]{Name: "g_variant_type_is_subtype_of"}
	g_variant_type_is_tuple          = cc.DlFunc[func(*GVariantType) int32, int32]{Name: "g_variant_type_is_tuple"}
	g_variant_type_is_variant        = cc.DlFunc[func(*GVariantType) int32, int32]{Name: "g_variant_type_is_variant"}
	g_variant_type_key               = cc.DlFunc[func(*GVariantType) *GVariantType, *GVariantType]{Name: "g_variant_type_key"}
	g_variant_type_n_items           = cc.DlFunc[func(*GVariantType) uint, uint]{Name: "g_variant_type_n_items"}
	g_variant_type_next              = cc.DlFunc[func(*GVariantType) *GVariantType, *GVariantType]{Name: "g_variant_type_next"}
	g_variant_type_peek_string       = cc.DlFunc[func(*GVariantType) cc.String, cc.String]{Name: "g_variant_type_peek_string"}
	g_variant_type_value             = cc.DlFunc[func(*GVariantType) *GVariantType, *GVariantType]{Name: "g_variant_type_value"}

	g_variant_new                  = cc.DlFunc[func(cc.String, ...any) *GVariant, *GVariant]{Name: "g_variant_new", Va: true}
	g_variant_new_array            = cc.DlFunc[func(*GVariantType, **GVariant, int64) *GVariant, *GVariant]{Name: "g_variant_new_array"}
	g_variant_new_boolean          = cc.DlFunc[func(int32) *GVariant, *GVariant]{Name: "g_variant_new_boolean"}
	g_variant_new_byte             = cc.DlFunc[func(byte) *GVariant, *GVariant]{Name: "g_variant_new_byte"}
	g_variant_new_bytestring       = cc.DlFunc[func(cc.String) *GVariant, *GVariant]{Name: "g_variant_new_bytestring"}
	g_variant_new_bytestring_array = cc.DlFunc[func(cc.Strings, int64) *GVariant, *GVariant]{Name: "g_variant_new_bytestring_array"}
	g_variant_new_dict_entry       = cc.DlFunc[func(*GVariant, *GVariant) *GVariant, *GVariant]{Name: "g_variant_new_dict_entry"}
	g_variant_new_double           = cc.DlFunc[func(float64) *GVariant, *GVariant]{Name: "g_variant_new_double"}
	g_variant_new_fixed_array      = cc.DlFunc[func(*GVariantType, UPtr, uint64, uint64) *GVariant, *GVariant]{Name: "g_variant_new_fixed_array"}
	g_variant_new_from_bytes       = cc.DlFunc[func(*GVariantType, *GBytes, int32) *GVariant, *GVariant]{Name: "g_variant_new_from_bytes"}
	// g_variant_new_from_data        func(*GVariantType, UPtr, int64, bool, GDestroyNotify, UPtr) *GVariant
	g_variant_new_handle      = cc.DlFunc[func(int32) *GVariant, *GVariant]{Name: "g_variant_new_handle"}
	g_variant_new_int16       = cc.DlFunc[func(int16) *GVariant, *GVariant]{Name: "g_variant_new_int16"}
	g_variant_new_int32       = cc.DlFunc[func(int32) *GVariant, *GVariant]{Name: "g_variant_new_int32"}
	g_variant_new_int64       = cc.DlFunc[func(int64) *GVariant, *GVariant]{Name: "g_variant_new_int64"}
	g_variant_new_maybe       = cc.DlFunc[func(*GVariantType, *GVariant) *GVariant, *GVariant]{Name: "g_variant_new_maybe"}
	g_variant_new_object_path = cc.DlFunc[func(cc.String) *GVariant, *GVariant]{Name: "g_variant_new_object_path"}
	g_variant_new_objv        = cc.DlFunc[func(cc.Strings, int64) *GVariant, *GVariant]{Name: "g_variant_new_objv"}
	g_variant_new_parsed      = cc.DlFunc[func(cc.String) *GVariant, *GVariant]{Name: "g_variant_new_parsed"}
	// g_variant_new_parsed_va       func(string, va_list) *GVariant
	g_variant_new_printf    = cc.DlFunc[func(cc.String) *GVariant, *GVariant]{Name: "g_variant_new_printf"}
	g_variant_new_signature = cc.DlFunc[func(cc.String) *GVariant, *GVariant]{Name: "g_variant_new_signature"}
	g_variant_new_string    = cc.DlFunc[func(cc.String) *GVariant, *GVariant]{Name: "g_variant_new_string"}
	g_variant_new_strv      = cc.DlFunc[func(cc.Strings, int64) *GVariant, *GVariant]{Name: "g_variant_new_strv"}
	// g_variant_new_take_string func(string) *GVariant
	g_variant_new_tuple   = cc.DlFunc[func(**GVariant, uint64) *GVariant, *GVariant]{Name: "g_variant_new_tuple"}
	g_variant_new_uint16  = cc.DlFunc[func(uint16) *GVariant, *GVariant]{Name: "g_variant_new_uint16"}
	g_variant_new_uint32  = cc.DlFunc[func(uint32) *GVariant, *GVariant]{Name: "g_variant_new_uint32"}
	g_variant_new_uint64  = cc.DlFunc[func(uint64) *GVariant, *GVariant]{Name: "g_variant_new_uint64"}
	g_variant_new_variant = cc.DlFunc[func(*GVariant) *GVariant, *GVariant]{Name: "g_variant_new_variant"}

	g_variant_byteswap             = cc.DlFunc[func(*GVariant) *GVariant, *GVariant]{Name: "g_variant_byteswap"}
	g_variant_check_format_string  = cc.DlFunc[func(*GVariant, cc.String, int32) int32, int32]{Name: "g_variant_check_format_string"}
	g_variant_classify             = cc.DlFunc[func(*GVariant) GVariantClass, GVariantClass]{Name: "g_variant_classify"}
	g_variant_compare              = cc.DlFunc[func(*GVariant, *GVariant) int, int]{Name: "g_variant_compare"}
	g_variant_dup_bytestring       = cc.DlFunc[func(*GVariant, *uint64) cc.String, cc.String]{Name: "g_variant_dup_bytestring"}
	g_variant_dup_bytestring_array = cc.DlFunc[func(*GVariant, *uint64) cc.Strings, cc.Strings]{Name: "g_variant_dup_bytestring_array"}
	g_variant_dup_objv             = cc.DlFunc[func(*GVariant, *uint64) cc.Strings, cc.Strings]{Name: "g_variant_dup_objv"}
	g_variant_dup_string           = cc.DlFunc[func(*GVariant, *uint64) cc.String, cc.String]{Name: "g_variant_dup_string"}
	g_variant_dup_strv             = cc.DlFunc[func(*GVariant, *uint64) cc.Strings, cc.Strings]{Name: "g_variant_dup_strv"}
	g_variant_equal                = cc.DlFunc[func(*GVariant, *GVariant) int32, int32]{Name: "g_variant_equal"}
	g_variant_get                  = cc.DlFunc[func(*GVariant, cc.String, UPtr), cc.Void]{Name: "g_variant_get"}
	g_variant_get_boolean          = cc.DlFunc[func(*GVariant) int32, int32]{Name: "g_variant_get_boolean"}
	g_variant_get_byte             = cc.DlFunc[func(*GVariant) byte, byte]{Name: "g_variant_get_byte"}
	g_variant_get_bytestring       = cc.DlFunc[func(*GVariant) string, string]{Name: "g_variant_get_bytestring"}
	g_variant_get_bytestring_array = cc.DlFunc[func(*GVariant, *uint64) cc.Strings, cc.Strings]{Name: "g_variant_get_bytestring_array"}
	g_variant_get_child            = cc.DlFunc[func(*GVariant, uint64, cc.String, UPtr), cc.Void]{Name: "g_variant_get_child"}
	g_variant_get_child_value      = cc.DlFunc[func(*GVariant, uint64) *GVariant, *GVariant]{Name: "g_variant_get_child_value"}
	g_variant_get_data             = cc.DlFunc[func(*GVariant) UPtr, UPtr]{Name: "g_variant_get_data"}
	g_variant_get_data_as_bytes    = cc.DlFunc[func(*GVariant) *GBytes, *GBytes]{Name: "g_variant_get_data_as_bytes"}
	g_variant_get_double           = cc.DlFunc[func(*GVariant) float64, float64]{Name: "g_variant_get_double"}
	g_variant_get_fixed_array      = cc.DlFunc[func(*GVariant, *uint64, uint64) UPtr, UPtr]{Name: "g_variant_get_fixed_array"}
	g_variant_get_handle           = cc.DlFunc[func(*GVariant) int32, int32]{Name: "g_variant_get_handle"}
	g_variant_get_int16            = cc.DlFunc[func(*GVariant) int16, int16]{Name: "g_variant_get_int16"}
	g_variant_get_int32            = cc.DlFunc[func(*GVariant) int32, int32]{Name: "g_variant_get_int32"}
	g_variant_get_int64            = cc.DlFunc[func(*GVariant) int64, int64]{Name: "g_variant_get_int64"}
	g_variant_get_maybe            = cc.DlFunc[func(*GVariant) *GVariant, *GVariant]{Name: "g_variant_get_maybe"}
	g_variant_get_normal_form      = cc.DlFunc[func(*GVariant) *GVariant, *GVariant]{Name: "g_variant_get_normal_form"}
	g_variant_get_objv             = cc.DlFunc[func(*GVariant, *uint64) cc.Strings, cc.Strings]{Name: "g_variant_get_objv"}
	g_variant_get_size             = cc.DlFunc[func(*GVariant) uint, uint]{Name: "g_variant_get_size"}
	g_variant_get_string           = cc.DlFunc[func(*GVariant, *uint64) cc.String, cc.String]{Name: "g_variant_get_string"}
	g_variant_get_strv             = cc.DlFunc[func(*GVariant, *uint64) cc.Strings, cc.Strings]{Name: "g_variant_get_strv"}
	g_variant_get_type             = cc.DlFunc[func(*GVariant) *GVariantType, *GVariantType]{Name: "g_variant_get_type"}
	g_variant_get_type_string      = cc.DlFunc[func(*GVariant) cc.String, cc.String]{Name: "g_variant_get_type_string"}
	g_variant_get_uint16           = cc.DlFunc[func(*GVariant) uint16, uint16]{Name: "g_variant_get_uint16"}
	g_variant_get_uint32           = cc.DlFunc[func(*GVariant) uint32, uint32]{Name: "g_variant_get_uint32"}
	g_variant_get_uint64           = cc.DlFunc[func(*GVariant) uint64, uint64]{Name: "g_variant_get_uint64"}
	g_variant_get_va               = cc.DlFunc[func(*GVariant, string, *string, *[]interface{}), cc.Void]{Name: "g_variant_get_va"}
	g_variant_get_variant          = cc.DlFunc[func(*GVariant) *GVariant, *GVariant]{Name: "g_variant_get_variant"}
	g_variant_hash                 = cc.DlFunc[func(*GVariant) uint64, uint64]{Name: "g_variant_hash"}
	g_variant_is_container         = cc.DlFunc[func(*GVariant) int32, int32]{Name: "g_variant_is_container"}
	g_variant_is_floating          = cc.DlFunc[func(*GVariant) int32, int32]{Name: "g_variant_is_floating"}
	g_variant_is_normal_form       = cc.DlFunc[func(*GVariant) int32, int32]{Name: "g_variant_is_normal_form"}
	g_variant_is_of_type           = cc.DlFunc[func(*GVariant, *GVariantType) int32, int32]{Name: "g_variant_is_of_type"}
	g_variant_lookup               = cc.DlFunc[func(*GVariant, cc.String, cc.String, ...interface{}) int32, int32]{Name: "g_variant_lookup", Va: true}
	g_variant_lookup_value         = cc.DlFunc[func(*GVariant, cc.String, *GVariantType) *GVariant, *GVariant]{Name: "g_variant_lookup_value"}
	g_variant_n_children           = cc.DlFunc[func(*GVariant) uint64, uint64]{Name: "g_variant_n_children"}
	g_variant_print                = cc.DlFunc[func(*GVariant, int32) cc.String, cc.String]{Name: "g_variant_print"}
	g_variant_print_string         = cc.DlFunc[func(*GVariant, *GString, int32) *GString, *GString]{Name: "g_variant_print_string"}
	g_variant_ref                  = cc.DlFunc[func(*GVariant) *GVariant, *GVariant]{Name: "g_variant_ref"}
	g_variant_ref_sink             = cc.DlFunc[func(*GVariant) *GVariant, *GVariant]{Name: "g_variant_ref_sink"}
	g_variant_store                = cc.DlFunc[func(*GVariant, UPtr), cc.Void]{Name: "g_variant_store"}
	g_variant_take_ref             = cc.DlFunc[func(*GVariant) *GVariant, *GVariant]{Name: "g_variant_take_ref"}
	g_variant_unref                = cc.DlFunc[func(*GVariant), cc.Void]{Name: "g_variant_unref"}

	g_variant_iter_new        = cc.DlFunc[func(*GVariant) *GVariantIter, *GVariantIter]{Name: "g_variant_iter_new"}
	g_variant_iter_init       = cc.DlFunc[func(*GVariantIter, *GVariant) uint64, uint64]{Name: "g_variant_iter_init"}
	g_variant_iter_copy       = cc.DlFunc[func(*GVariantIter) *GVariantIter, *GVariantIter]{Name: "g_variant_iter_copy"}
	g_variant_iter_n_children = cc.DlFunc[func(*GVariantIter) uint64, uint64]{Name: "g_variant_iter_n_children"}
	g_variant_iter_free       = cc.DlFunc[func(*GVariantIter), cc.Void]{Name: "g_variant_iter_free"}
	g_variant_iter_next_value = cc.DlFunc[func(*GVariantIter) *GVariant, *GVariant]{Name: "g_variant_iter_next_value"}
	g_variant_iter_next       = cc.DlFunc[func(*GVariantIter, cc.String, ...interface{}) int32, int32]{Name: "g_variant_iter_next", Va: true}
	g_variant_iter_loop       = cc.DlFunc[func(*GVariantIter, cc.String, ...interface{}) int32, int32]{Name: "g_variant_iter_loop", Va: true}

	g_variant_builder_new         = cc.DlFunc[func(*GVariantType) *GVariantBuilder, *GVariantBuilder]{Name: "g_variant_builder_new"}
	g_variant_builder_unref       = cc.DlFunc[func(*GVariantBuilder), cc.Void]{Name: "g_variant_builder_unref"}
	g_variant_builder_ref         = cc.DlFunc[func(*GVariantBuilder) *GVariantBuilder, *GVariantBuilder]{Name: "g_variant_builder_ref"}
	g_variant_builder_init        = cc.DlFunc[func(*GVariantBuilder, *GVariantType), cc.Void]{Name: "g_variant_builder_init"}
	g_variant_builder_init_static = cc.DlFunc[func(*GVariantBuilder, *GVariantType), cc.Void]{Name: "g_variant_builder_init_static"}
	g_variant_builder_end         = cc.DlFunc[func(*GVariantBuilder) *GVariant, *GVariant]{Name: "g_variant_builder_end"}
	g_variant_builder_clear       = cc.DlFunc[func(*GVariantBuilder), cc.Void]{Name: "g_variant_builder_clear"}
	g_variant_builder_open        = cc.DlFunc[func(*GVariantBuilder, *GVariantType), cc.Void]{Name: "g_variant_builder_open"}
	g_variant_builder_close       = cc.DlFunc[func(*GVariantBuilder), cc.Void]{Name: "g_variant_builder_close"}
	g_variant_builder_add_value   = cc.DlFunc[func(*GVariantBuilder, *GVariant), cc.Void]{Name: "g_variant_builder_add_value"}
	g_variant_builder_add         = cc.DlFunc[func(*GVariantBuilder, cc.String, ...interface{}), cc.Void]{Name: "g_variant_builder_add", Va: true}
	g_variant_builder_add_parsed  = cc.DlFunc[func(*GVariantBuilder, cc.String, ...interface{}), cc.Void]{Name: "g_variant_builder_add_parsed", Va: true}

	g_variant_dict_new          = cc.DlFunc[func(*GVariant) *GVariantDict, *GVariantDict]{Name: "g_variant_dict_new"}
	g_variant_dict_init         = cc.DlFunc[func(*GVariantDict, *GVariant), cc.Void]{Name: "g_variant_dict_init"}
	g_variant_dict_lookup       = cc.DlFunc[func(*GVariantDict, cc.String, cc.String, ...interface{}) int32, int32]{Name: "g_variant_dict_lookup", Va: true}
	g_variant_dict_lookup_value = cc.DlFunc[func(*GVariantDict, cc.String, *GVariantType) *GVariant, *GVariant]{Name: "g_variant_dict_lookup_value"}
	g_variant_dict_contains     = cc.DlFunc[func(*GVariantDict, cc.String) int32, int32]{Name: "g_variant_dict_contains"}
	g_variant_dict_insert       = cc.DlFunc[func(*GVariantDict, cc.String, cc.String, ...interface{}), cc.Void]{Name: "g_variant_dict_insert", Va: true}
	g_variant_dict_insert_value = cc.DlFunc[func(*GVariantDict, cc.String, *GVariant), cc.Void]{Name: "g_variant_dict_insert_value"}
	g_variant_dict_remove       = cc.DlFunc[func(*GVariantDict, cc.String) int32, int32]{Name: "g_variant_dict_remove"}
	g_variant_dict_clear        = cc.DlFunc[func(*GVariantDict), cc.Void]{Name: "g_variant_dict_clear"}
	g_variant_dict_end          = cc.DlFunc[func(*GVariantDict) *GVariant, *GVariant]{Name: "g_variant_dict_end"}
	g_variant_dict_ref          = cc.DlFunc[func(*GVariantDict) *GVariantDict, *GVariantDict]{Name: "g_variant_dict_ref"}
	g_variant_dict_unref        = cc.DlFunc[func(*GVariantDict), cc.Void]{Name: "g_variant_dict_unref"}
	// #endregion

)
