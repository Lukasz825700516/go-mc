package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
)

func processClass(f *os.File) {
	exp, err := regexp.Compile(`(\s*)public void write\(.*?\)`)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)
	var ending []byte
	var endingLen int
	inMethod := false

	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		if inMethod {
			ok := false
			for i, b := range l[0:endingLen] {
				ok = ok || (ending[i] != b)
			}
			ok = !ok

			if ok {
				inMethod = false
				break
			}

 		}

		if inMethod {
			fmt.Println(string(l))
		}

		mch := exp.FindSubmatchIndex(l)
		if mch == nil {
			continue
		}

		inMethod = true
		endingLen = mch[3] - mch[2] + 1
		ending = make([]byte, endingLen)
		copy(ending, l[mch[2]:mch[3]])
		ending[endingLen - 1] = '}'
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	i := 0

	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		s := string(l)

		if len(l) == 0 {
			i += 1
			continue
		}

		var t string
		switch i {
		case 0:
			t = "login"
		case 1:
			t = "status"
		case 2:
			t = "game"
		}
		
		p := path.Join("net/minecraft/network/protocol", t, s + "Packet.java")
		f, err := os.Open(p)
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}

		fmt.Printf("%s\n", s)
		processClass(f)
		fmt.Printf("\n")
	}
}
