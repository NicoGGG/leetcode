from typing import List


def searchMatrix(matrix: List[List[int]], target: int) -> bool:
    height = len(matrix) - 1
    i = 0
    while i <= height:
        width = len(matrix[0]) - 1
        mid_h = (i + height) // 2
        j = 0
        mid_w = (j + width) // 2
        while j <= width:
            mid_w = (j + width) // 2
            if matrix[mid_h][mid_w] == target:
                return True
            if matrix[mid_h][mid_w] < target:
                j = mid_w + 1
            else:
                width = mid_w - 1
        if matrix[mid_h][mid_w] < target:
            i = mid_h + 1
        else:
            height = mid_h - 1
    return False


matrix1 = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
target1 = 3
matrix2 = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
target2 = 13
matrix3 = [[1, 3]]
target3 = 3

print(True, searchMatrix(matrix1, target1))
print(False, searchMatrix(matrix2, target2))
print(True, searchMatrix(matrix3, target3))
