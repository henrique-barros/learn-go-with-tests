package application

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(text string) string {
	return strings.Replace(text, " wins", "", 1)
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
