# go-telegram-bot

Бот для мониторинга курсов валют по состоянию на текущий день с сайта ЦБ РФ.</br>
Возможности бота:
* вывод курса конкретной валюты;
* запрос у пользователя разрешения на отправку ежедневных сводок;
* отправка сводок по курсам валют для конкретных пользователей;
* автоматический запрос курсов валют с удаленного источника (по cron);
* при изменении курса, активные пользователи получат оповещения об изменении.

Команды бота:
* **/start** - приветствие
* **/summary** - настройка отправки оповещений
* **/valutes** - список доступных валют (запрос курса)
* **/help** - помощь и описание