#Proyecto Lab 3 - NetSpace

En este proyecto usarmos los datos de la aplicacion NetSpace a la cual le daremos permisos cuando se haga el OAuth.

##Integrantes

Oliva Benjamin,
Gali Asmuzi,
Ery Tello.

####Link OAuth: 

https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=4026193299288302&redirect_uri=http://localhost:8080/auth/code/

Aclaracion: Esta URL es utilizada para generar token en una cuenta. Usando el metodo OAuth se obtienen estos datos necesarios para luego usarse para obtener los datos de Items del usuario, Preguntas realizadas en sus Items, etc.

###Link Utilizado para obtener Datos del Usuario:

http://localhost:8080/dashboard

Aclaracion: Esta URL nos mostrara en el navegador los datos obtenidos por el codigo (obtenerDatos.go) en formato JSON la informacion.

#####- Items Obtenidos
#####- Preguntas Sin responder
#####- Ventas reaizadas ( items con Order.status = " paid ")

###Informacion Detallada:

Dentro de cada archivo de codigo se entra mas en detalle el funcionamiento en si del Sistema y el manejo de los datos en el mismo. 
 