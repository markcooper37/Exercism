def equilateral(sides):
    return is_triangle(sides) and sides[0] == sides[1] == sides[2]


def isosceles(sides):
    return is_triangle(sides) and sides[0] == sides[1] or sides[0] == sides[2] or sides[1] == sides[2]


def scalene(sides):
    return is_triangle(sides) and not isosceles(sides)


def is_triangle(sides):
    sides.sort()
    return sides[0] > 0 and sides[0] + sides[1] >= sides[2]