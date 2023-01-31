import re


def decode(string):
    return "".join([int(phrase[:-1]) * phrase[-1] if len(phrase) > 1 else phrase[0] for phrase in re.findall("\d*[a-zA-Z ]", string)])


def encode(string):
    letters, last_index = "", 0
    for index, letter in enumerate(string):
        if index == len(string) - 1 or letter != string[index + 1]:
            letters += (str(index + 1 - last_index) if index != last_index else "") + letter
            last_index = index + 1
    return letters
