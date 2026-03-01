-- +goose Up
CREATE SCHEMA IF NOT EXISTS itk;

CREATE TABLE IF NOT EXISTS itk.wallets (
       wallet_id UUID PRIMARY KEY,
       amount    BIGINT NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE IF EXISTS itk.wallets;
DROP SCHEMA IF EXISTS itk;