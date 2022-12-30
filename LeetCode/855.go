/*
855. 考场就坐
2022-12-30
link:https://leetcode.cn/problems/exam-room/
question:
	在考场里，一排有 N 个座位，分别编号为 0, 1, 2, ..., N-1 。
	当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。
	如果有多个这样的座位，他会坐在编号最小的座位上。
	(另外，如果考场里没有人，那么学生就坐在 0 号座位上。)
	返回 ExamRoom(int N) 类，它有两个公开的函数：其中，
	函数 ExamRoom.seat() 会返回一个 int （整型数据），代表学生坐的位置；
	函数 ExamRoom.leave(int p) 代表坐在座位 p 上的学生现在离开了考场。每次调用 ExamRoom.leave(p) 时都保证有学生坐在座位 p 上。
answer:
	动态问题：一般都用到有序数据结构，比如平衡二叉树和二叉堆。
	然后自己暂时还写不好。
	参考：使用切片保持有序插入和删除，替代有序哈希表
	https://leetcode.cn/problems/exam-room/solution/golang-li-yong-qie-pian-ti-dai-you-xu-ha-v2w1/
	代码看着简单，有些难理解！
*/
type ExamRoom struct {
	n   int
	arr []int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		n:   n,
		arr: []int{},
	}
}

func (this *ExamRoom) Seat() int { // 这里设计很关键
	student := 0 // 学生被安排的位置索引
	index := 0
	if len(this.arr) > 0 {
		distance := this.arr[0]
		pre := -1
		for i, val := range this.arr {
			if pre != -1 {
				temp := (val - pre) / 2
				if temp > distance {
					distance = temp
					student = pre + temp
					index = i
				}
			}
			pre = val
		}
		if this.n-1-this.arr[len(this.arr)-1] > distance {
			student = this.n - 1
			index = len(this.arr)
		}
	}
	this.arr = append(this.arr, 0)
	copy(this.arr[index+1:], this.arr[index:])
	this.arr[index] = student
	return student
}

func (this *ExamRoom) Leave(p int) {
	// 在切片中寻找
	index := 0
	for i := 0; i < len(this.arr); i++ {
		if this.arr[i] == p {
			index = i
			break
		}
	}
	// 删除
	this.arr = append(this.arr[:index], this.arr[index+1:]...) // 为什么这样删除呢？
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */