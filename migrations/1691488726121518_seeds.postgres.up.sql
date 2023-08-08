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

-- add employees to job positions (job_positions_in_organization_units)
INSERT INTO employees_in_organization_units 
    (id, user_account_id, user_profile_id, position_in_organization_unit_id, active, created_at, updated_at)
VALUES
    (1, 2, 2, 2, true, NOW(), NOW()),
    (2, 3, 3, 1, true, NOW(), NOW()),
    (3, 4, 4, 5, true, NOW(), NOW()),
    (4, 5, 5, 4, true, NOW(), NOW()),
    (5, 6, 6, 5, true, NOW(), NOW())


-- contracts
INSERT INTO employee_contracts (user_profile_id, contract_type_id, abbreviation, description, active, serial_number, net_salary, gross_salary, bank_account, bank_name, date_of_signature, date_of_eligibility, date_of_start, date_of_end, file_id, created_at, updated_at)
VALUES
    (2, 6, 'CT1', 'Predsjednik Niksic', true, '123456', '2000', '2500', '123456789', 'Banka A', '2023-08-01', '2023-09-01', '2023-09-15', null, 101, NOW(), NOW()),
    (3, 7, 'CT2', 'Sudija Niksic', true, '654321', '1800', '1200', '987654321', 'Banka B', '2023-08-02', '2023-09-02', '2023-09-20', '2024-09-19', 102, NOW(), NOW()),
    (4, 7, 'CT3', 'Predsjednik Berane', true, '987654', '2200', '2800', '456789123', 'Banka C', '2023-08-03', '2023-09-03', '2023-10-01', '2024-10-01', 103, NOW(), NOW()),
	(5, 6, 'CT5', 'Sudija Berane', true, '123456', '2000', '1200', '123456789', 'Banka A', '2023-08-01', '2023-09-01', '2023-09-15', null, 101, NOW(), NOW()),
    (6, 7, 'CT6', 'IT strucnjak', true, '654321', '1800', '600', '987654321', 'Banka B', '2023-08-02', '2023-09-02', '2023-09-20', '2024-09-19', 102, NOW(), NOW()),
	(7, 6, 'CT6', 'Sudija Berane neaktivan', false, '654321', '1800', '1000', '987654321', 'Banka B', '2022-08-02', '2022-09-02', '2022-09-20', '2022-10-19', 102, NOW(), NOW())

-- experiences for each employee
INSERT INTO employee_experiences (
    id, user_profile_id, relevant, organization_unit, organization_unit_id, amount_of_experience,
    amount_of_insured_experience, date_of_signature, date_of_start, date_of_end, file_id, created_at, updated_at)
VALUES
    (1, 1, true, 'Advokatska firma Beograd', NULL, 2, 2, '2015-01-01', '2013-01-01', '2015-01-01', 2001, NOW(), NOW()),
    (2, 1, true, NULL, 3, 3, 2, '2018-01-01', '2015-01-01', '2018-01-01', 2007, NOW(), NOW()),
    (3, 2, true, 'Notarski ured Tuzla', NULL, 4, 3, '2016-01-01', '2012-01-01', '2016-01-01', 2002, NOW(), NOW()),
    (4, 2, true, NULL, 2, 5, 4, '2017-01-01', '2012-01-01', '2017-01-01', 2008, NOW(), NOW()),
    (5, 3, true, NULL, 4, 6, 5, '2018-01-01', '2012-01-01', '2018-01-01', 2003, NOW(), NOW()),
    (6, 3, true, NULL, 2, 3, 2, '2014-01-01', '2011-01-01', '2014-01-01', 2009, NOW(), NOW()),
    (7, 4, true, 'Osnovni sud Herceg Novi', NULL, 4, 3, '2018-06-06', '2014-06-06', '2018-06-06', 2010, NOW(), NOW()),
    (8, 4, true, NULL, 3, 3, 3, '2018-04-04', '2015-04-04', '2018-04-04', 2005, NOW(), NOW()),
    (9, 5, true, 'Tužilaštvo Tivat', NULL, 2, 1, '2017-07-07', '2015-07-07', '2017-07-07', 2006, NOW(), NOW()),
    (10, 5, true, NULL, 4, 4, 4, '2017-04-04', '2013-04-04', '2017-04-04', 2011, NOW(), NOW()),
    (11, 6, true, 'IT Kompanija Bar', NULL, 3, 2, '2016-07-07', '2013-07-07', '2016-07-07', 2012, NOW(), NOW()),
    (12, 6, true, 'Zavod za informatičke tehnologije Bar', NULL, 4, 3, '2015-12-12', '2011-12-12', '2015-12-12', 2013, NOW(), NOW()),
    (13, 6, true, NULL, 5, 5, 4, '2016-12-12', '2011-12-12', '2016-12-12', 2014, NOW(), NOW());

-- educations
INSERT INTO employee_educations (
    id, user_profile_id, education_type_id, date_of_certification, date_of_start, date_of_end, academic_title, expertise_level, certificate_issuer, title, description, created_at, updated_at)
VALUES
    (1, 2, 16, '2020-01-01', '2016-01-01', '2019-01-01', 'Bachelor prava', '6', 'Univerzitet Crne Gore', 'Pravne studije', 'Specijalističke studije prava', NOW(), NOW()),
    (2, 3, 16, '2021-06-01', '2017-06-01', '2020-06-01', 'Bachelor kriminalistike', '6', 'Pravni fakultet', 'Kriminalistika', 'Studije kriminalistike', NOW(), NOW()),
    (3, 4, 17, '2019-05-01', '2015-05-01', '2018-05-01', 'Magistar prava', '7', 'Univerzitet Crne Gore', 'Pravo', 'Magistarske studije prava', NOW(), NOW()),
    (4, 5, 16, '2018-09-01', '2014-09-01', '2017-09-01', 'Bachelor', '6', 'Univerzitet Crne Gore', 'Pravne studije', 'Bachelor studije prava', NOW(), NOW()),
    (5, 6, 17, '2023-06-01', '2019-06-01', '2022-06-01', 'Master inženjer informacionih tehnologija', '7', 'Tehnički fakultet', 'IT u pravnom sistemu', 'Postdiplomske studije u IT i pravu', NOW(), NOW());


INSERT INTO employee_resolutions
    (id, resolution_type_id, user_profile_id, resolution_purpose, date_of_start, date_of_end, 
    file_id, created_at, updated_at)
VALUES
    (1, 8, 3, 'Svadja sa nadleznim', '2022-09-01', '2022-12-01', NULL, NOW(), NOW()),
    (2, 9, 4, 'Povecanje plate', '2022-09-01', '2022-12-01', NULL, NOW(), NOW()),
    (3, 10, 3, 'Nedozvoljene materije na radnom mjestu', '2022-09-01', '2022-12-01' ,NULL, NOW(), NOW())

-- absent types and absents
INSERT INTO absent_types (
    id, parent_id, title, abbreviation, accounting_days_off, relocation, description, color, icon, created_at, updated_at)
VALUES
    (1, NULL, 'Odsustvo zbog bolesti', 'OBB', false, false, 'Odsustvo zbog bolesti ili medicinskog tretmana', '#FF0000', 'icon-sick', NOW(), NOW()),
    (2, NULL, 'Godišnji odmor', 'GO', true, false, 'Zakonski godišnji odmor', '#00FF00', 'icon-vacation', NOW(), NOW()),
    (3, NULL, 'Sudsko odsustvo', 'SO', true, false, 'Odsustvo zbog sudskog postupka ili saslušanja', '#0000FF', 'icon-court', NOW(), NOW()),
    (4, NULL, 'Službeni put', 'SP', false, true, 'Odsustvo zbog službenog putovanja ili delegacije', '#FFFF00', 'icon-business-trip', NOW(), NOW()),
    (5, NULL, 'Odsustvo zbog obuke', 'OBO', false, false, 'Odsustvo zbog stručne obuke ili edukacije', '#00FFFF', 'icon-training', NOW(), NOW());

INSERT INTO employee_absents (
    id, absent_type_id, user_profile_id, target_organization_unit_id, description, date_of_start, date_of_end, location, file_id, created_at, updated_at)
VALUES
    (1, 1, 1, NULL, 'Odsustvo zbog gripa', '2023-08-01', '2023-08-07', 'Podgorica', NULL, NOW(), NOW()),
    (2, 2, 1, NULL, 'Godišnji odmor', '2023-07-01', '2023-07-15', 'Budva', NULL, NOW(), NOW()),
    
    (3, 3, 2, NULL, 'Sudsko saslušanje', '2023-08-05', '2023-08-05', 'Nikšić', NULL, NOW(), NOW()),
    (4, 1, 2, NULL, 'Odsustvo zbog gripa', '2022-01-01', '2022-04-01', 'Podgorica', NULL, NOW(), NOW()),
    
    (5, 4, 3, 3, 'Službeni put u inostranstvo', '2023-06-01', '2023-06-10', 'Budva', NULL, NOW(), NOW()),
    
    (6, 1, 4, NULL, 'Odsustvo zbog bolesti', '2023-05-10', '2023-05-20', 'Herceg Novi', NULL, NOW(), NOW()),
    (7, 2, 4, NULL, 'Godišnji odmor', '2023-09-01', '2023-09-15', 'Herceg Novi', NULL, NOW(), NOW()),
    
    (8, 4, 5, 3, 'Službeni put', '2023-04-01', '2023-04-05', 'Tivat', NULL, NOW(), NOW()),
    
    (9, 2, 6, NULL, 'Godišnji odmor', '2023-07-01', '2023-07-15', 'Bar', NULL, NOW(), NOW()),
    (10, 5, 6, NULL, 'Obuka za IT stručnjaka', '2023-05-15', '2023-05-20', 'Bar', NULL, NOW(), NOW());
