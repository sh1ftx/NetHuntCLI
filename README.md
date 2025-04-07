# NetHunterCLI

DevBook CLI é uma aplicação de linha de comando (CLI) desenvolvida em Go, que permite buscar informações sobre IPs e servidores na internet. Este projeto foi criado para ser simples, eficiente e útil para desenvolvedores e administradores de sistemas.

## 📋 Funcionalidades

- **Busca de IPs**: Retorna os endereços IP associados a um domínio.
- **Busca de servidores**: Retorna os servidores de nomes (DNS) associados a um domínio.

## 🚀 Tecnologias Utilizadas

- [Go](https://golang.org/) - Linguagem de programação principal.
- [urfave/cli](https://github.com/urfave/cli) - Biblioteca para criação de aplicações CLI.

## 📂 Estrutura do Projeto

```plaintext
NetHunterCLI/
├── Project/
│   └── App/
│       └── app.go        # Lógica principal da aplicação CLI
├── [main.go](http://_vscodecontentref_/0)               # Ponto de entrada da aplicação
├── [go.mod](http://_vscodecontentref_/1)                # Gerenciamento de dependências
├── license               # Licença do projeto
└── .gitignore            # Arquivos ignorados pelo Git
