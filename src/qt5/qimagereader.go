package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qimagereader.h
// dst-file: /src/gui/qimagereader.go
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
  // proto:  void QImageReader::setAutoDetectImageFormat(bool enabled);
extern void C_ZN12QImageReader24setAutoDetectImageFormatEb(void* qthis, bool arg0); // 4
  // proto:  QString QImageReader::text(const QString & key);
extern void* C_ZNK12QImageReader4textERK7QString(void* qthis, void* arg0); // 4
  // proto:  QRect QImageReader::currentImageRect();
extern void* C_ZNK12QImageReader16currentImageRectEv(void* qthis); // 4
  // proto: static QList<QByteArray> QImageReader::supportedMimeTypes();
extern void C_ZN12QImageReader18supportedMimeTypesEv(); // 4
  // proto:  QByteArray QImageReader::subType();
extern void* C_ZNK12QImageReader7subTypeEv(void* qthis); // 4
  // proto:  void QImageReader::setDecideFormatFromContent(bool ignored);
extern void C_ZN12QImageReader26setDecideFormatFromContentEb(void* qthis, bool arg0); // 4
  // proto:  void QImageReader::setScaledSize(const QSize & size);
extern void C_ZN12QImageReader13setScaledSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QImageReader::setScaledClipRect(const QRect & rect);
extern void C_ZN12QImageReader17setScaledClipRectERK5QRect(void* qthis, void* arg0); // 4
  // proto:  bool QImageReader::canRead();
extern bool C_ZNK12QImageReader7canReadEv(void* qthis); // 4
  // proto:  QString QImageReader::fileName();
extern void* C_ZNK12QImageReader8fileNameEv(void* qthis); // 4
  // proto:  int QImageReader::quality();
extern int32_t C_ZNK12QImageReader7qualityEv(void* qthis); // 4
  // proto:  int QImageReader::imageCount();
extern int32_t C_ZNK12QImageReader10imageCountEv(void* qthis); // 4
  // proto:  QSize QImageReader::size();
extern void* C_ZNK12QImageReader4sizeEv(void* qthis); // 4
  // proto:  void QImageReader::QImageReader(const QString & fileName, const QByteArray & format);
extern void* C_ZN12QImageReaderC2ERK7QStringRK10QByteArray(void* arg0, void* arg1); // 3
  // proto:  void QImageReader::QImageReader(QIODevice * device, const QByteArray & format);
extern void* C_ZN12QImageReaderC2EP9QIODeviceRK10QByteArray(void* arg0, void* arg1); // 3
  // proto:  void QImageReader::QImageReader();
extern void* C_ZN12QImageReaderC2Ev(); // 3
  // proto:  int QImageReader::loopCount();
extern int32_t C_ZNK12QImageReader9loopCountEv(void* qthis); // 4
  // proto:  QRect QImageReader::scaledClipRect();
extern void* C_ZNK12QImageReader14scaledClipRectEv(void* qthis); // 4
  // proto:  bool QImageReader::autoDetectImageFormat();
extern bool C_ZNK12QImageReader21autoDetectImageFormatEv(void* qthis); // 4
  // proto:  bool QImageReader::jumpToNextImage();
extern bool C_ZN12QImageReader15jumpToNextImageEv(void* qthis); // 4
  // proto:  QStringList QImageReader::textKeys();
extern void C_ZNK12QImageReader8textKeysEv(void* qthis); // 4
  // proto:  QColor QImageReader::backgroundColor();
extern void* C_ZNK12QImageReader15backgroundColorEv(void* qthis); // 4
  // proto:  void QImageReader::setDevice(QIODevice * device);
extern void C_ZN12QImageReader9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  QImageIOHandler::Transformations QImageReader::transformation();
extern void C_ZNK12QImageReader14transformationEv(void* qthis); // 4
  // proto:  bool QImageReader::autoTransform();
extern bool C_ZNK12QImageReader13autoTransformEv(void* qthis); // 4
  // proto:  bool QImageReader::jumpToImage(int imageNumber);
extern bool C_ZN12QImageReader11jumpToImageEi(void* qthis, int32_t arg0); // 4
  // proto: static QList<QByteArray> QImageReader::supportedImageFormats();
extern void C_ZN12QImageReader21supportedImageFormatsEv(); // 4
  // proto:  void QImageReader::setAutoTransform(bool enabled);
extern void C_ZN12QImageReader16setAutoTransformEb(void* qthis, bool arg0); // 4
  // proto:  QRect QImageReader::clipRect();
extern void* C_ZNK12QImageReader8clipRectEv(void* qthis); // 4
  // proto:  QString QImageReader::errorString();
extern void* C_ZNK12QImageReader11errorStringEv(void* qthis); // 4
  // proto:  QByteArray QImageReader::format();
extern void* C_ZNK12QImageReader6formatEv(void* qthis); // 4
  // proto:  QImage QImageReader::read();
extern void* C_ZN12QImageReader4readEv(void* qthis); // 4
  // proto:  bool QImageReader::read(QImage * image);
extern bool C_ZN12QImageReader4readEP6QImage(void* qthis, void* arg0); // 4
  // proto:  void QImageReader::setBackgroundColor(const QColor & color);
extern void C_ZN12QImageReader18setBackgroundColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  int QImageReader::currentImageNumber();
extern int32_t C_ZNK12QImageReader18currentImageNumberEv(void* qthis); // 4
  // proto:  QIODevice * QImageReader::device();
extern void* C_ZNK12QImageReader6deviceEv(void* qthis); // 4
  // proto:  bool QImageReader::supportsAnimation();
extern bool C_ZNK12QImageReader17supportsAnimationEv(void* qthis); // 4
  // proto:  QList<QByteArray> QImageReader::supportedSubTypes();
extern void C_ZNK12QImageReader17supportedSubTypesEv(void* qthis); // 4
  // proto:  void QImageReader::setClipRect(const QRect & rect);
extern void C_ZN12QImageReader11setClipRectERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QImageReader::setFileName(const QString & fileName);
extern void C_ZN12QImageReader11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QImageReader::nextImageDelay();
extern int32_t C_ZNK12QImageReader14nextImageDelayEv(void* qthis); // 4
  // proto:  void QImageReader::~QImageReader();
extern void C_ZN12QImageReaderD2Ev(void* qthis); // 4
  // proto:  void QImageReader::setFormat(const QByteArray & format);
extern void C_ZN12QImageReader9setFormatERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QImageReader::setQuality(int quality);
extern void C_ZN12QImageReader10setQualityEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QImageReader::scaledSize();
extern void* C_ZNK12QImageReader10scaledSizeEv(void* qthis); // 4
  // proto: static QByteArray QImageReader::imageFormat(QIODevice * device);
extern void* C_ZN12QImageReader11imageFormatEP9QIODevice(void* arg0); // 4
  // proto: static QByteArray QImageReader::imageFormat(const QString & fileName);
extern void* C_ZN12QImageReader11imageFormatERK7QString(void* arg0); // 4
  // proto:  QImage::Format QImageReader::imageFormat();
extern void C_ZNK12QImageReader11imageFormatEv(void* qthis); // 4
  // proto:  QImageReader::ImageReaderError QImageReader::error();
extern void C_ZNK12QImageReader5errorEv(void* qthis); // 4
  // proto:  bool QImageReader::decideFormatFromContent();
extern bool C_ZNK12QImageReader23decideFormatFromContentEv(void* qthis); // 4
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

// class sizeof(QImageReader)=8
type QImageReader struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setAutoDetectImageFormat(_Bool)
func (this *QImageReader) Setautodetectimageformat(args ...interface{}) () {
  // setAutoDetectImageFormat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader24setAutoDetectImageFormatEb
    // invoke: void setAutoDetectImageFormat(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader24setAutoDetectImageFormatEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setAutoDetectImageFormat", args)
  }

  return
}

// text(const class QString &)
func (this *QImageReader) Text(args ...interface{}) (ret interface{}) {
  // text(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader4textERK7QString
    // invoke: QString text(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QImageReader4textERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "text", args)
  }

  return
}

// currentImageRect()
func (this *QImageReader) Currentimagerect(args ...interface{}) (ret interface{}) {
  // currentImageRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader16currentImageRectEv
    // invoke: QRect currentImageRect()
    var ret0 = C.C_ZNK12QImageReader16currentImageRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "currentImageRect", args)
  }

  return
}

// supportedMimeTypes()
func (this *QImageReader) Supportedmimetypes_S(args ...interface{}) () {
  // supportedMimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader18supportedMimeTypesEv
    // invoke: QList<QByteArray> supportedMimeTypes()
    C.C_ZN12QImageReader18supportedMimeTypesEv()
  default:
    qtrt.ErrorResolve("QImageReader", "supportedMimeTypes", args)
  }

  return
}

// subType()
func (this *QImageReader) Subtype(args ...interface{}) (ret interface{}) {
  // subType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader7subTypeEv
    // invoke: QByteArray subType()
    var ret0 = C.C_ZNK12QImageReader7subTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "subType", args)
  }

  return
}

// setDecideFormatFromContent(_Bool)
func (this *QImageReader) Setdecideformatfromcontent(args ...interface{}) () {
  // setDecideFormatFromContent(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader26setDecideFormatFromContentEb
    // invoke: void setDecideFormatFromContent(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader26setDecideFormatFromContentEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setDecideFormatFromContent", args)
  }

  return
}

// setScaledSize(const class QSize &)
func (this *QImageReader) Setscaledsize(args ...interface{}) () {
  // setScaledSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader13setScaledSizeERK5QSize
    // invoke: void setScaledSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader13setScaledSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setScaledSize", args)
  }

  return
}

// setScaledClipRect(const class QRect &)
func (this *QImageReader) Setscaledcliprect(args ...interface{}) () {
  // setScaledClipRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader17setScaledClipRectERK5QRect
    // invoke: void setScaledClipRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader17setScaledClipRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setScaledClipRect", args)
  }

  return
}

// canRead()
func (this *QImageReader) Canread(args ...interface{}) (ret interface{}) {
  // canRead()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader7canReadEv
    // invoke: bool canRead()
    var ret0 = C.C_ZNK12QImageReader7canReadEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "canRead", args)
  }

  return
}

// fileName()
func (this *QImageReader) Filename(args ...interface{}) (ret interface{}) {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader8fileNameEv
    // invoke: QString fileName()
    var ret0 = C.C_ZNK12QImageReader8fileNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "fileName", args)
  }

  return
}

// quality()
func (this *QImageReader) Quality(args ...interface{}) (ret interface{}) {
  // quality()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader7qualityEv
    // invoke: int quality()
    var ret0 = C.C_ZNK12QImageReader7qualityEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "quality", args)
  }

  return
}

// imageCount()
func (this *QImageReader) Imagecount(args ...interface{}) (ret interface{}) {
  // imageCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader10imageCountEv
    // invoke: int imageCount()
    var ret0 = C.C_ZNK12QImageReader10imageCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "imageCount", args)
  }

  return
}

// size()
func (this *QImageReader) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader4sizeEv
    // invoke: QSize size()
    var ret0 = C.C_ZNK12QImageReader4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "size", args)
  }

  return
}

// QImageReader(const class QString &, const class QByteArray &)
func NewQImageReader(args ...interface{}) *QImageReader {
  // QImageReader(const class QString &, const class QByteArray &)
  // QImageReader(class QIODevice *, const class QByteArray &)
  // QImageReader()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReaderC1ERK7QStringRK10QByteArray
    // invoke: void QImageReader(const class QString &, const class QByteArray &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QImageReaderC2ERK7QStringRK10QByteArray(arg0, arg1)
    return &QImageReader{qclsinst:qthis}
  case 1:
    // invoke: _ZN12QImageReaderC1EP9QIODeviceRK10QByteArray
    // invoke: void QImageReader(class QIODevice *, const class QByteArray &)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QImageReaderC2EP9QIODeviceRK10QByteArray(arg0, arg1)
    return &QImageReader{qclsinst:qthis}
  case 2:
    // invoke: _ZN12QImageReaderC1Ev
    // invoke: void QImageReader()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QImageReaderC2Ev()
    return &QImageReader{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QImageReader", "QImageReader", args)
  }

  return nil // QImageReader{qclsinst:qthis}
}

// loopCount()
func (this *QImageReader) Loopcount(args ...interface{}) (ret interface{}) {
  // loopCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader9loopCountEv
    // invoke: int loopCount()
    var ret0 = C.C_ZNK12QImageReader9loopCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "loopCount", args)
  }

  return
}

// scaledClipRect()
func (this *QImageReader) Scaledcliprect(args ...interface{}) (ret interface{}) {
  // scaledClipRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader14scaledClipRectEv
    // invoke: QRect scaledClipRect()
    var ret0 = C.C_ZNK12QImageReader14scaledClipRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "scaledClipRect", args)
  }

  return
}

// autoDetectImageFormat()
func (this *QImageReader) Autodetectimageformat(args ...interface{}) (ret interface{}) {
  // autoDetectImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader21autoDetectImageFormatEv
    // invoke: bool autoDetectImageFormat()
    var ret0 = C.C_ZNK12QImageReader21autoDetectImageFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "autoDetectImageFormat", args)
  }

  return
}

// jumpToNextImage()
func (this *QImageReader) Jumptonextimage(args ...interface{}) (ret interface{}) {
  // jumpToNextImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader15jumpToNextImageEv
    // invoke: bool jumpToNextImage()
    var ret0 = C.C_ZN12QImageReader15jumpToNextImageEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "jumpToNextImage", args)
  }

  return
}

// textKeys()
func (this *QImageReader) Textkeys(args ...interface{}) () {
  // textKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader8textKeysEv
    // invoke: QStringList textKeys()
    C.C_ZNK12QImageReader8textKeysEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageReader", "textKeys", args)
  }

  return
}

// backgroundColor()
func (this *QImageReader) Backgroundcolor(args ...interface{}) (ret interface{}) {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader15backgroundColorEv
    // invoke: QColor backgroundColor()
    var ret0 = C.C_ZNK12QImageReader15backgroundColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "backgroundColor", args)
  }

  return
}

// setDevice(class QIODevice *)
func (this *QImageReader) Setdevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader9setDeviceEP9QIODevice
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader9setDeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setDevice", args)
  }

  return
}

// transformation()
func (this *QImageReader) Transformation(args ...interface{}) () {
  // transformation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader14transformationEv
    // invoke: QImageIOHandler::Transformations transformation()
    C.C_ZNK12QImageReader14transformationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageReader", "transformation", args)
  }

  return
}

// autoTransform()
func (this *QImageReader) Autotransform(args ...interface{}) (ret interface{}) {
  // autoTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader13autoTransformEv
    // invoke: bool autoTransform()
    var ret0 = C.C_ZNK12QImageReader13autoTransformEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "autoTransform", args)
  }

  return
}

// jumpToImage(int)
func (this *QImageReader) Jumptoimage(args ...interface{}) (ret interface{}) {
  // jumpToImage(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader11jumpToImageEi
    // invoke: bool jumpToImage(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QImageReader11jumpToImageEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "jumpToImage", args)
  }

  return
}

// supportedImageFormats()
func (this *QImageReader) Supportedimageformats_S(args ...interface{}) () {
  // supportedImageFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader21supportedImageFormatsEv
    // invoke: QList<QByteArray> supportedImageFormats()
    C.C_ZN12QImageReader21supportedImageFormatsEv()
  default:
    qtrt.ErrorResolve("QImageReader", "supportedImageFormats", args)
  }

  return
}

// setAutoTransform(_Bool)
func (this *QImageReader) Setautotransform(args ...interface{}) () {
  // setAutoTransform(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader16setAutoTransformEb
    // invoke: void setAutoTransform(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader16setAutoTransformEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setAutoTransform", args)
  }

  return
}

// clipRect()
func (this *QImageReader) Cliprect(args ...interface{}) (ret interface{}) {
  // clipRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader8clipRectEv
    // invoke: QRect clipRect()
    var ret0 = C.C_ZNK12QImageReader8clipRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "clipRect", args)
  }

  return
}

// errorString()
func (this *QImageReader) Errorstring(args ...interface{}) (ret interface{}) {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader11errorStringEv
    // invoke: QString errorString()
    var ret0 = C.C_ZNK12QImageReader11errorStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "errorString", args)
  }

  return
}

// format()
func (this *QImageReader) Format(args ...interface{}) (ret interface{}) {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader6formatEv
    // invoke: QByteArray format()
    var ret0 = C.C_ZNK12QImageReader6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "format", args)
  }

  return
}

// read()
func (this *QImageReader) Read(args ...interface{}) (ret interface{}) {
  // read()
  // read(class QImage *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QImage{}) // "QImage *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader4readEv
    // invoke: QImage read()
    var ret0 = C.C_ZN12QImageReader4readEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN12QImageReader4readEP6QImage
    // invoke: bool read(class QImage *)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QImageReader4readEP6QImage(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "read", args)
  }

  return
}

// setBackgroundColor(const class QColor &)
func (this *QImageReader) Setbackgroundcolor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader18setBackgroundColorERK6QColor
    // invoke: void setBackgroundColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader18setBackgroundColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setBackgroundColor", args)
  }

  return
}

// currentImageNumber()
func (this *QImageReader) Currentimagenumber(args ...interface{}) (ret interface{}) {
  // currentImageNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader18currentImageNumberEv
    // invoke: int currentImageNumber()
    var ret0 = C.C_ZNK12QImageReader18currentImageNumberEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "currentImageNumber", args)
  }

  return
}

// device()
func (this *QImageReader) Device(args ...interface{}) (ret interface{}) {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader6deviceEv
    // invoke: QIODevice * device()
    var ret0 = C.C_ZNK12QImageReader6deviceEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIODevice{}) // "QIODevice *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "device", args)
  }

  return
}

// supportsAnimation()
func (this *QImageReader) Supportsanimation(args ...interface{}) (ret interface{}) {
  // supportsAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader17supportsAnimationEv
    // invoke: bool supportsAnimation()
    var ret0 = C.C_ZNK12QImageReader17supportsAnimationEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "supportsAnimation", args)
  }

  return
}

// supportedSubTypes()
func (this *QImageReader) Supportedsubtypes(args ...interface{}) () {
  // supportedSubTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader17supportedSubTypesEv
    // invoke: QList<QByteArray> supportedSubTypes()
    C.C_ZNK12QImageReader17supportedSubTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageReader", "supportedSubTypes", args)
  }

  return
}

// setClipRect(const class QRect &)
func (this *QImageReader) Setcliprect(args ...interface{}) () {
  // setClipRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader11setClipRectERK5QRect
    // invoke: void setClipRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader11setClipRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setClipRect", args)
  }

  return
}

// setFileName(const class QString &)
func (this *QImageReader) Setfilename(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader11setFileNameERK7QString
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setFileName", args)
  }

  return
}

// nextImageDelay()
func (this *QImageReader) Nextimagedelay(args ...interface{}) (ret interface{}) {
  // nextImageDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader14nextImageDelayEv
    // invoke: int nextImageDelay()
    var ret0 = C.C_ZNK12QImageReader14nextImageDelayEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "nextImageDelay", args)
  }

  return
}

// ~QImageReader()
func (this *QImageReader) Freeqimagereader(args ...interface{}) () {
  // ~QImageReader()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReaderD0Ev
    // invoke: void ~QImageReader()
    C.C_ZN12QImageReaderD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageReader", "~QImageReader", args)
  }

  return
}

// setFormat(const class QByteArray &)
func (this *QImageReader) Setformat(args ...interface{}) () {
  // setFormat(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader9setFormatERK10QByteArray
    // invoke: void setFormat(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader9setFormatERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setFormat", args)
  }

  return
}

// setQuality(int)
func (this *QImageReader) Setquality(args ...interface{}) () {
  // setQuality(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader10setQualityEi
    // invoke: void setQuality(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageReader10setQualityEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageReader", "setQuality", args)
  }

  return
}

// scaledSize()
func (this *QImageReader) Scaledsize(args ...interface{}) (ret interface{}) {
  // scaledSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader10scaledSizeEv
    // invoke: QSize scaledSize()
    var ret0 = C.C_ZNK12QImageReader10scaledSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "scaledSize", args)
  }

  return
}

// imageFormat(class QIODevice *)
func (this *QImageReader) Imageformat_S(args ...interface{}) (ret interface{}) {
  // imageFormat(class QIODevice *)
  // imageFormat(const class QString &)
  // imageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader11imageFormatEP9QIODevice
    // invoke: QByteArray imageFormat(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QImageReader11imageFormatEP9QIODevice(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN12QImageReader11imageFormatERK7QString
    // invoke: QByteArray imageFormat(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QImageReader11imageFormatERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 2:
    // invoke: _ZNK12QImageReader11imageFormatEv
    // invoke: QImage::Format imageFormat()
    C.C_ZNK12QImageReader11imageFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageReader", "imageFormat", args)
  }

  return
}

// error()
func (this *QImageReader) Error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader5errorEv
    // invoke: QImageReader::ImageReaderError error()
    C.C_ZNK12QImageReader5errorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageReader", "error", args)
  }

  return
}

// decideFormatFromContent()
func (this *QImageReader) Decideformatfromcontent(args ...interface{}) (ret interface{}) {
  // decideFormatFromContent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader23decideFormatFromContentEv
    // invoke: bool decideFormatFromContent()
    var ret0 = C.C_ZNK12QImageReader23decideFormatFromContentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QImageReader", "decideFormatFromContent", args)
  }

  return
}

// <= body block end
