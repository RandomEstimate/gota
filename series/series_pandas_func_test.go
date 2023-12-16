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

func TestAdd(t *testing.T) {
	series1 := New([]float64{3.0, 4.0, 5.0}, Float, "COL")
	series2 := New([]float64{3.0, 4.0, 5.0}, Float, "COL")

	series := Add(series1, series2)
	if series.Elem(0).Float() != 6 {
		t.Errorf("无法相加")
	}
}

func TestSub(t *testing.T) {
	series1 := New([]float64{3.0, 4.0, 5.0}, Float, "COL")
	series2 := New([]float64{3.0, 4.0, 5.0}, Float, "COL")

	series := Sub(series1, series2)
	if series.Elem(0).Float() != 0 {
		t.Errorf("无法相减")
	}
}

func TestMultiply(t *testing.T) {
	series1 := New([]float64{3.0, 4.0, 5.0}, Float, "COL")
	series2 := New([]float64{3.0, 4.0, 5.0}, Float, "COL")

	series := Multiply(series1, series2)
	if series.Elem(0).Float() != 9 {
		t.Errorf("无法相乘")
	}
}

func TestDiv(t *testing.T) {
	series1 := New([]float64{3.0, 4.0, 5.0}, Float, "COL")
	series2 := New([]float64{3.0, 4.0, 5.0}, Float, "COL")

	series := Div(series1, series2)
	if series.Elem(0).Float() != 1 {
		t.Errorf("无法相除")
	}
}
