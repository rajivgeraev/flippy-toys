from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton


def get_welcome_keyboard(webapp_url: str) -> InlineKeyboardMarkup:
    keyboard = [
        [
            InlineKeyboardButton(
                text="🎮 Открыть Flippy Toys",
                web_app={"url": webapp_url}
            )
        ]
    ]
    return InlineKeyboardMarkup(inline_keyboard=keyboard)
