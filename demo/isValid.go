package demo

func isValid(s string) bool {
	hash := map[byte]byte{'{': '}', '[': ']', '(': ')'}
	stack := make([]byte, 0)
	str := []byte(s)
	for _, value := range str {
		if value == '{' || value == '[' || value == '(' { //写入栈
			stack = append(stack, hash[value])
		} else if len(stack) != 0 && value == stack[len(stack)-1] { //防止第一个是反向的，无法删除
			//以上语法：先判断stack长度，如果为0的话就不判断下一个语句，所以不会报错（数组越界）
			stack = stack[:len(stack)-1]
		} else { //不能删除，而且和之前的括号都反向，说明不符合要求
			return false
		}
	}
	//最后删除完是否还有剩余
	return len(stack) == 0
}
