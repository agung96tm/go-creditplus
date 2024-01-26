CREATE TABLE IF NOT EXISTS partners (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    type_partner VARCHAR(255)
);


INSERT INTO partners
    (name, type_partner)
VALUES
    ('mobilku', 'mechine'),
    ('motorku', 'mechine'),
    ('produkku', 'ecommerce');