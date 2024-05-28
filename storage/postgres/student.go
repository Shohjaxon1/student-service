package postgres

import (
    "database/sql"
    pb "student-service/genproto/student-service"
    "time"
    "github.com/spf13/cast"
)

// CreateStudent inserts a new student into the database
func (r *StudentRepo) CreateStudent(student *pb.StudentRequest) (*pb.Student, error) {
    query := `
        INSERT INTO students (id, first_name, last_name, date_of_birth, gender, phone_number, address, points, created_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
        RETURNING id, first_name, last_name, date_of_birth, gender, phone_number, address, points, created_at
    `
    newStudent := &pb.Student{}
    err := r.db.QueryRow(query, student.Id, student.FirstName, student.LastName, student.DateOfBirth, student.Gender, student.PhoneNumber, student.Address, student.Points, time.Now()).Scan(
        &newStudent.Id,
        &newStudent.FirstName,
        &newStudent.LastName,
        &newStudent.DateOfBirth,
        &newStudent.Gender,
        &newStudent.PhoneNumber,
        &newStudent.Address,
        &newStudent.Points,
        &newStudent.CreatedAt,
    )
    if err!= nil {
        return nil, err
    }
    return newStudent, nil
}

// GetById retrieves a student by their ID
func (r *StudentRepo) GetStudentById(req *pb.GetByIdRequest) (*pb.Student, error) {
    query := "SELECT id, first_name, last_name, date_of_birth, gender, phone_number, address, points, created_at, updated_at FROM students WHERE id = $1"
    row := r.db.QueryRow(query, req.Id)

    student := &pb.Student{}
    err := row.Scan(
        &student.Id,
        &student.FirstName,
        &student.LastName,
        &student.DateOfBirth,
        &student.Gender,
        &student.PhoneNumber,
        &student.Address,
        &student.Points,
        &student.CreatedAt,
        &student.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }
    return student, nil
}

// UpdateStudent updates a student's details in the database
func (r *StudentRepo) UpdateStudent(req *pb.StudentRequest) (*pb.Student, error) {
    query := `
        UPDATE students
        SET first_name = $2, last_name = $3, date_of_birth = $4, gender = $5, phone_number = $6, address = $7, points = $8, updated_at = $9
        WHERE id = $1
        RETURNING id, first_name, last_name, date_of_birth, gender, phone_number, address, points, created_at, updated_at
    `
    row := r.db.QueryRow(query, req.Id, req.FirstName, req.LastName, req.DateOfBirth, req.Gender, req.PhoneNumber, req.Address, req.Points, time.Now())
    updatedStudent := &pb.Student{}
    err := row.Scan(
        &updatedStudent.Id,
        &updatedStudent.FirstName,
        &updatedStudent.LastName,
        &updatedStudent.DateOfBirth,
        &updatedStudent.Gender,
        &updatedStudent.PhoneNumber,
        &updatedStudent.Address,
        &updatedStudent.Points,
        &updatedStudent.CreatedAt,
        &updatedStudent.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }
    return updatedStudent, nil
}

// DeleteStudent removes a student from the database
func (r *StudentRepo) DeleteStudent(req *pb.GetByIdRequest) (*pb.Student, error) {
    query := "DELETE FROM students WHERE id = $1 RETURNING id, first_name, last_name, date_of_birth, gender, phone_number, address, points, created_at, updated_at"
    row := r.db.QueryRow(query, req.Id)
    var updatedAt sql.NullTime
    deletedStudent := &pb.Student{}
    err := row.Scan(
        &deletedStudent.Id,
        &deletedStudent.FirstName,
        &deletedStudent.LastName,
        &deletedStudent.DateOfBirth,
        &deletedStudent.Gender,
        &deletedStudent.PhoneNumber,
        &deletedStudent.Address,
        &deletedStudent.Points,
        &deletedStudent.CreatedAt,
        &updatedAt,
    )
    if err != nil {
        return nil, err
    }
    if updatedAt.Valid {
        deletedStudent.UpdatedAt = updatedAt.Time.String()
    } else {
        deletedStudent.UpdatedAt = ""
    }
    return deletedStudent, nil
}

// GetAllStudents retrieves all students from the database
func (r *StudentRepo) GetAllStudents(req *pb.GetAllRequest) (*pb.AllStudents, error) {
    intLimit := cast.ToInt(req.Limit)
    intPage := cast.ToInt(req.Page)
    offset := (intPage - 1) * intLimit

    query := `
        SELECT id, first_name, last_name, date_of_birth, gender, phone_number, address, points, created_at, updated_at
        FROM students
        LIMIT $1 OFFSET $2
    `
    rows, err := r.db.Query(query, intLimit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    students := []*pb.Student{}

    for rows.Next() {
        var updatedAt sql.NullTime
        student := &pb.Student{}
        err := rows.Scan(
            &student.Id,
            &student.FirstName,
            &student.LastName,
            &student.DateOfBirth,
            &student.Gender,
            &student.PhoneNumber,
            &student.Address,
            &student.Points,
            &student.CreatedAt,
            &updatedAt,
        )
        if err != nil {
            return nil, err
        }
        if updatedAt.Valid {
            student.UpdatedAt = updatedAt.Time.String()
        } else {
            student.UpdatedAt = ""
        }
        students = append(students, student)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &pb.AllStudents{Students: students}, nil
}

