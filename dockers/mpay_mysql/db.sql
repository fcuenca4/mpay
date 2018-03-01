# In production you would almost certainly limit the replication user must be on the follower (slave) machine,
# to prevent other clients accessing the log from other machines. For example, 'replicator'@'follower.acme.com'.
#
# However, this grant is equivalent to specifying *any* hosts, which makes this easier since the docker host
# is not easily known to the Docker container. But don't do this in production.
#
GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'replicator' IDENTIFIED BY 'replpass';
GRANT SELECT, RELOAD, SHOW DATABASES, REPLICATION SLAVE, REPLICATION CLIENT  ON *.* TO 'debezium' IDENTIFIED BY 'dbz';

# Create the database that we'll use to populate data and watch the effect in the binlog
DROP DATABASE IF EXISTS mpay;
CREATE DATABASE mpay;
GRANT ALL PRIVILEGES ON mpay.* TO 'mysqluser'@'%';

# Switch to this database
USE mpay;

# Create and populate our products using a single insert with many rows
DROP TABLE IF EXISTS payments;
CREATE TABLE payments (
	id BIGINT NOT NULL AUTO_INCREMENT,
	collector BIGINT NULL,
	payer BIGINT NULL,
	creation_date DATE NULL,
	amount BIGINT NULL,
	status varchar(40) DEFAULT 'Pending' NOT NULL,
	status_detail varchar(40) NULL,
	metadata varchar(255) NULL,
	CONSTRAINT payments_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci ;
