create table orders(
    id int unsigned auto_increment,
    userID int(10) not null default 0,
    createTime int(11) not null default 0,
    name varchar(255) not null default '',
    total FLOAT,
    primary key(id)
)engine = InnoDB;