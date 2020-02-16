1) `docker-compose up -d ` - for up docker conrainer with sqlite3 db.          
          
2) in env you should set full path to you positions.db file           
           
3)           
- From bin            
`bin/sample-server`           
               
or           
           
- From source           
`(cd src/app/ && dep ensure)`           
`go run src/app/main.go`           

4) For checks routes            
http://localhost:3000/summary/apostrophied.co.uk           
http://localhost:3000/positions/apostrophied.co.uk?sortBy=keyword&page=3           
