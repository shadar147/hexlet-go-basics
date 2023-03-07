package ex30

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Возвращаемые ошибки принято проверять при каждом вызове:

func handleResponse() {
	response, err := DoHTTPCall()
	if err != nil {
		log.Println(err)
	}
	// только после проверки на ошибку можно делать что-то с объектом response

	fmt.Println(response)
}

func DoHTTPCall() (int, error) {
	return 1, nil
}

// При этом логика обработки отличается от места и типа ошибки. Ошибки можно оборачивать и прокидывать в функцию выше, логировать или делать любые фоллбек действия.
// Оборачивание ошибок — важная часть написания кода на Go. Это позволяет явно видеть трейс вызова и место возникновения ошибки. Для оборачивания используется функция fmt.Errorf:
// для простоты примера опускаем аргументы запроса и ответа
func DoHTTPCallError() error {
	err := SendTCP()
	if err != nil {
		// оборачивается в виде "[название метода]: %w". %w — это плейсхолдер для ошибки
		return fmt.Errorf("send tcp: %w", err)
	}

	return nil
}

var errTCPConnectionIssue = errors.New("TCP connect issue")

func SendTCP() error {
	return errTCPConnectionIssue
}

func handleDoHTTPCallError() {
	fmt.Println(DoHTTPCallError()) // send tcp: TCP connect issue
}

// В современном Go существуют функции для проверки типов конкретных ошибок. Например, ошибку из примера выше можно проверить с помощью функции errors.Is. В данном случае errTCPConnectionIssue обернута другой ошибкой, но функция errors.Is найдет ее при проверке:
func handleIsError() error {
	err := DoHTTPCallError()
	if err != nil {
		if errors.Is(err, errTCPConnectionIssue) {
			// в случае ошибки соединения ждем 1 секунду и пытаемся сделать запрос снова
			time.Sleep(1 * time.Second)
			return DoHTTPCallError()
		}

		// обработка неизвестной ошибки
		log.Println("unknown error on HTTP call", err)
	}

	return nil
}

// errors.Is подходит для проверки статичных ошибок, хранящихся в переменных. Иногда нужно проверить не конкретную ошибку, а целый тип. Для этого используется функция errors.As:
type ConnectionErr struct{}

func (e ConnectionErr) Error() string {
	return "connection err"
}

func handleAsError() {
	tries := 0
	for {
		if tries > 2 {
			log.Println("Can't connect to DB")
			break
		}

		err := connectDB()
		if err != nil {
			// если ошибка подключения, то ждем 1 секунду и пытаемся снова
			if errors.As(err, &ConnectionErr{}) {
				log.Println("Connection error. Trying to reconnect...")
				time.Sleep(1 * time.Second)
				tries++
				continue
			}

			// в противном случае ошибка критичная, логируем и выходим из цикла
			log.Println("connect DB critical error", err)
		}

		break
	}
}

// для простоты функция всегда возвращает ошибку подключения
func connectDB() error {
	return ConnectionErr{}
}

// Вывод программы спустя 3 секунды:
// Connection error. Trying to reconnect...
// Connection error. Trying to reconnect...
// Connection error. Trying to reconnect...
// Can't connect to DB
