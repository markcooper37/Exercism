package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	mu        sync.Mutex
	reader    io.Reader
	reads     int
	bytesRead int
}

type writeCounter struct {
	mu           sync.Mutex
	writer       io.Writer
	writes       int
	bytesWritten int
}

type readWriteCounter struct {
	*writeCounter
	*readCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer, writes: 0, bytesWritten: 0}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader, reads: 0, bytesRead: 0}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	readCounter := readCounter{reader: readwriter, reads: 0, bytesRead: 0}
	writeCounter := writeCounter{writer: readwriter, writes: 0, bytesWritten: 0}
	return &readWriteCounter{readCounter: &readCounter, writeCounter: &writeCounter}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	bytesRead, err := rc.reader.Read(p)
	rc.bytesRead += bytesRead
	rc.reads++
	return bytesRead, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return int64(rc.bytesRead), rc.reads
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	bytesWritten, err := wc.writer.Write(p)
	wc.bytesWritten += bytesWritten
	wc.writes++
	return bytesWritten, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return int64(wc.bytesWritten), wc.writes
}
