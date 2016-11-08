FROM scratch

ADD ./aircnc /aircnc
EXPOSE 8080
EXPOSE 4000
WORKDIR /

CMD ["./aircnc"]
