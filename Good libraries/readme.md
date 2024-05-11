# Packages' usages

This is a dictionary for golang packages and their 
usages . 

| Package  | usage                       | link |
|----------|-----------------------------|------|
| httptest | making mock http requests   | todo |
| mockgen  | creating mock of interfaces | todo |
| goqu     | generates queries           | https://github.com/doug-martin/goqu |
| Hugo | static web pages  | | 
| Viper | It is used for config giles | |
| cobra | command line | | 
| Task | Automations | https://taskfile.dev/ |
| Carbon | Managing datetimes | https://github.com/golang-module/carbon |
| Go-micro | for building microservices  | https://github.com/go-micro/go-micro
| Colly | web scraping | https://github.com/gocolly/colly | 
| semconv | | |
| Viper | Parsing config files | github.com/spf13/viper |
| i18n | Internationalization, But I only know this | |
| go breaker | implentation of circuit breaker pattern in go | https://github.com/sony/gobreaker |
| Throttled | Rate limiting | https://github.com/throttled/throttled |
| go-cron | Job scheduling package | https://github.com/go-co-op/gocron |
| pyroscope | Profiling | https://github.com/grafana/pyroscope |

### Built-in packages
| Package      | usage                                                                        |
|--------------|------------------------------------------------------------------------------|
| os/exec      | executing bash commands                                                      |
| os           | creating files                                                               |
| ioutil       | read from file and write to them , <br/> read directores and files' names    |
| regexp       | find matchings of a regex in a text ,<br/> check a texts matching to a regex |
| encoding/csv | reading and opening csv files                                                | 
|text/template | creating texts, accepting an object, the object's attributes are shown in text |
| crypto/x509 | Package x509 implements a subset of the X.509 standard. [link](https://pkg.go.dev/crypto/x509) |
| hash | Don't know what, example: using crc32 to compute `ChecksumIEEE` |
| runtime | I think this package helps to manage goroutines more efficiently and in a lower level. It has some methods to get information about goroutines and other functionalities such as terminating them, etc. [documentation](https://pkg.go.dev/runtime#GOMAXPROCS) |
