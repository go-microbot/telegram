package models

// GameHighScore represents one row of the high scores table for a game.
// https://core.telegram.org/bots/api#gamehighscore.
type GameHighScore struct {
	// Position in high score table for the game.
	Position int32 `json:"position"`
	// User.
	User User `json:"user"`
	// Score.
	Score int32 `json:"score"`
}
