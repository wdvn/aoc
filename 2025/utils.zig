const std = @import("std");

pub fn read_file(alloc: std.mem.Allocator, path: []const u8) ![]u8 {
    const file = try std.fs.cwd().openFile(path, .{ .mode = .read_only });
    defer file.close();
    const fileSize = try file.getEndPos();
    const content = file.readToEndAlloc(alloc, fileSize);
    return content;
}

pub const strings = struct {
    alloc: std.mem.Allocator,
    pub fn split(self: *const strings, s: []const u8, deter: []const u8) ![][]const u8 {
        var out = try std.ArrayList([]const u8).initCapacity(self.alloc, 0);
        var iter = std.mem.splitSequence(u8, s, deter);
        while (true) {
            const seq = iter.next();
            if (seq != null) {
                try out.append(self.alloc, @constCast(seq.?));
            } else {
                break;
            }
        }
        return out.items;
    }
};

pub const slice = struct {
    pub fn init(comptime T: type, alloc: std.mem.Allocator, cap: usize) std.ArrayList {
        var list = try std.ArrayList(T);
        return list.initCapacity(alloc, cap);
    }
    pub fn len(list: std.ArrayList) usize{
        return list.items.len;
    }
    pub fn last(comptime T: type, list: std.ArrayList) type {
        const last_index = list.items.len - 1;
        return @as(T, list.items[last_index]);
    }
    pub fn assign(list: std.ArrayList, i: usize, v: type) void {
        list.items[i] = v;
    }
};
