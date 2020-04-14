FROM scratch
ADD certificates/server.crt /server.crt
ADD certificates/server.key /server.key
ADD .env /
ADD server.out /
CMD ["/server.out"]
