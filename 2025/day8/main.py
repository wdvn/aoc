import itertools
import math
import re
from collections import deque


def read_file(path) -> str:
    with open(path, 'r') as file:
        content = file.read()
    return content


def extract():
    lines = read_file("input.txt").split("\n")
    out = []
    for line in lines:
        out.append(tuple([int(n) for n in re.findall(r'\d+', line)]))
    return out


def distance(p, q):
    return math.sqrt((p[0] - q[0]) ** 2 + (p[1] - q[1]) ** 2 + (p[2] - q[2]) ** 2)


def part1():
    points = extract()
    # make n(n-1) pairs from n points
    pairs = list(itertools.combinations(points, 2))

    # get a graph with length= len(points)/s with shortest distance
    pairs.sort(key=lambda x: distance(x[0], x[1]))
    pairs = pairs[:len(points)]

    # make graph from pairs
    graph = {}
    for pair in pairs:
        if pair[0] not in graph:
            graph[pair[0]] = []
        if pair[1] not in graph:
            graph[pair[1]] = []
        graph[pair[0]].append(pair[1])
        graph[pair[1]].append(pair[0])
    # using bfs to make all circuits
    visited = {}
    circuits = []
    for point in points:
        if point in visited:
            continue
        queue = deque([point])
        visited[point] = 1
        circuit = []
        while len(queue) > 0:
            current = queue.popleft()
            circuit.append(current)
            for neighbor in graph.get(current, []):
                if neighbor not in visited:
                    visited[neighbor] = 1
                    queue.append(neighbor)
        circuits.append(circuit)
    # sorted
    out = [len(c) for c in circuits]
    out.sort()
    print(math.prod(out[-3:]))


class UnionFind:
    def __init__(self, points: list[tuple]):
        self.parent = {point: point for point in points}
        self.size = len(points)

    def find(self, point):
        if self.parent[point] != point:
            self.parent[point] = self.find(self.parent[point])
        return self.parent[point]

    def union(self, p1, p2):
        leader1 = self.find(p1)
        leader2 = self.find(p2)
        if leader1 == leader2:
            return False
        self.parent[leader1] = leader2
        self.size -= 1
        return True


def part2():
    points = extract()
    # make n(n-1) pairs from n points
    pairs = list(itertools.combinations(points, 2))

    # sort by distance
    pairs.sort(key=lambda x: distance(x[0], x[1]))

    uf = UnionFind(points)

    # Using Kruskal's: Iterate through the sorted edges
    for p1, p2 in pairs:
        if uf.union(p1, p2):
            # finish when size = 1
            if uf.size == 1:
                print(p1[0] * p2[0])
                return


if __name__ == '__main__':
    part2()
