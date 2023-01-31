def find(search_list, value):
    lower_bound, upper_bound = 0, len(search_list) - 1
    while upper_bound >= lower_bound:
        index = (upper_bound + lower_bound) // 2
        if value > search_list[index]:
            lower_bound = index + 1
        elif value < search_list[index]:
            upper_bound = index - 1
        else:
            return index
    raise ValueError("value not in array")
            
