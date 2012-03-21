local cosmo = require 'cosmo'

local filename = "tmpl/person_list_lua.html"

local template_file = assert(io.open(filename, 'r'))

local template = template_file:read("*all")

local persons = {} 

for i = 1,100 do
    persons[i] = { name = "Name " .. i, last_name = "LastName " .. i} 
end

local compiled = cosmo.f(template) 

local values = { persons = persons }

for i = 1,10 do
    print("Executing ", i)
    print(compiled(values))
end
