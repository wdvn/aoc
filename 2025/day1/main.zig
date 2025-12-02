const std = @import("std");

const utils = @import("utils");
const allocator = std.heap.page_allocator;
pub fn main() !void {
    const detail = try utils.read_file("./input.txt", allocator);
    
    std.debug.print("{s}\n", .{detail});
    std.debug.print("{}", .{"1"[0]});
}
