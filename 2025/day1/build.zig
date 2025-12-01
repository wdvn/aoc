const std = @import("std");
const Build = std.Build;

pub fn build(b: *Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const main_module = b.createModule(.{ .optimize = optimize, .target = target, .root_source_file = b.path("main.zig") });
    const utils_module = b.addModule("utils", .{ .root_source_file = b.path("../utils.zig") });
    const exe = b.addExecutable(.{ .name = "luynt", .root_module = main_module });
    exe.root_module.addImport("utils", utils_module);
    // b.test_step.root_module.addImport("utils", utils_module);
    const run_cmd = b.addRunArtifact(exe);
    run_cmd.step.dependOn(b.getInstallStep());
    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    const run_step = b.step("run", "Run the application");
    run_step.dependOn(&run_cmd.step);
}
