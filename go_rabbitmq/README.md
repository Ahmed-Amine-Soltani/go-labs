go to get amqp module
``` bash
go get github.com/streadway/amqp
```

docker rabbitmq container:

``` bash
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```
