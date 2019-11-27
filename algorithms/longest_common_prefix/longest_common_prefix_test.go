package longest_common_prefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(LongestCommonPrefix([]string{}))
}
