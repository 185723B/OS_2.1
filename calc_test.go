package calc

import (
    "testing"
)

func TestAdd(t *testing.T) {
    patterns := []struct {
        a        int
        b        int
        expected int
    }{
        {1, 2, 4},
        {10, -3, 7},
        {-4, -12, -16},
    }

    for idx, pattern := range patterns {
        actual := Add(pattern.a, pattern.b)
        if pattern.expected != actual {
            t.Errorf("pattern %d: want %d, actual %d", idx, pattern.expected, actual)
        }
    }
}