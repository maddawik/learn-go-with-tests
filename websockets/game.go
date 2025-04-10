package poker

import "io"

type Game interface {
	Play(numberOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}
