-- Create Consumers Table
CREATE TABLE IF NOT EXISTS consumers (
    nik VARCHAR(20) PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    legal_name VARCHAR(100) NOT NULL,
    place_of_birth VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    salary DECIMAL(10, 2) NOT NULL,
    limit_1 DECIMAL(10, 2),
    limit_2 DECIMAL(10, 2),
    limit_3 DECIMAL(10, 2),
    limit_4 DECIMAL(10, 2)
);

-- Create Transactions Table
CREATE TABLE IF NOT EXISTS transactions (
    contract_number VARCHAR(20) PRIMARY KEY,
    nik VARCHAR(20) REFERENCES consumers(nik) ON DELETE CASCADE,
    otr DECIMAL(10, 2) NOT NULL,
    admin_fee DECIMAL(10, 2) NOT NULL,
    installment DECIMAL(10, 2) NOT NULL,
    interest DECIMAL(10, 2) NOT NULL,
    asset_name VARCHAR(100) NOT NULL
);