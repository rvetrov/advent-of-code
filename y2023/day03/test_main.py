import main

TEST_CASE_1 = """467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
"""

def test_solve_v1():
    res = main.solve_v1(TEST_CASE_1.split())
    assert res == 4361


def test_solve_v2():
    res = main.solve_v2(TEST_CASE_1.split())
    assert res == 467835
