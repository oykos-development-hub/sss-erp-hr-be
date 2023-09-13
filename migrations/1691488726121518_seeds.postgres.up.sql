INSERT INTO user_profiles (
    user_account_id, first_name, middle_name, last_name, birth_last_name, father_name, mother_name,
    mother_birth_last_name, date_of_birth, country_of_birth, city_of_birth, nationality, citizenship,
    address, bank_account, bank_name, official_personal_id, official_personal_document_number,
    official_personal_document_issuer, gender, single_parent, housing_done, housing_description,
    martial_status, date_of_taking_oath, date_of_becoming_judge, active_contract, revisor_role,
    engagement_type_id, created_at, updated_at)
VALUES
    (1, 'Nikola', NULL, 'Perović', NULL, 'Marko', 'Mila', NULL, '1990-05-05', 'Montenegro', 'Podgorica', 'Montenegrin', 'Montenegrin', 'Ulica Vladana Giljena 21', 'ME123456789', 'Erste Bank', 'ID12345', 'DOC67890', 'MUP Montenegro', 'M', false, true, 'Stan u Nikšiću', 'Oženjen', '2015-05-05', NULL, true, true, 1, NOW(), NOW()),
    (2, 'Jelena', NULL, 'Vuković', NULL, 'Igor', 'Sanja', NULL, '1985-03-15', 'Montenegro', 'Nikšić', 'Montenegrin', 'Montenegrin', 'Ulica Nikole Tesle bb', 'ME987654321', 'NLB Banka', 'ID67890', 'DOC12345', 'MUP Montenegro', 'F', false, true, 'Stan u Podgorica', 'Udata', '2014-05-05', '2008-05-05', true, true, 1, NOW(), NOW()),
    (3, 'Petar', NULL, 'Milošević', NULL, 'Nenad', 'Jelena', NULL, '1972-11-30', 'Montenegro', 'Budva', 'Montenegrin', 'Montenegrin', 'Ulica Mainska 10', 'ME112233445', 'Societe Generale', 'ID33333', 'DOC44444', 'MUP Montenegro', 'M', false, true, 'Kuća u Budvi', 'Razveden', '2000-01-01', '1995-01-01', true, false, 2, NOW(), NOW()),
    (4, 'Ana', NULL, 'Marković', NULL, 'Vladimir', 'Marija', NULL, '1993-09-23', 'Montenegro', 'Herceg Novi', 'Montenegrin', 'Montenegrin', 'Bulevar Dobre Vode 45', 'ME998877665', 'Lovćen Banka', 'ID22222', 'DOC55555', 'MUP Montenegro', 'F', true, true, 'Stan u Herceg Novom', 'Slobodna', '2018-06-06', '2015-06-06', true, true, 3, NOW(), NOW()),
    (5, 'Marko', NULL, 'Radulović', NULL, 'Zoran', 'Lana', NULL, '1988-04-17', 'Montenegro', 'Tivat', 'Montenegrin', 'Montenegrin', 'Ulica Parkova 3', 'ME554433221', 'Hipotekarna Banka', 'ID11111', 'DOC66666', 'MUP Montenegro', 'M', false, true, 'Stan u Tivtu', 'Oženjen', '2013-04-04', '2011-04-04', true, false, 2, NOW(), NOW()),
    (6, 'Marija', NULL, 'Kovačević', NULL, 'Dušan', 'Tina', NULL, '1991-12-12', 'Montenegro', 'Bar', 'Montenegrin', 'Montenegrin', 'Ulica Sunčanih Luka 7', 'ME332211009', 'Komercijalna Banka', 'ID99999', 'DOC77777', 'MUP Montenegro', 'F', false, false, 'Stan u Baru', 'Udata', '2016-07-07', NULL, true, false, 5, NOW(), NOW());

-- family members for all 6 users
INSERT INTO employee_family_members (
    user_profile_id, first_name, middle_name, last_name, father_name, mother_name,
    mother_birth_last_name, date_of_birth, country_of_birth, city_of_birth, nationality, citizenship,
    address, official_personal_id, gender, employee_relationship, insurance_coverage, handicapped_person,
    created_at, updated_at)
VALUES
    (1, 'Sara', NULL, 'Perović', 'Nikola', 'Jana', 'Ilić', '2010-08-05', 'Montenegro', 'Podgorica', 'Montenegrin', 'Montenegrin', 'Ulica Vladana Giljena 21', 'ID10001', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (1, 'Ivan', NULL, 'Perović', 'Nikola', 'Jana', 'Ilić', '2012-12-05', 'Montenegro', 'Podgorica', 'Montenegrin', 'Montenegrin', 'Ulica Vladana Giljena 21', 'ID10002', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW()),
    (2, 'Igor', NULL, 'Vuković', 'Igor', 'Jelena', 'Mirković', '2012-03-20', 'Montenegro', 'Nikšić', 'Montenegrin', 'Montenegrin', 'Ulica Nikole Tesle bb', 'ID20003', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW()),
    (2, 'Lena', NULL, 'Vuković', 'Igor', 'Jelena', 'Mirković', '2015-11-11', 'Montenegro', 'Nikšić', 'Montenegrin', 'Montenegrin', 'Ulica Nikole Tesle bb', 'ID20004', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (3, 'Maja', NULL, 'Milošević', 'Petar', 'Lena', 'Jovanović', '2005-10-15', 'Montenegro', 'Budva', 'Montenegrin', 'Montenegrin', 'Ulica Mainska 10', 'ID30005', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (3, 'Stefan', NULL, 'Milošević', 'Petar', 'Lena', 'Jovanović', '2010-02-05', 'Montenegro', 'Budva', 'Montenegrin', 'Montenegrin', 'Ulica Mainska 10', 'ID30006', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW()),
    (4, 'Ana', NULL, 'Marković', 'Ana', 'Mila', 'Đorđević', '2015-09-23', 'Montenegro', 'Herceg Novi', 'Montenegrin', 'Montenegrin', 'Bulevar Dobre Vode 45', 'ID40007', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (5, 'Luka', NULL, 'Radulović', 'Marko', 'Lana', 'Nikolić', '2009-01-22', 'Montenegro', 'Tivat', 'Montenegrin', 'Montenegrin', 'Ulica Parkova 3', 'ID50008', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW()),
    (5, 'Mila', NULL, 'Radulović', 'Marko', 'Lana', 'Nikolić', '2011-04-10', 'Montenegro', 'Tivat', 'Montenegrin', 'Montenegrin', 'Ulica Parkova 3', 'ID50009', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (6, 'Ema', NULL, 'Kovačević', 'Marija', 'Tina', 'Simić', '2013-12-05', 'Montenegro', 'Bar', 'Montenegrin', 'Montenegrin', 'Ulica Sunčanih Luka 7', 'ID60010', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (6, 'Davor', NULL, 'Kovačević', 'Marija', 'Tina', 'Simić', '2015-05-10', 'Montenegro', 'Bar', 'Montenegrin', 'Montenegrin', 'Ulica Sunčanih Luka 7', 'ID60011', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW());

-- organization_units - 5 jedinica, 3 suda (prva 3) i 2 pomocne jedinice (racunovodstvo i it sluzba koji pripadaju prvom sudu id:1)
INSERT INTO organization_units
    (parent_id, title, abbreviation, number_of_judges, color, icon, address, description, folder_id,
    created_at, updated_at)
VALUES
    (NULL, 'Sud Niksic', 'SNK', 10, '#3498DB', 'icon.png', 'Ulica Nikole Tesle bb', 'desc', 0, NOW(), NOW()),
    (NULL, 'Sud Berane', 'SBE', 7, '#3498DB', 'icon.png', 'Ulica Emira Kosute bb', 'desc', 0, NOW(), NOW()),
    (NULL, 'Sud Podgorica', 'SPG', 20, '#3498DB', 'icon.png', 'Ulica Vladana Giljena 21', 'desc', 0, NOW(), NOW()),
    (1, 'Racunovodstvo', 'RNK', 0, '#3498DB', 'icon.png', 'Ulica Nikole Tesle bb', 'desc', 0, NOW(), NOW()),
    (1, 'IT Sluzba', 'NKIT', 0, '#3498DB', 'icon.png', 'Ulica Nikole Tesle bb', 'desc', 0, NOW(), NOW());


-- systematizations - dodate su sistematizacije za sudove niksic i berane i it sluzbu i ekonomsku sluzbu u niksic sudu i jedna neaktivna
INSERT INTO systematizations 
    (user_profile_id, organization_unit_id, description, serial_number, active, date_of_activation,
    file_id, created_at, updated_at)
VALUES
    (1, 1, 'Sistematizacija za 2024 godinu', 'SN001', true, '2023-08-08 12:34:56', null, NOW(), NOW()),
    (2, 2, 'Sistematizacija za 2024 godinu', 'SN002', true, '2023-08-08 09:00:00', 101, NOW(), NOW()),
    (3, 1, 'Sistematizacija za 2022 godinu', 'SN003', false, null, null, NOW(), NOW()),
    (1, 4, 'Sistematizacija za 2024 godinu', 'SN004', true, '2023-08-08 12:34:56', null, NOW(), NOW()),
    (2, 5, 'Sistematizacija za 2024 godinu', 'SN005', true, '2023-08-08 09:00:00', 101, NOW(), NOW());


INSERT INTO job_positions
    (title, abbreviation, serial_number, requirements, description, is_judge, is_judge_president,
    color, icon, created_at, updated_at)
VALUES
    ('Sudija', 'SU', '53671d', 'Pravni fakultet, 5 godina iskustva', 'Odgovoran za presuđivanje u slučajevima', true, false, '#3498DB', 'icon.png', NOW(), NOW()),
    ('Predsjednik suda', 'PS', '53571d', 'Pravni fakultet, 10 godina iskustva', 'Vodi sud i nadgleda rad sudija', true, true, '#3498DB', 'icon.png', NOW(), NOW()),
    ('Ekonomista', 'EK', '53871d', 'Ekonomski fakultet, 3 godine iskustva', 'Bavi se ekonomskim pitanjima unutar suda', false, false, '#3498DB', 'icon.png', NOW(), NOW()),
    ('IT strucnjak', 'IT', '53971d', 'Fakultet tehničkih nauka, 2 godine iskustva', 'Odgovoran za IT infrastrukturu i softverska rješenja', false, false, '#3498DB', 'icon.png', NOW(), NOW());

-- job position in organization unit
INSERT INTO job_positions_in_organization_units
    (systematization_id, parent_organization_unit_id, job_position_id, available_slots)
VALUES
    (1, 1, 2, 3),
    (4, 4, 3, 3),
    (2, 2, 2, 3),
    (5, 5, 4, 3),
    (3, 1, 1, 3),
    (1, 1, 1, 3),
    (2, 2, 1, 3);


-- add employees to job positions (job_positions_in_organization_units)
INSERT INTO employees_in_organization_units 
    (user_account_id, user_profile_id, position_in_organization_unit_id, active, created_at, updated_at)
VALUES
    (2, 2, 2, true, NOW(), NOW()),
    (3, 3, 1, true, NOW(), NOW()),
    (4, 4, 5, true, NOW(), NOW()),
    (5, 5, 4, true, NOW(), NOW()),
    (6, 6, 5, true, NOW(), NOW());


-- contracts
INSERT INTO employee_contracts (user_profile_id, contract_type_id, organization_unit_id, organization_unit_department_id, abbreviation, description, active, serial_number, net_salary, gross_salary, bank_account, bank_name, date_of_signature, date_of_eligibility, date_of_start, date_of_end, file_id, created_at, updated_at)
VALUES
    (2, 6, 1, null, 'CT1', 'Predsjednik Niksic', true, '123456', '2000', '2500', '123456789', 'Banka A', '2023-08-01', '2023-09-01', '2023-09-15', null, 101, NOW(), NOW()),
    (3, 7, 2, null, 'CT2', 'Sudija Niksic', true, '654321', '1800', '1200', '987654321', 'Banka B', '2023-08-02', '2023-09-02', '2023-09-20', '2024-09-19', 102, NOW(), NOW()),
    (4, 7, 3, null, 'CT3', 'Predsjednik Berane', true, '987654', '2200', '2800', '456789123', 'Banka C', '2023-08-03', '2023-09-03', '2023-10-01', '2024-10-01', 103, NOW(), NOW()),
	(5, 6, 1, null, 'CT5', 'Sudija Berane', true, '123456', '2000', '1200', '123456789', 'Banka A', '2023-08-01', '2023-09-01', '2023-09-15', null, 101, NOW(), NOW()),
    (6, 7, 2, null, 'CT6', 'IT strucnjak', true, '654321', '1800', '600', '987654321', 'Banka B', '2023-08-02', '2023-09-02', '2023-09-20', '2024-09-19', 102, NOW(), NOW()),
	(5, 6, 3, null, 'CT6', 'Sudija Berane neaktivan', false, '654321', '1800', '1000', '987654321', 'Banka B', '2022-08-02', '2022-09-02', '2022-09-20', '2022-10-19', 102, NOW(), NOW());

-- experiences for each employee
INSERT INTO employee_experiences (
    user_profile_id, relevant, organization_unit, organization_unit_id, amount_of_experience,
    amount_of_insured_experience, date_of_signature, date_of_start, date_of_end, file_id, created_at, updated_at)
VALUES
    (1, true, 'Advokatska firma Beograd', NULL, 2, 2, '2015-01-01', '2013-01-01', '2015-01-01', 2001, NOW(), NOW()),
    (1, true, NULL, 3, 3, 2, '2018-01-01', '2015-01-01', '2018-01-01', 2007, NOW(), NOW()),
    (2, true, 'Notarski ured Tuzla', NULL, 4, 3, '2016-01-01', '2012-01-01', '2016-01-01', 2002, NOW(), NOW()),
    (2, true, NULL, 2, 5, 4, '2017-01-01', '2012-01-01', '2017-01-01', 2008, NOW(), NOW()),
    (3, true, NULL, 4, 6, 5, '2018-01-01', '2012-01-01', '2018-01-01', 2003, NOW(), NOW()),
    (3, true, NULL, 2, 3, 2, '2014-01-01', '2011-01-01', '2014-01-01', 2009, NOW(), NOW()),
    (4, true, 'Osnovni sud Herceg Novi', NULL, 4, 3, '2018-06-06', '2014-06-06', '2018-06-06', 2010, NOW(), NOW()),
    (4, true, NULL, 3, 3, 3, '2018-04-04', '2015-04-04', '2018-04-04', 2005, NOW(), NOW()),
    (5, true, 'Tužilaštvo Tivat', NULL, 2, 1, '2017-07-07', '2015-07-07', '2017-07-07', 2006, NOW(), NOW()),
    (5, true, NULL, 4, 4, 4, '2017-04-04', '2013-04-04', '2017-04-04', 2011, NOW(), NOW()),
    (6, true, 'IT Kompanija Bar', NULL, 3, 2, '2016-07-07', '2013-07-07', '2016-07-07', 2012, NOW(), NOW()),
    (6, true, 'Zavod za informatičke tehnologije Bar', NULL, 4, 3, '2015-12-12', '2011-12-12', '2015-12-12', 2013, NOW(), NOW()),
    (6, true, NULL, 5, 5, 4, '2016-12-12', '2011-12-12', '2016-12-12', 2014, NOW(), NOW());

-- educations
INSERT INTO employee_educations (
    user_profile_id, type_id, date_of_certification, date_of_start, date_of_end, academic_title, expertise_level, certificate_issuer, title, description, created_at, updated_at)
VALUES
    (2, 16, '2020-01-01', '2016-01-01', '2019-01-01', 'Bachelor prava', '6', 'Univerzitet Crne Gore', 'Pravne studije', 'Specijalističke studije prava', NOW(), NOW()),
    (3, 16, '2021-06-01', '2017-06-01', '2020-06-01', 'Bachelor kriminalistike', '6', 'Pravni fakultet', 'Kriminalistika', 'Studije kriminalistike', NOW(), NOW()),
    (4, 17, '2019-05-01', '2015-05-01', '2018-05-01', 'Magistar prava', '7', 'Univerzitet Crne Gore', 'Pravo', 'Magistarske studije prava', NOW(), NOW()),
    (5, 16, '2018-09-01', '2014-09-01', '2017-09-01', 'Bachelor', '6', 'Univerzitet Crne Gore', 'Pravne studije', 'Bachelor studije prava', NOW(), NOW()),
    (6, 17, '2023-06-01', '2019-06-01', '2022-06-01', 'Master inženjer informacionih tehnologija', '7', 'Tehnički fakultet', 'IT u pravnom sistemu', 'Postdiplomske studije u IT i pravu', NOW(), NOW());


INSERT INTO employee_resolutions
    (resolution_type_id, user_profile_id, resolution_purpose, date_of_start, date_of_end, 
    file_id, created_at, updated_at)
VALUES
    (8, 3, 'Svadja sa nadleznim', '2022-09-01', '2022-12-01', NULL, NOW(), NOW()),
    (9, 4, 'Povecanje plate', '2022-09-01', '2022-12-01', NULL, NOW(), NOW()),
    (10, 3, 'Nedozvoljene materije na radnom mjestu', '2022-09-01', '2022-12-01' ,NULL, NOW(), NOW());

-- absent types and absents
INSERT INTO absent_types (
    parent_id, title, abbreviation, accounting_days_off, relocation, description, color, icon, created_at, updated_at)
VALUES
    (NULL, 'Bolovanje.', 'OZB', false, false, 'Odsustvo zbog bolovanja.', '#FF0000', 'icon-sick', NOW(), NOW()),
    (NULL, 'Godišnji odmor', 'GO', true, false, 'Zakonski godišnji odmor', '#00FF00', 'icon-vacation', NOW(), NOW()),
    (NULL, 'Upućivanje u drugi državni organ', 'UD', false, true, 'Odsustvo zbog upućivanja u drugi državni organ.', '#FFFF00', 'icon-business-trip', NOW(), NOW()),
    (NULL, 'Roditeljstvo', 'OBO', false, false, 'Odsustvo zbog roditeljstva.', '#00FFFF', 'icon-training', NOW(), NOW());

INSERT INTO employee_absents (
    absent_type_id, user_profile_id, target_organization_unit_id, description, date_of_start, date_of_end, location, file_id, created_at, updated_at)
VALUES
    (1, 1, NULL, 'Odsustvo zbog gripa', '2023-08-01', '2023-08-07', 'Podgorica', NULL, NOW(), NOW()),
    (2, 1, NULL, 'Godišnji odmor', '2023-07-01', '2023-07-15', 'Budva', NULL, NOW(), NOW()),
    
    (3, 2, NULL, 'Sudsko saslušanje', '2023-08-05', '2023-08-05', 'Nikšić', NULL, NOW(), NOW()),
    (1, 2, NULL, 'Odsustvo zbog gripa', '2022-01-01', '2022-04-01', 'Podgorica', NULL, NOW(), NOW()),
    
    (4, 3, 3, 'Službeni put u inostranstvo', '2023-06-01', '2023-06-10', 'Budva', NULL, NOW(), NOW()),
    
    (1, 4, NULL, 'Odsustvo zbog bolesti', '2023-05-10', '2023-05-20', 'Herceg Novi', NULL, NOW(), NOW()),
    (2, 4, NULL, 'Godišnji odmor', '2023-09-01', '2023-09-15', 'Herceg Novi', NULL, NOW(), NOW()),
    
    (4, 5, 3, 'Službeni put', '2023-04-01', '2023-04-05', 'Tivat', NULL, NOW(), NOW()),
    
    (2, 6, NULL, 'Godišnji odmor', '2023-07-01', '2023-07-15', 'Bar', NULL, NOW(), NOW()),
    (5, 6, NULL, 'Obuka za IT stručnjaka', '2023-05-15', '2023-05-20', 'Bar', NULL, NOW(), NOW());

INSERT INTO evaluations (user_profile_id, evaluation_type_id, score, date_of_evaluation, evaluator, is_relevant, created_at, updated_at, file_id)
VALUES
    (3, 19, 'Dobar', '2023-08-01', 'Bozo Markovic', true, NOW(), NOW(), null),
    (4, 20, 'Zadovoljio', '2023-08-02', 'Bozo Markovic', true, NOW(), NOW(), null),
    (5, 21, 'Nije zadovoljio', '2023-08-03', 'Vladan Giljen', false, NOW(), NOW(), null);

INSERT INTO foreigners (
    user_profile_id,
    work_permit_number,
    work_permit_issuer,
    work_permit_date_of_start,
    work_permit_date_of_end,
    work_permit_indefinite_length,
    residence_permit_number,
    residence_permit_date_of_end,
    residence_permit_indefinite_length,
    country_of_origin,
    created_at,
    updated_at,
    work_permit_file_id,
    residence_permit_file_id
)
VALUES
    (3, 'WP123', 'Ministry of Labor', '2023-08-01', '2025-08-01', false, 'A123', '2025-08-01', false, 'Madjarska', NOW(), NOW(), null, null),
    (4, 'WP456', 'Immigration Office', '2023-08-15', '2024-08-15', false, 'A123', '2024-08-15', false, 'Madjarska', NOW(), NOW(), null, null);

INSERT INTO judge_number_resolutions (active, serial_number, year, created_at, updated_at)
VALUES
    (true, 'JNR001', '2023', NOW(), NOW()),
    (true, 'JNR002', '2023', NOW(), NOW()),
    (false, 'JNR003', '2023', NOW(), NOW());

INSERT INTO judge_number_resolution_organization_units (resolution_id, organization_unit_id, number_of_judges, number_of_presidents, created_at, updated_at)
VALUES
    (1, 1, 10, 1, NOW(), NOW()),
    (2, 2, 5, 1, NOW(), NOW()),
    (3, 1, 8, 2, NOW(), NOW());

INSERT INTO salaries (
    user_profile_id,
    organization_unit_id,
    benefited_track,
    without_raise,
    insurance_basis,
    salary_rank,
    daily_work_hours,
    weekly_work_hours,
    education_rank,
    education_naming,
    created_at,
    updated_at,
    user_resolution_id
)
VALUES
    (3, 1, false, true, 'Basis B', 'Rank 2', '8', '40', 'Degree 5', 'UCG', NOW(), NOW(), 3),
    (4, 1, true, false, 'Basis C', 'Rank 4', '7', '35', 'Degree 4', 'UDG', NOW(), NOW(), 2);


INSERT INTO user_norm_fulfilments (
    user_profile_id,
    topic,
    title,
    number_of_norm_decrease,
    number_of_items,
    number_of_items_solved,
    evaluation_id,
    date_of_evaluation,
    date_of_evaluation_validity,
    relocation_id,
    file_id,
    created_at,
    updated_at
)
VALUES
    (3, 'Topic 1', 'Title 1', 5, 10, 8, 1, '2023-08-01', '2023-08-31', 5, null, NOW(), NOW()),
    (4, 'Topic 2', 'Title 2', 3, 8, 7, 2, '2023-08-02', '2023-08-31', null, null, NOW(), NOW()),
    (5, 'Topic 3', 'Title 3', 1, 12, 3, 3, '2023-08-03', '2023-08-31', 8, null, NOW(), NOW());

INSERT INTO tender_types (title, is_judge, is_judge_president, abbreviation, description, color, value, icon, created_at, updated_at)
VALUES
    ('Konkurs za sudiju', true, false, 'KS1', 'Tender', '#3498DB', NULL, 'icon1.png', NOW(), NOW()),
    ('Konkurs za predsjednika', true, true, 'KP1', 'Tender', '#E74C3C', NULL, 'icon2.png', NOW(), NOW()),
    ('Konkurs za IT strucnjaka', false, false, 'KT1', 'Tender', '#2ECC71', NULL, 'icon3.png', NOW(), NOW());

INSERT INTO tenders_in_organization_units (
	position_in_organization_unit_id,
    type,
    date_of_start,
    date_of_end,
    description,
    serial_number,
    file_id,
    created_at,
    updated_at
)
VALUES
    (1,1, '2023-08-01', '2023-08-31', 'Konkurs za sudiju u Niksicu', 'TSN001', null, NOW(), NOW()),
    (2,2, '2023-08-15', '2023-09-15', 'Konkurs za predsjednika u Niksicu', 'TSN002', null, NOW(), NOW()),
    (3,3, '2023-08-20', '2023-09-30', 'Konkurs za ekonomistu', 'TSN003', null, NOW(), NOW()),
	(1, 1, '2022-08-01', '2022-08-31', 'Stari konkurs za sudiju u Niksicu', 'TSN011', null, NOW(), NOW()),
    (2, 2, '2022-08-15', '2022-09-15', 'Stari konkurs za predsjednika u Niksicu', 'TSN012', null, NOW(), NOW());

INSERT INTO tender_applications_in_organization_units (
    job_tender_id,
    user_profile_id,
    active,
    status,
    file_id,
    is_internal,
    created_at,
    updated_at
)
VALUES
    (1, 2, true, 'Na cekanju' , null, true, NOW(), NOW()),
    (1, 3, true, 'Na cekanju', null, true, NOW(), NOW()),
    (1, 4, false, 'Na cekanju', null, false, NOW(), NOW()),
	(2, 2, true, 'Na cekanju', null, false, NOW(), NOW()),
    (2, 3, true, 'Na cekanju', null, false, NOW(), NOW()),
    (2, 4, false, 'Na cekanju', null, true, NOW(), NOW()),
	(5, 2, false,'Na cekanju', null, true,  NOW(), NOW()),
    (5, 3, false, 'Na cekanju', null, true, NOW(), NOW()),
    (5, 4, false, 'Na cekanju', null, true, NOW(), NOW());