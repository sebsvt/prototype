from dataclasses import dataclass
from datetime import date

@dataclass(frozen=True)
class OrderLine:
	reference: str
	sku: str
	qty: int

class Batch:
	def __init__(self, ref: str, sku: str, qty: int, eta: date) -> None:
		self.reference = ref
		self.sku = sku
		self.eta = eta
		self._purchased_quantity = qty
		self._allocations: set[OrderLine] = set()

	@property
	def allocated_quantity(self):
		return sum(line.qty for line in self._allocations)

	@property
	def available_quantity(self):
		return self._purchased_quantity - self.allocated_quantity

	# I made this function because allocate function should do it own work
	def can_allocate(self, line: OrderLine):
		return self.available_quantity >= line.qty and self.sku == line.sku

	def allocate(self, line: OrderLine):
		if self.can_allocate(line):
			self._allocations.add(line)

	def deallocate(self, line: OrderLine):
		if line in self._allocations:
			self._allocations.remove(line)
