CREATE TABLE teacher
(
    id         serial  not null,
    name       varchar not null,
    surname    varchar not null,
    patronymic varchar not null,
    address    varchar not null
);

CREATE TABLE group
(
    id         int     not null primary key,
    department varchar not null,
    curator_id int     not null,
    CONSTRAINT fk_teachers_id FOREIGN KEY (curator_id)
        REFERENCES teacher (id) on delete cascade
);

CREATE TABLE student
(
    id                 serial  not null primary key,
    name               varchar not null,
    surname            varchar not null,
    patronymic         varchar not null,
    address            varchar not null,
    login              varchar not null unique,
    encrypted_password varchar not null,
    group_id           int     not null,
    CONSTRAINT fk_group_id FOREIGN KEY (group_id)
        REFERENCES group (id) on delete cascade
);

CREATE TABLE subject
(
    id   integer not null primary key,
    name varchar not null
);

CREATE TABLE journal
(
    id         serial  not null primary key,
    grade      integer null,
    subject_id integer not null,
    CONSTRAINT fk_subject_id FOREIGN KEY (subject_id)
        REFERENCES subject (id) on delete cascade,
    student_id integer not null,
    CONSTRAINT fk_student_id FOREIGN KEY (student_id)
        REFERENCES student (id) on delete cascade
);

CREATE TABLE teacher
(
    id         serial  not null primary key,
    name       varchar not null,
    subject_id integer not null,
    CONSTRAINT fk_subject_id FOREIGN KEY (subject_id)
        REFERENCES subject (id) on delete cascade
);