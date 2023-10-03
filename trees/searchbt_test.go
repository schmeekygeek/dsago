package trees

import (
	"fmt"
	"testing"
)

func TestSearchBt(t *testing.T) {
  node := CreateBT()
  fmt.Println(searchBT(&node, 5))
}
