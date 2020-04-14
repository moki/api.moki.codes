FROM scratch
ADD certificates/server.crt /server.crt
ADD certificates/server.key /server.key
ADD server.out /
CMD ["/server.out"]
