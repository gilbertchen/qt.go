package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
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
type QGesture struct {
	*qtcore.QObject
}
type QGesture_ITF interface {
	qtcore.QObject_ITF
	QGesture_PTR() *QGesture
}

func (ptr *QGesture) QGesture_PTR() *QGesture { return ptr }

func (this *QGesture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGesture) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGestureFromPointer(cthis unsafe.Pointer) *QGesture {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGesture{bcthis0}
}
func (*QGesture) NewFromPointer(cthis unsafe.Pointer) *QGesture {
	return NewQGestureFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGesture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgesture.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGesture(QObject *)

/*
Constructs a new gesture object with the given parent.

QGesture objects are created by gesture recognizers in the QGestureRecognizer::create() function.
*/
func (*QGesture) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QGesture {
	return NewQGesture(parent)
}
func NewQGesture(parent qtcore.QObject_ITF /*777 QObject **/) *QGesture {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGesture")
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGesture(QObject *)

/*
Constructs a new gesture object with the given parent.

QGesture objects are created by gesture recognizers in the QGestureRecognizer::create() function.
*/
func (*QGesture) NewForInherit__() *QGesture {
	return NewQGesture__()
}
func NewQGesture__() *QGesture {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGesture")
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGesture()

/*

 */
func DeleteQGesture(this *QGesture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGestureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::GestureType gestureType() const

/*

 */
func (this *QGesture) GestureType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGesture11gestureTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::GestureState state() const

/*

 */
func (this *QGesture) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGesture5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:80
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF hotSpot() const

/*

 */
func (this *QGesture) HotSpot() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGesture7hotSpotEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHotSpot(const QPointF &)

/*

 */
func (this *QGesture) SetHotSpot(value qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QPointF_PTR() != nil {
		convArg0 = value.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGesture10setHotSpotERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasHotSpot() const

/*

 */
func (this *QGesture) HasHotSpot() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGesture10hasHotSpotEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgesture.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetHotSpot()

/*

 */
func (this *QGesture) UnsetHotSpot() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGesture12unsetHotSpotEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGestureCancelPolicy(QGesture::GestureCancelPolicy)

/*

 */
func (this *QGesture) SetGestureCancelPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QGesture22setGestureCancelPolicyENS_19GestureCancelPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgesture.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] QGesture::GestureCancelPolicy gestureCancelPolicy() const

/*

 */
func (this *QGesture) GestureCancelPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QGesture19gestureCancelPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*
This enum describes how accepting a gesture can cancel other gestures automatically.


*/
type QGesture__GestureCancelPolicy = int

// On accepting this gesture no other gestures will be affected.
const QGesture__CancelNone QGesture__GestureCancelPolicy = 0

// On accepting this gesture all gestures that are active in the context (respecting the Qt::GestureFlag that were specified when subscribed to the gesture) will be cancelled.
const QGesture__CancelAllInContext QGesture__GestureCancelPolicy = 1

func (this *QGesture) GestureCancelPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGesture_GestureCancelPolicyItemName(val int) string {
	var nilthis *QGesture
	return nilthis.GestureCancelPolicyItemName(val)
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
