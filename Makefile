docker:
	docker build -t allenkaplan/ak-rest-api .

adduser:
	curl -X POST localhost:8080/user --data '{"Name":"Stephan"}' -H "Content-Type:application/json"