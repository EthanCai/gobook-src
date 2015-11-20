package mychapter1

import (
	"robpike.io/filter"
	"strconv"
)

type Student struct {
	Id          uint
	FirstName   string
	LastName    string
	Age         uint
	Teacher     string
	ClassNumber string
}

// 创建指定大小的Student List
func CreateStudentList(size uint) []Student {
	result := make([]Student, size)

	for i := range result {
		istr := strconv.FormatUint(uint64(i), 10)

		item := &result[i]
		(*item).Id = uint(i)
		(*item).FirstName = "firstName" + istr
		(*item).LastName = "lastName" + istr
		(*item).Age = uint(i)
		(*item).Teacher = "Maria Lee"
		(*item).ClassNumber = "Class" + strconv.FormatUint(uint64((i % 4)), 10)
	}

	return result
}

// 使用filter.Apply提取Student List的Id
func GetStudentIdsUsingApply(studentList []Student) []uint {
	result := filter.Apply(studentList, func(student Student) uint {
		return student.Id
	}).([]uint)

	return result
}

// 使用循环提取Student List的Id
func GetStudentIdsUsingLoop(studentList []Student) []uint {
	size := len(studentList)
	result := make([]uint, size)

	for i, student := range studentList {
		result[i] = student.Id
	}

	return result
}
