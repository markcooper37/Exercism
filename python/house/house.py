CHARACTERS = ["house that Jack built.", "malt", "rat", "cat", "dog", "cow with the crumpled horn", "maiden all forlorn", "man all tattered and torn", "priest all shaven and shorn", "rooster that crowed in the morn", "farmer sowing his corn", "horse and the hound and the horn"]


ACTIONS = ["lay in", "ate", "killed", "worried", "tossed", "milked", "kissed", "married", "woke", "kept", "belonged to"]


def recite(start_verse, end_verse):
    return [make_verse(verse_number) for verse_number in range(start_verse, end_verse + 1)]


def make_verse(verse_number):
    verse = "This is the " + CHARACTERS[verse_number - 1]
    for index in range(verse_number - 2, -1, -1):
        verse += " that " + ACTIONS[index] + " the " + CHARACTERS[index]
    return verse
