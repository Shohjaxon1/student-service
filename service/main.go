package service

import (
	l "student-service/pkg/logger"
	"student-service/storage"

	"github.com/jmoiron/sqlx"
)

type StudentRepo struct {
	storage storage.IStorage
	logger  l.Logger
}

type CourseRepo struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewStudentService(db *sqlx.DB, log l.Logger) *StudentRepo {
	return &StudentRepo{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func NewCourseService(db *sqlx.DB, log l.Logger) *CourseRepo {
	return &CourseRepo{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}