package pagerduty

import (
	"errors"

	"github.com/PagerDuty/go-pagerduty"
)

func isNotFoundError(err error) bool {
	var aerr pagerduty.APIError

	if errors.As(err, &aerr) {
		return aerr.NotFound()
	}
	return false
}
