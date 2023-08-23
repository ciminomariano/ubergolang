package tests

import (
	"testing"
	"ubergolang/mathoperations"
)

func TestSum(t *testing.T) {
	result := mathoperations.Sum(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Sum function returned %d, expected %d", result, expected)
	}
}
