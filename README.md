# Proyecto Lab 3 - NetSpace

Este Proyecto se basa en una plataforma dedicada a manejar distintos aspectos de la cuenta de MercadoLibre del usuario logueado. Dando herramientas para: Publicar Productos, Responder Preguntas, Ver ventas realizadas, Ver Productos Publicados, Guardar Informacion de Productos en una DB MySQL.
## Integrantes

Oliva Benjamin,
Asmuzi Gali,
Tello Eric.

### Estructura:
* La estructura de la plataforma comienza con el ingreso de usuario, que una vez logueado, pasara a un
dashboard de control e informacion del usuario, su cuenta, productos, ventas, etc.
* En este dashboard se pueden gestionar distintos parametros de la cuenta e informacion, y una vez terminado si lo desea puede desloquearse
con el boton "Salir".
* Representacion de la estructura de la Plataforma

* ![Alt text](pkg/assets/Plataforma%20NetSpace.png)


### Heramientas Utilizadas

* Se utilizaron servicios CDN para usar hojas de estilos y archivos JS de Boostrap 4
* Se utilizo el template `SB-Admin-2` de Boostrap 4 para el diseño UI, modals, cards, botones, etc.(URL al template: https://startbootstrap.com/theme/sb-admin-2)
* Se utilizo XAMPP para servir la DB en `MySQL`
* `GoLand` fue el IDE utilizado para ambiente de trabajdo debido a sus funciones y facilidades con `GitHub`.
* Se uso GitHub para seguimiento del proyecto, trabajo paralelo, y mantener a los developers al tanto de este proyecto.
* Se dio uso de la api de MercadoLibre y sus herramientas para trabajar junto con Go y dar funcionamiento a la plataforma.
  (URL a pagina de la api `MELI`: https://developers.mercadolibre.com.ar/es_ar/api-docs-es)
  
### Lenguajes Utilizados durante el Proyecto

* Para el desarrollo del proyecto se necesito conocimientos en:
 *HTML, CSS, JS, Go, SQL, jQuery* 

 ### Pasos para Ejecucion e ingreso a la Plataforma:
 
 * **1.** Ejecutar el archivo `main.go` para iniciar el servidor local en el puerto `:8080`, e iniciar el servicio de `Apache` y `MySQL` en XAMPP para usar la base de datos (Aclaracion: Se debe tener creada la Database `netspace` y la tabla `items` dentro de la misma para el correcto funcionamiento, ademas de la estructura misma de la tabla) .
 * **2.** Una vez que se este ejecutando el servidor local y XAMPP, iremos a: http://localhost:8080/ingresar .
 * **3.** Iniciaremos sesion en nuestra cuenta y daremos los permisos correspondientes a la aplicacion.
          (si usted ya esta logueado, se lo redijira directamente al dashboard)
 * **4.** Una vez hecha la autenticacion, se lo redirijira al dashboard donde en la parte superior izquierda estara su "nickname"
          de la plataforma, verificando asi que se haya logueado correctamente.

#### Funciones del DASHBOARD:

* **1.** Una vez en el dashboard podemos acceder a los links habilitados en la barra lateral izquierda como se ve en la siguiente imagen.
* ![Alt text](pkg/assets/netspace-platformSidebar.png "Opciones en Sidebar")
* **2.** El apartado de Preguntas no redijira a una seccion donde se mostraran las preguntas realizadas en los items del usuario actualmente logueado
  junto con un boton correspondiente a la pregunta para ser respondida dentro de la plataforma, seguido de un mensaje de confirmacion de envio de respuesta.
* **3.** En el Apartado de Estadisticas, veremos valores correspondientes a los datos guardados por el usuario en la base de datos.
        Esto varia segun el usuario logueado y la informacion guardada por el mismo.
* **4.**  Y por ultimo tenemos el apartado Productos que consta de 3 secciones.
* **4. 1.**  La primera es de publicaciones donde el usuario puede ver los items
publicados por el mismo, junto con un boton de guardado para almacenar en la base de datos.
* **4. 2.** La segunda seccion es para Crear una Publicacion en MercadoLibre desde la plataforma NetSpace. 
            Se despliega un formulario en la pantalla para completar los datos del itema a publicar, y una vez completado se publica mostranso posteriormente un mensaje
            de confirmacion.
* **4. 3.** Por ultimo, esta la seccion de Ventas donde el usuario puede acceder a ver las Ventas recibidas en su cuenta, ademas de los datos correspondientes a cada venta en particular.


 

