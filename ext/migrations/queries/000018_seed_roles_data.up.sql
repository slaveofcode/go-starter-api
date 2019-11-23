INSERT INTO "Roles" (name, scopes, "CreatedAt", "UpdatedAt")
VALUES('Owner', '{"team:read","team:write","invoice:read","invoice:write","subscription:read","subscription:write","apikey:read","apikey:write"}', now(), now()),
VALUES('Manager', '{"team:read","team:write","invoice:read","subscription:read","apikey:read","apikey:write"}', now(), now()),
VALUES('Staff', '{"team:read","subscription:read","apikey:read"}', now(), now());