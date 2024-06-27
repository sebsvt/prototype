from fastapi import APIRouter, Body, Depends
from starlette.status import HTTP_201_CREATED
from app.models.cleaning import CleaningPulic, CleaningCreate
from app.db.repositories.cleanings import CleaningsRepository
from app.api.dependencies.database import get_repository

router = APIRouter()

@router.get("/")
async def get_all_cleanings() -> list[dict]:
	cleanings = [
		{"id": 1, "name": "My house", "cleaning_type": "full_clean", "price_per_hour": 2.99},
		{"id": 2, "name": "Someone else's house", "cleaning_type": "spot_clean", "price_per_hour": 19.99},
	]
	return cleanings

@router.post("/", response_model=CleaningPulic, name="cleanings:create-cleaning", status_code=HTTP_201_CREATED)
async def create_new_cleaning(
	new_cleaning: CleaningCreate = Body(..., embed=True),
	cleaning_repo: CleaningsRepository = Depends(get_repository(CleaningsRepository))
) -> CleaningPulic:
	created_cleaning = await cleaning_repo.create_cleaning(new_cleaning=new_cleaning)
	return created_cleaning
