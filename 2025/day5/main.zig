const utils = @import("utils");
const std = @import("std");

const alloc = std.heap.page_allocator;
const strings = utils.strings{ .alloc = alloc };
const Input = struct { intervals: []const [2]i64, ids: []i64 };
fn extract(s: []u8) !Input {
    const iter = try strings.split(s, "\n\n");
    const rawIntervals = iter[0];
    const rawIds = iter[1];
    var intervals = try std.ArrayList([2]i64).initCapacity(alloc, 0);
    for (try strings.split(rawIntervals, "\n")) |rIterval| {
        const tmp = try strings.split(rIterval, "-");
        const min = try std.fmt.parseInt(i64, tmp[0], 10);
        const max = try std.fmt.parseInt(i64, tmp[1], 10);
        try intervals.append(alloc, [2]i64{ min, max });
    }
    var ids = try std.ArrayList(i64).initCapacity(alloc, 0);
    for (try strings.split(rawIds, "\n")) |id| {
        const n = try std.fmt.parseInt(i64, id, 10);
        try ids.append(alloc, n);
    }
    return .{ .intervals = intervals.items, .ids = ids.items };
}

fn part1(data: Input) void {
    var count: i64 = 0;
    for (data.ids) |id| {
        for (data.intervals) |interval| {
            if (interval[0] <= id and id <= interval[1]) {
                count += 1;
                break;
            }
        }
    }
    std.debug.print("{}", .{count});
}
fn less_fn(context: *anyopaque, a: [2]i64, b: [2]i64) bool {
    _ = context;
    if (a[0] != b[0]) a[0] < b[0];
    return a[1] < b[1];
}
fn sort_intervals(input_intervals: []const [2]i64) ![]const [2]i64 {
    const sorted: []const[2]i64 = input_intervals[0..];
    //TODO: fix compile error
    std.mem.sort([]const [2]i64, sorted, null, less_fn);
    return sorted;
}

fn merge_interval(intervals: []const [2]i64) ![][2]i64 {
    const sorted = try sort_intervals(intervals);
    var list = utils.slice.init([2]i64, alloc, 0);
    try list.append(alloc, sorted[0]);
    for (sorted[1..]) |interval| {
        const last = utils.slice.last([2]i64, list);
        if (interval[0] <= last[1]) {
            utils.slice.assign(list, list.items.len - 1, [2]i64{ last[0], @max(interval[0], last[1]) });
        } else {
            try list.append(alloc, [2]i64{ interval[0], interval[1] });
        }
    }
    return list.items;
}

fn part2(data: Input) void {
    const merged = try merge_interval(data.intervals);
    var count: i64 = 0;
    for (merged) |interval| {
        count += interval[1] - interval[0] + 1;
    }
    std.debug.print("{}", .{count});
}

pub fn main() !void {
    const detail = try utils.read_file(alloc, "input.txt");
    const input = try extract(detail);
    part1(input);
    part2(input);
}
