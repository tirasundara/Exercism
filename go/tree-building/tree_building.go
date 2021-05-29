package tree

import (
	"errors"
	"sort"
)

// Record is database record
type Record struct {
	ID     int
	Parent int
}

// Node is a tree node
type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree from records
func Build(records []Record) (*Node, error) {

	node := make(map[int]*Node, len(records))

	// Sort by ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for i, record := range records {
		if i != record.ID || record.Parent > record.ID || record.ID > 0 && record.Parent == record.ID {
			return nil, errors.New("Invalid record")
		}

		node[record.ID] = &Node{ID: record.ID}
		if record.ID != 0 {
			node[record.Parent].Children = append(node[record.Parent].Children, node[record.ID])
		}
	}

	return node[0], nil
}
