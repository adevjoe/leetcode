package zigzag_conversion

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	fmt.Println(splitGroup("ABCDE", 4))
	fmt.Println(Convert("ABCDE", 4))
}
