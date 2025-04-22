#  Histórico de Estudo: Aprendizado de Go com Projeto de Xadrez

Este documento acompanha o progresso diário de aprendizado da linguagem Go através do desenvolvimento de um simulador de tabuleiro de xadrez.

---

##  Dia 1 – Introdução à linguagem e estrutura de projeto

- Compreensão da estrutura básica de um programa Go:
  - `package main`, `import`, `func main()`
- Criação da matriz `[8][8]string` representando o tabuleiro
- Impressão simples do tabuleiro com `fmt.Print`, `fmt.Println`, `fmt.Printf`

---

##  Dia 2 – Iteração com `for` e `range`, formatação visual

- Uso de laços `for` e `range` para percorrer o tabuleiro
- Aplicação de `8 - i` para inverter visualmente a impressão
- Compreensão dos placeholders `%d` e `%s` no `Printf`

---

##  Dia 3 – Conversão de coordenadas de xadrez

- Transformar strings como `"d2"` em índices de matriz
- Uso de `rune`, conversão ASCII com `'d' - 'a'`
- Criação das funções auxiliares `letraParaIndice` e `numeroParaIndice`

---

##  Dia 4 – Movimento de peças

- Criação da função `moverPeca(origem, destino)`
- Atualização da matriz após o movimento
- Prevenção de movimentação a partir de casas vazias
- Adição de loop contínuo com `fmt.Scan` para entrada de jogadas

---

##  Dia 5 – Validação de entrada

- Implementação da função `entradaValida(pos string)`
- Verificação do formato da coordenada (`a2`, `h8`, etc.)
- Prevenção de erros causados por strings fora do padrão (ex: `7a`, `z9`)

---

##  Dia 6 – Regras básicas de movimentação

- Uso de `switch` para identificar a peça
- Implementação da lógica de movimentação para:
  - **Peão branco (♙)** – anda 1 ou 2 casas para frente
  - **Peão preto (♟)** – anda 1 ou 2 casas para trás
  - **Cavalo (♘, ♞)** – movimento em L
- Criação da função auxiliar `abs(int)` para validar distância

---

##  Próximos passos

- Adicionar regras de movimento para bispo, torre, rainha e rei
- Implementar lógica de captura
- Verificar turno (vez de quem jogar)
- Suporte a `check`, `checkmate`, e `empate`

---

Esse histórico mostra não só a evolução do projeto, mas também como a linguagem Go pode ser aprendida de forma prática e visual.

