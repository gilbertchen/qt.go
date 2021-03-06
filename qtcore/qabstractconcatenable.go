package qtcore

// /usr/include/qt/QtCore/qstringbuilder.h
// #include <qstringbuilder.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void convertFromAscii(const char *, int, QChar *&)
func (this *QAbstractConcatenable) InheritConvertFromAscii(f func(a string, len_ int, out *QChar) /*void*/) {
	qtrt.SetAllInheritCallback(this, "convertFromAscii", f)
}

// void appendLatin1To(const char *, int, QChar *)
func (this *QAbstractConcatenable) InheritAppendLatin1To(f func(a string, len_ int, out *QChar /*777 QChar **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "appendLatin1To", f)
}

/*

 */
type QAbstractConcatenable struct {
	*qtrt.CObject
}
type QAbstractConcatenable_ITF interface {
	QAbstractConcatenable_PTR() *QAbstractConcatenable
}

func (ptr *QAbstractConcatenable) QAbstractConcatenable_PTR() *QAbstractConcatenable { return ptr }

func (this *QAbstractConcatenable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAbstractConcatenable) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAbstractConcatenableFromPointer(cthis unsafe.Pointer) *QAbstractConcatenable {
	return &QAbstractConcatenable{&qtrt.CObject{cthis}}
}
func (*QAbstractConcatenable) NewFromPointer(cthis unsafe.Pointer) *QAbstractConcatenable {
	return NewQAbstractConcatenableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstringbuilder.h:61
// index:0
// Protected static Visibility=Default Availability=Available
// [-2] void convertFromAscii(const char *, int, QChar *&)

/*

 */
func (this *QAbstractConcatenable) ConvertFromAscii(a string, len_ int, out QChar_ITF) {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	var convArg2 unsafe.Pointer
	if out != nil && out.QChar_PTR() != nil {
		convArg2 = out.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractConcatenable16convertFromAsciiEPKciRP5QChar", qtrt.FFI_TYPE_POINTER, convArg0, len_, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QAbstractConcatenable_ConvertFromAscii(a string, len_ int, out QChar_ITF) {
	var nilthis *QAbstractConcatenable
	nilthis.ConvertFromAscii(a, len_, out)
}

// /usr/include/qt/QtCore/qstringbuilder.h:62
// index:1
// Protected static inline Visibility=Default Availability=Available
// [-2] void convertFromAscii(char, QChar *&)

/*

 */
func (this *QAbstractConcatenable) ConvertFromAscii_1(a byte, out QChar_ITF) {
	var convArg1 unsafe.Pointer
	if out != nil && out.QChar_PTR() != nil {
		convArg1 = out.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractConcatenable16convertFromAsciiEcRP5QChar", qtrt.FFI_TYPE_POINTER, a, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QAbstractConcatenable_ConvertFromAscii_1(a byte, out QChar_ITF) {
	var nilthis *QAbstractConcatenable
	nilthis.ConvertFromAscii_1(a, out)
}

// /usr/include/qt/QtCore/qstringbuilder.h:66
// index:0
// Protected static Visibility=Default Availability=Available
// [-2] void appendLatin1To(const char *, int, QChar *)

/*

 */
func (this *QAbstractConcatenable) AppendLatin1To(a string, len_ int, out QChar_ITF /*777 QChar **/) {
	var convArg0 = qtrt.CString(a)
	defer qtrt.FreeMem(convArg0)
	var convArg2 unsafe.Pointer
	if out != nil && out.QChar_PTR() != nil {
		convArg2 = out.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractConcatenable14appendLatin1ToEPKciP5QChar", qtrt.FFI_TYPE_POINTER, convArg0, len_, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QAbstractConcatenable_AppendLatin1To(a string, len_ int, out QChar_ITF /*777 QChar **/) {
	var nilthis *QAbstractConcatenable
	nilthis.AppendLatin1To(a, len_, out)
}

func DeleteQAbstractConcatenable(this *QAbstractConcatenable) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractConcatenableD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
