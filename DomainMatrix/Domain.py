from Right import Right
class Domain:
    def __init__(self, name):
        self.name = name
        self.object_rights_tuple_list = []
        self.past_domain = None


    def add_right_to_object(self, focus_object, right, switchable):
        tuple_in_list = self.search_object_in_list(focus_object)
        new_right = Right(right, switchable)
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