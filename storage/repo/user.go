package repo

import (
	pb "student-service/genproto/student-service"
)

// UserStorageI ...
type StudentStorageI interface {
	CreateStudent(*pb.StudentRequest) (*pb.Student, error)
	GetStudentById(*pb.GetByIdRequest) (*pb.Student, error)
	GetAllStudents(*pb.GetAllRequest) (*pb.AllStudents, error)
	UpdateStudent(*pb.StudentRequest) (*pb.Student, error)
	DeleteStudent(*pb.GetByIdRequest) (*pb.Student, error)
}

type CourseStorageI interface {
	CreateCourse(*pb.Course) (*pb.Course, error)
	GetCourseById(*pb.GetByIdRequest) (*pb.Course, error)
	GetAllCourses(*pb.GetAllRequest) (*pb.AllCourses, error)
	UpdateCourse(*pb.Course) (*pb.Course, error)
	DeleteCourse(*pb.GetByIdRequest) (*pb.Course, error)
}
