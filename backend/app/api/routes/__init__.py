from fastapi import APIRouter
from app.api.routes.softwares import router as softwares_router

router = APIRouter()

router.include_router(router=softwares_router, prefix="/softwares", tags=["softwares"])
