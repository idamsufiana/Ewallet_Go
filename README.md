# Ewallet_Go
Ewallet dalam golang


<img width="640" height="360" alt="image" src="https://github.com/user-attachments/assets/6e6fc979-2342-411a-b2e3-04311fc057e8" />

--


🧱 Database Schema
-- USERS
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- WALLETS
CREATE TABLE wallets (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    balance NUMERIC(18,2) NOT NULL DEFAULT 0
);

🌱 Seed Data
-- Insert user
INSERT INTO users (id, name)
VALUES (
    '11111111-1111-1111-1111-111111111111',
    'Alice'
);

-- Insert wallet
INSERT INTO wallets (id, user_id, balance)
VALUES (
    '22222222-2222-2222-2222-222222222222',
    '11111111-1111-1111-1111-111111111111',
    1000000.00
);

⚙️ Run Application
go run cmd/main.go


📬 API Endpoints
🔍 Get Balance
GET /balance/{user_id}

Example:

GET http://localhost:8080/balance/11111111-1111-1111-1111-111111111111

💸 Withdraw
POST /withdraw

Headers:

Content-Type: application/json

Body:

{
  "user_id": "11111111-1111-1111-1111-111111111111",
  "amount": 50000
}
