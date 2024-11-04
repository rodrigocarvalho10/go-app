# Filmes e Séries - API - v2

A API teve uma nova versão lançada, incluímos o cadastro agora em um banco de dados sqlite e mais funcionalidades:

## Lista de Endpoints - v2

- **Lista uma Obra:** `/production`
- **Inclui uma nova obra:** `/production`
- **Lista todas as obras:** `/productions`
- **Atualiza um cadastro:** `/updproduction`
- **Remove um cadastro:** `/rmproduction`

Segue o comando para usar e brincar com a aplicação
- **docker run -d --name go-app -p 8000:8000 rodrigocarvalho92/go-app-v2**

## Cadastrar uma série ou filme - v2

- **URL:** `/production`
- **Método:** `POST`
- **Body:**
  ```json
   {
   "Name":"Vikings", 
   "Producer":"MGM Television", 
   "Movie": false, 
   "Series": true,
   "Protagonist": "Ragnar Lothbrok",
   "Notice": 10,
   "Assessment": "Serie excelente, o Ragnar foi monstro"
   }
  
## Buscar um filme ou série

- **URL:** `/production`
- **Método:** `GET`
- **Params:**
  ```
  key: id
  value: 1

  ```

## Atualizar uma série ou filme - v2

- **URL:** `/updproduction`
- **Método:** `POST`
- **Body:**
- **Parameters:**
  ```
  key: id
  value: 1
  
  ```json
   {
   "Name":"Vikings", 
   "Producer":"MGM Television", 
   "Movie": false, 
   "Series": true,
   "Protagonist": "Ragnar Lothbrok",
   "Notice": 9, //Atualizamos a nota
   "Assessment": "Serie excelente, o Ragnar foi monstro"
   }

## Deletar uma série ou filme - v2

- **URL:** `/rmproduction`
- **Método:** `POST`
- **Body:**
- **Parameters:**
  ```
  key: id
  value: 1
  ```