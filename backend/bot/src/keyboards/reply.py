from aiogram.types import (
    InlineKeyboardMarkup,
    InlineKeyboardButton,
    ReplyKeyboardMarkup,
    KeyboardButton,
)


def get_welcome_keyboard(webapp_url: str) -> InlineKeyboardMarkup:
    """Клавиатура с кнопкой для открытия веб-приложения"""
    keyboard = [
        [
            InlineKeyboardButton(
                text="🎮 Открыть Flippy Toys", web_app={"url": webapp_url}
            )
        ]
    ]
    return InlineKeyboardMarkup(inline_keyboard=keyboard)


def get_location_keyboard() -> ReplyKeyboardMarkup:
    """Клавиатура с кнопкой запроса геопозиции"""
    keyboard = [
        [KeyboardButton(text="📍 Отправить местоположение", request_location=True)]
    ]
    return ReplyKeyboardMarkup(
        keyboard=keyboard, resize_keyboard=True, one_time_keyboard=True
    )
