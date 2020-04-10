FROM scratch
ADD www.google.com.crt /etc/ssl/certs/
ADD server.out /
CMD ["/server.out"]
