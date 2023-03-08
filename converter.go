package iconv

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

var (
	ErrNotSupportCharset = errors.New("not support this charset for now")
)

var (
	UTF8       = "UTF-8"
	GBK        = "GBK"
	GB18030    = "GB-18030"
	HZGB2312   = "HZ-GB2312"
	Big5       = "Big5"
	ISO88591   = "ISO-8859-1"
	EUCJP      = "EUC-JP"
	ShiftJIS   = "Shift_JIS"
	CP850      = "CP850"
	charsets   = []string{GBK, GB18030, Big5, ISO88591, EUCJP, ShiftJIS, HZGB2312, CP850}
	charsetMap = map[string]transform.Transformer{}
)

// Converter ...
type Converter struct {
	fromEncoding string
	toEncoding   string
}

func init() {
	for _, charset := range charsets {
		switch charset {
		case GBK:
			charsetMap[GBK+UTF8] = simplifiedchinese.GBK.NewDecoder()
			charsetMap[UTF8+GBK] = simplifiedchinese.GBK.NewEncoder()
		case GB18030:
			charsetMap[GB18030+UTF8] = simplifiedchinese.GB18030.NewDecoder()
			charsetMap[UTF8+GB18030] = simplifiedchinese.GB18030.NewEncoder()
		case Big5:
			charsetMap[Big5+UTF8] = traditionalchinese.Big5.NewDecoder()
			charsetMap[UTF8+Big5] = traditionalchinese.Big5.NewEncoder()
		case ISO88591:
			charsetMap[ISO88591+UTF8] = charmap.ISO8859_1.NewDecoder()
			charsetMap[UTF8+ISO88591] = charmap.ISO8859_1.NewEncoder()
		case EUCJP:
			charsetMap[EUCJP+UTF8] = japanese.EUCJP.NewDecoder()
			charsetMap[UTF8+EUCJP] = japanese.EUCJP.NewEncoder()
		case ShiftJIS:
			charsetMap[ShiftJIS+UTF8] = japanese.ShiftJIS.NewDecoder()
			charsetMap[UTF8+ShiftJIS] = japanese.ShiftJIS.NewEncoder()
		case HZGB2312:
			charsetMap[HZGB2312+UTF8] = simplifiedchinese.HZGB2312.NewDecoder()
			charsetMap[UTF8+HZGB2312] = simplifiedchinese.HZGB2312.NewEncoder()
		case CP850:
			charsetMap[CP850+UTF8] = charmap.CodePage850.NewDecoder()
			charsetMap[UTF8+CP850] = charmap.CodePage850.NewEncoder()
		}
	}
}

// NewConverter Initialize a new Converter. If fromEncoding or toEncoding are not supported
// then an error will be returned.
func NewConverter(fromEncoding string, toEncoding string) (*Converter, error) {
	_, fromOK := charsetMap[fromEncoding+UTF8]
	_, toOK := charsetMap[UTF8+toEncoding]
	if _, ok := charsetMap[fromEncoding+toEncoding]; ok || (fromOK && toOK) {
		return &Converter{fromEncoding: fromEncoding, toEncoding: toEncoding}, nil
	}
	return nil, ErrNotSupportCharset
}

// ConvertString Convert an input string
func (t *Converter) ConvertString(input string) (string, error) {
	res, err := t.ConvertBytes([]byte(input))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// ConvertBytes ...
func (t *Converter) ConvertBytes(input []byte) ([]byte, error) {
	reader := t.Convert(bytes.NewReader(input))
	return ioutil.ReadAll(reader)
}

// Convert ...
func (t *Converter) Convert(reader io.Reader) io.Reader {
	return t.convert(reader)
}

func (t *Converter) convert(reader io.Reader) io.Reader {
	srcToDst := t.fromEncoding + t.toEncoding
	if val, ok := charsetMap[srcToDst]; ok {
		return transform.NewReader(reader, val)
	}
	resReader := transform.NewReader(reader, charsetMap[t.fromEncoding+UTF8])
	resReader = transform.NewReader(resReader, charsetMap[UTF8+t.toEncoding])
	return resReader
}
