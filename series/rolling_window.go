package series

// RollingWindow is used for rolling window calculations.
type RollingWindow struct {
	window int
	series Series
}

// Rolling creates new RollingWindow
func (s Series) Rolling(window int) RollingWindow {
	return RollingWindow{
		window: window,
		series: s,
	}
}

// Mean returns the rolling mean.
func (r RollingWindow) Mean() (s Series) {
	s = New([]float64{}, Float, "Mean")
	for _, block := range r.getBlocks() {
		s.Append(block.Mean())
	}

	return
}

// StdDev returns the rolling mean.
func (r RollingWindow) StdDev() (s Series) {
	s = New([]float64{}, Float, "StdDev")
	for _, block := range r.getBlocks() {
		s.Append(block.StdDev())
	}

	return
}

// Quantile 计算分位数 p >= 0 <= 1
func (r RollingWindow) Quantile(p float64) (s Series) {
	s = New([]float64{}, Float, "Quantile")
	for _, block := range r.getBlocks() {
		s.Append(block.Quantile(p))
	}

	return
}

// Max 计算最大值
func (r RollingWindow) Max() (s Series) {
	s = New([]float64{}, Float, "Max")
	for _, block := range r.getBlocks() {
		s.Append(block.Max())
	}

	return
}

// Min 计算最大值
func (r RollingWindow) Min() (s Series) {
	s = New([]float64{}, Float, "Min")
	for _, block := range r.getBlocks() {
		s.Append(block.Min())
	}

	return
}

// Sum 计算最大值
func (r RollingWindow) Sum() (s Series) {
	s = New([]float64{}, Float, "Min")
	for _, block := range r.getBlocks() {
		s.Append(block.Sum())
	}

	return
}

func (r RollingWindow) getBlocks() (blocks []Series) {
	for i := 1; i <= r.series.Len(); i++ {
		if i < r.window {
			blocks = append(blocks, r.series.Empty())
			continue
		}

		index := []int{}
		for j := i - r.window; j < i; j++ {
			index = append(index, j)
		}
		blocks = append(blocks, r.series.Subset(index))
	}

	return
}

func (r RollingWindow) GetBlocks() (blocks []Series) {
	for i := 1; i <= r.series.Len(); i++ {
		if i < r.window {
			blocks = append(blocks, r.series.Empty())
			continue
		}

		index := []int{}
		for j := i - r.window; j < i; j++ {
			index = append(index, j)
		}
		blocks = append(blocks, r.series.Subset(index))
	}

	return
}
