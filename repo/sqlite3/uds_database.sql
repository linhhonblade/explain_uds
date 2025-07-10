-- Creating the services table to store UDS Service IDs and their descriptions
CREATE TABLE services (
    sid TEXT PRIMARY KEY, -- Service ID in hex (e.g., "0x22")
    name TEXT NOT NULL,   -- Service name (e.g., "ReadDataByIdentifier")
    description TEXT,      -- Detailed description of the SID
    has_subfunction BOOLEAN NOT NULL DEFAULT 0,
    positive_response INTEGER

);

-- Creating the sub_functions table to store sub-functions for specific SIDs
CREATE TABLE sub_functions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sid TEXT NOT NULL,    -- Foreign key to services
    value TEXT NOT NULL,  -- Sub-function value in hex (e.g., "0x01")
    name TEXT NOT NULL,   -- Sub-function name (e.g., "Default Session")
    description TEXT,     -- Detailed description
    FOREIGN KEY (sid) REFERENCES services(sid)
);

-- Creating the parameters table to store parameters like DIDs
CREATE TABLE parameters (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sid TEXT NOT NULL,    -- Foreign key to services
    type TEXT NOT NULL,   -- Parameter type (e.g., "DID", "Data")
    value TEXT,           -- Parameter value (e.g., "0xF190")
    description TEXT,     -- Detailed description
    FOREIGN KEY (sid) REFERENCES services(sid)
);

-- Inserting sample data into services table
INSERT INTO services (sid, name, description, has_subfunction) VALUES
('0x10', 'DiagnosticSessionControl', 'Initiates a diagnostic session with the ECU.', 1),
('0x22', 'ReadDataByIdentifier', 'Requests data from the ECU by specifying a Data Identifier (DID).', 0),
('0x2E', 'WriteDataByIdentifier', 'Writes data to the ECU using a Data Identifier (DID).', 0),
('0x3E', 'TesterPresent', 'Indicates that the tester is present and active.', 0),
('0x31', 'RoutineControl', 'Controls the execution of a routine in the ECU.', 1),
('0x3C', 'ReadMemoryByAddress', 'Reads memory from the ECU by specifying an address.', 0),
('0x2F', 'InputOutputControlByIdentifier', 'Controls input/output operations by specifying an identifier.', 1);

-- Inserting sample data into sub_functions table
INSERT INTO sub_functions (sid, value, name, description) VALUES
('0x10', '0x01', 'Default Session', 'Starts the default diagnostic session.'),
('0x10', '0x02', 'Programming Session', 'Enables programming mode for flashing or updates.'),
('0x10', '0x03', 'Extended Diagnostic Session', 'Enables extended diagnostic functions.');

-- Inserting sample data into parameters table
INSERT INTO parameters (sid, type, value, description) VALUES
('0x22', 'DID', '0xF190', 'Vehicle Identification Number (VIN): Retrieves the VIN from the ECU.'),
('0x22', 'DID', '0xF18C', 'ECU Serial Number: Retrieves the serial number of the ECU.'),
('0x2E', 'DID', '0xF190', 'Vehicle Identification Number (VIN): Writes a new VIN to the ECU.'),
('0x2E', 'Data', NULL, 'Data payload for writing to the specified DID (e.g., VIN string).');
