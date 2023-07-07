# vilnaCMS
Prototype of open-source CMS on Go

### For local development
We're running  application and postgresql db in docker containers, for persistent db records docker volume is used.
1. Install docker and docker compose.
2. Run `docker-compose build` && `docker-compose up -d`. 
3. Navigate to the browser and open http://0.0.0.0:8080.
4. Read containers logs here: http://0.0.0.0:9999
- NOTES: 
1. Default user credentials are `admin:admin`.
2. For persisting db data please DO NOT shut down docker compose with `-v` key.
