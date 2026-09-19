// Битовые флаги: свой тип, iota со сдвигом, три оператора.
//
// Запуск:
//
//	go run ./seminar-02-language/examples/09-bit-flags
//
// Набор независимых признаков можно хранить тремя полями bool, а можно —
// одним числом, где каждый бит отвечает за свой признак. Второй способ
// компактнее и позволяет передавать «несколько прав сразу» одним значением.
package main

import "fmt"

// Permission — собственный тип поверх uint8. Именно из-за него компилятор
// не даст случайно сложить права с обычным числом или с другим флаговым типом.
type Permission uint8

// Каждая константа — отдельный бит. 1 << iota даёт 001, 010, 100.
// Выражение справа повторяется само, поэтому писать его нужно один раз.
const (
	Read    Permission = 1 << iota // 001
	Write                          // 010
	Execute                        // 100
)

func main() {
	fmt.Printf("Read=%03b Write=%03b Execute=%03b\n\n", Read, Write, Execute)

	operators()
	combining()
	theTrap()
}

func operators() {
	fmt.Println("--- три оператора, которые нужны ---")

	current := Read

	// | — включить биты. Уже включённые не ломаются.
	granted := current | Write
	fmt.Printf("  %03b | %03b  = %03b   включить\n", current, Write, granted)

	// &^ — «and not», сбросить биты. Сбрасывать отсутствующий бит безопасно.
	revoked := granted &^ Read
	fmt.Printf("  %03b &^ %03b = %03b   сбросить\n", granted, Read, revoked)

	// & — оставить только общие биты. На нём строится проверка наличия.
	common := granted & (Read | Execute)
	fmt.Printf("  %03b &  %03b = %03b   пересечение\n", granted, Read|Execute, common)
}

func combining() {
	fmt.Println("--- проверка наличия ---")

	current := Read | Execute

	// ЛОВУШКА: current & required != 0 отвечает на вопрос «есть хотя бы одно
	// из прав», а не «есть все». Для «есть все» сравнивают с самим required.
	for _, required := range []Permission{Read, Write, Read | Execute, Read | Write} {
		anyOf := current&required != 0
		allOf := current&required == required
		fmt.Printf("  нужно %03b: хотя бы одно = %-5t  все = %t\n", required, anyOf, allOf)
	}
}

func theTrap() {
	fmt.Println("--- почему пустой набор прав всегда есть ---")

	current := Read | Execute
	var nothing Permission // 000

	// current & 000 == 000 — условие выполнено. Это не баг, а следствие
	// определения «есть все требуемые права»: требуемых прав ноль штук,
	// значит все они присутствуют. Так же ведёт себя пустое множество
	// в математике и, например, strings.Contains(s, "").
	fmt.Printf("  Has(%03b, %03b) = %t\n", current, nothing, current&nothing == nothing)
}
