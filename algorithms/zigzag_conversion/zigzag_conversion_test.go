package zigzag_conversion

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	fmt.Println(splitGroup("ABCDE", 4))
	fmt.Println(convert("ABCDE", 4))
}
