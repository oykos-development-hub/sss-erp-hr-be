insert into systematizations(user_profile_id, organization_unit_id, active, date_of_activation, serial_number)
values (1, 2, 2, NOW(), '1234');

insert into job_positions(title, is_judge, is_judge_president)
values('Administrator sistema', false, false);

insert into job_positions_in_organization_units(systematization_id, job_position_id, available_slots)
values(1, 1, 1);

insert into employees_in_organization_units(user_account_id, user_profile_id, position_in_organization_unit_id, active)
values(1,1,1,true);

