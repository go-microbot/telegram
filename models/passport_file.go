package models

// PassportFile represents a file uploaded to Telegram Passport.
// Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
// https://core.telegram.org/bots/api#passportfile.
type PassportFile struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// Unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUID string `json:"file_unique_id"`
	// File size.
	FileSize int32 `json:"file_size"`
	// Unix time when the file was uploaded.
	FileDate int32 `json:"file_date"`
}
