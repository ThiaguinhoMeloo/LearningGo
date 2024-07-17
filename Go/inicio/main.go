package main

import (
	// "inicio/exampleimports"
	"bufio"
	"fmt"
	"inicio/mongoconnect"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite qual ação deseja realizar: ")
	fmt.Println("1 - Insert (Inserir um cadastro)")
	fmt.Println("2 - Update (Atualizar um cadastro)")

	typeStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler a entrada Type:", err)
		return
	}
	typeStr = strings.TrimSpace(typeStr)
	typeData, err := strconv.Atoi(typeStr)
	if err != nil || (typeData != 1 && typeData != 2) {
		fmt.Println("Erro ao ler a entrada de ação a ser realizada: Esperado 1 ou 2. O número %s informado não existe na lista de ações acima.\n", typeStr)
		return
	}

	if typeData == 1 {

		mongoconnect.InsertMongoDb()

	} else if typeData == 2 {

		fmt.Println("Digite o seu nome de cadastro para prosseguir: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler a entrada Nome:", err)
			return
		}
		name = strings.TrimSpace(name)

		fmt.Println("Escolha que tipo de alteração deseja realizar: 1 para Nome e 2 para E-mail")
		typeStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler a entrada de ação a ser realizada: Esperado 1 ou 2", err)
			return
		}

		typeStr = strings.TrimSpace(typeStr)
		typeKey, err := strconv.Atoi(typeStr)
		if err != nil || (typeKey != 1 && typeKey != 2) {
			fmt.Println("Erro ao ler a entrada de ação a ser realizada: Esperado 1 ou 2")
			return
		}

		if typeKey == 1 {
			fmt.Println("Digite o seu novo nome: ")
			alter, err := reader.ReadString('\n')
			alter = strings.TrimSpace(alter)
			if err != nil {
				fmt.Println("Erro ao ler a entrada Nome: ", err)
				return
			}
			mongoconnect.UpdateData(name, alter, typeKey)
		} else if typeKey == 2 {
			fmt.Println("Digite o seu novo nome: ")
			alter, err := reader.ReadString('\n')
			alter = strings.TrimSpace(alter)
			if err != nil {
				fmt.Println("Erro ao ler a entrada E-mail: ", err)
				return
			}
			mongoconnect.UpdateData(name, alter, typeKey)
		} else {
			fmt.Println("Tipo de dado invalido.")
			return
		}
	}
}
