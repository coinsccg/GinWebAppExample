package main

import (
	"encoding/json"
	"fmt"
)

// Go语言 JSON序列化技巧其一
// 后端使用SnowFlake算法生成的ID值
// 前端JS拿到这么大的数字（uint64/int64）会丢失精度
// 1<<52          1<<63
// 前端JS中会数字溢出

type Question struct {
	QuestionID uint64 `json:"id,string"`
	Caption    string `json:"caption"`
}

func main() {
	// json marshal
	q1 := Question{
		QuestionID: 21833559413096449,
		Caption:    "你真的会学习吗？",
	}
	b, err := json.Marshal(q1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	// json unmarshal
	var q2 Question
	if err := json.Unmarshal(b, &q2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("q2.QuestionID:%v \n", q2.QuestionID)
}
