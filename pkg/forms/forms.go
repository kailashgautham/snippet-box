package forms

import (
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}
