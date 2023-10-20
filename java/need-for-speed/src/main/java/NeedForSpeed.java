class NeedForSpeed {

    private int speed;
    private int batteryDrain;
    private int distanceDriven = 0;
    private int battery = 100;
    
    NeedForSpeed(int speed, int batteryDrain) {
        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public boolean batteryDrained() {
        return this.battery <= 0;
    }

    public int distanceDriven() {
        return this.distanceDriven;
    }

    public void drive() {
        if (this.battery >= this.batteryDrain) {
            this.distanceDriven += this.speed;
            this.battery -= this.batteryDrain;
        }
    }

    public static NeedForSpeed nitro() {
        return new NeedForSpeed(50, 4);
    }
}

class RaceTrack {

    private int distance;
    
    RaceTrack(int distance) {
        this.distance = distance;
    }

    public boolean tryFinishTrack(NeedForSpeed car) {
        while (!car.batteryDrained()) {
            car.drive();
            if (car.distanceDriven() >= this.distance) {
                return true;
            }
        }
        return false;
    }
}
