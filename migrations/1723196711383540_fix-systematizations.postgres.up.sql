INSERT INTO organization_units (parent_id, title, abbreviation, pib, city, order_id, number_of_judges, color, icon, address, description, folder_id, bank_accounts, code, created_at, updated_at)
SELECT 
    id AS parent_id,
    'Odjeljenje' AS title,
    abbreviation,
    pib,
    city,
    order_id,
    number_of_judges,
    color,
    icon,
    address,
    description,
    folder_id,
    bank_accounts,
    code,
    NOW() AS created_at,
    NOW() AS updated_at
FROM 
    organization_units
WHERE 
    parent_id IS NULL
RETURNING id, parent_id;

UPDATE job_positions_in_organization_units jo
SET parent_organization_unit_id = new_ou.id
FROM (
    SELECT 
        ou.id AS old_parent_id, 
        new_ou.id AS id
    FROM 
        organization_units ou
    JOIN 
        organization_units new_ou ON new_ou.parent_id = ou.id
    WHERE 
        ou.parent_id IS NULL
) AS new_ou
WHERE 
    jo.parent_organization_unit_id = new_ou.old_parent_id;
