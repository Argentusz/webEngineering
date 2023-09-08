CREATE TABLE IF NOT EXISTS universities (
    id       INTEGER PRIMARY KEY,
    name     TEXT NOT NULL DEFAULT '',
    login    TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS courses (
    id           INTEGER PRIMARY KEY,
    universityID INTEGER REFERENCES universities(id),
    name         TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS students (
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS courses_to_students (
   courseID  INTEGER REFERENCES courses(id),
   studentID INTEGER REFERENCES students(id)
)