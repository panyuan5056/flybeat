package cacl

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/spf13/cast"
)

func TestCacl(t *testing.T) {
	t.Log("hello world")
	tmp := [][]byte{}
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		t := cast.ToString(time.Now().UnixNano())
		for _, tag := range []string{"c", "b", "a"} {
			tmp = append(tmp, []byte(tag+t))
		}
		sort.Slice(tmp, func(i int, j int) bool { return bytes.Compare(tmp[i], tmp[j]) == -1 })
		for _, a := range tmp {
			fmt.Println(string(a))
		}
		time.Sleep(time.Duration(10))
	}
}
