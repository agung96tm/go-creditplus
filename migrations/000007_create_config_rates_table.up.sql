CREATE TABLE IF NOT EXISTS config_rates (
    id INT PRIMARY KEY AUTO_INCREMENT,
    month INT NOT NULL,
    percentage DECIMAL(5, 2) NOT NULL
);

INSERT INTO config_rates
    (month, percentage)
VALUES
    (1, 2),
    (2, 2),
    (3, 2.75),
    (4, 2.75);