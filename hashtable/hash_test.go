package hashtable

import "testing"

func TestHash(t *testing.T) {
	t.Run("Hash1", func(t *testing.T) {
		key := "abc"
		hashSize := 10
		expected := 1
		actual := Hash1(key, hashSize)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Hash2", func(t *testing.T) {
		key := "abc"
		hashSize := 10
		expected := 3
		actual := Hash2(key, hashSize)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Hash3", func(t *testing.T) {
		key := "abc"
		hashSize := 10
		expected := 7
		actual := Hash3(key, hashSize)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("Hash4", func(t *testing.T) {
		key := "abc"
		hashSize := 15
		expected := 10
		actual := Hash4(key, hashSize)
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
}
