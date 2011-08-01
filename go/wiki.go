package main

import (
/*    "fmt"
*/    "http"
    "io/ioutil"
    "os"
    "template"
    "github.com/Philio/GoMySQL"
)

type Page struct {
    Title string
    Body []byte
}

const lenPath = len("/view/")

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    
    t, _ := template.ParseFile("view.html", nil)
    t.Execute(w, p)
    
    db, err := mysql.DialUnix(mysql.DEFAULT_SOCKET, "user", "password", "database")  
    if err != nil {  
        os.Exit(1)  
    }
    
    if db != nil {
        os.Exit(1)
    } 
    
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    
}

func (p *Page) save() os.Error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, os.Error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func main() {
/*  p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))*/
    
    /* Once you register, you should be able to create new checklists that have a title and 
        
        Log you in as the demo user the first time you sign up.
        
        You should be able to view all your checklists on your homepage */
    
    /* How do you want to store the user name and password combination? */
    
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/register", registerHandler)
    http.HandleFunc("/login", loginHandler)
    http.ListenAndServe(":8080", nil)
}