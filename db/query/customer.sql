-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = sqlc.arg('id')
LIMIT 1;

-- name: ListCustomer :many
WITH revenue AS (
    SELECT customer,
    COALESCE(SUM(total_price), 0)::float AS total_revenue,
    COALESCE(COUNT(id), 0)::int AS total_orders
    FROM orders
    GROUP BY customer
)
SELECT *, r.total_revenue, r.total_orders FROM customers c
LEFT JOIN revenue r ON c.id = r.customer 
WHERE c.company = sqlc.arg(company)::int
AND (
    c.full_name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    c.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    c.phone ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -c.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: CountCustomer :one
SELECT COUNT(*) FROM customers
WHERE company = sqlc.arg(company)::int;

-- name: CreateCustomer :one
INSERT INTO customers (
    full_name, code, company, address, email, phone, gender, license, issued_by, birthday, user_updated, user_created, "group", 
    title, license_date, contact_name, contact_title, contact_phone, contact_email, contact_address, account_number,
    bank_name, bank_branch
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23
) RETURNING *;

-- name: DetailCustomer :one
SELECT * FROM customers
WHERE id = $1
LIMIT 1;

-- name: UpdateCustomer :one
UPDATE customers
SET
    full_name = COALESCE(sqlc.narg(full_name)::varchar, full_name),
    code = COALESCE(sqlc.narg(code)::varchar, code),
    email = COALESCE(sqlc.narg(email)::varchar, email),
    phone = COALESCE(sqlc.narg(phone)::varchar, phone),
    license = COALESCE(sqlc.narg(license)::varchar, license),
    birthday = COALESCE(sqlc.narg(birthday)::timestamp, birthday),
    user_updated = COALESCE(sqlc.narg(user_updated)::int, user_updated),
    "group" = COALESCE(sqlc.narg('group')::int, "group"),
    title = COALESCE(sqlc.narg(title)::varchar, title),
    gender = COALESCE(sqlc.narg(gender)::gender, gender),
    license_date = COALESCE(sqlc.narg(license_date)::timestamp, license_date),
    contact_name = COALESCE(sqlc.narg(contact_name)::varchar, contact_name),
    contact_title = COALESCE(sqlc.narg(contact_title)::varchar, contact_title),
    contact_phone = COALESCE(sqlc.narg(contact_phone)::varchar, contact_phone),
    contact_email = COALESCE(sqlc.narg(contact_email)::varchar, contact_email),
    contact_address = COALESCE(sqlc.narg(contact_address)::int, contact_address),
    account_number = COALESCE(sqlc.narg(account_number)::varchar, account_number),
    bank_name = COALESCE(sqlc.narg(bank_name)::varchar, bank_name),
    bank_branch = COALESCE(sqlc.narg(bank_branch)::varchar, bank_branch)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: CreateCustomerGroup :one
INSERT INTO customer_group (
    code, name, company, note, user_created
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListCustomerGroup :many
SELECT * FROM customer_group cg
LEFT JOIN accounts ac ON ac.id = cg.user_created
LEFT JOIN accounts au ON au.id = cg.user_updated
WHERE cg.company = sqlc.arg('company')::int
AND (
    cg.name ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%' OR
    cg.code ILIKE '%' || COALESCE(sqlc.narg('search')::varchar, '') || '%'
)
ORDER BY -cg.id
LIMIT COALESCE(sqlc.narg('limit')::int, 10)
OFFSET (COALESCE(sqlc.narg('page')::int, 1) - 1) * COALESCE(sqlc.narg('limit')::int, 10);

-- name: DetailCustomerGroup :one
SELECT * FROM customer_group cg
LEFT JOIN accounts ac ON ac.id = cg.user_created
LEFT JOIN accounts au ON au.id = cg.user_updated
WHERE cg.id = $1;

-- name: UpdateCustomerGroup :one
UPDATE customer_group
SET
    name = COALESCE(sqlc.narg(name), name),
    code = COALESCE(sqlc.narg(code), code),
    note = COALESCE(sqlc.narg(note), note),
    user_updated = sqlc.arg(user_updated),
    updated_at = now()
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteCustomerGroup :one
DELETE FROM customer_group
WHERE id = $1 RETURNING *;

-- name: CreateMedicalRecordLink :one
INSERT INTO medical_record_link (
    uuid, "type", title, url, customer, appointment_schedule, user_created
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: ListMedicalRecordLink :many
SELECT * FROM medical_record_link
WHERE (sqlc.narg(customer)::int IS NULL OR sqlc.narg(customer)::int = customer)
AND (sqlc.narg(type_mrl)::medical_record_link_type IS NULL OR sqlc.narg(type_mrl)::medical_record_link_type = "type")
AND (sqlc.narg(schedule)::uuid IS NULL OR sqlc.narg(schedule)::uuid = appointment_schedule);
