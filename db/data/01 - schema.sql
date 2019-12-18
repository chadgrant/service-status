create database if not exists service_status;

use service_status;

create table if not exists environment (
    id int not null auto_increment,
    name varchar(150) character set ascii not null,
    active bit not null default 1,
    created datetime not null default current_timestamp,
    updated datetime null,
    constraint uc_name unique (name, active),
    primary key (id)
);

create table if not exists service (
    id bigint not null auto_increment,
    name varchar(150) character set ascii not null,
    friendly_name varchar(150) character set ascii not null,
    status varchar(150) not null default 'pending',
    active bit not null default 1,
    created datetime not null default current_timestamp,
    updated datetime null,
    constraint uc_name unique (name, active),
    primary key (id)
);

create table if not exists deployment (
    id bigint not null auto_increment,
    environment_id int not null,
    service_id bigint not null,
    build_number varchar(150) character set ascii not null,
    created datetime not null default current_timestamp,
    updated datetime null,
    index idx_environment_service (environment_id, service_id),
    foreign key (environment_id) references environment(id),
    primary key (id)
);

create table if not exists service_endpoint (
    id bigint not null auto_increment,
    environment_id int not null,
    service_id bigint not null,
    endpoint varchar(512) character set ascii not null,
    created datetime not null default current_timestamp,
    primary key (id),
    constraint uc_endpoint unique (environment_id, service_id, endpoint),
    index idx_environment_service (environment_id, service_id),
    foreign key (environment_id) references environment(id),
    foreign key (service_id) references service(id)
);

create table if not exists service_instance (
    id bigint not null auto_increment,
    environment_id int not null,
    service_id bigint not null,
    deployment_id bigint null,
    name varchar(150) character set ascii not null,
    endpoint varchar(512) character set ascii not null,
    status varchar(150) character set ascii not null default 'pending',
    created datetime not null default current_timestamp,
    primary key (id),
    index idx_environment (environment_id),
    constraint uc_service_environment_endpoint unique (environment_id, service_id, endpoint),
    constraint uc_service_environment_name unique (environment_id, service_id, name),
    foreign key (environment_id) references environment(id),
    foreign key (service_id) references service(id)
);

create table if not exists service_environment (
    id bigint not null auto_increment,
    environment_id int not null,
    service_id bigint not null,
    deployment_id bigint not null,
    created datetime not null default current_timestamp,
    primary key (id),
    index idx_environment (environment_id),
    constraint uc_service_environment unique (environment_id, service_id),
    foreign key (deployment_id) references deployment(id),
    foreign key (environment_id) references environment(id),
    foreign key (service_id) references service(id)
);