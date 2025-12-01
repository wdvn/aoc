const std = @import("std");
const allocator = std.heap.page_allocator;

pub fn read_file(path: []const u8) []u8 {
    const buf: []u8 = undefined;
    const file = std.fs.openFileAbsolute(path, .{}) catch |err| {
        std.debug.print("err={} \n", .{err});
        return "";
    };
    defer file.close();
    _ = file.read(buf) catch |err| {
        std.debug.print("err={} \n", .{err});
        return "";
    };
    
    return buf;
}
