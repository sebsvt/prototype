from fastapi import APIRouter, Body, Depends, HTTPException, Path
from starlette.status import HTTP_201_CREATED, HTTP_200_OK, HTTP_404_NOT_FOUND
from app.models.cleaning import CleaningPulic, CleaningCreate, CleaningUpdate
from app.db.repositories.cleanings import CleaningsRepository
from app.api.dependencies.database import get_repository

router = APIRouter()

@router.get("/", response_model=list[CleaningPulic], name="cleanings:get-all-cleanings", status_code=HTTP_200_OK)
async def get_all_cleanings(cleaning_repo: CleaningsRepository = Depends(get_repository(CleaningsRepository))) -> list[CleaningPulic]:
	cleanings = await cleaning_repo.get_all_cleanings()
	return cleanings

@router.get("/{id}", response_model=CleaningPulic, name="cleanings:get-cleaning-by-id", status_code=HTTP_200_OK)
async def get_cleaning_by_id(id: int, cleaning_repo: CleaningsRepository = Depends(get_repository(CleaningsRepository))) -> CleaningPulic:
	cleaning = await cleaning_repo.get_cleaning_by_id(id=id)
	if not cleaning:
		raise HTTPException(status_code=HTTP_404_NOT_FOUND, detail=f"cleaning id {id} does not exists")
	return cleaning

@router.post("/", response_model=CleaningPulic, name="cleanings:create-cleaning", status_code=HTTP_201_CREATED)
async def create_new_cleaning(
	new_cleaning: CleaningCreate = Body(..., embed=True),
	cleaning_repo: CleaningsRepository = Depends(get_repository(CleaningsRepository))
) -> CleaningPulic:
	created_cleaning = await cleaning_repo.create_cleaning(new_cleaning=new_cleaning)
	return created_cleaning

@router.put("/{id}", response_model=CleaningPulic, name="cleanings:update-cleaning-by-id")
async def update_cleaning_by_id(
	id: int = Path(..., ge=1, title="The ID of the cleanings to update."),
	cleaning_update: CleaningUpdate = Body(..., embed=True),
	cleaning_repo: CleaningsRepository = Depends(get_repository(CleaningsRepository))
) -> CleaningPulic:
	updated_cleaning = await cleaning_repo.update_cleaning(id=id, cleaning_update=cleaning_update)
	if not updated_cleaning:
		raise HTTPException(
			status_code=HTTP_404_NOT_FOUND,
			detail="No cleaning found with that id."
		)
	return updated_cleaning

@router.delete("/{id}", response_model=int, name="cleanings:delete-cleanings-by-id")
async def delete_cleaning_by_id(id: int = Path(..., ge=1, title="The ID of cleaning to delete."), cleaning_repo: CleaningsRepository = Depends(get_repository(CleaningsRepository))) -> int:
	deleted_id = await cleaning_repo.delete_cleaning_by_id(id=id)
	if not deleted_id:
		raise HTTPException(
			status_code=HTTP_404_NOT_FOUND,
			detail="No cleaning found with that id."
		)
	return deleted_id
