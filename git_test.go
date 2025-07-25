package main

import (
	"testing"
)

func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func Test_sumLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{

		{"simple add", args{
			l1: CreateList([]int{1, 2, 3}), l2: CreateList([]int{1, 2, 3}),
		},
			CreateList([]int{2, 4, 6}),
		},
		{"nulls add", args{
			l1: CreateList([]int{0, 0, 0}), l2: CreateList([]int{0, 0, 0}),
		},
			CreateList([]int{0, 0, 0}),
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := sumLists(tt.args.l1, tt.args.l2); !EqualLists(got, tt.want) {
				t.Errorf("sumLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func EqualLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	return l1 == nil && l2 == nil
}
