package binaryhelper

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

// CopyFixed reinterprets the little-endian encoding of src into dst.
// src may be a fixed-size value, a slice of fixed-size values, or a pointer
// to either. dst must be a slice or a pointer to a slice or fixed-size value.
// The encoded sizes must match. A slice's length determines its encoded size.
//
// Destination struct fields, including nested fields, must be exported or
// blank. Blank fields serve as padding. On error, dst retains its original value.
func CopyFixed(src, dst any) error {
	if !reflect.Indirect(reflect.ValueOf(src)).IsValid() {
		return errors.New("src is nil")
	}
	target := reflect.Indirect(reflect.ValueOf(dst))
	if !target.IsValid() {
		return errors.New("dst is nil")
	}
	if target.Kind() != reflect.Slice && !target.CanSet() {
		return errors.New("dst must be a pointer or a slice")
	}
	size := binary.Size(dst)
	if size < 0 {
		return fmt.Errorf("dst has unsupported binary type %T", dst)
	}
	if err := checkExportedFields(target.Type()); err != nil {
		return err
	}
	var b bytes.Buffer
	if err := binary.Write(&b, binary.LittleEndian, src); err != nil {
		return err
	}
	if b.Len() > size {
		return errors.New("src is larger than dst")
	}
	return binary.Read(&b, binary.LittleEndian, dst)
}

// checkExportedFields checks a type already accepted by binary.Size.
func checkExportedFields(t reflect.Type) error {
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		return checkExportedFields(t.Elem())
	case reflect.Struct:
		for field := range t.Fields() {
			if field.Name == "_" {
				continue
			}
			if !field.IsExported() {
				return fmt.Errorf("dst field %s.%s is unexported", t, field.Name)
			}
			if err := checkExportedFields(field.Type); err != nil {
				return err
			}
		}
	}
	return nil
}
