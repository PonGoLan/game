package pong

func BallPlayerCollision(ball *Ball, player *Player) {
	if ball.PredictedY >= player.Y && ball.PredictedY <= player.Y+player.Size {
		if ball.PredictedX == player.X {
			ball.VectorX *= -1
		}
	}
}
