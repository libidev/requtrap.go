package http

import (
	"sync"
	"time"

	"github.com/libidev/requtrap.go/cli/config"
	"github.com/libidev/requtrap.go/pkg/constants"
)

type CircuitBreaker struct {
	Circuit map[string]*Circuit
}

type Circuit struct {
	failureCount     int
	failureThreshold int
	state            string
	mutex            sync.Mutex
	timeout          time.Duration
	lastFailureTime  time.Time
}

func (cb *CircuitBreaker) Call(fn func(s config.Service) error, service config.Service, timeout time.Duration) error {

	circuit, ok := cb.Circuit[service.Upstream]
	if !ok {
		cb.Circuit[service.Upstream] = &Circuit{
			failureThreshold: 3, //failureThreshold,
			state:            "CLOSED",
			timeout:          timeout,
		}

		circuit = cb.Circuit[service.Upstream]
	}

	circuit.mutex.Lock()

	switch circuit.state {
	case "OPEN":
		// If in OPEN state, check if the timeout period has passed
		if time.Since(circuit.lastFailureTime) > circuit.timeout {
			circuit.state = "HALF-OPEN" // Move to HALF-OPEN state
		} else {
			circuit.mutex.Unlock()
			return constants.ERR_CIRCUIT_OPEN
		}
	case "HALF-OPEN":
		// Allow one request to test if service is back
		circuit.mutex.Unlock()
		err := fn(service)
		if err != nil {
			circuit.recordFailure()
			return err
		}
		circuit.reset()
		return nil
	}

	circuit.mutex.Unlock()

	err := fn(service)
	if err != nil {
		circuit.recordFailure()
		return err
	}

	circuit.reset()
	return nil
}

func (cb *Circuit) recordFailure() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.failureCount++
	if cb.failureCount >= cb.failureThreshold {
		cb.state = "OPEN"
		cb.lastFailureTime = time.Now()
	}
}

func (cb *Circuit) reset() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.failureCount = 0
	cb.state = "CLOSED"
}
