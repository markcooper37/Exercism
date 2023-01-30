import re


def parse(markdown):
    lines = markdown.split('\n')
    result = ''
    in_list = False
    for line in lines:
        line = add_italic(add_bold(line))
        if re.match('#{1,6} (.*)', line):
            line = f'<h{line.find(" ")}>{line[line.find(" ")+1:]}</h{line.find(" ")}>'
        elif line[0:2] == '* ':
            line = '<li>' + line[2:] + '</li>'
            if not in_list:
                in_list = True
                line = '<ul>' + line
        else:
            line = '<p>' + line + '</p>'
            if in_list:
                result += '</ul>'
                in_list = False
        result += line
    if in_list:
        result += '</ul>'
    return result


def add_bold(line):
    match = re.match('(.*)__(.*)__(.*)', line)
    if match:
        line = match.group(1) + '<strong>' + \
            match.group(2) + '</strong>' + match.group(3)
    return line


def add_italic(line):
    match = re.match('(.*)_(.*)_(.*)', line)
    if match:
        line = match.group(1) + '<em>' + match.group(2) + \
            '</em>' + match.group(3)
    return line
