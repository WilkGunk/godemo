package main

func main()  {

}



func sumBTree(root *TreeNode,preNum int)  int {
	if root == nil {
		return 0
	}
	sum := preNum*10+root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return sumBTree(root.Right,sum) + sumBTree(root.Left,sum)
}

