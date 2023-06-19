# go_hexagonal_microservice
BBDD structure:

CREATE SEQUENCE tasks_id_seq START WITH 1;

CREATE TABLE public."tasks" (
	id INTEGER DEFAULT nextval('tasks_id_seq'::regclass) PRIMARY KEY,
	title varchar NULL,
	description varchar NULL,
	priority int NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
)TABLESPACE pg_default;