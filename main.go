// this is for maintaining purposes
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	getNumbers *string
	userAnswer *string
)

func answerOne() {
	keyAnswer := "/etc/hostname"
	if keyAnswer != *userAnswer {
		fmt.Println("Sorry, That is not a valid answer.")
	} else {
		fmt.Println(`Correct!, Here's the flag : LabProyekFun{Ngonfig hostname penting bro, ngasi tau aja}`)
	}

}

func answerTwo() {
	keyAnswer := `~/directory-banyak-bgt-bjir/hayolo130/anjaymabarprofesional`
	if keyAnswer != *userAnswer {
		fmt.Println("Sorry, That is not a valid answer.")
	} else {
		fmt.Println(`Correct!, Here's the flag : LabProyekFun{Kenapa unik?, ya gapapa pengen aja bilang gitu}`)
	}

}

func answerThree() {
	filename := "answerFile.txt"
	fileStat, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	waktuModifikasi := fileStat.ModTime()
	keyAnswerFirst := waktuModifikasi.Format("01-02")
	valueFile, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer valueFile.Close()
	var massage string
	readData := bufio.NewReader(valueFile)
	for {
		dataCatch, _, err := readData.ReadLine()
		if err == io.EOF {
			break
		}
		massage += string(dataCatch)
	}
	keyAnswerSecond := "hallo gays, ini IG ku dnfpr_int, follow peliss"

	if keyAnswerFirst == *userAnswer {
		if keyAnswerSecond != massage {
			fmt.Println("sorry, that is not a valid answer for the file")
		} else {
			fmt.Println(`Correct!, Here's the flag : LabProyekFun{SELAMAT DATANG TEK 62 :)))))))!!!!!}`)
		}
	} else {
		fmt.Println("Sorry, That is not a valid answer for the date.")
	}

}
func main() {
	getNumbers = flag.String("n", "false", "use -n to choose soal")
	userAnswer = flag.String("j", "false", "use -j to answer the question")
	flag.Parse()

	if *getNumbers == "1" {
		answerOne()
	} else if *getNumbers == "2" {
		answerTwo()
	} else if *getNumbers == "3" {
		answerThree()
	}

}
