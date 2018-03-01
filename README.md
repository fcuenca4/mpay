# mpay

Como crear las instancias:

<dentro de dockers>
docker-compose -f docker-compose-mysql.yaml up

Como detener las instancias

docker-compose -f docker-compose-mysql.yaml up

Como eliminar las instancias

docker-compose -f docker-compose-mysql.yaml stop
docker-compose -f docker-compose-mysql.yaml rm

Luego, crear un nuevo connector y un topic:

curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d '{ "name": "mpay-connector", "config": { "connector.class": "io.debezium.connector.mysql.MySqlConnector", "tasks.max": "1", "database.hostname": "mysql", "database.port": "3306", "database.user": "debezium", "database.password": "dbz", "database.server.id": "184054", "database.server.name": "dbserver1", "database.whitelist": "mpay", "database.history.kafka.bootstrap.servers": "kafka:9092", "database.history.kafka.topic": "dbhistory.mpay" } }'

y finalmente montar el consumidor 

./consumer_example localhost:9092 topic dbserver1.mpay.payments



Para montar el servidor en GO:
Dependencias: github.com/gin-gonic/gin/, github.com/jinzhu/gorm, github.com/pilu/fresh, gopkg.in/go-playground/validator.v9

En el main folder: 

fresh

Curls:

Creación de un pago:

curl -i -X POST -H "Content-Type:application/json" localhost:8080/payments -d '{ "collector":1,"amount":33,"payer":3}'

Aprobación/Cancelación 
curl -i -X PUT -H "Content-Type:application/json" localhost:8080/payments/:paymentID -d '{"status":"Approved"}'
curl -i -X PUT -H "Content-Type:application/json" localhost:8080/payments/:paymentID -d '{"status":"Cancelled"}'

Recuperar un pago

curl localhost:8080/payments/:paymentID

