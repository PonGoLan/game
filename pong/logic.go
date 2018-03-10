package pong

func BallPlayerCollision(ball *Ball, player *Player) {
	// TODO fix this
	if ball.PredictedX == player.X || ball.PredictedX == player.X-1 || ball.PredictedX == player.X+1 {
		if ball.PredictedY >= player.Y && ball.PredictedY <= player.Y+player.Size {
			ball.VectorX *= -1
		}
	}
}
