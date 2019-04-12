FROM scratch
ADD http/ /http
ADD main /
CMD ["/main"]
