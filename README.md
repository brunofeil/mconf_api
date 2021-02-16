# api
- Build
  - docker build -t api .

- Para rodar
  - docker run --publish 8080:8080 api

# runner
- Build
  - docker build -t runner .

- Para rodar
  - docker run runner <"frase a ser pesquisada entre aspas duplas">

# Algumas considerações
- A api roda direto na porta 8080
