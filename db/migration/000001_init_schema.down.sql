ALTER TABLE IF EXISTS "transactions" DROP CONSTRAINT IF EXISTS "transactions_wallet_address_fkey";
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS blocks;