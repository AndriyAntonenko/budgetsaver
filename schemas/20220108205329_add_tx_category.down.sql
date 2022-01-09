ALTER TABLE budget_txs DROP CONSTRAINT budget_tx_category_fk;

ALTER TABLE budget_txs DROP COLUMN "category";

DROP TABLE tx_categories;
