def read_file(path) -> str:
    with open(path, 'r') as file:
        content = file.read()
    return content


def extract():
    lines = read_file("input.txt").split("\n")
    return lines


def part1():
    grid = extract()
    beams = grid[0].replace("S", "|", 1)
    count = 0
    for row in grid[1:]:
        tmp = list(beams)
        for i, char in enumerate(row):
            if char == '^' and beams[i] == '|':
                tmp[i] = '.'
                tmp[i + 1] = '|'
                tmp[i - 1] = '|'
                count += 1
        beams = "".join(tmp)
        print(row, count, row.count("^"))

    print(count)

def part2():
    grid = [list(line) for line in extract()]
    dp = [[0] * len(grid[0]) for _ in range(len(grid))]
    dp[0][grid[0].index("S")] = 1
    for i, row in enumerate(grid):
        for j, char in enumerate(row):
            if char == "^":
                dp[i][j + 1] += dp[i - 1][j]
                dp[i][j - 1] += dp[i - 1][j]
            else:
                dp[i][j] += dp[i - 1][j]
    print(sum(dp[-1]))


if __name__ == '__main__':
    part2()
