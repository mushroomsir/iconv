package iconv

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	os.Mkdir("output", 0755)
}
func TestConvertBytes(t *testing.T) {
	require := require.New(t)

	testCase := []struct {
		src  string
		from string
		want string
		to   string
	}{
		{"花间一壶酒，独酌无相亲。", UTF8, "\xbb\xa8\xbc\xe4\xd2\xbb\xba\xf8\xbe\xc6\xa3\xac\xb6\xc0\xd7\xc3" +
			"\xce\xde\xcf\xe0\xc7\xd7\xa1\xa3", GB18030},
		{"\u0081\u00de\u00df\u00e0\u00e1\u00e2\u00e3\uffff\U00010000", UTF8, "\x81\x30\x81\x31\x81\x30\x89\x37\x81\x30\x89\x38\xa8\xa4\xa8\xa2" +
			"\x81\x30\x89\x39\x81\x30\x8a\x30\x84\x31\xa4\x39\x90\x30\x81\x30", GB18030},
		{"花间一壶酒，独酌无相亲。", UTF8, "\xbb\xa8\xbc\xe4\xd2\xbb\xba\xf8\xbe\xc6\xa3\xac\xb6\xc0\xd7\xc3" +
			"\xce\xde\xcf\xe0\xc7\xd7\xa1\xa3", GBK},
		{"A\u3000\u554a\u4e02\u4e90\u72dc\u7349\u02ca\u2588Z€", UTF8, "A\xa1\xa1\xb0\xa1\x81\x40\x81\x80\xaa\x40\xaa\x80\xa8\x40\xa8\x80Z\x80", GBK},
	}
	for i, val := range testCase {

		dst, err := ConvertBytes([]byte(val.src), val.from, val.to)
		require.Nil(err)
		require.Equal(val.want, string(dst))

		fileWriter, err := os.Create(fmt.Sprintf("output/%v-%vTo%v.txt", i, val.from, val.to))
		require.Nil(err)
		fileWriter.Write(dst)
		fileWriter.Close()
	}
}
func TestConvertString(t *testing.T) {
	require := require.New(t)

	testCase := []struct {
		src  string
		from string
		want string
		to   string
	}{
		{"a\xfe\xfeb", GBK, "a\ufffdb", UTF8},
		{"\x80", GB18030, "€", UTF8},
	}
	for i, val := range testCase {
		dst, err := ConvertString(val.src, val.from, val.to)
		require.Nil(err)
		require.Equal(val.want, dst)

		fileWriter, err := os.Create(fmt.Sprintf("output/%v-%vTo%v.txt", i, val.from, val.to))
		require.Nil(err)
		fileWriter.WriteString(dst)
		fileWriter.Close()
	}
}
