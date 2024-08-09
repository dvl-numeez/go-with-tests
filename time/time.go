package time

import (
	"time"

	"github.com/dvl-numeez/go-with-tests/time/poker"
)

type SpyBlindAlerter struct {
	alerts []poker.ScheduledAlert
}
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts,poker.ScheduledAlert{duration,amount} )
}
