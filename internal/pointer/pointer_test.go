package pointer

import (
	"reflect"
	"testing"
	"time"
)

func TestBoolP(t *testing.T) {
	type args struct {
		val bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolP(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoalesceStrings(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoalesceStrings(tt.args.str...); got != tt.want {
				t.Errorf("CoalesceStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoalesceInts(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoalesceInts(tt.args.i...); got != tt.want {
				t.Errorf("CoalesceInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32P(t *testing.T) {
	type args struct {
		value float32
	}
	tests := []struct {
		name string
		args args
		want *float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32P(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32P() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32P(t *testing.T) {
	type args struct {
		value int32
	}
	tests := []struct {
		name string
		args args
		want *int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32P(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32P() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringP(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringP(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringDefault(t *testing.T) {
	type args struct {
		input        *string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringDefault(tt.args.input, tt.args.defaultValue); got != tt.want {
				t.Errorf("StringDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceContains(t *testing.T) {
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContains(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("SliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringPEqual(t *testing.T) {
	type args struct {
		a *string
		b *string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringPEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("StringPEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeP(t *testing.T) {
	type args struct {
		input time.Time
	}
	tests := []struct {
		name string
		args args
		want *time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeP(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimesPEqual(t *testing.T) {
	type args struct {
		a *time.Time
		b *time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimesPEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("TimesPEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxTime(t *testing.T) {
	type args struct {
		a time.Time
		b time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxTime(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
