package qtwidgets

// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QHBoxLayout struct {
	*QBoxLayout
}
type QHBoxLayout_ITF interface {
	QBoxLayout_ITF
	QHBoxLayout_PTR() *QHBoxLayout
}

func (ptr *QHBoxLayout) QHBoxLayout_PTR() *QHBoxLayout { return ptr }

func (this *QHBoxLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QBoxLayout.GetCthis()
	}
}
func (this *QHBoxLayout) SetCthis(cthis unsafe.Pointer) {
	this.QBoxLayout = NewQBoxLayoutFromPointer(cthis)
}
func NewQHBoxLayoutFromPointer(cthis unsafe.Pointer) *QHBoxLayout {
	bcthis0 := NewQBoxLayoutFromPointer(cthis)
	return &QHBoxLayout{bcthis0}
}
func (*QHBoxLayout) NewFromPointer(cthis unsafe.Pointer) *QHBoxLayout {
	return NewQHBoxLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QHBoxLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QHBoxLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qboxlayout.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHBoxLayout()

/*

 */
func (*QHBoxLayout) NewForInherit() *QHBoxLayout {
	return NewQHBoxLayout()
}
func NewQHBoxLayout() *QHBoxLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHBoxLayoutC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHBoxLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:118
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QHBoxLayout(QWidget *)

/*

 */
func (*QHBoxLayout) NewForInherit_1(parent QWidget_ITF /*777 QWidget **/) *QHBoxLayout {
	return NewQHBoxLayout_1(parent)
}
func NewQHBoxLayout_1(parent QWidget_ITF /*777 QWidget **/) *QHBoxLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHBoxLayoutC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHBoxLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QHBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHBoxLayout()

/*

 */
func DeleteQHBoxLayout(this *QHBoxLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QHBoxLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
