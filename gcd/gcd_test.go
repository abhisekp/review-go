package gcd

import "testing"

func TestGCD(t *testing.T) {
	t.Run("GCD", func(t *testing.T) {
		a := 10
		b := 20
		expected := 10
		actual := GCD(a, b)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("GCD with negative numbers", func(t *testing.T) {
		a := -10
		b := -20
		expected := 10
		actual := GCD(a, b)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("GCD with one negative number", func(t *testing.T) {
		a := -10
		b := 20
		expected := 10
		actual := GCD(a, b)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("GCD with one negative number", func(t *testing.T) {
		a := 10
		b := -20
		expected := 10
		actual := GCD(a, b)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("GCD with zero", func(t *testing.T) {
		a := 0
		b := 20
		expected := 20
		actual := GCD(a, b)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("GCD with one", func(t *testing.T) {
		a := 10
		b := 1
		expected := 1
		actual := GCD(a, b)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
}
