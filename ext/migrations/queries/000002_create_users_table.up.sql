CREATE TABLE IF NOT EXISTS "Users" (
   id serial PRIMARY KEY,
   name VARCHAR(60) NOT NULL,
   city VARCHAR(30) NULL,
   country VARCHAR(30) NULL,
   "avatarImgURL" VARCHAR (500) NULL,
   "lastLoginAt" TIMESTAMP WITH TIME ZONE NULL,
   "blockedAt" TIMESTAMP WITH TIME ZONE NULL,
   "verifiedAt" TIMESTAMP WITH TIME ZONE NULL,
   timezone VARCHAR(20) NOT NULL,
   "timezoneOffset" VARCHAR(3) NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "Users_city_idx" ON "Users" (city);
CREATE INDEX "Users_country_idx" ON "Users" (country);
CREATE INDEX "Users_lastLoginAt_idx" ON "Users" ("lastLoginAt");
CREATE INDEX "Users_blockedAt_idx" ON "Users" ("blockedAt");
CREATE INDEX "Users_verifiedAt_idx" ON "Users" ("verifiedAt");
CREATE INDEX "Users_timezone_idx" ON "Users" (timezone);
CREATE INDEX "Users_timezoneOffset_idx" ON "Users" ("timezoneOffset");

CREATE INDEX "Users_CreatedAt_idx" ON "Users" ("CreatedAt");
CREATE INDEX "Users_UpdatedAt_idx" ON "Users" ("UpdatedAt");
CREATE INDEX "Users_DeletedAt_idx" ON "Users" ("DeletedAt");