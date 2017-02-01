# sftp-client
golang sftp client, connect to multiple servers sftp and download a specific file


La idea de este proyecto es crear un cliente SFTP para que pueda conectarse a un servidor específico y descargar un archivo específico.

Adicionalmente en este cliente SFTP pueden agregarse una lista de configuraciones:
       - Servidor sftp, usuario/clave, path, archivo a descargar, path local, email-confirmación
       
Con esto este cliente SFTP podría conectarse a todos los servidores configurados y ejecutar la acción.

Parte de la funcionalidad es enviar por parámetro si solamente se conectará a uno, un grupo o todos los servidores.

Un plug-in adicional es que si falla en alguna conexión, se envíe un email a la cuenta registrada, para esto debería de poder configurarse los parámetros de un servidor SMTP


## Editado desde el ambiente de desarrollo ##
.
