class CalculatorConundrum {
    public String calculate(int operand1, int operand2, String operation) {
        if (operation == null) {
            throw new IllegalArgumentException("Operation cannot be null");
        } else if (operation == "") {
            throw new IllegalArgumentException("Operation cannot be empty");
        }
        String start = operand1 + " " + operation + " " + operand2 + " = ";
        switch (operation) {
            case "+":
                return start + (operand1 + operand2);
            case "*":
                return start + (operand1 * operand2);
            case "/":
                try {
                    return start + (operand1 / operand2);
                } catch (ArithmeticException e) {
                    throw new IllegalOperationException("Division by zero is not allowed", e);
                }
            default:
                throw new IllegalOperationException("Operation '" + operation + "' does not exist");
        }
    }
}
