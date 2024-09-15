# Event Management API

## Descripción

La API de Event Management permite la gestión de eventos en una plataforma web. Los usuarios pueden visualizar e inscribirse en eventos publicados, mientras que los administradores tienen la capacidad de gestionar (alta, baja, modificación) los eventos. La API está desarrollada en Golang y sigue los principios de clean architecture.

## Funcionalidades

### Gestión de Eventos

- **Creación, actualización y eliminación de eventos**: Solo los usuarios con rol de administrador pueden gestionar eventos.
- **Campos de evento**: 
  - Título
  - Descripción corta
  - Descripción larga
  - Fecha y hora
  - Organizador
  - Lugar
  - Estado (borrador o publicada)
- **Visibilidad**:
  - Los eventos en estado de borrador solo son visibles para los administradores.
  - Los eventos publicados pueden ser visualizados e inscritos por todos los usuarios, siempre que la fecha del evento sea futura.
  - Los eventos pasados pueden ser visualizados pero no inscritos.

## Endpoints

### 1. Listar Eventos

- **Endpoint**: `GET /events`
- **Descripción**: Obtiene una lista de eventos, filtrados por estado, fecha y título.
- **Parámetros de consulta**:
  - `status`: Estado del evento (publicado o borrador).
  - `date_from`: Fecha de inicio para filtrar eventos.
  - `date_to`: Fecha de fin para filtrar eventos.
  - `title`: Título del evento.
- **Respuesta**:
  - **200 OK**: Lista de eventos.
  - **Ejemplo**:
    ```json
    [
      {
        "id": 1,
        "title": "Evento Ejemplo",
        "description_short": "Descripción corta",
        "description_long": "Descripción larga",
        "date_time": "2024-09-21T15:00:00Z",
        "organizer": "Organizador",
        "location": "Ubicación",
        "status": "published"
      }
    ]
    ```

### 2. Obtener Eventos Inscritos

- **Endpoint**: `GET /user/events`
- **Descripción**: Obtiene una lista de eventos a los que el usuario está inscrito, filtrados por activos o completados.
- **Parámetros de consulta**:
  - `user_id`: ID del usuario.
  - `status`: Estado de los eventos (activos o completados).
- **Respuesta**:
  - **200 OK**: Lista de eventos inscritos.
  - **Ejemplo**:
    ```json
    [
      {
        "id": 1,
        "title": "Evento Inscrito",
        "description_short": "Descripción corta",
        "description_long": "Descripción larga",
        "date_time": "2024-09-21T15:00:00Z",
        "organizer": "Organizador",
        "location": "Ubicación",
        "status": "active"
      }
    ]
### 3. Crear Evento

- **Endpoint**: `POST /events`
- **Descripción**: Crea un nuevo evento.
- **Requiere Autenticación**: Sí, debe ser un administrador.
- **Cuerpo de la solicitud**:
  ```json
  {
    "title": "Evento Ejemplo",
    "description_short": "Descripción corta",
    "description_long": "Descripción larga",
    "date_time": "2024-09-21T15:00:00Z",
    "organizer": "Organizador",
    "location": "Ubicación",
    "status": "draft"
  }
- **Respuesta**:
  - **201 OK**:  Evento creado.

### 4. Actualizar Evento
- **Endpoint**: `PUT /events/{id}`
- **Descripción**:  Actualiza un evento existente.
- **Requiere Autenticación**: Sí, debe ser un administrador.
- **Cuerpo de la solicitud**:
    ```json
    {
  "title": "Evento Actualizado",
  "description_short": "Descripción corta actualizada",
  "description_long": "Descripción larga actualizada",
  "date_time": "2024-09-21T15:00:00Z",
  "organizer": "Organizador Actualizado",
  "location": "Ubicación Actualizada",
  "status": "published"
    }

- **Respuesta**:
  - **200 OK**:  Evento actualizado.

### 5. Eliminar Evento
- **Endpoint**: `DELETE /events/{id}`
- **Descripción**:   Elimina un evento existente..
- **Requiere Autenticación**: Sí, debe ser un administrador.
- **Respuesta**:
  - **201 OK**:  Evento eliminado.

  ## Middleware

### Autenticación

- **Middleware**: `AuthMiddleware`
- **Descripción**: Verifica si el usuario está autenticado. Usa un token simulado en el encabezado `Authorization` con el prefijo `Bearer`.

### Autorización

- **Middleware**: `AdminMiddleware`
- **Descripción**: Verifica si el usuario tiene el rol de administrador. Requiere el encabezado `Roles` que debe contener `admin`.

## Conexión a MySQL

### Configuración de la Base de Datos

La aplicación se conecta a una base de datos MySQL utilizando el paquete `gorm` para ORM.

### Configuración en el Código

La conexión a la base de datos está configurada en `infrastructure/database.go`:

Acceder a la API
La API estará disponible en http://localhost:8080