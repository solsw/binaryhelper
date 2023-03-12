package binaryhelper

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// CopyFixed copies bytes from 'src' to 'dst'.
// 'src' must meet requirements for 'data' from binary.Write.
// 'dst' must meet requirements for 'data' from binary.Read.
func CopyFixed(src, dst any) error {
	if src == nil || dst == nil {
		return errors.New("nil src or/and dst")
	}
	var b bytes.Buffer
	if err := binary.Write(&b, binary.LittleEndian, src); err != nil {
		return err
	}
	r := bytes.NewReader(b.Bytes())
	if err := binary.Read(r, binary.LittleEndian, dst); err != nil {
		return err
	}
	return nil
}
