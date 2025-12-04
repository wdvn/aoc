def read_file(path) -> str:
    with open(path, 'r') as file:
        content = file.read()
    return content


def total_adjacent(grid, i, j):
    count = 0
    directs = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]
    for d in directs:
        x = i + d[0]
        if x < 0 or x >= len(grid): continue
        y = j + d[1]
        if y >= len(grid[i]) or y < 0: continue
        if grid[x][y] == '@':
            count += 1
    return count


def part2():
    count, s, result = 0, "", 0
    while True:
        count, s = part1(s)
        if not count: break
        result += count

    print(result)


def part1(s: str):
    if not s:
        raw = read_file("input.txt")
    else:
        raw = s
    grid = raw.split("\n")
    count = 0
    replace = []
    for i, row in enumerate(grid):
        for j, c in enumerate(row):
            if c == '@' and total_adjacent(grid, i, j) < 4:
                count += 1
                replace.append((i, j))
    for item in replace:
        i, j = item
        grid[i] = grid[i][:j] + 'x' + grid[i][j + 1:]
    return count, "\n".join(grid)


def part2():
    count, s, result = 0, "", 0
    while True:
        count, s = part1(s)
        if not count: break
        result += count

    print(result)


if __name__ == '__main__':
    # part1()
    part2()
