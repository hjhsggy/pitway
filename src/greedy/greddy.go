package greedy

import (
	"encoding/json"
	"fmt"
)

func ChargeCoin(m []int, target int) []int {

	result := make([]int, 0, len(m))

	if len(m) == 0 || target == 0 {
		return result
	}

	// m数组是有序的, 有大到小
	for i := 0; i < len(m); i++ {
		result = append(result, target/m[i])
		target = target % m[i]
	}

	return result

}

func Unmarshal() {

	type ListTagResponse struct {
		TagId     int32  `protobuf:"varint,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
		TagName   string `protobuf:"bytes,2,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
		TagStatus int32  `protobuf:"varint,3,opt,name=tag_status,json=tagStatus,proto3" json:"tag_status,omitempty"`
	}
	var reply []ListTagResponse

	data := `[{"tag_id":1,"tag_name":"1","tag_status": 1},
	{"tag_id":2,"tag_name":"2","tag_status":1},
	{"tag_id":3,"tag_name":"3","tag_status":2}]`

	err := json.Unmarshal([]byte(data), &reply)
	if err != nil {
		fmt.Printf("Error unmarsh:%v", err)
		return
	}

	fmt.Println(reply)

}

func moveZeroes(nums []int) {

	var slow, fast int = 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}
