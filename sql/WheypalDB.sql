--drop table profile;
--drop table notification;
--drop table userInterest;
--drop table blockList;
--drop table matchRequest;
--drop table interest;
--drop table location;
--drop table users;

CREATE TABLE users (userId SERIAL PRIMARY KEY, username VARCHAR(32) NOT NULL,
	password VARCHAR(32) NOT NULL, email VARCHAR(50) NOT NULL);
CREATE TABLE interest (interestId SERIAL PRIMARY KEY, interestName VARCHAR(32) NOT NULL);
CREATE TABLE location (locationId SERIAL PRIMARY KEY, city VARCHAR(100) NOT NULL, state VARCHAR(100),
	country VARCHAR(75) NOT NULL);
CREATE TABLE matchRequest (matchRequestId SERIAL PRIMARY KEY, status VARCHAR(20) NOT NULL,
	userA SERIAL REFERENCES users(userId) ON DELETE CASCADE, userB SERIAL REFERENCES users(userId) ON DELETE CASCADE,
	CHECK (userA <> userB), CHECK (status = 'accepted' OR status = 'declined' OR status = 'pendingUserA' OR status = 'pendingUserB'));
CREATE TABLE profile (userId SERIAL REFERENCES users(userId) ON DELETE CASCADE,locationId SERIAL REFERENCES location(locationId) ON DELETE CASCADE, gender CHAR(1) NOT NULL, availability BOOLEAN NOT NULL, 
	birthDate DATE NOT NULL, firstName VARCHAR(50), lastName VARCHAR(50), profileDescription VARCHAR(10000),
	PRIMARY KEY (userId), CHECK (gender = 'M' OR gender = 'F' OR gender = 'X'));
CREATE TABLE notification (userId SERIAL REFERENCES users(userId) ON DELETE CASCADE, matchRequest BOOLEAN NOT NULL,
	mutualMatch BOOLEAN NOT NULL, unreadMessage BOOLEAN NOT NULL, PRIMARY KEY (userId));
CREATE TABLE userInterest (userId SERIAL REFERENCES users(userId) ON DELETE CASCADE,
	interestId SERIAL REFERENCES interest(interestId) ON DELETE CASCADE, PRIMARY KEY (userId, interestId));
CREATE TABLE blockList (blocker SERIAL REFERENCES users(userId) ON DELETE CASCADE,
	blockedBy SERIAL REFERENCES users(userId) ON DELETE CASCADE, PRIMARY KEY (blocker, blockedBy));
