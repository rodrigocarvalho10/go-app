# Personagens de Filmes e Séries - API

Esta é uma API escrita em Go para cadastrar personagens de filmes e séries, no exemplo cadastramos os personagens de Vikings e informamos se é um Filme ou Série.
A API está no ínicio e será utilizada como base para os meus estudos na linguagem.

# main.go

Publiquei o arquivo apesar de ter feito o build direto na imagem para que possam avaliar o código e colaborarem.

## Endpoints

- **Principal:** `/`
- **Cadastro:** `/cadastro`

### Criar Personagem

- **URL:** `/cadastro`
- **Método:** `POST`
- **Body:**
  ```json
   {
   "Name":"Bjorn Ironside", "Movie": false, "Serie": true
  }

# Arquivos YAML

Na pasta k8s temos dois arquivos yaml do k8s que faz o deploy da aplicação que foi publicada no docker hub e também cria o service para expor a porta da aplicação.
Uma observação é a sessão abaixo do service.yaml:

 type: LoadBalancer

  externalIPs:
  - 192.168.15.9 

O externalIPs eu apontei para o IP do meu cluster local, pois sem a chave o kubernetes ficou como pending o processo de atribuição do EXTERNAL IP, dessa forma fiquem atento ao de vocês caso forem subir no kubernetes igual o meu exemplo.


