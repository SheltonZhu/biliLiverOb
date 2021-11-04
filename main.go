package main

import (
	"flag"
	"fmt"
	"github.com/SheltonZhu/biliLiverOb/ob"
	"github.com/mattn/go-runewidth"
	"strconv"
	"unicode/utf8"
)

var mid = flag.Int("mid", 7706705, "bilibili uid")
var list = flag.Bool("l", false, "ob list")
var all = flag.Bool("all", false, "ob all")
var obList = map[string]int{
	"七海Nana7mi":   434334701,
	"召唤师yami":     487902,
	"阿梓从小就很可爱":    7706705,
	"星宮汐Official": 402417817,
}

func FormatWidth(name string) int {
	maxWidth := 20
	realWidth := runewidth.StringWidth(name)
	fakeWidth := utf8.RuneCount([]byte(name))
	fmtWidth := maxWidth + fakeWidth - realWidth
	return fmtWidth
}
func main() {
	flag.Parse()
	if *list {
		fmt.Printf("%-"+strconv.Itoa(FormatWidth("name"))+"s %-20s\n", "name", "mid")
		for name, mid := range obList {
			fmt.Printf("%-"+strconv.Itoa(FormatWidth(name))+"s %-20d\n", name, mid)
		}
		return
	}
	fmt.Printf("%-"+strconv.Itoa(FormatWidth("name"))+"s %-20s\n", "name", "isLiving")
	if *all {
		for _, mid := range obList {
			u, err := ob.NewOb(mid)
			if err != nil {
				panic(err)
			} else {
				if u.IsLiving() {
					fmt.Printf("\033[1;32;42m%-"+strconv.Itoa(FormatWidth(u.Name))+"s %v\033[0m\n", u.Name, u.IsLiving())
				} else {
					fmt.Printf("%-"+strconv.Itoa(FormatWidth(u.Name))+"s %v\n", u.Name, u.IsLiving())
				}
			}
		}
		return
	}

	u, err := ob.NewOb(*mid)
	if err != nil {
		panic(err)
	} else {
		if u.IsLiving() {
			fmt.Printf("\033[1;32;42m%-"+strconv.Itoa(FormatWidth(u.Name))+"s %v\033[0m\n", u.Name, u.IsLiving())
		} else {
			fmt.Printf("%-"+strconv.Itoa(FormatWidth(u.Name))+"s %v\n", u.Name, u.IsLiving())
		}
	}
}
