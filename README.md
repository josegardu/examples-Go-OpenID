# examples-Go-OpenID
Example of Oauth2.0 implementation with Go language.
## Go OAuth2 ejemplo con OpenID Connect

---------- Español:

Ejemplo de utilización de OAuth2 con API.

Para preparar el contexto:

 Para que goolge nos permita solicitar datos de usuarios deberemos registrar nuestra API y obtener un id de cliente para ella.
  con este dato y el id secret deberemos crear asignaciones a las variables de entorno del Sistema.
  
 Si para probar el ejemplo, no quereis tener que registraros en Google podéis emplear estos datos: 

 GOOGLE_OAUTH2_CLIENT_ID = 696706423016-f62l771ikt4soovips03r28sfouff10c.apps.googleusercontent.com
 GOOGLE_OAUTH2_CLIENT_SECRET = YuMeWw_iq2dM6Z4cs57wO7_X

NOTAS: Hay muchas formas para establecer los valores a las variables de entorno, y en muchas ocasiones depende del SO que tengáis, a modo de ayuda dejo este enlace por si os es necesario https://www.java.com/es/download/help/path.xml
       También os dejo un enlace al procedimiento para registrar la API en el servicio OAuth2 de google. https://docs.moodle.org/all/es/Servicio_OAuth_2_Google

Para iniciar el servidor:

1. Inicia el servidor desde la linea de comandos : `go run oauth2.go`
2. Con un Navegador a http://127.0.0.1:1736


------------- English:

Example of using OAuth2 with API.

To prepare the context:

In order for goolge to allow us to request user data, we must register our API and obtain a client id for it. With this data and the secret id we must create assignments to the System environment variables.

If to try the example, you do not want to have to register with Google you can use this data:

GOOGLE_OAUTH2_CLIENT_ID = 696706423016-f62l771ikt4soovips03r28sfouff10c.apps.googleusercontent.com 
GOOGLE_OAUTH2_CLIENT_SECRET = YuMeWw_iq2dM6Z4cs57wO7_X

NOTES: There are many ways to set the values ​​to the environment variables, and in many cases it depends on the OS you have, as a help I leave this link in case you need it https://www.java.com/download /help/path.xml I also leave a link to the procedure to register the API in the Google OAuth2 service. https://docs.moodle.org/all/es/OAuth_2_Google_Service

To start the server:

Start the server from the command line: go run oauth2.go
With a Browser at http://127.0.0.1:1736
