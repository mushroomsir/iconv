# iconv
[![Build Status](https://img.shields.io/travis/mushroomsir/iconv.svg?style=flat-square)](https://travis-ci.org/mushroomsir/iconv)
[![Coverage Status](http://img.shields.io/coveralls/mushroomsir/iconv.svg?style=flat-square)](https://coveralls.io/github/mushroomsir/iconv?branch=master)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/mushroomsir/iconv/blob/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mushroomsir/iconv)

# Installation

```sh
go get -u github.com/mushroomsir/iconv
```
## Support charset
- GBK 
- GB18030
- UTF8
- Big5
- ISO-8859-1
- EUCJP
- More coming soon

# Usage
```go
import (
	github.com/mushroomsir/iconv
)
```
## Converting string Values 

Converting a string can be done with two methods. First, there's
iconv.ConvertString(input, fromEncoding, toEncoding string) syntactic sugar.
```go
output,err := iconv.ConvertString("Hello World!", iconv.GBK, iconv.UTF8)
```

Alternatively, you can create a converter and use its ConvertString method.
Reuse of a Converter instance is recommended when doing many string conversions
between the same encodings.
```go
converter := iconv.NewConverter(iconv.GBK, iconv.UTF8)
output,err := converter.ConvertString("Hello World!")
```

## Converting []byte Values

Converting a []byte can similarly be done with two methods. First, there's
iconv.Convert(input []byte, fromEncoding, toEncoding string). 
```go
input := []byte("Hello World!")
output, err := iconv.ConvertBytes(input, iconv.GBK, iconv.UTF8)
```
Just like with ConvertString, there is also a Convert method on Converter that
can be used.
```go
convert,err := iconv.NewConverter(iconv.GBK, iconv.UTF8)
input := []byte("Hello World!")
output, err := converter.ConvertBytes(input)
```


## Converting an io.Reader

The iconv.Reader allows any other \*io.Reader to be wrapped and have its bytes
transcoded as they are read. 
```go
reader,err := iconv.Convert(strings.NewReader("Hello World!"),  iconv.GBK, iconv.UTF8)
```

## Licenses

All source code is licensed under the [MIT License](https://github.com/mushroomsir/iconv/blob/master/LICENSE).
