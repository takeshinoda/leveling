package leveling

import (
	"testing"
	"time"
)

func TestNewSplitSecond(t *testing.T) {
	actual := NewTimesPerSecond(nil, 10, 10000)
	expected := Writer{
		writer:        nil,
		onceWriteSize: 1000,
		interval:       time.Second / time.Duration(10),
	}

	if actual.onceWriteSize != expected.onceWriteSize || actual.interval != expected.interval {
		t.Errorf("It returns an unexpected writer: %v, expected %v", actual, expected)
	}
}

type recordWriter struct {
	eachByBytes [][]byte
	timestamps []time.Time
}

func (r *recordWriter) Write(b []byte) (int, error) {
	r.eachByBytes = append(r.eachByBytes, b)
	r.timestamps = append(r.timestamps, time.Now())
	return len(b), nil
}

func TestWriter_Write(t *testing.T) {
	t.Run("Writing is done at regular intervals.", func(t *testing.T) {
		recorder := &recordWriter{}
		splitNum := 10
		writer := NewTimesPerSecond(recorder, splitNum, 10000)
		data := make([]byte, 100000)

		n, err := writer.Write(data)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if n != len(data) {
			t.Errorf("unexpected written length %d bytes", n)
		}

		if len(recorder.eachByBytes) != 100 {
			t.Errorf("unexpected the times of written data = %d", len(recorder.eachByBytes))
		}

		interval := time.Second / time.Duration(splitNum)
		margin := interval / 10
		min := interval - margin
		max := interval + margin

		previous := recorder.timestamps[0]
		for i, timestamp := range recorder.timestamps[1:] {
			diff := timestamp.Sub(previous)
			if diff < min {
				t.Errorf("The %d times write interval is too short = %s, expected: >= %s", i+1, diff, min)
			}
			if diff > max {
				t.Errorf("The %d times write interval is too long = %s, expected: <= %s", i+1, diff, max)
			}
			previous = timestamp
		}
	})
}