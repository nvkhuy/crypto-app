create table if not exists `symbols` (
  `id` int not null auto_increment primary key,
  `name` varchar(20),
  `created_at` timestamp default current_timestamp,
  `updated_at` timestamp default current_timestamp on update current_timestamp
);
create table if not exists `bots` (
  `id` int not null auto_increment primary key,
  `name` varchar(20),
  `avatar_url` varchar(500),
  `token` varchar(500),
  `created_at` timestamp default current_timestamp,
  `updated_at` timestamp default current_timestamp on update current_timestamp
);
