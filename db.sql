-- satrio.accounts definition

-- Drop table

-- DROP TABLE accounts;

CREATE TABLE accounts (
	id uuid NOT NULL,
	nama text NOT NULL,
	nik text NOT NULL,
	no_hp text NOT NULL,
	no_rekening text NOT NULL,
	created_at int8 NOT NULL,
	updated_at int8 NULL,
	saldo numeric NOT NULL,
	CONSTRAINT nik_unique UNIQUE (nik),
	CONSTRAINT no_hp_unique UNIQUE (no_hp),
	CONSTRAINT no_rekening_unique UNIQUE (no_rekening),
	CONSTRAINT users_pk PRIMARY KEY (id)
);

-- satrio.transactions definition

-- Drop table

-- DROP TABLE transactions;

CREATE TABLE transactions (
	id uuid NOT NULL,
	nominal numeric NOT NULL,
	"type" text NOT NULL,
	created_at int8 NOT NULL,
	account_id uuid NOT NULL,
	CONSTRAINT transactions_pk PRIMARY KEY (id)
);


-- satrio.transactions foreign keys

ALTER TABLE satrio.transactions ADD CONSTRAINT transactions_users_fk FOREIGN KEY (account_id) REFERENCES accounts(id);