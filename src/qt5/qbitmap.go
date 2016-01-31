package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qbitmap.h
// dst-file: /src/gui/qbitmap.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QBitmap::clear();
extern void C_ZN7QBitmap5clearEv(void* qthis); // 2
  // proto:  void QBitmap::~QBitmap();
extern void C_ZN7QBitmapD2Ev(void* qthis); // 4
  // proto:  void QBitmap::QBitmap();
extern void* C_ZN7QBitmapC2Ev(); // 3
  // proto:  void QBitmap::QBitmap(const QPixmap & );
extern void* C_ZN7QBitmapC2ERK7QPixmap(void* arg0); // 3
  // proto:  void QBitmap::QBitmap(const QString & fileName, const char * format);
extern void* C_ZN7QBitmapC2ERK7QStringPKc(void* arg0, void* arg1); // 3
  // proto:  void QBitmap::QBitmap(const QSize & );
extern void* C_ZN7QBitmapC2ERK5QSize(void* arg0); // 3
  // proto:  void QBitmap::QBitmap(int w, int h);
extern void* C_ZN7QBitmapC2Eii(int32_t arg0, int32_t arg1); // 3
  // proto:  QBitmap QBitmap::transformed(const QTransform & matrix);
extern void* C_ZNK7QBitmap11transformedERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  QBitmap QBitmap::transformed(const QMatrix & );
extern void* C_ZNK7QBitmap11transformedERK7QMatrix(void* qthis, void* arg0); // 4
  // proto:  void QBitmap::swap(QBitmap & other);
extern void C_ZN7QBitmap4swapERS_(void* qthis, void* arg0); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QBitmap)=1
type QBitmap struct {
  /*qbase*/ QPixmap;
  qclsinst unsafe.Pointer /* *C.void */;
}

// clear()
func (this *QBitmap) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBitmap5clearEv
    // invoke: void clear()
    C.C_ZN7QBitmap5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBitmap", "clear", args)
  }

  return
}

// ~QBitmap()
func (this *QBitmap) Freeqbitmap(args ...interface{}) () {
  // ~QBitmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBitmapD0Ev
    // invoke: void ~QBitmap()
    C.C_ZN7QBitmapD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBitmap", "~QBitmap", args)
  }

  return
}

// QBitmap()
func NewQBitmap(args ...interface{}) *QBitmap {
  // QBitmap()
  // QBitmap(const class QPixmap &)
  // QBitmap(const class QString &, const char *)
  // QBitmap(const class QSize &)
  // QBitmap(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBitmapC1Ev
    // invoke: void QBitmap()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QBitmapC2Ev()
    return &QBitmap{qclsinst:qthis}
  case 1:
    // invoke: _ZN7QBitmapC1ERK7QPixmap
    // invoke: void QBitmap(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QBitmapC2ERK7QPixmap(arg0)
    return &QBitmap{qclsinst:qthis}
  case 2:
    // invoke: _ZN7QBitmapC1ERK7QStringPKc
    // invoke: void QBitmap(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[2][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QBitmapC2ERK7QStringPKc(arg0, arg1)
    return &QBitmap{qclsinst:qthis}
  case 3:
    // invoke: _ZN7QBitmapC1ERK5QSize
    // invoke: void QBitmap(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QBitmapC2ERK5QSize(arg0)
    return &QBitmap{qclsinst:qthis}
  case 4:
    // invoke: _ZN7QBitmapC1Eii
    // invoke: void QBitmap(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QBitmapC2Eii(arg0, arg1)
    return &QBitmap{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QBitmap", "QBitmap", args)
  }

  return nil // QBitmap{qclsinst:qthis}
}

// transformed(const class QTransform &)
func (this *QBitmap) Transformed(args ...interface{}) (ret interface{}) {
  // transformed(const class QTransform &)
  // transformed(const class QMatrix &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBitmap11transformedERK10QTransform
    // invoke: QBitmap transformed(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QBitmap11transformedERK10QTransform(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBitmap{}) // "QBitmap"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK7QBitmap11transformedERK7QMatrix
    // invoke: QBitmap transformed(const class QMatrix &)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QBitmap11transformedERK7QMatrix(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBitmap{}) // "QBitmap"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QBitmap", "transformed", args)
  }

  return
}

// swap(class QBitmap &)
func (this *QBitmap) Swap(args ...interface{}) () {
  // swap(class QBitmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitmap{}) // "QBitmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBitmap4swapERS_
    // invoke: void swap(class QBitmap &)
    var arg0 = args[0].(QBitmap).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QBitmap4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBitmap", "swap", args)
  }

  return
}

// <= body block end
