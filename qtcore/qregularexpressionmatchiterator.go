package qtcore

// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QRegularExpressionMatchIterator struct {
	*qtrt.CObject
}
type QRegularExpressionMatchIterator_ITF interface {
	QRegularExpressionMatchIterator_PTR() *QRegularExpressionMatchIterator
}

func (ptr *QRegularExpressionMatchIterator) QRegularExpressionMatchIterator_PTR() *QRegularExpressionMatchIterator {
	return ptr
}

func (this *QRegularExpressionMatchIterator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRegularExpressionMatchIterator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRegularExpressionMatchIteratorFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatchIterator {
	return &QRegularExpressionMatchIterator{&qtrt.CObject{cthis}}
}
func (*QRegularExpressionMatchIterator) NewFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatchIterator {
	return NewQRegularExpressionMatchIteratorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qregularexpression.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionMatchIterator()

/*

 */
func (*QRegularExpressionMatchIterator) NewForInherit() *QRegularExpressionMatchIterator {
	return NewQRegularExpressionMatchIterator()
}
func NewQRegularExpressionMatchIterator() *QRegularExpressionMatchIterator {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratorC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegularExpressionMatchIterator)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRegularExpressionMatchIterator()

/*

 */
func DeleteQRegularExpressionMatchIterator(this *QRegularExpressionMatchIterator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qregularexpression.h:252
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator & operator=(const QRegularExpressionMatchIterator &)

/*

 */
func (this *QRegularExpressionMatchIterator) Operator_equal(iterator QRegularExpressionMatchIterator_ITF) *QRegularExpressionMatchIterator {
	var convArg0 unsafe.Pointer
	if iterator != nil && iterator.QRegularExpressionMatchIterator_PTR() != nil {
		convArg0 = iterator.QRegularExpressionMatchIterator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratoraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:254
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QRegularExpressionMatchIterator & operator=(QRegularExpressionMatchIterator &&)

/*

 */
func (this *QRegularExpressionMatchIterator) Operator_equal_1(iterator unsafe.Pointer /*333*/) *QRegularExpressionMatchIterator {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratoraSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), iterator)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatchIterator)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegularExpressionMatchIterator &)

/*
Swaps the regular expression other with this regular expression. This operation is very fast and never fails.
*/
func (this *QRegularExpressionMatchIterator) Swap(other QRegularExpressionMatchIterator_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRegularExpressionMatchIterator_PTR() != nil {
		convArg0 = other.QRegularExpressionMatchIterator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIterator4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:259
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the regular expression is a valid regular expression (that is, it contains no syntax errors, etc.), or false otherwise. Use errorString() to obtain a textual description of the error.

See also errorString() and patternErrorOffset().
*/
func (this *QRegularExpressionMatchIterator) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNext() const

/*

 */
func (this *QRegularExpressionMatchIterator) HasNext() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator7hasNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:262
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch next()

/*

 */
func (this *QRegularExpressionMatchIterator) Next() *QRegularExpressionMatch /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIterator4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch peekNext() const

/*

 */
func (this *QRegularExpressionMatchIterator) PeekNext() *QRegularExpressionMatch /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator8peekNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:265
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression regularExpression() const

/*

 */
func (this *QRegularExpressionMatchIterator) RegularExpression() *QRegularExpression /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator17regularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:266
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::MatchType matchType() const

/*

 */
func (this *QRegularExpressionMatchIterator) MatchType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator9matchTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:267
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::MatchOptions matchOptions() const

/*

 */
func (this *QRegularExpressionMatchIterator) MatchOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator12matchOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
