import os
import sys
from time import gmtime, strftime


def exec(cmd):
    code = os.system(cmd)
    if code != 0:
        sys.exit(code)


def push():
    exec("go mod tidy")
    exec("gofumpt -s -w .")
    exec("git add .")
    exec("git commit --allow-empty-message -m ''")
    exec(
        f'gh release create v{strftime("%Y.%m.%d.%H.%M.%S", gmtime())} --notes "casual release"')
    exec("git pull --tags")


if __name__ == "__main__":
    push()
