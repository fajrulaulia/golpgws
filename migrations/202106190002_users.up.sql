CREATE TABLE users(
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    fullname VARCHAR(50),
    username VARCHAR(30),
    passkey TEXT,
    email VARCHAR(50),
    phonenumber CHAR(15),
    isverified BOOLEAN DEFAULT false,
    isdisabled BOOLEAN DEFAULT false,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);