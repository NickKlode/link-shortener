CREATE TABLE links (
    
    original_url TEXT NOT NULL UNIQUE,
    token VARCHAR(10) PRIMARY KEY
);