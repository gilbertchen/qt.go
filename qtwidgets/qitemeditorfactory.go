package qtwidgets

// /usr/include/qt/QtWidgets/qitemeditorfactory.h
// #include <qitemeditorfactory.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QItemEditorFactory struct {
	*qtrt.CObject
}
type QItemEditorFactory_ITF interface {
	QItemEditorFactory_PTR() *QItemEditorFactory
}

func (ptr *QItemEditorFactory) QItemEditorFactory_PTR() *QItemEditorFactory { return ptr }

func (this *QItemEditorFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QItemEditorFactory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQItemEditorFactoryFromPointer(cthis unsafe.Pointer) *QItemEditorFactory {
	return &QItemEditorFactory{&qtrt.CObject{cthis}}
}
func (*QItemEditorFactory) NewFromPointer(cthis unsafe.Pointer) *QItemEditorFactory {
	return NewQItemEditorFactoryFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QItemEditorFactory()

/*
Constructs a new item editor factory.
*/
func (*QItemEditorFactory) NewForInherit() *QItemEditorFactory {
	return NewQItemEditorFactory()
}
func NewQItemEditorFactory() *QItemEditorFactory {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactoryC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQItemEditorFactory)
	return gothis
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QItemEditorFactory()

/*

 */
func DeleteQItemEditorFactory(this *QItemEditorFactory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QWidget * createEditor(int, QWidget *) const

/*
Creates an editor widget with the given parent for the specified userType of data, and returns it as a QWidget.

See also registerEditor().
*/
func (this *QItemEditorFactory) CreateEditor(userType int, parent QWidget_ITF /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QItemEditorFactory12createEditorEiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), userType, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:102
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QByteArray valuePropertyName(int) const

/*
Returns the property name used to access data for the given userType of data.
*/
func (this *QItemEditorFactory) ValuePropertyName(userType int) *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QItemEditorFactory17valuePropertyNameEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), userType)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void registerEditor(int, QItemEditorCreatorBase *)

/*
Registers an item editor creator specified by creator for the given userType of data.

Note: The factory takes ownership of the item editor creator and will destroy it if a new creator for the same type is registered later.

See also createEditor().
*/
func (this *QItemEditorFactory) RegisterEditor(userType int, creator QItemEditorCreatorBase_ITF /*777 QItemEditorCreatorBase **/) {
	var convArg1 unsafe.Pointer
	if creator != nil && creator.QItemEditorCreatorBase_PTR() != nil {
		convArg1 = creator.QItemEditorCreatorBase_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactory14registerEditorEiP22QItemEditorCreatorBase", qtrt.FFI_TYPE_POINTER, this.GetCthis(), userType, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:106
// index:0
// Public static Visibility=Default Availability=Available
// [8] const QItemEditorFactory * defaultFactory()

/*
Returns the default item editor factory.

See also setDefaultFactory().
*/
func (this *QItemEditorFactory) DefaultFactory() *QItemEditorFactory /*777 const QItemEditorFactory **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactory14defaultFactoryEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQItemEditorFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QItemEditorFactory_DefaultFactory() *QItemEditorFactory /*777 const QItemEditorFactory **/ {
	var nilthis *QItemEditorFactory
	rv := nilthis.DefaultFactory()
	return rv
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:107
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultFactory(QItemEditorFactory *)

/*
Sets the default item editor factory to the given factory. Both new and existing delegates will use the new factory.

See also defaultFactory().
*/
func (this *QItemEditorFactory) SetDefaultFactory(factory QItemEditorFactory_ITF /*777 QItemEditorFactory **/) {
	var convArg0 unsafe.Pointer
	if factory != nil && factory.QItemEditorFactory_PTR() != nil {
		convArg0 = factory.QItemEditorFactory_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QItemEditorFactory17setDefaultFactoryEPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QItemEditorFactory_SetDefaultFactory(factory QItemEditorFactory_ITF /*777 QItemEditorFactory **/) {
	var nilthis *QItemEditorFactory
	nilthis.SetDefaultFactory(factory)
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
