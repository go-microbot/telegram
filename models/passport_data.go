package models

// PassportData information about Telegram Passport data shared with the bot by the user.
// https://core.telegram.org/bots/api#passportdata.
type PassportData struct {
	// Array with information about documents
	// and other Telegram Passport elements that was shared with the bot.
	Data []EncryptedPassportElement `json:"data"`
	// Encrypted credentials required to decrypt the data.
	Credentials EncryptedCredentials `json:"credentials"`
}
