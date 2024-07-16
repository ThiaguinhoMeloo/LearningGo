package exampleimports

import (
	"fmt"
	"time"
)

func Time() {
	fmt.Println("Exemplos da função Time: ")

	// Data Atual
	now := time.Now()
	fmt.Println("Data e Hora Atuais: ", now)

	//Formatação
	formattedTime := now.Format("02-01-2006 15:04:05")
	fmt.Println("Data e Hora formatadas: ", formattedTime)

	//Parse
	timeStr := "2024-07-16 12:34:56"
	parseTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("Erro ao analisar a data: ", err)
		return
	} else {
		fmt.Println("Data e Hora analisada: ", parseTime)
	}

	//Subtração
	time1 := time.Date(2024, time.July, 16, 12, 0, 0, 0, time.UTC)
	time2 := time.Date(2024, time.July, 15, 12, 0, 0, 0, time.UTC)

	diff := time1.Sub(time2)
	fmt.Printf("A diferença entre os tempos é de %v horas.\n", diff.Hours())

	//UTC e TIMEZONE

	nowUTC := time.Now().UTC()
	fmt.Println("Hora UTC atual: ", nowUTC)

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Erro ao carregar fusu horário: ", err)
		return
	}

	nowNY := nowUTC.In(loc)
	formattedTimeNY := nowNY.Format("2006-01-02 15:04")
	fmt.Println("Hora em Nova Yourk", formattedTimeNY)
}
