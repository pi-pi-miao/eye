package json

import (
	"fmt"
	"testing"
)

type student struct {
	Name  string  `json:"name"`
	Ege   int     `json:"ege"`
}

func NewStudent()*student{
	return &student{
		Name:"pyrene",
		Ege:18,
	}
}

func TestMarshal(t *testing.T){
	if stu,err := Marshal(NewStudent());err != nil {
		t.Errorf("marshal err:%v",err)
	}else{
		fmt.Printf("marshal stu  succ data is %v",string(stu))
	}
	TestUnmarshal(t)
}

func BenchmarkMarshal(b *testing.B) {
	for i := 0; i < b.N ; i ++ {
		Marshal(NewStudent())
	}
	BenchmarkUnmarshal(b)
}

func TestUnmarshal(t *testing.T) {
	if stu,err := Marshal(NewStudent());err != nil {
		t.Errorf("marshal err:%v",err)
	}else{
		stu1 := NewStudent()
		if err := Unmarshal(stu,stu1);err != nil {
			t.Errorf("unmarshal stu faild err is %v",err)
			return
		}
		fmt.Printf("unmarshal stu succ ege  is %v ,name is %v",stu1.Ege,stu1.Name)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N ; i ++ {
		if stu,err := Marshal(NewStudent());err == nil {
			stu1 := NewStudent()
			Unmarshal(stu,stu1)
		}else {
			b.Errorf("unmarshal err %v",err)
		}
	}
}