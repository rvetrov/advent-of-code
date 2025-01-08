def non_empty_lines(data: str) -> [str]:
    return data.strip().splitlines()


def parse_numbers(data: str, sep: str = None) -> [int]:
    return [int(_) for _ in data.split(sep)]
