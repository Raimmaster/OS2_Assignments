from Right import Right
class Domain:
	def __init__(self, name):
		self.name = name
		self.object_rights_tuple_list = []
		self.past_domain = None
		self.object_type = 'DOMAIN'
		self.switched = False
		self.inherited_rights_tuple_list = []

	def add_right_to_object(self, focus_object, right, switchable):
		tuple_in_list = self.search_object_in_list(focus_object)
		new_right = Right(right, switchable)
		print('Adding right %s' % new_right.name)
		print('Right switch: %r' % new_right.switchable)
		if tuple_in_list is not None:
			rights_array = tuple_in_list[1]
			right_exists = len(list(filter(lambda tup_right: tup_right.name == right, rights_array))) > 0
			if not right_exists:
				rights_array.append(new_right)
			return

		object_tuple = (focus_object, [new_right])
		self.object_rights_tuple_list.append(object_tuple)

	def search_object_in_list(self, focus_object):
		for obj_right_tuple in self.object_rights_tuple_list:
			if(focus_object is obj_right_tuple[0]):
				return obj_right_tuple
        
		return None

	def print_current_objects(self):
		for obj_right_tuple in self.object_rights_tuple_list:
			print('Object: %s' % obj_right_tuple[0].name)
			for right in obj_right_tuple[1]:
				print('With rights: %s' % right.name)

	def delete_object_from_tuples_array(self, object_to_delete):
		for obj_right_tuple in object_rights_tuple_list:
			objeto = obj_right_tuple[0]
			if objeto is object_to_delete:
				object_rights_tuple_list.remove(obj_right_tuple)
				print("Object removed from tuple list")

	def has_right_of_object(self, target_object, right_to_filter):
		if target_object is self:
			return True

		obj, rights = self.search_object_in_list(target_object)
		rights_filtered_list = list(filter(lambda right: right.name == right_to_filter, rights))
		return len(rights_filtered_list) > 0

	def get_objects_with_right(self, right):
		objects_tuple_list = list(filter(lambda obj_right_tuple: len(list(filter(lambda tuple_right: tuple_right.name == right, obj_right_tuple[1]))) > 0, self.object_rights_tuple_list))

		objects = []
		for obj_right_tuple in objects_tuple_list:
			objects.append(obj_right_tuple[0])

		return objects

	def remove_right_from_object(self, target_object, right_name):
		obj_right_tuple = self.search_object_in_list(target_object)
		right_to_remove = next(right for right in obj_right_tuple[1] if right.name == right_name)
		obj_right_tuple[1].remove(right_to_remove)
		self.print_current_objects()

	def get_rights_from_past_domain(self):
		switchable_rights = self.past_domain.get_switchable_rights()
		for switch_tuple in switchable_rights:
			objeto, rights = switchable_rights
			for right in rights:
				self.add_right_to_object(objeto.name, right.name, False)

		self.inherited_rights_tuple_list = switchable_rights

	def get_switchable_rights(self):
		switchable_rights = []
		for obj_right_tuple in self.object_rights_tuple_list:
			objeto, rights = obj_right_tuple
			rights_array = []
			for right in rights:
				if right.switchable:
					rights_array.append(right)

			if len(rights_array) > 0:
				obj_right_tuple = (objeto, rights_array)
				switchable_rights.append(obj_right_tuple)

		return switchable_right

	def remove_inherited_rights(self):
		for obj_right_tuple in inherited_rights_tuple_list:
			objeto, rights = obj_right_tuple
			for right in rights:
				self.remove_right_from_object(objeto, right.name)

		self.inherited_rights_tuple_list = []