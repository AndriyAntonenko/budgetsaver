CREATE TABLE tx_categories (
    "id" uuid not null unique default uuid_generate_v4(),
    "name" text not null,
    "creator" uuid null,
    "finance_group" uuid null,
    "created_at" timestamp with time zone default current_timestamp,
    "deleted_at" timestamp with time zone null,

    CONSTRAINT tx_category_creator_fk
        FOREIGN KEY("creator")
            REFERENCES users("id")
                ON DELETE NO ACTION,

    CONSTRAINT tx_category_finance_group_fk
        FOREIGN KEY("finance_group")
            REFERENCES finance_groups("id")
                ON DELETE NO ACTION
);

ALTER TABLE budget_txs ADD COLUMN "category" uuid null;

ALTER TABLE budget_txs ADD CONSTRAINT budget_tx_category_fk FOREIGN KEY("category") REFERENCES tx_categories("id");
