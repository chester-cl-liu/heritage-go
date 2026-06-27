package service

import "genealogy-system/model"

// BuildFamilyTree 将扁平的成员列表高效转化为多叉拓扑树
func BuildFamilyTree(members []model.Member) []*model.Node {
	// 1. 用 Map 建立索引，实现 O(N) 复杂度的快速查找
	nodeMap := make(map[int]*model.Node)
	var roots []*model.Node

	for _, m := range members {
		nodeMap[m.ID] = &model.Node{
			ID:         m.ID,
			Name:       m.Name,
			Spouse:     m.Spouse,
			Generation: m.Generation,
			Children:   []*model.Node{},
		}
	}

	// 2. 建立父子指针纽带
	for _, m := range members {
		currentNode := nodeMap[m.ID]
		if m.FatherID == 0 {
			// 没有父亲，说明是始祖支系
			roots = append(roots, currentNode)
		} else {
			// 找到父亲，并将自己塞入父亲的 children 列表
			if parentNode, exists := nodeMap[m.FatherID]; exists {
				parentNode.Children = append(parentNode.Children, currentNode)
			}
		}
	}
	return roots
}