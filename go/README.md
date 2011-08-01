In order to make a simple authenticated checklist. We can basically store the checklists in different directories for each user and then, retrieve only from directories that the user has access to.


Deployment
=========

DEPS
----

* autotools
* go
* https://github.com/Philio/GoMySQL


Setup the database
-------------------

TODO
====

* Good testing framework (http://golang.org/doc/code.html)
* Better Makefiles (You can do make install that will download your packages? WHAT.)
* 

-------------------

Database

// Connect to database  
db, err := mysql.DialUnix(mysql.DEFAULT_SOCKET, "user", "password", "database")  
if err != nil {  
    os.Exit(1)  
}  
// Perform query  
err = db.Query("select * from my_table")  
if err != nil {  
    os.Exit(1)  
}  
// Get result set  
result, err := db.UseResult()  
if err != nil {  
    os.Exit(1)  
}  
// Get each row from the result and perform some processing  
for {  
    row := result.FetchRow()  
    if row == nil {  
        break  
    }  
    // ADD SOME ROW PROCESSING HERE  
}  


Use database just for the users table and the cookies table.

Session checking only hits the session table.

User + password happens on every logon.