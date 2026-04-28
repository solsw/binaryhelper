package binaryhelper

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// CopyFixed reinterprets the binary representation of 'src' into 'dst'.
// Both 'src' and 'dst' must be fixed-size values (or pointers to them) as
// required by [binary.Write] and [binary.Read] respectively.
// Byte order does not affect in-process copies.
func CopyFixed(src, dst any) error {
	if src == nil {
		return errors.New("src is nil")
	}
	if dst == nil {
		return errors.New("dst is nil")
	}
	var b bytes.Buffer
	if err := binary.Write(&b, binary.LittleEndian, src); err != nil {
		return err
	}
	r := bytes.NewReader(b.Bytes())
	if err := binary.Read(r, binary.LittleEndian, dst); err != nil {
		return err
	}
	if r.Len() != 0 {
		return errors.New("src is larger than dst")
	}
	return nil
}
