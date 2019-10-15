# examples-Go-OpenID
Example of Oauth2.0 implementation with Go language.
## Go OAuth2 ejemplo con OpenID Connect


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

