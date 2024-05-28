package storage

import (
	"student-service/storage/postgres"
	"student-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Student() repo.StudentStorageI
	Course() repo.CourseStorageI
}

type storagePg struct {
	db       *sqlx.DB
	studentRepo repo.StudentStorageI
	courseRepo  repo.CourseStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		studentRepo: postgres.NewStudentRepo(db),
		courseRepo:  postgres.NewCourseRepo(db),
	}
}

func (s storagePg) Student() repo.StudentStorageI {
	return s.studentRepo
}

func (s storagePg) Course() repo.CourseStorageI {
	return s.courseRepo
}
