from Domain import Domain
from Objeto import Objeto
class Matrix:
    def __init__(self):
        self.domains = []
        self.objects = []
        self.rights = []
        self._init_rights()

        self.admin_domain = self.create_domain('Admin')
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
        self.domains.append(domain)

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
