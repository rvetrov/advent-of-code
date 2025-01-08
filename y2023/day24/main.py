import collections
import dataclasses
import random

import numpy as np
import scipy

from internal.py import task, task_io


@dataclasses.dataclass
class Point:
    x: int
    y: int
    z: int


@dataclasses.dataclass
class Vector:
    x: int
    y: int
    z: int


"""
https://www.reddit.com/r/adventofcode/comments/18q40he/2023_day_24_part_2_a_straightforward_nonsolver/
(dy'- dy) X + ( dx-dx') Y + (  y- y') DX + ( x'-  x) DY =  x' dy' -  y' dx' -   x  dy +   y  dx
(v2y-v1y) X + (v1x-v2x) Y + (p1y-p2y) DX + (p2x-p1x) DY = p2x v2y - p2y v2x - p1x v1y + p1y v1x
X, Y, DX, DY - initial position and velocity

Equation:
a0 X + a1 Y + a2 Z + a3 DX + a4 DY + a5 DZ = b
"""


def __make_equation(p1x, p1y, v1x, v1y, p2x, p2y, v2x, v2y) -> (list[int], int):
    a = [0] * 4
    a[0] = v2y - v1y
    a[1] = v1x - v2x
    a[2] = p1y - p2y
    a[3] = p2x - p1x
    b = p2x * v2y - p2y * v2x - p1x * v1y + p1y * v1x
    return a, b


def __make_equations(p1: Point, v1: Vector, p2: Point, v2: Vector) -> (list[list[int]], list[int]):
    a, b_xy = __make_equation(p1.x, p1.y, v1.x, v1.y, p2.x, p2.y, v2.x, v2.y)
    a_xy = [0] * 6
    a_xy[0], a_xy[1], a_xy[3], a_xy[4] = a

    a, b_xz = __make_equation(p1.x, p1.z, v1.x, v1.z, p2.x, p2.z, v2.x, v2.z)
    a_xz = [0] * 6
    a_xz[0], a_xz[2], a_xz[3], a_xz[5] = a
    return [a_xy, a_xz], [b_xy, b_xz]


def solve_v2(data: str) -> int:
    lines = task_io.non_empty_lines(data)

    points: list[Point] = []
    vecs: list[Vector] = []
    for line in lines:
        parts = line.split(" @ ")
        points.append(Point(*task_io.parse_numbers(parts[0], ", ")))
        vecs.append(Vector(*task_io.parse_numbers(parts[1], ", ")))

    results = collections.Counter()
    for _ in range(100):
        ind = random.sample(range(len(points)), k=4)
        aa, bb = [], []
        for i in range(1, 4):
            a, b = __make_equations(points[ind[0]], vecs[ind[0]], points[ind[i]], vecs[ind[i]])
            aa.extend(a)
            bb.extend(b)

        a = np.array(aa, dtype=np.longdouble)
        b = np.array(bb, dtype=np.longdouble)

        x = scipy.linalg.solve(a, b)
        res = int(np.round(x[0] + x[1] + x[2]))
        results[res] += 1

    # avoiding errors related to floating numbers accuracy
    return results.most_common(1)[0][0]


def solver(data_dir_path: str):
    task.solve(data_dir_path, solver_v2=solve_v2)


if __name__ == "__main__":
    solver("y2023/day24")
