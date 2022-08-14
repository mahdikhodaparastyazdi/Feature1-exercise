# restpapiwithredis
create a rest api and save /api/order with pub redis and another application use it with sub redis and store it to postgre sql
the project in mahdi's folder is a application that defines a point entry and store data in redis with pub on a specific channel
and the project in redissub's folder is another application that subcribes in that soecific channel and pop data into postgre sql.
