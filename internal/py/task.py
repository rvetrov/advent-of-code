import time
import typing


INPUT_FILE_NAME = "input.big"
OUTPUT_FILE_BASE_NAME = "output"
Solver = typing.Callable[[[str]], typing.Any]


def solve(dir_path: str, solver_v1: Solver = None, solver_v2: Solver = None):
    input_path = f"{dir_path}/{INPUT_FILE_NAME}"
    with open(input_path, "rt") as input_file:
        lines = input_file.readlines()

    for solver, file_name_ext in (
        (solver_v1, "v1"),
        (solver_v2, "v2"),
    ):
        if solver is None:
            continue

        start_ns = time.monotonic()
        res = solver(lines)
        duration = time.monotonic() - start_ns

        output_path = f"{dir_path}/{OUTPUT_FILE_BASE_NAME}.{file_name_ext}"
        with open(output_path, "wt") as result_file:
            print(res, file=result_file)
        print(f"{dir_path}: {input_path} -> {output_path}, {duration:.3f}s")
