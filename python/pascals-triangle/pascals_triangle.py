def rows(row_count):
    if row_count < 0:
        raise ValueError("number of rows is negative")
    if row_count == 0:
        return []
    if row_count == 1:
        return [[1]]
    all_rows = rows(row_count - 1)
    prev_row = [0] + all_rows[len(all_rows)-1] + [0]
    all_rows.append([sum(prev_row[i:i+2]) for i in range(row_count)])
    return all_rows
