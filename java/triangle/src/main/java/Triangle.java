class Triangle {

    private double a;
    private double b;
    private double c;

    Triangle(double side1, double side2, double side3) throws TriangleException {
        if (side1 + side2 < side3 || side1 + side3 < side2 || side2 + side3 < side1) {
            throw new TriangleException();
        } else if (side1 == 0 || side2 == 0 || side3 == 0) {
            throw new TriangleException();
        }
        this.a = side1;
        this.b = side2;
        this.c = side3;
    }

    boolean isEquilateral() {
        return this.a == this.b && this.b == this.c;
    }

    boolean isIsosceles() {
        return this.a == this.b || this.a == this.c || this.b == this.c;
    }

    boolean isScalene() {
        return this.a != this.b && this.a != this.c && this.b != this.c;
    }

}
