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

CREATE TABLE public.bookings (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	id_event int8 NULL,
	total int8 NULL,
	CONSTRAINT bookings_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_bookings_deleted_at ON public.bookings USING btree (deleted_at);