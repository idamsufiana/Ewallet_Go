# Ewallet_Go
Ewallet dalam golang


<img width="640" height="360" alt="image" src="https://github.com/user-attachments/assets/6e6fc979-2342-411a-b2e3-04311fc057e8" />

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE wallets (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    balance NUMERIC(18,2) NOT NULL DEFAULT 0
);

-- Insert a user
INSERT INTO users (id, name) 
VALUES ('11111111-1111-1111-1111-111111111111', 'Alice');

-- Insert a wallet for that user
INSERT INTO wallets (id, user_id, balance) 
VALUES ('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', 1000000.00);

go run cmd/main.go

Test 

postman request 'http://localhost:8080/balance/11111111-1111-1111-1111-111111111111'

postman request POST 'http://localhost:8080/withdraw' \
  --header 'Content-Type: application/json' \
  --body '{
  "user_id": "11111111-1111-1111-1111-111111111111",
  "amount": 50000
}'





