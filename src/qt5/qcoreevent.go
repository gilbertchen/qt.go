package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qcoreevent.h
// dst-file: /src/core/qcoreevent.go
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
  // proto:  void QDeferredDeleteEvent::QDeferredDeleteEvent();
extern void* C_ZN20QDeferredDeleteEventC2Ev(); // 3
  // proto:  void QDeferredDeleteEvent::~QDeferredDeleteEvent();
extern void C_ZN20QDeferredDeleteEventD2Ev(void* qthis); // 4
  // proto:  int QDeferredDeleteEvent::loopLevel();
extern int32_t C_ZNK20QDeferredDeleteEvent9loopLevelEv(void* qthis); // 2
  // proto:  void QDynamicPropertyChangeEvent::QDynamicPropertyChangeEvent(const QByteArray & name);
extern void* C_ZN27QDynamicPropertyChangeEventC2ERK10QByteArray(void* arg0); // 3
  // proto:  QByteArray QDynamicPropertyChangeEvent::propertyName();
extern void* C_ZNK27QDynamicPropertyChangeEvent12propertyNameEv(void* qthis); // 2
  // proto:  void QDynamicPropertyChangeEvent::~QDynamicPropertyChangeEvent();
extern void C_ZN27QDynamicPropertyChangeEventD2Ev(void* qthis); // 4
  // proto:  int QTimerEvent::timerId();
extern int32_t C_ZNK11QTimerEvent7timerIdEv(void* qthis); // 2
  // proto:  void QTimerEvent::~QTimerEvent();
extern void C_ZN11QTimerEventD2Ev(void* qthis); // 4
  // proto:  void QTimerEvent::QTimerEvent(int timerId);
extern void* C_ZN11QTimerEventC2Ei(int32_t arg0); // 3
  // proto:  bool QChildEvent::added();
extern bool C_ZNK11QChildEvent5addedEv(void* qthis); // 2
  // proto:  bool QChildEvent::polished();
extern bool C_ZNK11QChildEvent8polishedEv(void* qthis); // 2
  // proto:  void QChildEvent::~QChildEvent();
extern void C_ZN11QChildEventD2Ev(void* qthis); // 4
  // proto:  QObject * QChildEvent::child();
extern void* C_ZNK11QChildEvent5childEv(void* qthis); // 2
  // proto:  bool QChildEvent::removed();
extern bool C_ZNK11QChildEvent7removedEv(void* qthis); // 2
  // proto: static int QEvent::registerEventType(int hint);
extern int32_t C_ZN6QEvent17registerEventTypeEi(int32_t arg0); // 4
  // proto:  bool QEvent::spontaneous();
extern bool C_ZNK6QEvent11spontaneousEv(void* qthis); // 2
  // proto:  void QEvent::accept();
extern void C_ZN6QEvent6acceptEv(void* qthis); // 2
  // proto:  void QEvent::ignore();
extern void C_ZN6QEvent6ignoreEv(void* qthis); // 2
  // proto:  bool QEvent::isAccepted();
extern bool C_ZNK6QEvent10isAcceptedEv(void* qthis); // 2
  // proto:  void QEvent::setAccepted(bool accepted);
extern void C_ZN6QEvent11setAcceptedEb(void* qthis, bool arg0); // 2
  // proto:  QEvent::Type QEvent::type();
extern void C_ZNK6QEvent4typeEv(void* qthis); // 2
  // proto:  void QEvent::QEvent(const QEvent & other);
extern void* C_ZN6QEventC2ERKS_(void* arg0); // 3
  // proto:  void QEvent::~QEvent();
extern void C_ZN6QEventD2Ev(void* qthis); // 4
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

// class sizeof(QDeferredDeleteEvent)=24
type QDeferredDeleteEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QDynamicPropertyChangeEvent)=32
type QDynamicPropertyChangeEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTimerEvent)=24
type QTimerEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QChildEvent)=32
type QChildEvent struct {
  /*qbase*/ QEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QEvent)=24
type QEvent struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QDeferredDeleteEvent()
func NewQDeferredDeleteEvent(args ...interface{}) *QDeferredDeleteEvent {
  // QDeferredDeleteEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QDeferredDeleteEventC1Ev
    // invoke: void QDeferredDeleteEvent()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QDeferredDeleteEventC2Ev()
    return &QDeferredDeleteEvent{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "QDeferredDeleteEvent", args)
  }

  return nil // QDeferredDeleteEvent{qclsinst:qthis}
}

// ~QDeferredDeleteEvent()
func (this *QDeferredDeleteEvent) Freeqdeferreddeleteevent(args ...interface{}) () {
  // ~QDeferredDeleteEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QDeferredDeleteEventD0Ev
    // invoke: void ~QDeferredDeleteEvent()
    C.C_ZN20QDeferredDeleteEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "~QDeferredDeleteEvent", args)
  }

  return
}

// loopLevel()
func (this *QDeferredDeleteEvent) Looplevel(args ...interface{}) (ret interface{}) {
  // loopLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QDeferredDeleteEvent9loopLevelEv
    // invoke: int loopLevel()
    var ret0 = C.C_ZNK20QDeferredDeleteEvent9loopLevelEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDeferredDeleteEvent", "loopLevel", args)
  }

  return
}

// QDynamicPropertyChangeEvent(const class QByteArray &)
func NewQDynamicPropertyChangeEvent(args ...interface{}) *QDynamicPropertyChangeEvent {
  // QDynamicPropertyChangeEvent(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QDynamicPropertyChangeEventC1ERK10QByteArray
    // invoke: void QDynamicPropertyChangeEvent(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QDynamicPropertyChangeEventC2ERK10QByteArray(arg0)
    return &QDynamicPropertyChangeEvent{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "QDynamicPropertyChangeEvent", args)
  }

  return nil // QDynamicPropertyChangeEvent{qclsinst:qthis}
}

// propertyName()
func (this *QDynamicPropertyChangeEvent) Propertyname(args ...interface{}) (ret interface{}) {
  // propertyName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QDynamicPropertyChangeEvent12propertyNameEv
    // invoke: QByteArray propertyName()
    var ret0 = C.C_ZNK27QDynamicPropertyChangeEvent12propertyNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "propertyName", args)
  }

  return
}

// ~QDynamicPropertyChangeEvent()
func (this *QDynamicPropertyChangeEvent) Freeqdynamicpropertychangeevent(args ...interface{}) () {
  // ~QDynamicPropertyChangeEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QDynamicPropertyChangeEventD0Ev
    // invoke: void ~QDynamicPropertyChangeEvent()
    C.C_ZN27QDynamicPropertyChangeEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDynamicPropertyChangeEvent", "~QDynamicPropertyChangeEvent", args)
  }

  return
}

// timerId()
func (this *QTimerEvent) Timerid(args ...interface{}) (ret interface{}) {
  // timerId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTimerEvent7timerIdEv
    // invoke: int timerId()
    var ret0 = C.C_ZNK11QTimerEvent7timerIdEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTimerEvent", "timerId", args)
  }

  return
}

// ~QTimerEvent()
func (this *QTimerEvent) Freeqtimerevent(args ...interface{}) () {
  // ~QTimerEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTimerEventD0Ev
    // invoke: void ~QTimerEvent()
    C.C_ZN11QTimerEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTimerEvent", "~QTimerEvent", args)
  }

  return
}

// QTimerEvent(int)
func NewQTimerEvent(args ...interface{}) *QTimerEvent {
  // QTimerEvent(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTimerEventC1Ei
    // invoke: void QTimerEvent(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTimerEventC2Ei(arg0)
    return &QTimerEvent{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTimerEvent", "QTimerEvent", args)
  }

  return nil // QTimerEvent{qclsinst:qthis}
}

// added()
func (this *QChildEvent) Added(args ...interface{}) (ret interface{}) {
  // added()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent5addedEv
    // invoke: bool added()
    var ret0 = C.C_ZNK11QChildEvent5addedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QChildEvent", "added", args)
  }

  return
}

// polished()
func (this *QChildEvent) Polished(args ...interface{}) (ret interface{}) {
  // polished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent8polishedEv
    // invoke: bool polished()
    var ret0 = C.C_ZNK11QChildEvent8polishedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QChildEvent", "polished", args)
  }

  return
}

// ~QChildEvent()
func (this *QChildEvent) Freeqchildevent(args ...interface{}) () {
  // ~QChildEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QChildEventD0Ev
    // invoke: void ~QChildEvent()
    C.C_ZN11QChildEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QChildEvent", "~QChildEvent", args)
  }

  return
}

// child()
func (this *QChildEvent) Child(args ...interface{}) (ret interface{}) {
  // child()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent5childEv
    // invoke: QObject * child()
    var ret0 = C.C_ZNK11QChildEvent5childEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QChildEvent", "child", args)
  }

  return
}

// removed()
func (this *QChildEvent) Removed(args ...interface{}) (ret interface{}) {
  // removed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QChildEvent7removedEv
    // invoke: bool removed()
    var ret0 = C.C_ZNK11QChildEvent7removedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QChildEvent", "removed", args)
  }

  return
}

// registerEventType(int)
func (this *QEvent) Registereventtype_S(args ...interface{}) (ret interface{}) {
  // registerEventType(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent17registerEventTypeEi
    // invoke: int registerEventType(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN6QEvent17registerEventTypeEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QEvent", "registerEventType", args)
  }

  return
}

// spontaneous()
func (this *QEvent) Spontaneous(args ...interface{}) (ret interface{}) {
  // spontaneous()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QEvent11spontaneousEv
    // invoke: bool spontaneous()
    var ret0 = C.C_ZNK6QEvent11spontaneousEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QEvent", "spontaneous", args)
  }

  return
}

// accept()
func (this *QEvent) Accept(args ...interface{}) () {
  // accept()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent6acceptEv
    // invoke: void accept()
    C.C_ZN6QEvent6acceptEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "accept", args)
  }

  return
}

// ignore()
func (this *QEvent) Ignore(args ...interface{}) () {
  // ignore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent6ignoreEv
    // invoke: void ignore()
    C.C_ZN6QEvent6ignoreEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "ignore", args)
  }

  return
}

// isAccepted()
func (this *QEvent) Isaccepted(args ...interface{}) (ret interface{}) {
  // isAccepted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QEvent10isAcceptedEv
    // invoke: bool isAccepted()
    var ret0 = C.C_ZNK6QEvent10isAcceptedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QEvent", "isAccepted", args)
  }

  return
}

// setAccepted(_Bool)
func (this *QEvent) Setaccepted(args ...interface{}) () {
  // setAccepted(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEvent11setAcceptedEb
    // invoke: void setAccepted(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QEvent11setAcceptedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QEvent", "setAccepted", args)
  }

  return
}

// type()
func (this *QEvent) Type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QEvent4typeEv
    // invoke: QEvent::Type type()
    C.C_ZNK6QEvent4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "type", args)
  }

  return
}

// QEvent(const class QEvent &)
func NewQEvent(args ...interface{}) *QEvent {
  // QEvent(const class QEvent &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "const QEvent &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEventC1ERKS_
    // invoke: void QEvent(const class QEvent &)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QEventC2ERKS_(arg0)
    return &QEvent{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QEvent", "QEvent", args)
  }

  return nil // QEvent{qclsinst:qthis}
}

// ~QEvent()
func (this *QEvent) Freeqevent(args ...interface{}) () {
  // ~QEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QEventD0Ev
    // invoke: void ~QEvent()
    C.C_ZN6QEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QEvent", "~QEvent", args)
  }

  return
}

// <= body block end
