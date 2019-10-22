CREATE TABLE IF NOT EXISTS "SubscriptionInvoices" (
   id serial PRIMARY KEY,
   "subscriptionId" INTEGER NOT NULL REFERENCES "Subscriptions"(id),
   "invoiceId" INTEGER NOT NULL REFERENCES "Invoices"(id),
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "SubscriptionInvoices_subscriptionId_idx" ON "SubscriptionInvoices" ("subscriptionId");
CREATE INDEX "SubscriptionInvoices_invoiceId_idx" ON "SubscriptionInvoices" ("invoiceId");

CREATE INDEX "SubscriptionInvoices_CreatedAt_idx" ON "SubscriptionInvoices" ("CreatedAt");
CREATE INDEX "SubscriptionInvoices_UpdatedAt_idx" ON "SubscriptionInvoices" ("UpdatedAt");
CREATE INDEX "SubscriptionInvoices_DeletedAt_idx" ON "SubscriptionInvoices" ("DeletedAt");
