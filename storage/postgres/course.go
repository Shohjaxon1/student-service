package postgres

import (
	"database/sql"
	pb "student-service/genproto/student-service"
	"time"

	"github.com/jmoiron/sqlx"
)

type StudentRepo struct {
	db *sqlx.DB
}

type CourseRepo struct {
	db *sqlx.DB
}

// NewStudentRepo creates a new StudentRepo
func NewStudentRepo(db *sqlx.DB) *StudentRepo {
	return &StudentRepo{db: db}
}

// NewCourseRepo creates a new CourseRepo
func NewCourseRepo(db *sqlx.DB) *CourseRepo {
    return &CourseRepo{db: db}
}

// CreateCourse inserts a new course into the database
func (r *CourseRepo) CreateCourse(course *pb.Course) (*pb.Course, error) {
	var updatedAt sql.NullTime
    query := `
        INSERT INTO courses (id, course_name, department, instructor)
        VALUES ($1, $2, $3, $4)
        RETURNING id, course_name, department, instructor, created_at, updated_at
    `
    err := r.db.QueryRow(
        query,
        course.Id,
        course.CourseName,
        course.Department,
        course.Instructor,
    ).Scan(
        &course.Id,
        &course.CourseName,
        &course.Department,
        &course.Instructor,
        &course.CreatedAt,
        &updatedAt,
    )
    if err != nil {
        return nil, err
    }
	course.UpdatedAt = updatedAt.Time.String()
    return course, nil
}

// GetCourseById retrieves a course by its ID
func (r *CourseRepo) GetCourseById(req *pb.GetByIdRequest) (*pb.Course, error) {
    query := "SELECT id, course_name, department, instructor, created_at, updated_at FROM courses WHERE id = $1"
    row := r.db.QueryRow(query, req.Id)
	var updatedAt sql.NullTime
    course := &pb.Course{}
    err := row.Scan(
        &course.Id,
        &course.CourseName,
        &course.Department,
        &course.Instructor,
        &course.CreatedAt,
        &updatedAt,
    )
    if err != nil {
        return nil, err
    }
	course.UpdatedAt = updatedAt.Time.String()
    return course, nil
}

// GetAllCourses retrieves all courses with pagination
func (r *CourseRepo) GetAllCourses(req *pb.GetAllRequest) (*pb.AllCourses, error) {
    intLimit := req.Limit
    intPage := req.Page
    offset := (intPage - 1) * intLimit

    query := `
        SELECT id, course_name, department, instructor, created_at, updated_at
        FROM courses
        LIMIT $1 OFFSET $2
    `
    rows, err := r.db.Query(query, intLimit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    courses := []*pb.Course{}

    for rows.Next() {
        var updatedAt sql.NullTime
        course := &pb.Course{}
        err := rows.Scan(
            &course.Id,
            &course.CourseName,
            &course.Department,
            &course.Instructor,
            &course.CreatedAt,
            &updatedAt,
        )
        if err != nil {
            return nil, err
        }
        if updatedAt.Valid {
            course.UpdatedAt = updatedAt.Time.String()
        } else {
            course.UpdatedAt = ""
        }
        courses = append(courses, course)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &pb.AllCourses{Courses: courses}, nil
}

// UpdateCourse updates a course's details in the database
func (r *CourseRepo) UpdateCourse(req *pb.Course) (*pb.Course, error) {
    query := `
        UPDATE courses
        SET course_name = $2, department = $3, instructor = $4, updated_at = $5
        WHERE id = $1
        RETURNING id, course_name, department, instructor, created_at, updated_at
    `
    row := r.db.QueryRow(query, req.Id, req.CourseName, req.Department, req.Instructor, time.Now())
    updatedCourse := &pb.Course{}
    err := row.Scan(
        &updatedCourse.Id,
        &updatedCourse.CourseName,
        &updatedCourse.Department,
        &updatedCourse.Instructor,
        &updatedCourse.CreatedAt,
        &updatedCourse.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }
    return updatedCourse, nil
}

// DeleteCourse removes a course from the database
func (r *CourseRepo) DeleteCourse(req *pb.GetByIdRequest) (*pb.Course, error) {
    query := "DELETE FROM courses WHERE id = $1 RETURNING id, course_name, department, instructor, created_at, updated_at"
    row := r.db.QueryRow(query, req.Id)
    var updatedAt sql.NullTime
    deletedCourse := &pb.Course{}
    err := row.Scan(
        &deletedCourse.Id,
        &deletedCourse.CourseName,
        &deletedCourse.Department,
        &deletedCourse.Instructor,
        &deletedCourse.CreatedAt,
        &updatedAt,
    )
    if err != nil {
        return nil, err
    }
    if updatedAt.Valid {
        deletedCourse.UpdatedAt = updatedAt.Time.String()
    } else {
        deletedCourse.UpdatedAt = ""
    }
    return deletedCourse, nil
}
