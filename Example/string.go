package main

import s "strings"
import "fmt"

func main() {
	var p = fmt.Println

	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"a", "b", "c"}, "&"))
	p("Repeat: ", s.Repeat("abc", 3))
	p("Replace: ", s.Replace("good", "o", "0", 1))
	p("Split", s.Split("137-6608-3349", "-"))
	p("ToLower: ", s.ToLower("TesT"))
	p("ToUpper: ", s.ToUpper("tESt"));
}
