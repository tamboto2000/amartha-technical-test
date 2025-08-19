-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  "id" serial PRIMARY KEY NOT NULL,
  "name" varchar(20) NOT NULL,
  "email" varchar(100) NOT NULL,
  "password" bytea NOT NULL,
  "password_salt" bytea NOT NULL,
  "credibility_level" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "user_loans" (
  "id" serial PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "loan_product_id" int NOT NULL,
  "is_paid_off" bool NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "bills" (
  "id" serial PRIMARY KEY NOT NULL,
  "number" varchar NOT NULL,
  "user_id" int NOT NULL,
  "user_loan_id" int NOT NULL,
  "amount" numeric(9,2) NOT NULL,
  "due_date" date NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "repayments" (
  "id" serial PRIMARY KEY NOT NULL,
  "bill_id" int NOT NULL,
  "amount" numeric(9,2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "loan_products" (
  "id" serial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "description" text NOT NULL,
  "amount" numeric(9,2) NOT NULL,
  "interest" numeric(4,2) NOT NULL, -- percentage
  "installment" numeric(9,2) NOT NULL,
  "installment_period" int NOT NULL, -- in days
  "tenor" int NOT NULL, -- in days
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE "user_loans" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_loans" ADD FOREIGN KEY ("loan_product_id") REFERENCES "loan_products" ("id");

ALTER TABLE "bills" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "bills" ADD FOREIGN KEY ("user_loan_id") REFERENCES "user_loans" ("id");

ALTER TABLE "repayments" ADD FOREIGN KEY ("bill_id") REFERENCES "bills" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
