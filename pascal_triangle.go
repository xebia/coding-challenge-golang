package main

func triangleRow(n int) []int64 {
	if n == 0 {
		return []int64{1}
	}
	var previousRow = triangleRow(n - 1)
	row := make([]int64, 0, n+1)
	row = append(row, 1)
	for col := 1; col < n; col++ {
		row = append(row, previousRow[col-1]+previousRow[col]+1)
	}
	row = append(row, 1)
	return row
}
