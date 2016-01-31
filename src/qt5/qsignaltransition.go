package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qsignaltransition.h
// dst-file: /src/core/qsignaltransition.go
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
  // proto:  QByteArray QSignalTransition::signal();
extern void* C_ZNK17QSignalTransition6signalEv(void* qthis); // 4
  // proto:  void QSignalTransition::~QSignalTransition();
extern void C_ZN17QSignalTransitionD2Ev(void* qthis); // 4
  // proto:  void QSignalTransition::QSignalTransition(QState * sourceState);
extern void* C_ZN17QSignalTransitionC2EP6QState(void* arg0); // 3
  // proto:  void QSignalTransition::QSignalTransition(const QObject * sender, const char * signal, QState * sourceState);
extern void* C_ZN17QSignalTransitionC2EPK7QObjectPKcP6QState(void* arg0, void* arg1, void* arg2); // 3
  // proto:  const QMetaObject * QSignalTransition::metaObject();
extern void C_ZNK17QSignalTransition10metaObjectEv(void* qthis); // 4
  // proto:  void QSignalTransition::setSenderObject(const QObject * sender);
extern void C_ZN17QSignalTransition15setSenderObjectEPK7QObject(void* qthis, void* arg0); // 4
  // proto:  QObject * QSignalTransition::senderObject();
extern void* C_ZNK17QSignalTransition12senderObjectEv(void* qthis); // 4
  // proto:  void QSignalTransition::setSignal(const QByteArray & signal);
extern void C_ZN17QSignalTransition9setSignalERK10QByteArray(void* qthis, void* arg0); // 4
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

// class sizeof(QSignalTransition)=1
type QSignalTransition struct {
  /*qbase*/ QAbstractTransition;
  qclsinst unsafe.Pointer /* *C.void */;
//  _senderObjectChanged QSignalTransition_senderObjectChanged_signal;
//  _signalChanged QSignalTransition_signalChanged_signal;
}

// signal()
func (this *QSignalTransition) Signal(args ...interface{}) (ret interface{}) {
  // signal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QSignalTransition6signalEv
    // invoke: QByteArray signal()
    var ret0 = C.C_ZNK17QSignalTransition6signalEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSignalTransition", "signal", args)
  }

  return
}

// ~QSignalTransition()
func (this *QSignalTransition) Freeqsignaltransition(args ...interface{}) () {
  // ~QSignalTransition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransitionD0Ev
    // invoke: void ~QSignalTransition()
    C.C_ZN17QSignalTransitionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalTransition", "~QSignalTransition", args)
  }

  return
}

// QSignalTransition(class QState *)
func NewQSignalTransition(args ...interface{}) *QSignalTransition {
  // QSignalTransition(class QState *)
  // QSignalTransition(const class QObject *, const char *, class QState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QState{}) // "QState *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = reflect.TypeOf(QState{}) // "QState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransitionC1EP6QState
    // invoke: void QSignalTransition(class QState *)
    var arg0 = args[0].(QState).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QSignalTransitionC2EP6QState(arg0)
    return &QSignalTransition{qclsinst:qthis}
  case 1:
    // invoke: _ZN17QSignalTransitionC1EPK7QObjectPKcP6QState
    // invoke: void QSignalTransition(const class QObject *, const char *, class QState *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = args[2].(QState).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QSignalTransitionC2EPK7QObjectPKcP6QState(arg0, arg1, arg2)
    return &QSignalTransition{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSignalTransition", "QSignalTransition", args)
  }

  return nil // QSignalTransition{qclsinst:qthis}
}

// metaObject()
func (this *QSignalTransition) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QSignalTransition10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK17QSignalTransition10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSignalTransition", "metaObject", args)
  }

  return
}

// setSenderObject(const class QObject *)
func (this *QSignalTransition) Setsenderobject(args ...interface{}) () {
  // setSenderObject(const class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "const QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransition15setSenderObjectEPK7QObject
    // invoke: void setSenderObject(const class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QSignalTransition15setSenderObjectEPK7QObject(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSignalTransition", "setSenderObject", args)
  }

  return
}

// senderObject()
func (this *QSignalTransition) Senderobject(args ...interface{}) (ret interface{}) {
  // senderObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QSignalTransition12senderObjectEv
    // invoke: QObject * senderObject()
    var ret0 = C.C_ZNK17QSignalTransition12senderObjectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSignalTransition", "senderObject", args)
  }

  return
}

// setSignal(const class QByteArray &)
func (this *QSignalTransition) Setsignal(args ...interface{}) () {
  // setSignal(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QSignalTransition9setSignalERK10QByteArray
    // invoke: void setSignal(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QSignalTransition9setSignalERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSignalTransition", "setSignal", args)
  }

  return
}

// <= body block end
