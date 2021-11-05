FROM jrei/systemd-centos:7

COPY PermissionDoor.viper.yaml /etc/PermissionDoor/PermissionDoor.viper.yaml
COPY permission-door /usr/local/bin/permission-door

CMD ["permission-door", "run"]
