package interview

import "fmt"

/*
环形遍历
1,2,3,4
5,6,7,8
9,10,11,12
输出：1，2，3，4，8，12，11，10，9，5，6，7

a := [][]int{
		{1,2,3,4,5},
		{12,13,14,15,6},
		{11,10,9,8,7},
	}
clockwiseIter(a)
*/
func clockwiseIter(num [][]int) {
	row := 0
	var col int
	xLeft := 0
	xRight := len(num[0]) - 1
	yUp := 0
	count := 0
	yDown := len(num) - 1
	totalCount := len(num) * len(num[0])

	direction := 0

	for count < totalCount {
		if direction == 0 {
			for col = xLeft; col <= xRight; col++ {
				count++
				fmt.Print(num[row][col])
			}

			yUp += 1
			direction = 1
			col--
		} else if direction == 1 {
			for row = yUp; row <= yDown; row++ {
				count++
				fmt.Print(num[row][col])
			}
			xRight -= 1
			direction = 2
			row--
		} else if direction == 2 {
			for col = xRight; col >= xLeft; col-- {
				count++
				fmt.Print(num[row][col])
			}
			yDown -= 1
			direction = 3
			col++
		} else if direction == 3 {
			for row = yDown; row >= yUp; row-- {
				count++
				fmt.Print(num[row][col])
			}
			xLeft += 1
			direction = 0
			row++
		}
	}
}
