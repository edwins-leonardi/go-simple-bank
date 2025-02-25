alter table if exists "accounts" drop constraint if exists "owner_currency_key";

alter table if exists "accounts" drop constraint if exists "accounts_owner_fkey";

alter table if exists "accounts" drop constraint if exists "unique_owner_currency";

drop table if exists users;