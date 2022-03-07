-- AUTOGENERATED BY storj.io/dbx
-- DO NOT EDIT
CREATE TABLE nodes (
	cid text NOT NULL,
	expired_at timestamp with time zone NOT NULL,
	amount bigint NOT NULL,
	created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
	PRIMARY KEY ( cid )
);
CREATE TABLE pins (
	tx text NOT NULL,
	ix integer NOT NULL,
	cid text NOT NULL,
	amount bigint NOT NULL,
	created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
	PRIMARY KEY ( tx, ix )
);
