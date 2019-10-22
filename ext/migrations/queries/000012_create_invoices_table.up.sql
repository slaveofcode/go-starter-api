CREATE TABLE IF NOT EXISTS "Invoices" (
   id serial PRIMARY KEY,
   "userId" INTEGER NOT NULL REFERENCES "Users"(id),
   "planType" VARCHAR(15) NOT NULL,
   "discAmount" DECIMAL(12,2) NULL,
   "discPercentage" DECIMAL(12,2) NULL,
   "paymentType" VARCHAR(15) NOT NULL,
   amount DECIMAL(12,2) NOT NULL,
   tax DECIMAL(12,2) NULL,
   "totalAmount" DECIMAL(12,2) NOT NULL,
   "paidAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "Invoices_userId_idx" ON "Invoices" ("userId");
CREATE INDEX "Invoices_planType_idx" ON "Invoices" ("planType");
CREATE INDEX "Invoices_paymentType_idx" ON "Invoices" ("paymentType");
CREATE INDEX "Invoices_amount_idx" ON "Invoices" ("amount");
CREATE INDEX "Invoices_tax_idx" ON "Invoices" ("tax");
CREATE INDEX "Invoices_totalAmount_idx" ON "Invoices" ("totalAmount");
CREATE INDEX "Invoices_paidAt_idx" ON "Invoices" ("paidAt");

CREATE INDEX "Invoices_CreatedAt_idx" ON "Invoices" ("CreatedAt");
CREATE INDEX "Invoices_UpdatedAt_idx" ON "Invoices" ("UpdatedAt");
CREATE INDEX "Invoices_DeletedAt_idx" ON "Invoices" ("DeletedAt");
