CREATE TABLE IF NOT EXISTS users (
    id                  serial       NOT NULL,
    firebase_id         VARCHAR(28) NOT NULL,
    studying            VARCHAR(256) DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS course_made_by_user (
    user_id           VARCHAR(28) NOT NULL,
    course            VARCHAR(256) NOT NULL
);
