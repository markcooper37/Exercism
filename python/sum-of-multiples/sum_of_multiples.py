def sum_of_multiples(limit, multiples):
    return sum(set().union(*[range(multiple, limit, multiple) if multiple > 0 else [] for multiple in multiples]))
