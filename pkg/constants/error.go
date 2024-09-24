package constants

import "errors"

var (
	ERR_CIRCUIT_OPEN error = errors.New("circuit breaker is OPEN")
)
