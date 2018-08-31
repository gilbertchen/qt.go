package qtgui

// /usr/include/qt/QtGui/qpicture.h
// #include <qpicture.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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

// int metric(QPaintDevice::PaintDeviceMetric)
func (this *QPicture) InheritMetric(f func(m int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

/*

 */
type QPicture struct {
	*QPaintDevice
}
type QPicture_ITF interface {
	QPaintDevice_ITF
	QPicture_PTR() *QPicture
}

func (ptr *QPicture) QPicture_PTR() *QPicture { return ptr }

func (this *QPicture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDevice.GetCthis()
	}
}
func (this *QPicture) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDevice = NewQPaintDeviceFromPointer(cthis)
}
func NewQPictureFromPointer(cthis unsafe.Pointer) *QPicture {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QPicture{bcthis0}
}
func (*QPicture) NewFromPointer(cthis unsafe.Pointer) *QPicture {
	return NewQPictureFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpicture.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPicture(int)

/*
Constructs an empty picture.

The formatVersion parameter may be used to create a QPicture that can be read by applications that are compiled with earlier versions of Qt.

Note that the default formatVersion is -1 which signifies the current release, i.e. for Qt 4.0 a formatVersion of 7 is the same as the default formatVersion of -1.

Reading pictures generated by earlier versions of Qt is not supported in Qt 4.0.
*/
func (*QPicture) NewForInherit(formatVersion int) *QPicture {
	return NewQPicture(formatVersion)
}
func NewQPicture(formatVersion int) *QPicture {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPictureC2Ei", qtrt.FFI_TYPE_POINTER, formatVersion)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPictureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPicture)
	return gothis
}

// /usr/include/qt/QtGui/qpicture.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPicture(int)

/*
Constructs an empty picture.

The formatVersion parameter may be used to create a QPicture that can be read by applications that are compiled with earlier versions of Qt.

Note that the default formatVersion is -1 which signifies the current release, i.e. for Qt 4.0 a formatVersion of 7 is the same as the default formatVersion of -1.

Reading pictures generated by earlier versions of Qt is not supported in Qt 4.0.
*/
func (*QPicture) NewForInherit__() *QPicture {
	return NewQPicture__()
}
func NewQPicture__() *QPicture {
	// arg: 0, int=Int, =Invalid, , Invalid
	formatVersion := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPictureC2Ei", qtrt.FFI_TYPE_POINTER, formatVersion)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPictureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPicture)
	return gothis
}

// /usr/include/qt/QtGui/qpicture.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPicture()

/*

 */
func DeleteQPicture(this *QPicture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPictureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpicture.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the picture contains no data; otherwise returns false.
*/
func (this *QPicture) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType() const

/*

 */
func (this *QPicture) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpicture.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] uint size() const

/*
Returns the size of the picture data.

See also data().
*/
func (this *QPicture) Size() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qpicture.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * data() const

/*
Returns a pointer to the picture data. The pointer is only valid until the next non-const function is called on this picture. The returned pointer is 0 if the picture contains no data.

See also setData(), size(), and isNull().
*/
func (this *QPicture) Data() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qpicture.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setData(const char *, uint)

/*
Sets the picture data directly from data and size. This function copies the input data.

See also data() and size().
*/
func (this *QPicture) SetData(data string, size uint) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture7setDataEPKcj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool play(QPainter *)

/*
Replays the picture using painter, and returns true if successful; otherwise returns false.

This function does exactly the same as QPainter::drawPicture() with (x, y) = (0, 0).
*/
func (this *QPicture) Play(p QPainter_ITF /*777 QPainter **/) bool {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPainter_PTR() != nil {
		convArg0 = p.QPainter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4playEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(QIODevice *, const char *)

/*
Loads a picture from the file specified by fileName and returns true if successful; otherwise invalidates the picture and returns false.

Please note that the format parameter has been deprecated and will have no effect.

See also save().
*/
func (this *QPicture) Load(dev qtcore.QIODevice_ITF /*777 QIODevice **/, format string) bool {
	var convArg0 unsafe.Pointer
	if dev != nil && dev.QIODevice_PTR() != nil {
		convArg0 = dev.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4loadEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(QIODevice *, const char *)

/*
Loads a picture from the file specified by fileName and returns true if successful; otherwise invalidates the picture and returns false.

Please note that the format parameter has been deprecated and will have no effect.

See also save().
*/
func (this *QPicture) Load__(dev qtcore.QIODevice_ITF /*777 QIODevice **/) bool {
	var convArg0 unsafe.Pointer
	if dev != nil && dev.QIODevice_PTR() != nil {
		convArg0 = dev.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4loadEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:73
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *)

/*
Loads a picture from the file specified by fileName and returns true if successful; otherwise invalidates the picture and returns false.

Please note that the format parameter has been deprecated and will have no effect.

See also save().
*/
func (this *QPicture) Load_1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4loadERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:73
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const char *)

/*
Loads a picture from the file specified by fileName and returns true if successful; otherwise invalidates the picture and returns false.

Please note that the format parameter has been deprecated and will have no effect.

See also save().
*/
func (this *QPicture) Load_1_(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4loadERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *)

/*
Saves a picture to the file specified by fileName and returns true if successful; otherwise returns false.

Please note that the format parameter has been deprecated and will have no effect.

See also load().
*/
func (this *QPicture) Save(dev qtcore.QIODevice_ITF /*777 QIODevice **/, format string) bool {
	var convArg0 unsafe.Pointer
	if dev != nil && dev.QIODevice_PTR() != nil {
		convArg0 = dev.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4saveEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool save(QIODevice *, const char *)

/*
Saves a picture to the file specified by fileName and returns true if successful; otherwise returns false.

Please note that the format parameter has been deprecated and will have no effect.

See also load().
*/
func (this *QPicture) Save__(dev qtcore.QIODevice_ITF /*777 QIODevice **/) bool {
	var convArg0 unsafe.Pointer
	if dev != nil && dev.QIODevice_PTR() != nil {
		convArg0 = dev.QIODevice_PTR().GetCthis()
	}
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4saveEP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:75
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *)

/*
Saves a picture to the file specified by fileName and returns true if successful; otherwise returns false.

Please note that the format parameter has been deprecated and will have no effect.

See also load().
*/
func (this *QPicture) Save_1(fileName string, format string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4saveERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:75
// index:1
// Public Visibility=Default Availability=Available
// [1] bool save(const QString &, const char *)

/*
Saves a picture to the file specified by fileName and returns true if successful; otherwise returns false.

Please note that the format parameter has been deprecated and will have no effect.

See also load().
*/
func (this *QPicture) Save_1_(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const char *=Pointer, =Invalid, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4saveERK7QStringPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:77
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect boundingRect() const

/*
Returns the picture's bounding rectangle or an invalid rectangle if the picture contains no data.

See also setBoundingRect().
*/
func (this *QPicture) BoundingRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpicture.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBoundingRect(const QRect &)

/*
Sets the picture's bounding rectangle to r. The automatically calculated value is overridden.

See also boundingRect().
*/
func (this *QPicture) SetBoundingRect(r qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRect_PTR() != nil {
		convArg0 = r.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture15setBoundingRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:80
// index:0
// Public Visibility=Default Availability=Available
// [32] QPicture & operator=(const QPicture &)

/*

 */
func (this *QPicture) Operator_equal(p QPicture_ITF) *QPicture {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPicture_PTR() != nil {
		convArg0 = p.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPictureaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPictureFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPicture)
	return rv2
}

// /usr/include/qt/QtGui/qpicture.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QPicture & operator=(QPicture &&)

/*

 */
func (this *QPicture) Operator_equal_1(other unsafe.Pointer /*333*/) *QPicture {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPictureaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPictureFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPicture)
	return rv2
}

// /usr/include/qt/QtGui/qpicture.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPicture &)

/*
Swaps picture other with this picture. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func (this *QPicture) Swap(other QPicture_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPicture_PTR() != nil {
		convArg0 = other.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()

/*

 */
func (this *QPicture) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
func (this *QPicture) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [8] const char * pictureFormat(const QString &)

/*

 */
func (this *QPicture) PictureFormat(fileName string) string {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture13pictureFormatERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}
func QPicture_PictureFormat(fileName string) string {
	var nilthis *QPicture
	rv := nilthis.PictureFormat(fileName)
	return rv
}

// /usr/include/qt/QtGui/qpicture.h:97
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList inputFormatList()

/*

 */
func (this *QPicture) InputFormatList() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture15inputFormatListEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QPicture_InputFormatList() *qtcore.QStringList /*123*/ {
	var nilthis *QPicture
	rv := nilthis.InputFormatList()
	return rv
}

// /usr/include/qt/QtGui/qpicture.h:98
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList outputFormatList()

/*

 */
func (this *QPicture) OutputFormatList() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QPicture16outputFormatListEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QPicture_OutputFormatList() *qtcore.QStringList /*123*/ {
	var nilthis *QPicture
	rv := nilthis.OutputFormatList()
	return rv
}

// /usr/include/qt/QtGui/qpicture.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const

/*

 */
func (this *QPicture) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpicture.h:106
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(QPaintDevice::PaintDeviceMetric) const

/*

 */
func (this *QPicture) Metric(m int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QPicture6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), m)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
