from fastapi import HTTPException
from starlette.status import HTTP_400_BAD_REQUEST
from app.db.repositories.base import BaseRepository
from app.models.cleaning import CleaningCreate, CleaningInDB, CleaningUpdate

CREATE_CLEANING_QUERY = """
	INSERT INTO cleanings (name, description, price, cleaning_type)
	VALUES (:name, :description, :price, :cleaning_type)
	RETURNING id, name, description, price, cleaning_type;
"""

GET_CLEANING_QUERY = """
	SELECT id, name, description, price, cleaning_type
	FROM cleanings
	WHERE id = :id;
"""

GET_ALL_CLEANING_QUERY = """
	SELECT id, name, description, price, cleaning_type
	FROM cleanings
"""

UPDATE_CLEANING_BY_ID_QUERY = """
	UPDATE cleanings
	SET name 		  = :name,
		description	  = :description,
		price 		  = :price,
		cleaning_type = :cleaning_type
	WHERE id = :id
	RETURNING id, name, description, price, cleaning_type;
"""

DELETE_CLEANING_BY_ID_QUERY = """
	DELETE FROM cleanings
	WHERE id = :id
	RETURNING id;
"""

class CleaningsRepository(BaseRepository):
	async def create_cleaning(self, *, new_cleaning: CleaningCreate) -> CleaningInDB:
		query_values = new_cleaning.model_dump()
		cleaning = await self.db.fetch_one(query=CREATE_CLEANING_QUERY, values=query_values)

		return CleaningInDB(**cleaning)

	async def get_cleaning_by_id(self, *, id: int) -> CleaningInDB | None:
		cleaning = await self.db.fetch_one(query=GET_CLEANING_QUERY, values={"id": id})
		if not cleaning:
			return None
		return CleaningInDB(**cleaning)

	async def get_all_cleanings(self) -> CleaningInDB:
		cleanings = await self.db.fetch_all(query=GET_ALL_CLEANING_QUERY)
		return [CleaningInDB(**cln) for cln in cleanings]

	async def update_cleaning(
			self, *, id: int, cleaning_update: CleaningUpdate,
	) -> CleaningInDB | None:
		cleaning = await self.get_cleaning_by_id(id=id)
		if not cleaning:
			return None
		cleaning_update_params = cleaning.model_copy(update=cleaning_update.model_dump(exclude_unset=True))
		if cleaning_update_params.cleaning_type is None:
			raise HTTPException(
				status_code=HTTP_400_BAD_REQUEST,
				detail="Invalid cleaning Type. Cannot be None."
			)
		try:
			update_cleaning = await self.db.fetch_one(
				query=UPDATE_CLEANING_BY_ID_QUERY,
				values=cleaning_update_params.model_dump()
			)
			return CleaningInDB(**update_cleaning)
		except Exception as e:
			print(e)
			raise HTTPException(
				status_code=HTTP_400_BAD_REQUEST,
				detail="Invalid update params."
			)

	async def delete_cleaning_by_id(self, *, id: int) -> int | None:
		cleaning = await self.get_cleaning_by_id(id=id)
		if not cleaning:
			return None
		deleted_id = await self.db.execute(
			query=DELETE_CLEANING_BY_ID_QUERY,
			values={'id': id}
		)
		return deleted_id

