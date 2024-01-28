# GoLang
Estudos da linguagem GOLang

# Go Kelvin to Celsius Converter

## Descrição do Projeto

Este projeto consiste em um programa em Go para a conversão de escalas termométricas de Kelvin para Celsius. O código-fonte realiza a conversão usando a fórmula apropriada e fornece uma solução simples para esse tipo de conversão e informa qual a temperatura de ebulição da água em Kelvin e qual o valor em Celsius.

## Funcionalidades

- **Conversão de Kelvin para Celsius:** O programa utiliza a fórmula de conversão adequada para calcular e exibir a temperatura em Celsius a partir da entrada em Kelvin.

## Código Fonte

```go
package main

import "fmt"

// declaração da variável CONST do ponto de ebulição da água em Kelvin
const ebulicaoK = 373.15

func main() {
    tempKK := 32.0
    tempK := ebulicaoK
    tempC := tempK - tempKK // Convertendo K para C

    fmt.Printf("A temperatura de ebulição da água em K° = %v , a temperatura de ebulição em C° a partir de %v K  é = %v graus celsius ", tempK, tempKK, tempC)
}
```

## Pré-requisitos

- Go (Instalação: [https://golang.org/doc/install](https://golang.org/doc/install))

## Como Usar

1. Clone este repositório:
   ```bash
   git clone https://github.com/seu-usuario/nome-do-repositorio.git
   ```

2. Navegue até o diretório do projeto:
   ```bash
   cd nome-do-repositorio
   ```

3. Execute o programa, fornecendo a temperatura em Kelvin como argumento:
   ```bash
   go run main.go
   ```
   O programa usará a temperatura padrão definida no código. Você também pode modificar `tempKK` para fornecer uma temperatura diferente em Kelvin.

4. O programa exibirá a temperatura equivalente em Celsius:
   ```bash
   A temperatura de ebulição da água em K° = 373.15 , a temperatura de ebulição em C° a partir de 32 K  é = 341.15 graus celsius
   ```

## Contribuição

Contribuições são bem-vindas! Se você encontrar bugs, melhorias ou quiser adicionar novas funcionalidades, sinta-se à vontade para criar uma "issue" ou enviar um "pull request".

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).
