package errorz

import "errors"

var (
	InvalidPartnerFormat = errors.New("invalid partner format")
	PartnerNotFound      = errors.New("partner not found")
)
