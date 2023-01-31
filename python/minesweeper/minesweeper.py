def annotate(minefield):
    new_minefield = []
    for row_index, row in enumerate(minefield):
        new_row = ""
        if len(row) != len(minefield[0]):
            raise ValueError("The board is invalid with current input.")
        for column_index, value in enumerate(row):
            if value == "*":
                new_row += "*"
            elif value == " ":
                mine_count = count_mines(minefield, row_index, column_index)
                new_row += str(mine_count) if mine_count > 0 else " "
            else:
                raise ValueError("The board is invalid with current input.")
        new_minefield.append(new_row)
    return new_minefield


def count_mines(minefield, row, column):
    count = 0
    for new_row in range(row - 1, row + 2):
        for new_column in range(column - 1, column + 2):
            if new_row == row and new_column == column:
                continue
            if 0 <= new_row < len(minefield) and 0 <= new_column < len(minefield[0]) and minefield[new_row][new_column] == "*":
                    count += 1
    return count
