from Domain import Domain
from Objeto import Objeto
class Matrix:
	def __init__(self):
		self.domains = []
		self.objects = []
		self.rights = []
		self.current_domain = None
		self._init_rights()
		self.admin_domain = self.create_domain('Admin')
		self.admin_domain.past_domain = self.admin_domain
		self.current_domain = self.admin_domain
		self.domains_object = self.create_object('Domains', self.admin_domain)

		admin_rights = ['switch', 'create', 'owner', 'delete']
		switchable = False
		for right in admin_rights:
			self.admin_domain.add_right_to_object(self.domains_object, right, switchable)

	def _init_rights(self):
		self.rights = ['switch', 'create', 'owner', 'delete', 'read', 'write', 'execute', 'switch']

	def create_domain(self, domain_name):
		domain = Domain(domain_name)
		if self.current_domain is None:
			self.current_domain = domain
		self.domains.append(domain)
		self.objects.append(domain)
		self.current_domain.add_right_to_object(domain, 'control', False)
		return domain

	def create_object(self, object_name, owner_domain):
		objeto = Objeto(object_name)
		self.objects.append(objeto)
		owner_domain.add_right_to_object(objeto, 'owner', False)

		return objeto

	def set_current_domain(self, domain_name):
		domain = self.search_domain(domain_name)
		if domain is not None:
			domain.past_domain = self.current_domain
			self.current_domain = domain

	def get_current_domain(self):
		return self.current_domain

	def search_domain(self, domain_name):
		for domain in self.domains:
			if domain.name == domain_name:
				return domain

		return None

	def search_object(self, object_name):
		for obj in self.objects:
			if obj.name == object_name:
				return obj

		return None

	def delete_domain(self, name):
		domain = self.search_domain(name)
		self.domains.remove(domain)

	def exit(self):
		if self.current_domain.switched:
			self.current_domain.remove_inherited_rights()
			self.current_domain.switched = False
		self.current_domain = self.current_domain.past_domain

	def delete_object(self, name):
		objeto = self.search_object(name)
		if(objeto.object_type == 'DOMAIN'):
			self.delete_domain(name)
		self.delete_object_from_domains_tuples(objeto)
		self.objects.remove(objeto)

	def delete_object_from_domains_tuples(self, object_to_delete):
		for domain in self.domains:
			domain.delete_object_from_tuples_array(object_to_delete)

	def print_objects(self, object_list):
		for objeto in object_list:
			print(objeto.name)

	def set_access_right(self, target_domain, target_object_name, right, switchable):
		domain = self.search_domain(target_domain)
		objeto = self.search_object(target_object_name)
		is_owner = self.current_domain.has_right_of_object(objeto, 'owner')
		has_control = self.current_domain.has_right_of_object(domain, 'control')

		if is_owner or has_control:
			domain.add_right_to_object(objeto, right, switchable == 'true')

	def verify_access_right(self, right):
		objects = self.current_domain.get_objects_with_right(right)
		self.print_objects(objects)
		return objects

	def rm_access_right(self, target_domain, target_object, right):
		domain = self.search_domain(target_domain)
		objeto = self.search_object(target_object)
		is_owner = self.current_domain.has_right_of_object(objeto, 'owner')
		has_control = self.current_domain.has_right_of_object(domain, 'control')
		if(is_owner or has_control):
			domain.remove_right_from_object(objeto, right)

	def switch(self, target_domain):
		object_right_list = self.verify_access_right('switch')
		if len(object_right_list) > 0:
			domain = self.search_domain(target_domain)
			if domain in object_right_list:
				self.set_current_domain(target_domain)
				self.current_domain.get_rights_from_past_domain()
				self.current_domain.switched = True
		else:
			print('You have no right switch for that domain')
    			