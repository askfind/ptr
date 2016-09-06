// Package pointer provides helpers to get pointers to values of build-in types.
// Based on github.com/AlekSi/pointer
// This is too small helper. Ideally copy it to your project.
package ptr

import (
	"time"
)

func Bool(b bool) *bool                   { return &b }
func Byte(b byte) *byte                   { return &b }
func Complex128(c complex128) *complex128 { return &c }
func Complex64(c complex64) *complex64    { return &c }
func Error(e error) *error                { return &e }
func Float32(f float32) *float32          { return &f }
func Float64(f float64) *float64          { return &f }
func Int(i int) *int                      { return &i }
func Int16(i int16) *int16                { return &i }
func Int32(i int32) *int32                { return &i }
func Int64(i int64) *int64                { return &i }
func Int8(i int8) *int8                   { return &i }
func Rune(r rune) *rune                   { return &r }
func String(s string) *string             { return &s }
func Time(t time.Time) *time.Time         { return &t }
func Uint(u uint) *uint                   { return &u }
func Uint16(u uint16) *uint16             { return &u }
func Uint32(u uint32) *uint32             { return &u }
func Uint64(u uint64) *uint64             { return &u }
func Uint8(u uint8) *uint8                { return &u }
func Uintptr(u uintptr) *uintptr          { return &u }
