from aiogram import Router, types
from aiogram.filters import Command

from ..keyboards.reply import get_welcome_keyboard
from ..texts.messages import WELCOME_MESSAGE

router = Router()


@router.message(Command("start"))
async def command_start(message: types.Message, webapp_url: str) -> None:
    await message.answer(
        text=WELCOME_MESSAGE,
        reply_markup=get_welcome_keyboard(webapp_url),
        parse_mode="HTML"
    )
