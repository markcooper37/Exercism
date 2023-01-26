# Score categories.
# Change the values as you see fit.
YACHT = 'Yacht'
ONES = 1
TWOS = 2
THREES = 3
FOURS = 4
FIVES = 5
SIXES = 6
FULL_HOUSE = 'Full House'
FOUR_OF_A_KIND = 'Four of a Kind'
LITTLE_STRAIGHT = 'Little Straight'
BIG_STRAIGHT = 'Big Straight'
CHOICE = 'Choice'


def score(dice, category):
    dice.sort()
    if category == CHOICE:
        return sum(dice)
    elif category == YACHT:
        return (dice[0] == dice[4]) * 50
    elif category == FOUR_OF_A_KIND:
        return (dice[0] == dice[3] or dice[1] == dice[4]) * 4 * dice[3]
    elif category == FULL_HOUSE:
        return (3 * dice[0] + 2 * dice[4] if  dice[0] == dice[2] and dice[2] != dice[3] and dice[3] == dice[4]
        else 2 * dice[0] + 3 * dice[4] if dice[0] == dice[1] and dice[1] != dice[2] and dice[2] == dice[4]
        else 0)
    elif category == LITTLE_STRAIGHT:
        return (is_straight(dice) and dice[0] == 1) * 30
    elif category == BIG_STRAIGHT:
        return (is_straight(dice) and dice[0] == 2) * 30
    else:
        return len([die for die in dice if die == category]) * category


def is_straight(dice):
    return dice[0] + 1 == dice[1] and dice[1] + 1 == dice[2] and dice[2] + 1 == dice[3] and dice[3] + 1 == dice[4]
