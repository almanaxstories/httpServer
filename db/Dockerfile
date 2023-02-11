FROM postgres:15.1
RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates
RUN apt-get install make -y 
RUN make migrateup 