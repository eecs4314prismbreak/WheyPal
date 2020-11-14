build:
	docker build -t eecs4314prismbreak/wheypal .

pull:
	sudo docker pull eecs4314prismbreak/wheypal

push:
	sudo docker push eecs4314prismbreak/wheypal

run:
	sudo docker run  --rm -d -p 8080:8080 -e PORT='8080' \
		--name wheypal-backend eecs4314prismbreak/wheypal

kill:
	sudo docker kill wheypal-backend
