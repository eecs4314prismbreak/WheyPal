create table "userInterests"
(
	userid serial not null
		constraint userinterests_pk
			primary key
		constraint userinterests_profile_userid_fk
			references profile
				on update cascade on delete cascade,
	interestid int not null
		constraint userinterests_interests_interestid_fk
			references interests (interestid)
);

