DROP TRIGGER IF EXISTS activate_new_contract_trigger ON employee_contracts;

CREATE TRIGGER activate_new_contract_trigger
BEFORE INSERT ON employee_contracts
FOR EACH ROW
EXECUTE FUNCTION activate_new_contract();
