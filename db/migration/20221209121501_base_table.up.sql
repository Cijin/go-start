CREATE OR REPLACE FUNCTION updated()
 RETURNS timestamptz LANGUAGE sql IMMUTABLE STRICT
AS $function$
SELECT now();
$function$;

CREATE TABLE base (
  id uuid NOT NULL uuid_generate_v4 (),
  created_at NOT NULL DEFAULT now (),
  updated_at GENERATED ALWAYS AS (updated()) STORED
);
