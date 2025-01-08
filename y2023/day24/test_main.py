from y2023.day24 import main

TEST_CASE_1 = """
19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
"""


def test_solve_v2():
    res = main.solve_v2(TEST_CASE_1)
    assert res == 47
