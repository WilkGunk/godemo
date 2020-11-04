package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Print(validMountainArray([]int{0,3,2,1}))
}

//``````````````有效山脉数组``````````````````
func validMountainArray(A []int) bool {
	if len(A) <=2 {
		return false
	}
	max := 0
	index := 0
	for i, a := range A {
		if a > max {
			max = a
			index = i
		}
	}
	if index == len(A) -1 || index == 0 {
		return false
	}
	for i, _ := range A {
		if i < index {
			if A[i+1] <= A[i]{
				return false
			}
		}
		if i >= index && i < len(A) - 1 {
			if A[i+1] >= A[i]{
				return false
			}
		}
	}
	return true
}
//```````````````````````````````````````````

//`````````数组交集`````````````````
func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	m2 := map[int]int{}
	nums3:= make([]int,0)

	for i, v := range nums1 {
		m[v] = i
	}
	for _, num := range nums2 {
		if _,ok := m[num];ok{
			m2[num] = 1

		}
	}

	for k, _ := range m2 {
		nums3 = append(nums3,k)
	}
	return nums3
}
//``````````````````````````````````

//`````````````````停车系统````````````````````
type ParkingSystem struct {
	BigCar int
	mediumCar int
	SmallCar int


}


func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big,medium,small}
}


func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 {
		if 0 < this.BigCar {
			this.BigCar -= 1
			return true
		}else {
			return false
		}
	}
	if carType == 2 {
		if 0 < this.mediumCar {
			this.mediumCar -= 1
			return true
		}else {
			return false
		}
	}
	if carType == 3 {
		if 0 < this.SmallCar {
			this.SmallCar -= 1
			return true
		}else {
			return false
		}
	}
	return false
}
//````````````````````````````````````````````


//``````````````宝石与石头`````````````
func numJewelsInStones(J string, S string) int {
	count := 0
	for i := 0; i< len(J);i ++  {
		for x := 0; x< len(S);x ++ {
			if J[i] == S[x] {
				count ++
			}
		}

	}
	return count
}
//`````````````````````````````````````



//··················重排数组··················
func shuffle(nums []int, n int) []int {
	nNums := make([]int,0)
	snums := nums[:n]
	enums := nums[n:]
	for i:=0;i<n;i++{
		nNums = append(nNums,snums[i])
		nNums = append(nNums,enums[i])
	}
	return nNums
}
//``````````````````````````````````````````````

//````````糖果最多```````````
func kidsWithCandies(candies []int, extraCandies int) []bool {
	bs := make([]bool,0)
	maxCandies := 0
	for _, candies := range candies {
		if maxCandies < candies{
			maxCandies = candies
		}
	}

	for _, v := range candies {
		if v+extraCandies >= maxCandies {
			bs = append(bs,true)
		}else {
			bs = append(bs,false)
		}
	}


	return  bs
}
//``````````````````````````


//````````````好数对``````````
func numIdenticalPairs(nums []int) int {
	var count int
	for i, v := range nums {
		for i2, v2 := range nums {
			if v2 == v && i > i2 {
				count ++
			}
		}
	}
	return count
}

func numIdenticalPairs2(nums []int) int {
	m :=map[int]int{}
	count := 0
	for i, _ := range nums {
		if v,ok := m[nums[i]];ok{
			count+=v
			m[nums[i]] = v+1
		}else{
			m[nums[i]]= 1
		}
	}
	return count
}

//````````````````````````````


//`````````````左旋转字符串``````````````````
func reverseLeftWords(s string, n int) string {
	strs := strings.Split(s,"")
	var endStr,startStr string
	for i, v := range strs {
		if i < n {
			endStr += v
			continue
		}
		startStr += v
	}
	return startStr+endStr
}

func reverseLeftWords1(s string,n int) string {
	return s[n:] + s[:n]
}
//```````````````````````````````````````````

//···········动态和············
func runningSum(nums []int) []int{
	newNums := make([]int,0)
	for i, _ := range nums {
		var newNum int
		for i3, i4 := range nums {
			if i3 <= i {
				newNum += i4
			}
		}
		newNums = append(newNums,newNum)
	}
	return newNums
}

func runingSum1(nums []int) []int  {
	for i := range nums {
		if i > 0 {
			nums[i] = nums[i-1] +nums[i]
		}
	}
	return nums
}

//································



//````````删除元音```````````
func delAEIOU(string2 string) string {
	strs :=strings.Split(string2,"")
	var newStr string
	for _, str := range strs {
		if strings.ToLower(str) == "a" || strings.ToLower(str) == "e" || strings.ToLower(str) == "i" || strings.ToLower(str) == "o" || strings.ToLower(str) == "u" {
			continue
		}
		newStr += str
	}
	return newStr
}
//``````````````````````````

//````````````岛屿周长```````````````
func islandPerimeter(grid [][]int) int {
	perm := 0
	for x, val := range grid {
		for y, _ := range val {
			if grid[x][y] == 0 {
				continue
			}
			perm += 4
			if x > 0 && grid[x-1][y] == 1 {
				perm -= 2
			}

			if y > 0 && grid[x][y-1] == 1 {
				perm -= 2
			}
		}
	}
	return perm
}
//````````````````


//···········两数求和···············
func twoSum(nums []int, target int) []int {
	for index := 0; index <= len(nums)-1; index ++ {
		for index2 := index+1;index2 <= len(nums) -1; index2 ++ {
			if target == nums[index]+nums[index2] {
				return []int{index, index2}
			}
		}
	}
	return []int{0,0}
}

func twoSumForMap(nums []int,target int) []int{
	m :=map[int]int{}
	for i, v := range nums {
		if k,ok := m[target - v];ok {
			return []int{k,i}
		}
		m[v] = i
	}
	return nil
}
//······························
type ListNode struct {
	     Val int
	     Next *ListNode
	}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	//这里用一个result，只是为了后面返回节点方便，并无他用
	result := list
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{tmp % 10, nil}
		tmp = tmp / 10
		list = list.Next
	}
	return result.Next
}

//``````````````独一无二出现次数``````````````
func uniqueOccurrences(arr []int) bool {
	m := arr_To_Map(arr)
	for _,v := range arr {
		m[v] = m[v]+1
	}
	newm := map[int]int{}
	for k, v := range m {
		if _,ok := newm[v]; ok {
			return false
		}
		newm[v] = k
	}

	return true
}

func arr_To_Map(arr []int) (map[int]int) {
	mapObj := map[int]int{}
	for _,v := range arr {
		mapObj[v] = 0
	}
	return mapObj
}
//`````````````````````


//`````````回文数``````````````````
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0&& x!=0){
		return false
	}
	num := 0
	for x > num {
		num = num*10 + x%10
		x = x / 10
	}
   return  x == num || x == num/10
}
//```````````````````````

//·········根到叶子节点的数字之和(二叉树)···········
type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }

func dfs(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}
	sum := prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}
//··············

