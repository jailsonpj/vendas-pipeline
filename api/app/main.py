from fastapi import FastAPI 
from elastic import get_all, get_data, get_escrv, get_grp_merc, get_material
app = FastAPI()

@app.get("/vendas")
async def read_vendas():
    response = get_all()
    return response

@app.get("/vendas/{material}")
async def read_material(material):
    response = get_material(material)
    return response
