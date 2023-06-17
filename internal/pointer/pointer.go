package pointer

import (
	"fmt"
	"strings"
	"time"
)

// BoolP returns a pointer to the bool value passed in.
func BoolP(val bool) *bool {
	return &val
}

// CoalesceStrings returns the first non-empty string from the list of strings passed in.
func CoalesceStrings(str ...string) string {
	for _, value := range str {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}

	return ""
}

// CoalesceInts returns the first non-zero int from the list of ints passed in.
func CoalesceInts(i ...int) int {
	for _, value := range i {
		if value > 0 {
			return value
		}
	}

	return 0
}

// Map will take a slice of items and a callback function and return a new slice of items that have been transformed by the callback function.
func Map[A any, B any, F func(arg A) B](items []A, callback F) []B {
	result := make([]B, 0, len(items))
	for _, item := range items {
		result = append(result, callback(item))
	}

	return result
}

// Float32P returns a pointer to the float32 value passed in.
func Float32P(value float32) *float32 {
	return &value
}

// Float64P returns a pointer to the float64 value passed in.
func Float64P(value float64) *float64 {
	return &value
}

// Int32P returns a pointer to the int32 value passed in.
func Int32P(value int32) *int32 {
	return &value
}

// Number is an interface that represents a number type.
type Number interface {
	int | int32 | int64
}

// Max will return the maximum of two numbers.
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// Min will return the minimum of two numbers.
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}

// StringP returns a pointer to the string value passed in.
func StringP(input string) *string {
	return &input
}

// StringDefault will return the value of the string pointer if it is not nil.
// Otherwise it will return the default value.
func StringDefault(input *string, defaultValue string) string {
	if input != nil {
		return *input
	}

	return defaultValue
}

// SliceContains will return true if the slice contains the item.
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

// TimeP returns a pointer to the time.Time value passed in.
func TimeP(input time.Time) *time.Time {
	return &input
}

// TimesPEqual will compare whether or not two time.Time pointers are equal. If one or the other is nil then it will
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

// MaxTime will return the maximum of two times.
func MaxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	} else {
		return b
	}
}

// ToSliceOf will convert a slice of interface{} to a slice of the specified type.
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
