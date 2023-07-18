def saddle_points(matrix):
    points = []
    for row_index, row in enumerate(matrix):
        if len(row) != len(matrix[0]):
            raise ValueError("irregular matrix")
        for column_index, column in enumerate(row):
            if is_tallest_in_row(row_index, column_index, matrix) and is_smallest_in_column(row_index, column_index, matrix):
                points.append({"row": row_index + 1, "column": column_index + 1})
    return points
            
def is_tallest_in_row(row_index, column_index, matrix):
    for element in matrix[row_index]:
        if element > matrix[row_index][column_index]:
            return False
    return True


def is_smallest_in_column(row_index, column_index, matrix):
    for row in matrix:
        if row[column_index] < matrix[row_index][column_index]:
            return False
    return True
    