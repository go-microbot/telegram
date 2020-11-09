package models

// Dice represents an animated emoji that displays a random value.
// https://core.telegram.org/bots/api#dice.
type Dice struct {
	// Emoji on which the dice throw animation is based.
	Emoji string `json:"emoji"`
	// Value of the dice, 1-6 for “🎲” and “🎯” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji.
	Value int32 `json:"value"`
}
