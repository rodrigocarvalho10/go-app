# Vikings API

Esta é uma API REST escrita em Go para cadastrar personagens de filmes e séries, no exemplo cadastramos os personagens de Vikings e informamos se é um Filme ou Série.

## Endpoints

### Criar Personagem

- **URL:** `/personagens`
- **Método:** `POST`
- **Body:**
  ```json
   {
   "Name":"Bjorn Ironside", "Movie": false, "Serie": true
  }
