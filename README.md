# Final-Project
# psql code
CREATE DATABASE supplychain_db;

GRANT ALL PRIVILEGES ON DATABASE supplychain_db TO postgres;

\c supplychain_db;

CREATE TYPE user_role AS ENUM ('Farmer', 'Factory', 'Retailer', 'Logistics', 'Admin');

CREATE TABLE Users (
  UserID VARCHAR(255) PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  Role user_role NOT NULL
);

CREATE TABLE Location (
  LocationID SERIAL PRIMARY KEY,
  country VARCHAR(255),
  province VARCHAR(255),
  address TEXT
);

CREATE TABLE Farmer (
  FarmerID VARCHAR(255) PRIMARY KEY,
  UserID VARCHAR(255) REFERENCES Users(UserID) ON DELETE CASCADE,
  farmer_name VARCHAR(255),
  companyname VARCHAR(255),
  LocationID INT REFERENCES Location(LocationID) ON DELETE SET NULL,
  telephone VARCHAR(255),
  LineID VARCHAR(255),
  facebook VARCHAR(255),
  createdOn TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE DairyFactory (
  FactoryID VARCHAR(255) PRIMARY KEY,
  UserID VARCHAR(255) REFERENCES Users(UserID) ON DELETE CASCADE,
  companyname VARCHAR(255),
  LocationID INT REFERENCES Location(LocationID) ON DELETE SET NULL,
  telephone VARCHAR(255),
  createdOn TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE OrganicCertification (
  CertificationID VARCHAR(255) PRIMARY KEY,
  FarmerID VARCHAR(255) REFERENCES Farmer(FarmerID) ON DELETE CASCADE,
  FactoryID VARCHAR(255) REFERENCES DairyFactory(FactoryID) ON DELETE SET NULL,
  CertificationType VARCHAR(255),
  CertificationCID VARCHAR(255) UNIQUE,
  Effective_date DATE,
  Issued_date DATE,
  createdOn TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ExternalID (
  ExternalID VARCHAR(255) PRIMARY KEY,
  FactoryID VARCHAR(255) REFERENCES DairyFactory(FactoryID) ON DELETE CASCADE,
  LogisticName VARCHAR(255),
  SenderName VARCHAR(255),
  LogisticShippingDate DATE,
  LogisticDeliveryDate DATE,
  LogisticQualityCheck BOOLEAN,
  LogisticTemp FLOAT,
  RetailersReceiptDate DATE,
  RetailerQualityCheck BOOLEAN,
  RetailerTemp FLOAT,
  RetailerName VARCHAR(255),
  createdOn TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE LogisticsProvider (
  LogisticsID VARCHAR(255) PRIMARY KEY,
  UserID VARCHAR(255) REFERENCES Users(UserID) ON DELETE CASCADE,
  companyname VARCHAR(255),
  LocationID INT REFERENCES Location(LocationID) ON DELETE SET NULL,
  telephone VARCHAR(255),
  createdOn TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Retailer (
  RetailerID VARCHAR(255) PRIMARY KEY,
  UserID VARCHAR(255) REFERENCES Users(UserID) ON DELETE CASCADE,
  companyname VARCHAR(255),
  LocationID INT REFERENCES Location(LocationID) ON DELETE SET NULL,
  telephone VARCHAR(255),
  createdOn TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE SEQUENCE user_seq START 1;

CREATE FUNCTION generate_userid(role_prefix TEXT) RETURNS TEXT AS $$
DECLARE
  new_id INT;
  new_userid TEXT;
BEGIN
  SELECT NEXTVAL('user_seq') INTO new_id;
  new_userid := role_prefix || TO_CHAR(new_id, 'FM000');
  RETURN new_userid;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION generate_userid() RETURNS TEXT AS $$ 
DECLARE
  prefix TEXT;
  seq_num INTEGER;
  result TEXT;
BEGIN
  prefix := TO_CHAR(CURRENT_DATE, 'YY');
  seq_num := nextval('user_seq_' || prefix);
  result := prefix || TO_CHAR(seq_num, 'FM0000');
  RETURN result;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION create_yearly_sequence(year_prefix TEXT) RETURNS VOID AS $$
BEGIN
  EXECUTE 'CREATE SEQUENCE IF NOT EXISTS user_seq_' || year_prefix || ' START 1;';
END;
$$ LANGUAGE plpgsql;

ALTER TABLE users
ALTER COLUMN userid SET DEFAULT generate_userid();

ALTER TABLE farmer ADD COLUMN IF NOT EXISTS firstname VARCHAR(255);
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS lastname VARCHAR(255);
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS email TEXT UNIQUE;
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS address TEXT;
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS address2 TEXT;
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS areacode VARCHAR(10);
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS city VARCHAR(255);
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS post VARCHAR(10);
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS upload_certification TEXT;
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS location_link TEXT;
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS company_name VARCHAR(255);
ALTER TABLE farmer ADD COLUMN IF NOT EXISTS phone VARCHAR(20);
# Backend
