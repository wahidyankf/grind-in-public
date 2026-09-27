---
tldr: "Builds the shared PostgreSQL 18.6 schema and deterministic data used by every lesson."
when_to_use: "Use first to create a disposable database for following the course."
---

# Setup and Study Schema

Start the container from the course README, enter `psql`, and verify the server:

```sql
SELECT current_setting('server_version') AS version,
       current_database() AS database,
       current_user AS role;
```

The domain separates identities, financial events, detections, and human workflow. `tenant_id` appears in every
tenant-owned key so an index or foreign key cannot silently cross tenants.

```text
tenants -> customers -> accounts -> transactions
                        |              |
                        |              +-> alerts <- policy_versions
                        |                    |
                        +--------------------+-> case_alerts -> cases

transactional writes -------------------------------------> outbox_events
```

Create the schema:

```sql
CREATE TABLE tenants (
    tenant_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    slug text NOT NULL UNIQUE,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE customers (
    tenant_id bigint NOT NULL REFERENCES tenants,
    customer_id bigint GENERATED ALWAYS AS IDENTITY,
    external_ref text NOT NULL,
    full_name text NOT NULL,
    risk_level text NOT NULL DEFAULT 'standard'
        CHECK (risk_level IN ('standard', 'elevated', 'restricted')),
    created_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (tenant_id, customer_id),
    UNIQUE (tenant_id, external_ref)
);

CREATE TABLE accounts (
    tenant_id bigint NOT NULL,
    account_id bigint GENERATED ALWAYS AS IDENTITY,
    customer_id bigint NOT NULL,
    currency char(3) NOT NULL,
    status text NOT NULL CHECK (status IN ('open', 'frozen', 'closed')),
    opened_at timestamptz NOT NULL,
    PRIMARY KEY (tenant_id, account_id),
    FOREIGN KEY (tenant_id, customer_id)
        REFERENCES customers (tenant_id, customer_id)
);

CREATE TABLE policy_versions (
    tenant_id bigint NOT NULL REFERENCES tenants,
    policy_id bigint NOT NULL,
    version integer NOT NULL CHECK (version > 0),
    definition jsonb NOT NULL,
    active_from timestamptz NOT NULL,
    PRIMARY KEY (tenant_id, policy_id, version)
);

CREATE TABLE transactions (
    tenant_id bigint NOT NULL,
    transaction_id bigint GENERATED ALWAYS AS IDENTITY,
    account_id bigint NOT NULL,
    idempotency_key text NOT NULL,
    amount numeric(18, 2) NOT NULL CHECK (amount > 0),
    direction text NOT NULL CHECK (direction IN ('credit', 'debit')),
    counterparty text NOT NULL,
    occurred_at timestamptz NOT NULL,
    recorded_at timestamptz NOT NULL DEFAULT now(),
    metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
    PRIMARY KEY (tenant_id, transaction_id),
    UNIQUE (tenant_id, idempotency_key),
    FOREIGN KEY (tenant_id, account_id)
        REFERENCES accounts (tenant_id, account_id)
);

CREATE TABLE alerts (
    tenant_id bigint NOT NULL,
    alert_id bigint GENERATED ALWAYS AS IDENTITY,
    transaction_id bigint NOT NULL,
    policy_id bigint NOT NULL,
    policy_version integer NOT NULL,
    score numeric(5, 4) NOT NULL CHECK (score BETWEEN 0 AND 1),
    status text NOT NULL CHECK (status IN ('open', 'assigned', 'closed')),
    created_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (tenant_id, alert_id),
    FOREIGN KEY (tenant_id, transaction_id)
        REFERENCES transactions (tenant_id, transaction_id),
    FOREIGN KEY (tenant_id, policy_id, policy_version)
        REFERENCES policy_versions (tenant_id, policy_id, version)
);

CREATE TABLE cases (
    tenant_id bigint NOT NULL REFERENCES tenants,
    case_id bigint GENERATED ALWAYS AS IDENTITY,
    title text NOT NULL,
    status text NOT NULL CHECK (status IN ('open', 'review', 'closed')),
    assignee text,
    version integer NOT NULL DEFAULT 1,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (tenant_id, case_id)
);

CREATE TABLE case_alerts (
    tenant_id bigint NOT NULL,
    case_id bigint NOT NULL,
    alert_id bigint NOT NULL,
    added_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (tenant_id, case_id, alert_id),
    FOREIGN KEY (tenant_id, case_id) REFERENCES cases (tenant_id, case_id),
    FOREIGN KEY (tenant_id, alert_id) REFERENCES alerts (tenant_id, alert_id)
);

CREATE TABLE outbox_events (
    tenant_id bigint NOT NULL REFERENCES tenants,
    event_id bigint GENERATED ALWAYS AS IDENTITY,
    aggregate_type text NOT NULL,
    aggregate_id bigint NOT NULL,
    event_type text NOT NULL,
    payload jsonb NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    published_at timestamptz,
    PRIMARY KEY (tenant_id, event_id)
);
```

Seed a small deterministic dataset:

```sql
INSERT INTO tenants (slug) VALUES ('north'), ('south');

INSERT INTO customers (tenant_id, external_ref, full_name, risk_level) VALUES
    (1, 'C-100', 'Avery Stone', 'standard'),
    (1, 'C-101', 'Robin Vale', 'elevated'),
    (2, 'C-100', 'Sam River', 'standard');

INSERT INTO accounts (tenant_id, customer_id, currency, status, opened_at) VALUES
    (1, 1, 'USD', 'open', '2026-01-01T00:00:00Z'),
    (1, 2, 'EUR', 'open', '2026-01-02T00:00:00Z'),
    (2, 3, 'USD', 'open', '2026-01-03T00:00:00Z');

INSERT INTO policy_versions (tenant_id, policy_id, version, definition, active_from) VALUES
    (1, 1, 1, '{"threshold": 1000}', '2026-01-01T00:00:00Z'),
    (2, 1, 1, '{"threshold": 1500}', '2026-01-01T00:00:00Z');

INSERT INTO transactions
    (tenant_id, account_id, idempotency_key, amount, direction, counterparty, occurred_at, metadata)
VALUES
    (1, 1, 'n-001', 250.00, 'debit', 'Orchid Goods', '2026-01-10T09:00:00Z', '{"channel":"card"}'),
    (1, 1, 'n-002', 1800.00, 'debit', 'Cedar Export', '2026-01-10T09:05:00Z', '{"channel":"wire"}'),
    (1, 2, 'n-003', 700.00, 'credit', 'Maple Works', '2026-01-11T10:00:00Z', '{"channel":"wire"}'),
    (2, 3, 's-001', 2100.00, 'debit', 'Juniper Trade', '2026-01-11T11:00:00Z', '{"channel":"wire"}');
```

For performance labs, generate synthetic rows only after the small answers are understood:

```sql
INSERT INTO transactions
    (tenant_id, account_id, idempotency_key, amount, direction, counterparty, occurred_at)
SELECT 1, 1, 'load-' || g, 1 + (g % 5000),
       CASE WHEN g % 3 = 0 THEN 'credit' ELSE 'debit' END,
       'counterparty-' || (g % 1000),
       '2026-01-01T00:00:00Z'::timestamptz + (g || ' seconds')::interval
FROM generate_series(1, 100000) AS g;
```

Checkpoint: `SELECT tenant_id, count(*) FROM transactions GROUP BY tenant_id ORDER BY tenant_id;` returns `3` and `1`
before the optional load seed.
