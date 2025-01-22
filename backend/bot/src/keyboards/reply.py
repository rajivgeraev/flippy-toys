from aiogram.types import (
    InlineKeyboardMarkup,
    InlineKeyboardButton,
    ReplyKeyboardMarkup,
    KeyboardButton,
    WebAppInfo,
)
from typing import List
from ..models import Child


def get_welcome_keyboard(webapp_url: str) -> InlineKeyboardMarkup:
    keyboard = [
        [
            InlineKeyboardButton(
                text="🎮 Открыть Flippy Toys", web_app={"url": webapp_url}
            )
        ]
    ]
    return InlineKeyboardMarkup(inline_keyboard=keyboard)


def get_children_keyboard(children: List[Child], kids_url: str) -> InlineKeyboardMarkup:
    keyboard = [
        [
            InlineKeyboardButton(
                text=f"👶 {child.name} ({child.age} лет)",
                web_app={"url": f"{kids_url}/{child.id}"},
            )
        ]
        for child in children
    ]
    return InlineKeyboardMarkup(inline_keyboard=keyboard)
