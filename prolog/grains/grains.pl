square(SquareNumber, Value) :- (SquareNumber > 0, SquareNumber < 65 -> Value is 2**(SquareNumber-1)).

total(Value) :- Value is 2**64-1.
