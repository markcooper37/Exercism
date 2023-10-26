class Clock {

    private int minutes;

    Clock(int hours, int minutes) {
        this.minutes = (((60 * hours + minutes + 1440) % 1440) + 1440) % 1440;
    }

    void add(int minutes) {
        this.minutes = (((this.minutes + minutes) % 1440) + 1440) % 1440;
    }

    @Override
    public String toString() {
        return String.format("%02d:%02d", this.minutes / 60, this.minutes % 60);
    }

    @Override
    public boolean equals(Object obj) {
        if (obj == this) return true;
        if (obj == null) return false;
        if (getClass() != obj.getClass()) return false;
        Clock clock = (Clock) obj;
        return this.minutes == clock.minutes;
    }

    @Override
    public int hashCode() {
        return this.minutes;
    }

}