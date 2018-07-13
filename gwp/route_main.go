package main

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
										"templates/nabar.html",
										"templates/index.html",
					}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}