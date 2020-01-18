package diffcode

import "testing"

func TestWordCount(t *testing.T) {
    total := WordCount("5, 5")
    if total != 1 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 1)
    }
}