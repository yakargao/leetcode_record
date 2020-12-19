package bfs

import (
	"fmt"
	"testing"
)

func TestMinDepth(t *testing.T)  {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left: &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
			},
	}
	fmt.Println(minDepth(root))

	root = &TreeNode{Val: 2,
		Right: &TreeNode{Val: 3,
			Right: &TreeNode{Val: 4,
				Right: &TreeNode{Val: 5,
					Right: &TreeNode{Val: 6,
					},
				},
			},
		},
	}
	fmt.Println(minDepth(root))
}
