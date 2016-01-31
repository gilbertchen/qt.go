package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qfuturewatcher.h
// dst-file: /src/core/qfuturewatcher.go
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
  // proto:  bool QFutureWatcherBase::isStarted();
extern bool C_ZNK18QFutureWatcherBase9isStartedEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::cancel();
extern void C_ZN18QFutureWatcherBase6cancelEv(void* qthis); // 4
  // proto:  bool QFutureWatcherBase::event(QEvent * event);
extern bool C_ZN18QFutureWatcherBase5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QFutureWatcherBase::togglePaused();
extern void C_ZN18QFutureWatcherBase12togglePausedEv(void* qthis); // 4
  // proto:  int QFutureWatcherBase::progressMinimum();
extern int32_t C_ZNK18QFutureWatcherBase15progressMinimumEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::waitForFinished();
extern void C_ZN18QFutureWatcherBase15waitForFinishedEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::resume();
extern void C_ZN18QFutureWatcherBase6resumeEv(void* qthis); // 4
  // proto:  int QFutureWatcherBase::progressMaximum();
extern int32_t C_ZNK18QFutureWatcherBase15progressMaximumEv(void* qthis); // 4
  // proto:  bool QFutureWatcherBase::isFinished();
extern bool C_ZNK18QFutureWatcherBase10isFinishedEv(void* qthis); // 4
  // proto:  QString QFutureWatcherBase::progressText();
extern void* C_ZNK18QFutureWatcherBase12progressTextEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::setPendingResultsLimit(int limit);
extern void C_ZN18QFutureWatcherBase22setPendingResultsLimitEi(void* qthis, int32_t arg0); // 4
  // proto:  void QFutureWatcherBase::pause();
extern void C_ZN18QFutureWatcherBase5pauseEv(void* qthis); // 4
  // proto:  const QMetaObject * QFutureWatcherBase::metaObject();
extern void C_ZNK18QFutureWatcherBase10metaObjectEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::QFutureWatcherBase(QObject * parent);
extern void* C_ZN18QFutureWatcherBaseC2EP7QObject(void* arg0); // 3
  // proto:  bool QFutureWatcherBase::isRunning();
extern bool C_ZNK18QFutureWatcherBase9isRunningEv(void* qthis); // 4
  // proto:  bool QFutureWatcherBase::isCanceled();
extern bool C_ZNK18QFutureWatcherBase10isCanceledEv(void* qthis); // 4
  // proto:  bool QFutureWatcherBase::isPaused();
extern bool C_ZNK18QFutureWatcherBase8isPausedEv(void* qthis); // 4
  // proto:  int QFutureWatcherBase::progressValue();
extern int32_t C_ZNK18QFutureWatcherBase13progressValueEv(void* qthis); // 4
  // proto:  void QFutureWatcherBase::setPaused(bool paused);
extern void C_ZN18QFutureWatcherBase9setPausedEb(void* qthis, bool arg0); // 4
  // proto:  void QFutureWatcher<void>::~QFutureWatcher();
extern void C_ZN14QFutureWatcherIvED2Ev(void* qthis); // 4
  // proto:  void QFutureWatcher<void>::QFutureWatcher(QObject * _parent);
extern void* C_ZN14QFutureWatcherIvEC2EP7QObject(void* arg0); // 1
  // proto:  QFuture<void> QFutureWatcher<void>::future();
extern void C_ZNK14QFutureWatcherIvE6futureEv(void* qthis); // 2
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

// class sizeof(QFutureWatcherBase)=1
type QFutureWatcherBase struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _progressRangeChanged QFutureWatcherBase_progressRangeChanged_signal;
//  _resumed QFutureWatcherBase_resumed_signal;
//  _progressValueChanged QFutureWatcherBase_progressValueChanged_signal;
//  _started QFutureWatcherBase_started_signal;
//  _resultReadyAt QFutureWatcherBase_resultReadyAt_signal;
//  _resultsReadyAt QFutureWatcherBase_resultsReadyAt_signal;
//  _paused QFutureWatcherBase_paused_signal;
//  _canceled QFutureWatcherBase_canceled_signal;
//  _finished QFutureWatcherBase_finished_signal;
//  _progressTextChanged QFutureWatcherBase_progressTextChanged_signal;
}

// class sizeof(QFutureWatcher<void>)=1
type QFutureWatcherLvoidG struct {
  /*qbase*/ QFutureWatcherBase;
  qclsinst unsafe.Pointer /* *C.void */;
}

// isStarted()
func (this *QFutureWatcherBase) Isstarted(args ...interface{}) (ret interface{}) {
  // isStarted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase9isStartedEv
    // invoke: bool isStarted()
    var ret0 = C.C_ZNK18QFutureWatcherBase9isStartedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isStarted", args)
  }

  return
}

// cancel()
func (this *QFutureWatcherBase) Cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase6cancelEv
    // invoke: void cancel()
    C.C_ZN18QFutureWatcherBase6cancelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "cancel", args)
  }

  return
}

// event(class QEvent *)
func (this *QFutureWatcherBase) Event(args ...interface{}) (ret interface{}) {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QFutureWatcherBase5eventEP6QEvent(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "event", args)
  }

  return
}

// togglePaused()
func (this *QFutureWatcherBase) Togglepaused(args ...interface{}) () {
  // togglePaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase12togglePausedEv
    // invoke: void togglePaused()
    C.C_ZN18QFutureWatcherBase12togglePausedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "togglePaused", args)
  }

  return
}

// progressMinimum()
func (this *QFutureWatcherBase) Progressminimum(args ...interface{}) (ret interface{}) {
  // progressMinimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase15progressMinimumEv
    // invoke: int progressMinimum()
    var ret0 = C.C_ZNK18QFutureWatcherBase15progressMinimumEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressMinimum", args)
  }

  return
}

// waitForFinished()
func (this *QFutureWatcherBase) Waitforfinished(args ...interface{}) () {
  // waitForFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase15waitForFinishedEv
    // invoke: void waitForFinished()
    C.C_ZN18QFutureWatcherBase15waitForFinishedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "waitForFinished", args)
  }

  return
}

// resume()
func (this *QFutureWatcherBase) Resume(args ...interface{}) () {
  // resume()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase6resumeEv
    // invoke: void resume()
    C.C_ZN18QFutureWatcherBase6resumeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "resume", args)
  }

  return
}

// progressMaximum()
func (this *QFutureWatcherBase) Progressmaximum(args ...interface{}) (ret interface{}) {
  // progressMaximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase15progressMaximumEv
    // invoke: int progressMaximum()
    var ret0 = C.C_ZNK18QFutureWatcherBase15progressMaximumEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressMaximum", args)
  }

  return
}

// isFinished()
func (this *QFutureWatcherBase) Isfinished(args ...interface{}) (ret interface{}) {
  // isFinished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10isFinishedEv
    // invoke: bool isFinished()
    var ret0 = C.C_ZNK18QFutureWatcherBase10isFinishedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isFinished", args)
  }

  return
}

// progressText()
func (this *QFutureWatcherBase) Progresstext(args ...interface{}) (ret interface{}) {
  // progressText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase12progressTextEv
    // invoke: QString progressText()
    var ret0 = C.C_ZNK18QFutureWatcherBase12progressTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressText", args)
  }

  return
}

// setPendingResultsLimit(int)
func (this *QFutureWatcherBase) Setpendingresultslimit(args ...interface{}) () {
  // setPendingResultsLimit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase22setPendingResultsLimitEi
    // invoke: void setPendingResultsLimit(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QFutureWatcherBase22setPendingResultsLimitEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "setPendingResultsLimit", args)
  }

  return
}

// pause()
func (this *QFutureWatcherBase) Pause(args ...interface{}) () {
  // pause()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase5pauseEv
    // invoke: void pause()
    C.C_ZN18QFutureWatcherBase5pauseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "pause", args)
  }

  return
}

// metaObject()
func (this *QFutureWatcherBase) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QFutureWatcherBase10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "metaObject", args)
  }

  return
}

// QFutureWatcherBase(class QObject *)
func NewQFutureWatcherBase(args ...interface{}) *QFutureWatcherBase {
  // QFutureWatcherBase(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBaseC1EP7QObject
    // invoke: void QFutureWatcherBase(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QFutureWatcherBaseC2EP7QObject(arg0)
    return &QFutureWatcherBase{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "QFutureWatcherBase", args)
  }

  return nil // QFutureWatcherBase{qclsinst:qthis}
}

// isRunning()
func (this *QFutureWatcherBase) Isrunning(args ...interface{}) (ret interface{}) {
  // isRunning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase9isRunningEv
    // invoke: bool isRunning()
    var ret0 = C.C_ZNK18QFutureWatcherBase9isRunningEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isRunning", args)
  }

  return
}

// isCanceled()
func (this *QFutureWatcherBase) Iscanceled(args ...interface{}) (ret interface{}) {
  // isCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase10isCanceledEv
    // invoke: bool isCanceled()
    var ret0 = C.C_ZNK18QFutureWatcherBase10isCanceledEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isCanceled", args)
  }

  return
}

// isPaused()
func (this *QFutureWatcherBase) Ispaused(args ...interface{}) (ret interface{}) {
  // isPaused()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase8isPausedEv
    // invoke: bool isPaused()
    var ret0 = C.C_ZNK18QFutureWatcherBase8isPausedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "isPaused", args)
  }

  return
}

// progressValue()
func (this *QFutureWatcherBase) Progressvalue(args ...interface{}) (ret interface{}) {
  // progressValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QFutureWatcherBase13progressValueEv
    // invoke: int progressValue()
    var ret0 = C.C_ZNK18QFutureWatcherBase13progressValueEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "progressValue", args)
  }

  return
}

// setPaused(_Bool)
func (this *QFutureWatcherBase) Setpaused(args ...interface{}) () {
  // setPaused(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QFutureWatcherBase9setPausedEb
    // invoke: void setPaused(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN18QFutureWatcherBase9setPausedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFutureWatcherBase", "setPaused", args)
  }

  return
}

// ~QFutureWatcher()
func (this *QFutureWatcherLvoidG) Freeqfuturewatcherlvoidg(args ...interface{}) () {
  // ~QFutureWatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QFutureWatcherIvED0Ev
    // invoke: void ~QFutureWatcher()
    C.C_ZN14QFutureWatcherIvED2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcher<void>", "~QFutureWatcher", args)
  }

  return
}

// QFutureWatcher(class QObject *)
func NewQFutureWatcherLvoidG(args ...interface{}) *QFutureWatcherLvoidG {
  // QFutureWatcher(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QFutureWatcherIvEC1EP7QObject
    // invoke: void QFutureWatcher(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QFutureWatcherIvEC2EP7QObject(arg0)
    return &QFutureWatcherLvoidG{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFutureWatcher<void>", "QFutureWatcher", args)
  }

  return nil // QFutureWatcherLvoidG{qclsinst:qthis}
}

// future()
func (this *QFutureWatcherLvoidG) Future(args ...interface{}) () {
  // future()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QFutureWatcherIvE6futureEv
    // invoke: QFuture<void> future()
    C.C_ZNK14QFutureWatcherIvE6futureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFutureWatcher<void>", "future", args)
  }

  return
}

// <= body block end
