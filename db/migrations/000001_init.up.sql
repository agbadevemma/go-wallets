CREATE TABLE users(
    id   VARCHAR(36) PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);


CREATE TABLE wallets(
    id   VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    currency VARCHAR(10) NOT NULL DEFAULT 'USD',
    status VARCHAR(10) NOT NULL DEFAULT 'active', -- active, frozen
    version INTEGER NOT NULL DEFAULT 0, -- optimistic lock
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TABLE transactions(
    id   VARCHAR(36) PRIMARY KEY,
    sender_wallet_id VARCHAR(36) NOT NULL,
    receiver_wallet_id VARCHAR(36) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    description TEXT DEFAULT NULL,
    idempotency_key VARCHAR(255) NOT NULL UNIQUE,
    status VARCHAR(10) NOT NULL DEFAULT 'success', -- success,pending, failed
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_wallet_id) REFERENCES wallets(id) ,
    FOREIGN KEY (receiver_wallet_id) REFERENCES wallets(id)
);

CREATE TABLE ledger_entries(
    id   VARCHAR(36) PRIMARY KEY,
    transaction_id VARCHAR(36) NOT NULL,
    wallet_id VARCHAR(36) NOT NULL,
    entry_type VARCHAR(10) NOT NULL DEFAULT 'credit', -- credit, debit
    amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id) ,
    FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);