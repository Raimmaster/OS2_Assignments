class Domain:
    def __init__(self, name):
        self.name = name;
        self.object_rights_tuple_list = []

    def add_right_to_object(focus_object, right):
    	tuple_in_list = self.search_object_in_list(focus_object)
    	if(tuple_in_list is not None):
    		rights_array = tuple_in_list[1]
    		rights_array.append(right)
    		return

    	self.object_rights_tuple_list.append(focus_object, [right])

    def search_object_in_list(focus_object):
    	for obj_right_tuple in self.object_rights_tuple_list:
    		if focus_object is obj_right_tuple[0]:
    			return obj_right_tuple


    	return None

