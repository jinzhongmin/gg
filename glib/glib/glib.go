package glib

import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/jinzhongmin/gg/cc"
)

// #region Alloc

func GAlloc[T any](size uint64) *T {
	return (*T)(malloc.Fn()(size))
}
func GAlloc0[T any](size uint64) *T {
	p := malloc.Fn()(size)
	memset.Fn()(p, 0, size)
	return (*T)(p)
}
func GNew[T any]() *T {
	var a T
	return GAlloc[T](uint64(unsafe.Sizeof(a)))
}
func GNew0[T any]() *T {
	var a T
	return GAlloc0[T](uint64(unsafe.Sizeof(a)))
}

// #endregion

// #region Array

type GArray struct {
	Data UPtr
	Len  uint32
}
type GByteArray struct {
	Data UPtr
	Len  uint32
}
type GPtrArray struct {
	Data UPtr
	Len  uint32
}

func NewGArray(zeroTerminated, clear bool, elementSize uint32) *GArray {
	return g_array_new.Fn()(zeroTerminated, clear, elementSize)
}
func NewGArrayTake(data UPtr, length uint64, clear bool, elementSize uint64) *GArray {
	return g_array_new_take.Fn()(data, length, clear, elementSize)
}
func NewGArrayTakeZeroTerminated(data UPtr, clear bool, elementSize uint64) *GArray {
	return g_array_new_take_zero_terminated.Fn()(data, clear, elementSize)
}
func (array *GArray) Steal() (data UPtr, length uint64) {
	data = g_array_steal.Fn()(array, &length)
	return
}
func NewGArraySized(zeroTerminated, clear bool, elementSize, reservedSize uint32) *GArray {
	return g_array_sized_new.Fn()(zeroTerminated, clear, elementSize, reservedSize)
}
func (array *GArray) Copy() *GArray              { return g_array_copy.Fn()(array) }
func (array *GArray) Free(freeSegment bool) UPtr { return g_array_free.Fn()(array, freeSegment) }
func (array *GArray) Ref() *GArray               { return g_array_ref.Fn()(array) }
func (array *GArray) Unref()                     { g_array_unref.Fn()(array) }
func (array *GArray) GetElementSize() uint32 {
	return g_array_get_element_size.Fn()(array)
}
func (array *GArray) AppendVals(data interface{}, length uint32) *GArray {
	return g_array_append_vals.Fn()(array, anyptr(data), length)
}
func (array *GArray) PrependVals(data interface{}, length uint32) *GArray {
	return g_array_prepend_vals.Fn()(array, anyptr(data), length)
}
func (array *GArray) InsertVals(index uint32, data interface{}, length uint32) *GArray {
	return g_array_insert_vals.Fn()(array, index, anyptr(data), length)
}
func (array *GArray) SetSize(length uint32) *GArray {
	return g_array_set_size.Fn()(array, length)
}
func (array *GArray) RemoveIndex(index uint32) *GArray {
	return g_array_remove_index.Fn()(array, index)
}
func (array *GArray) RemoveIndexFast(index uint32) *GArray {
	return g_array_remove_index_fast.Fn()(array, index)
}
func (array *GArray) RemoveRange(index uint32, length uint32) *GArray {
	return g_array_remove_range.Fn()(array, index, length)
}
func (array *GArray) Sort(compareFunc func(a, b UPtr) int32) {
	g_array_sort.Fn()(array, vcbu(compareFunc))
}
func (array *GArray) BinarySearch(target interface{}, compareFunc func(a, b UPtr) int32) (found bool, matchIndex uint32) {
	found = g_array_binary_search.Fn()(array, anyptr(target), vcbu(compareFunc), &matchIndex)
	return
}
func (array *GArray) SetClearFunc(clearFunc func(data UPtr)) {
	g_array_set_clear_func.Fn()(array, vcbu(clearFunc))
}

func NewGPtrArray() *GPtrArray { return g_ptr_array_new.Fn()() }
func NewGPtrArrayWithFreeFunc(elementFreeFunc func(data UPtr)) *GPtrArray {
	return g_ptr_array_new_with_free_func.Fn()(vcbu(elementFreeFunc))
}
func NewGPtrArrayTake(data UPtr, length uint64, elementFreeFunc func(data UPtr)) *GPtrArray {
	return g_ptr_array_new_take.Fn()(data, length, vcbu(elementFreeFunc))
}
func NewGPtrArrayFromArray(data UPtr, length uint64,
	copyFunc func(src, copyFuncUserData UPtr) UPtr,
	copyFuncUserData interface{},
	elementFreeFunc func(UPtr) UPtr) *GPtrArray {
	return g_ptr_array_new_from_array.Fn()(data, length, vcbu(copyFunc), anyptr(copyFuncUserData), vcbu(elementFreeFunc))
}
func (a *GPtrArray) Steal() (data UPtr, length uint64) {
	data = g_ptr_array_steal.Fn()(a, &length)
	return
}
func (a *GPtrArray) Copy(copyFunc func(src, userData UPtr) UPtr, userData interface{}) *GPtrArray {
	return g_ptr_array_copy.Fn()(a, vcbu(copyFunc), anyptr(userData))
}
func NewGPtrArraySized(reservedSize uint32) *GPtrArray {
	return g_ptr_array_sized_new.Fn()(reservedSize)
}
func NewGPtrArrayFull(reservedSize uint32, elementFreeFunc func(data UPtr)) *GPtrArray {
	return g_ptr_array_new_full.Fn()(reservedSize, vcbu(elementFreeFunc))
}
func NewGPtrArrayNullTerminated(reservedSize uint32, elementFreeFunc func(data UPtr), nullTerminated bool) *GPtrArray {
	return g_ptr_array_new_null_terminated.Fn()(reservedSize, vcbu(elementFreeFunc), nullTerminated)
}
func NewGPtrArrayTakeNullTerminated(data UPtr, elementFreeFunc func(data UPtr)) *GPtrArray {
	return g_ptr_array_new_take_null_terminated.Fn()(data, vcbu(elementFreeFunc))
}
func NewGPtrArrayFromNullTerminatedArray(data UPtr,
	copyFunc func(src, copyFuncUserData UPtr) UPtr,
	copyFuncUserData interface{},
	elementFreeFunc func(UPtr) UPtr) *GPtrArray {
	return g_ptr_array_new_from_null_terminated_array.Fn()(data, vcbu(copyFunc), anyptr(copyFuncUserData), vcbu(elementFreeFunc))
}
func (a *GPtrArray) Free(freeSegment bool) UPtr { return g_ptr_array_free.Fn()(a, freeSegment) }
func (a *GPtrArray) Ref() *GPtrArray            { return g_ptr_array_ref.Fn()(a) }
func (a *GPtrArray) Unref()                     { g_ptr_array_unref.Fn()(a) }
func (a *GPtrArray) SetFreeFunc(elementFreeFunc func(data UPtr)) {
	g_ptr_array_set_free_func.Fn()(a, vcbu(elementFreeFunc))
}
func (a *GPtrArray) SetSize(length int32)          { g_ptr_array_set_size.Fn()(a, length) }
func (a *GPtrArray) RemoveIndex(index uint32) UPtr { return g_ptr_array_remove_index.Fn()(a, index) }
func (a *GPtrArray) RemoveIndexFast(index uint32) UPtr {
	return g_ptr_array_remove_index_fast.Fn()(a, index)
}
func (a *GPtrArray) StealIndex(index uint32) UPtr { return g_ptr_array_steal_index.Fn()(a, index) }
func (a *GPtrArray) StealIndexFast(index uint32) UPtr {
	return g_ptr_array_steal_index_fast.Fn()(a, index)
}
func (a *GPtrArray) Remove(data UPtr) bool     { return g_ptr_array_remove.Fn()(a, data) }
func (a *GPtrArray) RemoveFast(data UPtr) bool { return g_ptr_array_remove_fast.Fn()(a, data) }
func (a *GPtrArray) RemoveRange(index, length uint32) *GPtrArray {
	return g_ptr_array_remove_range.Fn()(a, index, length)
}
func (a *GPtrArray) Add(data UPtr) { g_ptr_array_add.Fn()(a, data) }
func (a *GPtrArray) Extend(array *GPtrArray, copyFunc func(src, userData UPtr) UPtr, userData UPtr) {
	g_ptr_array_extend.Fn()(a, array, vcbu(copyFunc), userData)
}
func (a *GPtrArray) ExtendAndSteal(array *GPtrArray) { g_ptr_array_extend_and_steal.Fn()(a, array) }
func (a *GPtrArray) Insert(index int32, data UPtr)   { g_ptr_array_insert.Fn()(a, index, data) }
func (a *GPtrArray) Sort(compareFunc func(a, b UPtr) int32) {
	g_ptr_array_sort.Fn()(a, vcbu(compareFunc))
}
func (a *GPtrArray) SortWithData(compareFunc func(a, b, userData UPtr) int32, userData UPtr) {
	g_ptr_array_sort_with_data.Fn()(a, vcbu(compareFunc), userData)
}
func (a *GPtrArray) SortValues(compareFunc func(a, b UPtr) int32) {
	g_ptr_array_sort_values.Fn()(a, vcbu(compareFunc))
}
func (a *GPtrArray) SortValuesWithData(compareFunc func(a, b, userData UPtr) int32, userData UPtr) {
	g_ptr_array_sort_values_with_data.Fn()(a, vcbu(compareFunc), userData)
}
func (a *GPtrArray) Foreach(fn func(data, userData UPtr), userData interface{}) {
	g_ptr_array_foreach.Fn()(a, vcbu(fn), anyptr(userData))
}
func (a *GPtrArray) Find(needle UPtr, index *uint32) bool {
	return g_ptr_array_find.Fn()(a, needle, index)
}
func (a *GPtrArray) FindWithEqualFunc(needle UPtr, equalFunc func(a, b UPtr) bool, index *uint32) bool {
	return g_ptr_array_find_with_equal_func.Fn()(a, needle, vcbu(equalFunc), index)
}
func (a *GPtrArray) IsNullTerminated() bool { return g_ptr_array_is_null_terminated.Fn()(a) }

func NewGByteArray() *GByteArray { return g_byte_array_new.Fn()() }
func NewGByteArrayTake(data UPtr, length uint64) *GByteArray {
	return g_byte_array_new_take.Fn()(data, length)
}
func (a *GByteArray) Steal() (data UPtr, length uint64) {
	data = g_byte_array_steal.Fn()(a, &length)
	return
}
func NewGByteArraySized(reservedSize uint32) *GByteArray {
	return g_byte_array_sized_new.Fn()(reservedSize)
}
func (a *GByteArray) Free(freeSegment bool) UPtr { return g_byte_array_free.Fn()(a, freeSegment) }
func (a *GByteArray) FreeToBytes() *GBytes       { return g_byte_array_free_to_bytes.Fn()(a) }
func (a *GByteArray) Ref() *GByteArray           { return g_byte_array_ref.Fn()(a) }
func (a *GByteArray) Unref()                     { g_byte_array_unref.Fn()(a) }
func (a *GByteArray) Append(data []byte) *GByteArray {
	return g_byte_array_append.Fn()(a, UPtr(&data[0]), uint32(len(data)))
}
func (a *GByteArray) Prepend(data []byte) *GByteArray {
	return g_byte_array_prepend.Fn()(a, UPtr(&data[0]), uint32(len(data)))
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
func (a *GByteArray) Sort(compareFunc func(a, b UPtr) int32) {
	g_byte_array_sort.Fn()(a, vcbu(compareFunc))
}
func (a *GByteArray) SortWithData(compareFunc func(a, b, userData UPtr) int32, userData UPtr) {
	g_byte_array_sort_with_data.Fn()(a, vcbu(compareFunc), userData)
}

// #endregion

// #region AsyncQueue

type GAsyncQueue struct{}

func NewGAsyncQueue() *GAsyncQueue { return g_async_queue_new.Fn()() }
func NewGAsyncQueueFull(itemFreeFunc func(data UPtr)) *GAsyncQueue {
	return g_async_queue_new_full.Fn()(vcbu(itemFreeFunc))
}
func (q *GAsyncQueue) Lock()                 { g_async_queue_lock.Fn()(q) }
func (q *GAsyncQueue) Unlock()               { g_async_queue_unlock.Fn()(q) }
func (q *GAsyncQueue) Ref() *GAsyncQueue     { return g_async_queue_ref.Fn()(q) }
func (q *GAsyncQueue) Unref()                { g_async_queue_unref.Fn()(q) }
func (q *GAsyncQueue) Push(data interface{}) { g_async_queue_push.Fn()(q, anyptr(data)) }
func (q *GAsyncQueue) PushUnlocked(data interface{}) {
	g_async_queue_push_unlocked.Fn()(q, anyptr(data))
}
func (q *GAsyncQueue) PushSorted(data interface{}, cmp func(a, b, userData UPtr) int32, userData interface{}) {
	g_async_queue_push_sorted.Fn()(q, anyptr(data), vcbu(cmp), anyptr(userData))
}
func (q *GAsyncQueue) PushSortedUnlocked(data interface{}, cmp func(a, b, userData UPtr) int32, userData interface{}) {
	g_async_queue_push_sorted_unlocked.Fn()(q, anyptr(data), vcbu(cmp), anyptr(userData))
}
func (q *GAsyncQueue) Pop() UPtr            { return g_async_queue_pop.Fn()(q) }
func (q *GAsyncQueue) PopUnlocked() UPtr    { return g_async_queue_pop_unlocked.Fn()(q) }
func (q *GAsyncQueue) TryPop() UPtr         { return g_async_queue_try_pop.Fn()(q) }
func (q *GAsyncQueue) TryPopUnlocked() UPtr { return g_async_queue_try_pop_unlocked.Fn()(q) }
func (q *GAsyncQueue) TimeoutPop(timeout uint64) UPtr {
	return g_async_queue_timeout_pop.Fn()(q, timeout)
}
func (q *GAsyncQueue) TimeoutPopUnlocked(timeout uint64) UPtr {
	return g_async_queue_timeout_pop_unlocked.Fn()(q, timeout)
}
func (q *GAsyncQueue) Length() int32         { return g_async_queue_length.Fn()(q) }
func (q *GAsyncQueue) LengthUnlocked() int32 { return g_async_queue_length_unlocked.Fn()(q) }
func (q *GAsyncQueue) Sort(cmp func(a, b, userData UPtr) int32, userData interface{}) {
	g_async_queue_sort.Fn()(q, vcbu(cmp), anyptr(userData))
}
func (q *GAsyncQueue) SortUnlocked(cmp func(a, b, userData UPtr) int32, userData interface{}) {
	g_async_queue_sort_unlocked.Fn()(q, vcbu(cmp), anyptr(userData))
}
func (q *GAsyncQueue) Remove(item interface{}) bool {
	return g_async_queue_remove.Fn()(q, anyptr(item))
}
func (q *GAsyncQueue) RemoveUnlocked(item interface{}) bool {
	return g_async_queue_remove_unlocked.Fn()(q, anyptr(item))
}
func (q *GAsyncQueue) PushFront(item interface{}) { g_async_queue_push_front.Fn()(q, anyptr(item)) }
func (q *GAsyncQueue) PushFrontUnlocked(item interface{}) {
	g_async_queue_push_front_unlocked.Fn()(q, anyptr(item))
}

// #endregion

// #region Bytes

type GBytes struct{}

func NewBytes(data []byte) *GBytes {
	if len(data) == 0 {
		return g_bytes_new.Fn()(nil, 0)
	}
	return g_bytes_new.Fn()(UPtr(&data[0]), uint(len(data)))
}

func NewBytesFromString(data string) *GBytes {
	n := len(data)
	if n == 0 {
		return g_bytes_new.Fn()(nil, 0)
	}
	ps := *(*[2]UPtr)(UPtr(&data))
	return g_bytes_new.Fn()(ps[0], uint(n))
}

func NewBytesStatic(data []byte) *GBytes {
	if len(data) == 0 {
		return g_bytes_new_static.Fn()(nil, 0)
	}
	return g_bytes_new_static.Fn()(UPtr(&data[0]), uint(len(data)))
}

func (b *GBytes) Compare(b2 *GBytes) int { return g_bytes_compare.Fn()(b, b2) }
func (b *GBytes) Equal(b2 *GBytes) bool  { return g_bytes_equal.Fn()(b, b2) }
func (b *GBytes) GetData() []byte {
	var size uint
	ptr := g_bytes_get_data.Fn()(b, &size)
	return *(*[]byte)(cc.Slice(ptr, int64(size)))
}
func (b *GBytes) GetRegion(offset, length uint) []byte {
	ptr := g_bytes_get_region.Fn()(b, offset, length)
	return *(*[]byte)(cc.Slice(ptr, int64(length)))
}

func (b *GBytes) GetSize() uint { return g_bytes_get_size.Fn()(b) }
func (b *GBytes) Hash() uint    { return g_bytes_hash.Fn()(b) }
func (b *GBytes) NewFromBytes(offset, length uint) *GBytes {
	return g_bytes_new_from_bytes.Fn()(b, offset, length)
}

func (b *GBytes) Ref() *GBytes              { return g_bytes_ref.Fn()(b) }
func (b *GBytes) Unref()                    { g_bytes_unref.Fn()(b) }
func (b *GBytes) UnrefToArray() *GByteArray { return g_bytes_unref_to_array.Fn()(b) }
func (b *GBytes) UnrefToData() []byte {
	var size uint
	ptr := g_bytes_unref_to_data.Fn()(b, &size)
	return *(*[]byte)(cc.Slice(ptr, int64(size)))
}

// #endregion

// #region DateTime

type GDateTime struct {
}

// #endregion

// #region Error

type GError struct {
	Domain  GQuark
	Code    int32
	Message cc.String
}

func NewError(domain GQuark, code int32, format string, a ...interface{}) *GError {
	return g_error_new.Fn()(domain, code, fmt.Sprintf(format, a...))
}
func ErrorDomainRegister(typeName string, privateSize uint,
	typeInit func(err *GError), typeCopy func(src, dst *GError), typeFree func(err *GError)) GQuark {
	init, copy, free := UPtr(nil), UPtr(nil), UPtr(nil)
	if typeInit != nil {
		init = vcbu(typeInit)
	}
	if typeCopy != nil {
		copy = vcbu(typeCopy)
	}
	if typeFree != nil {
		free = vcbu(typeFree)
	}
	return g_error_domain_register.Fn()(typeName, privateSize, init, copy, free)
}
func ErrorDomainRegisterStatic(typeName string, privateSize uint,
	typeInit func(err *GError), typeCopy func(src, dst *GError), typeFree func(err *GError)) GQuark {
	init, copy, free := UPtr(nil), UPtr(nil), UPtr(nil)
	if typeInit != nil {
		init = vcbu(typeInit)
	}
	if typeCopy != nil {
		copy = vcbu(typeCopy)
	}
	if typeFree != nil {
		free = vcbu(typeFree)
	}
	return g_error_domain_register_static.Fn()(typeName, privateSize, init, copy, free)
}

func (e *GError) Copy() *GError                     { return g_error_copy.Fn()(e) }
func (e *GError) Free()                             { g_error_free.Fn()(e) }
func (e *GError) Matches(domain GQuark, code int32) { g_error_matches.Fn()(e, domain, code) }
func (e *GError) Error() error {
	return errors.New(e.Message.String())
}

// #endregion

// #region List

type GList struct {
	Data UPtr
	Next *GList
	Prev *GList
}

func ListAlloc() *GList                            { return g_list_alloc.Fn()() }
func NewGList() *GList                             { return g_list_alloc.Fn()() }
func AllocList() *GList                            { return g_list_alloc.Fn()() }
func (l *GList) Free()                             { g_list_free.Fn()(l) }
func (l *GList) Free1()                            { g_list_free_1.Fn()(l) }
func (l *GList) FreeFull(freeFunc func(data UPtr)) { g_list_free_full.Fn()(l, vcbu(freeFunc)) }
func (l *GList) Append(data UPtr) *GList           { return g_list_append.Fn()(l, data) }
func (l *GList) Prepend(data UPtr) *GList          { return g_list_prepend.Fn()(l, data) }
func (l *GList) Insert(data UPtr, position int32) *GList {
	return g_list_insert.Fn()(l, data, position)
}
func (l *GList) InsertSorted(data UPtr, compare func(a, b UPtr) int32) *GList {
	return g_list_insert_sorted.Fn()(l, data, vcbu(compare))
}
func (l *GList) InsertSortedWithData(data UPtr, compare func(a, b, data UPtr) int32, userData UPtr) *GList {
	return g_list_insert_sorted_with_data.Fn()(l, data, vcbu(compare), userData)
}
func (l *GList) InsertBefore(sibling *GList, data UPtr) *GList {
	return g_list_insert_before.Fn()(l, sibling, data)
}
func (l *GList) InsertBeforeLink(sibling *GList, link *GList) *GList {
	return g_list_insert_before_link.Fn()(l, sibling, link)
}
func (l *GList) Concat(list2 *GList) *GList {
	return g_list_concat.Fn()(l, list2)
}
func (l *GList) Remove(data UPtr) *GList {
	return g_list_remove.Fn()(l, data)
}
func (l *GList) RemoveAll(data UPtr) *GList {
	return g_list_remove_all.Fn()(l, data)
}
func (l *GList) RemoveLink(link *GList) *GList {
	return g_list_remove_link.Fn()(l, link)
}
func (l *GList) DeleteLink(link *GList) *GList {
	return g_list_delete_link.Fn()(l, link)
}
func (l *GList) Reverse() *GList {
	return g_list_reverse.Fn()(l)
}

func (l *GList) Copy() *GList {
	return g_list_copy.Fn()(l)
}

func (l *GList) CopyDeep(copyFunc func(src, data UPtr) UPtr, userData UPtr) *GList {
	return g_list_copy_deep.Fn()(l, vcbu(copyFunc), userData)
}
func (l *GList) Nth(n uint32) *GList     { return g_list_nth.Fn()(l, n) }
func (l *GList) NthPrev(n uint32) *GList { return g_list_nth_prev.Fn()(l, n) }
func (l *GList) Find(data UPtr) *GList   { return g_list_find.Fn()(l, data) }
func (l *GList) FindCustom(data UPtr, compareFunc func(a, b UPtr) int32) *GList {
	return g_list_find_custom.Fn()(l, data, vcbu(compareFunc))
}
func (l *GList) Position(link *GList) int32 {
	return g_list_position.Fn()(l, link)
}
func (l *GList) Index(data UPtr) int32 {
	return g_list_index.Fn()(l, data)
}
func (l *GList) Last() *GList {
	return g_list_last.Fn()(l)
}
func (l *GList) First() *GList {
	return g_list_first.Fn()(l)
}
func (l *GList) Length() uint32 {
	return g_list_length.Fn()(l)
}
func (l *GList) ForEach(fn func(data, userData UPtr), userData UPtr) {
	g_list_foreach.Fn()(l, vcbu(fn), userData)
}

func (l *GList) Sort(compareFunc func(a, b UPtr) int32) *GList {
	return g_list_sort.Fn()(l, vcbu(compareFunc))
}

func (l *GList) SortWithData(compareFunc func(a, b, userData UPtr) int32, userData UPtr) *GList {
	return g_list_sort_with_data.Fn()(l, vcbu(compareFunc), userData)
}
func (l *GList) NthData(n uint32) UPtr { return g_list_nth_data.Fn()(l, n) }
func ClearList(list **GList, destroy func(data UPtr)) {
	g_clear_list.Fn()(list, vcbu(destroy))
}
func GListList[T any](l *GList, close func(*GList)) []T {
	type pinner struct {
		list  *GList
		slice *[]T
	}
	num := l.Length()
	if num == 0 {
		return nil
	}

	pin := new(pinner)
	pin.list = l
	slice := make([]T, num)
	pin.slice = &slice
	runtime.SetFinalizer(pin, func(p *pinner) {
		if close != nil {
			close(p.list)
			return
		}
		p.list.Free()

		_ = *p.slice
	})

	var t T
	switch any(t).(type) {
	case string:
		for i := uint32(0); i < num; i++ {
			a := ((cc.String)(l.Nth(i).Data)).String()
			(*pin.slice)[i] = any(a).(T)
		}
	default:
		for i := uint32(0); i < num; i++ {
			(*pin.slice)[i] = *(*T)(UPtr(&l.Nth(i).Data))
		}
	}

	return *pin.slice
}

// #endregion

// #region Main

type GMainContext struct{}
type GMainLoop struct{}

type GSource struct {
	CallBackData  UPtr
	CallbackFuncs *GSourceCallbackFuncs
	SourceFuncs   *GSourceFuncs
	RefCount      uint32
	Context       *GMainContext

	Priority int32
	Flags    uint32
	SourceID uint32

	PollFds *GSList
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
func NewGMainContextWithFlags(flags uint32) *GMainContext {
	return g_main_context_new_with_flags.Fn()(flags)
}
func (c *GMainContext) Ref() *GMainContext { return g_main_context_ref.Fn()(c) }
func (c *GMainContext) Unref()             { g_main_context_unref.Fn()(c) }
func DefaultMainContext() *GMainContext    { return g_main_context_default.Fn()() }
func (c *GMainContext) Iteration(mayBlock bool) bool {
	return g_main_context_iteration.Fn()(c, mayBlock)
}
func (c *GMainContext) Pending() bool { return g_main_context_pending.Fn()(c) }
func (c *GMainContext) FindSourceByID(sourceID uint32) *GSource {
	return g_main_context_find_source_by_id.Fn()(c, sourceID)
}
func (c *GMainContext) FindSourceByUserData(userData interface{}) *GSource {
	return g_main_context_find_source_by_user_data.Fn()(c, anyptr(userData))
}
func (c *GMainContext) FindSourceByFuncsUserData(funcs *GSourceFuncs, userData interface{}) *GSource {
	return g_main_context_find_source_by_funcs_user_data.Fn()(c, funcs, anyptr(userData))
}
func (c *GMainContext) Wakeup()                      { g_main_context_wakeup.Fn()(c) }
func (c *GMainContext) Acquire() bool                { return g_main_context_acquire.Fn()(c) }
func (c *GMainContext) Release()                     { g_main_context_release.Fn()(c) }
func (c *GMainContext) IsOwner() bool                { return g_main_context_is_owner.Fn()(c) }
func (c *GMainContext) Wait(cond, mutex UPtr) bool   { return g_main_context_wait.Fn()(c, cond, mutex) }
func (c *GMainContext) Prepare(priority *int32) bool { return g_main_context_prepare.Fn()(c, priority) }
func (c *GMainContext) Query(maxPriority int32, timeout *int32, fds *GPollFD, nFds int32) int32 {
	return g_main_context_query.Fn()(c, maxPriority, timeout, fds, nFds)
}
func (c *GMainContext) Check(maxPriority int32, fds *GPollFD, nFds int32) bool {
	return g_main_context_check.Fn()(c, maxPriority, fds, nFds)
}
func (c *GMainContext) Dispatch() { g_main_context_dispatch.Fn()(c) }
func (c *GMainContext) SetPollFunc(fn func(ufds *GPollFD, nfsd uint32, timeout int32) int32) {
	g_main_context_set_poll_func.Fn()(c, vcbu(fn))
}
func (c *GMainContext) GetPollFunc() UPtr { return g_main_context_get_poll_func.Fn()(c) }
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
func (c *GMainContext) InvokeFull(priority int32, fn func(_ UPtr) (Continue bool),
	data interface{}, notify func(_ UPtr)) {
	g_main_context_invoke_full.Fn()(c, priority, vcbu(fn), anyptr(data), vcbu(notify))
}
func (c *GMainContext) Invoke(fn func(_ UPtr) (Continue bool)) {
	g_main_context_invoke.Fn()(c, vcbu(fn), nil)
}

func NewGMainLoop(context *GMainContext, isRunning bool) *GMainLoop {
	return g_main_loop_new.Fn()(context, isRunning)
}
func (l *GMainLoop) Run()                      { g_main_loop_run.Fn()(l) }
func (l *GMainLoop) Quit()                     { g_main_loop_quit.Fn()(l) }
func (l *GMainLoop) Ref() *GMainLoop           { return g_main_loop_ref.Fn()(l) }
func (l *GMainLoop) Unref()                    { g_main_loop_unref.Fn()(l) }
func (l *GMainLoop) IsRunning() bool           { return g_main_loop_is_running.Fn()(l) }
func (l *GMainLoop) GetContext() *GMainContext { return g_main_loop_get_context.Fn()(l) }

func NewGSource(sourceFuncs *GSourceFuncs, structSize uint32) *GSource {
	return g_source_new.Fn()(sourceFuncs, structSize)
}
func (s *GSource) SetDisposeFunction(dispose func(s *GSource)) {
	g_source_set_dispose_function.Fn()(s, vcbu(dispose))
}
func (s *GSource) Ref() *GSource                       { return g_source_ref.Fn()(s) }
func (s *GSource) Unref()                              { g_source_unref.Fn()(s) }
func (s *GSource) Attach(context *GMainContext) uint32 { return g_source_attach.Fn()(s, context) }
func (s *GSource) Destroy()                            { g_source_destroy.Fn()(s) }
func (s *GSource) SetPriority(priority int32)          { g_source_set_priority.Fn()(s, priority) }
func (s *GSource) GetPriority() int32                  { return g_source_get_priority.Fn()(s) }
func (s *GSource) SetCanRecurse(canRecurse bool)       { g_source_set_can_recurse.Fn()(s, canRecurse) }
func (s *GSource) GetCanRecurse() bool                 { return g_source_get_can_recurse.Fn()(s) }
func (s *GSource) GetID() uint32                       { return g_source_get_id.Fn()(s) }
func (s *GSource) GetContext() *GMainContext           { return g_source_get_context.Fn()(s) }
func (s *GSource) SetCallback(fn func(_ UPtr) (Continue bool)) {
	g_source_set_callback.Fn()(s, vcbu(fn), nil, nil)
}
func (s *GSource) SetFuncs(funcs *GSourceFuncs) { g_source_set_funcs.Fn()(s, funcs) }
func (s *GSource) IsDestroyed() bool            { return g_source_is_destroyed.Fn()(s) }
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

func SourceRemove(tag uint32) bool { return g_source_remove.Fn()(tag) }
func SourceRemoveByUserData(userData interface{}) bool {
	return g_source_remove_by_user_data.Fn()(anyptr(userData))
}
func SourceRemoveByFuncsUserData(funcs *GSourceFuncs, userData interface{}) bool {
	return g_source_remove_by_funcs_user_data.Fn()(funcs, anyptr(userData))
}

func TimeoutAddFull(priority int32, ms uint32, fn func(_ UPtr) (Continue bool), data interface{}, notify func(_ UPtr)) uint32 {
	return g_timeout_add_full.Fn()(priority, ms, vcbu(fn), anyptr(data), vcbu(notify))
}
func TimeoutAdd(ms uint32, fn func() (Continue bool)) uint32 {
	var fnp UPtr
	fnp = vcbu(func(_ UPtr) (Continue bool) {
		b := fn()
		if !b {
			cc.CbkCloseLate(fnp)
		}
		return b
	})
	return g_timeout_add.Fn()(ms, fnp, nil)
}
func TimeoutAddOnce(ms uint32, fn func()) uint32 {
	return g_timeout_add_once.Fn()(ms, vcbu(fn), nil)
}
func TimeoutAddSecondsFull(priority int32, sec uint32, fn func(_ UPtr) (Continue bool),
	data interface{}, notify func(_ UPtr)) uint32 {
	return g_timeout_add_seconds_full.Fn()(priority, sec, vcbu(fn), anyptr(data), vcbu(notify))
}
func TimeoutAddSeconds(sec uint32, fn func(_ UPtr) (Continue bool)) uint32 {
	return g_timeout_add_seconds.Fn()(sec, vcbu(fn), nil)
}
func TimeoutAddSecondsOnce(sec uint32, fn func()) uint32 {
	return g_timeout_add_seconds_once.Fn()(sec, vcbu(fn), nil)
}
func ChildWatchAddFull(priority int32, pid GPid, fn func(pid GPid, waitStatus int32),
	data interface{}, notify func(_ UPtr)) uint32 {
	return g_child_watch_add_full.Fn()(priority, pid, vcbu(fn), anyptr(data), vcbu(notify))
}
func ChildWatchAdd(pid GPid, fn func(pid GPid, waitStatus int32)) uint32 {
	return g_child_watch_add.Fn()(pid, vcbu(fn), nil)
}
func IdleAdd(fn func() (Continue bool)) uint32 {
	var fnp UPtr
	fnp = vcbu(func(_ UPtr) (Continue bool) {
		b := fn()
		if !b {
			cc.CbkCloseLate(fnp)
		}
		return b
	})
	return g_idle_add.Fn()(fnp, nil)
}
func IdleAddFull(priority int32, fn func(_ UPtr) (Continue bool),
	data interface{}, notify func(_ UPtr)) uint32 {
	return g_idle_add_full.Fn()(priority, vcbu(fn), anyptr(data), vcbu(notify))
}
func IdleAddOnce(fn func()) uint32 {
	var fnp UPtr
	fnp = vcbu(func() {
		fn()
		cc.CbkCloseLate(fnp)
	})
	return g_idle_add_once.Fn()(fnp, nil)
}
func IdleRemoveByData(data interface{}) bool { return g_idle_remove_by_data.Fn()(anyptr(data)) }

// #endregion

// #region PollFD

type GPollFD struct{}

// #endregion

// #region Quark

type GQuark uint32

func GQuarkTryString(str string) GQuark        { return g_quark_try_string.Fn()(str) }
func GQuarkFromStaticString(str string) GQuark { return g_quark_from_static_string.Fn()(str) }
func GQuarkFromString(str string) GQuark       { return g_quark_from_string.Fn()(str) }
func (quark GQuark) ToString() string          { return g_quark_to_string.Fn()(quark) }
func (quark GQuark) String() string            { return g_quark_to_string.Fn()(quark) }

// #endregion

// #region SList

type GSList struct {
	Data UPtr
	Next *GSList
}

func NewGSList() *GSList                            { return g_slist_alloc.Fn()() }
func (l *GSList) Free()                             { g_slist_free.Fn()(l) }
func (l *GSList) Free1()                            { g_slist_free_1.Fn()(l) }
func (l *GSList) FreeFull(freeFunc func(data UPtr)) { g_slist_free_full.Fn()(l, vcbu(freeFunc)) }
func (l *GSList) Append(data interface{}) *GSList   { return g_slist_append.Fn()(l, anyptr(data)) }
func (l *GSList) Prepend(data interface{}) *GSList  { return g_slist_prepend.Fn()(l, anyptr(data)) }
func (l *GSList) Insert(data interface{}, position int32) *GSList {
	return g_slist_insert.Fn()(l, anyptr(data), position)
}
func (l *GSList) InsertSorted(data interface{}, compareFunc func(a, b UPtr) int32) *GSList {
	return g_slist_insert_sorted.Fn()(l, anyptr(data), vcbu(compareFunc))
}
func (l *GSList) InsertSortedWithData(data interface{}, compareDataFunc func(a, b, userData UPtr) int32, userData interface{}) *GSList {
	return g_slist_insert_sorted_with_data.Fn()(l, anyptr(data), vcbu(compareDataFunc), anyptr(userData))
}
func (l *GSList) InsertBefore(sibling *GSList, data interface{}) *GSList {
	return g_slist_insert_before.Fn()(l, sibling, anyptr(data))
}
func (l *GSList) Concat(list2 *GSList) *GSList       { return g_slist_concat.Fn()(l, list2) }
func (l *GSList) Remove(data interface{}) *GSList    { return g_slist_remove.Fn()(l, anyptr(data)) }
func (l *GSList) RemoveAll(data interface{}) *GSList { return g_slist_remove_all.Fn()(l, anyptr(data)) }
func (l *GSList) RemoveLink(link *GSList) *GSList    { return g_slist_remove_link.Fn()(l, link) }
func (l *GSList) DeleteLink(link *GSList) *GSList    { return g_slist_delete_link.Fn()(l, link) }
func (l *GSList) Reverse() *GSList                   { return g_slist_reverse.Fn()(l) }
func (l *GSList) Copy() *GSList                      { return g_slist_copy.Fn()(l) }
func (l *GSList) CopyDeep(copyFunc func(src, userData UPtr) UPtr, userData interface{}) *GSList {
	return g_slist_copy_deep.Fn()(l, vcbu(copyFunc), anyptr(userData))
}
func (l *GSList) Nth(n uint32) *GSList          { return g_slist_nth.Fn()(l, n) }
func (l *GSList) Find(data interface{}) *GSList { return g_slist_find.Fn()(l, anyptr(data)) }
func (l *GSList) FindCustom(data interface{}, compareFunc func(a, b UPtr) int32) *GSList {
	return g_slist_find_custom.Fn()(l, anyptr(data), vcbu(compareFunc))
}
func (l *GSList) Position(link *GSList) int32  { return g_slist_position.Fn()(l, link) }
func (l *GSList) Index(data interface{}) int32 { return g_slist_index.Fn()(l, anyptr(data)) }
func (l *GSList) Last() *GSList                { return g_slist_last.Fn()(l) }
func (l *GSList) Length() uint32               { return g_slist_length.Fn()(l) }
func (l *GSList) ForEach(fn func(data, userData UPtr), userData interface{}) {
	g_slist_foreach.Fn()(l, vcbu(fn), anyptr(userData))
}
func (l *GSList) Sort(compareFunc func(a, b UPtr) int32) *GSList {
	return g_slist_sort.Fn()(l, vcbu(compareFunc))
}
func (l *GSList) SortWithData(compareDataFunc func(a, b, userData UPtr) int32, userData interface{}) *GSList {
	return g_slist_sort_with_data.Fn()(l, vcbu(compareDataFunc), anyptr(userData))
}
func (l *GSList) NthData(n uint32) UPtr { return g_slist_nth_data.Fn()(l, n) }
func (l *GSList) Clear(destroy func(data UPtr)) {
	g_clear_slist.Fn()(&l, vcbu(destroy))
}
func GSListList[T any](l *GSList, close func(*GSList)) (list []T) {
	type pinner struct {
		list  *GSList
		slice *[]T
	}
	num := l.Length()
	if num == 0 {
		return nil
	}

	pin := new(pinner)
	pin.list = l
	slice := make([]T, num)
	pin.slice = &slice

	runtime.SetFinalizer(pin, func(p *pinner) {
		if close != nil {
			close(p.list)
			return
		}
		p.list.Free()

		_ = *p.slice
	})

	var t T
	switch any(t).(type) {
	case string:
		for i := uint32(0); i < num; i++ {
			a := ((cc.String)(l.Nth(i).Data)).String()
			(*pin.slice)[i] = any(a).(T)
		}
	default:
		for i := uint32(0); i < num; i++ {
			(*pin.slice)[i] = *(*T)(l.Nth(i).Data)
		}
	}

	return *pin.slice
}

// #endregion

// #region String

type GString struct {
	ptr UPtr
	len uint64
	alc uint64
}

func (s *GString) String() string { return *(*string)(UPtr(s)) }
func (s *GString) Ptr() UPtr      { return s.ptr }
func (s *GString) Len() uint64    { return s.len }
func (s *GString) Alc() uint64    { return s.alc }

func NewGString(init string) *GString             { return g_string_new.Fn()(init) }
func NewGStringLen(init string, len int) *GString { return g_string_new_len.Fn()(init, len) }
func NewGStringTake(ptr UPtr) *GString            { return g_string_new_take.Fn()(ptr) }
func NewGstringSized(dflSize uint) *GString       { return g_string_sized_new.Fn()(dflSize) }

func (s *GString) Append(val string) *GString { return g_string_append.Fn()(s, val) }
func (s *GString) AppendC(c byte) *GString    { return g_string_append_c.Fn()(s, c) }
func (s *GString) AppendLen(val string, len int) *GString {
	return g_string_append_len.Fn()(s, val, len)
}
func (s *GString) AppendPrintf(format string, args ...interface{}) {
	g_string_append_printf.Fn()(s, fmt.Sprintf(format, args...))
}
func (s *GString) AppendUnichar(c uint32) *GString { return g_string_append_unichar.Fn()(s, c) }
func (s *GString) AppendUriEscaped(unescaped, reservedCharsAllowed string, allowUtf8 bool) *GString {
	return g_string_append_uri_escaped.Fn()(s, unescaped, reservedCharsAllowed, allowUtf8)
}
func (s *GString) AsciiDown() *GString                 { return g_string_ascii_down.Fn()(s) }
func (s *GString) AsciiUp() *GString                   { return g_string_ascii_up.Fn()(s) }
func (s *GString) Assign(val string) *GString          { return g_string_assign.Fn()(s, val) }
func (s *GString) Down() *GString                      { return g_string_down.Fn()(s) }
func (s *GString) Equal(s2 *GString) bool              { return g_string_equal.Fn()(s, s2) }
func (s *GString) Erase(pos, len int) *GString         { return g_string_erase.Fn()(s, pos, len) }
func (s *GString) Free(freeSegment bool) UPtr          { return g_string_free.Fn()(s, freeSegment) }
func (s *GString) FreeAndSteal() UPtr                  { return g_string_free_and_steal.Fn()(s) }
func (s *GString) FreeToBytes() *GBytes                { return g_string_free_to_bytes.Fn()(s) }
func (s *GString) Hash() uint                          { return g_string_hash.Fn()(s) }
func (s *GString) Insert(pos int, val string) *GString { return g_string_insert.Fn()(s, pos, val) }
func (s *GString) InsertC(pos int, c byte) *GString    { return g_string_insert_c.Fn()(s, pos, c) }
func (s *GString) InsertLen(pos int, val string, len int) *GString {
	return g_string_insert_len.Fn()(s, pos, val, len)
}
func (s *GString) InsertUnichar(pos int, c uint32) *GString {
	return g_string_insert_unichar.Fn()(s, pos, c)
}
func (s *GString) Overwrite(pos uint, val string) *GString {
	return g_string_overwrite.Fn()(s, pos, val)
}
func (s *GString) OverwriteLen(pos uint, val string, len int) *GString {
	return g_string_overwrite_len.Fn()(s, pos, val, len)
}
func (s *GString) Prepend(val string) *GString { return g_string_prepend.Fn()(s, val) }
func (s *GString) PrependC(c byte) *GString    { return g_string_prepend_c.Fn()(s, c) }
func (s *GString) PrependLen(val string, len int) *GString {
	return g_string_prepend_len.Fn()(s, val, len)
}
func (s *GString) PrependUnichar(c uint32) *GString { return g_string_prepend_unichar.Fn()(s, c) }
func (s *GString) Printf(format string, args ...interface{}) {
	g_string_printf.Fn()(s, fmt.Sprintf(format, args...))
}
func (s *GString) Replace(find, replace string, limit uint32) uint32 {
	return g_string_replace.Fn()(s, find, replace, limit)
}
func (s *GString) SetSize(len uint) *GString  { return g_string_set_size.Fn()(s, len) }
func (s *GString) Truncate(len uint) *GString { return g_string_truncate.Fn()(s, len) }
func (s *GString) Up() *GString               { return g_string_up.Fn()(s) }

// #endregion

// #region Variant

type GVariantIter struct{}

type GVariantType struct{}

type nameSpaceGVariantType int

const (
	GVariantTypes nameSpaceGVariantType = 0
)

func (t nameSpaceGVariantType) New(typeString string) *GVariantType {
	return g_variant_type_new.Fn()(typeString)
}
func (t nameSpaceGVariantType) NewByClass(cls GVariantClass) *GVariantType {
	return g_variant_type_new.Fn()(string(byte(cls)))
}
func (t nameSpaceGVariantType) NewArray(elementType *GVariantType) *GVariantType {
	return g_variant_type_new_array.Fn()(elementType)
}
func (t nameSpaceGVariantType) NewDictEntry(keyType, valueType *GVariantType) *GVariantType {
	return g_variant_type_new_dict_entry.Fn()(keyType, valueType)
}
func (t nameSpaceGVariantType) NewMaybe(elementType *GVariantType) *GVariantType {
	return g_variant_type_new_maybe.Fn()(elementType)
}
func (t nameSpaceGVariantType) NewTuple(items []*GVariantType) *GVariantType {
	return g_variant_type_new_tuple.Fn()(&items[0], int32(len(items)))
}

func (t *GVariantType) Copy() *GVariantType            { return g_variant_type_copy.Fn()(t) }
func (t *GVariantType) DupString() string              { return g_variant_type_dup_string.Fn()(t) }
func (t *GVariantType) Element() *GVariantType         { return g_variant_type_element.Fn()(t) }
func (t *GVariantType) Equal(other *GVariantType) bool { return g_variant_type_equal.Fn()(t, other) }
func (t *GVariantType) First() *GVariantType           { return g_variant_type_first.Fn()(t) }
func (t *GVariantType) Free()                          { g_variant_type_free.Fn()(t) }
func (t *GVariantType) GetStringLength() uint          { return g_variant_type_get_string_length.Fn()(t) }
func (t *GVariantType) Hash() uint32                   { return g_variant_type_hash.Fn()(UPtr(t)) }
func (t *GVariantType) IsArray() bool                  { return g_variant_type_is_array.Fn()(t) }
func (t *GVariantType) IsBasic() bool                  { return g_variant_type_is_basic.Fn()(t) }
func (t *GVariantType) IsContainer() bool              { return g_variant_type_is_container.Fn()(t) }
func (t *GVariantType) IsDefinite() bool               { return g_variant_type_is_definite.Fn()(t) }
func (t *GVariantType) IsDictEntry() bool              { return g_variant_type_is_dict_entry.Fn()(t) }
func (t *GVariantType) IsMaybe() bool                  { return g_variant_type_is_maybe.Fn()(t) }
func (t *GVariantType) IsSubtypeOf(supertype *GVariantType) bool {
	return g_variant_type_is_subtype_of.Fn()(t, supertype)
}
func (t *GVariantType) IsTuple() bool        { return g_variant_type_is_tuple.Fn()(t) }
func (t *GVariantType) IsVariant() bool      { return g_variant_type_is_variant.Fn()(t) }
func (t *GVariantType) Key() *GVariantType   { return g_variant_type_key.Fn()(t) }
func (t *GVariantType) NItems() uint         { return g_variant_type_n_items.Fn()(t) }
func (t *GVariantType) Next() *GVariantType  { return g_variant_type_next.Fn()(t) }
func (t *GVariantType) PeekString() string   { return g_variant_type_peek_string.Fn()(t) }
func (t *GVariantType) Value() *GVariantType { return g_variant_type_value.Fn()(t) }

type GVariant struct{}

type nameSpaceGVariant int

const (
	GVariants nameSpaceGVariant = 0
)

func (ns nameSpaceGVariant) New(format string, args ...interface{}) *GVariant {
	return g_variant_new.Fn()(fmt.Sprintf(format, args...))
}

func (ns nameSpaceGVariant) NewArray(elementType *GVariantType, children []*GVariant) *GVariant {
	return g_variant_new_array.Fn()(elementType, &children[0], int64(len(children)))
}

func (ns nameSpaceGVariant) NewBoolean(value bool) *GVariant {
	return g_variant_new_boolean.Fn()(value)
}

func (ns nameSpaceGVariant) NewByte(value byte) *GVariant {
	return g_variant_new_byte.Fn()(value)
}

func (ns nameSpaceGVariant) NewByteString(str string) *GVariant {
	return g_variant_new_bytestring.Fn()(str)
}

func (ns nameSpaceGVariant) NewByteStringArray(strv []string) *GVariant {
	p := cc.NewStrings(strv)
	defer p.Free()
	return g_variant_new_bytestring_array.Fn()(p, int64(len(strv)))
}

func (ns nameSpaceGVariant) NewDictEntry(key, value *GVariant) *GVariant {
	return g_variant_new_dict_entry.Fn()(key, value)
}

func (ns nameSpaceGVariant) NewDouble(value float64) *GVariant {
	return g_variant_new_double.Fn()(value)
}

func (ns nameSpaceGVariant) NewFixedArray(elementType *GVariantType, elementsPtr interface{}, nElements, elementSize uint) *GVariant {
	return g_variant_new_fixed_array.Fn()(elementType, anyptr(elementsPtr), nElements, elementSize)
}

func (ns nameSpaceGVariant) NewFromBytes(type_ *GVariantType, bytes *GBytes, trusted bool) *GVariant {
	return g_variant_new_from_bytes.Fn()(type_, bytes, trusted)
}

// func (ns nameSpaceGVariant) NewFromData(type_ *GVariantType, data uptr, size int64, trusted bool, notify GDestroyNotify, userData uptr) *GVariant {
// 	return g_variant_new_from_data(type_, data, size, trusted, notify, userData)
// }

func (ns nameSpaceGVariant) NewHandle(value int32) *GVariant { return g_variant_new_handle.Fn()(value) }
func (ns nameSpaceGVariant) NewInt16(value int16) *GVariant  { return g_variant_new_int16.Fn()(value) }
func (ns nameSpaceGVariant) NewInt32(value int32) *GVariant  { return g_variant_new_int32.Fn()(value) }
func (ns nameSpaceGVariant) NewInt64(value int64) *GVariant  { return g_variant_new_int64.Fn()(value) }
func (ns nameSpaceGVariant) NewMaybe(childType *GVariantType, child *GVariant) *GVariant {
	return g_variant_new_maybe.Fn()(childType, child)
}
func (ns nameSpaceGVariant) NewObjectPath(objectPath string) *GVariant {
	return g_variant_new_object_path.Fn()(objectPath)
}

func (ns nameSpaceGVariant) NewObjv(strv []string) *GVariant {
	p := cc.NewStrings(strv)
	defer p.Free()
	return g_variant_new_objv.Fn()(p, int64(len(strv)))
}

func (ns nameSpaceGVariant) NewParsed(format string, args ...interface{}) *GVariant {
	return g_variant_new_parsed.Fn()(fmt.Sprintf(format, args...))
}

func (ns nameSpaceGVariant) NewPrintf(format string, args ...interface{}) *GVariant {
	return g_variant_new_printf.Fn()(fmt.Sprintf(format, args...))
}

func (ns nameSpaceGVariant) NewSignature(signature string) *GVariant {
	return g_variant_new_signature.Fn()(signature)
}

func (ns nameSpaceGVariant) NewString(str string) *GVariant {
	return g_variant_new_string.Fn()(str)
}

func (ns nameSpaceGVariant) NewStrv(strv []string) *GVariant {
	p := cc.NewStrings(strv)
	defer p.Free()
	return g_variant_new_strv.Fn()(p, int64(len(strv)))
}

func (ns nameSpaceGVariant) NewTuple(children []*GVariant) *GVariant {
	return g_variant_new_tuple.Fn()(&children[0], uint(len(children)))
}

func (ns nameSpaceGVariant) NewUint16(value uint16) *GVariant {
	return g_variant_new_uint16.Fn()(value)
}
func (ns nameSpaceGVariant) NewUint32(value uint32) *GVariant {
	return g_variant_new_uint32.Fn()(value)
}
func (ns nameSpaceGVariant) NewUint64(value uint64) *GVariant {
	return g_variant_new_uint64.Fn()(value)
}
func (ns nameSpaceGVariant) NewVariant(value *GVariant) *GVariant {
	return g_variant_new_variant.Fn()(value)
}

// GVariant 方法

func (v *GVariant) Byteswap() *GVariant { return g_variant_byteswap.Fn()(v) }
func (v *GVariant) CheckFormatString(format string, copyOnly bool) bool {
	return g_variant_check_format_string.Fn()(v, format, copyOnly)
}
func (v *GVariant) Classify() GVariantClass      { return g_variant_classify.Fn()(v) }
func (v *GVariant) Compare(other *GVariant) bool { return g_variant_compare.Fn()(v, other) != 0 }
func (v *GVariant) DupByteString() string {
	var length uint64
	p := g_variant_dup_bytestring.Fn()(v, &length)
	defer p.Free()
	return p.String(length)
}
func (v *GVariant) DupByteStringArray() []string {
	var length uint64
	cArray := g_variant_dup_bytestring_array.Fn()(v, &length)
	defer cArray.Free()
	return cArray.Strings(length)
}
func (v *GVariant) DupObjv() []string {
	var length uint64
	cArray := g_variant_dup_objv.Fn()(v, &length)
	defer cArray.Free()
	return cArray.Strings(length)
}
func (v *GVariant) DupString() string {
	var length uint64
	str := g_variant_dup_string.Fn()(v, &length)
	defer str.Free()
	return str.String(length)
}
func (v *GVariant) DupStrv() []string {
	var length uint64
	cArray := g_variant_dup_strv.Fn()(v, &length)
	defer cArray.Free()
	return cArray.Strings(length)
}
func (v *GVariant) Equal(other *GVariant) bool         { return g_variant_equal.Fn()(v, other) }
func (v *GVariant) Get(format string, ptr interface{}) { g_variant_get.Fn()(v, format, anyptr(ptr)) }
func (v *GVariant) GetBoolean() bool                   { return g_variant_get_boolean.Fn()(v) }
func (v *GVariant) GetByte() byte                      { return g_variant_get_byte.Fn()(v) }

func (v *GVariant) GetByteString() string { return g_variant_get_bytestring.Fn()(v) }

func (v *GVariant) GetByteStringArray() []string {
	var length uint64
	cArray := g_variant_get_bytestring_array.Fn()(v, &length)
	defer cc.Ptr(cArray).Free()
	return cArray.Ref(length)
}
func (v *GVariant) GetChild(index uint, format string, ptr interface{}) {
	g_variant_get_child.Fn()(v, index, format, anyptr(ptr))
}
func (v *GVariant) GetChildValue(index uint) *GVariant {
	return g_variant_get_child_value.Fn()(v, index)
}
func (v *GVariant) GetData() UPtr           { return g_variant_get_data.Fn()(v) }
func (v *GVariant) GetDataAsBytes() *GBytes { return g_variant_get_data_as_bytes.Fn()(v) }
func (v *GVariant) GetDouble() float64      { return g_variant_get_double.Fn()(v) }
func (v *GVariant) GetFixedArray(elementSize uint) (firstPtr UPtr, nElements uint) {
	firstPtr = g_variant_get_fixed_array.Fn()(v, &nElements, elementSize)
	return
}
func (v *GVariant) GetHandle() int32         { return g_variant_get_handle.Fn()(v) }
func (v *GVariant) GetInt16() int16          { return g_variant_get_int16.Fn()(v) }
func (v *GVariant) GetInt32() int32          { return g_variant_get_int32.Fn()(v) }
func (v *GVariant) GetInt64() int64          { return g_variant_get_int64.Fn()(v) }
func (v *GVariant) GetMaybe() *GVariant      { return g_variant_get_maybe.Fn()(v) }
func (v *GVariant) GetNormalForm() *GVariant { return g_variant_get_normal_form.Fn()(v) }
func (v *GVariant) GetObjv() []string {
	var length uint64
	cArray := g_variant_get_objv.Fn()(v, &length)
	return cArray.Ref(length)
}
func (v *GVariant) GetSize() uint { return g_variant_get_size.Fn()(v) }
func (v *GVariant) GetString() string {
	var length uint64
	str := g_variant_get_string.Fn()(v, &length)
	return str.String(length)
}

func (v *GVariant) GetStrv() []string {
	var length uint64
	cArray := g_variant_get_strv.Fn()(v, &length)
	return cArray.Strings(length)
}

func (v *GVariant) GetType() *GVariantType { return g_variant_get_type.Fn()(v) }

func (v *GVariant) GetTypeString() string { return g_variant_get_type_string.Fn()(v) }

func (v *GVariant) GetUint16() uint16 { return g_variant_get_uint16.Fn()(v) }
func (v *GVariant) GetUint32() uint32 { return g_variant_get_uint32.Fn()(v) }
func (v *GVariant) GetUint64() uint64 { return g_variant_get_uint64.Fn()(v) }

func (v *GVariant) GetVariant() *GVariant { return g_variant_get_variant.Fn()(v) }

func (v *GVariant) Hash() uint { return g_variant_hash.Fn()(v) }

func (v *GVariant) IsContainer() bool  { return g_variant_is_container.Fn()(v) }
func (v *GVariant) IsFloating() bool   { return g_variant_is_floating.Fn()(v) }
func (v *GVariant) IsNormalForm() bool { return g_variant_is_normal_form.Fn()(v) }

func (v *GVariant) IsOfType(typ *GVariantType) bool {
	return g_variant_is_of_type.Fn()(v, typ)
}

func (v *GVariant) IterNew() *GVariantIter { return g_variant_iter_new.Fn()(v) }

func (v *GVariant) Lookup(key string, format string, args ...interface{}) bool {
	return g_variant_lookup.FnVa()(v, key, format, args...)
}

func (v *GVariant) LookupValue(key string, typ *GVariantType) *GVariant {
	return g_variant_lookup_value.Fn()(v, key, typ)
}

func (v *GVariant) NChildren() uint { return g_variant_n_children.Fn()(v) }

func (v *GVariant) Print(typeAnnotate bool) string { return g_variant_print.Fn()(v, typeAnnotate) }

func (v *GVariant) PrintString(str *GString, typeAnnotate bool) *GString {
	return g_variant_print_string.Fn()(v, str, typeAnnotate)
}

func (v *GVariant) Ref() *GVariant     { return g_variant_ref.Fn()(v) }
func (v *GVariant) RefSink() *GVariant { return g_variant_ref_sink.Fn()(v) }

func (v *GVariant) Store(data UPtr) { g_variant_store.Fn()(v, data) }

func (v *GVariant) TakeRef() *GVariant { return g_variant_take_ref.Fn()(v) }
func (v *GVariant) Unref()             { g_variant_unref.Fn()(v) }

// #endregion

func GetEnv(env string) string                      { return g_getenv.Fn()(env) }
func SetEnv(env, value string, overwrite bool) bool { return g_setenv.Fn()(env, value, overwrite) }

func GFree(ptr interface{}) { g_free.Fn()(anyptr(ptr)) }

type GOptionGroup struct{}
