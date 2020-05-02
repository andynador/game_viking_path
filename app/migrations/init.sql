create table users
(
  id serial primary key,
  login character varying(64) not null
);

create table enemy_island
(
  id serial primary key,
  name character varying(64) not null
);

alter table enemy_island add constraint enemy_island_name_unique unique (name);

create type weapon_type as enum
(
  'chopping',
  'pricking',
  'universal'
);

create table weapon
(
  id serial primary key,
  name character varying(64) not null,
  type weapon_type not null,
  damage_value float not null
);

alter table weapon add constraint weapon_name_unique unique (name);

create type armor_type as enum
(
  'chopping',
  'pricking',
  'universal'
);

create table armor
(
  id serial primary key,
  type armor_type not null,
  protection_value float not null
);

alter table armor add constraint armor_type_protection_value_unique unique (type, protection_value);

create table warrior
(
  id serial primary key,
  name character varying(64) not null,
  health_value float not null,
  user_id integer,
  enemy_island_id integer,
  weapon_id integer not null,
  armor_id integer not null
);

alter table warrior add constraint warrior2users foreign key (user_id) references users(id) on delete set null on update cascade;
alter table warrior add constraint warrior2enemy_island foreign key (enemy_island_id) references enemy_island(id) on delete set null on update cascade;
alter table warrior add constraint warrior2weapon foreign key (weapon_id) references weapon(id) on delete restrict on update cascade;
alter table warrior add constraint warrior2armor foreign key (armor_id) references armor(id) on delete restrict on update cascade;
alter table warrior add constraint warrior_only_one_reference check(not (user_id is not null and enemy_island_id is not null));
alter table warrior add constraint warrior_name_unique unique (name);