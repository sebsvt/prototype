from fastapi import APIRouter

router = APIRouter()

@router.get("/")
async def get_all_cleanings() -> list[dict]:
	cleanings = [
		{"id": 1, "name": "My house", "cleaning_type": "full_clean", "price_per_hour": 2.99},
		{"id": 2, "name": "Someone else's house", "cleaning_type": "spot_clean", "price_per_hour": 19.99},
	]
	return cleanings
