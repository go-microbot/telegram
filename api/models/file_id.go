package models

import "github.com/go-microbot/telegram/query"

// FileID represents file ID query model.
type FileID struct {
	// File identifier to get info about.
	FileID query.ParamString `query:"file_id"`
}
