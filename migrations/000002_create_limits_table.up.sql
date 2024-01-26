CREATE TABLE IF NOT EXISTS limits (
    id INT PRIMARY KEY AUTO_INCREMENT,
    consumer_id INT,
    month INT,
    consumer_limit DECIMAL(10, 2),
    FOREIGN KEY (consumer_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO limits
(consumer_id, month, consumer_limit)
VALUES
    ((SELECT id FROM users WHERE nik = '1411502550123' LIMIT 1), 1, 100000),
    ((SELECT id FROM users WHERE nik = '1411502550123' LIMIT 1), 2, 200000),
    ((SELECT id FROM users WHERE nik = '1411502550123' LIMIT 1), 3, 500000),
    ((SELECT id FROM users WHERE nik = '1411502550123' LIMIT 1), 4, 700000),

    ((SELECT id FROM users WHERE nik = '1411502550827' LIMIT 1), 1, 1000000),
    ((SELECT id FROM users WHERE nik = '1411502550827' LIMIT 1), 2, 1200000),
    ((SELECT id FROM users WHERE nik = '1411502550827' LIMIT 1), 3, 1500000),
    ((SELECT id FROM users WHERE nik = '1411502550827' LIMIT 1), 4, 2000000);
