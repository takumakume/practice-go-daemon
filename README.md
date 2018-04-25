###### Usage

* clone

```sh
https://github.com/takumakume/practice-go-daemon
cd practice-go-daemon
```

* install systemd service

```
./takumakumed install
```

* run

```sh
# systemctl start takumakumed
# systemctl status takumakumed
● takumakumed.service - daemon of takumakume
   Loaded: loaded (/etc/systemd/system/takumakumed.service; enabled; vendor preset: disabled)
   Active: active (running) since Wed 2018-04-25 05:39:14 UTC; 3s ago
  Process: 8450 ExecStartPre=/bin/rm -f /var/run/takumakumed.pid (code=exited, status=0/SUCCESS)
 Main PID: 8452 (takumakumed)
   CGroup: /system.slice/takumakumed.service
           └─8452 /vagrant/go/src/github.com/takumakume/practice-go-daemon/takumakumed

Apr 25 05:39:14 agent systemd[1]: Starting daemon of takumakume...
Apr 25 05:39:14 agent systemd[1]: Started daemon of takumakume.
Apr 25 05:39:14 agent takumakumed[8452]: hoge
Apr 25 05:39:17 agent takumakumed[8452]: hoge
```
