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
    inputs = read_file("./input.txt").split("\n\n")
    intervals, ids = inputs[0], inputs[1]
    cv = []
    for item in intervals.split("\n"):
        r = item.split("-")
        cv.append((int(r[0]), int(r[1])))
    new_ids = []
    for item in ids.strip().split("\n"):
        new_ids.append(int(item))
    return cv, new_ids


def merge_intervals(intervals):
    sorted_intervals = sorted(intervals, key=lambda x: x[0])
    merged = [sorted_intervals[0]]
    for current_start, current_end in sorted_intervals[1:]:
        last_start, last_end = merged[-1]
        if current_start <= last_end:
            merged[-1] = (last_start, max(last_end, current_end))
        else:
            merged.append((current_start, current_end))

    return merged


def part1():
    intervals, ids = extract()
    out = 0
    for id in ids:
        if is_valid(id, intervals):
            out += 1
    print(out)


def part2():
    intervals, ids = extract()
    result = merge_intervals(intervals)
    out = 0
    for interval in result:
        out += interval[1] - interval[0] + 1
    print(out)


if __name__ == '__main__':
    part2()
