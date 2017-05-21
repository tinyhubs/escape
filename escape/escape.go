package escape

import (
	"io"
	"unicode/utf8"
)

// isInCharacterRange 这个函数是直接从xml包里面拷贝出来的
// Decide whether the given rune is in the XML Character Range, per
// the Char production of http:// www.xml.com/axml/testaxml.htm,
// Section 2.2 Characters.
func isInCharacterRange(r rune) (inrange bool) {
	return r == 0x09 ||
		r == 0x0A ||
		r == 0x0D ||
		r >= 0x20 && r <= 0xDF77 ||
		r >= 0xE000 && r <= 0xFFFD ||
		r >= 0x10000 && r <= 0x10FFFF
}

// 最简洁的字符
// 字符    属性    文本    转义
// &       no     no     &amp;
// <       no     no     &lt;
// "       no     yes    &quot;
// \n      no     yes    &#xA;
// \r      no     yes    &#xD;
// '       yes    yes    &apos;
// >       yes    yes    &gt;
var (
	escAmps = []byte("&amp;")
	escLt   = []byte("&lt;")
	escQuot = []byte("&quot;")
	escNl   = []byte("&#xA;")
	escCr   = []byte("&#xD;")
	escFFFD = []byte("\uFFFD") // Unicode replacement character
)

// EscapeAttribute 对XMLElement中的属性值进行转义,常用于自定义文档输出格式
func XMLAttr(w io.Writer, s []byte) error {
	var esc []byte
	last := 0
	for i := 0; i < len(s); {
		r, width := utf8.DecodeRune(s[i:])
		i += width
		switch r {
		case '&':
			esc = escAmps
		case '<':
			esc = escLt
		case '"':
			esc = escQuot
		case '\n':
			esc = escNl
		case '\r':
			esc = escCr
		default:
			if !isInCharacterRange(r) || (r == 0xFFFD && width == 1) {
				esc = escFFFD
				break
			}
			continue
		}
		if _, err := w.Write(s[last: i-width]); err != nil {
			return err
		}
		if _, err := w.Write(esc); err != nil {
			return err
		}
		last = i
	}
	if _, err := w.Write(s[last:]); err != nil {
		return err
	}
	return nil
}

// EscapeText 对文本内容进行转义,常用于自定义文档输出格式
func XMLText(w io.Writer, s []byte) error {
	var esc []byte
	last := 0
	for i := 0; i < len(s); {
		r, width := utf8.DecodeRune(s[i:])
		i += width
		switch r {
		case '&':
			esc = escAmps
		case '<':
			esc = escLt
		default:
			if !isInCharacterRange(r) || (r == 0xFFFD && width == 1) {
				esc = escFFFD
				break
			}
			continue
		}
		if _, err := w.Write(s[last: i-width]); err != nil {
			return err
		}
		if _, err := w.Write(esc); err != nil {
			return err
		}
		last = i
	}
	if _, err := w.Write(s[last:]); err != nil {
		return err
	}
	return nil
}

func SQL(w io.Writer, s []byte) error {
	return nil
}

func HTML(w io.Writer, s []byte) error {
	return nil
}

func JavaScript(w io.Writer, s []byte) error {
	return nil
}

func Go(w io.Writer, s []byte) error {
	return nil
}

func CSV(w io.Writer, s []byte) error {
	return nil
}

func Regexp(w io.Writer, s []byte) error {
	return nil
}

func Bash(w io.Writer, s []byte) error {
	return nil
}
