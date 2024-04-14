CREATE DATABASE IF NOT EXISTS avito;
USE avito;

CREATE TABLE content (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tag_id VARCHAR(255),
    feature_id VARCHAR(255),
    user_token VARCHAR(255),
    admin_token VARCHAR(255),
    content VARCHAR(255),
    is_active INT
);
