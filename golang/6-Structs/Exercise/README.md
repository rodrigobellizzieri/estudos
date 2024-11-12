# Exercicio de Structs em Go

Este exercício é para iniciantes que estão aprendendo sobre `structs` em Go. Ele consiste na criação de um programa que simula um cadastro simples de produtos.

## Objetivo

Praticar a criação e uso de `structs`, além de funções que recebem e modificam `structs` em Go.

## Descrição do Exercício

1. Crie um programa em Go que simule o cadastro de um produto.

2. Defina uma `struct` chamada `Produto` com os seguintes campos:
   - `Nome` (string): Nome do produto.
   - `Categoria` (string): Categoria do produto.
   - `Preco` (float64): Preço do produto.
   - `EmEstoque` (bool): Se o produto está disponível em estoque.

3. No `main`, faça o seguinte:
   - Crie uma instância da struct `Produto` para um produto qualquer (ex.: "Notebook", "Eletrônicos", 3500.00, true).
   - Mostre os detalhes do produto no console.

4. Adicione uma função chamada `exibirProduto` que receba uma variável do tipo `Produto` e exiba as informações do produto formatadas no console.

5. Adicione uma segunda função chamada `desconto` que:
   - Receba como parâmetros uma variável `Produto` e um valor de desconto em percentual (ex.: 10 para 10%).
   - Aplique o desconto ao campo `Preco` do produto e retorne o novo preço com desconto.

6. No `main`, use a função `desconto` para aplicar um desconto ao produto e exiba o novo preço no console.

## Exemplo de Saída Esperada

```plaintext
Produto: Notebook
Categoria: Eletrônicos
Preço Original: 3500.00
Em Estoque: Sim

Preço com desconto de 10%: 3150.00
