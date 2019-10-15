package main

import (
	"log"
	"net/http"
	"os"
    "fmt"
	oidc "github.com/coreos/go-oidc"  // paquete para la implementacion del protocolo OpenID Connect
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

    // Obtenemos del contexto de nuestra consola los valores de las variables indicadas
    // Estos valores deberan ser establecidos previamente a la ejecucion del ejemplo y obtenidos al registrar nuestra API en Google-Cloud
var (
	clientID     = os.Getenv("GOOGLE_OAUTH2_CLIENT_ID")
	clientSecret = os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET")
    URL = "127.0.0.1:1736"

	ctx = context.Background()
     
    googleprovider, err = oidc.NewProvider(ctx, "https://accounts.google.com")
	
	// Creamos la estructura para la peticion OAuth2
	googleconfig = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     googleprovider.Endpoint(),
		RedirectURL:  "http://"+URL+"/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

 	state = "foobar"
)

// Manejador de las peticiones a '/'
func handleMain(w http.ResponseWriter, r *http.Request) {
		var htmlBody = `<html>
		                    <head>
							<title>Api OpenID Connect-OAuth2</title>
							</head>
							<body>
							    Para el acceso deberemos hacer uso del recurso '/links' o seguir este enlace 
								<a href="/links">Ver los enlaces delegados</a>
							</body>
						</html>`

	    fmt.Fprintf(w, htmlBody)
		log.Printf(" Enviado al Front la informacion.")

	
	}

// Manejador de las peticiones a '/links'	
func handleLinks(w http.ResponseWriter, r *http.Request) {
		var htmlBody = `<html>
							<head>
								<title>Api OpenID Connect-OAuth2</title>
							</head>
							<body>
							<ul>
								<li><a href="/google/login">Acceso vía Google</a></li>
							</ul>	
							</body>
						</html>`

	    fmt.Fprintf(w, htmlBody)
		log.Printf(" Enviado al Front los Links a los accesos delegados.")

	
	}

// Manejador de las peticiones a '/google/login'	
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Establecemos como proveedor de acceso Google
			http.Redirect(w, r, googleconfig.AuthCodeURL(state), http.StatusFound)
	}



// Manejador de las llamadas a '/callback'
func handleCallback(w http.ResponseWriter, r *http.Request) {

		var htmlHead = `    <html>
		                    <head>
							<title>Api OpenID Connect-OAuth2</title>
							</head>							
						`
		
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state no encontrado", http.StatusBadRequest)
			return
		}

		oauth2Token, err := googleconfig.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Fallo al intercambiar el Token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Token obtenido")
		
		userInfo, err := googleprovider.UserInfo(ctx, oauth2.StaticTokenSource(oauth2Token))
		
		if err != nil {
			http.Error(w, "Fallo al recuperar los datos del usuario: "+err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Datos de usuario obtenidos del Servidor de Recursos")
	
		htmlBody := `  <body> ` + `Datos del usuario: <ul>` +`<li>` + userInfo.Subject + `</li>` +`<li>` + userInfo.Email + `</li>` + ` </ul></body> </html>`
		fmt.Fprintf(w, htmlHead + htmlBody)
		log.Printf("Datos de usuario enviados al front ")
	}		
				
// Función principal
func main() {

	// Asignamos los manejadores
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/links", handleLinks)
	http.HandleFunc("/google/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleCallback)

	log.Printf("Servidor activo en http://%s/", URL)
	log.Fatal(http.ListenAndServe(URL, nil))
}


