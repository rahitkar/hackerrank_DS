package main

func getHourGlassSum(hourGlassCoordinate map[string]int, arr [][]int32) int32 {
	x := hourGlassCoordinate["x"]
	y := hourGlassCoordinate["y"]
	sum := int32(0)
	counter := 0
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			counter++
			if counter != 4 && counter != 6 {
				sum += arr[i][j]
			}
		}
	}
	return sum
}

func hourglassSum(arr [][]int32) int32 {
	maxHourGlassSum := int32(-2147483648)
	for numOfHourGlass := 0; numOfHourGlass < 16; numOfHourGlass++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				sum := getHourGlassSum(map[string]int{"x": i, "y": j}, arr)
				if sum > maxHourGlassSum {
					maxHourGlassSum = sum
				}
			}
		}
	}

	return maxHourGlassSum
}

func main() {
	//hourGlasses := [][]int32{
	//	{1, 1, 1, 0, 0, 0},
	//	{0, 1, 0, 0, 0, 0},
	//	{1, 1, 1, 0, 0, 0},
	//	{0, 0, 2, 4, 4, 0},
	//	{0, 0, 0, 2, 0, 0},
	//	{0, 0, 1, 2, 4, 0},
	//}

	hourGlasses := [][]int32{
		{0, -4, -6, 0, -7, -6},
		{-1, -2, -6, -8, -3, -1},
		{-8, -4, -2, -8, -8, -6},
		{-3, -1, -2, -5, -7, -4},
		{-3, -5, -3, -6, -6, -6},
		{-3, -6, 0, -8, -6, -7},
	}
	maxHourGlassSum := hourglassSum(hourGlasses)
	println(maxHourGlassSum)
}
