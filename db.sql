-- Active: 1658515262971@@127.0.0.1@3306
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    `user_name` TEXT
);
CREATE TABLE posts(  
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    `id_user` INTEGER NOT NULL,
    `text` TEXT,
    create_date DATE,
    foreign key(id_user) references users(id)
);

CREATE TABLE comments(  
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id_posts INTEGER NOT NULL,
    user_name TEXT,
    `text` TEXT,
    create_date DATE,
    foreign key(id_posts) references posts(id)
);

CREATE TABLE reports_posts(  
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id_posts INTEGER,
    `user_name` TEXT,
    `text` TEXT,
    create_date DATE,
    foreign key(id_posts) references posts(id)
);

CREATE TABLE reports_comments(  
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id_comments INTEGER,
    `user_name` TEXT,
    `text` TEXT,
    create_date DATE,
    foreign key(id_comments) references comments(id)
);

INSERT INTO users (id, user_name) VALUES(1, "conacevincent@gmail.com")
INSERT INTO users (id, user_name) VALUES(2, "conace@gmail.com")
INSERT INTO posts (id, id_user, `text`, create_date) VALUES(1, 1, "test posts 1", "2022-07-26T15:25:59.586955503-03:00")
INSERT INTO comments (id, id_posts, `user_name`,`text`, create_date) VALUES(1, 1, "conacevincent@gmail.com", "test posts 1", "2022-07-26T15:25:59.586955503-03:00")