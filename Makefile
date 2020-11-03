docker:
	docker build -t eec4314prismbreak/wheypal .

adduser:
	curl -X POST localhost:8080/user --data '{"Name":"Stephan"}' -H "Content-Type:application/json"