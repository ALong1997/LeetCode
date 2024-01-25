package LeetCode

import (
	"math"
)

/*
解法: 确定有限状态机（deterministic finite automaton, DFA）
字符串处理的题目往往涉及复杂的流程以及条件情况，如果直接上手写程序，一不小心就会写出极其臃肿的代码。
因此，为了有条理地分析每个输入字符的处理方法，我们可以使用自动机这个概念：
我们的程序在每个时刻有一个状态 s，每次从序列中输入一个字符 c，并根据字符 c 转移到下一个状态 s'。
这样，我们只需要建立一个覆盖所有情况的从 s 与 c 映射到 s' 的表格即可解决题目中的问题。
另外自动机也需要记录当前已经输入的数字，只要在 s' 为 in_number 时，更新我们输入的数字，即可最终得到输入的数字。
	          ' '	    +/-	      number	   other
start	      start	    signed	  in_number	   end
signed	      end	    end	      in_number	   end
in_number	  end	    end	      in_number	   end
end	          end	    end	      end	       end


结果: 执行用时 :0 ms 内存消耗 :2.3 MB
*/

// 状态
const (
	startState    = iota // 开头空格
	signedState          // 判断符号
	inNumberState        // 数字处理
	endState             // 无效转换，返回0
)

var stateMachine = [][]int{
	startState:    {startState, signedState, inNumberState, endState},
	signedState:   {endState, endState, inNumberState, endState},
	inNumberState: {endState, endState, inNumberState, endState},
	endState:      {endState, endState, endState, endState}}

func referenceMyAtoi(str string) int {
	currentState := startState
	sign, res := 1, 0
	for i := 0; i < len(str); i++ {
		currentState = stateMachine[currentState][getState(str[i])]
		switch currentState {
		case startState:
			continue
		case signedState:
			if str[i] == '-' {
				sign = -1
			} else {
				sign = 1
			}
		case inNumberState:
			res = res*10 + int(str[i]-'0')
			if withSign := res * sign; withSign > math.MaxInt32 {
				return math.MaxInt32
			} else if withSign < math.MinInt32 {
				return math.MinInt32
			}
		case endState:
			break
		default:
			return 0
		}
	}
	return res * sign
}

func getState(c byte) int {
	if c == ' ' {
		return startState
	}
	if c == '+' || c == '-' {
		return signedState
	}
	if c-'0' <= 9 {
		return inNumberState
	}
	return endState
}
