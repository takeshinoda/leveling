package leveling

import (
	"io"
	"time"

	"golang.org/x/time/rate"
)

type Writer struct {
	writer io.Writer
	onceWriteSize int
	limiter *rate.Limiter
}

var _ io.Writer = (*Writer)(nil)

func New(writer io.Writer, interval time.Duration, onceWriteSize int) *Writer {
	limiter := rate.NewLimiter(rate.Every(interval), 1)

	return &Writer{
		writer:     writer,
		onceWriteSize: onceWriteSize,
		limiter: limiter,
	}
}

func NewTimesPerSecond(writer io.Writer, secSplitNum int, bytesPerSecond int ) *Writer {
	interval := time.Second / time.Duration(secSplitNum)
	onceWriteSize := bytesPerSecond / secSplitNum
	return New(writer, interval, onceWriteSize)
}

func (w *Writer) Write(p []byte) (int, error) {
	_ = w.limiter.Reserve()

	remaining := p
	total := 0
	for {
		r := w.limiter.Reserve()

		if len(remaining) <= w.onceWriteSize {
			n, err := w.writer.Write(remaining)
			return n+total, err
		}

		buf := remaining[0:w.onceWriteSize]
		n, err := w.writer.Write(buf)
		if err != nil {
			return total, err
		}

		total += n
		remaining = remaining[n:]

		time.Sleep(r.Delay())
	}
}
