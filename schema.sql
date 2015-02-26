CREATE TABLE users (
    id INT,
    name TEXT,
    PRIMARY KEY(id)
);

CREATE TABLE messages (
    id INT,
    content TEXT,
    sender_id INT,
    PRIMARY KEY(id)
    FOREIGN KEY(sender_id) REFERENCES users(id)
);
