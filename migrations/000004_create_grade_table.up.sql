CREATE TABLE IF NOT EXISTS grades (
    id UUID,
    student_course_id UUID NOT NULL,
    grade NUMERIC(3, 2),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (student_course_id) REFERENCES students_course(id)
);