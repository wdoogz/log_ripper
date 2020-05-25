## log_ripper
Program that scrapes logs for failures/errors/aborts etc.

[![Total alerts](https://img.shields.io/lgtm/alerts/g/wdoogz/log_ripper.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/wdoogz/log_ripper/alerts/) ![CI](https://github.com/wdoogz/log_ripper/workflows/CI/badge.svg)


---

#### How to use

1) Change script permissions to 755 and move to `/bin`
2) Create a crontab that runs every 1-5 mins and executes the following, change to match your influxdb values.
    >  `sudo crontab -e`

    >  `*/1 * * * * /bin/grafana_log_ripper_l /var/log/messages 192.168.1.10:8086 influxdb_dbname influxdb_username influxdb_password`
