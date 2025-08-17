package convert

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode/utf32"
	"golang.org/x/text/transform"
)

//var UTF32Encoder = utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM).NewEncoder()

var UTF32Encoder = &encoding.Encoder{
	Transformer: transform.Chain(encoding.UTF8Validator, utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM).NewEncoder()),
}

//var UTF32Decoder = utf32.UTF32(utf32.BigEndian, utf32.IgnoreBOM).NewDecoder()

func ToUTF32(s string) (string, error) {
	return UTF32Encoder.String(s)
}

//
//func FromUTF32(s string) (string, error) {
//	return UTF32Decoder.String(s)
//}
//
//func IsUTF32(s string) bool {
//	_, err := UTF32Decoder.String(s)
//	return err == nil
//}
