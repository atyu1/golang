package dog

import (
  "testing"
  "fmt"
)

func BenchmarkYears(b *testing.B) {
  for i:=0;i<b.N; i++ {
    Years(10)
  }
}

func TestYears(t *testing.T) {
  humanYears := Years(10)
  if humanYears != 70 {
    t.Error("Expected is 70 and we got",humanYears)
  }
}

func ExampleYears() {
  fmt.Println(Years(10))
  //Output: 70
}
