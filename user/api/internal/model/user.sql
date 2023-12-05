CREATE TABLE user (
                      id INT AUTO_INCREMENT ,
                      name VARCHAR(255) NOT NULL default '',
                      password VARCHAR(255) NOT NULL default '',
                      gender VARCHAR(10) NOT NULL default '0',
                      PRIMARY KEY(id)
) ENGINE=InnoDB;