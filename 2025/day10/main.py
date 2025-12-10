import itertools
import math
import re


def read_file(path) -> str:
    with open(path, 'r') as file:
        content = file.read()
    return content


def extract():
    lines = read_file("input.txt").split("\n")
    out =[]
    for line in lines:
        expect = re.findall(r'\[.*]',line)[0].strip('[]')
        buttons = re.findall(r'\(.+\)',line)[0]
        btns = []
        for btn in buttons.split(" "):
            btns.append([int(n) for n in re.findall(r'\d+',btn)])
        out.append((expect,btns))
    return out
def dfs(expected, chains, buttons, btn, out) -> int:
    for b in btn:
        if chains[b] == '.':
            chains[b] = '#'
        else:
            chains[b] = '.'
    if "".join(chains) == expected: return out
    for btn in buttons:
        return dfs(expected, chains, buttons, btn, out + 1)
    return out


def part1():
    out = 0
    for expected, buttons in extract():
        lm = 1e300 * 1e100
        for btn in buttons:
            t = dfs(expected, ["."] * len(expected), buttons, btn, 1)
            lm = min(lm, t)
            out+=min
    print(out)


# 
if __name__ == '__main__':
    part1()
