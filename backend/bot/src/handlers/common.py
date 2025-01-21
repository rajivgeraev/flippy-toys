from aiogram import Router, types, F
from aiogram.filters import Command
from aiogram.fsm.context import FSMContext

from ..keyboards.reply import get_welcome_keyboard, get_location_keyboard
from ..texts.messages import WELCOME_MESSAGE, LOCATION_REQUEST_MESSAGE

router = Router()


@router.message(Command("start"))
async def command_start(message: types.Message, webapp_url: str) -> None:
    """Обработчик команды /start"""
    # Отправляем приветственное сообщение с кнопкой веб-приложения
    await message.answer(
        text=WELCOME_MESSAGE,
        reply_markup=get_welcome_keyboard(webapp_url),
        parse_mode="HTML",
    )
    # Отправляем запрос на геолокацию
    await message.answer(
        text=LOCATION_REQUEST_MESSAGE,
        reply_markup=get_location_keyboard(),
    )


@router.message(F.location)
async def handle_location(message: types.Message) -> None:
    """Обработчик получения геопозиции"""
    lat = message.location.latitude
    lon = message.location.longitude

    # Здесь можно добавить логику сохранения локации в базу данных
    # или отправки её в основной бэкенд

    await message.answer(
        text=f"Спасибо! Ваши координаты получены:\nШирота: {lat}\nДолгота: {lon}",
        reply_markup=types.ReplyKeyboardRemove(),  # Убираем клавиатуру после получения локации
    )
