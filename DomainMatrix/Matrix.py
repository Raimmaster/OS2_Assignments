class Matrix:
    def __init__(self):
        self.domains = []
        self.objects = []

        self.admin_domain = self.create_domain('Admin')
        self.domains_object = self.create_object('Domains', self.admin_domain)

        
