CREATE TABLE IF NOT EXISTS users (
	"id" 		    SERIAL PRIMARY KEY,
    "email" 		VARCHAR(250) NOT NULL,
	"password" 		VARCHAR(250) NOT NULL,
	"created_at" 	TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX ix_users_email ON users USING btree (email);

CREATE TABLE IF NOT EXISTS keeper (
    "id" 		    BIGSERIAL PRIMARY KEY,
	"data" 		    TEXT NULL,
	"name" 			VARCHAR(250) NOT NULL,
	"type" 		    VARCHAR(20) NOT NULL,
	"object_id"     TEXT NULL,
	"filename"      TEXT NULL,
	"user_id"		INTEGER NOT NULL,
	"created_at" 	TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	"updated_at" 	TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
ALTER TABLE keeper ADD CONSTRAINT keeper_user_id_fkey FOREIGN KEY ("user_id") REFERENCES users("id");
