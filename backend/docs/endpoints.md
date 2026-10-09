# Documentação da API - Jardim

## 1. Autenticação

### `POST /auth/login`

* **Request:**
  ```json
  {
    "email": "usuario@email.com",
    "password": "senha123"
  }
  ```
* **Response (200 OK):**
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

## 2. Cliente Web (`/user`)
**Requer Header:** `Authorization: Bearer <JWT>`

### `GET /user/devices`
* **Response (200 OK):**
  ```json
  [
    { "id": 1, "uuid": "123e4567-e89b-12d3-a456-426614174000", "nome": "Jardim Varanda" }
  ]
  ```

### `POST /user/devices`
* **Request:**
  ```json
  {
    "uuid": "123e4567-e89b-12d3-a456-426614174000",
    "pin": "847291"
  }
  ```
* **Response (201 Created):**
  ```json
  { "status": "vinculado", "id": 1 }
  ```

### `GET /user/devices/{id}`
* **Response (200 OK):**
  ```json
  { "id": 1, "nome": "Jardim Varanda", "status": "online" }
  ```

### `PUT /user/devices/{id}`
* **Request:**
  ```json
  { "nome": "Jardim Fundos" }
  ```

### `DELETE /user/devices/{id}`
* **Request:** (Sem payload)

### `GET /user/devices/{id}/telemetry`
* **Response (200 OK):**
  ```json
  [
    {
      "timestamp": "2026-10-08T10:00:00Z",
      "temperatura_celsius": 25.5,
      "umidade_ar_percentual": 60.0,
      "umidade_solo_percentual": 45.0
    }
  ]
  ```

### `GET /user/devices/{id}/config`
* **Response (200 OK):**
  ```json
  {
    "rega_automatica": true,
    "limiar_umidade_solo_percentual": 30.0,
    "duracao_rega_segundos": 120
  }
  ```

### `PUT /user/devices/{id}/config`
* **Request:**
  ```json
  {
    "rega_automatica": false,
    "limiar_umidade_solo_percentual": 40.0,
    "duracao_rega_segundos": 60
  }
  ```

## 3. Hardware (`/hw`)
**Requer Header:** `X-Device-Token: <UUID>`

### `POST /hw/telemetry`
* **Request:**
  ```json
  {
    "temperatura_celsius": 26.2,
    "umidade_ar_percentual": 55.0,
    "umidade_solo_percentual": 38.5
  }
  ```
* **Response (201 Created):** (Sem payload)

### `GET /hw/config`
* **Response (200 OK):**
  ```json
  {
    "rega_automatica": false,
    "limiar_umidade_solo_percentual": 40.0,
    "duracao_rega_segundos": 60
  }
  ```

## 4. Admin (`/admin`)
**Requer Header:** `Authorization: Bearer <Admin_JWT>`

### `GET /admin/devices`
* **Response (200 OK):**
  ```json
  [
    { "id": 1, "uuid": "123e...", "mac": "00:1B:44:11:3A:B7", "user_id": 15 }
  ]
  ```

### `POST /admin/devices`
* **Request:**
  ```json
  {
    "mac": "00:1B:44:11:3A:B7",
    "pin": "847291"
  }
  ```
* **Response (201 Created):**
  ```json
  { "uuid": "123e4567-e89b-12d3-a456-426614174000" }
  ```

### `GET /admin/devices/{id}`
* **Response (200 OK):**
  ```json
  {
    "id": 1,
    "mac": "00:1B:44:11:3A:B7",
    "last_ip": "192.168.1.50",
    "status": "ativo"
  }
  ```

### `PUT /admin/devices/{id}`
* **Request:**
  ```json
  { "status": "suspenso" }
  ```

### `DELETE /admin/devices/{id}`
* **Request:** (Sem payload)

### `GET /admin/users`
* **Response (200 OK):**
  ```json
  [
    { "id": 15, "email": "usuario@email.com", "total_devices": 2 }
  ]
  ```

### `DELETE /admin/users/{id}`
* **Request:** (Sem payload)

### `DELETE /admin/telemetry?older_than=30d`
* **Request:** (Sem payload)
* **Response (200 OK):**
  ```json
  { "registros_apagados": 15420 }
  ```