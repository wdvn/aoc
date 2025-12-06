import re
import math

def read_file(path) -> str:
    with open(path, 'r') as file:
        content = file.read()
    return content


def is_valid(num, group):
    for interval in group:
        if interval[0] <= num <= interval[1]:
            return True
    return False


def extract():
    lines = read_file("./input.txt").split("\n")
    return lines


def part1():
    out = 0
    lines = extract()
    matrix = []
    for line in lines[:-1]:
        matrix.append([int(n) for n in re.findall(r"\d+", line)])
    ops= re.findall(r"\S",lines[-1])
    matrix=[list(row) for row in zip(*matrix)]
    for i,op in enumerate(ops):
        if op=="+":
            out+= sum(matrix[i])
        else:
            out+= math.prod(matrix[i])
    print(out)


def part2():
    out = 0
    lines = extract()
    while len(lines[-1]) > 0:
        size = 0
        op = ""
        for i in range(len(lines[-1]) - 1, -1, -1):
            size += 1
            if lines[-1][i] != " ":
                op = lines[-1][i]
                lines[-1] = lines[-1][:-size]
                break
        news = []
        for i in range(0, len(lines) - 1):
            line = lines[i]
            col = line[-size:]
            lines[i] = line[:-size]
            news.append(col)
        nums = [""] * (size + 1)
        for i in range(len(news[0]) - 1, -1, -1):
            for raw in news:
                if raw[i] != " ":
                    nums[i] += raw[i]
        vals=[int(n) for n in nums if n.isdigit()]
        val = vals[0]
        for v  in vals[1:]:
            val = val + v if op == "+" else val * v
        out += val
    print(out)

if __name__ == '__main__':
    part1()
    part2()