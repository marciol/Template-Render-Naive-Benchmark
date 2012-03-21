require 'erb'

filename = "tmpl/person_list_erb.html"

f = File.new(filename, 'r')

persons = [] 

100.times do |i|
  persons << {:name => 'Name ' + i.to_s, :last_name => 'LastName ' + i.to_s}
end

t = ERB.new(f.read)

10.times do |i|
  puts 'Executing %s' % i
  puts t.result(binding)
end
