package qtgui

// /usr/include/qt/QtGui/qbackingstore.h
// #include <qbackingstore.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 146
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QBackingStore struct {
	*qtrt.CObject
}
type QBackingStore_ITF interface {
	QBackingStore_PTR() *QBackingStore
}

func (ptr *QBackingStore) QBackingStore_PTR() *QBackingStore { return ptr }

func (this *QBackingStore) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBackingStore) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQBackingStoreFromPointer(cthis unsafe.Pointer) *QBackingStore {
	return &QBackingStore{&qtrt.CObject{cthis}}
}
func (*QBackingStore) NewFromPointer(cthis unsafe.Pointer) *QBackingStore {
	return NewQBackingStoreFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbackingstore.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBackingStore(QWindow *)

/*
Constructs an empty surface for the given top-level window.
*/
func (*QBackingStore) NewForInherit(window QWindow_ITF /*777 QWindow **/) *QBackingStore {
	return NewQBackingStore(window)
}
func NewQBackingStore(window QWindow_ITF /*777 QWindow **/) *QBackingStore {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStoreC2EP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBackingStoreFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBackingStore)
	return gothis
}

// /usr/include/qt/QtGui/qbackingstore.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QBackingStore()

/*

 */
func DeleteQBackingStore(this *QBackingStore) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStoreD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qbackingstore.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * window() const

/*
Returns a pointer to the top-level window associated with this surface.
*/
func (this *QBackingStore) Window() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QBackingStore6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qbackingstore.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QPaintDevice * paintDevice()

/*
Returns the paint device for this surface.

Warning: The device is only valid between calls to beginPaint() and endPaint(). You should not cache the returned value.
*/
func (this *QBackingStore) PaintDevice() *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStore11paintDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qbackingstore.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush(const QRegion &, QWindow *, const QPoint &)

/*
Flushes the given region from the specified window onto the screen.

The window must either be the top level window represented by this backingstore, or a non-transient child of that window. Passing nullptr falls back to using the backingstore's top level window.

If the window is a child window, the region should be in child window coordinates, and the offset should be the child window's offset in relation to the backingstore's top level window.

You should call this function after ending painting with endPaint().

See also QWindow::transientParent().
*/
func (this *QBackingStore) Flush(region QRegion_ITF, window QWindow_ITF /*777 QWindow **/, offset qtcore.QPoint_ITF) {
	var convArg0 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg0 = region.QRegion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg1 = window.QWindow_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg2 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush(const QRegion &, QWindow *, const QPoint &)

/*
Flushes the given region from the specified window onto the screen.

The window must either be the top level window represented by this backingstore, or a non-transient child of that window. Passing nullptr falls back to using the backingstore's top level window.

If the window is a child window, the region should be in child window coordinates, and the offset should be the child window's offset in relation to the backingstore's top level window.

You should call this function after ending painting with endPaint().

See also QWindow::transientParent().
*/
func (this *QBackingStore) Flush__(region QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg0 = region.QRegion_PTR().GetCthis()
	}
	// arg: 1, QWindow *=Pointer, QWindow=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, const QPoint &=LValueReference, QPoint=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flush(const QRegion &, QWindow *, const QPoint &)

/*
Flushes the given region from the specified window onto the screen.

The window must either be the top level window represented by this backingstore, or a non-transient child of that window. Passing nullptr falls back to using the backingstore's top level window.

If the window is a child window, the region should be in child window coordinates, and the offset should be the child window's offset in relation to the backingstore's top level window.

You should call this function after ending painting with endPaint().

See also QWindow::transientParent().
*/
func (this *QBackingStore) Flush__1(region QRegion_ITF, window QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg0 = region.QRegion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg1 = window.QWindow_PTR().GetCthis()
	}
	// arg: 2, const QPoint &=LValueReference, QPoint=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(const QSize &)

/*
Sets the size of the window surface to size.

See also size().
*/
func (this *QBackingStore) Resize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStore6resizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const

/*
Returns the current size of the window surface.
*/
func (this *QBackingStore) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QBackingStore4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qbackingstore.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool scroll(const QRegion &, int, int)

/*
Scrolls the given area dx pixels to the right and dy downward; both dx and dy may be negative.

Returns true if the area was scrolled successfully; false otherwise.
*/
func (this *QBackingStore) Scroll(area QRegion_ITF, dx int, dy int) bool {
	var convArg0 unsafe.Pointer
	if area != nil && area.QRegion_PTR() != nil {
		convArg0 = area.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStore6scrollERK7QRegionii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, dx, dy)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qbackingstore.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginPaint(const QRegion &)

/*
Begins painting on the backing store surface in the given region.

You should call this function before using the paintDevice() to paint.

See also endPaint() and paintDevice().
*/
func (this *QBackingStore) BeginPaint(arg0 QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRegion_PTR() != nil {
		convArg0 = arg0.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStore10beginPaintERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endPaint()

/*
Ends painting.

You should call this function after painting with the paintDevice() has ended.

See also beginPaint() and paintDevice().
*/
func (this *QBackingStore) EndPaint() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStore8endPaintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStaticContents(const QRegion &)

/*
Set region as the static contents of this window.

See also staticContents().
*/
func (this *QBackingStore) SetStaticContents(region QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg0 = region.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QBackingStore17setStaticContentsERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion staticContents() const

/*
Returns a QRegion representing the area of the window that has static contents.

See also setStaticContents().
*/
func (this *QBackingStore) StaticContents() *QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QBackingStore14staticContentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qbackingstore.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasStaticContents() const

/*
Returns a boolean indicating if this window has static contents or not.
*/
func (this *QBackingStore) HasStaticContents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QBackingStore17hasStaticContentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
}

//  keep block end
