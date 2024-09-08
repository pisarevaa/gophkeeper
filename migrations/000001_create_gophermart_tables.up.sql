CREATE TABLE IF NOT EXISTS user (
	"id" 		    INTEGER PRIMARY KEY,
    "email" 		VARCHAR(250) NOT NULL,
	"password" 		VARCHAR(250) NOT NULL,
	"created_at" 	TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX ix_gophkeeper_user_email ON gophkeeper."user" USING btree (email);

CREATE TABLE IF NOT EXISTS data_text (
    "id" 		    BIGINT PRIMARY KEY,
	"data" 		    TEXT NOT NULL,
	"user_id"		INTEGER NOT NULL,
	"created_at" 	TIMESTAMPTZ NOT NULL DEFAULT NOW()
	"updated_at" 	TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
ALTER TABLE data_text ADD CONSTRAINT data_text_user_id_fkey FOREIGN KEY ("user_id") REFERENCES user("id");

CREATE TABLE IF NOT EXISTS data_file (
    "id" 		    BIGINT PRIMARY KEY,
	"name" 		    VARCHAR(250) NOT NULL,
	"filepath" 		TEXT NOT NULL,
	"user_id"		INTEGER NOT NULL,
	"created_at" 	TIMESTAMPTZ NOT NULL DEFAULT NOW()
	"updated_at" 	TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
ALTER TABLE data_text ADD CONSTRAINT data_text_user_id_fkey FOREIGN KEY ("user_id") REFERENCES user("id");