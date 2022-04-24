package common

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	fmt.Println(PathConcat("1path", "/2path", "3path/", "/4path/", "/path5", "path6"))
}
