package controller

import (
	"shoppie/util"

	"github.com/gin-gonic/gin"
)

// Staff staff
type Staff struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Task         []Task `json:"tasks"`
	LuckyNumbers []int  `json:"lucky_numbers"`
}

// Task task
type Task struct {
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
}

// Index create staff
func Index(c *gin.Context) {
	// var staff model.ShopStaff
	staff := Staff{
		Name: "kevin",
		Age:  25,
	}
	staff.Task = append(staff.Task, Task{
		Name:     "clean room",
		Deadline: "2020-06-11",
	})
	staff.Task = append(staff.Task, Task{
		Name:     "cooking",
		Deadline: "2020-06-12",
	})
	courseTask := Task{
		Name:     "Youtube course about goalng, about 6 hours",
		Deadline: "2020-06-13",
	}
	frontendCourseTask := Task{
		Name:     "Youtube course about javascript and css",
		Deadline: "2020-06-13",
	}
	staff.addTask(courseTask)
	staff.addTask(frontendCourseTask)
	staff.LuckyNumbers = []int{1, 3, 5, 7}
	evenNumbers := []int{2, 4, 6, 8}
	staff.LuckyNumbers = append(evenNumbers, staff.LuckyNumbers...) // variadic append
	util.Success(c, staff)

}

func (staff *Staff) addTask(task Task) {
	staff.Task = append(staff.Task, task)
}
