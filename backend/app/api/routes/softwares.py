from fastapi import APIRouter
from datetime import datetime, timedelta

router = APIRouter()

@router.get("/")
async def get_all_softwares() -> list[dict]:
	return [
		{
			'id': "11",
			'name': "service 1",
			"tenant_id": 1,
			"created_at": datetime.today(),
			"expried_at": datetime.today() + timedelta(days=30)
		},
		{
			'id': "2",
			'name': "service 2",
			"tenant_id": 1,
			"created_at": datetime.today(),
			"expried_at": datetime.today() + timedelta(days=30)
		},
	]
