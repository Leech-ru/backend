package errorz

import "errors"

var (
	InvalidMainPageFormat = errors.New("invalid main page format")
	MainPageNotFound      = errors.New("main page not found")
)
