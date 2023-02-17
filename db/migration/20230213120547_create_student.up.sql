CREATE TABLE student (
                         id serial not null primary key,
                         login varchar not null unique,
                         encrypted_password varchar not null
);

CREATE TABLE subject (
                         id integer not null primary key,
                         name varchar not null
);

CREATE TABLE journal (
                        id serial not null primary key,
                        grade integer null,
                        subject_id integer not null,
                        CONSTRAINT fk_subject_id FOREIGN KEY (subject_id)
                            REFERENCES subject (id) on delete cascade,
                        student_id integer not null,
                        CONSTRAINT fk_student_id FOREIGN KEY (student_id)
                            REFERENCES student (id) on delete cascade
);

CREATE TABLE teacher (
                        id serial not null primary key,
                        name varchar not null,
                        subject_id integer not null,
                        CONSTRAINT fk_subject_id FOREIGN KEY (subject_id)
                            REFERENCES subject (id) on delete cascade
);