create table users
(
  id serial primary key,
  login character varying(64) not null
);

create table enemy_island
(
  id serial primary key
);

create table warrior
(
  id serial primary key,
  name character varying(64) not null,
  health_value float not null,
  user_id integer,
  enemy_island_id integer
);

alter table warrior add constraint warrior2users foreign key (user_id) references users(id) on delete set null on update cascade;
alter table warrior add constraint warrior2enemy_island foreign key (enemy_island_id) references enemy_island(id) on delete set null on update cascade;
alter table warrior add constraint warrior_only_one_reference check(not (user_id is not null and enemy_island_id is not null));

create type weapon_type as enum
(
  'chopping',
  'pricking',
  'universal'
);

create table weapon
(
  id serial primary key,
  type weapon_type not null,
  name character varying(64) not null,
  damage_value float not null
);

create table weapon_warrior
(
  weapon_id integer not null,
  warrior_id integer not null
);

alter table weapon_warrior add primary key(weapon_id, warrior_id);
alter table weapon_warrior add constraint weapon_warrior2weapon foreign key (weapon_id) references weapon(id) on delete cascade on update cascade;
alter table weapon_warrior add constraint weapon_warrior2warrior foreign key (warrior_id) references warrior(id) on delete cascade on update cascade;

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

create table armor_warrior
(
  armor_id integer not null,
  warrior_id integer not null
);

alter table armor_warrior add primary key(armor_id, warrior_id);
alter table armor_warrior add constraint armor_warrior2armor foreign key (armor_id) references armor(id) on delete cascade on update cascade;
alter table armor_warrior add constraint armor_warrior2warrior foreign key (warrior_id) references warrior(id) on delete cascade on update cascade;