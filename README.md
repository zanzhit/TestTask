# ЭТОТ КОД БЫЛ НУЖЕН ДЛЯ РЕШЕНИЯ ТЕСТОВОГО ЗАДАНИЯ, ДЛЯ ЗАПИСИ ВИДЕО ЛУЧШЕ ИСПОЛЬЗОВАТЬ https://github.com/zanzhit/gstreamer-go В НЕМ БОЛЕЕ ШИРОКИЙ ФУНКЦОНАЛ

Установка
1. git clone https://github.com/zanzhit/TestTask
2. cd TestTask
3. go mod download
4. Установите ffmpeg.
5. Установите OpenVPN с нужнык конфигом.

Использование
1. В функции recordStream, предназначенной для записи потока, в переменную  cmdPath нужно указать путь до ffmpeg.
2. В файле urls.txt указать необходимые URL каждый с новой строки.
3. Запустите OpenVPN.
4. go run learn.go в терминал.
5. Для прекращения записи нужно отправить что-нибудь в терминал (пустая строка, символ, предложение - неважно).
