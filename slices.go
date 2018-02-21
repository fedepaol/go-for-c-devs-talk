var a s[]int = []int{1,2,3}
s1 := a[1:3] // 2, 3
s2 := make([]int, 3) // {0, 0, 0}
var s3

s := append(s, 4) /
s := append(s, s1...) // 1,2,3,4,2,3

