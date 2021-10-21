package main

import (
	"flag"
	"fmt"
	"github.com/SheltonZhu/biliLiverOb/ob"
)

var mid = flag.Int("mid", 7706705, "bilibili uid")

func main() {
	flag.Parse()
	u, err := ob.NewOb(*mid)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s isLiving? %v\n", u.Name, u.IsLiving())
	}
}
