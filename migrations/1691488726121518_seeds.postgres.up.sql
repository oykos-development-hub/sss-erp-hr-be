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
    id, user_profile_id, first_name, middle_name, last_name, father_name, mother_name,
    mother_birth_last_name, date_of_birth, country_of_birth, city_of_birth, nationality, citizenship,
    address, official_personal_id, gender, employee_relationship, insurance_coverage, handicapped_person,
    created_at, updated_at)
VALUES
    (1, 1, 'Sara', NULL, 'Perović', 'Nikola', 'Jana', 'Ilić', '2010-08-05', 'Montenegro', 'Podgorica', 'Montenegrin', 'Montenegrin', 'Ulica Vladana Giljena 21', 'ID10001', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (2, 1, 'Ivan', NULL, 'Perović', 'Nikola', 'Jana', 'Ilić', '2012-12-05', 'Montenegro', 'Podgorica', 'Montenegrin', 'Montenegrin', 'Ulica Vladana Giljena 21', 'ID10002', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW()),
    (3, 2, 'Igor', NULL, 'Vuković', 'Igor', 'Jelena', 'Mirković', '2012-03-20', 'Montenegro', 'Nikšić', 'Montenegrin', 'Montenegrin', 'Ulica Nikole Tesle bb', 'ID20003', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW()),
    (4, 2, 'Lena', NULL, 'Vuković', 'Igor', 'Jelena', 'Mirković', '2015-11-11', 'Montenegro', 'Nikšić', 'Montenegrin', 'Montenegrin', 'Ulica Nikole Tesle bb', 'ID20004', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (5, 3, 'Maja', NULL, 'Milošević', 'Petar', 'Lena', 'Jovanović', '2005-10-15', 'Montenegro', 'Budva', 'Montenegrin', 'Montenegrin', 'Ulica Mainska 10', 'ID30005', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (6, 3, 'Stefan', NULL, 'Milošević', 'Petar', 'Lena', 'Jovanović', '2010-02-05', 'Montenegro', 'Budva', 'Montenegrin', 'Montenegrin', 'Ulica Mainska 10', 'ID30006', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW()),
    (7, 4, 'Ana', NULL, 'Marković', 'Ana', 'Mila', 'Đorđević', '2015-09-23', 'Montenegro', 'Herceg Novi', 'Montenegrin', 'Montenegrin', 'Bulevar Dobre Vode 45', 'ID40007', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (8, 5, 'Luka', NULL, 'Radulović', 'Marko', 'Lana', 'Nikolić', '2009-01-22', 'Montenegro', 'Tivat', 'Montenegrin', 'Montenegrin', 'Ulica Parkova 3', 'ID50008', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW()),
    (9, 5, 'Mila', NULL, 'Radulović', 'Marko', 'Lana', 'Nikolić', '2011-04-10', 'Montenegro', 'Tivat', 'Montenegrin', 'Montenegrin', 'Ulica Parkova 3', 'ID50009', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (10, 6, 'Ema', NULL, 'Kovačević', 'Marija', 'Tina', 'Simić', '2013-12-05', 'Montenegro', 'Bar', 'Montenegrin', 'Montenegrin', 'Ulica Sunčanih Luka 7', 'ID60010', 'F', 'Kći', 'Zakonsko', false, NOW(), NOW()),
    (11, 6, 'Davor', NULL, 'Kovačević', 'Marija', 'Tina', 'Simić', '2015-05-10', 'Montenegro', 'Bar', 'Montenegrin', 'Montenegrin', 'Ulica Sunčanih Luka 7', 'ID60011', 'M', 'Sin', 'Zakonsko', false, NOW(), NOW());

-- education for each employee
INSERT INTO employee_education (
    user_profile_id, education_type_id, date_of_certification, price, date_of_start, date_of_end,
    academic_title, expertise_level, certificate_issuer, title, description, file_id, created_at, updated_at)
VALUES
    (1, 1, 1, '2015-05-05', 500, '2010-09-01', '2015-05-05', 'Bachelor of Law', 'Intermediate', 'University of Montenegro', 'LLB', 'Completed law degree with honors', 1001, NOW(), NOW()),
    (2, 2, 2, '2014-05-05', 300, '2008-09-01', '2012-05-05', 'Master of Science in Economics', 'Advanced', 'Faculty of Economics, Podgorica', 'MSc', 'Specialization in Finance', 1002, NOW(), NOW()),
    (3, 3, 3, '1999-01-01', NULL, '1990-09-01', '1995-01-01', 'Ph.D. in History', 'Expert', 'University of Arts and Sciences, Budva', 'PhD', 'Thesis on Montenegrin Modern History', 1003, NOW(), NOW()),
    (4, 4, 4, '2018-06-06', 200, '2015-09-01', '2018-06-06', 'Master of Architecture', 'Advanced', 'Faculty of Architecture, Herceg Novi', 'MArch', 'Focus on urban planning', 1004, NOW(), NOW()),
    (5, 5, 5, '2013-04-04', 400, '2009-09-01', '2013-04-04', 'Bachelor of Medicine', 'Intermediate', 'Medical Faculty, Tivat', 'MBBS', 'Specialization in Pediatrics', 1005, NOW(), NOW()),
    (6, 6, 6, '2016-07-07', 350, '2012-09-01', '2016-07-07', 'Master of Fine Arts', 'Advanced', 'Academy of Fine Arts, Bar', 'MFA', 'Concentration in Visual Arts', 1006, NOW(), NOW());


-- organization_units - 5 jedinica, 3 suda (prva 3) i 2 pomocne jedinice (racunovodstvo i it sluzba koji pripadaju prvom sudu id:1)
INSERT INTO organization_units
    (id, parent_id, title, abbreviation, number_of_judges, color, icon, address, description, folder_id,
    created_at, updated_at)
VALUES
    (1, NULL, 'Sud Niksic', 'SNK', 10, '#3498DB', 'icon.png', 'Ulica Nikole Tesle bb', 'desc', 0, NOW(), NOW()),
    (2, NULL, 'Sud Berane', 'SBE', 7, '#3498DB', 'icon.png', 'Ulica Emira Kosute bb', 'desc', 0, NOW(), NOW()),
    (3, NULL, 'Sud Podgorica', 'SPG', 20, '#3498DB', 'icon.png', 'Ulica Vladana Giljena 21', 'desc', 0, NOW(), NOW()),
    (4, 1, 'Racunovodstvo', 'RNK', 0, '#3498DB', 'icon.png', 'Ulica Nikole Tesle bb', 'desc', 0, NOW(), NOW()),
    (5, 1, 'IT Sluzba', 'NKIT', 0, '#3498DB', 'icon.png', 'Ulica Nikole Tesle bb', 'desc', 0, NOW(), NOW());


-- systematizations - dodate su sistematizacije za sudove niksic i berane i it sluzbu i ekonomsku sluzbu u niksic sudu i jedna neaktivna
INSERT INTO systematizations 
    (id, user_profile_id, organization_unit_id, description, serial_number, active, date_of_activation,
    file_id, created_at, updated_at)
VALUES
    (1, 1, 1,'Sistematizacija za 2024 godinu', 'SN001', true, '2023-08-08 12:34:56', null, NOW(), NOW()),
    (2, 2, 2,'Sistematizacija za 2024 godinu', 'SN002', true, '2023-08-08 09:00:00', 101, NOW(), NOW()),
    (3, 3, 1,'Sistematizacija za 2022 godinu', 'SN003', false, null, null, NOW(), NOW());
    (4, 1, 4,'Sistematizacija za 2024 godinu', 'SN004', true, '2023-08-08 12:34:56', null, NOW(), NOW()),
    (5, 2, 5,'Sistematizacija za 2024 godinu', 'SN005', true, '2023-08-08 09:00:00', 101, NOW(), NOW()),


INSERT INTO job_positions
    (id, title, abbreviation, serial_number, requirements, description, is_judge, is_judge_president,
    color, icon, created_at, updated_at)
VALUES
    (1, 'Sudija', 'SU', '53671d', 'Pravni fakultet, 5 godina iskustva', 'Odgovoran za presuđivanje u slučajevima', true, false, '#3498DB', 'icon.png', NOW(), NOW()),
    (2, 'Predsjednik suda', 'PS', '53571d', 'Pravni fakultet, 10 godina iskustva', 'Vodi sud i nadgleda rad sudija', true, true, '#3498DB', 'icon.png', NOW(), NOW()),
    (3, 'Ekonomista', 'EK', '53871d', 'Ekonomski fakultet, 3 godine iskustva', 'Bavi se ekonomskim pitanjima unutar suda', false, false, '#3498DB', 'icon.png', NOW(), NOW()),
    (4, 'IT strucnjak', 'IT', '53971d', 'Fakultet tehničkih nauka, 2 godine iskustva', 'Odgovoran za IT infrastrukturu i softverska rješenja', false, false, '#3498DB', 'icon.png', NOW(), NOW());

-- job position in organization unit
INSERT INTO job_positions_in_organization_units
    (id, systematization_id, parent_organization_unit_id, job_position_id, parent_job_position_id, 
    system_permission_id, description, serial_number, available_slots, requirements, icon, 
    created_at, updated_at)
VALUES
    (2, 1, 1, 2, NULL, NULL, 'Predsjednik Niksic', 'ABC456', 3, NULL, 'icon.png', NOW(), NOW()),
    (3, 4, 4, 3, NULL, NULL, 'Ekonomista Niksic', 'ABC789', 3, NULL, 'icon.png', NOW(), NOW()),
    (5, 2, 2, 2, NULL, NULL, 'Predsjednik Berane', 'BABC123', 3, NULL, 'icon.png', NOW(), NOW()),
    (6, 5, 5, 4, NULL, NULL, 'IT Niksic', 'NABC789', 3, NULL, 'icon.png', NOW(), NOW()),
    (7, 3, 1, 1, 2, NULL, 'Stari sudija Niksic', 'NABC789', 3, NULL, 'icon.png', NOW(), NOW()),
    (1, 1, 1, 1, 2, NULL, 'Sudija Niksic', 'ABC123', 3, NULL, 'icon.png', NOW(), NOW()),
    (4, 2, 2, 1, 5, NULL, 'Sudija Berane', 'ABC124', 3, NULL, 'icon.png', NOW(), NOW())