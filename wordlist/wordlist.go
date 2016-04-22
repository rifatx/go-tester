package wordlist

import (
	"fmt"
)

func GenerateFrom(letters string, len int) {
	contains := func(l []rune, e rune) bool {
		for _, a := range l {
			if a == e {
				return true
			}
		}

		return false
	}

	var letterCombo func(l []rune, len int, p string, wc chan string)
	letterCombo = func(l []rune, len int, p string, wc chan string) {
		if len == 0 {
			wc <- p
			return
		}

		for _, r := range l {
			letterCombo(l, len-1, p+string(r), wc)
		}
	}

	comboWrapper := func(l []rune, len int, p string, wc chan string) {
		letterCombo(l, len, "", wc)
		close(wc)
	}

	l := make([]rune, 0)

	for _, r := range letters {
		if !contains(l, r) {
			l = append(l, r)
		}
	}

	r := make(chan string, 0)

	go comboWrapper(l, len, "", r)

	for w := range r {
		fmt.Println(w)
	}
}
