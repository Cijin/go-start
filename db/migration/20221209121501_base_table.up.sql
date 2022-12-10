CREATE OR REPLACE FUNCTION updated()
 RETURNS timestamptz LANGUAGE sql IMMUTABLE STRICT
AS $function$
SELECT now();
$function$;

CREATE TABLE IF NOT EXISTS base (
  created_at timestamptz DEFAULT now() NOT NULL,
  updated_at timestamptz GENERATED ALWAYS AS (updated()) STORED
);
