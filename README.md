На непредусмотренные сообщения бот отвечает: `Нет такой команды`


На `/help` выводит список команд


на `/start` выводит сообщение о запуске бота и сразу предлагает авторизоваться, если это еще не сделанно

На `/login` предлагает зарегистрироваться через _Яндекс ID_, _Github_ или _кодом_

__Выявленна критическая ошибка:__ при нажатии на inline кнопки, которые повляются вместе с командой `/login`, бот перестает работать падая в панику!