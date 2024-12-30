from internal.py import task


def nearby_symbols(lines: list[str], ii: int, jj: int) -> dict[str, set[tuple[int, int]]]:
    symbols = {}
    for i in range(-1, 2):
        for j in range(-1, 2):
            if i == 0 and j == 0:
                continue
            ni = ii + i
            nj = jj + j
            if 0 <= ni < len(lines) and 0 <= nj < len(lines[ni]):
                ch = lines[ni][nj]
                if ch != '.' and not ch.isdigit():
                    symbols.setdefault(ch, set()).add((ni, nj))
    return symbols


def solve_v1(lines: list[str]) -> int:
    res = 0
    lines = [_.strip() for _ in lines if _]
    for i, line in enumerate(lines):
        is_engine_part = False
        number = 0
        for j, ch in enumerate(line):
            if ch.isdigit():
                number = number * 10 + int(ch)
                if not is_engine_part and nearby_symbols(lines, i, j):
                    is_engine_part = True
            else:
                if is_engine_part:
                    res += number
                is_engine_part = False
                number = 0
        if is_engine_part:
            res += number

    return res


def solve_v2(lines: list[str]) -> int:
    lines = [_.strip() for _ in lines if _]
    gears = {}
    for i, line in enumerate(lines):
        number = 0
        nearby_gears = set()
        for j, ch in enumerate(line):
            if ch.isdigit():
                number = number * 10 + int(ch)
                symbols = nearby_symbols(lines, i, j)
                nearby_gears.update(symbols.get("*", []))
            else:
                for gear in nearby_gears:
                    gears.setdefault(gear, []).append(number)
                nearby_gears = set()
                number = 0
        for gear in nearby_gears:
            gears.setdefault(gear, []).append(number)

    res = 0
    for _, numbers in gears.items():
        if len(numbers) == 2:
            res += numbers[0] * numbers[1]
    return res


def solver(dir_path: str):
    task.solve(dir_path, solver_v1=solve_v1, solver_v2=solve_v2)


if __name__ == "__main__":
    solver("y2023/day03")
