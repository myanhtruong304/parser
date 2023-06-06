ALTER TABLE IF EXISTS "wallets" DROP CONSTRAINT IF EXISTS "transactions_wallet_address_fkey";
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS transactions;