"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language: https://en.wikipedia.org/wiki/Guido_van_Rossum
"""

EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2

def bake_time_remaining(elapsed_bake_time):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """

    return EXPECTED_BAKE_TIME - elapsed_bake_time

def preparation_time_in_minutes(number_of_layers):
    """Calculate the total preparation time.

    :param number_of_layers: int - number of layers in the lasagna.
    :return: int - total time to prepare all layers (in minutes) derived from 'PREPARATION_TIME'.

    Function that takes the number of layers of the lasagna as an argument and
    returns how many minutes it will take to prepare all layers based on the
    `PREPARATION_TIME`.
    """

    return number_of_layers * PREPARATION_TIME

def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """Calculate the time spent cooking so far.

    :param number_of_layers: int - number of layers in the lasagna.
    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - total time spent cooking (in minutes) derived from 'PREPARATION_TIME'.

    Function that takes the number of layers of the lasagna and the actual minutes the
    lasagna has been in the oven as arguments and returns the number of minutes spent
    cooking so far based on the `PREPARATION_TIME`.
    """

    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
