# Basic-CRUD
basic CRUD functionality that returns one of the jsons via the "quests" route, and returns the string "hello world" via the "/" route. Before use, you need to initialize the database through the resetdb.sql file. 
```
psql postgres -U username
CREATE DATABASE dbname
psql -U username -d dbname -f resetdb.sql
``
And configure the config.env file
