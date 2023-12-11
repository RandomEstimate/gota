package series

import (
	"fmt"
	"math"
)

// SetValue 对Series中的元素进行值修改
func (s Series) SetValue(index int, newValue interface{}) {

	// 检测是否超出 长度
	if s.Len()-1 > index || index < 0 {
		s.Err = fmt.Errorf("set error: index out of range")
	}

	switch s.Type() {
	case String:
		s.elements.(stringElements)[index].Set(newValue)
	case Int:
		s.elements.(intElements)[index].Set(newValue)
	case Float:
		s.elements.(floatElements)[index].Set(newValue)
	case Bool:
		s.elements.(boolElements)[index].Set(newValue)
	}
}

// Shift 对Series序列进行上下移动
func (s Series) Shift(n int) Series {

	newSeries := New([]float64{}, Float, "Shift")
	if int(math.Abs(float64(n))) > s.Len() {
		for i := 0; i < s.Len(); i++ {
			newSeries.Append(math.NaN())
		}

	} else if n >= 0 {
		for i := 0; i < n; i++ {
			newSeries.Append(math.NaN())
		}

		for i := n; i < s.Len(); i++ {
			newSeries.Append(s.Elem(i - n).Float())
		}
	} else if n < 0 {
		for i := int(math.Abs(float64(n))); i < s.Len(); i++ {
			newSeries.Append(s.Elem(i).Float())
		}

		for i := 0; i < int(math.Abs(float64(n))); i++ {
			newSeries.Append(math.NaN())
		}
	}

	return newSeries
}

// Diff 对Series序列进行差分
func (s Series) Diff(n int) Series {
	newSeries := New([]float64{}, Float, "Diff")
	for _, block := range s.Rolling(int(math.Abs(float64(n)) + 1)).getBlocks() {
		if block.Len() == 0 {
			newSeries.Append(math.NaN())
			continue
		}

		if n >= 0 {
			newSeries.Append(block.Elem(block.Len()-1).Float() - block.Elem(0).Float())
		} else if n < 0 {
			newSeries.Append(block.Elem(0).Float() - block.Elem(block.Len()-1).Float())
		}
	}

	if n < 0 {
		newSeries.Reverse()
	}

	return newSeries
}

// Deprecated: Apply 类似于pandas 的apply函数 原始函数中存在Map 这个方法
func (s Series) Apply(f func(element Element) Element) Series {
	newSeries := New([]Element{}, Float, "Apply")
	for i := 0; i < s.Len(); i++ {
		newSeries.Append(f(s.Elem(i)))
	}

	return newSeries
}

// Reverse 翻转series
func (s Series) Reverse() {
	i := 0
	j := s.Len() - 1
	for i < j {
		tmp := s.elements.(floatElements).Elem(i).Copy()
		s.elements.(floatElements).Elem(i).Set(s.elements.(floatElements).Elem(j))
		s.elements.(floatElements).Elem(j).Set(tmp)
		i++
		j--
	}
}
