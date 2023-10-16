
class BirdWatcher {
    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public int[] getLastWeek() {
        return this.birdsPerDay;
    }

    public int getToday() {
        return this.birdsPerDay[this.birdsPerDay.length-1];
    }

    public void incrementTodaysCount() {
        this.birdsPerDay[this.birdsPerDay.length-1] += 1;
    }

    public boolean hasDayWithoutBirds() {
        for (int count: this.birdsPerDay) {
            if (count == 0) {
                return true;
            }
        }
        return false;
    }

    public int getCountForFirstDays(int numberOfDays) {
        int count = 0;
        for (int i = 0; i < Math.min(numberOfDays, this.birdsPerDay.length); i++) {
            count += this.birdsPerDay[i];
        }
        return count;
    }

    public int getBusyDays() {
        int busyDays = 0;
        for (int count: this.birdsPerDay) {
            if (count >= 5) {
                busyDays += 1;
            }
        }
        return busyDays;
    }
}
