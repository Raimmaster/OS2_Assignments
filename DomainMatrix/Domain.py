from Right import Right
class Domain:
    def __init__(self, name):
        self.name = name
        self.object_rights_tuple_list = []
        self.past_domain = None
        self.object_type = 'DOMAIN'

    def add_right_to_object(self, focus_object, right, switchable):
        tuple_in_list = self.search_object_in_list(focus_object)
        new_right = Right(right, switchable)
        print('Adding right %s' % new_right.name)
        print('Right switch: %r' % new_right.switchable)
        if tuple_in_list is not None:
            rights_array = tuple_in_list[1]
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
            print('%s' % obj_right_tuple[0].name)
            for right in obj_right_tuple[1]:
                print('%s' % right.name)

    def delete_object_from_tuples_array(self, object_to_delete):
    	for obj_right_tuple in object_rights_tuple_list:
    		objeto = obj_right_tuple[0]
    		if objeto is object_to_delete:
    			object_rights_tuple_list.remove(obj_right_tuple)
    			print("Object removed from tuple list")

    def has_right_of_object(self, target_object, right_to_filter):
        obj, rights = self.search_object_in_list(target_object)
        rights_filtered_list = list(filter(lambda right: right.name == right_to_filter, rights))
        return len(rights_filtered_list) > 0

    def get_objects_with_right(self, right):
    	objects_tuple_list = list(
    		filter(
    			lambda obj_right_tuple: len(
    				list(
    					filter(
    						lambda tuple_right: tuple_right.name == right, obj_right_tuple[1]))) > 0,
    			self.object_rights_tuple_list))

    	objects = []
    	for obj_right_tuple in objects_tuple_list:
    		objects.append(obj_right_tuple[0])

    	return objects