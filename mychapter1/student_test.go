package mychapter1_test

import (
	"gobook-src/mychapter1"
	"testing"
)

func Test_CreateStudentList(t *testing.T) {

	var studentList []mychapter1.Student
	studentList = mychapter1.CreateStudentList(5)

	if studentList == nil || len(studentList) == 0 {
		t.Error("TestCreateSmallStudentList Failed")
	}
}

func Test_GetStudentIdsUsingApply(t *testing.T) {
	var studentList []mychapter1.Student
	studentList = mychapter1.CreateStudentList(5)
	var studentIds []uint
	studentIds = mychapter1.GetStudentIdsUsingApply(studentList)

	if studentIds == nil || len(studentIds) == 0 {
		t.Error("Test_GetStudentIdsUsingApply Failed")
	}
}

func Test_GetStudentIdsUsingLoop(t *testing.T) {
	var studentList []mychapter1.Student
	studentList = mychapter1.CreateStudentList(5)
	var studentIds []uint
	studentIds = mychapter1.GetStudentIdsUsingLoop(studentList)

	if studentIds == nil || len(studentIds) == 0 {
		t.Error("Test_GetStudentIdsUsingLoop Failed")
	}
}

func Benchmark_CreateStudentList(b *testing.B) {
	mychapter1.CreateStudentList(uint(b.N))
}

func Benchmark_GetStudentIdsUsingApply(b *testing.B) {
	var studentList []mychapter1.Student
	studentList = mychapter1.CreateStudentList(uint(b.N))

	b.ResetTimer()

	mychapter1.GetStudentIdsUsingApply(studentList)
}

func Benchmark_GetStudentIdsUsingLoop(b *testing.B) {
	var studentList []mychapter1.Student
	studentList = mychapter1.CreateStudentList(uint(b.N))

	b.ResetTimer()

	mychapter1.GetStudentIdsUsingLoop(studentList)
}

