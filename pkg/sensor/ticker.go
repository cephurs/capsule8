// Copyright 2017 Capsule8, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sensor

import (
	"fmt"
	"time"

	"github.com/capsule8/capsule8/pkg/expression"
	"github.com/capsule8/capsule8/pkg/sys"
	"github.com/capsule8/capsule8/pkg/sys/perf"
)

// TickerEventTypes defines the field types that can be used with filters on
// ticker telemetry events.
var TickerEventTypes = expression.FieldTypeMap{
	"seconds":     expression.ValueTypeSignedInt64,
	"nanoseconds": expression.ValueTypeSignedInt64,
}

// TickerTelemetryEvent is a telemetry event generated by the ticker event
// source.
type TickerTelemetryEvent struct {
	TelemetryEventData

	Seconds     int64
	Nanoseconds int64
}

// CommonTelemetryEventData returns the telemtry event data common to all
// telemetry events for a chargen telemetry event.
func (e TickerTelemetryEvent) CommonTelemetryEventData() TelemetryEventData {
	return e.TelemetryEventData
}

func (s *Subscription) decodeTickerEvent(
	sample *perf.SampleRecord,
	data perf.TraceEventSampleData,
) (interface{}, error) {
	var e TickerTelemetryEvent

	e.Init(s.sensor)
	e.Seconds = data["seconds"].(int64)
	e.Nanoseconds = data["nanoseconds"].(int64)

	return e, nil
}

// RegisterTickerEventFilter registers a ticker event filter with a
// subscription.
func (s *Subscription) RegisterTickerEventFilter(
	interval int64,
	filter *expression.Expression,
) {
	if interval <= 0 {
		s.logStatus(
			fmt.Sprintf("Invalid ticker interval: %d", interval))
		return
	}

	monitor := s.sensor.Monitor()
	eventID := monitor.RegisterExternalEvent("ticker",
		s.decodeTickerEvent)

	done := make(chan struct{})

	es, err := s.addEventSink(eventID, filter, TickerEventTypes)
	if err != nil {
		s.logStatus(
			fmt.Sprintf("Invalid filter expression for ticker filter: %v", err))
		monitor.UnregisterEvent(eventID)
		return
	}
	es.unregister = func(es *eventSink) {
		monitor.UnregisterEvent(es.eventID)
		close(done)
	}

	go func() {
		ticker := time.NewTicker(time.Duration(interval))
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				monoNow := sys.CurrentMonotonicRaw()
				sampleID := perf.SampleID{
					Time: uint64(monoNow),
				}
				now := time.Now()
				data := perf.TraceEventSampleData{
					"seconds":     int64(now.Unix()),
					"nanoseconds": int64(now.UnixNano()),
				}
				monitor.EnqueueExternalSample(
					eventID, sampleID, data)

			}
		}
	}()
}
