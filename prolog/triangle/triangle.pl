triangle_nondegenerate(Side1, Side2, Side3) :- Side1 > 0, Side2 > 0, Side3 > 0.

triangle_inequality(Side1, Side2, Side3) :- Side1 + Side2 > Side3, Side2 + Side3 > Side1, Side3 + Side1 > Side2.

triangle_valid(Side1, Side2, Side3) :- triangle_nondegenerate(Side1, Side2, Side3), triangle_inequality(Side1, Side2, Side3).

triangle_isosceles(Side1, Side2, Side3) :- Side1 =:= Side2; Side2 =:= Side3; Side1 =:= Side3.

triangle(Side1, Side2, Side3, "equilateral") :- triangle_valid(Side1, Side2, Side3), Side1 =:= Side2, Side2 =:= Side3.

triangle(Side1, Side2, Side3, "isosceles") :- triangle_valid(Side1, Side2, Side3), triangle_isosceles(Side1, Side2, Side3).

triangle(Side1, Side2, Side3, "scalene") :- triangle_valid(Side1, Side2, Side3), \+ triangle_isosceles(Side1, Side2, Side3).

