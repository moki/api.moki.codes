FROM scratch
ADD certificates /certificates
ADD .env /
ADD server.out /
CMD ["/server.out"]
