const std = @import("std");

const utils = @import("utils");

pub fn main() !void {
    const detail = utils.read_file("/home/mypc/projects/aoc/2025/day1/input.txt");
    std.debug.print("{s}", .{detail});
}
