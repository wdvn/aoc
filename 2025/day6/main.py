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
    matrix = extract()
    for col in range(0, len(matrix[-1])):
        sum = int(matrix[0][col])
        for i in range(1, len(matrix) - 1):
            if matrix[-1][col] == "+":
                sum += int(matrix[i][col])
            else:
                sum *= int(matrix[i][col])
        out += sum
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
        vals=[]
        for n in nums:
            try:
                vals.append(int(n))
            except:
                continue
        val = vals[0]
        for i  in range(1,len(vals)):
            if op == "+":
                val += vals[i]
            else:
                val *= vals[i]
        out += val
    print(out)


if __name__ == '__main__':
    # part1()
    part2()
