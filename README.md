# Proyecto Lab 3 - NetSpace

En este proyecto usarmos los datos de la aplicacion NetSpace a la cual le daremos permisos cuando se haga el OAuth.

## Integrantes

Oliva Benjamin,
Gali Asmuzi,
Ery Tello.

### Estructura:
* La estructura de la plataforma comienza con el ingreso de usuario, que una vez logueado, pasara a un
dashboard de control e informacion del usuario, su cuenta, productos, ventas, etc.
* En este dashboard se pueden gestionar distintos parametros de la cuenta e informacion, y una vez terminado si lo desea puede desloquearse
con el boton "Salir".
* [ Descripcion de estructura sujeto a desarrolar ]

 ### Pasos para Ejecucion e ingreso a la Plataforma:
 
 * **1.** Ejecutar el archivo `main.go` para iniciar el servidor local en el puerto `:8080`.
 * **2.** Una vez que se este ejecutando el servidor local, iremos a: http://localhost:8080/ingresar .
 * **3.** Iniciaremos sesion en nuestra cuenta y daremos los permisos correspondientes a la aplicacion.
          (si usted ya esta logueado, se lo redijira directamente al dashboard)
 * **4.** Una vez hecha la autenticacion, se lo redirijira al dashboard donde en la parte superior izquierda estara su "nickname"
          de la plataforma, verificando asi que se haya logueado correctamente.

#### Funciones del DASHBOARD:

* **1.** Una vez en el dashboard podemos acceder a los links habilitados en la barra lateral izquierda como se ve en la siguiente imagen.
![Alt text](pkg/server/templates/sidebar-readme.png "Opciones en Sidebar")
* **2.** Las opciones que no estan habilitadas o desarrolladas, aparecen sin la posibilidad de ser utilizadas hasta que se complete su desarrollo, acompañado de un 
         estilo mas tenue para que el usuario entienda que la opcion no esta disponible.
* **3.** Entre algunas de las opciones tenemos el apartado "Preguntas" donde se mostraran las consultas realizadas por otros usuarios en los productos publicados por el usuario logueado que aun no han sido respondidas (Aun en desarrollo)
* **4.** La otra opcion disponible es la de "Productos", donde se ven todos los items publicados por el usuario logueado y su respectiva informacion.
           ( Aun en desarrollo )         
 

