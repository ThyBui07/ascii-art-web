package drawing

import (
	"bufio"
	"log"
	"os"
)

type char struct {
	charDec int
	charArr [][]byte
}

func Display(s string, bannerName string) string {
	//get the dictionary
	bannerPath := "drawing/banners/" + bannerName + ".txt"
	dictionary := ReadBannerFile(bannerPath)
	//
	var display [8][]byte

	for i := 0; i < len(s); i++ {
		for _, v := range dictionary {
			if s[i] == byte(v.charDec) {
				for i := 0; i < 8; i++ {
					display[i] = append(display[i], v.charArr[i]...)
				}
			}
		}
	}

	res := ""
	for i := 0; i < 8; i++ {
		for _, v := range display[i] {
			res += string(v)
		}
		res += "\n"
	}

	return res

}

func ReadBannerFile(bannerPath string) []char {
	//handle 2nd os arg for txt name

	file, err := os.Open(bannerPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var line []byte
	var lines [][]byte

	//read txt file from beginning [except 1 and last \n line]
	for i := 1; i <= 855; i++ {
		line, err = r.ReadBytes('\n')
		if len(line) != 1 {
			num := len(line) - 1
			lines = append(lines, line[:num])
		} else {
			continue
		}
	}

	//create dictionary 95 chars (32-126)
	count := 0
	decN := 32
	var dictionary []char
	var charArr [][]byte
	for _, v := range lines {
		count++
		charArr = append(charArr, v)
		if count == 8 {
			x := char{decN, charArr}
			dictionary = append(dictionary, x)
			count = 0
			charArr = [][]byte{}
			decN += 1
		}
	}

	return dictionary

}
