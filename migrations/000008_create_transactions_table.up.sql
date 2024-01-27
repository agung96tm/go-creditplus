CREATE TABLE IF NOT EXISTS transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    contact_number VARCHAR(255) NOT NULL,
    amount_otr DECIMAL(10, 2) NOT NULL,
    amount_installment DECIMAL(10, 2) NOT NULL,
    amount_interest DECIMAL(10, 2) NOT NULL
);