package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestRow0(t *testing.T) {
	assert.Equal(t, []int64{1}, triangleRow(0))
}
func TestRow1(t *testing.T) {
	assert.Equal(t, []int64{1, 1}, triangleRow(1))
}
func TestRow2(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 1}, triangleRow(2))
}
func TestRow3(t *testing.T) {
	assert.Equal(t, []int64{1, 3, 3, 1}, triangleRow(3))
}
func TestRow4(t *testing.T) {
	assert.Equal(t, []int64{1, 4, 6, 4, 1}, triangleRow(4))
}
