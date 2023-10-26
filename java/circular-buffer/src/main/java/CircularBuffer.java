class CircularBuffer<T> {

    private Object[] buffer;
    private int readPosition = 0;
    private int writePosition = 0;

    CircularBuffer(final int size) {
        this.buffer = new Object[size];
    }

    T read() throws BufferIOException {
        if (buffer[readPosition] == null) {
            throw new BufferIOException("Tried to read from empty buffer");
        }
        T data = (T) buffer[this.readPosition];
        buffer[this.readPosition] = null;
        this.readPosition = (this.readPosition + 1) % buffer.length;
        return data;
    }

    void write(T data) throws BufferIOException {
        if (buffer[writePosition] != null) {
            throw new BufferIOException("Tried to write to full buffer");
        }
        buffer[writePosition] = data;
        this.writePosition = (this.writePosition + 1) % buffer.length;
    }

    void overwrite(T data) {
        if (buffer[this.writePosition] != null) {
            this.readPosition = (this.readPosition + 1) % buffer.length;
        }
        buffer[writePosition] = data;
        this.writePosition = (this.writePosition + 1) % buffer.length;
    }

    void clear() {
        this.buffer = new Object[this.buffer.length];
        this.readPosition = 0;
        this.writePosition = 0;
    }

}