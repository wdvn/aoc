const std = @import("std");

const utils = @import("utils");
const allocator = std.heap.page_allocator;

fn part1() !void {
    const contents = try utils.read_file("./input.txt", allocator);
    var iter = std.mem.splitScalar(u8, contents, '\n');
    var sum: i32 = 0;
    while (iter.next()) |line| {
        std.debug.print("{s}\n", .{line});

        var num: i32 = 0;
        for (line, 0..) |x, i| {
            var j = i + 1;
            while (j < line.len) {
                const y = line[j];
                j += 1;
                if (std.ascii.isDigit(x) and std.ascii.isDigit(y)) {
                    const s: i32 = (x - '0') * 10 + (y - '0');
                    if (num < s) {
                        num = s;
                    }
                }
            }
        }
        sum += num;
    }
    std.debug.print("{}", .{sum});
}

fn part2() !void {
    const contents = try utils.read_file("./input.txt", allocator);
    var iter = std.mem.splitScalar(u8, contents, '\n');
    var sum: i64 = 0;
    while (iter.next()) |line| {
        sum += try findNum(line);
    }
    std.debug.print("{}", .{sum});
}
fn findNum(line: []const u8) i64 {
    var num: [12]u8 = undefined;
    var start: usize = 0;

    for (0..11) |i| {
        const end = line.len - (12 - i);
        std.debug.print("{s}: {} {} => ", .{ line, start, end });

        const c = maxInRange(line, start, end);
        std.debug.print("{} {} \n", .{ c.@"0", c.@"1" });

        num[i] = c.@"0";
        start = c.@"1" + 1;
    }
    std.debug.print("{s}", .{num});
    const rest = try convert(num);
    return rest;
}

fn maxInRange(arr: []const u8, start: usize, end: usize) struct { u8, usize } {
    var n: u8 = arr[start];
    var index: usize = start;

    for (start..end) |j| {
        if (n < arr[j]) {
            n = arr[j];
            index = j;
        }
    }
    return .{ n, index };
}

fn convert(num: [12]u8) !i64 {
    const o: []const u8 = num[0..];
    const value = try std.fmt.parseInt(i64, o, 10); // Parses "12345" as a u64 in base 10
    return value;
}
pub fn main() !void {
    try part2();
}

test "a" {}
