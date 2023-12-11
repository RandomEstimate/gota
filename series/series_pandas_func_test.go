package series

import (
	"testing"
)

func TestSeries_SetValue(t *testing.T) {
	series := New([]float64{3.0, 4.0}, Float, "COL")

	series.SetValue(0, 100)

	if series.Elem(0).Float() != 100 {
		t.Errorf("无法直接修改 series的值")
	}
}

func TestSeries_Diff(t *testing.T) {
	series := New([]float64{3.0, 4.0, 5.0}, Float, "COL")

	diff := series.Diff(-2)
	if diff.Elem(0).Float() != -2 {
		t.Errorf("无法直接完成差分")
	}
}

func TestSeries_Shift(t *testing.T) {
	series := New([]float64{3.0, 4.0, 5.0}, Float, "COL")

	shift := series.Shift(2)
	if shift.Elem(2).Float() != 3 {
		t.Errorf("无法直接完成位移")
	}
}
