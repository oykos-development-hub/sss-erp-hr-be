INSERT INTO absent_types (parent_id, title, abbreviation, relocation, accounting_days_off, description, color, icon, created_at, updated_at) VALUES
(NULL, 'Godišnji odmor', 'GO', 'f', 't', 'Zakonski godišnji odmor', '#00FF00', 'icon-vacation', NOW(), NOW()),
(NULL, 'Bolovanje', '123', 'f', 'f', 'Odsustvo zbog bolovanja.', NULL, NULL, NOW(), NOW()),
(NULL, 'Roditeljstvo', '123', 'f', 'f', 'Odsustvo zbog roditeljstva.', NULL, NULL, NOW(), NOW()),
(NULL, 'Upućivanje u drugi državni organ', '123', 't', 'f', 'Odsustvo zbog upućivanja u drugi državni organ.', NULL, NULL, NOW(), NOW()),
(NULL, 'Neplaćeno odsustvo', '123', 'f', 'f', 'Neplaćeno odsustvo po zahtjevu radnika.', NULL, NULL,NOW(), NOW());
