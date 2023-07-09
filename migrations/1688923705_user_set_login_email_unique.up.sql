ALTER TABLE IF EXISTS users ADD CONSTRAINT users_login_unique UNIQUE(login);
ALTER TABLE IF EXISTS users ADD CONSTRAINT users_email_unique UNIQUE(email);