/*
	Migration file generated by protoc-gen-orm. DO NOT EDIT.
	source: ormpb/examples/postgres/example.proto
*/


/* Generated for Example */
CREATE TABLE examples (
	user_id SERIAL,
	name VARCHAR(128) PRIMARY KEY,
	email VARCHAR(255) NOT NULL UNIQUE,
	point NUMERIC DEFAULT 17.33,
	created_at TIMESTAMP
);