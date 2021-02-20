package leveling

import (
	"io"
	"time"
)

type Writer struct {
	writer io.Writer
	onceWriteSize int
	interval time.Duration
}

var _ io.Writer = (*Writer)(nil)

func New(writer io.Writer, interval time.Duration, onceWriteSize int) *Writer {
	return &Writer{
		writer:     writer,
		onceWriteSize: onceWriteSize,
		interval: interval,
	}
}

func (w *Writer) Write(p []byte) (int, error) {
	remaining := p
	total := 0

	latest := time.Now()
	sleepOverhead := time.Duration(0)
	for {
		if len(remaining) <= w.onceWriteSize {
			n, err := w.writer.Write(remaining)
			return n+total, err
		}

		buf := remaining[0:w.onceWriteSize]
		n, err := w.writer.Write(buf)
		if err != nil {
			return total, err
		}

		sleptAt := time.Now()
		wait := w.interval - sleptAt.Sub(latest) - sleepOverhead
		time.Sleep(wait)
		latest = time.Now()
		sleepOverhead = latest.Sub(sleptAt) - wait

		total += n
		remaining = remaining[n:]
	}
}
