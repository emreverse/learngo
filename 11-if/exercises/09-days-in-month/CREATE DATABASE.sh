CREATE DATABASE trtemplate ENCODING='UTF-8' LC_COLLATE = 'tr_TR.UTF-8' LC_CTYPE='tr_TR.UTF-8' TEMPLATE template0;

CREATE DATABASE emre_test TEMPLATE trtemplate;

CREATE ROLE emre_test NOLOGIN;

CREATE ROLE emre_test_owner password 'emre_test.2021!' IN ROLE emre_test;

CREATE ROLE emre_test_app NOLOGIN IN ROLE emre_test;

CREATE ROLE emre_test_readonly NOLOGIN IN ROLE emre_test;

CREATE ROLE emre_test_dwh NOLOGIN IN ROLE emre_test;

CREATE ROLE emre_test_developer NOLOGIN IN ROLE emre_test;

CREATE ROLE emre_test_report NOLOGIN IN ROLE emre_test;

\c emre_test;

GRANT CONNECT, TEMP ON DATABASE emre_test TO emre_test;

GRANT USAGE ON SCHEMA public TO emre_test;

ALTER DEFAULT PRIVILEGES FOR ROLE emre_test_owner GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO emre_test_app, emre_test_developer;

ALTER DEFAULT PRIVILEGES FOR ROLE emre_test_owner GRANT SELECT, UPDATE ON SEQUENCES TO emre_test_app, emre_test_developer;

ALTER DEFAULT PRIVILEGES FOR ROLE emre_test_owner GRANT EXECUTE ON FUNCTIONS TO emre_test_app, emre_test_developer;

ALTER DEFAULT PRIVILEGES FOR ROLE emre_test_owner GRANT USAGE ON TYPES TO emre_test_app, emre_test_developer;

ALTER DEFAULT PRIVILEGES FOR ROLE emre_test_owner GRANT USAGE ON SCHEMAS TO emre_test_app, emre_test_developer;

ALTER DEFAULT PRIVILEGES FOR ROLE emre_test_owner GRANT SELECT ON TABLES TO emre_test_readonly,emre_test_dwh,emre_test_report;

ALTER DEFAULT PRIVILEGES FOR ROLE emre_test_owner GRANT USAGE ON TYPES TO emre_test_readonly,emre_test_dwh,emre_test_report;

ALTER DEFAULT PRIVILEGES FOR ROLE emre_test_owner GRANT USAGE ON SCHEMAS TO emre_test_readonly,emre_test_dwh,emre_test_report;

ALTER DEFAULT PRIVILEGES FOR ROLE emre_test_owner GRANT SELECT ON SEQUENCES TO emre_test_readonly,emre_test_dwh,emre_test_report;

CREATE ROLE emre_test_appuser LOGIN IN ROLE emre_test_app;

ALTER ROLE emre_test_appuser PASSWORD '8HBXCE(>&Sn6B4eZ';