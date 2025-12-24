CREATE TABLE public.events (
	id serial NULL,
	description varchar(256) NULL,
	quota int4 NULL,
	start_date timestamptz NULL,
	end_date timestamptz NULL,
	created_at timestamptz DEFAULT now() NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT events_pk PRIMARY KEY (id)
);

