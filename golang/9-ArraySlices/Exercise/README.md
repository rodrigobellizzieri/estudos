# Exercício: Manipulação de Arrays e Slices em Go

## Descrição

Neste exercício, você irá trabalhar com **arrays** e **slices** em Go. O objetivo é praticar a criação, modificação e iteração sobre arrays e slices. Você também aprenderá a usar a função `append` para adicionar elementos a um slice e a acessar elementos de ambos.

## Objetivo

O objetivo deste exercício é:

- Criar e inicializar um **array**.
- Criar e manipular um **slice**.
- Usar a função `append` para adicionar novos elementos ao slice.
- Iterar sobre os elementos de um array e um slice.

## Estrutura do Exercício

1. **Parte 1: Array**
   - Crie um **array** de inteiros com 5 elementos e inicialize com valores (por exemplo: `10, 20, 30, 40, 50`).
   - Exiba o valor de todos os elementos do array.

2. **Parte 2: Slice**
   - Crie um **slice** de inteiros inicialmente vazio.
   - Use a função `append` para adicionar os números do array (da Parte 1) ao slice.
   - Exiba o valor do slice após adicionar os elementos.

3. **Parte 3: Manipulação**
   - Após adicionar os elementos ao slice, adicione mais dois valores ao slice (por exemplo, `60` e `70`).
   - Exiba o valor do slice novamente após as modificações.

4. **Parte 4: Iteração**
   - Use um laço `for` para iterar sobre o array e o slice, exibindo cada elemento.

## Exemplo de Saída Esperada

```plaintext
Valores do array:
10 20 30 40 50

Valores do slice após adicionar os elementos do array:
[10 20 30 40 50]

Valores do slice após adicionar 60 e 70:
[10 20 30 40 50 60 70]

Iterando sobre o array:
10
20
30
40
50

Iterando sobre o slice:
10
20
30
40
50
60
70
