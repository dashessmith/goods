import os
import sys
import subprocess
from time import gmtime, strftime


def exec(cmd):
    print(f'executing cmd: {cmd}')
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
    if len(sys.argv) > 1:
        exec(f'git commit -m "{sys.argv[1]}"')
    else:
        exec("git commit --allow-empty-message -m ''")

    exec("git push")

    exec(f'gh release create {newtag} -n ""')

    exec("git pull --tags")


if __name__ == "__main__":
    push()
