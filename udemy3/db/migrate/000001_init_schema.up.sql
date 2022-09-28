CREATE TABLE accounts (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE entries (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE transfers (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON udemy.accounts ("owner");

CREATE INDEX ON udemy.entries ("account_id");

CREATE INDEX ON udemy.transfers ("from_account_id");

CREATE INDEX ON udemy.transfers ("to_account_id");

CREATE INDEX ON udemy.transfers ("from_account_id", "to_account_id");

COMMENT ON COLUMN udemy.entries."amount" IS 'must be positive';

COMMENT ON COLUMN udemy.transfers."amount" IS 'must be positive';

ALTER TABLE udemy.entries ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE udemy.transfers ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE udemy.transfers ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
