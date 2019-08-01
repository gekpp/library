POSTGRES_TAG=11
POSTGRES_PASSWORD=library
POSTGRES_USER=librarian
POSTGRES_DB=library

.PHONY: postgres clean-postgres

compile: gen example build

gen:
	cd pkg && goa gen library/design

example:
	cd pkg && goa example library/design

build:
	cd pkg && \
	  go build -o ../bin/books ./cmd/books && \
	  go build -o ../bin/books-cli ./cmd/books-cli
	
run:
	bin/books -debug
		# ./books -dbHost 127.0.0.1 -dbUser ${POSTGRES_USER} -dbPassword ${POSTGRES_PASSWORD} -dbName ${POSTGRES_DB}


postgres: if-postgres-not-running
	docker run --name library-postgres \
		-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
		-e POSTGRES_USER=${POSTGRES_USER} \
		-e POSTGRES_DB=${POSTGRES_DB} \
		-v `pwd`/migrations:/docker-entrypoint-initdb.d \
		-p 55432:5432 \
		-d postgres:${POSTGRES_TAG}

if-postgres-not-running:
	! (docker ps -a --format "{{.Names}}" | grep 'library-postgres') 

clean-postgres:
	docker rm -f library-postgres

connect-postgres:
	docker exec -ti library-postgres psql -d library -U librarian

clean-gen:
	rm -Rf ./pkg/gen

# clean-all:
# 	rm -Rf bin/books bin/books-cli
# 	cd pkg && \
# 	rm -Rf cmd gen auther.go books.go bin/books bin/books-cli
