package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

const (
	PlayerPrompt         = "Please enter the number of players:"
	BadPlayerInputErrMsg = "Bad value receieved for number of players, please try again with a number"
	BadWinnerInputErrMsg = "Bad value receieved for winner, please try again with the format: `{PlayerName} wins`"
)

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Play(numberOfPlayers, cli.out)
	winner, err := extractWinner(cli.readLine())
	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputErrMsg)
		return
	}

	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	userInputSlice := strings.Split(userInput, " ")

	if len(userInputSlice) != 2 {
		return "", fmt.Errorf("wrong number of tokens for winner input, %v", BadWinnerInputErrMsg)
	}

	if userInputSlice[1] != "wins" {
		return "", fmt.Errorf("must use `wins` for second token in winner input, %v", BadWinnerInputErrMsg)
	}

	return userInputSlice[0], nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
