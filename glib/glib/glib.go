package glib

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

type ints interface {
	~int8 | ~int16 | ~int32 | ~int | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint | ~uint64
}

func carry[T any, E ints](slice []T) (*T, E) {
	p, l := (*T)(nil), uint64(len(slice))
	if l > 0 {
		p = &slice[0]
	}
	return p, E(l)
}
func cbool(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// #region Alloc

func GAlloc[T any](size uint64) *T {
	return (*T)(malloc.Fn()(size))
}
func GAlloc0[T any](size uint64) *T {
	p := malloc.Fn()(size)
	memset.Fn()(p, 0, size)
	return (*T)(p)
}
func GNew[T any](a T) *T {
	p := GAlloc[T](uint64(unsafe.Sizeof(a)))
	(*p) = a
	return p
}
func GNew0[T any]() *T {
	var a T
	return GAlloc0[T](uint64(unsafe.Sizeof(a)))
}

// #endregion

// #region Array

type GArray[T any] struct {
	ptr uptr
	Len uint32
}
type GByteArray struct {
	data uptr
	Len  uint32
}
type GPtrArray[T any] struct {
	data uptr
	Len  uint32
}

func NewGArray[T any](zeroTerminated, clear bool) *GArray[T] {
	var t T
	return (*GArray[T])(g_array_new.Fn()(cbool(zeroTerminated), cbool(clear), uint32(unsafe.Sizeof(t))))
}
func NewGArrayCopy[T any](data []T) *GArray[T] {
	var t T

	size := unsafe.Sizeof(t)

	dt := (*struct {
		p uptr
		l int
	})((uptr)(&data))

	p := GAlloc[T](uint64(size * uintptr(dt.l)))
	copy(unsafe.Slice(p, dt.l), data)

	return (*GArray[T])(g_array_new_take.Fn()(uptr(p), uint64(dt.l), cbool(false), uint64(unsafe.Sizeof(t))))
}
func NewGArrayTake[T any](data uptr, length uint64, clear bool) *GArray[T] {
	var t T
	return (*GArray[T])(g_array_new_take.Fn()(data, length, cbool(clear), uint64(unsafe.Sizeof(t))))
}
func NewGArrayTakeZeroTerminated[T any](data uptr, clear bool) *GArray[T] {
	var t T
	return (*GArray[T])(g_array_new_take_zero_terminated.Fn()(data, cbool(clear), uint64(unsafe.Sizeof(t))))
}
func (array *GArray[T]) Steal() (data uptr, length uint64) {
	data = g_array_steal.Fn()(uptr(array), &length)
	return
}
func (array *GArray[T]) StealCopy() []T {
	var length uint64
	data := g_array_steal.Fn()(uptr(array), &length)
	if length == 0 {
		return nil
	}
	out := make([]T, length)
	copy(out, unsafe.Slice((*T)(data), length))
	return out
}
func (array *GArray[T]) List() []T {
	return unsafe.Slice((*T)(array.ptr), array.Len)
}
func NewGArraySized[T any](zeroTerminated, clear bool, reservedSize uint32) *GArray[T] {
	var t T
	return (*GArray[T])(g_array_sized_new.Fn()(cbool(zeroTerminated), cbool(clear), uint32(unsafe.Sizeof(t)), reservedSize))
}
func (array *GArray[T]) Copy() *GArray[T] { return (*GArray[T])(g_array_copy.Fn()(uptr(array))) }
func (array *GArray[T]) Free(freeSegment bool) uptr {
	return g_array_free.Fn()(uptr(array), cbool(freeSegment))
}
func (array *GArray[T]) Ref() *GArray[T] { return (*GArray[T])(g_array_ref.Fn()(uptr(array))) }
func (array *GArray[T]) Unref()          { g_array_unref.Fn()(uptr(array)) }
func (array *GArray[T]) GetElementSize() uint32 {
	return g_array_get_element_size.Fn()(uptr(array))
}
func (array *GArray[T]) Append(data []T) *GArray[T] {
	p, l := (*T)(nil), uint32(len(data))
	if l > 0 {
		p = &data[0]
	}
	return (*GArray[T])(g_array_append_vals.Fn()(uptr(array), uptr(p), l))
}
func (array *GArray[T]) Prepend(data []T) *GArray[T] {
	p, l := (*T)(nil), uint32(len(data))
	if l > 0 {
		p = &data[0]
	}
	return (*GArray[T])(g_array_prepend_vals.Fn()(uptr(array), uptr(p), l))
}
func (array *GArray[T]) Insert(index uint32, data []T) *GArray[T] {
	p, l := (*T)(nil), uint32(len(data))
	if l > 0 {
		p = &data[0]
	}
	return (*GArray[T])(g_array_insert_vals.Fn()(uptr(array), index, uptr(p), l))
}
func (array *GArray[T]) SetSize(length uint32) *GArray[T] {
	return (*GArray[T])(g_array_set_size.Fn()(uptr(array), length))
}
func (array *GArray[T]) RemoveIndex(index uint32) *GArray[T] {
	return (*GArray[T])(g_array_remove_index.Fn()(uptr(array), index))
}
func (array *GArray[T]) RemoveIndexFast(index uint32) *GArray[T] {
	return (*GArray[T])(g_array_remove_index_fast.Fn()(uptr(array), index))
}
func (array *GArray[T]) RemoveRange(index uint32, length uint32) *GArray[T] {
	return (*GArray[T])(g_array_remove_range.Fn()(uptr(array), index, length))
}
func (array *GArray[T]) Sort(compareFunc func(a, b *T) int32) {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	g_array_sort.Fn()(uptr(array), cf)
}
func (array *GArray[T]) BinarySearch(target *T, compareFunc func(curr, target *T) int32) (found bool, matchIndex uint32) {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	found = g_array_binary_search.Fn()(uptr(array), anyptr(target), cf, &matchIndex) != 0
	return
}
func (array *GArray[T]) SetClearFunc(clearFunc cc.Func /* clearFunc func(elmAddr *T) */) {
	g_array_set_clear_func.Fn()(uptr(array), clearFunc)
}

func (a *GPtrArray[T]) List() []*T {
	return unsafe.Slice((**T)(a.data), a.Len)
}
func NewGPtrArray[T any]() *GPtrArray[T] { return (*GPtrArray[T])(g_ptr_array_new.Fn()()) }
func NewGPtrArrayWithFreeFunc[T any](elmFreeFunc cc.Func /* elmFreeFunc func(data *T) */) *GPtrArray[T] {
	return (*GPtrArray[T])(g_ptr_array_new_with_free_func.Fn()(elmFreeFunc))
}
func NewGPtrArrayTake[T any](data *T, length uint64,
	elementFreeFunc cc.Func /* elementFreeFunc func(data UPtr) */) *GPtrArray[T] {
	return (*GPtrArray[T])(g_ptr_array_new_take.Fn()(uptr(data), length, elementFreeFunc))
}
func NewGPtrArrayFromArray[T any](arry *GArray[*T],
	copyFunc func(src *T, _ uptr) *T, elmFreeFunc cc.Func /* elmFreeFunc func(*T) */) *GPtrArray[T] {
	cf := cc.Cbk(copyFunc)
	defer cc.CbkClose(cf)
	return (*GPtrArray[T])(g_ptr_array_new_from_array.Fn()(arry.ptr, uint64(arry.Len), cf, nil, elmFreeFunc))
}
func (a *GPtrArray[T]) Steal() (data *T, length uint64) {
	data = (*T)(g_ptr_array_steal.Fn()(uptr(a), &length))
	return
}
func (a *GPtrArray[T]) Copy(copyFunc func(src *T, _ uptr) uptr) uptr /* *GPtrArray[T] */ {
	cf := cc.Cbk(copyFunc)
	defer cc.CbkClose(cf)
	return g_ptr_array_copy.Fn()(uptr(a), cf, nil)
}
func NewGPtrArraySized[T any](reservedSize uint32) *GPtrArray[T] {
	return (*GPtrArray[T])(g_ptr_array_sized_new.Fn()(reservedSize))
}
func NewGPtrArrayFull[T any](reservedSize uint32, elmFreeFunc cc.Func /* elementFreeFunc func(data *T) */) *GPtrArray[T] {
	return (*GPtrArray[T])(g_ptr_array_new_full.Fn()(reservedSize, elmFreeFunc))
}
func NewGPtrArrayNullTerminated[T any](reservedSize uint32,
	elmFreeFunc cc.Func /* elmFreeFunc func(data *T) */, nullTerminated bool) *GPtrArray[T] {
	return (*GPtrArray[T])(g_ptr_array_new_null_terminated.Fn()(reservedSize, elmFreeFunc, cbool(nullTerminated)))
}
func NewGPtrArrayTakeNullTerminated[T any](data uptr, elmFreeFunc cc.Func /* func(data *T) */) *GPtrArray[T] {
	return (*GPtrArray[T])(g_ptr_array_new_take_null_terminated.Fn()(data, elmFreeFunc))
}
func NewGPtrArrayFromNullTerminatedArray[T any](data uptr,
	copyFunc func(src *T, _ uptr) *T,
	elmFreeFunc cc.Func /* elmFreeFunc func(UPtr) UPtr */) *GPtrArray[T] {
	cf := cc.Cbk(copyFunc)
	defer cc.CbkClose(cf)

	return (*GPtrArray[T])(g_ptr_array_new_from_null_terminated_array.Fn()(data, cf, nil, elmFreeFunc))
}
func (a *GPtrArray[T]) Free(freeSegment bool) uptr {
	return g_ptr_array_free.Fn()(uptr(a), cbool(freeSegment))
}
func (a *GPtrArray[T]) Ref() *GPtrArray[T] { return (*GPtrArray[T])(g_ptr_array_ref.Fn()(uptr(a))) }
func (a *GPtrArray[T]) Unref()             { g_ptr_array_unref.Fn()(uptr(a)) }
func (a *GPtrArray[T]) SetFreeFunc(elmFreeFunc cc.Func /* elmFreeFunc func(UPtr */) {
	g_ptr_array_set_free_func.Fn()(uptr(a), elmFreeFunc)
}
func (a *GPtrArray[T]) SetSize(length int32) { g_ptr_array_set_size.Fn()(uptr(a), length) }
func (a *GPtrArray[T]) RemoveIndex(index uint32) *T {
	return (*T)(g_ptr_array_remove_index.Fn()(uptr(a), index))
}
func (a *GPtrArray[T]) RemoveIndexFast(index uint32) *T {
	return (*T)(g_ptr_array_remove_index_fast.Fn()(uptr(a), index))
}
func (a *GPtrArray[T]) StealIndex(index uint32) *T {
	return (*T)(g_ptr_array_steal_index.Fn()(uptr(a), index))
}
func (a *GPtrArray[T]) StealIndexFast(index uint32) *T {
	return (*T)(g_ptr_array_steal_index_fast.Fn()(uptr(a), index))
}
func (a *GPtrArray[T]) Remove(data *T) bool { return g_ptr_array_remove.Fn()(uptr(a), uptr(data)) != 0 }
func (a *GPtrArray[T]) RemoveFast(data *T) bool {
	return g_ptr_array_remove_fast.Fn()(uptr(a), uptr(data)) != 0
}
func (a *GPtrArray[T]) RemoveRange(index, length uint32) *GPtrArray[T] {
	return (*GPtrArray[T])(g_ptr_array_remove_range.Fn()(uptr(a), index, length))
}
func (a *GPtrArray[T]) Add(data *T) { g_ptr_array_add.Fn()(uptr(a), uptr(data)) }
func (a *GPtrArray[T]) Extend(array *GPtrArray[T], copyFunc func(src *T, _ uptr) *T) {
	cf := cc.Cbk(copyFunc)
	defer cc.CbkClose(cf)
	g_ptr_array_extend.Fn()(uptr(a), uptr(array), cf, nil)
}
func (a *GPtrArray[T]) ExtendAndSteal(array *GPtrArray[T]) {
	g_ptr_array_extend_and_steal.Fn()(uptr(a), uptr(array))
}
func (a *GPtrArray[T]) Insert(index int32, data *T) {
	g_ptr_array_insert.Fn()(uptr(a), index, uptr(data))
}
func (a *GPtrArray[T]) Sort(compareFunc func(a, b *T) int32) {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	g_ptr_array_sort.Fn()(uptr(a), cf)
}

// func (a *GPtrArray[T]) SortWithData(compareFunc func(a, b *T, userData UPtr) int32, userData UPtr) {
// 	cf := cc.Cbk(compareFunc)
// 	defer cc.CbkClose(cf)
// 	g_ptr_array_sort_with_data.Fn()(UPtr(a), cc.Cbk(compareFunc), nil)
// }

func (a *GPtrArray[T]) SortValues(compareFunc func(a, b *T) int32) {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	g_ptr_array_sort_values.Fn()(uptr(a), cf)
}

// func (a *GPtrArray[T]) SortValuesWithData(compareFunc func(a, b, userData UPtr) int32, userData UPtr) {
// 	g_ptr_array_sort_values_with_data.Fn()(a, cc.Cbk(compareFunc), userData)
// }

func (a *GPtrArray[T]) Foreach(fn func(data *T)) {
	f := cc.Cbk(fn)
	defer cc.CbkClose(f)
	g_ptr_array_foreach.Fn()(uptr(a), f, nil)
}
func (a *GPtrArray[T]) Find(needle *T) (index uint32, ok bool) {
	ok = g_ptr_array_find.Fn()(uptr(a), uptr(needle), &index) != 0
	return
}
func (a *GPtrArray[T]) FindWithEqualFunc(needle *T, equalFunc func(a, b *T) bool) (index uint32, ok bool) {
	ef := cc.Cbk(equalFunc)
	defer cc.CbkClose(ef)
	ok = g_ptr_array_find_with_equal_func.Fn()(uptr(a), uptr(needle), ef, &index) != 0
	return
}
func (a *GPtrArray[T]) IsNullTerminated() bool {
	return g_ptr_array_is_null_terminated.Fn()(uptr(a)) != 0
}

func NewGByteArray() *GByteArray { return g_byte_array_new.Fn()() }
func NewGByteArrayFrom(data []byte) *GByteArray {
	l := len(data)
	p := GAlloc[byte](uint64(l))
	copy(unsafe.Slice(p, l), data)
	return g_byte_array_new_take.Fn()(uptr(p), uint64(l))
}
func NewGByteArrayTake(data uptr, length uint64) *GByteArray {
	return g_byte_array_new_take.Fn()(data, length)
}
func (a *GByteArray) List() []byte {
	return unsafe.Slice((*byte)(a.data), a.Len)
}
func (a *GByteArray) Steal() (data uptr, length uint64) {
	data = g_byte_array_steal.Fn()(a, &length)
	return
}
func NewGByteArraySized(reservedSize uint32) *GByteArray {
	return g_byte_array_sized_new.Fn()(reservedSize)
}
func (a *GByteArray) Free(freeSegment bool) *byte {
	return g_byte_array_free.Fn()(a, cbool(freeSegment))
}
func (a *GByteArray) FreeToBytes() *GBytes { return g_byte_array_free_to_bytes.Fn()(a) }
func (a *GByteArray) Ref() *GByteArray     { return g_byte_array_ref.Fn()(a) }
func (a *GByteArray) Unref()               { g_byte_array_unref.Fn()(a) }
func (a *GByteArray) Append(data []byte) *GByteArray {
	return g_byte_array_append.Fn()(a, uptr(&data[0]), uint32(len(data)))
}
func (a *GByteArray) Prepend(data []byte) *GByteArray {
	return g_byte_array_prepend.Fn()(a, uptr(&data[0]), uint32(len(data)))
}
func (a *GByteArray) SetSize(length uint32) *GByteArray { return g_byte_array_set_size.Fn()(a, length) }
func (a *GByteArray) RemoveIndex(index uint32) *GByteArray {
	return g_byte_array_remove_index.Fn()(a, index)
}
func (a *GByteArray) RemoveIndexFast(index uint32) *GByteArray {
	return g_byte_array_remove_index_fast.Fn()(a, index)
}
func (a *GByteArray) RemoveRange(index, length uint32) *GByteArray {
	return g_byte_array_remove_range.Fn()(a, index, length)
}
func (a *GByteArray) Sort(compareFunc func(a, b *byte) int32) {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	g_byte_array_sort.Fn()(a, cf)
}

// func (a *GByteArray) SortWithData(compareFunc func(a, b, userData UPtr) int32, userData UPtr) {
// 	g_byte_array_sort_with_data.Fn()(a, cc.Cbk(compareFunc), userData)
// }

// #endregion

// #region AsyncQueue

type GAsyncQueue struct{}

func NewGAsyncQueue() *GAsyncQueue {
	return g_async_queue_new.Fn()()
}
func NewGAsyncQueueFull(itemFreeFunc cc.Func /* itemFreeFunc func(data UPtr) */) *GAsyncQueue {
	return g_async_queue_new_full.Fn()(itemFreeFunc)
}
func (q *GAsyncQueue) Lock()             { g_async_queue_lock.Fn()(q) }
func (q *GAsyncQueue) Unlock()           { g_async_queue_unlock.Fn()(q) }
func (q *GAsyncQueue) Ref() *GAsyncQueue { return g_async_queue_ref.Fn()(q) }
func (q *GAsyncQueue) Unref()            { g_async_queue_unref.Fn()(q) }
func (q *GAsyncQueue) Push(data interface{}) {
	g_async_queue_push.Fn()(q, anyptr(data))
}
func (q *GAsyncQueue) PushUnlocked(data interface{}) {
	g_async_queue_push_unlocked.Fn()(q, anyptr(data))
}
func (q *GAsyncQueue) PushSorted(data interface{}, cmp func(a, b, _ uptr) int32) {
	cp := cc.Cbk(cmp)
	defer cc.CbkClose(cp)
	g_async_queue_push_sorted.Fn()(q, anyptr(data), cp, nil)
}
func (q *GAsyncQueue) PushSortedUnlocked(data interface{}, cmp func(a, b, _ uptr) int32) {
	cp := cc.Cbk(cmp)
	defer cc.CbkClose(cp)
	g_async_queue_push_sorted_unlocked.Fn()(q, anyptr(data), cp, nil)
}
func (q *GAsyncQueue) Pop() uptr            { return g_async_queue_pop.Fn()(q) }
func (q *GAsyncQueue) PopUnlocked() uptr    { return g_async_queue_pop_unlocked.Fn()(q) }
func (q *GAsyncQueue) TryPop() uptr         { return g_async_queue_try_pop.Fn()(q) }
func (q *GAsyncQueue) TryPopUnlocked() uptr { return g_async_queue_try_pop_unlocked.Fn()(q) }
func (q *GAsyncQueue) TimeoutPop(timeout uint64) uptr {
	return g_async_queue_timeout_pop.Fn()(q, timeout)
}
func (q *GAsyncQueue) TimeoutPopUnlocked(timeout uint64) uptr {
	return g_async_queue_timeout_pop_unlocked.Fn()(q, timeout)
}
func (q *GAsyncQueue) Length() int32         { return g_async_queue_length.Fn()(q) }
func (q *GAsyncQueue) LengthUnlocked() int32 { return g_async_queue_length_unlocked.Fn()(q) }
func (q *GAsyncQueue) Sort(cmp func(a, b, _ uptr) int32) {
	cp := cc.Cbk(cmp)
	defer cc.CbkClose(cp)
	g_async_queue_sort.Fn()(q, cp, nil)
}
func (q *GAsyncQueue) SortUnlocked(cmp func(a, b, _ uptr) int32) {
	cp := cc.Cbk(cmp)
	defer cc.CbkClose(cp)
	g_async_queue_sort_unlocked.Fn()(q, cp, nil)
}
func (q *GAsyncQueue) Remove(item interface{}) bool {
	return g_async_queue_remove.Fn()(q, anyptr(item)) != 0
}
func (q *GAsyncQueue) RemoveUnlocked(item interface{}) bool {
	return g_async_queue_remove_unlocked.Fn()(q, anyptr(item)) != 0
}
func (q *GAsyncQueue) PushFront(item interface{}) { g_async_queue_push_front.Fn()(q, anyptr(item)) }
func (q *GAsyncQueue) PushFrontUnlocked(item interface{}) {
	g_async_queue_push_front_unlocked.Fn()(q, anyptr(item))
}

// #endregion

// #region Bytes

type GBytes struct{}

func NewGBytes[T string | []byte](data T) *GBytes {
	dt := (*struct {
		p *byte
		l int
	})((uptr)(&data))

	return g_bytes_new.Fn()(dt.p, uint64(dt.l))
}

func (b *GBytes) Compare(b2 *GBytes) int { return g_bytes_compare.Fn()(b, b2) }
func (b *GBytes) Equal(b2 *GBytes) bool  { return g_bytes_equal.Fn()(b, b2) != 0 }
func (b *GBytes) GetData() []byte {
	var size uint64
	ptr := g_bytes_get_data.Fn()(b, &size)
	return unsafe.Slice(ptr, size)
}
func (b *GBytes) GetRegion(offset, length uint64) []byte {
	ptr := g_bytes_get_region.Fn()(b, offset, length)
	return unsafe.Slice(ptr, length)
}
func (b *GBytes) GetSize() uint64 { return g_bytes_get_size.Fn()(b) }
func (b *GBytes) Hash() uint32    { return g_bytes_hash.Fn()(b) }
func (b *GBytes) NewFromBytes(offset, length uint64) *GBytes {
	return g_bytes_new_from_bytes.Fn()(b, offset, length)
}
func (b *GBytes) Ref() *GBytes              { return g_bytes_ref.Fn()(b) }
func (b *GBytes) Unref()                    { g_bytes_unref.Fn()(b) }
func (b *GBytes) UnrefToArray() *GByteArray { return g_bytes_unref_to_array.Fn()(b) }
func (b *GBytes) UnrefToData() []byte {
	var size uint64
	ptr := g_bytes_unref_to_data.Fn()(b, &size)
	defer GFree(ptr)
	return unsafe.Slice(ptr, size)
}
func (b *GBytes) UnrefToString() string {
	var size uint64
	ptr := g_bytes_unref_to_data.Fn()(b, &size)
	return (cc.String)(uptr(ptr)).TakeString(size)
}

// #endregion

// #region DateTime

type GDateTime struct{}

func (dt *GDateTime) Unref()          { g_date_time_unref.Fn()(dt) }
func (dt *GDateTime) Ref() *GDateTime { return g_date_time_ref.Fn()(dt) }

func NewGDateTimeNow(tz *GTimeZone) *GDateTime { return g_date_time_new_now.Fn()(tz) }
func NewGDateTimeNowLocal() *GDateTime         { return g_date_time_new_now_local.Fn()() }
func NewGDateTimeNowUtc() *GDateTime           { return g_date_time_new_now_utc.Fn()() }

func NewGDateTimeFromUnixLocal(t int64) *GDateTime { return g_date_time_new_from_unix_local.Fn()(t) }
func NewGDateTimeFromUnixUtc(t int64) *GDateTime   { return g_date_time_new_from_unix_utc.Fn()(t) }
func NewGDateTimeFromUnixLocalUsec(usecs int64) *GDateTime {
	return g_date_time_new_from_unix_local_usec.Fn()(usecs)
}
func NewGDateTimeFromUnixUtcUsec(usecs int64) *GDateTime {
	return g_date_time_new_from_unix_utc_usec.Fn()(usecs)
}

func NewGDateTimeFromIso8601(text string, defaultTz *GTimeZone) *GDateTime {
	t := cc.NewString(text)
	defer t.Free()
	return g_date_time_new_from_iso8601.Fn()(t, defaultTz)
}
func NewGDateTime(tz *GTimeZone, year, month, day, hour, minute int, seconds float64) *GDateTime {
	return g_date_time_new.Fn()(tz, year, month, day, hour, minute, seconds)
}
func NewGDateTimeLocal(year, month, day, hour, minute int, seconds float64) *GDateTime {
	return g_date_time_new_local.Fn()(year, month, day, hour, minute, seconds)
}
func NewGDateTimeUtc(year, month, day, hour, minute int, seconds float64) *GDateTime {
	return g_date_time_new_utc.Fn()(year, month, day, hour, minute, seconds)
}

func (dt *GDateTime) Add(timespan GTimeSpan) *GDateTime { return g_date_time_add.Fn()(dt, timespan) }
func (dt *GDateTime) AddYears(years int) *GDateTime     { return g_date_time_add_years.Fn()(dt, years) }
func (dt *GDateTime) AddMonths(months int) *GDateTime   { return g_date_time_add_months.Fn()(dt, months) }
func (dt *GDateTime) AddWeeks(weeks int) *GDateTime     { return g_date_time_add_weeks.Fn()(dt, weeks) }
func (dt *GDateTime) AddDays(days int) *GDateTime       { return g_date_time_add_days.Fn()(dt, days) }
func (dt *GDateTime) AddHours(hours int) *GDateTime     { return g_date_time_add_hours.Fn()(dt, hours) }
func (dt *GDateTime) AddMinutes(minutes int) *GDateTime {
	return g_date_time_add_minutes.Fn()(dt, minutes)
}
func (dt *GDateTime) AddSeconds(seconds float64) *GDateTime {
	return g_date_time_add_seconds.Fn()(dt, seconds)
}
func (dt *GDateTime) AddFull(years, months, days, hours, minutes int, seconds float64) *GDateTime {
	return g_date_time_add_full.Fn()(dt, years, months, days, hours, minutes, seconds)
}

func (dt1 *GDateTime) Compare(dt2 *GDateTime) int {
	return g_date_time_compare.Fn()(dt1, dt2)
}
func (begin *GDateTime) Difference(end *GDateTime) GTimeSpan {
	return g_date_time_difference.Fn()(end, begin)
}
func (datetime *GDateTime) TimeHash() uint32 {
	return g_date_time_hash.Fn()(datetime)
}
func (dt *GDateTime) Equal(dt2 *GDateTime) bool {
	return g_date_time_equal.Fn()(dt, dt2) != 0
}
func (dt *GDateTime) GetYmd() (year, month, day int) {
	g_date_time_get_ymd.Fn()(dt, &year, &month, &day)
	return
}
func (dt *GDateTime) GetYear() int              { return g_date_time_get_year.Fn()(dt) }
func (dt *GDateTime) GetMonth() int             { return g_date_time_get_month.Fn()(dt) }
func (dt *GDateTime) GetDayOfMonth() int        { return g_date_time_get_day_of_month.Fn()(dt) }
func (dt *GDateTime) GetWeekNumberingYear() int { return g_date_time_get_week_numbering_year.Fn()(dt) }
func (dt *GDateTime) GetWeekOfYear() int        { return g_date_time_get_week_of_year.Fn()(dt) }
func (dt *GDateTime) GetDayOfWeek() int         { return g_date_time_get_day_of_week.Fn()(dt) }
func (dt *GDateTime) GetDayOfYear() int         { return g_date_time_get_day_of_year.Fn()(dt) }
func (dt *GDateTime) GetHour() int              { return g_date_time_get_hour.Fn()(dt) }
func (dt *GDateTime) GetMinute() int            { return g_date_time_get_minute.Fn()(dt) }
func (dt *GDateTime) GetSecond() int            { return g_date_time_get_second.Fn()(dt) }
func (dt *GDateTime) GetMicrosecond() int       { return g_date_time_get_microsecond.Fn()(dt) }
func (dt *GDateTime) GetSeconds() float64       { return g_date_time_get_seconds.Fn()(dt) }

func (dt *GDateTime) ToUnix() int64     { return g_date_time_to_unix.Fn()(dt) }
func (dt *GDateTime) ToUnixUsec() int64 { return g_date_time_to_unix_usec.Fn()(dt) }

func (dt *GDateTime) GetUtcOffset() GTimeSpan { return g_date_time_get_utc_offset.Fn()(dt) }
func (dt *GDateTime) GetTimezone() *GTimeZone { return g_date_time_get_timezone.Fn()(dt) }
func (dt *GDateTime) GetTimezoneAbbreviation() string {
	return g_date_time_get_timezone_abbreviation.Fn()(dt).String()
}
func (dt *GDateTime) IsDaylightSavings() bool { return g_date_time_is_daylight_savings.Fn()(dt) != 0 }

func (dt *GDateTime) ToTimezone(tz *GTimeZone) *GDateTime {
	return g_date_time_to_timezone.Fn()(dt, tz)
}
func (dt *GDateTime) ToLocal() *GDateTime { return g_date_time_to_local.Fn()(dt) }
func (dt *GDateTime) ToUtc() *GDateTime   { return g_date_time_to_utc.Fn()(dt) }

func (dt *GDateTime) Format(format string) string {
	f := cc.NewString(format)
	defer f.Free()
	return g_date_time_format.Fn()(dt, f).TakeString()
}
func (dt *GDateTime) FormatIso8601() string { return g_date_time_format_iso8601.Fn()(dt).TakeString() }

// #endregion

// #region Env Exports
func GGetenv(variable string) string {
	v := cc.NewString(variable)
	defer v.Free()
	return g_getenv.Fn()(v).String()
}
func GSetenv(variable, value string, overwrite bool) bool {
	vb, vl := cc.NewString(variable), cc.NewString(value)
	defer vb.Free()
	defer vl.Free()
	return g_setenv.Fn()(vb, vl, cbool(overwrite)) != 0
}
func GUnsetenv(variable string) {
	v := cc.NewString(variable)
	defer v.Free()
	g_unsetenv.Fn()(v)
}
func GListenv() []string    { return g_listenv.Fn()().TakeStrings() }
func GGetEnviron() []string { return g_get_environ.Fn()().TakeStrings() }
func GEnvironGetenv(envp []string, variable string) string {
	p, v := cc.NewStrings(envp), cc.NewString(variable)
	defer p.Free()
	return g_environ_getenv.Fn()(p, v).String()
}
func GEnvironSetenv(envp []string, variable, value string, overwrite bool) []string {
	p, vb, va := cc.NewStrings(envp), cc.NewString(variable), cc.NewString(value)
	defer p.Free()
	return g_environ_setenv.Fn()(p, vb, va, cbool(overwrite)).Strings()
}
func GEnvironUnsetenv(envp []string, variable string) []string {
	p, v := cc.NewStrings(envp), cc.NewString(variable)
	defer p.Free()
	return g_environ_unsetenv.Fn()(p, v).TakeStrings()
}

// #endregion

// #region Error

type GError struct {
	Domain  GQuark
	Code    int32
	Message cc.String
}

func NewGError(domain GQuark, code int32, format string, a ...interface{}) *GError {
	f := cc.NewString(fmt.Sprintf(format, a...))
	defer f.Free()
	return g_error_new.Fn()(domain, code, f)
}
func ErrorDomainRegister(typeName string, privateSize uint,
	typeInit func(err *GError), typeCopy func(src, dst *GError), typeFree func(err *GError)) GQuark {
	init, copy, free := cc.Cbk(typeInit), cc.Cbk(typeCopy), cc.Cbk(typeFree)
	tn := cc.NewString(typeName)
	defer tn.Free()
	return g_error_domain_register.Fn()(tn, privateSize, init, copy, free)
}
func ErrorDomainRegisterStatic(typeName string, privateSize uint,
	typeInit func(err *GError), typeCopy func(src, dst *GError), typeFree func(err *GError)) GQuark {
	init, copy, free := cc.Cbk(typeInit), cc.Cbk(typeCopy), cc.Cbk(typeFree)
	tn := cc.NewString(typeName)
	return g_error_domain_register_static.Fn()(tn, privateSize, init, copy, free)
}

func (e *GError) Copy() *GError                     { return g_error_copy.Fn()(e) }
func (e *GError) Free()                             { g_error_free.Fn()(e) }
func (e *GError) Matches(domain GQuark, code int32) { g_error_matches.Fn()(e, domain, code) }
func (e *GError) Error() error {
	if e == nil {
		return nil
	}
	return errors.New(e.Message.String())
}
func (e *GError) TakeError() error {
	if e == nil {
		return nil
	}
	defer e.Free()
	return errors.New(e.Message.String())
}

// #endregion

// #region List

type GList[T any] struct {
	Data *T
	Next *GList[T]
	Prev *GList[T]
}

func GListAlloc[T any]() *GList[T] { return (*GList[T])(g_list_alloc.Fn()()) }
func NewGList[T any]() *GList[T]   { return (*GList[T])(g_list_alloc.Fn()()) }

// func AllocGList[T any]() *GList[T] { return (*GList[T])(g_list_alloc.Fn()()) }

func (l *GList[T]) Free()  { g_list_free.Fn()(uptr(l)) }
func (l *GList[T]) Free1() { g_list_free_1.Fn()(uptr(l)) }
func (l *GList[T]) FreeFull(freeFunc func(data *T)) {
	ff := cc.Cbk(freeFunc)
	defer cc.CbkClose(ff)
	g_list_free_full.Fn()(uptr(l), ff)
}
func (l *GList[T]) Append(data *T) *GList[T] {
	return (*GList[T])(g_list_append.Fn()(uptr(l), uptr(data)))
}
func (l *GList[T]) Prepend(data *T) *GList[T] {
	return (*GList[T])(g_list_prepend.Fn()(uptr(l), uptr(data)))
}
func (l *GList[T]) Insert(data *T, position int32) *GList[T] {
	return (*GList[T])(g_list_insert.Fn()(uptr(l), uptr(data), position))
}
func (l *GList[T]) InsertSorted(data *T, compare func(a, b *T) int32) *GList[T] {
	cp := cc.Cbk(compare)
	defer cc.CbkClose(cp)
	return (*GList[T])(g_list_insert_sorted.Fn()(uptr(l), uptr(data), cp))
}

//	func (l *GList[T]) InsertSortedWithData(data *T, compare func(a, b, data UPtr) int32, userData UPtr) *GList {
//		return g_list_insert_sorted_with_data.Fn()(l, data, cc.Cbk(compare), userData)
//	}

func (l *GList[T]) InsertBefore(sibling *GList[T], data *T) *GList[T] {
	return (*GList[T])(g_list_insert_before.Fn()(uptr(l), uptr(sibling), uptr(data)))
}
func (l *GList[T]) InsertBeforeLink(sibling *GList[T], link *GList[T]) *GList[T] {
	return (*GList[T])(g_list_insert_before_link.Fn()(uptr(l), uptr(sibling), uptr(link)))
}
func (l *GList[T]) Concat(list2 *GList[T]) *GList[T] {
	return (*GList[T])(g_list_concat.Fn()(uptr(l), uptr(list2)))
}
func (l *GList[T]) Remove(data *T) *GList[T] {
	return (*GList[T])(g_list_remove.Fn()(uptr(l), uptr(data)))
}
func (l *GList[T]) RemoveAll(data *T) *GList[T] {
	return (*GList[T])(g_list_remove_all.Fn()(uptr(l), uptr(data)))
}
func (l *GList[T]) RemoveLink(link *GList[T]) *GList[T] {
	return (*GList[T])(g_list_remove_link.Fn()(uptr(l), uptr(link)))
}
func (l *GList[T]) DeleteLink(link *GList[T]) *GList[T] {
	return (*GList[T])(g_list_delete_link.Fn()(uptr(l), uptr(link)))
}
func (l *GList[T]) Reverse() *GList[T] {
	return (*GList[T])(g_list_reverse.Fn()(uptr(l)))
}
func (l *GList[T]) Copy() *GList[T] {
	return (*GList[T])(g_list_copy.Fn()(uptr(l)))
}
func (l *GList[T]) CopyDeep(copyFunc func(src *T, _ uptr) *T) *GList[T] {
	cp := cc.Cbk(copyFunc)
	defer cc.CbkClose(cp)
	return (*GList[T])(g_list_copy_deep.Fn()(uptr(l), cp, nil))
}
func (l *GList[T]) Nth(n uint32) *GList[T]     { return (*GList[T])(g_list_nth.Fn()(uptr(l), n)) }
func (l *GList[T]) NthPrev(n uint32) *GList[T] { return (*GList[T])(g_list_nth_prev.Fn()(uptr(l), n)) }
func (l *GList[T]) Find(data *T) *GList[T]     { return (*GList[T])(g_list_find.Fn()(uptr(l), uptr(data))) }
func (l *GList[T]) FindCustom(data *T, compareFunc func(curr, data *T) int32) *GList[T] {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	return (*GList[T])(g_list_find_custom.Fn()(uptr(l), uptr(data), cf))
}
func (l *GList[T]) Position(link *GList[T]) int32 {
	return g_list_position.Fn()(uptr(l), uptr(link))
}
func (l *GList[T]) Index(data *T) int32 {
	return g_list_index.Fn()(uptr(l), uptr(data))
}
func (l *GList[T]) Last() *GList[T] {
	return (*GList[T])(g_list_last.Fn()(uptr(l)))
}
func (l *GList[T]) First() *GList[T] {
	return (*GList[T])(g_list_first.Fn()(uptr(l)))
}
func (l *GList[T]) Length() uint32 {
	return g_list_length.Fn()(uptr(l))
}
func (l *GList[T]) ForEach(fn func(data *T, _ uptr)) {
	f := cc.Cbk(fn)
	defer cc.CbkClose(f)
	g_list_foreach.Fn()(uptr(l), f, nil)
}

func (l *GList[T]) Sort(compareFunc func(a, b *T) int32) *GList[T] {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	return (*GList[T])(g_list_sort.Fn()(uptr(l), cf))
}

// func (l *GList[T]) SortWithData(compareFunc func(a, b, userData UPtr) int32, userData UPtr) *GList[T] {
// 	return (*GList[T])(g_list_sort_with_data.Fn()(l, cc.Cbk(compareFunc), userData))
// }

func (l *GList[T]) NthData(n uint32) *T { return (*T)(g_list_nth_data.Fn()(uptr(l), n)) }
func ClearList[T any](list **GList[T], destroy func(data *T)) {
	d := cc.Cbk(destroy)
	defer cc.CbkClose(d)
	g_clear_list.Fn()(uptr(list), d)
}

// func GListList[T any](l *GList[T], close func(*GList[T])) []T {
// 	type pinner struct {
// 		list  *GList[T]
// 		slice *[]T
// 	}
// 	num := l.Length()
// 	if num == 0 {
// 		return nil
// 	}

// 	pin := new(pinner)
// 	pin.list = l
// 	slice := make([]T, num)
// 	pin.slice = &slice
// 	runtime.SetFinalizer(pin, func(p *pinner) {
// 		if close != nil {
// 			close(p.list)
// 			return
// 		}
// 		p.list.Free()

// 		_ = *p.slice
// 	})

// 	var t T
// 	switch any(t).(type) {
// 	case string:
// 		for i := uint32(0); i < num; i++ {
// 			a := ((cc.String)(UPtr(l.Nth(i).Data))).String()
// 			(*pin.slice)[i] = any(a).(T)
// 		}
// 	default:
// 		for i := uint32(0); i < num; i++ {
// 			(*pin.slice)[i] = *(*T)(UPtr(&l.Nth(i).Data))
// 		}
// 	}

//		return *pin.slice
//	}

func (l *GList[T]) List() (lst []*T) {
	lst = []*T{}
	l.ForEach(func(data *T, _ uptr) {
		lst = append(lst, data)
	})
	return
}

// #endregion

// #region Main

type GMainContext struct{}
type GMainLoop struct{}

type GSource struct {
	CallBackData  uptr
	CallbackFuncs *GSourceCallbackFuncs
	SourceFuncs   *GSourceFuncs
	RefCount      uint32
	Context       *GMainContext

	Priority int32
	Flags    uint32
	SourceID uint32

	PollFds *GSList[GPollFD]
	Prev    *GSource
	Next    *GSource

	Name cc.String
	Priv *GSourcePrivate
}
type GSourcePrivate struct{}

type GSourceCallbackFuncs struct {
	Ref   cc.Func // void (*ref) (gpointer cb_data);
	Unref cc.Func // void (*unref) (gpointer cb_data);
	Get   cc.Func // void (*get) (gpointer cb_data, GSource *source, GSourceFunc *func, gpointer *data);
}
type GSourceFuncs struct {
	Prepare  cc.Func // /* Can be NULL */ typedef gboolean (*GSourceFuncsPrepareFunc) (GSource *source, gint *timeout_);
	Check    cc.Func // /* Can be NULL */ typedef gboolean (*GSourceFuncsCheckFunc) (GSource *source);
	Dispatch cc.Func // typedef gboolean (*GSourceFuncsDispatchFunc) (GSource *source,GSourceFunc callback,gpointer user_data);
	Finalize cc.Func // /* Can be NULL */ typedef void (*GSourceFuncsFinalizeFunc) (GSource *source);

	ClosureCallback cc.Func // typedef gboolean (*GSourceFunc)(gpointer user_data);
	ClosureMarshal  cc.Func // typedef void (*GSourceDummyMarshal) (void);
}

type GPid int32

func NewGMainContext() *GMainContext { return g_main_context_new.Fn()() }
func NewGMainContextWithFlags(flags GMainContextFlags) *GMainContext {
	return g_main_context_new_with_flags.Fn()(flags)
}
func (c *GMainContext) Ref() *GMainContext { return g_main_context_ref.Fn()(c) }
func (c *GMainContext) Unref()             { g_main_context_unref.Fn()(c) }
func DefaultMainContext() *GMainContext    { return g_main_context_default.Fn()() }
func (c *GMainContext) Iteration(mayBlock bool) bool {
	return g_main_context_iteration.Fn()(c, cbool(mayBlock)) != 0
}
func (c *GMainContext) Pending() bool { return g_main_context_pending.Fn()(c) != 0 }
func (c *GMainContext) FindSourceByID(sourceID uint32) *GSource {
	return g_main_context_find_source_by_id.Fn()(c, sourceID)
}
func (c *GMainContext) FindSourceByUserData(userData interface{}) *GSource {
	return g_main_context_find_source_by_user_data.Fn()(c, anyptr(userData))
}
func (c *GMainContext) FindSourceByFuncsUserData(funcs *GSourceFuncs, userData interface{}) *GSource {
	return g_main_context_find_source_by_funcs_user_data.Fn()(c, funcs, anyptr(userData))
}
func (c *GMainContext) Wakeup()       { g_main_context_wakeup.Fn()(c) }
func (c *GMainContext) Acquire() bool { return g_main_context_acquire.Fn()(c) != 0 }
func (c *GMainContext) Release()      { g_main_context_release.Fn()(c) }
func (c *GMainContext) IsOwner() bool { return g_main_context_is_owner.Fn()(c) != 0 }

// func (c *GMainContext) Wait(cond, mutex UPtr) bool   { return g_main_context_wait.Fn()(c, cond, mutex) }

func (c *GMainContext) Prepare() (priority int32, ready bool) {
	ready = g_main_context_prepare.Fn()(c, &priority) != 0
	return
}
func (c *GMainContext) Query(maxPriority int32, timeout *int32, fds *GPollFD, nFds int32) int32 {
	return g_main_context_query.Fn()(c, maxPriority, timeout, fds, nFds)
}
func (c *GMainContext) Check(maxPriority int32, fds *GPollFD, nFds int32) bool {
	return g_main_context_check.Fn()(c, maxPriority, fds, nFds) != 0
}
func (c *GMainContext) Dispatch() { g_main_context_dispatch.Fn()(c) }
func (c *GMainContext) SetPollFunc(fn func(ufds *GPollFD, nfsd uint32, timeout int32) int32) {
	g_main_context_set_poll_func.Fn()(c, cc.Cbk(fn))
}
func (c *GMainContext) GetPollFunc() uptr { return g_main_context_get_poll_func.Fn()(c) }
func (c *GMainContext) AddPoll(fd *GPollFD, priority int32) {
	g_main_context_add_poll.Fn()(c, fd, priority)
}
func (c *GMainContext) RemovePoll(fd *GPollFD)   { g_main_context_remove_poll.Fn()(c, fd) }
func MainDepth() int32                           { return g_main_depth.Fn()() }
func MainCurrentSource() *GSource                { return g_main_current_source.Fn()() }
func (c *GMainContext) PushThreadDefault()       { g_main_context_push_thread_default.Fn()(c) }
func (c *GMainContext) PopThreadDefault()        { g_main_context_pop_thread_default.Fn()(c) }
func GetThreadDefaultMainContext() *GMainContext { return g_main_context_get_thread_default.Fn()() }
func RefThreadDefaultMainContext() *GMainContext { return g_main_context_ref_thread_default.Fn()() }
func (c *GMainContext) InvokeFull(priority int32, fn func(_ uptr) (Continue bool), notify func(_ uptr)) {
	var f, n uptr
	f = cc.Cbk(f)
	n = cc.Cbk(func(a uptr) {
		if notify != nil {
			notify(a)
		}
		if f != nil {
			cc.CbkClose(f)
		}
		cc.CbkCloseLate(n)
	})

	g_main_context_invoke_full.Fn()(c, priority, f, nil, n)
}
func (c *GMainContext) Invoke(fn func(_ uptr) (Continue bool)) {
	var f uptr
	f = cc.Cbk(func(a uptr) (Continue bool) {
		if fn(a) {
			return true
		}

		cc.CbkCloseLate(f)
		return false
	})
	g_main_context_invoke.Fn()(c, f, nil)
}

func NewGMainLoop(context *GMainContext, isRunning bool) *GMainLoop {
	return g_main_loop_new.Fn()(context, cbool(isRunning))
}
func (l *GMainLoop) Run()                      { g_main_loop_run.Fn()(l) }
func (l *GMainLoop) Quit()                     { g_main_loop_quit.Fn()(l) }
func (l *GMainLoop) Ref() *GMainLoop           { return g_main_loop_ref.Fn()(l) }
func (l *GMainLoop) Unref()                    { g_main_loop_unref.Fn()(l) }
func (l *GMainLoop) IsRunning() bool           { return g_main_loop_is_running.Fn()(l) != 0 }
func (l *GMainLoop) GetContext() *GMainContext { return g_main_loop_get_context.Fn()(l) }

func NewGSource(sourceFuncs *GSourceFuncs, structSize uint32) *GSource {
	return g_source_new.Fn()(sourceFuncs, structSize)
}
func (s *GSource) SetDisposeFunction(dispose func(s *GSource)) {
	var d uptr
	d = cc.Cbk(func(s *GSource) {
		dispose(s)
		cc.CbkCloseLate(d)
	})
	g_source_set_dispose_function.Fn()(s, d)
}
func (s *GSource) Ref() *GSource                       { return g_source_ref.Fn()(s) }
func (s *GSource) Unref()                              { g_source_unref.Fn()(s) }
func (s *GSource) Attach(context *GMainContext) uint32 { return g_source_attach.Fn()(s, context) }
func (s *GSource) Destroy()                            { g_source_destroy.Fn()(s) }
func (s *GSource) SetPriority(priority int32)          { g_source_set_priority.Fn()(s, priority) }
func (s *GSource) GetPriority() int32                  { return g_source_get_priority.Fn()(s) }
func (s *GSource) SetCanRecurse(canRecurse bool)       { g_source_set_can_recurse.Fn()(s, cbool(canRecurse)) }
func (s *GSource) GetCanRecurse() bool                 { return g_source_get_can_recurse.Fn()(s) != 0 }
func (s *GSource) GetID() uint32                       { return g_source_get_id.Fn()(s) }
func (s *GSource) GetContext() *GMainContext           { return g_source_get_context.Fn()(s) }
func (s *GSource) SetCallback(fn func(_ uptr) (Continue bool)) {
	g_source_set_callback.Fn()(s, cc.Cbk(fn), nil, nil)
}
func (s *GSource) SetFuncs(funcs *GSourceFuncs) { g_source_set_funcs.Fn()(s, funcs) }
func (s *GSource) IsDestroyed() bool            { return g_source_is_destroyed.Fn()(s) != 0 }
func (s *GSource) SetName(name string) {
	cs := cc.NewString(name)
	defer cs.Free()
	g_source_set_name.Fn()(s, cs)
}
func (s *GSource) SetStaticName(name string) {
	cs := cc.NewString(name)
	defer cs.Free()
	g_source_set_static_name.Fn()(s, cs)
}
func (s *GSource) GetName() string { return g_source_get_name.Fn()(s).String() }
func SourceSetNameByID(tag uint32, name string) {
	cs := cc.NewString(name)
	defer cs.Free()
	g_source_set_name_by_id.Fn()(tag, cs)
}
func (s *GSource) SetReadyTime(readyTime int64)     { g_source_set_ready_time.Fn()(s, readyTime) }
func (s *GSource) GetReadyTime() int64              { return g_source_get_ready_time.Fn()(s) }
func (s *GSource) AddPoll(fd *GPollFD)              { g_source_add_poll.Fn()(s, fd) }
func (s *GSource) RemovePoll(fd *GPollFD)           { g_source_remove_poll.Fn()(s, fd) }
func (s *GSource) AddChildSource(child *GSource)    { g_source_add_child_source.Fn()(s, child) }
func (s *GSource) RemoveChildSource(child *GSource) { g_source_remove_child_source.Fn()(s, child) }

// func (s *GSource) GetCurrentTime(timeval *GTimeVal) { g_source_get_current_time(s, timeval) }

func (s *GSource) GetTime() int64 { return g_source_get_time.Fn()(s) }

func NewIdleSource() *GSource                   { return g_idle_source_new.Fn()() }
func NewChildWatchSource(pid GPid) *GSource     { return g_child_watch_source_new.Fn()(pid) }
func NewTimeoutSource(interval uint32) *GSource { return g_timeout_source_new.Fn()(interval) }
func NewTimeoutSourceSeconds(interval uint32) *GSource {
	return g_timeout_source_new_seconds.Fn()(interval)
}

func SourceRemove(tag uint32) bool { return g_source_remove.Fn()(tag) != 0 }
func SourceRemoveByUserData(userData interface{}) bool {
	return g_source_remove_by_user_data.Fn()(anyptr(userData)) != 0
}
func SourceRemoveByFuncsUserData(funcs *GSourceFuncs, userData interface{}) bool {
	return g_source_remove_by_funcs_user_data.Fn()(funcs, anyptr(userData)) != 0
}

func TimeoutAddFull(priority int32, ms uint32, fn func(_ uptr) (Continue bool), data interface{}, notify func(_ uptr)) uint32 {
	return g_timeout_add_full.Fn()(priority, ms, cc.Cbk(fn), anyptr(data), cc.Cbk(notify))
}
func TimeoutAdd(ms uint32, fn func() (Continue bool)) uint32 {
	var fnp uptr
	fnp = cc.Cbk(func(_ uptr) (Continue bool) {
		b := fn()
		if !b {
			cc.CbkCloseLate(fnp)
		}
		return b
	})
	return g_timeout_add.Fn()(ms, fnp, nil)
}
func TimeoutAddOnce(ms uint32, fn func()) uint32 {
	return g_timeout_add_once.Fn()(ms, cc.Cbk(fn), nil)
}
func TimeoutAddSecondsFull(priority int32, sec uint32, fn func(_ uptr) (Continue bool),
	data interface{}, notify func(_ uptr)) uint32 {
	return g_timeout_add_seconds_full.Fn()(priority, sec, cc.Cbk(fn), anyptr(data), cc.Cbk(notify))
}
func TimeoutAddSeconds(sec uint32, fn func(_ uptr) (Continue bool)) uint32 {
	return g_timeout_add_seconds.Fn()(sec, cc.Cbk(fn), nil)
}
func TimeoutAddSecondsOnce(sec uint32, fn func()) uint32 {
	return g_timeout_add_seconds_once.Fn()(sec, cc.Cbk(fn), nil)
}
func ChildWatchAddFull(priority int32, pid GPid, fn func(pid GPid, waitStatus int32),
	data interface{}, notify func(_ uptr)) uint32 {
	return g_child_watch_add_full.Fn()(priority, pid, cc.Cbk(fn), anyptr(data), cc.Cbk(notify))
}
func ChildWatchAdd(pid GPid, fn func(pid GPid, waitStatus int32)) uint32 {
	return g_child_watch_add.Fn()(pid, cc.Cbk(fn), nil)
}
func IdleAdd[T any](fn func(data *T) (Continue bool), data *T) uint32 {
	var f uptr
	f = cc.CbkRaw[func(data uptr) int32](func(out, ins uptr) {
		data := cc.RawInAddr[T](ins, 0)
		ret := (*int32)(out)
		if fn(data) {
			(*ret) = 1
		} else {
			(*ret) = 0
			cc.CbkCloseLate(f)
		}
	})
	return g_idle_add.Fn()(f, uptr(data))
}
func IdleAddFull[T any](
	priority int32, fn func(data *T) (Continue bool),
	data *T, notify func(data *T)) uint32 {

	var f, n uptr
	f = cc.CbkRaw[func(data uptr) int32](func(out, ins uptr) {
		ret := (*int32)(out)
		if fn(cc.RawInAddr[T](ins, 0)) {
			(*ret) = 1
		} else {
			(*ret) = 0
		}
	})
	n = cc.CbkRaw[func(data uptr)](func(out, ins uptr) {
		notify(cc.RawInAddr[T](ins, 0))
		cc.CbkClose(f)
		cc.CbkCloseLate(n)
	})

	return g_idle_add_full.Fn()(priority, f, uptr(data), n)
}
func IdleAddOnce[T any](fn func(data *T), data *T) uint32 {
	var f uptr
	f = cc.CbkRaw[func(data uptr)](func(out, ins uptr) {
		fn(cc.RawInAddr[T](ins, 0))
		cc.CbkCloseLate(f)
	})
	return g_idle_add_once.Fn()(f, uptr(data))
}
func IdleRemoveByData[T any](data *T) bool { return g_idle_remove_by_data.Fn()(uptr(data)) != 0 }

// #endregion

// #region Mem
func GFree[T any](ptr *T)                   { g_free.Fn()(uptr(ptr)) }
func GFreeSized[T any](mem *T, size uint64) { g_free_sized.Fn()(uptr(mem), size) }

// func GClearPointer(pp *uptr, destroy cc.Func)    { g_clear_pointer.Fn()(pp, destroy) }

func GMalloc[T any](nBytes uint64) *T             { return (*T)(g_malloc.Fn()(nBytes)) }
func GMalloc0[T any](nBytes uint64) *T            { return (*T)(g_malloc0.Fn()(nBytes)) }
func GRealloc[T any](mem *T, nBytes uint64) *T    { return (*T)(g_realloc.Fn()(uptr(mem), nBytes)) }
func GTryMalloc[T any](nBytes uint64) *T          { return (*T)(g_try_malloc.Fn()(nBytes)) }
func GTryMalloc0[T any](nBytes uint64) *T         { return (*T)(g_try_malloc0.Fn()(nBytes)) }
func GTryRealloc[T any](mem *T, nBytes uint64) *T { return (*T)(g_try_realloc.Fn()(uptr(mem), nBytes)) }
func GMallocN[T any](nBlocks, nBlockBytes uint64) *T {
	return (*T)(g_malloc_n.Fn()(nBlocks, nBlockBytes))
}
func GMalloc0N[T any](nBlocks, nBlockBytes uint64) *T {
	return (*T)(g_malloc0_n.Fn()(nBlocks, nBlockBytes))
}
func GReallocN[T any](mem *T, nBlocks, nBlockBytes uint64) *T {
	return (*T)(g_realloc_n.Fn()(uptr(mem), nBlocks, nBlockBytes))
}
func GTryMallocN[T any](nBlocks, nBlockBytes uint64) *T {
	return (*T)(g_try_malloc_n.Fn()(nBlocks, nBlockBytes))
}
func GTryMalloc0N[T any](nBlocks, nBlockBytes uint64) *T {
	return (*T)(g_try_malloc0_n.Fn()(nBlocks, nBlockBytes))
}
func GTryReallocN[T any](mem *T, nBlocks, nBlockBytes uint64) *T {
	return (*T)(g_try_realloc_n.Fn()(uptr(mem), nBlocks, nBlockBytes))
}
func GAlignedAlloc[T any](nBlocks, nBlockBytes, alignment uint64) *T {
	return (*T)(g_aligned_alloc.Fn()(nBlocks, nBlockBytes, alignment))
}
func GAlignedAlloc0[T any](nBlocks, nBlockBytes, alignment uint64) *T {
	return (*T)(g_aligned_alloc0.Fn()(nBlocks, nBlockBytes, alignment))
}
func GAlignedFree[T any](mem *T) { g_aligned_free.Fn()(uptr(mem)) }
func GAlignedFreeSized[T any](mem *T, alignment, size uint64) {
	g_aligned_free_sized.Fn()(uptr(mem), alignment, size)
}

// #endregion

// #region OptionContext

type GOptionContext struct{}
type GOptionGroup struct{}
type GOptionEntry struct {
	LongName  cc.String
	ShortName byte
	Flags     GOptionFlags

	Arg     GOptionArg
	ArgData uptr

	Description    cc.String
	ArgDescription cc.String
}

func MakeGOptionEntry[T any](longName string, shortName byte,
	flags GOptionFlags, arg GOptionArg, argAddr *T, description, argDescription string) GOptionEntry {
	e := GOptionEntry{}
	e.LongName = cc.NewString(longName)
	e.ShortName = shortName
	e.Flags = flags
	e.Arg = arg
	e.ArgData = uptr(argAddr)
	e.Description = cc.NewString(description)
	e.ArgDescription = cc.NewString(argDescription)
	return e
}
func (e GOptionEntry) Unref() {
	e.LongName.Free()
	e.Description.Free()
	e.ArgDescription.Free()
}
func NewGOptionContext(parameterString string) *GOptionContext {
	ps := cc.NewString(parameterString)
	defer ps.Free()
	return g_option_context_new.Fn()(ps)
}
func (ctx *GOptionContext) SetSummary(summary string) {
	s := cc.NewString(summary)
	defer s.Free()
	g_option_context_set_summary.Fn()(ctx, s)
}
func (ctx *GOptionContext) GetSummary() string {
	return g_option_context_get_summary.Fn()(ctx).String()
}
func (ctx *GOptionContext) SetDescription(description string) {
	des := cc.NewString(description)
	defer des.Free()
	g_option_context_set_description.Fn()(ctx, des)
}
func (ctx *GOptionContext) GetDescription() string {
	return g_option_context_get_description.Fn()(ctx).String()
}
func (ctx *GOptionContext) Free() { g_option_context_free.Fn()(ctx) }
func (ctx *GOptionContext) SetHelpEnabled(helpEnabled bool) {
	g_option_context_set_help_enabled.Fn()(ctx, cbool(helpEnabled))
}
func (ctx *GOptionContext) GetHelpEnabled() bool {
	return g_option_context_get_help_enabled.Fn()(ctx) != 0
}
func (ctx *GOptionContext) SetIgnoreUnknownOptions(ignoreUnknown bool) {
	g_option_context_set_ignore_unknown_options.Fn()(ctx, cbool(ignoreUnknown))
}
func (ctx *GOptionContext) GetIgnoreUnknownOptions() bool {
	return g_option_context_get_ignore_unknown_options.Fn()(ctx) != 0
}
func (ctx *GOptionContext) SetStrictPosix(strictPosix bool) {
	g_option_context_set_strict_posix.Fn()(ctx, cbool(strictPosix))
}
func (ctx *GOptionContext) GetStrictPosix() bool {
	return g_option_context_get_strict_posix.Fn()(ctx) != 0
}
func (ctx *GOptionContext) AddMainEntries(entries []GOptionEntry, translationDomain string) {
	t := cc.NewString(translationDomain)
	defer t.Free()
	entries = append(entries, GOptionEntry{})
	e, _ := carry[GOptionEntry, int](entries)
	g_option_context_add_main_entries.Fn()(ctx, e, t)
	_ = entries
}
func (ctx *GOptionContext) Parse(args []string) error {
	argv, argc := cc.NewStringsLen(args)
	var gerr *GError
	if g_option_context_parse.Fn()(ctx, &argc, &argv, &gerr) == 0 {
		return gerr.TakeError()
	}
	return nil
}
func (ctx *GOptionContext) ParseStrv(args []string) error {
	argv := cc.NewStrings(args)
	var gerr *GError
	if g_option_context_parse_strv.Fn()(ctx, &argv, &gerr) == 0 {
		return gerr.TakeError()
	}
	return nil
}
func (ctx *GOptionContext) SetTranslateFunc(translate func(str cc.String, _ uptr) cc.String) {
	t := cc.CbkRaw[func(str cc.String, _ uptr) cc.String](func(out, ins uptr) {
		i := unsafe.Slice((*uptr)(ins), 2)
		i1 := (*cc.String)(i[0])
		i2 := (*uptr)(i[2])

		*(*cc.String)(out) = translate(*i1, *i2)
	})
	var des uptr
	des = cc.Cbk(func(_ uptr) {
		cc.CbkClose(t)
		cc.CbkCloseLate(des)
	})

	g_option_context_set_translate_func.Fn()(ctx, t, nil, des)
}
func (ctx *GOptionContext) SetTranslationDomain(domain string) {
	d := cc.NewString(domain)
	defer d.Free()
	g_option_context_set_translation_domain.Fn()(ctx, d)
}
func (ctx *GOptionContext) AddGroup(group *GOptionGroup) { g_option_context_add_group.Fn()(ctx, group) }
func (ctx *GOptionContext) SetMainGroup(group *GOptionGroup) {
	g_option_context_set_main_group.Fn()(ctx, group)
}
func (ctx *GOptionContext) GetMainGroup() *GOptionGroup {
	return g_option_context_get_main_group.Fn()(ctx)
}
func (ctx *GOptionContext) GetHelp(mainHelp bool, group *GOptionGroup) string {
	return g_option_context_get_help.Fn()(ctx, cbool(mainHelp), group).TakeString()
}

func NewGOptionGroup(name, description, helpDescription string, destroy func()) *GOptionGroup {
	des, n, d, h := uptr(nil), cc.NewString(name), cc.NewString(description), cc.NewString(helpDescription)
	cc.Cbk(func(_ uptr) {
		destroy()
		cc.CbkCloseLate(des)
	})
	return g_option_group_new.Fn()(n, d, h, nil, des)
}
func (group *GOptionGroup) SetParseHooks(preParseFunc, postParseFunc cc.Func,

/* gboolean(* GOptionParseFunc) (GOptionContext* context, GOptionGroup* group, gpointer data, GError** error) */) {
	g_option_group_set_parse_hooks.Fn()(group, preParseFunc, postParseFunc)
}
func (group *GOptionGroup) SetErrorHook(errorFunc cc.Func,

/* void (* GOptionErrorFunc) ( GOptionContext* context, GOptionGroup* group, gpointer data, GError** error) */) {
	g_option_group_set_error_hook.Fn()(group, errorFunc)
}
func (group *GOptionGroup) Ref() *GOptionGroup { return g_option_group_ref.Fn()(group) }
func (group *GOptionGroup) Unref()             { g_option_group_unref.Fn()(group) }
func (group *GOptionGroup) AddEntries(entries *GOptionEntry) {
	g_option_group_add_entries.Fn()(group, entries)
}
func (group *GOptionGroup) SetTranslateFunc(translate func(str cc.String, _ uptr) cc.String) {
	t := cc.CbkRaw[func(str cc.String, _ uptr) cc.String](func(out, ins uptr) {
		i := unsafe.Slice((*uptr)(ins), 2)
		i1 := (*cc.String)(i[0])
		i2 := (*uptr)(i[2])

		*(*cc.String)(out) = translate(*i1, *i2)
	})
	var des uptr
	des = cc.Cbk(func(_ uptr) {
		cc.CbkClose(t)
		cc.CbkCloseLate(des)
	})

	g_option_group_set_translate_func.Fn()(group, t, nil, des)
}
func (group *GOptionGroup) SetTranslationDomain(domain string) {
	d := cc.NewString(domain)
	defer d.Free()
	g_option_group_set_translation_domain.Fn()(group, d)
}

// #endregion

// #region PollFD

type GPollFD struct {
	FD      int
	Events  uint16
	REvents uint16
}

// #endregion

// #region Quark

type GQuark uint32

func GQuarkTryString(str string) GQuark {
	s := cc.NewString(str)
	defer s.Free()
	return g_quark_try_string.Fn()(s)
}
func GQuarkFromStaticString(str string) GQuark {
	s := cc.NewString(str)
	defer s.Free()
	return g_quark_from_static_string.Fn()(s)
}
func GQuarkFromString(str string) GQuark {
	s := cc.NewString(str)
	defer s.Free()
	return g_quark_from_string.Fn()(s)
}
func (quark GQuark) ToString() string { return g_quark_to_string.Fn()(quark).String() }
func (quark GQuark) String() string   { return g_quark_to_string.Fn()(quark).String() }

// #endregion

// #region SList

type GSList[T any] struct {
	Data *T
	Next *GSList[T]
}

func NewGSList[T any]() *GSList[T] { return (*GSList[T])(g_slist_alloc.Fn()()) }
func (l *GSList[T]) Free()         { g_slist_free.Fn()(uptr(l)) }
func (l *GSList[T]) Free1()        { g_slist_free_1.Fn()(uptr(l)) }
func (l *GSList[T]) FreeFull(freeFunc func(data *T)) {
	ff := cc.Cbk(freeFunc)
	defer cc.CbkClose(ff)
	g_slist_free_full.Fn()(uptr(l), ff)
}
func (l *GSList[T]) Append(data *T) *GSList[T] {
	return (*GSList[T])(g_slist_append.Fn()(uptr(l), uptr(data)))
}
func (l *GSList[T]) Prepend(data *T) *GSList[T] {
	return (*GSList[T])(g_slist_prepend.Fn()(uptr(l), uptr(data)))
}
func (l *GSList[T]) Insert(data *T, position int32) *GSList[T] {
	return (*GSList[T])(g_slist_insert.Fn()(uptr(l), uptr(data), position))
}
func (l *GSList[T]) InsertSorted(data *T, compareFunc func(a, b *T) int32) *GSList[T] {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	return (*GSList[T])(g_slist_insert_sorted.Fn()(uptr(l), uptr(data), cf))
}

//	func (l *GSList[T]) InsertSortedWithData(data interface{}, compareDataFunc func(a, b, userData UPtr) int32, userData interface{}) *GSList {
//		return g_slist_insert_sorted_with_data.Fn()(UPtr(l), anyptr(data), cc.Cbk(compareDataFunc), anyptr(userData))
//	}

func (l *GSList[T]) InsertBefore(sibling *GSList[T], data *T) *GSList[T] {
	return (*GSList[T])(g_slist_insert_before.Fn()(uptr(l), uptr(sibling), uptr(data)))
}
func (l *GSList[T]) Concat(list2 *GSList[T]) *GSList[T] {
	return (*GSList[T])(g_slist_concat.Fn()(uptr(l), uptr(list2)))
}
func (l *GSList[T]) Remove(data *T) *GSList[T] {
	return (*GSList[T])(g_slist_remove.Fn()(uptr(l), uptr(data)))
}
func (l *GSList[T]) RemoveAll(data *T) *GSList[T] {
	return (*GSList[T])(g_slist_remove_all.Fn()(uptr(l), uptr(data)))
}
func (l *GSList[T]) RemoveLink(link *GSList[T]) *GSList[T] {
	return (*GSList[T])(g_slist_remove_link.Fn()(uptr(l), uptr(link)))
}
func (l *GSList[T]) DeleteLink(link *GSList[T]) *GSList[T] {
	return (*GSList[T])(g_slist_delete_link.Fn()(uptr(l), uptr(link)))
}
func (l *GSList[T]) Reverse() *GSList[T] { return (*GSList[T])(g_slist_reverse.Fn()(uptr(l))) }
func (l *GSList[T]) Copy() *GSList[T]    { return (*GSList[T])(g_slist_copy.Fn()(uptr(l))) }
func (l *GSList[T]) CopyDeep(copyFunc func(src *T, _ uptr) *T) *GSList[T] {
	cf := cc.Cbk(copyFunc)
	defer cc.CbkClose(cf)
	return (*GSList[T])(g_slist_copy_deep.Fn()(uptr(l), cf, nil))
}
func (l *GSList[T]) Nth(n uint32) *GSList[T] { return (*GSList[T])(g_slist_nth.Fn()(uptr(l), n)) }
func (l *GSList[T]) Find(data *T) *GSList[T] {
	return (*GSList[T])(g_slist_find.Fn()(uptr(l), uptr(data)))
}
func (l *GSList[T]) FindCustom(data *T, compareFunc func(curr, data *T) int32) *GSList[T] {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	return (*GSList[T])(g_slist_find_custom.Fn()(uptr(l), uptr(data), cf))
}
func (l *GSList[T]) Position(link *GSList[T]) int32 {
	return g_slist_position.Fn()(uptr(l), uptr(link))
}
func (l *GSList[T]) Index(data *T) int32 { return g_slist_index.Fn()(uptr(l), uptr(data)) }
func (l *GSList[T]) Last() *GSList[T]    { return (*GSList[T])(g_slist_last.Fn()(uptr(l))) }
func (l *GSList[T]) Length() uint32      { return g_slist_length.Fn()(uptr(l)) }
func (l *GSList[T]) ForEach(fn func(data *T, _ uptr)) {
	f := cc.Cbk(fn)
	defer cc.CbkClose(f)
	g_slist_foreach.Fn()(uptr(l), f, nil)
}
func (l *GSList[T]) Sort(compareFunc func(a, b *T) int32) *GSList[T] {
	cf := cc.Cbk(compareFunc)
	defer cc.CbkClose(cf)
	return (*GSList[T])(g_slist_sort.Fn()(uptr(l), cf))
}

//	func (l *GSList[T]) SortWithData(compareDataFunc func(a, b, userData UPtr) int32, userData interface{}) *GSList {
//		return g_slist_sort_with_data.Fn()(UPtr(l), cc.Cbk(compareDataFunc), anyptr(userData))
//	}

func (l *GSList[T]) NthData(n uint32) uptr { return g_slist_nth_data.Fn()(uptr(l), n) }
func (l *GSList[T]) Clear(destroy func(data uptr)) {
	addr := uptr(l)
	g_clear_slist.Fn()(&addr, cc.Cbk(destroy))
}
func (l *GSList[T]) List() (lst []*T) {
	lst = []*T{}
	l.ForEach(func(data *T, _ uptr) {
		lst = append(lst, data)
	})
	return
}

// func GSListList[T any](l *GSList, close func(*GSList)) (list []T) {
// 	type pinner struct {
// 		list  *GSList
// 		slice *[]T
// 	}
// 	num := l.Length()
// 	if num == 0 {
// 		return nil
// 	}

// 	pin := new(pinner)
// 	pin.list = l
// 	slice := make([]T, num)
// 	pin.slice = &slice

// 	runtime.SetFinalizer(pin, func(p *pinner) {
// 		if close != nil {
// 			close(p.list)
// 			return
// 		}
// 		p.list.Free()

// 		_ = *p.slice
// 	})

// 	var t T
// 	switch any(t).(type) {
// 	case string:
// 		for i := uint32(0); i < num; i++ {
// 			a := ((cc.String)(l.Nth(i).Data)).String()
// 			(*pin.slice)[i] = any(a).(T)
// 		}
// 	default:
// 		for i := uint32(0); i < num; i++ {
// 			(*pin.slice)[i] = *(*T)(l.Nth(i).Data)
// 		}
// 	}

// 	return *pin.slice
// }

// #endregion

// #region String

type GString struct {
	ptr uptr
	len uint64
	alc uint64
}

func (s *GString) String() string { return *(*string)(uptr(s)) }
func (s *GString) Ptr() uptr      { return s.ptr }
func (s *GString) Len() uint64    { return s.len }
func (s *GString) Alc() uint64    { return s.alc }

func NewGString(init string) *GString {
	s := cc.NewString(init)
	defer s.Free()
	return g_string_new.Fn()(s)
}
func NewGStringLen(init string, len int) *GString {
	s := cc.NewString(init)
	defer s.Free()
	return g_string_new_len.Fn()(s, len)
}
func NewGStringTake(ptr cc.String) *GString { return g_string_new_take.Fn()(ptr) }
func NewGstringSized(dflSize uint) *GString { return g_string_sized_new.Fn()(dflSize) }

func (s *GString) Append(val string) *GString {
	v := cc.NewString(val)
	defer v.Free()
	return g_string_append.Fn()(s, v)
}
func (s *GString) AppendC(c byte) *GString { return g_string_append_c.Fn()(s, c) }
func (s *GString) AppendLen(val string, len int) *GString {
	ss := cc.NewString(val)
	defer ss.Free()
	return g_string_append_len.Fn()(s, ss, len)
}
func (s *GString) AppendPrintf(format string, args ...interface{}) {
	f := cc.NewString(fmt.Sprintf(format, args...))
	defer f.Free()
	g_string_append_printf.Fn()(s, f)
}
func (s *GString) AppendUnichar(c uint32) *GString { return g_string_append_unichar.Fn()(s, c) }
func (s *GString) AppendUriEscaped(unescaped, reservedCharsAllowed string, allowUtf8 bool) *GString {
	u, r := cc.NewString(unescaped), cc.NewString(reservedCharsAllowed)
	defer u.Free()
	defer r.Free()
	return g_string_append_uri_escaped.Fn()(s, u, r, cbool(allowUtf8))
}
func (s *GString) AsciiDown() *GString { return g_string_ascii_down.Fn()(s) }
func (s *GString) AsciiUp() *GString   { return g_string_ascii_up.Fn()(s) }
func (s *GString) Assign(val string) *GString {
	ss := cc.NewString(val)
	defer ss.Free()
	return g_string_assign.Fn()(s, ss)
}
func (s *GString) Down() *GString              { return g_string_down.Fn()(s) }
func (s *GString) Equal(s2 *GString) bool      { return g_string_equal.Fn()(s, s2) != 0 }
func (s *GString) Erase(pos, len int) *GString { return g_string_erase.Fn()(s, pos, len) }
func (s *GString) Free(freeSegment bool) uptr  { return g_string_free.Fn()(s, cbool(freeSegment)) }
func (s *GString) FreeAndSteal() uptr          { return g_string_free_and_steal.Fn()(s) }
func (s *GString) FreeToBytes() *GBytes        { return g_string_free_to_bytes.Fn()(s) }
func (s *GString) Hash() uint                  { return g_string_hash.Fn()(s) }
func (s *GString) Insert(pos int, val string) *GString {
	ss := cc.NewString(val)
	defer ss.Free()
	return g_string_insert.Fn()(s, pos, ss)
}

func (s *GString) InsertC(pos int, c byte) *GString { return g_string_insert_c.Fn()(s, pos, c) }
func (s *GString) InsertLen(pos int, val string, len int) *GString {
	ss := cc.NewString(val)
	defer ss.Free()
	return g_string_insert_len.Fn()(s, pos, ss, len)
}
func (s *GString) InsertUnichar(pos int, c uint32) *GString {
	return g_string_insert_unichar.Fn()(s, pos, c)
}
func (s *GString) Overwrite(pos uint, val string) *GString {
	ss := cc.NewString(val)
	defer ss.Free()
	return g_string_overwrite.Fn()(s, pos, ss)
}
func (s *GString) OverwriteLen(pos uint, val string, len int) *GString {
	ss := cc.NewString(val)
	defer ss.Free()
	return g_string_overwrite_len.Fn()(s, pos, ss, len)
}
func (s *GString) Prepend(val string) *GString {
	ss := cc.NewString(val)
	defer ss.Free()
	return g_string_prepend.Fn()(s, ss)
}
func (s *GString) PrependC(c byte) *GString { return g_string_prepend_c.Fn()(s, c) }
func (s *GString) PrependLen(val string, len int) *GString {
	ss := cc.NewString(val)
	defer ss.Free()
	return g_string_prepend_len.Fn()(s, ss, len)
}
func (s *GString) PrependUnichar(c uint32) *GString { return g_string_prepend_unichar.Fn()(s, c) }
func (s *GString) Printf(format string, args ...interface{}) {
	f := cc.NewString(fmt.Sprintf(format, args...))
	defer f.Free()
	g_string_printf.Fn()(s, f)
}
func (s *GString) Replace(find, replace string, limit uint32) uint32 {
	f, r := cc.NewString(find), cc.NewString(replace)
	defer f.Free()
	defer r.Free()
	return g_string_replace.Fn()(s, f, r, limit)
}
func (s *GString) SetSize(len uint) *GString  { return g_string_set_size.Fn()(s, len) }
func (s *GString) Truncate(len uint) *GString { return g_string_truncate.Fn()(s, len) }
func (s *GString) Up() *GString               { return g_string_up.Fn()(s) }

// #endregion

// #region TimeZone

type GTimeZone struct{}

// func NewGTimeZone(identifier string) *GTimeZone {
// 	idt := cc.NewString(identifier)
// 	defer idt.Free()
// 	return g_time_zone_new.Fn()(idt)
// }

func NewGTimeZoneIdentifier(identifier string) *GTimeZone {
	idt := cc.NewString(identifier)
	defer idt.Free()
	return g_time_zone_new_identifier.Fn()(idt)
}
func NewGTimeZoneUtc() *GTimeZone                 { return g_time_zone_new_utc.Fn()() }
func NewGTimeZoneLocal() *GTimeZone               { return g_time_zone_new_local.Fn()() }
func NewGTimeZoneOffset(seconds int32) *GTimeZone { return g_time_zone_new_offset.Fn()(seconds) }

func (tz *GTimeZone) Ref() *GTimeZone { return g_time_zone_ref.Fn()(tz) }
func (tz *GTimeZone) Unref()          { g_time_zone_unref.Fn()(tz) }

func (tz *GTimeZone) FindInterval(typ GTimeType, time int64) int {
	return g_time_zone_find_interval.Fn()(tz, typ, time)
}
func (tz *GTimeZone) AdjustTime(typ GTimeType, time *int64) int {
	return g_time_zone_adjust_time.Fn()(tz, typ, time)
}

func (tz *GTimeZone) GetAbbreviation(interval int) string {
	return g_time_zone_get_abbreviation.Fn()(tz, interval).String()
}
func (tz *GTimeZone) GetOffset(interval int) int32 { return g_time_zone_get_offset.Fn()(tz, interval) }
func (tz *GTimeZone) IsDst(interval int) bool      { return g_time_zone_is_dst.Fn()(tz, interval) != 0 }
func (tz *GTimeZone) GetIdentifier() string        { return g_time_zone_get_identifier.Fn()(tz).String() }

// #endregion

// #region Variant

type GVariantType struct{}

func NewGVariantType[T string | GVariantClass](typ T) *GVariantType {
	switch tp := any(typ).(type) {
	case string:
		t := cc.NewString(tp)
		defer t.Free()
		return g_variant_type_new.Fn()(t)
	case GVariantClass:
		t := cc.NewString(string(byte(tp)))
		defer t.Free()
		return g_variant_type_new.Fn()(t)
	}
	return nil
}

func NewGVariantTypeArray(elementType *GVariantType) *GVariantType {
	return g_variant_type_new_array.Fn()(elementType)
}
func NewGVariantTypeDictEntry(keyType, valueType *GVariantType) *GVariantType {
	return g_variant_type_new_dict_entry.Fn()(keyType, valueType)
}
func NewGVariantTypeMaybe(elementType *GVariantType) *GVariantType {
	return g_variant_type_new_maybe.Fn()(elementType)
}
func NewGVariantTypeTuple(items []*GVariantType) *GVariantType {
	return g_variant_type_new_tuple.Fn()(carry[*GVariantType, int32](items))
}

func (t *GVariantType) Copy() *GVariantType    { return g_variant_type_copy.Fn()(t) }
func (t *GVariantType) DupString() string      { return g_variant_type_dup_string.Fn()(t).TakeString() }
func (t *GVariantType) Element() *GVariantType { return g_variant_type_element.Fn()(t) }
func (t *GVariantType) Equal(other *GVariantType) bool {
	return g_variant_type_equal.Fn()(t, other) != 0
}
func (t *GVariantType) First() *GVariantType  { return g_variant_type_first.Fn()(t) }
func (t *GVariantType) Free()                 { g_variant_type_free.Fn()(t) }
func (t *GVariantType) GetStringLength() uint { return g_variant_type_get_string_length.Fn()(t) }
func (t *GVariantType) Hash() uint32          { return g_variant_type_hash.Fn()(uptr(t)) }
func (t *GVariantType) IsArray() bool         { return g_variant_type_is_array.Fn()(t) != 0 }
func (t *GVariantType) IsBasic() bool         { return g_variant_type_is_basic.Fn()(t) != 0 }
func (t *GVariantType) IsContainer() bool     { return g_variant_type_is_container.Fn()(t) != 0 }
func (t *GVariantType) IsDefinite() bool      { return g_variant_type_is_definite.Fn()(t) != 0 }
func (t *GVariantType) IsDictEntry() bool     { return g_variant_type_is_dict_entry.Fn()(t) != 0 }
func (t *GVariantType) IsMaybe() bool         { return g_variant_type_is_maybe.Fn()(t) != 0 }
func (t *GVariantType) IsSubtypeOf(supertype *GVariantType) bool {
	return g_variant_type_is_subtype_of.Fn()(t, supertype) != 0
}
func (t *GVariantType) IsTuple() bool        { return g_variant_type_is_tuple.Fn()(t) != 0 }
func (t *GVariantType) IsVariant() bool      { return g_variant_type_is_variant.Fn()(t) != 0 }
func (t *GVariantType) Key() *GVariantType   { return g_variant_type_key.Fn()(t) }
func (t *GVariantType) NItems() uint         { return g_variant_type_n_items.Fn()(t) }
func (t *GVariantType) Next() *GVariantType  { return g_variant_type_next.Fn()(t) }
func (t *GVariantType) PeekString() string   { return g_variant_type_peek_string.Fn()(t).String() }
func (t *GVariantType) Value() *GVariantType { return g_variant_type_value.Fn()(t) }

type GVariant struct{}

func NewGVariant(format string, args ...interface{}) *GVariant {
	f := cc.NewString(format)
	defer f.Free()
	return g_variant_new.FnVa()(f, args...)
}
func NewGVariantArray(elementType *GVariantType, children []*GVariant) *GVariant {
	c, l := carry[*GVariant, int64](children)
	return g_variant_new_array.Fn()(elementType, c, l)
}

func NewGVariantBoolean(value bool) *GVariant {
	return g_variant_new_boolean.Fn()(cbool(value))
}
func NewGVariantByte(value byte) *GVariant {
	return g_variant_new_byte.Fn()(value)
}

func NewGVariantByteString(str string) *GVariant {
	s := cc.NewString(str)
	defer s.Free()
	return g_variant_new_bytestring.Fn()(s)
}

func NewGVariantByteStringArray(strv []string) *GVariant {
	p, l := cc.NewStringsLen(strv)
	defer p.Free()
	return g_variant_new_bytestring_array.Fn()(p, int64(l))
}

func NewGVariantDictEntry(key, value *GVariant) *GVariant {
	return g_variant_new_dict_entry.Fn()(key, value)
}

func NewGVariantDouble(value float64) *GVariant {
	return g_variant_new_double.Fn()(value)
}

func NewGVariantFixedArray[T any](elementType *GVariantType, elements []T) *GVariant {
	var t T
	p, l := carry[T, uint64](elements)
	return g_variant_new_fixed_array.Fn()(elementType, uptr(p), l, uint64(unsafe.Sizeof(t)))
}

func NewGVariantFromBytes(typ *GVariantType, bytes *GBytes, trusted bool) *GVariant {
	return g_variant_new_from_bytes.Fn()(typ, bytes, cbool(trusted))
}

// func (ns nameSpaceGVariant) NewFromData(type_ *GVariantType, data UPtr, size int64, trusted bool, notify GDestroyNotify, userData UPtr) *GVariant {
// 	return g_variant_new_from_data(type_, data, size, trusted, notify, userData)
// }

func NewGVariantHandle(value int32) *GVariant { return g_variant_new_handle.Fn()(value) }
func NewGVariantInt16(value int16) *GVariant  { return g_variant_new_int16.Fn()(value) }
func NewGVariantInt32(value int32) *GVariant  { return g_variant_new_int32.Fn()(value) }
func NewGVariantInt64(value int64) *GVariant  { return g_variant_new_int64.Fn()(value) }
func NewGVariantMaybe(childType *GVariantType, child *GVariant) *GVariant {
	return g_variant_new_maybe.Fn()(childType, child)
}
func NewGVariantObjectPath(objectPath string) *GVariant {
	ob := cc.NewString(objectPath)
	defer ob.Free()
	return g_variant_new_object_path.Fn()(ob)
}

func NewGVariantObjv(strv []string) *GVariant {
	p, l := cc.NewStringsLen(strv)
	defer p.Free()
	return g_variant_new_objv.Fn()(p, int64(l))
}

func NewGVariantParsed(format string, args ...interface{}) *GVariant {
	f := cc.NewString(fmt.Sprintf(format, args...))
	defer f.Free()
	return g_variant_new_parsed.Fn()(f)
}

func NewGVariantPrintf(format string, args ...interface{}) *GVariant {
	f := cc.NewString(fmt.Sprintf(format, args...))
	defer f.Free()
	return g_variant_new_printf.Fn()(f)
}

func NewGVariantSignature(signature string) *GVariant {
	s := cc.NewString(signature)
	defer s.Free()
	return g_variant_new_signature.Fn()(s)
}

func NewGVariantString(str string) *GVariant {
	s := cc.NewString(str)
	defer s.Free()
	return g_variant_new_string.Fn()(s)
}

func NewGVariantStrv(strv []string) *GVariant {
	p, l := cc.NewStringsLen(strv)
	defer p.Free()
	return g_variant_new_strv.Fn()(p, int64(l))
}

func NewGVariantTuple(children []*GVariant) *GVariant {
	return g_variant_new_tuple.Fn()(carry[*GVariant, uint64](children))
}

func NewGVariantUint16(value uint16) *GVariant {
	return g_variant_new_uint16.Fn()(value)
}
func NewGVariantUint32(value uint32) *GVariant {
	return g_variant_new_uint32.Fn()(value)
}
func NewGVariantUint64(value uint64) *GVariant {
	return g_variant_new_uint64.Fn()(value)
}
func NewGVariantVariant(value *GVariant) *GVariant {
	return g_variant_new_variant.Fn()(value)
}

func (v *GVariant) Byteswap() *GVariant { return g_variant_byteswap.Fn()(v) }
func (v *GVariant) CheckFormatString(format string, copyOnly bool) bool {
	f := cc.NewString(format)
	defer f.Free()
	return g_variant_check_format_string.Fn()(v, f, cbool(copyOnly)) != 0
}
func (v *GVariant) Classify() GVariantClass      { return g_variant_classify.Fn()(v) }
func (v *GVariant) Compare(other *GVariant) bool { return g_variant_compare.Fn()(v, other) != 0 }
func (v *GVariant) DupByteString() string {
	var l uint64
	return g_variant_dup_bytestring.Fn()(v, &l).TakeString(l)
}
func (v *GVariant) DupByteStringArray() []string {
	var l uint64
	return g_variant_dup_bytestring_array.Fn()(v, &l).TakeStrings()
}
func (v *GVariant) DupObjv() []string {
	var l uint64
	return g_variant_dup_objv.Fn()(v, &l).TakeStrings(l)
}
func (v *GVariant) DupString() string {
	var l uint64
	return g_variant_dup_string.Fn()(v, &l).TakeString(l)
}
func (v *GVariant) DupStrv() []string {
	var l uint64
	return g_variant_dup_strv.Fn()(v, &l).TakeStrings(l)
}
func (v *GVariant) Equal(other *GVariant) bool { return g_variant_equal.Fn()(v, other) != 0 }

// func (v *GVariant) Get(format string, ptr interface{}) { g_variant_get.Fn()(v, format, anyptr(ptr)) }

func GVariantGet[T any](v *GVariant, format string, ptr *T) {
	f := cc.NewString(format)
	defer f.Free()
	g_variant_get.Fn()(v, f, uptr(ptr))
}
func (v *GVariant) GetBoolean() bool      { return g_variant_get_boolean.Fn()(v) != 0 }
func (v *GVariant) GetByte() byte         { return g_variant_get_byte.Fn()(v) }
func (v *GVariant) GetByteString() string { return g_variant_get_bytestring.Fn()(v).TakeString() }
func (v *GVariant) GetByteStringArray() []string {
	var l uint64
	return g_variant_get_bytestring_array.Fn()(v, &l).TakeStrings(l)
}
func GVariantGetChild[T any](v *GVariant, index uint64, format string, ptr *T) {
	f := cc.NewString(format)
	defer f.Free()
	g_variant_get_child.Fn()(v, index, f, uptr(ptr))
}
func (v *GVariant) GetChildValue(index uint64) *GVariant {
	return g_variant_get_child_value.Fn()(v, index)
}
func (v *GVariant) GetData() uptr           { return g_variant_get_data.Fn()(v) }
func (v *GVariant) GetDataAsBytes() *GBytes { return g_variant_get_data_as_bytes.Fn()(v) }
func (v *GVariant) GetDouble() float64      { return g_variant_get_double.Fn()(v) }
func (v *GVariant) GetFixedArray(elementSize uint64) (firstPtr uptr, nElements uint64) {
	firstPtr = g_variant_get_fixed_array.Fn()(v, &nElements, elementSize)
	return
}
func GVariantGetFixedArray[T any](v *GVariant) []T {
	var l uint64
	var t T
	p := g_variant_get_fixed_array.Fn()(v, &l, uint64(unsafe.Sizeof(t)))
	return unsafe.Slice((*T)(p), l)
}
func (v *GVariant) GetHandle() int32         { return g_variant_get_handle.Fn()(v) }
func (v *GVariant) GetInt16() int16          { return g_variant_get_int16.Fn()(v) }
func (v *GVariant) GetInt32() int32          { return g_variant_get_int32.Fn()(v) }
func (v *GVariant) GetInt64() int64          { return g_variant_get_int64.Fn()(v) }
func (v *GVariant) GetMaybe() *GVariant      { return g_variant_get_maybe.Fn()(v) }
func (v *GVariant) GetNormalForm() *GVariant { return g_variant_get_normal_form.Fn()(v) }
func (v *GVariant) GetObjv() []string {
	var l uint64
	cArray := g_variant_get_objv.Fn()(v, &l)
	defer cc.Free(uptr(cArray))
	return cArray.Strings(l)
}
func (v *GVariant) GetSize() uint { return g_variant_get_size.Fn()(v) }
func (v *GVariant) GetString() string {
	var length uint64
	return g_variant_get_string.Fn()(v, &length).String(length)
}
func (v *GVariant) GetStrv() []string {
	var length uint64
	return g_variant_get_strv.Fn()(v, &length).Strings(length)
}
func (v *GVariant) GetType() *GVariantType { return g_variant_get_type.Fn()(v) }
func (v *GVariant) GetTypeString() string  { return g_variant_get_type_string.Fn()(v).String() }
func (v *GVariant) GetUint16() uint16      { return g_variant_get_uint16.Fn()(v) }
func (v *GVariant) GetUint32() uint32      { return g_variant_get_uint32.Fn()(v) }
func (v *GVariant) GetUint64() uint64      { return g_variant_get_uint64.Fn()(v) }
func (v *GVariant) GetVariant() *GVariant  { return g_variant_get_variant.Fn()(v) }
func (v *GVariant) Hash() uint64           { return g_variant_hash.Fn()(v) }
func (v *GVariant) IsContainer() bool      { return g_variant_is_container.Fn()(v) != 0 }
func (v *GVariant) IsFloating() bool       { return g_variant_is_floating.Fn()(v) != 0 }
func (v *GVariant) IsNormalForm() bool     { return g_variant_is_normal_form.Fn()(v) != 0 }
func (v *GVariant) IsOfType(typ *GVariantType) bool {
	return g_variant_is_of_type.Fn()(v, typ) != 0
}

func (v *GVariant) Lookup(key string, format string, args ...interface{}) bool {
	k, f := cc.NewString(key), cc.NewString(format)
	defer k.Free()
	defer k.Free()
	return g_variant_lookup.FnVa()(v, k, f, args...) != 0
}
func (v *GVariant) LookupValue(key string, typ *GVariantType) *GVariant {
	k := cc.NewString(key)
	defer k.Free()
	return g_variant_lookup_value.Fn()(v, k, typ)
}
func (v *GVariant) NChildren() uint64 { return g_variant_n_children.Fn()(v) }
func (v *GVariant) Print(typeAnnotate bool) string {
	return g_variant_print.Fn()(v, cbool(typeAnnotate)).TakeString()
}
func (v *GVariant) PrintString(str *GString, typeAnnotate bool) *GString {
	return g_variant_print_string.Fn()(v, str, cbool(typeAnnotate))
}
func (v *GVariant) Ref() *GVariant     { return g_variant_ref.Fn()(v) }
func (v *GVariant) RefSink() *GVariant { return g_variant_ref_sink.Fn()(v) }
func (v *GVariant) Store(data uptr)    { g_variant_store.Fn()(v, data) }
func (v *GVariant) TakeRef() *GVariant { return g_variant_take_ref.Fn()(v) }
func (v *GVariant) Unref()             { g_variant_unref.Fn()(v) }

type GVariantIter struct{}

func NewGVariantIter(value *GVariant) *GVariantIter { return g_variant_iter_new.Fn()(value) }
func (iter *GVariantIter) Init(value *GVariant) uint64 {
	return g_variant_iter_init.Fn()(iter, value)
}
func (iter *GVariantIter) Copy() *GVariantIter  { return g_variant_iter_copy.Fn()(iter) }
func (iter *GVariantIter) NChildren() uint64    { return g_variant_iter_n_children.Fn()(iter) }
func (iter *GVariantIter) Free()                { g_variant_iter_free.Fn()(iter) }
func (iter *GVariantIter) NextValue() *GVariant { return g_variant_iter_next_value.Fn()(iter) }
func (iter *GVariantIter) Next(formatString string, args ...interface{}) bool {
	f := cc.NewString(formatString)
	defer f.Free()
	return g_variant_iter_next.FnVa()(iter, f, args...) != 0
}
func (iter *GVariantIter) Loop(formatString string, args ...interface{}) bool {
	f := cc.NewString(formatString)
	defer f.Free()
	return g_variant_iter_loop.FnVa()(iter, f, args...) != 0
}

type GVariantBuilder struct{}

func NewGVariantBuilder(typ *GVariantType) *GVariantBuilder { return g_variant_builder_new.Fn()(typ) }
func (builder *GVariantBuilder) Unref()                     { g_variant_builder_unref.Fn()(builder) }
func (builder *GVariantBuilder) Ref() *GVariantBuilder      { return g_variant_builder_ref.Fn()(builder) }
func (builder *GVariantBuilder) Init(typ *GVariantType)     { g_variant_builder_init.Fn()(builder, typ) }
func (builder *GVariantBuilder) InitStatic(typ *GVariantType) {
	g_variant_builder_init_static.Fn()(builder, typ)
}
func (builder *GVariantBuilder) End() *GVariant         { return g_variant_builder_end.Fn()(builder) }
func (builder *GVariantBuilder) Clear()                 { g_variant_builder_clear.Fn()(builder) }
func (builder *GVariantBuilder) Open(typ *GVariantType) { g_variant_builder_open.Fn()(builder, typ) }
func (builder *GVariantBuilder) Close()                 { g_variant_builder_close.Fn()(builder) }
func (builder *GVariantBuilder) AddValue(value *GVariant) {
	g_variant_builder_add_value.Fn()(builder, value)
}
func (builder *GVariantBuilder) Add(formatString string, args ...interface{}) {
	ss := cc.NewString(formatString)
	defer ss.Free()
	g_variant_builder_add.FnVa()(builder, ss, args...)
}
func (builder *GVariantBuilder) AddParsed(formatString string, args ...interface{}) {
	ss := cc.NewString(formatString)
	defer ss.Free()
	g_variant_builder_add_parsed.FnVa()(builder, ss, args...)
}

type GVariantDict struct{}

func NewGVariantDict(fromAsv *GVariant) *GVariantDict { return g_variant_dict_new.Fn()(fromAsv) }
func (dict *GVariantDict) Init(fromAsv *GVariant)     { g_variant_dict_init.Fn()(dict, fromAsv) }
func (dict *GVariantDict) Lookup(key string, formatString string, args ...interface{}) bool {
	k, f := cc.NewString(key), cc.NewString(formatString)
	defer k.Free()
	defer f.Free()
	return g_variant_dict_lookup.FnVa()(dict, k, f, args...) != 0
}
func (dict *GVariantDict) LookupValue(key string, expectedType *GVariantType) *GVariant {
	k := cc.NewString(key)
	defer k.Free()
	return g_variant_dict_lookup_value.Fn()(dict, k, expectedType)
}
func (dict *GVariantDict) Contains(key string) bool {
	k := cc.NewString(key)
	defer k.Free()
	return g_variant_dict_contains.Fn()(dict, k) != 0
}
func (dict *GVariantDict) Insert(key, formatString string, args ...interface{}) {
	k, f := cc.NewString(key), cc.NewString(formatString)
	defer k.Free()
	defer f.Free()
	g_variant_dict_insert.FnVa()(dict, k, f, args...)
}
func (dict *GVariantDict) InsertValue(key string, value *GVariant) {
	k := cc.NewString(key)
	defer k.Free()
	g_variant_dict_insert_value.Fn()(dict, k, value)
}
func (dict *GVariantDict) Remove(key string) bool {
	k := cc.NewString(key)
	defer k.Free()
	return g_variant_dict_remove.Fn()(dict, k) != 0
}
func (dict *GVariantDict) Clear()             { g_variant_dict_clear.Fn()(dict) }
func (dict *GVariantDict) End() *GVariant     { return g_variant_dict_end.Fn()(dict) }
func (dict *GVariantDict) Ref() *GVariantDict { return g_variant_dict_ref.Fn()(dict) }
func (dict *GVariantDict) Unref()             { g_variant_dict_unref.Fn()(dict) }

// #endregion
