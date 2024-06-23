from datetime import date
from model import Batch, OrderLine

def make_batch_and_orderline(sku: str, qty: int, line_qty: int):
	return (
		Batch(ref="batch-ref", sku=sku, qty=qty, eta=date.today()),
		OrderLine(reference="order-ref", sku=sku, qty=line_qty)
	)

def test_allocate_if_available_greater_than_required():
	batch, line = make_batch_and_orderline("RED-TABLE", 20, 2)
	batch.allocate(line)
	assert batch.available_quantity == 18

def test_can_not_allocate_if_available_smaller_than_required():
	batch, line = make_batch_and_orderline("RED-TABLE", 20, 21)
	batch.allocate(line)
	assert batch.available_quantity == 20

def test_can_allocate_if_available_equal_to_required():
	batch, line = make_batch_and_orderline("RED-TABLE", 20, 20)
	batch.allocate(line)
	assert batch.available_quantity == 0

def test_cannot_allocate_if_skus_do_not_match():
	batch, un_used = make_batch_and_orderline("RED-TABLE", 20, 1)
	line = OrderLine(reference="order-ref", sku="PYTHON", qty=1)
	batch.allocate(line)
	assert batch.available_quantity == 20

def test_can_only_deallocate_allocated_lines():
	batch, unallocated_line = make_batch_and_orderline("DECORATIVE-TRINKET", 20, 2)
	batch.allocate(unallocated_line)
	batch.deallocate(unallocated_line)
	assert batch.available_quantity == 20

def test_allocation_is_idempotent():
	batch, line = make_batch_and_orderline("ANGULAR-DESK", 20, 2)
	batch.allocate(line)
	batch.allocate(line)
	assert batch.available_quantity == 18
