package models

// InputFile represents the contents of a file to be uploaded.
// Must be posted using multipart/form-data in the usual way
// that files are uploaded via the browser.
// https://core.telegram.org/bots/api#inputfile.
type InputFile []byte
