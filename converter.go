package iconv

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"

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
	Big5       = "Big5"
	charsets   = []string{UTF8, GBK, GB18030}
	charsetMap = map[string]transform.Transformer{}
)

// Converter ...
type Converter struct {
	FromEncoding string
	ToEncoding   string
}

func init() {
	GBKToUTF8 := GBK + UTF8
	charsetMap[GBKToUTF8] = simplifiedchinese.GBK.NewDecoder()

	UTF8ToGBK := UTF8 + GBK
	charsetMap[UTF8ToGBK] = simplifiedchinese.GBK.NewEncoder()

	GB18030ToUTF8 := GB18030 + UTF8
	charsetMap[GB18030ToUTF8] = simplifiedchinese.GB18030.NewDecoder()

	UTF8ToGB18030 := UTF8 + GB18030
	charsetMap[UTF8ToGB18030] = simplifiedchinese.GB18030.NewEncoder()

	Big5ToUTF8 := Big5 + UTF8

	charsetMap[Big5ToUTF8] = traditionalchinese.Big5.NewDecoder()

	UTF8ToBig5 := UTF8 + Big5
	charsetMap[UTF8ToBig5] = traditionalchinese.Big5.NewEncoder()
}

// NewConverter Initialize a new Converter. If fromEncoding or toEncoding are not supported
// then an error will be returned.
func NewConverter(fromEncoding string, toEncoding string) (*Converter, error) {
	if _, ok := charsetMap[fromEncoding+toEncoding]; ok {
		return &Converter{FromEncoding: fromEncoding, ToEncoding: toEncoding}, nil
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
	srcToDst := string(t.FromEncoding) + string(t.ToEncoding)
	return transform.NewReader(reader, charsetMap[srcToDst])
}
