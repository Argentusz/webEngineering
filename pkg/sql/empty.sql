CREATE TABLE IF NOT EXISTS universities (
    id       SERIAL PRIMARY KEY,
    name     TEXT NOT NULL DEFAULT '',
    login    TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    lang     TEXT NOT NULL DEFAULT 'ru'
);

CREATE TABLE IF NOT EXISTS faculties (
    id           SERIAL PRIMARY KEY,
    universityID INTEGER REFERENCES universities(id),
    name         TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS students (
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS faculties_to_students (
   courseID  INTEGER REFERENCES faculties(id),
   studentID INTEGER REFERENCES students(id)
)