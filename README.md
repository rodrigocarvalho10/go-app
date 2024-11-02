# Personagens de Filmes e Séries - API - v2

A API teve uma nova versão lançada, incluímos o cadastro agora em um banco de dados sqlite e mais funcionalidades:

- **GET /production**
- **POST /production**
- **GET /productions**
- **DELETE /rmproduction**
- **PUT /updproduction**

A imagem está sendo atualizada, a mesma ainda está na versão anterior
- **docker run -d --name go-app -p 8000:8000 rodrigocarvalho92/go-app**

## Endpoints

- **Lista uma Obra:** `/production`
- **Lista todas as obras:** `/productions`
- **Atualiza um cadastro:** `/updproduction`
- **Remove um cadastro:** `/rmproduction`

### Criar Personagem - v2

- **URL:** `/cadastro`
- **Método:** `POST`
- **Body:**
  ```json
   {
   "Name":"Bjorn Ironside", "Movie": false, "Serie": true
  }
