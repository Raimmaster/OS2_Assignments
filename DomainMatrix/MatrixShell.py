import os
import sys
import cmd
import signal
from Matrix import Matrix

class MatrixShell(cmd.Cmd):
    intro = 'Welcome to the access matrix shell.\tType help or ? to list commands.\n'
    matrix = Matrix()
    prompt = '(%s)>' % matrix.get_current_domain().name
    # ----- basic turtle commands -----
    def do_su(self, name):
        'Set the current domain: su D1'
        self.matrix.set_current_domain(name)
        self.prompt = '(%s)>' % self.matrix.get_current_domain().name
    
    def do_create_domain(self, arg):
        'Create a domain: create_domain D1'
        self.matrix.create_domain(arg)
    
    def do_create_object(self, arg):
        'Create an object: create_object F1'
        objeto = self.matrix.create_object(arg, self.matrix.get_current_domain())
        #print('%s' % objeto.name)
        #self.matrix.get_current_domain().print_current_objects()
    
    def do_delete_domain(self, arg):
        'Delete a domain: delete_domain D1'
        self.matrix.delete_domain(arg)
    
    def do_delete_object(self, arg):
        'Delete an object: delete_object F1'
        self.matrix.delete_object(arg)
    
    def do_set_access_right(self, arg):
        'Set a domain/object access right: set_access_right target_domain target_object right_name switchable'
        target_domain, target_object, right, switchable  = arg.split()
        self.matrix.set_access_right(target_domain, target_object, right, switchable)
    
    def do_rm_access_right(self, arg):
        'Remove a domain/object access right: rm_access_right target_domain target_object right_name'
        target_domain, target_object, right = arg.split()
        self.matrix.rm_access_right(target_domain, target_object, right)
    
    def do_verify_access_right(self, arg):
        "Verify a domain's access right, and get the objects where it can use it: verify_access_right read"
        self.matrix.verify_access_right(arg)

    def do_switch(self, arg):
        'Switch to another domain, if allowed: switch target_domain'
        self.matrix.switch(arg)
        self.prompt = '(%s)>' % self.matrix.get_current_domain().name
    
    def do_exit(self, arg):
        'Return to previous domain: exit'
        self.matrix.exit()
        self.prompt = '(%s)>' % self.matrix.get_current_domain().name
    
    def do_print_domains_objects(self, arg):
        'Prints every domain and its respective objects with rights'
        self.matrix.print_domains_and_objects()

    def do_print_objects(self, arg):
        'Prints all objects in matrix'
        self.matrix.print_objects_()

    def do_quit(self, arg):
        'Exit from the program  '
        quit()

if __name__ == '__main__':
    MatrixShell().cmdloop()
