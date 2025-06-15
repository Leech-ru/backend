package errorz

import "errors"

var (
	InvalidCosmeticsFormat = errors.New("invalid cosmetics format; WARNING: if you see this message, contact in support")
	CosmeticsNotFound      = errors.New("cosmetics not found")
)
