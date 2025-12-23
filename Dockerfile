FROM registry.access.redhat.com/ubi9/ubi-micro:latest

WORKDIR /app

# Copy the binary you just built
COPY beego-postgresql .

# Copy config (Required for Beego)
COPY conf/ ./conf/

EXPOSE 8080

CMD ["./beego-postgresql"]
