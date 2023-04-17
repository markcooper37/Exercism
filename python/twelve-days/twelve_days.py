POSITION = ('first', 'second', 'third', 'fourth', 'fifth', 'sixth', 'seventh', 'eighth', 'ninth', 'tenth', 'eleventh', 'twelfth')
GIFTS = ('a Partridge in a Pear Tree.', 'two Turtle Doves, ', 'three French Hens, ', 'four Calling Birds, ', 'five Gold Rings, ', 'six Geese-a-Laying, ', 'seven Swans-a-Swimming, ', 'eight Maids-a-Milking, ', 'nine Ladies Dancing, ', 'ten Lords-a-Leaping, ', 'eleven Pipers Piping, ', 'twelve Drummers Drumming, ')

def recite(start_verse, end_verse):
    verses = []
    for verse_number in range (start_verse, end_verse + 1):
        verse = f'On the {POSITION[verse_number - 1]} day of Christmas my true love gave to me: '
        for value in range (verse_number - 1, 0, -1):
            verse += GIFTS[value]
        if verse_number != 1:
            verse += 'and '
        verse += GIFTS[0]
        verses.append(verse)
    return verses
