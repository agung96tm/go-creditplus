CREATE TABLE IF NOT EXISTS transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    contact_number VARCHAR(255) NOT NULL,
    amount_admin_fee DECIMAL(10, 2) NOT NULL,
    amount_otr DECIMAL(10, 2) NOT NULL,
    amount_installment DECIMAL(10, 2) NOT NULL,
    amount_interest DECIMAL(10, 2) NOT NULL,
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    asset_name VARCHAR(255) NOT NULL,

    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);