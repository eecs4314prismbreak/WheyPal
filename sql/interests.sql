create table interests
(
	interestid int not null,
	interest varchar(32)
);

create unique index interests_interestid_uindex
	on interests (interestid);

alter table interests
	add constraint interests_pk
		primary key (interestid);