"""
This exercise stub and the test suite contain several enumerated constants.

Since Python 2 does not have the enum module, the idiomatic way to write
enumerated constants has traditionally been a NAME assigned to an arbitrary,
but unique value. An integer is traditionally used because it’s memory
efficient.
It is a common practice to export both constants and functions that work with
those constants (ex. the constants in the os, subprocess and re modules).

You can learn more here: https://en.wikipedia.org/wiki/Enumerated_type
"""

SUBLIST = 0
SUPERLIST = 1
EQUAL = 2
UNEQUAL = 3


def sublist(list_one, list_two):
    if list_one == list_two:
        return EQUAL
    if contains(list_two, list_one):
        return SUBLIST
    if contains(list_one, list_two):
        return SUPERLIST
    return UNEQUAL


def contains(list_one, list_two):
    for i in range(len(list_one) - len(list_two) + 1):
        if list_one[slice(i, i + len(list_two))] == list_two:
            return True
    return False
