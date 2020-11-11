docker:
	docker build -t eecs4314prismbreak/wheypal .

adduser:
	curl -v -H "Content-Type:application/json" -X POST -d '{"Name": "kevin", "Email":"ke@v.in", "Password":"k3v1n", "Birthday":"Kevin 1st, 1998", "Location":"21 Kevin Street", "Interest":"Kevinterests"}' localhost:8081/user
