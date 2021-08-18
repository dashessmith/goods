import os
import sys
import subprocess
from time import gmtime, strftime


def exec(cmd):
    code = os.system(cmd)
    if code != 0:
        sys.exit(code)


def push():
    maxtag = subprocess.check_output(
        "git describe --abbrev=0 --tags").decode('utf8').strip()
    tokens = maxtag.split('.')
    lasttoken = int(tokens[-1]) + 1
    newtag = ".".join(tokens[:-1] + [str(lasttoken)])
    exec("go mod tidy")
    exec("gofumpt -s -w .")
    exec("git add .")
    exec("git commit --allow-empty-message -m ''")
    exec("git push")
    exec(
        f'gh release create {newtag} --notes "casual release"')
    exec("git pull --tags")


if __name__ == "__main__":
    push()
