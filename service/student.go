package service

import (
    "context"
    pb "student-service/genproto/student-service"
)

func (s *StudentRepo) CreateStudent(ctx context.Context, req *pb.StudentRequest) (*pb.Student, error) {
    return s.storage.Student().CreateStudent(req)
}

func (s *StudentRepo) GetStudentById(ctx context.Context, req *pb.GetByIdRequest) (*pb.Student, error) {
    return s.storage.Student().GetStudentById(req)
}

func (s *StudentRepo) GetAllStudents(ctx context.Context, req *pb.GetAllRequest) (*pb.AllStudents, error) {
    return s.storage.Student().GetAllStudents(req)
}

func (s *StudentRepo) UpdateStudent(ctx context.Context, req *pb.StudentRequest) (*pb.Student, error) {
    return s.storage.Student().UpdateStudent(req)
}

func (s *StudentRepo) DeleteStudent(ctx context.Context, req *pb.GetByIdRequest) (*pb.Student, error) {
    return s.storage.Student().DeleteStudent(req)
}
