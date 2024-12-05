CREATE TABLE IF NOT EXISTS author (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
    );

CREATE TABLE IF NOT EXISTS book (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
    );

CREATE TABLE IF NOT EXISTS author_book (
    author_id INT NOT NULL,
    book_id INT NOT NULL,
    PRIMARY KEY (author_id, book_id),
    FOREIGN KEY (author_id) REFERENCES author(id),
    FOREIGN KEY (book_id) REFERENCES book(id)
    );

CREATE TABLE IF NOT EXISTS stack (
    id INT AUTO_INCREMENT PRIMARY KEY,
    book_id INT NOT NULL,
    stock INT NOT NULL,
    quality VARCHAR(50),
    FOREIGN KEY (book_id) REFERENCES book(id)
    );
