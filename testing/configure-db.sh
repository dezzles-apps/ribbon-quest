mysql -h localhost -P 3306 --protocol=tcp -u root -psomepassword -e "DROP SCHEMA IF EXISTS ribbonquesttest;"
mysql -h localhost -P 3306 --protocol=tcp -u root -psomepassword -e "CREATE database ribbonquesttest;"
mysql -h localhost -P 3306 --protocol=tcp -u root -psomepassword ribbonquesttest < ../sql/001-prod-clone.sql
mysql -h localhost -P 3306 --protocol=tcp -u root -psomepassword ribbonquesttest < ./setup.sql