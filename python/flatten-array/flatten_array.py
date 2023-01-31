def flatten(iterable):
    new_list = []
    for item in iterable:
        if type(item) is list:
            new_list.extend(flatten(item))
        elif item is not None:
            new_list.append(item)
    return new_list
