CREATE TABLE IF NOT EXISTS products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    partner_id INT,
    description TEXT,


    FOREIGN KEY (partner_id) REFERENCES partners(id) ON DELETE CASCADE
);

INSERT INTO products
    (name, partner_id, price, description)
VALUES
    ('BMW', (SELECT id FROM partners WHERE name = 'mobilku' LIMIT 1), 20000000, 'Mobile Description'),
    ('HONDA BEAT', (SELECT id FROM partners WHERE name = 'motorku' LIMIT 1), 1000000, 'Motorcycle Description'),
    ('Momogi', (SELECT id FROM partners WHERE name = 'produkku' LIMIT 1), 1200000, 'Snack & Drinks'),
    ('Aqua Bottle', (SELECT id FROM partners WHERE name = 'produkku' LIMIT 1), 500000, 'Mineral Water')
