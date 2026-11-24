# Cycle in Directed Graph (BFS / Kahn)

> Striver SDE Sheet 158 · Day 23 Graph — peel indegree-zero vertices; a cycle is precisely what the peeling can never reach, so count the processed. Go, stdlib only, tested.

## Problem

Directed cycle detection again — this time without DFS, using the indegree-peeling idea that also yields topological sort.

## Approach

A vertex with indegree zero cannot be part of any cycle (nothing feeds it), so it can be safely removed, decrementing its targets' indegrees and possibly freeing them too. Repeat with a queue. If this cascade consumes all n vertices, the graph is a DAG; vertices inside or downstream-locked-behind a cycle never reach indegree zero, so `processed < n` *is* the cycle verdict — no path tracking, no colors. The contrast with #157 is the sheet's point: the DFS version finds a cycle by walking into it; Kahn detects it by exhaustion, as the residue peeling can't dissolve.

**Complexity:** O(V+E) time, O(V) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
