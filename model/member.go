package model

// Member 代表谱系中的一位家族成员
type Member struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name" gorm:"size:64;not null"`
	Spouse     string `json:"spouse" gorm:"size:64"` // 配偶姓名
	FatherID   int    `json:"father_id" gorm:"index"` // 父亲ID，顶层始祖为 0
	Generation int    `json:"generation"`             // 代际代数
}

// Node 用于前端 G6 树图渲染的无损拓扑节点
type Node struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Spouse     string  `json:"spouse,omitempty"`
	Generation int     `json:"generation"`
	Children   []*Node `json:"children,omitempty"` // 子孙切片
}