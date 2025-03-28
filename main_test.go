package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKazuateGame(t *testing.T) {

	g := NewKazuateGame(1)

	assert.EqualValues(t, &KazuateGame{Ans: 1}, g)
}

func TestKazuateGameInput(t *testing.T) {
	g := NewKazuateGame(1)

	assert.Equal(t, 2, g.Input(0))
	assert.Equal(t, 1, g.Input(1))
	assert.Equal(t, 0, g.Input(2))
	assert.Equal(t, 0, g.Input(3))
}
func TestKazuateGameRun(t *testing.T) {
	g := NewKazuateGame(1)

	stdout_buf := new(bytes.Buffer)
	g.Run(stdout_buf,
		bufio.NewScanner(strings.NewReader("0\n3\n2\n1\n3\n")))
	buf_str := stdout_buf.String()

	assert.Equal(
		t,
		"入力してください>0より大きいです\n入力してください>3より小さいです\n入力してください>2より小さいです\n入力してください>1で正解です\n",
		buf_str,
	)
}
