import itertools
import re


def read_file(path) -> str:
    with open(path, 'r') as file:
        content = file.read()
    return content


def extract():
    lines = read_file("input.txt").split("\n")
    out = []
    for line in lines:
        out.append(tuple([int(n) for n in re.findall(r'\d+', line)]))
    return out


def part1():
    points = extract()
    # make n(n-1) pairs from n points
    pairs = list(itertools.combinations(points, 2))

    out = 0
    # make graph from pairs
    for pair in pairs:
        area = abs(pair[0][0] - pair[1][0] + 1) * abs(pair[0][1] - pair[1][1] + 1)
        if area > out:
            out = area
    print(out)


def part2():
    points = extract()
    pairs = list(itertools.combinations(points, 2))

    out = 0
    # make graph from pairs
    for pair in pairs:
        area = abs(pair[0][0] - pair[1][0] + 1) * abs(pair[0][1] - pair[1][1] + 1)
        if area > out:
            out = area
    print(out)
if __name__ == '__main__':
    part2()
