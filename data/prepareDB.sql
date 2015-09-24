create table "Users" ("Name" varchar(100) primary key, "Password" varchar(100), "Admin" integer, "CreationDate" timestamp, "Data" blob);
create table "Sessions" ("Name" varchar(100), "SessionKey" varchar(100), "LastRequest" date, primary key ("Name", "SessionKey"));
commit;
insert into "Users" ("Name", "Admin", "CreationDate") values ('admin', 1, current_timestamp);
