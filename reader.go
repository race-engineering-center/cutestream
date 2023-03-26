package cutestream

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"time"
	"unicode/utf16"
)

// QMetaType represents a Qt metatype.
type QMetaType int

// From qtbase/corelib/kernel/qmetatype.h.
const (
	QMetaTypeBool               QMetaType = 1
	QMetaTypeInt                QMetaType = 2
	QMetaTypeUInt               QMetaType = 3
	QMetaTypeLongLong           QMetaType = 4
	QMetaTypeULongLong          QMetaType = 5
	QMetaTypeDouble             QMetaType = 6
	QMetaTypeQChar              QMetaType = 7
	QMetaTypeQVariantMap        QMetaType = 8
	QMetaTypeQVariantList       QMetaType = 9
	QMetaTypeQString            QMetaType = 10
	QMetaTypeQStringList        QMetaType = 11
	QMetaTypeQByteArray         QMetaType = 12
	QMetaTypeQBitArray          QMetaType = 13
	QMetaTypeQDate              QMetaType = 14
	QMetaTypeQTime              QMetaType = 15
	QMetaTypeQDateTime          QMetaType = 16
	QMetaTypeQUrl               QMetaType = 17
	QMetaTypeQLocale            QMetaType = 18
	QMetaTypeQRect              QMetaType = 19
	QMetaTypeQRectF             QMetaType = 20
	QMetaTypeQSize              QMetaType = 21
	QMetaTypeQSizeF             QMetaType = 22
	QMetaTypeQLine              QMetaType = 23
	QMetaTypeQLineF             QMetaType = 24
	QMetaTypeQPoint             QMetaType = 25
	QMetaTypeQPointF            QMetaType = 26
	QMetaTypeQRegExp            QMetaType = 27
	QMetaTypeQVariantHash       QMetaType = 28
	QMetaTypeQEasingCurve       QMetaType = 29
	QMetaTypeQUuid              QMetaType = 30
	QMetaTypeVoidStar           QMetaType = 31
	QMetaTypeLong               QMetaType = 32
	QMetaTypeShort              QMetaType = 33
	QMetaTypeChar               QMetaType = 34
	QMetaTypeULong              QMetaType = 35
	QMetaTypeUShort             QMetaType = 36
	QMetaTypeUChar              QMetaType = 37
	QMetaTypeFloat              QMetaType = 38
	QMetaTypeQObjectStar        QMetaType = 39
	QMetaTypeSChar              QMetaType = 40
	QMetaTypeVoid               QMetaType = 43
	QMetaTypeQVariant           QMetaType = 41
	QMetaTypeQModelIndex        QMetaType = 42
	QMetaTypeQRegularExpression QMetaType = 44
	QMetaTypeQJsonValue         QMetaType = 45
	QMetaTypeQJsonObject        QMetaType = 46
	QMetaTypeQJsonArray         QMetaType = 47
	QMetaTypeQJsonDocument      QMetaType = 48
	QMetaTypeQFont              QMetaType = 64
	QMetaTypeQPixmap            QMetaType = 65
	QMetaTypeQBrush             QMetaType = 66
	QMetaTypeQColor             QMetaType = 67
	QMetaTypeQPalette           QMetaType = 68
	QMetaTypeQIcon              QMetaType = 69
	QMetaTypeQImage             QMetaType = 70
	QMetaTypeQPolygon           QMetaType = 71
	QMetaTypeQRegion            QMetaType = 72
	QMetaTypeQBitmap            QMetaType = 73
	QMetaTypeQCursor            QMetaType = 74
	QMetaTypeQKeySequence       QMetaType = 75
	QMetaTypeQPen               QMetaType = 76
	QMetaTypeQTextLength        QMetaType = 77
	QMetaTypeQTextFormat        QMetaType = 78
	QMetaTypeQMatrix            QMetaType = 79
	QMetaTypeQTransform         QMetaType = 80
	QMetaTypeQMatrix4x4         QMetaType = 81
	QMetaTypeQVector2D          QMetaType = 82
	QMetaTypeQVector3D          QMetaType = 83
	QMetaTypeQVector4D          QMetaType = 84
	QMetaTypeQQuaternion        QMetaType = 85
	QMetaTypeQPolygonF          QMetaType = 86
	QMetaTypeQSizePolicy        QMetaType = 121
)

type Reader struct {
	Reader          io.Reader
	ByteOrder       binary.ByteOrder
	version         int
	DoublePrecision bool // Use Double precision for floats. Set to `false` to use Single precision
}

// NewReader creates a new Reader object with the specified underlying reader,
// big endian byte order and enabled double precision
func NewReader(reader io.Reader) Reader {
	return Reader{
		Reader:          reader,
		ByteOrder:       binary.BigEndian,
		version:         19,
		DoublePrecision: false,
	}
}

func NewReaderWithVersion(reader io.Reader, version int) (Reader, error) {
	r := Reader{
		Reader:          reader,
		ByteOrder:       binary.BigEndian,
		version:         19,
		DoublePrecision: false,
	}
	err := r.SetVersion(version)
	if err != nil {
		return Reader{}, err
	}
	return r, nil
}

func (r *Reader) SetVersion(version int) error {
	acceptedVersions := []int{19, 20}
	found := false
	for _, v := range acceptedVersions {
		if v == version {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("%d is not a supported version, %v", version, acceptedVersions)
	}
	r.version = version
	return nil
}

func (r *Reader) ReadBool() (bool, error) {
	var v uint8
	if err := binary.Read(r.Reader, r.ByteOrder, &v); err != nil {
		return false, err
	}
	return v != 0, nil
}

func ReadNumber[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](reader *Reader) (T, error) {
	var v T
	if err := binary.Read(reader.Reader, reader.ByteOrder, &v); err != nil {
		return 0, err
	}
	return v, nil
}

func (r *Reader) ReadInt8() (int8, error) {
	return ReadNumber[int8](r)
}

func (r *Reader) ReadInt16() (int16, error) {
	return ReadNumber[int16](r)
}

func (r *Reader) ReadInt32() (int32, error) {
	return ReadNumber[int32](r)
}

func (r *Reader) ReadInt64() (int64, error) {
	return ReadNumber[int64](r)
}

func (r *Reader) ReadUint8() (uint8, error) {
	return ReadNumber[uint8](r)
}

func (r *Reader) ReadUint16() (uint16, error) {
	return ReadNumber[uint16](r)
}

func (r *Reader) ReadUint32() (uint32, error) {
	return ReadNumber[uint32](r)
}

func (r *Reader) ReadUint64() (uint64, error) {
	return ReadNumber[uint64](r)
}

func (r *Reader) ReadFloat() (float32, error) {
	if r.DoublePrecision {
		val, err := ReadNumber[float64](r)
		if err != nil {
			return 0, err
		}
		return float32(val), nil
	}
	return ReadNumber[float32](r)
}

func (r *Reader) ReadDouble() (float64, error) {
	if !r.DoublePrecision {
		val, err := ReadNumber[float32](r)
		if err != nil {
			return 0, err
		}
		return float64(val), nil
	}
	return ReadNumber[float64](r)
}

func (r *Reader) ReadCString() (string, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return "", err
	}
	buf := make([]byte, n)
	if err := binary.Read(r.Reader, r.ByteOrder, &buf); err != nil {
		return "", err
	}
	return string(buf), nil
}

func (r *Reader) ReadQBitArray() ([]bool, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}
	buf := make([]byte, (n+7)/8)
	if err := binary.Read(r.Reader, r.ByteOrder, &buf); err != nil {
		return nil, err
	}
	bits := make([]bool, n)
	for i := n - 1; i >= 0; i-- {
		bits[i] = (((buf[i/8]) >> (7 - i%8)) & 0x1) == 0x1
	}
	return bits, nil
}

func (r *Reader) ReadQByteArray() ([]byte, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}
	if n == 0xFFFFFFFF {
		return nil, nil
	}
	buf := make([]byte, n)
	if err := binary.Read(r.Reader, r.ByteOrder, &buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func (r *Reader) ReadQDate() (time.Time, error) {
	julian, err := r.ReadUint64()
	if err != nil {
		return time.Time{}, err
	}
	// ported from qdatetime.cpp
	floordiv := func(a, b int) int {
		var x int
		if a < 0 {
			x = b - 1
		}
		return (a - x) / b
	}
	var a, b int64
	a = int64(julian) + 32044
	b = int64(floordiv(int(4*a+3), 146097))
	var c, d, e, m, day, month, year int
	c = int(a) - floordiv(146097*int(b), 4)
	d = floordiv(4*c+3, 1461)
	e = c - floordiv(1461*d, 4)
	m = floordiv(5*e+2, 153)
	day = e - floordiv(153*m+2, 5) + 1
	month = m + 3 - 12*floordiv(m, 10)
	year = 100*int(b) + d - 4800 + floordiv(m, 10)
	if year <= 0 {
		year--
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func (r *Reader) ReadQString() (string, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return "", err
	}
	if n == 0xFFFFFFFF {
		return "", nil
	}
	buf := make([]uint16, n/2)
	if err := binary.Read(r.Reader, r.ByteOrder, &buf); err != nil {
		return "", err
	}
	return string(utf16.Decode(buf)), nil
}

func (r *Reader) ReadQTime() (time.Duration, error) { // msecs past midnight
	msecsMidnight, err := r.ReadUint32()
	if err != nil {
		return 0, err
	}
	return time.Millisecond * time.Duration(msecsMidnight), nil
}

func (r *Reader) ReadQUrl() (*url.URL, error) { // msecs past midnight
	str, err := r.ReadQString()
	if err != nil {
		return nil, err
	}
	if str == "" {
		return nil, nil
	}
	return url.Parse(str)
}

func (r *Reader) ReadQVariant() (QMetaType, interface{}, error) { // msecs past midnight
	t, err := r.ReadUint32()
	if err != nil {
		return 0, nil, err
	}
	if null, err := r.ReadBool(); err != nil {
		return 0, nil, err
	} else if null {
		return QMetaType(t), nil, nil
	}

	var v interface{}
	err = nil
	switch QMetaType(t) {
	case QMetaTypeBool:
		v, err = r.ReadBool()
	case QMetaTypeInt:
		v, err = r.ReadInt32()
	case QMetaTypeUInt:
		v, err = r.ReadUint32()
	case QMetaTypeLongLong:
		v, err = r.ReadInt64()
	case QMetaTypeULongLong:
		v, err = r.ReadUint64()
	case QMetaTypeDouble:
		v, err = r.ReadDouble()
	case QMetaTypeFloat:
		v, err = r.ReadFloat()
	case QMetaTypeQChar, QMetaTypeChar:
		v, err = r.ReadUint8()
	case QMetaTypeSChar:
		v, err = r.ReadInt8()
	case QMetaTypeShort:
		v, err = r.ReadInt16()
	case QMetaTypeUShort:
		v, err = r.ReadUint16()
	case QMetaTypeQBitArray:
		v, err = r.ReadQBitArray()
	case QMetaTypeQVariantMap, QMetaTypeQVariantHash:
		v, err = r.ReadQStringQVariantAssociative()
	case QMetaTypeQUuid:
		v, err = r.ReadQUuid()
	case QMetaTypeQVariantList:
		v, err = r.ReadQStringQVariantList()
	case QMetaTypeQByteArray:
		v, err = r.ReadQByteArray()
	case QMetaTypeQString:
		v, err = r.ReadQString()
	case QMetaTypeQStringList:
		v, err = r.ReadQStringQStringList()
	case QMetaTypeQDate:
		v, err = r.ReadQDate()
	case QMetaTypeQTime:
		v, err = r.ReadQTime()
	case QMetaTypeQDateTime:
		v, err = r.ReadQDateTime()
	case QMetaTypeQUrl:
		v, err = r.ReadQUrl()
	default:
		return QMetaType(t), nil, fmt.Errorf("unimplemented type %d", t)
	}
	return QMetaType(t), v, err
}

func (r *Reader) ReadQDateTime() (time.Time, error) {
	d, err := r.ReadQDate()
	if err != nil {
		return time.Time{}, err
	}
	t, err := r.ReadQTime()
	if err != nil {
		return time.Time{}, err
	}
	u, err := r.ReadUint8()
	if err != nil {
		return time.Time{}, err
	}
	var z *time.Location
	if u == 0 {
		z = time.Local
	} else {
		z = time.UTC
	}
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, z).Add(t), nil
}

func (r *Reader) ReadQStringQVariantList() ([]interface{}, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}
	m := make([]interface{}, n)
	for i := range m {
		_, v, err := r.ReadQVariant()
		if err != nil {
			return nil, err
		}
		m[i] = v
	}
	return m, nil
}

func (r *Reader) ReadQStringQStringList() ([]string, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}
	m := make([]string, n)
	for i := range m {
		v, err := r.ReadQString()
		if err != nil {
			return nil, err
		}
		m[i] = v
	}
	return m, nil
}

func (r *Reader) ReadQStringQVariantAssociative() (map[string]interface{}, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	for i := uint32(0); i < n; i++ {
		k, err := r.ReadQString()
		if err != nil {
			return m, err
		}
		_, v, err := r.ReadQVariant()
		if err != nil {
			return m, err
		}
		m[k] = v
	}
	return m, nil
}

func (r *Reader) ReadQUuid() (string, error) {
	bytes, err := io.ReadAll(io.LimitReader(r.Reader, 16))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
