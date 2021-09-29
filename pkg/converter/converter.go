package converter

import (
	"time"
)

type Converter interface {
	StringPointerToString(s *string) string
	TimePointerToTime(t *time.Time) time.Time
	Int32PointerToInt32(n *int32) int32
	Float32PointerToFloat32(n *float32) float32
	BoolPointerToBool(n *bool) bool
}

type converter struct{}

func New() Converter {
	return &converter{}
}

func (c *converter) StringPointerToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func (c *converter) TimePointerToTime(t *time.Time) time.Time {
	if t != nil {
		return *t
	}
	return time.Time{}
}

func (c *converter) Int32PointerToInt32(n *int32) int32 {
	if n != nil {
		return *n
	}
	return 0
}

func (c *converter) Float32PointerToFloat32(n *float32) float32 {
	if n != nil {
		return *n
	}
	return 0.0
}

func (c *converter) BoolPointerToBool(n *bool) bool {
	if n != nil {
		return *n
	}
	return false
}
