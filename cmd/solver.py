#!.venv/bin/python

import importlib
import os
import sys
__project_root = os.path.realpath(os.path.join(os.path.dirname(__file__), os.path.pardir))
sys.path = [__project_root] + sys.path  # noqa


def main(args: [str]):
    task_paths = args[1:]
    task_modules = []

    for path in task_paths:
        python_path = path.replace("/", ".") + ".main"

        try:
            task_modules.append(importlib.import_module(python_path))
        except ImportError:
            print(python_path)
            print(sys.path)
            raise

    for path, mod in zip(task_paths, task_modules):
        mod.solver(path)


if __name__ == "__main__":
    main(args=sys.argv)
