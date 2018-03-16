package instances

import (
	"log"
	"time"

	pong "github.com/PonGoLan/game/shared"
)

type Instance struct {
	// instance informations
	ticks                    int8
	numberOfPlayersConnected int8

	// game simulation
	game *pong.Game
}

func (i *Instance) GetGame() *pong.Game {
	return i.game
}

// Run : runs an instance
func (i *Instance) Run() {
	second := time.Tick(time.Second)
	aTick := time.Tick(time.Second / 64)

	ball := i.game.Ball
	player1 := i.game.Players[0]
	player2 := i.game.Players[1]

	for 1 == 1 {
		select {
		case <-second:
			log.Printf("[TICKS/s]: %d\n", i.ticks)
			i.ticks = 0
		case <-aTick:
			i.ticks++
			pong.BallPlayerCollision(ball, player1)
			pong.BallPlayerCollision(ball, player2)
			ball.Move(i.game)
		default:
		}
	}
}

func (i *Instance) AddPlayer() int8 {
	currentPlayer := i.numberOfPlayersConnected

	i.numberOfPlayersConnected++

	return currentPlayer
}

// CreateInstance : Creates & Initialize a new game instance
func CreateInstance() *Instance {
	instance := new(Instance)

	game := pong.NewGame()

	board := pong.NewBoard()
	player1 := pong.NewPlayer(0, board)
	player2 := pong.NewPlayer(1, board)

	game.Players[0] = player1
	game.Players[1] = player2

	ball := pong.NewBall(board)

	game.Ball = ball

	instance.game = game

	return instance
}
