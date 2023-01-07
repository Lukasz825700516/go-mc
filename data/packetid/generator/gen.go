package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
)

func processWrite(body [][]byte) {
	expVarInt, err := regexp.Compile(`^\s*\$\$0\.writeVarInt\((.*?)\);`)
	expInt, err := regexp.Compile(`^\s*\$\$0\.writeInt\((.*?)\);`)
	expFloat, err := regexp.Compile(`^\s*\$\$0\.writeFloat\((.*?)\);`)
	expDouble, err := regexp.Compile(`^\s*\$\$0\.writeDouble\((.*?)\);`)
	expByte, err := regexp.Compile(`^\s*\$\$0\.writeByte\((.*?)\);`)
	expLong, err := regexp.Compile(`^\s*\$\$0\.writeLong\((.*?)\);`)
	expEnum, err := regexp.Compile(`^\s*\$\$0\.writeEnum\((.*?)\);`)
	expUUID, err := regexp.Compile(`^\s*\$\$0\.writeUUID\((.*?)\);`)
	expBlockPos, err := regexp.Compile(`^\s*\$\$0\.writeBlockPos\((.*?)\);`)
	expVarLong, err := regexp.Compile(`^\s*\$\$0\.writeBlockPos\((.*?)\);`)
	expUtf, err := regexp.Compile(`^\s*\$\$0\.writeUtf\((.*?)\);`)
	expShort, err := regexp.Compile(`^\s*\$\$0\.writeShort\((.*?)\);`)
	expByteArray, err := regexp.Compile(`^\s*\$\$0\.writeByteArray\((.*?)\);`)
	expVarIntArray, err := regexp.Compile(`^\s*\$\$0\.writeVarIntArray\((.*?)\);`)
	expResourceLocation, err := regexp.Compile(`^\s*\$\$0\.writeResourceLocation\((.*?)\);`)
	expResourceKey, err := regexp.Compile(`^\s*\$\$0\.writeResourceKey\((.*?)\);`)
	expBytes, err := regexp.Compile(`^\s*\$\$0\.writeBytes\((.*?)\);`)
	expId, err := regexp.Compile(`^\s*\$\$0\.writeId\((.*?)\);`)
	expNbt, err := regexp.Compile(`^\s*\$\$0\.writeNbt\((.*?)\);`)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range body {
		var mch []int

		mch = expVarInt.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found varint: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expInt.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found int: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expFloat.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found float: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expDouble.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found double: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expByte.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found byte: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expLong.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found long: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expEnum.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found enum: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expUUID.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found UUID: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expBlockPos.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found BlockPos: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expVarLong.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found VarLong: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expUtf.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found Utf: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expShort.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found Short: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expByteArray.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found ByteArray: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expVarIntArray.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found VarIntArray: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expResourceLocation.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found ResourceLocation: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expResourceKey.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found ResourceKey: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expBytes.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found Bytes: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expId.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found Id: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
		mch = expNbt.FindSubmatchIndex(l)
		if mch != nil {
			fmt.Printf("found Nbt: %s\n", string(l[mch[2]:mch[3]]))
			continue
		}
	}
}

func processClass(f *os.File) {
	exp, err := regexp.Compile(`(\s*)public void write\(.*?\)`)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)
	var ending []byte
	var endingLen int
	inMethod := false

	var body [][]byte

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
			lCopy := make([]byte, len(l))
			copy(lCopy, l)

			fmt.Println("to parse: ", string(l))
			body = append(body, lCopy)
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

		body = make([][]byte, 0, 16)
	}

	processWrite(body)
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
