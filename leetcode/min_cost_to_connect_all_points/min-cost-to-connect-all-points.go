package min_cost_to_connect_all_points

import "sort"

type UnionFindSet struct { // 并查集
	parent []int // index -> only parent

	// TODO: necessary?
	// rank []int // all are 1 when init, as parent, its rank is size of 连通分量
}

func NewUnionFindSet(size int) *UnionFindSet {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFindSet{parent}
}

func (ufs *UnionFindSet) findAgent(x int) int {
	parent := ufs.parent[x]
	if parent == x {
		return parent
	}
	return ufs.findAgent(parent)
}

// return canMergeAndMerged
func (ufs *UnionFindSet) unit(a int, b int) bool {
	fa, fb := ufs.findAgent(a), ufs.findAgent(b)
	// if findAgent(b) == findAgent(a) return
	if fa == fb {
		return false
	}
	// findAgent(b) -> findAgent(a)
	ufs.parent[fa] = fb
	// rank[findAgent(a)] += rank[findAgent(b)]
	return true
}

type Edge struct {
	a    int
	b    int
	dist int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func dist(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

// 最小生成树
func minCostConnectPoints(points [][]int) int {
	// edge statistic
	edges := []Edge{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{a: i, b: j, dist: dist(points[i], points[j])})
		}
	}

	// sort all edge
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})
	// model the graph
	ufs := NewUnionFindSet(len(points))
	var result int
	edgesLeftCount := len(points) - 1
	for _, e := range edges {
		// if (edge not appeared in tree, in other words not united) {
		// 		unit 2 point
		// }
		if ufs.unit(e.a, e.b) {
			result += e.dist
			edgesLeftCount--
		}
		if edgesLeftCount <= 0 {
			break
		}
	}
	return result
}
