const std = @import("std");


pub fn read_file(path: []const u8, allocator: std.mem.Allocator) ![]align(@alignOf(@Vector(std.simd.suggestVectorLength(u8).?, u8))) u8 {
    var pathBuff: [std.fs.max_path_bytes]u8 = undefined;
    const fpath = try std.fs.realpath(path, &pathBuff);
    const file = try std.fs.openFileAbsolute(fpath, .{ .mode = .read_only });
    defer file.close();
    const fileSize = try file.getEndPos();

    const aligned_size = std.mem.alignForward(usize, fileSize, std.simd.suggestVectorLength(u8).?);

    const buff = try allocator.alignedAlloc(u8, std.mem.Alignment.of(@Vector(std.simd.suggestVectorLength(u8).?, u8)), aligned_size);

    @memset(buff, 0);

    _ = try file.read(buff);

    return buff;
}