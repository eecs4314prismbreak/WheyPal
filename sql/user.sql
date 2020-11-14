create table "user"
(
	userid serial not null
		constraint user_logins_userid_fk
			references logins
				on update cascade on delete cascade,
	"firstName" varchar(32) not null,
	"lastName" varchar(32) not null,
	birthdate date not null,
	location varchar(64) not null,
    interest varchar(32) not null
);

create unique index user_userid_uindex
	on "user" (userid);

alter table "user"
	add constraint user_pk
		primary key (userid);

