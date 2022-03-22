package models

const (
	CommandStart   = "start"
	CommandValutes = "valutes"
	CommandSummary = "summary"

	DifferenceGrown     = "вырос ⬆️"
	DifferenceDecreased = "опустился ⬇️"

	ReplyWelcome          = "Привет, %s!\nВас приветствует бот для отслеживания курсов валют.\nВы можете отслеживать, как отдельную валюту сами, или настроить ежедневное оповещение.\n\nСейчас мы отслеживаем курсы %s, %s и %s."
	ReplyValute           = "Текущий курс %s: <strong>%.2f руб.</strong>"
	ReplySelectValute     = "Выберите валюту"
	ReplyUndefined        = "К сожалению, я не знаю, что тебе ответить."
	ReplySummaryRequest   = "Вы можете включить ежедневную отправку сводки курсов.\nПожалуйста, выберите желаемое время получения."
	ReplySelectTime       = "Принято! Вы будете получать ежедневные сводки в %s."
	ReplySummaryNotActive = "Принято! Вам больше не будут приходить автоматические сводки."
	ReplyEveryDaySummary  = "Ежедневная сводка за %s.\n\n%s:  %.2f руб.\n%s:  %.2f руб.\n%s:  %.2f руб."
	ReplyChangeCourseData = "Изменение курса❗️\nКурс %s %s на %.2f руб.\nСейчас: %.2f руб.\nРанее: %.2f руб."
)
