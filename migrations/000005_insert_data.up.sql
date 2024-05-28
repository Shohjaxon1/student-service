-- Insert mock data into students
INSERT INTO students (id, first_name, last_name, date_of_birth, gender, phone_number, address, points)
VALUES
    ('123e4567-e89b-12d3-a456-426614174001', 'Alice', 'Smith', '2000-01-15', 'female', '555-1234', '123 Maple St', 90),
    ('123e4567-e89b-12d3-a456-426614174002', 'Bob', 'Johnson', '1999-07-22', 'male', '555-5678', '456 Oak St', 85),
    ('123e4567-e89b-12d3-a456-426614174003', 'Charlie', 'Williams', '2001-03-30', 'male', '555-8765', '789 Pine St', 88),
    ('123e4567-e89b-12d3-a456-426614174004', 'Diana', 'Brown', '2000-06-12', 'female', '555-4321', '101 Elm St', 92),
    ('123e4567-e89b-12d3-a456-426614174005', 'Eve', 'Davis', '2002-02-28', 'female', '555-6789', '202 Birch St', 87),
    ('123e4567-e89b-12d3-a456-426614174006', 'Frank', 'Miller', '2001-08-05', 'male', '555-3456', '303 Cedar St', 89),
    ('123e4567-e89b-12d3-a456-426614174007', 'Grace', 'Wilson', '1998-12-17', 'female', '555-7890', '404 Spruce St', 93),
    ('123e4567-e89b-12d3-a456-426614174008', 'Hank', 'Moore', '2000-04-25', 'male', '555-9012', '505 Ash St', 86),
    ('123e4567-e89b-12d3-a456-426614174009', 'Ivy', 'Taylor', '1999-11-11', 'female', '555-2345', '606 Walnut St', 91),
    ('123e4567-e89b-12d3-a456-42661417400a', 'Jack', 'Anderson', '2001-09-09', 'male', '555-3457', '707 Cherry St', 84);

-- Insert mock data into courses
INSERT INTO courses (id, course_name, department, instructor)
VALUES
    ('123e4567-e89b-12d3-a456-42661417400b', 'Math 101', 'Mathematics', 'Dr. Alan Turing'),
    ('123e4567-e89b-12d3-a456-42661417400c', 'Biology 201', 'Biology', 'Dr. Marie Curie'),
    ('123e4567-e89b-12d3-a456-42661417400d', 'History 301', 'History', 'Dr. George Washington'),
    ('123e4567-e89b-12d3-a456-42661417400e', 'Physics 401', 'Physics', 'Dr. Albert Einstein'),
    ('123e4567-e89b-12d3-a456-42661417400f', 'Chemistry 101', 'Chemistry', 'Dr. Rosalind Franklin'),
    ('123e4567-e89b-12d3-a456-426614174010', 'Literature 201', 'Literature', 'Dr. William Shakespeare');

-- Insert mock data into students_course
INSERT INTO students_course (id, student_id, course_id)
VALUES
    ('123e4567-e89b-12d3-a456-426614174011', '123e4567-e89b-12d3-a456-426614174001', '123e4567-e89b-12d3-a456-42661417400b'),
    ('123e4567-e89b-12d3-a456-426614174012', '123e4567-e89b-12d3-a456-426614174002', '123e4567-e89b-12d3-a456-42661417400c'),
    ('123e4567-e89b-12d3-a456-426614174013', '123e4567-e89b-12d3-a456-426614174003', '123e4567-e89b-12d3-a456-42661417400d'),
    ('123e4567-e89b-12d3-a456-426614174014', '123e4567-e89b-12d3-a456-426614174004', '123e4567-e89b-12d3-a456-42661417400e'),
    ('123e4567-e89b-12d3-a456-426614174015', '123e4567-e89b-12d3-a456-426614174005', '123e4567-e89b-12d3-a456-42661417400f'),
    ('123e4567-e89b-12d3-a456-426614174016', '123e4567-e89b-12d3-a456-426614174006', '123e4567-e89b-12d3-a456-426614174010'),
    ('123e4567-e89b-12d3-a456-426614174017', '123e4567-e89b-12d3-a456-426614174007', '123e4567-e89b-12d3-a456-42661417400b'),
    ('123e4567-e89b-12d3-a456-426614174018', '123e4567-e89b-12d3-a456-426614174008', '123e4567-e89b-12d3-a456-42661417400c'),
    ('123e4567-e89b-12d3-a456-426614174019', '123e4567-e89b-12d3-a456-426614174009', '123e4567-e89b-12d3-a456-42661417400d'),
    ('123e4567-e89b-12d3-a456-42661417401a', '123e4567-e89b-12d3-a456-42661417400a', '123e4567-e89b-12d3-a456-42661417400e');

-- Insert mock data into grades
INSERT INTO grades (id, student_course_id, grade)
VALUES
    ('123e4567-e89b-12d3-a456-42661417401b', '123e4567-e89b-12d3-a456-426614174011', 3.5),
    ('123e4567-e89b-12d3-a456-42661417401c', '123e4567-e89b-12d3-a456-426614174012', 3.8),
    ('123e4567-e89b-12d3-a456-42661417401d', '123e4567-e89b-12d3-a456-426614174013', 3.2),
    ('123e4567-e89b-12d3-a456-42661417401e', '123e4567-e89b-12d3-a456-426614174014', 3.7),
    ('123e4567-e89b-12d3-a456-42661417401f', '123e4567-e89b-12d3-a456-426614174015', 3.6),
    ('123e4567-e89b-12d3-a456-426614174020', '123e4567-e89b-12d3-a456-426614174016', 3.9),
    ('123e4567-e89b-12d3-a456-426614174021', '123e4567-e89b-12d3-a456-426614174017', 3.4),
    ('123e4567-e89b-12d3-a456-426614174022', '123e4567-e89b-12d3-a456-426614174018', 3.1),
    ('123e4567-e89b-12d3-a456-426614174023', '123e4567-e89b-12d3-a456-426614174019', 3.7),
    ('123e4567-e89b-12d3-a456-426614174024', '123e4567-e89b-12d3-a456-42661417401a', 3.5);
