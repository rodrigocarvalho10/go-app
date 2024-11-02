# Personagens de Filmes e Séries - API - v2

A API teve uma nova versão lançada, incluímos o cadastro agora em um banco de dados sqlite e mais funcionalidades:

- **GET /production**
- **POST /production**
- **GET /productions**
- **DELETE /rmproduction**
- **PUT /updproduction**

# main.go

Publiquei o arquivo apesar de ter feito o build direto na imagem para que possam avaliar o código e colaborarem.

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
