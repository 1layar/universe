package breaker

import (
	"github.com/1layar/universe/pkg/logger"
	"github.com/sony/gobreaker/v2"
)

func New() *gobreaker.CircuitBreaker[[]byte] {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		logger := logger.GetLogger()
		isTrip := counts.Requests >= 3 && failureRatio >= 0.6

		if isTrip {
			logger.With(
				"failureRatio", failureRatio,
			).Info("circuit breaker is trip")
		}
		return isTrip
	}

	return gobreaker.NewCircuitBreaker[[]byte](st)
}
