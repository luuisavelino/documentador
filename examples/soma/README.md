# Documentação de Código

## Requerimentos

* [Golang](https://golang.org/)

## Instrução de uso

* Para executar o programa, você pode compilá-lo e executá-lo manualmente ou usar o comando `go run`:

```bash
$ go run main.go
A soma dos numeros 10 + 5 é: 15
```

* Também é possível compilar o código para obter um executável:

```bash
$ go build main.go
$ ./main
A soma dos numeros 10 + 5 é: 15
```

## Descrição do Código

O programa começa importando a biblioteca padrão `fmt`, que fornece funções para formatação e saída de texto. Ele então define a função `main`, que será executada quando o programa iniciar.

A variável `a` é inicializada com o valor 10 e `b` com 5. O valor de `c` é então atribuído à soma de `a` e `b`. Por fim, o programa usa a função `Printf` da biblioteca `fmt` para imprimir os valores de `a`, `b` e `c` com alguma formatação.
