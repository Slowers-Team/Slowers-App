CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.last_modified = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER set_timestamp
AFTER UPDATE ON Users
FOR EACH ROW
WHEN (((pg_trigger_depth() = 0) ))
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE OR REPLACE TRIGGER set_timestamp
AFTER UPDATE ON Businesses
FOR EACH ROW
WHEN (((pg_trigger_depth() = 0) ))
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE OR REPLACE TRIGGER set_timestamp
AFTER UPDATE ON Memberships
FOR EACH ROW
WHEN (((pg_trigger_depth() = 0) ))
EXECUTE PROCEDURE trigger_set_timestamp();
