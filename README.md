# Personagens de Filmes e Séries - API - v2

A API teve uma nova versão lançada, incluímos o cadastro agora em um banco de dados sqlite e mais funcionalidades:

# Lista de Endpoints - v2

- **Lista uma Obra:** `/production`
- **Inclui uma nova obra:** `/production`
- **Lista todas as obras:** `/productions`
- **Atualiza um cadastro:** `/updproduction`
- **Remove um cadastro:** `/rmproduction`

A imagem está sendo atualizada, a mesma ainda está na versão anterior
- **docker run -d --name go-app -p 8000:8000 rodrigocarvalho92/go-app**

### Criar Personagem - v2

- **URL:** `/cadastro`
- **Método:** `POST`
- **Body:**
  ```json
   {
   "Name":"Bjorn Ironside", 
   "Producer":"MGM Television", 
   "Movie": false, 
   "Serie": true,
   "Protagonist": "Ragnar Lothbrok",
   "Notice": 10,
   "Assessment": "Serie excelente, o Ragnar foi o principal"
   }