package qtgui

// /usr/include/qt/QtGui/qtextlayout.h
// #include <qtextlayout.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 41
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
type QTextLine struct {
	*qtrt.CObject
}
type QTextLine_ITF interface {
	QTextLine_PTR() *QTextLine
}

func (ptr *QTextLine) QTextLine_PTR() *QTextLine { return ptr }

func (this *QTextLine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextLine) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextLineFromPointer(cthis unsafe.Pointer) *QTextLine {
	return &QTextLine{&qtrt.CObject{cthis}}
}
func (*QTextLine) NewFromPointer(cthis unsafe.Pointer) *QTextLine {
	return NewQTextLineFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextlayout.h:213
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextLine()

/*

 */
func (*QTextLine) NewForInherit() *QTextLine {
	return NewQTextLine()
}
func NewQTextLine() *QTextLine {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLineC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextLineFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLine)
	return gothis
}

// /usr/include/qt/QtGui/qtextlayout.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QTextLine) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:216
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect() const

/*

 */
func (this *QTextLine) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:217
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal x() const

/*

 */
func (this *QTextLine) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:218
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal y() const

/*

 */
func (this *QTextLine) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width() const

/*

 */
func (this *QTextLine) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:220
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent() const

/*

 */
func (this *QTextLine) Ascent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine6ascentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:221
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent() const

/*

 */
func (this *QTextLine) Descent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine7descentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal height() const

/*

 */
func (this *QTextLine) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:223
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leading() const

/*

 */
func (this *QTextLine) Leading() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine7leadingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLeadingIncluded(bool)

/*

 */
func (this *QTextLine) SetLeadingIncluded(included bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine18setLeadingIncludedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), included)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:226
// index:0
// Public Visibility=Default Availability=Available
// [1] bool leadingIncluded() const

/*

 */
func (this *QTextLine) LeadingIncluded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine15leadingIncludedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextlayout.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal naturalTextWidth() const

/*

 */
func (this *QTextLine) NaturalTextWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine16naturalTextWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal horizontalAdvance() const

/*

 */
func (this *QTextLine) HorizontalAdvance() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine17horizontalAdvanceEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:230
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF naturalTextRect() const

/*

 */
func (this *QTextLine) NaturalTextRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine15naturalTextRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:242
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal cursorToX(int *, QTextLine::Edge) const

/*

 */
func (this *QTextLine) CursorToX(cursorPos unsafe.Pointer /*666*/, edge int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEPiNS_4EdgeE", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), cursorPos, edge)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:242
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal cursorToX(int *, QTextLine::Edge) const

/*

 */
func (this *QTextLine) CursorToX__(cursorPos unsafe.Pointer /*666*/) float64 {
	// arg: 1, QTextLine::Edge=Enum, QTextLine::Edge=Enum, , Invalid
	edge := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEPiNS_4EdgeE", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), cursorPos, edge)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:243
// index:1
// Public inline Visibility=Default Availability=Available
// [8] qreal cursorToX(int, QTextLine::Edge) const

/*

 */
func (this *QTextLine) CursorToX_1(cursorPos int, edge int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEiNS_4EdgeE", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), cursorPos, edge)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:243
// index:1
// Public inline Visibility=Default Availability=Available
// [8] qreal cursorToX(int, QTextLine::Edge) const

/*

 */
func (this *QTextLine) CursorToX_1_(cursorPos int) float64 {
	// arg: 1, QTextLine::Edge=Enum, QTextLine::Edge=Enum, , Invalid
	edge := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9cursorToXEiNS_4EdgeE", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), cursorPos, edge)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:244
// index:0
// Public Visibility=Default Availability=Available
// [4] int xToCursor(qreal, QTextLine::CursorPosition) const

/*

 */
func (this *QTextLine) XToCursor(x float64, arg1 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9xToCursorEdNS_14CursorPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, arg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:244
// index:0
// Public Visibility=Default Availability=Available
// [4] int xToCursor(qreal, QTextLine::CursorPosition) const

/*

 */
func (this *QTextLine) XToCursor__(x float64) int {
	// arg: 1, QTextLine::CursorPosition=Enum, QTextLine::CursorPosition=Enum, , Invalid
	arg1 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9xToCursorEdNS_14CursorPositionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, arg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLineWidth(qreal)

/*

 */
func (this *QTextLine) SetLineWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine12setLineWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNumColumns(int)

/*

 */
func (this *QTextLine) SetNumColumns(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine13setNumColumnsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:248
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setNumColumns(int, qreal)

/*

 */
func (this *QTextLine) SetNumColumns_1(columns int, alignmentWidth float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine13setNumColumnsEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns, alignmentWidth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)

/*
Moves the text layout to point p.

See also position().
*/
func (this *QTextLine) SetPosition(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLine11setPositionERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextlayout.h:251
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position() const

/*
The global position of the layout. This is independent of the bounding rectangle and of the layout process.

This function was introduced in  Qt 4.2.

See also setPosition().
*/
func (this *QTextLine) Position() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qtextlayout.h:253
// index:0
// Public Visibility=Default Availability=Available
// [4] int textStart() const

/*

 */
func (this *QTextLine) TextStart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine9textStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:254
// index:0
// Public Visibility=Default Availability=Available
// [4] int textLength() const

/*

 */
func (this *QTextLine) TextLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine10textLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextlayout.h:256
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lineNumber() const

/*

 */
func (this *QTextLine) LineNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTextLine10lineNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQTextLine(this *QTextLine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTextLineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QTextLine__Edge = int

//
const QTextLine__Leading QTextLine__Edge = 0

//
const QTextLine__Trailing QTextLine__Edge = 1

func (this *QTextLine) EdgeItemName(val int) string {
	switch val {
	case QTextLine__Leading: // 0
		return "Leading"
	case QTextLine__Trailing: // 1
		return "Trailing"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextLine_EdgeItemName(val int) string {
	var nilthis *QTextLine
	return nilthis.EdgeItemName(val)
}

/*


 */
type QTextLine__CursorPosition = int

//
const QTextLine__CursorBetweenCharacters QTextLine__CursorPosition = 0

//
const QTextLine__CursorOnCharacter QTextLine__CursorPosition = 1

func (this *QTextLine) CursorPositionItemName(val int) string {
	switch val {
	case QTextLine__CursorBetweenCharacters: // 0
		return "CursorBetweenCharacters"
	case QTextLine__CursorOnCharacter: // 1
		return "CursorOnCharacter"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextLine_CursorPositionItemName(val int) string {
	var nilthis *QTextLine
	return nilthis.CursorPositionItemName(val)
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
