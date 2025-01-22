from dataclasses import dataclass
from datetime import datetime


@dataclass
class Child:
    id: str
    parent_id: str
    name: str
    age: int
    gender: str
    created_at: datetime
    updated_at: datetime
