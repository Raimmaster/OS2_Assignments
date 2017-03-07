class Matrix:
    def __init__(self):
        self.domains = []
        self.objects = []
        self.rights = []
        self._init_rights()

        self.admin_domain = self.create_domain('Admin')
        self.domains_object = self.create_object('Domains', self.admin_domain)

        admin_rights = ['switch', 'create', 'owner', 'delete']
        
        for right in admin_rights:
        	print(right)
        	self.admin_domain.add_right_to_object(domains_object, right)


    def create_domain(domain_name):
    	self.domains.append(domain_name)

    def create_object(object_name, owner_domain):
    	

