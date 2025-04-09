package poker

type Game interface {
	Play(int)
	Finish(string)
}
