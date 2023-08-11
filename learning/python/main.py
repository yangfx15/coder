from fastapi import FastAPI
from pydantic import BaseModel
# 新增一个书籍Book的数据结构，继承自BaseModel类
class Book(BaseModel):
    name: str
    desp: str
    price: float

app = FastAPI()
# 定义一个post接口，接收一本书的信息
@app.post("/books/")
async def create_book(book: Book):
    return book