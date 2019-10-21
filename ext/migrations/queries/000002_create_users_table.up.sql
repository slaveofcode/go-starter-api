CREATE TABLE IF NOT EXISTS "Users" (
   id serial PRIMARY KEY,
   name VARCHAR(60) NOT NULL,
   city VARCHAR(30) NULL,
   country VARCHAR(30) NULL,
   "avatarImgURL" VARCHAR (500) NULL,
   "lastLoginAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "blockedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "verifiedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   timezone VARCHAR(20) NOT NULL,
   "timezoneOffset" VARCHAR(3) NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX city_idx ON "Users" (city);
CREATE INDEX country_idx ON "Users" (country);
CREATE INDEX "lastLoginAt_idx" ON "Users" ("lastLoginAt");
CREATE INDEX "blockedAt_idx" ON "Users" ("blockedAt");
CREATE INDEX "verifiedAt_idx" ON "Users" ("verifiedAt");
CREATE INDEX timezone_idx ON "Users" (timezone);
CREATE INDEX "timezoneOffset_idx" ON "Users" ("timezoneOffset");

CREATE INDEX "CreatedAt_idx" ON "Users" ("CreatedAt");
CREATE INDEX "UpdatedAt_idx" ON "Users" ("UpdatedAt");
CREATE INDEX "DeletedAt_idx" ON "Users" ("DeletedAt");