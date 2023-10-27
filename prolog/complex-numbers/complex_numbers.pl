real((R1, _), R) :- R = R1.
imaginary((_, I1), I) :- I = I1.

conjugate((R1, I1), Conjugate) :- Conjugate = (R1, -I1).
abs((R1, I1), Abs) :- Abs = sqrt(R1 * R1 + I1 * I1).

add((R1, I1), (R2, I2), CAdd) :- CAdd = (R1 + R2, I1 + I2).
sub((R1, I1), (R2, I2), CSub) :- CSub = (R1 - R2, I1 - I2).

mul((R1, I1), (R2, I2), CMul) :- CMul = (R1 * R2 - I1 * I2, R1 * I2 + R2 * I1).
div((R1, I1), (R2, I2), CDiv) :- CDiv = ((R1 * R2 + I1 * I2) / (R2 * R2 + I2 * I2), (R2 * I1 - R1 * I2) / (R2 * R2 + I2 * I2)).
