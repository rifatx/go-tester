package eksisozluk

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"runtime"
	"strings"
	"sync"
)

func DumpEntries(urlPattern string, values []string) {
	mp := runtime.GOMAXPROCS(4)
	defer runtime.GOMAXPROCS(mp)

	var wg sync.WaitGroup

	for _, v := range values {
		wg.Add(1)

		go func(v string) {
			getHtmlAttrib := func(l []html.Attribute, key string) (string, bool) {
				for _, a := range l {
					if a.Key == key {
						return a.Val, true
					}
				}

				return "", false
			}

			defer wg.Done()

			doc, err := goquery.NewDocument(fmt.Sprintf(urlPattern, v))

			if err != nil {
				fmt.Println(err)
				return
			}

			var sbr bytes.Buffer
			s := doc.Find("div.content *")

			for i := 0; i < len(s.Nodes); {
				n := s.Nodes[i]
				t := s.Eq(i).Text()

				//fmt.Println("[", n.Type, "]", n.Attr, t)

				switch n.Type {
				case 1:
					sbr.WriteString(t)
					i++
				case 3:
					c, ok := getHtmlAttrib(n.Attr, "class")

					switch {
					case ok && c == "b":
						sbr.WriteString("`" + t + "`")
						i = i + 2
					case ok && c == "ab":
						i++
						n := s.Nodes[i]
						t := s.Eq(i).Text()

						if t == "*" {
							t = ""
						}

						c, _ = getHtmlAttrib(n.Attr, "data-query")

						sbr.WriteString("`" + t + ":" + c + "`")
						i = i + 2
					case ok && c == "url":
						c, _ = getHtmlAttrib(n.Attr, "href")
						sbr.WriteString("[" + c + " " + t + "]")
						i = i + 2
					default:
						sbr.WriteString("\n")
						i++
					}
				}
			}

			u := doc.Find("a.entry-author").Text()
			t := doc.Find("span[itemprop='name']").Text()
			d := doc.Find("a.entry-date").Text()

			fmt.Println(t)
			fmt.Println(strings.TrimSpace(sbr.String()))
			fmt.Println(u)
			fmt.Println(d)
		}(v)
	}

	wg.Wait()
}
