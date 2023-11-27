TestTask

Установка
1. git clone https://github.com/zanzhit/TestTask
2. cd TestTask
3. go mod download
4. Установите ffmpeg.

Использование
1. В функции recordStream, предназначенной для записи потока, в переменную  cmdPath нужно указать путь до ffmpeg.
2. В файле urls.txt указать необходимые URL каждый с новой строки.
3. go run learn.go в терминал.
4. Для прекращения записи нужно отправить что-нибудь в терминал (пустая строка, символ, предложение - неважно).