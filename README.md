# Proyecto Lab 3 - NetSpace

En este proyecto usarmos los datos de la aplicacion NetSpace a la cual le daremos permisos cuando se haga el OAuth.

## Integrantes

Oliva Benjamin,
Gali Asmuzi,
Ery Tello.

<a name="top"></a>
#### Link OAuth: 

https://auth.mercadolibre.com.ar/authorization?response_type=code&client_id=4026193299288302&redirect_uri=http://localhost:8080/auth/code/

Aclaracion: Esta URL es utilizada para generar token en una cuenta. Usando el metodo OAuth se obtienen estos datos necesarios para luego usarse para obtener los datos de Items del usuario, Preguntas realizadas en sus Items, etc.

### Link Utilizado para obtener Datos del Usuario:

http://localhost:8080/dashboard

***Aclaracion:** Esta URL nos mostrara en el navegador los datos obtenidos por el codigo (obtenerDatos.go) en formato JSON la informacion.*

##### - Items Obtenidos
##### - Preguntas Sin responder
##### - Ventas reaizadas ( items con Order.status = " paid ")

### Informacion Detallada:

Dentro de cada archivo de codigo se entra mas en detalle el funcionamiento en si del Sistema y el manejo de los datos en el mismo. 
 
 ### Pasos para Ejecucion:
 
 * **1.** Ejecutar el archivo `main.go` para iniciar el servidor local en el puerto `:8080`.
 * **2.** Una vez que se este ejecutando el servidor local, abrimos el [Link OAuth](#top).
 * **3.** Una vez que dimos permisos a la aplicacion NetSpace, y ya tengamos nuestro JSON en pantalla mostrando los datos, podremos usar los endpoints.
 * **4.** Para usar el endpoint para ver todos los datos del usuario ir a: http://localhost:8080/dashboard.
 * **5.** Para usar el endpoint para ver solo items del usuario ir a http://localhost:8080/items.
 
 

