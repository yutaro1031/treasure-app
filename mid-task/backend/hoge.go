package main

import (
	"fmt"
	"github.com/yutaro1031/mid-task/httputil"
)

func main() {
	hoge, _ := httputil.Search("https://example.com")
	fmt.Println(hoge)
}
