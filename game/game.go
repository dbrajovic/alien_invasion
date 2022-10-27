package game

type Game struct {
	//	engine
}

func New() *Game {
	println("new game created")
	return nil
}

func (g *Game) Run() {
	println("game run")
}
