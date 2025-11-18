package model

// Student 学生模型
type Student struct {
	ID    int    `json:"student_id" gorm:"primary_key;auto_increment;column:student_id"`
	Name  string `json:"name" gorm:"column:name"`
	Grade string `json:"grade" gorm:"column:grade"`
	Class string `json:"class" gorm:"column:class"`
	//关联
	Enrollments []Enrollment `gorm:"foreignKey:StudentID"`
}

// Course 课程模型
type Course struct {
	ID      int    `json:"course_id" gorm:"primary_key;auto_increment;column:course_id"`
	Name    string `json:"name" gorm:"column:name"`
	Capital int    `json:"capital" gorm:"column:capital"`
	Enroll  int    `json:"enroll" gorm:"column:enroll"`
	//关联
	Enrollments []Enrollment `gorm:"foreignKey:CourseID"`
}

// Enrollment 选课记录模型
type Enrollment struct {
	StudentID int `gorm:"column:student_id"`
	CourseID  int `gorm:"column:course_id"`
	//定义外键关联
	Student Student `gorm:"foreignKey:StudentID;references:ID"`
	Course  Course  `gorm:"foreignKey:CourseID;references:ID"`
}
