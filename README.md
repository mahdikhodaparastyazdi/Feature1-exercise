we want to write an app to manage order process
design a restful api that expose 1 endpoint
POST api/order input : {“order_id”:10,”price”:1000,”title”:”burger”}
in the receiving point send the received data to redis (queue)
create another app to read data from the above redis and process the order
the process should goes like this :
use mysql or any other databases to save the retrieved data into orders table
