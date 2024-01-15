CREATE TABLE "users" (
  "user_id" SERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "loans" (
  "loan_id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "total_loan" numeric(15,2) NOT NULL,
  "status" varchar NOT NULL,
  "tenor" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "approved_at" timestamptz,
  "completion_at" timestamptz
);

ALTER TABLE "loans" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");