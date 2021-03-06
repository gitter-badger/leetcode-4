package algorithms

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	if !reflect.DeepEqual(twoSum([]int{1, 3, 2}, 3), []int{0, 2}) {
		t.Fatal("result must be 0, 2")
	}
	if len(twoSum([]int{}, 1)) != 0 {
		t.Fatal("result must be []")
	}
	if len(twoSum([]int{1, 3, 2}, 6)) != 0 {
		t.Fatal("result must be []")
	}
}

func Test1V1(t *testing.T) {
	if !reflect.DeepEqual(twoSumV0([]int{1, 2, 3, 5, 6}, 5), []int{1, 2}) {
		t.Fatal("result must be 1, 2")
	}
	if len(twoSumV0([]int{}, 1)) != 0 {
		t.Fatal("result must be []")
	}
	if len(twoSumV0([]int{1, 3, 2}, 6)) != 0 {
		t.Fatal("result must be []")
	}
}

func Test2(t *testing.T) {
	result := addTwoNumbers(
		&ListNode{5, nil},
		&ListNode{5, nil},
	)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func Test3(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func TestFindKth(t *testing.T) {
	fmt.Println(findKth([]int{1, 2, 3, 4}, []int{2, 3, 4, 7, 9}, 3))
}

func Test4(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{3}))
}

func TestSym(t *testing.T) {
	fmt.Println(sym("abbbc"))
}

func Test5(t *testing.T) {
	fmt.Println(longestPalindrome("aba"))
}

func Test6(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

func Test7(t *testing.T) {
	fmt.Println(reverse(2147483648))
}

func Test8(t *testing.T) {
	fmt.Println(myAtoi("-2147483648"))
}

func Test9(t *testing.T) {
	fmt.Println(isPalindrome(121))
}

func Test10(t *testing.T) {
	if isMatch("abc", "a*c") {
		t.Fatal("should be false")
	}
}

func Test10V1(t *testing.T) {
	if isMatchV1("aa", "a") {
		t.Fatal("should be false")
	}
	if !isMatchV1("aa", "aa") {
		t.Fatal("should be true")
	}
	if isMatchV1("aaa", "aa") {
		t.Fatal("should be false")
	}
	if !isMatchV1("aa", "a*") {
		t.Fatal("should be true")
	}
	if !isMatchV1("aa", ".*") {
		t.Fatal("should be true")
	}
	if !isMatchV1("ab", ".*") {
		t.Fatal("should be true")
	}
	if !isMatchV1("aab", "c*a*b") {
		t.Fatal("should be true")
	}
}

func Test11(t *testing.T) {
	fmt.Println(maxArea([]int{1, 1}))
}

func Test12(t *testing.T) {
	fmt.Println(intToRoman(112))
}

func Test13(t *testing.T) {
	fmt.Println(romanToInt("IV"))
}

func Test14(t *testing.T) {
	a := longestCommonPrefix([]string{"ab", "a", "ac", "abd"})
	if a != "a" {
		t.Fatal("result should be a")
	}
	none := longestCommonPrefix([]string{})
	if none != "" {
		t.Fatal("result should be ")
	}
}

func Test15(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func Test16(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func Test17(t *testing.T) {
	fmt.Println(letterCombinations("0234"))
}

func Test18(t *testing.T) {
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
}

func Test19(t *testing.T) {
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, nil}}, 2))
}

func Test20(t *testing.T) {
	fmt.Println(isValid("(())"))
}

func Test21(t *testing.T) {
	spew.Println(mergeTwoLists(&ListNode{2, nil}, &ListNode{1, nil}))
}

func Test22(t *testing.T) {
	fmt.Println(generateParenthesis(4))
}

func Test23(t *testing.T) {
	spew.Println(mergeKLists([]*ListNode{{2, nil}, {1, nil}}))
}

func Test26(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

func Test27(t *testing.T) {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}

func Test28(t *testing.T) {
	fmt.Println(strStr("a", "a"))
}

func Test29(t *testing.T) {
	fmt.Println(divide(6, 0))
}

func Test30(t *testing.T) {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}

func Test31(t *testing.T) {
	i := []int{1, 3, 2}
	nextPermutation(i)
	fmt.Println(i)
}

func Test32(t *testing.T) {
	fmt.Println(longestValidParentheses("()())("))
}

func Test33(t *testing.T) {
	fmt.Println(search([]int{3, 1}, 1))
}

func Test34(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func Test35(t *testing.T) {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}

func Test38(t *testing.T) {
	fmt.Println(countAndSay(4))
}

func Test44(t *testing.T) {
	fmt.Println(wildCardIsMatch("aab", "c*a*b"))
}

func Test53(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func Test56(t *testing.T) {
	fmt.Println(merge([]Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func Test62(t *testing.T) {
	fmt.Println(uniquePaths(1, 2))
}

func Test63(t *testing.T) {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}

func Test64(t *testing.T) {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

func Test67(t *testing.T) {
	fmt.Println(addBinary("11", "11"))
}

func Test70(t *testing.T) {
	fmt.Println(climbStairs(10))
}

func Test72(t *testing.T) {
	fmt.Println(minDistance("abce", "cde"))
}

func Test81(t *testing.T) {
	fmt.Println(search2([]int{3, 1}, 1))
}

func Test84(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func Test85(t *testing.T) {
	fmt.Println(naiveMaximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
}

func Test91(t *testing.T) {
	decodings := numDecodings("1345")
	if decodings != 2 {
		t.Fatal("decodings should be 2")
	}
}

func Test94(t *testing.T) {
	r := inorderTraversal(&TreeNode{1,
		nil, &TreeNode{Val: 2,
			Left: &TreeNode{Val: 3}, Right: nil}})
	if !reflect.DeepEqual(r, []int{1, 3, 2}) {
		t.Fatal(r)
	}
}

func Test96(t *testing.T) {
	trees := numTrees(3)
	if trees != 5 {
		t.Fatal("trees should be 5")
	}
}

func Test100(t *testing.T) {
	isSameOne := isSameTree(
		&TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}},
		&TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}},
	)

	if !isSameOne {
		t.Fatal("[1, 2, 3] and [1, 2, 3] are same")
	}

	isSameTwo := isSameTree(
		&TreeNode{1, &TreeNode{Val: 2}, nil},
		&TreeNode{1, nil, &TreeNode{Val: 3}},
	)

	if isSameTwo {
		t.Fatal("[1, 2] and [1, null, 3] are not same")
	}

	isSameThree := isSameTree(
		&TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 1}},
		&TreeNode{1, &TreeNode{Val: 1}, &TreeNode{Val: 2}},
	)

	if isSameThree {
		t.Fatal("[1, 1, 2] and [1, 2, 1] are not same")
	}
}

func Test206(t *testing.T) {
	if !reflect.DeepEqual(
		reverseList(&ListNode{1, &ListNode{2, nil}}),
		&ListNode{2, &ListNode{1, nil}},
	) {
		t.Fatal("not equal")
	}
}
