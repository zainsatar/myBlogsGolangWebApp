package requesthandlers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"main/controllers"
	"main/models"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func ShowSignupPage(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	temp, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	temp.ExecuteTemplate(w, "signup.gohtml", msg)
}
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formMap := r.Form
	//fmt.Println(formMap)
	s1 := models.Signup{formMap["firstName"][0], formMap["lastName"][0], formMap["email"][0], formMap["password"][0], ""}
	flag, msg := controllers.SignupController(s1)
	if flag {
		http.Redirect(w, r, "/login?msg="+msg, 302)
	} else {
		http.Redirect(w, r, "/signup?msg="+msg, 302)
	}
}
func ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["id"]
	if ok {
		http.Redirect(w, r, "/home?msg=Please logout first.", 302)
	} else {
		msg := r.URL.Query().Get("msg")
		temp, err := template.ParseGlob("templates/*.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		temp.ExecuteTemplate(w, "login.gohtml", msg)
	}
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formMap := r.Form
	l1 := models.Login{formMap["email"][0], formMap["password"][0]}
	flag, msg, userInfo := controllers.LoginController(l1)
	if flag {
		session, _ := store.Get(r, "session") //getting the session object

		session.Values["firstName"] = userInfo.FirstName
		session.Values["lastName"] = userInfo.LastName
		session.Values["email"] = userInfo.Email
		session.Values["bio"] = userInfo.Bio
		id := userInfo.ID
		session.Values["id"] = id.String()
		err := session.Save(r, w)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/home?msg=Login Successful. Welcome "+msg, 302)
	} else {
		http.Redirect(w, r, "/login?msg="+msg, 302)
	}
}
func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	id, ok := session.Values["id"]
	// fmt.Println(fmt.Sprint(session.Values["picture"]))
	// fmt.Println("")
	// fmt.Println("-----------------------------------------------------------------------------------------------------")
	filename := fmt.Sprint(id)
	if !ok {
		http.Redirect(w, r, "/login?msg=Please login first.", 302)
	} else {
		// msg := r.URL.Query().Get("msg")
		fname := session.Values["firstName"]
		lname := session.Values["lastName"]
		b := session.Values["bio"]
		firstName := fmt.Sprint(fname)
		lastName := fmt.Sprint(lname)
		bio := fmt.Sprint(b)
		temp, err := template.ParseGlob("templates/*.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		posts := controllers.GetAllPosts()
		var postsReverse []primitive.M
		for i := len(posts) - 1; i >= 0; i-- {
			postsReverse = append(postsReverse, posts[i])
		}
		picture := controllers.RetrieveImage(filename)
		homeData := models.Home{picture, firstName, lastName, bio, postsReverse}
		temp.ExecuteTemplate(w, "home.gohtml", homeData)
	}
}
func ShowAddNewPostPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["id"]
	if !ok {
		http.Redirect(w, r, "/login?msg=Please login first.", 302)
	} else {
		temp, err := template.ParseGlob("templates/*.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		temp.ExecuteTemplate(w, "addNewPost.gohtml", nil)
	}
}

func AddNewPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	id, ok := session.Values["id"]
	ID := fmt.Sprintf("%v", id)
	if !ok {
		http.Redirect(w, r, "/login?msg=Please login first.", 302)
	} else {
		fName := session.Values["firstName"]
		lName := session.Values["lastName"]
		firstName := fmt.Sprintf("%v", fName)
		lastName := fmt.Sprintf("%v", lName)
		name := firstName + " " + lastName
		r.ParseForm()
		formMap := r.Form
		post := models.Post{name, ID, formMap["title"][0], formMap["content"][0], time.Now().Format("02 Jan 2006")}
		_, msg := controllers.AddNewPostController(post)
		http.Redirect(w, r, "/home?msg="+msg, 302)

	}

}
func ShowMyPostsPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	idInterface, ok := session.Values["id"]
	id := fmt.Sprintf("%v", idInterface)
	if !ok {
		http.Redirect(w, r, "/login?msg=Please login first.", 302)
	} else {
		posts := controllers.GetMyPosts(id)
		var postsReverse []primitive.M
		for i := len(posts) - 1; i >= 0; i-- {
			postsReverse = append(postsReverse, posts[i])
		}
		temp, err := template.ParseGlob("templates/*.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		temp.ExecuteTemplate(w, "myPosts.gohtml", postsReverse)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["id"]
	if !ok {
		http.Redirect(w, r, "/login", 302)
	} else {
		delete(session.Values, "id")
		delete(session.Values, "email")
		delete(session.Values, "firstNmae")
		delete(session.Values, "lastName")
		delete(session.Values, "bio")
		delete(session.Values, "picture")
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
	}
}
func ShowSettingsPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	idInterface, ok := session.Values["id"]
	id := fmt.Sprint(idInterface)
	if !ok {
		http.Redirect(w, r, "/login?msg=Please login first.", 302)
	} else {
		e := session.Values["email"]
		fName := session.Values["firstName"]
		lName := session.Values["lastName"]
		b := session.Values["bio"]
		firstName := fmt.Sprintf("%v", fName)
		lastName := fmt.Sprintf("%v", lName)
		email := fmt.Sprintf("%v", e)
		bio := fmt.Sprintf("%v", b)
		temp, err := template.ParseGlob("templates/*.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		picture := controllers.RetrieveImage(id)
		profile := models.UserDetails{picture, firstName, lastName, email, "", bio}
		temp.ExecuteTemplate(w, "myAccount.gohtml", profile)
	}
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["id"]
	if !ok {
		http.Redirect(w, r, "/login?msg=Please login first.", 302)
	} else {
		e := session.Values["email"]
		sessionEmail := fmt.Sprintf("%v", e)
		r.ParseForm()
		file, _, err := r.FormFile("profilePic")
		if err == nil {
			buf := bytes.NewBuffer(nil)
			_, err = io.Copy(buf, file)
			if err != nil {
				fmt.Println("could not copy file to buffer")
				log.Fatal(err)
			}
			userid := session.Values["id"]
			filename := fmt.Sprint(userid)
			name := fmt.Sprintf("%v", filename)
			controllers.UploadFile(buf.Bytes(), name)
			delete(session.Values, "picture")
		}
		formMap := r.Form
		cPassword := formMap["currentPassword"][0]
		newPassword := formMap["newPassword"][0]
		var password string
		if newPassword != "" {
			password = newPassword
		} else {
			password = cPassword
		}
		profile := models.Signup{formMap["firstName"][0], formMap["lastName"][0], formMap["email"][0], password, formMap["bio"][0]}
		flag, msg := controllers.UpdateProfileController(sessionEmail, cPassword, profile)
		if flag {
			session.Values["firstName"] = formMap["firstName"][0]
			session.Values["lastName"] = formMap["lastName"][0]
			session.Values["email"] = formMap["email"][0]
			session.Values["bio"] = formMap["bio"][0]
			session.Save(r, w)
			http.Redirect(w, r, "/home", 302)
		} else {
			http.Redirect(w, r, "/home?msg="+msg, 302)
		}

	}
}
