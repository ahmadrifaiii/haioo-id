###########################################################################
# Stage 1 Start
###########################################################################
FROM golang AS build-golang


RUN export GO111MODULE=on
RUN export GOPROXY=direct
RUN export GOSUMDB=off

################################
# Build Service:
################################
WORKDIR /usr/share/telkomIndonesia/the-shopping-cart

COPY  . .

RUN make deploy

###########################################################################
# Stage 2 Start
###########################################################################
FROM ubuntu:20.04

# Change Repository ke kambing.ui:
RUN sed -i 's*archive.ubuntu.com*kambing.ui.ac.id*g' /etc/apt/sources.list

RUN apt-get update

RUN apt-get install -y ca-certificates

# Copy Binary
COPY --from=build-golang /usr/share/telkomIndonesia/the-shopping-cart/bin /usr/share/telkomIndonesia/the-shopping-cart/bin/

WORKDIR /usr/share/telkomIndonesia/the-shopping-cart

# Create group and user to the group
RUN groupadd -r the-shopping-cart && useradd -r -s /bin/false -g the-shopping-cart the-shopping-cart

# Set ownership golang directory
RUN chown -R the-shopping-cart:the-shopping-cart /usr/share/telkomIndonesia/the-shopping-cart

# Make docker container rootless
USER the-shopping-cart

# EXPOSE 8080

# ENTRYPOINT [ "./service" ]
# CMD [ "./rest" ]