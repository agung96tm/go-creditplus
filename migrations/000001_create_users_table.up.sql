CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nik VARCHAR(100) UNIQUE,
    full_name VARCHAR(255),
    legal_name VARCHAR(255),
    place_birth VARCHAR(255),
    date_birth DATE,
    salary DECIMAL(10, 2),
    id_card_photo VARCHAR(255),
    selfie_photo VARCHAR(255),
    password VARCHAR(255)
);

INSERT INTO users
    (nik, full_name, legal_name, place_birth, date_birth, salary, id_card_photo, selfie_photo, password)
VALUES
    ('1411502550123', 'Budi', 'Budi Suntoyo', 'Tangerang', '1996-08-08', 3000000, 'https://picsum.photos/200/300', 'https://picsum.photos/200/300', 'a075d17f3d453073853f813838c15b8023b8c487038436354fe599c3942e1f95'),
    ('1411502550827', 'Annisa', 'Annisa Yulistia', 'Jakarta', '1990-08-25', 6000000, 'https://picsum.photos/200/300', 'https://picsum.photos/200/300', 'a075d17f3d453073853f813838c15b8023b8c487038436354fe599c3942e1f95');