package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 使用數組來模擬一個棧的使用
type Stack struct {
	MaxTop int     // 表示棧最大可以存放個數
	Top    int     // 表示棧頂，因為棧頂固定，因此直接使用Top
	arr    [20]int // 數組模擬棧
}

// 入棧
func (s *Stack) Push(val int) (err error) {
	// 先判斷棧是否滿了
	if s.Top == s.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	s.Top++

	// 存放數據
	s.arr[s.Top] = val
	return
}

// 出棧
func (s *Stack) Pop() (val int, err error) {
	// 判斷棧是否為空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	// 先取值，再s.Top--
	val = s.arr[s.Top]
	s.Top--
	return val, nil
}

// 遍歷棧，注意需要從棧頂開始遍歷
func (s *Stack) List() {
	// 判斷棧是否為空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	fmt.Println("棧的情況如下: ")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d \n", i, s.arr[i])
	}
}

// 判斷一個字符是不是一個運算符 [+ , - , * , /]
func (s *Stack) IsOper(val int) bool {

	// 判斷ASCII碼
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 運算的方法
func (s *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0

	switch oper {
	case 42: // 乘
		res = num2 * num1
	case 43: // 加
		res = num2 + num1
	case 45: // 減
		res = num2 - num1
	case 47: // 除
		res = num2 / num1
	default:
		fmt.Println("運算符錯誤.")
	}
	return res
}

// 返回某個運算符的優先級
func (s *Stack) Priority(oper int) int {

	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {

	// 數棧
	numStack := &Stack{
		MaxTop: 20, // 表示最多存放20個數到棧中
		Top:    -1, // 當前棧頂為-1，表示棧為空
	}

	// 符號棧
	operStack := &Stack{
		MaxTop: 20, // 表示最多存放20個數到棧中
		Top:    -1, // 當前棧頂為-1，表示棧為空
	}

	//exp := "3+2*6-2"
	//exp := "3+3*6-4"
	//exp := "30+3*6-4"
	//exp := "30+30*6-4"
	exp := "300+30*6-4"

	// 定義一個index，幫助掃描exp
	index := 0

	// 為了配合運算，定義需要的變量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		// 這邊增加一個邏輯，處理多位數的問題

		ch := exp[index : index+1] // 字符串

		// ch ==> "+" ==> "43"
		temp := int([]byte(ch)[0]) // 就是字符對應的ASCII碼

		if operStack.IsOper(temp) { // 說明是符號
			if operStack.Top == -1 { // 空棧
				// 如果operStack是一個空棧，直接入棧
				operStack.Push(temp)
			} else {
				// 如果發現operStack棧頂的運算符的優先級大於等於當前準備入棧的運算符的優先級，
				// 就以符號棧pop出，並從數字棧也pop兩個數，進行運算，
				// 運算後的結果，再重新入棧到數字棧，符號再入符號棧
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)

					// 將計算結果重新入數棧
					numStack.Push(result)

					// 當前的符號壓入符號棧
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else { // 說明是數
			// 處理多位數的思路
			// 1. 定義一個變量keepNum string，作拼接
			keepNum += ch

			// 2. 每次要向index的前面字符測試一下，看看是不是運算符，然後處理

			// 如果已經到表達式最後，直接將keepNum
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				// 向index後面測試看看是不是運算符[index]
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}

			// 將字串轉換成整數
			// val, _ := strconv.ParseInt(ch, 10, 64)
			// numStack.Push(int(val))
		}

		// 繼續掃描
		// 先判斷index是否已經掃描到計算表達式的最後
		if index+1 == len(exp) {
			break
		}
		index++
	}

	// 如果掃描表達式完畢，依次從符號棧取出符號，然後從數字棧取出兩個數，
	// 運算後的結果，入數字棧，直到符號棧為空
	for {
		if operStack.Top == -1 {
			break // 退出條件
		}

		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)

		// 將計算結果重新入數棧
		numStack.Push(result)
	}

	// 如果算法沒有問題，表達式也是正確的，則結果就是numStack最後數
	res, _ := numStack.Pop()
	fmt.Printf("表達式 %s = %v", exp, res)
}
