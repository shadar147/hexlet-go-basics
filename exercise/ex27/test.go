package ex27

func Test() {
	list := &ListNode{
		Next: &ListNode{
			Next: nil,
			Val:  2,
		},
		Val: 1,
	}

	list.Reverse().Display()
	list.Display()

	l := &ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: &ListNode{
					Next: nil,
					Val:  40,
				},
				Val: 30,
			},
			Val: 20,
		},
		Val: 10,
	}

	l.Reverse().Display()
	l.Display()
}
