-- files
create table files (
    `id` char(32) NOT NULL PRIMARY KEY,
    `key` char(32) NOT NULL,
    created_at datetime not null default CURRENT_TIMESTAMP,
    expiration_time datetime null,
    account_id bigint unsigned null,
    `name` varchar(64) not null,
    `type` varchar(32) not null,
    `extension` varchar(6) not null,
    `path` varchar(256) not null,
    `meta` json null,
    index files_created_at_index (created_at),
    index files_expiration_time_index (expiration_time),
    index files_account_id_index (account_id),
    index files_type_index (type)
);