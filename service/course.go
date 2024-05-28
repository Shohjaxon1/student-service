package service

import (
    "context"
    pb "student-service/genproto/student-service"
)

func (s *CourseRepo) CreateCourse(ctx context.Context, req *pb.Course) (*pb.Course, error) {
    return s.storage.Course().CreateCourse(req)
}

func (s *CourseRepo) GetCourseById(ctx context.Context, req *pb.GetByIdRequest) (*pb.Course, error) {
    return s.storage.Course().GetCourseById(req)
}

func (s *CourseRepo) GetAllCourses(ctx context.Context, req *pb.GetAllRequest) (*pb.AllCourses, error) {
    return s.storage.Course().GetAllCourses(req)
}

func (s *CourseRepo) UpdateCourse(ctx context.Context, req *pb.Course) (*pb.Course, error) {
    return s.storage.Course().UpdateCourse(req)
}

func (s *CourseRepo) DeleteCourse(ctx context.Context, req *pb.GetByIdRequest) (*pb.Course, error) {
    return s.storage.Course().DeleteCourse(req)
}
