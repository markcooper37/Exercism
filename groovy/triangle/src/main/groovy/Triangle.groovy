class Triangle {

    def a
    def b
    def c

    Triangle(a, b, c) {
        this.a = a
        this.b = b
        this.c = c
    }

    boolean isEquilateral() {
        return isValid() && (a == b && b == c)
    }

    boolean isIsosceles() {
        return isValid() && (a == b || a == c || b == c)
    }

    boolean isScalene() {
        return isValid() && !isIsosceles()
    }

    private boolean isValid() {
        return (a > 0 && b > 0 && c > 0) && (a + b > c && a + c > b && b + c > a)
    }

}
