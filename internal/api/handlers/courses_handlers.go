package handlers

import (
	"GoGin/internal/model"
	"GoGin/internal/services"
	"GoGin/internal/util"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	CourseService *services.CourseService
}

func NewCourseHandler(courseService *services.CourseService) *CourseHandler {
	return &CourseHandler{CourseService: courseService}
}

// Info 获取课程列表 Get
func (h *CourseHandler) Info(c *gin.Context) {
	//调用服务层
	courses, err := h.CourseService.GetInfo()
	if err != nil {
		util.Error(c, 500, err.Error())
		return
	}

	// 返回响应
	util.Success(c, gin.H{
		"courses": courses,
	}, "Courses Information")
}

// EnrollmentInfo 获取已选课程列表 Get
func (h *CourseHandler) EnrollmentInfo(c *gin.Context) {
	//捕获数据
	userID, _ := c.Get("user_id")

	//调用服务层
	courses, err := h.CourseService.GetEnrollmentInfo(userID.(int))
	if err != nil {
		util.Error(c, 500, err.Error())
		return
	}

	// 返回响应
	util.Success(c, gin.H{
		"courses": courses,
	}, "Your Enrollment Information")
}

// PickCourse 选课
func (h *CourseHandler) PickCourse(c *gin.Context) {
	// 捕获数据
	var req model.PickRequest
	if err := c.ShouldBind(&req); err != nil {
		util.Error(c, 400, err.Error())
	}
	studentID, _ := c.Get("user_id")

	//调用服务层
	course, err := h.CourseService.PickCourse(studentID.(int), req.CourseID)
	if err != nil {
		util.Error(c, 500, err.Error())
	}

	//返回响应
	util.Success(c, gin.H{
		"course": course,
	}, "Course Picked")
}

// DropCourse 退课
func (h *CourseHandler) DropCourse(c *gin.Context) {
	// 捕获数据
	var req model.DropRequest
	if err := c.ShouldBind(&req); err != nil {
		util.Error(c, 400, err.Error())
	}
	studentID, _ := c.Get("user_id")

	//调用服务层
	course, err := h.CourseService.DropCourse(studentID.(int), req.CourseID)
	if err != nil {
		util.Error(c, 500, err.Error())
	}

	//返回响应
	util.Success(c, gin.H{
		"course": course,
	}, "Course Dropped")
}

// AddCourse 新增课程 (admin)
func (h *CourseHandler) AddCourse(c *gin.Context) {
	//捕获数据
	var req model.AddCourseRequest
	if err := c.ShouldBind(&req); err != nil {
		util.Error(c, 400, err.Error())
	}

	//调用服务层
	course, err := h.CourseService.AddCourse(req.Name, req.Capital)
	if err != nil {
		util.Error(c, 500, err.Error())
	}

	//返回响应
	util.Success(c, gin.H{
		"course": course,
	}, "Course Added")
}
