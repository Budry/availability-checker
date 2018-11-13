FROM golang:alpine as builder

# Install git
RUN apk --no-cache --update add git ca-certificates
RUN go get -u github.com/golang/dep/cmd/dep

RUN update-ca-certificates

COPY . $GOPATH/src/bitbucket.org/Budry/availability-checker/
WORKDIR $GOPATH/src/bitbucket.org/Budry/availability-checker/

# Install dependencies
RUN dep ensure

#build the binary
RUN go build -o /go/bin/availability-checker
# STEP 2 build a small image


##########################################################################


# start from scratch
FROM alpine

RUN apk --no-cache --update add ca-certificates && update-ca-certificates

# Copy our static executable
COPY --from=builder /go/bin/availability-checker /go/bin/availability-checker

CMD ["/go/bin/availability-checker"]