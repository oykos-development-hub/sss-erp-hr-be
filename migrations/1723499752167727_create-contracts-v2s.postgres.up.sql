INSERT INTO employee_contracts (
    user_profile_id,
    contract_type_id,
    organization_unit_id,
    organization_unit_department_id,
    number_of_conference,
    abbreviation,
    description,
    active,
    serial_number,
    net_salary,
    gross_salary,
    bank_account,
    bank_name,
    date_of_signature,
    date_of_eligibility,
    date_of_start,
    date_of_end,
    file_id,
    created_at,
    updated_at
)
VALUES
    (1, 7, 2, 31, 'CON12345', 'ABBR', 'Opis', TRUE, 'SN123456', '1000.00', '1500.00', '1234567890123456', 'Banka', NOW()::DATE, NOW()::DATE, NOW()::DATE, NULL, NULL, NOW(), NOW()),

    (2, 7, 1, 30, 'CON12346', 'ABBR', 'Opis', TRUE, 'SN123457', '1000.00', '1500.00', '1234567890123457', 'Banka', NOW()::DATE, NOW()::DATE, NOW()::DATE, NULL, NULL, NOW(), NOW()),

    (3, 7, 2, 31, 'CON12347', 'ABBR', 'Opis', TRUE, 'SN123458', '1000.00', '1500.00', '1234567890123458', 'Banka', NOW()::DATE, NOW()::DATE, NOW()::DATE, NULL, NULL, NOW(), NOW());

INSERT INTO employee_contracts (
    user_profile_id,
    contract_type_id,
    organization_unit_id,
    organization_unit_department_id,
    number_of_conference,
    abbreviation,
    description,
    active,
    serial_number,
    net_salary,
    gross_salary,
    bank_account,
    bank_name,
    date_of_signature,
    date_of_eligibility,
    date_of_start,
    date_of_end,
    file_id,
    created_at,
    updated_at
)
SELECT
    id AS user_profile_id,
    7 AS contract_type_id,
    id AS organization_unit_id, 
    id + 25 AS organization_unit_department_id, 
    'CON1234' || id AS number_of_conference,
    'ABBR' AS abbreviation,
    'Opis' AS description,
    TRUE AS active,
    'SN12345' || id AS serial_number,
    '1000.00' AS net_salary,
    '1500.00' AS gross_salary,
    '1234567890123' || id AS bank_account,
    'Banka' AS bank_name,
    NOW()::DATE AS date_of_signature,
    NOW()::DATE AS date_of_eligibility,
    NOW()::DATE AS date_of_start,
    NULL AS date_of_end,
    NULL AS file_id,
    NOW() AS created_at,
    NOW() AS updated_at
FROM
    user_profiles
WHERE
    id > 3 AND id < 28;
