package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type KazuateGame struct {
	Ans int
}

var mes []string = []string{"%dより小さいです", "%dで正解です", "%dより大きいです"}

func NewKazuateGame(Ans int) *KazuateGame {
	return &KazuateGame{Ans: Ans}
}

func (self *KazuateGame) Input(v int) int {
	if self.Ans < v {
		return 0
	} else if self.Ans == v {
		return 1
	} else {
		return 2
	}
}

func (self *KazuateGame) Run(Output_Writer io.Writer, input_Reader *bufio.Scanner) {
	for {
		fmt.Fprintf(Output_Writer, "入力してください>")
		input_Reader.Scan()
		input_text := input_Reader.Text()
		input_text_int, err := strconv.Atoi(input_text)
		if err != nil {
			log.Fatalln(err)
		}
		result_code := self.Input(input_text_int)
		fmt.Fprintf(Output_Writer, mes[result_code]+"\n", input_text_int)
		if result_code == 1 {
			break
		}
	}
}

func main() {
	g := NewKazuateGame(rand.Intn(100))
	g.Run(os.Stdout, bufio.NewScanner(os.Stdin))
}
