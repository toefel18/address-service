-- Example data that needs to be stored:
--
-- aobjectid  "0003010000126006"
-- nraandid   "0003200000134086"
-- oruimteid  "0003300000117006"
-- straat     Burgemeester Lewe van Aduardstraat
-- huisnummer 52
-- huisletter "g"
-- huisnrtoev "5"
-- postcode   "9902NN"
-- wnpcode    "3386"
-- woonplaats "Appingedam"
-- gemcode    "0003"
-- gemeente   "Appingedam"
-- provcode   "20"
-- provincie  "Groningen"
-- buurtcode  "00030002"
-- buurtnr    30002
-- straatnen  "Burg Lewe van Aduardstr"
-- aotype     "v"
-- status     "Verblijfsobject in gebruik"
-- oppvlakte  10
-- gebrksdoel "overige gebruiksfunctie"
-- x_rd       253463.07
-- y_rd       592651.11
-- lat        53.311231855738
-- long       6.86450306650747

-- test instance with docker:
-- docker run --name address-service-go-postgres -p 5432:5432 -e POSTGRES_DB=addressservice -e POSTGRES_PASSWORD=root -d postgres
-- sudo docker run --name address-service-go-postgres -p 5432:5432 -v /opt/address-service-pgdata:/var/lib/postgresql/data:rw -e POSTGRES_PASSWORD=root -d postgres

CREATE TABLE addressesnetherlands (
  aobjectid VARCHAR(16) PRIMARY KEY NOT NULL,
  kixcode VARCHAR(24),
  nraandid VARCHAR(16),
  oruimteid VARCHAR(16),
  straat VARCHAR(64),
  huisnummer INTEGER,
  huisletter VARCHAR(6),
  huisnrtoev VARCHAR(12),
  postcode VARCHAR(6),
  wnpcode VARCHAR(4),
  woonplaats VARCHAR(64),
  gemcode VARCHAR(4),
  gemeente VARCHAR(64),
  provcode VARCHAR(2),
  provincie VARCHAR(128),
  buurtcode VARCHAR(8),
  buurtnr INTEGER,
  straatnen VARCHAR(64),
  aotype VARCHAR(4),
  status VARCHAR(64),
  oppvlakte INTEGER,
  gebrksdoel VARCHAR(64),
  x_rd DOUBLE PRECISION,
  y_rd DOUBLE PRECISION,
  lat DOUBLE PRECISION,
  long DOUBLE PRECISION
);

CREATE INDEX index_addr_kixcode ON AddressesNetherlands (kixcode);
CREATE INDEX index_addr_postcode ON AddressesNetherlands (postcode);
CREATE INDEX index_addr_huisnummer ON AddressesNetherlands (huisnummer);

ALTER TABLE addressesnetherlands ADD CONSTRAINT kixcode_unique UNIQUE (kixcode);
