package pointer

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSliceContains(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}

	// Test when the item is present in the slice
	result := SliceContains(slice, "banana")
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test when the item is not present in the slice
	result = SliceContains(slice, "pear")
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestStringPEqual(t *testing.T) {
	str1 := "Hello"
	str2 := "Hello"
	result := StringPEqual(StringP(str1), StringP(str2))
	expected := true
	if result != expected {
		t.Errorf("Unexpected result. Got: %t, want: %t", result, expected)
	}

	str3 := "World"
	result = StringPEqual(StringP(str1), StringP(str3))
	expected = false
	if result != expected {
		t.Errorf("Unexpected result. Got: %t, want: %t", result, expected)
	}

	result = StringPEqual(StringP(str1), nil)
	expected = false
	if result != expected {
		t.Errorf("Unexpected result. Got: %t, want: %t", result, expected)
	}

	result = StringPEqual(nil, StringP(str1))
	expected = false
	if result != expected {
		t.Errorf("Unexpected result. Got: %t, want: %t", result, expected)
	}

	result = StringPEqual(nil, nil)
	expected = true
	if result != expected {
		t.Errorf("Unexpected result. Got: %t, want: %t", result, expected)
	}
}

func TestTimeP(t *testing.T) {
	input := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	result := TimeP(input)

	// Check if the pointer is not nil
	if result == nil {
		t.Error("Unexpected nil pointer")
	}

	// Check if the pointer points to the correct time.Time value
	if *result != input {
		t.Errorf("Unexpected result. Got: %v, want: %v", *result, input)
	}
}

func TestCoalesceStrings(t *testing.T) {
	result := CoalesceStrings("", "foo", "bar")
	expected := "foo"
	if result != expected {
		t.Errorf("Unexpected result. Got: %s, want: %s", result, expected)
	}
}

func TestCoalesceInts(t *testing.T) {
	result := CoalesceInts(0, -10, 20)
	expected := 20
	if result != expected {
		t.Errorf("Unexpected result. Got: %d, want: %d", result, expected)
	}
}

func TestStringDefault(t *testing.T) {
	input := StringP("Hello")
	result := StringDefault(input, "Default")
	expected := "Hello"
	if result != expected {
		t.Errorf("Unexpected result. Got: %s, want: %s", result, expected)
	}

	input = nil
	result = StringDefault(input, "Default")
	expected = "Default"
	if result != expected {
		t.Errorf("Unexpected result. Got: %s, want: %s", result, expected)
	}
}

func TestTimesPEqual(t *testing.T) {
	t1 := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	result := TimesPEqual(TimeP(t1), TimeP(t2))
	expected := true
	if result != expected {
		t.Errorf("Unexpected result. Got: %t, want: %t", result, expected)
	}

	result = TimesPEqual(TimeP(t1), nil)
	expected = false
	if result != expected {
		t.Errorf("Unexpected result. Got: %t, want: %t", result, expected)
	}
}

func TestToSliceOf(t *testing.T) {
	input := []interface{}{1, 2, 3, 4, 5}
	result, err := ToSliceOf[int](input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result. Got: %v, want: %v", result, expected)
	}

	_, err = ToSliceOf[string](input)
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}

func TestBoolP(t *testing.T) {
	result := BoolP(true)
	expected := true
	if *result != expected {
		t.Errorf("Unexpected result. Got: %t, want: %t", *result, expected)
	}

	result = BoolP(false)
	expected = false
	if *result != expected {
		t.Errorf("Unexpected result. Got: %t, want: %t", *result, expected)
	}
}

func TestMap(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	callback := func(arg int) string {
		return fmt.Sprintf("Item: %d", arg)
	}

	result := Map(items, callback)
	expected := []string{"Item: 1", "Item: 2", "Item: 3", "Item: 4", "Item: 5"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unexpected result. Got: %v, want: %v", result, expected)
	}
}

func TestFloat32P(t *testing.T) {
	result := Float32P(3.14)
	expected := float32(3.14)
	if *result != expected {
		t.Errorf("Unexpected result. Got: %f, want: %f", *result, expected)
	}

	result = Float32P(0.0)
	expected = float32(0.0)
	if *result != expected {
		t.Errorf("Unexpected result. Got: %f, want: %f", *result, expected)
	}
}

func TestFloat64P(t *testing.T) {
	result := Float64P(3.14159)
	expected := 3.14159
	if *result != expected {
		t.Errorf("Unexpected result. Got: %f, want: %f", *result, expected)
	}

	result = Float64P(0.0)
	expected = 0.0
	if *result != expected {
		t.Errorf("Unexpected result. Got: %f, want: %f", *result, expected)
	}
}

func TestInt32P(t *testing.T) {
	result := Int32P(42)
	expected := int32(42)
	if *result != expected {
		t.Errorf("Unexpected result. Got: %d, want: %d", *result, expected)
	}

	result = Int32P(0)
	expected = int32(0)
	if *result != expected {
		t.Errorf("Unexpected result. Got: %d, want: %d", *result, expected)
	}
}

func TestMax(t *testing.T) {
	result := Max(10, 20)
	expected := 20
	if result != expected {
		t.Errorf("Unexpected result. Got: %d, want: %d", result, expected)
	}

	result = Max(-5, -10)
	expected = -5
	if result != expected {
		t.Errorf("Unexpected result. Got: %d, want: %d", result, expected)
	}

	result = Max(0, 0)
	expected = 0
	if result != expected {
		t.Errorf("Unexpected result. Got: %d, want: %d", result, expected)
	}
}

func TestMin(t *testing.T) {
	result := Min(10, 20)
	expected := 10
	if result != expected {
		t.Errorf("Unexpected result. Got: %d, want: %d", result, expected)
	}

	result = Min(-5, -10)
	expected = -10
	if result != expected {
		t.Errorf("Unexpected result. Got: %d, want: %d", result, expected)
	}

	result = Min(0, 0)
	expected = 0
	if result != expected {
		t.Errorf("Unexpected result. Got: %d, want: %d", result, expected)
	}
}

func TestStringP(t *testing.T) {
	result := StringP("Hello")
	expected := "Hello"
	if *result != expected {
		t.Errorf("Unexpected result. Got: %s, want: %s", *result, expected)
	}

	result = StringP("")
	expected = ""
	if *result != expected {
		t.Errorf("Unexpected result. Got: %s, want: %s", *result, expected)
	}
}

func TestMaxTime(t *testing.T) {
	t1 := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC)
	result := MaxTime(t1, t2)
	expected := t2
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, want: %v", result, expected)
	}

	result = MaxTime(t1, t1)
	expected = t1
	if result != expected {
		t.Errorf("Unexpected result. Got: %v, want: %v", result, expected)
	}
}
