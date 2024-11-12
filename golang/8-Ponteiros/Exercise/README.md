# Exercício: Manipulação de Ponteiros em Go

## Descrição

Neste exercício, você criará uma função simples que usa ponteiros para atualizar o valor de uma variável. O objetivo é entender como os ponteiros funcionam em Go, permitindo que uma função altere diretamente o valor de uma variável fora de seu escopo.

## Objetivo

Aprender a:
- Criar uma variável e obter seu ponteiro.
- Passar ponteiros para funções para modificar o valor original de uma variável.

## Estrutura do Exercício

1. Crie uma função `incrementar(valor *int)`.
   - Essa função deve receber um ponteiro para uma variável `int`.
   - Dentro da função, incremente o valor apontado por `valor` em 1.
   
2. No `main()`:
   - Crie uma variável `numero` do tipo `int` e defina um valor inicial (por exemplo, `10`).
   - Exiba o valor inicial de `numero`.
   - Chame a função `incrementar` passando o ponteiro de `numero`.
   - Exiba o valor de `numero` novamente para confirmar que ele foi incrementado.

## Exemplo de Saída Esperada

```plaintext
Valor inicial: 10
Valor após incremento: 11
