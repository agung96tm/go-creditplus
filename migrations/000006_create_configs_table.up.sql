CREATE TABLE IF NOT EXISTS configs (
    id INT PRIMARY KEY AUTO_INCREMENT,
    admin_fee DECIMAL(5, 2) NOT NULL
);

INSERT INTO configs (admin_fee) VALUES (1.2);