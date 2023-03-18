package pointer

import (
	"fmt"
	"strings"
	"time"
)

func BoolP(val bool) *bool {
	return &val
}

func CoalesceStrings(str ...string) string {
	for _, value := range str {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}

	return ""
}

func CoalesceInts(i ...int) int {
	for _, value := range i {
		if value > 0 {
			return value
		}
	}

	return 0
}

func Map[A any, B any, F func(arg A) B](items []A, callback F) []B {
	result := make([]B, 0, len(items))
	for _, item := range items {
		result = append(result, callback(item))
	}

	return result
}

func Float32P(value float32) *float32 {
	return &value
}

func Int32P(value int32) *int32 {
	return &value
}

type Number interface {
	int | int32 | int64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func StringP(input string) *string {
	return &input
}

func StringDefault(input *string, defaultValue string) string {
	if input != nil {
		return *input
	}

	return defaultValue
}

func SliceContains(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}

	return false
}

// StringPEqual will compare whether or not two string pointers are equal. If one or the other is nil then it will
// return false. Otherwise it will compare their values as strings not as pointers.
func StringPEqual(a, b *string) bool {
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	if a == nil && b == nil {
		return true
	}

	return *a == *b
}

func TimeP(input time.Time) *time.Time {
	return &input
}

func TimesPEqual(a, b *time.Time) bool {
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	if a == nil && b == nil {
		return true
	}

	// Just to silence editor warning. Neither should be nil at this point.
	if a == nil || b == nil {
		return false
	}

	return a.Equal(*b)
}

func MaxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	} else {
		return b
	}
}

func ToSliceOf[T interface{}](arg []interface{}) ([]T, error) {
	s := make([]T, len(arg))
	for i, v := range arg {
		rec, ok := v.(T)
		if !ok {
			return nil, fmt.Errorf("cannot convert")
		}
		s[i] = rec
	}

	return s, nil
}
