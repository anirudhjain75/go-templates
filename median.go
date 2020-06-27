package main

type MedianFinder struct {
	Input []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		Input: make([]int, 0),
	}
}


func (this *MedianFinder) AddNum(num int)  {
	if len(this.Input) < 1 {
		this.Input = append(this.Input, num)
	} else {
		addToIndex := -1

		// Linear Search for Position
		for i, v := range this.Input {
			if num < v {
				addToIndex = i
				break
			}
		}

		// Append at position i
		if addToIndex != -1 {
			t := make([]int, 1)
			t[0] = num
			this.Input = append(this.Input[:addToIndex], append(t, this.Input[addToIndex:]...)...)
		} else {
			this.Input = append(this.Input, num)
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	lenOfStream := len(this.Input)

	if lenOfStream % 2 == 0 {
		return float64(this.Input[lenOfStream / 2] + this.Input[(lenOfStream / 2) - 1])/2.0
	} else {
		return float64(this.Input[lenOfStream / 2])
	}
}